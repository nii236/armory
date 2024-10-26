package armory

import (
	"armory/contracts"
	"context"
	"crypto/ecdsa"
	"database/sql"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

type Service struct {
	ethRPC *ethclient.Client
	db     *DB
	*contracts.WithdrawalQueueSession
	ownerAddr common.Address
	ownerKey  *ecdsa.PrivateKey
}

func New(
	ethRPC *ethclient.Client,
	db *DB,
) *Service {
	return &Service{
		db: db,
	}
}

// RequestWithdrawal requests an exit of the user's balance to the user's wallet
// The yield from the LST does not get withdrawn, it stays in protocol
// This is done through the server instead of triggered on-chain because the LST is stored in the single admin address
func (s *Service) RequestWithdrawal(ctx context.Context, userAddress common.Address, amt decimal.Decimal) error {
	stats, err := s.db.UserStatsByAddress(ctx, userAddress)
	if err != nil {
		return fmt.Errorf("user stats: %w", err)
	}
	if stats.TotalBalance.LessThan(amt) {
		return fmt.Errorf("insufficient balance: %s < %s", stats.TotalBalance.String(), amt.String())
	}
	nonce, err := s.ethRPC.NonceAt(ctx, s.ownerAddr, nil)
	if err != nil {
		return fmt.Errorf("nonce: %w", err)
	}

	block, err := s.ethRPC.BlockByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("block: %w", err)
	}

	blockTimestamp := block.Time()

	deadlineUnix := blockTimestamp + 300
	permitInput, err := CreatePermitInput(s.ownerAddr.Hex(), WithdrawalQueueAddress.Hex(), amt.BigInt(), nonce, big.NewInt(int64(deadlineUnix)), s.ownerKey)
	if err != nil {
		return fmt.Errorf("create permit input: %w", err)
	}

	contractPermitInput := contracts.WithdrawalQueuePermitInput{
		Value:    permitInput.Value,
		Deadline: permitInput.Deadline,
		V:        permitInput.V,
		R:        permitInput.R,
		S:        permitInput.S,
	}

	tx, err := s.WithdrawalQueueSession.RequestWithdrawalsWithPermit([]*big.Int{amt.BigInt()}, s.ownerAddr, contractPermitInput)
	if err != nil {
		return fmt.Errorf("request withdrawals: %w", err)
	}

	receipt, err := bind.WaitMined(ctx, s.ethRPC, tx)
	if err != nil {
		return fmt.Errorf("wait mined: %w", err)
	}

	// TODO: save in db
	return ErrNotImplemented
}

// Balance of the user's game wallet
func (s *Service) Balance(ctx context.Context, userAddress common.Address) (decimal.Decimal, error) {
	stats, err := s.db.UserStatsByAddress(ctx, userAddress)
	if err != nil {
		return decimal.Zero, fmt.Errorf("user stats: %w", err)
	}
	return stats.TotalBalance, nil
}

// BalanceSeconds returns the balance of the user's game wallet in eth seconds
// This is useful for measuring value of locked collateral
// Equivalent to airdrop points
func (s *Service) BalanceSeconds(ctx context.Context, userAddress common.Address) (decimal.Decimal, error) {
	stats, err := s.db.UserStatsByAddress(ctx, userAddress)
	if err != nil {
		return decimal.Zero, fmt.Errorf("user stats: %w", err)
	}
	return stats.TotalBalanceSeconds, nil
}

// Yield of the user's collateral
// After deposit, the collateral is converted to a rebasing LST which accumulates interest
// It is the difference between the user's on-chain LST balance and the user's total deposit
func (s *Service) Yield(ctx context.Context, userAddress common.Address) (decimal.Decimal, error) {
	stats, err := s.db.UserStatsByAddress(ctx, userAddress)
	if err != nil {
		return decimal.Zero, fmt.Errorf("user stats: %w", err)
	}
	return stats.TotalYield, nil
}

// Wallet returns the wallet for the user
// If it already exists, it returns the existing wallet
// If it doesn't exist, it creates a new wallet and returns it
func (s *Service) Wallet(ctx context.Context, userAddress common.Address) (common.Address, error) {
	user, err := s.db.UserByAddress(ctx, userAddress)
	if err != nil && err != sql.ErrNoRows {
		return common.Address{}, fmt.Errorf("user: %w", err)
	}

	key, err := crypto.GenerateKey()
	if err != nil {
		return common.Address{}, fmt.Errorf("generate key: %w", err)
	}

	walletAddress := crypto.PubkeyToAddress(key.PublicKey)

	if err == sql.ErrNoRows {
		err = s.db.InsertUser(ctx, user.ID, walletAddress, userAddress, key)
		if err != nil {
			return common.Address{}, fmt.Errorf("insert user: %w", err)
		}
	}

	return walletAddress, nil
}

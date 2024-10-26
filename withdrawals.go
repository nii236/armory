package armory

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// CompleteWithdraw completes an exit of the user's game balance to the user's wallet
func (s *Service) completeWithdrawal(ctx context.Context, userAddress common.Address, tokenID int) error {
	return ErrNotImplemented
}

// QueueWithdrawal queues an exit of the user's game balance to the user's wallet
func (s *Service) queueWithdrawal(ctx context.Context, userAddress common.Address, amt decimal.Decimal) error {
	return ErrNotImplemented
}

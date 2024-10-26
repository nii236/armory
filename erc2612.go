package armory

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

type PermitInput struct {
	Value    *big.Int
	Deadline *big.Int
	V        uint8
	R        [32]byte
	S        [32]byte
}

func CreatePermitInput(owner, spender string, value *big.Int, nonce uint64, deadline *big.Int, privateKey *ecdsa.PrivateKey) (*PermitInput, error) {
	domainSeparator := crypto.Keccak256Hash(
		[]byte("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)"),
	)
	addressType, _ := abi.NewType("address", "", nil)
	uint256Type, _ := abi.NewType("uint256", "", nil)
	data := abi.Arguments{
		{Type: addressType}, // Owner
		{Type: addressType}, // Spender
		{Type: uint256Type}, // Value
		{Type: uint256Type}, // Nonce
		{Type: uint256Type}, // Deadline
	}
	encodedData, err := data.Pack(owner, spender, value, nonce, deadline)
	if err != nil {
		return nil, fmt.Errorf("failed to encode data: %w", err)
	}

	permitHash := crypto.Keccak256Hash(append(domainSeparator.Bytes(), encodedData...))

	signature, err := secp256k1.Sign(permitHash.Bytes(), crypto.FromECDSA(privateKey))
	if err != nil {
		return nil, fmt.Errorf("failed to sign data: %w", err)
	}

	r := [32]byte{}
	s := [32]byte{}
	copy(r[:], signature[:32])
	copy(s[:], signature[32:64])
	v := uint8(signature[64] + 27)

	return &PermitInput{
		Value:    value,
		Deadline: deadline,
		V:        v,
		R:        r,
		S:        s,
	}, nil
}

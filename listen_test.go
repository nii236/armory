package armory_test

import (
	"armory"
	"armory/contracts"
	"context"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/test-go/testify/require"
)

func PrepareRPC(t *testing.T) (*ethclient.Client, *contracts.WithdrawalQueueFilterer, *contracts.LidoFilterer) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	RPC_URL := os.Getenv("RPC_URL")
	executionRPC, err := ethclient.Dial(RPC_URL)
	require.NoError(t, err)

	withdrawalQueueAddr := "0x889edC2eDab5f40e902b864aD4d7AdE8E412F9B1"
	withdrawalQueue, err := contracts.NewWithdrawalQueueFilterer(common.HexToAddress(withdrawalQueueAddr), executionRPC)
	require.NoError(t, err)

	lidoAddr := "0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84"
	lido, err := contracts.NewLidoFilterer(common.HexToAddress(lidoAddr), executionRPC)
	require.NoError(t, err)

	return executionRPC, withdrawalQueue, lido
}

func TestBackfill(t *testing.T) {
	executionRPC, withdrawalQueue, lido := PrepareRPC(t)

	l := armory.NewListener(withdrawalQueue, lido, common.Address{})
	require.NotNil(t, l)

	var err error

	optFns := []armory.WatchOptionFn{
		armory.WithFilterSenders(nil),
		armory.WithFilterOwners(nil),
		armory.WithFilterRequesters(nil),
		armory.WithFilterFrom(nil),
		armory.WithFilterTo(nil),
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)

	latestBlock, err := executionRPC.BlockNumber(ctx)
	require.NoError(t, err)

	fromBlock := latestBlock - 1000

	defer cancel()
	// go func() {
	// 	err = l.BackfillSubmitteds(ctx, fromBlock, optFns...)
	// 	require.NoError(t, err)
	// }()
	// go func() {
	// 	err = l.BackfillWithdrawalRequests(ctx, fromBlock, optFns...)
	// 	require.NoError(t, err)
	// }()
	go func() {
		err = l.BackfillWithdrawalFinalisations(ctx, fromBlock, optFns...)
		require.NoError(t, err)
	}()
	// go func() {
	// 	err = l.BackfillWithdrawalClaimeds(ctx, fromBlock, optFns...)
	// 	require.NoError(t, err)
	// }()
	<-ctx.Done()
}

func TestListen(t *testing.T) {
	_, withdrawalQueue, lido := PrepareRPC(t)

	l := armory.NewListener(withdrawalQueue, lido, common.Address{})
	require.NotNil(t, l)

	var err error

	optFns := []armory.WatchOptionFn{
		armory.WithFilterSenders(nil),
		armory.WithFilterOwners(nil),
		armory.WithFilterRequesters(nil),
		armory.WithFilterFrom(nil),
		armory.WithFilterTo(nil),
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	go func() {
		err = l.ListenForSubmitteds(ctx, optFns...)
		require.NoError(t, err)
	}()
	go func() {
		err = l.ListenForWithdrawalRequests(ctx, optFns...)
		require.NoError(t, err)
	}()
	go func() {
		err = l.ListenForWithdrawalFinalisations(ctx, optFns...)
		require.NoError(t, err)
	}()
	go func() {
		err = l.ListenForWithdrawalClaimeds(ctx, optFns...)
		require.NoError(t, err)
	}()
	<-ctx.Done()
}

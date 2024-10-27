package armory

import (
	"armory/onchain"
	"context"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type Listener struct {
	WithdrawalQueueFilterer *onchain.WithdrawalQueueFilterer
	LidoFilterer            *onchain.LidoFilterer
	OwnerAddr               common.Address
}

func NewListener(
	withdrawalQueueFilterer *onchain.WithdrawalQueueFilterer,
	lidoFilterer *onchain.LidoFilterer,
	ownerAddr common.Address,
) *Listener {
	return &Listener{
		WithdrawalQueueFilterer: withdrawalQueueFilterer,
		LidoFilterer:            lidoFilterer,
		OwnerAddr:               ownerAddr,
	}
}

type WatchOptionFn func(*WatchOptions)

type WatchOptions struct {
	FilterOwners     []common.Address
	FilterSenders    []common.Address
	FilterRequesters []common.Address
	FilterFrom       []*big.Int
	FilterTo         []*big.Int
}

func WithFilterOwners(addrs []common.Address) WatchOptionFn {
	return func(opts *WatchOptions) {
		opts.FilterOwners = addrs
	}
}

func WithFilterSenders(addrs []common.Address) WatchOptionFn {
	return func(opts *WatchOptions) {
		opts.FilterSenders = addrs
	}
}

func WithFilterRequesters(addrs []common.Address) WatchOptionFn {
	return func(opts *WatchOptions) {
		opts.FilterRequesters = addrs
	}
}

func WithFilterFrom(indices []*big.Int) WatchOptionFn {
	return func(opts *WatchOptions) {
		opts.FilterFrom = indices
	}
}

func WithFilterTo(indices []*big.Int) WatchOptionFn {
	return func(opts *WatchOptions) {
		opts.FilterTo = indices
	}
}

func (l *Listener) ListenForSubmitteds(ctx context.Context, optionFns ...WatchOptionFn) error {
	slog.Info("listening for Submitted events")
	opts := &WatchOptions{
		FilterOwners: []common.Address{l.OwnerAddr},
	}
	for _, optFn := range optionFns {
		optFn(opts)
	}

	sink := make(chan *onchain.LidoSubmitted)
	sub, err := l.LidoFilterer.WatchSubmitted(nil, sink, opts.FilterSenders)
	if err != nil {
		return fmt.Errorf("watch submitted: %w", err)
	}
	defer sub.Unsubscribe()

	select {
	case event := <-sink:
		err = l.ProcessSubmitteds(event)
		if err != nil {
			return fmt.Errorf("process submitted: %w", err)
		}

	case err := <-sub.Err():
		return fmt.Errorf("subscription error: %w", err)

	case <-ctx.Done():
		return nil
	}

	return nil
}
func (l *Listener) ListenForWithdrawalRequests(ctx context.Context, optionFns ...WatchOptionFn) error {
	slog.Info("listening for WithdrawalRequest events")
	opts := &WatchOptions{
		FilterOwners: []common.Address{l.OwnerAddr},
	}

	for _, optFn := range optionFns {
		optFn(opts)
	}
	sink := make(chan *onchain.WithdrawalQueueWithdrawalRequested)
	sub, err := l.WithdrawalQueueFilterer.WatchWithdrawalRequested(nil, sink, nil, opts.FilterRequesters, opts.FilterOwners)
	if err != nil {
		return fmt.Errorf("watch withdrawal requested: %w", err)
	}
	defer sub.Unsubscribe()

	select {
	case event := <-sink:
		err = l.ProcessWithdrawalRequests(event)
		if err != nil {
			return fmt.Errorf("process withdrawal requested: %w", err)
		}

	case err := <-sub.Err():
		return fmt.Errorf("subscription error: %w", err)

	case <-ctx.Done():
		return nil
	}

	return nil
}
func (l *Listener) ListenForWithdrawalFinalisations(ctx context.Context, optionFns ...WatchOptionFn) error {
	slog.Info("listening for WithdrawalFinalisation events")
	opts := &WatchOptions{
		FilterOwners: []common.Address{l.OwnerAddr},
	}
	for _, optFn := range optionFns {
		optFn(opts)
	}
	sink := make(chan *onchain.WithdrawalQueueWithdrawalsFinalized)
	sub, err := l.WithdrawalQueueFilterer.WatchWithdrawalsFinalized(nil, sink, opts.FilterFrom, opts.FilterTo)
	if err != nil {
		return fmt.Errorf("watch withdrawal finalized: %w", err)
	}
	defer sub.Unsubscribe()

	select {
	case event := <-sink:
		err = l.ProcessWithdrawalFinalisations(event)
		if err != nil {
			return fmt.Errorf("process withdrawal finalized: %w", err)
		}

	case err := <-sub.Err():
		return fmt.Errorf("subscription error: %w", err)

	case <-ctx.Done():
		return nil
	}

	return nil
}
func (l *Listener) ListenForWithdrawalClaimeds(ctx context.Context, optionFns ...WatchOptionFn) error {
	slog.Info("listening for WithdrawalClaimed events")
	opts := &WatchOptions{
		FilterOwners: []common.Address{l.OwnerAddr},
	}

	for _, optFn := range optionFns {
		optFn(opts)
	}
	sink := make(chan *onchain.WithdrawalQueueWithdrawalClaimed)
	sub, err := l.WithdrawalQueueFilterer.WatchWithdrawalClaimed(nil, sink, nil, opts.FilterOwners, nil)
	if err != nil {
		return fmt.Errorf("watch withdrawal claimed: %w", err)
	}
	defer sub.Unsubscribe()

	select {
	case event := <-sink:
		err = l.ProcessWithdrawalClaimeds(event)
		if err != nil {
			return fmt.Errorf("process withdrawal claimed: %w", err)
		}

	case err := <-sub.Err():
		return fmt.Errorf("subscription error: %w", err)

	case <-ctx.Done():
		return nil
	}

	return nil
}

func (l *Listener) BackfillSubmitteds(ctx context.Context, from uint64, optionFns ...WatchOptionFn) error {
	slog.Info("backfilling Submitted events")
	opts := &WatchOptions{
		FilterOwners: []common.Address{l.OwnerAddr},
	}

	for _, optFn := range optionFns {
		optFn(opts)
	}
	iterator, err := l.LidoFilterer.FilterSubmitted(&bind.FilterOpts{
		Start: from,
		End:   nil,
	}, opts.FilterSenders)
	if err != nil {
		return fmt.Errorf("filter deposit transferred: %w", err)
	}
	defer iterator.Close()

	for iterator.Next() {
		if err := iterator.Error(); err != nil {
			return fmt.Errorf("iterator error: %w", err)
		}

		err = l.ProcessSubmitteds(iterator.Event)
		if err != nil {
			return fmt.Errorf("process Submitteds: %w", err)
		}
	}

	return nil
}

func (l *Listener) BackfillWithdrawalRequests(ctx context.Context, from uint64, optionFns ...WatchOptionFn) error {
	slog.Info("backfilling WithdrawalRequests events")
	opts := &WatchOptions{
		FilterOwners: []common.Address{l.OwnerAddr},
	}

	for _, optFn := range optionFns {
		optFn(opts)
	}
	iterator, err := l.WithdrawalQueueFilterer.FilterWithdrawalRequested(&bind.FilterOpts{
		Start: from,
		End:   nil,
	}, []*big.Int{}, nil, opts.FilterRequesters)
	if err != nil {
		return fmt.Errorf("filter withdrawal requested: %w", err)
	}
	defer iterator.Close()

	for iterator.Next() {
		if err := iterator.Error(); err != nil {
			return fmt.Errorf("iterator error: %w", err)
		}

		err = l.ProcessWithdrawalRequests(iterator.Event)
		if err != nil {
			return fmt.Errorf("process withdrawal requests: %w", err)
		}
	}

	return nil
}

func (l *Listener) BackfillWithdrawalFinalisations(ctx context.Context, from uint64, optionFns ...WatchOptionFn) error {
	slog.Info("backfilling WithdrawalFinalisations events")
	opts := &WatchOptions{
		FilterOwners: []common.Address{l.OwnerAddr},
	}

	for _, optFn := range optionFns {
		optFn(opts)
	}
	iterator, err := l.WithdrawalQueueFilterer.FilterWithdrawalsFinalized(&bind.FilterOpts{
		Start: from,
		End:   nil,
	}, opts.FilterFrom, opts.FilterTo)
	if err != nil {
		return fmt.Errorf("filter withdrawal requested: %w", err)
	}
	defer iterator.Close()

	for iterator.Next() {
		if err := iterator.Error(); err != nil {
			return fmt.Errorf("iterator error: %w", err)
		}

		err = l.ProcessWithdrawalFinalisations(iterator.Event)
		if err != nil {
			return fmt.Errorf("process withdrawal requests: %w", err)
		}
	}

	return nil
}

func (l *Listener) BackfillWithdrawalClaimeds(ctx context.Context, from uint64, optionFns ...WatchOptionFn) error {
	slog.Info("backfilling WithdrawalClaimeds events")
	opts := &WatchOptions{
		FilterOwners: []common.Address{l.OwnerAddr},
	}

	for _, optFn := range optionFns {
		optFn(opts)
	}
	iterator, err := l.WithdrawalQueueFilterer.FilterWithdrawalClaimed(&bind.FilterOpts{
		Start: from,
		End:   nil,
	}, []*big.Int{}, opts.FilterOwners, nil)
	if err != nil {
		return fmt.Errorf("filter withdrawal requested: %w", err)
	}
	defer iterator.Close()

	for iterator.Next() {
		if err := iterator.Error(); err != nil {
			return fmt.Errorf("iterator error: %w", err)
		}

		err = l.ProcessWithdrawalClaimeds(iterator.Event)
		if err != nil {
			return fmt.Errorf("process withdrawal requests: %w", err)
		}
	}

	return nil
}

func (l *Listener) ProcessSubmitteds(in *onchain.LidoSubmitted) error {
	slog.Info("submitted", "index", in.Raw.Index, "tx", in.Raw.TxHash.String(), "sender", in.Sender.Hex(), "amount", in.Amount.String())
	return nil
}
func (l *Listener) ProcessWithdrawalRequests(in *onchain.WithdrawalQueueWithdrawalRequested) error {
	slog.Info("withdrawal requested", "index", in.Raw.Index, "tx", in.Raw.TxHash.String(), "requester", in.Requestor.Hex(), "owner", in.Owner.Hex(), "amount_shares", in.AmountOfShares.String(), "amount_steth", in.AmountOfStETH.String())
	return nil
}
func (l *Listener) ProcessWithdrawalFinalisations(in *onchain.WithdrawalQueueWithdrawalsFinalized) error {
	slog.Info("withdrawal finalized", "from", in.From.String(), "to", in.To.String(), "eth_locked", in.AmountOfETHLocked.String(), "shares_to_burn", in.SharesToBurn.String())
	return nil
}
func (l *Listener) ProcessWithdrawalClaimeds(in *onchain.WithdrawalQueueWithdrawalClaimed) error {
	slog.Info("withdrawal claimed", "index", in.Raw.Index, "tx", in.Raw.TxHash.String(), "receiver", in.Receiver.Hex(), "amount_eth", in.AmountOfETH.String())
	return nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// WithdrawalQueuePermitInput is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalQueuePermitInput struct {
	Value    *big.Int
	Deadline *big.Int
	V        uint8
	R        [32]byte
	S        [32]byte
}

// WithdrawalQueueMetaData contains all meta data concerning the WithdrawalQueue contract.
var WithdrawalQueueMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOfETH\",\"type\":\"uint256\"}],\"name\":\"WithdrawalClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOfStETH\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOfShares\",\"type\":\"uint256\"}],\"name\":\"WithdrawalRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOfETHLocked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sharesToBurn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"WithdrawalsFinalized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_requestIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_hints\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"claimWithdrawalsTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_requestIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_firstIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lastIndex\",\"type\":\"uint256\"}],\"name\":\"findCheckpointHints\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"hintIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structWithdrawalQueue.PermitInput\",\"name\":\"_permit\",\"type\":\"tuple\"}],\"name\":\"requestWithdrawalsWithPermit\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"requestIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WithdrawalQueueABI is the input ABI used to generate the binding from.
// Deprecated: Use WithdrawalQueueMetaData.ABI instead.
var WithdrawalQueueABI = WithdrawalQueueMetaData.ABI

// WithdrawalQueue is an auto generated Go binding around an Ethereum contract.
type WithdrawalQueue struct {
	WithdrawalQueueCaller     // Read-only binding to the contract
	WithdrawalQueueTransactor // Write-only binding to the contract
	WithdrawalQueueFilterer   // Log filterer for contract events
}

// WithdrawalQueueCaller is an auto generated read-only Go binding around an Ethereum contract.
type WithdrawalQueueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawalQueueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WithdrawalQueueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawalQueueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WithdrawalQueueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawalQueueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WithdrawalQueueSession struct {
	Contract     *WithdrawalQueue  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WithdrawalQueueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WithdrawalQueueCallerSession struct {
	Contract *WithdrawalQueueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// WithdrawalQueueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WithdrawalQueueTransactorSession struct {
	Contract     *WithdrawalQueueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// WithdrawalQueueRaw is an auto generated low-level Go binding around an Ethereum contract.
type WithdrawalQueueRaw struct {
	Contract *WithdrawalQueue // Generic contract binding to access the raw methods on
}

// WithdrawalQueueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WithdrawalQueueCallerRaw struct {
	Contract *WithdrawalQueueCaller // Generic read-only contract binding to access the raw methods on
}

// WithdrawalQueueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WithdrawalQueueTransactorRaw struct {
	Contract *WithdrawalQueueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWithdrawalQueue creates a new instance of WithdrawalQueue, bound to a specific deployed contract.
func NewWithdrawalQueue(address common.Address, backend bind.ContractBackend) (*WithdrawalQueue, error) {
	contract, err := bindWithdrawalQueue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WithdrawalQueue{WithdrawalQueueCaller: WithdrawalQueueCaller{contract: contract}, WithdrawalQueueTransactor: WithdrawalQueueTransactor{contract: contract}, WithdrawalQueueFilterer: WithdrawalQueueFilterer{contract: contract}}, nil
}

// NewWithdrawalQueueCaller creates a new read-only instance of WithdrawalQueue, bound to a specific deployed contract.
func NewWithdrawalQueueCaller(address common.Address, caller bind.ContractCaller) (*WithdrawalQueueCaller, error) {
	contract, err := bindWithdrawalQueue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawalQueueCaller{contract: contract}, nil
}

// NewWithdrawalQueueTransactor creates a new write-only instance of WithdrawalQueue, bound to a specific deployed contract.
func NewWithdrawalQueueTransactor(address common.Address, transactor bind.ContractTransactor) (*WithdrawalQueueTransactor, error) {
	contract, err := bindWithdrawalQueue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawalQueueTransactor{contract: contract}, nil
}

// NewWithdrawalQueueFilterer creates a new log filterer instance of WithdrawalQueue, bound to a specific deployed contract.
func NewWithdrawalQueueFilterer(address common.Address, filterer bind.ContractFilterer) (*WithdrawalQueueFilterer, error) {
	contract, err := bindWithdrawalQueue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WithdrawalQueueFilterer{contract: contract}, nil
}

// bindWithdrawalQueue binds a generic wrapper to an already deployed contract.
func bindWithdrawalQueue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WithdrawalQueueMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WithdrawalQueue *WithdrawalQueueRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WithdrawalQueue.Contract.WithdrawalQueueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WithdrawalQueue *WithdrawalQueueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithdrawalQueue.Contract.WithdrawalQueueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WithdrawalQueue *WithdrawalQueueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WithdrawalQueue.Contract.WithdrawalQueueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WithdrawalQueue *WithdrawalQueueCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WithdrawalQueue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WithdrawalQueue *WithdrawalQueueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithdrawalQueue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WithdrawalQueue *WithdrawalQueueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WithdrawalQueue.Contract.contract.Transact(opts, method, params...)
}

// FindCheckpointHints is a free data retrieval call binding the contract method 0x62abe3fa.
//
// Solidity: function findCheckpointHints(uint256[] _requestIds, uint256 _firstIndex, uint256 _lastIndex) view returns(uint256[] hintIds)
func (_WithdrawalQueue *WithdrawalQueueCaller) FindCheckpointHints(opts *bind.CallOpts, _requestIds []*big.Int, _firstIndex *big.Int, _lastIndex *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _WithdrawalQueue.contract.Call(opts, &out, "findCheckpointHints", _requestIds, _firstIndex, _lastIndex)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// FindCheckpointHints is a free data retrieval call binding the contract method 0x62abe3fa.
//
// Solidity: function findCheckpointHints(uint256[] _requestIds, uint256 _firstIndex, uint256 _lastIndex) view returns(uint256[] hintIds)
func (_WithdrawalQueue *WithdrawalQueueSession) FindCheckpointHints(_requestIds []*big.Int, _firstIndex *big.Int, _lastIndex *big.Int) ([]*big.Int, error) {
	return _WithdrawalQueue.Contract.FindCheckpointHints(&_WithdrawalQueue.CallOpts, _requestIds, _firstIndex, _lastIndex)
}

// FindCheckpointHints is a free data retrieval call binding the contract method 0x62abe3fa.
//
// Solidity: function findCheckpointHints(uint256[] _requestIds, uint256 _firstIndex, uint256 _lastIndex) view returns(uint256[] hintIds)
func (_WithdrawalQueue *WithdrawalQueueCallerSession) FindCheckpointHints(_requestIds []*big.Int, _firstIndex *big.Int, _lastIndex *big.Int) ([]*big.Int, error) {
	return _WithdrawalQueue.Contract.FindCheckpointHints(&_WithdrawalQueue.CallOpts, _requestIds, _firstIndex, _lastIndex)
}

// ClaimWithdrawalsTo is a paid mutator transaction binding the contract method 0x5e7eead9.
//
// Solidity: function claimWithdrawalsTo(uint256[] _requestIds, uint256[] _hints, address _recipient) returns()
func (_WithdrawalQueue *WithdrawalQueueTransactor) ClaimWithdrawalsTo(opts *bind.TransactOpts, _requestIds []*big.Int, _hints []*big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _WithdrawalQueue.contract.Transact(opts, "claimWithdrawalsTo", _requestIds, _hints, _recipient)
}

// ClaimWithdrawalsTo is a paid mutator transaction binding the contract method 0x5e7eead9.
//
// Solidity: function claimWithdrawalsTo(uint256[] _requestIds, uint256[] _hints, address _recipient) returns()
func (_WithdrawalQueue *WithdrawalQueueSession) ClaimWithdrawalsTo(_requestIds []*big.Int, _hints []*big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _WithdrawalQueue.Contract.ClaimWithdrawalsTo(&_WithdrawalQueue.TransactOpts, _requestIds, _hints, _recipient)
}

// ClaimWithdrawalsTo is a paid mutator transaction binding the contract method 0x5e7eead9.
//
// Solidity: function claimWithdrawalsTo(uint256[] _requestIds, uint256[] _hints, address _recipient) returns()
func (_WithdrawalQueue *WithdrawalQueueTransactorSession) ClaimWithdrawalsTo(_requestIds []*big.Int, _hints []*big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _WithdrawalQueue.Contract.ClaimWithdrawalsTo(&_WithdrawalQueue.TransactOpts, _requestIds, _hints, _recipient)
}

// RequestWithdrawalsWithPermit is a paid mutator transaction binding the contract method 0xacf41e4d.
//
// Solidity: function requestWithdrawalsWithPermit(uint256[] _amounts, address _owner, (uint256,uint256,uint8,bytes32,bytes32) _permit) returns(uint256[] requestIds)
func (_WithdrawalQueue *WithdrawalQueueTransactor) RequestWithdrawalsWithPermit(opts *bind.TransactOpts, _amounts []*big.Int, _owner common.Address, _permit WithdrawalQueuePermitInput) (*types.Transaction, error) {
	return _WithdrawalQueue.contract.Transact(opts, "requestWithdrawalsWithPermit", _amounts, _owner, _permit)
}

// RequestWithdrawalsWithPermit is a paid mutator transaction binding the contract method 0xacf41e4d.
//
// Solidity: function requestWithdrawalsWithPermit(uint256[] _amounts, address _owner, (uint256,uint256,uint8,bytes32,bytes32) _permit) returns(uint256[] requestIds)
func (_WithdrawalQueue *WithdrawalQueueSession) RequestWithdrawalsWithPermit(_amounts []*big.Int, _owner common.Address, _permit WithdrawalQueuePermitInput) (*types.Transaction, error) {
	return _WithdrawalQueue.Contract.RequestWithdrawalsWithPermit(&_WithdrawalQueue.TransactOpts, _amounts, _owner, _permit)
}

// RequestWithdrawalsWithPermit is a paid mutator transaction binding the contract method 0xacf41e4d.
//
// Solidity: function requestWithdrawalsWithPermit(uint256[] _amounts, address _owner, (uint256,uint256,uint8,bytes32,bytes32) _permit) returns(uint256[] requestIds)
func (_WithdrawalQueue *WithdrawalQueueTransactorSession) RequestWithdrawalsWithPermit(_amounts []*big.Int, _owner common.Address, _permit WithdrawalQueuePermitInput) (*types.Transaction, error) {
	return _WithdrawalQueue.Contract.RequestWithdrawalsWithPermit(&_WithdrawalQueue.TransactOpts, _amounts, _owner, _permit)
}

// WithdrawalQueueWithdrawalClaimedIterator is returned from FilterWithdrawalClaimed and is used to iterate over the raw logs and unpacked data for WithdrawalClaimed events raised by the WithdrawalQueue contract.
type WithdrawalQueueWithdrawalClaimedIterator struct {
	Event *WithdrawalQueueWithdrawalClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawalQueueWithdrawalClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalQueueWithdrawalClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawalQueueWithdrawalClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawalQueueWithdrawalClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalQueueWithdrawalClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalQueueWithdrawalClaimed represents a WithdrawalClaimed event raised by the WithdrawalQueue contract.
type WithdrawalQueueWithdrawalClaimed struct {
	RequestId   *big.Int
	Owner       common.Address
	Receiver    common.Address
	AmountOfETH *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalClaimed is a free log retrieval operation binding the contract event 0x6ad26c5e238e7d002799f9a5db07e81ef14e37386ae03496d7a7ef04713e145b.
//
// Solidity: event WithdrawalClaimed(uint256 indexed requestId, address indexed owner, address indexed receiver, uint256 amountOfETH)
func (_WithdrawalQueue *WithdrawalQueueFilterer) FilterWithdrawalClaimed(opts *bind.FilterOpts, requestId []*big.Int, owner []common.Address, receiver []common.Address) (*WithdrawalQueueWithdrawalClaimedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _WithdrawalQueue.contract.FilterLogs(opts, "WithdrawalClaimed", requestIdRule, ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawalQueueWithdrawalClaimedIterator{contract: _WithdrawalQueue.contract, event: "WithdrawalClaimed", logs: logs, sub: sub}, nil
}

// WatchWithdrawalClaimed is a free log subscription operation binding the contract event 0x6ad26c5e238e7d002799f9a5db07e81ef14e37386ae03496d7a7ef04713e145b.
//
// Solidity: event WithdrawalClaimed(uint256 indexed requestId, address indexed owner, address indexed receiver, uint256 amountOfETH)
func (_WithdrawalQueue *WithdrawalQueueFilterer) WatchWithdrawalClaimed(opts *bind.WatchOpts, sink chan<- *WithdrawalQueueWithdrawalClaimed, requestId []*big.Int, owner []common.Address, receiver []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _WithdrawalQueue.contract.WatchLogs(opts, "WithdrawalClaimed", requestIdRule, ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalQueueWithdrawalClaimed)
				if err := _WithdrawalQueue.contract.UnpackLog(event, "WithdrawalClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalClaimed is a log parse operation binding the contract event 0x6ad26c5e238e7d002799f9a5db07e81ef14e37386ae03496d7a7ef04713e145b.
//
// Solidity: event WithdrawalClaimed(uint256 indexed requestId, address indexed owner, address indexed receiver, uint256 amountOfETH)
func (_WithdrawalQueue *WithdrawalQueueFilterer) ParseWithdrawalClaimed(log types.Log) (*WithdrawalQueueWithdrawalClaimed, error) {
	event := new(WithdrawalQueueWithdrawalClaimed)
	if err := _WithdrawalQueue.contract.UnpackLog(event, "WithdrawalClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalQueueWithdrawalRequestedIterator is returned from FilterWithdrawalRequested and is used to iterate over the raw logs and unpacked data for WithdrawalRequested events raised by the WithdrawalQueue contract.
type WithdrawalQueueWithdrawalRequestedIterator struct {
	Event *WithdrawalQueueWithdrawalRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawalQueueWithdrawalRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalQueueWithdrawalRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawalQueueWithdrawalRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawalQueueWithdrawalRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalQueueWithdrawalRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalQueueWithdrawalRequested represents a WithdrawalRequested event raised by the WithdrawalQueue contract.
type WithdrawalQueueWithdrawalRequested struct {
	RequestId      *big.Int
	Requestor      common.Address
	Owner          common.Address
	AmountOfStETH  *big.Int
	AmountOfShares *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequested is a free log retrieval operation binding the contract event 0xf0cb471f23fb74ea44b8252eb1881a2dca546288d9f6e90d1a0e82fe0ed342ab.
//
// Solidity: event WithdrawalRequested(uint256 indexed requestId, address indexed requestor, address indexed owner, uint256 amountOfStETH, uint256 amountOfShares)
func (_WithdrawalQueue *WithdrawalQueueFilterer) FilterWithdrawalRequested(opts *bind.FilterOpts, requestId []*big.Int, requestor []common.Address, owner []common.Address) (*WithdrawalQueueWithdrawalRequestedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requestorRule []interface{}
	for _, requestorItem := range requestor {
		requestorRule = append(requestorRule, requestorItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WithdrawalQueue.contract.FilterLogs(opts, "WithdrawalRequested", requestIdRule, requestorRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawalQueueWithdrawalRequestedIterator{contract: _WithdrawalQueue.contract, event: "WithdrawalRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequested is a free log subscription operation binding the contract event 0xf0cb471f23fb74ea44b8252eb1881a2dca546288d9f6e90d1a0e82fe0ed342ab.
//
// Solidity: event WithdrawalRequested(uint256 indexed requestId, address indexed requestor, address indexed owner, uint256 amountOfStETH, uint256 amountOfShares)
func (_WithdrawalQueue *WithdrawalQueueFilterer) WatchWithdrawalRequested(opts *bind.WatchOpts, sink chan<- *WithdrawalQueueWithdrawalRequested, requestId []*big.Int, requestor []common.Address, owner []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requestorRule []interface{}
	for _, requestorItem := range requestor {
		requestorRule = append(requestorRule, requestorItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WithdrawalQueue.contract.WatchLogs(opts, "WithdrawalRequested", requestIdRule, requestorRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalQueueWithdrawalRequested)
				if err := _WithdrawalQueue.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalRequested is a log parse operation binding the contract event 0xf0cb471f23fb74ea44b8252eb1881a2dca546288d9f6e90d1a0e82fe0ed342ab.
//
// Solidity: event WithdrawalRequested(uint256 indexed requestId, address indexed requestor, address indexed owner, uint256 amountOfStETH, uint256 amountOfShares)
func (_WithdrawalQueue *WithdrawalQueueFilterer) ParseWithdrawalRequested(log types.Log) (*WithdrawalQueueWithdrawalRequested, error) {
	event := new(WithdrawalQueueWithdrawalRequested)
	if err := _WithdrawalQueue.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalQueueWithdrawalsFinalizedIterator is returned from FilterWithdrawalsFinalized and is used to iterate over the raw logs and unpacked data for WithdrawalsFinalized events raised by the WithdrawalQueue contract.
type WithdrawalQueueWithdrawalsFinalizedIterator struct {
	Event *WithdrawalQueueWithdrawalsFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawalQueueWithdrawalsFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalQueueWithdrawalsFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawalQueueWithdrawalsFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawalQueueWithdrawalsFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalQueueWithdrawalsFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalQueueWithdrawalsFinalized represents a WithdrawalsFinalized event raised by the WithdrawalQueue contract.
type WithdrawalQueueWithdrawalsFinalized struct {
	From              *big.Int
	To                *big.Int
	AmountOfETHLocked *big.Int
	SharesToBurn      *big.Int
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalsFinalized is a free log retrieval operation binding the contract event 0x197874c72af6a06fb0aa4fab45fd39c7cb61ac0992159872dc3295207da7e9eb.
//
// Solidity: event WithdrawalsFinalized(uint256 indexed from, uint256 indexed to, uint256 amountOfETHLocked, uint256 sharesToBurn, uint256 timestamp)
func (_WithdrawalQueue *WithdrawalQueueFilterer) FilterWithdrawalsFinalized(opts *bind.FilterOpts, from []*big.Int, to []*big.Int) (*WithdrawalQueueWithdrawalsFinalizedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WithdrawalQueue.contract.FilterLogs(opts, "WithdrawalsFinalized", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawalQueueWithdrawalsFinalizedIterator{contract: _WithdrawalQueue.contract, event: "WithdrawalsFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawalsFinalized is a free log subscription operation binding the contract event 0x197874c72af6a06fb0aa4fab45fd39c7cb61ac0992159872dc3295207da7e9eb.
//
// Solidity: event WithdrawalsFinalized(uint256 indexed from, uint256 indexed to, uint256 amountOfETHLocked, uint256 sharesToBurn, uint256 timestamp)
func (_WithdrawalQueue *WithdrawalQueueFilterer) WatchWithdrawalsFinalized(opts *bind.WatchOpts, sink chan<- *WithdrawalQueueWithdrawalsFinalized, from []*big.Int, to []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WithdrawalQueue.contract.WatchLogs(opts, "WithdrawalsFinalized", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalQueueWithdrawalsFinalized)
				if err := _WithdrawalQueue.contract.UnpackLog(event, "WithdrawalsFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalsFinalized is a log parse operation binding the contract event 0x197874c72af6a06fb0aa4fab45fd39c7cb61ac0992159872dc3295207da7e9eb.
//
// Solidity: event WithdrawalsFinalized(uint256 indexed from, uint256 indexed to, uint256 amountOfETHLocked, uint256 sharesToBurn, uint256 timestamp)
func (_WithdrawalQueue *WithdrawalQueueFilterer) ParseWithdrawalsFinalized(log types.Log) (*WithdrawalQueueWithdrawalsFinalized, error) {
	event := new(WithdrawalQueueWithdrawalsFinalized)
	if err := _WithdrawalQueue.contract.UnpackLog(event, "WithdrawalsFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

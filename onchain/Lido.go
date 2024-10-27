// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package onchain

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

// LidoMetaData contains all meta data concerning the Lido contract.
var LidoMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referral\",\"type\":\"address\"}],\"name\":\"Submitted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_referral\",\"type\":\"address\"}],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// LidoABI is the input ABI used to generate the binding from.
// Deprecated: Use LidoMetaData.ABI instead.
var LidoABI = LidoMetaData.ABI

// Lido is an auto generated Go binding around an Ethereum contract.
type Lido struct {
	LidoCaller     // Read-only binding to the contract
	LidoTransactor // Write-only binding to the contract
	LidoFilterer   // Log filterer for contract events
}

// LidoCaller is an auto generated read-only Go binding around an Ethereum contract.
type LidoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LidoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LidoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LidoSession struct {
	Contract     *Lido             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LidoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LidoCallerSession struct {
	Contract *LidoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LidoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LidoTransactorSession struct {
	Contract     *LidoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LidoRaw is an auto generated low-level Go binding around an Ethereum contract.
type LidoRaw struct {
	Contract *Lido // Generic contract binding to access the raw methods on
}

// LidoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LidoCallerRaw struct {
	Contract *LidoCaller // Generic read-only contract binding to access the raw methods on
}

// LidoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LidoTransactorRaw struct {
	Contract *LidoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLido creates a new instance of Lido, bound to a specific deployed contract.
func NewLido(address common.Address, backend bind.ContractBackend) (*Lido, error) {
	contract, err := bindLido(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lido{LidoCaller: LidoCaller{contract: contract}, LidoTransactor: LidoTransactor{contract: contract}, LidoFilterer: LidoFilterer{contract: contract}}, nil
}

// NewLidoCaller creates a new read-only instance of Lido, bound to a specific deployed contract.
func NewLidoCaller(address common.Address, caller bind.ContractCaller) (*LidoCaller, error) {
	contract, err := bindLido(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LidoCaller{contract: contract}, nil
}

// NewLidoTransactor creates a new write-only instance of Lido, bound to a specific deployed contract.
func NewLidoTransactor(address common.Address, transactor bind.ContractTransactor) (*LidoTransactor, error) {
	contract, err := bindLido(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LidoTransactor{contract: contract}, nil
}

// NewLidoFilterer creates a new log filterer instance of Lido, bound to a specific deployed contract.
func NewLidoFilterer(address common.Address, filterer bind.ContractFilterer) (*LidoFilterer, error) {
	contract, err := bindLido(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LidoFilterer{contract: contract}, nil
}

// bindLido binds a generic wrapper to an already deployed contract.
func bindLido(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LidoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lido *LidoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lido.Contract.LidoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lido *LidoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lido.Contract.LidoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lido *LidoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lido.Contract.LidoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lido *LidoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lido.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lido *LidoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lido.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lido *LidoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lido.Contract.contract.Transact(opts, method, params...)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_Lido *LidoTransactor) Submit(opts *bind.TransactOpts, _referral common.Address) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "submit", _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_Lido *LidoSession) Submit(_referral common.Address) (*types.Transaction, error) {
	return _Lido.Contract.Submit(&_Lido.TransactOpts, _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_Lido *LidoTransactorSession) Submit(_referral common.Address) (*types.Transaction, error) {
	return _Lido.Contract.Submit(&_Lido.TransactOpts, _referral)
}

// LidoSubmittedIterator is returned from FilterSubmitted and is used to iterate over the raw logs and unpacked data for Submitted events raised by the Lido contract.
type LidoSubmittedIterator struct {
	Event *LidoSubmitted // Event containing the contract specifics and raw log

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
func (it *LidoSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoSubmitted)
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
		it.Event = new(LidoSubmitted)
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
func (it *LidoSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoSubmitted represents a Submitted event raised by the Lido contract.
type LidoSubmitted struct {
	Sender   common.Address
	Amount   *big.Int
	Referral common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubmitted is a free log retrieval operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_Lido *LidoFilterer) FilterSubmitted(opts *bind.FilterOpts, sender []common.Address) (*LidoSubmittedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Lido.contract.FilterLogs(opts, "Submitted", senderRule)
	if err != nil {
		return nil, err
	}
	return &LidoSubmittedIterator{contract: _Lido.contract, event: "Submitted", logs: logs, sub: sub}, nil
}

// WatchSubmitted is a free log subscription operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_Lido *LidoFilterer) WatchSubmitted(opts *bind.WatchOpts, sink chan<- *LidoSubmitted, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Lido.contract.WatchLogs(opts, "Submitted", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoSubmitted)
				if err := _Lido.contract.UnpackLog(event, "Submitted", log); err != nil {
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

// ParseSubmitted is a log parse operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_Lido *LidoFilterer) ParseSubmitted(log types.Log) (*LidoSubmitted, error) {
	event := new(LidoSubmitted)
	if err := _Lido.contract.UnpackLog(event, "Submitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

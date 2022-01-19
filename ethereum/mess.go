// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = klaytn.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StructsGuardianSet is an auto generated low-level Go binding around an user-defined struct.
type StructsGuardianSet struct {
	Keys           []common.Address
	ExpirationTime uint32
}

// StructsSignature is an auto generated low-level Go binding around an user-defined struct.
type StructsSignature struct {
	R             [32]byte
	S             [32]byte
	V             uint8
	GuardianIndex uint8
}

// StructsVM is an auto generated low-level Go binding around an user-defined struct.
type StructsVM struct {
	Version          uint8
	Timestamp        uint32
	Nonce            uint32
	EmitterChainId   uint16
	EmitterAddress   [32]byte
	Sequence         uint64
	ConsistencyLevel uint8
	Payload          []byte
	GuardianSetIndex uint32
	Signatures       []StructsSignature
	Hash             [32]byte
}

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const BytesLibBinRuntime = `73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122035b9208ed4ef97e2ae1508d5146c6344988abf7df29932228b08e2e424ce29a164736f6c637828302e382e31322d646576656c6f702e323032322e312e31322b636f6d6d69742e61373131393639390059`

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x607c6037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122035b9208ed4ef97e2ae1508d5146c6344988abf7df29932228b08e2e424ce29a164736f6c637828302e382e31322d646576656c6f702e323032322e312e31322b636f6d6d69742e61373131393639390059"

// DeployBytesLib deploys a new Klaytn contract, binding an instance of BytesLib to it.
func DeployBytesLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesLib, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BytesLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// BytesLib is an auto generated Go binding around a Klaytn contract.
type BytesLib struct {
	BytesLibCaller     // Read-only binding to the contract
	BytesLibTransactor // Write-only binding to the contract
	BytesLibFilterer   // Log filterer for contract events
}

// BytesLibCaller is an auto generated read-only Go binding around a Klaytn contract.
type BytesLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibTransactor is an auto generated write-only Go binding around a Klaytn contract.
type BytesLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type BytesLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type BytesLibSession struct {
	Contract     *BytesLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesLibCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type BytesLibCallerSession struct {
	Contract *BytesLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BytesLibTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type BytesLibTransactorSession struct {
	Contract     *BytesLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BytesLibRaw is an auto generated low-level Go binding around a Klaytn contract.
type BytesLibRaw struct {
	Contract *BytesLib // Generic contract binding to access the raw methods on
}

// BytesLibCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type BytesLibCallerRaw struct {
	Contract *BytesLibCaller // Generic read-only contract binding to access the raw methods on
}

// BytesLibTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type BytesLibTransactorRaw struct {
	Contract *BytesLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesLib creates a new instance of BytesLib, bound to a specific deployed contract.
func NewBytesLib(address common.Address, backend bind.ContractBackend) (*BytesLib, error) {
	contract, err := bindBytesLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// NewBytesLibCaller creates a new read-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibCaller(address common.Address, caller bind.ContractCaller) (*BytesLibCaller, error) {
	contract, err := bindBytesLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibCaller{contract: contract}, nil
}

// NewBytesLibTransactor creates a new write-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesLibTransactor, error) {
	contract, err := bindBytesLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibTransactor{contract: contract}, nil
}

// NewBytesLibFilterer creates a new log filterer instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesLibFilterer, error) {
	contract, err := bindBytesLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesLibFilterer{contract: contract}, nil
}

// bindBytesLib binds a generic wrapper to an already deployed contract.
func bindBytesLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.BytesLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transact(opts, method, params...)
}

// EventsABI is the input ABI used to generate the binding from.
const EventsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"oldGuardianIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newGuardianIndex\",\"type\":\"uint32\"}],\"name\":\"LogGuardianSetChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"emitter_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"LogMessagePublished\",\"type\":\"event\"}]"

// EventsBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const EventsBinRuntime = `6080604052600080fdfea26469706673582212204bed3c58e6363cb8d43d1a39427e9010b16a5e43dba85333c5e6e9dca0ac043364736f6c637828302e382e31322d646576656c6f702e323032322e312e31322b636f6d6d69742e61373131393639390059`

// EventsBin is the compiled bytecode used for deploying new contracts.
var EventsBin = "0x6080604052348015600f57600080fd5b50606580601d6000396000f3fe6080604052600080fdfea26469706673582212204bed3c58e6363cb8d43d1a39427e9010b16a5e43dba85333c5e6e9dca0ac043364736f6c637828302e382e31322d646576656c6f702e323032322e312e31322b636f6d6d69742e61373131393639390059"

// DeployEvents deploys a new Klaytn contract, binding an instance of Events to it.
func DeployEvents(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Events, error) {
	parsed, err := abi.JSON(strings.NewReader(EventsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EventsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Events{EventsCaller: EventsCaller{contract: contract}, EventsTransactor: EventsTransactor{contract: contract}, EventsFilterer: EventsFilterer{contract: contract}}, nil
}

// Events is an auto generated Go binding around a Klaytn contract.
type Events struct {
	EventsCaller     // Read-only binding to the contract
	EventsTransactor // Write-only binding to the contract
	EventsFilterer   // Log filterer for contract events
}

// EventsCaller is an auto generated read-only Go binding around a Klaytn contract.
type EventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsTransactor is an auto generated write-only Go binding around a Klaytn contract.
type EventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type EventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type EventsSession struct {
	Contract     *Events           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventsCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type EventsCallerSession struct {
	Contract *EventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EventsTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type EventsTransactorSession struct {
	Contract     *EventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventsRaw is an auto generated low-level Go binding around a Klaytn contract.
type EventsRaw struct {
	Contract *Events // Generic contract binding to access the raw methods on
}

// EventsCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type EventsCallerRaw struct {
	Contract *EventsCaller // Generic read-only contract binding to access the raw methods on
}

// EventsTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type EventsTransactorRaw struct {
	Contract *EventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEvents creates a new instance of Events, bound to a specific deployed contract.
func NewEvents(address common.Address, backend bind.ContractBackend) (*Events, error) {
	contract, err := bindEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Events{EventsCaller: EventsCaller{contract: contract}, EventsTransactor: EventsTransactor{contract: contract}, EventsFilterer: EventsFilterer{contract: contract}}, nil
}

// NewEventsCaller creates a new read-only instance of Events, bound to a specific deployed contract.
func NewEventsCaller(address common.Address, caller bind.ContractCaller) (*EventsCaller, error) {
	contract, err := bindEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventsCaller{contract: contract}, nil
}

// NewEventsTransactor creates a new write-only instance of Events, bound to a specific deployed contract.
func NewEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*EventsTransactor, error) {
	contract, err := bindEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventsTransactor{contract: contract}, nil
}

// NewEventsFilterer creates a new log filterer instance of Events, bound to a specific deployed contract.
func NewEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*EventsFilterer, error) {
	contract, err := bindEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventsFilterer{contract: contract}, nil
}

// bindEvents binds a generic wrapper to an already deployed contract.
func bindEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Events *EventsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Events.Contract.EventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Events *EventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Events.Contract.EventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Events *EventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Events.Contract.EventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Events *EventsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Events.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Events *EventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Events.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Events *EventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Events.Contract.contract.Transact(opts, method, params...)
}

// EventsLogGuardianSetChangedIterator is returned from FilterLogGuardianSetChanged and is used to iterate over the raw logs and unpacked data for LogGuardianSetChanged events raised by the Events contract.
type EventsLogGuardianSetChangedIterator struct {
	Event *EventsLogGuardianSetChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsLogGuardianSetChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsLogGuardianSetChanged)
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
		it.Event = new(EventsLogGuardianSetChanged)
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
func (it *EventsLogGuardianSetChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsLogGuardianSetChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsLogGuardianSetChanged represents a LogGuardianSetChanged event raised by the Events contract.
type EventsLogGuardianSetChanged struct {
	OldGuardianIndex uint32
	NewGuardianIndex uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogGuardianSetChanged is a free log retrieval operation binding the contract event 0xdfb80683934199683861bf00b64ecdf0984bbaf661bf27983dba382e99297a62.
//
// Solidity: event LogGuardianSetChanged(uint32 oldGuardianIndex, uint32 newGuardianIndex)
func (_Events *EventsFilterer) FilterLogGuardianSetChanged(opts *bind.FilterOpts) (*EventsLogGuardianSetChangedIterator, error) {

	logs, sub, err := _Events.contract.FilterLogs(opts, "LogGuardianSetChanged")
	if err != nil {
		return nil, err
	}
	return &EventsLogGuardianSetChangedIterator{contract: _Events.contract, event: "LogGuardianSetChanged", logs: logs, sub: sub}, nil
}

// WatchLogGuardianSetChanged is a free log subscription operation binding the contract event 0xdfb80683934199683861bf00b64ecdf0984bbaf661bf27983dba382e99297a62.
//
// Solidity: event LogGuardianSetChanged(uint32 oldGuardianIndex, uint32 newGuardianIndex)
func (_Events *EventsFilterer) WatchLogGuardianSetChanged(opts *bind.WatchOpts, sink chan<- *EventsLogGuardianSetChanged) (event.Subscription, error) {

	logs, sub, err := _Events.contract.WatchLogs(opts, "LogGuardianSetChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsLogGuardianSetChanged)
				if err := _Events.contract.UnpackLog(event, "LogGuardianSetChanged", log); err != nil {
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

// ParseLogGuardianSetChanged is a log parse operation binding the contract event 0xdfb80683934199683861bf00b64ecdf0984bbaf661bf27983dba382e99297a62.
//
// Solidity: event LogGuardianSetChanged(uint32 oldGuardianIndex, uint32 newGuardianIndex)
func (_Events *EventsFilterer) ParseLogGuardianSetChanged(log types.Log) (*EventsLogGuardianSetChanged, error) {
	event := new(EventsLogGuardianSetChanged)
	if err := _Events.contract.UnpackLog(event, "LogGuardianSetChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsLogMessagePublishedIterator is returned from FilterLogMessagePublished and is used to iterate over the raw logs and unpacked data for LogMessagePublished events raised by the Events contract.
type EventsLogMessagePublishedIterator struct {
	Event *EventsLogMessagePublished // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsLogMessagePublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsLogMessagePublished)
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
		it.Event = new(EventsLogMessagePublished)
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
func (it *EventsLogMessagePublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsLogMessagePublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsLogMessagePublished represents a LogMessagePublished event raised by the Events contract.
type EventsLogMessagePublished struct {
	EmitterAddress common.Address
	Nonce          uint32
	Payload        []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogMessagePublished is a free log retrieval operation binding the contract event 0x95aa04882088c741573df04f9988e171e0d07d0b36f6518f4e03837b7ea0a023.
//
// Solidity: event LogMessagePublished(address emitter_address, uint32 nonce, bytes payload)
func (_Events *EventsFilterer) FilterLogMessagePublished(opts *bind.FilterOpts) (*EventsLogMessagePublishedIterator, error) {

	logs, sub, err := _Events.contract.FilterLogs(opts, "LogMessagePublished")
	if err != nil {
		return nil, err
	}
	return &EventsLogMessagePublishedIterator{contract: _Events.contract, event: "LogMessagePublished", logs: logs, sub: sub}, nil
}

// WatchLogMessagePublished is a free log subscription operation binding the contract event 0x95aa04882088c741573df04f9988e171e0d07d0b36f6518f4e03837b7ea0a023.
//
// Solidity: event LogMessagePublished(address emitter_address, uint32 nonce, bytes payload)
func (_Events *EventsFilterer) WatchLogMessagePublished(opts *bind.WatchOpts, sink chan<- *EventsLogMessagePublished) (event.Subscription, error) {

	logs, sub, err := _Events.contract.WatchLogs(opts, "LogMessagePublished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsLogMessagePublished)
				if err := _Events.contract.UnpackLog(event, "LogMessagePublished", log); err != nil {
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

// ParseLogMessagePublished is a log parse operation binding the contract event 0x95aa04882088c741573df04f9988e171e0d07d0b36f6518f4e03837b7ea0a023.
//
// Solidity: event LogMessagePublished(address emitter_address, uint32 nonce, bytes payload)
func (_Events *EventsFilterer) ParseLogMessagePublished(log types.Log) (*EventsLogMessagePublished, error) {
	event := new(EventsLogMessagePublished)
	if err := _Events.contract.UnpackLog(event, "LogMessagePublished", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GettersABI is the input ABI used to generate the binding from.
const GettersABI = "[{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentGuardianSetIndex\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"getGuardianSet\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"keys\",\"type\":\"address[]\"},{\"internalType\":\"uint32\",\"name\":\"expirationTime\",\"type\":\"uint32\"}],\"internalType\":\"structStructs.GuardianSet\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGuardianSetExpiry\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"governanceActionIsConsumed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceChainId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceContract\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"emitter\",\"type\":\"address\"}],\"name\":\"nextSequence\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// GettersBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const GettersBinRuntime = `608060405234801561001057600080fd5b506004361061009e5760003560e01c8063b172b22211610066578063b172b22214610173578063d60b347f1461017b578063eb8d3f12146101a7578063f951975a146101bd578063fbe3c2cd146101dd57600080fd5b80631a90a219146100a35780631cfe7951146100ba5780632c3c02a4146100d95780634cf842b51461010c5780639a8a059214610158575b600080fd5b6007545b6040519081526020015b60405180910390f35b60035463ffffffff165b60405163ffffffff90911681526020016100b1565b6100fc6100e736600461028e565b60009081526005602052604090205460ff1690565b60405190151581526020016100b1565b61013f61011a3660046102a7565b6001600160a01b031660009081526004602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff90911681526020016100b1565b60005461ffff165b60405161ffff90911681526020016100b1565b6001546100a7565b6100fc6101893660046102a7565b6001600160a01b031660009081526006602052604090205460ff1690565b600354640100000000900463ffffffff166100c4565b6101d06101cb3660046102d7565b6101ef565b6040516100b191906102fd565b60005462010000900461ffff16610160565b60408051808201825260608082526000602080840182905263ffffffff8616825260028152908490208451815492830281018401865294850182815293949390928492849184018282801561026d57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161024f575b50505091835250506001919091015463ffffffff1660209091015292915050565b6000602082840312156102a057600080fd5b5035919050565b6000602082840312156102b957600080fd5b81356001600160a01b03811681146102d057600080fd5b9392505050565b6000602082840312156102e957600080fd5b813563ffffffff811681146102d057600080fd5b6020808252825160408383015280516060840181905260009291820190839060808601905b8083101561034b5783516001600160a01b03168252928401926001929092019190840190610322565b5063ffffffff84880151166040870152809450505050509291505056fea26469706673582212209fef9e3bb81e105e38dee74a6cdb5fef49e7e26c67ce17ebabbc97472774bec664736f6c637828302e382e31322d646576656c6f702e323032322e312e31322b636f6d6d69742e61373131393639390059`

// GettersFuncSigs maps the 4-byte function signature to its string representation.
var GettersFuncSigs = map[string]string{
	"9a8a0592": "chainId()",
	"1cfe7951": "getCurrentGuardianSetIndex()",
	"f951975a": "getGuardianSet(uint32)",
	"eb8d3f12": "getGuardianSetExpiry()",
	"2c3c02a4": "governanceActionIsConsumed(bytes32)",
	"fbe3c2cd": "governanceChainId()",
	"b172b222": "governanceContract()",
	"d60b347f": "isInitialized(address)",
	"1a90a219": "messageFee()",
	"4cf842b5": "nextSequence(address)",
}

// GettersBin is the compiled bytecode used for deploying new contracts.
var GettersBin = "0x608060405234801561001057600080fd5b506103c4806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063b172b22211610066578063b172b22214610173578063d60b347f1461017b578063eb8d3f12146101a7578063f951975a146101bd578063fbe3c2cd146101dd57600080fd5b80631a90a219146100a35780631cfe7951146100ba5780632c3c02a4146100d95780634cf842b51461010c5780639a8a059214610158575b600080fd5b6007545b6040519081526020015b60405180910390f35b60035463ffffffff165b60405163ffffffff90911681526020016100b1565b6100fc6100e736600461028e565b60009081526005602052604090205460ff1690565b60405190151581526020016100b1565b61013f61011a3660046102a7565b6001600160a01b031660009081526004602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff90911681526020016100b1565b60005461ffff165b60405161ffff90911681526020016100b1565b6001546100a7565b6100fc6101893660046102a7565b6001600160a01b031660009081526006602052604090205460ff1690565b600354640100000000900463ffffffff166100c4565b6101d06101cb3660046102d7565b6101ef565b6040516100b191906102fd565b60005462010000900461ffff16610160565b60408051808201825260608082526000602080840182905263ffffffff8616825260028152908490208451815492830281018401865294850182815293949390928492849184018282801561026d57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161024f575b50505091835250506001919091015463ffffffff1660209091015292915050565b6000602082840312156102a057600080fd5b5035919050565b6000602082840312156102b957600080fd5b81356001600160a01b03811681146102d057600080fd5b9392505050565b6000602082840312156102e957600080fd5b813563ffffffff811681146102d057600080fd5b6020808252825160408383015280516060840181905260009291820190839060808601905b8083101561034b5783516001600160a01b03168252928401926001929092019190840190610322565b5063ffffffff84880151166040870152809450505050509291505056fea26469706673582212209fef9e3bb81e105e38dee74a6cdb5fef49e7e26c67ce17ebabbc97472774bec664736f6c637828302e382e31322d646576656c6f702e323032322e312e31322b636f6d6d69742e61373131393639390059"

// DeployGetters deploys a new Klaytn contract, binding an instance of Getters to it.
func DeployGetters(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Getters, error) {
	parsed, err := abi.JSON(strings.NewReader(GettersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GettersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Getters{GettersCaller: GettersCaller{contract: contract}, GettersTransactor: GettersTransactor{contract: contract}, GettersFilterer: GettersFilterer{contract: contract}}, nil
}

// Getters is an auto generated Go binding around a Klaytn contract.
type Getters struct {
	GettersCaller     // Read-only binding to the contract
	GettersTransactor // Write-only binding to the contract
	GettersFilterer   // Log filterer for contract events
}

// GettersCaller is an auto generated read-only Go binding around a Klaytn contract.
type GettersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GettersTransactor is an auto generated write-only Go binding around a Klaytn contract.
type GettersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GettersFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type GettersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GettersSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type GettersSession struct {
	Contract     *Getters          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GettersCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type GettersCallerSession struct {
	Contract *GettersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GettersTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type GettersTransactorSession struct {
	Contract     *GettersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GettersRaw is an auto generated low-level Go binding around a Klaytn contract.
type GettersRaw struct {
	Contract *Getters // Generic contract binding to access the raw methods on
}

// GettersCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type GettersCallerRaw struct {
	Contract *GettersCaller // Generic read-only contract binding to access the raw methods on
}

// GettersTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type GettersTransactorRaw struct {
	Contract *GettersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGetters creates a new instance of Getters, bound to a specific deployed contract.
func NewGetters(address common.Address, backend bind.ContractBackend) (*Getters, error) {
	contract, err := bindGetters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Getters{GettersCaller: GettersCaller{contract: contract}, GettersTransactor: GettersTransactor{contract: contract}, GettersFilterer: GettersFilterer{contract: contract}}, nil
}

// NewGettersCaller creates a new read-only instance of Getters, bound to a specific deployed contract.
func NewGettersCaller(address common.Address, caller bind.ContractCaller) (*GettersCaller, error) {
	contract, err := bindGetters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GettersCaller{contract: contract}, nil
}

// NewGettersTransactor creates a new write-only instance of Getters, bound to a specific deployed contract.
func NewGettersTransactor(address common.Address, transactor bind.ContractTransactor) (*GettersTransactor, error) {
	contract, err := bindGetters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GettersTransactor{contract: contract}, nil
}

// NewGettersFilterer creates a new log filterer instance of Getters, bound to a specific deployed contract.
func NewGettersFilterer(address common.Address, filterer bind.ContractFilterer) (*GettersFilterer, error) {
	contract, err := bindGetters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GettersFilterer{contract: contract}, nil
}

// bindGetters binds a generic wrapper to an already deployed contract.
func bindGetters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GettersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Getters *GettersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Getters.Contract.GettersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Getters *GettersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Getters.Contract.GettersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Getters *GettersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Getters.Contract.GettersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Getters *GettersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Getters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Getters *GettersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Getters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Getters *GettersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Getters.Contract.contract.Transact(opts, method, params...)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_Getters *GettersCaller) ChainId(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Getters.contract.Call(opts, out, "chainId")
	return *ret0, err
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_Getters *GettersSession) ChainId() (uint16, error) {
	return _Getters.Contract.ChainId(&_Getters.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_Getters *GettersCallerSession) ChainId() (uint16, error) {
	return _Getters.Contract.ChainId(&_Getters.CallOpts)
}

// GetCurrentGuardianSetIndex is a free data retrieval call binding the contract method 0x1cfe7951.
//
// Solidity: function getCurrentGuardianSetIndex() view returns(uint32)
func (_Getters *GettersCaller) GetCurrentGuardianSetIndex(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Getters.contract.Call(opts, out, "getCurrentGuardianSetIndex")
	return *ret0, err
}

// GetCurrentGuardianSetIndex is a free data retrieval call binding the contract method 0x1cfe7951.
//
// Solidity: function getCurrentGuardianSetIndex() view returns(uint32)
func (_Getters *GettersSession) GetCurrentGuardianSetIndex() (uint32, error) {
	return _Getters.Contract.GetCurrentGuardianSetIndex(&_Getters.CallOpts)
}

// GetCurrentGuardianSetIndex is a free data retrieval call binding the contract method 0x1cfe7951.
//
// Solidity: function getCurrentGuardianSetIndex() view returns(uint32)
func (_Getters *GettersCallerSession) GetCurrentGuardianSetIndex() (uint32, error) {
	return _Getters.Contract.GetCurrentGuardianSetIndex(&_Getters.CallOpts)
}

// GetGuardianSet is a free data retrieval call binding the contract method 0xf951975a.
//
// Solidity: function getGuardianSet(uint32 index) view returns((address[],uint32))
func (_Getters *GettersCaller) GetGuardianSet(opts *bind.CallOpts, index uint32) (StructsGuardianSet, error) {
	var (
		ret0 = new(StructsGuardianSet)
	)
	out := ret0
	err := _Getters.contract.Call(opts, out, "getGuardianSet", index)
	return *ret0, err
}

// GetGuardianSet is a free data retrieval call binding the contract method 0xf951975a.
//
// Solidity: function getGuardianSet(uint32 index) view returns((address[],uint32))
func (_Getters *GettersSession) GetGuardianSet(index uint32) (StructsGuardianSet, error) {
	return _Getters.Contract.GetGuardianSet(&_Getters.CallOpts, index)
}

// GetGuardianSet is a free data retrieval call binding the contract method 0xf951975a.
//
// Solidity: function getGuardianSet(uint32 index) view returns((address[],uint32))
func (_Getters *GettersCallerSession) GetGuardianSet(index uint32) (StructsGuardianSet, error) {
	return _Getters.Contract.GetGuardianSet(&_Getters.CallOpts, index)
}

// GetGuardianSetExpiry is a free data retrieval call binding the contract method 0xeb8d3f12.
//
// Solidity: function getGuardianSetExpiry() view returns(uint32)
func (_Getters *GettersCaller) GetGuardianSetExpiry(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Getters.contract.Call(opts, out, "getGuardianSetExpiry")
	return *ret0, err
}

// GetGuardianSetExpiry is a free data retrieval call binding the contract method 0xeb8d3f12.
//
// Solidity: function getGuardianSetExpiry() view returns(uint32)
func (_Getters *GettersSession) GetGuardianSetExpiry() (uint32, error) {
	return _Getters.Contract.GetGuardianSetExpiry(&_Getters.CallOpts)
}

// GetGuardianSetExpiry is a free data retrieval call binding the contract method 0xeb8d3f12.
//
// Solidity: function getGuardianSetExpiry() view returns(uint32)
func (_Getters *GettersCallerSession) GetGuardianSetExpiry() (uint32, error) {
	return _Getters.Contract.GetGuardianSetExpiry(&_Getters.CallOpts)
}

// GovernanceActionIsConsumed is a free data retrieval call binding the contract method 0x2c3c02a4.
//
// Solidity: function governanceActionIsConsumed(bytes32 hash) view returns(bool)
func (_Getters *GettersCaller) GovernanceActionIsConsumed(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Getters.contract.Call(opts, out, "governanceActionIsConsumed", hash)
	return *ret0, err
}

// GovernanceActionIsConsumed is a free data retrieval call binding the contract method 0x2c3c02a4.
//
// Solidity: function governanceActionIsConsumed(bytes32 hash) view returns(bool)
func (_Getters *GettersSession) GovernanceActionIsConsumed(hash [32]byte) (bool, error) {
	return _Getters.Contract.GovernanceActionIsConsumed(&_Getters.CallOpts, hash)
}

// GovernanceActionIsConsumed is a free data retrieval call binding the contract method 0x2c3c02a4.
//
// Solidity: function governanceActionIsConsumed(bytes32 hash) view returns(bool)
func (_Getters *GettersCallerSession) GovernanceActionIsConsumed(hash [32]byte) (bool, error) {
	return _Getters.Contract.GovernanceActionIsConsumed(&_Getters.CallOpts, hash)
}

// GovernanceChainId is a free data retrieval call binding the contract method 0xfbe3c2cd.
//
// Solidity: function governanceChainId() view returns(uint16)
func (_Getters *GettersCaller) GovernanceChainId(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Getters.contract.Call(opts, out, "governanceChainId")
	return *ret0, err
}

// GovernanceChainId is a free data retrieval call binding the contract method 0xfbe3c2cd.
//
// Solidity: function governanceChainId() view returns(uint16)
func (_Getters *GettersSession) GovernanceChainId() (uint16, error) {
	return _Getters.Contract.GovernanceChainId(&_Getters.CallOpts)
}

// GovernanceChainId is a free data retrieval call binding the contract method 0xfbe3c2cd.
//
// Solidity: function governanceChainId() view returns(uint16)
func (_Getters *GettersCallerSession) GovernanceChainId() (uint16, error) {
	return _Getters.Contract.GovernanceChainId(&_Getters.CallOpts)
}

// GovernanceContract is a free data retrieval call binding the contract method 0xb172b222.
//
// Solidity: function governanceContract() view returns(bytes32)
func (_Getters *GettersCaller) GovernanceContract(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Getters.contract.Call(opts, out, "governanceContract")
	return *ret0, err
}

// GovernanceContract is a free data retrieval call binding the contract method 0xb172b222.
//
// Solidity: function governanceContract() view returns(bytes32)
func (_Getters *GettersSession) GovernanceContract() ([32]byte, error) {
	return _Getters.Contract.GovernanceContract(&_Getters.CallOpts)
}

// GovernanceContract is a free data retrieval call binding the contract method 0xb172b222.
//
// Solidity: function governanceContract() view returns(bytes32)
func (_Getters *GettersCallerSession) GovernanceContract() ([32]byte, error) {
	return _Getters.Contract.GovernanceContract(&_Getters.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0xd60b347f.
//
// Solidity: function isInitialized(address impl) view returns(bool)
func (_Getters *GettersCaller) IsInitialized(opts *bind.CallOpts, impl common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Getters.contract.Call(opts, out, "isInitialized", impl)
	return *ret0, err
}

// IsInitialized is a free data retrieval call binding the contract method 0xd60b347f.
//
// Solidity: function isInitialized(address impl) view returns(bool)
func (_Getters *GettersSession) IsInitialized(impl common.Address) (bool, error) {
	return _Getters.Contract.IsInitialized(&_Getters.CallOpts, impl)
}

// IsInitialized is a free data retrieval call binding the contract method 0xd60b347f.
//
// Solidity: function isInitialized(address impl) view returns(bool)
func (_Getters *GettersCallerSession) IsInitialized(impl common.Address) (bool, error) {
	return _Getters.Contract.IsInitialized(&_Getters.CallOpts, impl)
}

// MessageFee is a free data retrieval call binding the contract method 0x1a90a219.
//
// Solidity: function messageFee() view returns(uint256)
func (_Getters *GettersCaller) MessageFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Getters.contract.Call(opts, out, "messageFee")
	return *ret0, err
}

// MessageFee is a free data retrieval call binding the contract method 0x1a90a219.
//
// Solidity: function messageFee() view returns(uint256)
func (_Getters *GettersSession) MessageFee() (*big.Int, error) {
	return _Getters.Contract.MessageFee(&_Getters.CallOpts)
}

// MessageFee is a free data retrieval call binding the contract method 0x1a90a219.
//
// Solidity: function messageFee() view returns(uint256)
func (_Getters *GettersCallerSession) MessageFee() (*big.Int, error) {
	return _Getters.Contract.MessageFee(&_Getters.CallOpts)
}

// NextSequence is a free data retrieval call binding the contract method 0x4cf842b5.
//
// Solidity: function nextSequence(address emitter) view returns(uint64)
func (_Getters *GettersCaller) NextSequence(opts *bind.CallOpts, emitter common.Address) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Getters.contract.Call(opts, out, "nextSequence", emitter)
	return *ret0, err
}

// NextSequence is a free data retrieval call binding the contract method 0x4cf842b5.
//
// Solidity: function nextSequence(address emitter) view returns(uint64)
func (_Getters *GettersSession) NextSequence(emitter common.Address) (uint64, error) {
	return _Getters.Contract.NextSequence(&_Getters.CallOpts, emitter)
}

// NextSequence is a free data retrieval call binding the contract method 0x4cf842b5.
//
// Solidity: function nextSequence(address emitter) view returns(uint64)
func (_Getters *GettersCallerSession) NextSequence(emitter common.Address) (uint64, error) {
	return _Getters.Contract.NextSequence(&_Getters.CallOpts, emitter)
}

// MessagesABI is the input ABI used to generate the binding from.
const MessagesABI = "[{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentGuardianSetIndex\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"getGuardianSet\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"keys\",\"type\":\"address[]\"},{\"internalType\":\"uint32\",\"name\":\"expirationTime\",\"type\":\"uint32\"}],\"internalType\":\"structStructs.GuardianSet\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGuardianSetExpiry\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"governanceActionIsConsumed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceChainId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceContract\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"emitter\",\"type\":\"address\"}],\"name\":\"nextSequence\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedVM\",\"type\":\"bytes\"}],\"name\":\"parseAndVerifyVM\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"emitterChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"emitterAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"consistencyLevel\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"guardianSetIndex\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"guardianIndex\",\"type\":\"uint8\"}],\"internalType\":\"structStructs.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"internalType\":\"structStructs.VM\",\"name\":\"vm\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedVM\",\"type\":\"bytes\"}],\"name\":\"parseVM\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"emitterChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"emitterAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"consistencyLevel\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"guardianSetIndex\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"guardianIndex\",\"type\":\"uint8\"}],\"internalType\":\"structStructs.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"internalType\":\"structStructs.VM\",\"name\":\"vm\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"guardianIndex\",\"type\":\"uint8\"}],\"internalType\":\"structStructs.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"keys\",\"type\":\"address[]\"},{\"internalType\":\"uint32\",\"name\":\"expirationTime\",\"type\":\"uint32\"}],\"internalType\":\"structStructs.GuardianSet\",\"name\":\"guardianSet\",\"type\":\"tuple\"}],\"name\":\"verifySignatures\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"emitterChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"emitterAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"consistencyLevel\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"guardianSetIndex\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"guardianIndex\",\"type\":\"uint8\"}],\"internalType\":\"structStructs.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"internalType\":\"structStructs.VM\",\"name\":\"vm\",\"type\":\"tuple\"}],\"name\":\"verifyVM\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MessagesBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const MessagesBinRuntime = `608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063a9e118931161008c578063d60b347f11610066578063d60b347f1461023b578063eb8d3f1214610267578063f951975a1461027d578063fbe3c2cd1461029d57600080fd5b8063a9e11893146101f1578063b172b22214610211578063c0fd8bde1461021957600080fd5b80634cf842b5116100c85780634cf842b514610158578063875be02a146101a25780639a8a0592146101c3578063a0cce1b3146101de57600080fd5b80631a90a219146100ef5780631cfe7951146101065780632c3c02a414610125575b600080fd5b6007545b6040519081526020015b60405180910390f35b60035463ffffffff165b60405163ffffffff90911681526020016100fd565b610148610133366004610ddb565b60009081526005602052604090205460ff1690565b60405190151581526020016100fd565b61018a610166366004610e10565b6001600160a01b03166000908152600460205260409020546001600160401b031690565b6040516001600160401b0390911681526020016100fd565b6101b56101b0366004611076565b6102af565b6040516100fd9291906111e7565b60005461ffff165b60405161ffff90911681526020016100fd565b6101b56101ec36600461120a565b610443565b6102046101ff366004611312565b610609565b6040516100fd919061149a565b6001546100f3565b61022c6102273660046114ad565b6109a1565b6040516100fd9392919061151e565b610148610249366004610e10565b6001600160a01b031660009081526006602052604090205460ff1690565b600354640100000000900463ffffffff16610110565b61029061028b366004611555565b610a03565b6040516100fd9190611570565b60005462010000900461ffff166101cb565b6000606060006102c3846101000151610a03565b805151909150610306576000604051806040016040528060148152602001731a5b9d985b1a590819dd585c991a585b881cd95d60621b8152509250925050915091565b60035463ffffffff1663ffffffff1684610100015163ffffffff1614158015610338575042816020015163ffffffff16105b1561037f5760006040518060400160405280601881526020017f677561726469616e2073657420686173206578706972656400000000000000008152509250925050915091565b61012084015151815151600a9060039061039990836115f1565b6103a39190611610565b6103ae9060026115f1565b6103b89190611610565b6103c3906001611632565b11156103f7576000604051806040016040528060098152602001686e6f2071756f72756d60b81b8152509250925050915091565b60008061040f86610140015187610120015185610443565b9150915081610425576000969095509350505050565b60016040518060200160405280600081525094509450505050915091565b600060606000805b85518110156105e85760008682815181106104685761046861164a565b60200260200101519050816000148061048a57508260ff16816060015160ff16115b6104e75760405162461bcd60e51b815260206004820152602360248201527f7369676e617475726520696e6469636573206d75737420626520617363656e64604482015262696e6760e81b60648201526084015b60405180910390fd5b6060810151865180519194509060ff85169081106105075761050761164a565b60200260200101516001600160a01b031660018983604001518460000151856020015160405160008152602001604052604051610560949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015610582573d6000803e3d6000fd5b505050602060405103516001600160a01b0316146105d5576000604051806040016040528060148152602001731593481cda59db985d1d5c99481a5b9d985b1a5960621b81525094509450505050610601565b50806105e081611660565b91505061044b565b5060016040518060200160405280600081525092509250505b935093915050565b610611610d80565b600061061d8382610aa2565b60ff16825261062d600182611632565b9050816000015160ff166001146106865760405162461bcd60e51b815260206004820152601760248201527f564d2076657273696f6e20696e636f6d70617469626c6500000000000000000060448201526064016104de565b6106908382610afe565b63ffffffff166101008301526106a7600482611632565b905060006106b58483610aa2565b60ff1690506106c5600183611632565b9150806001600160401b038111156106df576106df610e32565b60405190808252806020026020018201604052801561073157816020015b6040805160808101825260008082526020808301829052928201819052606082015282526000199092019101816106fd5790505b5061012084015260005b818110156108685761074d8584610aa2565b84610120015182815181106107645761076461164a565b602090810291909101015160ff909116606090910152610785600184611632565b92506107918584610b5b565b84610120015182815181106107a8576107a861164a565b602002602001015160000181815250506020836107c59190611632565b92506107d18584610b5b565b84610120015182815181106107e8576107e861164a565b602002602001015160200181815250506020836108059190611632565b92506108118584610aa2565b61081c90601b61167b565b84610120015182815181106108335761083361164a565b602090810291909101015160ff909116604090910152610854600184611632565b92508061086081611660565b91505061073b565b5060006108848384875161087c91906116a0565b879190610bb9565b905080805190602001206040516020016108a091815260200190565b60408051601f1981840301815291905280516020909101206101408501526108c88584610afe565b63ffffffff1660208501526108de600484611632565b92506108ea8584610afe565b63ffffffff166040850152610900600484611632565b925061090c8584610cc6565b61ffff166060850152610920600284611632565b925061092c8584610b5b565b608085015261093c602084611632565b92506109488584610d23565b6001600160401b031660a0850152610961600884611632565b925061096d8584610aa2565b60ff1660c0850152610980600184611632565b92506109938384875161087c91906116a0565b60e085015250919392505050565b6109a9610d80565b600060606109ec85858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061060992505050565b92506109f7836102af565b93969095509293505050565b60408051808201825260608082526000602080840182905263ffffffff86168252600281529084902084518154928302810184018652948501828152939493909284928491840182828015610a8157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610a63575b50505091835250506001919091015463ffffffff1660209091015292915050565b6000610aaf826001611632565b83511015610af55760405162461bcd60e51b8152602060048201526013602482015272746f55696e74385f6f75744f66426f756e647360681b60448201526064016104de565b50016001015190565b6000610b0b826004611632565b83511015610b525760405162461bcd60e51b8152602060048201526014602482015273746f55696e7433325f6f75744f66426f756e647360601b60448201526064016104de565b50016004015190565b6000610b68826020611632565b83511015610bb05760405162461bcd60e51b8152602060048201526015602482015274746f427974657333325f6f75744f66426f756e647360581b60448201526064016104de565b50016020015190565b606081610bc781601f611632565b1015610c065760405162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b60448201526064016104de565b610c108284611632565b84511015610c545760405162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b60448201526064016104de565b606082158015610c735760405191506000825260208201604052610cbd565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015610cac578051835260209283019201610c94565b5050858452601f01601f1916604052505b50949350505050565b6000610cd3826002611632565b83511015610d1a5760405162461bcd60e51b8152602060048201526014602482015273746f55696e7431365f6f75744f66426f756e647360601b60448201526064016104de565b50016002015190565b6000610d30826008611632565b83511015610d775760405162461bcd60e51b8152602060048201526014602482015273746f55696e7436345f6f75744f66426f756e647360601b60448201526064016104de565b50016008015190565b604080516101608101825260008082526020820181905291810182905260608082018390526080820183905260a0820183905260c0820183905260e08201819052610100820183905261012082015261014081019190915290565b600060208284031215610ded57600080fd5b5035919050565b80356001600160a01b0381168114610e0b57600080fd5b919050565b600060208284031215610e2257600080fd5b610e2b82610df4565b9392505050565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b0381118282101715610e6a57610e6a610e32565b60405290565b60405161016081016001600160401b0381118282101715610e6a57610e6a610e32565b604080519081016001600160401b0381118282101715610e6a57610e6a610e32565b604051601f8201601f191681016001600160401b0381118282101715610edd57610edd610e32565b604052919050565b803560ff81168114610e0b57600080fd5b803563ffffffff81168114610e0b57600080fd5b803561ffff81168114610e0b57600080fd5b80356001600160401b0381168114610e0b57600080fd5b600082601f830112610f4457600080fd5b81356001600160401b03811115610f5d57610f5d610e32565b610f70601f8201601f1916602001610eb5565b818152846020838601011115610f8557600080fd5b816020850160208301376000918101602001919091529392505050565b60006001600160401b03821115610fbb57610fbb610e32565b5060051b60200190565b600082601f830112610fd657600080fd5b81356020610feb610fe683610fa2565b610eb5565b82815260079290921b8401810191818101908684111561100a57600080fd5b8286015b8481101561106b57608081890312156110275760008081fd5b61102f610e48565b8135815284820135858201526040611048818401610ee5565b908201526060611059838201610ee5565b9082015283529183019160800161100e565b509695505050505050565b60006020828403121561108857600080fd5b81356001600160401b038082111561109f57600080fd5b9083019061016082860312156110b457600080fd5b6110bc610e70565b6110c583610ee5565b81526110d360208401610ef6565b60208201526110e460408401610ef6565b60408201526110f560608401610f0a565b60608201526080830135608082015261111060a08401610f1c565b60a082015261112160c08401610ee5565b60c082015260e08301358281111561113857600080fd5b61114487828601610f33565b60e083015250610100611158818501610ef6565b90820152610120838101358381111561117057600080fd5b61117c88828701610fc5565b91830191909152506101409283013592810192909252509392505050565b6000815180845260005b818110156111c0576020818501810151868301820152016111a4565b818111156111d2576000602083870101525b50601f01601f19169290920160200192915050565b8215158152604060208201526000611202604083018461119a565b949350505050565b60008060006060848603121561121f57600080fd5b833592506020808501356001600160401b038082111561123e57600080fd5b61124a88838901610fc5565b9450604087013591508082111561126057600080fd5b908601906040828903121561127457600080fd5b61127c610e93565b82358281111561128b57600080fd5b83019150601f8201891361129e57600080fd5b81356112ac610fe682610fa2565b81815260059190911b8301850190858101908b8311156112cb57600080fd5b938601935b828510156112f0576112e185610df4565b825293860193908601906112d0565b8352506113009050838501610ef6565b84820152809450505050509250925092565b60006020828403121561132457600080fd5b81356001600160401b0381111561133a57600080fd5b61120284828501610f33565b600081518084526020808501945080840160005b8381101561139f57815180518852838101518489015260408082015160ff908116918a019190915260609182015116908801526080909601959082019060010161135a565b509495945050505050565b805160ff168252600061016060208301516113cd602086018263ffffffff169052565b5060408301516113e5604086018263ffffffff169052565b5060608301516113fb606086018261ffff169052565b506080830151608085015260a083015161142060a08601826001600160401b03169052565b5060c083015161143560c086018260ff169052565b5060e08301518160e086015261144d8286018261119a565b915050610100808401516114688287018263ffffffff169052565b505061012080840151858303828701526114828382611346565b61014095860151969095019590955250919392505050565b602081526000610e2b60208301846113aa565b600080602083850312156114c057600080fd5b82356001600160401b03808211156114d757600080fd5b818501915085601f8301126114eb57600080fd5b8135818111156114fa57600080fd5b86602082850101111561150c57600080fd5b60209290920196919550909350505050565b60608152600061153160608301866113aa565b8415156020840152828103604084015261154b818561119a565b9695505050505050565b60006020828403121561156757600080fd5b610e2b82610ef6565b6020808252825160408383015280516060840181905260009291820190839060808601905b808310156115be5783516001600160a01b03168252928401926001929092019190840190611595565b5063ffffffff848801511660408701528094505050505092915050565b634e487b7160e01b600052601160045260246000fd5b600081600019048311821515161561160b5761160b6115db565b500290565b60008261162d57634e487b7160e01b600052601260045260246000fd5b500490565b60008219821115611645576116456115db565b500190565b634e487b7160e01b600052603260045260246000fd5b6000600019821415611674576116746115db565b5060010190565b600060ff821660ff84168060ff03821115611698576116986115db565b019392505050565b6000828210156116b2576116b26115db565b50039056fea2646970667358221220cef4d6830fcc39dfbfdd1164532ba1552c048f762e89b857b4666014a9057ad064736f6c637828302e382e31322d646576656c6f702e323032322e312e31322b636f6d6d69742e61373131393639390059`

// MessagesFuncSigs maps the 4-byte function signature to its string representation.
var MessagesFuncSigs = map[string]string{
	"9a8a0592": "chainId()",
	"1cfe7951": "getCurrentGuardianSetIndex()",
	"f951975a": "getGuardianSet(uint32)",
	"eb8d3f12": "getGuardianSetExpiry()",
	"2c3c02a4": "governanceActionIsConsumed(bytes32)",
	"fbe3c2cd": "governanceChainId()",
	"b172b222": "governanceContract()",
	"d60b347f": "isInitialized(address)",
	"1a90a219": "messageFee()",
	"4cf842b5": "nextSequence(address)",
	"c0fd8bde": "parseAndVerifyVM(bytes)",
	"a9e11893": "parseVM(bytes)",
	"a0cce1b3": "verifySignatures(bytes32,(bytes32,bytes32,uint8,uint8)[],(address[],uint32))",
	"875be02a": "verifyVM((uint8,uint32,uint32,uint16,bytes32,uint64,uint8,bytes,uint32,(bytes32,bytes32,uint8,uint8)[],bytes32))",
}

// MessagesBin is the compiled bytecode used for deploying new contracts.
var MessagesBin = "0x608060405234801561001057600080fd5b50611713806100206000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063a9e118931161008c578063d60b347f11610066578063d60b347f1461023b578063eb8d3f1214610267578063f951975a1461027d578063fbe3c2cd1461029d57600080fd5b8063a9e11893146101f1578063b172b22214610211578063c0fd8bde1461021957600080fd5b80634cf842b5116100c85780634cf842b514610158578063875be02a146101a25780639a8a0592146101c3578063a0cce1b3146101de57600080fd5b80631a90a219146100ef5780631cfe7951146101065780632c3c02a414610125575b600080fd5b6007545b6040519081526020015b60405180910390f35b60035463ffffffff165b60405163ffffffff90911681526020016100fd565b610148610133366004610ddb565b60009081526005602052604090205460ff1690565b60405190151581526020016100fd565b61018a610166366004610e10565b6001600160a01b03166000908152600460205260409020546001600160401b031690565b6040516001600160401b0390911681526020016100fd565b6101b56101b0366004611076565b6102af565b6040516100fd9291906111e7565b60005461ffff165b60405161ffff90911681526020016100fd565b6101b56101ec36600461120a565b610443565b6102046101ff366004611312565b610609565b6040516100fd919061149a565b6001546100f3565b61022c6102273660046114ad565b6109a1565b6040516100fd9392919061151e565b610148610249366004610e10565b6001600160a01b031660009081526006602052604090205460ff1690565b600354640100000000900463ffffffff16610110565b61029061028b366004611555565b610a03565b6040516100fd9190611570565b60005462010000900461ffff166101cb565b6000606060006102c3846101000151610a03565b805151909150610306576000604051806040016040528060148152602001731a5b9d985b1a590819dd585c991a585b881cd95d60621b8152509250925050915091565b60035463ffffffff1663ffffffff1684610100015163ffffffff1614158015610338575042816020015163ffffffff16105b1561037f5760006040518060400160405280601881526020017f677561726469616e2073657420686173206578706972656400000000000000008152509250925050915091565b61012084015151815151600a9060039061039990836115f1565b6103a39190611610565b6103ae9060026115f1565b6103b89190611610565b6103c3906001611632565b11156103f7576000604051806040016040528060098152602001686e6f2071756f72756d60b81b8152509250925050915091565b60008061040f86610140015187610120015185610443565b9150915081610425576000969095509350505050565b60016040518060200160405280600081525094509450505050915091565b600060606000805b85518110156105e85760008682815181106104685761046861164a565b60200260200101519050816000148061048a57508260ff16816060015160ff16115b6104e75760405162461bcd60e51b815260206004820152602360248201527f7369676e617475726520696e6469636573206d75737420626520617363656e64604482015262696e6760e81b60648201526084015b60405180910390fd5b6060810151865180519194509060ff85169081106105075761050761164a565b60200260200101516001600160a01b031660018983604001518460000151856020015160405160008152602001604052604051610560949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015610582573d6000803e3d6000fd5b505050602060405103516001600160a01b0316146105d5576000604051806040016040528060148152602001731593481cda59db985d1d5c99481a5b9d985b1a5960621b81525094509450505050610601565b50806105e081611660565b91505061044b565b5060016040518060200160405280600081525092509250505b935093915050565b610611610d80565b600061061d8382610aa2565b60ff16825261062d600182611632565b9050816000015160ff166001146106865760405162461bcd60e51b815260206004820152601760248201527f564d2076657273696f6e20696e636f6d70617469626c6500000000000000000060448201526064016104de565b6106908382610afe565b63ffffffff166101008301526106a7600482611632565b905060006106b58483610aa2565b60ff1690506106c5600183611632565b9150806001600160401b038111156106df576106df610e32565b60405190808252806020026020018201604052801561073157816020015b6040805160808101825260008082526020808301829052928201819052606082015282526000199092019101816106fd5790505b5061012084015260005b818110156108685761074d8584610aa2565b84610120015182815181106107645761076461164a565b602090810291909101015160ff909116606090910152610785600184611632565b92506107918584610b5b565b84610120015182815181106107a8576107a861164a565b602002602001015160000181815250506020836107c59190611632565b92506107d18584610b5b565b84610120015182815181106107e8576107e861164a565b602002602001015160200181815250506020836108059190611632565b92506108118584610aa2565b61081c90601b61167b565b84610120015182815181106108335761083361164a565b602090810291909101015160ff909116604090910152610854600184611632565b92508061086081611660565b91505061073b565b5060006108848384875161087c91906116a0565b879190610bb9565b905080805190602001206040516020016108a091815260200190565b60408051601f1981840301815291905280516020909101206101408501526108c88584610afe565b63ffffffff1660208501526108de600484611632565b92506108ea8584610afe565b63ffffffff166040850152610900600484611632565b925061090c8584610cc6565b61ffff166060850152610920600284611632565b925061092c8584610b5b565b608085015261093c602084611632565b92506109488584610d23565b6001600160401b031660a0850152610961600884611632565b925061096d8584610aa2565b60ff1660c0850152610980600184611632565b92506109938384875161087c91906116a0565b60e085015250919392505050565b6109a9610d80565b600060606109ec85858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061060992505050565b92506109f7836102af565b93969095509293505050565b60408051808201825260608082526000602080840182905263ffffffff86168252600281529084902084518154928302810184018652948501828152939493909284928491840182828015610a8157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610a63575b50505091835250506001919091015463ffffffff1660209091015292915050565b6000610aaf826001611632565b83511015610af55760405162461bcd60e51b8152602060048201526013602482015272746f55696e74385f6f75744f66426f756e647360681b60448201526064016104de565b50016001015190565b6000610b0b826004611632565b83511015610b525760405162461bcd60e51b8152602060048201526014602482015273746f55696e7433325f6f75744f66426f756e647360601b60448201526064016104de565b50016004015190565b6000610b68826020611632565b83511015610bb05760405162461bcd60e51b8152602060048201526015602482015274746f427974657333325f6f75744f66426f756e647360581b60448201526064016104de565b50016020015190565b606081610bc781601f611632565b1015610c065760405162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b60448201526064016104de565b610c108284611632565b84511015610c545760405162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b60448201526064016104de565b606082158015610c735760405191506000825260208201604052610cbd565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015610cac578051835260209283019201610c94565b5050858452601f01601f1916604052505b50949350505050565b6000610cd3826002611632565b83511015610d1a5760405162461bcd60e51b8152602060048201526014602482015273746f55696e7431365f6f75744f66426f756e647360601b60448201526064016104de565b50016002015190565b6000610d30826008611632565b83511015610d775760405162461bcd60e51b8152602060048201526014602482015273746f55696e7436345f6f75744f66426f756e647360601b60448201526064016104de565b50016008015190565b604080516101608101825260008082526020820181905291810182905260608082018390526080820183905260a0820183905260c0820183905260e08201819052610100820183905261012082015261014081019190915290565b600060208284031215610ded57600080fd5b5035919050565b80356001600160a01b0381168114610e0b57600080fd5b919050565b600060208284031215610e2257600080fd5b610e2b82610df4565b9392505050565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b0381118282101715610e6a57610e6a610e32565b60405290565b60405161016081016001600160401b0381118282101715610e6a57610e6a610e32565b604080519081016001600160401b0381118282101715610e6a57610e6a610e32565b604051601f8201601f191681016001600160401b0381118282101715610edd57610edd610e32565b604052919050565b803560ff81168114610e0b57600080fd5b803563ffffffff81168114610e0b57600080fd5b803561ffff81168114610e0b57600080fd5b80356001600160401b0381168114610e0b57600080fd5b600082601f830112610f4457600080fd5b81356001600160401b03811115610f5d57610f5d610e32565b610f70601f8201601f1916602001610eb5565b818152846020838601011115610f8557600080fd5b816020850160208301376000918101602001919091529392505050565b60006001600160401b03821115610fbb57610fbb610e32565b5060051b60200190565b600082601f830112610fd657600080fd5b81356020610feb610fe683610fa2565b610eb5565b82815260079290921b8401810191818101908684111561100a57600080fd5b8286015b8481101561106b57608081890312156110275760008081fd5b61102f610e48565b8135815284820135858201526040611048818401610ee5565b908201526060611059838201610ee5565b9082015283529183019160800161100e565b509695505050505050565b60006020828403121561108857600080fd5b81356001600160401b038082111561109f57600080fd5b9083019061016082860312156110b457600080fd5b6110bc610e70565b6110c583610ee5565b81526110d360208401610ef6565b60208201526110e460408401610ef6565b60408201526110f560608401610f0a565b60608201526080830135608082015261111060a08401610f1c565b60a082015261112160c08401610ee5565b60c082015260e08301358281111561113857600080fd5b61114487828601610f33565b60e083015250610100611158818501610ef6565b90820152610120838101358381111561117057600080fd5b61117c88828701610fc5565b91830191909152506101409283013592810192909252509392505050565b6000815180845260005b818110156111c0576020818501810151868301820152016111a4565b818111156111d2576000602083870101525b50601f01601f19169290920160200192915050565b8215158152604060208201526000611202604083018461119a565b949350505050565b60008060006060848603121561121f57600080fd5b833592506020808501356001600160401b038082111561123e57600080fd5b61124a88838901610fc5565b9450604087013591508082111561126057600080fd5b908601906040828903121561127457600080fd5b61127c610e93565b82358281111561128b57600080fd5b83019150601f8201891361129e57600080fd5b81356112ac610fe682610fa2565b81815260059190911b8301850190858101908b8311156112cb57600080fd5b938601935b828510156112f0576112e185610df4565b825293860193908601906112d0565b8352506113009050838501610ef6565b84820152809450505050509250925092565b60006020828403121561132457600080fd5b81356001600160401b0381111561133a57600080fd5b61120284828501610f33565b600081518084526020808501945080840160005b8381101561139f57815180518852838101518489015260408082015160ff908116918a019190915260609182015116908801526080909601959082019060010161135a565b509495945050505050565b805160ff168252600061016060208301516113cd602086018263ffffffff169052565b5060408301516113e5604086018263ffffffff169052565b5060608301516113fb606086018261ffff169052565b506080830151608085015260a083015161142060a08601826001600160401b03169052565b5060c083015161143560c086018260ff169052565b5060e08301518160e086015261144d8286018261119a565b915050610100808401516114688287018263ffffffff169052565b505061012080840151858303828701526114828382611346565b61014095860151969095019590955250919392505050565b602081526000610e2b60208301846113aa565b600080602083850312156114c057600080fd5b82356001600160401b03808211156114d757600080fd5b818501915085601f8301126114eb57600080fd5b8135818111156114fa57600080fd5b86602082850101111561150c57600080fd5b60209290920196919550909350505050565b60608152600061153160608301866113aa565b8415156020840152828103604084015261154b818561119a565b9695505050505050565b60006020828403121561156757600080fd5b610e2b82610ef6565b6020808252825160408383015280516060840181905260009291820190839060808601905b808310156115be5783516001600160a01b03168252928401926001929092019190840190611595565b5063ffffffff848801511660408701528094505050505092915050565b634e487b7160e01b600052601160045260246000fd5b600081600019048311821515161561160b5761160b6115db565b500290565b60008261162d57634e487b7160e01b600052601260045260246000fd5b500490565b60008219821115611645576116456115db565b500190565b634e487b7160e01b600052603260045260246000fd5b6000600019821415611674576116746115db565b5060010190565b600060ff821660ff84168060ff03821115611698576116986115db565b019392505050565b6000828210156116b2576116b26115db565b50039056fea2646970667358221220cef4d6830fcc39dfbfdd1164532ba1552c048f762e89b857b4666014a9057ad064736f6c637828302e382e31322d646576656c6f702e323032322e312e31322b636f6d6d69742e61373131393639390059"

// DeployMessages deploys a new Klaytn contract, binding an instance of Messages to it.
func DeployMessages(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Messages, error) {
	parsed, err := abi.JSON(strings.NewReader(MessagesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MessagesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Messages{MessagesCaller: MessagesCaller{contract: contract}, MessagesTransactor: MessagesTransactor{contract: contract}, MessagesFilterer: MessagesFilterer{contract: contract}}, nil
}

// Messages is an auto generated Go binding around a Klaytn contract.
type Messages struct {
	MessagesCaller     // Read-only binding to the contract
	MessagesTransactor // Write-only binding to the contract
	MessagesFilterer   // Log filterer for contract events
}

// MessagesCaller is an auto generated read-only Go binding around a Klaytn contract.
type MessagesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesTransactor is an auto generated write-only Go binding around a Klaytn contract.
type MessagesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type MessagesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type MessagesSession struct {
	Contract     *Messages         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessagesCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type MessagesCallerSession struct {
	Contract *MessagesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MessagesTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type MessagesTransactorSession struct {
	Contract     *MessagesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MessagesRaw is an auto generated low-level Go binding around a Klaytn contract.
type MessagesRaw struct {
	Contract *Messages // Generic contract binding to access the raw methods on
}

// MessagesCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type MessagesCallerRaw struct {
	Contract *MessagesCaller // Generic read-only contract binding to access the raw methods on
}

// MessagesTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type MessagesTransactorRaw struct {
	Contract *MessagesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessages creates a new instance of Messages, bound to a specific deployed contract.
func NewMessages(address common.Address, backend bind.ContractBackend) (*Messages, error) {
	contract, err := bindMessages(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Messages{MessagesCaller: MessagesCaller{contract: contract}, MessagesTransactor: MessagesTransactor{contract: contract}, MessagesFilterer: MessagesFilterer{contract: contract}}, nil
}

// NewMessagesCaller creates a new read-only instance of Messages, bound to a specific deployed contract.
func NewMessagesCaller(address common.Address, caller bind.ContractCaller) (*MessagesCaller, error) {
	contract, err := bindMessages(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessagesCaller{contract: contract}, nil
}

// NewMessagesTransactor creates a new write-only instance of Messages, bound to a specific deployed contract.
func NewMessagesTransactor(address common.Address, transactor bind.ContractTransactor) (*MessagesTransactor, error) {
	contract, err := bindMessages(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessagesTransactor{contract: contract}, nil
}

// NewMessagesFilterer creates a new log filterer instance of Messages, bound to a specific deployed contract.
func NewMessagesFilterer(address common.Address, filterer bind.ContractFilterer) (*MessagesFilterer, error) {
	contract, err := bindMessages(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessagesFilterer{contract: contract}, nil
}

// bindMessages binds a generic wrapper to an already deployed contract.
func bindMessages(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessagesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Messages *MessagesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Messages.Contract.MessagesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Messages *MessagesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Messages.Contract.MessagesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Messages *MessagesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Messages.Contract.MessagesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Messages *MessagesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Messages.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Messages *MessagesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Messages.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Messages *MessagesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Messages.Contract.contract.Transact(opts, method, params...)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_Messages *MessagesCaller) ChainId(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Messages.contract.Call(opts, out, "chainId")
	return *ret0, err
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_Messages *MessagesSession) ChainId() (uint16, error) {
	return _Messages.Contract.ChainId(&_Messages.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_Messages *MessagesCallerSession) ChainId() (uint16, error) {
	return _Messages.Contract.ChainId(&_Messages.CallOpts)
}

// GetCurrentGuardianSetIndex is a free data retrieval call binding the contract method 0x1cfe7951.
//
// Solidity: function getCurrentGuardianSetIndex() view returns(uint32)
func (_Messages *MessagesCaller) GetCurrentGuardianSetIndex(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Messages.contract.Call(opts, out, "getCurrentGuardianSetIndex")
	return *ret0, err
}

// GetCurrentGuardianSetIndex is a free data retrieval call binding the contract method 0x1cfe7951.
//
// Solidity: function getCurrentGuardianSetIndex() view returns(uint32)
func (_Messages *MessagesSession) GetCurrentGuardianSetIndex() (uint32, error) {
	return _Messages.Contract.GetCurrentGuardianSetIndex(&_Messages.CallOpts)
}

// GetCurrentGuardianSetIndex is a free data retrieval call binding the contract method 0x1cfe7951.
//
// Solidity: function getCurrentGuardianSetIndex() view returns(uint32)
func (_Messages *MessagesCallerSession) GetCurrentGuardianSetIndex() (uint32, error) {
	return _Messages.Contract.GetCurrentGuardianSetIndex(&_Messages.CallOpts)
}

// GetGuardianSet is a free data retrieval call binding the contract method 0xf951975a.
//
// Solidity: function getGuardianSet(uint32 index) view returns((address[],uint32))
func (_Messages *MessagesCaller) GetGuardianSet(opts *bind.CallOpts, index uint32) (StructsGuardianSet, error) {
	var (
		ret0 = new(StructsGuardianSet)
	)
	out := ret0
	err := _Messages.contract.Call(opts, out, "getGuardianSet", index)
	return *ret0, err
}

// GetGuardianSet is a free data retrieval call binding the contract method 0xf951975a.
//
// Solidity: function getGuardianSet(uint32 index) view returns((address[],uint32))
func (_Messages *MessagesSession) GetGuardianSet(index uint32) (StructsGuardianSet, error) {
	return _Messages.Contract.GetGuardianSet(&_Messages.CallOpts, index)
}

// GetGuardianSet is a free data retrieval call binding the contract method 0xf951975a.
//
// Solidity: function getGuardianSet(uint32 index) view returns((address[],uint32))
func (_Messages *MessagesCallerSession) GetGuardianSet(index uint32) (StructsGuardianSet, error) {
	return _Messages.Contract.GetGuardianSet(&_Messages.CallOpts, index)
}

// GetGuardianSetExpiry is a free data retrieval call binding the contract method 0xeb8d3f12.
//
// Solidity: function getGuardianSetExpiry() view returns(uint32)
func (_Messages *MessagesCaller) GetGuardianSetExpiry(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Messages.contract.Call(opts, out, "getGuardianSetExpiry")
	return *ret0, err
}

// GetGuardianSetExpiry is a free data retrieval call binding the contract method 0xeb8d3f12.
//
// Solidity: function getGuardianSetExpiry() view returns(uint32)
func (_Messages *MessagesSession) GetGuardianSetExpiry() (uint32, error) {
	return _Messages.Contract.GetGuardianSetExpiry(&_Messages.CallOpts)
}

// GetGuardianSetExpiry is a free data retrieval call binding the contract method 0xeb8d3f12.
//
// Solidity: function getGuardianSetExpiry() view returns(uint32)
func (_Messages *MessagesCallerSession) GetGuardianSetExpiry() (uint32, error) {
	return _Messages.Contract.GetGuardianSetExpiry(&_Messages.CallOpts)
}

// GovernanceActionIsConsumed is a free data retrieval call binding the contract method 0x2c3c02a4.
//
// Solidity: function governanceActionIsConsumed(bytes32 hash) view returns(bool)
func (_Messages *MessagesCaller) GovernanceActionIsConsumed(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Messages.contract.Call(opts, out, "governanceActionIsConsumed", hash)
	return *ret0, err
}

// GovernanceActionIsConsumed is a free data retrieval call binding the contract method 0x2c3c02a4.
//
// Solidity: function governanceActionIsConsumed(bytes32 hash) view returns(bool)
func (_Messages *MessagesSession) GovernanceActionIsConsumed(hash [32]byte) (bool, error) {
	return _Messages.Contract.GovernanceActionIsConsumed(&_Messages.CallOpts, hash)
}

// GovernanceActionIsConsumed is a free data retrieval call binding the contract method 0x2c3c02a4.
//
// Solidity: function governanceActionIsConsumed(bytes32 hash) view returns(bool)
func (_Messages *MessagesCallerSession) GovernanceActionIsConsumed(hash [32]byte) (bool, error) {
	return _Messages.Contract.GovernanceActionIsConsumed(&_Messages.CallOpts, hash)
}

// GovernanceChainId is a free data retrieval call binding the contract method 0xfbe3c2cd.
//
// Solidity: function governanceChainId() view returns(uint16)
func (_Messages *MessagesCaller) GovernanceChainId(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Messages.contract.Call(opts, out, "governanceChainId")
	return *ret0, err
}

// GovernanceChainId is a free data retrieval call binding the contract method 0xfbe3c2cd.
//
// Solidity: function governanceChainId() view returns(uint16)
func (_Messages *MessagesSession) GovernanceChainId() (uint16, error) {
	return _Messages.Contract.GovernanceChainId(&_Messages.CallOpts)
}

// GovernanceChainId is a free data retrieval call binding the contract method 0xfbe3c2cd.
//
// Solidity: function governanceChainId() view returns(uint16)
func (_Messages *MessagesCallerSession) GovernanceChainId() (uint16, error) {
	return _Messages.Contract.GovernanceChainId(&_Messages.CallOpts)
}

// GovernanceContract is a free data retrieval call binding the contract method 0xb172b222.
//
// Solidity: function governanceContract() view returns(bytes32)
func (_Messages *MessagesCaller) GovernanceContract(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Messages.contract.Call(opts, out, "governanceContract")
	return *ret0, err
}

// GovernanceContract is a free data retrieval call binding the contract method 0xb172b222.
//
// Solidity: function governanceContract() view returns(bytes32)
func (_Messages *MessagesSession) GovernanceContract() ([32]byte, error) {
	return _Messages.Contract.GovernanceContract(&_Messages.CallOpts)
}

// GovernanceContract is a free data retrieval call binding the contract method 0xb172b222.
//
// Solidity: function governanceContract() view returns(bytes32)
func (_Messages *MessagesCallerSession) GovernanceContract() ([32]byte, error) {
	return _Messages.Contract.GovernanceContract(&_Messages.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0xd60b347f.
//
// Solidity: function isInitialized(address impl) view returns(bool)
func (_Messages *MessagesCaller) IsInitialized(opts *bind.CallOpts, impl common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Messages.contract.Call(opts, out, "isInitialized", impl)
	return *ret0, err
}

// IsInitialized is a free data retrieval call binding the contract method 0xd60b347f.
//
// Solidity: function isInitialized(address impl) view returns(bool)
func (_Messages *MessagesSession) IsInitialized(impl common.Address) (bool, error) {
	return _Messages.Contract.IsInitialized(&_Messages.CallOpts, impl)
}

// IsInitialized is a free data retrieval call binding the contract method 0xd60b347f.
//
// Solidity: function isInitialized(address impl) view returns(bool)
func (_Messages *MessagesCallerSession) IsInitialized(impl common.Address) (bool, error) {
	return _Messages.Contract.IsInitialized(&_Messages.CallOpts, impl)
}

// MessageFee is a free data retrieval call binding the contract method 0x1a90a219.
//
// Solidity: function messageFee() view returns(uint256)
func (_Messages *MessagesCaller) MessageFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Messages.contract.Call(opts, out, "messageFee")
	return *ret0, err
}

// MessageFee is a free data retrieval call binding the contract method 0x1a90a219.
//
// Solidity: function messageFee() view returns(uint256)
func (_Messages *MessagesSession) MessageFee() (*big.Int, error) {
	return _Messages.Contract.MessageFee(&_Messages.CallOpts)
}

// MessageFee is a free data retrieval call binding the contract method 0x1a90a219.
//
// Solidity: function messageFee() view returns(uint256)
func (_Messages *MessagesCallerSession) MessageFee() (*big.Int, error) {
	return _Messages.Contract.MessageFee(&_Messages.CallOpts)
}

// NextSequence is a free data retrieval call binding the contract method 0x4cf842b5.
//
// Solidity: function nextSequence(address emitter) view returns(uint64)
func (_Messages *MessagesCaller) NextSequence(opts *bind.CallOpts, emitter common.Address) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Messages.contract.Call(opts, out, "nextSequence", emitter)
	return *ret0, err
}

// NextSequence is a free data retrieval call binding the contract method 0x4cf842b5.
//
// Solidity: function nextSequence(address emitter) view returns(uint64)
func (_Messages *MessagesSession) NextSequence(emitter common.Address) (uint64, error) {
	return _Messages.Contract.NextSequence(&_Messages.CallOpts, emitter)
}

// NextSequence is a free data retrieval call binding the contract method 0x4cf842b5.
//
// Solidity: function nextSequence(address emitter) view returns(uint64)
func (_Messages *MessagesCallerSession) NextSequence(emitter common.Address) (uint64, error) {
	return _Messages.Contract.NextSequence(&_Messages.CallOpts, emitter)
}

// ParseAndVerifyVM is a free data retrieval call binding the contract method 0xc0fd8bde.
//
// Solidity: function parseAndVerifyVM(bytes encodedVM) view returns((uint8,uint32,uint32,uint16,bytes32,uint64,uint8,bytes,uint32,(bytes32,bytes32,uint8,uint8)[],bytes32) vm, bool valid, string reason)
func (_Messages *MessagesCaller) ParseAndVerifyVM(opts *bind.CallOpts, encodedVM []byte) (struct {
	Vm     StructsVM
	Valid  bool
	Reason string
}, error) {
	ret := new(struct {
		Vm     StructsVM
		Valid  bool
		Reason string
	})
	out := ret
	err := _Messages.contract.Call(opts, out, "parseAndVerifyVM", encodedVM)
	return *ret, err
}

// ParseAndVerifyVM is a free data retrieval call binding the contract method 0xc0fd8bde.
//
// Solidity: function parseAndVerifyVM(bytes encodedVM) view returns((uint8,uint32,uint32,uint16,bytes32,uint64,uint8,bytes,uint32,(bytes32,bytes32,uint8,uint8)[],bytes32) vm, bool valid, string reason)
func (_Messages *MessagesSession) ParseAndVerifyVM(encodedVM []byte) (struct {
	Vm     StructsVM
	Valid  bool
	Reason string
}, error) {
	return _Messages.Contract.ParseAndVerifyVM(&_Messages.CallOpts, encodedVM)
}

// ParseAndVerifyVM is a free data retrieval call binding the contract method 0xc0fd8bde.
//
// Solidity: function parseAndVerifyVM(bytes encodedVM) view returns((uint8,uint32,uint32,uint16,bytes32,uint64,uint8,bytes,uint32,(bytes32,bytes32,uint8,uint8)[],bytes32) vm, bool valid, string reason)
func (_Messages *MessagesCallerSession) ParseAndVerifyVM(encodedVM []byte) (struct {
	Vm     StructsVM
	Valid  bool
	Reason string
}, error) {
	return _Messages.Contract.ParseAndVerifyVM(&_Messages.CallOpts, encodedVM)
}

// ParseVM is a free data retrieval call binding the contract method 0xa9e11893.
//
// Solidity: function parseVM(bytes encodedVM) pure returns((uint8,uint32,uint32,uint16,bytes32,uint64,uint8,bytes,uint32,(bytes32,bytes32,uint8,uint8)[],bytes32) vm)
func (_Messages *MessagesCaller) ParseVM(opts *bind.CallOpts, encodedVM []byte) (StructsVM, error) {
	var (
		ret0 = new(StructsVM)
	)
	out := ret0
	err := _Messages.contract.Call(opts, out, "parseVM", encodedVM)
	return *ret0, err
}

// ParseVM is a free data retrieval call binding the contract method 0xa9e11893.
//
// Solidity: function parseVM(bytes encodedVM) pure returns((uint8,uint32,uint32,uint16,bytes32,uint64,uint8,bytes,uint32,(bytes32,bytes32,uint8,uint8)[],bytes32) vm)
func (_Messages *MessagesSession) ParseVM(encodedVM []byte) (StructsVM, error) {
	return _Messages.Contract.ParseVM(&_Messages.CallOpts, encodedVM)
}

// ParseVM is a free data retrieval call binding the contract method 0xa9e11893.
//
// Solidity: function parseVM(bytes encodedVM) pure returns((uint8,uint32,uint32,uint16,bytes32,uint64,uint8,bytes,uint32,(bytes32,bytes32,uint8,uint8)[],bytes32) vm)
func (_Messages *MessagesCallerSession) ParseVM(encodedVM []byte) (StructsVM, error) {
	return _Messages.Contract.ParseVM(&_Messages.CallOpts, encodedVM)
}

// VerifySignatures is a free data retrieval call binding the contract method 0xa0cce1b3.
//
// Solidity: function verifySignatures(bytes32 hash, (bytes32,bytes32,uint8,uint8)[] signatures, (address[],uint32) guardianSet) pure returns(bool valid, string reason)
func (_Messages *MessagesCaller) VerifySignatures(opts *bind.CallOpts, hash [32]byte, signatures []StructsSignature, guardianSet StructsGuardianSet) (struct {
	Valid  bool
	Reason string
}, error) {
	ret := new(struct {
		Valid  bool
		Reason string
	})
	out := ret
	err := _Messages.contract.Call(opts, out, "verifySignatures", hash, signatures, guardianSet)
	return *ret, err
}

// VerifySignatures is a free data retrieval call binding the contract method 0xa0cce1b3.
//
// Solidity: function verifySignatures(bytes32 hash, (bytes32,bytes32,uint8,uint8)[] signatures, (address[],uint32) guardianSet) pure returns(bool valid, string reason)
func (_Messages *MessagesSession) VerifySignatures(hash [32]byte, signatures []StructsSignature, guardianSet StructsGuardianSet) (struct {
	Valid  bool
	Reason string
}, error) {
	return _Messages.Contract.VerifySignatures(&_Messages.CallOpts, hash, signatures, guardianSet)
}

// VerifySignatures is a free data retrieval call binding the contract method 0xa0cce1b3.
//
// Solidity: function verifySignatures(bytes32 hash, (bytes32,bytes32,uint8,uint8)[] signatures, (address[],uint32) guardianSet) pure returns(bool valid, string reason)
func (_Messages *MessagesCallerSession) VerifySignatures(hash [32]byte, signatures []StructsSignature, guardianSet StructsGuardianSet) (struct {
	Valid  bool
	Reason string
}, error) {
	return _Messages.Contract.VerifySignatures(&_Messages.CallOpts, hash, signatures, guardianSet)
}

// VerifyVM is a free data retrieval call binding the contract method 0x875be02a.
//
// Solidity: function verifyVM((uint8,uint32,uint32,uint16,bytes32,uint64,uint8,bytes,uint32,(bytes32,bytes32,uint8,uint8)[],bytes32) vm) view returns(bool valid, string reason)
func (_Messages *MessagesCaller) VerifyVM(opts *bind.CallOpts, vm StructsVM) (struct {
	Valid  bool
	Reason string
}, error) {
	ret := new(struct {
		Valid  bool
		Reason string
	})
	out := ret
	err := _Messages.contract.Call(opts, out, "verifyVM", vm)
	return *ret, err
}

// VerifyVM is a free data retrieval call binding the contract method 0x875be02a.
//
// Solidity: function verifyVM((uint8,uint32,uint32,uint16,bytes32,uint64,uint8,bytes,uint32,(bytes32,bytes32,uint8,uint8)[],bytes32) vm) view returns(bool valid, string reason)
func (_Messages *MessagesSession) VerifyVM(vm StructsVM) (struct {
	Valid  bool
	Reason string
}, error) {
	return _Messages.Contract.VerifyVM(&_Messages.CallOpts, vm)
}

// VerifyVM is a free data retrieval call binding the contract method 0x875be02a.
//
// Solidity: function verifyVM((uint8,uint32,uint32,uint16,bytes32,uint64,uint8,bytes,uint32,(bytes32,bytes32,uint8,uint8)[],bytes32) vm) view returns(bool valid, string reason)
func (_Messages *MessagesCallerSession) VerifyVM(vm StructsVM) (struct {
	Valid  bool
	Reason string
}, error) {
	return _Messages.Contract.VerifyVM(&_Messages.CallOpts, vm)
}

// StateABI is the input ABI used to generate the binding from.
const StateABI = "[]"

// StateBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const StateBinRuntime = `6080604052600080fdfea26469706673582212204f4e286d4bb295ff9991551a502832ab24a0463129f1cdbd32efbce5a250dbc964736f6c637828302e382e31322d646576656c6f702e323032322e312e31322b636f6d6d69742e61373131393639390059`

// StateBin is the compiled bytecode used for deploying new contracts.
var StateBin = "0x6080604052348015600f57600080fd5b50606580601d6000396000f3fe6080604052600080fdfea26469706673582212204f4e286d4bb295ff9991551a502832ab24a0463129f1cdbd32efbce5a250dbc964736f6c637828302e382e31322d646576656c6f702e323032322e312e31322b636f6d6d69742e61373131393639390059"

// DeployState deploys a new Klaytn contract, binding an instance of State to it.
func DeployState(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *State, error) {
	parsed, err := abi.JSON(strings.NewReader(StateABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &State{StateCaller: StateCaller{contract: contract}, StateTransactor: StateTransactor{contract: contract}, StateFilterer: StateFilterer{contract: contract}}, nil
}

// State is an auto generated Go binding around a Klaytn contract.
type State struct {
	StateCaller     // Read-only binding to the contract
	StateTransactor // Write-only binding to the contract
	StateFilterer   // Log filterer for contract events
}

// StateCaller is an auto generated read-only Go binding around a Klaytn contract.
type StateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateTransactor is an auto generated write-only Go binding around a Klaytn contract.
type StateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type StateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type StateSession struct {
	Contract     *State            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type StateCallerSession struct {
	Contract *StateCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StateTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type StateTransactorSession struct {
	Contract     *StateTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateRaw is an auto generated low-level Go binding around a Klaytn contract.
type StateRaw struct {
	Contract *State // Generic contract binding to access the raw methods on
}

// StateCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type StateCallerRaw struct {
	Contract *StateCaller // Generic read-only contract binding to access the raw methods on
}

// StateTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type StateTransactorRaw struct {
	Contract *StateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewState creates a new instance of State, bound to a specific deployed contract.
func NewState(address common.Address, backend bind.ContractBackend) (*State, error) {
	contract, err := bindState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &State{StateCaller: StateCaller{contract: contract}, StateTransactor: StateTransactor{contract: contract}, StateFilterer: StateFilterer{contract: contract}}, nil
}

// NewStateCaller creates a new read-only instance of State, bound to a specific deployed contract.
func NewStateCaller(address common.Address, caller bind.ContractCaller) (*StateCaller, error) {
	contract, err := bindState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateCaller{contract: contract}, nil
}

// NewStateTransactor creates a new write-only instance of State, bound to a specific deployed contract.
func NewStateTransactor(address common.Address, transactor bind.ContractTransactor) (*StateTransactor, error) {
	contract, err := bindState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateTransactor{contract: contract}, nil
}

// NewStateFilterer creates a new log filterer instance of State, bound to a specific deployed contract.
func NewStateFilterer(address common.Address, filterer bind.ContractFilterer) (*StateFilterer, error) {
	contract, err := bindState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateFilterer{contract: contract}, nil
}

// bindState binds a generic wrapper to an already deployed contract.
func bindState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_State *StateRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _State.Contract.StateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_State *StateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _State.Contract.StateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_State *StateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _State.Contract.StateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_State *StateCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _State.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_State *StateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _State.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_State *StateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _State.Contract.contract.Transact(opts, method, params...)
}

// StorageABI is the input ABI used to generate the binding from.
const StorageABI = "[]"

// StorageBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const StorageBinRuntime = `6080604052600080fdfea264697066735822122092658ac640b819394f0858ff37b31a190d6bc3a18191f8581ce8e819f1eb2aa164736f6c637828302e382e31322d646576656c6f702e323032322e312e31322b636f6d6d69742e61373131393639390059`

// StorageBin is the compiled bytecode used for deploying new contracts.
var StorageBin = "0x6080604052348015600f57600080fd5b50606580601d6000396000f3fe6080604052600080fdfea264697066735822122092658ac640b819394f0858ff37b31a190d6bc3a18191f8581ce8e819f1eb2aa164736f6c637828302e382e31322d646576656c6f702e323032322e312e31322b636f6d6d69742e61373131393639390059"

// DeployStorage deploys a new Klaytn contract, binding an instance of Storage to it.
func DeployStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Storage, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// Storage is an auto generated Go binding around a Klaytn contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around a Klaytn contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around a Klaytn contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around a Klaytn contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// StructsABI is the input ABI used to generate the binding from.
const StructsABI = "[]"

// StructsBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const StructsBinRuntime = ``

// Structs is an auto generated Go binding around a Klaytn contract.
type Structs struct {
	StructsCaller     // Read-only binding to the contract
	StructsTransactor // Write-only binding to the contract
	StructsFilterer   // Log filterer for contract events
}

// StructsCaller is an auto generated read-only Go binding around a Klaytn contract.
type StructsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructsTransactor is an auto generated write-only Go binding around a Klaytn contract.
type StructsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructsFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type StructsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructsSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type StructsSession struct {
	Contract     *Structs          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StructsCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type StructsCallerSession struct {
	Contract *StructsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StructsTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type StructsTransactorSession struct {
	Contract     *StructsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StructsRaw is an auto generated low-level Go binding around a Klaytn contract.
type StructsRaw struct {
	Contract *Structs // Generic contract binding to access the raw methods on
}

// StructsCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type StructsCallerRaw struct {
	Contract *StructsCaller // Generic read-only contract binding to access the raw methods on
}

// StructsTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type StructsTransactorRaw struct {
	Contract *StructsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStructs creates a new instance of Structs, bound to a specific deployed contract.
func NewStructs(address common.Address, backend bind.ContractBackend) (*Structs, error) {
	contract, err := bindStructs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Structs{StructsCaller: StructsCaller{contract: contract}, StructsTransactor: StructsTransactor{contract: contract}, StructsFilterer: StructsFilterer{contract: contract}}, nil
}

// NewStructsCaller creates a new read-only instance of Structs, bound to a specific deployed contract.
func NewStructsCaller(address common.Address, caller bind.ContractCaller) (*StructsCaller, error) {
	contract, err := bindStructs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StructsCaller{contract: contract}, nil
}

// NewStructsTransactor creates a new write-only instance of Structs, bound to a specific deployed contract.
func NewStructsTransactor(address common.Address, transactor bind.ContractTransactor) (*StructsTransactor, error) {
	contract, err := bindStructs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StructsTransactor{contract: contract}, nil
}

// NewStructsFilterer creates a new log filterer instance of Structs, bound to a specific deployed contract.
func NewStructsFilterer(address common.Address, filterer bind.ContractFilterer) (*StructsFilterer, error) {
	contract, err := bindStructs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StructsFilterer{contract: contract}, nil
}

// bindStructs binds a generic wrapper to an already deployed contract.
func bindStructs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StructsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Structs *StructsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Structs.Contract.StructsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Structs *StructsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Structs.Contract.StructsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Structs *StructsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Structs.Contract.StructsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Structs *StructsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Structs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Structs *StructsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Structs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Structs *StructsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Structs.Contract.contract.Transact(opts, method, params...)
}

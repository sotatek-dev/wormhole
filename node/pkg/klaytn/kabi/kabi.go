// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kabi

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

// GovernanceStructsContractUpgrade is an auto generated low-level Go binding around an user-defined struct.
type GovernanceStructsContractUpgrade struct {
	Module      [32]byte
	Action      uint8
	Chain       uint16
	NewContract common.Address
}

// GovernanceStructsGuardianSetUpgrade is an auto generated low-level Go binding around an user-defined struct.
type GovernanceStructsGuardianSetUpgrade struct {
	Module              [32]byte
	Action              uint8
	Chain               uint16
	NewGuardianSet      StructsGuardianSet
	NewGuardianSetIndex uint32
}

// GovernanceStructsSetMessageFee is an auto generated low-level Go binding around an user-defined struct.
type GovernanceStructsSetMessageFee struct {
	Module     [32]byte
	Action     uint8
	Chain      uint16
	MessageFee *big.Int
}

// GovernanceStructsTransferFees is an auto generated low-level Go binding around an user-defined struct.
type GovernanceStructsTransferFees struct {
	Module    [32]byte
	Action    uint8
	Chain     uint16
	Amount    *big.Int
	Recipient [32]byte
}

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

// KabiABI is the input ABI used to generate the binding from.
const KabiABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"GuardianSetAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"consistencyLevel\",\"type\":\"uint8\"}],\"name\":\"LogMessagePublished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentGuardianSetIndex\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"getGuardianSet\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"keys\",\"type\":\"address[]\"},{\"internalType\":\"uint32\",\"name\":\"expirationTime\",\"type\":\"uint32\"}],\"internalType\":\"structStructs.GuardianSet\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGuardianSetExpiry\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"governanceActionIsConsumed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceChainId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceContract\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"emitter\",\"type\":\"address\"}],\"name\":\"nextSequence\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedVM\",\"type\":\"bytes\"}],\"name\":\"parseAndVerifyVM\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"emitterChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"emitterAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"consistencyLevel\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"guardianSetIndex\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"guardianIndex\",\"type\":\"uint8\"}],\"internalType\":\"structStructs.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"internalType\":\"structStructs.VM\",\"name\":\"vm\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedUpgrade\",\"type\":\"bytes\"}],\"name\":\"parseContractUpgrade\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"module\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"chain\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"newContract\",\"type\":\"address\"}],\"internalType\":\"structGovernanceStructs.ContractUpgrade\",\"name\":\"cu\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedUpgrade\",\"type\":\"bytes\"}],\"name\":\"parseGuardianSetUpgrade\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"module\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"chain\",\"type\":\"uint16\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"keys\",\"type\":\"address[]\"},{\"internalType\":\"uint32\",\"name\":\"expirationTime\",\"type\":\"uint32\"}],\"internalType\":\"structStructs.GuardianSet\",\"name\":\"newGuardianSet\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"newGuardianSetIndex\",\"type\":\"uint32\"}],\"internalType\":\"structGovernanceStructs.GuardianSetUpgrade\",\"name\":\"gsu\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedSetMessageFee\",\"type\":\"bytes\"}],\"name\":\"parseSetMessageFee\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"module\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"chain\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"messageFee\",\"type\":\"uint256\"}],\"internalType\":\"structGovernanceStructs.SetMessageFee\",\"name\":\"smf\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTransferFees\",\"type\":\"bytes\"}],\"name\":\"parseTransferFees\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"module\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"chain\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"}],\"internalType\":\"structGovernanceStructs.TransferFees\",\"name\":\"tf\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedVM\",\"type\":\"bytes\"}],\"name\":\"parseVM\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"emitterChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"emitterAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"consistencyLevel\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"guardianSetIndex\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"guardianIndex\",\"type\":\"uint8\"}],\"internalType\":\"structStructs.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"internalType\":\"structStructs.VM\",\"name\":\"vm\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"consistencyLevel\",\"type\":\"uint8\"}],\"name\":\"publishMessage\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_vm\",\"type\":\"bytes\"}],\"name\":\"submitContractUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_vm\",\"type\":\"bytes\"}],\"name\":\"submitNewGuardianSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_vm\",\"type\":\"bytes\"}],\"name\":\"submitSetMessageFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_vm\",\"type\":\"bytes\"}],\"name\":\"submitTransferFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"guardianIndex\",\"type\":\"uint8\"}],\"internalType\":\"structStructs.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"keys\",\"type\":\"address[]\"},{\"internalType\":\"uint32\",\"name\":\"expirationTime\",\"type\":\"uint32\"}],\"internalType\":\"structStructs.GuardianSet\",\"name\":\"guardianSet\",\"type\":\"tuple\"}],\"name\":\"verifySignatures\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"emitterChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"emitterAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"consistencyLevel\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"guardianSetIndex\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"guardianIndex\",\"type\":\"uint8\"}],\"internalType\":\"structStructs.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"internalType\":\"structStructs.VM\",\"name\":\"vm\",\"type\":\"tuple\"}],\"name\":\"verifyVM\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// KabiBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const KabiBinRuntime = ``

// Kabi is an auto generated Go binding around a Klaytn contract.
type Kabi struct {
	KabiCaller     // Read-only binding to the contract
	KabiTransactor // Write-only binding to the contract
	KabiFilterer   // Log filterer for contract events
}

// KabiCaller is an auto generated read-only Go binding around a Klaytn contract.
type KabiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KabiTransactor is an auto generated write-only Go binding around a Klaytn contract.
type KabiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KabiFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type KabiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KabiSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type KabiSession struct {
	Contract     *Kabi             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KabiCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type KabiCallerSession struct {
	Contract *KabiCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// KabiTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type KabiTransactorSession struct {
	Contract     *KabiTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KabiRaw is an auto generated low-level Go binding around a Klaytn contract.
type KabiRaw struct {
	Contract *Kabi // Generic contract binding to access the raw methods on
}

// KabiCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type KabiCallerRaw struct {
	Contract *KabiCaller // Generic read-only contract binding to access the raw methods on
}

// KabiTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type KabiTransactorRaw struct {
	Contract *KabiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKabi creates a new instance of Kabi, bound to a specific deployed contract.
func NewKabi(address common.Address, backend bind.ContractBackend) (*Kabi, error) {
	contract, err := bindKabi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kabi{KabiCaller: KabiCaller{contract: contract}, KabiTransactor: KabiTransactor{contract: contract}, KabiFilterer: KabiFilterer{contract: contract}}, nil
}

// NewKabiCaller creates a new read-only instance of Kabi, bound to a specific deployed contract.
func NewKabiCaller(address common.Address, caller bind.ContractCaller) (*KabiCaller, error) {
	contract, err := bindKabi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KabiCaller{contract: contract}, nil
}

// NewKabiTransactor creates a new write-only instance of Kabi, bound to a specific deployed contract.
func NewKabiTransactor(address common.Address, transactor bind.ContractTransactor) (*KabiTransactor, error) {
	contract, err := bindKabi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KabiTransactor{contract: contract}, nil
}

// NewKabiFilterer creates a new log filterer instance of Kabi, bound to a specific deployed contract.
func NewKabiFilterer(address common.Address, filterer bind.ContractFilterer) (*KabiFilterer, error) {
	contract, err := bindKabi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KabiFilterer{contract: contract}, nil
}

// bindKabi binds a generic wrapper to an already deployed contract.
func bindKabi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KabiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kabi *KabiRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Kabi.Contract.KabiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kabi *KabiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kabi.Contract.KabiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kabi *KabiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kabi.Contract.KabiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kabi *KabiCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Kabi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kabi *KabiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kabi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kabi *KabiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kabi.Contract.contract.Transact(opts, method, params...)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_Kabi *KabiCaller) ChainId(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Kabi.contract.Call(opts, out, "chainId")
	return *ret0, err
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_Kabi *KabiSession) ChainId() (uint16, error) {
	return _Kabi.Contract.ChainId(&_Kabi.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_Kabi *KabiCallerSession) ChainId() (uint16, error) {
	return _Kabi.Contract.ChainId(&_Kabi.CallOpts)
}

// GetCurrentGuardianSetIndex is a free data retrieval call binding the contract method 0x1cfe7951.
//
// Solidity: function getCurrentGuardianSetIndex() view returns(uint32)
func (_Kabi *KabiCaller) GetCurrentGuardianSetIndex(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Kabi.contract.Call(opts, out, "getCurrentGuardianSetIndex")
	return *ret0, err
}

// GetCurrentGuardianSetIndex is a free data retrieval call binding the contract method 0x1cfe7951.
//
// Solidity: function getCurrentGuardianSetIndex() view returns(uint32)
func (_Kabi *KabiSession) GetCurrentGuardianSetIndex() (uint32, error) {
	return _Kabi.Contract.GetCurrentGuardianSetIndex(&_Kabi.CallOpts)
}

// GetCurrentGuardianSetIndex is a free data retrieval call binding the contract method 0x1cfe7951.
//
// Solidity: function getCurrentGuardianSetIndex() view returns(uint32)
func (_Kabi *KabiCallerSession) GetCurrentGuardianSetIndex() (uint32, error) {
	return _Kabi.Contract.GetCurrentGuardianSetIndex(&_Kabi.CallOpts)
}

// GetGuardianSet is a free data retrieval call binding the contract method 0xf951975a.
//
// Solidity: function getGuardianSet(uint32 index) view returns((address[],uint32))
func (_Kabi *KabiCaller) GetGuardianSet(opts *bind.CallOpts, index uint32) (StructsGuardianSet, error) {
	var (
		ret0 = new(StructsGuardianSet)
	)
	out := ret0
	err := _Kabi.contract.Call(opts, out, "getGuardianSet", index)
	return *ret0, err
}

// GetGuardianSet is a free data retrieval call binding the contract method 0xf951975a.
//
// Solidity: function getGuardianSet(uint32 index) view returns((address[],uint32))
func (_Kabi *KabiSession) GetGuardianSet(index uint32) (StructsGuardianSet, error) {
	return _Kabi.Contract.GetGuardianSet(&_Kabi.CallOpts, index)
}

// GetGuardianSet is a free data retrieval call binding the contract method 0xf951975a.
//
// Solidity: function getGuardianSet(uint32 index) view returns((address[],uint32))
func (_Kabi *KabiCallerSession) GetGuardianSet(index uint32) (StructsGuardianSet, error) {
	return _Kabi.Contract.GetGuardianSet(&_Kabi.CallOpts, index)
}

// GetGuardianSetExpiry is a free data retrieval call binding the contract method 0xeb8d3f12.
//
// Solidity: function getGuardianSetExpiry() view returns(uint32)
func (_Kabi *KabiCaller) GetGuardianSetExpiry(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Kabi.contract.Call(opts, out, "getGuardianSetExpiry")
	return *ret0, err
}

// GetGuardianSetExpiry is a free data retrieval call binding the contract method 0xeb8d3f12.
//
// Solidity: function getGuardianSetExpiry() view returns(uint32)
func (_Kabi *KabiSession) GetGuardianSetExpiry() (uint32, error) {
	return _Kabi.Contract.GetGuardianSetExpiry(&_Kabi.CallOpts)
}

// GetGuardianSetExpiry is a free data retrieval call binding the contract method 0xeb8d3f12.
//
// Solidity: function getGuardianSetExpiry() view returns(uint32)
func (_Kabi *KabiCallerSession) GetGuardianSetExpiry() (uint32, error) {
	return _Kabi.Contract.GetGuardianSetExpiry(&_Kabi.CallOpts)
}

// GovernanceActionIsConsumed is a free data retrieval call binding the contract method 0x2c3c02a4.
//
// Solidity: function governanceActionIsConsumed(bytes32 hash) view returns(bool)
func (_Kabi *KabiCaller) GovernanceActionIsConsumed(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Kabi.contract.Call(opts, out, "governanceActionIsConsumed", hash)
	return *ret0, err
}

// GovernanceActionIsConsumed is a free data retrieval call binding the contract method 0x2c3c02a4.
//
// Solidity: function governanceActionIsConsumed(bytes32 hash) view returns(bool)
func (_Kabi *KabiSession) GovernanceActionIsConsumed(hash [32]byte) (bool, error) {
	return _Kabi.Contract.GovernanceActionIsConsumed(&_Kabi.CallOpts, hash)
}

// GovernanceActionIsConsumed is a free data retrieval call binding the contract method 0x2c3c02a4.
//
// Solidity: function governanceActionIsConsumed(bytes32 hash) view returns(bool)
func (_Kabi *KabiCallerSession) GovernanceActionIsConsumed(hash [32]byte) (bool, error) {
	return _Kabi.Contract.GovernanceActionIsConsumed(&_Kabi.CallOpts, hash)
}

// GovernanceChainId is a free data retrieval call binding the contract method 0xfbe3c2cd.
//
// Solidity: function governanceChainId() view returns(uint16)
func (_Kabi *KabiCaller) GovernanceChainId(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Kabi.contract.Call(opts, out, "governanceChainId")
	return *ret0, err
}

// GovernanceChainId is a free data retrieval call binding the contract method 0xfbe3c2cd.
//
// Solidity: function governanceChainId() view returns(uint16)
func (_Kabi *KabiSession) GovernanceChainId() (uint16, error) {
	return _Kabi.Contract.GovernanceChainId(&_Kabi.CallOpts)
}

// GovernanceChainId is a free data retrieval call binding the contract method 0xfbe3c2cd.
//
// Solidity: function governanceChainId() view returns(uint16)
func (_Kabi *KabiCallerSession) GovernanceChainId() (uint16, error) {
	return _Kabi.Contract.GovernanceChainId(&_Kabi.CallOpts)
}

// GovernanceContract is a free data retrieval call binding the contract method 0xb172b222.
//
// Solidity: function governanceContract() view returns(bytes32)
func (_Kabi *KabiCaller) GovernanceContract(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Kabi.contract.Call(opts, out, "governanceContract")
	return *ret0, err
}

// GovernanceContract is a free data retrieval call binding the contract method 0xb172b222.
//
// Solidity: function governanceContract() view returns(bytes32)
func (_Kabi *KabiSession) GovernanceContract() ([32]byte, error) {
	return _Kabi.Contract.GovernanceContract(&_Kabi.CallOpts)
}

// GovernanceContract is a free data retrieval call binding the contract method 0xb172b222.
//
// Solidity: function governanceContract() view returns(bytes32)
func (_Kabi *KabiCallerSession) GovernanceContract() ([32]byte, error) {
	return _Kabi.Contract.GovernanceContract(&_Kabi.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0xd60b347f.
//
// Solidity: function isInitialized(address impl) view returns(bool)
func (_Kabi *KabiCaller) IsInitialized(opts *bind.CallOpts, impl common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Kabi.contract.Call(opts, out, "isInitialized", impl)
	return *ret0, err
}

// IsInitialized is a free data retrieval call binding the contract method 0xd60b347f.
//
// Solidity: function isInitialized(address impl) view returns(bool)
func (_Kabi *KabiSession) IsInitialized(impl common.Address) (bool, error) {
	return _Kabi.Contract.IsInitialized(&_Kabi.CallOpts, impl)
}

// IsInitialized is a free data retrieval call binding the contract method 0xd60b347f.
//
// Solidity: function isInitialized(address impl) view returns(bool)
func (_Kabi *KabiCallerSession) IsInitialized(impl common.Address) (bool, error) {
	return _Kabi.Contract.IsInitialized(&_Kabi.CallOpts, impl)
}

// MessageFee is a free data retrieval call binding the contract method 0x1a90a219.
//
// Solidity: function messageFee() view returns(uint256)
func (_Kabi *KabiCaller) MessageFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Kabi.contract.Call(opts, out, "messageFee")
	return *ret0, err
}

// MessageFee is a free data retrieval call binding the contract method 0x1a90a219.
//
// Solidity: function messageFee() view returns(uint256)
func (_Kabi *KabiSession) MessageFee() (*big.Int, error) {
	return _Kabi.Contract.MessageFee(&_Kabi.CallOpts)
}

// MessageFee is a free data retrieval call binding the contract method 0x1a90a219.
//
// Solidity: function messageFee() view returns(uint256)
func (_Kabi *KabiCallerSession) MessageFee() (*big.Int, error) {
	return _Kabi.Contract.MessageFee(&_Kabi.CallOpts)
}

// NextSequence is a free data retrieval call binding the contract method 0x4cf842b5.
//
// Solidity: function nextSequence(address emitter) view returns(uint64)
func (_Kabi *KabiCaller) NextSequence(opts *bind.CallOpts, emitter common.Address) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Kabi.contract.Call(opts, out, "nextSequence", emitter)
	return *ret0, err
}

// NextSequence is a free data retrieval call binding the contract method 0x4cf842b5.
//
// Solidity: function nextSequence(address emitter) view returns(uint64)
func (_Kabi *KabiSession) NextSequence(emitter common.Address) (uint64, error) {
	return _Kabi.Contract.NextSequence(&_Kabi.CallOpts, emitter)
}

// NextSequence is a free data retrieval call binding the contract method 0x4cf842b5.
//
// Solidity: function nextSequence(address emitter) view returns(uint64)
func (_Kabi *KabiCallerSession) NextSequence(emitter common.Address) (uint64, error) {
	return _Kabi.Contract.NextSequence(&_Kabi.CallOpts, emitter)
}

// ParseAndVerifyVM is a free data retrieval call binding the contract method 0xc0fd8bde.
//
// Solidity: function parseAndVerifyVM(bytes encodedVM) view returns((uint8,uint32,uint32,uint16,bytes32,uint64,uint8,bytes,uint32,(bytes32,bytes32,uint8,uint8)[],bytes32) vm, bool valid, string reason)
func (_Kabi *KabiCaller) ParseAndVerifyVM(opts *bind.CallOpts, encodedVM []byte) (struct {
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
	err := _Kabi.contract.Call(opts, out, "parseAndVerifyVM", encodedVM)
	return *ret, err
}

// ParseAndVerifyVM is a free data retrieval call binding the contract method 0xc0fd8bde.
//
// Solidity: function parseAndVerifyVM(bytes encodedVM) view returns((uint8,uint32,uint32,uint16,bytes32,uint64,uint8,bytes,uint32,(bytes32,bytes32,uint8,uint8)[],bytes32) vm, bool valid, string reason)
func (_Kabi *KabiSession) ParseAndVerifyVM(encodedVM []byte) (struct {
	Vm     StructsVM
	Valid  bool
	Reason string
}, error) {
	return _Kabi.Contract.ParseAndVerifyVM(&_Kabi.CallOpts, encodedVM)
}

// ParseAndVerifyVM is a free data retrieval call binding the contract method 0xc0fd8bde.
//
// Solidity: function parseAndVerifyVM(bytes encodedVM) view returns((uint8,uint32,uint32,uint16,bytes32,uint64,uint8,bytes,uint32,(bytes32,bytes32,uint8,uint8)[],bytes32) vm, bool valid, string reason)
func (_Kabi *KabiCallerSession) ParseAndVerifyVM(encodedVM []byte) (struct {
	Vm     StructsVM
	Valid  bool
	Reason string
}, error) {
	return _Kabi.Contract.ParseAndVerifyVM(&_Kabi.CallOpts, encodedVM)
}

// ParseContractUpgrade is a free data retrieval call binding the contract method 0x4fdc60fa.
//
// Solidity: function parseContractUpgrade(bytes encodedUpgrade) pure returns((bytes32,uint8,uint16,address) cu)
func (_Kabi *KabiCaller) ParseContractUpgrade(opts *bind.CallOpts, encodedUpgrade []byte) (GovernanceStructsContractUpgrade, error) {
	var (
		ret0 = new(GovernanceStructsContractUpgrade)
	)
	out := ret0
	err := _Kabi.contract.Call(opts, out, "parseContractUpgrade", encodedUpgrade)
	return *ret0, err
}

// ParseContractUpgrade is a free data retrieval call binding the contract method 0x4fdc60fa.
//
// Solidity: function parseContractUpgrade(bytes encodedUpgrade) pure returns((bytes32,uint8,uint16,address) cu)
func (_Kabi *KabiSession) ParseContractUpgrade(encodedUpgrade []byte) (GovernanceStructsContractUpgrade, error) {
	return _Kabi.Contract.ParseContractUpgrade(&_Kabi.CallOpts, encodedUpgrade)
}

// ParseContractUpgrade is a free data retrieval call binding the contract method 0x4fdc60fa.
//
// Solidity: function parseContractUpgrade(bytes encodedUpgrade) pure returns((bytes32,uint8,uint16,address) cu)
func (_Kabi *KabiCallerSession) ParseContractUpgrade(encodedUpgrade []byte) (GovernanceStructsContractUpgrade, error) {
	return _Kabi.Contract.ParseContractUpgrade(&_Kabi.CallOpts, encodedUpgrade)
}

// ParseGuardianSetUpgrade is a free data retrieval call binding the contract method 0x04ca84cf.
//
// Solidity: function parseGuardianSetUpgrade(bytes encodedUpgrade) pure returns((bytes32,uint8,uint16,(address[],uint32),uint32) gsu)
func (_Kabi *KabiCaller) ParseGuardianSetUpgrade(opts *bind.CallOpts, encodedUpgrade []byte) (GovernanceStructsGuardianSetUpgrade, error) {
	var (
		ret0 = new(GovernanceStructsGuardianSetUpgrade)
	)
	out := ret0
	err := _Kabi.contract.Call(opts, out, "parseGuardianSetUpgrade", encodedUpgrade)
	return *ret0, err
}

// ParseGuardianSetUpgrade is a free data retrieval call binding the contract method 0x04ca84cf.
//
// Solidity: function parseGuardianSetUpgrade(bytes encodedUpgrade) pure returns((bytes32,uint8,uint16,(address[],uint32),uint32) gsu)
func (_Kabi *KabiSession) ParseGuardianSetUpgrade(encodedUpgrade []byte) (GovernanceStructsGuardianSetUpgrade, error) {
	return _Kabi.Contract.ParseGuardianSetUpgrade(&_Kabi.CallOpts, encodedUpgrade)
}

// ParseGuardianSetUpgrade is a free data retrieval call binding the contract method 0x04ca84cf.
//
// Solidity: function parseGuardianSetUpgrade(bytes encodedUpgrade) pure returns((bytes32,uint8,uint16,(address[],uint32),uint32) gsu)
func (_Kabi *KabiCallerSession) ParseGuardianSetUpgrade(encodedUpgrade []byte) (GovernanceStructsGuardianSetUpgrade, error) {
	return _Kabi.Contract.ParseGuardianSetUpgrade(&_Kabi.CallOpts, encodedUpgrade)
}

// ParseSetMessageFee is a free data retrieval call binding the contract method 0x515f3247.
//
// Solidity: function parseSetMessageFee(bytes encodedSetMessageFee) pure returns((bytes32,uint8,uint16,uint256) smf)
func (_Kabi *KabiCaller) ParseSetMessageFee(opts *bind.CallOpts, encodedSetMessageFee []byte) (GovernanceStructsSetMessageFee, error) {
	var (
		ret0 = new(GovernanceStructsSetMessageFee)
	)
	out := ret0
	err := _Kabi.contract.Call(opts, out, "parseSetMessageFee", encodedSetMessageFee)
	return *ret0, err
}

// ParseSetMessageFee is a free data retrieval call binding the contract method 0x515f3247.
//
// Solidity: function parseSetMessageFee(bytes encodedSetMessageFee) pure returns((bytes32,uint8,uint16,uint256) smf)
func (_Kabi *KabiSession) ParseSetMessageFee(encodedSetMessageFee []byte) (GovernanceStructsSetMessageFee, error) {
	return _Kabi.Contract.ParseSetMessageFee(&_Kabi.CallOpts, encodedSetMessageFee)
}

// ParseSetMessageFee is a free data retrieval call binding the contract method 0x515f3247.
//
// Solidity: function parseSetMessageFee(bytes encodedSetMessageFee) pure returns((bytes32,uint8,uint16,uint256) smf)
func (_Kabi *KabiCallerSession) ParseSetMessageFee(encodedSetMessageFee []byte) (GovernanceStructsSetMessageFee, error) {
	return _Kabi.Contract.ParseSetMessageFee(&_Kabi.CallOpts, encodedSetMessageFee)
}

// ParseTransferFees is a free data retrieval call binding the contract method 0x0319e59c.
//
// Solidity: function parseTransferFees(bytes encodedTransferFees) pure returns((bytes32,uint8,uint16,uint256,bytes32) tf)
func (_Kabi *KabiCaller) ParseTransferFees(opts *bind.CallOpts, encodedTransferFees []byte) (GovernanceStructsTransferFees, error) {
	var (
		ret0 = new(GovernanceStructsTransferFees)
	)
	out := ret0
	err := _Kabi.contract.Call(opts, out, "parseTransferFees", encodedTransferFees)
	return *ret0, err
}

// ParseTransferFees is a free data retrieval call binding the contract method 0x0319e59c.
//
// Solidity: function parseTransferFees(bytes encodedTransferFees) pure returns((bytes32,uint8,uint16,uint256,bytes32) tf)
func (_Kabi *KabiSession) ParseTransferFees(encodedTransferFees []byte) (GovernanceStructsTransferFees, error) {
	return _Kabi.Contract.ParseTransferFees(&_Kabi.CallOpts, encodedTransferFees)
}

// ParseTransferFees is a free data retrieval call binding the contract method 0x0319e59c.
//
// Solidity: function parseTransferFees(bytes encodedTransferFees) pure returns((bytes32,uint8,uint16,uint256,bytes32) tf)
func (_Kabi *KabiCallerSession) ParseTransferFees(encodedTransferFees []byte) (GovernanceStructsTransferFees, error) {
	return _Kabi.Contract.ParseTransferFees(&_Kabi.CallOpts, encodedTransferFees)
}

// ParseVM is a free data retrieval call binding the contract method 0xa9e11893.
//
// Solidity: function parseVM(bytes encodedVM) pure returns((uint8,uint32,uint32,uint16,bytes32,uint64,uint8,bytes,uint32,(bytes32,bytes32,uint8,uint8)[],bytes32) vm)
func (_Kabi *KabiCaller) ParseVM(opts *bind.CallOpts, encodedVM []byte) (StructsVM, error) {
	var (
		ret0 = new(StructsVM)
	)
	out := ret0
	err := _Kabi.contract.Call(opts, out, "parseVM", encodedVM)
	return *ret0, err
}

// ParseVM is a free data retrieval call binding the contract method 0xa9e11893.
//
// Solidity: function parseVM(bytes encodedVM) pure returns((uint8,uint32,uint32,uint16,bytes32,uint64,uint8,bytes,uint32,(bytes32,bytes32,uint8,uint8)[],bytes32) vm)
func (_Kabi *KabiSession) ParseVM(encodedVM []byte) (StructsVM, error) {
	return _Kabi.Contract.ParseVM(&_Kabi.CallOpts, encodedVM)
}

// ParseVM is a free data retrieval call binding the contract method 0xa9e11893.
//
// Solidity: function parseVM(bytes encodedVM) pure returns((uint8,uint32,uint32,uint16,bytes32,uint64,uint8,bytes,uint32,(bytes32,bytes32,uint8,uint8)[],bytes32) vm)
func (_Kabi *KabiCallerSession) ParseVM(encodedVM []byte) (StructsVM, error) {
	return _Kabi.Contract.ParseVM(&_Kabi.CallOpts, encodedVM)
}

// VerifySignatures is a free data retrieval call binding the contract method 0xa0cce1b3.
//
// Solidity: function verifySignatures(bytes32 hash, (bytes32,bytes32,uint8,uint8)[] signatures, (address[],uint32) guardianSet) pure returns(bool valid, string reason)
func (_Kabi *KabiCaller) VerifySignatures(opts *bind.CallOpts, hash [32]byte, signatures []StructsSignature, guardianSet StructsGuardianSet) (struct {
	Valid  bool
	Reason string
}, error) {
	ret := new(struct {
		Valid  bool
		Reason string
	})
	out := ret
	err := _Kabi.contract.Call(opts, out, "verifySignatures", hash, signatures, guardianSet)
	return *ret, err
}

// VerifySignatures is a free data retrieval call binding the contract method 0xa0cce1b3.
//
// Solidity: function verifySignatures(bytes32 hash, (bytes32,bytes32,uint8,uint8)[] signatures, (address[],uint32) guardianSet) pure returns(bool valid, string reason)
func (_Kabi *KabiSession) VerifySignatures(hash [32]byte, signatures []StructsSignature, guardianSet StructsGuardianSet) (struct {
	Valid  bool
	Reason string
}, error) {
	return _Kabi.Contract.VerifySignatures(&_Kabi.CallOpts, hash, signatures, guardianSet)
}

// VerifySignatures is a free data retrieval call binding the contract method 0xa0cce1b3.
//
// Solidity: function verifySignatures(bytes32 hash, (bytes32,bytes32,uint8,uint8)[] signatures, (address[],uint32) guardianSet) pure returns(bool valid, string reason)
func (_Kabi *KabiCallerSession) VerifySignatures(hash [32]byte, signatures []StructsSignature, guardianSet StructsGuardianSet) (struct {
	Valid  bool
	Reason string
}, error) {
	return _Kabi.Contract.VerifySignatures(&_Kabi.CallOpts, hash, signatures, guardianSet)
}

// VerifyVM is a free data retrieval call binding the contract method 0x875be02a.
//
// Solidity: function verifyVM((uint8,uint32,uint32,uint16,bytes32,uint64,uint8,bytes,uint32,(bytes32,bytes32,uint8,uint8)[],bytes32) vm) view returns(bool valid, string reason)
func (_Kabi *KabiCaller) VerifyVM(opts *bind.CallOpts, vm StructsVM) (struct {
	Valid  bool
	Reason string
}, error) {
	ret := new(struct {
		Valid  bool
		Reason string
	})
	out := ret
	err := _Kabi.contract.Call(opts, out, "verifyVM", vm)
	return *ret, err
}

// VerifyVM is a free data retrieval call binding the contract method 0x875be02a.
//
// Solidity: function verifyVM((uint8,uint32,uint32,uint16,bytes32,uint64,uint8,bytes,uint32,(bytes32,bytes32,uint8,uint8)[],bytes32) vm) view returns(bool valid, string reason)
func (_Kabi *KabiSession) VerifyVM(vm StructsVM) (struct {
	Valid  bool
	Reason string
}, error) {
	return _Kabi.Contract.VerifyVM(&_Kabi.CallOpts, vm)
}

// VerifyVM is a free data retrieval call binding the contract method 0x875be02a.
//
// Solidity: function verifyVM((uint8,uint32,uint32,uint16,bytes32,uint64,uint8,bytes,uint32,(bytes32,bytes32,uint8,uint8)[],bytes32) vm) view returns(bool valid, string reason)
func (_Kabi *KabiCallerSession) VerifyVM(vm StructsVM) (struct {
	Valid  bool
	Reason string
}, error) {
	return _Kabi.Contract.VerifyVM(&_Kabi.CallOpts, vm)
}

// PublishMessage is a paid mutator transaction binding the contract method 0xb19a437e.
//
// Solidity: function publishMessage(uint32 nonce, bytes payload, uint8 consistencyLevel) payable returns(uint64 sequence)
func (_Kabi *KabiTransactor) PublishMessage(opts *bind.TransactOpts, nonce uint32, payload []byte, consistencyLevel uint8) (*types.Transaction, error) {
	return _Kabi.contract.Transact(opts, "publishMessage", nonce, payload, consistencyLevel)
}

// PublishMessage is a paid mutator transaction binding the contract method 0xb19a437e.
//
// Solidity: function publishMessage(uint32 nonce, bytes payload, uint8 consistencyLevel) payable returns(uint64 sequence)
func (_Kabi *KabiSession) PublishMessage(nonce uint32, payload []byte, consistencyLevel uint8) (*types.Transaction, error) {
	return _Kabi.Contract.PublishMessage(&_Kabi.TransactOpts, nonce, payload, consistencyLevel)
}

// PublishMessage is a paid mutator transaction binding the contract method 0xb19a437e.
//
// Solidity: function publishMessage(uint32 nonce, bytes payload, uint8 consistencyLevel) payable returns(uint64 sequence)
func (_Kabi *KabiTransactorSession) PublishMessage(nonce uint32, payload []byte, consistencyLevel uint8) (*types.Transaction, error) {
	return _Kabi.Contract.PublishMessage(&_Kabi.TransactOpts, nonce, payload, consistencyLevel)
}

// SubmitContractUpgrade is a paid mutator transaction binding the contract method 0x5cb8cae2.
//
// Solidity: function submitContractUpgrade(bytes _vm) returns()
func (_Kabi *KabiTransactor) SubmitContractUpgrade(opts *bind.TransactOpts, _vm []byte) (*types.Transaction, error) {
	return _Kabi.contract.Transact(opts, "submitContractUpgrade", _vm)
}

// SubmitContractUpgrade is a paid mutator transaction binding the contract method 0x5cb8cae2.
//
// Solidity: function submitContractUpgrade(bytes _vm) returns()
func (_Kabi *KabiSession) SubmitContractUpgrade(_vm []byte) (*types.Transaction, error) {
	return _Kabi.Contract.SubmitContractUpgrade(&_Kabi.TransactOpts, _vm)
}

// SubmitContractUpgrade is a paid mutator transaction binding the contract method 0x5cb8cae2.
//
// Solidity: function submitContractUpgrade(bytes _vm) returns()
func (_Kabi *KabiTransactorSession) SubmitContractUpgrade(_vm []byte) (*types.Transaction, error) {
	return _Kabi.Contract.SubmitContractUpgrade(&_Kabi.TransactOpts, _vm)
}

// SubmitNewGuardianSet is a paid mutator transaction binding the contract method 0x6606b4e0.
//
// Solidity: function submitNewGuardianSet(bytes _vm) returns()
func (_Kabi *KabiTransactor) SubmitNewGuardianSet(opts *bind.TransactOpts, _vm []byte) (*types.Transaction, error) {
	return _Kabi.contract.Transact(opts, "submitNewGuardianSet", _vm)
}

// SubmitNewGuardianSet is a paid mutator transaction binding the contract method 0x6606b4e0.
//
// Solidity: function submitNewGuardianSet(bytes _vm) returns()
func (_Kabi *KabiSession) SubmitNewGuardianSet(_vm []byte) (*types.Transaction, error) {
	return _Kabi.Contract.SubmitNewGuardianSet(&_Kabi.TransactOpts, _vm)
}

// SubmitNewGuardianSet is a paid mutator transaction binding the contract method 0x6606b4e0.
//
// Solidity: function submitNewGuardianSet(bytes _vm) returns()
func (_Kabi *KabiTransactorSession) SubmitNewGuardianSet(_vm []byte) (*types.Transaction, error) {
	return _Kabi.Contract.SubmitNewGuardianSet(&_Kabi.TransactOpts, _vm)
}

// SubmitSetMessageFee is a paid mutator transaction binding the contract method 0xf42bc641.
//
// Solidity: function submitSetMessageFee(bytes _vm) returns()
func (_Kabi *KabiTransactor) SubmitSetMessageFee(opts *bind.TransactOpts, _vm []byte) (*types.Transaction, error) {
	return _Kabi.contract.Transact(opts, "submitSetMessageFee", _vm)
}

// SubmitSetMessageFee is a paid mutator transaction binding the contract method 0xf42bc641.
//
// Solidity: function submitSetMessageFee(bytes _vm) returns()
func (_Kabi *KabiSession) SubmitSetMessageFee(_vm []byte) (*types.Transaction, error) {
	return _Kabi.Contract.SubmitSetMessageFee(&_Kabi.TransactOpts, _vm)
}

// SubmitSetMessageFee is a paid mutator transaction binding the contract method 0xf42bc641.
//
// Solidity: function submitSetMessageFee(bytes _vm) returns()
func (_Kabi *KabiTransactorSession) SubmitSetMessageFee(_vm []byte) (*types.Transaction, error) {
	return _Kabi.Contract.SubmitSetMessageFee(&_Kabi.TransactOpts, _vm)
}

// SubmitTransferFees is a paid mutator transaction binding the contract method 0x93df337e.
//
// Solidity: function submitTransferFees(bytes _vm) returns()
func (_Kabi *KabiTransactor) SubmitTransferFees(opts *bind.TransactOpts, _vm []byte) (*types.Transaction, error) {
	return _Kabi.contract.Transact(opts, "submitTransferFees", _vm)
}

// SubmitTransferFees is a paid mutator transaction binding the contract method 0x93df337e.
//
// Solidity: function submitTransferFees(bytes _vm) returns()
func (_Kabi *KabiSession) SubmitTransferFees(_vm []byte) (*types.Transaction, error) {
	return _Kabi.Contract.SubmitTransferFees(&_Kabi.TransactOpts, _vm)
}

// SubmitTransferFees is a paid mutator transaction binding the contract method 0x93df337e.
//
// Solidity: function submitTransferFees(bytes _vm) returns()
func (_Kabi *KabiTransactorSession) SubmitTransferFees(_vm []byte) (*types.Transaction, error) {
	return _Kabi.Contract.SubmitTransferFees(&_Kabi.TransactOpts, _vm)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Kabi *KabiTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Kabi.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Kabi *KabiSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Kabi.Contract.Fallback(&_Kabi.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Kabi *KabiTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Kabi.Contract.Fallback(&_Kabi.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Kabi *KabiTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kabi.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Kabi *KabiSession) Receive() (*types.Transaction, error) {
	return _Kabi.Contract.Receive(&_Kabi.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Kabi *KabiTransactorSession) Receive() (*types.Transaction, error) {
	return _Kabi.Contract.Receive(&_Kabi.TransactOpts)
}

// KabiAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Kabi contract.
type KabiAdminChangedIterator struct {
	Event *KabiAdminChanged // Event containing the contract specifics and raw log

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
func (it *KabiAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KabiAdminChanged)
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
		it.Event = new(KabiAdminChanged)
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
func (it *KabiAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KabiAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KabiAdminChanged represents a AdminChanged event raised by the Kabi contract.
type KabiAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Kabi *KabiFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*KabiAdminChangedIterator, error) {

	logs, sub, err := _Kabi.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &KabiAdminChangedIterator{contract: _Kabi.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Kabi *KabiFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *KabiAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Kabi.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KabiAdminChanged)
				if err := _Kabi.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Kabi *KabiFilterer) ParseAdminChanged(log types.Log) (*KabiAdminChanged, error) {
	event := new(KabiAdminChanged)
	if err := _Kabi.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KabiBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Kabi contract.
type KabiBeaconUpgradedIterator struct {
	Event *KabiBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *KabiBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KabiBeaconUpgraded)
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
		it.Event = new(KabiBeaconUpgraded)
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
func (it *KabiBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KabiBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KabiBeaconUpgraded represents a BeaconUpgraded event raised by the Kabi contract.
type KabiBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Kabi *KabiFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*KabiBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Kabi.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &KabiBeaconUpgradedIterator{contract: _Kabi.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Kabi *KabiFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *KabiBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Kabi.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KabiBeaconUpgraded)
				if err := _Kabi.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Kabi *KabiFilterer) ParseBeaconUpgraded(log types.Log) (*KabiBeaconUpgraded, error) {
	event := new(KabiBeaconUpgraded)
	if err := _Kabi.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KabiContractUpgradedIterator is returned from FilterContractUpgraded and is used to iterate over the raw logs and unpacked data for ContractUpgraded events raised by the Kabi contract.
type KabiContractUpgradedIterator struct {
	Event *KabiContractUpgraded // Event containing the contract specifics and raw log

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
func (it *KabiContractUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KabiContractUpgraded)
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
		it.Event = new(KabiContractUpgraded)
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
func (it *KabiContractUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KabiContractUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KabiContractUpgraded represents a ContractUpgraded event raised by the Kabi contract.
type KabiContractUpgraded struct {
	OldContract common.Address
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgraded is a free log retrieval operation binding the contract event 0x2e4cc16c100f0b55e2df82ab0b1a7e294aa9cbd01b48fbaf622683fbc0507a49.
//
// Solidity: event ContractUpgraded(address indexed oldContract, address indexed newContract)
func (_Kabi *KabiFilterer) FilterContractUpgraded(opts *bind.FilterOpts, oldContract []common.Address, newContract []common.Address) (*KabiContractUpgradedIterator, error) {

	var oldContractRule []interface{}
	for _, oldContractItem := range oldContract {
		oldContractRule = append(oldContractRule, oldContractItem)
	}
	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _Kabi.contract.FilterLogs(opts, "ContractUpgraded", oldContractRule, newContractRule)
	if err != nil {
		return nil, err
	}
	return &KabiContractUpgradedIterator{contract: _Kabi.contract, event: "ContractUpgraded", logs: logs, sub: sub}, nil
}

// WatchContractUpgraded is a free log subscription operation binding the contract event 0x2e4cc16c100f0b55e2df82ab0b1a7e294aa9cbd01b48fbaf622683fbc0507a49.
//
// Solidity: event ContractUpgraded(address indexed oldContract, address indexed newContract)
func (_Kabi *KabiFilterer) WatchContractUpgraded(opts *bind.WatchOpts, sink chan<- *KabiContractUpgraded, oldContract []common.Address, newContract []common.Address) (event.Subscription, error) {

	var oldContractRule []interface{}
	for _, oldContractItem := range oldContract {
		oldContractRule = append(oldContractRule, oldContractItem)
	}
	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _Kabi.contract.WatchLogs(opts, "ContractUpgraded", oldContractRule, newContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KabiContractUpgraded)
				if err := _Kabi.contract.UnpackLog(event, "ContractUpgraded", log); err != nil {
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

// ParseContractUpgraded is a log parse operation binding the contract event 0x2e4cc16c100f0b55e2df82ab0b1a7e294aa9cbd01b48fbaf622683fbc0507a49.
//
// Solidity: event ContractUpgraded(address indexed oldContract, address indexed newContract)
func (_Kabi *KabiFilterer) ParseContractUpgraded(log types.Log) (*KabiContractUpgraded, error) {
	event := new(KabiContractUpgraded)
	if err := _Kabi.contract.UnpackLog(event, "ContractUpgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KabiGuardianSetAddedIterator is returned from FilterGuardianSetAdded and is used to iterate over the raw logs and unpacked data for GuardianSetAdded events raised by the Kabi contract.
type KabiGuardianSetAddedIterator struct {
	Event *KabiGuardianSetAdded // Event containing the contract specifics and raw log

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
func (it *KabiGuardianSetAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KabiGuardianSetAdded)
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
		it.Event = new(KabiGuardianSetAdded)
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
func (it *KabiGuardianSetAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KabiGuardianSetAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KabiGuardianSetAdded represents a GuardianSetAdded event raised by the Kabi contract.
type KabiGuardianSetAdded struct {
	Index uint32
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardianSetAdded is a free log retrieval operation binding the contract event 0x2384dbc52f7b617fb7c5aa71e5455a21ff21d58604bb6daef6af2bb44aadebdd.
//
// Solidity: event GuardianSetAdded(uint32 indexed index)
func (_Kabi *KabiFilterer) FilterGuardianSetAdded(opts *bind.FilterOpts, index []uint32) (*KabiGuardianSetAddedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Kabi.contract.FilterLogs(opts, "GuardianSetAdded", indexRule)
	if err != nil {
		return nil, err
	}
	return &KabiGuardianSetAddedIterator{contract: _Kabi.contract, event: "GuardianSetAdded", logs: logs, sub: sub}, nil
}

// WatchGuardianSetAdded is a free log subscription operation binding the contract event 0x2384dbc52f7b617fb7c5aa71e5455a21ff21d58604bb6daef6af2bb44aadebdd.
//
// Solidity: event GuardianSetAdded(uint32 indexed index)
func (_Kabi *KabiFilterer) WatchGuardianSetAdded(opts *bind.WatchOpts, sink chan<- *KabiGuardianSetAdded, index []uint32) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Kabi.contract.WatchLogs(opts, "GuardianSetAdded", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KabiGuardianSetAdded)
				if err := _Kabi.contract.UnpackLog(event, "GuardianSetAdded", log); err != nil {
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

// ParseGuardianSetAdded is a log parse operation binding the contract event 0x2384dbc52f7b617fb7c5aa71e5455a21ff21d58604bb6daef6af2bb44aadebdd.
//
// Solidity: event GuardianSetAdded(uint32 indexed index)
func (_Kabi *KabiFilterer) ParseGuardianSetAdded(log types.Log) (*KabiGuardianSetAdded, error) {
	event := new(KabiGuardianSetAdded)
	if err := _Kabi.contract.UnpackLog(event, "GuardianSetAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KabiLogMessagePublishedIterator is returned from FilterLogMessagePublished and is used to iterate over the raw logs and unpacked data for LogMessagePublished events raised by the Kabi contract.
type KabiLogMessagePublishedIterator struct {
	Event *KabiLogMessagePublished // Event containing the contract specifics and raw log

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
func (it *KabiLogMessagePublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KabiLogMessagePublished)
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
		it.Event = new(KabiLogMessagePublished)
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
func (it *KabiLogMessagePublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KabiLogMessagePublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KabiLogMessagePublished represents a LogMessagePublished event raised by the Kabi contract.
type KabiLogMessagePublished struct {
	Sender           common.Address
	Sequence         uint64
	Nonce            uint32
	Payload          []byte
	ConsistencyLevel uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogMessagePublished is a free log retrieval operation binding the contract event 0x6eb224fb001ed210e379b335e35efe88672a8ce935d981a6896b27ffdf52a3b2.
//
// Solidity: event LogMessagePublished(address indexed sender, uint64 sequence, uint32 nonce, bytes payload, uint8 consistencyLevel)
func (_Kabi *KabiFilterer) FilterLogMessagePublished(opts *bind.FilterOpts, sender []common.Address) (*KabiLogMessagePublishedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Kabi.contract.FilterLogs(opts, "LogMessagePublished", senderRule)
	if err != nil {
		return nil, err
	}
	return &KabiLogMessagePublishedIterator{contract: _Kabi.contract, event: "LogMessagePublished", logs: logs, sub: sub}, nil
}

// WatchLogMessagePublished is a free log subscription operation binding the contract event 0x6eb224fb001ed210e379b335e35efe88672a8ce935d981a6896b27ffdf52a3b2.
//
// Solidity: event LogMessagePublished(address indexed sender, uint64 sequence, uint32 nonce, bytes payload, uint8 consistencyLevel)
func (_Kabi *KabiFilterer) WatchLogMessagePublished(opts *bind.WatchOpts, sink chan<- *KabiLogMessagePublished, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Kabi.contract.WatchLogs(opts, "LogMessagePublished", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KabiLogMessagePublished)
				if err := _Kabi.contract.UnpackLog(event, "LogMessagePublished", log); err != nil {
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

// ParseLogMessagePublished is a log parse operation binding the contract event 0x6eb224fb001ed210e379b335e35efe88672a8ce935d981a6896b27ffdf52a3b2.
//
// Solidity: event LogMessagePublished(address indexed sender, uint64 sequence, uint32 nonce, bytes payload, uint8 consistencyLevel)
func (_Kabi *KabiFilterer) ParseLogMessagePublished(log types.Log) (*KabiLogMessagePublished, error) {
	event := new(KabiLogMessagePublished)
	if err := _Kabi.contract.UnpackLog(event, "LogMessagePublished", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KabiUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Kabi contract.
type KabiUpgradedIterator struct {
	Event *KabiUpgraded // Event containing the contract specifics and raw log

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
func (it *KabiUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KabiUpgraded)
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
		it.Event = new(KabiUpgraded)
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
func (it *KabiUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KabiUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KabiUpgraded represents a Upgraded event raised by the Kabi contract.
type KabiUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Kabi *KabiFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*KabiUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Kabi.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &KabiUpgradedIterator{contract: _Kabi.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Kabi *KabiFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *KabiUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Kabi.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KabiUpgraded)
				if err := _Kabi.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Kabi *KabiFilterer) ParseUpgraded(log types.Log) (*KabiUpgraded, error) {
	event := new(KabiUpgraded)
	if err := _Kabi.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}

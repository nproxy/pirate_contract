// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HopShopABI is the input ABI used to generate the binding from.
const HopShopABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"ChangeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"}],\"name\":\"SetLevelRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TokenNoToSellInLevel1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_adminWallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_rateLevel1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_rateLevel2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"remainingTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HopShop is an auto generated Go binding around an Ethereum contract.
type HopShop struct {
	HopShopCaller     // Read-only binding to the contract
	HopShopTransactor // Write-only binding to the contract
	HopShopFilterer   // Log filterer for contract events
}

// HopShopCaller is an auto generated read-only Go binding around an Ethereum contract.
type HopShopCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HopShopTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HopShopTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HopShopFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HopShopFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HopShopSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HopShopSession struct {
	Contract     *HopShop          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HopShopCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HopShopCallerSession struct {
	Contract *HopShopCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// HopShopTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HopShopTransactorSession struct {
	Contract     *HopShopTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// HopShopRaw is an auto generated low-level Go binding around an Ethereum contract.
type HopShopRaw struct {
	Contract *HopShop // Generic contract binding to access the raw methods on
}

// HopShopCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HopShopCallerRaw struct {
	Contract *HopShopCaller // Generic read-only contract binding to access the raw methods on
}

// HopShopTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HopShopTransactorRaw struct {
	Contract *HopShopTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHopShop creates a new instance of HopShop, bound to a specific deployed contract.
func NewHopShop(address common.Address, backend bind.ContractBackend) (*HopShop, error) {
	contract, err := bindHopShop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HopShop{HopShopCaller: HopShopCaller{contract: contract}, HopShopTransactor: HopShopTransactor{contract: contract}, HopShopFilterer: HopShopFilterer{contract: contract}}, nil
}

// NewHopShopCaller creates a new read-only instance of HopShop, bound to a specific deployed contract.
func NewHopShopCaller(address common.Address, caller bind.ContractCaller) (*HopShopCaller, error) {
	contract, err := bindHopShop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HopShopCaller{contract: contract}, nil
}

// NewHopShopTransactor creates a new write-only instance of HopShop, bound to a specific deployed contract.
func NewHopShopTransactor(address common.Address, transactor bind.ContractTransactor) (*HopShopTransactor, error) {
	contract, err := bindHopShop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HopShopTransactor{contract: contract}, nil
}

// NewHopShopFilterer creates a new log filterer instance of HopShop, bound to a specific deployed contract.
func NewHopShopFilterer(address common.Address, filterer bind.ContractFilterer) (*HopShopFilterer, error) {
	contract, err := bindHopShop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HopShopFilterer{contract: contract}, nil
}

// bindHopShop binds a generic wrapper to an already deployed contract.
func bindHopShop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HopShopABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HopShop *HopShopRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HopShop.Contract.HopShopCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HopShop *HopShopRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HopShop.Contract.HopShopTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HopShop *HopShopRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HopShop.Contract.HopShopTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HopShop *HopShopCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HopShop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HopShop *HopShopTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HopShop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HopShop *HopShopTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HopShop.Contract.contract.Transact(opts, method, params...)
}

// TokenNoToSellInLevel1 is a free data retrieval call binding the contract method 0x86bd3b70.
//
// Solidity: function TokenNoToSellInLevel1() view returns(uint256)
func (_HopShop *HopShopCaller) TokenNoToSellInLevel1(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HopShop.contract.Call(opts, out, "TokenNoToSellInLevel1")
	return *ret0, err
}

// TokenNoToSellInLevel1 is a free data retrieval call binding the contract method 0x86bd3b70.
//
// Solidity: function TokenNoToSellInLevel1() view returns(uint256)
func (_HopShop *HopShopSession) TokenNoToSellInLevel1() (*big.Int, error) {
	return _HopShop.Contract.TokenNoToSellInLevel1(&_HopShop.CallOpts)
}

// TokenNoToSellInLevel1 is a free data retrieval call binding the contract method 0x86bd3b70.
//
// Solidity: function TokenNoToSellInLevel1() view returns(uint256)
func (_HopShop *HopShopCallerSession) TokenNoToSellInLevel1() (*big.Int, error) {
	return _HopShop.Contract.TokenNoToSellInLevel1(&_HopShop.CallOpts)
}

// AdminWallet is a free data retrieval call binding the contract method 0x0ed22083.
//
// Solidity: function _adminWallet() view returns(address)
func (_HopShop *HopShopCaller) AdminWallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HopShop.contract.Call(opts, out, "_adminWallet")
	return *ret0, err
}

// AdminWallet is a free data retrieval call binding the contract method 0x0ed22083.
//
// Solidity: function _adminWallet() view returns(address)
func (_HopShop *HopShopSession) AdminWallet() (common.Address, error) {
	return _HopShop.Contract.AdminWallet(&_HopShop.CallOpts)
}

// AdminWallet is a free data retrieval call binding the contract method 0x0ed22083.
//
// Solidity: function _adminWallet() view returns(address)
func (_HopShop *HopShopCallerSession) AdminWallet() (common.Address, error) {
	return _HopShop.Contract.AdminWallet(&_HopShop.CallOpts)
}

// RateLevel1 is a free data retrieval call binding the contract method 0x93952cc2.
//
// Solidity: function _rateLevel1() view returns(uint256)
func (_HopShop *HopShopCaller) RateLevel1(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HopShop.contract.Call(opts, out, "_rateLevel1")
	return *ret0, err
}

// RateLevel1 is a free data retrieval call binding the contract method 0x93952cc2.
//
// Solidity: function _rateLevel1() view returns(uint256)
func (_HopShop *HopShopSession) RateLevel1() (*big.Int, error) {
	return _HopShop.Contract.RateLevel1(&_HopShop.CallOpts)
}

// RateLevel1 is a free data retrieval call binding the contract method 0x93952cc2.
//
// Solidity: function _rateLevel1() view returns(uint256)
func (_HopShop *HopShopCallerSession) RateLevel1() (*big.Int, error) {
	return _HopShop.Contract.RateLevel1(&_HopShop.CallOpts)
}

// RateLevel2 is a free data retrieval call binding the contract method 0x386800bf.
//
// Solidity: function _rateLevel2() view returns(uint256)
func (_HopShop *HopShopCaller) RateLevel2(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HopShop.contract.Call(opts, out, "_rateLevel2")
	return *ret0, err
}

// RateLevel2 is a free data retrieval call binding the contract method 0x386800bf.
//
// Solidity: function _rateLevel2() view returns(uint256)
func (_HopShop *HopShopSession) RateLevel2() (*big.Int, error) {
	return _HopShop.Contract.RateLevel2(&_HopShop.CallOpts)
}

// RateLevel2 is a free data retrieval call binding the contract method 0x386800bf.
//
// Solidity: function _rateLevel2() view returns(uint256)
func (_HopShop *HopShopCallerSession) RateLevel2() (*big.Int, error) {
	return _HopShop.Contract.RateLevel2(&_HopShop.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xecd0c0c3.
//
// Solidity: function _token() view returns(address)
func (_HopShop *HopShopCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HopShop.contract.Call(opts, out, "_token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xecd0c0c3.
//
// Solidity: function _token() view returns(address)
func (_HopShop *HopShopSession) Token() (common.Address, error) {
	return _HopShop.Contract.Token(&_HopShop.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xecd0c0c3.
//
// Solidity: function _token() view returns(address)
func (_HopShop *HopShopCallerSession) Token() (common.Address, error) {
	return _HopShop.Contract.Token(&_HopShop.CallOpts)
}

// HasRaised is a free data retrieval call binding the contract method 0xec1d66c0.
//
// Solidity: function hasRaised() view returns(uint256)
func (_HopShop *HopShopCaller) HasRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HopShop.contract.Call(opts, out, "hasRaised")
	return *ret0, err
}

// HasRaised is a free data retrieval call binding the contract method 0xec1d66c0.
//
// Solidity: function hasRaised() view returns(uint256)
func (_HopShop *HopShopSession) HasRaised() (*big.Int, error) {
	return _HopShop.Contract.HasRaised(&_HopShop.CallOpts)
}

// HasRaised is a free data retrieval call binding the contract method 0xec1d66c0.
//
// Solidity: function hasRaised() view returns(uint256)
func (_HopShop *HopShopCallerSession) HasRaised() (*big.Int, error) {
	return _HopShop.Contract.HasRaised(&_HopShop.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HopShop *HopShopCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HopShop.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HopShop *HopShopSession) Owner() (common.Address, error) {
	return _HopShop.Contract.Owner(&_HopShop.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HopShop *HopShopCallerSession) Owner() (common.Address, error) {
	return _HopShop.Contract.Owner(&_HopShop.CallOpts)
}

// RemainingTokens is a free data retrieval call binding the contract method 0xbf583903.
//
// Solidity: function remainingTokens() view returns(uint256)
func (_HopShop *HopShopCaller) RemainingTokens(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HopShop.contract.Call(opts, out, "remainingTokens")
	return *ret0, err
}

// RemainingTokens is a free data retrieval call binding the contract method 0xbf583903.
//
// Solidity: function remainingTokens() view returns(uint256)
func (_HopShop *HopShopSession) RemainingTokens() (*big.Int, error) {
	return _HopShop.Contract.RemainingTokens(&_HopShop.CallOpts)
}

// RemainingTokens is a free data retrieval call binding the contract method 0xbf583903.
//
// Solidity: function remainingTokens() view returns(uint256)
func (_HopShop *HopShopCallerSession) RemainingTokens() (*big.Int, error) {
	return _HopShop.Contract.RemainingTokens(&_HopShop.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x927cc064.
//
// Solidity: function ChangeAdmin(address wallet) returns()
func (_HopShop *HopShopTransactor) ChangeAdmin(opts *bind.TransactOpts, wallet common.Address) (*types.Transaction, error) {
	return _HopShop.contract.Transact(opts, "ChangeAdmin", wallet)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x927cc064.
//
// Solidity: function ChangeAdmin(address wallet) returns()
func (_HopShop *HopShopSession) ChangeAdmin(wallet common.Address) (*types.Transaction, error) {
	return _HopShop.Contract.ChangeAdmin(&_HopShop.TransactOpts, wallet)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x927cc064.
//
// Solidity: function ChangeAdmin(address wallet) returns()
func (_HopShop *HopShopTransactorSession) ChangeAdmin(wallet common.Address) (*types.Transaction, error) {
	return _HopShop.Contract.ChangeAdmin(&_HopShop.TransactOpts, wallet)
}

// SetLevelRate is a paid mutator transaction binding the contract method 0x9f377b78.
//
// Solidity: function SetLevelRate(uint256 newRate, uint8 level) returns()
func (_HopShop *HopShopTransactor) SetLevelRate(opts *bind.TransactOpts, newRate *big.Int, level uint8) (*types.Transaction, error) {
	return _HopShop.contract.Transact(opts, "SetLevelRate", newRate, level)
}

// SetLevelRate is a paid mutator transaction binding the contract method 0x9f377b78.
//
// Solidity: function SetLevelRate(uint256 newRate, uint8 level) returns()
func (_HopShop *HopShopSession) SetLevelRate(newRate *big.Int, level uint8) (*types.Transaction, error) {
	return _HopShop.Contract.SetLevelRate(&_HopShop.TransactOpts, newRate, level)
}

// SetLevelRate is a paid mutator transaction binding the contract method 0x9f377b78.
//
// Solidity: function SetLevelRate(uint256 newRate, uint8 level) returns()
func (_HopShop *HopShopTransactorSession) SetLevelRate(newRate *big.Int, level uint8) (*types.Transaction, error) {
	return _HopShop.Contract.SetLevelRate(&_HopShop.TransactOpts, newRate, level)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HopShop *HopShopTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HopShop.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HopShop *HopShopSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HopShop.Contract.TransferOwnership(&_HopShop.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HopShop *HopShopTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HopShop.Contract.TransferOwnership(&_HopShop.TransactOpts, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_HopShop *HopShopTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _HopShop.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_HopShop *HopShopSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _HopShop.Contract.Fallback(&_HopShop.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_HopShop *HopShopTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _HopShop.Contract.Fallback(&_HopShop.TransactOpts, calldata)
}

// HopShopTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the HopShop contract.
type HopShopTokensPurchasedIterator struct {
	Event *HopShopTokensPurchased // Event containing the contract specifics and raw log

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
func (it *HopShopTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HopShopTokensPurchased)
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
		it.Event = new(HopShopTokensPurchased)
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
func (it *HopShopTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HopShopTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HopShopTokensPurchased represents a TokensPurchased event raised by the HopShop contract.
type HopShopTokensPurchased struct {
	Purchaser common.Address
	Value     *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x8fafebcaf9d154343dad25669bfa277f4fbacd7ac6b0c4fed522580e040a0f33.
//
// Solidity: event TokensPurchased(address indexed purchaser, uint256 value, uint256 amount)
func (_HopShop *HopShopFilterer) FilterTokensPurchased(opts *bind.FilterOpts, purchaser []common.Address) (*HopShopTokensPurchasedIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}

	logs, sub, err := _HopShop.contract.FilterLogs(opts, "TokensPurchased", purchaserRule)
	if err != nil {
		return nil, err
	}
	return &HopShopTokensPurchasedIterator{contract: _HopShop.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x8fafebcaf9d154343dad25669bfa277f4fbacd7ac6b0c4fed522580e040a0f33.
//
// Solidity: event TokensPurchased(address indexed purchaser, uint256 value, uint256 amount)
func (_HopShop *HopShopFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *HopShopTokensPurchased, purchaser []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}

	logs, sub, err := _HopShop.contract.WatchLogs(opts, "TokensPurchased", purchaserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HopShopTokensPurchased)
				if err := _HopShop.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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

// ParseTokensPurchased is a log parse operation binding the contract event 0x8fafebcaf9d154343dad25669bfa277f4fbacd7ac6b0c4fed522580e040a0f33.
//
// Solidity: event TokensPurchased(address indexed purchaser, uint256 value, uint256 amount)
func (_HopShop *HopShopFilterer) ParseTokensPurchased(log types.Log) (*HopShopTokensPurchased, error) {
	event := new(HopShopTokensPurchased)
	if err := _HopShop.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
		return nil, err
	}
	return event, nil
}

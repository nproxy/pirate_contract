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

// ApplyTokenABI is the input ABI used to generate the binding from.
const ApplyTokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"p\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"TokenNoPerApplication\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"applyEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"applyFreeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPayer\",\"type\":\"address\"}],\"name\":\"changePayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ApplyToken is an auto generated Go binding around an Ethereum contract.
type ApplyToken struct {
	ApplyTokenCaller     // Read-only binding to the contract
	ApplyTokenTransactor // Write-only binding to the contract
	ApplyTokenFilterer   // Log filterer for contract events
}

// ApplyTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApplyTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApplyTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApplyTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApplyTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApplyTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApplyTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApplyTokenSession struct {
	Contract     *ApplyToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApplyTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApplyTokenCallerSession struct {
	Contract *ApplyTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ApplyTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApplyTokenTransactorSession struct {
	Contract     *ApplyTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ApplyTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApplyTokenRaw struct {
	Contract *ApplyToken // Generic contract binding to access the raw methods on
}

// ApplyTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApplyTokenCallerRaw struct {
	Contract *ApplyTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ApplyTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApplyTokenTransactorRaw struct {
	Contract *ApplyTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApplyToken creates a new instance of ApplyToken, bound to a specific deployed contract.
func NewApplyToken(address common.Address, backend bind.ContractBackend) (*ApplyToken, error) {
	contract, err := bindApplyToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ApplyToken{ApplyTokenCaller: ApplyTokenCaller{contract: contract}, ApplyTokenTransactor: ApplyTokenTransactor{contract: contract}, ApplyTokenFilterer: ApplyTokenFilterer{contract: contract}}, nil
}

// NewApplyTokenCaller creates a new read-only instance of ApplyToken, bound to a specific deployed contract.
func NewApplyTokenCaller(address common.Address, caller bind.ContractCaller) (*ApplyTokenCaller, error) {
	contract, err := bindApplyToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApplyTokenCaller{contract: contract}, nil
}

// NewApplyTokenTransactor creates a new write-only instance of ApplyToken, bound to a specific deployed contract.
func NewApplyTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ApplyTokenTransactor, error) {
	contract, err := bindApplyToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApplyTokenTransactor{contract: contract}, nil
}

// NewApplyTokenFilterer creates a new log filterer instance of ApplyToken, bound to a specific deployed contract.
func NewApplyTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ApplyTokenFilterer, error) {
	contract, err := bindApplyToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApplyTokenFilterer{contract: contract}, nil
}

// bindApplyToken binds a generic wrapper to an already deployed contract.
func bindApplyToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApplyTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ApplyToken *ApplyTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ApplyToken.Contract.ApplyTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ApplyToken *ApplyTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApplyToken.Contract.ApplyTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ApplyToken *ApplyTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ApplyToken.Contract.ApplyTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ApplyToken *ApplyTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ApplyToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ApplyToken *ApplyTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApplyToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ApplyToken *ApplyTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ApplyToken.Contract.contract.Transact(opts, method, params...)
}

// TokenNoPerApplication is a free data retrieval call binding the contract method 0x5f98d6f4.
//
// Solidity: function TokenNoPerApplication() view returns(uint256)
func (_ApplyToken *ApplyTokenCaller) TokenNoPerApplication(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ApplyToken.contract.Call(opts, &out, "TokenNoPerApplication")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenNoPerApplication is a free data retrieval call binding the contract method 0x5f98d6f4.
//
// Solidity: function TokenNoPerApplication() view returns(uint256)
func (_ApplyToken *ApplyTokenSession) TokenNoPerApplication() (*big.Int, error) {
	return _ApplyToken.Contract.TokenNoPerApplication(&_ApplyToken.CallOpts)
}

// TokenNoPerApplication is a free data retrieval call binding the contract method 0x5f98d6f4.
//
// Solidity: function TokenNoPerApplication() view returns(uint256)
func (_ApplyToken *ApplyTokenCallerSession) TokenNoPerApplication() (*big.Int, error) {
	return _ApplyToken.Contract.TokenNoPerApplication(&_ApplyToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ApplyToken *ApplyTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ApplyToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ApplyToken *ApplyTokenSession) Owner() (common.Address, error) {
	return _ApplyToken.Contract.Owner(&_ApplyToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ApplyToken *ApplyTokenCallerSession) Owner() (common.Address, error) {
	return _ApplyToken.Contract.Owner(&_ApplyToken.CallOpts)
}

// Payer is a free data retrieval call binding the contract method 0x123119cd.
//
// Solidity: function payer() view returns(address)
func (_ApplyToken *ApplyTokenCaller) Payer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ApplyToken.contract.Call(opts, &out, "payer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Payer is a free data retrieval call binding the contract method 0x123119cd.
//
// Solidity: function payer() view returns(address)
func (_ApplyToken *ApplyTokenSession) Payer() (common.Address, error) {
	return _ApplyToken.Contract.Payer(&_ApplyToken.CallOpts)
}

// Payer is a free data retrieval call binding the contract method 0x123119cd.
//
// Solidity: function payer() view returns(address)
func (_ApplyToken *ApplyTokenCallerSession) Payer() (common.Address, error) {
	return _ApplyToken.Contract.Payer(&_ApplyToken.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ApplyToken *ApplyTokenCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ApplyToken.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ApplyToken *ApplyTokenSession) Token() (common.Address, error) {
	return _ApplyToken.Contract.Token(&_ApplyToken.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ApplyToken *ApplyTokenCallerSession) Token() (common.Address, error) {
	return _ApplyToken.Contract.Token(&_ApplyToken.CallOpts)
}

// ApplyEth is a paid mutator transaction binding the contract method 0x83542bcb.
//
// Solidity: function applyEth(address userAddress) returns()
func (_ApplyToken *ApplyTokenTransactor) ApplyEth(opts *bind.TransactOpts, userAddress common.Address) (*types.Transaction, error) {
	return _ApplyToken.contract.Transact(opts, "applyEth", userAddress)
}

// ApplyEth is a paid mutator transaction binding the contract method 0x83542bcb.
//
// Solidity: function applyEth(address userAddress) returns()
func (_ApplyToken *ApplyTokenSession) ApplyEth(userAddress common.Address) (*types.Transaction, error) {
	return _ApplyToken.Contract.ApplyEth(&_ApplyToken.TransactOpts, userAddress)
}

// ApplyEth is a paid mutator transaction binding the contract method 0x83542bcb.
//
// Solidity: function applyEth(address userAddress) returns()
func (_ApplyToken *ApplyTokenTransactorSession) ApplyEth(userAddress common.Address) (*types.Transaction, error) {
	return _ApplyToken.Contract.ApplyEth(&_ApplyToken.TransactOpts, userAddress)
}

// ApplyFreeToken is a paid mutator transaction binding the contract method 0x1597a113.
//
// Solidity: function applyFreeToken(address userAddress) returns()
func (_ApplyToken *ApplyTokenTransactor) ApplyFreeToken(opts *bind.TransactOpts, userAddress common.Address) (*types.Transaction, error) {
	return _ApplyToken.contract.Transact(opts, "applyFreeToken", userAddress)
}

// ApplyFreeToken is a paid mutator transaction binding the contract method 0x1597a113.
//
// Solidity: function applyFreeToken(address userAddress) returns()
func (_ApplyToken *ApplyTokenSession) ApplyFreeToken(userAddress common.Address) (*types.Transaction, error) {
	return _ApplyToken.Contract.ApplyFreeToken(&_ApplyToken.TransactOpts, userAddress)
}

// ApplyFreeToken is a paid mutator transaction binding the contract method 0x1597a113.
//
// Solidity: function applyFreeToken(address userAddress) returns()
func (_ApplyToken *ApplyTokenTransactorSession) ApplyFreeToken(userAddress common.Address) (*types.Transaction, error) {
	return _ApplyToken.Contract.ApplyFreeToken(&_ApplyToken.TransactOpts, userAddress)
}

// ChangePayer is a paid mutator transaction binding the contract method 0xd6ea53b1.
//
// Solidity: function changePayer(address newPayer) returns()
func (_ApplyToken *ApplyTokenTransactor) ChangePayer(opts *bind.TransactOpts, newPayer common.Address) (*types.Transaction, error) {
	return _ApplyToken.contract.Transact(opts, "changePayer", newPayer)
}

// ChangePayer is a paid mutator transaction binding the contract method 0xd6ea53b1.
//
// Solidity: function changePayer(address newPayer) returns()
func (_ApplyToken *ApplyTokenSession) ChangePayer(newPayer common.Address) (*types.Transaction, error) {
	return _ApplyToken.Contract.ChangePayer(&_ApplyToken.TransactOpts, newPayer)
}

// ChangePayer is a paid mutator transaction binding the contract method 0xd6ea53b1.
//
// Solidity: function changePayer(address newPayer) returns()
func (_ApplyToken *ApplyTokenTransactorSession) ChangePayer(newPayer common.Address) (*types.Transaction, error) {
	return _ApplyToken.Contract.ChangePayer(&_ApplyToken.TransactOpts, newPayer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ApplyToken *ApplyTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ApplyToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ApplyToken *ApplyTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ApplyToken.Contract.TransferOwnership(&_ApplyToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ApplyToken *ApplyTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ApplyToken.Contract.TransferOwnership(&_ApplyToken.TransactOpts, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ApplyToken *ApplyTokenTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ApplyToken.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ApplyToken *ApplyTokenSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ApplyToken.Contract.Fallback(&_ApplyToken.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ApplyToken *ApplyTokenTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ApplyToken.Contract.Fallback(&_ApplyToken.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ApplyToken *ApplyTokenTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApplyToken.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ApplyToken *ApplyTokenSession) Receive() (*types.Transaction, error) {
	return _ApplyToken.Contract.Receive(&_ApplyToken.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ApplyToken *ApplyTokenTransactorSession) Receive() (*types.Transaction, error) {
	return _ApplyToken.Contract.Receive(&_ApplyToken.TransactOpts)
}

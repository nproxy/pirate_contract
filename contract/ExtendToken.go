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

// ExtendTokenABI is the input ABI used to generate the binding from.
const ExtendTokenABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimal\",\"type\":\"uint8\"},{\"internalType\":\"bytes8\",\"name\":\"symbol\",\"type\":\"bytes8\"}],\"name\":\"AddNewToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"AllPayContractList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"AllTokenData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentContract\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenI\",\"type\":\"address\"},{\"internalType\":\"bytes8\",\"name\":\"symbol\",\"type\":\"bytes8\"},{\"internalType\":\"uint256\",\"name\":\"decimal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"AllTokenList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"DelToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PayContractsList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"TokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"TokenBalances\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TokensList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ExtendToken is an auto generated Go binding around an Ethereum contract.
type ExtendToken struct {
	ExtendTokenCaller     // Read-only binding to the contract
	ExtendTokenTransactor // Write-only binding to the contract
	ExtendTokenFilterer   // Log filterer for contract events
}

// ExtendTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExtendTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtendTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExtendTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtendTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExtendTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtendTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExtendTokenSession struct {
	Contract     *ExtendToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExtendTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExtendTokenCallerSession struct {
	Contract *ExtendTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ExtendTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExtendTokenTransactorSession struct {
	Contract     *ExtendTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ExtendTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExtendTokenRaw struct {
	Contract *ExtendToken // Generic contract binding to access the raw methods on
}

// ExtendTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExtendTokenCallerRaw struct {
	Contract *ExtendTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ExtendTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExtendTokenTransactorRaw struct {
	Contract *ExtendTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExtendToken creates a new instance of ExtendToken, bound to a specific deployed contract.
func NewExtendToken(address common.Address, backend bind.ContractBackend) (*ExtendToken, error) {
	contract, err := bindExtendToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExtendToken{ExtendTokenCaller: ExtendTokenCaller{contract: contract}, ExtendTokenTransactor: ExtendTokenTransactor{contract: contract}, ExtendTokenFilterer: ExtendTokenFilterer{contract: contract}}, nil
}

// NewExtendTokenCaller creates a new read-only instance of ExtendToken, bound to a specific deployed contract.
func NewExtendTokenCaller(address common.Address, caller bind.ContractCaller) (*ExtendTokenCaller, error) {
	contract, err := bindExtendToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExtendTokenCaller{contract: contract}, nil
}

// NewExtendTokenTransactor creates a new write-only instance of ExtendToken, bound to a specific deployed contract.
func NewExtendTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ExtendTokenTransactor, error) {
	contract, err := bindExtendToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExtendTokenTransactor{contract: contract}, nil
}

// NewExtendTokenFilterer creates a new log filterer instance of ExtendToken, bound to a specific deployed contract.
func NewExtendTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ExtendTokenFilterer, error) {
	contract, err := bindExtendToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExtendTokenFilterer{contract: contract}, nil
}

// bindExtendToken binds a generic wrapper to an already deployed contract.
func bindExtendToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExtendTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtendToken *ExtendTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExtendToken.Contract.ExtendTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtendToken *ExtendTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtendToken.Contract.ExtendTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtendToken *ExtendTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtendToken.Contract.ExtendTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtendToken *ExtendTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExtendToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtendToken *ExtendTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtendToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtendToken *ExtendTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtendToken.Contract.contract.Transact(opts, method, params...)
}

// AllPayContractList is a free data retrieval call binding the contract method 0x0c59fdea.
//
// Solidity: function AllPayContractList(uint256 ) view returns(address)
func (_ExtendToken *ExtendTokenCaller) AllPayContractList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ExtendToken.contract.Call(opts, out, "AllPayContractList", arg0)
	return *ret0, err
}

// AllPayContractList is a free data retrieval call binding the contract method 0x0c59fdea.
//
// Solidity: function AllPayContractList(uint256 ) view returns(address)
func (_ExtendToken *ExtendTokenSession) AllPayContractList(arg0 *big.Int) (common.Address, error) {
	return _ExtendToken.Contract.AllPayContractList(&_ExtendToken.CallOpts, arg0)
}

// AllPayContractList is a free data retrieval call binding the contract method 0x0c59fdea.
//
// Solidity: function AllPayContractList(uint256 ) view returns(address)
func (_ExtendToken *ExtendTokenCallerSession) AllPayContractList(arg0 *big.Int) (common.Address, error) {
	return _ExtendToken.Contract.AllPayContractList(&_ExtendToken.CallOpts, arg0)
}

// AllTokenData is a free data retrieval call binding the contract method 0x31462773.
//
// Solidity: function AllTokenData(address ) view returns(address admin, address paymentContract, address tokenI, bytes8 symbol, uint256 decimal, uint256 index)
func (_ExtendToken *ExtendTokenCaller) AllTokenData(opts *bind.CallOpts, arg0 common.Address) (struct {
	Admin           common.Address
	PaymentContract common.Address
	TokenI          common.Address
	Symbol          [8]byte
	Decimal         *big.Int
	Index           *big.Int
}, error) {
	ret := new(struct {
		Admin           common.Address
		PaymentContract common.Address
		TokenI          common.Address
		Symbol          [8]byte
		Decimal         *big.Int
		Index           *big.Int
	})
	out := ret
	err := _ExtendToken.contract.Call(opts, out, "AllTokenData", arg0)
	return *ret, err
}

// AllTokenData is a free data retrieval call binding the contract method 0x31462773.
//
// Solidity: function AllTokenData(address ) view returns(address admin, address paymentContract, address tokenI, bytes8 symbol, uint256 decimal, uint256 index)
func (_ExtendToken *ExtendTokenSession) AllTokenData(arg0 common.Address) (struct {
	Admin           common.Address
	PaymentContract common.Address
	TokenI          common.Address
	Symbol          [8]byte
	Decimal         *big.Int
	Index           *big.Int
}, error) {
	return _ExtendToken.Contract.AllTokenData(&_ExtendToken.CallOpts, arg0)
}

// AllTokenData is a free data retrieval call binding the contract method 0x31462773.
//
// Solidity: function AllTokenData(address ) view returns(address admin, address paymentContract, address tokenI, bytes8 symbol, uint256 decimal, uint256 index)
func (_ExtendToken *ExtendTokenCallerSession) AllTokenData(arg0 common.Address) (struct {
	Admin           common.Address
	PaymentContract common.Address
	TokenI          common.Address
	Symbol          [8]byte
	Decimal         *big.Int
	Index           *big.Int
}, error) {
	return _ExtendToken.Contract.AllTokenData(&_ExtendToken.CallOpts, arg0)
}

// AllTokenList is a free data retrieval call binding the contract method 0x7a09fa61.
//
// Solidity: function AllTokenList(uint256 ) view returns(address)
func (_ExtendToken *ExtendTokenCaller) AllTokenList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ExtendToken.contract.Call(opts, out, "AllTokenList", arg0)
	return *ret0, err
}

// AllTokenList is a free data retrieval call binding the contract method 0x7a09fa61.
//
// Solidity: function AllTokenList(uint256 ) view returns(address)
func (_ExtendToken *ExtendTokenSession) AllTokenList(arg0 *big.Int) (common.Address, error) {
	return _ExtendToken.Contract.AllTokenList(&_ExtendToken.CallOpts, arg0)
}

// AllTokenList is a free data retrieval call binding the contract method 0x7a09fa61.
//
// Solidity: function AllTokenList(uint256 ) view returns(address)
func (_ExtendToken *ExtendTokenCallerSession) AllTokenList(arg0 *big.Int) (common.Address, error) {
	return _ExtendToken.Contract.AllTokenList(&_ExtendToken.CallOpts, arg0)
}

// PayContractsList is a free data retrieval call binding the contract method 0x1b191b1f.
//
// Solidity: function PayContractsList() view returns(address[])
func (_ExtendToken *ExtendTokenCaller) PayContractsList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ExtendToken.contract.Call(opts, out, "PayContractsList")
	return *ret0, err
}

// PayContractsList is a free data retrieval call binding the contract method 0x1b191b1f.
//
// Solidity: function PayContractsList() view returns(address[])
func (_ExtendToken *ExtendTokenSession) PayContractsList() ([]common.Address, error) {
	return _ExtendToken.Contract.PayContractsList(&_ExtendToken.CallOpts)
}

// PayContractsList is a free data retrieval call binding the contract method 0x1b191b1f.
//
// Solidity: function PayContractsList() view returns(address[])
func (_ExtendToken *ExtendTokenCallerSession) PayContractsList() ([]common.Address, error) {
	return _ExtendToken.Contract.PayContractsList(&_ExtendToken.CallOpts)
}

// TokenBalance is a free data retrieval call binding the contract method 0x47ea5004.
//
// Solidity: function TokenBalance(address userAddr, address tokenAddr) view returns(uint256)
func (_ExtendToken *ExtendTokenCaller) TokenBalance(opts *bind.CallOpts, userAddr common.Address, tokenAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExtendToken.contract.Call(opts, out, "TokenBalance", userAddr, tokenAddr)
	return *ret0, err
}

// TokenBalance is a free data retrieval call binding the contract method 0x47ea5004.
//
// Solidity: function TokenBalance(address userAddr, address tokenAddr) view returns(uint256)
func (_ExtendToken *ExtendTokenSession) TokenBalance(userAddr common.Address, tokenAddr common.Address) (*big.Int, error) {
	return _ExtendToken.Contract.TokenBalance(&_ExtendToken.CallOpts, userAddr, tokenAddr)
}

// TokenBalance is a free data retrieval call binding the contract method 0x47ea5004.
//
// Solidity: function TokenBalance(address userAddr, address tokenAddr) view returns(uint256)
func (_ExtendToken *ExtendTokenCallerSession) TokenBalance(userAddr common.Address, tokenAddr common.Address) (*big.Int, error) {
	return _ExtendToken.Contract.TokenBalance(&_ExtendToken.CallOpts, userAddr, tokenAddr)
}

// TokenBalances is a free data retrieval call binding the contract method 0xfead20ea.
//
// Solidity: function TokenBalances(address userAddr) view returns(address[], uint256[])
func (_ExtendToken *ExtendTokenCaller) TokenBalances(opts *bind.CallOpts, userAddr common.Address) ([]common.Address, []*big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ExtendToken.contract.Call(opts, out, "TokenBalances", userAddr)
	return *ret0, *ret1, err
}

// TokenBalances is a free data retrieval call binding the contract method 0xfead20ea.
//
// Solidity: function TokenBalances(address userAddr) view returns(address[], uint256[])
func (_ExtendToken *ExtendTokenSession) TokenBalances(userAddr common.Address) ([]common.Address, []*big.Int, error) {
	return _ExtendToken.Contract.TokenBalances(&_ExtendToken.CallOpts, userAddr)
}

// TokenBalances is a free data retrieval call binding the contract method 0xfead20ea.
//
// Solidity: function TokenBalances(address userAddr) view returns(address[], uint256[])
func (_ExtendToken *ExtendTokenCallerSession) TokenBalances(userAddr common.Address) ([]common.Address, []*big.Int, error) {
	return _ExtendToken.Contract.TokenBalances(&_ExtendToken.CallOpts, userAddr)
}

// TokensList is a free data retrieval call binding the contract method 0xf2a3c846.
//
// Solidity: function TokensList() view returns(address[])
func (_ExtendToken *ExtendTokenCaller) TokensList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ExtendToken.contract.Call(opts, out, "TokensList")
	return *ret0, err
}

// TokensList is a free data retrieval call binding the contract method 0xf2a3c846.
//
// Solidity: function TokensList() view returns(address[])
func (_ExtendToken *ExtendTokenSession) TokensList() ([]common.Address, error) {
	return _ExtendToken.Contract.TokensList(&_ExtendToken.CallOpts)
}

// TokensList is a free data retrieval call binding the contract method 0xf2a3c846.
//
// Solidity: function TokensList() view returns(address[])
func (_ExtendToken *ExtendTokenCallerSession) TokensList() ([]common.Address, error) {
	return _ExtendToken.Contract.TokensList(&_ExtendToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExtendToken *ExtendTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ExtendToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExtendToken *ExtendTokenSession) Owner() (common.Address, error) {
	return _ExtendToken.Contract.Owner(&_ExtendToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExtendToken *ExtendTokenCallerSession) Owner() (common.Address, error) {
	return _ExtendToken.Contract.Owner(&_ExtendToken.CallOpts)
}

// AddNewToken is a paid mutator transaction binding the contract method 0xe94af2b8.
//
// Solidity: function AddNewToken(address admin, address payContract, address tokenAddr, uint8 decimal, bytes8 symbol) returns()
func (_ExtendToken *ExtendTokenTransactor) AddNewToken(opts *bind.TransactOpts, admin common.Address, payContract common.Address, tokenAddr common.Address, decimal uint8, symbol [8]byte) (*types.Transaction, error) {
	return _ExtendToken.contract.Transact(opts, "AddNewToken", admin, payContract, tokenAddr, decimal, symbol)
}

// AddNewToken is a paid mutator transaction binding the contract method 0xe94af2b8.
//
// Solidity: function AddNewToken(address admin, address payContract, address tokenAddr, uint8 decimal, bytes8 symbol) returns()
func (_ExtendToken *ExtendTokenSession) AddNewToken(admin common.Address, payContract common.Address, tokenAddr common.Address, decimal uint8, symbol [8]byte) (*types.Transaction, error) {
	return _ExtendToken.Contract.AddNewToken(&_ExtendToken.TransactOpts, admin, payContract, tokenAddr, decimal, symbol)
}

// AddNewToken is a paid mutator transaction binding the contract method 0xe94af2b8.
//
// Solidity: function AddNewToken(address admin, address payContract, address tokenAddr, uint8 decimal, bytes8 symbol) returns()
func (_ExtendToken *ExtendTokenTransactorSession) AddNewToken(admin common.Address, payContract common.Address, tokenAddr common.Address, decimal uint8, symbol [8]byte) (*types.Transaction, error) {
	return _ExtendToken.Contract.AddNewToken(&_ExtendToken.TransactOpts, admin, payContract, tokenAddr, decimal, symbol)
}

// DelToken is a paid mutator transaction binding the contract method 0x842a8bf6.
//
// Solidity: function DelToken(address tokenAddr) returns()
func (_ExtendToken *ExtendTokenTransactor) DelToken(opts *bind.TransactOpts, tokenAddr common.Address) (*types.Transaction, error) {
	return _ExtendToken.contract.Transact(opts, "DelToken", tokenAddr)
}

// DelToken is a paid mutator transaction binding the contract method 0x842a8bf6.
//
// Solidity: function DelToken(address tokenAddr) returns()
func (_ExtendToken *ExtendTokenSession) DelToken(tokenAddr common.Address) (*types.Transaction, error) {
	return _ExtendToken.Contract.DelToken(&_ExtendToken.TransactOpts, tokenAddr)
}

// DelToken is a paid mutator transaction binding the contract method 0x842a8bf6.
//
// Solidity: function DelToken(address tokenAddr) returns()
func (_ExtendToken *ExtendTokenTransactorSession) DelToken(tokenAddr common.Address) (*types.Transaction, error) {
	return _ExtendToken.Contract.DelToken(&_ExtendToken.TransactOpts, tokenAddr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExtendToken *ExtendTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ExtendToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExtendToken *ExtendTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExtendToken.Contract.TransferOwnership(&_ExtendToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExtendToken *ExtendTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExtendToken.Contract.TransferOwnership(&_ExtendToken.TransactOpts, newOwner)
}

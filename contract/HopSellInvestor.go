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

// HopSellInvestorABI is the input ABI used to generate the binding from.
const HopSellInvestorABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"HOPAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"USDTAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"HOPAmount\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"HOPAmount\",\"type\":\"uint256\"}],\"name\":\"Exchange\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"EmergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HOP\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HOP_FUND\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDT\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"addBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceDetail\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mutiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimable\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fundAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneAddr\",\"type\":\"address\"}],\"name\":\"changeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimHOP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"inputUSDT\",\"type\":\"uint256\"}],\"name\":\"exchangeForHOP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"inputUSDT\",\"type\":\"uint256\"}],\"name\":\"receiveUSDT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"releaseFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"setRateFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLocked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HopSellInvestor is an auto generated Go binding around an Ethereum contract.
type HopSellInvestor struct {
	HopSellInvestorCaller     // Read-only binding to the contract
	HopSellInvestorTransactor // Write-only binding to the contract
	HopSellInvestorFilterer   // Log filterer for contract events
}

// HopSellInvestorCaller is an auto generated read-only Go binding around an Ethereum contract.
type HopSellInvestorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HopSellInvestorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HopSellInvestorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HopSellInvestorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HopSellInvestorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HopSellInvestorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HopSellInvestorSession struct {
	Contract     *HopSellInvestor  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HopSellInvestorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HopSellInvestorCallerSession struct {
	Contract *HopSellInvestorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// HopSellInvestorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HopSellInvestorTransactorSession struct {
	Contract     *HopSellInvestorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// HopSellInvestorRaw is an auto generated low-level Go binding around an Ethereum contract.
type HopSellInvestorRaw struct {
	Contract *HopSellInvestor // Generic contract binding to access the raw methods on
}

// HopSellInvestorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HopSellInvestorCallerRaw struct {
	Contract *HopSellInvestorCaller // Generic read-only contract binding to access the raw methods on
}

// HopSellInvestorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HopSellInvestorTransactorRaw struct {
	Contract *HopSellInvestorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHopSellInvestor creates a new instance of HopSellInvestor, bound to a specific deployed contract.
func NewHopSellInvestor(address common.Address, backend bind.ContractBackend) (*HopSellInvestor, error) {
	contract, err := bindHopSellInvestor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HopSellInvestor{HopSellInvestorCaller: HopSellInvestorCaller{contract: contract}, HopSellInvestorTransactor: HopSellInvestorTransactor{contract: contract}, HopSellInvestorFilterer: HopSellInvestorFilterer{contract: contract}}, nil
}

// NewHopSellInvestorCaller creates a new read-only instance of HopSellInvestor, bound to a specific deployed contract.
func NewHopSellInvestorCaller(address common.Address, caller bind.ContractCaller) (*HopSellInvestorCaller, error) {
	contract, err := bindHopSellInvestor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HopSellInvestorCaller{contract: contract}, nil
}

// NewHopSellInvestorTransactor creates a new write-only instance of HopSellInvestor, bound to a specific deployed contract.
func NewHopSellInvestorTransactor(address common.Address, transactor bind.ContractTransactor) (*HopSellInvestorTransactor, error) {
	contract, err := bindHopSellInvestor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HopSellInvestorTransactor{contract: contract}, nil
}

// NewHopSellInvestorFilterer creates a new log filterer instance of HopSellInvestor, bound to a specific deployed contract.
func NewHopSellInvestorFilterer(address common.Address, filterer bind.ContractFilterer) (*HopSellInvestorFilterer, error) {
	contract, err := bindHopSellInvestor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HopSellInvestorFilterer{contract: contract}, nil
}

// bindHopSellInvestor binds a generic wrapper to an already deployed contract.
func bindHopSellInvestor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HopSellInvestorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HopSellInvestor *HopSellInvestorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HopSellInvestor.Contract.HopSellInvestorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HopSellInvestor *HopSellInvestorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.HopSellInvestorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HopSellInvestor *HopSellInvestorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.HopSellInvestorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HopSellInvestor *HopSellInvestorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HopSellInvestor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HopSellInvestor *HopSellInvestorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HopSellInvestor *HopSellInvestorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.contract.Transact(opts, method, params...)
}

// HOP is a free data retrieval call binding the contract method 0xe85a1b0a.
//
// Solidity: function HOP() view returns(address)
func (_HopSellInvestor *HopSellInvestorCaller) HOP(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HopSellInvestor.contract.Call(opts, &out, "HOP")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HOP is a free data retrieval call binding the contract method 0xe85a1b0a.
//
// Solidity: function HOP() view returns(address)
func (_HopSellInvestor *HopSellInvestorSession) HOP() (common.Address, error) {
	return _HopSellInvestor.Contract.HOP(&_HopSellInvestor.CallOpts)
}

// HOP is a free data retrieval call binding the contract method 0xe85a1b0a.
//
// Solidity: function HOP() view returns(address)
func (_HopSellInvestor *HopSellInvestorCallerSession) HOP() (common.Address, error) {
	return _HopSellInvestor.Contract.HOP(&_HopSellInvestor.CallOpts)
}

// HOPFUND is a free data retrieval call binding the contract method 0x6078fbd6.
//
// Solidity: function HOP_FUND() view returns(address)
func (_HopSellInvestor *HopSellInvestorCaller) HOPFUND(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HopSellInvestor.contract.Call(opts, &out, "HOP_FUND")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HOPFUND is a free data retrieval call binding the contract method 0x6078fbd6.
//
// Solidity: function HOP_FUND() view returns(address)
func (_HopSellInvestor *HopSellInvestorSession) HOPFUND() (common.Address, error) {
	return _HopSellInvestor.Contract.HOPFUND(&_HopSellInvestor.CallOpts)
}

// HOPFUND is a free data retrieval call binding the contract method 0x6078fbd6.
//
// Solidity: function HOP_FUND() view returns(address)
func (_HopSellInvestor *HopSellInvestorCallerSession) HOPFUND() (common.Address, error) {
	return _HopSellInvestor.Contract.HOPFUND(&_HopSellInvestor.CallOpts)
}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_HopSellInvestor *HopSellInvestorCaller) USDT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HopSellInvestor.contract.Call(opts, &out, "USDT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_HopSellInvestor *HopSellInvestorSession) USDT() (common.Address, error) {
	return _HopSellInvestor.Contract.USDT(&_HopSellInvestor.CallOpts)
}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_HopSellInvestor *HopSellInvestorCallerSession) USDT() (common.Address, error) {
	return _HopSellInvestor.Contract.USDT(&_HopSellInvestor.CallOpts)
}

// BalanceDetail is a free data retrieval call binding the contract method 0x3104e7ce.
//
// Solidity: function balanceDetail(address ) view returns(uint256 balance, uint256 mutiplier, uint256 claimable)
func (_HopSellInvestor *HopSellInvestorCaller) BalanceDetail(opts *bind.CallOpts, arg0 common.Address) (struct {
	Balance   *big.Int
	Mutiplier *big.Int
	Claimable *big.Int
}, error) {
	var out []interface{}
	err := _HopSellInvestor.contract.Call(opts, &out, "balanceDetail", arg0)

	outstruct := new(struct {
		Balance   *big.Int
		Mutiplier *big.Int
		Claimable *big.Int
	})

	outstruct.Balance = out[0].(*big.Int)
	outstruct.Mutiplier = out[1].(*big.Int)
	outstruct.Claimable = out[2].(*big.Int)

	return *outstruct, err

}

// BalanceDetail is a free data retrieval call binding the contract method 0x3104e7ce.
//
// Solidity: function balanceDetail(address ) view returns(uint256 balance, uint256 mutiplier, uint256 claimable)
func (_HopSellInvestor *HopSellInvestorSession) BalanceDetail(arg0 common.Address) (struct {
	Balance   *big.Int
	Mutiplier *big.Int
	Claimable *big.Int
}, error) {
	return _HopSellInvestor.Contract.BalanceDetail(&_HopSellInvestor.CallOpts, arg0)
}

// BalanceDetail is a free data retrieval call binding the contract method 0x3104e7ce.
//
// Solidity: function balanceDetail(address ) view returns(uint256 balance, uint256 mutiplier, uint256 claimable)
func (_HopSellInvestor *HopSellInvestorCallerSession) BalanceDetail(arg0 common.Address) (struct {
	Balance   *big.Int
	Mutiplier *big.Int
	Claimable *big.Int
}, error) {
	return _HopSellInvestor.Contract.BalanceDetail(&_HopSellInvestor.CallOpts, arg0)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_HopSellInvestor *HopSellInvestorCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HopSellInvestor.contract.Call(opts, &out, "beneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_HopSellInvestor *HopSellInvestorSession) Beneficiary() (common.Address, error) {
	return _HopSellInvestor.Contract.Beneficiary(&_HopSellInvestor.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_HopSellInvestor *HopSellInvestorCallerSession) Beneficiary() (common.Address, error) {
	return _HopSellInvestor.Contract.Beneficiary(&_HopSellInvestor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HopSellInvestor *HopSellInvestorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HopSellInvestor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HopSellInvestor *HopSellInvestorSession) Owner() (common.Address, error) {
	return _HopSellInvestor.Contract.Owner(&_HopSellInvestor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HopSellInvestor *HopSellInvestorCallerSession) Owner() (common.Address, error) {
	return _HopSellInvestor.Contract.Owner(&_HopSellInvestor.CallOpts)
}

// TotalLocked is a free data retrieval call binding the contract method 0x56891412.
//
// Solidity: function totalLocked() view returns(uint256)
func (_HopSellInvestor *HopSellInvestorCaller) TotalLocked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HopSellInvestor.contract.Call(opts, &out, "totalLocked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLocked is a free data retrieval call binding the contract method 0x56891412.
//
// Solidity: function totalLocked() view returns(uint256)
func (_HopSellInvestor *HopSellInvestorSession) TotalLocked() (*big.Int, error) {
	return _HopSellInvestor.Contract.TotalLocked(&_HopSellInvestor.CallOpts)
}

// TotalLocked is a free data retrieval call binding the contract method 0x56891412.
//
// Solidity: function totalLocked() view returns(uint256)
func (_HopSellInvestor *HopSellInvestorCallerSession) TotalLocked() (*big.Int, error) {
	return _HopSellInvestor.Contract.TotalLocked(&_HopSellInvestor.CallOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xbb43d2e6.
//
// Solidity: function EmergencyWithdraw(uint256 amount, address to) returns()
func (_HopSellInvestor *HopSellInvestorTransactor) EmergencyWithdraw(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _HopSellInvestor.contract.Transact(opts, "EmergencyWithdraw", amount, to)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xbb43d2e6.
//
// Solidity: function EmergencyWithdraw(uint256 amount, address to) returns()
func (_HopSellInvestor *HopSellInvestorSession) EmergencyWithdraw(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.EmergencyWithdraw(&_HopSellInvestor.TransactOpts, amount, to)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xbb43d2e6.
//
// Solidity: function EmergencyWithdraw(uint256 amount, address to) returns()
func (_HopSellInvestor *HopSellInvestorTransactorSession) EmergencyWithdraw(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.EmergencyWithdraw(&_HopSellInvestor.TransactOpts, amount, to)
}

// AddBalance is a paid mutator transaction binding the contract method 0x2415aeeb.
//
// Solidity: function addBalance(address[] addrs, uint256[] values) returns()
func (_HopSellInvestor *HopSellInvestorTransactor) AddBalance(opts *bind.TransactOpts, addrs []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _HopSellInvestor.contract.Transact(opts, "addBalance", addrs, values)
}

// AddBalance is a paid mutator transaction binding the contract method 0x2415aeeb.
//
// Solidity: function addBalance(address[] addrs, uint256[] values) returns()
func (_HopSellInvestor *HopSellInvestorSession) AddBalance(addrs []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.AddBalance(&_HopSellInvestor.TransactOpts, addrs, values)
}

// AddBalance is a paid mutator transaction binding the contract method 0x2415aeeb.
//
// Solidity: function addBalance(address[] addrs, uint256[] values) returns()
func (_HopSellInvestor *HopSellInvestorTransactorSession) AddBalance(addrs []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.AddBalance(&_HopSellInvestor.TransactOpts, addrs, values)
}

// ChangeAddress is a paid mutator transaction binding the contract method 0xefe08a7d.
//
// Solidity: function changeAddress(address fundAddr, address beneAddr) returns()
func (_HopSellInvestor *HopSellInvestorTransactor) ChangeAddress(opts *bind.TransactOpts, fundAddr common.Address, beneAddr common.Address) (*types.Transaction, error) {
	return _HopSellInvestor.contract.Transact(opts, "changeAddress", fundAddr, beneAddr)
}

// ChangeAddress is a paid mutator transaction binding the contract method 0xefe08a7d.
//
// Solidity: function changeAddress(address fundAddr, address beneAddr) returns()
func (_HopSellInvestor *HopSellInvestorSession) ChangeAddress(fundAddr common.Address, beneAddr common.Address) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.ChangeAddress(&_HopSellInvestor.TransactOpts, fundAddr, beneAddr)
}

// ChangeAddress is a paid mutator transaction binding the contract method 0xefe08a7d.
//
// Solidity: function changeAddress(address fundAddr, address beneAddr) returns()
func (_HopSellInvestor *HopSellInvestorTransactorSession) ChangeAddress(fundAddr common.Address, beneAddr common.Address) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.ChangeAddress(&_HopSellInvestor.TransactOpts, fundAddr, beneAddr)
}

// ClaimHOP is a paid mutator transaction binding the contract method 0x51ee5ae7.
//
// Solidity: function claimHOP(uint256 amount) returns()
func (_HopSellInvestor *HopSellInvestorTransactor) ClaimHOP(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _HopSellInvestor.contract.Transact(opts, "claimHOP", amount)
}

// ClaimHOP is a paid mutator transaction binding the contract method 0x51ee5ae7.
//
// Solidity: function claimHOP(uint256 amount) returns()
func (_HopSellInvestor *HopSellInvestorSession) ClaimHOP(amount *big.Int) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.ClaimHOP(&_HopSellInvestor.TransactOpts, amount)
}

// ClaimHOP is a paid mutator transaction binding the contract method 0x51ee5ae7.
//
// Solidity: function claimHOP(uint256 amount) returns()
func (_HopSellInvestor *HopSellInvestorTransactorSession) ClaimHOP(amount *big.Int) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.ClaimHOP(&_HopSellInvestor.TransactOpts, amount)
}

// ExchangeForHOP is a paid mutator transaction binding the contract method 0x37ed8096.
//
// Solidity: function exchangeForHOP(uint256 inputUSDT) returns()
func (_HopSellInvestor *HopSellInvestorTransactor) ExchangeForHOP(opts *bind.TransactOpts, inputUSDT *big.Int) (*types.Transaction, error) {
	return _HopSellInvestor.contract.Transact(opts, "exchangeForHOP", inputUSDT)
}

// ExchangeForHOP is a paid mutator transaction binding the contract method 0x37ed8096.
//
// Solidity: function exchangeForHOP(uint256 inputUSDT) returns()
func (_HopSellInvestor *HopSellInvestorSession) ExchangeForHOP(inputUSDT *big.Int) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.ExchangeForHOP(&_HopSellInvestor.TransactOpts, inputUSDT)
}

// ExchangeForHOP is a paid mutator transaction binding the contract method 0x37ed8096.
//
// Solidity: function exchangeForHOP(uint256 inputUSDT) returns()
func (_HopSellInvestor *HopSellInvestorTransactorSession) ExchangeForHOP(inputUSDT *big.Int) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.ExchangeForHOP(&_HopSellInvestor.TransactOpts, inputUSDT)
}

// FundWithdraw is a paid mutator transaction binding the contract method 0x4ab53c24.
//
// Solidity: function fundWithdraw(uint256 amount) returns()
func (_HopSellInvestor *HopSellInvestorTransactor) FundWithdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _HopSellInvestor.contract.Transact(opts, "fundWithdraw", amount)
}

// FundWithdraw is a paid mutator transaction binding the contract method 0x4ab53c24.
//
// Solidity: function fundWithdraw(uint256 amount) returns()
func (_HopSellInvestor *HopSellInvestorSession) FundWithdraw(amount *big.Int) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.FundWithdraw(&_HopSellInvestor.TransactOpts, amount)
}

// FundWithdraw is a paid mutator transaction binding the contract method 0x4ab53c24.
//
// Solidity: function fundWithdraw(uint256 amount) returns()
func (_HopSellInvestor *HopSellInvestorTransactorSession) FundWithdraw(amount *big.Int) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.FundWithdraw(&_HopSellInvestor.TransactOpts, amount)
}

// ReceiveUSDT is a paid mutator transaction binding the contract method 0x38d51da8.
//
// Solidity: function receiveUSDT(uint256 inputUSDT) returns()
func (_HopSellInvestor *HopSellInvestorTransactor) ReceiveUSDT(opts *bind.TransactOpts, inputUSDT *big.Int) (*types.Transaction, error) {
	return _HopSellInvestor.contract.Transact(opts, "receiveUSDT", inputUSDT)
}

// ReceiveUSDT is a paid mutator transaction binding the contract method 0x38d51da8.
//
// Solidity: function receiveUSDT(uint256 inputUSDT) returns()
func (_HopSellInvestor *HopSellInvestorSession) ReceiveUSDT(inputUSDT *big.Int) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.ReceiveUSDT(&_HopSellInvestor.TransactOpts, inputUSDT)
}

// ReceiveUSDT is a paid mutator transaction binding the contract method 0x38d51da8.
//
// Solidity: function receiveUSDT(uint256 inputUSDT) returns()
func (_HopSellInvestor *HopSellInvestorTransactorSession) ReceiveUSDT(inputUSDT *big.Int) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.ReceiveUSDT(&_HopSellInvestor.TransactOpts, inputUSDT)
}

// ReleaseFor is a paid mutator transaction binding the contract method 0xbeb96be5.
//
// Solidity: function releaseFor(address account, uint256 amount) returns()
func (_HopSellInvestor *HopSellInvestorTransactor) ReleaseFor(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HopSellInvestor.contract.Transact(opts, "releaseFor", account, amount)
}

// ReleaseFor is a paid mutator transaction binding the contract method 0xbeb96be5.
//
// Solidity: function releaseFor(address account, uint256 amount) returns()
func (_HopSellInvestor *HopSellInvestorSession) ReleaseFor(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.ReleaseFor(&_HopSellInvestor.TransactOpts, account, amount)
}

// ReleaseFor is a paid mutator transaction binding the contract method 0xbeb96be5.
//
// Solidity: function releaseFor(address account, uint256 amount) returns()
func (_HopSellInvestor *HopSellInvestorTransactorSession) ReleaseFor(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.ReleaseFor(&_HopSellInvestor.TransactOpts, account, amount)
}

// SetRateFor is a paid mutator transaction binding the contract method 0x29dc6e34.
//
// Solidity: function setRateFor(address account, uint256 r) returns()
func (_HopSellInvestor *HopSellInvestorTransactor) SetRateFor(opts *bind.TransactOpts, account common.Address, r *big.Int) (*types.Transaction, error) {
	return _HopSellInvestor.contract.Transact(opts, "setRateFor", account, r)
}

// SetRateFor is a paid mutator transaction binding the contract method 0x29dc6e34.
//
// Solidity: function setRateFor(address account, uint256 r) returns()
func (_HopSellInvestor *HopSellInvestorSession) SetRateFor(account common.Address, r *big.Int) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.SetRateFor(&_HopSellInvestor.TransactOpts, account, r)
}

// SetRateFor is a paid mutator transaction binding the contract method 0x29dc6e34.
//
// Solidity: function setRateFor(address account, uint256 r) returns()
func (_HopSellInvestor *HopSellInvestorTransactorSession) SetRateFor(account common.Address, r *big.Int) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.SetRateFor(&_HopSellInvestor.TransactOpts, account, r)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HopSellInvestor *HopSellInvestorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HopSellInvestor.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HopSellInvestor *HopSellInvestorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.TransferOwnership(&_HopSellInvestor.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HopSellInvestor *HopSellInvestorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HopSellInvestor.Contract.TransferOwnership(&_HopSellInvestor.TransactOpts, newOwner)
}

// HopSellInvestorClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the HopSellInvestor contract.
type HopSellInvestorClaimIterator struct {
	Event *HopSellInvestorClaim // Event containing the contract specifics and raw log

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
func (it *HopSellInvestorClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HopSellInvestorClaim)
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
		it.Event = new(HopSellInvestorClaim)
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
func (it *HopSellInvestorClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HopSellInvestorClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HopSellInvestorClaim represents a Claim event raised by the HopSellInvestor contract.
type HopSellInvestorClaim struct {
	User      common.Address
	HOPAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address user, uint256 HOPAmount)
func (_HopSellInvestor *HopSellInvestorFilterer) FilterClaim(opts *bind.FilterOpts) (*HopSellInvestorClaimIterator, error) {

	logs, sub, err := _HopSellInvestor.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &HopSellInvestorClaimIterator{contract: _HopSellInvestor.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address user, uint256 HOPAmount)
func (_HopSellInvestor *HopSellInvestorFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *HopSellInvestorClaim) (event.Subscription, error) {

	logs, sub, err := _HopSellInvestor.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HopSellInvestorClaim)
				if err := _HopSellInvestor.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address user, uint256 HOPAmount)
func (_HopSellInvestor *HopSellInvestorFilterer) ParseClaim(log types.Log) (*HopSellInvestorClaim, error) {
	event := new(HopSellInvestorClaim)
	if err := _HopSellInvestor.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HopSellInvestorExchangeIterator is returned from FilterExchange and is used to iterate over the raw logs and unpacked data for Exchange events raised by the HopSellInvestor contract.
type HopSellInvestorExchangeIterator struct {
	Event *HopSellInvestorExchange // Event containing the contract specifics and raw log

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
func (it *HopSellInvestorExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HopSellInvestorExchange)
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
		it.Event = new(HopSellInvestorExchange)
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
func (it *HopSellInvestorExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HopSellInvestorExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HopSellInvestorExchange represents a Exchange event raised by the HopSellInvestor contract.
type HopSellInvestorExchange struct {
	User      common.Address
	HOPAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExchange is a free log retrieval operation binding the contract event 0x5988e4c12f4844b895de0739f562558435dca9602fd8b970720ee3cf8dff39be.
//
// Solidity: event Exchange(address indexed user, uint256 HOPAmount)
func (_HopSellInvestor *HopSellInvestorFilterer) FilterExchange(opts *bind.FilterOpts, user []common.Address) (*HopSellInvestorExchangeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _HopSellInvestor.contract.FilterLogs(opts, "Exchange", userRule)
	if err != nil {
		return nil, err
	}
	return &HopSellInvestorExchangeIterator{contract: _HopSellInvestor.contract, event: "Exchange", logs: logs, sub: sub}, nil
}

// WatchExchange is a free log subscription operation binding the contract event 0x5988e4c12f4844b895de0739f562558435dca9602fd8b970720ee3cf8dff39be.
//
// Solidity: event Exchange(address indexed user, uint256 HOPAmount)
func (_HopSellInvestor *HopSellInvestorFilterer) WatchExchange(opts *bind.WatchOpts, sink chan<- *HopSellInvestorExchange, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _HopSellInvestor.contract.WatchLogs(opts, "Exchange", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HopSellInvestorExchange)
				if err := _HopSellInvestor.contract.UnpackLog(event, "Exchange", log); err != nil {
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

// ParseExchange is a log parse operation binding the contract event 0x5988e4c12f4844b895de0739f562558435dca9602fd8b970720ee3cf8dff39be.
//
// Solidity: event Exchange(address indexed user, uint256 HOPAmount)
func (_HopSellInvestor *HopSellInvestorFilterer) ParseExchange(log types.Log) (*HopSellInvestorExchange, error) {
	event := new(HopSellInvestorExchange)
	if err := _HopSellInvestor.contract.UnpackLog(event, "Exchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

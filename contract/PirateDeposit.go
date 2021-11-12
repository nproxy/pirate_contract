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

// PirateDepositABI is the input ABI used to generate the binding from.
const PirateDepositABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"m\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardTime\",\"type\":\"uint256\"}],\"name\":\"AddRewardEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profits\",\"type\":\"uint256\"}],\"name\":\"PoolDrawRewardEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositTime\",\"type\":\"uint256\"}],\"name\":\"UserDepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastIndex\",\"type\":\"uint256\"}],\"name\":\"UserDrawRewardEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositTime\",\"type\":\"uint256\"}],\"name\":\"UserWithDrawDepositEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Decimal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"DepositDatas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastRateIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"DrawRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"drawRateTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftReward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"drawTime\",\"type\":\"uint256\"}],\"name\":\"addReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coorD\",\"type\":\"address\"}],\"name\":\"changeCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"changeDepositInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"}],\"name\":\"changeMinDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"changePoolDrawIntval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"changeRateDecimal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"changeRewardInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"coordinator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"doUserDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"poolIndex\",\"type\":\"uint256\"}],\"name\":\"drawReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"market\",\"outputs\":[{\"internalType\":\"contractTrafficMarket\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDepositInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minOneMonth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolDrawInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolIndex\",\"type\":\"uint256\"}],\"name\":\"poolDrawReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rateDecimal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"recordTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolIndex\",\"type\":\"uint256\"}],\"name\":\"userWithDrawDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PirateDeposit is an auto generated Go binding around an Ethereum contract.
type PirateDeposit struct {
	PirateDepositCaller     // Read-only binding to the contract
	PirateDepositTransactor // Write-only binding to the contract
	PirateDepositFilterer   // Log filterer for contract events
}

// PirateDepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type PirateDepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PirateDepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PirateDepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PirateDepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PirateDepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PirateDepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PirateDepositSession struct {
	Contract     *PirateDeposit    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PirateDepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PirateDepositCallerSession struct {
	Contract *PirateDepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PirateDepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PirateDepositTransactorSession struct {
	Contract     *PirateDepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PirateDepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type PirateDepositRaw struct {
	Contract *PirateDeposit // Generic contract binding to access the raw methods on
}

// PirateDepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PirateDepositCallerRaw struct {
	Contract *PirateDepositCaller // Generic read-only contract binding to access the raw methods on
}

// PirateDepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PirateDepositTransactorRaw struct {
	Contract *PirateDepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPirateDeposit creates a new instance of PirateDeposit, bound to a specific deployed contract.
func NewPirateDeposit(address common.Address, backend bind.ContractBackend) (*PirateDeposit, error) {
	contract, err := bindPirateDeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PirateDeposit{PirateDepositCaller: PirateDepositCaller{contract: contract}, PirateDepositTransactor: PirateDepositTransactor{contract: contract}, PirateDepositFilterer: PirateDepositFilterer{contract: contract}}, nil
}

// NewPirateDepositCaller creates a new read-only instance of PirateDeposit, bound to a specific deployed contract.
func NewPirateDepositCaller(address common.Address, caller bind.ContractCaller) (*PirateDepositCaller, error) {
	contract, err := bindPirateDeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PirateDepositCaller{contract: contract}, nil
}

// NewPirateDepositTransactor creates a new write-only instance of PirateDeposit, bound to a specific deployed contract.
func NewPirateDepositTransactor(address common.Address, transactor bind.ContractTransactor) (*PirateDepositTransactor, error) {
	contract, err := bindPirateDeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PirateDepositTransactor{contract: contract}, nil
}

// NewPirateDepositFilterer creates a new log filterer instance of PirateDeposit, bound to a specific deployed contract.
func NewPirateDepositFilterer(address common.Address, filterer bind.ContractFilterer) (*PirateDepositFilterer, error) {
	contract, err := bindPirateDeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PirateDepositFilterer{contract: contract}, nil
}

// bindPirateDeposit binds a generic wrapper to an already deployed contract.
func bindPirateDeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PirateDepositABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PirateDeposit *PirateDepositRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PirateDeposit.Contract.PirateDepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PirateDeposit *PirateDepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PirateDeposit.Contract.PirateDepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PirateDeposit *PirateDepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PirateDeposit.Contract.PirateDepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PirateDeposit *PirateDepositCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PirateDeposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PirateDeposit *PirateDepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PirateDeposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PirateDeposit *PirateDepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PirateDeposit.Contract.contract.Transact(opts, method, params...)
}

// Decimal is a free data retrieval call binding the contract method 0x1fe50c39.
//
// Solidity: function Decimal() view returns(uint256)
func (_PirateDeposit *PirateDepositCaller) Decimal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PirateDeposit.contract.Call(opts, &out, "Decimal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimal is a free data retrieval call binding the contract method 0x1fe50c39.
//
// Solidity: function Decimal() view returns(uint256)
func (_PirateDeposit *PirateDepositSession) Decimal() (*big.Int, error) {
	return _PirateDeposit.Contract.Decimal(&_PirateDeposit.CallOpts)
}

// Decimal is a free data retrieval call binding the contract method 0x1fe50c39.
//
// Solidity: function Decimal() view returns(uint256)
func (_PirateDeposit *PirateDepositCallerSession) Decimal() (*big.Int, error) {
	return _PirateDeposit.Contract.Decimal(&_PirateDeposit.CallOpts)
}

// DepositDatas is a free data retrieval call binding the contract method 0xe1bdb727.
//
// Solidity: function DepositDatas(address , address ) view returns(uint256 lastRateIndex)
func (_PirateDeposit *PirateDepositCaller) DepositDatas(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PirateDeposit.contract.Call(opts, &out, "DepositDatas", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositDatas is a free data retrieval call binding the contract method 0xe1bdb727.
//
// Solidity: function DepositDatas(address , address ) view returns(uint256 lastRateIndex)
func (_PirateDeposit *PirateDepositSession) DepositDatas(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _PirateDeposit.Contract.DepositDatas(&_PirateDeposit.CallOpts, arg0, arg1)
}

// DepositDatas is a free data retrieval call binding the contract method 0xe1bdb727.
//
// Solidity: function DepositDatas(address , address ) view returns(uint256 lastRateIndex)
func (_PirateDeposit *PirateDepositCallerSession) DepositDatas(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _PirateDeposit.Contract.DepositDatas(&_PirateDeposit.CallOpts, arg0, arg1)
}

// DrawRates is a free data retrieval call binding the contract method 0x759e7f8a.
//
// Solidity: function DrawRates(address , uint256 ) view returns(uint256 rate, uint256 drawRateTime, uint256 totalReward, uint256 leftReward)
func (_PirateDeposit *PirateDepositCaller) DrawRates(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Rate         *big.Int
	DrawRateTime *big.Int
	TotalReward  *big.Int
	LeftReward   *big.Int
}, error) {
	var out []interface{}
	err := _PirateDeposit.contract.Call(opts, &out, "DrawRates", arg0, arg1)

	outstruct := new(struct {
		Rate         *big.Int
		DrawRateTime *big.Int
		TotalReward  *big.Int
		LeftReward   *big.Int
	})

	outstruct.Rate = out[0].(*big.Int)
	outstruct.DrawRateTime = out[1].(*big.Int)
	outstruct.TotalReward = out[2].(*big.Int)
	outstruct.LeftReward = out[3].(*big.Int)

	return *outstruct, err

}

// DrawRates is a free data retrieval call binding the contract method 0x759e7f8a.
//
// Solidity: function DrawRates(address , uint256 ) view returns(uint256 rate, uint256 drawRateTime, uint256 totalReward, uint256 leftReward)
func (_PirateDeposit *PirateDepositSession) DrawRates(arg0 common.Address, arg1 *big.Int) (struct {
	Rate         *big.Int
	DrawRateTime *big.Int
	TotalReward  *big.Int
	LeftReward   *big.Int
}, error) {
	return _PirateDeposit.Contract.DrawRates(&_PirateDeposit.CallOpts, arg0, arg1)
}

// DrawRates is a free data retrieval call binding the contract method 0x759e7f8a.
//
// Solidity: function DrawRates(address , uint256 ) view returns(uint256 rate, uint256 drawRateTime, uint256 totalReward, uint256 leftReward)
func (_PirateDeposit *PirateDepositCallerSession) DrawRates(arg0 common.Address, arg1 *big.Int) (struct {
	Rate         *big.Int
	DrawRateTime *big.Int
	TotalReward  *big.Int
	LeftReward   *big.Int
}, error) {
	return _PirateDeposit.Contract.DrawRates(&_PirateDeposit.CallOpts, arg0, arg1)
}

// Coordinator is a free data retrieval call binding the contract method 0x0a009097.
//
// Solidity: function coordinator() view returns(address)
func (_PirateDeposit *PirateDepositCaller) Coordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PirateDeposit.contract.Call(opts, &out, "coordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coordinator is a free data retrieval call binding the contract method 0x0a009097.
//
// Solidity: function coordinator() view returns(address)
func (_PirateDeposit *PirateDepositSession) Coordinator() (common.Address, error) {
	return _PirateDeposit.Contract.Coordinator(&_PirateDeposit.CallOpts)
}

// Coordinator is a free data retrieval call binding the contract method 0x0a009097.
//
// Solidity: function coordinator() view returns(address)
func (_PirateDeposit *PirateDepositCallerSession) Coordinator() (common.Address, error) {
	return _PirateDeposit.Contract.Coordinator(&_PirateDeposit.CallOpts)
}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_PirateDeposit *PirateDepositCaller) Market(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PirateDeposit.contract.Call(opts, &out, "market")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_PirateDeposit *PirateDepositSession) Market() (common.Address, error) {
	return _PirateDeposit.Contract.Market(&_PirateDeposit.CallOpts)
}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_PirateDeposit *PirateDepositCallerSession) Market() (common.Address, error) {
	return _PirateDeposit.Contract.Market(&_PirateDeposit.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_PirateDeposit *PirateDepositCaller) MinDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PirateDeposit.contract.Call(opts, &out, "minDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_PirateDeposit *PirateDepositSession) MinDeposit() (*big.Int, error) {
	return _PirateDeposit.Contract.MinDeposit(&_PirateDeposit.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_PirateDeposit *PirateDepositCallerSession) MinDeposit() (*big.Int, error) {
	return _PirateDeposit.Contract.MinDeposit(&_PirateDeposit.CallOpts)
}

// MinDepositInterval is a free data retrieval call binding the contract method 0xd2040687.
//
// Solidity: function minDepositInterval() view returns(uint256)
func (_PirateDeposit *PirateDepositCaller) MinDepositInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PirateDeposit.contract.Call(opts, &out, "minDepositInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDepositInterval is a free data retrieval call binding the contract method 0xd2040687.
//
// Solidity: function minDepositInterval() view returns(uint256)
func (_PirateDeposit *PirateDepositSession) MinDepositInterval() (*big.Int, error) {
	return _PirateDeposit.Contract.MinDepositInterval(&_PirateDeposit.CallOpts)
}

// MinDepositInterval is a free data retrieval call binding the contract method 0xd2040687.
//
// Solidity: function minDepositInterval() view returns(uint256)
func (_PirateDeposit *PirateDepositCallerSession) MinDepositInterval() (*big.Int, error) {
	return _PirateDeposit.Contract.MinDepositInterval(&_PirateDeposit.CallOpts)
}

// MinOneMonth is a free data retrieval call binding the contract method 0x3e7754c1.
//
// Solidity: function minOneMonth() view returns(uint256)
func (_PirateDeposit *PirateDepositCaller) MinOneMonth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PirateDeposit.contract.Call(opts, &out, "minOneMonth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinOneMonth is a free data retrieval call binding the contract method 0x3e7754c1.
//
// Solidity: function minOneMonth() view returns(uint256)
func (_PirateDeposit *PirateDepositSession) MinOneMonth() (*big.Int, error) {
	return _PirateDeposit.Contract.MinOneMonth(&_PirateDeposit.CallOpts)
}

// MinOneMonth is a free data retrieval call binding the contract method 0x3e7754c1.
//
// Solidity: function minOneMonth() view returns(uint256)
func (_PirateDeposit *PirateDepositCallerSession) MinOneMonth() (*big.Int, error) {
	return _PirateDeposit.Contract.MinOneMonth(&_PirateDeposit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PirateDeposit *PirateDepositCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PirateDeposit.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PirateDeposit *PirateDepositSession) Owner() (common.Address, error) {
	return _PirateDeposit.Contract.Owner(&_PirateDeposit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PirateDeposit *PirateDepositCallerSession) Owner() (common.Address, error) {
	return _PirateDeposit.Contract.Owner(&_PirateDeposit.CallOpts)
}

// PoolDrawInterval is a free data retrieval call binding the contract method 0xb4ec753e.
//
// Solidity: function poolDrawInterval() view returns(uint256)
func (_PirateDeposit *PirateDepositCaller) PoolDrawInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PirateDeposit.contract.Call(opts, &out, "poolDrawInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolDrawInterval is a free data retrieval call binding the contract method 0xb4ec753e.
//
// Solidity: function poolDrawInterval() view returns(uint256)
func (_PirateDeposit *PirateDepositSession) PoolDrawInterval() (*big.Int, error) {
	return _PirateDeposit.Contract.PoolDrawInterval(&_PirateDeposit.CallOpts)
}

// PoolDrawInterval is a free data retrieval call binding the contract method 0xb4ec753e.
//
// Solidity: function poolDrawInterval() view returns(uint256)
func (_PirateDeposit *PirateDepositCallerSession) PoolDrawInterval() (*big.Int, error) {
	return _PirateDeposit.Contract.PoolDrawInterval(&_PirateDeposit.CallOpts)
}

// RateDecimal is a free data retrieval call binding the contract method 0x31809c7e.
//
// Solidity: function rateDecimal() view returns(uint256)
func (_PirateDeposit *PirateDepositCaller) RateDecimal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PirateDeposit.contract.Call(opts, &out, "rateDecimal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RateDecimal is a free data retrieval call binding the contract method 0x31809c7e.
//
// Solidity: function rateDecimal() view returns(uint256)
func (_PirateDeposit *PirateDepositSession) RateDecimal() (*big.Int, error) {
	return _PirateDeposit.Contract.RateDecimal(&_PirateDeposit.CallOpts)
}

// RateDecimal is a free data retrieval call binding the contract method 0x31809c7e.
//
// Solidity: function rateDecimal() view returns(uint256)
func (_PirateDeposit *PirateDepositCallerSession) RateDecimal() (*big.Int, error) {
	return _PirateDeposit.Contract.RateDecimal(&_PirateDeposit.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_PirateDeposit *PirateDepositCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PirateDeposit.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_PirateDeposit *PirateDepositSession) Token() (common.Address, error) {
	return _PirateDeposit.Contract.Token(&_PirateDeposit.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_PirateDeposit *PirateDepositCallerSession) Token() (common.Address, error) {
	return _PirateDeposit.Contract.Token(&_PirateDeposit.CallOpts)
}

// AddReward is a paid mutator transaction binding the contract method 0x70dc5ca5.
//
// Solidity: function addReward(address pool, uint256 rate, uint256 totalReward, uint256 drawTime) returns()
func (_PirateDeposit *PirateDepositTransactor) AddReward(opts *bind.TransactOpts, pool common.Address, rate *big.Int, totalReward *big.Int, drawTime *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.contract.Transact(opts, "addReward", pool, rate, totalReward, drawTime)
}

// AddReward is a paid mutator transaction binding the contract method 0x70dc5ca5.
//
// Solidity: function addReward(address pool, uint256 rate, uint256 totalReward, uint256 drawTime) returns()
func (_PirateDeposit *PirateDepositSession) AddReward(pool common.Address, rate *big.Int, totalReward *big.Int, drawTime *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.Contract.AddReward(&_PirateDeposit.TransactOpts, pool, rate, totalReward, drawTime)
}

// AddReward is a paid mutator transaction binding the contract method 0x70dc5ca5.
//
// Solidity: function addReward(address pool, uint256 rate, uint256 totalReward, uint256 drawTime) returns()
func (_PirateDeposit *PirateDepositTransactorSession) AddReward(pool common.Address, rate *big.Int, totalReward *big.Int, drawTime *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.Contract.AddReward(&_PirateDeposit.TransactOpts, pool, rate, totalReward, drawTime)
}

// ChangeCoordinator is a paid mutator transaction binding the contract method 0x86ab2395.
//
// Solidity: function changeCoordinator(address coorD) returns()
func (_PirateDeposit *PirateDepositTransactor) ChangeCoordinator(opts *bind.TransactOpts, coorD common.Address) (*types.Transaction, error) {
	return _PirateDeposit.contract.Transact(opts, "changeCoordinator", coorD)
}

// ChangeCoordinator is a paid mutator transaction binding the contract method 0x86ab2395.
//
// Solidity: function changeCoordinator(address coorD) returns()
func (_PirateDeposit *PirateDepositSession) ChangeCoordinator(coorD common.Address) (*types.Transaction, error) {
	return _PirateDeposit.Contract.ChangeCoordinator(&_PirateDeposit.TransactOpts, coorD)
}

// ChangeCoordinator is a paid mutator transaction binding the contract method 0x86ab2395.
//
// Solidity: function changeCoordinator(address coorD) returns()
func (_PirateDeposit *PirateDepositTransactorSession) ChangeCoordinator(coorD common.Address) (*types.Transaction, error) {
	return _PirateDeposit.Contract.ChangeCoordinator(&_PirateDeposit.TransactOpts, coorD)
}

// ChangeDepositInterval is a paid mutator transaction binding the contract method 0xd91edb8e.
//
// Solidity: function changeDepositInterval(uint256 d) returns()
func (_PirateDeposit *PirateDepositTransactor) ChangeDepositInterval(opts *bind.TransactOpts, d *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.contract.Transact(opts, "changeDepositInterval", d)
}

// ChangeDepositInterval is a paid mutator transaction binding the contract method 0xd91edb8e.
//
// Solidity: function changeDepositInterval(uint256 d) returns()
func (_PirateDeposit *PirateDepositSession) ChangeDepositInterval(d *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.Contract.ChangeDepositInterval(&_PirateDeposit.TransactOpts, d)
}

// ChangeDepositInterval is a paid mutator transaction binding the contract method 0xd91edb8e.
//
// Solidity: function changeDepositInterval(uint256 d) returns()
func (_PirateDeposit *PirateDepositTransactorSession) ChangeDepositInterval(d *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.Contract.ChangeDepositInterval(&_PirateDeposit.TransactOpts, d)
}

// ChangeMinDeposit is a paid mutator transaction binding the contract method 0x4b2a12da.
//
// Solidity: function changeMinDeposit(uint256 min) returns()
func (_PirateDeposit *PirateDepositTransactor) ChangeMinDeposit(opts *bind.TransactOpts, min *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.contract.Transact(opts, "changeMinDeposit", min)
}

// ChangeMinDeposit is a paid mutator transaction binding the contract method 0x4b2a12da.
//
// Solidity: function changeMinDeposit(uint256 min) returns()
func (_PirateDeposit *PirateDepositSession) ChangeMinDeposit(min *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.Contract.ChangeMinDeposit(&_PirateDeposit.TransactOpts, min)
}

// ChangeMinDeposit is a paid mutator transaction binding the contract method 0x4b2a12da.
//
// Solidity: function changeMinDeposit(uint256 min) returns()
func (_PirateDeposit *PirateDepositTransactorSession) ChangeMinDeposit(min *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.Contract.ChangeMinDeposit(&_PirateDeposit.TransactOpts, min)
}

// ChangePoolDrawIntval is a paid mutator transaction binding the contract method 0xf0f605eb.
//
// Solidity: function changePoolDrawIntval(uint256 d) returns()
func (_PirateDeposit *PirateDepositTransactor) ChangePoolDrawIntval(opts *bind.TransactOpts, d *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.contract.Transact(opts, "changePoolDrawIntval", d)
}

// ChangePoolDrawIntval is a paid mutator transaction binding the contract method 0xf0f605eb.
//
// Solidity: function changePoolDrawIntval(uint256 d) returns()
func (_PirateDeposit *PirateDepositSession) ChangePoolDrawIntval(d *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.Contract.ChangePoolDrawIntval(&_PirateDeposit.TransactOpts, d)
}

// ChangePoolDrawIntval is a paid mutator transaction binding the contract method 0xf0f605eb.
//
// Solidity: function changePoolDrawIntval(uint256 d) returns()
func (_PirateDeposit *PirateDepositTransactorSession) ChangePoolDrawIntval(d *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.Contract.ChangePoolDrawIntval(&_PirateDeposit.TransactOpts, d)
}

// ChangeRateDecimal is a paid mutator transaction binding the contract method 0x7c543100.
//
// Solidity: function changeRateDecimal(uint256 d) returns()
func (_PirateDeposit *PirateDepositTransactor) ChangeRateDecimal(opts *bind.TransactOpts, d *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.contract.Transact(opts, "changeRateDecimal", d)
}

// ChangeRateDecimal is a paid mutator transaction binding the contract method 0x7c543100.
//
// Solidity: function changeRateDecimal(uint256 d) returns()
func (_PirateDeposit *PirateDepositSession) ChangeRateDecimal(d *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.Contract.ChangeRateDecimal(&_PirateDeposit.TransactOpts, d)
}

// ChangeRateDecimal is a paid mutator transaction binding the contract method 0x7c543100.
//
// Solidity: function changeRateDecimal(uint256 d) returns()
func (_PirateDeposit *PirateDepositTransactorSession) ChangeRateDecimal(d *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.Contract.ChangeRateDecimal(&_PirateDeposit.TransactOpts, d)
}

// ChangeRewardInterval is a paid mutator transaction binding the contract method 0x61ea2392.
//
// Solidity: function changeRewardInterval(uint256 d) returns()
func (_PirateDeposit *PirateDepositTransactor) ChangeRewardInterval(opts *bind.TransactOpts, d *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.contract.Transact(opts, "changeRewardInterval", d)
}

// ChangeRewardInterval is a paid mutator transaction binding the contract method 0x61ea2392.
//
// Solidity: function changeRewardInterval(uint256 d) returns()
func (_PirateDeposit *PirateDepositSession) ChangeRewardInterval(d *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.Contract.ChangeRewardInterval(&_PirateDeposit.TransactOpts, d)
}

// ChangeRewardInterval is a paid mutator transaction binding the contract method 0x61ea2392.
//
// Solidity: function changeRewardInterval(uint256 d) returns()
func (_PirateDeposit *PirateDepositTransactorSession) ChangeRewardInterval(d *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.Contract.ChangeRewardInterval(&_PirateDeposit.TransactOpts, d)
}

// DoUserDeposit is a paid mutator transaction binding the contract method 0x67cf3e8f.
//
// Solidity: function doUserDeposit(address pool, uint256 tokenNumber, uint256 index) returns()
func (_PirateDeposit *PirateDepositTransactor) DoUserDeposit(opts *bind.TransactOpts, pool common.Address, tokenNumber *big.Int, index *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.contract.Transact(opts, "doUserDeposit", pool, tokenNumber, index)
}

// DoUserDeposit is a paid mutator transaction binding the contract method 0x67cf3e8f.
//
// Solidity: function doUserDeposit(address pool, uint256 tokenNumber, uint256 index) returns()
func (_PirateDeposit *PirateDepositSession) DoUserDeposit(pool common.Address, tokenNumber *big.Int, index *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.Contract.DoUserDeposit(&_PirateDeposit.TransactOpts, pool, tokenNumber, index)
}

// DoUserDeposit is a paid mutator transaction binding the contract method 0x67cf3e8f.
//
// Solidity: function doUserDeposit(address pool, uint256 tokenNumber, uint256 index) returns()
func (_PirateDeposit *PirateDepositTransactorSession) DoUserDeposit(pool common.Address, tokenNumber *big.Int, index *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.Contract.DoUserDeposit(&_PirateDeposit.TransactOpts, pool, tokenNumber, index)
}

// DrawReward is a paid mutator transaction binding the contract method 0x83e069b6.
//
// Solidity: function drawReward(address pool, uint256 poolIndex) returns()
func (_PirateDeposit *PirateDepositTransactor) DrawReward(opts *bind.TransactOpts, pool common.Address, poolIndex *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.contract.Transact(opts, "drawReward", pool, poolIndex)
}

// DrawReward is a paid mutator transaction binding the contract method 0x83e069b6.
//
// Solidity: function drawReward(address pool, uint256 poolIndex) returns()
func (_PirateDeposit *PirateDepositSession) DrawReward(pool common.Address, poolIndex *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.Contract.DrawReward(&_PirateDeposit.TransactOpts, pool, poolIndex)
}

// DrawReward is a paid mutator transaction binding the contract method 0x83e069b6.
//
// Solidity: function drawReward(address pool, uint256 poolIndex) returns()
func (_PirateDeposit *PirateDepositTransactorSession) DrawReward(pool common.Address, poolIndex *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.Contract.DrawReward(&_PirateDeposit.TransactOpts, pool, poolIndex)
}

// PoolDrawReward is a paid mutator transaction binding the contract method 0xa33e9754.
//
// Solidity: function poolDrawReward(uint256 poolIndex) returns()
func (_PirateDeposit *PirateDepositTransactor) PoolDrawReward(opts *bind.TransactOpts, poolIndex *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.contract.Transact(opts, "poolDrawReward", poolIndex)
}

// PoolDrawReward is a paid mutator transaction binding the contract method 0xa33e9754.
//
// Solidity: function poolDrawReward(uint256 poolIndex) returns()
func (_PirateDeposit *PirateDepositSession) PoolDrawReward(poolIndex *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.Contract.PoolDrawReward(&_PirateDeposit.TransactOpts, poolIndex)
}

// PoolDrawReward is a paid mutator transaction binding the contract method 0xa33e9754.
//
// Solidity: function poolDrawReward(uint256 poolIndex) returns()
func (_PirateDeposit *PirateDepositTransactorSession) PoolDrawReward(poolIndex *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.Contract.PoolDrawReward(&_PirateDeposit.TransactOpts, poolIndex)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PirateDeposit *PirateDepositTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PirateDeposit.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PirateDeposit *PirateDepositSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PirateDeposit.Contract.TransferOwnership(&_PirateDeposit.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PirateDeposit *PirateDepositTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PirateDeposit.Contract.TransferOwnership(&_PirateDeposit.TransactOpts, newOwner)
}

// UserWithDrawDeposit is a paid mutator transaction binding the contract method 0xe8c4e1cd.
//
// Solidity: function userWithDrawDeposit(address pool, uint256 recordTime, uint256 index, uint256 poolIndex) returns()
func (_PirateDeposit *PirateDepositTransactor) UserWithDrawDeposit(opts *bind.TransactOpts, pool common.Address, recordTime *big.Int, index *big.Int, poolIndex *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.contract.Transact(opts, "userWithDrawDeposit", pool, recordTime, index, poolIndex)
}

// UserWithDrawDeposit is a paid mutator transaction binding the contract method 0xe8c4e1cd.
//
// Solidity: function userWithDrawDeposit(address pool, uint256 recordTime, uint256 index, uint256 poolIndex) returns()
func (_PirateDeposit *PirateDepositSession) UserWithDrawDeposit(pool common.Address, recordTime *big.Int, index *big.Int, poolIndex *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.Contract.UserWithDrawDeposit(&_PirateDeposit.TransactOpts, pool, recordTime, index, poolIndex)
}

// UserWithDrawDeposit is a paid mutator transaction binding the contract method 0xe8c4e1cd.
//
// Solidity: function userWithDrawDeposit(address pool, uint256 recordTime, uint256 index, uint256 poolIndex) returns()
func (_PirateDeposit *PirateDepositTransactorSession) UserWithDrawDeposit(pool common.Address, recordTime *big.Int, index *big.Int, poolIndex *big.Int) (*types.Transaction, error) {
	return _PirateDeposit.Contract.UserWithDrawDeposit(&_PirateDeposit.TransactOpts, pool, recordTime, index, poolIndex)
}

// PirateDepositAddRewardEventIterator is returned from FilterAddRewardEvent and is used to iterate over the raw logs and unpacked data for AddRewardEvent events raised by the PirateDeposit contract.
type PirateDepositAddRewardEventIterator struct {
	Event *PirateDepositAddRewardEvent // Event containing the contract specifics and raw log

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
func (it *PirateDepositAddRewardEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PirateDepositAddRewardEvent)
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
		it.Event = new(PirateDepositAddRewardEvent)
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
func (it *PirateDepositAddRewardEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PirateDepositAddRewardEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PirateDepositAddRewardEvent represents a AddRewardEvent event raised by the PirateDeposit contract.
type PirateDepositAddRewardEvent struct {
	Pool        common.Address
	Rate        *big.Int
	TotalReward *big.Int
	RewardTime  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddRewardEvent is a free log retrieval operation binding the contract event 0x77d89f931b457f29d8507338a1f00da3300127bf09adf9f5e28cda890fa3959c.
//
// Solidity: event AddRewardEvent(address pool, uint256 rate, uint256 totalReward, uint256 rewardTime)
func (_PirateDeposit *PirateDepositFilterer) FilterAddRewardEvent(opts *bind.FilterOpts) (*PirateDepositAddRewardEventIterator, error) {

	logs, sub, err := _PirateDeposit.contract.FilterLogs(opts, "AddRewardEvent")
	if err != nil {
		return nil, err
	}
	return &PirateDepositAddRewardEventIterator{contract: _PirateDeposit.contract, event: "AddRewardEvent", logs: logs, sub: sub}, nil
}

// WatchAddRewardEvent is a free log subscription operation binding the contract event 0x77d89f931b457f29d8507338a1f00da3300127bf09adf9f5e28cda890fa3959c.
//
// Solidity: event AddRewardEvent(address pool, uint256 rate, uint256 totalReward, uint256 rewardTime)
func (_PirateDeposit *PirateDepositFilterer) WatchAddRewardEvent(opts *bind.WatchOpts, sink chan<- *PirateDepositAddRewardEvent) (event.Subscription, error) {

	logs, sub, err := _PirateDeposit.contract.WatchLogs(opts, "AddRewardEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PirateDepositAddRewardEvent)
				if err := _PirateDeposit.contract.UnpackLog(event, "AddRewardEvent", log); err != nil {
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

// ParseAddRewardEvent is a log parse operation binding the contract event 0x77d89f931b457f29d8507338a1f00da3300127bf09adf9f5e28cda890fa3959c.
//
// Solidity: event AddRewardEvent(address pool, uint256 rate, uint256 totalReward, uint256 rewardTime)
func (_PirateDeposit *PirateDepositFilterer) ParseAddRewardEvent(log types.Log) (*PirateDepositAddRewardEvent, error) {
	event := new(PirateDepositAddRewardEvent)
	if err := _PirateDeposit.contract.UnpackLog(event, "AddRewardEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PirateDepositPoolDrawRewardEventIterator is returned from FilterPoolDrawRewardEvent and is used to iterate over the raw logs and unpacked data for PoolDrawRewardEvent events raised by the PirateDeposit contract.
type PirateDepositPoolDrawRewardEventIterator struct {
	Event *PirateDepositPoolDrawRewardEvent // Event containing the contract specifics and raw log

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
func (it *PirateDepositPoolDrawRewardEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PirateDepositPoolDrawRewardEvent)
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
		it.Event = new(PirateDepositPoolDrawRewardEvent)
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
func (it *PirateDepositPoolDrawRewardEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PirateDepositPoolDrawRewardEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PirateDepositPoolDrawRewardEvent represents a PoolDrawRewardEvent event raised by the PirateDeposit contract.
type PirateDepositPoolDrawRewardEvent struct {
	Pool    common.Address
	Profits *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPoolDrawRewardEvent is a free log retrieval operation binding the contract event 0xc81302b4d924f185feb5092738f4359012bdd23fb0f77f0fd40b99d40f5a8cc9.
//
// Solidity: event PoolDrawRewardEvent(address pool, uint256 profits)
func (_PirateDeposit *PirateDepositFilterer) FilterPoolDrawRewardEvent(opts *bind.FilterOpts) (*PirateDepositPoolDrawRewardEventIterator, error) {

	logs, sub, err := _PirateDeposit.contract.FilterLogs(opts, "PoolDrawRewardEvent")
	if err != nil {
		return nil, err
	}
	return &PirateDepositPoolDrawRewardEventIterator{contract: _PirateDeposit.contract, event: "PoolDrawRewardEvent", logs: logs, sub: sub}, nil
}

// WatchPoolDrawRewardEvent is a free log subscription operation binding the contract event 0xc81302b4d924f185feb5092738f4359012bdd23fb0f77f0fd40b99d40f5a8cc9.
//
// Solidity: event PoolDrawRewardEvent(address pool, uint256 profits)
func (_PirateDeposit *PirateDepositFilterer) WatchPoolDrawRewardEvent(opts *bind.WatchOpts, sink chan<- *PirateDepositPoolDrawRewardEvent) (event.Subscription, error) {

	logs, sub, err := _PirateDeposit.contract.WatchLogs(opts, "PoolDrawRewardEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PirateDepositPoolDrawRewardEvent)
				if err := _PirateDeposit.contract.UnpackLog(event, "PoolDrawRewardEvent", log); err != nil {
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

// ParsePoolDrawRewardEvent is a log parse operation binding the contract event 0xc81302b4d924f185feb5092738f4359012bdd23fb0f77f0fd40b99d40f5a8cc9.
//
// Solidity: event PoolDrawRewardEvent(address pool, uint256 profits)
func (_PirateDeposit *PirateDepositFilterer) ParsePoolDrawRewardEvent(log types.Log) (*PirateDepositPoolDrawRewardEvent, error) {
	event := new(PirateDepositPoolDrawRewardEvent)
	if err := _PirateDeposit.contract.UnpackLog(event, "PoolDrawRewardEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PirateDepositUserDepositEventIterator is returned from FilterUserDepositEvent and is used to iterate over the raw logs and unpacked data for UserDepositEvent events raised by the PirateDeposit contract.
type PirateDepositUserDepositEventIterator struct {
	Event *PirateDepositUserDepositEvent // Event containing the contract specifics and raw log

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
func (it *PirateDepositUserDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PirateDepositUserDepositEvent)
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
		it.Event = new(PirateDepositUserDepositEvent)
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
func (it *PirateDepositUserDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PirateDepositUserDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PirateDepositUserDepositEvent represents a UserDepositEvent event raised by the PirateDeposit contract.
type PirateDepositUserDepositEvent struct {
	User        common.Address
	Pool        common.Address
	TokenNumber *big.Int
	DepositTime *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUserDepositEvent is a free log retrieval operation binding the contract event 0x0422644439dc2134e423bc8f35c714ff41e2998a7c44c1210ec2f14a389ea8f9.
//
// Solidity: event UserDepositEvent(address indexed user, address indexed pool, uint256 tokenNumber, uint256 depositTime)
func (_PirateDeposit *PirateDepositFilterer) FilterUserDepositEvent(opts *bind.FilterOpts, user []common.Address, pool []common.Address) (*PirateDepositUserDepositEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _PirateDeposit.contract.FilterLogs(opts, "UserDepositEvent", userRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &PirateDepositUserDepositEventIterator{contract: _PirateDeposit.contract, event: "UserDepositEvent", logs: logs, sub: sub}, nil
}

// WatchUserDepositEvent is a free log subscription operation binding the contract event 0x0422644439dc2134e423bc8f35c714ff41e2998a7c44c1210ec2f14a389ea8f9.
//
// Solidity: event UserDepositEvent(address indexed user, address indexed pool, uint256 tokenNumber, uint256 depositTime)
func (_PirateDeposit *PirateDepositFilterer) WatchUserDepositEvent(opts *bind.WatchOpts, sink chan<- *PirateDepositUserDepositEvent, user []common.Address, pool []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _PirateDeposit.contract.WatchLogs(opts, "UserDepositEvent", userRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PirateDepositUserDepositEvent)
				if err := _PirateDeposit.contract.UnpackLog(event, "UserDepositEvent", log); err != nil {
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

// ParseUserDepositEvent is a log parse operation binding the contract event 0x0422644439dc2134e423bc8f35c714ff41e2998a7c44c1210ec2f14a389ea8f9.
//
// Solidity: event UserDepositEvent(address indexed user, address indexed pool, uint256 tokenNumber, uint256 depositTime)
func (_PirateDeposit *PirateDepositFilterer) ParseUserDepositEvent(log types.Log) (*PirateDepositUserDepositEvent, error) {
	event := new(PirateDepositUserDepositEvent)
	if err := _PirateDeposit.contract.UnpackLog(event, "UserDepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PirateDepositUserDrawRewardEventIterator is returned from FilterUserDrawRewardEvent and is used to iterate over the raw logs and unpacked data for UserDrawRewardEvent events raised by the PirateDeposit contract.
type PirateDepositUserDrawRewardEventIterator struct {
	Event *PirateDepositUserDrawRewardEvent // Event containing the contract specifics and raw log

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
func (it *PirateDepositUserDrawRewardEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PirateDepositUserDrawRewardEvent)
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
		it.Event = new(PirateDepositUserDrawRewardEvent)
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
func (it *PirateDepositUserDrawRewardEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PirateDepositUserDrawRewardEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PirateDepositUserDrawRewardEvent represents a UserDrawRewardEvent event raised by the PirateDeposit contract.
type PirateDepositUserDrawRewardEvent struct {
	User      common.Address
	Pool      common.Address
	Reward    *big.Int
	LastIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUserDrawRewardEvent is a free log retrieval operation binding the contract event 0xc5c1035a162fd77d3905e71c364d5d061b6c1984d5894813c0e63cfcacb09784.
//
// Solidity: event UserDrawRewardEvent(address indexed user, address indexed pool, uint256 reward, uint256 lastIndex)
func (_PirateDeposit *PirateDepositFilterer) FilterUserDrawRewardEvent(opts *bind.FilterOpts, user []common.Address, pool []common.Address) (*PirateDepositUserDrawRewardEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _PirateDeposit.contract.FilterLogs(opts, "UserDrawRewardEvent", userRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &PirateDepositUserDrawRewardEventIterator{contract: _PirateDeposit.contract, event: "UserDrawRewardEvent", logs: logs, sub: sub}, nil
}

// WatchUserDrawRewardEvent is a free log subscription operation binding the contract event 0xc5c1035a162fd77d3905e71c364d5d061b6c1984d5894813c0e63cfcacb09784.
//
// Solidity: event UserDrawRewardEvent(address indexed user, address indexed pool, uint256 reward, uint256 lastIndex)
func (_PirateDeposit *PirateDepositFilterer) WatchUserDrawRewardEvent(opts *bind.WatchOpts, sink chan<- *PirateDepositUserDrawRewardEvent, user []common.Address, pool []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _PirateDeposit.contract.WatchLogs(opts, "UserDrawRewardEvent", userRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PirateDepositUserDrawRewardEvent)
				if err := _PirateDeposit.contract.UnpackLog(event, "UserDrawRewardEvent", log); err != nil {
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

// ParseUserDrawRewardEvent is a log parse operation binding the contract event 0xc5c1035a162fd77d3905e71c364d5d061b6c1984d5894813c0e63cfcacb09784.
//
// Solidity: event UserDrawRewardEvent(address indexed user, address indexed pool, uint256 reward, uint256 lastIndex)
func (_PirateDeposit *PirateDepositFilterer) ParseUserDrawRewardEvent(log types.Log) (*PirateDepositUserDrawRewardEvent, error) {
	event := new(PirateDepositUserDrawRewardEvent)
	if err := _PirateDeposit.contract.UnpackLog(event, "UserDrawRewardEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PirateDepositUserWithDrawDepositEventIterator is returned from FilterUserWithDrawDepositEvent and is used to iterate over the raw logs and unpacked data for UserWithDrawDepositEvent events raised by the PirateDeposit contract.
type PirateDepositUserWithDrawDepositEventIterator struct {
	Event *PirateDepositUserWithDrawDepositEvent // Event containing the contract specifics and raw log

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
func (it *PirateDepositUserWithDrawDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PirateDepositUserWithDrawDepositEvent)
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
		it.Event = new(PirateDepositUserWithDrawDepositEvent)
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
func (it *PirateDepositUserWithDrawDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PirateDepositUserWithDrawDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PirateDepositUserWithDrawDepositEvent represents a UserWithDrawDepositEvent event raised by the PirateDeposit contract.
type PirateDepositUserWithDrawDepositEvent struct {
	User        common.Address
	Pool        common.Address
	DepositTime *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUserWithDrawDepositEvent is a free log retrieval operation binding the contract event 0x80cdbfe7313fa67c29d74c1df0f788d97f25564b7a8aa0deee2cea5a962c2df6.
//
// Solidity: event UserWithDrawDepositEvent(address indexed user, address indexed pool, uint256 depositTime)
func (_PirateDeposit *PirateDepositFilterer) FilterUserWithDrawDepositEvent(opts *bind.FilterOpts, user []common.Address, pool []common.Address) (*PirateDepositUserWithDrawDepositEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _PirateDeposit.contract.FilterLogs(opts, "UserWithDrawDepositEvent", userRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &PirateDepositUserWithDrawDepositEventIterator{contract: _PirateDeposit.contract, event: "UserWithDrawDepositEvent", logs: logs, sub: sub}, nil
}

// WatchUserWithDrawDepositEvent is a free log subscription operation binding the contract event 0x80cdbfe7313fa67c29d74c1df0f788d97f25564b7a8aa0deee2cea5a962c2df6.
//
// Solidity: event UserWithDrawDepositEvent(address indexed user, address indexed pool, uint256 depositTime)
func (_PirateDeposit *PirateDepositFilterer) WatchUserWithDrawDepositEvent(opts *bind.WatchOpts, sink chan<- *PirateDepositUserWithDrawDepositEvent, user []common.Address, pool []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _PirateDeposit.contract.WatchLogs(opts, "UserWithDrawDepositEvent", userRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PirateDepositUserWithDrawDepositEvent)
				if err := _PirateDeposit.contract.UnpackLog(event, "UserWithDrawDepositEvent", log); err != nil {
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

// ParseUserWithDrawDepositEvent is a log parse operation binding the contract event 0x80cdbfe7313fa67c29d74c1df0f788d97f25564b7a8aa0deee2cea5a962c2df6.
//
// Solidity: event UserWithDrawDepositEvent(address indexed user, address indexed pool, uint256 depositTime)
func (_PirateDeposit *PirateDepositFilterer) ParseUserWithDrawDepositEvent(log types.Log) (*PirateDepositUserWithDrawDepositEvent, error) {
	event := new(PirateDepositUserWithDrawDepositEvent)
	if err := _PirateDeposit.contract.UnpackLog(event, "UserWithDrawDepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

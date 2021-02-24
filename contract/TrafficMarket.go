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

// TrafficMarketABI is the input ABI used to generate the binding from.
const TrafficMarketABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"trafficAmount\",\"type\":\"uint256\"}],\"name\":\"Charge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subAddr\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr2\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"eventType\",\"type\":\"uint8\"}],\"name\":\"MinerEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minerUsedPacket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minerPacket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimedBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolTotalPacket\",\"type\":\"uint256\"}],\"name\":\"PoolClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"eventType\",\"type\":\"uint8\"}],\"name\":\"PoolReg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SettingsChanged\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"LegalPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MBytesPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MinerDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PoolDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Pools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"UserData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalChargeBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTraffic\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blockChainSettings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"subAddr\",\"type\":\"bytes32\"}],\"name\":\"changePool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pDpos\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mDpos\",\"type\":\"uint256\"}],\"name\":\"changeSettings\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenNo\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"poolAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"charge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"destroyPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"emergency\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPoolList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"subAddr\",\"type\":\"bytes32\"}],\"name\":\"joinPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"payerForMiner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minerCredit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedTraffic\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"pclaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"regPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"subAddr\",\"type\":\"bytes32\"}],\"name\":\"retireFromPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"tokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TrafficMarket is an auto generated Go binding around an Ethereum contract.
type TrafficMarket struct {
	TrafficMarketCaller     // Read-only binding to the contract
	TrafficMarketTransactor // Write-only binding to the contract
	TrafficMarketFilterer   // Log filterer for contract events
}

// TrafficMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrafficMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrafficMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrafficMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrafficMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TrafficMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrafficMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrafficMarketSession struct {
	Contract     *TrafficMarket    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TrafficMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrafficMarketCallerSession struct {
	Contract *TrafficMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TrafficMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrafficMarketTransactorSession struct {
	Contract     *TrafficMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TrafficMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrafficMarketRaw struct {
	Contract *TrafficMarket // Generic contract binding to access the raw methods on
}

// TrafficMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrafficMarketCallerRaw struct {
	Contract *TrafficMarketCaller // Generic read-only contract binding to access the raw methods on
}

// TrafficMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrafficMarketTransactorRaw struct {
	Contract *TrafficMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrafficMarket creates a new instance of TrafficMarket, bound to a specific deployed contract.
func NewTrafficMarket(address common.Address, backend bind.ContractBackend) (*TrafficMarket, error) {
	contract, err := bindTrafficMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TrafficMarket{TrafficMarketCaller: TrafficMarketCaller{contract: contract}, TrafficMarketTransactor: TrafficMarketTransactor{contract: contract}, TrafficMarketFilterer: TrafficMarketFilterer{contract: contract}}, nil
}

// NewTrafficMarketCaller creates a new read-only instance of TrafficMarket, bound to a specific deployed contract.
func NewTrafficMarketCaller(address common.Address, caller bind.ContractCaller) (*TrafficMarketCaller, error) {
	contract, err := bindTrafficMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TrafficMarketCaller{contract: contract}, nil
}

// NewTrafficMarketTransactor creates a new write-only instance of TrafficMarket, bound to a specific deployed contract.
func NewTrafficMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*TrafficMarketTransactor, error) {
	contract, err := bindTrafficMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TrafficMarketTransactor{contract: contract}, nil
}

// NewTrafficMarketFilterer creates a new log filterer instance of TrafficMarket, bound to a specific deployed contract.
func NewTrafficMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*TrafficMarketFilterer, error) {
	contract, err := bindTrafficMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TrafficMarketFilterer{contract: contract}, nil
}

// bindTrafficMarket binds a generic wrapper to an already deployed contract.
func bindTrafficMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TrafficMarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrafficMarket *TrafficMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrafficMarket.Contract.TrafficMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrafficMarket *TrafficMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrafficMarket.Contract.TrafficMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrafficMarket *TrafficMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrafficMarket.Contract.TrafficMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrafficMarket *TrafficMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrafficMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrafficMarket *TrafficMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrafficMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrafficMarket *TrafficMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrafficMarket.Contract.contract.Transact(opts, method, params...)
}

// LegalPool is a free data retrieval call binding the contract method 0x8df854aa.
//
// Solidity: function LegalPool(address poolAddr, uint256 index) view returns(bool)
func (_TrafficMarket *TrafficMarketCaller) LegalPool(opts *bind.CallOpts, poolAddr common.Address, index *big.Int) (bool, error) {
	var out []interface{}
	err := _TrafficMarket.contract.Call(opts, &out, "LegalPool", poolAddr, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LegalPool is a free data retrieval call binding the contract method 0x8df854aa.
//
// Solidity: function LegalPool(address poolAddr, uint256 index) view returns(bool)
func (_TrafficMarket *TrafficMarketSession) LegalPool(poolAddr common.Address, index *big.Int) (bool, error) {
	return _TrafficMarket.Contract.LegalPool(&_TrafficMarket.CallOpts, poolAddr, index)
}

// LegalPool is a free data retrieval call binding the contract method 0x8df854aa.
//
// Solidity: function LegalPool(address poolAddr, uint256 index) view returns(bool)
func (_TrafficMarket *TrafficMarketCallerSession) LegalPool(poolAddr common.Address, index *big.Int) (bool, error) {
	return _TrafficMarket.Contract.LegalPool(&_TrafficMarket.CallOpts, poolAddr, index)
}

// MBytesPerToken is a free data retrieval call binding the contract method 0xfb44f734.
//
// Solidity: function MBytesPerToken() view returns(uint256)
func (_TrafficMarket *TrafficMarketCaller) MBytesPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TrafficMarket.contract.Call(opts, &out, "MBytesPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MBytesPerToken is a free data retrieval call binding the contract method 0xfb44f734.
//
// Solidity: function MBytesPerToken() view returns(uint256)
func (_TrafficMarket *TrafficMarketSession) MBytesPerToken() (*big.Int, error) {
	return _TrafficMarket.Contract.MBytesPerToken(&_TrafficMarket.CallOpts)
}

// MBytesPerToken is a free data retrieval call binding the contract method 0xfb44f734.
//
// Solidity: function MBytesPerToken() view returns(uint256)
func (_TrafficMarket *TrafficMarketCallerSession) MBytesPerToken() (*big.Int, error) {
	return _TrafficMarket.Contract.MBytesPerToken(&_TrafficMarket.CallOpts)
}

// MinerDeposit is a free data retrieval call binding the contract method 0x9a3c7309.
//
// Solidity: function MinerDeposit() view returns(uint256)
func (_TrafficMarket *TrafficMarketCaller) MinerDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TrafficMarket.contract.Call(opts, &out, "MinerDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinerDeposit is a free data retrieval call binding the contract method 0x9a3c7309.
//
// Solidity: function MinerDeposit() view returns(uint256)
func (_TrafficMarket *TrafficMarketSession) MinerDeposit() (*big.Int, error) {
	return _TrafficMarket.Contract.MinerDeposit(&_TrafficMarket.CallOpts)
}

// MinerDeposit is a free data retrieval call binding the contract method 0x9a3c7309.
//
// Solidity: function MinerDeposit() view returns(uint256)
func (_TrafficMarket *TrafficMarketCallerSession) MinerDeposit() (*big.Int, error) {
	return _TrafficMarket.Contract.MinerDeposit(&_TrafficMarket.CallOpts)
}

// PoolDeposit is a free data retrieval call binding the contract method 0x040912f6.
//
// Solidity: function PoolDeposit() view returns(uint256)
func (_TrafficMarket *TrafficMarketCaller) PoolDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TrafficMarket.contract.Call(opts, &out, "PoolDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolDeposit is a free data retrieval call binding the contract method 0x040912f6.
//
// Solidity: function PoolDeposit() view returns(uint256)
func (_TrafficMarket *TrafficMarketSession) PoolDeposit() (*big.Int, error) {
	return _TrafficMarket.Contract.PoolDeposit(&_TrafficMarket.CallOpts)
}

// PoolDeposit is a free data retrieval call binding the contract method 0x040912f6.
//
// Solidity: function PoolDeposit() view returns(uint256)
func (_TrafficMarket *TrafficMarketCallerSession) PoolDeposit() (*big.Int, error) {
	return _TrafficMarket.Contract.PoolDeposit(&_TrafficMarket.CallOpts)
}

// Pools is a free data retrieval call binding the contract method 0x73dfccca.
//
// Solidity: function Pools(uint256 ) view returns(address)
func (_TrafficMarket *TrafficMarketCaller) Pools(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TrafficMarket.contract.Call(opts, &out, "Pools", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pools is a free data retrieval call binding the contract method 0x73dfccca.
//
// Solidity: function Pools(uint256 ) view returns(address)
func (_TrafficMarket *TrafficMarketSession) Pools(arg0 *big.Int) (common.Address, error) {
	return _TrafficMarket.Contract.Pools(&_TrafficMarket.CallOpts, arg0)
}

// Pools is a free data retrieval call binding the contract method 0x73dfccca.
//
// Solidity: function Pools(uint256 ) view returns(address)
func (_TrafficMarket *TrafficMarketCallerSession) Pools(arg0 *big.Int) (common.Address, error) {
	return _TrafficMarket.Contract.Pools(&_TrafficMarket.CallOpts, arg0)
}

// UserData is a free data retrieval call binding the contract method 0x5a903303.
//
// Solidity: function UserData(address , address ) view returns(uint256 totalChargeBalance, uint256 totalTraffic)
func (_TrafficMarket *TrafficMarketCaller) UserData(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	TotalChargeBalance *big.Int
	TotalTraffic       *big.Int
}, error) {
	var out []interface{}
	err := _TrafficMarket.contract.Call(opts, &out, "UserData", arg0, arg1)

	outstruct := new(struct {
		TotalChargeBalance *big.Int
		TotalTraffic       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalChargeBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalTraffic = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserData is a free data retrieval call binding the contract method 0x5a903303.
//
// Solidity: function UserData(address , address ) view returns(uint256 totalChargeBalance, uint256 totalTraffic)
func (_TrafficMarket *TrafficMarketSession) UserData(arg0 common.Address, arg1 common.Address) (struct {
	TotalChargeBalance *big.Int
	TotalTraffic       *big.Int
}, error) {
	return _TrafficMarket.Contract.UserData(&_TrafficMarket.CallOpts, arg0, arg1)
}

// UserData is a free data retrieval call binding the contract method 0x5a903303.
//
// Solidity: function UserData(address , address ) view returns(uint256 totalChargeBalance, uint256 totalTraffic)
func (_TrafficMarket *TrafficMarketCallerSession) UserData(arg0 common.Address, arg1 common.Address) (struct {
	TotalChargeBalance *big.Int
	TotalTraffic       *big.Int
}, error) {
	return _TrafficMarket.Contract.UserData(&_TrafficMarket.CallOpts, arg0, arg1)
}

// BlockChainSettings is a free data retrieval call binding the contract method 0x01d2eba2.
//
// Solidity: function blockChainSettings() view returns(uint256, uint256, uint256)
func (_TrafficMarket *TrafficMarketCaller) BlockChainSettings(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _TrafficMarket.contract.Call(opts, &out, "blockChainSettings")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// BlockChainSettings is a free data retrieval call binding the contract method 0x01d2eba2.
//
// Solidity: function blockChainSettings() view returns(uint256, uint256, uint256)
func (_TrafficMarket *TrafficMarketSession) BlockChainSettings() (*big.Int, *big.Int, *big.Int, error) {
	return _TrafficMarket.Contract.BlockChainSettings(&_TrafficMarket.CallOpts)
}

// BlockChainSettings is a free data retrieval call binding the contract method 0x01d2eba2.
//
// Solidity: function blockChainSettings() view returns(uint256, uint256, uint256)
func (_TrafficMarket *TrafficMarketCallerSession) BlockChainSettings() (*big.Int, *big.Int, *big.Int, error) {
	return _TrafficMarket.Contract.BlockChainSettings(&_TrafficMarket.CallOpts)
}

// Decimal is a free data retrieval call binding the contract method 0x76809ce3.
//
// Solidity: function decimal() view returns(uint256)
func (_TrafficMarket *TrafficMarketCaller) Decimal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TrafficMarket.contract.Call(opts, &out, "decimal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimal is a free data retrieval call binding the contract method 0x76809ce3.
//
// Solidity: function decimal() view returns(uint256)
func (_TrafficMarket *TrafficMarketSession) Decimal() (*big.Int, error) {
	return _TrafficMarket.Contract.Decimal(&_TrafficMarket.CallOpts)
}

// Decimal is a free data retrieval call binding the contract method 0x76809ce3.
//
// Solidity: function decimal() view returns(uint256)
func (_TrafficMarket *TrafficMarketCallerSession) Decimal() (*big.Int, error) {
	return _TrafficMarket.Contract.Decimal(&_TrafficMarket.CallOpts)
}

// GetPoolList is a free data retrieval call binding the contract method 0xd41dcbea.
//
// Solidity: function getPoolList() view returns(address[])
func (_TrafficMarket *TrafficMarketCaller) GetPoolList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TrafficMarket.contract.Call(opts, &out, "getPoolList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPoolList is a free data retrieval call binding the contract method 0xd41dcbea.
//
// Solidity: function getPoolList() view returns(address[])
func (_TrafficMarket *TrafficMarketSession) GetPoolList() ([]common.Address, error) {
	return _TrafficMarket.Contract.GetPoolList(&_TrafficMarket.CallOpts)
}

// GetPoolList is a free data retrieval call binding the contract method 0xd41dcbea.
//
// Solidity: function getPoolList() view returns(address[])
func (_TrafficMarket *TrafficMarketCallerSession) GetPoolList() ([]common.Address, error) {
	return _TrafficMarket.Contract.GetPoolList(&_TrafficMarket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TrafficMarket *TrafficMarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TrafficMarket.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TrafficMarket *TrafficMarketSession) Owner() (common.Address, error) {
	return _TrafficMarket.Contract.Owner(&_TrafficMarket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TrafficMarket *TrafficMarketCallerSession) Owner() (common.Address, error) {
	return _TrafficMarket.Contract.Owner(&_TrafficMarket.CallOpts)
}

// PayerForMiner is a free data retrieval call binding the contract method 0xce53c9bf.
//
// Solidity: function payerForMiner(address , bytes32 ) view returns(address)
func (_TrafficMarket *TrafficMarketCaller) PayerForMiner(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TrafficMarket.contract.Call(opts, &out, "payerForMiner", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PayerForMiner is a free data retrieval call binding the contract method 0xce53c9bf.
//
// Solidity: function payerForMiner(address , bytes32 ) view returns(address)
func (_TrafficMarket *TrafficMarketSession) PayerForMiner(arg0 common.Address, arg1 [32]byte) (common.Address, error) {
	return _TrafficMarket.Contract.PayerForMiner(&_TrafficMarket.CallOpts, arg0, arg1)
}

// PayerForMiner is a free data retrieval call binding the contract method 0xce53c9bf.
//
// Solidity: function payerForMiner(address , bytes32 ) view returns(address)
func (_TrafficMarket *TrafficMarketCallerSession) PayerForMiner(arg0 common.Address, arg1 [32]byte) (common.Address, error) {
	return _TrafficMarket.Contract.PayerForMiner(&_TrafficMarket.CallOpts, arg0, arg1)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TrafficMarket *TrafficMarketCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TrafficMarket.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TrafficMarket *TrafficMarketSession) Token() (common.Address, error) {
	return _TrafficMarket.Contract.Token(&_TrafficMarket.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TrafficMarket *TrafficMarketCallerSession) Token() (common.Address, error) {
	return _TrafficMarket.Contract.Token(&_TrafficMarket.CallOpts)
}

// TokenBalance is a free data retrieval call binding the contract method 0xeedc966a.
//
// Solidity: function tokenBalance(address userAddr) view returns(uint256, uint256, uint256)
func (_TrafficMarket *TrafficMarketCaller) TokenBalance(opts *bind.CallOpts, userAddr common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _TrafficMarket.contract.Call(opts, &out, "tokenBalance", userAddr)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// TokenBalance is a free data retrieval call binding the contract method 0xeedc966a.
//
// Solidity: function tokenBalance(address userAddr) view returns(uint256, uint256, uint256)
func (_TrafficMarket *TrafficMarketSession) TokenBalance(userAddr common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _TrafficMarket.Contract.TokenBalance(&_TrafficMarket.CallOpts, userAddr)
}

// TokenBalance is a free data retrieval call binding the contract method 0xeedc966a.
//
// Solidity: function tokenBalance(address userAddr) view returns(uint256, uint256, uint256)
func (_TrafficMarket *TrafficMarketCallerSession) TokenBalance(userAddr common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _TrafficMarket.Contract.TokenBalance(&_TrafficMarket.CallOpts, userAddr)
}

// ChangePool is a paid mutator transaction binding the contract method 0x6e84d055.
//
// Solidity: function changePool(address from, address to, bytes32 subAddr) returns()
func (_TrafficMarket *TrafficMarketTransactor) ChangePool(opts *bind.TransactOpts, from common.Address, to common.Address, subAddr [32]byte) (*types.Transaction, error) {
	return _TrafficMarket.contract.Transact(opts, "changePool", from, to, subAddr)
}

// ChangePool is a paid mutator transaction binding the contract method 0x6e84d055.
//
// Solidity: function changePool(address from, address to, bytes32 subAddr) returns()
func (_TrafficMarket *TrafficMarketSession) ChangePool(from common.Address, to common.Address, subAddr [32]byte) (*types.Transaction, error) {
	return _TrafficMarket.Contract.ChangePool(&_TrafficMarket.TransactOpts, from, to, subAddr)
}

// ChangePool is a paid mutator transaction binding the contract method 0x6e84d055.
//
// Solidity: function changePool(address from, address to, bytes32 subAddr) returns()
func (_TrafficMarket *TrafficMarketTransactorSession) ChangePool(from common.Address, to common.Address, subAddr [32]byte) (*types.Transaction, error) {
	return _TrafficMarket.Contract.ChangePool(&_TrafficMarket.TransactOpts, from, to, subAddr)
}

// ChangeSettings is a paid mutator transaction binding the contract method 0xa626c089.
//
// Solidity: function changeSettings(uint256 price, uint256 pDpos, uint256 mDpos) returns()
func (_TrafficMarket *TrafficMarketTransactor) ChangeSettings(opts *bind.TransactOpts, price *big.Int, pDpos *big.Int, mDpos *big.Int) (*types.Transaction, error) {
	return _TrafficMarket.contract.Transact(opts, "changeSettings", price, pDpos, mDpos)
}

// ChangeSettings is a paid mutator transaction binding the contract method 0xa626c089.
//
// Solidity: function changeSettings(uint256 price, uint256 pDpos, uint256 mDpos) returns()
func (_TrafficMarket *TrafficMarketSession) ChangeSettings(price *big.Int, pDpos *big.Int, mDpos *big.Int) (*types.Transaction, error) {
	return _TrafficMarket.Contract.ChangeSettings(&_TrafficMarket.TransactOpts, price, pDpos, mDpos)
}

// ChangeSettings is a paid mutator transaction binding the contract method 0xa626c089.
//
// Solidity: function changeSettings(uint256 price, uint256 pDpos, uint256 mDpos) returns()
func (_TrafficMarket *TrafficMarketTransactorSession) ChangeSettings(price *big.Int, pDpos *big.Int, mDpos *big.Int) (*types.Transaction, error) {
	return _TrafficMarket.Contract.ChangeSettings(&_TrafficMarket.TransactOpts, price, pDpos, mDpos)
}

// Charge is a paid mutator transaction binding the contract method 0xde5175cb.
//
// Solidity: function charge(address user, uint256 tokenNo, address poolAddr, uint256 index) returns()
func (_TrafficMarket *TrafficMarketTransactor) Charge(opts *bind.TransactOpts, user common.Address, tokenNo *big.Int, poolAddr common.Address, index *big.Int) (*types.Transaction, error) {
	return _TrafficMarket.contract.Transact(opts, "charge", user, tokenNo, poolAddr, index)
}

// Charge is a paid mutator transaction binding the contract method 0xde5175cb.
//
// Solidity: function charge(address user, uint256 tokenNo, address poolAddr, uint256 index) returns()
func (_TrafficMarket *TrafficMarketSession) Charge(user common.Address, tokenNo *big.Int, poolAddr common.Address, index *big.Int) (*types.Transaction, error) {
	return _TrafficMarket.Contract.Charge(&_TrafficMarket.TransactOpts, user, tokenNo, poolAddr, index)
}

// Charge is a paid mutator transaction binding the contract method 0xde5175cb.
//
// Solidity: function charge(address user, uint256 tokenNo, address poolAddr, uint256 index) returns()
func (_TrafficMarket *TrafficMarketTransactorSession) Charge(user common.Address, tokenNo *big.Int, poolAddr common.Address, index *big.Int) (*types.Transaction, error) {
	return _TrafficMarket.Contract.Charge(&_TrafficMarket.TransactOpts, user, tokenNo, poolAddr, index)
}

// DestroyPool is a paid mutator transaction binding the contract method 0x1bff1fcf.
//
// Solidity: function destroyPool(uint256 index) returns()
func (_TrafficMarket *TrafficMarketTransactor) DestroyPool(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _TrafficMarket.contract.Transact(opts, "destroyPool", index)
}

// DestroyPool is a paid mutator transaction binding the contract method 0x1bff1fcf.
//
// Solidity: function destroyPool(uint256 index) returns()
func (_TrafficMarket *TrafficMarketSession) DestroyPool(index *big.Int) (*types.Transaction, error) {
	return _TrafficMarket.Contract.DestroyPool(&_TrafficMarket.TransactOpts, index)
}

// DestroyPool is a paid mutator transaction binding the contract method 0x1bff1fcf.
//
// Solidity: function destroyPool(uint256 index) returns()
func (_TrafficMarket *TrafficMarketTransactorSession) DestroyPool(index *big.Int) (*types.Transaction, error) {
	return _TrafficMarket.Contract.DestroyPool(&_TrafficMarket.TransactOpts, index)
}

// Emergency is a paid mutator transaction binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() returns()
func (_TrafficMarket *TrafficMarketTransactor) Emergency(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrafficMarket.contract.Transact(opts, "emergency")
}

// Emergency is a paid mutator transaction binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() returns()
func (_TrafficMarket *TrafficMarketSession) Emergency() (*types.Transaction, error) {
	return _TrafficMarket.Contract.Emergency(&_TrafficMarket.TransactOpts)
}

// Emergency is a paid mutator transaction binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() returns()
func (_TrafficMarket *TrafficMarketTransactorSession) Emergency() (*types.Transaction, error) {
	return _TrafficMarket.Contract.Emergency(&_TrafficMarket.TransactOpts)
}

// JoinPool is a paid mutator transaction binding the contract method 0x94238a0b.
//
// Solidity: function joinPool(address poolAddr, uint256 index, bytes32 subAddr) returns()
func (_TrafficMarket *TrafficMarketTransactor) JoinPool(opts *bind.TransactOpts, poolAddr common.Address, index *big.Int, subAddr [32]byte) (*types.Transaction, error) {
	return _TrafficMarket.contract.Transact(opts, "joinPool", poolAddr, index, subAddr)
}

// JoinPool is a paid mutator transaction binding the contract method 0x94238a0b.
//
// Solidity: function joinPool(address poolAddr, uint256 index, bytes32 subAddr) returns()
func (_TrafficMarket *TrafficMarketSession) JoinPool(poolAddr common.Address, index *big.Int, subAddr [32]byte) (*types.Transaction, error) {
	return _TrafficMarket.Contract.JoinPool(&_TrafficMarket.TransactOpts, poolAddr, index, subAddr)
}

// JoinPool is a paid mutator transaction binding the contract method 0x94238a0b.
//
// Solidity: function joinPool(address poolAddr, uint256 index, bytes32 subAddr) returns()
func (_TrafficMarket *TrafficMarketTransactorSession) JoinPool(poolAddr common.Address, index *big.Int, subAddr [32]byte) (*types.Transaction, error) {
	return _TrafficMarket.Contract.JoinPool(&_TrafficMarket.TransactOpts, poolAddr, index, subAddr)
}

// Pclaim is a paid mutator transaction binding the contract method 0xc89531ef.
//
// Solidity: function pclaim(address user, address pool, uint256 minerCredit, uint256 minerAmount, uint256 usedTraffic, bytes signature) returns()
func (_TrafficMarket *TrafficMarketTransactor) Pclaim(opts *bind.TransactOpts, user common.Address, pool common.Address, minerCredit *big.Int, minerAmount *big.Int, usedTraffic *big.Int, signature []byte) (*types.Transaction, error) {
	return _TrafficMarket.contract.Transact(opts, "pclaim", user, pool, minerCredit, minerAmount, usedTraffic, signature)
}

// Pclaim is a paid mutator transaction binding the contract method 0xc89531ef.
//
// Solidity: function pclaim(address user, address pool, uint256 minerCredit, uint256 minerAmount, uint256 usedTraffic, bytes signature) returns()
func (_TrafficMarket *TrafficMarketSession) Pclaim(user common.Address, pool common.Address, minerCredit *big.Int, minerAmount *big.Int, usedTraffic *big.Int, signature []byte) (*types.Transaction, error) {
	return _TrafficMarket.Contract.Pclaim(&_TrafficMarket.TransactOpts, user, pool, minerCredit, minerAmount, usedTraffic, signature)
}

// Pclaim is a paid mutator transaction binding the contract method 0xc89531ef.
//
// Solidity: function pclaim(address user, address pool, uint256 minerCredit, uint256 minerAmount, uint256 usedTraffic, bytes signature) returns()
func (_TrafficMarket *TrafficMarketTransactorSession) Pclaim(user common.Address, pool common.Address, minerCredit *big.Int, minerAmount *big.Int, usedTraffic *big.Int, signature []byte) (*types.Transaction, error) {
	return _TrafficMarket.Contract.Pclaim(&_TrafficMarket.TransactOpts, user, pool, minerCredit, minerAmount, usedTraffic, signature)
}

// RegPool is a paid mutator transaction binding the contract method 0xbb219cb8.
//
// Solidity: function regPool() returns()
func (_TrafficMarket *TrafficMarketTransactor) RegPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrafficMarket.contract.Transact(opts, "regPool")
}

// RegPool is a paid mutator transaction binding the contract method 0xbb219cb8.
//
// Solidity: function regPool() returns()
func (_TrafficMarket *TrafficMarketSession) RegPool() (*types.Transaction, error) {
	return _TrafficMarket.Contract.RegPool(&_TrafficMarket.TransactOpts)
}

// RegPool is a paid mutator transaction binding the contract method 0xbb219cb8.
//
// Solidity: function regPool() returns()
func (_TrafficMarket *TrafficMarketTransactorSession) RegPool() (*types.Transaction, error) {
	return _TrafficMarket.Contract.RegPool(&_TrafficMarket.TransactOpts)
}

// RetireFromPool is a paid mutator transaction binding the contract method 0xd35aa93b.
//
// Solidity: function retireFromPool(address from, bytes32 subAddr) returns()
func (_TrafficMarket *TrafficMarketTransactor) RetireFromPool(opts *bind.TransactOpts, from common.Address, subAddr [32]byte) (*types.Transaction, error) {
	return _TrafficMarket.contract.Transact(opts, "retireFromPool", from, subAddr)
}

// RetireFromPool is a paid mutator transaction binding the contract method 0xd35aa93b.
//
// Solidity: function retireFromPool(address from, bytes32 subAddr) returns()
func (_TrafficMarket *TrafficMarketSession) RetireFromPool(from common.Address, subAddr [32]byte) (*types.Transaction, error) {
	return _TrafficMarket.Contract.RetireFromPool(&_TrafficMarket.TransactOpts, from, subAddr)
}

// RetireFromPool is a paid mutator transaction binding the contract method 0xd35aa93b.
//
// Solidity: function retireFromPool(address from, bytes32 subAddr) returns()
func (_TrafficMarket *TrafficMarketTransactorSession) RetireFromPool(from common.Address, subAddr [32]byte) (*types.Transaction, error) {
	return _TrafficMarket.Contract.RetireFromPool(&_TrafficMarket.TransactOpts, from, subAddr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TrafficMarket *TrafficMarketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TrafficMarket.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TrafficMarket *TrafficMarketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TrafficMarket.Contract.TransferOwnership(&_TrafficMarket.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TrafficMarket *TrafficMarketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TrafficMarket.Contract.TransferOwnership(&_TrafficMarket.TransactOpts, newOwner)
}

// TrafficMarketChargeIterator is returned from FilterCharge and is used to iterate over the raw logs and unpacked data for Charge events raised by the TrafficMarket contract.
type TrafficMarketChargeIterator struct {
	Event *TrafficMarketCharge // Event containing the contract specifics and raw log

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
func (it *TrafficMarketChargeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrafficMarketCharge)
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
		it.Event = new(TrafficMarketCharge)
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
func (it *TrafficMarketChargeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrafficMarketChargeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrafficMarketCharge represents a Charge event raised by the TrafficMarket contract.
type TrafficMarketCharge struct {
	User          common.Address
	Pool          common.Address
	TokenAmount   *big.Int
	TrafficAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCharge is a free log retrieval operation binding the contract event 0xeb140eb0b2b0c31ca6e7085dc4b6ffec039ba17c46444eb1c1e196b9d857330c.
//
// Solidity: event Charge(address indexed user, address indexed pool, uint256 tokenAmount, uint256 trafficAmount)
func (_TrafficMarket *TrafficMarketFilterer) FilterCharge(opts *bind.FilterOpts, user []common.Address, pool []common.Address) (*TrafficMarketChargeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _TrafficMarket.contract.FilterLogs(opts, "Charge", userRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &TrafficMarketChargeIterator{contract: _TrafficMarket.contract, event: "Charge", logs: logs, sub: sub}, nil
}

// WatchCharge is a free log subscription operation binding the contract event 0xeb140eb0b2b0c31ca6e7085dc4b6ffec039ba17c46444eb1c1e196b9d857330c.
//
// Solidity: event Charge(address indexed user, address indexed pool, uint256 tokenAmount, uint256 trafficAmount)
func (_TrafficMarket *TrafficMarketFilterer) WatchCharge(opts *bind.WatchOpts, sink chan<- *TrafficMarketCharge, user []common.Address, pool []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _TrafficMarket.contract.WatchLogs(opts, "Charge", userRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrafficMarketCharge)
				if err := _TrafficMarket.contract.UnpackLog(event, "Charge", log); err != nil {
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

// ParseCharge is a log parse operation binding the contract event 0xeb140eb0b2b0c31ca6e7085dc4b6ffec039ba17c46444eb1c1e196b9d857330c.
//
// Solidity: event Charge(address indexed user, address indexed pool, uint256 tokenAmount, uint256 trafficAmount)
func (_TrafficMarket *TrafficMarketFilterer) ParseCharge(log types.Log) (*TrafficMarketCharge, error) {
	event := new(TrafficMarketCharge)
	if err := _TrafficMarket.contract.UnpackLog(event, "Charge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrafficMarketMinerEventIterator is returned from FilterMinerEvent and is used to iterate over the raw logs and unpacked data for MinerEvent events raised by the TrafficMarket contract.
type TrafficMarketMinerEventIterator struct {
	Event *TrafficMarketMinerEvent // Event containing the contract specifics and raw log

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
func (it *TrafficMarketMinerEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrafficMarketMinerEvent)
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
		it.Event = new(TrafficMarketMinerEvent)
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
func (it *TrafficMarketMinerEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrafficMarketMinerEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrafficMarketMinerEvent represents a MinerEvent event raised by the TrafficMarket contract.
type TrafficMarketMinerEvent struct {
	SubAddr   [32]byte
	Addr1     common.Address
	Addr2     common.Address
	PayAddr   common.Address
	EventType uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinerEvent is a free log retrieval operation binding the contract event 0x4123878230821c2969dff198af302c656e21384764d292cac85b9ed0cc880aba.
//
// Solidity: event MinerEvent(bytes32 indexed subAddr, address indexed addr1, address addr2, address payAddr, uint8 eventType)
func (_TrafficMarket *TrafficMarketFilterer) FilterMinerEvent(opts *bind.FilterOpts, subAddr [][32]byte, addr1 []common.Address) (*TrafficMarketMinerEventIterator, error) {

	var subAddrRule []interface{}
	for _, subAddrItem := range subAddr {
		subAddrRule = append(subAddrRule, subAddrItem)
	}
	var addr1Rule []interface{}
	for _, addr1Item := range addr1 {
		addr1Rule = append(addr1Rule, addr1Item)
	}

	logs, sub, err := _TrafficMarket.contract.FilterLogs(opts, "MinerEvent", subAddrRule, addr1Rule)
	if err != nil {
		return nil, err
	}
	return &TrafficMarketMinerEventIterator{contract: _TrafficMarket.contract, event: "MinerEvent", logs: logs, sub: sub}, nil
}

// WatchMinerEvent is a free log subscription operation binding the contract event 0x4123878230821c2969dff198af302c656e21384764d292cac85b9ed0cc880aba.
//
// Solidity: event MinerEvent(bytes32 indexed subAddr, address indexed addr1, address addr2, address payAddr, uint8 eventType)
func (_TrafficMarket *TrafficMarketFilterer) WatchMinerEvent(opts *bind.WatchOpts, sink chan<- *TrafficMarketMinerEvent, subAddr [][32]byte, addr1 []common.Address) (event.Subscription, error) {

	var subAddrRule []interface{}
	for _, subAddrItem := range subAddr {
		subAddrRule = append(subAddrRule, subAddrItem)
	}
	var addr1Rule []interface{}
	for _, addr1Item := range addr1 {
		addr1Rule = append(addr1Rule, addr1Item)
	}

	logs, sub, err := _TrafficMarket.contract.WatchLogs(opts, "MinerEvent", subAddrRule, addr1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrafficMarketMinerEvent)
				if err := _TrafficMarket.contract.UnpackLog(event, "MinerEvent", log); err != nil {
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

// ParseMinerEvent is a log parse operation binding the contract event 0x4123878230821c2969dff198af302c656e21384764d292cac85b9ed0cc880aba.
//
// Solidity: event MinerEvent(bytes32 indexed subAddr, address indexed addr1, address addr2, address payAddr, uint8 eventType)
func (_TrafficMarket *TrafficMarketFilterer) ParseMinerEvent(log types.Log) (*TrafficMarketMinerEvent, error) {
	event := new(TrafficMarketMinerEvent)
	if err := _TrafficMarket.contract.UnpackLog(event, "MinerEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrafficMarketPoolClaimIterator is returned from FilterPoolClaim and is used to iterate over the raw logs and unpacked data for PoolClaim events raised by the TrafficMarket contract.
type TrafficMarketPoolClaimIterator struct {
	Event *TrafficMarketPoolClaim // Event containing the contract specifics and raw log

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
func (it *TrafficMarketPoolClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrafficMarketPoolClaim)
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
		it.Event = new(TrafficMarketPoolClaim)
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
func (it *TrafficMarketPoolClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrafficMarketPoolClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrafficMarketPoolClaim represents a PoolClaim event raised by the TrafficMarket contract.
type TrafficMarketPoolClaim struct {
	Pool            common.Address
	User            common.Address
	MinerUsedPacket *big.Int
	MinerPacket     *big.Int
	ClaimedBalance  *big.Int
	PoolTotalPacket *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPoolClaim is a free log retrieval operation binding the contract event 0xe249550089b34fd8690cecdc0bd02ac66196461e455ee12c3024e64b17d7f355.
//
// Solidity: event PoolClaim(address indexed pool, address indexed user, uint256 minerUsedPacket, uint256 minerPacket, uint256 claimedBalance, uint256 poolTotalPacket)
func (_TrafficMarket *TrafficMarketFilterer) FilterPoolClaim(opts *bind.FilterOpts, pool []common.Address, user []common.Address) (*TrafficMarketPoolClaimIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _TrafficMarket.contract.FilterLogs(opts, "PoolClaim", poolRule, userRule)
	if err != nil {
		return nil, err
	}
	return &TrafficMarketPoolClaimIterator{contract: _TrafficMarket.contract, event: "PoolClaim", logs: logs, sub: sub}, nil
}

// WatchPoolClaim is a free log subscription operation binding the contract event 0xe249550089b34fd8690cecdc0bd02ac66196461e455ee12c3024e64b17d7f355.
//
// Solidity: event PoolClaim(address indexed pool, address indexed user, uint256 minerUsedPacket, uint256 minerPacket, uint256 claimedBalance, uint256 poolTotalPacket)
func (_TrafficMarket *TrafficMarketFilterer) WatchPoolClaim(opts *bind.WatchOpts, sink chan<- *TrafficMarketPoolClaim, pool []common.Address, user []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _TrafficMarket.contract.WatchLogs(opts, "PoolClaim", poolRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrafficMarketPoolClaim)
				if err := _TrafficMarket.contract.UnpackLog(event, "PoolClaim", log); err != nil {
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

// ParsePoolClaim is a log parse operation binding the contract event 0xe249550089b34fd8690cecdc0bd02ac66196461e455ee12c3024e64b17d7f355.
//
// Solidity: event PoolClaim(address indexed pool, address indexed user, uint256 minerUsedPacket, uint256 minerPacket, uint256 claimedBalance, uint256 poolTotalPacket)
func (_TrafficMarket *TrafficMarketFilterer) ParsePoolClaim(log types.Log) (*TrafficMarketPoolClaim, error) {
	event := new(TrafficMarketPoolClaim)
	if err := _TrafficMarket.contract.UnpackLog(event, "PoolClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrafficMarketPoolRegIterator is returned from FilterPoolReg and is used to iterate over the raw logs and unpacked data for PoolReg events raised by the TrafficMarket contract.
type TrafficMarketPoolRegIterator struct {
	Event *TrafficMarketPoolReg // Event containing the contract specifics and raw log

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
func (it *TrafficMarketPoolRegIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrafficMarketPoolReg)
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
		it.Event = new(TrafficMarketPoolReg)
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
func (it *TrafficMarketPoolRegIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrafficMarketPoolRegIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrafficMarketPoolReg represents a PoolReg event raised by the TrafficMarket contract.
type TrafficMarketPoolReg struct {
	PoolAddr  common.Address
	EventType uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPoolReg is a free log retrieval operation binding the contract event 0xaee0e2cc630d522504357ec1c379562e52418bdac77905d8ea072a3dd3d9bba7.
//
// Solidity: event PoolReg(address poolAddr, uint8 eventType)
func (_TrafficMarket *TrafficMarketFilterer) FilterPoolReg(opts *bind.FilterOpts) (*TrafficMarketPoolRegIterator, error) {

	logs, sub, err := _TrafficMarket.contract.FilterLogs(opts, "PoolReg")
	if err != nil {
		return nil, err
	}
	return &TrafficMarketPoolRegIterator{contract: _TrafficMarket.contract, event: "PoolReg", logs: logs, sub: sub}, nil
}

// WatchPoolReg is a free log subscription operation binding the contract event 0xaee0e2cc630d522504357ec1c379562e52418bdac77905d8ea072a3dd3d9bba7.
//
// Solidity: event PoolReg(address poolAddr, uint8 eventType)
func (_TrafficMarket *TrafficMarketFilterer) WatchPoolReg(opts *bind.WatchOpts, sink chan<- *TrafficMarketPoolReg) (event.Subscription, error) {

	logs, sub, err := _TrafficMarket.contract.WatchLogs(opts, "PoolReg")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrafficMarketPoolReg)
				if err := _TrafficMarket.contract.UnpackLog(event, "PoolReg", log); err != nil {
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

// ParsePoolReg is a log parse operation binding the contract event 0xaee0e2cc630d522504357ec1c379562e52418bdac77905d8ea072a3dd3d9bba7.
//
// Solidity: event PoolReg(address poolAddr, uint8 eventType)
func (_TrafficMarket *TrafficMarketFilterer) ParsePoolReg(log types.Log) (*TrafficMarketPoolReg, error) {
	event := new(TrafficMarketPoolReg)
	if err := _TrafficMarket.contract.UnpackLog(event, "PoolReg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrafficMarketSettingsChangedIterator is returned from FilterSettingsChanged and is used to iterate over the raw logs and unpacked data for SettingsChanged events raised by the TrafficMarket contract.
type TrafficMarketSettingsChangedIterator struct {
	Event *TrafficMarketSettingsChanged // Event containing the contract specifics and raw log

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
func (it *TrafficMarketSettingsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrafficMarketSettingsChanged)
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
		it.Event = new(TrafficMarketSettingsChanged)
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
func (it *TrafficMarketSettingsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrafficMarketSettingsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrafficMarketSettingsChanged represents a SettingsChanged event raised by the TrafficMarket contract.
type TrafficMarketSettingsChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSettingsChanged is a free log retrieval operation binding the contract event 0x85a52584e2a604db198ab2f447cba2fc038a972048381c24b40219ef7a83df70.
//
// Solidity: event SettingsChanged()
func (_TrafficMarket *TrafficMarketFilterer) FilterSettingsChanged(opts *bind.FilterOpts) (*TrafficMarketSettingsChangedIterator, error) {

	logs, sub, err := _TrafficMarket.contract.FilterLogs(opts, "SettingsChanged")
	if err != nil {
		return nil, err
	}
	return &TrafficMarketSettingsChangedIterator{contract: _TrafficMarket.contract, event: "SettingsChanged", logs: logs, sub: sub}, nil
}

// WatchSettingsChanged is a free log subscription operation binding the contract event 0x85a52584e2a604db198ab2f447cba2fc038a972048381c24b40219ef7a83df70.
//
// Solidity: event SettingsChanged()
func (_TrafficMarket *TrafficMarketFilterer) WatchSettingsChanged(opts *bind.WatchOpts, sink chan<- *TrafficMarketSettingsChanged) (event.Subscription, error) {

	logs, sub, err := _TrafficMarket.contract.WatchLogs(opts, "SettingsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrafficMarketSettingsChanged)
				if err := _TrafficMarket.contract.UnpackLog(event, "SettingsChanged", log); err != nil {
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

// ParseSettingsChanged is a log parse operation binding the contract event 0x85a52584e2a604db198ab2f447cba2fc038a972048381c24b40219ef7a83df70.
//
// Solidity: event SettingsChanged()
func (_TrafficMarket *TrafficMarketFilterer) ParseSettingsChanged(log types.Log) (*TrafficMarketSettingsChanged, error) {
	event := new(TrafficMarketSettingsChanged)
	if err := _TrafficMarket.contract.UnpackLog(event, "SettingsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

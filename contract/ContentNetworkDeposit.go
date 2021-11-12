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

// ContentNetworkDepositABI is the input ABI used to generate the binding from.
const ContentNetworkDepositABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenNo\",\"type\":\"uint256\"}],\"name\":\"PoolWithDraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenNo\",\"type\":\"uint256\"}],\"name\":\"UserDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenNo\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"UserWithDraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CurrentBID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TokenDecimal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"billDetails\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenNo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dueDay\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"billLifeInDays\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"billListOfPool\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"billListOfUser\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenNo\",\"type\":\"uint256\"}],\"name\":\"depositForPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"hopBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"op\",\"type\":\"bool\"}],\"name\":\"openOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolIncomeWithDraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolRewardFor100Hops\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolsBillIDList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolsFeeIncome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"queryBillInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDays\",\"type\":\"uint256\"}],\"name\":\"setBillLifeInDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"}],\"name\":\"setPoolRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"}],\"name\":\"setUserRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userBillIDList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userRewardFor100Hops\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"withDrawFromPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ContentNetworkDeposit is an auto generated Go binding around an Ethereum contract.
type ContentNetworkDeposit struct {
	ContentNetworkDepositCaller     // Read-only binding to the contract
	ContentNetworkDepositTransactor // Write-only binding to the contract
	ContentNetworkDepositFilterer   // Log filterer for contract events
}

// ContentNetworkDepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContentNetworkDepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContentNetworkDepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContentNetworkDepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContentNetworkDepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContentNetworkDepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContentNetworkDepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContentNetworkDepositSession struct {
	Contract     *ContentNetworkDeposit // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ContentNetworkDepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContentNetworkDepositCallerSession struct {
	Contract *ContentNetworkDepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ContentNetworkDepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContentNetworkDepositTransactorSession struct {
	Contract     *ContentNetworkDepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ContentNetworkDepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContentNetworkDepositRaw struct {
	Contract *ContentNetworkDeposit // Generic contract binding to access the raw methods on
}

// ContentNetworkDepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContentNetworkDepositCallerRaw struct {
	Contract *ContentNetworkDepositCaller // Generic read-only contract binding to access the raw methods on
}

// ContentNetworkDepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContentNetworkDepositTransactorRaw struct {
	Contract *ContentNetworkDepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContentNetworkDeposit creates a new instance of ContentNetworkDeposit, bound to a specific deployed contract.
func NewContentNetworkDeposit(address common.Address, backend bind.ContractBackend) (*ContentNetworkDeposit, error) {
	contract, err := bindContentNetworkDeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContentNetworkDeposit{ContentNetworkDepositCaller: ContentNetworkDepositCaller{contract: contract}, ContentNetworkDepositTransactor: ContentNetworkDepositTransactor{contract: contract}, ContentNetworkDepositFilterer: ContentNetworkDepositFilterer{contract: contract}}, nil
}

// NewContentNetworkDepositCaller creates a new read-only instance of ContentNetworkDeposit, bound to a specific deployed contract.
func NewContentNetworkDepositCaller(address common.Address, caller bind.ContractCaller) (*ContentNetworkDepositCaller, error) {
	contract, err := bindContentNetworkDeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContentNetworkDepositCaller{contract: contract}, nil
}

// NewContentNetworkDepositTransactor creates a new write-only instance of ContentNetworkDeposit, bound to a specific deployed contract.
func NewContentNetworkDepositTransactor(address common.Address, transactor bind.ContractTransactor) (*ContentNetworkDepositTransactor, error) {
	contract, err := bindContentNetworkDeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContentNetworkDepositTransactor{contract: contract}, nil
}

// NewContentNetworkDepositFilterer creates a new log filterer instance of ContentNetworkDeposit, bound to a specific deployed contract.
func NewContentNetworkDepositFilterer(address common.Address, filterer bind.ContractFilterer) (*ContentNetworkDepositFilterer, error) {
	contract, err := bindContentNetworkDeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContentNetworkDepositFilterer{contract: contract}, nil
}

// bindContentNetworkDeposit binds a generic wrapper to an already deployed contract.
func bindContentNetworkDeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContentNetworkDepositABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContentNetworkDeposit *ContentNetworkDepositRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContentNetworkDeposit.Contract.ContentNetworkDepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContentNetworkDeposit *ContentNetworkDepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContentNetworkDeposit.Contract.ContentNetworkDepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContentNetworkDeposit *ContentNetworkDepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContentNetworkDeposit.Contract.ContentNetworkDepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContentNetworkDeposit *ContentNetworkDepositCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContentNetworkDeposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContentNetworkDeposit *ContentNetworkDepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContentNetworkDeposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContentNetworkDeposit *ContentNetworkDepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContentNetworkDeposit.Contract.contract.Transact(opts, method, params...)
}

// CurrentBID is a free data retrieval call binding the contract method 0x734a5251.
//
// Solidity: function CurrentBID() view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositCaller) CurrentBID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContentNetworkDeposit.contract.Call(opts, &out, "CurrentBID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentBID is a free data retrieval call binding the contract method 0x734a5251.
//
// Solidity: function CurrentBID() view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositSession) CurrentBID() (*big.Int, error) {
	return _ContentNetworkDeposit.Contract.CurrentBID(&_ContentNetworkDeposit.CallOpts)
}

// CurrentBID is a free data retrieval call binding the contract method 0x734a5251.
//
// Solidity: function CurrentBID() view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositCallerSession) CurrentBID() (*big.Int, error) {
	return _ContentNetworkDeposit.Contract.CurrentBID(&_ContentNetworkDeposit.CallOpts)
}

// TokenDecimal is a free data retrieval call binding the contract method 0xa969d1de.
//
// Solidity: function TokenDecimal() view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositCaller) TokenDecimal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContentNetworkDeposit.contract.Call(opts, &out, "TokenDecimal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenDecimal is a free data retrieval call binding the contract method 0xa969d1de.
//
// Solidity: function TokenDecimal() view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositSession) TokenDecimal() (*big.Int, error) {
	return _ContentNetworkDeposit.Contract.TokenDecimal(&_ContentNetworkDeposit.CallOpts)
}

// TokenDecimal is a free data retrieval call binding the contract method 0xa969d1de.
//
// Solidity: function TokenDecimal() view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositCallerSession) TokenDecimal() (*big.Int, error) {
	return _ContentNetworkDeposit.Contract.TokenDecimal(&_ContentNetworkDeposit.CallOpts)
}

// BillDetails is a free data retrieval call binding the contract method 0x994b6e8d.
//
// Solidity: function billDetails(uint256 ) view returns(address user, address pool, uint256 tokenNo, uint256 dueDay)
func (_ContentNetworkDeposit *ContentNetworkDepositCaller) BillDetails(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User    common.Address
	Pool    common.Address
	TokenNo *big.Int
	DueDay  *big.Int
}, error) {
	var out []interface{}
	err := _ContentNetworkDeposit.contract.Call(opts, &out, "billDetails", arg0)

	outstruct := new(struct {
		User    common.Address
		Pool    common.Address
		TokenNo *big.Int
		DueDay  *big.Int
	})

	outstruct.User = out[0].(common.Address)
	outstruct.Pool = out[1].(common.Address)
	outstruct.TokenNo = out[2].(*big.Int)
	outstruct.DueDay = out[3].(*big.Int)

	return *outstruct, err

}

// BillDetails is a free data retrieval call binding the contract method 0x994b6e8d.
//
// Solidity: function billDetails(uint256 ) view returns(address user, address pool, uint256 tokenNo, uint256 dueDay)
func (_ContentNetworkDeposit *ContentNetworkDepositSession) BillDetails(arg0 *big.Int) (struct {
	User    common.Address
	Pool    common.Address
	TokenNo *big.Int
	DueDay  *big.Int
}, error) {
	return _ContentNetworkDeposit.Contract.BillDetails(&_ContentNetworkDeposit.CallOpts, arg0)
}

// BillDetails is a free data retrieval call binding the contract method 0x994b6e8d.
//
// Solidity: function billDetails(uint256 ) view returns(address user, address pool, uint256 tokenNo, uint256 dueDay)
func (_ContentNetworkDeposit *ContentNetworkDepositCallerSession) BillDetails(arg0 *big.Int) (struct {
	User    common.Address
	Pool    common.Address
	TokenNo *big.Int
	DueDay  *big.Int
}, error) {
	return _ContentNetworkDeposit.Contract.BillDetails(&_ContentNetworkDeposit.CallOpts, arg0)
}

// BillLifeInDays is a free data retrieval call binding the contract method 0x3ef4c0c8.
//
// Solidity: function billLifeInDays() view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositCaller) BillLifeInDays(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContentNetworkDeposit.contract.Call(opts, &out, "billLifeInDays")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BillLifeInDays is a free data retrieval call binding the contract method 0x3ef4c0c8.
//
// Solidity: function billLifeInDays() view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositSession) BillLifeInDays() (*big.Int, error) {
	return _ContentNetworkDeposit.Contract.BillLifeInDays(&_ContentNetworkDeposit.CallOpts)
}

// BillLifeInDays is a free data retrieval call binding the contract method 0x3ef4c0c8.
//
// Solidity: function billLifeInDays() view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositCallerSession) BillLifeInDays() (*big.Int, error) {
	return _ContentNetworkDeposit.Contract.BillLifeInDays(&_ContentNetworkDeposit.CallOpts)
}

// BillListOfPool is a free data retrieval call binding the contract method 0x722d3146.
//
// Solidity: function billListOfPool(address pool) view returns(uint256[])
func (_ContentNetworkDeposit *ContentNetworkDepositCaller) BillListOfPool(opts *bind.CallOpts, pool common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ContentNetworkDeposit.contract.Call(opts, &out, "billListOfPool", pool)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BillListOfPool is a free data retrieval call binding the contract method 0x722d3146.
//
// Solidity: function billListOfPool(address pool) view returns(uint256[])
func (_ContentNetworkDeposit *ContentNetworkDepositSession) BillListOfPool(pool common.Address) ([]*big.Int, error) {
	return _ContentNetworkDeposit.Contract.BillListOfPool(&_ContentNetworkDeposit.CallOpts, pool)
}

// BillListOfPool is a free data retrieval call binding the contract method 0x722d3146.
//
// Solidity: function billListOfPool(address pool) view returns(uint256[])
func (_ContentNetworkDeposit *ContentNetworkDepositCallerSession) BillListOfPool(pool common.Address) ([]*big.Int, error) {
	return _ContentNetworkDeposit.Contract.BillListOfPool(&_ContentNetworkDeposit.CallOpts, pool)
}

// BillListOfUser is a free data retrieval call binding the contract method 0x349901f7.
//
// Solidity: function billListOfUser(address usr) view returns(uint256[])
func (_ContentNetworkDeposit *ContentNetworkDepositCaller) BillListOfUser(opts *bind.CallOpts, usr common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ContentNetworkDeposit.contract.Call(opts, &out, "billListOfUser", usr)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BillListOfUser is a free data retrieval call binding the contract method 0x349901f7.
//
// Solidity: function billListOfUser(address usr) view returns(uint256[])
func (_ContentNetworkDeposit *ContentNetworkDepositSession) BillListOfUser(usr common.Address) ([]*big.Int, error) {
	return _ContentNetworkDeposit.Contract.BillListOfUser(&_ContentNetworkDeposit.CallOpts, usr)
}

// BillListOfUser is a free data retrieval call binding the contract method 0x349901f7.
//
// Solidity: function billListOfUser(address usr) view returns(uint256[])
func (_ContentNetworkDeposit *ContentNetworkDepositCallerSession) BillListOfUser(usr common.Address) ([]*big.Int, error) {
	return _ContentNetworkDeposit.Contract.BillListOfUser(&_ContentNetworkDeposit.CallOpts, usr)
}

// HopBalanceOf is a free data retrieval call binding the contract method 0x6bf324d6.
//
// Solidity: function hopBalanceOf(address addr) view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositCaller) HopBalanceOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContentNetworkDeposit.contract.Call(opts, &out, "hopBalanceOf", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HopBalanceOf is a free data retrieval call binding the contract method 0x6bf324d6.
//
// Solidity: function hopBalanceOf(address addr) view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositSession) HopBalanceOf(addr common.Address) (*big.Int, error) {
	return _ContentNetworkDeposit.Contract.HopBalanceOf(&_ContentNetworkDeposit.CallOpts, addr)
}

// HopBalanceOf is a free data retrieval call binding the contract method 0x6bf324d6.
//
// Solidity: function hopBalanceOf(address addr) view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositCallerSession) HopBalanceOf(addr common.Address) (*big.Int, error) {
	return _ContentNetworkDeposit.Contract.HopBalanceOf(&_ContentNetworkDeposit.CallOpts, addr)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_ContentNetworkDeposit *ContentNetworkDepositCaller) IsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContentNetworkDeposit.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_ContentNetworkDeposit *ContentNetworkDepositSession) IsOpen() (bool, error) {
	return _ContentNetworkDeposit.Contract.IsOpen(&_ContentNetworkDeposit.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_ContentNetworkDeposit *ContentNetworkDepositCallerSession) IsOpen() (bool, error) {
	return _ContentNetworkDeposit.Contract.IsOpen(&_ContentNetworkDeposit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContentNetworkDeposit *ContentNetworkDepositCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContentNetworkDeposit.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContentNetworkDeposit *ContentNetworkDepositSession) Owner() (common.Address, error) {
	return _ContentNetworkDeposit.Contract.Owner(&_ContentNetworkDeposit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContentNetworkDeposit *ContentNetworkDepositCallerSession) Owner() (common.Address, error) {
	return _ContentNetworkDeposit.Contract.Owner(&_ContentNetworkDeposit.CallOpts)
}

// PoolRewardFor100Hops is a free data retrieval call binding the contract method 0xac4f7bcb.
//
// Solidity: function poolRewardFor100Hops() view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositCaller) PoolRewardFor100Hops(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContentNetworkDeposit.contract.Call(opts, &out, "poolRewardFor100Hops")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolRewardFor100Hops is a free data retrieval call binding the contract method 0xac4f7bcb.
//
// Solidity: function poolRewardFor100Hops() view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositSession) PoolRewardFor100Hops() (*big.Int, error) {
	return _ContentNetworkDeposit.Contract.PoolRewardFor100Hops(&_ContentNetworkDeposit.CallOpts)
}

// PoolRewardFor100Hops is a free data retrieval call binding the contract method 0xac4f7bcb.
//
// Solidity: function poolRewardFor100Hops() view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositCallerSession) PoolRewardFor100Hops() (*big.Int, error) {
	return _ContentNetworkDeposit.Contract.PoolRewardFor100Hops(&_ContentNetworkDeposit.CallOpts)
}

// PoolsBillIDList is a free data retrieval call binding the contract method 0xee7977f8.
//
// Solidity: function poolsBillIDList(address , uint256 ) view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositCaller) PoolsBillIDList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContentNetworkDeposit.contract.Call(opts, &out, "poolsBillIDList", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolsBillIDList is a free data retrieval call binding the contract method 0xee7977f8.
//
// Solidity: function poolsBillIDList(address , uint256 ) view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositSession) PoolsBillIDList(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _ContentNetworkDeposit.Contract.PoolsBillIDList(&_ContentNetworkDeposit.CallOpts, arg0, arg1)
}

// PoolsBillIDList is a free data retrieval call binding the contract method 0xee7977f8.
//
// Solidity: function poolsBillIDList(address , uint256 ) view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositCallerSession) PoolsBillIDList(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _ContentNetworkDeposit.Contract.PoolsBillIDList(&_ContentNetworkDeposit.CallOpts, arg0, arg1)
}

// PoolsFeeIncome is a free data retrieval call binding the contract method 0x735cd16a.
//
// Solidity: function poolsFeeIncome(address ) view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositCaller) PoolsFeeIncome(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContentNetworkDeposit.contract.Call(opts, &out, "poolsFeeIncome", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolsFeeIncome is a free data retrieval call binding the contract method 0x735cd16a.
//
// Solidity: function poolsFeeIncome(address ) view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositSession) PoolsFeeIncome(arg0 common.Address) (*big.Int, error) {
	return _ContentNetworkDeposit.Contract.PoolsFeeIncome(&_ContentNetworkDeposit.CallOpts, arg0)
}

// PoolsFeeIncome is a free data retrieval call binding the contract method 0x735cd16a.
//
// Solidity: function poolsFeeIncome(address ) view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositCallerSession) PoolsFeeIncome(arg0 common.Address) (*big.Int, error) {
	return _ContentNetworkDeposit.Contract.PoolsFeeIncome(&_ContentNetworkDeposit.CallOpts, arg0)
}

// QueryBillInterest is a free data retrieval call binding the contract method 0x1ca4ff5e.
//
// Solidity: function queryBillInterest(uint256 bid) view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositCaller) QueryBillInterest(opts *bind.CallOpts, bid *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContentNetworkDeposit.contract.Call(opts, &out, "queryBillInterest", bid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryBillInterest is a free data retrieval call binding the contract method 0x1ca4ff5e.
//
// Solidity: function queryBillInterest(uint256 bid) view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositSession) QueryBillInterest(bid *big.Int) (*big.Int, error) {
	return _ContentNetworkDeposit.Contract.QueryBillInterest(&_ContentNetworkDeposit.CallOpts, bid)
}

// QueryBillInterest is a free data retrieval call binding the contract method 0x1ca4ff5e.
//
// Solidity: function queryBillInterest(uint256 bid) view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositCallerSession) QueryBillInterest(bid *big.Int) (*big.Int, error) {
	return _ContentNetworkDeposit.Contract.QueryBillInterest(&_ContentNetworkDeposit.CallOpts, bid)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ContentNetworkDeposit *ContentNetworkDepositCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContentNetworkDeposit.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ContentNetworkDeposit *ContentNetworkDepositSession) Token() (common.Address, error) {
	return _ContentNetworkDeposit.Contract.Token(&_ContentNetworkDeposit.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ContentNetworkDeposit *ContentNetworkDepositCallerSession) Token() (common.Address, error) {
	return _ContentNetworkDeposit.Contract.Token(&_ContentNetworkDeposit.CallOpts)
}

// UserBillIDList is a free data retrieval call binding the contract method 0x5998ef68.
//
// Solidity: function userBillIDList(address , uint256 ) view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositCaller) UserBillIDList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContentNetworkDeposit.contract.Call(opts, &out, "userBillIDList", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBillIDList is a free data retrieval call binding the contract method 0x5998ef68.
//
// Solidity: function userBillIDList(address , uint256 ) view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositSession) UserBillIDList(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _ContentNetworkDeposit.Contract.UserBillIDList(&_ContentNetworkDeposit.CallOpts, arg0, arg1)
}

// UserBillIDList is a free data retrieval call binding the contract method 0x5998ef68.
//
// Solidity: function userBillIDList(address , uint256 ) view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositCallerSession) UserBillIDList(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _ContentNetworkDeposit.Contract.UserBillIDList(&_ContentNetworkDeposit.CallOpts, arg0, arg1)
}

// UserRewardFor100Hops is a free data retrieval call binding the contract method 0x8847c68c.
//
// Solidity: function userRewardFor100Hops() view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositCaller) UserRewardFor100Hops(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContentNetworkDeposit.contract.Call(opts, &out, "userRewardFor100Hops")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardFor100Hops is a free data retrieval call binding the contract method 0x8847c68c.
//
// Solidity: function userRewardFor100Hops() view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositSession) UserRewardFor100Hops() (*big.Int, error) {
	return _ContentNetworkDeposit.Contract.UserRewardFor100Hops(&_ContentNetworkDeposit.CallOpts)
}

// UserRewardFor100Hops is a free data retrieval call binding the contract method 0x8847c68c.
//
// Solidity: function userRewardFor100Hops() view returns(uint256)
func (_ContentNetworkDeposit *ContentNetworkDepositCallerSession) UserRewardFor100Hops() (*big.Int, error) {
	return _ContentNetworkDeposit.Contract.UserRewardFor100Hops(&_ContentNetworkDeposit.CallOpts)
}

// DepositForPool is a paid mutator transaction binding the contract method 0xe3ef49eb.
//
// Solidity: function depositForPool(address poolAddress, uint256 tokenNo) returns()
func (_ContentNetworkDeposit *ContentNetworkDepositTransactor) DepositForPool(opts *bind.TransactOpts, poolAddress common.Address, tokenNo *big.Int) (*types.Transaction, error) {
	return _ContentNetworkDeposit.contract.Transact(opts, "depositForPool", poolAddress, tokenNo)
}

// DepositForPool is a paid mutator transaction binding the contract method 0xe3ef49eb.
//
// Solidity: function depositForPool(address poolAddress, uint256 tokenNo) returns()
func (_ContentNetworkDeposit *ContentNetworkDepositSession) DepositForPool(poolAddress common.Address, tokenNo *big.Int) (*types.Transaction, error) {
	return _ContentNetworkDeposit.Contract.DepositForPool(&_ContentNetworkDeposit.TransactOpts, poolAddress, tokenNo)
}

// DepositForPool is a paid mutator transaction binding the contract method 0xe3ef49eb.
//
// Solidity: function depositForPool(address poolAddress, uint256 tokenNo) returns()
func (_ContentNetworkDeposit *ContentNetworkDepositTransactorSession) DepositForPool(poolAddress common.Address, tokenNo *big.Int) (*types.Transaction, error) {
	return _ContentNetworkDeposit.Contract.DepositForPool(&_ContentNetworkDeposit.TransactOpts, poolAddress, tokenNo)
}

// OpenOperation is a paid mutator transaction binding the contract method 0x4a446b75.
//
// Solidity: function openOperation(bool op) returns()
func (_ContentNetworkDeposit *ContentNetworkDepositTransactor) OpenOperation(opts *bind.TransactOpts, op bool) (*types.Transaction, error) {
	return _ContentNetworkDeposit.contract.Transact(opts, "openOperation", op)
}

// OpenOperation is a paid mutator transaction binding the contract method 0x4a446b75.
//
// Solidity: function openOperation(bool op) returns()
func (_ContentNetworkDeposit *ContentNetworkDepositSession) OpenOperation(op bool) (*types.Transaction, error) {
	return _ContentNetworkDeposit.Contract.OpenOperation(&_ContentNetworkDeposit.TransactOpts, op)
}

// OpenOperation is a paid mutator transaction binding the contract method 0x4a446b75.
//
// Solidity: function openOperation(bool op) returns()
func (_ContentNetworkDeposit *ContentNetworkDepositTransactorSession) OpenOperation(op bool) (*types.Transaction, error) {
	return _ContentNetworkDeposit.Contract.OpenOperation(&_ContentNetworkDeposit.TransactOpts, op)
}

// PoolIncomeWithDraw is a paid mutator transaction binding the contract method 0xe2edfdb0.
//
// Solidity: function poolIncomeWithDraw() returns()
func (_ContentNetworkDeposit *ContentNetworkDepositTransactor) PoolIncomeWithDraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContentNetworkDeposit.contract.Transact(opts, "poolIncomeWithDraw")
}

// PoolIncomeWithDraw is a paid mutator transaction binding the contract method 0xe2edfdb0.
//
// Solidity: function poolIncomeWithDraw() returns()
func (_ContentNetworkDeposit *ContentNetworkDepositSession) PoolIncomeWithDraw() (*types.Transaction, error) {
	return _ContentNetworkDeposit.Contract.PoolIncomeWithDraw(&_ContentNetworkDeposit.TransactOpts)
}

// PoolIncomeWithDraw is a paid mutator transaction binding the contract method 0xe2edfdb0.
//
// Solidity: function poolIncomeWithDraw() returns()
func (_ContentNetworkDeposit *ContentNetworkDepositTransactorSession) PoolIncomeWithDraw() (*types.Transaction, error) {
	return _ContentNetworkDeposit.Contract.PoolIncomeWithDraw(&_ContentNetworkDeposit.TransactOpts)
}

// SetBillLifeInDay is a paid mutator transaction binding the contract method 0xe0514f2e.
//
// Solidity: function setBillLifeInDay(uint256 newDays) returns()
func (_ContentNetworkDeposit *ContentNetworkDepositTransactor) SetBillLifeInDay(opts *bind.TransactOpts, newDays *big.Int) (*types.Transaction, error) {
	return _ContentNetworkDeposit.contract.Transact(opts, "setBillLifeInDay", newDays)
}

// SetBillLifeInDay is a paid mutator transaction binding the contract method 0xe0514f2e.
//
// Solidity: function setBillLifeInDay(uint256 newDays) returns()
func (_ContentNetworkDeposit *ContentNetworkDepositSession) SetBillLifeInDay(newDays *big.Int) (*types.Transaction, error) {
	return _ContentNetworkDeposit.Contract.SetBillLifeInDay(&_ContentNetworkDeposit.TransactOpts, newDays)
}

// SetBillLifeInDay is a paid mutator transaction binding the contract method 0xe0514f2e.
//
// Solidity: function setBillLifeInDay(uint256 newDays) returns()
func (_ContentNetworkDeposit *ContentNetworkDepositTransactorSession) SetBillLifeInDay(newDays *big.Int) (*types.Transaction, error) {
	return _ContentNetworkDeposit.Contract.SetBillLifeInDay(&_ContentNetworkDeposit.TransactOpts, newDays)
}

// SetPoolRate is a paid mutator transaction binding the contract method 0xb95aff28.
//
// Solidity: function setPoolRate(uint256 newRate) returns()
func (_ContentNetworkDeposit *ContentNetworkDepositTransactor) SetPoolRate(opts *bind.TransactOpts, newRate *big.Int) (*types.Transaction, error) {
	return _ContentNetworkDeposit.contract.Transact(opts, "setPoolRate", newRate)
}

// SetPoolRate is a paid mutator transaction binding the contract method 0xb95aff28.
//
// Solidity: function setPoolRate(uint256 newRate) returns()
func (_ContentNetworkDeposit *ContentNetworkDepositSession) SetPoolRate(newRate *big.Int) (*types.Transaction, error) {
	return _ContentNetworkDeposit.Contract.SetPoolRate(&_ContentNetworkDeposit.TransactOpts, newRate)
}

// SetPoolRate is a paid mutator transaction binding the contract method 0xb95aff28.
//
// Solidity: function setPoolRate(uint256 newRate) returns()
func (_ContentNetworkDeposit *ContentNetworkDepositTransactorSession) SetPoolRate(newRate *big.Int) (*types.Transaction, error) {
	return _ContentNetworkDeposit.Contract.SetPoolRate(&_ContentNetworkDeposit.TransactOpts, newRate)
}

// SetUserRate is a paid mutator transaction binding the contract method 0x36422e54.
//
// Solidity: function setUserRate(uint256 newRate) returns()
func (_ContentNetworkDeposit *ContentNetworkDepositTransactor) SetUserRate(opts *bind.TransactOpts, newRate *big.Int) (*types.Transaction, error) {
	return _ContentNetworkDeposit.contract.Transact(opts, "setUserRate", newRate)
}

// SetUserRate is a paid mutator transaction binding the contract method 0x36422e54.
//
// Solidity: function setUserRate(uint256 newRate) returns()
func (_ContentNetworkDeposit *ContentNetworkDepositSession) SetUserRate(newRate *big.Int) (*types.Transaction, error) {
	return _ContentNetworkDeposit.Contract.SetUserRate(&_ContentNetworkDeposit.TransactOpts, newRate)
}

// SetUserRate is a paid mutator transaction binding the contract method 0x36422e54.
//
// Solidity: function setUserRate(uint256 newRate) returns()
func (_ContentNetworkDeposit *ContentNetworkDepositTransactorSession) SetUserRate(newRate *big.Int) (*types.Transaction, error) {
	return _ContentNetworkDeposit.Contract.SetUserRate(&_ContentNetworkDeposit.TransactOpts, newRate)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContentNetworkDeposit *ContentNetworkDepositTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContentNetworkDeposit.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContentNetworkDeposit *ContentNetworkDepositSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContentNetworkDeposit.Contract.TransferOwnership(&_ContentNetworkDeposit.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContentNetworkDeposit *ContentNetworkDepositTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContentNetworkDeposit.Contract.TransferOwnership(&_ContentNetworkDeposit.TransactOpts, newOwner)
}

// WithDrawFromPool is a paid mutator transaction binding the contract method 0x0f335151.
//
// Solidity: function withDrawFromPool(uint256 bid) returns()
func (_ContentNetworkDeposit *ContentNetworkDepositTransactor) WithDrawFromPool(opts *bind.TransactOpts, bid *big.Int) (*types.Transaction, error) {
	return _ContentNetworkDeposit.contract.Transact(opts, "withDrawFromPool", bid)
}

// WithDrawFromPool is a paid mutator transaction binding the contract method 0x0f335151.
//
// Solidity: function withDrawFromPool(uint256 bid) returns()
func (_ContentNetworkDeposit *ContentNetworkDepositSession) WithDrawFromPool(bid *big.Int) (*types.Transaction, error) {
	return _ContentNetworkDeposit.Contract.WithDrawFromPool(&_ContentNetworkDeposit.TransactOpts, bid)
}

// WithDrawFromPool is a paid mutator transaction binding the contract method 0x0f335151.
//
// Solidity: function withDrawFromPool(uint256 bid) returns()
func (_ContentNetworkDeposit *ContentNetworkDepositTransactorSession) WithDrawFromPool(bid *big.Int) (*types.Transaction, error) {
	return _ContentNetworkDeposit.Contract.WithDrawFromPool(&_ContentNetworkDeposit.TransactOpts, bid)
}

// ContentNetworkDepositPoolWithDrawIterator is returned from FilterPoolWithDraw and is used to iterate over the raw logs and unpacked data for PoolWithDraw events raised by the ContentNetworkDeposit contract.
type ContentNetworkDepositPoolWithDrawIterator struct {
	Event *ContentNetworkDepositPoolWithDraw // Event containing the contract specifics and raw log

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
func (it *ContentNetworkDepositPoolWithDrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentNetworkDepositPoolWithDraw)
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
		it.Event = new(ContentNetworkDepositPoolWithDraw)
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
func (it *ContentNetworkDepositPoolWithDrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentNetworkDepositPoolWithDrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentNetworkDepositPoolWithDraw represents a PoolWithDraw event raised by the ContentNetworkDeposit contract.
type ContentNetworkDepositPoolWithDraw struct {
	Pool    common.Address
	TokenNo *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPoolWithDraw is a free log retrieval operation binding the contract event 0x4de6c88a1e2a02c5d8b364c864d831610041d05c6c588bf38893571abca215f8.
//
// Solidity: event PoolWithDraw(address indexed pool, uint256 tokenNo)
func (_ContentNetworkDeposit *ContentNetworkDepositFilterer) FilterPoolWithDraw(opts *bind.FilterOpts, pool []common.Address) (*ContentNetworkDepositPoolWithDrawIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _ContentNetworkDeposit.contract.FilterLogs(opts, "PoolWithDraw", poolRule)
	if err != nil {
		return nil, err
	}
	return &ContentNetworkDepositPoolWithDrawIterator{contract: _ContentNetworkDeposit.contract, event: "PoolWithDraw", logs: logs, sub: sub}, nil
}

// WatchPoolWithDraw is a free log subscription operation binding the contract event 0x4de6c88a1e2a02c5d8b364c864d831610041d05c6c588bf38893571abca215f8.
//
// Solidity: event PoolWithDraw(address indexed pool, uint256 tokenNo)
func (_ContentNetworkDeposit *ContentNetworkDepositFilterer) WatchPoolWithDraw(opts *bind.WatchOpts, sink chan<- *ContentNetworkDepositPoolWithDraw, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _ContentNetworkDeposit.contract.WatchLogs(opts, "PoolWithDraw", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentNetworkDepositPoolWithDraw)
				if err := _ContentNetworkDeposit.contract.UnpackLog(event, "PoolWithDraw", log); err != nil {
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

// ParsePoolWithDraw is a log parse operation binding the contract event 0x4de6c88a1e2a02c5d8b364c864d831610041d05c6c588bf38893571abca215f8.
//
// Solidity: event PoolWithDraw(address indexed pool, uint256 tokenNo)
func (_ContentNetworkDeposit *ContentNetworkDepositFilterer) ParsePoolWithDraw(log types.Log) (*ContentNetworkDepositPoolWithDraw, error) {
	event := new(ContentNetworkDepositPoolWithDraw)
	if err := _ContentNetworkDeposit.contract.UnpackLog(event, "PoolWithDraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContentNetworkDepositUserDepositIterator is returned from FilterUserDeposit and is used to iterate over the raw logs and unpacked data for UserDeposit events raised by the ContentNetworkDeposit contract.
type ContentNetworkDepositUserDepositIterator struct {
	Event *ContentNetworkDepositUserDeposit // Event containing the contract specifics and raw log

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
func (it *ContentNetworkDepositUserDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentNetworkDepositUserDeposit)
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
		it.Event = new(ContentNetworkDepositUserDeposit)
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
func (it *ContentNetworkDepositUserDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentNetworkDepositUserDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentNetworkDepositUserDeposit represents a UserDeposit event raised by the ContentNetworkDeposit contract.
type ContentNetworkDepositUserDeposit struct {
	Pool    common.Address
	User    common.Address
	TokenNo *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUserDeposit is a free log retrieval operation binding the contract event 0x3bc57f469ad6d10d7723ea226cd22bd2b9e527def2b529f6ab44645a16689582.
//
// Solidity: event UserDeposit(address indexed pool, address indexed user, uint256 tokenNo)
func (_ContentNetworkDeposit *ContentNetworkDepositFilterer) FilterUserDeposit(opts *bind.FilterOpts, pool []common.Address, user []common.Address) (*ContentNetworkDepositUserDepositIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContentNetworkDeposit.contract.FilterLogs(opts, "UserDeposit", poolRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ContentNetworkDepositUserDepositIterator{contract: _ContentNetworkDeposit.contract, event: "UserDeposit", logs: logs, sub: sub}, nil
}

// WatchUserDeposit is a free log subscription operation binding the contract event 0x3bc57f469ad6d10d7723ea226cd22bd2b9e527def2b529f6ab44645a16689582.
//
// Solidity: event UserDeposit(address indexed pool, address indexed user, uint256 tokenNo)
func (_ContentNetworkDeposit *ContentNetworkDepositFilterer) WatchUserDeposit(opts *bind.WatchOpts, sink chan<- *ContentNetworkDepositUserDeposit, pool []common.Address, user []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContentNetworkDeposit.contract.WatchLogs(opts, "UserDeposit", poolRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentNetworkDepositUserDeposit)
				if err := _ContentNetworkDeposit.contract.UnpackLog(event, "UserDeposit", log); err != nil {
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

// ParseUserDeposit is a log parse operation binding the contract event 0x3bc57f469ad6d10d7723ea226cd22bd2b9e527def2b529f6ab44645a16689582.
//
// Solidity: event UserDeposit(address indexed pool, address indexed user, uint256 tokenNo)
func (_ContentNetworkDeposit *ContentNetworkDepositFilterer) ParseUserDeposit(log types.Log) (*ContentNetworkDepositUserDeposit, error) {
	event := new(ContentNetworkDepositUserDeposit)
	if err := _ContentNetworkDeposit.contract.UnpackLog(event, "UserDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContentNetworkDepositUserWithDrawIterator is returned from FilterUserWithDraw and is used to iterate over the raw logs and unpacked data for UserWithDraw events raised by the ContentNetworkDeposit contract.
type ContentNetworkDepositUserWithDrawIterator struct {
	Event *ContentNetworkDepositUserWithDraw // Event containing the contract specifics and raw log

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
func (it *ContentNetworkDepositUserWithDrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentNetworkDepositUserWithDraw)
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
		it.Event = new(ContentNetworkDepositUserWithDraw)
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
func (it *ContentNetworkDepositUserWithDrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentNetworkDepositUserWithDrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentNetworkDepositUserWithDraw represents a UserWithDraw event raised by the ContentNetworkDeposit contract.
type ContentNetworkDepositUserWithDraw struct {
	Pool    common.Address
	User    common.Address
	TokenNo *big.Int
	Rewards *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUserWithDraw is a free log retrieval operation binding the contract event 0x0406c2d1e9159dd57e0577331e73a14f5a5249d1caea0a272021b41233e6d371.
//
// Solidity: event UserWithDraw(address indexed pool, address indexed user, uint256 tokenNo, uint256 rewards)
func (_ContentNetworkDeposit *ContentNetworkDepositFilterer) FilterUserWithDraw(opts *bind.FilterOpts, pool []common.Address, user []common.Address) (*ContentNetworkDepositUserWithDrawIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContentNetworkDeposit.contract.FilterLogs(opts, "UserWithDraw", poolRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ContentNetworkDepositUserWithDrawIterator{contract: _ContentNetworkDeposit.contract, event: "UserWithDraw", logs: logs, sub: sub}, nil
}

// WatchUserWithDraw is a free log subscription operation binding the contract event 0x0406c2d1e9159dd57e0577331e73a14f5a5249d1caea0a272021b41233e6d371.
//
// Solidity: event UserWithDraw(address indexed pool, address indexed user, uint256 tokenNo, uint256 rewards)
func (_ContentNetworkDeposit *ContentNetworkDepositFilterer) WatchUserWithDraw(opts *bind.WatchOpts, sink chan<- *ContentNetworkDepositUserWithDraw, pool []common.Address, user []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContentNetworkDeposit.contract.WatchLogs(opts, "UserWithDraw", poolRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentNetworkDepositUserWithDraw)
				if err := _ContentNetworkDeposit.contract.UnpackLog(event, "UserWithDraw", log); err != nil {
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

// ParseUserWithDraw is a log parse operation binding the contract event 0x0406c2d1e9159dd57e0577331e73a14f5a5249d1caea0a272021b41233e6d371.
//
// Solidity: event UserWithDraw(address indexed pool, address indexed user, uint256 tokenNo, uint256 rewards)
func (_ContentNetworkDeposit *ContentNetworkDepositFilterer) ParseUserWithDraw(log types.Log) (*ContentNetworkDepositUserWithDraw, error) {
	event := new(ContentNetworkDepositUserWithDraw)
	if err := _ContentNetworkDeposit.contract.UnpackLog(event, "UserWithDraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

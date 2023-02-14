// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package slite

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
)

// StcImDaoNode is an auto generated low-level Go binding around an user-defined struct.
type StcImDaoNode struct {
	NodeId   string
	NodeInfo string
	IsUsed   bool
}

// SliteMetaData contains all meta data concerning the Slite contract.
var SliteMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getNodeList\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"nodeInfo\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isUsed\",\"type\":\"bool\"}],\"internalType\":\"structStcImDao.Node[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"}],\"name\":\"isStake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"nodeInfo\",\"type\":\"string\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SliteABI is the input ABI used to generate the binding from.
// Deprecated: Use SliteMetaData.ABI instead.
var SliteABI = SliteMetaData.ABI

// Slite is an auto generated Go binding around an Ethereum contract.
type Slite struct {
	SliteCaller     // Read-only binding to the contract
	SliteTransactor // Write-only binding to the contract
	SliteFilterer   // Log filterer for contract events
}

// SliteCaller is an auto generated read-only Go binding around an Ethereum contract.
type SliteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SliteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SliteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SliteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SliteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SliteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SliteSession struct {
	Contract     *Slite            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SliteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SliteCallerSession struct {
	Contract *SliteCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SliteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SliteTransactorSession struct {
	Contract     *SliteTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SliteRaw is an auto generated low-level Go binding around an Ethereum contract.
type SliteRaw struct {
	Contract *Slite // Generic contract binding to access the raw methods on
}

// SliteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SliteCallerRaw struct {
	Contract *SliteCaller // Generic read-only contract binding to access the raw methods on
}

// SliteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SliteTransactorRaw struct {
	Contract *SliteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlite creates a new instance of Slite, bound to a specific deployed contract.
func NewSlite(address common.Address, backend bind.ContractBackend) (*Slite, error) {
	contract, err := bindSlite(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Slite{SliteCaller: SliteCaller{contract: contract}, SliteTransactor: SliteTransactor{contract: contract}, SliteFilterer: SliteFilterer{contract: contract}}, nil
}

// NewSliteCaller creates a new read-only instance of Slite, bound to a specific deployed contract.
func NewSliteCaller(address common.Address, caller bind.ContractCaller) (*SliteCaller, error) {
	contract, err := bindSlite(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SliteCaller{contract: contract}, nil
}

// NewSliteTransactor creates a new write-only instance of Slite, bound to a specific deployed contract.
func NewSliteTransactor(address common.Address, transactor bind.ContractTransactor) (*SliteTransactor, error) {
	contract, err := bindSlite(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SliteTransactor{contract: contract}, nil
}

// NewSliteFilterer creates a new log filterer instance of Slite, bound to a specific deployed contract.
func NewSliteFilterer(address common.Address, filterer bind.ContractFilterer) (*SliteFilterer, error) {
	contract, err := bindSlite(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SliteFilterer{contract: contract}, nil
}

// bindSlite binds a generic wrapper to an already deployed contract.
func bindSlite(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SliteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Slite *SliteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Slite.Contract.SliteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Slite *SliteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Slite.Contract.SliteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Slite *SliteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Slite.Contract.SliteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Slite *SliteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Slite.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Slite *SliteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Slite.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Slite *SliteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Slite.Contract.contract.Transact(opts, method, params...)
}

// GetNodeList is a free data retrieval call binding the contract method 0x53f3b713.
//
// Solidity: function getNodeList() view returns((string,string,bool)[])
func (_Slite *SliteCaller) GetNodeList(opts *bind.CallOpts) ([]StcImDaoNode, error) {
	var out []interface{}
	err := _Slite.contract.Call(opts, &out, "getNodeList")

	if err != nil {
		return *new([]StcImDaoNode), err
	}

	out0 := *abi.ConvertType(out[0], new([]StcImDaoNode)).(*[]StcImDaoNode)

	return out0, err

}

// GetNodeList is a free data retrieval call binding the contract method 0x53f3b713.
//
// Solidity: function getNodeList() view returns((string,string,bool)[])
func (_Slite *SliteSession) GetNodeList() ([]StcImDaoNode, error) {
	return _Slite.Contract.GetNodeList(&_Slite.CallOpts)
}

// GetNodeList is a free data retrieval call binding the contract method 0x53f3b713.
//
// Solidity: function getNodeList() view returns((string,string,bool)[])
func (_Slite *SliteCallerSession) GetNodeList() ([]StcImDaoNode, error) {
	return _Slite.Contract.GetNodeList(&_Slite.CallOpts)
}

// IsStake is a free data retrieval call binding the contract method 0x49c7cc35.
//
// Solidity: function isStake(string nodeId) view returns(bool)
func (_Slite *SliteCaller) IsStake(opts *bind.CallOpts, nodeId string) (bool, error) {
	var out []interface{}
	err := _Slite.contract.Call(opts, &out, "isStake", nodeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStake is a free data retrieval call binding the contract method 0x49c7cc35.
//
// Solidity: function isStake(string nodeId) view returns(bool)
func (_Slite *SliteSession) IsStake(nodeId string) (bool, error) {
	return _Slite.Contract.IsStake(&_Slite.CallOpts, nodeId)
}

// IsStake is a free data retrieval call binding the contract method 0x49c7cc35.
//
// Solidity: function isStake(string nodeId) view returns(bool)
func (_Slite *SliteCallerSession) IsStake(nodeId string) (bool, error) {
	return _Slite.Contract.IsStake(&_Slite.CallOpts, nodeId)
}

// Stake is a paid mutator transaction binding the contract method 0x19b36b1d.
//
// Solidity: function stake(string nodeId, string nodeInfo) returns()
func (_Slite *SliteTransactor) Stake(opts *bind.TransactOpts, nodeId string, nodeInfo string) (*types.Transaction, error) {
	return _Slite.contract.Transact(opts, "stake", nodeId, nodeInfo)
}

// Stake is a paid mutator transaction binding the contract method 0x19b36b1d.
//
// Solidity: function stake(string nodeId, string nodeInfo) returns()
func (_Slite *SliteSession) Stake(nodeId string, nodeInfo string) (*types.Transaction, error) {
	return _Slite.Contract.Stake(&_Slite.TransactOpts, nodeId, nodeInfo)
}

// Stake is a paid mutator transaction binding the contract method 0x19b36b1d.
//
// Solidity: function stake(string nodeId, string nodeInfo) returns()
func (_Slite *SliteTransactorSession) Stake(nodeId string, nodeInfo string) (*types.Transaction, error) {
	return _Slite.Contract.Stake(&_Slite.TransactOpts, nodeId, nodeInfo)
}

// Withdraw is a paid mutator transaction binding the contract method 0x31fb67c2.
//
// Solidity: function withdraw(string nodeId) returns()
func (_Slite *SliteTransactor) Withdraw(opts *bind.TransactOpts, nodeId string) (*types.Transaction, error) {
	return _Slite.contract.Transact(opts, "withdraw", nodeId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x31fb67c2.
//
// Solidity: function withdraw(string nodeId) returns()
func (_Slite *SliteSession) Withdraw(nodeId string) (*types.Transaction, error) {
	return _Slite.Contract.Withdraw(&_Slite.TransactOpts, nodeId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x31fb67c2.
//
// Solidity: function withdraw(string nodeId) returns()
func (_Slite *SliteTransactorSession) Withdraw(nodeId string) (*types.Transaction, error) {
	return _Slite.Contract.Withdraw(&_Slite.TransactOpts, nodeId)
}

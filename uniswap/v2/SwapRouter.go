// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package v2

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
	_ = abi.ConvertType
)

// SwapRouterMetaData contains all meta data concerning the SwapRouter contract.
var SwapRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"}],\"name\":\"v2SwapExactInput\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"}],\"name\":\"v2SwapExactOutput\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SwapRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapRouterMetaData.ABI instead.
var SwapRouterABI = SwapRouterMetaData.ABI

// SwapRouter is an auto generated Go binding around an Ethereum contract.
type SwapRouter struct {
	SwapRouterCaller     // Read-only binding to the contract
	SwapRouterTransactor // Write-only binding to the contract
	SwapRouterFilterer   // Log filterer for contract events
}

// SwapRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapRouterSession struct {
	Contract     *SwapRouter       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapRouterCallerSession struct {
	Contract *SwapRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SwapRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapRouterTransactorSession struct {
	Contract     *SwapRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SwapRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapRouterRaw struct {
	Contract *SwapRouter // Generic contract binding to access the raw methods on
}

// SwapRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapRouterCallerRaw struct {
	Contract *SwapRouterCaller // Generic read-only contract binding to access the raw methods on
}

// SwapRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapRouterTransactorRaw struct {
	Contract *SwapRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapRouter creates a new instance of SwapRouter, bound to a specific deployed contract.
func NewSwapRouter(address common.Address, backend bind.ContractBackend) (*SwapRouter, error) {
	contract, err := bindSwapRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapRouter{SwapRouterCaller: SwapRouterCaller{contract: contract}, SwapRouterTransactor: SwapRouterTransactor{contract: contract}, SwapRouterFilterer: SwapRouterFilterer{contract: contract}}, nil
}

// NewSwapRouterCaller creates a new read-only instance of SwapRouter, bound to a specific deployed contract.
func NewSwapRouterCaller(address common.Address, caller bind.ContractCaller) (*SwapRouterCaller, error) {
	contract, err := bindSwapRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapRouterCaller{contract: contract}, nil
}

// NewSwapRouterTransactor creates a new write-only instance of SwapRouter, bound to a specific deployed contract.
func NewSwapRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapRouterTransactor, error) {
	contract, err := bindSwapRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapRouterTransactor{contract: contract}, nil
}

// NewSwapRouterFilterer creates a new log filterer instance of SwapRouter, bound to a specific deployed contract.
func NewSwapRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapRouterFilterer, error) {
	contract, err := bindSwapRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapRouterFilterer{contract: contract}, nil
}

// bindSwapRouter binds a generic wrapper to an already deployed contract.
func bindSwapRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwapRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapRouter *SwapRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapRouter.Contract.SwapRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapRouter *SwapRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapRouter.Contract.SwapRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapRouter *SwapRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapRouter.Contract.SwapRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapRouter *SwapRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapRouter *SwapRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapRouter *SwapRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapRouter.Contract.contract.Transact(opts, method, params...)
}

// V2SwapExactInput is a paid mutator transaction binding the contract method 0xb1542aef.
//
// Solidity: function v2SwapExactInput(address recipient, uint256 amountIn, uint256 amountOutMinimum, address[] path, address payer) returns()
func (_SwapRouter *SwapRouterTransactor) V2SwapExactInput(opts *bind.TransactOpts, recipient common.Address, amountIn *big.Int, amountOutMinimum *big.Int, path []common.Address, payer common.Address) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "v2SwapExactInput", recipient, amountIn, amountOutMinimum, path, payer)
}

// V2SwapExactInput is a paid mutator transaction binding the contract method 0xb1542aef.
//
// Solidity: function v2SwapExactInput(address recipient, uint256 amountIn, uint256 amountOutMinimum, address[] path, address payer) returns()
func (_SwapRouter *SwapRouterSession) V2SwapExactInput(recipient common.Address, amountIn *big.Int, amountOutMinimum *big.Int, path []common.Address, payer common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.V2SwapExactInput(&_SwapRouter.TransactOpts, recipient, amountIn, amountOutMinimum, path, payer)
}

// V2SwapExactInput is a paid mutator transaction binding the contract method 0xb1542aef.
//
// Solidity: function v2SwapExactInput(address recipient, uint256 amountIn, uint256 amountOutMinimum, address[] path, address payer) returns()
func (_SwapRouter *SwapRouterTransactorSession) V2SwapExactInput(recipient common.Address, amountIn *big.Int, amountOutMinimum *big.Int, path []common.Address, payer common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.V2SwapExactInput(&_SwapRouter.TransactOpts, recipient, amountIn, amountOutMinimum, path, payer)
}

// V2SwapExactOutput is a paid mutator transaction binding the contract method 0x343e4d26.
//
// Solidity: function v2SwapExactOutput(address recipient, uint256 amountOut, uint256 amountInMaximum, address[] path, address payer) returns()
func (_SwapRouter *SwapRouterTransactor) V2SwapExactOutput(opts *bind.TransactOpts, recipient common.Address, amountOut *big.Int, amountInMaximum *big.Int, path []common.Address, payer common.Address) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "v2SwapExactOutput", recipient, amountOut, amountInMaximum, path, payer)
}

// V2SwapExactOutput is a paid mutator transaction binding the contract method 0x343e4d26.
//
// Solidity: function v2SwapExactOutput(address recipient, uint256 amountOut, uint256 amountInMaximum, address[] path, address payer) returns()
func (_SwapRouter *SwapRouterSession) V2SwapExactOutput(recipient common.Address, amountOut *big.Int, amountInMaximum *big.Int, path []common.Address, payer common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.V2SwapExactOutput(&_SwapRouter.TransactOpts, recipient, amountOut, amountInMaximum, path, payer)
}

// V2SwapExactOutput is a paid mutator transaction binding the contract method 0x343e4d26.
//
// Solidity: function v2SwapExactOutput(address recipient, uint256 amountOut, uint256 amountInMaximum, address[] path, address payer) returns()
func (_SwapRouter *SwapRouterTransactorSession) V2SwapExactOutput(recipient common.Address, amountOut *big.Int, amountInMaximum *big.Int, path []common.Address, payer common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.V2SwapExactOutput(&_SwapRouter.TransactOpts, recipient, amountOut, amountInMaximum, path, payer)
}

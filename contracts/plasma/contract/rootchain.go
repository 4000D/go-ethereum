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

// ByteUtilsABI is the input ABI used to generate the binding from.
const ByteUtilsABI = "[]"

// ByteUtilsBin is the compiled bytecode used for deploying new contracts.
const ByteUtilsBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820404e9bea48bd5ed6854c499932789d08a7721b793a5f3e96eb0feb351fc289de0029`

// DeployByteUtils deploys a new Ethereum contract, binding an instance of ByteUtils to it.
func DeployByteUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ByteUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(ByteUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ByteUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ByteUtils{ByteUtilsCaller: ByteUtilsCaller{contract: contract}, ByteUtilsTransactor: ByteUtilsTransactor{contract: contract}, ByteUtilsFilterer: ByteUtilsFilterer{contract: contract}}, nil
}

// ByteUtils is an auto generated Go binding around an Ethereum contract.
type ByteUtils struct {
	ByteUtilsCaller     // Read-only binding to the contract
	ByteUtilsTransactor // Write-only binding to the contract
	ByteUtilsFilterer   // Log filterer for contract events
}

// ByteUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ByteUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ByteUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ByteUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ByteUtilsSession struct {
	Contract     *ByteUtils        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ByteUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ByteUtilsCallerSession struct {
	Contract *ByteUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ByteUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ByteUtilsTransactorSession struct {
	Contract     *ByteUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ByteUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ByteUtilsRaw struct {
	Contract *ByteUtils // Generic contract binding to access the raw methods on
}

// ByteUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ByteUtilsCallerRaw struct {
	Contract *ByteUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// ByteUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ByteUtilsTransactorRaw struct {
	Contract *ByteUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewByteUtils creates a new instance of ByteUtils, bound to a specific deployed contract.
func NewByteUtils(address common.Address, backend bind.ContractBackend) (*ByteUtils, error) {
	contract, err := bindByteUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ByteUtils{ByteUtilsCaller: ByteUtilsCaller{contract: contract}, ByteUtilsTransactor: ByteUtilsTransactor{contract: contract}, ByteUtilsFilterer: ByteUtilsFilterer{contract: contract}}, nil
}

// NewByteUtilsCaller creates a new read-only instance of ByteUtils, bound to a specific deployed contract.
func NewByteUtilsCaller(address common.Address, caller bind.ContractCaller) (*ByteUtilsCaller, error) {
	contract, err := bindByteUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ByteUtilsCaller{contract: contract}, nil
}

// NewByteUtilsTransactor creates a new write-only instance of ByteUtils, bound to a specific deployed contract.
func NewByteUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*ByteUtilsTransactor, error) {
	contract, err := bindByteUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ByteUtilsTransactor{contract: contract}, nil
}

// NewByteUtilsFilterer creates a new log filterer instance of ByteUtils, bound to a specific deployed contract.
func NewByteUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*ByteUtilsFilterer, error) {
	contract, err := bindByteUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ByteUtilsFilterer{contract: contract}, nil
}

// bindByteUtils binds a generic wrapper to an already deployed contract.
func bindByteUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ByteUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ByteUtils *ByteUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ByteUtils.Contract.ByteUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ByteUtils *ByteUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ByteUtils.Contract.ByteUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ByteUtils *ByteUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ByteUtils.Contract.ByteUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ByteUtils *ByteUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ByteUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ByteUtils *ByteUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ByteUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ByteUtils *ByteUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ByteUtils.Contract.contract.Transact(opts, method, params...)
}

// ECRecoveryABI is the input ABI used to generate the binding from.
const ECRecoveryABI = "[]"

// ECRecoveryBin is the compiled bytecode used for deploying new contracts.
const ECRecoveryBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582091df73330e81da77db0532d84525b55327ef06f4e97bdaf2834660d8a5631be90029`

// DeployECRecovery deploys a new Ethereum contract, binding an instance of ECRecovery to it.
func DeployECRecovery(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECRecovery, error) {
	parsed, err := abi.JSON(strings.NewReader(ECRecoveryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECRecoveryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECRecovery{ECRecoveryCaller: ECRecoveryCaller{contract: contract}, ECRecoveryTransactor: ECRecoveryTransactor{contract: contract}, ECRecoveryFilterer: ECRecoveryFilterer{contract: contract}}, nil
}

// ECRecovery is an auto generated Go binding around an Ethereum contract.
type ECRecovery struct {
	ECRecoveryCaller     // Read-only binding to the contract
	ECRecoveryTransactor // Write-only binding to the contract
	ECRecoveryFilterer   // Log filterer for contract events
}

// ECRecoveryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECRecoveryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoveryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECRecoveryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoveryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECRecoveryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoverySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECRecoverySession struct {
	Contract     *ECRecovery       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECRecoveryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECRecoveryCallerSession struct {
	Contract *ECRecoveryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ECRecoveryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECRecoveryTransactorSession struct {
	Contract     *ECRecoveryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ECRecoveryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECRecoveryRaw struct {
	Contract *ECRecovery // Generic contract binding to access the raw methods on
}

// ECRecoveryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECRecoveryCallerRaw struct {
	Contract *ECRecoveryCaller // Generic read-only contract binding to access the raw methods on
}

// ECRecoveryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECRecoveryTransactorRaw struct {
	Contract *ECRecoveryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECRecovery creates a new instance of ECRecovery, bound to a specific deployed contract.
func NewECRecovery(address common.Address, backend bind.ContractBackend) (*ECRecovery, error) {
	contract, err := bindECRecovery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECRecovery{ECRecoveryCaller: ECRecoveryCaller{contract: contract}, ECRecoveryTransactor: ECRecoveryTransactor{contract: contract}, ECRecoveryFilterer: ECRecoveryFilterer{contract: contract}}, nil
}

// NewECRecoveryCaller creates a new read-only instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryCaller(address common.Address, caller bind.ContractCaller) (*ECRecoveryCaller, error) {
	contract, err := bindECRecovery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryCaller{contract: contract}, nil
}

// NewECRecoveryTransactor creates a new write-only instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryTransactor(address common.Address, transactor bind.ContractTransactor) (*ECRecoveryTransactor, error) {
	contract, err := bindECRecovery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryTransactor{contract: contract}, nil
}

// NewECRecoveryFilterer creates a new log filterer instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryFilterer(address common.Address, filterer bind.ContractFilterer) (*ECRecoveryFilterer, error) {
	contract, err := bindECRecovery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryFilterer{contract: contract}, nil
}

// bindECRecovery binds a generic wrapper to an already deployed contract.
func bindECRecovery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECRecoveryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECRecovery *ECRecoveryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECRecovery.Contract.ECRecoveryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECRecovery *ECRecoveryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECRecovery.Contract.ECRecoveryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECRecovery *ECRecoveryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECRecovery.Contract.ECRecoveryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECRecovery *ECRecoveryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECRecovery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECRecovery *ECRecoveryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECRecovery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECRecovery *ECRecoveryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECRecovery.Contract.contract.Transact(opts, method, params...)
}

// MathABI is the input ABI used to generate the binding from.
const MathABI = "[]"

// MathBin is the compiled bytecode used for deploying new contracts.
const MathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582093c6f0865f73208d3868408da47a0b254e78fda541b25402612a9fe3198b9e4a0029`

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// MerkleABI is the input ABI used to generate the binding from.
const MerkleABI = "[]"

// MerkleBin is the compiled bytecode used for deploying new contracts.
const MerkleBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820b0fa05b6092eeb294bf1b3f8764c189f8508b9c4a8a6b2773e14bb892c388a510029`

// DeployMerkle deploys a new Ethereum contract, binding an instance of Merkle to it.
func DeployMerkle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Merkle, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MerkleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Merkle{MerkleCaller: MerkleCaller{contract: contract}, MerkleTransactor: MerkleTransactor{contract: contract}, MerkleFilterer: MerkleFilterer{contract: contract}}, nil
}

// Merkle is an auto generated Go binding around an Ethereum contract.
type Merkle struct {
	MerkleCaller     // Read-only binding to the contract
	MerkleTransactor // Write-only binding to the contract
	MerkleFilterer   // Log filterer for contract events
}

// MerkleCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleSession struct {
	Contract     *Merkle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleCallerSession struct {
	Contract *MerkleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MerkleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleTransactorSession struct {
	Contract     *MerkleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleRaw struct {
	Contract *Merkle // Generic contract binding to access the raw methods on
}

// MerkleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleCallerRaw struct {
	Contract *MerkleCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleTransactorRaw struct {
	Contract *MerkleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkle creates a new instance of Merkle, bound to a specific deployed contract.
func NewMerkle(address common.Address, backend bind.ContractBackend) (*Merkle, error) {
	contract, err := bindMerkle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Merkle{MerkleCaller: MerkleCaller{contract: contract}, MerkleTransactor: MerkleTransactor{contract: contract}, MerkleFilterer: MerkleFilterer{contract: contract}}, nil
}

// NewMerkleCaller creates a new read-only instance of Merkle, bound to a specific deployed contract.
func NewMerkleCaller(address common.Address, caller bind.ContractCaller) (*MerkleCaller, error) {
	contract, err := bindMerkle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleCaller{contract: contract}, nil
}

// NewMerkleTransactor creates a new write-only instance of Merkle, bound to a specific deployed contract.
func NewMerkleTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleTransactor, error) {
	contract, err := bindMerkle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTransactor{contract: contract}, nil
}

// NewMerkleFilterer creates a new log filterer instance of Merkle, bound to a specific deployed contract.
func NewMerkleFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleFilterer, error) {
	contract, err := bindMerkle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleFilterer{contract: contract}, nil
}

// bindMerkle binds a generic wrapper to an already deployed contract.
func bindMerkle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Merkle *MerkleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Merkle.Contract.MerkleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Merkle *MerkleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merkle.Contract.MerkleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Merkle *MerkleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Merkle.Contract.MerkleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Merkle *MerkleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Merkle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Merkle *MerkleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merkle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Merkle *MerkleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Merkle.Contract.contract.Transact(opts, method, params...)
}

// PlasmaRLPABI is the input ABI used to generate the binding from.
const PlasmaRLPABI = "[]"

// PlasmaRLPBin is the compiled bytecode used for deploying new contracts.
const PlasmaRLPBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820457b3669961f5ca16daf22cf1580338bc1ba994cc2f42e4f59c2e9ceccf3607f0029`

// DeployPlasmaRLP deploys a new Ethereum contract, binding an instance of PlasmaRLP to it.
func DeployPlasmaRLP(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PlasmaRLP, error) {
	parsed, err := abi.JSON(strings.NewReader(PlasmaRLPABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PlasmaRLPBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PlasmaRLP{PlasmaRLPCaller: PlasmaRLPCaller{contract: contract}, PlasmaRLPTransactor: PlasmaRLPTransactor{contract: contract}, PlasmaRLPFilterer: PlasmaRLPFilterer{contract: contract}}, nil
}

// PlasmaRLP is an auto generated Go binding around an Ethereum contract.
type PlasmaRLP struct {
	PlasmaRLPCaller     // Read-only binding to the contract
	PlasmaRLPTransactor // Write-only binding to the contract
	PlasmaRLPFilterer   // Log filterer for contract events
}

// PlasmaRLPCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlasmaRLPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaRLPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlasmaRLPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaRLPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlasmaRLPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaRLPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlasmaRLPSession struct {
	Contract     *PlasmaRLP        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlasmaRLPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlasmaRLPCallerSession struct {
	Contract *PlasmaRLPCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PlasmaRLPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlasmaRLPTransactorSession struct {
	Contract     *PlasmaRLPTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PlasmaRLPRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlasmaRLPRaw struct {
	Contract *PlasmaRLP // Generic contract binding to access the raw methods on
}

// PlasmaRLPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlasmaRLPCallerRaw struct {
	Contract *PlasmaRLPCaller // Generic read-only contract binding to access the raw methods on
}

// PlasmaRLPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlasmaRLPTransactorRaw struct {
	Contract *PlasmaRLPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlasmaRLP creates a new instance of PlasmaRLP, bound to a specific deployed contract.
func NewPlasmaRLP(address common.Address, backend bind.ContractBackend) (*PlasmaRLP, error) {
	contract, err := bindPlasmaRLP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlasmaRLP{PlasmaRLPCaller: PlasmaRLPCaller{contract: contract}, PlasmaRLPTransactor: PlasmaRLPTransactor{contract: contract}, PlasmaRLPFilterer: PlasmaRLPFilterer{contract: contract}}, nil
}

// NewPlasmaRLPCaller creates a new read-only instance of PlasmaRLP, bound to a specific deployed contract.
func NewPlasmaRLPCaller(address common.Address, caller bind.ContractCaller) (*PlasmaRLPCaller, error) {
	contract, err := bindPlasmaRLP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlasmaRLPCaller{contract: contract}, nil
}

// NewPlasmaRLPTransactor creates a new write-only instance of PlasmaRLP, bound to a specific deployed contract.
func NewPlasmaRLPTransactor(address common.Address, transactor bind.ContractTransactor) (*PlasmaRLPTransactor, error) {
	contract, err := bindPlasmaRLP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlasmaRLPTransactor{contract: contract}, nil
}

// NewPlasmaRLPFilterer creates a new log filterer instance of PlasmaRLP, bound to a specific deployed contract.
func NewPlasmaRLPFilterer(address common.Address, filterer bind.ContractFilterer) (*PlasmaRLPFilterer, error) {
	contract, err := bindPlasmaRLP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlasmaRLPFilterer{contract: contract}, nil
}

// bindPlasmaRLP binds a generic wrapper to an already deployed contract.
func bindPlasmaRLP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlasmaRLPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlasmaRLP *PlasmaRLPRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PlasmaRLP.Contract.PlasmaRLPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlasmaRLP *PlasmaRLPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlasmaRLP.Contract.PlasmaRLPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlasmaRLP *PlasmaRLPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlasmaRLP.Contract.PlasmaRLPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlasmaRLP *PlasmaRLPCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PlasmaRLP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlasmaRLP *PlasmaRLPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlasmaRLP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlasmaRLP *PlasmaRLPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlasmaRLP.Contract.contract.Transact(opts, method, params...)
}

// PriorityQueueABI is the input ABI used to generate the binding from.
const PriorityQueueABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"minChild\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"uint256\"}],\"name\":\"insert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"delMin\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMin\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PriorityQueueBin is the compiled bytecode used for deploying new contracts.
const PriorityQueueBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a03191633178155604080516020810190915290815261003e9060019081610049565b5060006002556100b6565b828054828255906000526020600020908101928215610089579160200282015b82811115610089578251829060ff16905591602001919060010190610069565b50610095929150610099565b5090565b6100b391905b80821115610095576000815560010161009f565b90565b6105ba806100c56000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632dcdcd0c811461007157806390b5561d1461009b578063b07576ac146100b5578063bda1504b146100ca578063d6362e97146100df575b600080fd5b34801561007d57600080fd5b506100896004356100f4565b60408051918252519081900360200190f35b3480156100a757600080fd5b506100b36004356101c3565b005b3480156100c157600080fd5b5061008961023d565b3480156100d657600080fd5b5061008961031c565b3480156100eb57600080fd5b50610089610322565b600060025461011e600161011260028661034490919063ffffffff16565b9063ffffffff61037a16565b111561013c5761013582600263ffffffff61034416565b90506101be565b60016101538161011285600263ffffffff61034416565b8154811061015d57fe5b600091825260209091200154600161017c84600263ffffffff61034416565b8154811061018657fe5b906000526020600020015410156101a85761013582600263ffffffff61034416565b610135600161011284600263ffffffff61034416565b919050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146101e757600080fd5b60018054808201825560008290527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60182905560025461022c9163ffffffff61037a16565b600281905561023a90610389565b50565b60008054819073ffffffffffffffffffffffffffffffffffffffff16331461026457600080fd5b600180548190811061027257fe5b90600052602060002001549050600160025481548110151561029057fe5b90600052602060002001546001808154811015156102aa57fe5b906000526020600020018190555060016002548154811015156102c957fe5b60009182526020822001556002546102e890600163ffffffff61045d16565b6002556102f5600161046f565b600180546103089163ffffffff61045d16565b610313600182610555565b508091505b5090565b60025481565b600060018081548110151561033357fe5b906000526020600020015490505b90565b6000808315156103575760009150610373565b5082820282848281151561036757fe5b041461036f57fe5b8091505b5092915050565b60008282018381101561036f57fe5b6001805482916000918390811061039c57fe5b906000526020600020015490505b60016103bd84600263ffffffff61053e16565b815481106103c757fe5b90600052602060002001548110156104345760016103ec84600263ffffffff61053e16565b815481106103f657fe5b906000526020600020015460018481548110151561041057fe5b60009182526020909120015561042d83600263ffffffff61053e16565b92506103aa565b828214610458578060018481548110151561044b57fe5b6000918252602090912001555b505050565b60008282111561046957fe5b50900390565b600080600083925060018481548110151561048657fe5b9060005260206000200154915061049c846100f4565b90505b60025481111580156104c8575060018054829081106104ba57fe5b906000526020600020015482115b156105145760018054829081106104db57fe5b90600052602060002001546001858154811015156104f557fe5b60009182526020909120015592508261050d816100f4565b905061049f565b838314610538578160018581548110151561052b57fe5b6000918252602090912001555b50505050565b600080828481151561054c57fe5b04949350505050565b8154818355818111156104585760008381526020902061045891810190830161034191905b80821115610318576000815560010161057a5600a165627a7a72305820f8c97a2e1dbea811473e5db707c2b25ee279c2092443d7f8cc3673ce063e62060029`

// DeployPriorityQueue deploys a new Ethereum contract, binding an instance of PriorityQueue to it.
func DeployPriorityQueue(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PriorityQueue, error) {
	parsed, err := abi.JSON(strings.NewReader(PriorityQueueABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PriorityQueueBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PriorityQueue{PriorityQueueCaller: PriorityQueueCaller{contract: contract}, PriorityQueueTransactor: PriorityQueueTransactor{contract: contract}, PriorityQueueFilterer: PriorityQueueFilterer{contract: contract}}, nil
}

// PriorityQueue is an auto generated Go binding around an Ethereum contract.
type PriorityQueue struct {
	PriorityQueueCaller     // Read-only binding to the contract
	PriorityQueueTransactor // Write-only binding to the contract
	PriorityQueueFilterer   // Log filterer for contract events
}

// PriorityQueueCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriorityQueueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriorityQueueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriorityQueueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriorityQueueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriorityQueueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriorityQueueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriorityQueueSession struct {
	Contract     *PriorityQueue    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriorityQueueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriorityQueueCallerSession struct {
	Contract *PriorityQueueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PriorityQueueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriorityQueueTransactorSession struct {
	Contract     *PriorityQueueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PriorityQueueRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriorityQueueRaw struct {
	Contract *PriorityQueue // Generic contract binding to access the raw methods on
}

// PriorityQueueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriorityQueueCallerRaw struct {
	Contract *PriorityQueueCaller // Generic read-only contract binding to access the raw methods on
}

// PriorityQueueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriorityQueueTransactorRaw struct {
	Contract *PriorityQueueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriorityQueue creates a new instance of PriorityQueue, bound to a specific deployed contract.
func NewPriorityQueue(address common.Address, backend bind.ContractBackend) (*PriorityQueue, error) {
	contract, err := bindPriorityQueue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriorityQueue{PriorityQueueCaller: PriorityQueueCaller{contract: contract}, PriorityQueueTransactor: PriorityQueueTransactor{contract: contract}, PriorityQueueFilterer: PriorityQueueFilterer{contract: contract}}, nil
}

// NewPriorityQueueCaller creates a new read-only instance of PriorityQueue, bound to a specific deployed contract.
func NewPriorityQueueCaller(address common.Address, caller bind.ContractCaller) (*PriorityQueueCaller, error) {
	contract, err := bindPriorityQueue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriorityQueueCaller{contract: contract}, nil
}

// NewPriorityQueueTransactor creates a new write-only instance of PriorityQueue, bound to a specific deployed contract.
func NewPriorityQueueTransactor(address common.Address, transactor bind.ContractTransactor) (*PriorityQueueTransactor, error) {
	contract, err := bindPriorityQueue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriorityQueueTransactor{contract: contract}, nil
}

// NewPriorityQueueFilterer creates a new log filterer instance of PriorityQueue, bound to a specific deployed contract.
func NewPriorityQueueFilterer(address common.Address, filterer bind.ContractFilterer) (*PriorityQueueFilterer, error) {
	contract, err := bindPriorityQueue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriorityQueueFilterer{contract: contract}, nil
}

// bindPriorityQueue binds a generic wrapper to an already deployed contract.
func bindPriorityQueue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriorityQueueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriorityQueue *PriorityQueueRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriorityQueue.Contract.PriorityQueueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriorityQueue *PriorityQueueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriorityQueue.Contract.PriorityQueueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriorityQueue *PriorityQueueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriorityQueue.Contract.PriorityQueueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriorityQueue *PriorityQueueCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriorityQueue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriorityQueue *PriorityQueueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriorityQueue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriorityQueue *PriorityQueueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriorityQueue.Contract.contract.Transact(opts, method, params...)
}

// CurrentSize is a free data retrieval call binding the contract method 0xbda1504b.
//
// Solidity: function currentSize() constant returns(uint256)
func (_PriorityQueue *PriorityQueueCaller) CurrentSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriorityQueue.contract.Call(opts, out, "currentSize")
	return *ret0, err
}

// CurrentSize is a free data retrieval call binding the contract method 0xbda1504b.
//
// Solidity: function currentSize() constant returns(uint256)
func (_PriorityQueue *PriorityQueueSession) CurrentSize() (*big.Int, error) {
	return _PriorityQueue.Contract.CurrentSize(&_PriorityQueue.CallOpts)
}

// CurrentSize is a free data retrieval call binding the contract method 0xbda1504b.
//
// Solidity: function currentSize() constant returns(uint256)
func (_PriorityQueue *PriorityQueueCallerSession) CurrentSize() (*big.Int, error) {
	return _PriorityQueue.Contract.CurrentSize(&_PriorityQueue.CallOpts)
}

// GetMin is a free data retrieval call binding the contract method 0xd6362e97.
//
// Solidity: function getMin() constant returns(uint256)
func (_PriorityQueue *PriorityQueueCaller) GetMin(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriorityQueue.contract.Call(opts, out, "getMin")
	return *ret0, err
}

// GetMin is a free data retrieval call binding the contract method 0xd6362e97.
//
// Solidity: function getMin() constant returns(uint256)
func (_PriorityQueue *PriorityQueueSession) GetMin() (*big.Int, error) {
	return _PriorityQueue.Contract.GetMin(&_PriorityQueue.CallOpts)
}

// GetMin is a free data retrieval call binding the contract method 0xd6362e97.
//
// Solidity: function getMin() constant returns(uint256)
func (_PriorityQueue *PriorityQueueCallerSession) GetMin() (*big.Int, error) {
	return _PriorityQueue.Contract.GetMin(&_PriorityQueue.CallOpts)
}

// MinChild is a free data retrieval call binding the contract method 0x2dcdcd0c.
//
// Solidity: function minChild(i uint256) constant returns(uint256)
func (_PriorityQueue *PriorityQueueCaller) MinChild(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriorityQueue.contract.Call(opts, out, "minChild", i)
	return *ret0, err
}

// MinChild is a free data retrieval call binding the contract method 0x2dcdcd0c.
//
// Solidity: function minChild(i uint256) constant returns(uint256)
func (_PriorityQueue *PriorityQueueSession) MinChild(i *big.Int) (*big.Int, error) {
	return _PriorityQueue.Contract.MinChild(&_PriorityQueue.CallOpts, i)
}

// MinChild is a free data retrieval call binding the contract method 0x2dcdcd0c.
//
// Solidity: function minChild(i uint256) constant returns(uint256)
func (_PriorityQueue *PriorityQueueCallerSession) MinChild(i *big.Int) (*big.Int, error) {
	return _PriorityQueue.Contract.MinChild(&_PriorityQueue.CallOpts, i)
}

// DelMin is a paid mutator transaction binding the contract method 0xb07576ac.
//
// Solidity: function delMin() returns(uint256)
func (_PriorityQueue *PriorityQueueTransactor) DelMin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriorityQueue.contract.Transact(opts, "delMin")
}

// DelMin is a paid mutator transaction binding the contract method 0xb07576ac.
//
// Solidity: function delMin() returns(uint256)
func (_PriorityQueue *PriorityQueueSession) DelMin() (*types.Transaction, error) {
	return _PriorityQueue.Contract.DelMin(&_PriorityQueue.TransactOpts)
}

// DelMin is a paid mutator transaction binding the contract method 0xb07576ac.
//
// Solidity: function delMin() returns(uint256)
func (_PriorityQueue *PriorityQueueTransactorSession) DelMin() (*types.Transaction, error) {
	return _PriorityQueue.Contract.DelMin(&_PriorityQueue.TransactOpts)
}

// Insert is a paid mutator transaction binding the contract method 0x90b5561d.
//
// Solidity: function insert(k uint256) returns()
func (_PriorityQueue *PriorityQueueTransactor) Insert(opts *bind.TransactOpts, k *big.Int) (*types.Transaction, error) {
	return _PriorityQueue.contract.Transact(opts, "insert", k)
}

// Insert is a paid mutator transaction binding the contract method 0x90b5561d.
//
// Solidity: function insert(k uint256) returns()
func (_PriorityQueue *PriorityQueueSession) Insert(k *big.Int) (*types.Transaction, error) {
	return _PriorityQueue.Contract.Insert(&_PriorityQueue.TransactOpts, k)
}

// Insert is a paid mutator transaction binding the contract method 0x90b5561d.
//
// Solidity: function insert(k uint256) returns()
func (_PriorityQueue *PriorityQueueTransactorSession) Insert(k *big.Int) (*types.Transaction, error) {
	return _PriorityQueue.Contract.Insert(&_PriorityQueue.TransactOpts, k)
}

// RootChainABI is the input ABI used to generate the binding from.
const RootChainABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"utxoPos\",\"type\":\"uint256\"},{\"name\":\"txBytes\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"sigs\",\"type\":\"bytes\"}],\"name\":\"startExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNextExit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cUtxoPos\",\"type\":\"uint256\"},{\"name\":\"eUtxoIndex\",\"type\":\"uint256\"},{\"name\":\"txBytes\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"sigs\",\"type\":\"bytes\"},{\"name\":\"confirmationSig\",\"type\":\"bytes\"}],\"name\":\"challengeExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exits\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"childBlockInterval\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"depositPos\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"startDepositExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentChildBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getChildChain\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentDepositBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"startFeeExit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"submitBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDepositBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalizeExits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentFeeExit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"utxoPos\",\"type\":\"uint256\"}],\"name\":\"getExit\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"childChain\",\"outputs\":[{\"name\":\"root\",\"type\":\"bytes32\"},{\"name\":\"created_at\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"depositBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"submitBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"Submit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"exitor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"utxoPos\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ExitStarted\",\"type\":\"event\"}]"

// RootChainBin is the compiled bytecode used for deploying new contracts.
const RootChainBin = `0x608060405234801561001057600080fd5b5060038054600160a060020a031916331790556103e8600681905560045560016005819055600755610040610082565b604051809103906000f08015801561005c573d6000803e3d6000fd5b5060028054600160a060020a031916600160a060020a0392909216919091179055610092565b60405161067f8061178d83390190565b6116ec806100a16000396000f3006080604052600436106100e25763ffffffff60e060020a6000350416631c91a6b981146100e75780632fd56a25146101c357806332773ba3146101f1578063342de1791461030b57806338a9e0bc1461034657806370e4abf61461036d5780637a95f1e81461038857806385444de31461039d578063a98c7f2c146103b5578063ae80d8c8146103ca578063baa47694146103e2578063bcd59261146103fa578063bf7e214f1461040f578063c6ab44cd14610440578063cfe7d85514610455578063d0e30db01461046a578063e60f1ff114610472578063f95643b11461048a575b600080fd5b3480156100f357600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101c195833595369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506104a29650505050505050565b005b3480156101cf57600080fd5b506101d86106c0565b6040805192835260208301919091528051918290030190f35b3480156101fd57600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526101c194823594602480359536959460649492019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506107799650505050505050565b34801561031757600080fd5b50610323600435610928565b60408051600160a060020a03909316835260208301919091528051918290030190f35b34801561035257600080fd5b5061035b61094d565b60408051918252519081900360200190f35b34801561037957600080fd5b506101c1600435602435610953565b34801561039457600080fd5b5061035b6109e7565b3480156103a957600080fd5b506101d86004356109ed565b3480156103c157600080fd5b5061035b610a07565b3480156103d657600080fd5b5061035b600435610a0d565b3480156103ee57600080fd5b506101c1600435610a54565b34801561040657600080fd5b5061035b610aba565b34801561041b57600080fd5b50610424610aea565b60408051600160a060020a039092168252519081900360200190f35b34801561044c57600080fd5b506101c1610af9565b34801561046157600080fd5b5061035b610d93565b6101c1610d99565b34801561047e57600080fd5b50610323600435610e61565b34801561049657600080fd5b506101d8600435610e85565b60008060006104af61165d565b600080633b9aca008a049550612710633b9aca008b068115156104ce57fe5b0494506127108502633b9aca0087028b030393506104f489600b8663ffffffff610e9e16565b6020810151909350600160a060020a0316331461051057600080fd5b600080878152602001908152602001600020600001549150886040518082805190602001908083835b602083106105585780518252601f199092019160209182019101610539565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206105928860006082610f57565b6040518281528151602080830191908401908083835b602083106105c75780518252601f1990920191602091820191016105a8565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405180910390209050610663896040518082805190602001908083835b6020831061062a5780518252601f19909201916020918201910161060b565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390208385604001518a610f9e565b151561066e57600080fd5b6106808186848b63ffffffff6110a916565b151561068b57600080fd5b6106b48a846020015185600001516000808b815260200190815260200160002060010154611135565b50505050505050505050565b6000806000806000600260009054906101000a9004600160a060020a0316600160a060020a031663d6362e976040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561071b57600080fd5b505af115801561072f573d6000803e3d6000fd5b505050506040513d602081101561074557600080fd5b50516fffffffffffffffffffffffffffffffff81169670010000000000000000000000000000000090910495509350505050565b60008080808080806107938b600b8e63ffffffff6112aa16565b9650612710633b9aca008e06049550600080633b9aca008f0481526020019081526020016000206000015494508a6040518082805190602001908083835b602083106107f05780518252601f1990920191602091820191016107d1565b51815160001960209485036101000a01908116901991909116179052604080519490920184900384208085528482018c905282519485900390920184208285528f51929a5098508995508e945083810192508401908083835b602083106108685780518252601f199092019160209182019101610849565b51815160209384036101000a6000190180199092169116179052604080519290940182900390912060008e815260019092529290205491965050600160a060020a031693506108bd92508591508a9050611310565b600160a060020a038281169116146108d457600080fd5b6108e68287878d63ffffffff6110a916565b15156108f157600080fd5b5050506000938452505060016020525060409020805473ffffffffffffffffffffffffffffffffffffffff19169055505050505050565b60016020819052600091825260409091208054910154600160a060020a039091169082565b60065481565b60008080633b9aca00850492506006548381151561096d57fe5b06151561097957600080fd5b5050600081815260208190526040908190205481516c0100000000000000000000000033028152601481018590529151918290036034019091208082146109bf57600080fd5b6109e085338660008088815260200190815260200160002060010154611135565b5050505050565b60045481565b600090815260208190526040902080546001909101549091565b60055481565b600354600090600160a060020a03163314610a2757600080fd5b610a38600754338442600101611135565b600754610a4c90600163ffffffff6113e516565b600755919050565b600354600160a060020a03163314610a6b57600080fd5b6040805180820182528281524260208083019182526004805460009081529182905293902091518255516001909101556006549054610aaf9163ffffffff6113e516565b600455506001600555565b6000610ae5600554610ad96006546004546113fb90919063ffffffff16565b9063ffffffff6113e516565b905090565b600354600160a060020a031681565b600080610b04611688565b610b0c6106c0565b60008281526001602081815260409283902083518085019094528054600160a060020a031684529091015490820152919450925090505b4282108015610bda5750600254604080517fbda1504b0000000000000000000000000000000000000000000000000000000081529051600092600160a060020a03169163bda1504b91600480830192602092919082900301818787803b158015610bac57600080fd5b505af1158015610bc0573d6000803e3d6000fd5b505050506040513d6020811015610bd657600080fd5b5051115b15610d8e57506000828152600160208181526040808420815180830183528154600160a060020a0316808252919094015492840183905290519293909282156108fc029291818181858888f19350505050158015610c3c573d6000803e3d6000fd5b50600260009054906101000a9004600160a060020a0316600160a060020a031663b07576ac6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610c9057600080fd5b505af1158015610ca4573d6000803e3d6000fd5b505050506040513d6020811015610cba57600080fd5b50506000838152600160209081526040808320805473ffffffffffffffffffffffffffffffffffffffff1916905560025481517fbda1504b0000000000000000000000000000000000000000000000000000000081529151600160a060020a039091169263bda1504b926004808201939182900301818787803b158015610d4057600080fd5b505af1158015610d54573d6000803e3d6000fd5b505050506040513d6020811015610d6a57600080fd5b50511115610d8457610d7a6106c0565b9093509150610d89565b610d8e565b610b43565b505050565b60075481565b600080600654600554101515610dae57600080fd5b604080516c010000000000000000000000003302815234601482015290519081900360340190209150610ddf610aba565b60408051808201825284815242602080830191825260008581529081905292909220905181559051600191820155600554919250610e23919063ffffffff6113e516565b600555604080513481529051829133917f90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a159181900360200190a35050565b60009081526001602081905260409091208054910154600160a060020a0390911691565b6000602081905290815260409020805460019091015482565b610ea661165d565b6060610eba610eb486611412565b85611440565b9050606060405190810160405280610eee8386600202600701815181101515610edf57fe5b906020019060200201516114ce565b8152602001610f198386600202600601815181101515610f0a57fe5b906020019060200201516114f2565b600160a060020a03168152602001610f39836003815181101515610edf57fe5b610f4b846000815181101515610edf57fe5b02905295945050505050565b604051606090601f831680820180850187830187015b81831015610f85578051835260209283019201610f6d565b5050848352601f01601f19166040525090509392505050565b600060608060606000806000606060418951811515610fb957fe5b06158015610fca5750610104895111155b1515610fd557600080fd5b610fe28960006041610f57565b9650610ff089604180610f57565b9550610fff8960826041610f57565b604080518e8152602081018e9052815190819003909101902090955093506001925082915061102e8486611310565b600160a060020a03166110418d89611310565b600160a060020a031614925060008a111561108f576110638960c36041610f57565b905061106f8482611310565b600160a060020a03166110828d88611310565b600160a060020a03161491505b8280156110995750815b9c9b505050505050505050505050565b60008060008084516102001415156110c057600080fd5b5086905060205b6102008111611127578481015192506002870615156110fe5760408051928352602083018490528051928390030190912090611119565b60408051848152602081019390935280519283900301909120905b6002870496506020016110c7565b509390931495945050505050565b633b9aca00840460008061115362127500850162093a80420161151a565b915050700100000000000000000000000000000000810286176000851161117957600080fd5b600087815260016020819052604090912001541561119657600080fd5b600254604080517f90b5561d000000000000000000000000000000000000000000000000000000008152600481018490529051600160a060020a03909216916390b5561d9160248082019260009290919082900301818387803b1580156111fc57600080fd5b505af1158015611210573d6000803e3d6000fd5b5050604080518082018252600160a060020a038a8116825260208083018b815260008e81526001808452908690209451855473ffffffffffffffffffffffffffffffffffffffff19169416939093178455519290910191909155815189815291518b94503393507fd62f48e64bfb0b95c15737946c829e7fe6378cd679a268c1ca5993dc235d900b9281900390910190a350505050505050565b6000606060006112c26112bc87611412565b86611440565b91508360030290506112de8282600201815181101515610edf57fe5b6112f28383600101815181101515610edf57fe5b835161130490859085908110610edf57fe5b01019695505050505050565b6000806000808451604114151561132a57600093506113dc565b50505060208201516040830151606084015160001a601b60ff8216101561134f57601b015b8060ff16601b1415801561136757508060ff16601c14155b1561137557600093506113dc565b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af11580156113cf573d6000803e3d6000fd5b5050506020604051035193505b50505092915050565b6000828201838110156113f457fe5b9392505050565b60008282111561140757fe5b508082035b92915050565b61141a611688565b5080516040805180820190915260208084018083529082018390529091905b5050919050565b606061144a61169f565b60008360405190808252806020026020018201604052801561148657816020015b611473611688565b81526020019060019003908161146b5790505b50925061149285611531565b91505b838110156114c6576114a682611556565b83828151811015156114b457fe5b60209081029091010152600101611495565b505092915050565b60008060006114dc84611588565b90516020919091036101000a9004949350505050565b600080600061150084611588565b50516c010000000000000000000000009004949350505050565b60008183111561152b57508161140c565b50919050565b61153961169f565b6000611544836115eb565b83519383529092016020820152919050565b61155e611688565b6020820151600061156e82611631565b828452602080850182905292019390910192909252919050565b805180516000918291821a908260808310156115aa57819450600193506115e3565b60b88310156115c857600186602001510393508160010194506115e3565b60b78303905080600187602001510303935080820160010194505b505050915091565b8051805160009190821a9060808210156116085760009250611439565b60b8821080611623575060c08210158015611623575060f882105b156114395760019250611439565b8051600090811a608081101561164a576001915061152b565b60b881101561152b57607e190192915050565b606060405190810160405280600081526020016000600160a060020a03168152602001600081525090565b604080518082019091526000808252602082015290565b6060604051908101604052806116b3611688565b81526020016000815250905600a165627a7a7230582034a4957d9f07742a7e3dfb706d4fb6eed6871151232ddc3fa703999a9129314d0029608060405234801561001057600080fd5b5060008054600160a060020a03191633178155604080516020810190915290815261003e9060019081610049565b5060006002556100b6565b828054828255906000526020600020908101928215610089579160200282015b82811115610089578251829060ff16905591602001919060010190610069565b50610095929150610099565b5090565b6100b391905b80821115610095576000815560010161009f565b90565b6105ba806100c56000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632dcdcd0c811461007157806390b5561d1461009b578063b07576ac146100b5578063bda1504b146100ca578063d6362e97146100df575b600080fd5b34801561007d57600080fd5b506100896004356100f4565b60408051918252519081900360200190f35b3480156100a757600080fd5b506100b36004356101c3565b005b3480156100c157600080fd5b5061008961023d565b3480156100d657600080fd5b5061008961031c565b3480156100eb57600080fd5b50610089610322565b600060025461011e600161011260028661034490919063ffffffff16565b9063ffffffff61037a16565b111561013c5761013582600263ffffffff61034416565b90506101be565b60016101538161011285600263ffffffff61034416565b8154811061015d57fe5b600091825260209091200154600161017c84600263ffffffff61034416565b8154811061018657fe5b906000526020600020015410156101a85761013582600263ffffffff61034416565b610135600161011284600263ffffffff61034416565b919050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146101e757600080fd5b60018054808201825560008290527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60182905560025461022c9163ffffffff61037a16565b600281905561023a90610389565b50565b60008054819073ffffffffffffffffffffffffffffffffffffffff16331461026457600080fd5b600180548190811061027257fe5b90600052602060002001549050600160025481548110151561029057fe5b90600052602060002001546001808154811015156102aa57fe5b906000526020600020018190555060016002548154811015156102c957fe5b60009182526020822001556002546102e890600163ffffffff61045d16565b6002556102f5600161046f565b600180546103089163ffffffff61045d16565b610313600182610555565b508091505b5090565b60025481565b600060018081548110151561033357fe5b906000526020600020015490505b90565b6000808315156103575760009150610373565b5082820282848281151561036757fe5b041461036f57fe5b8091505b5092915050565b60008282018381101561036f57fe5b6001805482916000918390811061039c57fe5b906000526020600020015490505b60016103bd84600263ffffffff61053e16565b815481106103c757fe5b90600052602060002001548110156104345760016103ec84600263ffffffff61053e16565b815481106103f657fe5b906000526020600020015460018481548110151561041057fe5b60009182526020909120015561042d83600263ffffffff61053e16565b92506103aa565b828214610458578060018481548110151561044b57fe5b6000918252602090912001555b505050565b60008282111561046957fe5b50900390565b600080600083925060018481548110151561048657fe5b9060005260206000200154915061049c846100f4565b90505b60025481111580156104c8575060018054829081106104ba57fe5b906000526020600020015482115b156105145760018054829081106104db57fe5b90600052602060002001546001858154811015156104f557fe5b60009182526020909120015592508261050d816100f4565b905061049f565b838314610538578160018581548110151561052b57fe5b6000918252602090912001555b50505050565b600080828481151561054c57fe5b04949350505050565b8154818355818111156104585760008381526020902061045891810190830161034191905b80821115610318576000815560010161057a5600a165627a7a72305820f8c97a2e1dbea811473e5db707c2b25ee279c2092443d7f8cc3673ce063e62060029`

// DeployRootChain deploys a new Ethereum contract, binding an instance of RootChain to it.
func DeployRootChain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RootChain, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RootChainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RootChain{RootChainCaller: RootChainCaller{contract: contract}, RootChainTransactor: RootChainTransactor{contract: contract}, RootChainFilterer: RootChainFilterer{contract: contract}}, nil
}

// RootChain is an auto generated Go binding around an Ethereum contract.
type RootChain struct {
	RootChainCaller     // Read-only binding to the contract
	RootChainTransactor // Write-only binding to the contract
	RootChainFilterer   // Log filterer for contract events
}

// RootChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootChainSession struct {
	Contract     *RootChain        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootChainCallerSession struct {
	Contract *RootChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RootChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootChainTransactorSession struct {
	Contract     *RootChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RootChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootChainRaw struct {
	Contract *RootChain // Generic contract binding to access the raw methods on
}

// RootChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootChainCallerRaw struct {
	Contract *RootChainCaller // Generic read-only contract binding to access the raw methods on
}

// RootChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootChainTransactorRaw struct {
	Contract *RootChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootChain creates a new instance of RootChain, bound to a specific deployed contract.
func NewRootChain(address common.Address, backend bind.ContractBackend) (*RootChain, error) {
	contract, err := bindRootChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootChain{RootChainCaller: RootChainCaller{contract: contract}, RootChainTransactor: RootChainTransactor{contract: contract}, RootChainFilterer: RootChainFilterer{contract: contract}}, nil
}

// NewRootChainCaller creates a new read-only instance of RootChain, bound to a specific deployed contract.
func NewRootChainCaller(address common.Address, caller bind.ContractCaller) (*RootChainCaller, error) {
	contract, err := bindRootChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainCaller{contract: contract}, nil
}

// NewRootChainTransactor creates a new write-only instance of RootChain, bound to a specific deployed contract.
func NewRootChainTransactor(address common.Address, transactor bind.ContractTransactor) (*RootChainTransactor, error) {
	contract, err := bindRootChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainTransactor{contract: contract}, nil
}

// NewRootChainFilterer creates a new log filterer instance of RootChain, bound to a specific deployed contract.
func NewRootChainFilterer(address common.Address, filterer bind.ContractFilterer) (*RootChainFilterer, error) {
	contract, err := bindRootChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootChainFilterer{contract: contract}, nil
}

// bindRootChain binds a generic wrapper to an already deployed contract.
func bindRootChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChain *RootChainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChain.Contract.RootChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChain *RootChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.Contract.RootChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChain *RootChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChain.Contract.RootChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChain *RootChainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChain *RootChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChain *RootChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChain.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_RootChain *RootChainCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_RootChain *RootChainSession) Authority() (common.Address, error) {
	return _RootChain.Contract.Authority(&_RootChain.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_RootChain *RootChainCallerSession) Authority() (common.Address, error) {
	return _RootChain.Contract.Authority(&_RootChain.CallOpts)
}

// ChildBlockInterval is a free data retrieval call binding the contract method 0x38a9e0bc.
//
// Solidity: function childBlockInterval() constant returns(uint256)
func (_RootChain *RootChainCaller) ChildBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "childBlockInterval")
	return *ret0, err
}

// ChildBlockInterval is a free data retrieval call binding the contract method 0x38a9e0bc.
//
// Solidity: function childBlockInterval() constant returns(uint256)
func (_RootChain *RootChainSession) ChildBlockInterval() (*big.Int, error) {
	return _RootChain.Contract.ChildBlockInterval(&_RootChain.CallOpts)
}

// ChildBlockInterval is a free data retrieval call binding the contract method 0x38a9e0bc.
//
// Solidity: function childBlockInterval() constant returns(uint256)
func (_RootChain *RootChainCallerSession) ChildBlockInterval() (*big.Int, error) {
	return _RootChain.Contract.ChildBlockInterval(&_RootChain.CallOpts)
}

// ChildChain is a free data retrieval call binding the contract method 0xf95643b1.
//
// Solidity: function childChain( uint256) constant returns(root bytes32, created_at uint256)
func (_RootChain *RootChainCaller) ChildChain(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Root      [32]byte
	CreatedAt *big.Int
}, error) {
	ret := new(struct {
		Root      [32]byte
		CreatedAt *big.Int
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "childChain", arg0)
	return *ret, err
}

// ChildChain is a free data retrieval call binding the contract method 0xf95643b1.
//
// Solidity: function childChain( uint256) constant returns(root bytes32, created_at uint256)
func (_RootChain *RootChainSession) ChildChain(arg0 *big.Int) (struct {
	Root      [32]byte
	CreatedAt *big.Int
}, error) {
	return _RootChain.Contract.ChildChain(&_RootChain.CallOpts, arg0)
}

// ChildChain is a free data retrieval call binding the contract method 0xf95643b1.
//
// Solidity: function childChain( uint256) constant returns(root bytes32, created_at uint256)
func (_RootChain *RootChainCallerSession) ChildChain(arg0 *big.Int) (struct {
	Root      [32]byte
	CreatedAt *big.Int
}, error) {
	return _RootChain.Contract.ChildChain(&_RootChain.CallOpts, arg0)
}

// CurrentChildBlock is a free data retrieval call binding the contract method 0x7a95f1e8.
//
// Solidity: function currentChildBlock() constant returns(uint256)
func (_RootChain *RootChainCaller) CurrentChildBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "currentChildBlock")
	return *ret0, err
}

// CurrentChildBlock is a free data retrieval call binding the contract method 0x7a95f1e8.
//
// Solidity: function currentChildBlock() constant returns(uint256)
func (_RootChain *RootChainSession) CurrentChildBlock() (*big.Int, error) {
	return _RootChain.Contract.CurrentChildBlock(&_RootChain.CallOpts)
}

// CurrentChildBlock is a free data retrieval call binding the contract method 0x7a95f1e8.
//
// Solidity: function currentChildBlock() constant returns(uint256)
func (_RootChain *RootChainCallerSession) CurrentChildBlock() (*big.Int, error) {
	return _RootChain.Contract.CurrentChildBlock(&_RootChain.CallOpts)
}

// CurrentDepositBlock is a free data retrieval call binding the contract method 0xa98c7f2c.
//
// Solidity: function currentDepositBlock() constant returns(uint256)
func (_RootChain *RootChainCaller) CurrentDepositBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "currentDepositBlock")
	return *ret0, err
}

// CurrentDepositBlock is a free data retrieval call binding the contract method 0xa98c7f2c.
//
// Solidity: function currentDepositBlock() constant returns(uint256)
func (_RootChain *RootChainSession) CurrentDepositBlock() (*big.Int, error) {
	return _RootChain.Contract.CurrentDepositBlock(&_RootChain.CallOpts)
}

// CurrentDepositBlock is a free data retrieval call binding the contract method 0xa98c7f2c.
//
// Solidity: function currentDepositBlock() constant returns(uint256)
func (_RootChain *RootChainCallerSession) CurrentDepositBlock() (*big.Int, error) {
	return _RootChain.Contract.CurrentDepositBlock(&_RootChain.CallOpts)
}

// CurrentFeeExit is a free data retrieval call binding the contract method 0xcfe7d855.
//
// Solidity: function currentFeeExit() constant returns(uint256)
func (_RootChain *RootChainCaller) CurrentFeeExit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "currentFeeExit")
	return *ret0, err
}

// CurrentFeeExit is a free data retrieval call binding the contract method 0xcfe7d855.
//
// Solidity: function currentFeeExit() constant returns(uint256)
func (_RootChain *RootChainSession) CurrentFeeExit() (*big.Int, error) {
	return _RootChain.Contract.CurrentFeeExit(&_RootChain.CallOpts)
}

// CurrentFeeExit is a free data retrieval call binding the contract method 0xcfe7d855.
//
// Solidity: function currentFeeExit() constant returns(uint256)
func (_RootChain *RootChainCallerSession) CurrentFeeExit() (*big.Int, error) {
	return _RootChain.Contract.CurrentFeeExit(&_RootChain.CallOpts)
}

// Exits is a free data retrieval call binding the contract method 0x342de179.
//
// Solidity: function exits( uint256) constant returns(owner address, amount uint256)
func (_RootChain *RootChainCaller) Exits(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner  common.Address
	Amount *big.Int
}, error) {
	ret := new(struct {
		Owner  common.Address
		Amount *big.Int
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "exits", arg0)
	return *ret, err
}

// Exits is a free data retrieval call binding the contract method 0x342de179.
//
// Solidity: function exits( uint256) constant returns(owner address, amount uint256)
func (_RootChain *RootChainSession) Exits(arg0 *big.Int) (struct {
	Owner  common.Address
	Amount *big.Int
}, error) {
	return _RootChain.Contract.Exits(&_RootChain.CallOpts, arg0)
}

// Exits is a free data retrieval call binding the contract method 0x342de179.
//
// Solidity: function exits( uint256) constant returns(owner address, amount uint256)
func (_RootChain *RootChainCallerSession) Exits(arg0 *big.Int) (struct {
	Owner  common.Address
	Amount *big.Int
}, error) {
	return _RootChain.Contract.Exits(&_RootChain.CallOpts, arg0)
}

// GetChildChain is a free data retrieval call binding the contract method 0x85444de3.
//
// Solidity: function getChildChain(blockNumber uint256) constant returns(bytes32, uint256)
func (_RootChain *RootChainCaller) GetChildChain(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _RootChain.contract.Call(opts, out, "getChildChain", blockNumber)
	return *ret0, *ret1, err
}

// GetChildChain is a free data retrieval call binding the contract method 0x85444de3.
//
// Solidity: function getChildChain(blockNumber uint256) constant returns(bytes32, uint256)
func (_RootChain *RootChainSession) GetChildChain(blockNumber *big.Int) ([32]byte, *big.Int, error) {
	return _RootChain.Contract.GetChildChain(&_RootChain.CallOpts, blockNumber)
}

// GetChildChain is a free data retrieval call binding the contract method 0x85444de3.
//
// Solidity: function getChildChain(blockNumber uint256) constant returns(bytes32, uint256)
func (_RootChain *RootChainCallerSession) GetChildChain(blockNumber *big.Int) ([32]byte, *big.Int, error) {
	return _RootChain.Contract.GetChildChain(&_RootChain.CallOpts, blockNumber)
}

// GetDepositBlock is a free data retrieval call binding the contract method 0xbcd59261.
//
// Solidity: function getDepositBlock() constant returns(uint256)
func (_RootChain *RootChainCaller) GetDepositBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getDepositBlock")
	return *ret0, err
}

// GetDepositBlock is a free data retrieval call binding the contract method 0xbcd59261.
//
// Solidity: function getDepositBlock() constant returns(uint256)
func (_RootChain *RootChainSession) GetDepositBlock() (*big.Int, error) {
	return _RootChain.Contract.GetDepositBlock(&_RootChain.CallOpts)
}

// GetDepositBlock is a free data retrieval call binding the contract method 0xbcd59261.
//
// Solidity: function getDepositBlock() constant returns(uint256)
func (_RootChain *RootChainCallerSession) GetDepositBlock() (*big.Int, error) {
	return _RootChain.Contract.GetDepositBlock(&_RootChain.CallOpts)
}

// GetExit is a free data retrieval call binding the contract method 0xe60f1ff1.
//
// Solidity: function getExit(utxoPos uint256) constant returns(address, uint256)
func (_RootChain *RootChainCaller) GetExit(opts *bind.CallOpts, utxoPos *big.Int) (common.Address, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _RootChain.contract.Call(opts, out, "getExit", utxoPos)
	return *ret0, *ret1, err
}

// GetExit is a free data retrieval call binding the contract method 0xe60f1ff1.
//
// Solidity: function getExit(utxoPos uint256) constant returns(address, uint256)
func (_RootChain *RootChainSession) GetExit(utxoPos *big.Int) (common.Address, *big.Int, error) {
	return _RootChain.Contract.GetExit(&_RootChain.CallOpts, utxoPos)
}

// GetExit is a free data retrieval call binding the contract method 0xe60f1ff1.
//
// Solidity: function getExit(utxoPos uint256) constant returns(address, uint256)
func (_RootChain *RootChainCallerSession) GetExit(utxoPos *big.Int) (common.Address, *big.Int, error) {
	return _RootChain.Contract.GetExit(&_RootChain.CallOpts, utxoPos)
}

// GetNextExit is a free data retrieval call binding the contract method 0x2fd56a25.
//
// Solidity: function getNextExit() constant returns(uint256, uint256)
func (_RootChain *RootChainCaller) GetNextExit(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _RootChain.contract.Call(opts, out, "getNextExit")
	return *ret0, *ret1, err
}

// GetNextExit is a free data retrieval call binding the contract method 0x2fd56a25.
//
// Solidity: function getNextExit() constant returns(uint256, uint256)
func (_RootChain *RootChainSession) GetNextExit() (*big.Int, *big.Int, error) {
	return _RootChain.Contract.GetNextExit(&_RootChain.CallOpts)
}

// GetNextExit is a free data retrieval call binding the contract method 0x2fd56a25.
//
// Solidity: function getNextExit() constant returns(uint256, uint256)
func (_RootChain *RootChainCallerSession) GetNextExit() (*big.Int, *big.Int, error) {
	return _RootChain.Contract.GetNextExit(&_RootChain.CallOpts)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x32773ba3.
//
// Solidity: function challengeExit(cUtxoPos uint256, eUtxoIndex uint256, txBytes bytes, proof bytes, sigs bytes, confirmationSig bytes) returns()
func (_RootChain *RootChainTransactor) ChallengeExit(opts *bind.TransactOpts, cUtxoPos *big.Int, eUtxoIndex *big.Int, txBytes []byte, proof []byte, sigs []byte, confirmationSig []byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "challengeExit", cUtxoPos, eUtxoIndex, txBytes, proof, sigs, confirmationSig)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x32773ba3.
//
// Solidity: function challengeExit(cUtxoPos uint256, eUtxoIndex uint256, txBytes bytes, proof bytes, sigs bytes, confirmationSig bytes) returns()
func (_RootChain *RootChainSession) ChallengeExit(cUtxoPos *big.Int, eUtxoIndex *big.Int, txBytes []byte, proof []byte, sigs []byte, confirmationSig []byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeExit(&_RootChain.TransactOpts, cUtxoPos, eUtxoIndex, txBytes, proof, sigs, confirmationSig)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x32773ba3.
//
// Solidity: function challengeExit(cUtxoPos uint256, eUtxoIndex uint256, txBytes bytes, proof bytes, sigs bytes, confirmationSig bytes) returns()
func (_RootChain *RootChainTransactorSession) ChallengeExit(cUtxoPos *big.Int, eUtxoIndex *big.Int, txBytes []byte, proof []byte, sigs []byte, confirmationSig []byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeExit(&_RootChain.TransactOpts, cUtxoPos, eUtxoIndex, txBytes, proof, sigs, confirmationSig)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_RootChain *RootChainTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_RootChain *RootChainSession) Deposit() (*types.Transaction, error) {
	return _RootChain.Contract.Deposit(&_RootChain.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_RootChain *RootChainTransactorSession) Deposit() (*types.Transaction, error) {
	return _RootChain.Contract.Deposit(&_RootChain.TransactOpts)
}

// FinalizeExits is a paid mutator transaction binding the contract method 0xc6ab44cd.
//
// Solidity: function finalizeExits() returns()
func (_RootChain *RootChainTransactor) FinalizeExits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "finalizeExits")
}

// FinalizeExits is a paid mutator transaction binding the contract method 0xc6ab44cd.
//
// Solidity: function finalizeExits() returns()
func (_RootChain *RootChainSession) FinalizeExits() (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeExits(&_RootChain.TransactOpts)
}

// FinalizeExits is a paid mutator transaction binding the contract method 0xc6ab44cd.
//
// Solidity: function finalizeExits() returns()
func (_RootChain *RootChainTransactorSession) FinalizeExits() (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeExits(&_RootChain.TransactOpts)
}

// StartDepositExit is a paid mutator transaction binding the contract method 0x70e4abf6.
//
// Solidity: function startDepositExit(depositPos uint256, amount uint256) returns()
func (_RootChain *RootChainTransactor) StartDepositExit(opts *bind.TransactOpts, depositPos *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "startDepositExit", depositPos, amount)
}

// StartDepositExit is a paid mutator transaction binding the contract method 0x70e4abf6.
//
// Solidity: function startDepositExit(depositPos uint256, amount uint256) returns()
func (_RootChain *RootChainSession) StartDepositExit(depositPos *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.StartDepositExit(&_RootChain.TransactOpts, depositPos, amount)
}

// StartDepositExit is a paid mutator transaction binding the contract method 0x70e4abf6.
//
// Solidity: function startDepositExit(depositPos uint256, amount uint256) returns()
func (_RootChain *RootChainTransactorSession) StartDepositExit(depositPos *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.StartDepositExit(&_RootChain.TransactOpts, depositPos, amount)
}

// StartExit is a paid mutator transaction binding the contract method 0x1c91a6b9.
//
// Solidity: function startExit(utxoPos uint256, txBytes bytes, proof bytes, sigs bytes) returns()
func (_RootChain *RootChainTransactor) StartExit(opts *bind.TransactOpts, utxoPos *big.Int, txBytes []byte, proof []byte, sigs []byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "startExit", utxoPos, txBytes, proof, sigs)
}

// StartExit is a paid mutator transaction binding the contract method 0x1c91a6b9.
//
// Solidity: function startExit(utxoPos uint256, txBytes bytes, proof bytes, sigs bytes) returns()
func (_RootChain *RootChainSession) StartExit(utxoPos *big.Int, txBytes []byte, proof []byte, sigs []byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, utxoPos, txBytes, proof, sigs)
}

// StartExit is a paid mutator transaction binding the contract method 0x1c91a6b9.
//
// Solidity: function startExit(utxoPos uint256, txBytes bytes, proof bytes, sigs bytes) returns()
func (_RootChain *RootChainTransactorSession) StartExit(utxoPos *big.Int, txBytes []byte, proof []byte, sigs []byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, utxoPos, txBytes, proof, sigs)
}

// StartFeeExit is a paid mutator transaction binding the contract method 0xae80d8c8.
//
// Solidity: function startFeeExit(amount uint256) returns(uint256)
func (_RootChain *RootChainTransactor) StartFeeExit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "startFeeExit", amount)
}

// StartFeeExit is a paid mutator transaction binding the contract method 0xae80d8c8.
//
// Solidity: function startFeeExit(amount uint256) returns(uint256)
func (_RootChain *RootChainSession) StartFeeExit(amount *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.StartFeeExit(&_RootChain.TransactOpts, amount)
}

// StartFeeExit is a paid mutator transaction binding the contract method 0xae80d8c8.
//
// Solidity: function startFeeExit(amount uint256) returns(uint256)
func (_RootChain *RootChainTransactorSession) StartFeeExit(amount *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.StartFeeExit(&_RootChain.TransactOpts, amount)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(root bytes32) returns()
func (_RootChain *RootChainTransactor) SubmitBlock(opts *bind.TransactOpts, root [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "submitBlock", root)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(root bytes32) returns()
func (_RootChain *RootChainSession) SubmitBlock(root [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitBlock(&_RootChain.TransactOpts, root)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(root bytes32) returns()
func (_RootChain *RootChainTransactorSession) SubmitBlock(root [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitBlock(&_RootChain.TransactOpts, root)
}

// RootChainDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the RootChain contract.
type RootChainDepositIterator struct {
	Event *RootChainDeposit // Event containing the contract specifics and raw log

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
func (it *RootChainDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainDeposit)
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
		it.Event = new(RootChainDeposit)
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
func (it *RootChainDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainDeposit represents a Deposit event raised by the RootChain contract.
type RootChainDeposit struct {
	Depositor    common.Address
	DepositBlock *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(depositor indexed address, depositBlock indexed uint256, amount uint256)
func (_RootChain *RootChainFilterer) FilterDeposit(opts *bind.FilterOpts, depositor []common.Address, depositBlock []*big.Int) (*RootChainDepositIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var depositBlockRule []interface{}
	for _, depositBlockItem := range depositBlock {
		depositBlockRule = append(depositBlockRule, depositBlockItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "Deposit", depositorRule, depositBlockRule)
	if err != nil {
		return nil, err
	}
	return &RootChainDepositIterator{contract: _RootChain.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(depositor indexed address, depositBlock indexed uint256, amount uint256)
func (_RootChain *RootChainFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *RootChainDeposit, depositor []common.Address, depositBlock []*big.Int) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var depositBlockRule []interface{}
	for _, depositBlockItem := range depositBlock {
		depositBlockRule = append(depositBlockRule, depositBlockItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "Deposit", depositorRule, depositBlockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainDeposit)
				if err := _RootChain.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// RootChainExitStartedIterator is returned from FilterExitStarted and is used to iterate over the raw logs and unpacked data for ExitStarted events raised by the RootChain contract.
type RootChainExitStartedIterator struct {
	Event *RootChainExitStarted // Event containing the contract specifics and raw log

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
func (it *RootChainExitStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainExitStarted)
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
		it.Event = new(RootChainExitStarted)
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
func (it *RootChainExitStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainExitStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainExitStarted represents a ExitStarted event raised by the RootChain contract.
type RootChainExitStarted struct {
	Exitor  common.Address
	UtxoPos *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExitStarted is a free log retrieval operation binding the contract event 0xd62f48e64bfb0b95c15737946c829e7fe6378cd679a268c1ca5993dc235d900b.
//
// Solidity: event ExitStarted(exitor indexed address, utxoPos indexed uint256, amount uint256)
func (_RootChain *RootChainFilterer) FilterExitStarted(opts *bind.FilterOpts, exitor []common.Address, utxoPos []*big.Int) (*RootChainExitStartedIterator, error) {

	var exitorRule []interface{}
	for _, exitorItem := range exitor {
		exitorRule = append(exitorRule, exitorItem)
	}
	var utxoPosRule []interface{}
	for _, utxoPosItem := range utxoPos {
		utxoPosRule = append(utxoPosRule, utxoPosItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "ExitStarted", exitorRule, utxoPosRule)
	if err != nil {
		return nil, err
	}
	return &RootChainExitStartedIterator{contract: _RootChain.contract, event: "ExitStarted", logs: logs, sub: sub}, nil
}

// WatchExitStarted is a free log subscription operation binding the contract event 0xd62f48e64bfb0b95c15737946c829e7fe6378cd679a268c1ca5993dc235d900b.
//
// Solidity: event ExitStarted(exitor indexed address, utxoPos indexed uint256, amount uint256)
func (_RootChain *RootChainFilterer) WatchExitStarted(opts *bind.WatchOpts, sink chan<- *RootChainExitStarted, exitor []common.Address, utxoPos []*big.Int) (event.Subscription, error) {

	var exitorRule []interface{}
	for _, exitorItem := range exitor {
		exitorRule = append(exitorRule, exitorItem)
	}
	var utxoPosRule []interface{}
	for _, utxoPosItem := range utxoPos {
		utxoPosRule = append(utxoPosRule, utxoPosItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "ExitStarted", exitorRule, utxoPosRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainExitStarted)
				if err := _RootChain.contract.UnpackLog(event, "ExitStarted", log); err != nil {
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

// RootChainSubmitIterator is returned from FilterSubmit and is used to iterate over the raw logs and unpacked data for Submit events raised by the RootChain contract.
type RootChainSubmitIterator struct {
	Event *RootChainSubmit // Event containing the contract specifics and raw log

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
func (it *RootChainSubmitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainSubmit)
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
		it.Event = new(RootChainSubmit)
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
func (it *RootChainSubmitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainSubmitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainSubmit represents a Submit event raised by the RootChain contract.
type RootChainSubmit struct {
	SubmitBlock *big.Int
	Root        [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSubmit is a free log retrieval operation binding the contract event 0x03e03b35d08c09e37050f76f1a914a334810af275603f96ed8ad6f2cc3617eba.
//
// Solidity: event Submit(submitBlock indexed uint256, root bytes32)
func (_RootChain *RootChainFilterer) FilterSubmit(opts *bind.FilterOpts, submitBlock []*big.Int) (*RootChainSubmitIterator, error) {

	var submitBlockRule []interface{}
	for _, submitBlockItem := range submitBlock {
		submitBlockRule = append(submitBlockRule, submitBlockItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "Submit", submitBlockRule)
	if err != nil {
		return nil, err
	}
	return &RootChainSubmitIterator{contract: _RootChain.contract, event: "Submit", logs: logs, sub: sub}, nil
}

// WatchSubmit is a free log subscription operation binding the contract event 0x03e03b35d08c09e37050f76f1a914a334810af275603f96ed8ad6f2cc3617eba.
//
// Solidity: event Submit(submitBlock indexed uint256, root bytes32)
func (_RootChain *RootChainFilterer) WatchSubmit(opts *bind.WatchOpts, sink chan<- *RootChainSubmit, submitBlock []*big.Int) (event.Subscription, error) {

	var submitBlockRule []interface{}
	for _, submitBlockItem := range submitBlock {
		submitBlockRule = append(submitBlockRule, submitBlockItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "Submit", submitBlockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainSubmit)
				if err := _RootChain.contract.UnpackLog(event, "Submit", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820117122f6316c1af0e5aa769cc3004989356f14473ece36926dcb37b714a9c40b0029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// ValidateABI is the input ABI used to generate the binding from.
const ValidateABI = "[]"

// ValidateBin is the compiled bytecode used for deploying new contracts.
const ValidateBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582039cc8f2cf483ea8891a9e81cb7d4ebd4b210f2ea7827a567efdf6127a38e4c0e0029`

// DeployValidate deploys a new Ethereum contract, binding an instance of Validate to it.
func DeployValidate(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Validate, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidateABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Validate{ValidateCaller: ValidateCaller{contract: contract}, ValidateTransactor: ValidateTransactor{contract: contract}, ValidateFilterer: ValidateFilterer{contract: contract}}, nil
}

// Validate is an auto generated Go binding around an Ethereum contract.
type Validate struct {
	ValidateCaller     // Read-only binding to the contract
	ValidateTransactor // Write-only binding to the contract
	ValidateFilterer   // Log filterer for contract events
}

// ValidateCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidateSession struct {
	Contract     *Validate         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidateCallerSession struct {
	Contract *ValidateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ValidateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidateTransactorSession struct {
	Contract     *ValidateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ValidateRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidateRaw struct {
	Contract *Validate // Generic contract binding to access the raw methods on
}

// ValidateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidateCallerRaw struct {
	Contract *ValidateCaller // Generic read-only contract binding to access the raw methods on
}

// ValidateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidateTransactorRaw struct {
	Contract *ValidateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidate creates a new instance of Validate, bound to a specific deployed contract.
func NewValidate(address common.Address, backend bind.ContractBackend) (*Validate, error) {
	contract, err := bindValidate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validate{ValidateCaller: ValidateCaller{contract: contract}, ValidateTransactor: ValidateTransactor{contract: contract}, ValidateFilterer: ValidateFilterer{contract: contract}}, nil
}

// NewValidateCaller creates a new read-only instance of Validate, bound to a specific deployed contract.
func NewValidateCaller(address common.Address, caller bind.ContractCaller) (*ValidateCaller, error) {
	contract, err := bindValidate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidateCaller{contract: contract}, nil
}

// NewValidateTransactor creates a new write-only instance of Validate, bound to a specific deployed contract.
func NewValidateTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidateTransactor, error) {
	contract, err := bindValidate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidateTransactor{contract: contract}, nil
}

// NewValidateFilterer creates a new log filterer instance of Validate, bound to a specific deployed contract.
func NewValidateFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidateFilterer, error) {
	contract, err := bindValidate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidateFilterer{contract: contract}, nil
}

// bindValidate binds a generic wrapper to an already deployed contract.
func bindValidate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validate *ValidateRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validate.Contract.ValidateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validate *ValidateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validate.Contract.ValidateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validate *ValidateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validate.Contract.ValidateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validate *ValidateCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validate *ValidateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validate *ValidateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validate.Contract.contract.Transact(opts, method, params...)
}

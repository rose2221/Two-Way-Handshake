// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ContractBMetaData contains all meta data concerning the ContractB contract.
var ContractBMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_counterparty\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_relayer\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acknowledge\",\"inputs\":[{\"name\":\"sessionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"confirm\",\"inputs\":[{\"name\":\"sessionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"counterparty\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"handshakeComplete\",\"inputs\":[{\"name\":\"sessionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"relayer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sessions\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumContractB.Status\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setRelayer\",\"inputs\":[{\"name\":\"newRelayer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Ack\",\"inputs\":[{\"name\":\"sessionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Done\",\"inputs\":[{\"name\":\"sessionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyRelayer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b50604051610e81380380610e81833981810160405281019061003191906102f3565b335f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100a2575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016100999190610340565b60405180910390fd5b6100b1816101d460201b60201c565b505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415801561011a57505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b610159576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610150906103b3565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506103d1565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102c282610299565b9050919050565b6102d2816102b8565b81146102dc575f5ffd5b50565b5f815190506102ed816102c9565b92915050565b5f5f6040838503121561030957610308610295565b5b5f610316858286016102df565b9250506020610327858286016102df565b9150509250929050565b61033a816102b8565b82525050565b5f6020820190506103535f830184610331565b92915050565b5f82825260208201905092915050565b7f7a65726f000000000000000000000000000000000000000000000000000000005f82015250565b5f61039d600483610359565b91506103a882610369565b602082019050919050565b5f6020820190508181035f8301526103ca81610391565b9050919050565b608051610a986103e95f395f6106270152610a985ff3fe608060405234801561000f575f5ffd5b506004361061009c575f3560e01c80638406c079116100645780638406c079146101425780638da5cb5b14610160578063ba0179b51461017e578063c3df65461461019a578063f2fde38b146101b85761009c565b8063403e3fba146100a05780634b5809e1146100d05780636548e9bc146100ec578063715018a61461010857806383c4b7a314610112575b5f5ffd5b6100ba60048036038101906100b59190610853565b6101d4565b6040516100c79190610898565b60405180910390f35b6100ea60048036038101906100e59190610853565b610222565b005b6101066004803603810190610101919061090b565b610384565b005b61011061043d565b005b61012c60048036038101906101279190610853565b610450565b60405161013991906109a9565b60405180910390f35b61014a610476565b60405161015791906109d1565b60405180910390f35b61016861049b565b60405161017591906109d1565b60405180910390f35b61019860048036038101906101939190610853565b6104c2565b005b6101a2610625565b6040516101af91906109d1565b60405180910390f35b6101d260048036038101906101cd919061090b565b610649565b005b5f6002808111156101e8576101e7610936565b5b60015f8481526020019081526020015f205f015f9054906101000a900460ff16600281111561021a57610219610936565b5b149050919050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102a8576040517f4578ddb800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60015f8381526020019081526020015f2090505f60028111156102cf576102ce610936565b5b815f015f9054906101000a900460ff1660028111156102f1576102f0610936565b5b14610328576040517fbaf3f0f700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001815f015f6101000a81548160ff0219169083600281111561034e5761034d610936565b5b0217905550817fb6f6b52300386ec96e95334e72607a9cc8e2cb41f75aac41eb81fda43a43fde860405160405180910390a25050565b61038c6106cd565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036103fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f190610a44565b60405180910390fd5b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6104456106cd565b61044e5f610754565b565b6001602052805f5260405f205f91509050805f015f9054906101000a900460ff16905081565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610548576040517f4578ddb800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60015f8381526020019081526020015f209050600160028111156105705761056f610936565b5b815f015f9054906101000a900460ff16600281111561059257610591610936565b5b146105c9576040517fbaf3f0f700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002815f015f6101000a81548160ff021916908360028111156105ef576105ee610936565b5b0217905550817f6bb841348c5a71169a2db8779d29699afa576c107c1bf7c33c3193ae1e980ba260405160405180910390a25050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6106516106cd565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036106c1575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016106b891906109d1565b60405180910390fd5b6106ca81610754565b50565b6106d5610815565b73ffffffffffffffffffffffffffffffffffffffff166106f361049b565b73ffffffffffffffffffffffffffffffffffffffff161461075257610716610815565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161074991906109d1565b60405180910390fd5b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f33905090565b5f5ffd5b5f819050919050565b61083281610820565b811461083c575f5ffd5b50565b5f8135905061084d81610829565b92915050565b5f602082840312156108685761086761081c565b5b5f6108758482850161083f565b91505092915050565b5f8115159050919050565b6108928161087e565b82525050565b5f6020820190506108ab5f830184610889565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6108da826108b1565b9050919050565b6108ea816108d0565b81146108f4575f5ffd5b50565b5f81359050610905816108e1565b92915050565b5f602082840312156109205761091f61081c565b5b5f61092d848285016108f7565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6003811061097457610973610936565b5b50565b5f81905061098482610963565b919050565b5f61099382610977565b9050919050565b6109a381610989565b82525050565b5f6020820190506109bc5f83018461099a565b92915050565b6109cb816108d0565b82525050565b5f6020820190506109e45f8301846109c2565b92915050565b5f82825260208201905092915050565b7f7a65726f000000000000000000000000000000000000000000000000000000005f82015250565b5f610a2e6004836109ea565b9150610a39826109fa565b602082019050919050565b5f6020820190508181035f830152610a5b81610a22565b905091905056fea26469706673582212208b6c9e3fbabec9db08df5ab9bcd42233e6c826bac6a09017eeacc5aac0a684a564736f6c634300081c0033",
}

// ContractBABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractBMetaData.ABI instead.
var ContractBABI = ContractBMetaData.ABI

// ContractBBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractBMetaData.Bin instead.
var ContractBBin = ContractBMetaData.Bin

// DeployContractB deploys a new Ethereum contract, binding an instance of ContractB to it.
func DeployContractB(auth *bind.TransactOpts, backend bind.ContractBackend, _counterparty common.Address, _relayer common.Address) (common.Address, *types.Transaction, *ContractB, error) {
	parsed, err := ContractBMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBBin), backend, _counterparty, _relayer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractB{ContractBCaller: ContractBCaller{contract: contract}, ContractBTransactor: ContractBTransactor{contract: contract}, ContractBFilterer: ContractBFilterer{contract: contract}}, nil
}

// ContractB is an auto generated Go binding around an Ethereum contract.
type ContractB struct {
	ContractBCaller     // Read-only binding to the contract
	ContractBTransactor // Write-only binding to the contract
	ContractBFilterer   // Log filterer for contract events
}

// ContractBCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractBCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractBTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractBFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractBSession struct {
	Contract     *ContractB        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractBCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractBCallerSession struct {
	Contract *ContractBCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractBTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractBTransactorSession struct {
	Contract     *ContractBTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractBRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractBRaw struct {
	Contract *ContractB // Generic contract binding to access the raw methods on
}

// ContractBCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractBCallerRaw struct {
	Contract *ContractBCaller // Generic read-only contract binding to access the raw methods on
}

// ContractBTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractBTransactorRaw struct {
	Contract *ContractBTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractB creates a new instance of ContractB, bound to a specific deployed contract.
func NewContractB(address common.Address, backend bind.ContractBackend) (*ContractB, error) {
	contract, err := bindContractB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractB{ContractBCaller: ContractBCaller{contract: contract}, ContractBTransactor: ContractBTransactor{contract: contract}, ContractBFilterer: ContractBFilterer{contract: contract}}, nil
}

// NewContractBCaller creates a new read-only instance of ContractB, bound to a specific deployed contract.
func NewContractBCaller(address common.Address, caller bind.ContractCaller) (*ContractBCaller, error) {
	contract, err := bindContractB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractBCaller{contract: contract}, nil
}

// NewContractBTransactor creates a new write-only instance of ContractB, bound to a specific deployed contract.
func NewContractBTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractBTransactor, error) {
	contract, err := bindContractB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractBTransactor{contract: contract}, nil
}

// NewContractBFilterer creates a new log filterer instance of ContractB, bound to a specific deployed contract.
func NewContractBFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractBFilterer, error) {
	contract, err := bindContractB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractBFilterer{contract: contract}, nil
}

// bindContractB binds a generic wrapper to an already deployed contract.
func bindContractB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractBMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractB *ContractBRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractB.Contract.ContractBCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractB *ContractBRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractB.Contract.ContractBTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractB *ContractBRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractB.Contract.ContractBTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractB *ContractBCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractB *ContractBTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractB *ContractBTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractB.Contract.contract.Transact(opts, method, params...)
}

// Counterparty is a free data retrieval call binding the contract method 0xc3df6546.
//
// Solidity: function counterparty() view returns(address)
func (_ContractB *ContractBCaller) Counterparty(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractB.contract.Call(opts, &out, "counterparty")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterparty is a free data retrieval call binding the contract method 0xc3df6546.
//
// Solidity: function counterparty() view returns(address)
func (_ContractB *ContractBSession) Counterparty() (common.Address, error) {
	return _ContractB.Contract.Counterparty(&_ContractB.CallOpts)
}

// Counterparty is a free data retrieval call binding the contract method 0xc3df6546.
//
// Solidity: function counterparty() view returns(address)
func (_ContractB *ContractBCallerSession) Counterparty() (common.Address, error) {
	return _ContractB.Contract.Counterparty(&_ContractB.CallOpts)
}

// HandshakeComplete is a free data retrieval call binding the contract method 0x403e3fba.
//
// Solidity: function handshakeComplete(uint256 sessionId) view returns(bool)
func (_ContractB *ContractBCaller) HandshakeComplete(opts *bind.CallOpts, sessionId *big.Int) (bool, error) {
	var out []interface{}
	err := _ContractB.contract.Call(opts, &out, "handshakeComplete", sessionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HandshakeComplete is a free data retrieval call binding the contract method 0x403e3fba.
//
// Solidity: function handshakeComplete(uint256 sessionId) view returns(bool)
func (_ContractB *ContractBSession) HandshakeComplete(sessionId *big.Int) (bool, error) {
	return _ContractB.Contract.HandshakeComplete(&_ContractB.CallOpts, sessionId)
}

// HandshakeComplete is a free data retrieval call binding the contract method 0x403e3fba.
//
// Solidity: function handshakeComplete(uint256 sessionId) view returns(bool)
func (_ContractB *ContractBCallerSession) HandshakeComplete(sessionId *big.Int) (bool, error) {
	return _ContractB.Contract.HandshakeComplete(&_ContractB.CallOpts, sessionId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractB *ContractBCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractB.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractB *ContractBSession) Owner() (common.Address, error) {
	return _ContractB.Contract.Owner(&_ContractB.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractB *ContractBCallerSession) Owner() (common.Address, error) {
	return _ContractB.Contract.Owner(&_ContractB.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_ContractB *ContractBCaller) Relayer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractB.contract.Call(opts, &out, "relayer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_ContractB *ContractBSession) Relayer() (common.Address, error) {
	return _ContractB.Contract.Relayer(&_ContractB.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_ContractB *ContractBCallerSession) Relayer() (common.Address, error) {
	return _ContractB.Contract.Relayer(&_ContractB.CallOpts)
}

// Sessions is a free data retrieval call binding the contract method 0x83c4b7a3.
//
// Solidity: function sessions(uint256 ) view returns(uint8 status)
func (_ContractB *ContractBCaller) Sessions(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _ContractB.contract.Call(opts, &out, "sessions", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Sessions is a free data retrieval call binding the contract method 0x83c4b7a3.
//
// Solidity: function sessions(uint256 ) view returns(uint8 status)
func (_ContractB *ContractBSession) Sessions(arg0 *big.Int) (uint8, error) {
	return _ContractB.Contract.Sessions(&_ContractB.CallOpts, arg0)
}

// Sessions is a free data retrieval call binding the contract method 0x83c4b7a3.
//
// Solidity: function sessions(uint256 ) view returns(uint8 status)
func (_ContractB *ContractBCallerSession) Sessions(arg0 *big.Int) (uint8, error) {
	return _ContractB.Contract.Sessions(&_ContractB.CallOpts, arg0)
}

// Acknowledge is a paid mutator transaction binding the contract method 0x4b5809e1.
//
// Solidity: function acknowledge(uint256 sessionId) returns()
func (_ContractB *ContractBTransactor) Acknowledge(opts *bind.TransactOpts, sessionId *big.Int) (*types.Transaction, error) {
	return _ContractB.contract.Transact(opts, "acknowledge", sessionId)
}

// Acknowledge is a paid mutator transaction binding the contract method 0x4b5809e1.
//
// Solidity: function acknowledge(uint256 sessionId) returns()
func (_ContractB *ContractBSession) Acknowledge(sessionId *big.Int) (*types.Transaction, error) {
	return _ContractB.Contract.Acknowledge(&_ContractB.TransactOpts, sessionId)
}

// Acknowledge is a paid mutator transaction binding the contract method 0x4b5809e1.
//
// Solidity: function acknowledge(uint256 sessionId) returns()
func (_ContractB *ContractBTransactorSession) Acknowledge(sessionId *big.Int) (*types.Transaction, error) {
	return _ContractB.Contract.Acknowledge(&_ContractB.TransactOpts, sessionId)
}

// Confirm is a paid mutator transaction binding the contract method 0xba0179b5.
//
// Solidity: function confirm(uint256 sessionId) returns()
func (_ContractB *ContractBTransactor) Confirm(opts *bind.TransactOpts, sessionId *big.Int) (*types.Transaction, error) {
	return _ContractB.contract.Transact(opts, "confirm", sessionId)
}

// Confirm is a paid mutator transaction binding the contract method 0xba0179b5.
//
// Solidity: function confirm(uint256 sessionId) returns()
func (_ContractB *ContractBSession) Confirm(sessionId *big.Int) (*types.Transaction, error) {
	return _ContractB.Contract.Confirm(&_ContractB.TransactOpts, sessionId)
}

// Confirm is a paid mutator transaction binding the contract method 0xba0179b5.
//
// Solidity: function confirm(uint256 sessionId) returns()
func (_ContractB *ContractBTransactorSession) Confirm(sessionId *big.Int) (*types.Transaction, error) {
	return _ContractB.Contract.Confirm(&_ContractB.TransactOpts, sessionId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractB *ContractBTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractB.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractB *ContractBSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractB.Contract.RenounceOwnership(&_ContractB.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractB *ContractBTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractB.Contract.RenounceOwnership(&_ContractB.TransactOpts)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address newRelayer) returns()
func (_ContractB *ContractBTransactor) SetRelayer(opts *bind.TransactOpts, newRelayer common.Address) (*types.Transaction, error) {
	return _ContractB.contract.Transact(opts, "setRelayer", newRelayer)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address newRelayer) returns()
func (_ContractB *ContractBSession) SetRelayer(newRelayer common.Address) (*types.Transaction, error) {
	return _ContractB.Contract.SetRelayer(&_ContractB.TransactOpts, newRelayer)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address newRelayer) returns()
func (_ContractB *ContractBTransactorSession) SetRelayer(newRelayer common.Address) (*types.Transaction, error) {
	return _ContractB.Contract.SetRelayer(&_ContractB.TransactOpts, newRelayer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractB *ContractBTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractB.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractB *ContractBSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractB.Contract.TransferOwnership(&_ContractB.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractB *ContractBTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractB.Contract.TransferOwnership(&_ContractB.TransactOpts, newOwner)
}

// ContractBAckIterator is returned from FilterAck and is used to iterate over the raw logs and unpacked data for Ack events raised by the ContractB contract.
type ContractBAckIterator struct {
	Event *ContractBAck // Event containing the contract specifics and raw log

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
func (it *ContractBAckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBAck)
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
		it.Event = new(ContractBAck)
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
func (it *ContractBAckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBAckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBAck represents a Ack event raised by the ContractB contract.
type ContractBAck struct {
	SessionId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAck is a free log retrieval operation binding the contract event 0xb6f6b52300386ec96e95334e72607a9cc8e2cb41f75aac41eb81fda43a43fde8.
//
// Solidity: event Ack(uint256 indexed sessionId)
func (_ContractB *ContractBFilterer) FilterAck(opts *bind.FilterOpts, sessionId []*big.Int) (*ContractBAckIterator, error) {

	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}

	logs, sub, err := _ContractB.contract.FilterLogs(opts, "Ack", sessionIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractBAckIterator{contract: _ContractB.contract, event: "Ack", logs: logs, sub: sub}, nil
}

// WatchAck is a free log subscription operation binding the contract event 0xb6f6b52300386ec96e95334e72607a9cc8e2cb41f75aac41eb81fda43a43fde8.
//
// Solidity: event Ack(uint256 indexed sessionId)
func (_ContractB *ContractBFilterer) WatchAck(opts *bind.WatchOpts, sink chan<- *ContractBAck, sessionId []*big.Int) (event.Subscription, error) {

	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}

	logs, sub, err := _ContractB.contract.WatchLogs(opts, "Ack", sessionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBAck)
				if err := _ContractB.contract.UnpackLog(event, "Ack", log); err != nil {
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

// ParseAck is a log parse operation binding the contract event 0xb6f6b52300386ec96e95334e72607a9cc8e2cb41f75aac41eb81fda43a43fde8.
//
// Solidity: event Ack(uint256 indexed sessionId)
func (_ContractB *ContractBFilterer) ParseAck(log types.Log) (*ContractBAck, error) {
	event := new(ContractBAck)
	if err := _ContractB.contract.UnpackLog(event, "Ack", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBDoneIterator is returned from FilterDone and is used to iterate over the raw logs and unpacked data for Done events raised by the ContractB contract.
type ContractBDoneIterator struct {
	Event *ContractBDone // Event containing the contract specifics and raw log

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
func (it *ContractBDoneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBDone)
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
		it.Event = new(ContractBDone)
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
func (it *ContractBDoneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBDoneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBDone represents a Done event raised by the ContractB contract.
type ContractBDone struct {
	SessionId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDone is a free log retrieval operation binding the contract event 0x6bb841348c5a71169a2db8779d29699afa576c107c1bf7c33c3193ae1e980ba2.
//
// Solidity: event Done(uint256 indexed sessionId)
func (_ContractB *ContractBFilterer) FilterDone(opts *bind.FilterOpts, sessionId []*big.Int) (*ContractBDoneIterator, error) {

	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}

	logs, sub, err := _ContractB.contract.FilterLogs(opts, "Done", sessionIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractBDoneIterator{contract: _ContractB.contract, event: "Done", logs: logs, sub: sub}, nil
}

// WatchDone is a free log subscription operation binding the contract event 0x6bb841348c5a71169a2db8779d29699afa576c107c1bf7c33c3193ae1e980ba2.
//
// Solidity: event Done(uint256 indexed sessionId)
func (_ContractB *ContractBFilterer) WatchDone(opts *bind.WatchOpts, sink chan<- *ContractBDone, sessionId []*big.Int) (event.Subscription, error) {

	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}

	logs, sub, err := _ContractB.contract.WatchLogs(opts, "Done", sessionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBDone)
				if err := _ContractB.contract.UnpackLog(event, "Done", log); err != nil {
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

// ParseDone is a log parse operation binding the contract event 0x6bb841348c5a71169a2db8779d29699afa576c107c1bf7c33c3193ae1e980ba2.
//
// Solidity: event Done(uint256 indexed sessionId)
func (_ContractB *ContractBFilterer) ParseDone(log types.Log) (*ContractBDone, error) {
	event := new(ContractBDone)
	if err := _ContractB.contract.UnpackLog(event, "Done", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractB contract.
type ContractBOwnershipTransferredIterator struct {
	Event *ContractBOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractBOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBOwnershipTransferred)
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
		it.Event = new(ContractBOwnershipTransferred)
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
func (it *ContractBOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBOwnershipTransferred represents a OwnershipTransferred event raised by the ContractB contract.
type ContractBOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractB *ContractBFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractBOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractB.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractBOwnershipTransferredIterator{contract: _ContractB.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractB *ContractBFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractBOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractB.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBOwnershipTransferred)
				if err := _ContractB.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractB *ContractBFilterer) ParseOwnershipTransferred(log types.Log) (*ContractBOwnershipTransferred, error) {
	event := new(ContractBOwnershipTransferred)
	if err := _ContractB.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

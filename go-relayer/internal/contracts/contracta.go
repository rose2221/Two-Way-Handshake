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

// ContractAMetaData contains all meta data concerning the ContractA contract.
var ContractAMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_counterparty\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_relayer\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acknowledge\",\"inputs\":[{\"name\":\"sessionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"counterparty\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"handshakeComplete\",\"inputs\":[{\"name\":\"sessionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initiateHandshake\",\"inputs\":[{\"name\":\"sessionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"relayer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sessions\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumContractA.Status\"},{\"name\":\"initiator\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setRelayer\",\"inputs\":[{\"name\":\"newRelayer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Syn\",\"inputs\":[{\"name\":\"sessionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"initiator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SynAck\",\"inputs\":[{\"name\":\"sessionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyRelayer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SessionExists\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b50604051610edb380380610edb833981810160405281019061003191906102f3565b335f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100a2575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016100999190610340565b60405180910390fd5b6100b1816101d460201b60201c565b505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415801561011a57505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b610159576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610150906103b3565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506103d1565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102c282610299565b9050919050565b6102d2816102b8565b81146102dc575f5ffd5b50565b5f815190506102ed816102c9565b92915050565b5f5f6040838503121561030957610308610295565b5b5f610316858286016102df565b9250506020610327858286016102df565b9150509250929050565b61033a816102b8565b82525050565b5f6020820190506103535f830184610331565b92915050565b5f82825260208201905092915050565b7f7a65726f000000000000000000000000000000000000000000000000000000005f82015250565b5f61039d600483610359565b91506103a882610369565b602082019050919050565b5f6020820190508181035f8301526103ca81610391565b9050919050565b608051610af26103e95f395f6106730152610af25ff3fe608060405234801561000f575f5ffd5b506004361061009c575f3560e01c806383c4b7a31161006457806383c4b7a31461012e5780638406c0791461015f5780638da5cb5b1461017d578063c3df65461461019b578063f2fde38b146101b95761009c565b806319f99ad3146100a0578063403e3fba146100bc5780634b5809e1146100ec5780636548e9bc14610108578063715018a614610124575b5f5ffd5b6100ba60048036038101906100b5919061089f565b6101d5565b005b6100d660048036038101906100d1919061089f565b61035d565b6040516100e391906108e4565b60405180910390f35b6101066004803603810190610101919061089f565b6103ab565b005b610122600480360381019061011d9190610957565b61050e565b005b61012c6105c7565b005b6101486004803603810190610143919061089f565b6105da565b604051610156929190610a04565b60405180910390f35b610167610625565b6040516101749190610a2b565b60405180910390f35b61018561064a565b6040516101929190610a2b565b60405180910390f35b6101a3610671565b6040516101b09190610a2b565b60405180910390f35b6101d360048036038101906101ce9190610957565b610695565b005b5f60038111156101e8576101e7610982565b5b60015f8381526020019081526020015f205f015f9054906101000a900460ff16600381111561021a57610219610982565b5b14610251576040517faf0dd7f000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518060400160405280600160038111156102705761026f610982565b5b81526020013373ffffffffffffffffffffffffffffffffffffffff1681525060015f8381526020019081526020015f205f820151815f015f6101000a81548160ff021916908360038111156102c8576102c7610982565b5b02179055506020820151815f0160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050503373ffffffffffffffffffffffffffffffffffffffff16817f4f4922f70c486c95340cec94bcfc794500438c5a6a4033deb6533cd74ca60e8d60405160405180910390a350565b5f60038081111561037157610370610982565b5b60015f8481526020019081526020015f205f015f9054906101000a900460ff1660038111156103a3576103a2610982565b5b149050919050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610431576040517f4578ddb800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60015f8381526020019081526020015f2090506001600381111561045957610458610982565b5b815f015f9054906101000a900460ff16600381111561047b5761047a610982565b5b146104b2576040517fbaf3f0f700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6003815f015f6101000a81548160ff021916908360038111156104d8576104d7610982565b5b0217905550817fec71d58db74e9fc73a8e4e7332027b119c6a36ff6f4fcf16447f25cd30c6b5ca60405160405180910390a25050565b610516610719565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610584576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161057b90610a9e565b60405180910390fd5b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6105cf610719565b6105d85f6107a0565b565b6001602052805f5260405f205f91509050805f015f9054906101000a900460ff1690805f0160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b61069d610719565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361070d575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016107049190610a2b565b60405180910390fd5b610716816107a0565b50565b610721610861565b73ffffffffffffffffffffffffffffffffffffffff1661073f61064a565b73ffffffffffffffffffffffffffffffffffffffff161461079e57610762610861565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016107959190610a2b565b60405180910390fd5b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f33905090565b5f5ffd5b5f819050919050565b61087e8161086c565b8114610888575f5ffd5b50565b5f8135905061089981610875565b92915050565b5f602082840312156108b4576108b3610868565b5b5f6108c18482850161088b565b91505092915050565b5f8115159050919050565b6108de816108ca565b82525050565b5f6020820190506108f75f8301846108d5565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610926826108fd565b9050919050565b6109368161091c565b8114610940575f5ffd5b50565b5f813590506109518161092d565b92915050565b5f6020828403121561096c5761096b610868565b5b5f61097984828501610943565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600481106109c0576109bf610982565b5b50565b5f8190506109d0826109af565b919050565b5f6109df826109c3565b9050919050565b6109ef816109d5565b82525050565b6109fe8161091c565b82525050565b5f604082019050610a175f8301856109e6565b610a2460208301846109f5565b9392505050565b5f602082019050610a3e5f8301846109f5565b92915050565b5f82825260208201905092915050565b7f7a65726f000000000000000000000000000000000000000000000000000000005f82015250565b5f610a88600483610a44565b9150610a9382610a54565b602082019050919050565b5f6020820190508181035f830152610ab581610a7c565b905091905056fea26469706673582212206dcaffb64b1afaba7fa479fefb41325444f617ded1733eff99af7443388a62a164736f6c634300081c0033",
}

// ContractAABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAMetaData.ABI instead.
var ContractAABI = ContractAMetaData.ABI

// ContractABin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractAMetaData.Bin instead.
var ContractABin = ContractAMetaData.Bin

// DeployContractA deploys a new Ethereum contract, binding an instance of ContractA to it.
func DeployContractA(auth *bind.TransactOpts, backend bind.ContractBackend, _counterparty common.Address, _relayer common.Address) (common.Address, *types.Transaction, *ContractA, error) {
	parsed, err := ContractAMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractABin), backend, _counterparty, _relayer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractA{ContractACaller: ContractACaller{contract: contract}, ContractATransactor: ContractATransactor{contract: contract}, ContractAFilterer: ContractAFilterer{contract: contract}}, nil
}

// ContractA is an auto generated Go binding around an Ethereum contract.
type ContractA struct {
	ContractACaller     // Read-only binding to the contract
	ContractATransactor // Write-only binding to the contract
	ContractAFilterer   // Log filterer for contract events
}

// ContractACaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractASession struct {
	Contract     *ContractA        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractACallerSession struct {
	Contract *ContractACaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractATransactorSession struct {
	Contract     *ContractATransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractARaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractARaw struct {
	Contract *ContractA // Generic contract binding to access the raw methods on
}

// ContractACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractACallerRaw struct {
	Contract *ContractACaller // Generic read-only contract binding to access the raw methods on
}

// ContractATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractATransactorRaw struct {
	Contract *ContractATransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractA creates a new instance of ContractA, bound to a specific deployed contract.
func NewContractA(address common.Address, backend bind.ContractBackend) (*ContractA, error) {
	contract, err := bindContractA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractA{ContractACaller: ContractACaller{contract: contract}, ContractATransactor: ContractATransactor{contract: contract}, ContractAFilterer: ContractAFilterer{contract: contract}}, nil
}

// NewContractACaller creates a new read-only instance of ContractA, bound to a specific deployed contract.
func NewContractACaller(address common.Address, caller bind.ContractCaller) (*ContractACaller, error) {
	contract, err := bindContractA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractACaller{contract: contract}, nil
}

// NewContractATransactor creates a new write-only instance of ContractA, bound to a specific deployed contract.
func NewContractATransactor(address common.Address, transactor bind.ContractTransactor) (*ContractATransactor, error) {
	contract, err := bindContractA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractATransactor{contract: contract}, nil
}

// NewContractAFilterer creates a new log filterer instance of ContractA, bound to a specific deployed contract.
func NewContractAFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAFilterer, error) {
	contract, err := bindContractA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAFilterer{contract: contract}, nil
}

// bindContractA binds a generic wrapper to an already deployed contract.
func bindContractA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractAMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractA *ContractARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractA.Contract.ContractACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractA *ContractARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractA.Contract.ContractATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractA *ContractARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractA.Contract.ContractATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractA *ContractACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractA *ContractATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractA *ContractATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractA.Contract.contract.Transact(opts, method, params...)
}

// Counterparty is a free data retrieval call binding the contract method 0xc3df6546.
//
// Solidity: function counterparty() view returns(address)
func (_ContractA *ContractACaller) Counterparty(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractA.contract.Call(opts, &out, "counterparty")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterparty is a free data retrieval call binding the contract method 0xc3df6546.
//
// Solidity: function counterparty() view returns(address)
func (_ContractA *ContractASession) Counterparty() (common.Address, error) {
	return _ContractA.Contract.Counterparty(&_ContractA.CallOpts)
}

// Counterparty is a free data retrieval call binding the contract method 0xc3df6546.
//
// Solidity: function counterparty() view returns(address)
func (_ContractA *ContractACallerSession) Counterparty() (common.Address, error) {
	return _ContractA.Contract.Counterparty(&_ContractA.CallOpts)
}

// HandshakeComplete is a free data retrieval call binding the contract method 0x403e3fba.
//
// Solidity: function handshakeComplete(uint256 sessionId) view returns(bool)
func (_ContractA *ContractACaller) HandshakeComplete(opts *bind.CallOpts, sessionId *big.Int) (bool, error) {
	var out []interface{}
	err := _ContractA.contract.Call(opts, &out, "handshakeComplete", sessionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HandshakeComplete is a free data retrieval call binding the contract method 0x403e3fba.
//
// Solidity: function handshakeComplete(uint256 sessionId) view returns(bool)
func (_ContractA *ContractASession) HandshakeComplete(sessionId *big.Int) (bool, error) {
	return _ContractA.Contract.HandshakeComplete(&_ContractA.CallOpts, sessionId)
}

// HandshakeComplete is a free data retrieval call binding the contract method 0x403e3fba.
//
// Solidity: function handshakeComplete(uint256 sessionId) view returns(bool)
func (_ContractA *ContractACallerSession) HandshakeComplete(sessionId *big.Int) (bool, error) {
	return _ContractA.Contract.HandshakeComplete(&_ContractA.CallOpts, sessionId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractA *ContractACaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractA.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractA *ContractASession) Owner() (common.Address, error) {
	return _ContractA.Contract.Owner(&_ContractA.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractA *ContractACallerSession) Owner() (common.Address, error) {
	return _ContractA.Contract.Owner(&_ContractA.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_ContractA *ContractACaller) Relayer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractA.contract.Call(opts, &out, "relayer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_ContractA *ContractASession) Relayer() (common.Address, error) {
	return _ContractA.Contract.Relayer(&_ContractA.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_ContractA *ContractACallerSession) Relayer() (common.Address, error) {
	return _ContractA.Contract.Relayer(&_ContractA.CallOpts)
}

// Sessions is a free data retrieval call binding the contract method 0x83c4b7a3.
//
// Solidity: function sessions(uint256 ) view returns(uint8 status, address initiator)
func (_ContractA *ContractACaller) Sessions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Status    uint8
	Initiator common.Address
}, error) {
	var out []interface{}
	err := _ContractA.contract.Call(opts, &out, "sessions", arg0)

	outstruct := new(struct {
		Status    uint8
		Initiator common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Initiator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Sessions is a free data retrieval call binding the contract method 0x83c4b7a3.
//
// Solidity: function sessions(uint256 ) view returns(uint8 status, address initiator)
func (_ContractA *ContractASession) Sessions(arg0 *big.Int) (struct {
	Status    uint8
	Initiator common.Address
}, error) {
	return _ContractA.Contract.Sessions(&_ContractA.CallOpts, arg0)
}

// Sessions is a free data retrieval call binding the contract method 0x83c4b7a3.
//
// Solidity: function sessions(uint256 ) view returns(uint8 status, address initiator)
func (_ContractA *ContractACallerSession) Sessions(arg0 *big.Int) (struct {
	Status    uint8
	Initiator common.Address
}, error) {
	return _ContractA.Contract.Sessions(&_ContractA.CallOpts, arg0)
}

// Acknowledge is a paid mutator transaction binding the contract method 0x4b5809e1.
//
// Solidity: function acknowledge(uint256 sessionId) returns()
func (_ContractA *ContractATransactor) Acknowledge(opts *bind.TransactOpts, sessionId *big.Int) (*types.Transaction, error) {
	return _ContractA.contract.Transact(opts, "acknowledge", sessionId)
}

// Acknowledge is a paid mutator transaction binding the contract method 0x4b5809e1.
//
// Solidity: function acknowledge(uint256 sessionId) returns()
func (_ContractA *ContractASession) Acknowledge(sessionId *big.Int) (*types.Transaction, error) {
	return _ContractA.Contract.Acknowledge(&_ContractA.TransactOpts, sessionId)
}

// Acknowledge is a paid mutator transaction binding the contract method 0x4b5809e1.
//
// Solidity: function acknowledge(uint256 sessionId) returns()
func (_ContractA *ContractATransactorSession) Acknowledge(sessionId *big.Int) (*types.Transaction, error) {
	return _ContractA.Contract.Acknowledge(&_ContractA.TransactOpts, sessionId)
}

// InitiateHandshake is a paid mutator transaction binding the contract method 0x19f99ad3.
//
// Solidity: function initiateHandshake(uint256 sessionId) returns()
func (_ContractA *ContractATransactor) InitiateHandshake(opts *bind.TransactOpts, sessionId *big.Int) (*types.Transaction, error) {
	return _ContractA.contract.Transact(opts, "initiateHandshake", sessionId)
}

// InitiateHandshake is a paid mutator transaction binding the contract method 0x19f99ad3.
//
// Solidity: function initiateHandshake(uint256 sessionId) returns()
func (_ContractA *ContractASession) InitiateHandshake(sessionId *big.Int) (*types.Transaction, error) {
	return _ContractA.Contract.InitiateHandshake(&_ContractA.TransactOpts, sessionId)
}

// InitiateHandshake is a paid mutator transaction binding the contract method 0x19f99ad3.
//
// Solidity: function initiateHandshake(uint256 sessionId) returns()
func (_ContractA *ContractATransactorSession) InitiateHandshake(sessionId *big.Int) (*types.Transaction, error) {
	return _ContractA.Contract.InitiateHandshake(&_ContractA.TransactOpts, sessionId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractA *ContractATransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractA.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractA *ContractASession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractA.Contract.RenounceOwnership(&_ContractA.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractA *ContractATransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractA.Contract.RenounceOwnership(&_ContractA.TransactOpts)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address newRelayer) returns()
func (_ContractA *ContractATransactor) SetRelayer(opts *bind.TransactOpts, newRelayer common.Address) (*types.Transaction, error) {
	return _ContractA.contract.Transact(opts, "setRelayer", newRelayer)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address newRelayer) returns()
func (_ContractA *ContractASession) SetRelayer(newRelayer common.Address) (*types.Transaction, error) {
	return _ContractA.Contract.SetRelayer(&_ContractA.TransactOpts, newRelayer)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address newRelayer) returns()
func (_ContractA *ContractATransactorSession) SetRelayer(newRelayer common.Address) (*types.Transaction, error) {
	return _ContractA.Contract.SetRelayer(&_ContractA.TransactOpts, newRelayer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractA *ContractATransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractA.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractA *ContractASession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractA.Contract.TransferOwnership(&_ContractA.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractA *ContractATransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractA.Contract.TransferOwnership(&_ContractA.TransactOpts, newOwner)
}

// ContractAOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractA contract.
type ContractAOwnershipTransferredIterator struct {
	Event *ContractAOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractAOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAOwnershipTransferred)
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
		it.Event = new(ContractAOwnershipTransferred)
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
func (it *ContractAOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAOwnershipTransferred represents a OwnershipTransferred event raised by the ContractA contract.
type ContractAOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractA *ContractAFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractAOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractA.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAOwnershipTransferredIterator{contract: _ContractA.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractA *ContractAFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractAOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractA.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAOwnershipTransferred)
				if err := _ContractA.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractA *ContractAFilterer) ParseOwnershipTransferred(log types.Log) (*ContractAOwnershipTransferred, error) {
	event := new(ContractAOwnershipTransferred)
	if err := _ContractA.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractASynIterator is returned from FilterSyn and is used to iterate over the raw logs and unpacked data for Syn events raised by the ContractA contract.
type ContractASynIterator struct {
	Event *ContractASyn // Event containing the contract specifics and raw log

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
func (it *ContractASynIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractASyn)
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
		it.Event = new(ContractASyn)
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
func (it *ContractASynIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractASynIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractASyn represents a Syn event raised by the ContractA contract.
type ContractASyn struct {
	SessionId *big.Int
	Initiator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSyn is a free log retrieval operation binding the contract event 0x4f4922f70c486c95340cec94bcfc794500438c5a6a4033deb6533cd74ca60e8d.
//
// Solidity: event Syn(uint256 indexed sessionId, address indexed initiator)
func (_ContractA *ContractAFilterer) FilterSyn(opts *bind.FilterOpts, sessionId []*big.Int, initiator []common.Address) (*ContractASynIterator, error) {

	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}
	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _ContractA.contract.FilterLogs(opts, "Syn", sessionIdRule, initiatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractASynIterator{contract: _ContractA.contract, event: "Syn", logs: logs, sub: sub}, nil
}

// WatchSyn is a free log subscription operation binding the contract event 0x4f4922f70c486c95340cec94bcfc794500438c5a6a4033deb6533cd74ca60e8d.
//
// Solidity: event Syn(uint256 indexed sessionId, address indexed initiator)
func (_ContractA *ContractAFilterer) WatchSyn(opts *bind.WatchOpts, sink chan<- *ContractASyn, sessionId []*big.Int, initiator []common.Address) (event.Subscription, error) {

	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}
	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _ContractA.contract.WatchLogs(opts, "Syn", sessionIdRule, initiatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractASyn)
				if err := _ContractA.contract.UnpackLog(event, "Syn", log); err != nil {
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

// ParseSyn is a log parse operation binding the contract event 0x4f4922f70c486c95340cec94bcfc794500438c5a6a4033deb6533cd74ca60e8d.
//
// Solidity: event Syn(uint256 indexed sessionId, address indexed initiator)
func (_ContractA *ContractAFilterer) ParseSyn(log types.Log) (*ContractASyn, error) {
	event := new(ContractASyn)
	if err := _ContractA.contract.UnpackLog(event, "Syn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractASynAckIterator is returned from FilterSynAck and is used to iterate over the raw logs and unpacked data for SynAck events raised by the ContractA contract.
type ContractASynAckIterator struct {
	Event *ContractASynAck // Event containing the contract specifics and raw log

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
func (it *ContractASynAckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractASynAck)
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
		it.Event = new(ContractASynAck)
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
func (it *ContractASynAckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractASynAckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractASynAck represents a SynAck event raised by the ContractA contract.
type ContractASynAck struct {
	SessionId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSynAck is a free log retrieval operation binding the contract event 0xec71d58db74e9fc73a8e4e7332027b119c6a36ff6f4fcf16447f25cd30c6b5ca.
//
// Solidity: event SynAck(uint256 indexed sessionId)
func (_ContractA *ContractAFilterer) FilterSynAck(opts *bind.FilterOpts, sessionId []*big.Int) (*ContractASynAckIterator, error) {

	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}

	logs, sub, err := _ContractA.contract.FilterLogs(opts, "SynAck", sessionIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractASynAckIterator{contract: _ContractA.contract, event: "SynAck", logs: logs, sub: sub}, nil
}

// WatchSynAck is a free log subscription operation binding the contract event 0xec71d58db74e9fc73a8e4e7332027b119c6a36ff6f4fcf16447f25cd30c6b5ca.
//
// Solidity: event SynAck(uint256 indexed sessionId)
func (_ContractA *ContractAFilterer) WatchSynAck(opts *bind.WatchOpts, sink chan<- *ContractASynAck, sessionId []*big.Int) (event.Subscription, error) {

	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}

	logs, sub, err := _ContractA.contract.WatchLogs(opts, "SynAck", sessionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractASynAck)
				if err := _ContractA.contract.UnpackLog(event, "SynAck", log); err != nil {
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

// ParseSynAck is a log parse operation binding the contract event 0xec71d58db74e9fc73a8e4e7332027b119c6a36ff6f4fcf16447f25cd30c6b5ca.
//
// Solidity: event SynAck(uint256 indexed sessionId)
func (_ContractA *ContractAFilterer) ParseSynAck(log types.Log) (*ContractASynAck, error) {
	event := new(ContractASynAck)
	if err := _ContractA.contract.UnpackLog(event, "SynAck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package main

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SignerABI is the input ABI used to generate the binding from.
const SignerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAgreement\",\"outputs\":[{\"name\":\"stringToAgreeOn\",\"type\":\"string\"},{\"name\":\"signed\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stringToAgreeOn\",\"type\":\"string\"},{\"name\":\"customer\",\"type\":\"address\"}],\"name\":\"createAgreement\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"signAgreement\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"type\":\"function\"}]"

// SignerBin is the compiled bytecode used for deploying new contracts.
const SignerBin = `0x606060405260008054600160a060020a03191633600160a060020a0316179055341561002a57600080fd5b5b6104868061003a6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b811461005e578063bca60cf51461008d578063bd791a6e14610121578063c939329b14610186575b600080fd5b341561006957600080fd5b6100716101a2565b604051600160a060020a03909116815260200160405180910390f35b341561009857600080fd5b6100a06101b1565b604051811515602082015260408082528190810184818151815260200191508051906020019080838360005b838110156100e55780820151818401525b6020016100cc565b50505050905090810190601f1680156101125780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b61017260046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965050509235600160a060020a031692506102a4915050565b604051901515815260200160405180910390f35b61017261034b565b604051901515815260200160405180910390f35b600054600160a060020a031681565b6101b96103a8565b33600160a060020a03166000908152600160208190526040822080820154909161010090910460ff161515146101ee57600080fd5b6001808201548254839260ff909216918391600260001992821615610100029290920116046020601f820181900481020160405190810160405280929190818152602001828054600181600116156101000203166002900480156102935780601f1061026857610100808354040283529160200191610293565b820191906000526020600020905b81548152906001019060200180831161027657829003601f168201915b50505050509150925092505b509091565b60008054600160a060020a0390811690331681146102c157600080fd5b60606040519081016040908152858252600060208084018290526001838501819052600160a060020a03881683529052208151819080516103069291602001906103ba565b50602082015160018201805460ff19169115159190911790556040820151600191820180549115156101000261ff00199092169190911790559250505b5b5092915050565b33600160a060020a03166000908152600160208190526040822080820154909161010090910460ff1615151461038057600080fd5b600181015460ff161561039257600080fd5b6001818101805460ff19168217905591505b5090565b60206040519081016040526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106103fb57805160ff1916838001178555610428565b82800160010185558215610428579182015b8281111561042857825182559160200191906001019061040d565b5b506103a4929150610439565b5090565b61045791905b808211156103a4576000815560010161043f565b5090565b905600a165627a7a7230582087e4089cb1673a2669b34cf8bd4662258ea9b8c26e97e329bf8d8a90483a631c0029`

// DeploySigner deploys a new Ethereum contract, binding an instance of Signer to it.
func DeploySigner(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Signer, error) {
	parsed, err := abi.JSON(strings.NewReader(SignerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SignerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Signer{SignerCaller: SignerCaller{contract: contract}, SignerTransactor: SignerTransactor{contract: contract}}, nil
}

// Signer is an auto generated Go binding around an Ethereum contract.
type Signer struct {
	SignerCaller     // Read-only binding to the contract
	SignerTransactor // Write-only binding to the contract
}

// SignerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignerSession struct {
	Contract     *Signer           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignerCallerSession struct {
	Contract *SignerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SignerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignerTransactorSession struct {
	Contract     *SignerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignerRaw struct {
	Contract *Signer // Generic contract binding to access the raw methods on
}

// SignerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignerCallerRaw struct {
	Contract *SignerCaller // Generic read-only contract binding to access the raw methods on
}

// SignerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignerTransactorRaw struct {
	Contract *SignerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSigner creates a new instance of Signer, bound to a specific deployed contract.
func NewSigner(address common.Address, backend bind.ContractBackend) (*Signer, error) {
	contract, err := bindSigner(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Signer{SignerCaller: SignerCaller{contract: contract}, SignerTransactor: SignerTransactor{contract: contract}}, nil
}

// NewSignerCaller creates a new read-only instance of Signer, bound to a specific deployed contract.
func NewSignerCaller(address common.Address, caller bind.ContractCaller) (*SignerCaller, error) {
	contract, err := bindSigner(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SignerCaller{contract: contract}, nil
}

// NewSignerTransactor creates a new write-only instance of Signer, bound to a specific deployed contract.
func NewSignerTransactor(address common.Address, transactor bind.ContractTransactor) (*SignerTransactor, error) {
	contract, err := bindSigner(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SignerTransactor{contract: contract}, nil
}

// bindSigner binds a generic wrapper to an already deployed contract.
func bindSigner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SignerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Signer *SignerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Signer.Contract.SignerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Signer *SignerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Signer.Contract.SignerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Signer *SignerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Signer.Contract.SignerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Signer *SignerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Signer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Signer *SignerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Signer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Signer *SignerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Signer.Contract.contract.Transact(opts, method, params...)
}

// GetAgreement is a free data retrieval call binding the contract method 0xbca60cf5.
//
// Solidity: function getAgreement() constant returns(stringToAgreeOn string, signed bool)
func (_Signer *SignerCaller) GetAgreement(opts *bind.CallOpts) (struct {
	StringToAgreeOn string
	Signed          bool
}, error) {
	ret := new(struct {
		StringToAgreeOn string
		Signed          bool
	})
	out := ret
	err := _Signer.contract.Call(opts, out, "getAgreement")
	return *ret, err
}

// GetAgreement is a free data retrieval call binding the contract method 0xbca60cf5.
//
// Solidity: function getAgreement() constant returns(stringToAgreeOn string, signed bool)
func (_Signer *SignerSession) GetAgreement() (struct {
	StringToAgreeOn string
	Signed          bool
}, error) {
	return _Signer.Contract.GetAgreement(&_Signer.CallOpts)
}

// GetAgreement is a free data retrieval call binding the contract method 0xbca60cf5.
//
// Solidity: function getAgreement() constant returns(stringToAgreeOn string, signed bool)
func (_Signer *SignerCallerSession) GetAgreement() (struct {
	StringToAgreeOn string
	Signed          bool
}, error) {
	return _Signer.Contract.GetAgreement(&_Signer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Signer *SignerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Signer.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Signer *SignerSession) Owner() (common.Address, error) {
	return _Signer.Contract.Owner(&_Signer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Signer *SignerCallerSession) Owner() (common.Address, error) {
	return _Signer.Contract.Owner(&_Signer.CallOpts)
}

// CreateAgreement is a paid mutator transaction binding the contract method 0xbd791a6e.
//
// Solidity: function createAgreement(_stringToAgreeOn string, customer address) returns(success bool)
func (_Signer *SignerTransactor) CreateAgreement(opts *bind.TransactOpts, _stringToAgreeOn string, customer common.Address) (*types.Transaction, error) {
	return _Signer.contract.Transact(opts, "createAgreement", _stringToAgreeOn, customer)
}

// CreateAgreement is a paid mutator transaction binding the contract method 0xbd791a6e.
//
// Solidity: function createAgreement(_stringToAgreeOn string, customer address) returns(success bool)
func (_Signer *SignerSession) CreateAgreement(_stringToAgreeOn string, customer common.Address) (*types.Transaction, error) {
	return _Signer.Contract.CreateAgreement(&_Signer.TransactOpts, _stringToAgreeOn, customer)
}

// CreateAgreement is a paid mutator transaction binding the contract method 0xbd791a6e.
//
// Solidity: function createAgreement(_stringToAgreeOn string, customer address) returns(success bool)
func (_Signer *SignerTransactorSession) CreateAgreement(_stringToAgreeOn string, customer common.Address) (*types.Transaction, error) {
	return _Signer.Contract.CreateAgreement(&_Signer.TransactOpts, _stringToAgreeOn, customer)
}

// SignAgreement is a paid mutator transaction binding the contract method 0xc939329b.
//
// Solidity: function signAgreement() returns(success bool)
func (_Signer *SignerTransactor) SignAgreement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Signer.contract.Transact(opts, "signAgreement")
}

// SignAgreement is a paid mutator transaction binding the contract method 0xc939329b.
//
// Solidity: function signAgreement() returns(success bool)
func (_Signer *SignerSession) SignAgreement() (*types.Transaction, error) {
	return _Signer.Contract.SignAgreement(&_Signer.TransactOpts)
}

// SignAgreement is a paid mutator transaction binding the contract method 0xc939329b.
//
// Solidity: function signAgreement() returns(success bool)
func (_Signer *SignerTransactorSession) SignAgreement() (*types.Transaction, error) {
	return _Signer.Contract.SignAgreement(&_Signer.TransactOpts)
}

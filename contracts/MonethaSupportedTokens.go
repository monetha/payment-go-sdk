// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// MonethaSupportedTokensContractABI is the input ABI used to generate the binding from.
const MonethaSupportedTokensContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"tokenId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isMonethaAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"token_acronym\",\"type\":\"bytes32\"},{\"name\":\"token_address\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_isMonethaAddress\",\"type\":\"bool\"}],\"name\":\"setMonethaAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_isMonethaAddress\",\"type\":\"bool\"}],\"name\":\"MonethaAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAcronym\",\"type\":\"bytes32\"},{\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"deleteToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAll\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MonethaSupportedTokensContractBin is the compiled bytecode used for deploying new contracts.
const MonethaSupportedTokensContractBin = `0x608060405260008054600160a060020a03191633179055610845806100256000396000f3006080604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166317d70f7c81146100a857806331d41325146100cf5780634f64b2be1461010457806353ed51431461013d5780636297c16c146101eb578063715018a6146102055780638da5cb5b1461021a5780639776aacf1461024b578063c07e33911461026f578063f2fde38b14610295575b600080fd5b3480156100b457600080fd5b506100bd6102b6565b60408051918252519081900360200190f35b3480156100db57600080fd5b506100f0600160a060020a03600435166102bc565b604080519115158252519081900360200190f35b34801561011057600080fd5b5061011c6004356102d1565b60408051928352600160a060020a0390911660208301528051918290030190f35b34801561014957600080fd5b506101526102f3565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561019657818101518382015260200161017e565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156101d55781810151838201526020016101bd565b5050505090500194505050505060405180910390f35b3480156101f757600080fd5b506102036004356103b1565b005b34801561021157600080fd5b5061020361053e565b34801561022657600080fd5b5061022f6105aa565b60408051600160a060020a039092168252519081900360200190f35b34801561025757600080fd5b50610203600435600160a060020a03602435166105b9565b34801561027b57600080fd5b50610203600160a060020a036004351660243515156106b4565b3480156102a157600080fd5b50610203600160a060020a036004351661072f565b60035481565b60016020526000908152604090205460ff1681565b60026020526000908152604090208054600190910154600160a060020a031682565b606080600460058180548060200260200160405190810160405280929190818152602001828054801561034f57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610331575b50505050509150808054806020026020016040519081016040528092919081815260200182805480156103a257602002820191906000526020600020905b8154815260019091019060200180831161038d575b50505050509050915091509091565b3360009081526001602052604081205460ff1615156103cf57600080fd5b5060038054600090815260026020526040808220600190810154858452828420918201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909216919091179055925482528120549083905290556004805490600019820182811061043c57fe5b60009182526020909120015460048054600160a060020a0390921691600019850190811061046657fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790556005805460001983019081106104ae57fe5b90600052602060002001546005600184038154811015156104cb57fe5b60009182526020909120015560048054906104ea9060001983016107cf565b5060058054906104fe9060001983016107cf565b5050600380546000908152600260205260408120908155600101805473ffffffffffffffffffffffffffffffffffffffff19169055805460001901905550565b600054600160a060020a0316331461055557600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b3360009081526001602052604090205460ff1615156105d757600080fd5b600160a060020a03811615156105ec57600080fd5b604080518082018252838152600160a060020a0392831660208083018281526003805460019081019182905560009182526002909352948520935184555192810180549390951673ffffffffffffffffffffffffffffffffffffffff1993841617909455600480548086019091557f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b01805490921617905560058054928301815590527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00155565b600054600160a060020a031633146106cb57600080fd5b600160a060020a038216600081815260016020908152604091829020805460ff191685151590811790915582519384529083015280517fa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e6929281900390910190a15050565b600054600160a060020a0316331461074657600080fd5b61074f81610752565b50565b600160a060020a038116151561076757600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b8154818355818111156107f3576000838152602090206107f39181019083016107f8565b505050565b61081691905b8082111561081257600081556001016107fe565b5090565b905600a165627a7a72305820ca567d0cea47817acc0043dfce64088c7ff121ca541599e0a065739c19775dc40029`

// DeployMonethaSupportedTokensContract deploys a new Ethereum contract, binding an instance of MonethaSupportedTokensContract to it.
func DeployMonethaSupportedTokensContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MonethaSupportedTokensContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MonethaSupportedTokensContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MonethaSupportedTokensContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MonethaSupportedTokensContract{MonethaSupportedTokensContractCaller: MonethaSupportedTokensContractCaller{contract: contract}, MonethaSupportedTokensContractTransactor: MonethaSupportedTokensContractTransactor{contract: contract}, MonethaSupportedTokensContractFilterer: MonethaSupportedTokensContractFilterer{contract: contract}}, nil
}

// MonethaSupportedTokensContract is an auto generated Go binding around an Ethereum contract.
type MonethaSupportedTokensContract struct {
	MonethaSupportedTokensContractCaller     // Read-only binding to the contract
	MonethaSupportedTokensContractTransactor // Write-only binding to the contract
	MonethaSupportedTokensContractFilterer   // Log filterer for contract events
}

// MonethaSupportedTokensContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MonethaSupportedTokensContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonethaSupportedTokensContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MonethaSupportedTokensContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonethaSupportedTokensContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MonethaSupportedTokensContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonethaSupportedTokensContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MonethaSupportedTokensContractSession struct {
	Contract     *MonethaSupportedTokensContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// MonethaSupportedTokensContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MonethaSupportedTokensContractCallerSession struct {
	Contract *MonethaSupportedTokensContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// MonethaSupportedTokensContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MonethaSupportedTokensContractTransactorSession struct {
	Contract     *MonethaSupportedTokensContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// MonethaSupportedTokensContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MonethaSupportedTokensContractRaw struct {
	Contract *MonethaSupportedTokensContract // Generic contract binding to access the raw methods on
}

// MonethaSupportedTokensContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MonethaSupportedTokensContractCallerRaw struct {
	Contract *MonethaSupportedTokensContractCaller // Generic read-only contract binding to access the raw methods on
}

// MonethaSupportedTokensContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MonethaSupportedTokensContractTransactorRaw struct {
	Contract *MonethaSupportedTokensContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMonethaSupportedTokensContract creates a new instance of MonethaSupportedTokensContract, bound to a specific deployed contract.
func NewMonethaSupportedTokensContract(address common.Address, backend bind.ContractBackend) (*MonethaSupportedTokensContract, error) {
	contract, err := bindMonethaSupportedTokensContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MonethaSupportedTokensContract{MonethaSupportedTokensContractCaller: MonethaSupportedTokensContractCaller{contract: contract}, MonethaSupportedTokensContractTransactor: MonethaSupportedTokensContractTransactor{contract: contract}, MonethaSupportedTokensContractFilterer: MonethaSupportedTokensContractFilterer{contract: contract}}, nil
}

// NewMonethaSupportedTokensContractCaller creates a new read-only instance of MonethaSupportedTokensContract, bound to a specific deployed contract.
func NewMonethaSupportedTokensContractCaller(address common.Address, caller bind.ContractCaller) (*MonethaSupportedTokensContractCaller, error) {
	contract, err := bindMonethaSupportedTokensContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MonethaSupportedTokensContractCaller{contract: contract}, nil
}

// NewMonethaSupportedTokensContractTransactor creates a new write-only instance of MonethaSupportedTokensContract, bound to a specific deployed contract.
func NewMonethaSupportedTokensContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MonethaSupportedTokensContractTransactor, error) {
	contract, err := bindMonethaSupportedTokensContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MonethaSupportedTokensContractTransactor{contract: contract}, nil
}

// NewMonethaSupportedTokensContractFilterer creates a new log filterer instance of MonethaSupportedTokensContract, bound to a specific deployed contract.
func NewMonethaSupportedTokensContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MonethaSupportedTokensContractFilterer, error) {
	contract, err := bindMonethaSupportedTokensContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MonethaSupportedTokensContractFilterer{contract: contract}, nil
}

// bindMonethaSupportedTokensContract binds a generic wrapper to an already deployed contract.
func bindMonethaSupportedTokensContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MonethaSupportedTokensContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MonethaSupportedTokensContract.Contract.MonethaSupportedTokensContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaSupportedTokensContract.Contract.MonethaSupportedTokensContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonethaSupportedTokensContract.Contract.MonethaSupportedTokensContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MonethaSupportedTokensContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaSupportedTokensContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonethaSupportedTokensContract.Contract.contract.Transact(opts, method, params...)
}

// GetAll is a free data retrieval call binding the contract method 0x53ed5143.
//
// Solidity: function getAll() constant returns(address[], bytes32[])
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractCaller) GetAll(opts *bind.CallOpts) ([]common.Address, [][32]byte, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([][32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _MonethaSupportedTokensContract.contract.Call(opts, out, "getAll")
	return *ret0, *ret1, err
}

// GetAll is a free data retrieval call binding the contract method 0x53ed5143.
//
// Solidity: function getAll() constant returns(address[], bytes32[])
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractSession) GetAll() ([]common.Address, [][32]byte, error) {
	return _MonethaSupportedTokensContract.Contract.GetAll(&_MonethaSupportedTokensContract.CallOpts)
}

// GetAll is a free data retrieval call binding the contract method 0x53ed5143.
//
// Solidity: function getAll() constant returns(address[], bytes32[])
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractCallerSession) GetAll() ([]common.Address, [][32]byte, error) {
	return _MonethaSupportedTokensContract.Contract.GetAll(&_MonethaSupportedTokensContract.CallOpts)
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractCaller) IsMonethaAddress(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MonethaSupportedTokensContract.contract.Call(opts, out, "isMonethaAddress", arg0)
	return *ret0, err
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractSession) IsMonethaAddress(arg0 common.Address) (bool, error) {
	return _MonethaSupportedTokensContract.Contract.IsMonethaAddress(&_MonethaSupportedTokensContract.CallOpts, arg0)
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractCallerSession) IsMonethaAddress(arg0 common.Address) (bool, error) {
	return _MonethaSupportedTokensContract.Contract.IsMonethaAddress(&_MonethaSupportedTokensContract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonethaSupportedTokensContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractSession) Owner() (common.Address, error) {
	return _MonethaSupportedTokensContract.Contract.Owner(&_MonethaSupportedTokensContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractCallerSession) Owner() (common.Address, error) {
	return _MonethaSupportedTokensContract.Contract.Owner(&_MonethaSupportedTokensContract.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() constant returns(uint256)
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractCaller) TokenId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaSupportedTokensContract.contract.Call(opts, out, "tokenId")
	return *ret0, err
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() constant returns(uint256)
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractSession) TokenId() (*big.Int, error) {
	return _MonethaSupportedTokensContract.Contract.TokenId(&_MonethaSupportedTokensContract.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() constant returns(uint256)
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractCallerSession) TokenId() (*big.Int, error) {
	return _MonethaSupportedTokensContract.Contract.TokenId(&_MonethaSupportedTokensContract.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens( uint256) constant returns(token_acronym bytes32, token_address address)
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenAcronym [32]byte
	TokenAddress common.Address
}, error) {
	ret := new(struct {
		TokenAcronym [32]byte
		TokenAddress common.Address
	})
	out := ret
	err := _MonethaSupportedTokensContract.contract.Call(opts, out, "tokens", arg0)
	return *ret, err
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens( uint256) constant returns(token_acronym bytes32, token_address address)
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractSession) Tokens(arg0 *big.Int) (struct {
	TokenAcronym [32]byte
	TokenAddress common.Address
}, error) {
	return _MonethaSupportedTokensContract.Contract.Tokens(&_MonethaSupportedTokensContract.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens( uint256) constant returns(token_acronym bytes32, token_address address)
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractCallerSession) Tokens(arg0 *big.Int) (struct {
	TokenAcronym [32]byte
	TokenAddress common.Address
}, error) {
	return _MonethaSupportedTokensContract.Contract.Tokens(&_MonethaSupportedTokensContract.CallOpts, arg0)
}

// AddToken is a paid mutator transaction binding the contract method 0x9776aacf.
//
// Solidity: function addToken(_tokenAcronym bytes32, _tokenAddress address) returns()
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractTransactor) AddToken(opts *bind.TransactOpts, _tokenAcronym [32]byte, _tokenAddress common.Address) (*types.Transaction, error) {
	return _MonethaSupportedTokensContract.contract.Transact(opts, "addToken", _tokenAcronym, _tokenAddress)
}

// AddToken is a paid mutator transaction binding the contract method 0x9776aacf.
//
// Solidity: function addToken(_tokenAcronym bytes32, _tokenAddress address) returns()
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractSession) AddToken(_tokenAcronym [32]byte, _tokenAddress common.Address) (*types.Transaction, error) {
	return _MonethaSupportedTokensContract.Contract.AddToken(&_MonethaSupportedTokensContract.TransactOpts, _tokenAcronym, _tokenAddress)
}

// AddToken is a paid mutator transaction binding the contract method 0x9776aacf.
//
// Solidity: function addToken(_tokenAcronym bytes32, _tokenAddress address) returns()
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractTransactorSession) AddToken(_tokenAcronym [32]byte, _tokenAddress common.Address) (*types.Transaction, error) {
	return _MonethaSupportedTokensContract.Contract.AddToken(&_MonethaSupportedTokensContract.TransactOpts, _tokenAcronym, _tokenAddress)
}

// DeleteToken is a paid mutator transaction binding the contract method 0x6297c16c.
//
// Solidity: function deleteToken(_tokenId uint256) returns()
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractTransactor) DeleteToken(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _MonethaSupportedTokensContract.contract.Transact(opts, "deleteToken", _tokenId)
}

// DeleteToken is a paid mutator transaction binding the contract method 0x6297c16c.
//
// Solidity: function deleteToken(_tokenId uint256) returns()
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractSession) DeleteToken(_tokenId *big.Int) (*types.Transaction, error) {
	return _MonethaSupportedTokensContract.Contract.DeleteToken(&_MonethaSupportedTokensContract.TransactOpts, _tokenId)
}

// DeleteToken is a paid mutator transaction binding the contract method 0x6297c16c.
//
// Solidity: function deleteToken(_tokenId uint256) returns()
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractTransactorSession) DeleteToken(_tokenId *big.Int) (*types.Transaction, error) {
	return _MonethaSupportedTokensContract.Contract.DeleteToken(&_MonethaSupportedTokensContract.TransactOpts, _tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaSupportedTokensContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _MonethaSupportedTokensContract.Contract.RenounceOwnership(&_MonethaSupportedTokensContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MonethaSupportedTokensContract.Contract.RenounceOwnership(&_MonethaSupportedTokensContract.TransactOpts)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractTransactor) SetMonethaAddress(opts *bind.TransactOpts, _address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _MonethaSupportedTokensContract.contract.Transact(opts, "setMonethaAddress", _address, _isMonethaAddress)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractSession) SetMonethaAddress(_address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _MonethaSupportedTokensContract.Contract.SetMonethaAddress(&_MonethaSupportedTokensContract.TransactOpts, _address, _isMonethaAddress)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractTransactorSession) SetMonethaAddress(_address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _MonethaSupportedTokensContract.Contract.SetMonethaAddress(&_MonethaSupportedTokensContract.TransactOpts, _address, _isMonethaAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _MonethaSupportedTokensContract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _MonethaSupportedTokensContract.Contract.TransferOwnership(&_MonethaSupportedTokensContract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _MonethaSupportedTokensContract.Contract.TransferOwnership(&_MonethaSupportedTokensContract.TransactOpts, _newOwner)
}

// MonethaSupportedTokensContractMonethaAddressSetIterator is returned from FilterMonethaAddressSet and is used to iterate over the raw logs and unpacked data for MonethaAddressSet events raised by the MonethaSupportedTokensContract contract.
type MonethaSupportedTokensContractMonethaAddressSetIterator struct {
	Event *MonethaSupportedTokensContractMonethaAddressSet // Event containing the contract specifics and raw log

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
func (it *MonethaSupportedTokensContractMonethaAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaSupportedTokensContractMonethaAddressSet)
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
		it.Event = new(MonethaSupportedTokensContractMonethaAddressSet)
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
func (it *MonethaSupportedTokensContractMonethaAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaSupportedTokensContractMonethaAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaSupportedTokensContractMonethaAddressSet represents a MonethaAddressSet event raised by the MonethaSupportedTokensContract contract.
type MonethaSupportedTokensContractMonethaAddressSet struct {
	Address          common.Address
	IsMonethaAddress bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMonethaAddressSet is a free log retrieval operation binding the contract event 0xa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e692.
//
// Solidity: e MonethaAddressSet(_address address, _isMonethaAddress bool)
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractFilterer) FilterMonethaAddressSet(opts *bind.FilterOpts) (*MonethaSupportedTokensContractMonethaAddressSetIterator, error) {

	logs, sub, err := _MonethaSupportedTokensContract.contract.FilterLogs(opts, "MonethaAddressSet")
	if err != nil {
		return nil, err
	}
	return &MonethaSupportedTokensContractMonethaAddressSetIterator{contract: _MonethaSupportedTokensContract.contract, event: "MonethaAddressSet", logs: logs, sub: sub}, nil
}

// WatchMonethaAddressSet is a free log subscription operation binding the contract event 0xa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e692.
//
// Solidity: e MonethaAddressSet(_address address, _isMonethaAddress bool)
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractFilterer) WatchMonethaAddressSet(opts *bind.WatchOpts, sink chan<- *MonethaSupportedTokensContractMonethaAddressSet) (event.Subscription, error) {

	logs, sub, err := _MonethaSupportedTokensContract.contract.WatchLogs(opts, "MonethaAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaSupportedTokensContractMonethaAddressSet)
				if err := _MonethaSupportedTokensContract.contract.UnpackLog(event, "MonethaAddressSet", log); err != nil {
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

// MonethaSupportedTokensContractOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the MonethaSupportedTokensContract contract.
type MonethaSupportedTokensContractOwnershipRenouncedIterator struct {
	Event *MonethaSupportedTokensContractOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *MonethaSupportedTokensContractOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaSupportedTokensContractOwnershipRenounced)
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
		it.Event = new(MonethaSupportedTokensContractOwnershipRenounced)
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
func (it *MonethaSupportedTokensContractOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaSupportedTokensContractOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaSupportedTokensContractOwnershipRenounced represents a OwnershipRenounced event raised by the MonethaSupportedTokensContract contract.
type MonethaSupportedTokensContractOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*MonethaSupportedTokensContractOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _MonethaSupportedTokensContract.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MonethaSupportedTokensContractOwnershipRenouncedIterator{contract: _MonethaSupportedTokensContract.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *MonethaSupportedTokensContractOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _MonethaSupportedTokensContract.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaSupportedTokensContractOwnershipRenounced)
				if err := _MonethaSupportedTokensContract.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// MonethaSupportedTokensContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MonethaSupportedTokensContract contract.
type MonethaSupportedTokensContractOwnershipTransferredIterator struct {
	Event *MonethaSupportedTokensContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MonethaSupportedTokensContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaSupportedTokensContractOwnershipTransferred)
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
		it.Event = new(MonethaSupportedTokensContractOwnershipTransferred)
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
func (it *MonethaSupportedTokensContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaSupportedTokensContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaSupportedTokensContractOwnershipTransferred represents a OwnershipTransferred event raised by the MonethaSupportedTokensContract contract.
type MonethaSupportedTokensContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MonethaSupportedTokensContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MonethaSupportedTokensContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MonethaSupportedTokensContractOwnershipTransferredIterator{contract: _MonethaSupportedTokensContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_MonethaSupportedTokensContract *MonethaSupportedTokensContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MonethaSupportedTokensContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MonethaSupportedTokensContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaSupportedTokensContractOwnershipTransferred)
				if err := _MonethaSupportedTokensContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

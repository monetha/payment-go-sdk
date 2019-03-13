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

// MonethaTokenContractABI is the input ABI used to generate the binding from.
const MonethaTokenContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"standard\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ico\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lockedAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokensForIco\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lockReleaseDate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_icoAddress\",\"type\":\"address\"}],\"name\":\"setICO\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newStart\",\"type\":\"uint256\"}],\"name\":\"setStart\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reservedAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_ownerAddr\",\"type\":\"address\"},{\"name\":\"_startTime\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"}]"

// MonethaTokenContractBin is the compiled bytecode used for deploying new contracts.
const MonethaTokenContractBin = `0x60606040526524991ae7e000600055341561001957600080fd5b604051604080610c5283398101604052808051919060200180519150505b60018054600160a060020a031916600160a060020a03848116919091179182905560038390556301e13380830160045560008054929091168152600660205260409020555b50505b610bc48061008e6000396000f3006060604052361561010f5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde038114610114578063095ea7b31461019f57806318160ddd146101d557806323b872dd146101fa578063313ce5671461023657806344df8e701461025f5780635a3b7e42146102745780635d452201146102ff5780636ab28bc81461032e57806370a082311461035357806378e979251461038457806382ea97b3146103a95780638da5cb5b146103ce57806395d89b41146103fd578063a9059cbb14610488578063ac4abae1146104be578063b6f50c29146104e3578063dd62ed3e14610504578063f6a03ebf1461053b578063f92c45b7146103a9575b600080fd5b341561011f57600080fd5b610127610578565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156101645780820151818401525b60200161014b565b50505050905090810190601f1680156101915780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101aa57600080fd5b6101c1600160a060020a03600435166024356105af565b604051901515815260200160405180910390f35b34156101e057600080fd5b6101e86105c4565b60405190815260200160405180910390f35b341561020557600080fd5b6101c1600160a060020a03600435811690602435166044356105ca565b604051901515815260200160405180910390f35b341561024157600080fd5b61024961074e565b60405160ff909116815260200160405180910390f35b341561026a57600080fd5b610272610753565b005b341561027f57600080fd5b610127610813565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156101645780820151818401525b60200161014b565b50505050905090810190601f1680156101915780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561030a57600080fd5b61031261084a565b604051600160a060020a03909116815260200160405180910390f35b341561033957600080fd5b6101e8610859565b60405190815260200160405180910390f35b341561035e57600080fd5b6101e8600160a060020a0360043516610863565b60405190815260200160405180910390f35b341561038f57600080fd5b6101e8610875565b60405190815260200160405180910390f35b34156103b457600080fd5b6101e861087b565b60405190815260200160405180910390f35b34156103d957600080fd5b610312610885565b604051600160a060020a03909116815260200160405180910390f35b341561040857600080fd5b610127610894565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156101645780820151818401525b60200161014b565b50505050905090810190601f1680156101915780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561049357600080fd5b6101c1600160a060020a03600435166024356108cb565b604051901515815260200160405180910390f35b34156104c957600080fd5b6101e86109ec565b60405190815260200160405180910390f35b34156104ee57600080fd5b610272600160a060020a03600435166109f2565b005b341561050f57600080fd5b6101e8600160a060020a0360043581169060243516610a58565b60405190815260200160405180910390f35b341561054657600080fd5b610272600435610a75565b005b34156103b457600080fd5b6101e861087b565b60405190815260200160405180910390f35b60408051908101604052600781527f4d6f6e6574686100000000000000000000000000000000000000000000000000602082015281565b60006105bb8383610ab2565b90505b92915050565b60005481565b6000806003544210156105f157600154600160a060020a038681169116146105f157600080fd5b5b600154600160a060020a038681169116148015610610575060045442105b1561064b57600160a060020a038516600090815260066020526040902054650de8428b5000906106409085610b59565b101561064b57600080fd5b5b50600160a060020a0380851660008181526007602090815260408083203390951683529381528382205492825260069052919091205461068c9084610b59565b600160a060020a0380871660009081526006602052604080822093909355908616815220546106bb9084610b70565b600160a060020a0385166000908152600660205260409020556106de8184610b59565b600160a060020a03808716600081815260076020908152604080832033861684529091529081902093909355908616917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9086905190815260200160405180910390a3600191505b509392505050565b600581565b60055460009060ff1615801561076a575060035442115b1561080f57600154600160a060020a031660009081526006602052604090205461079a9065124c8d73f000610b59565b600154600160a060020a0316600090815260066020526040812065124c8d73f0009055549091506107cb9082610b59565b6000556005805460ff191660011790557fd83c63197e8e676d80ab0122beba9a9d20f3828839e9a1d6fe81d242e9cd7e6e8160405190815260200160405180910390a15b5b50565b60408051908101604052600581527f4552433230000000000000000000000000000000000000000000000000000000602082015281565b600254600160a060020a031681565b650de8428b500081565b60066020526000908152604090205481565b60035481565b65124c8d73f00081565b600154600160a060020a031681565b60408051908101604052600381527f4d54480000000000000000000000000000000000000000000000000000000000602082015281565b6003546000904210156108dd57600080fd5b60015433600160a060020a0390811691161480156108fc575060045442105b1561093757600160a060020a033316600090815260066020526040902054650de8428b50009061092c9084610b59565b101561093757600080fd5b5b600160a060020a03331660009081526006602052604090205461095b9083610b59565b600160a060020a03338116600090815260066020526040808220939093559085168152205461098a9083610b70565b600160a060020a0380851660008181526006602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060015b92915050565b60045481565b60015433600160a060020a03908116911614610a0d57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055610a4c911665124c8d73f000610ab2565b151561080f57fe5b5b50565b600760209081526000928352604080842090915290825290205481565b60025433600160a060020a039081169116148015610a94575060035481105b1515610a9f57600080fd5b60038190555b50565b65124c8d73f00081565b6000811580610ae45750600160a060020a03338116600090815260076020908152604080832093871683529290522054155b1515610aef57600080fd5b600160a060020a03338116600081815260076020908152604080832094881680845294909152908190208590557f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a35060015b92915050565b600082821115610b6557fe5b508082035b92915050565b6000828201838110801590610b855750828110155b1515610b8d57fe5b8091505b50929150505600a165627a7a723058201ece16a4849d6f063ac16d4723492df94b137bb9f49ec010ea71aa0475f8a1f80029`

// DeployMonethaTokenContract deploys a new Ethereum contract, binding an instance of MonethaTokenContract to it.
func DeployMonethaTokenContract(auth *bind.TransactOpts, backend bind.ContractBackend, _ownerAddr common.Address, _startTime *big.Int) (common.Address, *types.Transaction, *MonethaTokenContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MonethaTokenContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MonethaTokenContractBin), backend, _ownerAddr, _startTime)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MonethaTokenContract{MonethaTokenContractCaller: MonethaTokenContractCaller{contract: contract}, MonethaTokenContractTransactor: MonethaTokenContractTransactor{contract: contract}, MonethaTokenContractFilterer: MonethaTokenContractFilterer{contract: contract}}, nil
}

// MonethaTokenContract is an auto generated Go binding around an Ethereum contract.
type MonethaTokenContract struct {
	MonethaTokenContractCaller     // Read-only binding to the contract
	MonethaTokenContractTransactor // Write-only binding to the contract
	MonethaTokenContractFilterer   // Log filterer for contract events
}

// MonethaTokenContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MonethaTokenContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonethaTokenContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MonethaTokenContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonethaTokenContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MonethaTokenContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonethaTokenContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MonethaTokenContractSession struct {
	Contract     *MonethaTokenContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MonethaTokenContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MonethaTokenContractCallerSession struct {
	Contract *MonethaTokenContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// MonethaTokenContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MonethaTokenContractTransactorSession struct {
	Contract     *MonethaTokenContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// MonethaTokenContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MonethaTokenContractRaw struct {
	Contract *MonethaTokenContract // Generic contract binding to access the raw methods on
}

// MonethaTokenContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MonethaTokenContractCallerRaw struct {
	Contract *MonethaTokenContractCaller // Generic read-only contract binding to access the raw methods on
}

// MonethaTokenContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MonethaTokenContractTransactorRaw struct {
	Contract *MonethaTokenContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMonethaTokenContract creates a new instance of MonethaTokenContract, bound to a specific deployed contract.
func NewMonethaTokenContract(address common.Address, backend bind.ContractBackend) (*MonethaTokenContract, error) {
	contract, err := bindMonethaTokenContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MonethaTokenContract{MonethaTokenContractCaller: MonethaTokenContractCaller{contract: contract}, MonethaTokenContractTransactor: MonethaTokenContractTransactor{contract: contract}, MonethaTokenContractFilterer: MonethaTokenContractFilterer{contract: contract}}, nil
}

// NewMonethaTokenContractCaller creates a new read-only instance of MonethaTokenContract, bound to a specific deployed contract.
func NewMonethaTokenContractCaller(address common.Address, caller bind.ContractCaller) (*MonethaTokenContractCaller, error) {
	contract, err := bindMonethaTokenContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MonethaTokenContractCaller{contract: contract}, nil
}

// NewMonethaTokenContractTransactor creates a new write-only instance of MonethaTokenContract, bound to a specific deployed contract.
func NewMonethaTokenContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MonethaTokenContractTransactor, error) {
	contract, err := bindMonethaTokenContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MonethaTokenContractTransactor{contract: contract}, nil
}

// NewMonethaTokenContractFilterer creates a new log filterer instance of MonethaTokenContract, bound to a specific deployed contract.
func NewMonethaTokenContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MonethaTokenContractFilterer, error) {
	contract, err := bindMonethaTokenContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MonethaTokenContractFilterer{contract: contract}, nil
}

// bindMonethaTokenContract binds a generic wrapper to an already deployed contract.
func bindMonethaTokenContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MonethaTokenContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonethaTokenContract *MonethaTokenContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MonethaTokenContract.Contract.MonethaTokenContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonethaTokenContract *MonethaTokenContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaTokenContract.Contract.MonethaTokenContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonethaTokenContract *MonethaTokenContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonethaTokenContract.Contract.MonethaTokenContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonethaTokenContract *MonethaTokenContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MonethaTokenContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonethaTokenContract *MonethaTokenContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaTokenContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonethaTokenContract *MonethaTokenContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonethaTokenContract.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaTokenContract.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MonethaTokenContract.Contract.Allowance(&_MonethaTokenContract.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MonethaTokenContract.Contract.Allowance(&_MonethaTokenContract.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaTokenContract.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _MonethaTokenContract.Contract.BalanceOf(&_MonethaTokenContract.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _MonethaTokenContract.Contract.BalanceOf(&_MonethaTokenContract.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MonethaTokenContract *MonethaTokenContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MonethaTokenContract.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MonethaTokenContract *MonethaTokenContractSession) Decimals() (uint8, error) {
	return _MonethaTokenContract.Contract.Decimals(&_MonethaTokenContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MonethaTokenContract *MonethaTokenContractCallerSession) Decimals() (uint8, error) {
	return _MonethaTokenContract.Contract.Decimals(&_MonethaTokenContract.CallOpts)
}

// Ico is a free data retrieval call binding the contract method 0x5d452201.
//
// Solidity: function ico() constant returns(address)
func (_MonethaTokenContract *MonethaTokenContractCaller) Ico(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonethaTokenContract.contract.Call(opts, out, "ico")
	return *ret0, err
}

// Ico is a free data retrieval call binding the contract method 0x5d452201.
//
// Solidity: function ico() constant returns(address)
func (_MonethaTokenContract *MonethaTokenContractSession) Ico() (common.Address, error) {
	return _MonethaTokenContract.Contract.Ico(&_MonethaTokenContract.CallOpts)
}

// Ico is a free data retrieval call binding the contract method 0x5d452201.
//
// Solidity: function ico() constant returns(address)
func (_MonethaTokenContract *MonethaTokenContractCallerSession) Ico() (common.Address, error) {
	return _MonethaTokenContract.Contract.Ico(&_MonethaTokenContract.CallOpts)
}

// LockReleaseDate is a free data retrieval call binding the contract method 0xac4abae1.
//
// Solidity: function lockReleaseDate() constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractCaller) LockReleaseDate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaTokenContract.contract.Call(opts, out, "lockReleaseDate")
	return *ret0, err
}

// LockReleaseDate is a free data retrieval call binding the contract method 0xac4abae1.
//
// Solidity: function lockReleaseDate() constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractSession) LockReleaseDate() (*big.Int, error) {
	return _MonethaTokenContract.Contract.LockReleaseDate(&_MonethaTokenContract.CallOpts)
}

// LockReleaseDate is a free data retrieval call binding the contract method 0xac4abae1.
//
// Solidity: function lockReleaseDate() constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractCallerSession) LockReleaseDate() (*big.Int, error) {
	return _MonethaTokenContract.Contract.LockReleaseDate(&_MonethaTokenContract.CallOpts)
}

// LockedAmount is a free data retrieval call binding the contract method 0x6ab28bc8.
//
// Solidity: function lockedAmount() constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractCaller) LockedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaTokenContract.contract.Call(opts, out, "lockedAmount")
	return *ret0, err
}

// LockedAmount is a free data retrieval call binding the contract method 0x6ab28bc8.
//
// Solidity: function lockedAmount() constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractSession) LockedAmount() (*big.Int, error) {
	return _MonethaTokenContract.Contract.LockedAmount(&_MonethaTokenContract.CallOpts)
}

// LockedAmount is a free data retrieval call binding the contract method 0x6ab28bc8.
//
// Solidity: function lockedAmount() constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractCallerSession) LockedAmount() (*big.Int, error) {
	return _MonethaTokenContract.Contract.LockedAmount(&_MonethaTokenContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MonethaTokenContract *MonethaTokenContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MonethaTokenContract.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MonethaTokenContract *MonethaTokenContractSession) Name() (string, error) {
	return _MonethaTokenContract.Contract.Name(&_MonethaTokenContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MonethaTokenContract *MonethaTokenContractCallerSession) Name() (string, error) {
	return _MonethaTokenContract.Contract.Name(&_MonethaTokenContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MonethaTokenContract *MonethaTokenContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonethaTokenContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MonethaTokenContract *MonethaTokenContractSession) Owner() (common.Address, error) {
	return _MonethaTokenContract.Contract.Owner(&_MonethaTokenContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MonethaTokenContract *MonethaTokenContractCallerSession) Owner() (common.Address, error) {
	return _MonethaTokenContract.Contract.Owner(&_MonethaTokenContract.CallOpts)
}

// ReservedAmount is a free data retrieval call binding the contract method 0xf92c45b7.
//
// Solidity: function reservedAmount() constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractCaller) ReservedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaTokenContract.contract.Call(opts, out, "reservedAmount")
	return *ret0, err
}

// ReservedAmount is a free data retrieval call binding the contract method 0xf92c45b7.
//
// Solidity: function reservedAmount() constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractSession) ReservedAmount() (*big.Int, error) {
	return _MonethaTokenContract.Contract.ReservedAmount(&_MonethaTokenContract.CallOpts)
}

// ReservedAmount is a free data retrieval call binding the contract method 0xf92c45b7.
//
// Solidity: function reservedAmount() constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractCallerSession) ReservedAmount() (*big.Int, error) {
	return _MonethaTokenContract.Contract.ReservedAmount(&_MonethaTokenContract.CallOpts)
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() constant returns(string)
func (_MonethaTokenContract *MonethaTokenContractCaller) Standard(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MonethaTokenContract.contract.Call(opts, out, "standard")
	return *ret0, err
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() constant returns(string)
func (_MonethaTokenContract *MonethaTokenContractSession) Standard() (string, error) {
	return _MonethaTokenContract.Contract.Standard(&_MonethaTokenContract.CallOpts)
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() constant returns(string)
func (_MonethaTokenContract *MonethaTokenContractCallerSession) Standard() (string, error) {
	return _MonethaTokenContract.Contract.Standard(&_MonethaTokenContract.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaTokenContract.contract.Call(opts, out, "startTime")
	return *ret0, err
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractSession) StartTime() (*big.Int, error) {
	return _MonethaTokenContract.Contract.StartTime(&_MonethaTokenContract.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractCallerSession) StartTime() (*big.Int, error) {
	return _MonethaTokenContract.Contract.StartTime(&_MonethaTokenContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MonethaTokenContract *MonethaTokenContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MonethaTokenContract.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MonethaTokenContract *MonethaTokenContractSession) Symbol() (string, error) {
	return _MonethaTokenContract.Contract.Symbol(&_MonethaTokenContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MonethaTokenContract *MonethaTokenContractCallerSession) Symbol() (string, error) {
	return _MonethaTokenContract.Contract.Symbol(&_MonethaTokenContract.CallOpts)
}

// TokensForIco is a free data retrieval call binding the contract method 0x82ea97b3.
//
// Solidity: function tokensForIco() constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractCaller) TokensForIco(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaTokenContract.contract.Call(opts, out, "tokensForIco")
	return *ret0, err
}

// TokensForIco is a free data retrieval call binding the contract method 0x82ea97b3.
//
// Solidity: function tokensForIco() constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractSession) TokensForIco() (*big.Int, error) {
	return _MonethaTokenContract.Contract.TokensForIco(&_MonethaTokenContract.CallOpts)
}

// TokensForIco is a free data retrieval call binding the contract method 0x82ea97b3.
//
// Solidity: function tokensForIco() constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractCallerSession) TokensForIco() (*big.Int, error) {
	return _MonethaTokenContract.Contract.TokensForIco(&_MonethaTokenContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaTokenContract.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractSession) TotalSupply() (*big.Int, error) {
	return _MonethaTokenContract.Contract.TotalSupply(&_MonethaTokenContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MonethaTokenContract *MonethaTokenContractCallerSession) TotalSupply() (*big.Int, error) {
	return _MonethaTokenContract.Contract.TotalSupply(&_MonethaTokenContract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_MonethaTokenContract *MonethaTokenContractTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaTokenContract.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_MonethaTokenContract *MonethaTokenContractSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaTokenContract.Contract.Approve(&_MonethaTokenContract.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_MonethaTokenContract *MonethaTokenContractTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaTokenContract.Contract.Approve(&_MonethaTokenContract.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns()
func (_MonethaTokenContract *MonethaTokenContractTransactor) Burn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaTokenContract.contract.Transact(opts, "burn")
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns()
func (_MonethaTokenContract *MonethaTokenContractSession) Burn() (*types.Transaction, error) {
	return _MonethaTokenContract.Contract.Burn(&_MonethaTokenContract.TransactOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns()
func (_MonethaTokenContract *MonethaTokenContractTransactorSession) Burn() (*types.Transaction, error) {
	return _MonethaTokenContract.Contract.Burn(&_MonethaTokenContract.TransactOpts)
}

// SetICO is a paid mutator transaction binding the contract method 0xb6f50c29.
//
// Solidity: function setICO(_icoAddress address) returns()
func (_MonethaTokenContract *MonethaTokenContractTransactor) SetICO(opts *bind.TransactOpts, _icoAddress common.Address) (*types.Transaction, error) {
	return _MonethaTokenContract.contract.Transact(opts, "setICO", _icoAddress)
}

// SetICO is a paid mutator transaction binding the contract method 0xb6f50c29.
//
// Solidity: function setICO(_icoAddress address) returns()
func (_MonethaTokenContract *MonethaTokenContractSession) SetICO(_icoAddress common.Address) (*types.Transaction, error) {
	return _MonethaTokenContract.Contract.SetICO(&_MonethaTokenContract.TransactOpts, _icoAddress)
}

// SetICO is a paid mutator transaction binding the contract method 0xb6f50c29.
//
// Solidity: function setICO(_icoAddress address) returns()
func (_MonethaTokenContract *MonethaTokenContractTransactorSession) SetICO(_icoAddress common.Address) (*types.Transaction, error) {
	return _MonethaTokenContract.Contract.SetICO(&_MonethaTokenContract.TransactOpts, _icoAddress)
}

// SetStart is a paid mutator transaction binding the contract method 0xf6a03ebf.
//
// Solidity: function setStart(_newStart uint256) returns()
func (_MonethaTokenContract *MonethaTokenContractTransactor) SetStart(opts *bind.TransactOpts, _newStart *big.Int) (*types.Transaction, error) {
	return _MonethaTokenContract.contract.Transact(opts, "setStart", _newStart)
}

// SetStart is a paid mutator transaction binding the contract method 0xf6a03ebf.
//
// Solidity: function setStart(_newStart uint256) returns()
func (_MonethaTokenContract *MonethaTokenContractSession) SetStart(_newStart *big.Int) (*types.Transaction, error) {
	return _MonethaTokenContract.Contract.SetStart(&_MonethaTokenContract.TransactOpts, _newStart)
}

// SetStart is a paid mutator transaction binding the contract method 0xf6a03ebf.
//
// Solidity: function setStart(_newStart uint256) returns()
func (_MonethaTokenContract *MonethaTokenContractTransactorSession) SetStart(_newStart *big.Int) (*types.Transaction, error) {
	return _MonethaTokenContract.Contract.SetStart(&_MonethaTokenContract.TransactOpts, _newStart)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_MonethaTokenContract *MonethaTokenContractTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaTokenContract.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_MonethaTokenContract *MonethaTokenContractSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaTokenContract.Contract.Transfer(&_MonethaTokenContract.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_MonethaTokenContract *MonethaTokenContractTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaTokenContract.Contract.Transfer(&_MonethaTokenContract.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_MonethaTokenContract *MonethaTokenContractTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaTokenContract.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_MonethaTokenContract *MonethaTokenContractSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaTokenContract.Contract.TransferFrom(&_MonethaTokenContract.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_MonethaTokenContract *MonethaTokenContractTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaTokenContract.Contract.TransferFrom(&_MonethaTokenContract.TransactOpts, _from, _to, _value)
}

// MonethaTokenContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MonethaTokenContract contract.
type MonethaTokenContractApprovalIterator struct {
	Event *MonethaTokenContractApproval // Event containing the contract specifics and raw log

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
func (it *MonethaTokenContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaTokenContractApproval)
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
		it.Event = new(MonethaTokenContractApproval)
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
func (it *MonethaTokenContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaTokenContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaTokenContractApproval represents a Approval event raised by the MonethaTokenContract contract.
type MonethaTokenContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, spender indexed address, value uint256)
func (_MonethaTokenContract *MonethaTokenContractFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, spender []common.Address) (*MonethaTokenContractApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MonethaTokenContract.contract.FilterLogs(opts, "Approval", _ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MonethaTokenContractApprovalIterator{contract: _MonethaTokenContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, spender indexed address, value uint256)
func (_MonethaTokenContract *MonethaTokenContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MonethaTokenContractApproval, _owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MonethaTokenContract.contract.WatchLogs(opts, "Approval", _ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaTokenContractApproval)
				if err := _MonethaTokenContract.contract.UnpackLog(event, "Approval", log); err != nil {
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

// MonethaTokenContractBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the MonethaTokenContract contract.
type MonethaTokenContractBurnedIterator struct {
	Event *MonethaTokenContractBurned // Event containing the contract specifics and raw log

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
func (it *MonethaTokenContractBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaTokenContractBurned)
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
		it.Event = new(MonethaTokenContractBurned)
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
func (it *MonethaTokenContractBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaTokenContractBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaTokenContractBurned represents a Burned event raised by the MonethaTokenContract contract.
type MonethaTokenContractBurned struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0xd83c63197e8e676d80ab0122beba9a9d20f3828839e9a1d6fe81d242e9cd7e6e.
//
// Solidity: e Burned(amount uint256)
func (_MonethaTokenContract *MonethaTokenContractFilterer) FilterBurned(opts *bind.FilterOpts) (*MonethaTokenContractBurnedIterator, error) {

	logs, sub, err := _MonethaTokenContract.contract.FilterLogs(opts, "Burned")
	if err != nil {
		return nil, err
	}
	return &MonethaTokenContractBurnedIterator{contract: _MonethaTokenContract.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0xd83c63197e8e676d80ab0122beba9a9d20f3828839e9a1d6fe81d242e9cd7e6e.
//
// Solidity: e Burned(amount uint256)
func (_MonethaTokenContract *MonethaTokenContractFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *MonethaTokenContractBurned) (event.Subscription, error) {

	logs, sub, err := _MonethaTokenContract.contract.WatchLogs(opts, "Burned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaTokenContractBurned)
				if err := _MonethaTokenContract.contract.UnpackLog(event, "Burned", log); err != nil {
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

// MonethaTokenContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MonethaTokenContract contract.
type MonethaTokenContractTransferIterator struct {
	Event *MonethaTokenContractTransfer // Event containing the contract specifics and raw log

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
func (it *MonethaTokenContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaTokenContractTransfer)
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
		it.Event = new(MonethaTokenContractTransfer)
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
func (it *MonethaTokenContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaTokenContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaTokenContractTransfer represents a Transfer event raised by the MonethaTokenContract contract.
type MonethaTokenContractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_MonethaTokenContract *MonethaTokenContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MonethaTokenContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MonethaTokenContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MonethaTokenContractTransferIterator{contract: _MonethaTokenContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_MonethaTokenContract *MonethaTokenContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MonethaTokenContractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MonethaTokenContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaTokenContractTransfer)
				if err := _MonethaTokenContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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

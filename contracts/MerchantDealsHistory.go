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

// MerchantDealsHistoryContractABI is the input ABI used to generate the binding from.
const MerchantDealsHistoryContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isMonethaAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contactInformation\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_info\",\"type\":\"string\"}],\"name\":\"setContactInformation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_isMonethaAddress\",\"type\":\"bool\"}],\"name\":\"setMonethaAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"merchantIdHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_merchantId\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"clientAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"clientReputation\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"merchantReputation\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"successful\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"dealHash\",\"type\":\"uint256\"}],\"name\":\"DealCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"clientAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"clientReputation\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"merchantReputation\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"dealHash\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"cancelReason\",\"type\":\"string\"}],\"name\":\"DealCancelationReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"clientAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"clientReputation\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"merchantReputation\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"dealHash\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"refundReason\",\"type\":\"string\"}],\"name\":\"DealRefundReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_isMonethaAddress\",\"type\":\"bool\"}],\"name\":\"MonethaAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\"},{\"name\":\"_clientAddress\",\"type\":\"address\"},{\"name\":\"_clientReputation\",\"type\":\"uint32\"},{\"name\":\"_merchantReputation\",\"type\":\"uint32\"},{\"name\":\"_isSuccess\",\"type\":\"bool\"},{\"name\":\"_dealHash\",\"type\":\"uint256\"}],\"name\":\"recordDeal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\"},{\"name\":\"_clientAddress\",\"type\":\"address\"},{\"name\":\"_clientReputation\",\"type\":\"uint32\"},{\"name\":\"_merchantReputation\",\"type\":\"uint32\"},{\"name\":\"_dealHash\",\"type\":\"uint256\"},{\"name\":\"_cancelReason\",\"type\":\"string\"}],\"name\":\"recordDealCancelReason\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\"},{\"name\":\"_clientAddress\",\"type\":\"address\"},{\"name\":\"_clientReputation\",\"type\":\"uint32\"},{\"name\":\"_merchantReputation\",\"type\":\"uint32\"},{\"name\":\"_dealHash\",\"type\":\"uint256\"},{\"name\":\"_refundReason\",\"type\":\"string\"}],\"name\":\"recordDealRefundReason\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MerchantDealsHistoryContractBin is the compiled bytecode used for deploying new contracts.
const MerchantDealsHistoryContractBin = `0x608060405234801561001057600080fd5b506040516109ba3803806109ba83398101604052805160008054600160a060020a031916331781559101805190911061004857600080fd5b806040516020018082805190602001908083835b6020831061007b5780518252601f19909201916020918201910161005c565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b602083106100de5780518252601f1990920191602091820191016100bf565b5181516000196020949094036101000a9390930192831692191691909117905260405192018290039091206003555050505061089b8061011f6000396000f3006080604052600436106100ae5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663188c668c81146100b357806331d41325146100fd57806336f7ab5e146101325780635d818e6b146101bc578063715018a6146101f75780638da5cb5b1461020c578063a0e5e8211461023d578063b967a52e14610285578063c07e3391146102de578063f0daba0114610304578063f2fde38b1461032b575b600080fd5b3480156100bf57600080fd5b506100fb600480359060248035600160a060020a03169160443563ffffffff90811692606435909116916084359160a43590810191013561034c565b005b34801561010957600080fd5b5061011e600160a060020a0360043516610415565b604080519115158252519081900360200190f35b34801561013e57600080fd5b5061014761042a565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610181578181015183820152602001610169565b50505050905090810190601f1680156101ae5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101c857600080fd5b506100fb600435600160a060020a036024351663ffffffff60443581169060643516608435151560a4356104b7565b34801561020357600080fd5b506100fb610541565b34801561021857600080fd5b506102216105ad565b60408051600160a060020a039092168252519081900360200190f35b34801561024957600080fd5b506100fb600480359060248035600160a060020a03169160443563ffffffff90811692606435909116916084359160a4359081019101356105bc565b34801561029157600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100fb9436949293602493928401919081908401838280828437509497506106859650505050505050565b3480156102ea57600080fd5b506100fb600160a060020a036004351660243515156106b3565b34801561031057600080fd5b5061031961072e565b60408051918252519081900360200190f35b34801561033757600080fd5b506100fb600160a060020a0360043516610734565b3360009081526002602052604090205460ff16151561036a57600080fd5b7f917df6bc35eda1274dcb5342bd54cab99299c55e524d65ae584b721e5567f3ef878787878787876040518088815260200187600160a060020a0316600160a060020a031681526020018663ffffffff1663ffffffff1681526020018563ffffffff1663ffffffff16815260200184815260200180602001828103825284848281815260200192508082843760405192018290039a509098505050505050505050a150505050505050565b60026020526000908152604090205460ff1681565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104af5780601f10610484576101008083540402835291602001916104af565b820191906000526020600020905b81548152906001019060200180831161049257829003601f168201915b505050505081565b3360009081526002602052604090205460ff1615156104d557600080fd5b60408051878152600160a060020a038716602082015263ffffffff8087168284015285166060820152831515608082015260a0810183905290517f864fff0aa0ac03e3e06afa833d4d03ab86079fcbc1c78d2b6bbfadaa69c1cf249181900360c00190a1505050505050565b600054600160a060020a0316331461055857600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b3360009081526002602052604090205460ff1615156105da57600080fd5b7f843296d6cc322bb9877b7eb21ae9064f88bcdc32ee2a2f491e80ca247ad0de7c878787878787876040518088815260200187600160a060020a0316600160a060020a031681526020018663ffffffff1663ffffffff1681526020018563ffffffff1663ffffffff16815260200184815260200180602001828103825284848281815260200192508082843760405192018290039a509098505050505050505050a150505050505050565b600054600160a060020a0316331461069c57600080fd5b80516106af9060019060208401906107d4565b5050565b600054600160a060020a031633146106ca57600080fd5b600160a060020a038216600081815260026020908152604091829020805460ff191685151590811790915582519384529083015280517fa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e6929281900390910190a15050565b60035481565b600054600160a060020a0316331461074b57600080fd5b61075481610757565b50565b600160a060020a038116151561076c57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061081557805160ff1916838001178555610842565b82800160010185558215610842579182015b82811115610842578251825591602001919060010190610827565b5061084e929150610852565b5090565b61086c91905b8082111561084e5760008155600101610858565b905600a165627a7a72305820ac98c13514c47169341a4fc82fbca344338da342c57330d8c68005aebd9db2dc0029`

// DeployMerchantDealsHistoryContract deploys a new Ethereum contract, binding an instance of MerchantDealsHistoryContract to it.
func DeployMerchantDealsHistoryContract(auth *bind.TransactOpts, backend bind.ContractBackend, _merchantId string) (common.Address, *types.Transaction, *MerchantDealsHistoryContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerchantDealsHistoryContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MerchantDealsHistoryContractBin), backend, _merchantId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerchantDealsHistoryContract{MerchantDealsHistoryContractCaller: MerchantDealsHistoryContractCaller{contract: contract}, MerchantDealsHistoryContractTransactor: MerchantDealsHistoryContractTransactor{contract: contract}, MerchantDealsHistoryContractFilterer: MerchantDealsHistoryContractFilterer{contract: contract}}, nil
}

// MerchantDealsHistoryContract is an auto generated Go binding around an Ethereum contract.
type MerchantDealsHistoryContract struct {
	MerchantDealsHistoryContractCaller     // Read-only binding to the contract
	MerchantDealsHistoryContractTransactor // Write-only binding to the contract
	MerchantDealsHistoryContractFilterer   // Log filterer for contract events
}

// MerchantDealsHistoryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerchantDealsHistoryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerchantDealsHistoryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerchantDealsHistoryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerchantDealsHistoryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerchantDealsHistoryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerchantDealsHistoryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerchantDealsHistoryContractSession struct {
	Contract     *MerchantDealsHistoryContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MerchantDealsHistoryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerchantDealsHistoryContractCallerSession struct {
	Contract *MerchantDealsHistoryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// MerchantDealsHistoryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerchantDealsHistoryContractTransactorSession struct {
	Contract     *MerchantDealsHistoryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// MerchantDealsHistoryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerchantDealsHistoryContractRaw struct {
	Contract *MerchantDealsHistoryContract // Generic contract binding to access the raw methods on
}

// MerchantDealsHistoryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerchantDealsHistoryContractCallerRaw struct {
	Contract *MerchantDealsHistoryContractCaller // Generic read-only contract binding to access the raw methods on
}

// MerchantDealsHistoryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerchantDealsHistoryContractTransactorRaw struct {
	Contract *MerchantDealsHistoryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerchantDealsHistoryContract creates a new instance of MerchantDealsHistoryContract, bound to a specific deployed contract.
func NewMerchantDealsHistoryContract(address common.Address, backend bind.ContractBackend) (*MerchantDealsHistoryContract, error) {
	contract, err := bindMerchantDealsHistoryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerchantDealsHistoryContract{MerchantDealsHistoryContractCaller: MerchantDealsHistoryContractCaller{contract: contract}, MerchantDealsHistoryContractTransactor: MerchantDealsHistoryContractTransactor{contract: contract}, MerchantDealsHistoryContractFilterer: MerchantDealsHistoryContractFilterer{contract: contract}}, nil
}

// NewMerchantDealsHistoryContractCaller creates a new read-only instance of MerchantDealsHistoryContract, bound to a specific deployed contract.
func NewMerchantDealsHistoryContractCaller(address common.Address, caller bind.ContractCaller) (*MerchantDealsHistoryContractCaller, error) {
	contract, err := bindMerchantDealsHistoryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerchantDealsHistoryContractCaller{contract: contract}, nil
}

// NewMerchantDealsHistoryContractTransactor creates a new write-only instance of MerchantDealsHistoryContract, bound to a specific deployed contract.
func NewMerchantDealsHistoryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MerchantDealsHistoryContractTransactor, error) {
	contract, err := bindMerchantDealsHistoryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerchantDealsHistoryContractTransactor{contract: contract}, nil
}

// NewMerchantDealsHistoryContractFilterer creates a new log filterer instance of MerchantDealsHistoryContract, bound to a specific deployed contract.
func NewMerchantDealsHistoryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MerchantDealsHistoryContractFilterer, error) {
	contract, err := bindMerchantDealsHistoryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerchantDealsHistoryContractFilterer{contract: contract}, nil
}

// bindMerchantDealsHistoryContract binds a generic wrapper to an already deployed contract.
func bindMerchantDealsHistoryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerchantDealsHistoryContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerchantDealsHistoryContract.Contract.MerchantDealsHistoryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.Contract.MerchantDealsHistoryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.Contract.MerchantDealsHistoryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerchantDealsHistoryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.Contract.contract.Transact(opts, method, params...)
}

// ContactInformation is a free data retrieval call binding the contract method 0x36f7ab5e.
//
// Solidity: function contactInformation() constant returns(string)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractCaller) ContactInformation(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MerchantDealsHistoryContract.contract.Call(opts, out, "contactInformation")
	return *ret0, err
}

// ContactInformation is a free data retrieval call binding the contract method 0x36f7ab5e.
//
// Solidity: function contactInformation() constant returns(string)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractSession) ContactInformation() (string, error) {
	return _MerchantDealsHistoryContract.Contract.ContactInformation(&_MerchantDealsHistoryContract.CallOpts)
}

// ContactInformation is a free data retrieval call binding the contract method 0x36f7ab5e.
//
// Solidity: function contactInformation() constant returns(string)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractCallerSession) ContactInformation() (string, error) {
	return _MerchantDealsHistoryContract.Contract.ContactInformation(&_MerchantDealsHistoryContract.CallOpts)
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractCaller) IsMonethaAddress(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MerchantDealsHistoryContract.contract.Call(opts, out, "isMonethaAddress", arg0)
	return *ret0, err
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractSession) IsMonethaAddress(arg0 common.Address) (bool, error) {
	return _MerchantDealsHistoryContract.Contract.IsMonethaAddress(&_MerchantDealsHistoryContract.CallOpts, arg0)
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractCallerSession) IsMonethaAddress(arg0 common.Address) (bool, error) {
	return _MerchantDealsHistoryContract.Contract.IsMonethaAddress(&_MerchantDealsHistoryContract.CallOpts, arg0)
}

// MerchantIdHash is a free data retrieval call binding the contract method 0xf0daba01.
//
// Solidity: function merchantIdHash() constant returns(bytes32)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractCaller) MerchantIdHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MerchantDealsHistoryContract.contract.Call(opts, out, "merchantIdHash")
	return *ret0, err
}

// MerchantIdHash is a free data retrieval call binding the contract method 0xf0daba01.
//
// Solidity: function merchantIdHash() constant returns(bytes32)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractSession) MerchantIdHash() ([32]byte, error) {
	return _MerchantDealsHistoryContract.Contract.MerchantIdHash(&_MerchantDealsHistoryContract.CallOpts)
}

// MerchantIdHash is a free data retrieval call binding the contract method 0xf0daba01.
//
// Solidity: function merchantIdHash() constant returns(bytes32)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractCallerSession) MerchantIdHash() ([32]byte, error) {
	return _MerchantDealsHistoryContract.Contract.MerchantIdHash(&_MerchantDealsHistoryContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MerchantDealsHistoryContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractSession) Owner() (common.Address, error) {
	return _MerchantDealsHistoryContract.Contract.Owner(&_MerchantDealsHistoryContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractCallerSession) Owner() (common.Address, error) {
	return _MerchantDealsHistoryContract.Contract.Owner(&_MerchantDealsHistoryContract.CallOpts)
}

// RecordDeal is a paid mutator transaction binding the contract method 0x5d818e6b.
//
// Solidity: function recordDeal(_orderId uint256, _clientAddress address, _clientReputation uint32, _merchantReputation uint32, _isSuccess bool, _dealHash uint256) returns()
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractTransactor) RecordDeal(opts *bind.TransactOpts, _orderId *big.Int, _clientAddress common.Address, _clientReputation uint32, _merchantReputation uint32, _isSuccess bool, _dealHash *big.Int) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.contract.Transact(opts, "recordDeal", _orderId, _clientAddress, _clientReputation, _merchantReputation, _isSuccess, _dealHash)
}

// RecordDeal is a paid mutator transaction binding the contract method 0x5d818e6b.
//
// Solidity: function recordDeal(_orderId uint256, _clientAddress address, _clientReputation uint32, _merchantReputation uint32, _isSuccess bool, _dealHash uint256) returns()
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractSession) RecordDeal(_orderId *big.Int, _clientAddress common.Address, _clientReputation uint32, _merchantReputation uint32, _isSuccess bool, _dealHash *big.Int) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.Contract.RecordDeal(&_MerchantDealsHistoryContract.TransactOpts, _orderId, _clientAddress, _clientReputation, _merchantReputation, _isSuccess, _dealHash)
}

// RecordDeal is a paid mutator transaction binding the contract method 0x5d818e6b.
//
// Solidity: function recordDeal(_orderId uint256, _clientAddress address, _clientReputation uint32, _merchantReputation uint32, _isSuccess bool, _dealHash uint256) returns()
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractTransactorSession) RecordDeal(_orderId *big.Int, _clientAddress common.Address, _clientReputation uint32, _merchantReputation uint32, _isSuccess bool, _dealHash *big.Int) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.Contract.RecordDeal(&_MerchantDealsHistoryContract.TransactOpts, _orderId, _clientAddress, _clientReputation, _merchantReputation, _isSuccess, _dealHash)
}

// RecordDealCancelReason is a paid mutator transaction binding the contract method 0xa0e5e821.
//
// Solidity: function recordDealCancelReason(_orderId uint256, _clientAddress address, _clientReputation uint32, _merchantReputation uint32, _dealHash uint256, _cancelReason string) returns()
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractTransactor) RecordDealCancelReason(opts *bind.TransactOpts, _orderId *big.Int, _clientAddress common.Address, _clientReputation uint32, _merchantReputation uint32, _dealHash *big.Int, _cancelReason string) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.contract.Transact(opts, "recordDealCancelReason", _orderId, _clientAddress, _clientReputation, _merchantReputation, _dealHash, _cancelReason)
}

// RecordDealCancelReason is a paid mutator transaction binding the contract method 0xa0e5e821.
//
// Solidity: function recordDealCancelReason(_orderId uint256, _clientAddress address, _clientReputation uint32, _merchantReputation uint32, _dealHash uint256, _cancelReason string) returns()
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractSession) RecordDealCancelReason(_orderId *big.Int, _clientAddress common.Address, _clientReputation uint32, _merchantReputation uint32, _dealHash *big.Int, _cancelReason string) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.Contract.RecordDealCancelReason(&_MerchantDealsHistoryContract.TransactOpts, _orderId, _clientAddress, _clientReputation, _merchantReputation, _dealHash, _cancelReason)
}

// RecordDealCancelReason is a paid mutator transaction binding the contract method 0xa0e5e821.
//
// Solidity: function recordDealCancelReason(_orderId uint256, _clientAddress address, _clientReputation uint32, _merchantReputation uint32, _dealHash uint256, _cancelReason string) returns()
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractTransactorSession) RecordDealCancelReason(_orderId *big.Int, _clientAddress common.Address, _clientReputation uint32, _merchantReputation uint32, _dealHash *big.Int, _cancelReason string) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.Contract.RecordDealCancelReason(&_MerchantDealsHistoryContract.TransactOpts, _orderId, _clientAddress, _clientReputation, _merchantReputation, _dealHash, _cancelReason)
}

// RecordDealRefundReason is a paid mutator transaction binding the contract method 0x188c668c.
//
// Solidity: function recordDealRefundReason(_orderId uint256, _clientAddress address, _clientReputation uint32, _merchantReputation uint32, _dealHash uint256, _refundReason string) returns()
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractTransactor) RecordDealRefundReason(opts *bind.TransactOpts, _orderId *big.Int, _clientAddress common.Address, _clientReputation uint32, _merchantReputation uint32, _dealHash *big.Int, _refundReason string) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.contract.Transact(opts, "recordDealRefundReason", _orderId, _clientAddress, _clientReputation, _merchantReputation, _dealHash, _refundReason)
}

// RecordDealRefundReason is a paid mutator transaction binding the contract method 0x188c668c.
//
// Solidity: function recordDealRefundReason(_orderId uint256, _clientAddress address, _clientReputation uint32, _merchantReputation uint32, _dealHash uint256, _refundReason string) returns()
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractSession) RecordDealRefundReason(_orderId *big.Int, _clientAddress common.Address, _clientReputation uint32, _merchantReputation uint32, _dealHash *big.Int, _refundReason string) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.Contract.RecordDealRefundReason(&_MerchantDealsHistoryContract.TransactOpts, _orderId, _clientAddress, _clientReputation, _merchantReputation, _dealHash, _refundReason)
}

// RecordDealRefundReason is a paid mutator transaction binding the contract method 0x188c668c.
//
// Solidity: function recordDealRefundReason(_orderId uint256, _clientAddress address, _clientReputation uint32, _merchantReputation uint32, _dealHash uint256, _refundReason string) returns()
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractTransactorSession) RecordDealRefundReason(_orderId *big.Int, _clientAddress common.Address, _clientReputation uint32, _merchantReputation uint32, _dealHash *big.Int, _refundReason string) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.Contract.RecordDealRefundReason(&_MerchantDealsHistoryContract.TransactOpts, _orderId, _clientAddress, _clientReputation, _merchantReputation, _dealHash, _refundReason)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.Contract.RenounceOwnership(&_MerchantDealsHistoryContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.Contract.RenounceOwnership(&_MerchantDealsHistoryContract.TransactOpts)
}

// SetContactInformation is a paid mutator transaction binding the contract method 0xb967a52e.
//
// Solidity: function setContactInformation(_info string) returns()
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractTransactor) SetContactInformation(opts *bind.TransactOpts, _info string) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.contract.Transact(opts, "setContactInformation", _info)
}

// SetContactInformation is a paid mutator transaction binding the contract method 0xb967a52e.
//
// Solidity: function setContactInformation(_info string) returns()
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractSession) SetContactInformation(_info string) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.Contract.SetContactInformation(&_MerchantDealsHistoryContract.TransactOpts, _info)
}

// SetContactInformation is a paid mutator transaction binding the contract method 0xb967a52e.
//
// Solidity: function setContactInformation(_info string) returns()
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractTransactorSession) SetContactInformation(_info string) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.Contract.SetContactInformation(&_MerchantDealsHistoryContract.TransactOpts, _info)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractTransactor) SetMonethaAddress(opts *bind.TransactOpts, _address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.contract.Transact(opts, "setMonethaAddress", _address, _isMonethaAddress)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractSession) SetMonethaAddress(_address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.Contract.SetMonethaAddress(&_MerchantDealsHistoryContract.TransactOpts, _address, _isMonethaAddress)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractTransactorSession) SetMonethaAddress(_address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.Contract.SetMonethaAddress(&_MerchantDealsHistoryContract.TransactOpts, _address, _isMonethaAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.Contract.TransferOwnership(&_MerchantDealsHistoryContract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _MerchantDealsHistoryContract.Contract.TransferOwnership(&_MerchantDealsHistoryContract.TransactOpts, _newOwner)
}

// MerchantDealsHistoryContractDealCancelationReasonIterator is returned from FilterDealCancelationReason and is used to iterate over the raw logs and unpacked data for DealCancelationReason events raised by the MerchantDealsHistoryContract contract.
type MerchantDealsHistoryContractDealCancelationReasonIterator struct {
	Event *MerchantDealsHistoryContractDealCancelationReason // Event containing the contract specifics and raw log

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
func (it *MerchantDealsHistoryContractDealCancelationReasonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantDealsHistoryContractDealCancelationReason)
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
		it.Event = new(MerchantDealsHistoryContractDealCancelationReason)
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
func (it *MerchantDealsHistoryContractDealCancelationReasonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantDealsHistoryContractDealCancelationReasonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantDealsHistoryContractDealCancelationReason represents a DealCancelationReason event raised by the MerchantDealsHistoryContract contract.
type MerchantDealsHistoryContractDealCancelationReason struct {
	OrderId            *big.Int
	ClientAddress      common.Address
	ClientReputation   uint32
	MerchantReputation uint32
	DealHash           *big.Int
	CancelReason       string
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDealCancelationReason is a free log retrieval operation binding the contract event 0x843296d6cc322bb9877b7eb21ae9064f88bcdc32ee2a2f491e80ca247ad0de7c.
//
// Solidity: e DealCancelationReason(orderId uint256, clientAddress address, clientReputation uint32, merchantReputation uint32, dealHash uint256, cancelReason string)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractFilterer) FilterDealCancelationReason(opts *bind.FilterOpts) (*MerchantDealsHistoryContractDealCancelationReasonIterator, error) {

	logs, sub, err := _MerchantDealsHistoryContract.contract.FilterLogs(opts, "DealCancelationReason")
	if err != nil {
		return nil, err
	}
	return &MerchantDealsHistoryContractDealCancelationReasonIterator{contract: _MerchantDealsHistoryContract.contract, event: "DealCancelationReason", logs: logs, sub: sub}, nil
}

// WatchDealCancelationReason is a free log subscription operation binding the contract event 0x843296d6cc322bb9877b7eb21ae9064f88bcdc32ee2a2f491e80ca247ad0de7c.
//
// Solidity: e DealCancelationReason(orderId uint256, clientAddress address, clientReputation uint32, merchantReputation uint32, dealHash uint256, cancelReason string)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractFilterer) WatchDealCancelationReason(opts *bind.WatchOpts, sink chan<- *MerchantDealsHistoryContractDealCancelationReason) (event.Subscription, error) {

	logs, sub, err := _MerchantDealsHistoryContract.contract.WatchLogs(opts, "DealCancelationReason")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantDealsHistoryContractDealCancelationReason)
				if err := _MerchantDealsHistoryContract.contract.UnpackLog(event, "DealCancelationReason", log); err != nil {
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

// MerchantDealsHistoryContractDealCompletedIterator is returned from FilterDealCompleted and is used to iterate over the raw logs and unpacked data for DealCompleted events raised by the MerchantDealsHistoryContract contract.
type MerchantDealsHistoryContractDealCompletedIterator struct {
	Event *MerchantDealsHistoryContractDealCompleted // Event containing the contract specifics and raw log

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
func (it *MerchantDealsHistoryContractDealCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantDealsHistoryContractDealCompleted)
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
		it.Event = new(MerchantDealsHistoryContractDealCompleted)
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
func (it *MerchantDealsHistoryContractDealCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantDealsHistoryContractDealCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantDealsHistoryContractDealCompleted represents a DealCompleted event raised by the MerchantDealsHistoryContract contract.
type MerchantDealsHistoryContractDealCompleted struct {
	OrderId            *big.Int
	ClientAddress      common.Address
	ClientReputation   uint32
	MerchantReputation uint32
	Successful         bool
	DealHash           *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDealCompleted is a free log retrieval operation binding the contract event 0x864fff0aa0ac03e3e06afa833d4d03ab86079fcbc1c78d2b6bbfadaa69c1cf24.
//
// Solidity: e DealCompleted(orderId uint256, clientAddress address, clientReputation uint32, merchantReputation uint32, successful bool, dealHash uint256)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractFilterer) FilterDealCompleted(opts *bind.FilterOpts) (*MerchantDealsHistoryContractDealCompletedIterator, error) {

	logs, sub, err := _MerchantDealsHistoryContract.contract.FilterLogs(opts, "DealCompleted")
	if err != nil {
		return nil, err
	}
	return &MerchantDealsHistoryContractDealCompletedIterator{contract: _MerchantDealsHistoryContract.contract, event: "DealCompleted", logs: logs, sub: sub}, nil
}

// WatchDealCompleted is a free log subscription operation binding the contract event 0x864fff0aa0ac03e3e06afa833d4d03ab86079fcbc1c78d2b6bbfadaa69c1cf24.
//
// Solidity: e DealCompleted(orderId uint256, clientAddress address, clientReputation uint32, merchantReputation uint32, successful bool, dealHash uint256)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractFilterer) WatchDealCompleted(opts *bind.WatchOpts, sink chan<- *MerchantDealsHistoryContractDealCompleted) (event.Subscription, error) {

	logs, sub, err := _MerchantDealsHistoryContract.contract.WatchLogs(opts, "DealCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantDealsHistoryContractDealCompleted)
				if err := _MerchantDealsHistoryContract.contract.UnpackLog(event, "DealCompleted", log); err != nil {
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

// MerchantDealsHistoryContractDealRefundReasonIterator is returned from FilterDealRefundReason and is used to iterate over the raw logs and unpacked data for DealRefundReason events raised by the MerchantDealsHistoryContract contract.
type MerchantDealsHistoryContractDealRefundReasonIterator struct {
	Event *MerchantDealsHistoryContractDealRefundReason // Event containing the contract specifics and raw log

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
func (it *MerchantDealsHistoryContractDealRefundReasonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantDealsHistoryContractDealRefundReason)
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
		it.Event = new(MerchantDealsHistoryContractDealRefundReason)
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
func (it *MerchantDealsHistoryContractDealRefundReasonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantDealsHistoryContractDealRefundReasonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantDealsHistoryContractDealRefundReason represents a DealRefundReason event raised by the MerchantDealsHistoryContract contract.
type MerchantDealsHistoryContractDealRefundReason struct {
	OrderId            *big.Int
	ClientAddress      common.Address
	ClientReputation   uint32
	MerchantReputation uint32
	DealHash           *big.Int
	RefundReason       string
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDealRefundReason is a free log retrieval operation binding the contract event 0x917df6bc35eda1274dcb5342bd54cab99299c55e524d65ae584b721e5567f3ef.
//
// Solidity: e DealRefundReason(orderId uint256, clientAddress address, clientReputation uint32, merchantReputation uint32, dealHash uint256, refundReason string)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractFilterer) FilterDealRefundReason(opts *bind.FilterOpts) (*MerchantDealsHistoryContractDealRefundReasonIterator, error) {

	logs, sub, err := _MerchantDealsHistoryContract.contract.FilterLogs(opts, "DealRefundReason")
	if err != nil {
		return nil, err
	}
	return &MerchantDealsHistoryContractDealRefundReasonIterator{contract: _MerchantDealsHistoryContract.contract, event: "DealRefundReason", logs: logs, sub: sub}, nil
}

// WatchDealRefundReason is a free log subscription operation binding the contract event 0x917df6bc35eda1274dcb5342bd54cab99299c55e524d65ae584b721e5567f3ef.
//
// Solidity: e DealRefundReason(orderId uint256, clientAddress address, clientReputation uint32, merchantReputation uint32, dealHash uint256, refundReason string)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractFilterer) WatchDealRefundReason(opts *bind.WatchOpts, sink chan<- *MerchantDealsHistoryContractDealRefundReason) (event.Subscription, error) {

	logs, sub, err := _MerchantDealsHistoryContract.contract.WatchLogs(opts, "DealRefundReason")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantDealsHistoryContractDealRefundReason)
				if err := _MerchantDealsHistoryContract.contract.UnpackLog(event, "DealRefundReason", log); err != nil {
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

// MerchantDealsHistoryContractMonethaAddressSetIterator is returned from FilterMonethaAddressSet and is used to iterate over the raw logs and unpacked data for MonethaAddressSet events raised by the MerchantDealsHistoryContract contract.
type MerchantDealsHistoryContractMonethaAddressSetIterator struct {
	Event *MerchantDealsHistoryContractMonethaAddressSet // Event containing the contract specifics and raw log

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
func (it *MerchantDealsHistoryContractMonethaAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantDealsHistoryContractMonethaAddressSet)
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
		it.Event = new(MerchantDealsHistoryContractMonethaAddressSet)
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
func (it *MerchantDealsHistoryContractMonethaAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantDealsHistoryContractMonethaAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantDealsHistoryContractMonethaAddressSet represents a MonethaAddressSet event raised by the MerchantDealsHistoryContract contract.
type MerchantDealsHistoryContractMonethaAddressSet struct {
	Address          common.Address
	IsMonethaAddress bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMonethaAddressSet is a free log retrieval operation binding the contract event 0xa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e692.
//
// Solidity: e MonethaAddressSet(_address address, _isMonethaAddress bool)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractFilterer) FilterMonethaAddressSet(opts *bind.FilterOpts) (*MerchantDealsHistoryContractMonethaAddressSetIterator, error) {

	logs, sub, err := _MerchantDealsHistoryContract.contract.FilterLogs(opts, "MonethaAddressSet")
	if err != nil {
		return nil, err
	}
	return &MerchantDealsHistoryContractMonethaAddressSetIterator{contract: _MerchantDealsHistoryContract.contract, event: "MonethaAddressSet", logs: logs, sub: sub}, nil
}

// WatchMonethaAddressSet is a free log subscription operation binding the contract event 0xa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e692.
//
// Solidity: e MonethaAddressSet(_address address, _isMonethaAddress bool)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractFilterer) WatchMonethaAddressSet(opts *bind.WatchOpts, sink chan<- *MerchantDealsHistoryContractMonethaAddressSet) (event.Subscription, error) {

	logs, sub, err := _MerchantDealsHistoryContract.contract.WatchLogs(opts, "MonethaAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantDealsHistoryContractMonethaAddressSet)
				if err := _MerchantDealsHistoryContract.contract.UnpackLog(event, "MonethaAddressSet", log); err != nil {
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

// MerchantDealsHistoryContractOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the MerchantDealsHistoryContract contract.
type MerchantDealsHistoryContractOwnershipRenouncedIterator struct {
	Event *MerchantDealsHistoryContractOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *MerchantDealsHistoryContractOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantDealsHistoryContractOwnershipRenounced)
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
		it.Event = new(MerchantDealsHistoryContractOwnershipRenounced)
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
func (it *MerchantDealsHistoryContractOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantDealsHistoryContractOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantDealsHistoryContractOwnershipRenounced represents a OwnershipRenounced event raised by the MerchantDealsHistoryContract contract.
type MerchantDealsHistoryContractOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*MerchantDealsHistoryContractOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _MerchantDealsHistoryContract.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MerchantDealsHistoryContractOwnershipRenouncedIterator{contract: _MerchantDealsHistoryContract.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *MerchantDealsHistoryContractOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _MerchantDealsHistoryContract.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantDealsHistoryContractOwnershipRenounced)
				if err := _MerchantDealsHistoryContract.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// MerchantDealsHistoryContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MerchantDealsHistoryContract contract.
type MerchantDealsHistoryContractOwnershipTransferredIterator struct {
	Event *MerchantDealsHistoryContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MerchantDealsHistoryContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantDealsHistoryContractOwnershipTransferred)
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
		it.Event = new(MerchantDealsHistoryContractOwnershipTransferred)
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
func (it *MerchantDealsHistoryContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantDealsHistoryContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantDealsHistoryContractOwnershipTransferred represents a OwnershipTransferred event raised by the MerchantDealsHistoryContract contract.
type MerchantDealsHistoryContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MerchantDealsHistoryContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MerchantDealsHistoryContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MerchantDealsHistoryContractOwnershipTransferredIterator{contract: _MerchantDealsHistoryContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_MerchantDealsHistoryContract *MerchantDealsHistoryContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MerchantDealsHistoryContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MerchantDealsHistoryContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantDealsHistoryContractOwnershipTransferred)
				if err := _MerchantDealsHistoryContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

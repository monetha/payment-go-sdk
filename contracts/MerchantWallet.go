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

// MerchantWalletContractABI is the input ABI used to generate the binding from.
const MerchantWalletContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isMonethaAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contactInformation\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"merchantAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"merchantFundAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_info\",\"type\":\"string\"}],\"name\":\"setContactInformation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_isMonethaAddress\",\"type\":\"bool\"}],\"name\":\"setMonethaAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"merchantIdHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"REPUTATION_DECIMALS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_merchantAccount\",\"type\":\"address\"},{\"name\":\"_merchantId\",\"type\":\"string\"},{\"name\":\"_fundAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_isMonethaAddress\",\"type\":\"bool\"}],\"name\":\"MonethaAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"profile\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"paymentSettings\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"compositeReputation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"profileKey\",\"type\":\"string\"},{\"name\":\"profileValue\",\"type\":\"string\"},{\"name\":\"repKey\",\"type\":\"string\"},{\"name\":\"repValue\",\"type\":\"uint32\"}],\"name\":\"setProfile\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setPaymentSettings\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"uint32\"}],\"name\":\"setCompositeReputation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"depositAccount\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToExchange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"depositAccount\",\"type\":\"address\"},{\"name\":\"min_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAllToExchange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_depositAccount\",\"type\":\"address\"},{\"name\":\"_minAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawAllTokensToExchange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAccount\",\"type\":\"address\"}],\"name\":\"changeMerchantAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newFundAddress\",\"type\":\"address\"}],\"name\":\"changeFundAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MerchantWalletContractBin is the compiled bytecode used for deploying new contracts.
const MerchantWalletContractBin = `0x60806040526000805460a060020a60ff021916905534801561002057600080fd5b50604051611384380380611384833981016040908152815160208301519183015160008054600160a060020a031916331790559092919091019080803b80156100ca57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f736f7272792068756d616e73206f6e6c79000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a03851615156100df57600080fd5b83516000106100ed57600080fd5b60038054600160a060020a031916600160a060020a0387161790556040518451859160209081019182918401908083835b6020831061013d5780518252601f19909201916020918201910161011e565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b602083106101a05780518252601f199092019160209182019101610181565b5181516000196020949094036101000a939093019283169219169190911790526040519201829003909120600555505060048054600160a060020a03909516600160a060020a031990951694909417909355505050505061117e806102066000396000f3006080604052600436106101695763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663061b3245811461016b5780630cf1d0501461018f578063205c2878146101b35780632aa0da47146101d75780632e1a7d4d1461026c57806331d413251461028457806336f7ab5e146102b9578063370fed6e146102ce5780633f4ba83a1461030f5780635c975abb146103245780636339782514610339578063715018a61461035a5780637dd917341461036f57806383197ef0146103a85780638456cb59146103bd5780638da5cb5b146103d2578063957f050b14610403578063969596d614610418578063b967a52e1461042d578063c0462ec314610486578063c07e3391146104b0578063e25a51b6146104d6578063e3577e7114610501578063e8c0485f1461052d578063f0daba011461054d578063f2cba6a614610574578063f2fde38b1461059f578063f9271fd3146105c0575b005b34801561017757600080fd5b50610169600160a060020a03600435166024356105e1565b34801561019b57600080fd5b50610169600160a060020a0360043516602435610639565b3480156101bf57600080fd5b50610169600160a060020a036004351660243561069c565b3480156101e357600080fd5b506101f760048035602481019101356106b3565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610231578181015183820152602001610219565b50505050905090810190601f16801561025e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561027857600080fd5b50610169600435610768565b34801561029057600080fd5b506102a5600160a060020a036004351661078c565b604080519115158252519081900360200190f35b3480156102c557600080fd5b506101f76107a1565b3480156102da57600080fd5b50610169602460048035828101929082013591813580830192908201359160443591820191013563ffffffff6064351661082e565b34801561031b57600080fd5b506101696108cb565b34801561033057600080fd5b506102a5610941565b34801561034557600080fd5b50610169600160a060020a0360043516610951565b34801561036657600080fd5b506101696109ae565b34801561037b57600080fd5b5061038f6004803560248101910135610a1a565b6040805163ffffffff9092168252519081900360200190f35b3480156103b457600080fd5b50610169610a50565b3480156103c957600080fd5b50610169610a81565b3480156103de57600080fd5b506103e7610afc565b60408051600160a060020a039092168252519081900360200190f35b34801561040f57600080fd5b506103e7610b0b565b34801561042457600080fd5b506103e7610b1a565b34801561043957600080fd5b506040805160206004803580820135601f8101849004840285018401909552848452610169943694929360249392840191908190840183828082843750949750610b299650505050505050565b34801561049257600080fd5b50610169600160a060020a0360043581169060243516604435610b53565b3480156104bc57600080fd5b50610169600160a060020a03600435166024351515610ceb565b3480156104e257600080fd5b50610169602460048035828101929101359063ffffffff903516610d66565b34801561050d57600080fd5b506101696024600480358281019290820135918135918201910135610dca565b34801561053957600080fd5b506101f76004803560248101910135610e1c565b34801561055957600080fd5b50610562610e99565b60408051918252519081900360200190f35b34801561058057600080fd5b50610589610e9f565b6040805160ff9092168252519081900360200190f35b3480156105ab57600080fd5b50610169600160a060020a0360043516610ea4565b3480156105cc57600080fd5b50610169600160a060020a0360043516610ec4565b600354600160a060020a031633148061060957503360009081526002602052604090205460ff165b151561061457600080fd5b60005460a060020a900460ff161561062b57600080fd5b6106358282610f7c565b5050565b600354600160a060020a031633148061066157503360009081526002602052604090205460ff165b151561066c57600080fd5b60005460a060020a900460ff161561068357600080fd5b303181111561069157600080fd5b610635823031610f7c565b600354600160a060020a0316331461061457600080fd5b606060078383604051808383808284379190910194855250506040805160209481900385018120805460026001821615610100026000190190911604601f8101879004870283018701909352828252909490935090915083018282801561075b5780601f106107305761010080835404028352916020019161075b565b820191906000526020600020905b81548152906001019060200180831161073e57829003601f168201915b5050505050905092915050565b600354600160a060020a0316331461077f57600080fd5b610789338261069c565b50565b60026020526000908152604090205460ff1681565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156108265780601f106107fb57610100808354040283529160200191610826565b820191906000526020600020905b81548152906001019060200180831161080957829003601f168201915b505050505081565b600054600160a060020a0316331461084557600080fd5b8484600689896040518083838082843782019150509250505090815260200160405180910390209190610879929190611049565b5081156108c257806008848460405180838380828437909101948552505060405192839003602001909220805463ffffffff9490941663ffffffff199094169390931790925550505b50505050505050565b600054600160a060020a031633146108e257600080fd5b60005460a060020a900460ff1615156108fa57600080fd5b6000805474ff0000000000000000000000000000000000000000191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a1565b60005460a060020a900460ff1681565b600354600160a060020a0316331461096857600080fd5b60005460a060020a900460ff161561097f57600080fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a031633146109c557600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600060088383604051808383808284379091019485525050604051928390036020019092205463ffffffff169250505092915050565b600054600160a060020a03163314610a6757600080fd5b303115610a7357600080fd5b600054600160a060020a0316ff5b600054600160a060020a03163314610a9857600080fd5b60005460a060020a900460ff1615610aaf57600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a1565b600054600160a060020a031681565b600354600160a060020a031681565b600454600160a060020a031681565b600054600160a060020a03163314610b4057600080fd5b80516106359060019060208401906110c7565b600354600090600160a060020a0316331480610b7e57503360009081526002602052604090205460ff165b1515610b8957600080fd5b60005460a060020a900460ff1615610ba057600080fd5b600160a060020a0384161515610bb557600080fd5b604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051600160a060020a038616916370a082319160248083019260209291908290030181600087803b158015610c1657600080fd5b505af1158015610c2a573d6000803e3d6000fd5b505050506040513d6020811015610c4057600080fd5b5051905081811015610c5157600080fd5b83600160a060020a031663a9059cbb84836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050600060405180830381600087803b158015610ccd57600080fd5b505af1158015610ce1573d6000803e3d6000fd5b5050505050505050565b600054600160a060020a03163314610d0257600080fd5b600160a060020a038216600081815260026020908152604091829020805460ff191685151590811790915582519384529083015280517fa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e6929281900390910190a15050565b3360009081526002602052604090205460ff161515610d8457600080fd5b806008848460405180838380828437909101948552505060405192839003602001909220805463ffffffff9490941663ffffffff19909416939093179092555050505050565b600054600160a060020a03163314610de157600080fd5b8181600786866040518083838082843782019150509250505090815260200160405180910390209190610e15929190611049565b5050505050565b606060068383604051808383808284379190910194855250506040805160209481900385018120805460026001821615610100026000190190911604601f8101879004870283018701909352828252909490935090915083018282801561075b5780601f106107305761010080835404028352916020019161075b565b60055481565b600481565b600054600160a060020a03163314610ebb57600080fd5b61078981610fcc565b600354600160a060020a03163314610edb57600080fd5b80803b8015610f4b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f736f7272792068756d616e73206f6e6c79000000000000000000000000000000604482015290519081900360640190fd5b50506004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600160a060020a0382161515610f9157600080fd5b604051600160a060020a0383169082156108fc029083906000818181858888f19350505050158015610fc7573d6000803e3d6000fd5b505050565b600160a060020a0381161515610fe157600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061108a5782800160ff198235161785556110b7565b828001600101855582156110b7579182015b828111156110b757823582559160200191906001019061109c565b506110c3929150611135565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061110857805160ff19168380011785556110b7565b828001600101855582156110b7579182015b828111156110b757825182559160200191906001019061111a565b61114f91905b808211156110c3576000815560010161113b565b905600a165627a7a72305820fcd0552feacca9c336924a52d9c6c48a522877def1fd8726f5ff152574fd12de0029`

// DeployMerchantWalletContract deploys a new Ethereum contract, binding an instance of MerchantWalletContract to it.
func DeployMerchantWalletContract(auth *bind.TransactOpts, backend bind.ContractBackend, _merchantAccount common.Address, _merchantId string, _fundAddress common.Address) (common.Address, *types.Transaction, *MerchantWalletContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerchantWalletContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MerchantWalletContractBin), backend, _merchantAccount, _merchantId, _fundAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerchantWalletContract{MerchantWalletContractCaller: MerchantWalletContractCaller{contract: contract}, MerchantWalletContractTransactor: MerchantWalletContractTransactor{contract: contract}, MerchantWalletContractFilterer: MerchantWalletContractFilterer{contract: contract}}, nil
}

// MerchantWalletContract is an auto generated Go binding around an Ethereum contract.
type MerchantWalletContract struct {
	MerchantWalletContractCaller     // Read-only binding to the contract
	MerchantWalletContractTransactor // Write-only binding to the contract
	MerchantWalletContractFilterer   // Log filterer for contract events
}

// MerchantWalletContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerchantWalletContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerchantWalletContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerchantWalletContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerchantWalletContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerchantWalletContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerchantWalletContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerchantWalletContractSession struct {
	Contract     *MerchantWalletContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MerchantWalletContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerchantWalletContractCallerSession struct {
	Contract *MerchantWalletContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// MerchantWalletContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerchantWalletContractTransactorSession struct {
	Contract     *MerchantWalletContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// MerchantWalletContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerchantWalletContractRaw struct {
	Contract *MerchantWalletContract // Generic contract binding to access the raw methods on
}

// MerchantWalletContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerchantWalletContractCallerRaw struct {
	Contract *MerchantWalletContractCaller // Generic read-only contract binding to access the raw methods on
}

// MerchantWalletContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerchantWalletContractTransactorRaw struct {
	Contract *MerchantWalletContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerchantWalletContract creates a new instance of MerchantWalletContract, bound to a specific deployed contract.
func NewMerchantWalletContract(address common.Address, backend bind.ContractBackend) (*MerchantWalletContract, error) {
	contract, err := bindMerchantWalletContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerchantWalletContract{MerchantWalletContractCaller: MerchantWalletContractCaller{contract: contract}, MerchantWalletContractTransactor: MerchantWalletContractTransactor{contract: contract}, MerchantWalletContractFilterer: MerchantWalletContractFilterer{contract: contract}}, nil
}

// NewMerchantWalletContractCaller creates a new read-only instance of MerchantWalletContract, bound to a specific deployed contract.
func NewMerchantWalletContractCaller(address common.Address, caller bind.ContractCaller) (*MerchantWalletContractCaller, error) {
	contract, err := bindMerchantWalletContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerchantWalletContractCaller{contract: contract}, nil
}

// NewMerchantWalletContractTransactor creates a new write-only instance of MerchantWalletContract, bound to a specific deployed contract.
func NewMerchantWalletContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MerchantWalletContractTransactor, error) {
	contract, err := bindMerchantWalletContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerchantWalletContractTransactor{contract: contract}, nil
}

// NewMerchantWalletContractFilterer creates a new log filterer instance of MerchantWalletContract, bound to a specific deployed contract.
func NewMerchantWalletContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MerchantWalletContractFilterer, error) {
	contract, err := bindMerchantWalletContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerchantWalletContractFilterer{contract: contract}, nil
}

// bindMerchantWalletContract binds a generic wrapper to an already deployed contract.
func bindMerchantWalletContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerchantWalletContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerchantWalletContract *MerchantWalletContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerchantWalletContract.Contract.MerchantWalletContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerchantWalletContract *MerchantWalletContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.MerchantWalletContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerchantWalletContract *MerchantWalletContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.MerchantWalletContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerchantWalletContract *MerchantWalletContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerchantWalletContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerchantWalletContract *MerchantWalletContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerchantWalletContract *MerchantWalletContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.contract.Transact(opts, method, params...)
}

// REPUTATIONDECIMALS is a free data retrieval call binding the contract method 0xf2cba6a6.
//
// Solidity: function REPUTATION_DECIMALS() constant returns(uint8)
func (_MerchantWalletContract *MerchantWalletContractCaller) REPUTATIONDECIMALS(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MerchantWalletContract.contract.Call(opts, out, "REPUTATION_DECIMALS")
	return *ret0, err
}

// REPUTATIONDECIMALS is a free data retrieval call binding the contract method 0xf2cba6a6.
//
// Solidity: function REPUTATION_DECIMALS() constant returns(uint8)
func (_MerchantWalletContract *MerchantWalletContractSession) REPUTATIONDECIMALS() (uint8, error) {
	return _MerchantWalletContract.Contract.REPUTATIONDECIMALS(&_MerchantWalletContract.CallOpts)
}

// REPUTATIONDECIMALS is a free data retrieval call binding the contract method 0xf2cba6a6.
//
// Solidity: function REPUTATION_DECIMALS() constant returns(uint8)
func (_MerchantWalletContract *MerchantWalletContractCallerSession) REPUTATIONDECIMALS() (uint8, error) {
	return _MerchantWalletContract.Contract.REPUTATIONDECIMALS(&_MerchantWalletContract.CallOpts)
}

// CompositeReputation is a free data retrieval call binding the contract method 0x7dd91734.
//
// Solidity: function compositeReputation(key string) constant returns(uint32)
func (_MerchantWalletContract *MerchantWalletContractCaller) CompositeReputation(opts *bind.CallOpts, key string) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _MerchantWalletContract.contract.Call(opts, out, "compositeReputation", key)
	return *ret0, err
}

// CompositeReputation is a free data retrieval call binding the contract method 0x7dd91734.
//
// Solidity: function compositeReputation(key string) constant returns(uint32)
func (_MerchantWalletContract *MerchantWalletContractSession) CompositeReputation(key string) (uint32, error) {
	return _MerchantWalletContract.Contract.CompositeReputation(&_MerchantWalletContract.CallOpts, key)
}

// CompositeReputation is a free data retrieval call binding the contract method 0x7dd91734.
//
// Solidity: function compositeReputation(key string) constant returns(uint32)
func (_MerchantWalletContract *MerchantWalletContractCallerSession) CompositeReputation(key string) (uint32, error) {
	return _MerchantWalletContract.Contract.CompositeReputation(&_MerchantWalletContract.CallOpts, key)
}

// ContactInformation is a free data retrieval call binding the contract method 0x36f7ab5e.
//
// Solidity: function contactInformation() constant returns(string)
func (_MerchantWalletContract *MerchantWalletContractCaller) ContactInformation(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MerchantWalletContract.contract.Call(opts, out, "contactInformation")
	return *ret0, err
}

// ContactInformation is a free data retrieval call binding the contract method 0x36f7ab5e.
//
// Solidity: function contactInformation() constant returns(string)
func (_MerchantWalletContract *MerchantWalletContractSession) ContactInformation() (string, error) {
	return _MerchantWalletContract.Contract.ContactInformation(&_MerchantWalletContract.CallOpts)
}

// ContactInformation is a free data retrieval call binding the contract method 0x36f7ab5e.
//
// Solidity: function contactInformation() constant returns(string)
func (_MerchantWalletContract *MerchantWalletContractCallerSession) ContactInformation() (string, error) {
	return _MerchantWalletContract.Contract.ContactInformation(&_MerchantWalletContract.CallOpts)
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_MerchantWalletContract *MerchantWalletContractCaller) IsMonethaAddress(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MerchantWalletContract.contract.Call(opts, out, "isMonethaAddress", arg0)
	return *ret0, err
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_MerchantWalletContract *MerchantWalletContractSession) IsMonethaAddress(arg0 common.Address) (bool, error) {
	return _MerchantWalletContract.Contract.IsMonethaAddress(&_MerchantWalletContract.CallOpts, arg0)
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_MerchantWalletContract *MerchantWalletContractCallerSession) IsMonethaAddress(arg0 common.Address) (bool, error) {
	return _MerchantWalletContract.Contract.IsMonethaAddress(&_MerchantWalletContract.CallOpts, arg0)
}

// MerchantAccount is a free data retrieval call binding the contract method 0x957f050b.
//
// Solidity: function merchantAccount() constant returns(address)
func (_MerchantWalletContract *MerchantWalletContractCaller) MerchantAccount(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MerchantWalletContract.contract.Call(opts, out, "merchantAccount")
	return *ret0, err
}

// MerchantAccount is a free data retrieval call binding the contract method 0x957f050b.
//
// Solidity: function merchantAccount() constant returns(address)
func (_MerchantWalletContract *MerchantWalletContractSession) MerchantAccount() (common.Address, error) {
	return _MerchantWalletContract.Contract.MerchantAccount(&_MerchantWalletContract.CallOpts)
}

// MerchantAccount is a free data retrieval call binding the contract method 0x957f050b.
//
// Solidity: function merchantAccount() constant returns(address)
func (_MerchantWalletContract *MerchantWalletContractCallerSession) MerchantAccount() (common.Address, error) {
	return _MerchantWalletContract.Contract.MerchantAccount(&_MerchantWalletContract.CallOpts)
}

// MerchantFundAddress is a free data retrieval call binding the contract method 0x969596d6.
//
// Solidity: function merchantFundAddress() constant returns(address)
func (_MerchantWalletContract *MerchantWalletContractCaller) MerchantFundAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MerchantWalletContract.contract.Call(opts, out, "merchantFundAddress")
	return *ret0, err
}

// MerchantFundAddress is a free data retrieval call binding the contract method 0x969596d6.
//
// Solidity: function merchantFundAddress() constant returns(address)
func (_MerchantWalletContract *MerchantWalletContractSession) MerchantFundAddress() (common.Address, error) {
	return _MerchantWalletContract.Contract.MerchantFundAddress(&_MerchantWalletContract.CallOpts)
}

// MerchantFundAddress is a free data retrieval call binding the contract method 0x969596d6.
//
// Solidity: function merchantFundAddress() constant returns(address)
func (_MerchantWalletContract *MerchantWalletContractCallerSession) MerchantFundAddress() (common.Address, error) {
	return _MerchantWalletContract.Contract.MerchantFundAddress(&_MerchantWalletContract.CallOpts)
}

// MerchantIdHash is a free data retrieval call binding the contract method 0xf0daba01.
//
// Solidity: function merchantIdHash() constant returns(bytes32)
func (_MerchantWalletContract *MerchantWalletContractCaller) MerchantIdHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MerchantWalletContract.contract.Call(opts, out, "merchantIdHash")
	return *ret0, err
}

// MerchantIdHash is a free data retrieval call binding the contract method 0xf0daba01.
//
// Solidity: function merchantIdHash() constant returns(bytes32)
func (_MerchantWalletContract *MerchantWalletContractSession) MerchantIdHash() ([32]byte, error) {
	return _MerchantWalletContract.Contract.MerchantIdHash(&_MerchantWalletContract.CallOpts)
}

// MerchantIdHash is a free data retrieval call binding the contract method 0xf0daba01.
//
// Solidity: function merchantIdHash() constant returns(bytes32)
func (_MerchantWalletContract *MerchantWalletContractCallerSession) MerchantIdHash() ([32]byte, error) {
	return _MerchantWalletContract.Contract.MerchantIdHash(&_MerchantWalletContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MerchantWalletContract *MerchantWalletContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MerchantWalletContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MerchantWalletContract *MerchantWalletContractSession) Owner() (common.Address, error) {
	return _MerchantWalletContract.Contract.Owner(&_MerchantWalletContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MerchantWalletContract *MerchantWalletContractCallerSession) Owner() (common.Address, error) {
	return _MerchantWalletContract.Contract.Owner(&_MerchantWalletContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MerchantWalletContract *MerchantWalletContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MerchantWalletContract.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MerchantWalletContract *MerchantWalletContractSession) Paused() (bool, error) {
	return _MerchantWalletContract.Contract.Paused(&_MerchantWalletContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MerchantWalletContract *MerchantWalletContractCallerSession) Paused() (bool, error) {
	return _MerchantWalletContract.Contract.Paused(&_MerchantWalletContract.CallOpts)
}

// PaymentSettings is a free data retrieval call binding the contract method 0x2aa0da47.
//
// Solidity: function paymentSettings(key string) constant returns(string)
func (_MerchantWalletContract *MerchantWalletContractCaller) PaymentSettings(opts *bind.CallOpts, key string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MerchantWalletContract.contract.Call(opts, out, "paymentSettings", key)
	return *ret0, err
}

// PaymentSettings is a free data retrieval call binding the contract method 0x2aa0da47.
//
// Solidity: function paymentSettings(key string) constant returns(string)
func (_MerchantWalletContract *MerchantWalletContractSession) PaymentSettings(key string) (string, error) {
	return _MerchantWalletContract.Contract.PaymentSettings(&_MerchantWalletContract.CallOpts, key)
}

// PaymentSettings is a free data retrieval call binding the contract method 0x2aa0da47.
//
// Solidity: function paymentSettings(key string) constant returns(string)
func (_MerchantWalletContract *MerchantWalletContractCallerSession) PaymentSettings(key string) (string, error) {
	return _MerchantWalletContract.Contract.PaymentSettings(&_MerchantWalletContract.CallOpts, key)
}

// Profile is a free data retrieval call binding the contract method 0xe8c0485f.
//
// Solidity: function profile(key string) constant returns(string)
func (_MerchantWalletContract *MerchantWalletContractCaller) Profile(opts *bind.CallOpts, key string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MerchantWalletContract.contract.Call(opts, out, "profile", key)
	return *ret0, err
}

// Profile is a free data retrieval call binding the contract method 0xe8c0485f.
//
// Solidity: function profile(key string) constant returns(string)
func (_MerchantWalletContract *MerchantWalletContractSession) Profile(key string) (string, error) {
	return _MerchantWalletContract.Contract.Profile(&_MerchantWalletContract.CallOpts, key)
}

// Profile is a free data retrieval call binding the contract method 0xe8c0485f.
//
// Solidity: function profile(key string) constant returns(string)
func (_MerchantWalletContract *MerchantWalletContractCallerSession) Profile(key string) (string, error) {
	return _MerchantWalletContract.Contract.Profile(&_MerchantWalletContract.CallOpts, key)
}

// ChangeFundAddress is a paid mutator transaction binding the contract method 0xf9271fd3.
//
// Solidity: function changeFundAddress(newFundAddress address) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactor) ChangeFundAddress(opts *bind.TransactOpts, newFundAddress common.Address) (*types.Transaction, error) {
	return _MerchantWalletContract.contract.Transact(opts, "changeFundAddress", newFundAddress)
}

// ChangeFundAddress is a paid mutator transaction binding the contract method 0xf9271fd3.
//
// Solidity: function changeFundAddress(newFundAddress address) returns()
func (_MerchantWalletContract *MerchantWalletContractSession) ChangeFundAddress(newFundAddress common.Address) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.ChangeFundAddress(&_MerchantWalletContract.TransactOpts, newFundAddress)
}

// ChangeFundAddress is a paid mutator transaction binding the contract method 0xf9271fd3.
//
// Solidity: function changeFundAddress(newFundAddress address) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactorSession) ChangeFundAddress(newFundAddress common.Address) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.ChangeFundAddress(&_MerchantWalletContract.TransactOpts, newFundAddress)
}

// ChangeMerchantAccount is a paid mutator transaction binding the contract method 0x63397825.
//
// Solidity: function changeMerchantAccount(newAccount address) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactor) ChangeMerchantAccount(opts *bind.TransactOpts, newAccount common.Address) (*types.Transaction, error) {
	return _MerchantWalletContract.contract.Transact(opts, "changeMerchantAccount", newAccount)
}

// ChangeMerchantAccount is a paid mutator transaction binding the contract method 0x63397825.
//
// Solidity: function changeMerchantAccount(newAccount address) returns()
func (_MerchantWalletContract *MerchantWalletContractSession) ChangeMerchantAccount(newAccount common.Address) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.ChangeMerchantAccount(&_MerchantWalletContract.TransactOpts, newAccount)
}

// ChangeMerchantAccount is a paid mutator transaction binding the contract method 0x63397825.
//
// Solidity: function changeMerchantAccount(newAccount address) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactorSession) ChangeMerchantAccount(newAccount common.Address) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.ChangeMerchantAccount(&_MerchantWalletContract.TransactOpts, newAccount)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_MerchantWalletContract *MerchantWalletContractTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantWalletContract.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_MerchantWalletContract *MerchantWalletContractSession) Destroy() (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.Destroy(&_MerchantWalletContract.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_MerchantWalletContract *MerchantWalletContractTransactorSession) Destroy() (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.Destroy(&_MerchantWalletContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MerchantWalletContract *MerchantWalletContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantWalletContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MerchantWalletContract *MerchantWalletContractSession) Pause() (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.Pause(&_MerchantWalletContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MerchantWalletContract *MerchantWalletContractTransactorSession) Pause() (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.Pause(&_MerchantWalletContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MerchantWalletContract *MerchantWalletContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantWalletContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MerchantWalletContract *MerchantWalletContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.RenounceOwnership(&_MerchantWalletContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MerchantWalletContract *MerchantWalletContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.RenounceOwnership(&_MerchantWalletContract.TransactOpts)
}

// SetCompositeReputation is a paid mutator transaction binding the contract method 0xe25a51b6.
//
// Solidity: function setCompositeReputation(key string, value uint32) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactor) SetCompositeReputation(opts *bind.TransactOpts, key string, value uint32) (*types.Transaction, error) {
	return _MerchantWalletContract.contract.Transact(opts, "setCompositeReputation", key, value)
}

// SetCompositeReputation is a paid mutator transaction binding the contract method 0xe25a51b6.
//
// Solidity: function setCompositeReputation(key string, value uint32) returns()
func (_MerchantWalletContract *MerchantWalletContractSession) SetCompositeReputation(key string, value uint32) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.SetCompositeReputation(&_MerchantWalletContract.TransactOpts, key, value)
}

// SetCompositeReputation is a paid mutator transaction binding the contract method 0xe25a51b6.
//
// Solidity: function setCompositeReputation(key string, value uint32) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactorSession) SetCompositeReputation(key string, value uint32) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.SetCompositeReputation(&_MerchantWalletContract.TransactOpts, key, value)
}

// SetContactInformation is a paid mutator transaction binding the contract method 0xb967a52e.
//
// Solidity: function setContactInformation(_info string) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactor) SetContactInformation(opts *bind.TransactOpts, _info string) (*types.Transaction, error) {
	return _MerchantWalletContract.contract.Transact(opts, "setContactInformation", _info)
}

// SetContactInformation is a paid mutator transaction binding the contract method 0xb967a52e.
//
// Solidity: function setContactInformation(_info string) returns()
func (_MerchantWalletContract *MerchantWalletContractSession) SetContactInformation(_info string) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.SetContactInformation(&_MerchantWalletContract.TransactOpts, _info)
}

// SetContactInformation is a paid mutator transaction binding the contract method 0xb967a52e.
//
// Solidity: function setContactInformation(_info string) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactorSession) SetContactInformation(_info string) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.SetContactInformation(&_MerchantWalletContract.TransactOpts, _info)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactor) SetMonethaAddress(opts *bind.TransactOpts, _address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _MerchantWalletContract.contract.Transact(opts, "setMonethaAddress", _address, _isMonethaAddress)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_MerchantWalletContract *MerchantWalletContractSession) SetMonethaAddress(_address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.SetMonethaAddress(&_MerchantWalletContract.TransactOpts, _address, _isMonethaAddress)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactorSession) SetMonethaAddress(_address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.SetMonethaAddress(&_MerchantWalletContract.TransactOpts, _address, _isMonethaAddress)
}

// SetPaymentSettings is a paid mutator transaction binding the contract method 0xe3577e71.
//
// Solidity: function setPaymentSettings(key string, value string) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactor) SetPaymentSettings(opts *bind.TransactOpts, key string, value string) (*types.Transaction, error) {
	return _MerchantWalletContract.contract.Transact(opts, "setPaymentSettings", key, value)
}

// SetPaymentSettings is a paid mutator transaction binding the contract method 0xe3577e71.
//
// Solidity: function setPaymentSettings(key string, value string) returns()
func (_MerchantWalletContract *MerchantWalletContractSession) SetPaymentSettings(key string, value string) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.SetPaymentSettings(&_MerchantWalletContract.TransactOpts, key, value)
}

// SetPaymentSettings is a paid mutator transaction binding the contract method 0xe3577e71.
//
// Solidity: function setPaymentSettings(key string, value string) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactorSession) SetPaymentSettings(key string, value string) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.SetPaymentSettings(&_MerchantWalletContract.TransactOpts, key, value)
}

// SetProfile is a paid mutator transaction binding the contract method 0x370fed6e.
//
// Solidity: function setProfile(profileKey string, profileValue string, repKey string, repValue uint32) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactor) SetProfile(opts *bind.TransactOpts, profileKey string, profileValue string, repKey string, repValue uint32) (*types.Transaction, error) {
	return _MerchantWalletContract.contract.Transact(opts, "setProfile", profileKey, profileValue, repKey, repValue)
}

// SetProfile is a paid mutator transaction binding the contract method 0x370fed6e.
//
// Solidity: function setProfile(profileKey string, profileValue string, repKey string, repValue uint32) returns()
func (_MerchantWalletContract *MerchantWalletContractSession) SetProfile(profileKey string, profileValue string, repKey string, repValue uint32) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.SetProfile(&_MerchantWalletContract.TransactOpts, profileKey, profileValue, repKey, repValue)
}

// SetProfile is a paid mutator transaction binding the contract method 0x370fed6e.
//
// Solidity: function setProfile(profileKey string, profileValue string, repKey string, repValue uint32) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactorSession) SetProfile(profileKey string, profileValue string, repKey string, repValue uint32) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.SetProfile(&_MerchantWalletContract.TransactOpts, profileKey, profileValue, repKey, repValue)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _MerchantWalletContract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MerchantWalletContract *MerchantWalletContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.TransferOwnership(&_MerchantWalletContract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.TransferOwnership(&_MerchantWalletContract.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MerchantWalletContract *MerchantWalletContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantWalletContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MerchantWalletContract *MerchantWalletContractSession) Unpause() (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.Unpause(&_MerchantWalletContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MerchantWalletContract *MerchantWalletContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.Unpause(&_MerchantWalletContract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MerchantWalletContract.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_MerchantWalletContract *MerchantWalletContractSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.Withdraw(&_MerchantWalletContract.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.Withdraw(&_MerchantWalletContract.TransactOpts, amount)
}

// WithdrawAllToExchange is a paid mutator transaction binding the contract method 0x0cf1d050.
//
// Solidity: function withdrawAllToExchange(depositAccount address, min_amount uint256) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactor) WithdrawAllToExchange(opts *bind.TransactOpts, depositAccount common.Address, min_amount *big.Int) (*types.Transaction, error) {
	return _MerchantWalletContract.contract.Transact(opts, "withdrawAllToExchange", depositAccount, min_amount)
}

// WithdrawAllToExchange is a paid mutator transaction binding the contract method 0x0cf1d050.
//
// Solidity: function withdrawAllToExchange(depositAccount address, min_amount uint256) returns()
func (_MerchantWalletContract *MerchantWalletContractSession) WithdrawAllToExchange(depositAccount common.Address, min_amount *big.Int) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.WithdrawAllToExchange(&_MerchantWalletContract.TransactOpts, depositAccount, min_amount)
}

// WithdrawAllToExchange is a paid mutator transaction binding the contract method 0x0cf1d050.
//
// Solidity: function withdrawAllToExchange(depositAccount address, min_amount uint256) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactorSession) WithdrawAllToExchange(depositAccount common.Address, min_amount *big.Int) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.WithdrawAllToExchange(&_MerchantWalletContract.TransactOpts, depositAccount, min_amount)
}

// WithdrawAllTokensToExchange is a paid mutator transaction binding the contract method 0xc0462ec3.
//
// Solidity: function withdrawAllTokensToExchange(_tokenAddress address, _depositAccount address, _minAmount uint256) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactor) WithdrawAllTokensToExchange(opts *bind.TransactOpts, _tokenAddress common.Address, _depositAccount common.Address, _minAmount *big.Int) (*types.Transaction, error) {
	return _MerchantWalletContract.contract.Transact(opts, "withdrawAllTokensToExchange", _tokenAddress, _depositAccount, _minAmount)
}

// WithdrawAllTokensToExchange is a paid mutator transaction binding the contract method 0xc0462ec3.
//
// Solidity: function withdrawAllTokensToExchange(_tokenAddress address, _depositAccount address, _minAmount uint256) returns()
func (_MerchantWalletContract *MerchantWalletContractSession) WithdrawAllTokensToExchange(_tokenAddress common.Address, _depositAccount common.Address, _minAmount *big.Int) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.WithdrawAllTokensToExchange(&_MerchantWalletContract.TransactOpts, _tokenAddress, _depositAccount, _minAmount)
}

// WithdrawAllTokensToExchange is a paid mutator transaction binding the contract method 0xc0462ec3.
//
// Solidity: function withdrawAllTokensToExchange(_tokenAddress address, _depositAccount address, _minAmount uint256) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactorSession) WithdrawAllTokensToExchange(_tokenAddress common.Address, _depositAccount common.Address, _minAmount *big.Int) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.WithdrawAllTokensToExchange(&_MerchantWalletContract.TransactOpts, _tokenAddress, _depositAccount, _minAmount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(beneficiary address, amount uint256) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactor) WithdrawTo(opts *bind.TransactOpts, beneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MerchantWalletContract.contract.Transact(opts, "withdrawTo", beneficiary, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(beneficiary address, amount uint256) returns()
func (_MerchantWalletContract *MerchantWalletContractSession) WithdrawTo(beneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.WithdrawTo(&_MerchantWalletContract.TransactOpts, beneficiary, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(beneficiary address, amount uint256) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactorSession) WithdrawTo(beneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.WithdrawTo(&_MerchantWalletContract.TransactOpts, beneficiary, amount)
}

// WithdrawToExchange is a paid mutator transaction binding the contract method 0x061b3245.
//
// Solidity: function withdrawToExchange(depositAccount address, amount uint256) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactor) WithdrawToExchange(opts *bind.TransactOpts, depositAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MerchantWalletContract.contract.Transact(opts, "withdrawToExchange", depositAccount, amount)
}

// WithdrawToExchange is a paid mutator transaction binding the contract method 0x061b3245.
//
// Solidity: function withdrawToExchange(depositAccount address, amount uint256) returns()
func (_MerchantWalletContract *MerchantWalletContractSession) WithdrawToExchange(depositAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.WithdrawToExchange(&_MerchantWalletContract.TransactOpts, depositAccount, amount)
}

// WithdrawToExchange is a paid mutator transaction binding the contract method 0x061b3245.
//
// Solidity: function withdrawToExchange(depositAccount address, amount uint256) returns()
func (_MerchantWalletContract *MerchantWalletContractTransactorSession) WithdrawToExchange(depositAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MerchantWalletContract.Contract.WithdrawToExchange(&_MerchantWalletContract.TransactOpts, depositAccount, amount)
}

// MerchantWalletContractMonethaAddressSetIterator is returned from FilterMonethaAddressSet and is used to iterate over the raw logs and unpacked data for MonethaAddressSet events raised by the MerchantWalletContract contract.
type MerchantWalletContractMonethaAddressSetIterator struct {
	Event *MerchantWalletContractMonethaAddressSet // Event containing the contract specifics and raw log

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
func (it *MerchantWalletContractMonethaAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantWalletContractMonethaAddressSet)
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
		it.Event = new(MerchantWalletContractMonethaAddressSet)
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
func (it *MerchantWalletContractMonethaAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantWalletContractMonethaAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantWalletContractMonethaAddressSet represents a MonethaAddressSet event raised by the MerchantWalletContract contract.
type MerchantWalletContractMonethaAddressSet struct {
	Address          common.Address
	IsMonethaAddress bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMonethaAddressSet is a free log retrieval operation binding the contract event 0xa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e692.
//
// Solidity: e MonethaAddressSet(_address address, _isMonethaAddress bool)
func (_MerchantWalletContract *MerchantWalletContractFilterer) FilterMonethaAddressSet(opts *bind.FilterOpts) (*MerchantWalletContractMonethaAddressSetIterator, error) {

	logs, sub, err := _MerchantWalletContract.contract.FilterLogs(opts, "MonethaAddressSet")
	if err != nil {
		return nil, err
	}
	return &MerchantWalletContractMonethaAddressSetIterator{contract: _MerchantWalletContract.contract, event: "MonethaAddressSet", logs: logs, sub: sub}, nil
}

// WatchMonethaAddressSet is a free log subscription operation binding the contract event 0xa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e692.
//
// Solidity: e MonethaAddressSet(_address address, _isMonethaAddress bool)
func (_MerchantWalletContract *MerchantWalletContractFilterer) WatchMonethaAddressSet(opts *bind.WatchOpts, sink chan<- *MerchantWalletContractMonethaAddressSet) (event.Subscription, error) {

	logs, sub, err := _MerchantWalletContract.contract.WatchLogs(opts, "MonethaAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantWalletContractMonethaAddressSet)
				if err := _MerchantWalletContract.contract.UnpackLog(event, "MonethaAddressSet", log); err != nil {
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

// MerchantWalletContractOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the MerchantWalletContract contract.
type MerchantWalletContractOwnershipRenouncedIterator struct {
	Event *MerchantWalletContractOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *MerchantWalletContractOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantWalletContractOwnershipRenounced)
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
		it.Event = new(MerchantWalletContractOwnershipRenounced)
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
func (it *MerchantWalletContractOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantWalletContractOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantWalletContractOwnershipRenounced represents a OwnershipRenounced event raised by the MerchantWalletContract contract.
type MerchantWalletContractOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_MerchantWalletContract *MerchantWalletContractFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*MerchantWalletContractOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _MerchantWalletContract.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MerchantWalletContractOwnershipRenouncedIterator{contract: _MerchantWalletContract.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_MerchantWalletContract *MerchantWalletContractFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *MerchantWalletContractOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _MerchantWalletContract.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantWalletContractOwnershipRenounced)
				if err := _MerchantWalletContract.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// MerchantWalletContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MerchantWalletContract contract.
type MerchantWalletContractOwnershipTransferredIterator struct {
	Event *MerchantWalletContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MerchantWalletContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantWalletContractOwnershipTransferred)
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
		it.Event = new(MerchantWalletContractOwnershipTransferred)
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
func (it *MerchantWalletContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantWalletContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantWalletContractOwnershipTransferred represents a OwnershipTransferred event raised by the MerchantWalletContract contract.
type MerchantWalletContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_MerchantWalletContract *MerchantWalletContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MerchantWalletContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MerchantWalletContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MerchantWalletContractOwnershipTransferredIterator{contract: _MerchantWalletContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_MerchantWalletContract *MerchantWalletContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MerchantWalletContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MerchantWalletContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantWalletContractOwnershipTransferred)
				if err := _MerchantWalletContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// MerchantWalletContractPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the MerchantWalletContract contract.
type MerchantWalletContractPauseIterator struct {
	Event *MerchantWalletContractPause // Event containing the contract specifics and raw log

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
func (it *MerchantWalletContractPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantWalletContractPause)
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
		it.Event = new(MerchantWalletContractPause)
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
func (it *MerchantWalletContractPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantWalletContractPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantWalletContractPause represents a Pause event raised by the MerchantWalletContract contract.
type MerchantWalletContractPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_MerchantWalletContract *MerchantWalletContractFilterer) FilterPause(opts *bind.FilterOpts) (*MerchantWalletContractPauseIterator, error) {

	logs, sub, err := _MerchantWalletContract.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &MerchantWalletContractPauseIterator{contract: _MerchantWalletContract.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_MerchantWalletContract *MerchantWalletContractFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *MerchantWalletContractPause) (event.Subscription, error) {

	logs, sub, err := _MerchantWalletContract.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantWalletContractPause)
				if err := _MerchantWalletContract.contract.UnpackLog(event, "Pause", log); err != nil {
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

// MerchantWalletContractUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the MerchantWalletContract contract.
type MerchantWalletContractUnpauseIterator struct {
	Event *MerchantWalletContractUnpause // Event containing the contract specifics and raw log

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
func (it *MerchantWalletContractUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantWalletContractUnpause)
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
		it.Event = new(MerchantWalletContractUnpause)
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
func (it *MerchantWalletContractUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantWalletContractUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantWalletContractUnpause represents a Unpause event raised by the MerchantWalletContract contract.
type MerchantWalletContractUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_MerchantWalletContract *MerchantWalletContractFilterer) FilterUnpause(opts *bind.FilterOpts) (*MerchantWalletContractUnpauseIterator, error) {

	logs, sub, err := _MerchantWalletContract.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &MerchantWalletContractUnpauseIterator{contract: _MerchantWalletContract.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_MerchantWalletContract *MerchantWalletContractFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *MerchantWalletContractUnpause) (event.Subscription, error) {

	logs, sub, err := _MerchantWalletContract.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantWalletContractUnpause)
				if err := _MerchantWalletContract.contract.UnpackLog(event, "Unpause", log); err != nil {
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

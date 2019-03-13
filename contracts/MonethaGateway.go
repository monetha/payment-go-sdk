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

// MonethaGatewayContractABI is the input ABI used to generate the binding from.
const MonethaGatewayContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isMonethaAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contactInformation\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"monethaVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MaxDiscountPermille\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"monethaVoucher\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_info\",\"type\":\"string\"}],\"name\":\"setContactInformation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMILLE_COEFFICIENT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_PERMILLE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"destroyAndSend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_monethaVault\",\"type\":\"address\"},{\"name\":\"_admin\",\"type\":\"address\"},{\"name\":\"_monethaVoucher\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"merchantWallet\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"merchantIncome\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"monethaIncome\",\"type\":\"uint256\"}],\"name\":\"PaymentProcessedEther\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"merchantWallet\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"merchantIncome\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"monethaIncome\",\"type\":\"uint256\"}],\"name\":\"PaymentProcessedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousMonethaVoucher\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newMonethaVoucher\",\"type\":\"address\"}],\"name\":\"MonethaVoucherChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"prevPermilleValue\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPermilleValue\",\"type\":\"uint256\"}],\"name\":\"MaxDiscountPermilleChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_isMonethaAddress\",\"type\":\"bool\"}],\"name\":\"MonethaAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_merchantWallet\",\"type\":\"address\"},{\"name\":\"_monethaFee\",\"type\":\"uint256\"},{\"name\":\"_customerAddress\",\"type\":\"address\"},{\"name\":\"_vouchersApply\",\"type\":\"uint256\"},{\"name\":\"_paybackPermille\",\"type\":\"uint256\"}],\"name\":\"acceptPayment\",\"outputs\":[{\"name\":\"discountWei\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_merchantWallet\",\"type\":\"address\"},{\"name\":\"_monethaFee\",\"type\":\"uint256\"},{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"acceptTokenPayment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newVault\",\"type\":\"address\"}],\"name\":\"changeMonethaVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_isMonethaAddress\",\"type\":\"bool\"}],\"name\":\"setMonethaAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_monethaVoucher\",\"type\":\"address\"}],\"name\":\"setMonethaVoucher\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maxDiscountPermille\",\"type\":\"uint256\"}],\"name\":\"setMaxDiscountPermille\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MonethaGatewayContractBin is the compiled bytecode used for deploying new contracts.
const MonethaGatewayContractBin = `0x60806040526000805460a060020a60ff02191690553480156200002157600080fd5b506040516060806200137c83398101604090815281516020830151919092015160008054600160a060020a03191633179055600160a060020a03831615156200006957600080fd5b60038054600160a060020a031916600160a060020a0385161790556200009882640100000000620000cb810204565b620000ac816401000000006200011b810204565b620000c26102bc640100000000620001a2810204565b5050506200020c565b600054600160a060020a03163314620000e357600080fd5b600160a060020a0381161515620000f957600080fd5b60048054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a031633146200013357600080fd5b600554600160a060020a038281169116146200019f57600554604051600160a060020a038084169216907fa638012f0992e8cb7e3af1700e4643463c20e5c6247964ff335411a3d369bd7590600090a360058054600160a060020a031916600160a060020a0383161790555b50565b600054600160a060020a03163314620001ba57600080fd5b6103e8811115620001ca57600080fd5b600654604080519182526020820183905280517f99f2762d865733730352be67c3db5ed72997381e375bac723590700e29f58a3e9281900390910190a1600655565b611160806200021c6000396000f30060806040526004361061013d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663102deb9c811461014257806331d413251461016557806336f7ab5e1461019a5780633f4ba83a14610224578063552548b5146102395780635c975abb14610251578063704b6c0214610266578063715018a6146102875780637adc580c1461029c578063816dbae4146102bd57806383197ef0146102ee5780638456cb59146103035780638da5cb5b14610318578063a4a876421461032d578063a7abdf0314610354578063b967a52e14610369578063c07e3391146103c2578063c4cc3f4d146103e8578063cc4fbc43146103fd578063d21c39a11461042b578063de8248fb1461044f578063f2fde38b14610464578063f5074f4114610485578063f851a440146104a6575b600080fd5b34801561014e57600080fd5b50610163600160a060020a03600435166104bb565b005b34801561017157600080fd5b50610186600160a060020a0360043516610518565b604080519115158252519081900360200190f35b3480156101a657600080fd5b506101af61052d565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101e95781810151838201526020016101d1565b50505050905090810190601f1680156102165780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561023057600080fd5b506101636105ba565b34801561024557600080fd5b50610163600435610630565b34801561025d57600080fd5b50610186610698565b34801561027257600080fd5b50610163600160a060020a03600435166106a8565b34801561029357600080fd5b50610163610703565b3480156102a857600080fd5b50610163600160a060020a036004351661076f565b3480156102c957600080fd5b506102d2610801565b60408051600160a060020a039092168252519081900360200190f35b3480156102fa57600080fd5b50610163610810565b34801561030f57600080fd5b50610163610835565b34801561032457600080fd5b506102d26108b0565b34801561033957600080fd5b506103426108bf565b60408051918252519081900360200190f35b34801561036057600080fd5b506102d26108c5565b34801561037557600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101639436949293602493928401919081908401838280828437509497506108d49650505050505050565b3480156103ce57600080fd5b50610163600160a060020a03600435166024351515610902565b3480156103f457600080fd5b50610342610994565b34801561040957600080fd5b50610163600160a060020a03600435811690602435906044351660643561099a565b610342600160a060020a036004358116906024359060443516606435608435610ba7565b34801561045b57600080fd5b50610342610f6f565b34801561047057600080fd5b50610163600160a060020a0360043516610f74565b34801561049157600080fd5b50610163600160a060020a0360043516610f94565b3480156104b257600080fd5b506102d2610fb7565b600054600160a060020a031633146104d257600080fd5b60005460a060020a900460ff16156104e957600080fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60026020526000908152604090205460ff1681565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156105b25780601f10610587576101008083540402835291602001916105b2565b820191906000526020600020905b81548152906001019060200180831161059557829003601f168201915b505050505081565b600054600160a060020a031633146105d157600080fd5b60005460a060020a900460ff1615156105e957600080fd5b6000805474ff0000000000000000000000000000000000000000191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a1565b600054600160a060020a0316331461064757600080fd5b6103e881111561065657600080fd5b600654604080519182526020820183905280517f99f2762d865733730352be67c3db5ed72997381e375bac723590700e29f58a3e9281900390910190a1600655565b60005460a060020a900460ff1681565b600054600160a060020a031633146106bf57600080fd5b600160a060020a03811615156106d457600080fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a0316331461071a57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a0316331461078657600080fd5b600554600160a060020a038281169116146107fe57600554604051600160a060020a038084169216907fa638012f0992e8cb7e3af1700e4643463c20e5c6247964ff335411a3d369bd7590600090a36005805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b600354600160a060020a031681565b600054600160a060020a0316331461082757600080fd5b600054600160a060020a0316ff5b600054600160a060020a0316331461084c57600080fd5b60005460a060020a900460ff161561086357600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a1565b600054600160a060020a031681565b60065481565b600554600160a060020a031681565b600054600160a060020a031633146108eb57600080fd5b80516108fe906001906020840190611099565b5050565b600454600160a060020a03163314806109255750600054600160a060020a031633145b151561093057600080fd5b600160a060020a038216600081815260026020908152604091829020805460ff191685151590811790915582519384529083015280517fa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e6929281900390910190a15050565b6103e881565b3360009081526002602052604081205460ff1615156109b857600080fd5b60005460a060020a900460ff16156109cf57600080fd5b600160a060020a03851615156109e457600080fd5b60008410158015610a175750610a136103e8610a07600f8563ffffffff610fc616565b9063ffffffff610ff516565b8411155b1515610a2257600080fd5b610a32828563ffffffff61100a16565b905082600160a060020a031663a9059cbb86836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050600060405180830381600087803b158015610ab057600080fd5b505af1158015610ac4573d6000803e3d6000fd5b5050600354604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a039283166004820152602481018990529051918716935063a9059cbb925060448082019260009290919082900301818387803b158015610b3657600080fd5b505af1158015610b4a573d6000803e3d6000fd5b505060408051600160a060020a038088168252891660208201528082018590526060810188905290517f618577a9afe9fd437dc6026a30af8cbadd167797636883a859979d637cfe35619350908190036080019150a15050505050565b3360009081526002602052604081205481908190819081908190819060ff161515610bd157600080fd5b60005460a060020a900460ff1615610be857600080fd5b600160a060020a038c161515610bfd57600080fd5b34955060008b10158015610c275750610c236103e8610a07600f8963ffffffff610fc616565b8b11155b1515610c3257600080fd5b60055460009750600160a060020a031615610e9357600089118015610c5957506000600654115b15610dbd57610c796103e8610a0760065489610fc690919063ffffffff16565b600554604080517fe65d1522000000000000000000000000000000000000000000000000000000008152600481018490529051929750600160a060020a039091169163e65d1522916024808201926020929091908290030181600087803b158015610ce357600080fd5b505af1158015610cf7573d6000803e3d6000fd5b505050506040513d6020811015610d0d57600080fd5b5051935088925083831115610d20578392505b600554604080517f6a2b171a000000000000000000000000000000000000000000000000000000008152600160a060020a038d81166004830152602482018790528251931692636a2b171a926044808401939192918290030181600087803b158015610d8b57600080fd5b505af1158015610d9f573d6000803e3d6000fd5b505050506040513d6040811015610db557600080fd5b506020015196505b6000881115610e9357610dec6103e8610a078a610de08a8c63ffffffff61100a16565b9063ffffffff610fc616565b91506000821115610e9357600554604080517f6daa212f000000000000000000000000000000000000000000000000000000008152600160a060020a038d811660048301526024820186905291519190921691636daa212f9160448083019260209291908290030181600087803b158015610e6657600080fd5b505af1158015610e7a573d6000803e3d6000fd5b505050506040513d6020811015610e9057600080fd5b50505b610ea3868c63ffffffff61100a16565b604051909150600160a060020a038d169082156108fc029083906000818181858888f19350505050158015610edc573d6000803e3d6000fd5b50600354604051600160a060020a03909116908c156108fc02908d906000818181858888f19350505050158015610f17573d6000803e3d6000fd5b5060408051600160a060020a038e168152602081018390528082018d905290517e79faf592f96faaec8c2d96ccc01876425fd8ef52a0837decaef1c007c5fa099181900360600190a150505050505095945050505050565b600f81565b600054600160a060020a03163314610f8b57600080fd5b6107fe8161101c565b600054600160a060020a03163314610fab57600080fd5b80600160a060020a0316ff5b600454600160a060020a031681565b6000821515610fd757506000610fef565b50818102818382811515610fe757fe5b0414610fef57fe5b92915050565b6000818381151561100257fe5b049392505050565b60008282111561101657fe5b50900390565b600160a060020a038116151561103157600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106110da57805160ff1916838001178555611107565b82800160010185558215611107579182015b828111156111075782518255916020019190600101906110ec565b50611113929150611117565b5090565b61113191905b80821115611113576000815560010161111d565b905600a165627a7a7230582043e98be71d792d8330be02cd66c5aa7c57adb06beab818727f7000c2992d4b230029`

// DeployMonethaGatewayContract deploys a new Ethereum contract, binding an instance of MonethaGatewayContract to it.
func DeployMonethaGatewayContract(auth *bind.TransactOpts, backend bind.ContractBackend, _monethaVault common.Address, _admin common.Address, _monethaVoucher common.Address) (common.Address, *types.Transaction, *MonethaGatewayContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MonethaGatewayContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MonethaGatewayContractBin), backend, _monethaVault, _admin, _monethaVoucher)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MonethaGatewayContract{MonethaGatewayContractCaller: MonethaGatewayContractCaller{contract: contract}, MonethaGatewayContractTransactor: MonethaGatewayContractTransactor{contract: contract}, MonethaGatewayContractFilterer: MonethaGatewayContractFilterer{contract: contract}}, nil
}

// MonethaGatewayContract is an auto generated Go binding around an Ethereum contract.
type MonethaGatewayContract struct {
	MonethaGatewayContractCaller     // Read-only binding to the contract
	MonethaGatewayContractTransactor // Write-only binding to the contract
	MonethaGatewayContractFilterer   // Log filterer for contract events
}

// MonethaGatewayContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MonethaGatewayContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonethaGatewayContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MonethaGatewayContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonethaGatewayContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MonethaGatewayContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonethaGatewayContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MonethaGatewayContractSession struct {
	Contract     *MonethaGatewayContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MonethaGatewayContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MonethaGatewayContractCallerSession struct {
	Contract *MonethaGatewayContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// MonethaGatewayContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MonethaGatewayContractTransactorSession struct {
	Contract     *MonethaGatewayContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// MonethaGatewayContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MonethaGatewayContractRaw struct {
	Contract *MonethaGatewayContract // Generic contract binding to access the raw methods on
}

// MonethaGatewayContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MonethaGatewayContractCallerRaw struct {
	Contract *MonethaGatewayContractCaller // Generic read-only contract binding to access the raw methods on
}

// MonethaGatewayContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MonethaGatewayContractTransactorRaw struct {
	Contract *MonethaGatewayContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMonethaGatewayContract creates a new instance of MonethaGatewayContract, bound to a specific deployed contract.
func NewMonethaGatewayContract(address common.Address, backend bind.ContractBackend) (*MonethaGatewayContract, error) {
	contract, err := bindMonethaGatewayContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MonethaGatewayContract{MonethaGatewayContractCaller: MonethaGatewayContractCaller{contract: contract}, MonethaGatewayContractTransactor: MonethaGatewayContractTransactor{contract: contract}, MonethaGatewayContractFilterer: MonethaGatewayContractFilterer{contract: contract}}, nil
}

// NewMonethaGatewayContractCaller creates a new read-only instance of MonethaGatewayContract, bound to a specific deployed contract.
func NewMonethaGatewayContractCaller(address common.Address, caller bind.ContractCaller) (*MonethaGatewayContractCaller, error) {
	contract, err := bindMonethaGatewayContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MonethaGatewayContractCaller{contract: contract}, nil
}

// NewMonethaGatewayContractTransactor creates a new write-only instance of MonethaGatewayContract, bound to a specific deployed contract.
func NewMonethaGatewayContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MonethaGatewayContractTransactor, error) {
	contract, err := bindMonethaGatewayContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MonethaGatewayContractTransactor{contract: contract}, nil
}

// NewMonethaGatewayContractFilterer creates a new log filterer instance of MonethaGatewayContract, bound to a specific deployed contract.
func NewMonethaGatewayContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MonethaGatewayContractFilterer, error) {
	contract, err := bindMonethaGatewayContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MonethaGatewayContractFilterer{contract: contract}, nil
}

// bindMonethaGatewayContract binds a generic wrapper to an already deployed contract.
func bindMonethaGatewayContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MonethaGatewayContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonethaGatewayContract *MonethaGatewayContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MonethaGatewayContract.Contract.MonethaGatewayContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonethaGatewayContract *MonethaGatewayContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.MonethaGatewayContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonethaGatewayContract *MonethaGatewayContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.MonethaGatewayContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonethaGatewayContract *MonethaGatewayContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MonethaGatewayContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonethaGatewayContract *MonethaGatewayContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonethaGatewayContract *MonethaGatewayContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.contract.Transact(opts, method, params...)
}

// FEEPERMILLE is a free data retrieval call binding the contract method 0xde8248fb.
//
// Solidity: function FEE_PERMILLE() constant returns(uint256)
func (_MonethaGatewayContract *MonethaGatewayContractCaller) FEEPERMILLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaGatewayContract.contract.Call(opts, out, "FEE_PERMILLE")
	return *ret0, err
}

// FEEPERMILLE is a free data retrieval call binding the contract method 0xde8248fb.
//
// Solidity: function FEE_PERMILLE() constant returns(uint256)
func (_MonethaGatewayContract *MonethaGatewayContractSession) FEEPERMILLE() (*big.Int, error) {
	return _MonethaGatewayContract.Contract.FEEPERMILLE(&_MonethaGatewayContract.CallOpts)
}

// FEEPERMILLE is a free data retrieval call binding the contract method 0xde8248fb.
//
// Solidity: function FEE_PERMILLE() constant returns(uint256)
func (_MonethaGatewayContract *MonethaGatewayContractCallerSession) FEEPERMILLE() (*big.Int, error) {
	return _MonethaGatewayContract.Contract.FEEPERMILLE(&_MonethaGatewayContract.CallOpts)
}

// MaxDiscountPermille is a free data retrieval call binding the contract method 0xa4a87642.
//
// Solidity: function MaxDiscountPermille() constant returns(uint256)
func (_MonethaGatewayContract *MonethaGatewayContractCaller) MaxDiscountPermille(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaGatewayContract.contract.Call(opts, out, "MaxDiscountPermille")
	return *ret0, err
}

// MaxDiscountPermille is a free data retrieval call binding the contract method 0xa4a87642.
//
// Solidity: function MaxDiscountPermille() constant returns(uint256)
func (_MonethaGatewayContract *MonethaGatewayContractSession) MaxDiscountPermille() (*big.Int, error) {
	return _MonethaGatewayContract.Contract.MaxDiscountPermille(&_MonethaGatewayContract.CallOpts)
}

// MaxDiscountPermille is a free data retrieval call binding the contract method 0xa4a87642.
//
// Solidity: function MaxDiscountPermille() constant returns(uint256)
func (_MonethaGatewayContract *MonethaGatewayContractCallerSession) MaxDiscountPermille() (*big.Int, error) {
	return _MonethaGatewayContract.Contract.MaxDiscountPermille(&_MonethaGatewayContract.CallOpts)
}

// PERMILLECOEFFICIENT is a free data retrieval call binding the contract method 0xc4cc3f4d.
//
// Solidity: function PERMILLE_COEFFICIENT() constant returns(uint256)
func (_MonethaGatewayContract *MonethaGatewayContractCaller) PERMILLECOEFFICIENT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaGatewayContract.contract.Call(opts, out, "PERMILLE_COEFFICIENT")
	return *ret0, err
}

// PERMILLECOEFFICIENT is a free data retrieval call binding the contract method 0xc4cc3f4d.
//
// Solidity: function PERMILLE_COEFFICIENT() constant returns(uint256)
func (_MonethaGatewayContract *MonethaGatewayContractSession) PERMILLECOEFFICIENT() (*big.Int, error) {
	return _MonethaGatewayContract.Contract.PERMILLECOEFFICIENT(&_MonethaGatewayContract.CallOpts)
}

// PERMILLECOEFFICIENT is a free data retrieval call binding the contract method 0xc4cc3f4d.
//
// Solidity: function PERMILLE_COEFFICIENT() constant returns(uint256)
func (_MonethaGatewayContract *MonethaGatewayContractCallerSession) PERMILLECOEFFICIENT() (*big.Int, error) {
	return _MonethaGatewayContract.Contract.PERMILLECOEFFICIENT(&_MonethaGatewayContract.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_MonethaGatewayContract *MonethaGatewayContractCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonethaGatewayContract.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_MonethaGatewayContract *MonethaGatewayContractSession) Admin() (common.Address, error) {
	return _MonethaGatewayContract.Contract.Admin(&_MonethaGatewayContract.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_MonethaGatewayContract *MonethaGatewayContractCallerSession) Admin() (common.Address, error) {
	return _MonethaGatewayContract.Contract.Admin(&_MonethaGatewayContract.CallOpts)
}

// ContactInformation is a free data retrieval call binding the contract method 0x36f7ab5e.
//
// Solidity: function contactInformation() constant returns(string)
func (_MonethaGatewayContract *MonethaGatewayContractCaller) ContactInformation(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MonethaGatewayContract.contract.Call(opts, out, "contactInformation")
	return *ret0, err
}

// ContactInformation is a free data retrieval call binding the contract method 0x36f7ab5e.
//
// Solidity: function contactInformation() constant returns(string)
func (_MonethaGatewayContract *MonethaGatewayContractSession) ContactInformation() (string, error) {
	return _MonethaGatewayContract.Contract.ContactInformation(&_MonethaGatewayContract.CallOpts)
}

// ContactInformation is a free data retrieval call binding the contract method 0x36f7ab5e.
//
// Solidity: function contactInformation() constant returns(string)
func (_MonethaGatewayContract *MonethaGatewayContractCallerSession) ContactInformation() (string, error) {
	return _MonethaGatewayContract.Contract.ContactInformation(&_MonethaGatewayContract.CallOpts)
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_MonethaGatewayContract *MonethaGatewayContractCaller) IsMonethaAddress(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MonethaGatewayContract.contract.Call(opts, out, "isMonethaAddress", arg0)
	return *ret0, err
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_MonethaGatewayContract *MonethaGatewayContractSession) IsMonethaAddress(arg0 common.Address) (bool, error) {
	return _MonethaGatewayContract.Contract.IsMonethaAddress(&_MonethaGatewayContract.CallOpts, arg0)
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_MonethaGatewayContract *MonethaGatewayContractCallerSession) IsMonethaAddress(arg0 common.Address) (bool, error) {
	return _MonethaGatewayContract.Contract.IsMonethaAddress(&_MonethaGatewayContract.CallOpts, arg0)
}

// MonethaVault is a free data retrieval call binding the contract method 0x816dbae4.
//
// Solidity: function monethaVault() constant returns(address)
func (_MonethaGatewayContract *MonethaGatewayContractCaller) MonethaVault(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonethaGatewayContract.contract.Call(opts, out, "monethaVault")
	return *ret0, err
}

// MonethaVault is a free data retrieval call binding the contract method 0x816dbae4.
//
// Solidity: function monethaVault() constant returns(address)
func (_MonethaGatewayContract *MonethaGatewayContractSession) MonethaVault() (common.Address, error) {
	return _MonethaGatewayContract.Contract.MonethaVault(&_MonethaGatewayContract.CallOpts)
}

// MonethaVault is a free data retrieval call binding the contract method 0x816dbae4.
//
// Solidity: function monethaVault() constant returns(address)
func (_MonethaGatewayContract *MonethaGatewayContractCallerSession) MonethaVault() (common.Address, error) {
	return _MonethaGatewayContract.Contract.MonethaVault(&_MonethaGatewayContract.CallOpts)
}

// MonethaVoucher is a free data retrieval call binding the contract method 0xa7abdf03.
//
// Solidity: function monethaVoucher() constant returns(address)
func (_MonethaGatewayContract *MonethaGatewayContractCaller) MonethaVoucher(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonethaGatewayContract.contract.Call(opts, out, "monethaVoucher")
	return *ret0, err
}

// MonethaVoucher is a free data retrieval call binding the contract method 0xa7abdf03.
//
// Solidity: function monethaVoucher() constant returns(address)
func (_MonethaGatewayContract *MonethaGatewayContractSession) MonethaVoucher() (common.Address, error) {
	return _MonethaGatewayContract.Contract.MonethaVoucher(&_MonethaGatewayContract.CallOpts)
}

// MonethaVoucher is a free data retrieval call binding the contract method 0xa7abdf03.
//
// Solidity: function monethaVoucher() constant returns(address)
func (_MonethaGatewayContract *MonethaGatewayContractCallerSession) MonethaVoucher() (common.Address, error) {
	return _MonethaGatewayContract.Contract.MonethaVoucher(&_MonethaGatewayContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MonethaGatewayContract *MonethaGatewayContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonethaGatewayContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MonethaGatewayContract *MonethaGatewayContractSession) Owner() (common.Address, error) {
	return _MonethaGatewayContract.Contract.Owner(&_MonethaGatewayContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MonethaGatewayContract *MonethaGatewayContractCallerSession) Owner() (common.Address, error) {
	return _MonethaGatewayContract.Contract.Owner(&_MonethaGatewayContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MonethaGatewayContract *MonethaGatewayContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MonethaGatewayContract.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MonethaGatewayContract *MonethaGatewayContractSession) Paused() (bool, error) {
	return _MonethaGatewayContract.Contract.Paused(&_MonethaGatewayContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MonethaGatewayContract *MonethaGatewayContractCallerSession) Paused() (bool, error) {
	return _MonethaGatewayContract.Contract.Paused(&_MonethaGatewayContract.CallOpts)
}

// AcceptPayment is a paid mutator transaction binding the contract method 0xd21c39a1.
//
// Solidity: function acceptPayment(_merchantWallet address, _monethaFee uint256, _customerAddress address, _vouchersApply uint256, _paybackPermille uint256) returns(discountWei uint256)
func (_MonethaGatewayContract *MonethaGatewayContractTransactor) AcceptPayment(opts *bind.TransactOpts, _merchantWallet common.Address, _monethaFee *big.Int, _customerAddress common.Address, _vouchersApply *big.Int, _paybackPermille *big.Int) (*types.Transaction, error) {
	return _MonethaGatewayContract.contract.Transact(opts, "acceptPayment", _merchantWallet, _monethaFee, _customerAddress, _vouchersApply, _paybackPermille)
}

// AcceptPayment is a paid mutator transaction binding the contract method 0xd21c39a1.
//
// Solidity: function acceptPayment(_merchantWallet address, _monethaFee uint256, _customerAddress address, _vouchersApply uint256, _paybackPermille uint256) returns(discountWei uint256)
func (_MonethaGatewayContract *MonethaGatewayContractSession) AcceptPayment(_merchantWallet common.Address, _monethaFee *big.Int, _customerAddress common.Address, _vouchersApply *big.Int, _paybackPermille *big.Int) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.AcceptPayment(&_MonethaGatewayContract.TransactOpts, _merchantWallet, _monethaFee, _customerAddress, _vouchersApply, _paybackPermille)
}

// AcceptPayment is a paid mutator transaction binding the contract method 0xd21c39a1.
//
// Solidity: function acceptPayment(_merchantWallet address, _monethaFee uint256, _customerAddress address, _vouchersApply uint256, _paybackPermille uint256) returns(discountWei uint256)
func (_MonethaGatewayContract *MonethaGatewayContractTransactorSession) AcceptPayment(_merchantWallet common.Address, _monethaFee *big.Int, _customerAddress common.Address, _vouchersApply *big.Int, _paybackPermille *big.Int) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.AcceptPayment(&_MonethaGatewayContract.TransactOpts, _merchantWallet, _monethaFee, _customerAddress, _vouchersApply, _paybackPermille)
}

// AcceptTokenPayment is a paid mutator transaction binding the contract method 0xcc4fbc43.
//
// Solidity: function acceptTokenPayment(_merchantWallet address, _monethaFee uint256, _tokenAddress address, _value uint256) returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactor) AcceptTokenPayment(opts *bind.TransactOpts, _merchantWallet common.Address, _monethaFee *big.Int, _tokenAddress common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaGatewayContract.contract.Transact(opts, "acceptTokenPayment", _merchantWallet, _monethaFee, _tokenAddress, _value)
}

// AcceptTokenPayment is a paid mutator transaction binding the contract method 0xcc4fbc43.
//
// Solidity: function acceptTokenPayment(_merchantWallet address, _monethaFee uint256, _tokenAddress address, _value uint256) returns()
func (_MonethaGatewayContract *MonethaGatewayContractSession) AcceptTokenPayment(_merchantWallet common.Address, _monethaFee *big.Int, _tokenAddress common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.AcceptTokenPayment(&_MonethaGatewayContract.TransactOpts, _merchantWallet, _monethaFee, _tokenAddress, _value)
}

// AcceptTokenPayment is a paid mutator transaction binding the contract method 0xcc4fbc43.
//
// Solidity: function acceptTokenPayment(_merchantWallet address, _monethaFee uint256, _tokenAddress address, _value uint256) returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactorSession) AcceptTokenPayment(_merchantWallet common.Address, _monethaFee *big.Int, _tokenAddress common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.AcceptTokenPayment(&_MonethaGatewayContract.TransactOpts, _merchantWallet, _monethaFee, _tokenAddress, _value)
}

// ChangeMonethaVault is a paid mutator transaction binding the contract method 0x102deb9c.
//
// Solidity: function changeMonethaVault(newVault address) returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactor) ChangeMonethaVault(opts *bind.TransactOpts, newVault common.Address) (*types.Transaction, error) {
	return _MonethaGatewayContract.contract.Transact(opts, "changeMonethaVault", newVault)
}

// ChangeMonethaVault is a paid mutator transaction binding the contract method 0x102deb9c.
//
// Solidity: function changeMonethaVault(newVault address) returns()
func (_MonethaGatewayContract *MonethaGatewayContractSession) ChangeMonethaVault(newVault common.Address) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.ChangeMonethaVault(&_MonethaGatewayContract.TransactOpts, newVault)
}

// ChangeMonethaVault is a paid mutator transaction binding the contract method 0x102deb9c.
//
// Solidity: function changeMonethaVault(newVault address) returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactorSession) ChangeMonethaVault(newVault common.Address) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.ChangeMonethaVault(&_MonethaGatewayContract.TransactOpts, newVault)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaGatewayContract.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_MonethaGatewayContract *MonethaGatewayContractSession) Destroy() (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.Destroy(&_MonethaGatewayContract.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactorSession) Destroy() (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.Destroy(&_MonethaGatewayContract.TransactOpts)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactor) DestroyAndSend(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _MonethaGatewayContract.contract.Transact(opts, "destroyAndSend", _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_MonethaGatewayContract *MonethaGatewayContractSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.DestroyAndSend(&_MonethaGatewayContract.TransactOpts, _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactorSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.DestroyAndSend(&_MonethaGatewayContract.TransactOpts, _recipient)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaGatewayContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonethaGatewayContract *MonethaGatewayContractSession) Pause() (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.Pause(&_MonethaGatewayContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactorSession) Pause() (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.Pause(&_MonethaGatewayContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaGatewayContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MonethaGatewayContract *MonethaGatewayContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.RenounceOwnership(&_MonethaGatewayContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.RenounceOwnership(&_MonethaGatewayContract.TransactOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_admin address) returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _MonethaGatewayContract.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_admin address) returns()
func (_MonethaGatewayContract *MonethaGatewayContractSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.SetAdmin(&_MonethaGatewayContract.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_admin address) returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.SetAdmin(&_MonethaGatewayContract.TransactOpts, _admin)
}

// SetContactInformation is a paid mutator transaction binding the contract method 0xb967a52e.
//
// Solidity: function setContactInformation(_info string) returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactor) SetContactInformation(opts *bind.TransactOpts, _info string) (*types.Transaction, error) {
	return _MonethaGatewayContract.contract.Transact(opts, "setContactInformation", _info)
}

// SetContactInformation is a paid mutator transaction binding the contract method 0xb967a52e.
//
// Solidity: function setContactInformation(_info string) returns()
func (_MonethaGatewayContract *MonethaGatewayContractSession) SetContactInformation(_info string) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.SetContactInformation(&_MonethaGatewayContract.TransactOpts, _info)
}

// SetContactInformation is a paid mutator transaction binding the contract method 0xb967a52e.
//
// Solidity: function setContactInformation(_info string) returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactorSession) SetContactInformation(_info string) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.SetContactInformation(&_MonethaGatewayContract.TransactOpts, _info)
}

// SetMaxDiscountPermille is a paid mutator transaction binding the contract method 0x552548b5.
//
// Solidity: function setMaxDiscountPermille(_maxDiscountPermille uint256) returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactor) SetMaxDiscountPermille(opts *bind.TransactOpts, _maxDiscountPermille *big.Int) (*types.Transaction, error) {
	return _MonethaGatewayContract.contract.Transact(opts, "setMaxDiscountPermille", _maxDiscountPermille)
}

// SetMaxDiscountPermille is a paid mutator transaction binding the contract method 0x552548b5.
//
// Solidity: function setMaxDiscountPermille(_maxDiscountPermille uint256) returns()
func (_MonethaGatewayContract *MonethaGatewayContractSession) SetMaxDiscountPermille(_maxDiscountPermille *big.Int) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.SetMaxDiscountPermille(&_MonethaGatewayContract.TransactOpts, _maxDiscountPermille)
}

// SetMaxDiscountPermille is a paid mutator transaction binding the contract method 0x552548b5.
//
// Solidity: function setMaxDiscountPermille(_maxDiscountPermille uint256) returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactorSession) SetMaxDiscountPermille(_maxDiscountPermille *big.Int) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.SetMaxDiscountPermille(&_MonethaGatewayContract.TransactOpts, _maxDiscountPermille)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactor) SetMonethaAddress(opts *bind.TransactOpts, _address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _MonethaGatewayContract.contract.Transact(opts, "setMonethaAddress", _address, _isMonethaAddress)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_MonethaGatewayContract *MonethaGatewayContractSession) SetMonethaAddress(_address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.SetMonethaAddress(&_MonethaGatewayContract.TransactOpts, _address, _isMonethaAddress)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactorSession) SetMonethaAddress(_address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.SetMonethaAddress(&_MonethaGatewayContract.TransactOpts, _address, _isMonethaAddress)
}

// SetMonethaVoucher is a paid mutator transaction binding the contract method 0x7adc580c.
//
// Solidity: function setMonethaVoucher(_monethaVoucher address) returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactor) SetMonethaVoucher(opts *bind.TransactOpts, _monethaVoucher common.Address) (*types.Transaction, error) {
	return _MonethaGatewayContract.contract.Transact(opts, "setMonethaVoucher", _monethaVoucher)
}

// SetMonethaVoucher is a paid mutator transaction binding the contract method 0x7adc580c.
//
// Solidity: function setMonethaVoucher(_monethaVoucher address) returns()
func (_MonethaGatewayContract *MonethaGatewayContractSession) SetMonethaVoucher(_monethaVoucher common.Address) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.SetMonethaVoucher(&_MonethaGatewayContract.TransactOpts, _monethaVoucher)
}

// SetMonethaVoucher is a paid mutator transaction binding the contract method 0x7adc580c.
//
// Solidity: function setMonethaVoucher(_monethaVoucher address) returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactorSession) SetMonethaVoucher(_monethaVoucher common.Address) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.SetMonethaVoucher(&_MonethaGatewayContract.TransactOpts, _monethaVoucher)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _MonethaGatewayContract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MonethaGatewayContract *MonethaGatewayContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.TransferOwnership(&_MonethaGatewayContract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.TransferOwnership(&_MonethaGatewayContract.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaGatewayContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonethaGatewayContract *MonethaGatewayContractSession) Unpause() (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.Unpause(&_MonethaGatewayContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonethaGatewayContract *MonethaGatewayContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _MonethaGatewayContract.Contract.Unpause(&_MonethaGatewayContract.TransactOpts)
}

// MonethaGatewayContractMaxDiscountPermilleChangedIterator is returned from FilterMaxDiscountPermilleChanged and is used to iterate over the raw logs and unpacked data for MaxDiscountPermilleChanged events raised by the MonethaGatewayContract contract.
type MonethaGatewayContractMaxDiscountPermilleChangedIterator struct {
	Event *MonethaGatewayContractMaxDiscountPermilleChanged // Event containing the contract specifics and raw log

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
func (it *MonethaGatewayContractMaxDiscountPermilleChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaGatewayContractMaxDiscountPermilleChanged)
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
		it.Event = new(MonethaGatewayContractMaxDiscountPermilleChanged)
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
func (it *MonethaGatewayContractMaxDiscountPermilleChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaGatewayContractMaxDiscountPermilleChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaGatewayContractMaxDiscountPermilleChanged represents a MaxDiscountPermilleChanged event raised by the MonethaGatewayContract contract.
type MonethaGatewayContractMaxDiscountPermilleChanged struct {
	PrevPermilleValue *big.Int
	NewPermilleValue  *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMaxDiscountPermilleChanged is a free log retrieval operation binding the contract event 0x99f2762d865733730352be67c3db5ed72997381e375bac723590700e29f58a3e.
//
// Solidity: e MaxDiscountPermilleChanged(prevPermilleValue uint256, newPermilleValue uint256)
func (_MonethaGatewayContract *MonethaGatewayContractFilterer) FilterMaxDiscountPermilleChanged(opts *bind.FilterOpts) (*MonethaGatewayContractMaxDiscountPermilleChangedIterator, error) {

	logs, sub, err := _MonethaGatewayContract.contract.FilterLogs(opts, "MaxDiscountPermilleChanged")
	if err != nil {
		return nil, err
	}
	return &MonethaGatewayContractMaxDiscountPermilleChangedIterator{contract: _MonethaGatewayContract.contract, event: "MaxDiscountPermilleChanged", logs: logs, sub: sub}, nil
}

// WatchMaxDiscountPermilleChanged is a free log subscription operation binding the contract event 0x99f2762d865733730352be67c3db5ed72997381e375bac723590700e29f58a3e.
//
// Solidity: e MaxDiscountPermilleChanged(prevPermilleValue uint256, newPermilleValue uint256)
func (_MonethaGatewayContract *MonethaGatewayContractFilterer) WatchMaxDiscountPermilleChanged(opts *bind.WatchOpts, sink chan<- *MonethaGatewayContractMaxDiscountPermilleChanged) (event.Subscription, error) {

	logs, sub, err := _MonethaGatewayContract.contract.WatchLogs(opts, "MaxDiscountPermilleChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaGatewayContractMaxDiscountPermilleChanged)
				if err := _MonethaGatewayContract.contract.UnpackLog(event, "MaxDiscountPermilleChanged", log); err != nil {
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

// MonethaGatewayContractMonethaAddressSetIterator is returned from FilterMonethaAddressSet and is used to iterate over the raw logs and unpacked data for MonethaAddressSet events raised by the MonethaGatewayContract contract.
type MonethaGatewayContractMonethaAddressSetIterator struct {
	Event *MonethaGatewayContractMonethaAddressSet // Event containing the contract specifics and raw log

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
func (it *MonethaGatewayContractMonethaAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaGatewayContractMonethaAddressSet)
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
		it.Event = new(MonethaGatewayContractMonethaAddressSet)
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
func (it *MonethaGatewayContractMonethaAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaGatewayContractMonethaAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaGatewayContractMonethaAddressSet represents a MonethaAddressSet event raised by the MonethaGatewayContract contract.
type MonethaGatewayContractMonethaAddressSet struct {
	Address          common.Address
	IsMonethaAddress bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMonethaAddressSet is a free log retrieval operation binding the contract event 0xa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e692.
//
// Solidity: e MonethaAddressSet(_address address, _isMonethaAddress bool)
func (_MonethaGatewayContract *MonethaGatewayContractFilterer) FilterMonethaAddressSet(opts *bind.FilterOpts) (*MonethaGatewayContractMonethaAddressSetIterator, error) {

	logs, sub, err := _MonethaGatewayContract.contract.FilterLogs(opts, "MonethaAddressSet")
	if err != nil {
		return nil, err
	}
	return &MonethaGatewayContractMonethaAddressSetIterator{contract: _MonethaGatewayContract.contract, event: "MonethaAddressSet", logs: logs, sub: sub}, nil
}

// WatchMonethaAddressSet is a free log subscription operation binding the contract event 0xa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e692.
//
// Solidity: e MonethaAddressSet(_address address, _isMonethaAddress bool)
func (_MonethaGatewayContract *MonethaGatewayContractFilterer) WatchMonethaAddressSet(opts *bind.WatchOpts, sink chan<- *MonethaGatewayContractMonethaAddressSet) (event.Subscription, error) {

	logs, sub, err := _MonethaGatewayContract.contract.WatchLogs(opts, "MonethaAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaGatewayContractMonethaAddressSet)
				if err := _MonethaGatewayContract.contract.UnpackLog(event, "MonethaAddressSet", log); err != nil {
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

// MonethaGatewayContractMonethaVoucherChangedIterator is returned from FilterMonethaVoucherChanged and is used to iterate over the raw logs and unpacked data for MonethaVoucherChanged events raised by the MonethaGatewayContract contract.
type MonethaGatewayContractMonethaVoucherChangedIterator struct {
	Event *MonethaGatewayContractMonethaVoucherChanged // Event containing the contract specifics and raw log

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
func (it *MonethaGatewayContractMonethaVoucherChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaGatewayContractMonethaVoucherChanged)
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
		it.Event = new(MonethaGatewayContractMonethaVoucherChanged)
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
func (it *MonethaGatewayContractMonethaVoucherChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaGatewayContractMonethaVoucherChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaGatewayContractMonethaVoucherChanged represents a MonethaVoucherChanged event raised by the MonethaGatewayContract contract.
type MonethaGatewayContractMonethaVoucherChanged struct {
	PreviousMonethaVoucher common.Address
	NewMonethaVoucher      common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterMonethaVoucherChanged is a free log retrieval operation binding the contract event 0xa638012f0992e8cb7e3af1700e4643463c20e5c6247964ff335411a3d369bd75.
//
// Solidity: e MonethaVoucherChanged(previousMonethaVoucher indexed address, newMonethaVoucher indexed address)
func (_MonethaGatewayContract *MonethaGatewayContractFilterer) FilterMonethaVoucherChanged(opts *bind.FilterOpts, previousMonethaVoucher []common.Address, newMonethaVoucher []common.Address) (*MonethaGatewayContractMonethaVoucherChangedIterator, error) {

	var previousMonethaVoucherRule []interface{}
	for _, previousMonethaVoucherItem := range previousMonethaVoucher {
		previousMonethaVoucherRule = append(previousMonethaVoucherRule, previousMonethaVoucherItem)
	}
	var newMonethaVoucherRule []interface{}
	for _, newMonethaVoucherItem := range newMonethaVoucher {
		newMonethaVoucherRule = append(newMonethaVoucherRule, newMonethaVoucherItem)
	}

	logs, sub, err := _MonethaGatewayContract.contract.FilterLogs(opts, "MonethaVoucherChanged", previousMonethaVoucherRule, newMonethaVoucherRule)
	if err != nil {
		return nil, err
	}
	return &MonethaGatewayContractMonethaVoucherChangedIterator{contract: _MonethaGatewayContract.contract, event: "MonethaVoucherChanged", logs: logs, sub: sub}, nil
}

// WatchMonethaVoucherChanged is a free log subscription operation binding the contract event 0xa638012f0992e8cb7e3af1700e4643463c20e5c6247964ff335411a3d369bd75.
//
// Solidity: e MonethaVoucherChanged(previousMonethaVoucher indexed address, newMonethaVoucher indexed address)
func (_MonethaGatewayContract *MonethaGatewayContractFilterer) WatchMonethaVoucherChanged(opts *bind.WatchOpts, sink chan<- *MonethaGatewayContractMonethaVoucherChanged, previousMonethaVoucher []common.Address, newMonethaVoucher []common.Address) (event.Subscription, error) {

	var previousMonethaVoucherRule []interface{}
	for _, previousMonethaVoucherItem := range previousMonethaVoucher {
		previousMonethaVoucherRule = append(previousMonethaVoucherRule, previousMonethaVoucherItem)
	}
	var newMonethaVoucherRule []interface{}
	for _, newMonethaVoucherItem := range newMonethaVoucher {
		newMonethaVoucherRule = append(newMonethaVoucherRule, newMonethaVoucherItem)
	}

	logs, sub, err := _MonethaGatewayContract.contract.WatchLogs(opts, "MonethaVoucherChanged", previousMonethaVoucherRule, newMonethaVoucherRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaGatewayContractMonethaVoucherChanged)
				if err := _MonethaGatewayContract.contract.UnpackLog(event, "MonethaVoucherChanged", log); err != nil {
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

// MonethaGatewayContractOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the MonethaGatewayContract contract.
type MonethaGatewayContractOwnershipRenouncedIterator struct {
	Event *MonethaGatewayContractOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *MonethaGatewayContractOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaGatewayContractOwnershipRenounced)
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
		it.Event = new(MonethaGatewayContractOwnershipRenounced)
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
func (it *MonethaGatewayContractOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaGatewayContractOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaGatewayContractOwnershipRenounced represents a OwnershipRenounced event raised by the MonethaGatewayContract contract.
type MonethaGatewayContractOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_MonethaGatewayContract *MonethaGatewayContractFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*MonethaGatewayContractOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _MonethaGatewayContract.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MonethaGatewayContractOwnershipRenouncedIterator{contract: _MonethaGatewayContract.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_MonethaGatewayContract *MonethaGatewayContractFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *MonethaGatewayContractOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _MonethaGatewayContract.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaGatewayContractOwnershipRenounced)
				if err := _MonethaGatewayContract.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// MonethaGatewayContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MonethaGatewayContract contract.
type MonethaGatewayContractOwnershipTransferredIterator struct {
	Event *MonethaGatewayContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MonethaGatewayContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaGatewayContractOwnershipTransferred)
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
		it.Event = new(MonethaGatewayContractOwnershipTransferred)
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
func (it *MonethaGatewayContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaGatewayContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaGatewayContractOwnershipTransferred represents a OwnershipTransferred event raised by the MonethaGatewayContract contract.
type MonethaGatewayContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_MonethaGatewayContract *MonethaGatewayContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MonethaGatewayContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MonethaGatewayContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MonethaGatewayContractOwnershipTransferredIterator{contract: _MonethaGatewayContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_MonethaGatewayContract *MonethaGatewayContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MonethaGatewayContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MonethaGatewayContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaGatewayContractOwnershipTransferred)
				if err := _MonethaGatewayContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// MonethaGatewayContractPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the MonethaGatewayContract contract.
type MonethaGatewayContractPauseIterator struct {
	Event *MonethaGatewayContractPause // Event containing the contract specifics and raw log

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
func (it *MonethaGatewayContractPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaGatewayContractPause)
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
		it.Event = new(MonethaGatewayContractPause)
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
func (it *MonethaGatewayContractPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaGatewayContractPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaGatewayContractPause represents a Pause event raised by the MonethaGatewayContract contract.
type MonethaGatewayContractPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_MonethaGatewayContract *MonethaGatewayContractFilterer) FilterPause(opts *bind.FilterOpts) (*MonethaGatewayContractPauseIterator, error) {

	logs, sub, err := _MonethaGatewayContract.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &MonethaGatewayContractPauseIterator{contract: _MonethaGatewayContract.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_MonethaGatewayContract *MonethaGatewayContractFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *MonethaGatewayContractPause) (event.Subscription, error) {

	logs, sub, err := _MonethaGatewayContract.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaGatewayContractPause)
				if err := _MonethaGatewayContract.contract.UnpackLog(event, "Pause", log); err != nil {
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

// MonethaGatewayContractPaymentProcessedEtherIterator is returned from FilterPaymentProcessedEther and is used to iterate over the raw logs and unpacked data for PaymentProcessedEther events raised by the MonethaGatewayContract contract.
type MonethaGatewayContractPaymentProcessedEtherIterator struct {
	Event *MonethaGatewayContractPaymentProcessedEther // Event containing the contract specifics and raw log

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
func (it *MonethaGatewayContractPaymentProcessedEtherIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaGatewayContractPaymentProcessedEther)
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
		it.Event = new(MonethaGatewayContractPaymentProcessedEther)
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
func (it *MonethaGatewayContractPaymentProcessedEtherIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaGatewayContractPaymentProcessedEtherIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaGatewayContractPaymentProcessedEther represents a PaymentProcessedEther event raised by the MonethaGatewayContract contract.
type MonethaGatewayContractPaymentProcessedEther struct {
	MerchantWallet common.Address
	MerchantIncome *big.Int
	MonethaIncome  *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPaymentProcessedEther is a free log retrieval operation binding the contract event 0x0079faf592f96faaec8c2d96ccc01876425fd8ef52a0837decaef1c007c5fa09.
//
// Solidity: e PaymentProcessedEther(merchantWallet address, merchantIncome uint256, monethaIncome uint256)
func (_MonethaGatewayContract *MonethaGatewayContractFilterer) FilterPaymentProcessedEther(opts *bind.FilterOpts) (*MonethaGatewayContractPaymentProcessedEtherIterator, error) {

	logs, sub, err := _MonethaGatewayContract.contract.FilterLogs(opts, "PaymentProcessedEther")
	if err != nil {
		return nil, err
	}
	return &MonethaGatewayContractPaymentProcessedEtherIterator{contract: _MonethaGatewayContract.contract, event: "PaymentProcessedEther", logs: logs, sub: sub}, nil
}

// WatchPaymentProcessedEther is a free log subscription operation binding the contract event 0x0079faf592f96faaec8c2d96ccc01876425fd8ef52a0837decaef1c007c5fa09.
//
// Solidity: e PaymentProcessedEther(merchantWallet address, merchantIncome uint256, monethaIncome uint256)
func (_MonethaGatewayContract *MonethaGatewayContractFilterer) WatchPaymentProcessedEther(opts *bind.WatchOpts, sink chan<- *MonethaGatewayContractPaymentProcessedEther) (event.Subscription, error) {

	logs, sub, err := _MonethaGatewayContract.contract.WatchLogs(opts, "PaymentProcessedEther")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaGatewayContractPaymentProcessedEther)
				if err := _MonethaGatewayContract.contract.UnpackLog(event, "PaymentProcessedEther", log); err != nil {
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

// MonethaGatewayContractPaymentProcessedTokenIterator is returned from FilterPaymentProcessedToken and is used to iterate over the raw logs and unpacked data for PaymentProcessedToken events raised by the MonethaGatewayContract contract.
type MonethaGatewayContractPaymentProcessedTokenIterator struct {
	Event *MonethaGatewayContractPaymentProcessedToken // Event containing the contract specifics and raw log

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
func (it *MonethaGatewayContractPaymentProcessedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaGatewayContractPaymentProcessedToken)
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
		it.Event = new(MonethaGatewayContractPaymentProcessedToken)
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
func (it *MonethaGatewayContractPaymentProcessedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaGatewayContractPaymentProcessedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaGatewayContractPaymentProcessedToken represents a PaymentProcessedToken event raised by the MonethaGatewayContract contract.
type MonethaGatewayContractPaymentProcessedToken struct {
	TokenAddress   common.Address
	MerchantWallet common.Address
	MerchantIncome *big.Int
	MonethaIncome  *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPaymentProcessedToken is a free log retrieval operation binding the contract event 0x618577a9afe9fd437dc6026a30af8cbadd167797636883a859979d637cfe3561.
//
// Solidity: e PaymentProcessedToken(tokenAddress address, merchantWallet address, merchantIncome uint256, monethaIncome uint256)
func (_MonethaGatewayContract *MonethaGatewayContractFilterer) FilterPaymentProcessedToken(opts *bind.FilterOpts) (*MonethaGatewayContractPaymentProcessedTokenIterator, error) {

	logs, sub, err := _MonethaGatewayContract.contract.FilterLogs(opts, "PaymentProcessedToken")
	if err != nil {
		return nil, err
	}
	return &MonethaGatewayContractPaymentProcessedTokenIterator{contract: _MonethaGatewayContract.contract, event: "PaymentProcessedToken", logs: logs, sub: sub}, nil
}

// WatchPaymentProcessedToken is a free log subscription operation binding the contract event 0x618577a9afe9fd437dc6026a30af8cbadd167797636883a859979d637cfe3561.
//
// Solidity: e PaymentProcessedToken(tokenAddress address, merchantWallet address, merchantIncome uint256, monethaIncome uint256)
func (_MonethaGatewayContract *MonethaGatewayContractFilterer) WatchPaymentProcessedToken(opts *bind.WatchOpts, sink chan<- *MonethaGatewayContractPaymentProcessedToken) (event.Subscription, error) {

	logs, sub, err := _MonethaGatewayContract.contract.WatchLogs(opts, "PaymentProcessedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaGatewayContractPaymentProcessedToken)
				if err := _MonethaGatewayContract.contract.UnpackLog(event, "PaymentProcessedToken", log); err != nil {
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

// MonethaGatewayContractUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the MonethaGatewayContract contract.
type MonethaGatewayContractUnpauseIterator struct {
	Event *MonethaGatewayContractUnpause // Event containing the contract specifics and raw log

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
func (it *MonethaGatewayContractUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaGatewayContractUnpause)
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
		it.Event = new(MonethaGatewayContractUnpause)
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
func (it *MonethaGatewayContractUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaGatewayContractUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaGatewayContractUnpause represents a Unpause event raised by the MonethaGatewayContract contract.
type MonethaGatewayContractUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_MonethaGatewayContract *MonethaGatewayContractFilterer) FilterUnpause(opts *bind.FilterOpts) (*MonethaGatewayContractUnpauseIterator, error) {

	logs, sub, err := _MonethaGatewayContract.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &MonethaGatewayContractUnpauseIterator{contract: _MonethaGatewayContract.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_MonethaGatewayContract *MonethaGatewayContractFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *MonethaGatewayContractUnpause) (event.Subscription, error) {

	logs, sub, err := _MonethaGatewayContract.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaGatewayContractUnpause)
				if err := _MonethaGatewayContract.contract.UnpackLog(event, "Unpause", log); err != nil {
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

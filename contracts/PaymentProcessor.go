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

// PaymentProcessorContractABI is the input ABI used to generate the binding from.
const PaymentProcessorContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isMonethaAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contactInformation\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"merchantHistory\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PAYBACK_PERMILLE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"monethaGateway\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"name\":\"state\",\"type\":\"uint8\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"paymentAcceptor\",\"type\":\"address\"},{\"name\":\"originAddress\",\"type\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"vouchersApply\",\"type\":\"uint256\"},{\"name\":\"discount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_info\",\"type\":\"string\"}],\"name\":\"setContactInformation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"merchantWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_isMonethaAddress\",\"type\":\"bool\"}],\"name\":\"setMonethaAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMILLE_COEFFICIENT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_PERMILLE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"merchantIdHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"destroyAndSend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_merchantId\",\"type\":\"string\"},{\"name\":\"_merchantHistory\",\"type\":\"address\"},{\"name\":\"_monethaGateway\",\"type\":\"address\"},{\"name\":\"_merchantWallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_isMonethaAddress\",\"type\":\"bool\"}],\"name\":\"MonethaAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_paymentAcceptor\",\"type\":\"address\"},{\"name\":\"_originAddress\",\"type\":\"address\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_vouchersApply\",\"type\":\"uint256\"}],\"name\":\"addOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"securePay\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"secureTokenPay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\"},{\"name\":\"_clientReputation\",\"type\":\"uint32\"},{\"name\":\"_merchantReputation\",\"type\":\"uint32\"},{\"name\":\"_dealHash\",\"type\":\"uint256\"},{\"name\":\"_cancelReason\",\"type\":\"string\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\"},{\"name\":\"_clientReputation\",\"type\":\"uint32\"},{\"name\":\"_merchantReputation\",\"type\":\"uint32\"},{\"name\":\"_dealHash\",\"type\":\"uint256\"},{\"name\":\"_refundReason\",\"type\":\"string\"}],\"name\":\"refundPayment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"withdrawRefund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"withdrawTokenRefund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\"},{\"name\":\"_clientReputation\",\"type\":\"uint32\"},{\"name\":\"_merchantReputation\",\"type\":\"uint32\"},{\"name\":\"_dealHash\",\"type\":\"uint256\"}],\"name\":\"processPayment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newGateway\",\"type\":\"address\"}],\"name\":\"setMonethaGateway\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newWallet\",\"type\":\"address\"}],\"name\":\"setMerchantWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_merchantHistory\",\"type\":\"address\"}],\"name\":\"setMerchantDealsHistory\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PaymentProcessorContractBin is the compiled bytecode used for deploying new contracts.
const PaymentProcessorContractBin = `0x60806040526000805460a060020a60ff02191690553480156200002157600080fd5b506040516200223b3803806200223b8339810160409081528151602083015191830151606084015160008054600160a060020a0319163317815592909401805190949192106200007057600080fd5b836040516020018082805190602001908083835b60208310620000a55780518252601f19909201916020918201910162000084565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b602083106200010a5780518252601f199092019160209182019101620000e9565b5181516020939093036101000a60001901801990911692169190911790526040519201829003909120600655506200014f915083905064010000000062000181810204565b6200016381640100000000620001d1810204565b6200017783640100000000620002b9810204565b50505050620003a1565b600054600160a060020a031633146200019957600080fd5b600160a060020a0381161515620001af57600080fd5b60038054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a03163314620001e957600080fd5b600160a060020a0381161515620001ff57600080fd5b6006546000191681600160a060020a031663f0daba016040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156200025e57600080fd5b505af115801562000273573d6000803e3d6000fd5b505050506040513d60208110156200028a57600080fd5b5051146200029757600080fd5b60058054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a03163314620002d157600080fd5b600160a060020a0381161515620002e757600080fd5b6006546000191681600160a060020a031663f0daba016040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156200034657600080fd5b505af11580156200035b573d6000803e3d6000fd5b505050506040513d60208110156200037257600080fd5b5051146200037f57600080fd5b60048054600160a060020a031916600160a060020a0392909216919091179055565b611e8a80620003b16000396000f3006080604052600436106101715763ffffffff60e060020a60003504166331d41325811461017657806336f7ab5e146101ab5780633d17a2d8146102355780633f4ba83a146102665780634153d65b1461027d578063444b6048146102885780634baf43e4146102a95780635c975abb146102d05780636137412c146102e5578063715018a6146102fa57806383197ef01461030f5780638456cb59146103245780638467d9cf146103395780638b11fb3e146103735780638da5cb5b1461038b5780639d153495146103a0578063a85c38ef146103b8578063aab9f16514610430578063ac1a13fb1461045a578063b440bf3914610494578063b967a52e146104b5578063bc85e0641461050e578063c07e339114610523578063c48aab6d14610549578063c4cc3f4d14610583578063ddda66db14610598578063de8248fb146105b9578063e5b8d6e0146105ce578063f0daba01146105e6578063f2fde38b146105fb578063f5074f411461061c575b600080fd5b34801561018257600080fd5b50610197600160a060020a036004351661063d565b604080519115158252519081900360200190f35b3480156101b757600080fd5b506101c0610652565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101fa5781810151838201526020016101e2565b50505050905090810190601f1680156102275780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561024157600080fd5b5061024a6106df565b60408051600160a060020a039092168252519081900360200190f35b34801561027257600080fd5b5061027b6106ee565b005b61027b600435610764565b34801561029457600080fd5b5061027b600160a060020a036004351661083d565b3480156102b557600080fd5b506102be610913565b60408051918252519081900360200190f35b3480156102dc57600080fd5b50610197610918565b3480156102f157600080fd5b5061024a610928565b34801561030657600080fd5b5061027b610937565b34801561031b57600080fd5b5061027b6109a3565b34801561033057600080fd5b5061027b6109c8565b34801561034557600080fd5b5061027b600480359063ffffffff602480358216926044359092169160643591608435908101910135610a43565b34801561037f57600080fd5b5061027b600435610c0f565b34801561039757600080fd5b5061024a610d72565b3480156103ac57600080fd5b5061027b600435610d81565b3480156103c457600080fd5b506103d0600435610e7a565b604051808960068111156103e057fe5b60ff168152602081019890985250604080880196909652600160a060020a0394851660608801529284166080870152921660a085015260c084019190915260e08301525190819003610100019150f35b34801561043c57600080fd5b5061027b60043563ffffffff60243581169060443516606435610ed2565b34801561046657600080fd5b5061027b600480359063ffffffff60248035821692604435909216916064359160843590810191013561146a565b3480156104a057600080fd5b5061027b600160a060020a03600435166115e8565b3480156104c157600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261027b9436949293602493928401919081908401838280828437509497506116439650505050505050565b34801561051a57600080fd5b5061024a611671565b34801561052f57600080fd5b5061027b600160a060020a03600435166024351515611680565b34801561055557600080fd5b5061027b600435602435600160a060020a036044358116906064358116906084359060a4351660c4356116fb565b34801561058f57600080fd5b506102be61190e565b3480156105a457600080fd5b5061027b600160a060020a0360043516611914565b3480156105c557600080fd5b506102be6119ea565b3480156105da57600080fd5b5061027b6004356119ef565b3480156105f257600080fd5b506102be611b34565b34801561060757600080fd5b5061027b600160a060020a0360043516611b3a565b34801561062857600080fd5b5061027b600160a060020a0360043516611b5d565b60026020526000908152604090205460ff1681565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106d75780601f106106ac576101008083540402835291602001916106d7565b820191906000526020600020905b8154815290600101906020018083116106ba57829003601f168201915b505050505081565b600454600160a060020a031681565b600054600160a060020a0316331461070557600080fd5b60005460a060020a900460ff16151561071d57600080fd5b6000805474ff0000000000000000000000000000000000000000191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a1565b6000805460a060020a900460ff161561077c57600080fd5b600082815260076020526040902054829060019060ff16600681111561079e57fe5b8160068111156107aa57fe5b146107b457600080fd5b600084815260076020526040902060058101549093508490600290600160a060020a0316156107e257600080fd5b6003850154600160a060020a031633146107fb57600080fd5b6001850154341461080b57600080fd5b6000828152600760205260409020805482919060ff1916600183600681111561083057fe5b0217905550505050505050565b600054600160a060020a0316331461085457600080fd5b600160a060020a038116151561086957600080fd5b6006546000191681600160a060020a031663f0daba016040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156108ae57600080fd5b505af11580156108c2573d6000803e3d6000fd5b505050506040513d60208110156108d857600080fd5b5051146108e457600080fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600281565b60005460a060020a900460ff1681565b600354600160a060020a031681565b600054600160a060020a0316331461094e57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031633146109ba57600080fd5b600054600160a060020a0316ff5b600054600160a060020a031633146109df57600080fd5b60005460a060020a900460ff16156109f657600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a1565b3360009081526002602052604081205460ff161515610a6157600080fd5b60005460a060020a900460ff1615610a7857600080fd5b600087815260076020526040902054879060029060ff166006811115610a9a57fe5b816006811115610aa657fe5b14610ab057600080fd5b88600460008611610ac057600080fd5b60008b81526007602052604081209550610ae0908c908c908c908c611b80565b600460009054906101000a9004600160a060020a0316600160a060020a031663188c668c8c8760040160009054906101000a9004600160a060020a03168d8d8d8d8d6040518863ffffffff1660e060020a0281526004018088815260200187600160a060020a0316600160a060020a031681526020018663ffffffff1663ffffffff1681526020018563ffffffff1663ffffffff168152602001848152602001806020018281038252848482818152602001925080828437820191505098505050505050505050600060405180830381600087803b158015610bc157600080fd5b505af1158015610bd5573d6000803e3d6000fd5b5050506000838152600760205260409020805483925060ff19166001836006811115610bfd57fe5b02179055505050505050505050505050565b6000805460a060020a900460ff1615610c2757600080fd5b600082815260076020526040902054829060019060ff166006811115610c4957fe5b816006811115610c5557fe5b14610c5f57600080fd5b600084815260076020526040902060038101549093508490600290600160a060020a03163314610c8e57600080fd5b6005850154600160a060020a03161515610ca757600080fd5b60058501546001860154604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481019290925251600160a060020a03909216916323b872dd916064808201926020929091908290030181600087803b158015610d2157600080fd5b505af1158015610d35573d6000803e3d6000fd5b505050506040513d6020811015610d4b57600080fd5b50506000828152600760205260409020805482919060ff1916600183600681111561083057fe5b600054600160a060020a031681565b6000805460a060020a900460ff1615610d9957600080fd5b600082815260076020526040902054829060049060ff166006811115610dbb57fe5b816006811115610dc757fe5b14610dd157600080fd5b60008481526007602052604090206005818101549194508591600160a060020a031615610dfd57600080fd5b600485015460078601546001870154600160a060020a03909216916108fc91610e2c919063ffffffff611cf116565b6040518115909202916000818181858888f19350505050158015610e54573d6000803e3d6000fd5b506000828152600760205260409020805482919060ff1916600183600681111561083057fe5b60076020819052600091825260409091208054600182015460028301546003840154600485015460058601546006870154969097015460ff9095169693959294600160a060020a039283169491831693919092169188565b336000908152600260205260408120548190819060ff161515610ef457600080fd5b60005460a060020a900460ff1615610f0b57600080fd5b600087815260076020526040902054879060029060ff166006811115610f2d57fe5b816006811115610f3957fe5b14610f4357600080fd5b886003600760008c81526020019081526020016000209650600560009054906101000a9004600160a060020a0316600160a060020a031663969596d66040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610fae57600080fd5b505af1158015610fc2573d6000803e3d6000fd5b505050506040513d6020811015610fd857600080fd5b50516005880154909650600160a060020a03161561127357600160a060020a038616156111395760058701546003546001890154604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a039384166004820152602481019290925251919092169163a9059cbb91604480830192600092919082900301818387803b15801561107757600080fd5b505af115801561108b573d6000803e3d6000fd5b505060035460028a015460058b015460018c0154604080517fcc4fbc43000000000000000000000000000000000000000000000000000000008152600160a060020a038e811660048301526024820195909552928416604484015260648301919091525191909216935063cc4fbc439250608480830192600092919082900301818387803b15801561111c57600080fd5b505af1158015611130573d6000803e3d6000fd5b5050505061126e565b60058701546003546001890154604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a039384166004820152602481019290925251919092169163a9059cbb91604480830192600092919082900301818387803b1580156111b157600080fd5b505af11580156111c5573d6000803e3d6000fd5b50506003546005805460028c0154918c015460018d0154604080517fcc4fbc43000000000000000000000000000000000000000000000000000000008152600160a060020a039485166004820152602481019590955291831660448501526064840152519216935063cc4fbc43925060848082019260009290919082900301818387803b15801561125557600080fd5b505af1158015611269573d6000803e3d6000fd5b505050505b611437565b60009450600160a060020a038616156113575760035460018801546002808a01546004808c015460068d0154604080517fd21c39a1000000000000000000000000000000000000000000000000000000008152600160a060020a038f81169582019590955260248101959095529183166044850152606484015260848301939093529151919093169263d21c39a1929160a480830192602092919082900301818588803b15801561132357600080fd5b505af1158015611337573d6000803e3d6000fd5b50505050506040513d602081101561134e57600080fd5b50519450611426565b60035460018801546005546002808b01546004808d015460068e0154604080517fd21c39a1000000000000000000000000000000000000000000000000000000008152600160a060020a0397881694810194909452602484019490945290851660448301526064820152608481019290925251919093169263d21c39a1929160a480830192602092919082900301818588803b1580156113f657600080fd5b505af115801561140a573d6000803e3d6000fd5b50505050506040513d602081101561142157600080fd5b505194505b600085111561143757600787018590555b6114458b8b8b60018c611b80565b6000828152600760205260409020805482919060ff19166001836006811115610bfd57fe5b3360009081526002602052604081205460ff16151561148857600080fd5b60005460a060020a900460ff161561149f57600080fd5b600087815260076020526040902054879060019060ff1660068111156114c157fe5b8160068111156114cd57fe5b146114d757600080fd5b886006600086116114e757600080fd5b60008b81526007602052604081209550611507908c908c908c908c611b80565b600460009054906101000a9004600160a060020a0316600160a060020a031663a0e5e8218c8760040160009054906101000a9004600160a060020a03168d8d8d8d8d6040518863ffffffff1660e060020a0281526004018088815260200187600160a060020a0316600160a060020a031681526020018663ffffffff1663ffffffff1681526020018563ffffffff1663ffffffff168152602001848152602001806020018281038252848482818152602001925080828437820191505098505050505050505050600060405180830381600087803b158015610bc157600080fd5b600054600160a060020a031633146115ff57600080fd5b600160a060020a038116151561161457600080fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a0316331461165a57600080fd5b805161166d906001906020840190611dc3565b5050565b600554600160a060020a031681565b600054600160a060020a0316331461169757600080fd5b600160a060020a038216600081815260026020908152604091829020805460ff191685151590811790915582519384529083015280517fa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e6929281900390910190a15050565b60005460a060020a900460ff161561171257600080fd5b60008781526007602052604081205488919060ff16600681111561173257fe5b81600681111561173e57fe5b1461174857600080fd5b6000891161175557600080fd5b6000881161176257600080fd5b6000851015801561179557506117916103e8611785600f8b63ffffffff611d0816565b9063ffffffff611d3116565b8511155b15156117a057600080fd5b600160a060020a03871615156117b557600080fd5b600160a060020a03861615156117ca57600080fd5b6000898152600760205260409020600101541580156117f85750600089815260076020526040902060020154155b151561180357600080fd5b604080516101008101909152806001815260208082018b90526040808301899052600160a060020a03808c1660608501528a81166080850152881660a084015260c08301879052600060e09093018390528c83526007909152902081518154829060ff1916600183600681111561187657fe5b021790555060208201516001820155604082015160028201556060820151600382018054600160a060020a0392831673ffffffffffffffffffffffffffffffffffffffff1991821617909155608084015160048401805491841691831691909117905560a084015160058401805491909316911617905560c0820151600682015560e090910151600790910155505050505050505050565b6103e881565b600054600160a060020a0316331461192b57600080fd5b600160a060020a038116151561194057600080fd5b6006546000191681600160a060020a031663f0daba016040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561198557600080fd5b505af1158015611999573d6000803e3d6000fd5b505050506040513d60208110156119af57600080fd5b5051146119bb57600080fd5b6005805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600f81565b60005460a060020a900460ff1615611a0657600080fd5b600081815260076020526040902054819060049060ff166006811115611a2857fe5b816006811115611a3457fe5b14611a3e57600080fd5b6000838152600760205260409020600590810154849190600160a060020a03161515611a6957600080fd5b600085815260076020526040808220600581015460048083015460019093015484517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a03948516928101929092526024820152925191169263a9059cbb926044808201939182900301818387803b158015611aec57600080fd5b505af1158015611b00573d6000803e3d6000fd5b5050506000838152600760205260409020805483925060ff19166001836006811115611b2857fe5b02179055505050505050565b60065481565b600054600160a060020a03163314611b5157600080fd5b611b5a81611d46565b50565b600054600160a060020a03163314611b7457600080fd5b80600160a060020a0316ff5b6004805460008781526007602052604080822084015481517f5d818e6b0000000000000000000000000000000000000000000000000000000081529485018a9052600160a060020a03908116602486015263ffffffff808a16604487015288166064860152861515608486015260a485018690529051921692635d818e6b9260c4808301939282900301818387803b158015611c1b57600080fd5b505af1158015611c2f573d6000803e3d6000fd5b505060058054604080517fe25a51b600000000000000000000000000000000000000000000000000000000815263ffffffff891660248201526004810182905260448101939093527f746f74616c000000000000000000000000000000000000000000000000000000606484015251600160a060020a03909116935063e25a51b69250608480830192600092919082900301818387803b158015611cd257600080fd5b505af1158015611ce6573d6000803e3d6000fd5b505050505050505050565b600082821115611cfd57fe5b508082035b92915050565b6000821515611d1957506000611d02565b50818102818382811515611d2957fe5b0414611d0257fe5b60008183811515611d3e57fe5b049392505050565b600160a060020a0381161515611d5b57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611e0457805160ff1916838001178555611e31565b82800160010185558215611e31579182015b82811115611e31578251825591602001919060010190611e16565b50611e3d929150611e41565b5090565b611e5b91905b80821115611e3d5760008155600101611e47565b905600a165627a7a72305820425d83036b103995a4b1fb5fa0bc1e5a0b2584d9eff20ea952e670fe3943f9c20029`

// DeployPaymentProcessorContract deploys a new Ethereum contract, binding an instance of PaymentProcessorContract to it.
func DeployPaymentProcessorContract(auth *bind.TransactOpts, backend bind.ContractBackend, _merchantId string, _merchantHistory common.Address, _monethaGateway common.Address, _merchantWallet common.Address) (common.Address, *types.Transaction, *PaymentProcessorContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentProcessorContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PaymentProcessorContractBin), backend, _merchantId, _merchantHistory, _monethaGateway, _merchantWallet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PaymentProcessorContract{PaymentProcessorContractCaller: PaymentProcessorContractCaller{contract: contract}, PaymentProcessorContractTransactor: PaymentProcessorContractTransactor{contract: contract}, PaymentProcessorContractFilterer: PaymentProcessorContractFilterer{contract: contract}}, nil
}

// PaymentProcessorContract is an auto generated Go binding around an Ethereum contract.
type PaymentProcessorContract struct {
	PaymentProcessorContractCaller     // Read-only binding to the contract
	PaymentProcessorContractTransactor // Write-only binding to the contract
	PaymentProcessorContractFilterer   // Log filterer for contract events
}

// PaymentProcessorContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentProcessorContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentProcessorContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentProcessorContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentProcessorContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentProcessorContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentProcessorContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentProcessorContractSession struct {
	Contract     *PaymentProcessorContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PaymentProcessorContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentProcessorContractCallerSession struct {
	Contract *PaymentProcessorContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// PaymentProcessorContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentProcessorContractTransactorSession struct {
	Contract     *PaymentProcessorContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// PaymentProcessorContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentProcessorContractRaw struct {
	Contract *PaymentProcessorContract // Generic contract binding to access the raw methods on
}

// PaymentProcessorContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentProcessorContractCallerRaw struct {
	Contract *PaymentProcessorContractCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentProcessorContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentProcessorContractTransactorRaw struct {
	Contract *PaymentProcessorContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaymentProcessorContract creates a new instance of PaymentProcessorContract, bound to a specific deployed contract.
func NewPaymentProcessorContract(address common.Address, backend bind.ContractBackend) (*PaymentProcessorContract, error) {
	contract, err := bindPaymentProcessorContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorContract{PaymentProcessorContractCaller: PaymentProcessorContractCaller{contract: contract}, PaymentProcessorContractTransactor: PaymentProcessorContractTransactor{contract: contract}, PaymentProcessorContractFilterer: PaymentProcessorContractFilterer{contract: contract}}, nil
}

// NewPaymentProcessorContractCaller creates a new read-only instance of PaymentProcessorContract, bound to a specific deployed contract.
func NewPaymentProcessorContractCaller(address common.Address, caller bind.ContractCaller) (*PaymentProcessorContractCaller, error) {
	contract, err := bindPaymentProcessorContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorContractCaller{contract: contract}, nil
}

// NewPaymentProcessorContractTransactor creates a new write-only instance of PaymentProcessorContract, bound to a specific deployed contract.
func NewPaymentProcessorContractTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentProcessorContractTransactor, error) {
	contract, err := bindPaymentProcessorContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorContractTransactor{contract: contract}, nil
}

// NewPaymentProcessorContractFilterer creates a new log filterer instance of PaymentProcessorContract, bound to a specific deployed contract.
func NewPaymentProcessorContractFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentProcessorContractFilterer, error) {
	contract, err := bindPaymentProcessorContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorContractFilterer{contract: contract}, nil
}

// bindPaymentProcessorContract binds a generic wrapper to an already deployed contract.
func bindPaymentProcessorContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentProcessorContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentProcessorContract *PaymentProcessorContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PaymentProcessorContract.Contract.PaymentProcessorContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentProcessorContract *PaymentProcessorContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.PaymentProcessorContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentProcessorContract *PaymentProcessorContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.PaymentProcessorContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentProcessorContract *PaymentProcessorContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PaymentProcessorContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentProcessorContract *PaymentProcessorContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentProcessorContract *PaymentProcessorContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.contract.Transact(opts, method, params...)
}

// FEEPERMILLE is a free data retrieval call binding the contract method 0xde8248fb.
//
// Solidity: function FEE_PERMILLE() constant returns(uint256)
func (_PaymentProcessorContract *PaymentProcessorContractCaller) FEEPERMILLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentProcessorContract.contract.Call(opts, out, "FEE_PERMILLE")
	return *ret0, err
}

// FEEPERMILLE is a free data retrieval call binding the contract method 0xde8248fb.
//
// Solidity: function FEE_PERMILLE() constant returns(uint256)
func (_PaymentProcessorContract *PaymentProcessorContractSession) FEEPERMILLE() (*big.Int, error) {
	return _PaymentProcessorContract.Contract.FEEPERMILLE(&_PaymentProcessorContract.CallOpts)
}

// FEEPERMILLE is a free data retrieval call binding the contract method 0xde8248fb.
//
// Solidity: function FEE_PERMILLE() constant returns(uint256)
func (_PaymentProcessorContract *PaymentProcessorContractCallerSession) FEEPERMILLE() (*big.Int, error) {
	return _PaymentProcessorContract.Contract.FEEPERMILLE(&_PaymentProcessorContract.CallOpts)
}

// PAYBACKPERMILLE is a free data retrieval call binding the contract method 0x4baf43e4.
//
// Solidity: function PAYBACK_PERMILLE() constant returns(uint256)
func (_PaymentProcessorContract *PaymentProcessorContractCaller) PAYBACKPERMILLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentProcessorContract.contract.Call(opts, out, "PAYBACK_PERMILLE")
	return *ret0, err
}

// PAYBACKPERMILLE is a free data retrieval call binding the contract method 0x4baf43e4.
//
// Solidity: function PAYBACK_PERMILLE() constant returns(uint256)
func (_PaymentProcessorContract *PaymentProcessorContractSession) PAYBACKPERMILLE() (*big.Int, error) {
	return _PaymentProcessorContract.Contract.PAYBACKPERMILLE(&_PaymentProcessorContract.CallOpts)
}

// PAYBACKPERMILLE is a free data retrieval call binding the contract method 0x4baf43e4.
//
// Solidity: function PAYBACK_PERMILLE() constant returns(uint256)
func (_PaymentProcessorContract *PaymentProcessorContractCallerSession) PAYBACKPERMILLE() (*big.Int, error) {
	return _PaymentProcessorContract.Contract.PAYBACKPERMILLE(&_PaymentProcessorContract.CallOpts)
}

// PERMILLECOEFFICIENT is a free data retrieval call binding the contract method 0xc4cc3f4d.
//
// Solidity: function PERMILLE_COEFFICIENT() constant returns(uint256)
func (_PaymentProcessorContract *PaymentProcessorContractCaller) PERMILLECOEFFICIENT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentProcessorContract.contract.Call(opts, out, "PERMILLE_COEFFICIENT")
	return *ret0, err
}

// PERMILLECOEFFICIENT is a free data retrieval call binding the contract method 0xc4cc3f4d.
//
// Solidity: function PERMILLE_COEFFICIENT() constant returns(uint256)
func (_PaymentProcessorContract *PaymentProcessorContractSession) PERMILLECOEFFICIENT() (*big.Int, error) {
	return _PaymentProcessorContract.Contract.PERMILLECOEFFICIENT(&_PaymentProcessorContract.CallOpts)
}

// PERMILLECOEFFICIENT is a free data retrieval call binding the contract method 0xc4cc3f4d.
//
// Solidity: function PERMILLE_COEFFICIENT() constant returns(uint256)
func (_PaymentProcessorContract *PaymentProcessorContractCallerSession) PERMILLECOEFFICIENT() (*big.Int, error) {
	return _PaymentProcessorContract.Contract.PERMILLECOEFFICIENT(&_PaymentProcessorContract.CallOpts)
}

// ContactInformation is a free data retrieval call binding the contract method 0x36f7ab5e.
//
// Solidity: function contactInformation() constant returns(string)
func (_PaymentProcessorContract *PaymentProcessorContractCaller) ContactInformation(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PaymentProcessorContract.contract.Call(opts, out, "contactInformation")
	return *ret0, err
}

// ContactInformation is a free data retrieval call binding the contract method 0x36f7ab5e.
//
// Solidity: function contactInformation() constant returns(string)
func (_PaymentProcessorContract *PaymentProcessorContractSession) ContactInformation() (string, error) {
	return _PaymentProcessorContract.Contract.ContactInformation(&_PaymentProcessorContract.CallOpts)
}

// ContactInformation is a free data retrieval call binding the contract method 0x36f7ab5e.
//
// Solidity: function contactInformation() constant returns(string)
func (_PaymentProcessorContract *PaymentProcessorContractCallerSession) ContactInformation() (string, error) {
	return _PaymentProcessorContract.Contract.ContactInformation(&_PaymentProcessorContract.CallOpts)
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_PaymentProcessorContract *PaymentProcessorContractCaller) IsMonethaAddress(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PaymentProcessorContract.contract.Call(opts, out, "isMonethaAddress", arg0)
	return *ret0, err
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_PaymentProcessorContract *PaymentProcessorContractSession) IsMonethaAddress(arg0 common.Address) (bool, error) {
	return _PaymentProcessorContract.Contract.IsMonethaAddress(&_PaymentProcessorContract.CallOpts, arg0)
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_PaymentProcessorContract *PaymentProcessorContractCallerSession) IsMonethaAddress(arg0 common.Address) (bool, error) {
	return _PaymentProcessorContract.Contract.IsMonethaAddress(&_PaymentProcessorContract.CallOpts, arg0)
}

// MerchantHistory is a free data retrieval call binding the contract method 0x3d17a2d8.
//
// Solidity: function merchantHistory() constant returns(address)
func (_PaymentProcessorContract *PaymentProcessorContractCaller) MerchantHistory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PaymentProcessorContract.contract.Call(opts, out, "merchantHistory")
	return *ret0, err
}

// MerchantHistory is a free data retrieval call binding the contract method 0x3d17a2d8.
//
// Solidity: function merchantHistory() constant returns(address)
func (_PaymentProcessorContract *PaymentProcessorContractSession) MerchantHistory() (common.Address, error) {
	return _PaymentProcessorContract.Contract.MerchantHistory(&_PaymentProcessorContract.CallOpts)
}

// MerchantHistory is a free data retrieval call binding the contract method 0x3d17a2d8.
//
// Solidity: function merchantHistory() constant returns(address)
func (_PaymentProcessorContract *PaymentProcessorContractCallerSession) MerchantHistory() (common.Address, error) {
	return _PaymentProcessorContract.Contract.MerchantHistory(&_PaymentProcessorContract.CallOpts)
}

// MerchantIdHash is a free data retrieval call binding the contract method 0xf0daba01.
//
// Solidity: function merchantIdHash() constant returns(bytes32)
func (_PaymentProcessorContract *PaymentProcessorContractCaller) MerchantIdHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PaymentProcessorContract.contract.Call(opts, out, "merchantIdHash")
	return *ret0, err
}

// MerchantIdHash is a free data retrieval call binding the contract method 0xf0daba01.
//
// Solidity: function merchantIdHash() constant returns(bytes32)
func (_PaymentProcessorContract *PaymentProcessorContractSession) MerchantIdHash() ([32]byte, error) {
	return _PaymentProcessorContract.Contract.MerchantIdHash(&_PaymentProcessorContract.CallOpts)
}

// MerchantIdHash is a free data retrieval call binding the contract method 0xf0daba01.
//
// Solidity: function merchantIdHash() constant returns(bytes32)
func (_PaymentProcessorContract *PaymentProcessorContractCallerSession) MerchantIdHash() ([32]byte, error) {
	return _PaymentProcessorContract.Contract.MerchantIdHash(&_PaymentProcessorContract.CallOpts)
}

// MerchantWallet is a free data retrieval call binding the contract method 0xbc85e064.
//
// Solidity: function merchantWallet() constant returns(address)
func (_PaymentProcessorContract *PaymentProcessorContractCaller) MerchantWallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PaymentProcessorContract.contract.Call(opts, out, "merchantWallet")
	return *ret0, err
}

// MerchantWallet is a free data retrieval call binding the contract method 0xbc85e064.
//
// Solidity: function merchantWallet() constant returns(address)
func (_PaymentProcessorContract *PaymentProcessorContractSession) MerchantWallet() (common.Address, error) {
	return _PaymentProcessorContract.Contract.MerchantWallet(&_PaymentProcessorContract.CallOpts)
}

// MerchantWallet is a free data retrieval call binding the contract method 0xbc85e064.
//
// Solidity: function merchantWallet() constant returns(address)
func (_PaymentProcessorContract *PaymentProcessorContractCallerSession) MerchantWallet() (common.Address, error) {
	return _PaymentProcessorContract.Contract.MerchantWallet(&_PaymentProcessorContract.CallOpts)
}

// MonethaGateway is a free data retrieval call binding the contract method 0x6137412c.
//
// Solidity: function monethaGateway() constant returns(address)
func (_PaymentProcessorContract *PaymentProcessorContractCaller) MonethaGateway(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PaymentProcessorContract.contract.Call(opts, out, "monethaGateway")
	return *ret0, err
}

// MonethaGateway is a free data retrieval call binding the contract method 0x6137412c.
//
// Solidity: function monethaGateway() constant returns(address)
func (_PaymentProcessorContract *PaymentProcessorContractSession) MonethaGateway() (common.Address, error) {
	return _PaymentProcessorContract.Contract.MonethaGateway(&_PaymentProcessorContract.CallOpts)
}

// MonethaGateway is a free data retrieval call binding the contract method 0x6137412c.
//
// Solidity: function monethaGateway() constant returns(address)
func (_PaymentProcessorContract *PaymentProcessorContractCallerSession) MonethaGateway() (common.Address, error) {
	return _PaymentProcessorContract.Contract.MonethaGateway(&_PaymentProcessorContract.CallOpts)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders( uint256) constant returns(state uint8, price uint256, fee uint256, paymentAcceptor address, originAddress address, tokenAddress address, vouchersApply uint256, discount uint256)
func (_PaymentProcessorContract *PaymentProcessorContractCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	State           uint8
	Price           *big.Int
	Fee             *big.Int
	PaymentAcceptor common.Address
	OriginAddress   common.Address
	TokenAddress    common.Address
	VouchersApply   *big.Int
	Discount        *big.Int
}, error) {
	ret := new(struct {
		State           uint8
		Price           *big.Int
		Fee             *big.Int
		PaymentAcceptor common.Address
		OriginAddress   common.Address
		TokenAddress    common.Address
		VouchersApply   *big.Int
		Discount        *big.Int
	})
	out := ret
	err := _PaymentProcessorContract.contract.Call(opts, out, "orders", arg0)
	return *ret, err
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders( uint256) constant returns(state uint8, price uint256, fee uint256, paymentAcceptor address, originAddress address, tokenAddress address, vouchersApply uint256, discount uint256)
func (_PaymentProcessorContract *PaymentProcessorContractSession) Orders(arg0 *big.Int) (struct {
	State           uint8
	Price           *big.Int
	Fee             *big.Int
	PaymentAcceptor common.Address
	OriginAddress   common.Address
	TokenAddress    common.Address
	VouchersApply   *big.Int
	Discount        *big.Int
}, error) {
	return _PaymentProcessorContract.Contract.Orders(&_PaymentProcessorContract.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders( uint256) constant returns(state uint8, price uint256, fee uint256, paymentAcceptor address, originAddress address, tokenAddress address, vouchersApply uint256, discount uint256)
func (_PaymentProcessorContract *PaymentProcessorContractCallerSession) Orders(arg0 *big.Int) (struct {
	State           uint8
	Price           *big.Int
	Fee             *big.Int
	PaymentAcceptor common.Address
	OriginAddress   common.Address
	TokenAddress    common.Address
	VouchersApply   *big.Int
	Discount        *big.Int
}, error) {
	return _PaymentProcessorContract.Contract.Orders(&_PaymentProcessorContract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PaymentProcessorContract *PaymentProcessorContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PaymentProcessorContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PaymentProcessorContract *PaymentProcessorContractSession) Owner() (common.Address, error) {
	return _PaymentProcessorContract.Contract.Owner(&_PaymentProcessorContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PaymentProcessorContract *PaymentProcessorContractCallerSession) Owner() (common.Address, error) {
	return _PaymentProcessorContract.Contract.Owner(&_PaymentProcessorContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PaymentProcessorContract *PaymentProcessorContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PaymentProcessorContract.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PaymentProcessorContract *PaymentProcessorContractSession) Paused() (bool, error) {
	return _PaymentProcessorContract.Contract.Paused(&_PaymentProcessorContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PaymentProcessorContract *PaymentProcessorContractCallerSession) Paused() (bool, error) {
	return _PaymentProcessorContract.Contract.Paused(&_PaymentProcessorContract.CallOpts)
}

// AddOrder is a paid mutator transaction binding the contract method 0xc48aab6d.
//
// Solidity: function addOrder(_orderId uint256, _price uint256, _paymentAcceptor address, _originAddress address, _fee uint256, _tokenAddress address, _vouchersApply uint256) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactor) AddOrder(opts *bind.TransactOpts, _orderId *big.Int, _price *big.Int, _paymentAcceptor common.Address, _originAddress common.Address, _fee *big.Int, _tokenAddress common.Address, _vouchersApply *big.Int) (*types.Transaction, error) {
	return _PaymentProcessorContract.contract.Transact(opts, "addOrder", _orderId, _price, _paymentAcceptor, _originAddress, _fee, _tokenAddress, _vouchersApply)
}

// AddOrder is a paid mutator transaction binding the contract method 0xc48aab6d.
//
// Solidity: function addOrder(_orderId uint256, _price uint256, _paymentAcceptor address, _originAddress address, _fee uint256, _tokenAddress address, _vouchersApply uint256) returns()
func (_PaymentProcessorContract *PaymentProcessorContractSession) AddOrder(_orderId *big.Int, _price *big.Int, _paymentAcceptor common.Address, _originAddress common.Address, _fee *big.Int, _tokenAddress common.Address, _vouchersApply *big.Int) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.AddOrder(&_PaymentProcessorContract.TransactOpts, _orderId, _price, _paymentAcceptor, _originAddress, _fee, _tokenAddress, _vouchersApply)
}

// AddOrder is a paid mutator transaction binding the contract method 0xc48aab6d.
//
// Solidity: function addOrder(_orderId uint256, _price uint256, _paymentAcceptor address, _originAddress address, _fee uint256, _tokenAddress address, _vouchersApply uint256) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactorSession) AddOrder(_orderId *big.Int, _price *big.Int, _paymentAcceptor common.Address, _originAddress common.Address, _fee *big.Int, _tokenAddress common.Address, _vouchersApply *big.Int) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.AddOrder(&_PaymentProcessorContract.TransactOpts, _orderId, _price, _paymentAcceptor, _originAddress, _fee, _tokenAddress, _vouchersApply)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xac1a13fb.
//
// Solidity: function cancelOrder(_orderId uint256, _clientReputation uint32, _merchantReputation uint32, _dealHash uint256, _cancelReason string) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactor) CancelOrder(opts *bind.TransactOpts, _orderId *big.Int, _clientReputation uint32, _merchantReputation uint32, _dealHash *big.Int, _cancelReason string) (*types.Transaction, error) {
	return _PaymentProcessorContract.contract.Transact(opts, "cancelOrder", _orderId, _clientReputation, _merchantReputation, _dealHash, _cancelReason)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xac1a13fb.
//
// Solidity: function cancelOrder(_orderId uint256, _clientReputation uint32, _merchantReputation uint32, _dealHash uint256, _cancelReason string) returns()
func (_PaymentProcessorContract *PaymentProcessorContractSession) CancelOrder(_orderId *big.Int, _clientReputation uint32, _merchantReputation uint32, _dealHash *big.Int, _cancelReason string) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.CancelOrder(&_PaymentProcessorContract.TransactOpts, _orderId, _clientReputation, _merchantReputation, _dealHash, _cancelReason)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xac1a13fb.
//
// Solidity: function cancelOrder(_orderId uint256, _clientReputation uint32, _merchantReputation uint32, _dealHash uint256, _cancelReason string) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactorSession) CancelOrder(_orderId *big.Int, _clientReputation uint32, _merchantReputation uint32, _dealHash *big.Int, _cancelReason string) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.CancelOrder(&_PaymentProcessorContract.TransactOpts, _orderId, _clientReputation, _merchantReputation, _dealHash, _cancelReason)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentProcessorContract.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_PaymentProcessorContract *PaymentProcessorContractSession) Destroy() (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.Destroy(&_PaymentProcessorContract.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactorSession) Destroy() (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.Destroy(&_PaymentProcessorContract.TransactOpts)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactor) DestroyAndSend(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _PaymentProcessorContract.contract.Transact(opts, "destroyAndSend", _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_PaymentProcessorContract *PaymentProcessorContractSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.DestroyAndSend(&_PaymentProcessorContract.TransactOpts, _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactorSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.DestroyAndSend(&_PaymentProcessorContract.TransactOpts, _recipient)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentProcessorContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PaymentProcessorContract *PaymentProcessorContractSession) Pause() (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.Pause(&_PaymentProcessorContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactorSession) Pause() (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.Pause(&_PaymentProcessorContract.TransactOpts)
}

// ProcessPayment is a paid mutator transaction binding the contract method 0xaab9f165.
//
// Solidity: function processPayment(_orderId uint256, _clientReputation uint32, _merchantReputation uint32, _dealHash uint256) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactor) ProcessPayment(opts *bind.TransactOpts, _orderId *big.Int, _clientReputation uint32, _merchantReputation uint32, _dealHash *big.Int) (*types.Transaction, error) {
	return _PaymentProcessorContract.contract.Transact(opts, "processPayment", _orderId, _clientReputation, _merchantReputation, _dealHash)
}

// ProcessPayment is a paid mutator transaction binding the contract method 0xaab9f165.
//
// Solidity: function processPayment(_orderId uint256, _clientReputation uint32, _merchantReputation uint32, _dealHash uint256) returns()
func (_PaymentProcessorContract *PaymentProcessorContractSession) ProcessPayment(_orderId *big.Int, _clientReputation uint32, _merchantReputation uint32, _dealHash *big.Int) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.ProcessPayment(&_PaymentProcessorContract.TransactOpts, _orderId, _clientReputation, _merchantReputation, _dealHash)
}

// ProcessPayment is a paid mutator transaction binding the contract method 0xaab9f165.
//
// Solidity: function processPayment(_orderId uint256, _clientReputation uint32, _merchantReputation uint32, _dealHash uint256) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactorSession) ProcessPayment(_orderId *big.Int, _clientReputation uint32, _merchantReputation uint32, _dealHash *big.Int) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.ProcessPayment(&_PaymentProcessorContract.TransactOpts, _orderId, _clientReputation, _merchantReputation, _dealHash)
}

// RefundPayment is a paid mutator transaction binding the contract method 0x8467d9cf.
//
// Solidity: function refundPayment(_orderId uint256, _clientReputation uint32, _merchantReputation uint32, _dealHash uint256, _refundReason string) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactor) RefundPayment(opts *bind.TransactOpts, _orderId *big.Int, _clientReputation uint32, _merchantReputation uint32, _dealHash *big.Int, _refundReason string) (*types.Transaction, error) {
	return _PaymentProcessorContract.contract.Transact(opts, "refundPayment", _orderId, _clientReputation, _merchantReputation, _dealHash, _refundReason)
}

// RefundPayment is a paid mutator transaction binding the contract method 0x8467d9cf.
//
// Solidity: function refundPayment(_orderId uint256, _clientReputation uint32, _merchantReputation uint32, _dealHash uint256, _refundReason string) returns()
func (_PaymentProcessorContract *PaymentProcessorContractSession) RefundPayment(_orderId *big.Int, _clientReputation uint32, _merchantReputation uint32, _dealHash *big.Int, _refundReason string) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.RefundPayment(&_PaymentProcessorContract.TransactOpts, _orderId, _clientReputation, _merchantReputation, _dealHash, _refundReason)
}

// RefundPayment is a paid mutator transaction binding the contract method 0x8467d9cf.
//
// Solidity: function refundPayment(_orderId uint256, _clientReputation uint32, _merchantReputation uint32, _dealHash uint256, _refundReason string) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactorSession) RefundPayment(_orderId *big.Int, _clientReputation uint32, _merchantReputation uint32, _dealHash *big.Int, _refundReason string) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.RefundPayment(&_PaymentProcessorContract.TransactOpts, _orderId, _clientReputation, _merchantReputation, _dealHash, _refundReason)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentProcessorContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PaymentProcessorContract *PaymentProcessorContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.RenounceOwnership(&_PaymentProcessorContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.RenounceOwnership(&_PaymentProcessorContract.TransactOpts)
}

// SecurePay is a paid mutator transaction binding the contract method 0x4153d65b.
//
// Solidity: function securePay(_orderId uint256) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactor) SecurePay(opts *bind.TransactOpts, _orderId *big.Int) (*types.Transaction, error) {
	return _PaymentProcessorContract.contract.Transact(opts, "securePay", _orderId)
}

// SecurePay is a paid mutator transaction binding the contract method 0x4153d65b.
//
// Solidity: function securePay(_orderId uint256) returns()
func (_PaymentProcessorContract *PaymentProcessorContractSession) SecurePay(_orderId *big.Int) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.SecurePay(&_PaymentProcessorContract.TransactOpts, _orderId)
}

// SecurePay is a paid mutator transaction binding the contract method 0x4153d65b.
//
// Solidity: function securePay(_orderId uint256) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactorSession) SecurePay(_orderId *big.Int) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.SecurePay(&_PaymentProcessorContract.TransactOpts, _orderId)
}

// SecureTokenPay is a paid mutator transaction binding the contract method 0x8b11fb3e.
//
// Solidity: function secureTokenPay(_orderId uint256) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactor) SecureTokenPay(opts *bind.TransactOpts, _orderId *big.Int) (*types.Transaction, error) {
	return _PaymentProcessorContract.contract.Transact(opts, "secureTokenPay", _orderId)
}

// SecureTokenPay is a paid mutator transaction binding the contract method 0x8b11fb3e.
//
// Solidity: function secureTokenPay(_orderId uint256) returns()
func (_PaymentProcessorContract *PaymentProcessorContractSession) SecureTokenPay(_orderId *big.Int) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.SecureTokenPay(&_PaymentProcessorContract.TransactOpts, _orderId)
}

// SecureTokenPay is a paid mutator transaction binding the contract method 0x8b11fb3e.
//
// Solidity: function secureTokenPay(_orderId uint256) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactorSession) SecureTokenPay(_orderId *big.Int) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.SecureTokenPay(&_PaymentProcessorContract.TransactOpts, _orderId)
}

// SetContactInformation is a paid mutator transaction binding the contract method 0xb967a52e.
//
// Solidity: function setContactInformation(_info string) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactor) SetContactInformation(opts *bind.TransactOpts, _info string) (*types.Transaction, error) {
	return _PaymentProcessorContract.contract.Transact(opts, "setContactInformation", _info)
}

// SetContactInformation is a paid mutator transaction binding the contract method 0xb967a52e.
//
// Solidity: function setContactInformation(_info string) returns()
func (_PaymentProcessorContract *PaymentProcessorContractSession) SetContactInformation(_info string) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.SetContactInformation(&_PaymentProcessorContract.TransactOpts, _info)
}

// SetContactInformation is a paid mutator transaction binding the contract method 0xb967a52e.
//
// Solidity: function setContactInformation(_info string) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactorSession) SetContactInformation(_info string) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.SetContactInformation(&_PaymentProcessorContract.TransactOpts, _info)
}

// SetMerchantDealsHistory is a paid mutator transaction binding the contract method 0x444b6048.
//
// Solidity: function setMerchantDealsHistory(_merchantHistory address) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactor) SetMerchantDealsHistory(opts *bind.TransactOpts, _merchantHistory common.Address) (*types.Transaction, error) {
	return _PaymentProcessorContract.contract.Transact(opts, "setMerchantDealsHistory", _merchantHistory)
}

// SetMerchantDealsHistory is a paid mutator transaction binding the contract method 0x444b6048.
//
// Solidity: function setMerchantDealsHistory(_merchantHistory address) returns()
func (_PaymentProcessorContract *PaymentProcessorContractSession) SetMerchantDealsHistory(_merchantHistory common.Address) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.SetMerchantDealsHistory(&_PaymentProcessorContract.TransactOpts, _merchantHistory)
}

// SetMerchantDealsHistory is a paid mutator transaction binding the contract method 0x444b6048.
//
// Solidity: function setMerchantDealsHistory(_merchantHistory address) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactorSession) SetMerchantDealsHistory(_merchantHistory common.Address) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.SetMerchantDealsHistory(&_PaymentProcessorContract.TransactOpts, _merchantHistory)
}

// SetMerchantWallet is a paid mutator transaction binding the contract method 0xddda66db.
//
// Solidity: function setMerchantWallet(_newWallet address) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactor) SetMerchantWallet(opts *bind.TransactOpts, _newWallet common.Address) (*types.Transaction, error) {
	return _PaymentProcessorContract.contract.Transact(opts, "setMerchantWallet", _newWallet)
}

// SetMerchantWallet is a paid mutator transaction binding the contract method 0xddda66db.
//
// Solidity: function setMerchantWallet(_newWallet address) returns()
func (_PaymentProcessorContract *PaymentProcessorContractSession) SetMerchantWallet(_newWallet common.Address) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.SetMerchantWallet(&_PaymentProcessorContract.TransactOpts, _newWallet)
}

// SetMerchantWallet is a paid mutator transaction binding the contract method 0xddda66db.
//
// Solidity: function setMerchantWallet(_newWallet address) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactorSession) SetMerchantWallet(_newWallet common.Address) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.SetMerchantWallet(&_PaymentProcessorContract.TransactOpts, _newWallet)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactor) SetMonethaAddress(opts *bind.TransactOpts, _address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _PaymentProcessorContract.contract.Transact(opts, "setMonethaAddress", _address, _isMonethaAddress)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_PaymentProcessorContract *PaymentProcessorContractSession) SetMonethaAddress(_address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.SetMonethaAddress(&_PaymentProcessorContract.TransactOpts, _address, _isMonethaAddress)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactorSession) SetMonethaAddress(_address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.SetMonethaAddress(&_PaymentProcessorContract.TransactOpts, _address, _isMonethaAddress)
}

// SetMonethaGateway is a paid mutator transaction binding the contract method 0xb440bf39.
//
// Solidity: function setMonethaGateway(_newGateway address) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactor) SetMonethaGateway(opts *bind.TransactOpts, _newGateway common.Address) (*types.Transaction, error) {
	return _PaymentProcessorContract.contract.Transact(opts, "setMonethaGateway", _newGateway)
}

// SetMonethaGateway is a paid mutator transaction binding the contract method 0xb440bf39.
//
// Solidity: function setMonethaGateway(_newGateway address) returns()
func (_PaymentProcessorContract *PaymentProcessorContractSession) SetMonethaGateway(_newGateway common.Address) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.SetMonethaGateway(&_PaymentProcessorContract.TransactOpts, _newGateway)
}

// SetMonethaGateway is a paid mutator transaction binding the contract method 0xb440bf39.
//
// Solidity: function setMonethaGateway(_newGateway address) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactorSession) SetMonethaGateway(_newGateway common.Address) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.SetMonethaGateway(&_PaymentProcessorContract.TransactOpts, _newGateway)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _PaymentProcessorContract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_PaymentProcessorContract *PaymentProcessorContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.TransferOwnership(&_PaymentProcessorContract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.TransferOwnership(&_PaymentProcessorContract.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentProcessorContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PaymentProcessorContract *PaymentProcessorContractSession) Unpause() (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.Unpause(&_PaymentProcessorContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.Unpause(&_PaymentProcessorContract.TransactOpts)
}

// WithdrawRefund is a paid mutator transaction binding the contract method 0x9d153495.
//
// Solidity: function withdrawRefund(_orderId uint256) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactor) WithdrawRefund(opts *bind.TransactOpts, _orderId *big.Int) (*types.Transaction, error) {
	return _PaymentProcessorContract.contract.Transact(opts, "withdrawRefund", _orderId)
}

// WithdrawRefund is a paid mutator transaction binding the contract method 0x9d153495.
//
// Solidity: function withdrawRefund(_orderId uint256) returns()
func (_PaymentProcessorContract *PaymentProcessorContractSession) WithdrawRefund(_orderId *big.Int) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.WithdrawRefund(&_PaymentProcessorContract.TransactOpts, _orderId)
}

// WithdrawRefund is a paid mutator transaction binding the contract method 0x9d153495.
//
// Solidity: function withdrawRefund(_orderId uint256) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactorSession) WithdrawRefund(_orderId *big.Int) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.WithdrawRefund(&_PaymentProcessorContract.TransactOpts, _orderId)
}

// WithdrawTokenRefund is a paid mutator transaction binding the contract method 0xe5b8d6e0.
//
// Solidity: function withdrawTokenRefund(_orderId uint256) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactor) WithdrawTokenRefund(opts *bind.TransactOpts, _orderId *big.Int) (*types.Transaction, error) {
	return _PaymentProcessorContract.contract.Transact(opts, "withdrawTokenRefund", _orderId)
}

// WithdrawTokenRefund is a paid mutator transaction binding the contract method 0xe5b8d6e0.
//
// Solidity: function withdrawTokenRefund(_orderId uint256) returns()
func (_PaymentProcessorContract *PaymentProcessorContractSession) WithdrawTokenRefund(_orderId *big.Int) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.WithdrawTokenRefund(&_PaymentProcessorContract.TransactOpts, _orderId)
}

// WithdrawTokenRefund is a paid mutator transaction binding the contract method 0xe5b8d6e0.
//
// Solidity: function withdrawTokenRefund(_orderId uint256) returns()
func (_PaymentProcessorContract *PaymentProcessorContractTransactorSession) WithdrawTokenRefund(_orderId *big.Int) (*types.Transaction, error) {
	return _PaymentProcessorContract.Contract.WithdrawTokenRefund(&_PaymentProcessorContract.TransactOpts, _orderId)
}

// PaymentProcessorContractMonethaAddressSetIterator is returned from FilterMonethaAddressSet and is used to iterate over the raw logs and unpacked data for MonethaAddressSet events raised by the PaymentProcessorContract contract.
type PaymentProcessorContractMonethaAddressSetIterator struct {
	Event *PaymentProcessorContractMonethaAddressSet // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorContractMonethaAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorContractMonethaAddressSet)
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
		it.Event = new(PaymentProcessorContractMonethaAddressSet)
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
func (it *PaymentProcessorContractMonethaAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorContractMonethaAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorContractMonethaAddressSet represents a MonethaAddressSet event raised by the PaymentProcessorContract contract.
type PaymentProcessorContractMonethaAddressSet struct {
	Address          common.Address
	IsMonethaAddress bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMonethaAddressSet is a free log retrieval operation binding the contract event 0xa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e692.
//
// Solidity: e MonethaAddressSet(_address address, _isMonethaAddress bool)
func (_PaymentProcessorContract *PaymentProcessorContractFilterer) FilterMonethaAddressSet(opts *bind.FilterOpts) (*PaymentProcessorContractMonethaAddressSetIterator, error) {

	logs, sub, err := _PaymentProcessorContract.contract.FilterLogs(opts, "MonethaAddressSet")
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorContractMonethaAddressSetIterator{contract: _PaymentProcessorContract.contract, event: "MonethaAddressSet", logs: logs, sub: sub}, nil
}

// WatchMonethaAddressSet is a free log subscription operation binding the contract event 0xa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e692.
//
// Solidity: e MonethaAddressSet(_address address, _isMonethaAddress bool)
func (_PaymentProcessorContract *PaymentProcessorContractFilterer) WatchMonethaAddressSet(opts *bind.WatchOpts, sink chan<- *PaymentProcessorContractMonethaAddressSet) (event.Subscription, error) {

	logs, sub, err := _PaymentProcessorContract.contract.WatchLogs(opts, "MonethaAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorContractMonethaAddressSet)
				if err := _PaymentProcessorContract.contract.UnpackLog(event, "MonethaAddressSet", log); err != nil {
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

// PaymentProcessorContractOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the PaymentProcessorContract contract.
type PaymentProcessorContractOwnershipRenouncedIterator struct {
	Event *PaymentProcessorContractOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorContractOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorContractOwnershipRenounced)
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
		it.Event = new(PaymentProcessorContractOwnershipRenounced)
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
func (it *PaymentProcessorContractOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorContractOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorContractOwnershipRenounced represents a OwnershipRenounced event raised by the PaymentProcessorContract contract.
type PaymentProcessorContractOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_PaymentProcessorContract *PaymentProcessorContractFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*PaymentProcessorContractOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PaymentProcessorContract.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorContractOwnershipRenouncedIterator{contract: _PaymentProcessorContract.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_PaymentProcessorContract *PaymentProcessorContractFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *PaymentProcessorContractOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PaymentProcessorContract.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorContractOwnershipRenounced)
				if err := _PaymentProcessorContract.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// PaymentProcessorContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PaymentProcessorContract contract.
type PaymentProcessorContractOwnershipTransferredIterator struct {
	Event *PaymentProcessorContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorContractOwnershipTransferred)
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
		it.Event = new(PaymentProcessorContractOwnershipTransferred)
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
func (it *PaymentProcessorContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorContractOwnershipTransferred represents a OwnershipTransferred event raised by the PaymentProcessorContract contract.
type PaymentProcessorContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PaymentProcessorContract *PaymentProcessorContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PaymentProcessorContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PaymentProcessorContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorContractOwnershipTransferredIterator{contract: _PaymentProcessorContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PaymentProcessorContract *PaymentProcessorContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PaymentProcessorContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PaymentProcessorContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorContractOwnershipTransferred)
				if err := _PaymentProcessorContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PaymentProcessorContractPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the PaymentProcessorContract contract.
type PaymentProcessorContractPauseIterator struct {
	Event *PaymentProcessorContractPause // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorContractPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorContractPause)
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
		it.Event = new(PaymentProcessorContractPause)
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
func (it *PaymentProcessorContractPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorContractPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorContractPause represents a Pause event raised by the PaymentProcessorContract contract.
type PaymentProcessorContractPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_PaymentProcessorContract *PaymentProcessorContractFilterer) FilterPause(opts *bind.FilterOpts) (*PaymentProcessorContractPauseIterator, error) {

	logs, sub, err := _PaymentProcessorContract.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorContractPauseIterator{contract: _PaymentProcessorContract.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_PaymentProcessorContract *PaymentProcessorContractFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PaymentProcessorContractPause) (event.Subscription, error) {

	logs, sub, err := _PaymentProcessorContract.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorContractPause)
				if err := _PaymentProcessorContract.contract.UnpackLog(event, "Pause", log); err != nil {
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

// PaymentProcessorContractUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the PaymentProcessorContract contract.
type PaymentProcessorContractUnpauseIterator struct {
	Event *PaymentProcessorContractUnpause // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorContractUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorContractUnpause)
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
		it.Event = new(PaymentProcessorContractUnpause)
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
func (it *PaymentProcessorContractUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorContractUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorContractUnpause represents a Unpause event raised by the PaymentProcessorContract contract.
type PaymentProcessorContractUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_PaymentProcessorContract *PaymentProcessorContractFilterer) FilterUnpause(opts *bind.FilterOpts) (*PaymentProcessorContractUnpauseIterator, error) {

	logs, sub, err := _PaymentProcessorContract.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorContractUnpauseIterator{contract: _PaymentProcessorContract.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_PaymentProcessorContract *PaymentProcessorContractFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PaymentProcessorContractUnpause) (event.Subscription, error) {

	logs, sub, err := _PaymentProcessorContract.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorContractUnpause)
				if err := _PaymentProcessorContract.contract.UnpackLog(event, "Unpause", log); err != nil {
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

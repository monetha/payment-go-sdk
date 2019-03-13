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

// PrivatePaymentProcessorContractABI is the input ABI used to generate the binding from.
const PrivatePaymentProcessorContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isMonethaAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contactInformation\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PAYBACK_PERMILLE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawals\",\"outputs\":[{\"name\":\"state\",\"type\":\"uint8\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"clientAddress\",\"type\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"monethaGateway\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_info\",\"type\":\"string\"}],\"name\":\"setContactInformation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"merchantWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_isMonethaAddress\",\"type\":\"bool\"}],\"name\":\"setMonethaAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"merchantIdHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"destroyAndSend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_merchantId\",\"type\":\"string\"},{\"name\":\"_monethaGateway\",\"type\":\"address\"},{\"name\":\"_merchantWallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_originAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_monethaFee\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_discount\",\"type\":\"uint256\"}],\"name\":\"OrderPaidInEther\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_originAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_monethaFee\",\"type\":\"uint256\"}],\"name\":\"OrderPaidInToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_merchantAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"PaymentsProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_clientAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_refundReason\",\"type\":\"string\"}],\"name\":\"PaymentRefunding\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_clientAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaymentWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_isMonethaAddress\",\"type\":\"bool\"}],\"name\":\"MonethaAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\"},{\"name\":\"_originAddress\",\"type\":\"address\"},{\"name\":\"_monethaFee\",\"type\":\"uint256\"},{\"name\":\"_vouchersApply\",\"type\":\"uint256\"}],\"name\":\"payForOrder\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\"},{\"name\":\"_originAddress\",\"type\":\"address\"},{\"name\":\"_monethaFee\",\"type\":\"uint256\"},{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_orderValue\",\"type\":\"uint256\"}],\"name\":\"payForOrderInTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\"},{\"name\":\"_clientAddress\",\"type\":\"address\"},{\"name\":\"_refundReason\",\"type\":\"string\"}],\"name\":\"refundPayment\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\"},{\"name\":\"_clientAddress\",\"type\":\"address\"},{\"name\":\"_refundReason\",\"type\":\"string\"},{\"name\":\"_orderValue\",\"type\":\"uint256\"},{\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"refundTokenPayment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"withdrawRefund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\"},{\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"withdrawTokenRefund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newGateway\",\"type\":\"address\"}],\"name\":\"setMonethaGateway\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newWallet\",\"type\":\"address\"}],\"name\":\"setMerchantWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PrivatePaymentProcessorContractBin is the compiled bytecode used for deploying new contracts.
const PrivatePaymentProcessorContractBin = `0x60806040526000805460a060020a60ff02191690553480156200002157600080fd5b5060405162001ace38038062001ace833981016040908152815160208301519183015160008054600160a060020a03191633178155919093018051909391106200006a57600080fd5b826040516020018082805190602001908083835b602083106200009f5780518252601f1990920191602091820191016200007e565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b60208310620001045780518252601f199092019160209182019101620000e3565b5181516020939093036101000a600019018019909116921691909117905260405192018290039091206005555062000149915083905064010000000062000166810204565b6200015d81640100000000620001b6810204565b5050506200029e565b600054600160a060020a031633146200017e57600080fd5b600160a060020a03811615156200019457600080fd5b60038054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a03163314620001ce57600080fd5b600160a060020a0381161515620001e457600080fd5b6005546000191681600160a060020a031663f0daba016040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156200024357600080fd5b505af115801562000258573d6000803e3d6000fd5b505050506040513d60208110156200026f57600080fd5b5051146200027c57600080fd5b60048054600160a060020a031916600160a060020a0392909216919091179055565b61182080620002ae6000396000f3006080604052600436106101485763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631ee37896811461014d57806331d413251461016c57806331d4fad4146101a157806336f7ab5e146101c55780633f4ba83a1461024f5780634baf43e4146102645780635c975abb1461028b5780635cc07076146102a05780636137412c146102fd578063715018a61461032e57806383197ef0146103435780638456cb59146103585780638da5cb5b1461036d5780639600f294146103825780639d153495146103be5780639da30467146103d6578063b440bf3914610407578063b967a52e14610428578063bc85e06414610481578063c07e339114610496578063cee749bc146104bc578063ddda66db146104e0578063f0daba0114610501578063f2fde38b14610516578063f5074f4114610537575b600080fd5b61016a600435600160a060020a0360243516604435606435610558565b005b34801561017857600080fd5b5061018d600160a060020a0360043516610808565b604080519115158252519081900360200190f35b61016a600480359060248035600160a060020a03169160443591820191013561081d565b3480156101d157600080fd5b506101da6109c0565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102145781810151838201526020016101fc565b50505050905090810190601f1680156102415780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561025b57600080fd5b5061016a610a4d565b34801561027057600080fd5b50610279610ac3565b60408051918252519081900360200190f35b34801561029757600080fd5b5061018d610ac8565b3480156102ac57600080fd5b506102b8600435610ad8565b604051808560028111156102c857fe5b60ff168152602081019490945250600160a060020a039182166040808501919091529116606083015251908190036080019150f35b34801561030957600080fd5b50610312610b10565b60408051600160a060020a039092168252519081900360200190f35b34801561033a57600080fd5b5061016a610b1f565b34801561034f57600080fd5b5061016a610b8b565b34801561036457600080fd5b5061016a610bb0565b34801561037957600080fd5b50610312610c2b565b34801561038e57600080fd5b5061016a6004803590600160a060020a036024803582169260443591820192910135906064359060843516610c3a565b3480156103ca57600080fd5b5061016a600435610e8c565b3480156103e257600080fd5b5061016a600435600160a060020a036024358116906044359060643516608435610f86565b34801561041357600080fd5b5061016a600160a060020a036004351661132e565b34801561043457600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261016a9436949293602493928401919081908401838280828437509497506113899650505050505050565b34801561048d57600080fd5b506103126113b7565b3480156104a257600080fd5b5061016a600160a060020a036004351660243515156113c6565b3480156104c857600080fd5b5061016a600435600160a060020a0360243516611441565b3480156104ec57600080fd5b5061016a600160a060020a03600435166115a1565b34801561050d57600080fd5b50610279611690565b34801561052257600080fd5b5061016a600160a060020a0360043516611696565b34801561054357600080fd5b5061016a600160a060020a03600435166116b9565b60008054819060a060020a900460ff161561057257600080fd5b6000861161057f57600080fd5b600160a060020a038516151561059457600080fd5b600034116105a157600080fd5b60048054604080517f969596d60000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169263969596d69282820192602092908290030181600087803b1580156105fe57600080fd5b505af1158015610612573d6000803e3d6000fd5b505050506040513d602081101561062857600080fd5b5051915060009050600160a060020a038216156106fa57600354604080517fd21c39a1000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015260248201889052888116604483015260648201879052600260848301529151919092169163d21c39a191349160a48082019260209290919082900301818588803b1580156106c657600080fd5b505af11580156106da573d6000803e3d6000fd5b50505050506040513d60208110156106f157600080fd5b505190506107b3565b60035460048054604080517fd21c39a1000000000000000000000000000000000000000000000000000000008152600160a060020a039283169381019390935260248301889052888216604484015260648301879052600260848401525192169163d21c39a191349160a480830192602092919082900301818588803b15801561078357600080fd5b505af1158015610797573d6000803e3d6000fd5b50505050506040513d60208110156107ae57600080fd5b505190505b60408051348152602081018690528082018390529051600160a060020a0387169188917ffb3b61c9f30df0c105c673af85f07796032c77c3bf21f256bce83ba0853f630f9181900360600190a3505050505050565b60026020526000908152604090205460ff1681565b3360009081526002602052604090205460ff16151561083b57600080fd5b60005460a060020a900460ff161561085257600080fd5b6000841161085f57600080fd5b600160a060020a038316151561087457600080fd5b6000341161088157600080fd5b60008481526006602052604090205460ff16600281111561089e57fe5b156108a857600080fd5b6040805160808101909152806001815234602080830191909152600160a060020a038616604080840191909152600060609093018390528783526006909152902081518154829060ff1916600183600281111561090157fe5b0217905550602082810151600183015560408084015160028401805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a039384161790915560609586015160039095018054909116948216949094179093558051348082529281018281529181018690529287169388937f26e77179a69c2db5e1f39af4e228bc8c2205384ba14b8c1e3339049db4ee42c5939288928892919082018484808284376040519201829003965090945050505050a350505050565b60018054604080516020600284861615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610a455780601f10610a1a57610100808354040283529160200191610a45565b820191906000526020600020905b815481529060010190602001808311610a2857829003601f168201915b505050505081565b600054600160a060020a03163314610a6457600080fd5b60005460a060020a900460ff161515610a7c57600080fd5b6000805474ff0000000000000000000000000000000000000000191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a1565b600281565b60005460a060020a900460ff1681565b600660205260009081526040902080546001820154600283015460039093015460ff909216929091600160a060020a03918216911684565b600354600160a060020a031681565b600054600160a060020a03163314610b3657600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a03163314610ba257600080fd5b600054600160a060020a0316ff5b600054600160a060020a03163314610bc757600080fd5b60005460a060020a900460ff1615610bde57600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a1565b600054600160a060020a031681565b3360009081526002602052604090205460ff161515610c5857600080fd5b60005460a060020a900460ff1615610c6f57600080fd5b60008611610c7c57600080fd5b600160a060020a0385161515610c9157600080fd5b60008211610c9e57600080fd5b600160a060020a0381161515610cb357600080fd5b60008681526006602052604090205460ff166002811115610cd057fe5b15610cda57600080fd5b604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018490529051600160a060020a038316916323b872dd9160648083019260209291908290030181600087803b158015610d4857600080fd5b505af1158015610d5c573d6000803e3d6000fd5b505050506040513d6020811015610d7257600080fd5b50506040805160808101825260018082526020808301869052600160a060020a03898116848601528516606084015260008a815260069091529290922081518154929391929091839160ff191690836002811115610dcc57fe5b0217905550602082810151600183015560408084015160028401805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a0393841617909155606095860151600390950180549091169482169490941790935580518681529182018181529082018790529188169289927f26e77179a69c2db5e1f39af4e228bc8c2205384ba14b8c1e3339049db4ee42c59287928a928a929182018484808284376040519201829003965090945050505050a3505050505050565b600080548190819060a060020a900460ff1615610ea857600080fd5b6000848152600660205260409020805490935060ff166002811115610ec957fe5b600114610ed557600080fd5b6003830154600160a060020a031615610eed57600080fd5b50506002818101546001830154835460ff19169092178355604051600160a060020a03909116919082906108fc8315029083906000818181858888f19350505050158015610f3f573d6000803e3d6000fd5b50604080518281529051600160a060020a0384169186917fbe85bf3b0a1e335a22c461f84cf759dfe589ec1539caf4dce60f999d72dd8e239181900360200190a350505050565b6000805460a060020a900460ff1615610f9e57600080fd5b60008611610fab57600080fd5b600160a060020a0385161515610fc057600080fd5b60008211610fcd57600080fd5b600160a060020a0383161515610fe257600080fd5b60048054604080517f969596d60000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169263969596d69282820192602092908290030181600087803b15801561103f57600080fd5b505af1158015611053573d6000803e3d6000fd5b505050506040513d602081101561106957600080fd5b5051604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590529051919250600160a060020a038516916323b872dd916064808201926020929091908290030181600087803b1580156110dd57600080fd5b505af11580156110f1573d6000803e3d6000fd5b505050506040513d602081101561110757600080fd5b5050600354604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a0392831660048201526024810185905290519185169163a9059cbb9160448082019260009290919082900301818387803b15801561117757600080fd5b505af115801561118b573d6000803e3d6000fd5b50505050600160a060020a0381161561123857600354604080517fcc4fbc43000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152602482018890528681166044830152606482018690529151919092169163cc4fbc4391608480830192600092919082900301818387803b15801561121b57600080fd5b505af115801561122f573d6000803e3d6000fd5b505050506112d2565b60035460048054604080517fcc4fbc43000000000000000000000000000000000000000000000000000000008152600160a060020a0392831693810193909352602483018890528682166044840152606483018690525192169163cc4fbc439160848082019260009290919082900301818387803b1580156112b957600080fd5b505af11580156112cd573d6000803e3d6000fd5b505050505b82600160a060020a031685600160a060020a0316877fd6dc98331ad06baebe39c90f4fd554341ad121d55e4384bd046def391501a00f8588604051808381526020018281526020019250505060405180910390a4505050505050565b600054600160a060020a0316331461134557600080fd5b600160a060020a038116151561135a57600080fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a031633146113a057600080fd5b80516113b3906001906020840190611759565b5050565b600454600160a060020a031681565b600054600160a060020a031633146113dd57600080fd5b600160a060020a038216600081815260026020908152604091829020805460ff191685151590811790915582519384529083015280517fa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e6929281900390910190a15050565b600080548190819060a060020a900460ff161561145d57600080fd5b600160a060020a038416151561147257600080fd5b6000858152600660205260409020805490935060ff16600281111561149357fe5b60011461149f57600080fd5b6003830154600160a060020a038581169116146114bb57600080fd5b50506002818101546001830154835460ff19169092178355604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a039283166004820181905260248201859052915191939286169163a9059cbb9160448082019260009290919082900301818387803b15801561154257600080fd5b505af1158015611556573d6000803e3d6000fd5b5050604080518481529051600160a060020a03861693508892507fbe85bf3b0a1e335a22c461f84cf759dfe589ec1539caf4dce60f999d72dd8e239181900360200190a35050505050565b600054600160a060020a031633146115b857600080fd5b600160a060020a03811615156115cd57600080fd5b6005546000191681600160a060020a031663f0daba016040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561162b57600080fd5b505af115801561163f573d6000803e3d6000fd5b505050506040513d602081101561165557600080fd5b50511461166157600080fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60055481565b600054600160a060020a031633146116ad57600080fd5b6116b6816116dc565b50565b600054600160a060020a031633146116d057600080fd5b80600160a060020a0316ff5b600160a060020a03811615156116f157600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061179a57805160ff19168380011785556117c7565b828001600101855582156117c7579182015b828111156117c75782518255916020019190600101906117ac565b506117d39291506117d7565b5090565b6117f191905b808211156117d357600081556001016117dd565b905600a165627a7a72305820b5debb2623b86cd4f7ffd25775500238675501750ea0fe8648caeb063ff4d1ba0029`

// DeployPrivatePaymentProcessorContract deploys a new Ethereum contract, binding an instance of PrivatePaymentProcessorContract to it.
func DeployPrivatePaymentProcessorContract(auth *bind.TransactOpts, backend bind.ContractBackend, _merchantId string, _monethaGateway common.Address, _merchantWallet common.Address) (common.Address, *types.Transaction, *PrivatePaymentProcessorContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PrivatePaymentProcessorContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PrivatePaymentProcessorContractBin), backend, _merchantId, _monethaGateway, _merchantWallet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PrivatePaymentProcessorContract{PrivatePaymentProcessorContractCaller: PrivatePaymentProcessorContractCaller{contract: contract}, PrivatePaymentProcessorContractTransactor: PrivatePaymentProcessorContractTransactor{contract: contract}, PrivatePaymentProcessorContractFilterer: PrivatePaymentProcessorContractFilterer{contract: contract}}, nil
}

// PrivatePaymentProcessorContract is an auto generated Go binding around an Ethereum contract.
type PrivatePaymentProcessorContract struct {
	PrivatePaymentProcessorContractCaller     // Read-only binding to the contract
	PrivatePaymentProcessorContractTransactor // Write-only binding to the contract
	PrivatePaymentProcessorContractFilterer   // Log filterer for contract events
}

// PrivatePaymentProcessorContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrivatePaymentProcessorContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivatePaymentProcessorContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrivatePaymentProcessorContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivatePaymentProcessorContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrivatePaymentProcessorContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivatePaymentProcessorContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrivatePaymentProcessorContractSession struct {
	Contract     *PrivatePaymentProcessorContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// PrivatePaymentProcessorContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrivatePaymentProcessorContractCallerSession struct {
	Contract *PrivatePaymentProcessorContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// PrivatePaymentProcessorContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrivatePaymentProcessorContractTransactorSession struct {
	Contract     *PrivatePaymentProcessorContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// PrivatePaymentProcessorContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrivatePaymentProcessorContractRaw struct {
	Contract *PrivatePaymentProcessorContract // Generic contract binding to access the raw methods on
}

// PrivatePaymentProcessorContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrivatePaymentProcessorContractCallerRaw struct {
	Contract *PrivatePaymentProcessorContractCaller // Generic read-only contract binding to access the raw methods on
}

// PrivatePaymentProcessorContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrivatePaymentProcessorContractTransactorRaw struct {
	Contract *PrivatePaymentProcessorContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrivatePaymentProcessorContract creates a new instance of PrivatePaymentProcessorContract, bound to a specific deployed contract.
func NewPrivatePaymentProcessorContract(address common.Address, backend bind.ContractBackend) (*PrivatePaymentProcessorContract, error) {
	contract, err := bindPrivatePaymentProcessorContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrivatePaymentProcessorContract{PrivatePaymentProcessorContractCaller: PrivatePaymentProcessorContractCaller{contract: contract}, PrivatePaymentProcessorContractTransactor: PrivatePaymentProcessorContractTransactor{contract: contract}, PrivatePaymentProcessorContractFilterer: PrivatePaymentProcessorContractFilterer{contract: contract}}, nil
}

// NewPrivatePaymentProcessorContractCaller creates a new read-only instance of PrivatePaymentProcessorContract, bound to a specific deployed contract.
func NewPrivatePaymentProcessorContractCaller(address common.Address, caller bind.ContractCaller) (*PrivatePaymentProcessorContractCaller, error) {
	contract, err := bindPrivatePaymentProcessorContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrivatePaymentProcessorContractCaller{contract: contract}, nil
}

// NewPrivatePaymentProcessorContractTransactor creates a new write-only instance of PrivatePaymentProcessorContract, bound to a specific deployed contract.
func NewPrivatePaymentProcessorContractTransactor(address common.Address, transactor bind.ContractTransactor) (*PrivatePaymentProcessorContractTransactor, error) {
	contract, err := bindPrivatePaymentProcessorContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrivatePaymentProcessorContractTransactor{contract: contract}, nil
}

// NewPrivatePaymentProcessorContractFilterer creates a new log filterer instance of PrivatePaymentProcessorContract, bound to a specific deployed contract.
func NewPrivatePaymentProcessorContractFilterer(address common.Address, filterer bind.ContractFilterer) (*PrivatePaymentProcessorContractFilterer, error) {
	contract, err := bindPrivatePaymentProcessorContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrivatePaymentProcessorContractFilterer{contract: contract}, nil
}

// bindPrivatePaymentProcessorContract binds a generic wrapper to an already deployed contract.
func bindPrivatePaymentProcessorContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PrivatePaymentProcessorContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PrivatePaymentProcessorContract.Contract.PrivatePaymentProcessorContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.PrivatePaymentProcessorContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.PrivatePaymentProcessorContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PrivatePaymentProcessorContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.contract.Transact(opts, method, params...)
}

// PAYBACKPERMILLE is a free data retrieval call binding the contract method 0x4baf43e4.
//
// Solidity: function PAYBACK_PERMILLE() constant returns(uint256)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractCaller) PAYBACKPERMILLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivatePaymentProcessorContract.contract.Call(opts, out, "PAYBACK_PERMILLE")
	return *ret0, err
}

// PAYBACKPERMILLE is a free data retrieval call binding the contract method 0x4baf43e4.
//
// Solidity: function PAYBACK_PERMILLE() constant returns(uint256)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) PAYBACKPERMILLE() (*big.Int, error) {
	return _PrivatePaymentProcessorContract.Contract.PAYBACKPERMILLE(&_PrivatePaymentProcessorContract.CallOpts)
}

// PAYBACKPERMILLE is a free data retrieval call binding the contract method 0x4baf43e4.
//
// Solidity: function PAYBACK_PERMILLE() constant returns(uint256)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractCallerSession) PAYBACKPERMILLE() (*big.Int, error) {
	return _PrivatePaymentProcessorContract.Contract.PAYBACKPERMILLE(&_PrivatePaymentProcessorContract.CallOpts)
}

// ContactInformation is a free data retrieval call binding the contract method 0x36f7ab5e.
//
// Solidity: function contactInformation() constant returns(string)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractCaller) ContactInformation(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PrivatePaymentProcessorContract.contract.Call(opts, out, "contactInformation")
	return *ret0, err
}

// ContactInformation is a free data retrieval call binding the contract method 0x36f7ab5e.
//
// Solidity: function contactInformation() constant returns(string)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) ContactInformation() (string, error) {
	return _PrivatePaymentProcessorContract.Contract.ContactInformation(&_PrivatePaymentProcessorContract.CallOpts)
}

// ContactInformation is a free data retrieval call binding the contract method 0x36f7ab5e.
//
// Solidity: function contactInformation() constant returns(string)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractCallerSession) ContactInformation() (string, error) {
	return _PrivatePaymentProcessorContract.Contract.ContactInformation(&_PrivatePaymentProcessorContract.CallOpts)
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractCaller) IsMonethaAddress(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PrivatePaymentProcessorContract.contract.Call(opts, out, "isMonethaAddress", arg0)
	return *ret0, err
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) IsMonethaAddress(arg0 common.Address) (bool, error) {
	return _PrivatePaymentProcessorContract.Contract.IsMonethaAddress(&_PrivatePaymentProcessorContract.CallOpts, arg0)
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractCallerSession) IsMonethaAddress(arg0 common.Address) (bool, error) {
	return _PrivatePaymentProcessorContract.Contract.IsMonethaAddress(&_PrivatePaymentProcessorContract.CallOpts, arg0)
}

// MerchantIdHash is a free data retrieval call binding the contract method 0xf0daba01.
//
// Solidity: function merchantIdHash() constant returns(bytes32)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractCaller) MerchantIdHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PrivatePaymentProcessorContract.contract.Call(opts, out, "merchantIdHash")
	return *ret0, err
}

// MerchantIdHash is a free data retrieval call binding the contract method 0xf0daba01.
//
// Solidity: function merchantIdHash() constant returns(bytes32)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) MerchantIdHash() ([32]byte, error) {
	return _PrivatePaymentProcessorContract.Contract.MerchantIdHash(&_PrivatePaymentProcessorContract.CallOpts)
}

// MerchantIdHash is a free data retrieval call binding the contract method 0xf0daba01.
//
// Solidity: function merchantIdHash() constant returns(bytes32)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractCallerSession) MerchantIdHash() ([32]byte, error) {
	return _PrivatePaymentProcessorContract.Contract.MerchantIdHash(&_PrivatePaymentProcessorContract.CallOpts)
}

// MerchantWallet is a free data retrieval call binding the contract method 0xbc85e064.
//
// Solidity: function merchantWallet() constant returns(address)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractCaller) MerchantWallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PrivatePaymentProcessorContract.contract.Call(opts, out, "merchantWallet")
	return *ret0, err
}

// MerchantWallet is a free data retrieval call binding the contract method 0xbc85e064.
//
// Solidity: function merchantWallet() constant returns(address)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) MerchantWallet() (common.Address, error) {
	return _PrivatePaymentProcessorContract.Contract.MerchantWallet(&_PrivatePaymentProcessorContract.CallOpts)
}

// MerchantWallet is a free data retrieval call binding the contract method 0xbc85e064.
//
// Solidity: function merchantWallet() constant returns(address)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractCallerSession) MerchantWallet() (common.Address, error) {
	return _PrivatePaymentProcessorContract.Contract.MerchantWallet(&_PrivatePaymentProcessorContract.CallOpts)
}

// MonethaGateway is a free data retrieval call binding the contract method 0x6137412c.
//
// Solidity: function monethaGateway() constant returns(address)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractCaller) MonethaGateway(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PrivatePaymentProcessorContract.contract.Call(opts, out, "monethaGateway")
	return *ret0, err
}

// MonethaGateway is a free data retrieval call binding the contract method 0x6137412c.
//
// Solidity: function monethaGateway() constant returns(address)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) MonethaGateway() (common.Address, error) {
	return _PrivatePaymentProcessorContract.Contract.MonethaGateway(&_PrivatePaymentProcessorContract.CallOpts)
}

// MonethaGateway is a free data retrieval call binding the contract method 0x6137412c.
//
// Solidity: function monethaGateway() constant returns(address)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractCallerSession) MonethaGateway() (common.Address, error) {
	return _PrivatePaymentProcessorContract.Contract.MonethaGateway(&_PrivatePaymentProcessorContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PrivatePaymentProcessorContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) Owner() (common.Address, error) {
	return _PrivatePaymentProcessorContract.Contract.Owner(&_PrivatePaymentProcessorContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractCallerSession) Owner() (common.Address, error) {
	return _PrivatePaymentProcessorContract.Contract.Owner(&_PrivatePaymentProcessorContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PrivatePaymentProcessorContract.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) Paused() (bool, error) {
	return _PrivatePaymentProcessorContract.Contract.Paused(&_PrivatePaymentProcessorContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractCallerSession) Paused() (bool, error) {
	return _PrivatePaymentProcessorContract.Contract.Paused(&_PrivatePaymentProcessorContract.CallOpts)
}

// Withdrawals is a free data retrieval call binding the contract method 0x5cc07076.
//
// Solidity: function withdrawals( uint256) constant returns(state uint8, amount uint256, clientAddress address, tokenAddress address)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractCaller) Withdrawals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	State         uint8
	Amount        *big.Int
	ClientAddress common.Address
	TokenAddress  common.Address
}, error) {
	ret := new(struct {
		State         uint8
		Amount        *big.Int
		ClientAddress common.Address
		TokenAddress  common.Address
	})
	out := ret
	err := _PrivatePaymentProcessorContract.contract.Call(opts, out, "withdrawals", arg0)
	return *ret, err
}

// Withdrawals is a free data retrieval call binding the contract method 0x5cc07076.
//
// Solidity: function withdrawals( uint256) constant returns(state uint8, amount uint256, clientAddress address, tokenAddress address)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) Withdrawals(arg0 *big.Int) (struct {
	State         uint8
	Amount        *big.Int
	ClientAddress common.Address
	TokenAddress  common.Address
}, error) {
	return _PrivatePaymentProcessorContract.Contract.Withdrawals(&_PrivatePaymentProcessorContract.CallOpts, arg0)
}

// Withdrawals is a free data retrieval call binding the contract method 0x5cc07076.
//
// Solidity: function withdrawals( uint256) constant returns(state uint8, amount uint256, clientAddress address, tokenAddress address)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractCallerSession) Withdrawals(arg0 *big.Int) (struct {
	State         uint8
	Amount        *big.Int
	ClientAddress common.Address
	TokenAddress  common.Address
}, error) {
	return _PrivatePaymentProcessorContract.Contract.Withdrawals(&_PrivatePaymentProcessorContract.CallOpts, arg0)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) Destroy() (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.Destroy(&_PrivatePaymentProcessorContract.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactorSession) Destroy() (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.Destroy(&_PrivatePaymentProcessorContract.TransactOpts)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactor) DestroyAndSend(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.contract.Transact(opts, "destroyAndSend", _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.DestroyAndSend(&_PrivatePaymentProcessorContract.TransactOpts, _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactorSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.DestroyAndSend(&_PrivatePaymentProcessorContract.TransactOpts, _recipient)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) Pause() (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.Pause(&_PrivatePaymentProcessorContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactorSession) Pause() (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.Pause(&_PrivatePaymentProcessorContract.TransactOpts)
}

// PayForOrder is a paid mutator transaction binding the contract method 0x1ee37896.
//
// Solidity: function payForOrder(_orderId uint256, _originAddress address, _monethaFee uint256, _vouchersApply uint256) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactor) PayForOrder(opts *bind.TransactOpts, _orderId *big.Int, _originAddress common.Address, _monethaFee *big.Int, _vouchersApply *big.Int) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.contract.Transact(opts, "payForOrder", _orderId, _originAddress, _monethaFee, _vouchersApply)
}

// PayForOrder is a paid mutator transaction binding the contract method 0x1ee37896.
//
// Solidity: function payForOrder(_orderId uint256, _originAddress address, _monethaFee uint256, _vouchersApply uint256) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) PayForOrder(_orderId *big.Int, _originAddress common.Address, _monethaFee *big.Int, _vouchersApply *big.Int) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.PayForOrder(&_PrivatePaymentProcessorContract.TransactOpts, _orderId, _originAddress, _monethaFee, _vouchersApply)
}

// PayForOrder is a paid mutator transaction binding the contract method 0x1ee37896.
//
// Solidity: function payForOrder(_orderId uint256, _originAddress address, _monethaFee uint256, _vouchersApply uint256) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactorSession) PayForOrder(_orderId *big.Int, _originAddress common.Address, _monethaFee *big.Int, _vouchersApply *big.Int) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.PayForOrder(&_PrivatePaymentProcessorContract.TransactOpts, _orderId, _originAddress, _monethaFee, _vouchersApply)
}

// PayForOrderInTokens is a paid mutator transaction binding the contract method 0x9da30467.
//
// Solidity: function payForOrderInTokens(_orderId uint256, _originAddress address, _monethaFee uint256, _tokenAddress address, _orderValue uint256) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactor) PayForOrderInTokens(opts *bind.TransactOpts, _orderId *big.Int, _originAddress common.Address, _monethaFee *big.Int, _tokenAddress common.Address, _orderValue *big.Int) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.contract.Transact(opts, "payForOrderInTokens", _orderId, _originAddress, _monethaFee, _tokenAddress, _orderValue)
}

// PayForOrderInTokens is a paid mutator transaction binding the contract method 0x9da30467.
//
// Solidity: function payForOrderInTokens(_orderId uint256, _originAddress address, _monethaFee uint256, _tokenAddress address, _orderValue uint256) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) PayForOrderInTokens(_orderId *big.Int, _originAddress common.Address, _monethaFee *big.Int, _tokenAddress common.Address, _orderValue *big.Int) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.PayForOrderInTokens(&_PrivatePaymentProcessorContract.TransactOpts, _orderId, _originAddress, _monethaFee, _tokenAddress, _orderValue)
}

// PayForOrderInTokens is a paid mutator transaction binding the contract method 0x9da30467.
//
// Solidity: function payForOrderInTokens(_orderId uint256, _originAddress address, _monethaFee uint256, _tokenAddress address, _orderValue uint256) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactorSession) PayForOrderInTokens(_orderId *big.Int, _originAddress common.Address, _monethaFee *big.Int, _tokenAddress common.Address, _orderValue *big.Int) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.PayForOrderInTokens(&_PrivatePaymentProcessorContract.TransactOpts, _orderId, _originAddress, _monethaFee, _tokenAddress, _orderValue)
}

// RefundPayment is a paid mutator transaction binding the contract method 0x31d4fad4.
//
// Solidity: function refundPayment(_orderId uint256, _clientAddress address, _refundReason string) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactor) RefundPayment(opts *bind.TransactOpts, _orderId *big.Int, _clientAddress common.Address, _refundReason string) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.contract.Transact(opts, "refundPayment", _orderId, _clientAddress, _refundReason)
}

// RefundPayment is a paid mutator transaction binding the contract method 0x31d4fad4.
//
// Solidity: function refundPayment(_orderId uint256, _clientAddress address, _refundReason string) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) RefundPayment(_orderId *big.Int, _clientAddress common.Address, _refundReason string) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.RefundPayment(&_PrivatePaymentProcessorContract.TransactOpts, _orderId, _clientAddress, _refundReason)
}

// RefundPayment is a paid mutator transaction binding the contract method 0x31d4fad4.
//
// Solidity: function refundPayment(_orderId uint256, _clientAddress address, _refundReason string) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactorSession) RefundPayment(_orderId *big.Int, _clientAddress common.Address, _refundReason string) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.RefundPayment(&_PrivatePaymentProcessorContract.TransactOpts, _orderId, _clientAddress, _refundReason)
}

// RefundTokenPayment is a paid mutator transaction binding the contract method 0x9600f294.
//
// Solidity: function refundTokenPayment(_orderId uint256, _clientAddress address, _refundReason string, _orderValue uint256, _tokenAddress address) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactor) RefundTokenPayment(opts *bind.TransactOpts, _orderId *big.Int, _clientAddress common.Address, _refundReason string, _orderValue *big.Int, _tokenAddress common.Address) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.contract.Transact(opts, "refundTokenPayment", _orderId, _clientAddress, _refundReason, _orderValue, _tokenAddress)
}

// RefundTokenPayment is a paid mutator transaction binding the contract method 0x9600f294.
//
// Solidity: function refundTokenPayment(_orderId uint256, _clientAddress address, _refundReason string, _orderValue uint256, _tokenAddress address) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) RefundTokenPayment(_orderId *big.Int, _clientAddress common.Address, _refundReason string, _orderValue *big.Int, _tokenAddress common.Address) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.RefundTokenPayment(&_PrivatePaymentProcessorContract.TransactOpts, _orderId, _clientAddress, _refundReason, _orderValue, _tokenAddress)
}

// RefundTokenPayment is a paid mutator transaction binding the contract method 0x9600f294.
//
// Solidity: function refundTokenPayment(_orderId uint256, _clientAddress address, _refundReason string, _orderValue uint256, _tokenAddress address) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactorSession) RefundTokenPayment(_orderId *big.Int, _clientAddress common.Address, _refundReason string, _orderValue *big.Int, _tokenAddress common.Address) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.RefundTokenPayment(&_PrivatePaymentProcessorContract.TransactOpts, _orderId, _clientAddress, _refundReason, _orderValue, _tokenAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.RenounceOwnership(&_PrivatePaymentProcessorContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.RenounceOwnership(&_PrivatePaymentProcessorContract.TransactOpts)
}

// SetContactInformation is a paid mutator transaction binding the contract method 0xb967a52e.
//
// Solidity: function setContactInformation(_info string) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactor) SetContactInformation(opts *bind.TransactOpts, _info string) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.contract.Transact(opts, "setContactInformation", _info)
}

// SetContactInformation is a paid mutator transaction binding the contract method 0xb967a52e.
//
// Solidity: function setContactInformation(_info string) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) SetContactInformation(_info string) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.SetContactInformation(&_PrivatePaymentProcessorContract.TransactOpts, _info)
}

// SetContactInformation is a paid mutator transaction binding the contract method 0xb967a52e.
//
// Solidity: function setContactInformation(_info string) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactorSession) SetContactInformation(_info string) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.SetContactInformation(&_PrivatePaymentProcessorContract.TransactOpts, _info)
}

// SetMerchantWallet is a paid mutator transaction binding the contract method 0xddda66db.
//
// Solidity: function setMerchantWallet(_newWallet address) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactor) SetMerchantWallet(opts *bind.TransactOpts, _newWallet common.Address) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.contract.Transact(opts, "setMerchantWallet", _newWallet)
}

// SetMerchantWallet is a paid mutator transaction binding the contract method 0xddda66db.
//
// Solidity: function setMerchantWallet(_newWallet address) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) SetMerchantWallet(_newWallet common.Address) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.SetMerchantWallet(&_PrivatePaymentProcessorContract.TransactOpts, _newWallet)
}

// SetMerchantWallet is a paid mutator transaction binding the contract method 0xddda66db.
//
// Solidity: function setMerchantWallet(_newWallet address) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactorSession) SetMerchantWallet(_newWallet common.Address) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.SetMerchantWallet(&_PrivatePaymentProcessorContract.TransactOpts, _newWallet)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactor) SetMonethaAddress(opts *bind.TransactOpts, _address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.contract.Transact(opts, "setMonethaAddress", _address, _isMonethaAddress)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) SetMonethaAddress(_address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.SetMonethaAddress(&_PrivatePaymentProcessorContract.TransactOpts, _address, _isMonethaAddress)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactorSession) SetMonethaAddress(_address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.SetMonethaAddress(&_PrivatePaymentProcessorContract.TransactOpts, _address, _isMonethaAddress)
}

// SetMonethaGateway is a paid mutator transaction binding the contract method 0xb440bf39.
//
// Solidity: function setMonethaGateway(_newGateway address) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactor) SetMonethaGateway(opts *bind.TransactOpts, _newGateway common.Address) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.contract.Transact(opts, "setMonethaGateway", _newGateway)
}

// SetMonethaGateway is a paid mutator transaction binding the contract method 0xb440bf39.
//
// Solidity: function setMonethaGateway(_newGateway address) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) SetMonethaGateway(_newGateway common.Address) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.SetMonethaGateway(&_PrivatePaymentProcessorContract.TransactOpts, _newGateway)
}

// SetMonethaGateway is a paid mutator transaction binding the contract method 0xb440bf39.
//
// Solidity: function setMonethaGateway(_newGateway address) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactorSession) SetMonethaGateway(_newGateway common.Address) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.SetMonethaGateway(&_PrivatePaymentProcessorContract.TransactOpts, _newGateway)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.TransferOwnership(&_PrivatePaymentProcessorContract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.TransferOwnership(&_PrivatePaymentProcessorContract.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) Unpause() (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.Unpause(&_PrivatePaymentProcessorContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.Unpause(&_PrivatePaymentProcessorContract.TransactOpts)
}

// WithdrawRefund is a paid mutator transaction binding the contract method 0x9d153495.
//
// Solidity: function withdrawRefund(_orderId uint256) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactor) WithdrawRefund(opts *bind.TransactOpts, _orderId *big.Int) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.contract.Transact(opts, "withdrawRefund", _orderId)
}

// WithdrawRefund is a paid mutator transaction binding the contract method 0x9d153495.
//
// Solidity: function withdrawRefund(_orderId uint256) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) WithdrawRefund(_orderId *big.Int) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.WithdrawRefund(&_PrivatePaymentProcessorContract.TransactOpts, _orderId)
}

// WithdrawRefund is a paid mutator transaction binding the contract method 0x9d153495.
//
// Solidity: function withdrawRefund(_orderId uint256) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactorSession) WithdrawRefund(_orderId *big.Int) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.WithdrawRefund(&_PrivatePaymentProcessorContract.TransactOpts, _orderId)
}

// WithdrawTokenRefund is a paid mutator transaction binding the contract method 0xcee749bc.
//
// Solidity: function withdrawTokenRefund(_orderId uint256, _tokenAddress address) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactor) WithdrawTokenRefund(opts *bind.TransactOpts, _orderId *big.Int, _tokenAddress common.Address) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.contract.Transact(opts, "withdrawTokenRefund", _orderId, _tokenAddress)
}

// WithdrawTokenRefund is a paid mutator transaction binding the contract method 0xcee749bc.
//
// Solidity: function withdrawTokenRefund(_orderId uint256, _tokenAddress address) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractSession) WithdrawTokenRefund(_orderId *big.Int, _tokenAddress common.Address) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.WithdrawTokenRefund(&_PrivatePaymentProcessorContract.TransactOpts, _orderId, _tokenAddress)
}

// WithdrawTokenRefund is a paid mutator transaction binding the contract method 0xcee749bc.
//
// Solidity: function withdrawTokenRefund(_orderId uint256, _tokenAddress address) returns()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractTransactorSession) WithdrawTokenRefund(_orderId *big.Int, _tokenAddress common.Address) (*types.Transaction, error) {
	return _PrivatePaymentProcessorContract.Contract.WithdrawTokenRefund(&_PrivatePaymentProcessorContract.TransactOpts, _orderId, _tokenAddress)
}

// PrivatePaymentProcessorContractMonethaAddressSetIterator is returned from FilterMonethaAddressSet and is used to iterate over the raw logs and unpacked data for MonethaAddressSet events raised by the PrivatePaymentProcessorContract contract.
type PrivatePaymentProcessorContractMonethaAddressSetIterator struct {
	Event *PrivatePaymentProcessorContractMonethaAddressSet // Event containing the contract specifics and raw log

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
func (it *PrivatePaymentProcessorContractMonethaAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivatePaymentProcessorContractMonethaAddressSet)
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
		it.Event = new(PrivatePaymentProcessorContractMonethaAddressSet)
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
func (it *PrivatePaymentProcessorContractMonethaAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivatePaymentProcessorContractMonethaAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivatePaymentProcessorContractMonethaAddressSet represents a MonethaAddressSet event raised by the PrivatePaymentProcessorContract contract.
type PrivatePaymentProcessorContractMonethaAddressSet struct {
	Address          common.Address
	IsMonethaAddress bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMonethaAddressSet is a free log retrieval operation binding the contract event 0xa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e692.
//
// Solidity: e MonethaAddressSet(_address address, _isMonethaAddress bool)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractFilterer) FilterMonethaAddressSet(opts *bind.FilterOpts) (*PrivatePaymentProcessorContractMonethaAddressSetIterator, error) {

	logs, sub, err := _PrivatePaymentProcessorContract.contract.FilterLogs(opts, "MonethaAddressSet")
	if err != nil {
		return nil, err
	}
	return &PrivatePaymentProcessorContractMonethaAddressSetIterator{contract: _PrivatePaymentProcessorContract.contract, event: "MonethaAddressSet", logs: logs, sub: sub}, nil
}

// WatchMonethaAddressSet is a free log subscription operation binding the contract event 0xa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e692.
//
// Solidity: e MonethaAddressSet(_address address, _isMonethaAddress bool)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractFilterer) WatchMonethaAddressSet(opts *bind.WatchOpts, sink chan<- *PrivatePaymentProcessorContractMonethaAddressSet) (event.Subscription, error) {

	logs, sub, err := _PrivatePaymentProcessorContract.contract.WatchLogs(opts, "MonethaAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivatePaymentProcessorContractMonethaAddressSet)
				if err := _PrivatePaymentProcessorContract.contract.UnpackLog(event, "MonethaAddressSet", log); err != nil {
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

// PrivatePaymentProcessorContractOrderPaidInEtherIterator is returned from FilterOrderPaidInEther and is used to iterate over the raw logs and unpacked data for OrderPaidInEther events raised by the PrivatePaymentProcessorContract contract.
type PrivatePaymentProcessorContractOrderPaidInEtherIterator struct {
	Event *PrivatePaymentProcessorContractOrderPaidInEther // Event containing the contract specifics and raw log

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
func (it *PrivatePaymentProcessorContractOrderPaidInEtherIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivatePaymentProcessorContractOrderPaidInEther)
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
		it.Event = new(PrivatePaymentProcessorContractOrderPaidInEther)
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
func (it *PrivatePaymentProcessorContractOrderPaidInEtherIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivatePaymentProcessorContractOrderPaidInEtherIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivatePaymentProcessorContractOrderPaidInEther represents a OrderPaidInEther event raised by the PrivatePaymentProcessorContract contract.
type PrivatePaymentProcessorContractOrderPaidInEther struct {
	OrderId       *big.Int
	OriginAddress common.Address
	Price         *big.Int
	MonethaFee    *big.Int
	Discount      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOrderPaidInEther is a free log retrieval operation binding the contract event 0xfb3b61c9f30df0c105c673af85f07796032c77c3bf21f256bce83ba0853f630f.
//
// Solidity: e OrderPaidInEther(_orderId indexed uint256, _originAddress indexed address, _price uint256, _monethaFee uint256, _discount uint256)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractFilterer) FilterOrderPaidInEther(opts *bind.FilterOpts, _orderId []*big.Int, _originAddress []common.Address) (*PrivatePaymentProcessorContractOrderPaidInEtherIterator, error) {

	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}
	var _originAddressRule []interface{}
	for _, _originAddressItem := range _originAddress {
		_originAddressRule = append(_originAddressRule, _originAddressItem)
	}

	logs, sub, err := _PrivatePaymentProcessorContract.contract.FilterLogs(opts, "OrderPaidInEther", _orderIdRule, _originAddressRule)
	if err != nil {
		return nil, err
	}
	return &PrivatePaymentProcessorContractOrderPaidInEtherIterator{contract: _PrivatePaymentProcessorContract.contract, event: "OrderPaidInEther", logs: logs, sub: sub}, nil
}

// WatchOrderPaidInEther is a free log subscription operation binding the contract event 0xfb3b61c9f30df0c105c673af85f07796032c77c3bf21f256bce83ba0853f630f.
//
// Solidity: e OrderPaidInEther(_orderId indexed uint256, _originAddress indexed address, _price uint256, _monethaFee uint256, _discount uint256)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractFilterer) WatchOrderPaidInEther(opts *bind.WatchOpts, sink chan<- *PrivatePaymentProcessorContractOrderPaidInEther, _orderId []*big.Int, _originAddress []common.Address) (event.Subscription, error) {

	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}
	var _originAddressRule []interface{}
	for _, _originAddressItem := range _originAddress {
		_originAddressRule = append(_originAddressRule, _originAddressItem)
	}

	logs, sub, err := _PrivatePaymentProcessorContract.contract.WatchLogs(opts, "OrderPaidInEther", _orderIdRule, _originAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivatePaymentProcessorContractOrderPaidInEther)
				if err := _PrivatePaymentProcessorContract.contract.UnpackLog(event, "OrderPaidInEther", log); err != nil {
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

// PrivatePaymentProcessorContractOrderPaidInTokenIterator is returned from FilterOrderPaidInToken and is used to iterate over the raw logs and unpacked data for OrderPaidInToken events raised by the PrivatePaymentProcessorContract contract.
type PrivatePaymentProcessorContractOrderPaidInTokenIterator struct {
	Event *PrivatePaymentProcessorContractOrderPaidInToken // Event containing the contract specifics and raw log

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
func (it *PrivatePaymentProcessorContractOrderPaidInTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivatePaymentProcessorContractOrderPaidInToken)
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
		it.Event = new(PrivatePaymentProcessorContractOrderPaidInToken)
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
func (it *PrivatePaymentProcessorContractOrderPaidInTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivatePaymentProcessorContractOrderPaidInTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivatePaymentProcessorContractOrderPaidInToken represents a OrderPaidInToken event raised by the PrivatePaymentProcessorContract contract.
type PrivatePaymentProcessorContractOrderPaidInToken struct {
	OrderId       *big.Int
	OriginAddress common.Address
	TokenAddress  common.Address
	Price         *big.Int
	MonethaFee    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOrderPaidInToken is a free log retrieval operation binding the contract event 0xd6dc98331ad06baebe39c90f4fd554341ad121d55e4384bd046def391501a00f.
//
// Solidity: e OrderPaidInToken(_orderId indexed uint256, _originAddress indexed address, _tokenAddress indexed address, _price uint256, _monethaFee uint256)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractFilterer) FilterOrderPaidInToken(opts *bind.FilterOpts, _orderId []*big.Int, _originAddress []common.Address, _tokenAddress []common.Address) (*PrivatePaymentProcessorContractOrderPaidInTokenIterator, error) {

	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}
	var _originAddressRule []interface{}
	for _, _originAddressItem := range _originAddress {
		_originAddressRule = append(_originAddressRule, _originAddressItem)
	}
	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}

	logs, sub, err := _PrivatePaymentProcessorContract.contract.FilterLogs(opts, "OrderPaidInToken", _orderIdRule, _originAddressRule, _tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &PrivatePaymentProcessorContractOrderPaidInTokenIterator{contract: _PrivatePaymentProcessorContract.contract, event: "OrderPaidInToken", logs: logs, sub: sub}, nil
}

// WatchOrderPaidInToken is a free log subscription operation binding the contract event 0xd6dc98331ad06baebe39c90f4fd554341ad121d55e4384bd046def391501a00f.
//
// Solidity: e OrderPaidInToken(_orderId indexed uint256, _originAddress indexed address, _tokenAddress indexed address, _price uint256, _monethaFee uint256)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractFilterer) WatchOrderPaidInToken(opts *bind.WatchOpts, sink chan<- *PrivatePaymentProcessorContractOrderPaidInToken, _orderId []*big.Int, _originAddress []common.Address, _tokenAddress []common.Address) (event.Subscription, error) {

	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}
	var _originAddressRule []interface{}
	for _, _originAddressItem := range _originAddress {
		_originAddressRule = append(_originAddressRule, _originAddressItem)
	}
	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}

	logs, sub, err := _PrivatePaymentProcessorContract.contract.WatchLogs(opts, "OrderPaidInToken", _orderIdRule, _originAddressRule, _tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivatePaymentProcessorContractOrderPaidInToken)
				if err := _PrivatePaymentProcessorContract.contract.UnpackLog(event, "OrderPaidInToken", log); err != nil {
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

// PrivatePaymentProcessorContractOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the PrivatePaymentProcessorContract contract.
type PrivatePaymentProcessorContractOwnershipRenouncedIterator struct {
	Event *PrivatePaymentProcessorContractOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *PrivatePaymentProcessorContractOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivatePaymentProcessorContractOwnershipRenounced)
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
		it.Event = new(PrivatePaymentProcessorContractOwnershipRenounced)
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
func (it *PrivatePaymentProcessorContractOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivatePaymentProcessorContractOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivatePaymentProcessorContractOwnershipRenounced represents a OwnershipRenounced event raised by the PrivatePaymentProcessorContract contract.
type PrivatePaymentProcessorContractOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*PrivatePaymentProcessorContractOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PrivatePaymentProcessorContract.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PrivatePaymentProcessorContractOwnershipRenouncedIterator{contract: _PrivatePaymentProcessorContract.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *PrivatePaymentProcessorContractOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PrivatePaymentProcessorContract.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivatePaymentProcessorContractOwnershipRenounced)
				if err := _PrivatePaymentProcessorContract.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// PrivatePaymentProcessorContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PrivatePaymentProcessorContract contract.
type PrivatePaymentProcessorContractOwnershipTransferredIterator struct {
	Event *PrivatePaymentProcessorContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PrivatePaymentProcessorContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivatePaymentProcessorContractOwnershipTransferred)
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
		it.Event = new(PrivatePaymentProcessorContractOwnershipTransferred)
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
func (it *PrivatePaymentProcessorContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivatePaymentProcessorContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivatePaymentProcessorContractOwnershipTransferred represents a OwnershipTransferred event raised by the PrivatePaymentProcessorContract contract.
type PrivatePaymentProcessorContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PrivatePaymentProcessorContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PrivatePaymentProcessorContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PrivatePaymentProcessorContractOwnershipTransferredIterator{contract: _PrivatePaymentProcessorContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PrivatePaymentProcessorContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PrivatePaymentProcessorContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivatePaymentProcessorContractOwnershipTransferred)
				if err := _PrivatePaymentProcessorContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PrivatePaymentProcessorContractPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the PrivatePaymentProcessorContract contract.
type PrivatePaymentProcessorContractPauseIterator struct {
	Event *PrivatePaymentProcessorContractPause // Event containing the contract specifics and raw log

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
func (it *PrivatePaymentProcessorContractPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivatePaymentProcessorContractPause)
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
		it.Event = new(PrivatePaymentProcessorContractPause)
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
func (it *PrivatePaymentProcessorContractPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivatePaymentProcessorContractPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivatePaymentProcessorContractPause represents a Pause event raised by the PrivatePaymentProcessorContract contract.
type PrivatePaymentProcessorContractPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractFilterer) FilterPause(opts *bind.FilterOpts) (*PrivatePaymentProcessorContractPauseIterator, error) {

	logs, sub, err := _PrivatePaymentProcessorContract.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PrivatePaymentProcessorContractPauseIterator{contract: _PrivatePaymentProcessorContract.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PrivatePaymentProcessorContractPause) (event.Subscription, error) {

	logs, sub, err := _PrivatePaymentProcessorContract.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivatePaymentProcessorContractPause)
				if err := _PrivatePaymentProcessorContract.contract.UnpackLog(event, "Pause", log); err != nil {
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

// PrivatePaymentProcessorContractPaymentRefundingIterator is returned from FilterPaymentRefunding and is used to iterate over the raw logs and unpacked data for PaymentRefunding events raised by the PrivatePaymentProcessorContract contract.
type PrivatePaymentProcessorContractPaymentRefundingIterator struct {
	Event *PrivatePaymentProcessorContractPaymentRefunding // Event containing the contract specifics and raw log

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
func (it *PrivatePaymentProcessorContractPaymentRefundingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivatePaymentProcessorContractPaymentRefunding)
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
		it.Event = new(PrivatePaymentProcessorContractPaymentRefunding)
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
func (it *PrivatePaymentProcessorContractPaymentRefundingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivatePaymentProcessorContractPaymentRefundingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivatePaymentProcessorContractPaymentRefunding represents a PaymentRefunding event raised by the PrivatePaymentProcessorContract contract.
type PrivatePaymentProcessorContractPaymentRefunding struct {
	OrderId       *big.Int
	ClientAddress common.Address
	Amount        *big.Int
	RefundReason  string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPaymentRefunding is a free log retrieval operation binding the contract event 0x26e77179a69c2db5e1f39af4e228bc8c2205384ba14b8c1e3339049db4ee42c5.
//
// Solidity: e PaymentRefunding(_orderId indexed uint256, _clientAddress indexed address, _amount uint256, _refundReason string)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractFilterer) FilterPaymentRefunding(opts *bind.FilterOpts, _orderId []*big.Int, _clientAddress []common.Address) (*PrivatePaymentProcessorContractPaymentRefundingIterator, error) {

	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}
	var _clientAddressRule []interface{}
	for _, _clientAddressItem := range _clientAddress {
		_clientAddressRule = append(_clientAddressRule, _clientAddressItem)
	}

	logs, sub, err := _PrivatePaymentProcessorContract.contract.FilterLogs(opts, "PaymentRefunding", _orderIdRule, _clientAddressRule)
	if err != nil {
		return nil, err
	}
	return &PrivatePaymentProcessorContractPaymentRefundingIterator{contract: _PrivatePaymentProcessorContract.contract, event: "PaymentRefunding", logs: logs, sub: sub}, nil
}

// WatchPaymentRefunding is a free log subscription operation binding the contract event 0x26e77179a69c2db5e1f39af4e228bc8c2205384ba14b8c1e3339049db4ee42c5.
//
// Solidity: e PaymentRefunding(_orderId indexed uint256, _clientAddress indexed address, _amount uint256, _refundReason string)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractFilterer) WatchPaymentRefunding(opts *bind.WatchOpts, sink chan<- *PrivatePaymentProcessorContractPaymentRefunding, _orderId []*big.Int, _clientAddress []common.Address) (event.Subscription, error) {

	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}
	var _clientAddressRule []interface{}
	for _, _clientAddressItem := range _clientAddress {
		_clientAddressRule = append(_clientAddressRule, _clientAddressItem)
	}

	logs, sub, err := _PrivatePaymentProcessorContract.contract.WatchLogs(opts, "PaymentRefunding", _orderIdRule, _clientAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivatePaymentProcessorContractPaymentRefunding)
				if err := _PrivatePaymentProcessorContract.contract.UnpackLog(event, "PaymentRefunding", log); err != nil {
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

// PrivatePaymentProcessorContractPaymentWithdrawnIterator is returned from FilterPaymentWithdrawn and is used to iterate over the raw logs and unpacked data for PaymentWithdrawn events raised by the PrivatePaymentProcessorContract contract.
type PrivatePaymentProcessorContractPaymentWithdrawnIterator struct {
	Event *PrivatePaymentProcessorContractPaymentWithdrawn // Event containing the contract specifics and raw log

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
func (it *PrivatePaymentProcessorContractPaymentWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivatePaymentProcessorContractPaymentWithdrawn)
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
		it.Event = new(PrivatePaymentProcessorContractPaymentWithdrawn)
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
func (it *PrivatePaymentProcessorContractPaymentWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivatePaymentProcessorContractPaymentWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivatePaymentProcessorContractPaymentWithdrawn represents a PaymentWithdrawn event raised by the PrivatePaymentProcessorContract contract.
type PrivatePaymentProcessorContractPaymentWithdrawn struct {
	OrderId       *big.Int
	ClientAddress common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPaymentWithdrawn is a free log retrieval operation binding the contract event 0xbe85bf3b0a1e335a22c461f84cf759dfe589ec1539caf4dce60f999d72dd8e23.
//
// Solidity: e PaymentWithdrawn(_orderId indexed uint256, _clientAddress indexed address, amount uint256)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractFilterer) FilterPaymentWithdrawn(opts *bind.FilterOpts, _orderId []*big.Int, _clientAddress []common.Address) (*PrivatePaymentProcessorContractPaymentWithdrawnIterator, error) {

	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}
	var _clientAddressRule []interface{}
	for _, _clientAddressItem := range _clientAddress {
		_clientAddressRule = append(_clientAddressRule, _clientAddressItem)
	}

	logs, sub, err := _PrivatePaymentProcessorContract.contract.FilterLogs(opts, "PaymentWithdrawn", _orderIdRule, _clientAddressRule)
	if err != nil {
		return nil, err
	}
	return &PrivatePaymentProcessorContractPaymentWithdrawnIterator{contract: _PrivatePaymentProcessorContract.contract, event: "PaymentWithdrawn", logs: logs, sub: sub}, nil
}

// WatchPaymentWithdrawn is a free log subscription operation binding the contract event 0xbe85bf3b0a1e335a22c461f84cf759dfe589ec1539caf4dce60f999d72dd8e23.
//
// Solidity: e PaymentWithdrawn(_orderId indexed uint256, _clientAddress indexed address, amount uint256)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractFilterer) WatchPaymentWithdrawn(opts *bind.WatchOpts, sink chan<- *PrivatePaymentProcessorContractPaymentWithdrawn, _orderId []*big.Int, _clientAddress []common.Address) (event.Subscription, error) {

	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}
	var _clientAddressRule []interface{}
	for _, _clientAddressItem := range _clientAddress {
		_clientAddressRule = append(_clientAddressRule, _clientAddressItem)
	}

	logs, sub, err := _PrivatePaymentProcessorContract.contract.WatchLogs(opts, "PaymentWithdrawn", _orderIdRule, _clientAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivatePaymentProcessorContractPaymentWithdrawn)
				if err := _PrivatePaymentProcessorContract.contract.UnpackLog(event, "PaymentWithdrawn", log); err != nil {
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

// PrivatePaymentProcessorContractPaymentsProcessedIterator is returned from FilterPaymentsProcessed and is used to iterate over the raw logs and unpacked data for PaymentsProcessed events raised by the PrivatePaymentProcessorContract contract.
type PrivatePaymentProcessorContractPaymentsProcessedIterator struct {
	Event *PrivatePaymentProcessorContractPaymentsProcessed // Event containing the contract specifics and raw log

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
func (it *PrivatePaymentProcessorContractPaymentsProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivatePaymentProcessorContractPaymentsProcessed)
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
		it.Event = new(PrivatePaymentProcessorContractPaymentsProcessed)
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
func (it *PrivatePaymentProcessorContractPaymentsProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivatePaymentProcessorContractPaymentsProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivatePaymentProcessorContractPaymentsProcessed represents a PaymentsProcessed event raised by the PrivatePaymentProcessorContract contract.
type PrivatePaymentProcessorContractPaymentsProcessed struct {
	MerchantAddress common.Address
	Amount          *big.Int
	Fee             *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaymentsProcessed is a free log retrieval operation binding the contract event 0x81a22b411fd70ce0518d51057fbb3af960d3f09f0d82d1e684d4484444674367.
//
// Solidity: e PaymentsProcessed(_merchantAddress indexed address, _amount uint256, _fee uint256)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractFilterer) FilterPaymentsProcessed(opts *bind.FilterOpts, _merchantAddress []common.Address) (*PrivatePaymentProcessorContractPaymentsProcessedIterator, error) {

	var _merchantAddressRule []interface{}
	for _, _merchantAddressItem := range _merchantAddress {
		_merchantAddressRule = append(_merchantAddressRule, _merchantAddressItem)
	}

	logs, sub, err := _PrivatePaymentProcessorContract.contract.FilterLogs(opts, "PaymentsProcessed", _merchantAddressRule)
	if err != nil {
		return nil, err
	}
	return &PrivatePaymentProcessorContractPaymentsProcessedIterator{contract: _PrivatePaymentProcessorContract.contract, event: "PaymentsProcessed", logs: logs, sub: sub}, nil
}

// WatchPaymentsProcessed is a free log subscription operation binding the contract event 0x81a22b411fd70ce0518d51057fbb3af960d3f09f0d82d1e684d4484444674367.
//
// Solidity: e PaymentsProcessed(_merchantAddress indexed address, _amount uint256, _fee uint256)
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractFilterer) WatchPaymentsProcessed(opts *bind.WatchOpts, sink chan<- *PrivatePaymentProcessorContractPaymentsProcessed, _merchantAddress []common.Address) (event.Subscription, error) {

	var _merchantAddressRule []interface{}
	for _, _merchantAddressItem := range _merchantAddress {
		_merchantAddressRule = append(_merchantAddressRule, _merchantAddressItem)
	}

	logs, sub, err := _PrivatePaymentProcessorContract.contract.WatchLogs(opts, "PaymentsProcessed", _merchantAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivatePaymentProcessorContractPaymentsProcessed)
				if err := _PrivatePaymentProcessorContract.contract.UnpackLog(event, "PaymentsProcessed", log); err != nil {
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

// PrivatePaymentProcessorContractUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the PrivatePaymentProcessorContract contract.
type PrivatePaymentProcessorContractUnpauseIterator struct {
	Event *PrivatePaymentProcessorContractUnpause // Event containing the contract specifics and raw log

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
func (it *PrivatePaymentProcessorContractUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivatePaymentProcessorContractUnpause)
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
		it.Event = new(PrivatePaymentProcessorContractUnpause)
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
func (it *PrivatePaymentProcessorContractUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivatePaymentProcessorContractUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivatePaymentProcessorContractUnpause represents a Unpause event raised by the PrivatePaymentProcessorContract contract.
type PrivatePaymentProcessorContractUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractFilterer) FilterUnpause(opts *bind.FilterOpts) (*PrivatePaymentProcessorContractUnpauseIterator, error) {

	logs, sub, err := _PrivatePaymentProcessorContract.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PrivatePaymentProcessorContractUnpauseIterator{contract: _PrivatePaymentProcessorContract.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_PrivatePaymentProcessorContract *PrivatePaymentProcessorContractFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PrivatePaymentProcessorContractUnpause) (event.Subscription, error) {

	logs, sub, err := _PrivatePaymentProcessorContract.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivatePaymentProcessorContractUnpause)
				if err := _PrivatePaymentProcessorContract.contract.UnpackLog(event, "Unpause", log); err != nil {
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

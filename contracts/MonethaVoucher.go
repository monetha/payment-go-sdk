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

// MonethaVoucherContractABI is the input ABI used to generate the binding from.
const MonethaVoucherContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"reclaimEtherTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isMonethaAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mthEthRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"purchased\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"standard\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"totalDistributedIn\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mthToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reclaimEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"reclaimTokenTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RATE_COEFFICIENT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voucherMthRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_isMonethaAddress\",\"type\":\"bool\"}],\"name\":\"setMonethaAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint16\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"distributed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalPurchased\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_voucherMthRate\",\"type\":\"uint256\"},{\"name\":\"_mthEthRate\",\"type\":\"uint256\"},{\"name\":\"_mthToken\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"releasedVouchers\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amountWeiTransferred\",\"type\":\"uint256\"}],\"name\":\"DiscountApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"addedVouchers\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amountWeiEquivalent\",\"type\":\"uint256\"}],\"name\":\"PaybackApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"vouchersBought\",\"type\":\"uint256\"}],\"name\":\"VouchersBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"vouchersSold\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amountWeiTransferred\",\"type\":\"uint256\"}],\"name\":\"VouchersSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldVoucherMthRate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newVoucherMthRate\",\"type\":\"uint256\"}],\"name\":\"VoucherMthRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldMthEthRate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newMthEthRate\",\"type\":\"uint256\"}],\"name\":\"MthEthRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"vouchersAdded\",\"type\":\"uint256\"}],\"name\":\"VouchersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"releasedVoucher\",\"type\":\"uint256\"}],\"name\":\"VoucherReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"vouchers\",\"type\":\"uint256\"}],\"name\":\"PurchasedVouchersReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReclaimTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReclaimEther\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_isMonethaAddress\",\"type\":\"bool\"}],\"name\":\"MonethaAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalInSharedPool\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalDistributed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"toWei\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"fromWei\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_for\",\"type\":\"address\"},{\"name\":\"_vouchers\",\"type\":\"uint256\"}],\"name\":\"applyDiscount\",\"outputs\":[{\"name\":\"amountVouchers\",\"type\":\"uint256\"},{\"name\":\"amountWei\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_for\",\"type\":\"address\"},{\"name\":\"_amountWei\",\"type\":\"uint256\"}],\"name\":\"applyPayback\",\"outputs\":[{\"name\":\"amountVouchers\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vouchers\",\"type\":\"uint256\"}],\"name\":\"buyVouchers\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vouchers\",\"type\":\"uint256\"}],\"name\":\"sellVouchers\",\"outputs\":[{\"name\":\"weis\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"releasePurchasedTo\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"purchasedBy\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voucherMthRate\",\"type\":\"uint256\"}],\"name\":\"updateVoucherMthRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_mthEthRate\",\"type\":\"uint256\"}],\"name\":\"updateMthEthRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MonethaVoucherContractBin is the compiled bytecode used for deploying new contracts.
const MonethaVoucherContractBin = `0x60806040526002805460ff191690553480156200001b57600080fd5b506040516060806200219583398101604090815281516020830151919092015160008054600160a060020a031916331781558311620000e157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f766f75636865724d7468526174652073686f756c64206265206772656174657260448201527f207468616e203000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600082116200017757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f6d7468457468526174652073686f756c6420626520677265617465722074686160448201527f6e20300000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a0381161515620001ef57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f6d7573742062652076616c696420636f6e747261637400000000000000000000604482015290519081900360640190fd5b6003839055600482905560068054600160a060020a031916600160a060020a0383161790556200022764010000000062000230810204565b50505062000287565b6004546003546200024f9164010000000062001d246200025482021704565b600555565b6000821515620002675750600062000281565b508181028183828115156200027857fe5b04146200028157fe5b92915050565b611efe80620002976000396000f3006080604052600436106102035763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde038114610223578063095ea7b3146102ad5780630ce11ded146102e557806317ffc320146102fd57806318160ddd1461031e57806323b872dd146103455780632e84a3741461036f578063313ce5671461039357806331d41325146103be5780633f4ba83a146103df57806349ff7a62146103f4578063522fe98e146104095780635a3b7e421461042a5780635c975abb1461043f5780636481f08114610454578063669dafe81461046c5780636a2b171a146104845780636daa212f146104c157806370a08231146104e5578063715018a614610506578063732a935c1461051b5780638456cb59146105375780638d8c36e91461054c5780638da5cb5b1461056d57806393306c431461059e57806395d89b41146105b35780639f727c27146105c85780639fe98bda146105dd578063a05e0f6514610607578063a9059cbb146102ad578063b04087211461061c578063b2bea9c114610627578063b6f817241461064b578063b9c5eb9014610660578063c07e339114610678578063c3fdbef81461069e578063d9582cde146106b3578063dd62ed3e146106db578063e632c2f314610702578063e65d152214610717578063efca2eed1461072f578063f2fde38b14610744575b3360009081526001602052604090205460ff16151561022157600080fd5b005b34801561022f57600080fd5b50610238610765565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561027257818101518382015260200161025a565b50505050905090810190601f16801561029f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102b957600080fd5b506102d1600160a060020a036004351660243561079c565b604080519115158252519081900360200190f35b3480156102f157600080fd5b506102216004356107a2565b34801561030957600080fd5b50610221600160a060020a03600435166108bf565b34801561032a57600080fd5b506103336109cd565b60408051918252519081900360200190f35b34801561035157600080fd5b506102d1600160a060020a036004358116906024351660443561079c565b34801561037b57600080fd5b50610221600160a060020a03600435166024356109dc565b34801561039f57600080fd5b506103a8610abb565b6040805160ff9092168252519081900360200190f35b3480156103ca57600080fd5b506102d1600160a060020a0360043516610ac0565b3480156103eb57600080fd5b50610221610ad5565b34801561040057600080fd5b50610333610b32565b34801561041557600080fd5b50610333600160a060020a0360043516610b38565b34801561043657600080fd5b50610238610b4a565b34801561044b57600080fd5b506102d1610b81565b34801561046057600080fd5b50610333600435610b8a565b34801561047857600080fd5b50610333600435610c92565b34801561049057600080fd5b506104a8600160a060020a0360043516602435610ca3565b6040805192835260208301919091528051918290030190f35b3480156104cd57600080fd5b50610333600160a060020a0360043516602435610e22565b3480156104f157600080fd5b50610333600160a060020a0360043516610ef4565b34801561051257600080fd5b50610221610f2e565b34801561052757600080fd5b5061033361ffff60043516610f9a565b34801561054357600080fd5b50610221610fac565b34801561055857600080fd5b50610333600160a060020a036004351661100b565b34801561057957600080fd5b50610582611026565b60408051600160a060020a039092168252519081900360200190f35b3480156105aa57600080fd5b50610582611035565b3480156105bf57600080fd5b50610238611044565b3480156105d457600080fd5b5061022161107b565b3480156105e957600080fd5b50610221600160a060020a0360043581169060243516604435611116565b34801561061357600080fd5b506103336111d9565b6102216004356111e5565b34801561063357600080fd5b506102d1600160a060020a0360043516602435611310565b34801561065757600080fd5b50610333611453565b34801561066c57600080fd5b50610221600435611459565b34801561068457600080fd5b50610221600160a060020a03600435166024351515611576565b3480156106aa57600080fd5b506103336115f1565b3480156106bf57600080fd5b5061033361ffff60043516600160a060020a0360243516611603565b3480156106e757600080fd5b50610333600160a060020a0360043581169060243516611620565b34801561070e57600080fd5b50610333611628565b34801561072357600080fd5b5061033360043561162e565b34801561073b57600080fd5b50610333611639565b34801561075057600080fd5b50610221600160a060020a036004351661164b565b60408051808201909152600f81527f4d6f6e6574686120566f75636865720000000000000000000000000000000000602082015281565b60008080fd5b3360009081526001602052604090205460ff1615156107c057600080fd5b60008111610818576040805160e560020a62461bcd02815260206004820152601860248201527f73686f756c642062652067726561746572207468616e20300000000000000000604482015290519081900360640190fd5b600454811415610872576040805160e560020a62461bcd02815260206004820152601660248201527f73616d652061732070726576696f75732076616c756500000000000000000000604482015290519081900360640190fd5b600481905561087f61166e565b600454604080519182526020820183905280517ff376950c4c79df1231c4d2cc604341aff0685bc9b6c3a1661200a6d4f319d4e49281900390910190a150565b60008054600160a060020a031633146108d757600080fd5b604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051600160a060020a038416916370a082319160248083019260209291908290030181600087803b15801561093857600080fd5b505af115801561094c573d6000803e3d6000fd5b505050506040513d602081101561096257600080fd5b505160005490915061098790600160a060020a0384811691168363ffffffff61168816565b600054604080518381529051600160a060020a03909216917f355069f20974db323c9dcd100e8bf13fb2acc1884e5ec05b0a89c09e15ce810f9181900360200190a25050565b60006109d7611740565b905090565b600054600160a060020a031633146109f357600080fd5b600160a060020a0382161515610a41576040805160e560020a62461bcd02815260206004820152601b6024820152600080516020611eb3833981519152604482015290519081900360640190fd5b604051600160a060020a0383169082156108fc029083906000818181858888f19350505050158015610a77573d6000803e3d6000fd5b50604080518281529051600160a060020a038416917fb54913b2b58b2e96ea9b4e96ba2353cf13426af9d3f252e0c17899a93c4ce125919081900360200190a25050565b600581565b60016020526000908152604090205460ff1681565b600054600160a060020a03163314610aec57600080fd5b60025460ff161515610afd57600080fd5b6002805460ff191690556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3390600090a1565b60045481565b60076020526000908152604090205481565b60408051808201909152600581527f4552433230000000000000000000000000000000000000000000000000000000602082015281565b60025460ff1681565b3360009081526001602052604081205460ff161515610ba857600080fd5b33600090815260076020526040902054821115610c0f576040805160e560020a62461bcd02815260206004820152601560248201527f496e73756666696369656e7420766f7563686572730000000000000000000000604482015290519081900360640190fd5b610c1933836117dd565b610c2282611839565b604051909150339082156108fc029083906000818181858888f19350505050158015610c52573d6000803e3d6000fd5b506040805183815260208101839052815133927f4f4ec1e156ff2a5d94d15690727a0c7c94705b9dd53e5e0abf3fa16195b106e2928290030190a2919050565b6000610c9d82611839565b92915050565b3360009081526001602052604081205481908190819060ff161515610cc757600080fd5b600160a060020a0386161515610d15576040805160e560020a62461bcd02815260206004820152601b6024820152600080516020611eb3833981519152604482015290519081900360640190fd5b610d1f868661186e565b9150811515610d345760009350839250610e19565b610d3d82611839565b90503031811115610d98576040805160e560020a62461bcd02815260206004820152601260248201527f696e73756666696369656e742066756e64730000000000000000000000000000604482015290519081900360640190fd5b604051600160a060020a0387169082156108fc029083906000818181858888f19350505050158015610dce573d6000803e3d6000fd5b5060408051838152602081018390528151600160a060020a038916927fa5d355764530fdfd305e1709a5c34997434bd0032f3afe650f362d1d1544b947928290030190a28181935093505b50509250929050565b3360009081526001602052604081205460ff161515610e4057600080fd5b610e4982611977565b9050610e55838261199c565b1515610eab576040805160e560020a62461bcd02815260206004820152601660248201527f766f756368657273206d75737420626520616464656400000000000000000000604482015290519081900360640190fd5b60408051828152602081018490528151600160a060020a038616927f289266f290897c2e82d0cee6b6f74f7d5e62d30244d4bfcf390abc4af4cd5666928290030190a292915050565b600160a060020a038116600090815260076020526040812054610c9d90610f2284610f1d611b56565b611b5f565b9063ffffffff611bd716565b600054600160a060020a03163314610f4557600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b60096020526000908152604090205481565b600054600160a060020a03163314610fc357600080fd5b60025460ff1615610fd357600080fd5b6002805460ff191660011790556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62590600090a1565b600160a060020a031660009081526007602052604090205490565b600054600160a060020a031681565b600654600160a060020a031681565b60408051808201909152600481527f4d54485600000000000000000000000000000000000000000000000000000000602082015281565b60008054600160a060020a0316331461109357600080fd5b5060008054604051303192600160a060020a03909216916108fc841502918491818181858888f193505050501580156110d0573d6000803e3d6000fd5b50600054604080518381529051600160a060020a03909216917fb54913b2b58b2e96ea9b4e96ba2353cf13426af9d3f252e0c17899a93c4ce1259181900360200190a250565b600054600160a060020a0316331461112d57600080fd5b600160a060020a038216151561117b576040805160e560020a62461bcd02815260206004820152601b6024820152600080516020611eb3833981519152604482015290519081900360640190fd5b611195600160a060020a038416838363ffffffff61168816565b604080518281529051600160a060020a038416917f355069f20974db323c9dcd100e8bf13fb2acc1884e5ec05b0a89c09e15ce810f919081900360200190a2505050565b670de0b6b3a764000081565b3360009081526001602052604081205460ff16151561120357600080fd5b61120b611b56565b90508161121782611be4565b101561126d576040805160e560020a62461bcd02815260206004820152601d60248201527f696e73756666696369656e7420766f7563686572732070726573656e74000000604482015290519081900360640190fd5b61127682611839565b34146112cc576040805160e560020a62461bcd02815260206004820152601260248201527f696e73756666696369656e742066756e64730000000000000000000000000000604482015290519081900360640190fd5b6112d63383611c0c565b60408051838152905133917fad6ca73eec0e00d9089e39fad2962bb18992dd5feeae519dcb4cef2aabd93e30919081900360200190a25050565b3360009081526001602052604081205460ff16151561132e57600080fd5b33600090815260076020526040902054821115611395576040805160e560020a62461bcd02815260206004820152601560248201527f496e73756666696369656e7420566f7563686572730000000000000000000000604482015290519081900360640190fd5b600160a060020a03831615156113f5576040805160e560020a62461bcd02815260206004820152601760248201527f616464726573732073686f756c642062652076616c6964000000000000000000604482015290519081900360640190fd5b6113ff33836117dd565b611409838361199c565b50604080518381529051600160a060020a0385169133917fd49597d4844c4d37387a6053e4a1832bd054dee0aaf72fcd84da7e2aa9f8cbf09181900360200190a350600192915050565b60035481565b3360009081526001602052604090205460ff16151561147757600080fd5b600081116114cf576040805160e560020a62461bcd02815260206004820152601860248201527f73686f756c642062652067726561746572207468616e20300000000000000000604482015290519081900360640190fd5b600354811415611529576040805160e560020a62461bcd02815260206004820152601660248201527f73616d652061732070726576696f75732076616c756500000000000000000000604482015290519081900360640190fd5b600381905561153661166e565b600354604080519182526020820183905280517f5f8dd084970454563d6e013038a3a2204c50015701c39e62c3429738026fa7849281900390910190a150565b600054600160a060020a0316331461158d57600080fd5b600160a060020a038216600081815260016020908152604091829020805460ff191685151590811790915582519384529083015280517fa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e6929281900390910190a15050565b60006109d76115fe611b56565b611be4565b600a60209081526000928352604080842090915290825290205481565b600092915050565b60085481565b6000610c9d82611977565b60006109d7611646611b56565b611c61565b600054600160a060020a0316331461166257600080fd5b61166b81611ca7565b50565b6004546003546116839163ffffffff611d2416565b600555565b82600160a060020a031663a9059cbb83836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561170457600080fd5b505af1158015611718573d6000803e3d6000fd5b505050506040513d602081101561172e57600080fd5b5051151561173b57600080fd5b505050565b600654604080517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290516000926109d792600160a060020a03909116916370a082319160248082019260209290919082900301818887803b1580156117ac57600080fd5b505af11580156117c0573d6000803e3d6000fd5b505050506040513d60208110156117d657600080fd5b5051611d4d565b600160a060020a038216600090815260076020526040902054611806908263ffffffff611d7016565b600160a060020a038316600090815260076020526040902055600854611832908263ffffffff611d7016565b6008555050565b600554600090610c9d90611862846ec097ce7bc90715b34b9f100000000063ffffffff611d2416565b9063ffffffff611d8216565b60008080600160a060020a03851615156118d2576040805160e560020a62461bcd02815260206004820152601560248201527f6d7573742062652076616c696420616464726573730000000000000000000000604482015290519081900360640190fd5b6118da611b56565b91506000905060008261ffff1611156119205761190b6118fe868660018603611d97565b829063ffffffff611bd716565b905061191d848263ffffffff611d7016565b93505b61192e6118fe868685611d97565b604080518281529051919250600160a060020a038716917f5c2b936c2e3089a9ba68c1146fc951f328fd49b4b1a5c00156babd126b227a459181900360200190a2949350505050565b6000610c9d670de0b6b3a7640000800261186260055485611d2490919063ffffffff16565b6000808080600160a060020a03861615156119ef576040805160e560020a62461bcd02815260206004820152601b6024820152600080516020611eb3833981519152604482015290519081900360640190fd5b6119f7611b56565b925084611a0384611be4565b1015611a7f576040805160e560020a62461bcd02815260206004820152603a60248201527f6d757374206265206c657373206f7220657175616c207468616e20766f75636860448201527f6572732070726573656e7420696e2073686172656420706f6f6c000000000000606482015290519081900360840190fd5b61ffff83166000908152600960205260409020549150611aa5828663ffffffff611bd716565b61ffff8416600090815260096020908152604080832093909355600a8152828220600160a060020a038a168352905220549050611ae8818663ffffffff611bd716565b61ffff84166000908152600a60209081526040808320600160a060020a038b1680855290835292819020939093558251888152925191927fa4822228ba5999ec50d98259975e2e7e3b56a666bfd3cae8c7722ee90af63a28929081900390910190a250600195945050505050565b62f0c3f0420490565b61ffff81166000818152600a60209081526040808320600160a060020a03871684529091528120549091821015611bd05761ffff6000198401166000908152600a60209081526040808320600160a060020a0388168452909152902054611bcd90829063ffffffff611bd716565b90505b9392505050565b81810182811015610c9d57fe5b6000610c9d600854611c00611bf885611c61565b611c00611740565b9063ffffffff611d7016565b600160a060020a038216600090815260076020526040902054611c35908263ffffffff611bd716565b600160a060020a038316600090815260076020526040902055600854611832908263ffffffff611bd716565b61ffff81166000818152600960205260408120549091821015610c9d5761ffff600019840116600090815260096020526040902054611bd090829063ffffffff611bd716565b600160a060020a0381161515611cbc57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000821515611d3557506000610c9d565b50818102818382811515611d4557fe5b0414610c9d57fe5b6000610c9d670de0b6b3a764000061186260035485611d2490919063ffffffff16565b600082821115611d7c57fe5b50900390565b60008183811515611d8f57fe5b049392505050565b6000808080851515611dac5760009350611ea8565b61ffff85166000908152600a60209081526040808320600160a060020a038b1684529091529020549250859150818311611e105761ffff85166000908152600a60209081526040808320600160a060020a038b168452909152812055829150611e49565b611e20838763ffffffff611d7016565b61ffff86166000908152600a60209081526040808320600160a060020a038c1684529091529020555b5061ffff841660009081526009602052604090205481811415611e7f5761ffff8516600090815260096020526040812055611ea4565b611e8f818363ffffffff611d7016565b61ffff86166000908152600960205260409020555b8193505b505050939250505056007a65726f2061646472657373206973206e6f7420616c6c6f7765640000000000a165627a7a723058206316be0f4a6182a4b2630543472bd763c5a9e29ab520cef2f96c5ffd0e3ef5e10029`

// DeployMonethaVoucherContract deploys a new Ethereum contract, binding an instance of MonethaVoucherContract to it.
func DeployMonethaVoucherContract(auth *bind.TransactOpts, backend bind.ContractBackend, _voucherMthRate *big.Int, _mthEthRate *big.Int, _mthToken common.Address) (common.Address, *types.Transaction, *MonethaVoucherContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MonethaVoucherContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MonethaVoucherContractBin), backend, _voucherMthRate, _mthEthRate, _mthToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MonethaVoucherContract{MonethaVoucherContractCaller: MonethaVoucherContractCaller{contract: contract}, MonethaVoucherContractTransactor: MonethaVoucherContractTransactor{contract: contract}, MonethaVoucherContractFilterer: MonethaVoucherContractFilterer{contract: contract}}, nil
}

// MonethaVoucherContract is an auto generated Go binding around an Ethereum contract.
type MonethaVoucherContract struct {
	MonethaVoucherContractCaller     // Read-only binding to the contract
	MonethaVoucherContractTransactor // Write-only binding to the contract
	MonethaVoucherContractFilterer   // Log filterer for contract events
}

// MonethaVoucherContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MonethaVoucherContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonethaVoucherContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MonethaVoucherContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonethaVoucherContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MonethaVoucherContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonethaVoucherContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MonethaVoucherContractSession struct {
	Contract     *MonethaVoucherContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MonethaVoucherContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MonethaVoucherContractCallerSession struct {
	Contract *MonethaVoucherContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// MonethaVoucherContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MonethaVoucherContractTransactorSession struct {
	Contract     *MonethaVoucherContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// MonethaVoucherContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MonethaVoucherContractRaw struct {
	Contract *MonethaVoucherContract // Generic contract binding to access the raw methods on
}

// MonethaVoucherContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MonethaVoucherContractCallerRaw struct {
	Contract *MonethaVoucherContractCaller // Generic read-only contract binding to access the raw methods on
}

// MonethaVoucherContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MonethaVoucherContractTransactorRaw struct {
	Contract *MonethaVoucherContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMonethaVoucherContract creates a new instance of MonethaVoucherContract, bound to a specific deployed contract.
func NewMonethaVoucherContract(address common.Address, backend bind.ContractBackend) (*MonethaVoucherContract, error) {
	contract, err := bindMonethaVoucherContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContract{MonethaVoucherContractCaller: MonethaVoucherContractCaller{contract: contract}, MonethaVoucherContractTransactor: MonethaVoucherContractTransactor{contract: contract}, MonethaVoucherContractFilterer: MonethaVoucherContractFilterer{contract: contract}}, nil
}

// NewMonethaVoucherContractCaller creates a new read-only instance of MonethaVoucherContract, bound to a specific deployed contract.
func NewMonethaVoucherContractCaller(address common.Address, caller bind.ContractCaller) (*MonethaVoucherContractCaller, error) {
	contract, err := bindMonethaVoucherContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContractCaller{contract: contract}, nil
}

// NewMonethaVoucherContractTransactor creates a new write-only instance of MonethaVoucherContract, bound to a specific deployed contract.
func NewMonethaVoucherContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MonethaVoucherContractTransactor, error) {
	contract, err := bindMonethaVoucherContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContractTransactor{contract: contract}, nil
}

// NewMonethaVoucherContractFilterer creates a new log filterer instance of MonethaVoucherContract, bound to a specific deployed contract.
func NewMonethaVoucherContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MonethaVoucherContractFilterer, error) {
	contract, err := bindMonethaVoucherContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContractFilterer{contract: contract}, nil
}

// bindMonethaVoucherContract binds a generic wrapper to an already deployed contract.
func bindMonethaVoucherContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MonethaVoucherContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonethaVoucherContract *MonethaVoucherContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MonethaVoucherContract.Contract.MonethaVoucherContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonethaVoucherContract *MonethaVoucherContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.MonethaVoucherContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonethaVoucherContract *MonethaVoucherContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.MonethaVoucherContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonethaVoucherContract *MonethaVoucherContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MonethaVoucherContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonethaVoucherContract *MonethaVoucherContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonethaVoucherContract *MonethaVoucherContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.contract.Transact(opts, method, params...)
}

// RATECOEFFICIENT is a free data retrieval call binding the contract method 0xa05e0f65.
//
// Solidity: function RATE_COEFFICIENT() constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) RATECOEFFICIENT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "RATE_COEFFICIENT")
	return *ret0, err
}

// RATECOEFFICIENT is a free data retrieval call binding the contract method 0xa05e0f65.
//
// Solidity: function RATE_COEFFICIENT() constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractSession) RATECOEFFICIENT() (*big.Int, error) {
	return _MonethaVoucherContract.Contract.RATECOEFFICIENT(&_MonethaVoucherContract.CallOpts)
}

// RATECOEFFICIENT is a free data retrieval call binding the contract method 0xa05e0f65.
//
// Solidity: function RATE_COEFFICIENT() constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) RATECOEFFICIENT() (*big.Int, error) {
	return _MonethaVoucherContract.Contract.RATECOEFFICIENT(&_MonethaVoucherContract.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MonethaVoucherContract.Contract.Allowance(&_MonethaVoucherContract.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MonethaVoucherContract.Contract.Allowance(&_MonethaVoucherContract.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MonethaVoucherContract.Contract.BalanceOf(&_MonethaVoucherContract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MonethaVoucherContract.Contract.BalanceOf(&_MonethaVoucherContract.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MonethaVoucherContract *MonethaVoucherContractSession) Decimals() (uint8, error) {
	return _MonethaVoucherContract.Contract.Decimals(&_MonethaVoucherContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) Decimals() (uint8, error) {
	return _MonethaVoucherContract.Contract.Decimals(&_MonethaVoucherContract.CallOpts)
}

// Distributed is a free data retrieval call binding the contract method 0xd9582cde.
//
// Solidity: function distributed( uint16,  address) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) Distributed(opts *bind.CallOpts, arg0 uint16, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "distributed", arg0, arg1)
	return *ret0, err
}

// Distributed is a free data retrieval call binding the contract method 0xd9582cde.
//
// Solidity: function distributed( uint16,  address) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractSession) Distributed(arg0 uint16, arg1 common.Address) (*big.Int, error) {
	return _MonethaVoucherContract.Contract.Distributed(&_MonethaVoucherContract.CallOpts, arg0, arg1)
}

// Distributed is a free data retrieval call binding the contract method 0xd9582cde.
//
// Solidity: function distributed( uint16,  address) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) Distributed(arg0 uint16, arg1 common.Address) (*big.Int, error) {
	return _MonethaVoucherContract.Contract.Distributed(&_MonethaVoucherContract.CallOpts, arg0, arg1)
}

// FromWei is a free data retrieval call binding the contract method 0xe65d1522.
//
// Solidity: function fromWei(_value uint256) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) FromWei(opts *bind.CallOpts, _value *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "fromWei", _value)
	return *ret0, err
}

// FromWei is a free data retrieval call binding the contract method 0xe65d1522.
//
// Solidity: function fromWei(_value uint256) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractSession) FromWei(_value *big.Int) (*big.Int, error) {
	return _MonethaVoucherContract.Contract.FromWei(&_MonethaVoucherContract.CallOpts, _value)
}

// FromWei is a free data retrieval call binding the contract method 0xe65d1522.
//
// Solidity: function fromWei(_value uint256) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) FromWei(_value *big.Int) (*big.Int, error) {
	return _MonethaVoucherContract.Contract.FromWei(&_MonethaVoucherContract.CallOpts, _value)
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) IsMonethaAddress(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "isMonethaAddress", arg0)
	return *ret0, err
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_MonethaVoucherContract *MonethaVoucherContractSession) IsMonethaAddress(arg0 common.Address) (bool, error) {
	return _MonethaVoucherContract.Contract.IsMonethaAddress(&_MonethaVoucherContract.CallOpts, arg0)
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) IsMonethaAddress(arg0 common.Address) (bool, error) {
	return _MonethaVoucherContract.Contract.IsMonethaAddress(&_MonethaVoucherContract.CallOpts, arg0)
}

// MthEthRate is a free data retrieval call binding the contract method 0x49ff7a62.
//
// Solidity: function mthEthRate() constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) MthEthRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "mthEthRate")
	return *ret0, err
}

// MthEthRate is a free data retrieval call binding the contract method 0x49ff7a62.
//
// Solidity: function mthEthRate() constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractSession) MthEthRate() (*big.Int, error) {
	return _MonethaVoucherContract.Contract.MthEthRate(&_MonethaVoucherContract.CallOpts)
}

// MthEthRate is a free data retrieval call binding the contract method 0x49ff7a62.
//
// Solidity: function mthEthRate() constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) MthEthRate() (*big.Int, error) {
	return _MonethaVoucherContract.Contract.MthEthRate(&_MonethaVoucherContract.CallOpts)
}

// MthToken is a free data retrieval call binding the contract method 0x93306c43.
//
// Solidity: function mthToken() constant returns(address)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) MthToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "mthToken")
	return *ret0, err
}

// MthToken is a free data retrieval call binding the contract method 0x93306c43.
//
// Solidity: function mthToken() constant returns(address)
func (_MonethaVoucherContract *MonethaVoucherContractSession) MthToken() (common.Address, error) {
	return _MonethaVoucherContract.Contract.MthToken(&_MonethaVoucherContract.CallOpts)
}

// MthToken is a free data retrieval call binding the contract method 0x93306c43.
//
// Solidity: function mthToken() constant returns(address)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) MthToken() (common.Address, error) {
	return _MonethaVoucherContract.Contract.MthToken(&_MonethaVoucherContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MonethaVoucherContract *MonethaVoucherContractSession) Name() (string, error) {
	return _MonethaVoucherContract.Contract.Name(&_MonethaVoucherContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) Name() (string, error) {
	return _MonethaVoucherContract.Contract.Name(&_MonethaVoucherContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MonethaVoucherContract *MonethaVoucherContractSession) Owner() (common.Address, error) {
	return _MonethaVoucherContract.Contract.Owner(&_MonethaVoucherContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) Owner() (common.Address, error) {
	return _MonethaVoucherContract.Contract.Owner(&_MonethaVoucherContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MonethaVoucherContract *MonethaVoucherContractSession) Paused() (bool, error) {
	return _MonethaVoucherContract.Contract.Paused(&_MonethaVoucherContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) Paused() (bool, error) {
	return _MonethaVoucherContract.Contract.Paused(&_MonethaVoucherContract.CallOpts)
}

// Purchased is a free data retrieval call binding the contract method 0x522fe98e.
//
// Solidity: function purchased( address) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) Purchased(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "purchased", arg0)
	return *ret0, err
}

// Purchased is a free data retrieval call binding the contract method 0x522fe98e.
//
// Solidity: function purchased( address) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractSession) Purchased(arg0 common.Address) (*big.Int, error) {
	return _MonethaVoucherContract.Contract.Purchased(&_MonethaVoucherContract.CallOpts, arg0)
}

// Purchased is a free data retrieval call binding the contract method 0x522fe98e.
//
// Solidity: function purchased( address) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) Purchased(arg0 common.Address) (*big.Int, error) {
	return _MonethaVoucherContract.Contract.Purchased(&_MonethaVoucherContract.CallOpts, arg0)
}

// PurchasedBy is a free data retrieval call binding the contract method 0x8d8c36e9.
//
// Solidity: function purchasedBy(owner address) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) PurchasedBy(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "purchasedBy", owner)
	return *ret0, err
}

// PurchasedBy is a free data retrieval call binding the contract method 0x8d8c36e9.
//
// Solidity: function purchasedBy(owner address) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractSession) PurchasedBy(owner common.Address) (*big.Int, error) {
	return _MonethaVoucherContract.Contract.PurchasedBy(&_MonethaVoucherContract.CallOpts, owner)
}

// PurchasedBy is a free data retrieval call binding the contract method 0x8d8c36e9.
//
// Solidity: function purchasedBy(owner address) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) PurchasedBy(owner common.Address) (*big.Int, error) {
	return _MonethaVoucherContract.Contract.PurchasedBy(&_MonethaVoucherContract.CallOpts, owner)
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() constant returns(string)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) Standard(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "standard")
	return *ret0, err
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() constant returns(string)
func (_MonethaVoucherContract *MonethaVoucherContractSession) Standard() (string, error) {
	return _MonethaVoucherContract.Contract.Standard(&_MonethaVoucherContract.CallOpts)
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() constant returns(string)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) Standard() (string, error) {
	return _MonethaVoucherContract.Contract.Standard(&_MonethaVoucherContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MonethaVoucherContract *MonethaVoucherContractSession) Symbol() (string, error) {
	return _MonethaVoucherContract.Contract.Symbol(&_MonethaVoucherContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) Symbol() (string, error) {
	return _MonethaVoucherContract.Contract.Symbol(&_MonethaVoucherContract.CallOpts)
}

// ToWei is a free data retrieval call binding the contract method 0x669dafe8.
//
// Solidity: function toWei(_value uint256) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) ToWei(opts *bind.CallOpts, _value *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "toWei", _value)
	return *ret0, err
}

// ToWei is a free data retrieval call binding the contract method 0x669dafe8.
//
// Solidity: function toWei(_value uint256) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractSession) ToWei(_value *big.Int) (*big.Int, error) {
	return _MonethaVoucherContract.Contract.ToWei(&_MonethaVoucherContract.CallOpts, _value)
}

// ToWei is a free data retrieval call binding the contract method 0x669dafe8.
//
// Solidity: function toWei(_value uint256) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) ToWei(_value *big.Int) (*big.Int, error) {
	return _MonethaVoucherContract.Contract.ToWei(&_MonethaVoucherContract.CallOpts, _value)
}

// TotalDistributed is a free data retrieval call binding the contract method 0xefca2eed.
//
// Solidity: function totalDistributed() constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) TotalDistributed(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "totalDistributed")
	return *ret0, err
}

// TotalDistributed is a free data retrieval call binding the contract method 0xefca2eed.
//
// Solidity: function totalDistributed() constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractSession) TotalDistributed() (*big.Int, error) {
	return _MonethaVoucherContract.Contract.TotalDistributed(&_MonethaVoucherContract.CallOpts)
}

// TotalDistributed is a free data retrieval call binding the contract method 0xefca2eed.
//
// Solidity: function totalDistributed() constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) TotalDistributed() (*big.Int, error) {
	return _MonethaVoucherContract.Contract.TotalDistributed(&_MonethaVoucherContract.CallOpts)
}

// TotalDistributedIn is a free data retrieval call binding the contract method 0x732a935c.
//
// Solidity: function totalDistributedIn( uint16) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) TotalDistributedIn(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "totalDistributedIn", arg0)
	return *ret0, err
}

// TotalDistributedIn is a free data retrieval call binding the contract method 0x732a935c.
//
// Solidity: function totalDistributedIn( uint16) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractSession) TotalDistributedIn(arg0 uint16) (*big.Int, error) {
	return _MonethaVoucherContract.Contract.TotalDistributedIn(&_MonethaVoucherContract.CallOpts, arg0)
}

// TotalDistributedIn is a free data retrieval call binding the contract method 0x732a935c.
//
// Solidity: function totalDistributedIn( uint16) constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) TotalDistributedIn(arg0 uint16) (*big.Int, error) {
	return _MonethaVoucherContract.Contract.TotalDistributedIn(&_MonethaVoucherContract.CallOpts, arg0)
}

// TotalInSharedPool is a free data retrieval call binding the contract method 0xc3fdbef8.
//
// Solidity: function totalInSharedPool() constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) TotalInSharedPool(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "totalInSharedPool")
	return *ret0, err
}

// TotalInSharedPool is a free data retrieval call binding the contract method 0xc3fdbef8.
//
// Solidity: function totalInSharedPool() constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractSession) TotalInSharedPool() (*big.Int, error) {
	return _MonethaVoucherContract.Contract.TotalInSharedPool(&_MonethaVoucherContract.CallOpts)
}

// TotalInSharedPool is a free data retrieval call binding the contract method 0xc3fdbef8.
//
// Solidity: function totalInSharedPool() constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) TotalInSharedPool() (*big.Int, error) {
	return _MonethaVoucherContract.Contract.TotalInSharedPool(&_MonethaVoucherContract.CallOpts)
}

// TotalPurchased is a free data retrieval call binding the contract method 0xe632c2f3.
//
// Solidity: function totalPurchased() constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) TotalPurchased(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "totalPurchased")
	return *ret0, err
}

// TotalPurchased is a free data retrieval call binding the contract method 0xe632c2f3.
//
// Solidity: function totalPurchased() constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractSession) TotalPurchased() (*big.Int, error) {
	return _MonethaVoucherContract.Contract.TotalPurchased(&_MonethaVoucherContract.CallOpts)
}

// TotalPurchased is a free data retrieval call binding the contract method 0xe632c2f3.
//
// Solidity: function totalPurchased() constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) TotalPurchased() (*big.Int, error) {
	return _MonethaVoucherContract.Contract.TotalPurchased(&_MonethaVoucherContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractSession) TotalSupply() (*big.Int, error) {
	return _MonethaVoucherContract.Contract.TotalSupply(&_MonethaVoucherContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) TotalSupply() (*big.Int, error) {
	return _MonethaVoucherContract.Contract.TotalSupply(&_MonethaVoucherContract.CallOpts)
}

// VoucherMthRate is a free data retrieval call binding the contract method 0xb6f81724.
//
// Solidity: function voucherMthRate() constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCaller) VoucherMthRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaVoucherContract.contract.Call(opts, out, "voucherMthRate")
	return *ret0, err
}

// VoucherMthRate is a free data retrieval call binding the contract method 0xb6f81724.
//
// Solidity: function voucherMthRate() constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractSession) VoucherMthRate() (*big.Int, error) {
	return _MonethaVoucherContract.Contract.VoucherMthRate(&_MonethaVoucherContract.CallOpts)
}

// VoucherMthRate is a free data retrieval call binding the contract method 0xb6f81724.
//
// Solidity: function voucherMthRate() constant returns(uint256)
func (_MonethaVoucherContract *MonethaVoucherContractCallerSession) VoucherMthRate() (*big.Int, error) {
	return _MonethaVoucherContract.Contract.VoucherMthRate(&_MonethaVoucherContract.CallOpts)
}

// ApplyDiscount is a paid mutator transaction binding the contract method 0x6a2b171a.
//
// Solidity: function applyDiscount(_for address, _vouchers uint256) returns(amountVouchers uint256, amountWei uint256)
func (_MonethaVoucherContract *MonethaVoucherContractTransactor) ApplyDiscount(opts *bind.TransactOpts, _for common.Address, _vouchers *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.contract.Transact(opts, "applyDiscount", _for, _vouchers)
}

// ApplyDiscount is a paid mutator transaction binding the contract method 0x6a2b171a.
//
// Solidity: function applyDiscount(_for address, _vouchers uint256) returns(amountVouchers uint256, amountWei uint256)
func (_MonethaVoucherContract *MonethaVoucherContractSession) ApplyDiscount(_for common.Address, _vouchers *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.ApplyDiscount(&_MonethaVoucherContract.TransactOpts, _for, _vouchers)
}

// ApplyDiscount is a paid mutator transaction binding the contract method 0x6a2b171a.
//
// Solidity: function applyDiscount(_for address, _vouchers uint256) returns(amountVouchers uint256, amountWei uint256)
func (_MonethaVoucherContract *MonethaVoucherContractTransactorSession) ApplyDiscount(_for common.Address, _vouchers *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.ApplyDiscount(&_MonethaVoucherContract.TransactOpts, _for, _vouchers)
}

// ApplyPayback is a paid mutator transaction binding the contract method 0x6daa212f.
//
// Solidity: function applyPayback(_for address, _amountWei uint256) returns(amountVouchers uint256)
func (_MonethaVoucherContract *MonethaVoucherContractTransactor) ApplyPayback(opts *bind.TransactOpts, _for common.Address, _amountWei *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.contract.Transact(opts, "applyPayback", _for, _amountWei)
}

// ApplyPayback is a paid mutator transaction binding the contract method 0x6daa212f.
//
// Solidity: function applyPayback(_for address, _amountWei uint256) returns(amountVouchers uint256)
func (_MonethaVoucherContract *MonethaVoucherContractSession) ApplyPayback(_for common.Address, _amountWei *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.ApplyPayback(&_MonethaVoucherContract.TransactOpts, _for, _amountWei)
}

// ApplyPayback is a paid mutator transaction binding the contract method 0x6daa212f.
//
// Solidity: function applyPayback(_for address, _amountWei uint256) returns(amountVouchers uint256)
func (_MonethaVoucherContract *MonethaVoucherContractTransactorSession) ApplyPayback(_for common.Address, _amountWei *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.ApplyPayback(&_MonethaVoucherContract.TransactOpts, _for, _amountWei)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_MonethaVoucherContract *MonethaVoucherContractTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_MonethaVoucherContract *MonethaVoucherContractSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.Approve(&_MonethaVoucherContract.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_MonethaVoucherContract *MonethaVoucherContractTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.Approve(&_MonethaVoucherContract.TransactOpts, spender, value)
}

// BuyVouchers is a paid mutator transaction binding the contract method 0xb0408721.
//
// Solidity: function buyVouchers(_vouchers uint256) returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactor) BuyVouchers(opts *bind.TransactOpts, _vouchers *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.contract.Transact(opts, "buyVouchers", _vouchers)
}

// BuyVouchers is a paid mutator transaction binding the contract method 0xb0408721.
//
// Solidity: function buyVouchers(_vouchers uint256) returns()
func (_MonethaVoucherContract *MonethaVoucherContractSession) BuyVouchers(_vouchers *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.BuyVouchers(&_MonethaVoucherContract.TransactOpts, _vouchers)
}

// BuyVouchers is a paid mutator transaction binding the contract method 0xb0408721.
//
// Solidity: function buyVouchers(_vouchers uint256) returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactorSession) BuyVouchers(_vouchers *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.BuyVouchers(&_MonethaVoucherContract.TransactOpts, _vouchers)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaVoucherContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonethaVoucherContract *MonethaVoucherContractSession) Pause() (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.Pause(&_MonethaVoucherContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactorSession) Pause() (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.Pause(&_MonethaVoucherContract.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactor) ReclaimEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaVoucherContract.contract.Transact(opts, "reclaimEther")
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_MonethaVoucherContract *MonethaVoucherContractSession) ReclaimEther() (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.ReclaimEther(&_MonethaVoucherContract.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactorSession) ReclaimEther() (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.ReclaimEther(&_MonethaVoucherContract.TransactOpts)
}

// ReclaimEtherTo is a paid mutator transaction binding the contract method 0x2e84a374.
//
// Solidity: function reclaimEtherTo(_to address, _value uint256) returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactor) ReclaimEtherTo(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.contract.Transact(opts, "reclaimEtherTo", _to, _value)
}

// ReclaimEtherTo is a paid mutator transaction binding the contract method 0x2e84a374.
//
// Solidity: function reclaimEtherTo(_to address, _value uint256) returns()
func (_MonethaVoucherContract *MonethaVoucherContractSession) ReclaimEtherTo(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.ReclaimEtherTo(&_MonethaVoucherContract.TransactOpts, _to, _value)
}

// ReclaimEtherTo is a paid mutator transaction binding the contract method 0x2e84a374.
//
// Solidity: function reclaimEtherTo(_to address, _value uint256) returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactorSession) ReclaimEtherTo(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.ReclaimEtherTo(&_MonethaVoucherContract.TransactOpts, _to, _value)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(_token address) returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactor) ReclaimToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _MonethaVoucherContract.contract.Transact(opts, "reclaimToken", _token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(_token address) returns()
func (_MonethaVoucherContract *MonethaVoucherContractSession) ReclaimToken(_token common.Address) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.ReclaimToken(&_MonethaVoucherContract.TransactOpts, _token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(_token address) returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactorSession) ReclaimToken(_token common.Address) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.ReclaimToken(&_MonethaVoucherContract.TransactOpts, _token)
}

// ReclaimTokenTo is a paid mutator transaction binding the contract method 0x9fe98bda.
//
// Solidity: function reclaimTokenTo(_token address, _to address, _value uint256) returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactor) ReclaimTokenTo(opts *bind.TransactOpts, _token common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.contract.Transact(opts, "reclaimTokenTo", _token, _to, _value)
}

// ReclaimTokenTo is a paid mutator transaction binding the contract method 0x9fe98bda.
//
// Solidity: function reclaimTokenTo(_token address, _to address, _value uint256) returns()
func (_MonethaVoucherContract *MonethaVoucherContractSession) ReclaimTokenTo(_token common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.ReclaimTokenTo(&_MonethaVoucherContract.TransactOpts, _token, _to, _value)
}

// ReclaimTokenTo is a paid mutator transaction binding the contract method 0x9fe98bda.
//
// Solidity: function reclaimTokenTo(_token address, _to address, _value uint256) returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactorSession) ReclaimTokenTo(_token common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.ReclaimTokenTo(&_MonethaVoucherContract.TransactOpts, _token, _to, _value)
}

// ReleasePurchasedTo is a paid mutator transaction binding the contract method 0xb2bea9c1.
//
// Solidity: function releasePurchasedTo(_to address, _value uint256) returns(bool)
func (_MonethaVoucherContract *MonethaVoucherContractTransactor) ReleasePurchasedTo(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.contract.Transact(opts, "releasePurchasedTo", _to, _value)
}

// ReleasePurchasedTo is a paid mutator transaction binding the contract method 0xb2bea9c1.
//
// Solidity: function releasePurchasedTo(_to address, _value uint256) returns(bool)
func (_MonethaVoucherContract *MonethaVoucherContractSession) ReleasePurchasedTo(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.ReleasePurchasedTo(&_MonethaVoucherContract.TransactOpts, _to, _value)
}

// ReleasePurchasedTo is a paid mutator transaction binding the contract method 0xb2bea9c1.
//
// Solidity: function releasePurchasedTo(_to address, _value uint256) returns(bool)
func (_MonethaVoucherContract *MonethaVoucherContractTransactorSession) ReleasePurchasedTo(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.ReleasePurchasedTo(&_MonethaVoucherContract.TransactOpts, _to, _value)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaVoucherContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MonethaVoucherContract *MonethaVoucherContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.RenounceOwnership(&_MonethaVoucherContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.RenounceOwnership(&_MonethaVoucherContract.TransactOpts)
}

// SellVouchers is a paid mutator transaction binding the contract method 0x6481f081.
//
// Solidity: function sellVouchers(_vouchers uint256) returns(weis uint256)
func (_MonethaVoucherContract *MonethaVoucherContractTransactor) SellVouchers(opts *bind.TransactOpts, _vouchers *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.contract.Transact(opts, "sellVouchers", _vouchers)
}

// SellVouchers is a paid mutator transaction binding the contract method 0x6481f081.
//
// Solidity: function sellVouchers(_vouchers uint256) returns(weis uint256)
func (_MonethaVoucherContract *MonethaVoucherContractSession) SellVouchers(_vouchers *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.SellVouchers(&_MonethaVoucherContract.TransactOpts, _vouchers)
}

// SellVouchers is a paid mutator transaction binding the contract method 0x6481f081.
//
// Solidity: function sellVouchers(_vouchers uint256) returns(weis uint256)
func (_MonethaVoucherContract *MonethaVoucherContractTransactorSession) SellVouchers(_vouchers *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.SellVouchers(&_MonethaVoucherContract.TransactOpts, _vouchers)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactor) SetMonethaAddress(opts *bind.TransactOpts, _address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _MonethaVoucherContract.contract.Transact(opts, "setMonethaAddress", _address, _isMonethaAddress)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_MonethaVoucherContract *MonethaVoucherContractSession) SetMonethaAddress(_address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.SetMonethaAddress(&_MonethaVoucherContract.TransactOpts, _address, _isMonethaAddress)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactorSession) SetMonethaAddress(_address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.SetMonethaAddress(&_MonethaVoucherContract.TransactOpts, _address, _isMonethaAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_MonethaVoucherContract *MonethaVoucherContractTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_MonethaVoucherContract *MonethaVoucherContractSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.Transfer(&_MonethaVoucherContract.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_MonethaVoucherContract *MonethaVoucherContractTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.Transfer(&_MonethaVoucherContract.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_MonethaVoucherContract *MonethaVoucherContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_MonethaVoucherContract *MonethaVoucherContractSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.TransferFrom(&_MonethaVoucherContract.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_MonethaVoucherContract *MonethaVoucherContractTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.TransferFrom(&_MonethaVoucherContract.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _MonethaVoucherContract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MonethaVoucherContract *MonethaVoucherContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.TransferOwnership(&_MonethaVoucherContract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.TransferOwnership(&_MonethaVoucherContract.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaVoucherContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonethaVoucherContract *MonethaVoucherContractSession) Unpause() (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.Unpause(&_MonethaVoucherContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.Unpause(&_MonethaVoucherContract.TransactOpts)
}

// UpdateMthEthRate is a paid mutator transaction binding the contract method 0x0ce11ded.
//
// Solidity: function updateMthEthRate(_mthEthRate uint256) returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactor) UpdateMthEthRate(opts *bind.TransactOpts, _mthEthRate *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.contract.Transact(opts, "updateMthEthRate", _mthEthRate)
}

// UpdateMthEthRate is a paid mutator transaction binding the contract method 0x0ce11ded.
//
// Solidity: function updateMthEthRate(_mthEthRate uint256) returns()
func (_MonethaVoucherContract *MonethaVoucherContractSession) UpdateMthEthRate(_mthEthRate *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.UpdateMthEthRate(&_MonethaVoucherContract.TransactOpts, _mthEthRate)
}

// UpdateMthEthRate is a paid mutator transaction binding the contract method 0x0ce11ded.
//
// Solidity: function updateMthEthRate(_mthEthRate uint256) returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactorSession) UpdateMthEthRate(_mthEthRate *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.UpdateMthEthRate(&_MonethaVoucherContract.TransactOpts, _mthEthRate)
}

// UpdateVoucherMthRate is a paid mutator transaction binding the contract method 0xb9c5eb90.
//
// Solidity: function updateVoucherMthRate(_voucherMthRate uint256) returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactor) UpdateVoucherMthRate(opts *bind.TransactOpts, _voucherMthRate *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.contract.Transact(opts, "updateVoucherMthRate", _voucherMthRate)
}

// UpdateVoucherMthRate is a paid mutator transaction binding the contract method 0xb9c5eb90.
//
// Solidity: function updateVoucherMthRate(_voucherMthRate uint256) returns()
func (_MonethaVoucherContract *MonethaVoucherContractSession) UpdateVoucherMthRate(_voucherMthRate *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.UpdateVoucherMthRate(&_MonethaVoucherContract.TransactOpts, _voucherMthRate)
}

// UpdateVoucherMthRate is a paid mutator transaction binding the contract method 0xb9c5eb90.
//
// Solidity: function updateVoucherMthRate(_voucherMthRate uint256) returns()
func (_MonethaVoucherContract *MonethaVoucherContractTransactorSession) UpdateVoucherMthRate(_voucherMthRate *big.Int) (*types.Transaction, error) {
	return _MonethaVoucherContract.Contract.UpdateVoucherMthRate(&_MonethaVoucherContract.TransactOpts, _voucherMthRate)
}

// MonethaVoucherContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MonethaVoucherContract contract.
type MonethaVoucherContractApprovalIterator struct {
	Event *MonethaVoucherContractApproval // Event containing the contract specifics and raw log

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
func (it *MonethaVoucherContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaVoucherContractApproval)
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
		it.Event = new(MonethaVoucherContractApproval)
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
func (it *MonethaVoucherContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaVoucherContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaVoucherContractApproval represents a Approval event raised by the MonethaVoucherContract contract.
type MonethaVoucherContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MonethaVoucherContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContractApprovalIterator{contract: _MonethaVoucherContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MonethaVoucherContractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaVoucherContractApproval)
				if err := _MonethaVoucherContract.contract.UnpackLog(event, "Approval", log); err != nil {
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

// MonethaVoucherContractDiscountAppliedIterator is returned from FilterDiscountApplied and is used to iterate over the raw logs and unpacked data for DiscountApplied events raised by the MonethaVoucherContract contract.
type MonethaVoucherContractDiscountAppliedIterator struct {
	Event *MonethaVoucherContractDiscountApplied // Event containing the contract specifics and raw log

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
func (it *MonethaVoucherContractDiscountAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaVoucherContractDiscountApplied)
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
		it.Event = new(MonethaVoucherContractDiscountApplied)
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
func (it *MonethaVoucherContractDiscountAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaVoucherContractDiscountAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaVoucherContractDiscountApplied represents a DiscountApplied event raised by the MonethaVoucherContract contract.
type MonethaVoucherContractDiscountApplied struct {
	User                 common.Address
	ReleasedVouchers     *big.Int
	AmountWeiTransferred *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDiscountApplied is a free log retrieval operation binding the contract event 0xa5d355764530fdfd305e1709a5c34997434bd0032f3afe650f362d1d1544b947.
//
// Solidity: e DiscountApplied(user indexed address, releasedVouchers uint256, amountWeiTransferred uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) FilterDiscountApplied(opts *bind.FilterOpts, user []common.Address) (*MonethaVoucherContractDiscountAppliedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.FilterLogs(opts, "DiscountApplied", userRule)
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContractDiscountAppliedIterator{contract: _MonethaVoucherContract.contract, event: "DiscountApplied", logs: logs, sub: sub}, nil
}

// WatchDiscountApplied is a free log subscription operation binding the contract event 0xa5d355764530fdfd305e1709a5c34997434bd0032f3afe650f362d1d1544b947.
//
// Solidity: e DiscountApplied(user indexed address, releasedVouchers uint256, amountWeiTransferred uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) WatchDiscountApplied(opts *bind.WatchOpts, sink chan<- *MonethaVoucherContractDiscountApplied, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.WatchLogs(opts, "DiscountApplied", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaVoucherContractDiscountApplied)
				if err := _MonethaVoucherContract.contract.UnpackLog(event, "DiscountApplied", log); err != nil {
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

// MonethaVoucherContractMonethaAddressSetIterator is returned from FilterMonethaAddressSet and is used to iterate over the raw logs and unpacked data for MonethaAddressSet events raised by the MonethaVoucherContract contract.
type MonethaVoucherContractMonethaAddressSetIterator struct {
	Event *MonethaVoucherContractMonethaAddressSet // Event containing the contract specifics and raw log

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
func (it *MonethaVoucherContractMonethaAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaVoucherContractMonethaAddressSet)
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
		it.Event = new(MonethaVoucherContractMonethaAddressSet)
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
func (it *MonethaVoucherContractMonethaAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaVoucherContractMonethaAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaVoucherContractMonethaAddressSet represents a MonethaAddressSet event raised by the MonethaVoucherContract contract.
type MonethaVoucherContractMonethaAddressSet struct {
	Address          common.Address
	IsMonethaAddress bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMonethaAddressSet is a free log retrieval operation binding the contract event 0xa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e692.
//
// Solidity: e MonethaAddressSet(_address address, _isMonethaAddress bool)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) FilterMonethaAddressSet(opts *bind.FilterOpts) (*MonethaVoucherContractMonethaAddressSetIterator, error) {

	logs, sub, err := _MonethaVoucherContract.contract.FilterLogs(opts, "MonethaAddressSet")
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContractMonethaAddressSetIterator{contract: _MonethaVoucherContract.contract, event: "MonethaAddressSet", logs: logs, sub: sub}, nil
}

// WatchMonethaAddressSet is a free log subscription operation binding the contract event 0xa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e692.
//
// Solidity: e MonethaAddressSet(_address address, _isMonethaAddress bool)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) WatchMonethaAddressSet(opts *bind.WatchOpts, sink chan<- *MonethaVoucherContractMonethaAddressSet) (event.Subscription, error) {

	logs, sub, err := _MonethaVoucherContract.contract.WatchLogs(opts, "MonethaAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaVoucherContractMonethaAddressSet)
				if err := _MonethaVoucherContract.contract.UnpackLog(event, "MonethaAddressSet", log); err != nil {
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

// MonethaVoucherContractMthEthRateUpdatedIterator is returned from FilterMthEthRateUpdated and is used to iterate over the raw logs and unpacked data for MthEthRateUpdated events raised by the MonethaVoucherContract contract.
type MonethaVoucherContractMthEthRateUpdatedIterator struct {
	Event *MonethaVoucherContractMthEthRateUpdated // Event containing the contract specifics and raw log

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
func (it *MonethaVoucherContractMthEthRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaVoucherContractMthEthRateUpdated)
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
		it.Event = new(MonethaVoucherContractMthEthRateUpdated)
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
func (it *MonethaVoucherContractMthEthRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaVoucherContractMthEthRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaVoucherContractMthEthRateUpdated represents a MthEthRateUpdated event raised by the MonethaVoucherContract contract.
type MonethaVoucherContractMthEthRateUpdated struct {
	OldMthEthRate *big.Int
	NewMthEthRate *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMthEthRateUpdated is a free log retrieval operation binding the contract event 0xf376950c4c79df1231c4d2cc604341aff0685bc9b6c3a1661200a6d4f319d4e4.
//
// Solidity: e MthEthRateUpdated(oldMthEthRate uint256, newMthEthRate uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) FilterMthEthRateUpdated(opts *bind.FilterOpts) (*MonethaVoucherContractMthEthRateUpdatedIterator, error) {

	logs, sub, err := _MonethaVoucherContract.contract.FilterLogs(opts, "MthEthRateUpdated")
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContractMthEthRateUpdatedIterator{contract: _MonethaVoucherContract.contract, event: "MthEthRateUpdated", logs: logs, sub: sub}, nil
}

// WatchMthEthRateUpdated is a free log subscription operation binding the contract event 0xf376950c4c79df1231c4d2cc604341aff0685bc9b6c3a1661200a6d4f319d4e4.
//
// Solidity: e MthEthRateUpdated(oldMthEthRate uint256, newMthEthRate uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) WatchMthEthRateUpdated(opts *bind.WatchOpts, sink chan<- *MonethaVoucherContractMthEthRateUpdated) (event.Subscription, error) {

	logs, sub, err := _MonethaVoucherContract.contract.WatchLogs(opts, "MthEthRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaVoucherContractMthEthRateUpdated)
				if err := _MonethaVoucherContract.contract.UnpackLog(event, "MthEthRateUpdated", log); err != nil {
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

// MonethaVoucherContractOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the MonethaVoucherContract contract.
type MonethaVoucherContractOwnershipRenouncedIterator struct {
	Event *MonethaVoucherContractOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *MonethaVoucherContractOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaVoucherContractOwnershipRenounced)
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
		it.Event = new(MonethaVoucherContractOwnershipRenounced)
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
func (it *MonethaVoucherContractOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaVoucherContractOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaVoucherContractOwnershipRenounced represents a OwnershipRenounced event raised by the MonethaVoucherContract contract.
type MonethaVoucherContractOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*MonethaVoucherContractOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContractOwnershipRenouncedIterator{contract: _MonethaVoucherContract.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *MonethaVoucherContractOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaVoucherContractOwnershipRenounced)
				if err := _MonethaVoucherContract.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// MonethaVoucherContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MonethaVoucherContract contract.
type MonethaVoucherContractOwnershipTransferredIterator struct {
	Event *MonethaVoucherContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MonethaVoucherContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaVoucherContractOwnershipTransferred)
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
		it.Event = new(MonethaVoucherContractOwnershipTransferred)
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
func (it *MonethaVoucherContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaVoucherContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaVoucherContractOwnershipTransferred represents a OwnershipTransferred event raised by the MonethaVoucherContract contract.
type MonethaVoucherContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MonethaVoucherContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContractOwnershipTransferredIterator{contract: _MonethaVoucherContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MonethaVoucherContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaVoucherContractOwnershipTransferred)
				if err := _MonethaVoucherContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// MonethaVoucherContractPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the MonethaVoucherContract contract.
type MonethaVoucherContractPauseIterator struct {
	Event *MonethaVoucherContractPause // Event containing the contract specifics and raw log

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
func (it *MonethaVoucherContractPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaVoucherContractPause)
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
		it.Event = new(MonethaVoucherContractPause)
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
func (it *MonethaVoucherContractPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaVoucherContractPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaVoucherContractPause represents a Pause event raised by the MonethaVoucherContract contract.
type MonethaVoucherContractPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) FilterPause(opts *bind.FilterOpts) (*MonethaVoucherContractPauseIterator, error) {

	logs, sub, err := _MonethaVoucherContract.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContractPauseIterator{contract: _MonethaVoucherContract.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *MonethaVoucherContractPause) (event.Subscription, error) {

	logs, sub, err := _MonethaVoucherContract.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaVoucherContractPause)
				if err := _MonethaVoucherContract.contract.UnpackLog(event, "Pause", log); err != nil {
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

// MonethaVoucherContractPaybackAppliedIterator is returned from FilterPaybackApplied and is used to iterate over the raw logs and unpacked data for PaybackApplied events raised by the MonethaVoucherContract contract.
type MonethaVoucherContractPaybackAppliedIterator struct {
	Event *MonethaVoucherContractPaybackApplied // Event containing the contract specifics and raw log

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
func (it *MonethaVoucherContractPaybackAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaVoucherContractPaybackApplied)
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
		it.Event = new(MonethaVoucherContractPaybackApplied)
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
func (it *MonethaVoucherContractPaybackAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaVoucherContractPaybackAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaVoucherContractPaybackApplied represents a PaybackApplied event raised by the MonethaVoucherContract contract.
type MonethaVoucherContractPaybackApplied struct {
	User                common.Address
	AddedVouchers       *big.Int
	AmountWeiEquivalent *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterPaybackApplied is a free log retrieval operation binding the contract event 0x289266f290897c2e82d0cee6b6f74f7d5e62d30244d4bfcf390abc4af4cd5666.
//
// Solidity: e PaybackApplied(user indexed address, addedVouchers uint256, amountWeiEquivalent uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) FilterPaybackApplied(opts *bind.FilterOpts, user []common.Address) (*MonethaVoucherContractPaybackAppliedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.FilterLogs(opts, "PaybackApplied", userRule)
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContractPaybackAppliedIterator{contract: _MonethaVoucherContract.contract, event: "PaybackApplied", logs: logs, sub: sub}, nil
}

// WatchPaybackApplied is a free log subscription operation binding the contract event 0x289266f290897c2e82d0cee6b6f74f7d5e62d30244d4bfcf390abc4af4cd5666.
//
// Solidity: e PaybackApplied(user indexed address, addedVouchers uint256, amountWeiEquivalent uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) WatchPaybackApplied(opts *bind.WatchOpts, sink chan<- *MonethaVoucherContractPaybackApplied, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.WatchLogs(opts, "PaybackApplied", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaVoucherContractPaybackApplied)
				if err := _MonethaVoucherContract.contract.UnpackLog(event, "PaybackApplied", log); err != nil {
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

// MonethaVoucherContractPurchasedVouchersReleasedIterator is returned from FilterPurchasedVouchersReleased and is used to iterate over the raw logs and unpacked data for PurchasedVouchersReleased events raised by the MonethaVoucherContract contract.
type MonethaVoucherContractPurchasedVouchersReleasedIterator struct {
	Event *MonethaVoucherContractPurchasedVouchersReleased // Event containing the contract specifics and raw log

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
func (it *MonethaVoucherContractPurchasedVouchersReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaVoucherContractPurchasedVouchersReleased)
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
		it.Event = new(MonethaVoucherContractPurchasedVouchersReleased)
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
func (it *MonethaVoucherContractPurchasedVouchersReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaVoucherContractPurchasedVouchersReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaVoucherContractPurchasedVouchersReleased represents a PurchasedVouchersReleased event raised by the MonethaVoucherContract contract.
type MonethaVoucherContractPurchasedVouchersReleased struct {
	From     common.Address
	To       common.Address
	Vouchers *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPurchasedVouchersReleased is a free log retrieval operation binding the contract event 0xd49597d4844c4d37387a6053e4a1832bd054dee0aaf72fcd84da7e2aa9f8cbf0.
//
// Solidity: e PurchasedVouchersReleased(from indexed address, to indexed address, vouchers uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) FilterPurchasedVouchersReleased(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MonethaVoucherContractPurchasedVouchersReleasedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.FilterLogs(opts, "PurchasedVouchersReleased", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContractPurchasedVouchersReleasedIterator{contract: _MonethaVoucherContract.contract, event: "PurchasedVouchersReleased", logs: logs, sub: sub}, nil
}

// WatchPurchasedVouchersReleased is a free log subscription operation binding the contract event 0xd49597d4844c4d37387a6053e4a1832bd054dee0aaf72fcd84da7e2aa9f8cbf0.
//
// Solidity: e PurchasedVouchersReleased(from indexed address, to indexed address, vouchers uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) WatchPurchasedVouchersReleased(opts *bind.WatchOpts, sink chan<- *MonethaVoucherContractPurchasedVouchersReleased, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.WatchLogs(opts, "PurchasedVouchersReleased", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaVoucherContractPurchasedVouchersReleased)
				if err := _MonethaVoucherContract.contract.UnpackLog(event, "PurchasedVouchersReleased", log); err != nil {
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

// MonethaVoucherContractReclaimEtherIterator is returned from FilterReclaimEther and is used to iterate over the raw logs and unpacked data for ReclaimEther events raised by the MonethaVoucherContract contract.
type MonethaVoucherContractReclaimEtherIterator struct {
	Event *MonethaVoucherContractReclaimEther // Event containing the contract specifics and raw log

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
func (it *MonethaVoucherContractReclaimEtherIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaVoucherContractReclaimEther)
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
		it.Event = new(MonethaVoucherContractReclaimEther)
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
func (it *MonethaVoucherContractReclaimEtherIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaVoucherContractReclaimEtherIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaVoucherContractReclaimEther represents a ReclaimEther event raised by the MonethaVoucherContract contract.
type MonethaVoucherContractReclaimEther struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReclaimEther is a free log retrieval operation binding the contract event 0xb54913b2b58b2e96ea9b4e96ba2353cf13426af9d3f252e0c17899a93c4ce125.
//
// Solidity: e ReclaimEther(to indexed address, amount uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) FilterReclaimEther(opts *bind.FilterOpts, to []common.Address) (*MonethaVoucherContractReclaimEtherIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.FilterLogs(opts, "ReclaimEther", toRule)
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContractReclaimEtherIterator{contract: _MonethaVoucherContract.contract, event: "ReclaimEther", logs: logs, sub: sub}, nil
}

// WatchReclaimEther is a free log subscription operation binding the contract event 0xb54913b2b58b2e96ea9b4e96ba2353cf13426af9d3f252e0c17899a93c4ce125.
//
// Solidity: e ReclaimEther(to indexed address, amount uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) WatchReclaimEther(opts *bind.WatchOpts, sink chan<- *MonethaVoucherContractReclaimEther, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.WatchLogs(opts, "ReclaimEther", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaVoucherContractReclaimEther)
				if err := _MonethaVoucherContract.contract.UnpackLog(event, "ReclaimEther", log); err != nil {
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

// MonethaVoucherContractReclaimTokensIterator is returned from FilterReclaimTokens and is used to iterate over the raw logs and unpacked data for ReclaimTokens events raised by the MonethaVoucherContract contract.
type MonethaVoucherContractReclaimTokensIterator struct {
	Event *MonethaVoucherContractReclaimTokens // Event containing the contract specifics and raw log

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
func (it *MonethaVoucherContractReclaimTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaVoucherContractReclaimTokens)
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
		it.Event = new(MonethaVoucherContractReclaimTokens)
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
func (it *MonethaVoucherContractReclaimTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaVoucherContractReclaimTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaVoucherContractReclaimTokens represents a ReclaimTokens event raised by the MonethaVoucherContract contract.
type MonethaVoucherContractReclaimTokens struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReclaimTokens is a free log retrieval operation binding the contract event 0x355069f20974db323c9dcd100e8bf13fb2acc1884e5ec05b0a89c09e15ce810f.
//
// Solidity: e ReclaimTokens(to indexed address, amount uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) FilterReclaimTokens(opts *bind.FilterOpts, to []common.Address) (*MonethaVoucherContractReclaimTokensIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.FilterLogs(opts, "ReclaimTokens", toRule)
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContractReclaimTokensIterator{contract: _MonethaVoucherContract.contract, event: "ReclaimTokens", logs: logs, sub: sub}, nil
}

// WatchReclaimTokens is a free log subscription operation binding the contract event 0x355069f20974db323c9dcd100e8bf13fb2acc1884e5ec05b0a89c09e15ce810f.
//
// Solidity: e ReclaimTokens(to indexed address, amount uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) WatchReclaimTokens(opts *bind.WatchOpts, sink chan<- *MonethaVoucherContractReclaimTokens, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.WatchLogs(opts, "ReclaimTokens", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaVoucherContractReclaimTokens)
				if err := _MonethaVoucherContract.contract.UnpackLog(event, "ReclaimTokens", log); err != nil {
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

// MonethaVoucherContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MonethaVoucherContract contract.
type MonethaVoucherContractTransferIterator struct {
	Event *MonethaVoucherContractTransfer // Event containing the contract specifics and raw log

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
func (it *MonethaVoucherContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaVoucherContractTransfer)
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
		it.Event = new(MonethaVoucherContractTransfer)
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
func (it *MonethaVoucherContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaVoucherContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaVoucherContractTransfer represents a Transfer event raised by the MonethaVoucherContract contract.
type MonethaVoucherContractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MonethaVoucherContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContractTransferIterator{contract: _MonethaVoucherContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MonethaVoucherContractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaVoucherContractTransfer)
				if err := _MonethaVoucherContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// MonethaVoucherContractUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the MonethaVoucherContract contract.
type MonethaVoucherContractUnpauseIterator struct {
	Event *MonethaVoucherContractUnpause // Event containing the contract specifics and raw log

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
func (it *MonethaVoucherContractUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaVoucherContractUnpause)
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
		it.Event = new(MonethaVoucherContractUnpause)
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
func (it *MonethaVoucherContractUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaVoucherContractUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaVoucherContractUnpause represents a Unpause event raised by the MonethaVoucherContract contract.
type MonethaVoucherContractUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) FilterUnpause(opts *bind.FilterOpts) (*MonethaVoucherContractUnpauseIterator, error) {

	logs, sub, err := _MonethaVoucherContract.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContractUnpauseIterator{contract: _MonethaVoucherContract.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *MonethaVoucherContractUnpause) (event.Subscription, error) {

	logs, sub, err := _MonethaVoucherContract.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaVoucherContractUnpause)
				if err := _MonethaVoucherContract.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// MonethaVoucherContractVoucherMthRateUpdatedIterator is returned from FilterVoucherMthRateUpdated and is used to iterate over the raw logs and unpacked data for VoucherMthRateUpdated events raised by the MonethaVoucherContract contract.
type MonethaVoucherContractVoucherMthRateUpdatedIterator struct {
	Event *MonethaVoucherContractVoucherMthRateUpdated // Event containing the contract specifics and raw log

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
func (it *MonethaVoucherContractVoucherMthRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaVoucherContractVoucherMthRateUpdated)
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
		it.Event = new(MonethaVoucherContractVoucherMthRateUpdated)
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
func (it *MonethaVoucherContractVoucherMthRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaVoucherContractVoucherMthRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaVoucherContractVoucherMthRateUpdated represents a VoucherMthRateUpdated event raised by the MonethaVoucherContract contract.
type MonethaVoucherContractVoucherMthRateUpdated struct {
	OldVoucherMthRate *big.Int
	NewVoucherMthRate *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterVoucherMthRateUpdated is a free log retrieval operation binding the contract event 0x5f8dd084970454563d6e013038a3a2204c50015701c39e62c3429738026fa784.
//
// Solidity: e VoucherMthRateUpdated(oldVoucherMthRate uint256, newVoucherMthRate uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) FilterVoucherMthRateUpdated(opts *bind.FilterOpts) (*MonethaVoucherContractVoucherMthRateUpdatedIterator, error) {

	logs, sub, err := _MonethaVoucherContract.contract.FilterLogs(opts, "VoucherMthRateUpdated")
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContractVoucherMthRateUpdatedIterator{contract: _MonethaVoucherContract.contract, event: "VoucherMthRateUpdated", logs: logs, sub: sub}, nil
}

// WatchVoucherMthRateUpdated is a free log subscription operation binding the contract event 0x5f8dd084970454563d6e013038a3a2204c50015701c39e62c3429738026fa784.
//
// Solidity: e VoucherMthRateUpdated(oldVoucherMthRate uint256, newVoucherMthRate uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) WatchVoucherMthRateUpdated(opts *bind.WatchOpts, sink chan<- *MonethaVoucherContractVoucherMthRateUpdated) (event.Subscription, error) {

	logs, sub, err := _MonethaVoucherContract.contract.WatchLogs(opts, "VoucherMthRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaVoucherContractVoucherMthRateUpdated)
				if err := _MonethaVoucherContract.contract.UnpackLog(event, "VoucherMthRateUpdated", log); err != nil {
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

// MonethaVoucherContractVoucherReleasedIterator is returned from FilterVoucherReleased and is used to iterate over the raw logs and unpacked data for VoucherReleased events raised by the MonethaVoucherContract contract.
type MonethaVoucherContractVoucherReleasedIterator struct {
	Event *MonethaVoucherContractVoucherReleased // Event containing the contract specifics and raw log

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
func (it *MonethaVoucherContractVoucherReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaVoucherContractVoucherReleased)
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
		it.Event = new(MonethaVoucherContractVoucherReleased)
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
func (it *MonethaVoucherContractVoucherReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaVoucherContractVoucherReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaVoucherContractVoucherReleased represents a VoucherReleased event raised by the MonethaVoucherContract contract.
type MonethaVoucherContractVoucherReleased struct {
	User            common.Address
	ReleasedVoucher *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVoucherReleased is a free log retrieval operation binding the contract event 0x5c2b936c2e3089a9ba68c1146fc951f328fd49b4b1a5c00156babd126b227a45.
//
// Solidity: e VoucherReleased(user indexed address, releasedVoucher uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) FilterVoucherReleased(opts *bind.FilterOpts, user []common.Address) (*MonethaVoucherContractVoucherReleasedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.FilterLogs(opts, "VoucherReleased", userRule)
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContractVoucherReleasedIterator{contract: _MonethaVoucherContract.contract, event: "VoucherReleased", logs: logs, sub: sub}, nil
}

// WatchVoucherReleased is a free log subscription operation binding the contract event 0x5c2b936c2e3089a9ba68c1146fc951f328fd49b4b1a5c00156babd126b227a45.
//
// Solidity: e VoucherReleased(user indexed address, releasedVoucher uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) WatchVoucherReleased(opts *bind.WatchOpts, sink chan<- *MonethaVoucherContractVoucherReleased, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.WatchLogs(opts, "VoucherReleased", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaVoucherContractVoucherReleased)
				if err := _MonethaVoucherContract.contract.UnpackLog(event, "VoucherReleased", log); err != nil {
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

// MonethaVoucherContractVouchersAddedIterator is returned from FilterVouchersAdded and is used to iterate over the raw logs and unpacked data for VouchersAdded events raised by the MonethaVoucherContract contract.
type MonethaVoucherContractVouchersAddedIterator struct {
	Event *MonethaVoucherContractVouchersAdded // Event containing the contract specifics and raw log

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
func (it *MonethaVoucherContractVouchersAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaVoucherContractVouchersAdded)
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
		it.Event = new(MonethaVoucherContractVouchersAdded)
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
func (it *MonethaVoucherContractVouchersAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaVoucherContractVouchersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaVoucherContractVouchersAdded represents a VouchersAdded event raised by the MonethaVoucherContract contract.
type MonethaVoucherContractVouchersAdded struct {
	User          common.Address
	VouchersAdded *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVouchersAdded is a free log retrieval operation binding the contract event 0xa4822228ba5999ec50d98259975e2e7e3b56a666bfd3cae8c7722ee90af63a28.
//
// Solidity: e VouchersAdded(user indexed address, vouchersAdded uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) FilterVouchersAdded(opts *bind.FilterOpts, user []common.Address) (*MonethaVoucherContractVouchersAddedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.FilterLogs(opts, "VouchersAdded", userRule)
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContractVouchersAddedIterator{contract: _MonethaVoucherContract.contract, event: "VouchersAdded", logs: logs, sub: sub}, nil
}

// WatchVouchersAdded is a free log subscription operation binding the contract event 0xa4822228ba5999ec50d98259975e2e7e3b56a666bfd3cae8c7722ee90af63a28.
//
// Solidity: e VouchersAdded(user indexed address, vouchersAdded uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) WatchVouchersAdded(opts *bind.WatchOpts, sink chan<- *MonethaVoucherContractVouchersAdded, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.WatchLogs(opts, "VouchersAdded", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaVoucherContractVouchersAdded)
				if err := _MonethaVoucherContract.contract.UnpackLog(event, "VouchersAdded", log); err != nil {
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

// MonethaVoucherContractVouchersBoughtIterator is returned from FilterVouchersBought and is used to iterate over the raw logs and unpacked data for VouchersBought events raised by the MonethaVoucherContract contract.
type MonethaVoucherContractVouchersBoughtIterator struct {
	Event *MonethaVoucherContractVouchersBought // Event containing the contract specifics and raw log

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
func (it *MonethaVoucherContractVouchersBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaVoucherContractVouchersBought)
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
		it.Event = new(MonethaVoucherContractVouchersBought)
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
func (it *MonethaVoucherContractVouchersBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaVoucherContractVouchersBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaVoucherContractVouchersBought represents a VouchersBought event raised by the MonethaVoucherContract contract.
type MonethaVoucherContractVouchersBought struct {
	User           common.Address
	VouchersBought *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVouchersBought is a free log retrieval operation binding the contract event 0xad6ca73eec0e00d9089e39fad2962bb18992dd5feeae519dcb4cef2aabd93e30.
//
// Solidity: e VouchersBought(user indexed address, vouchersBought uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) FilterVouchersBought(opts *bind.FilterOpts, user []common.Address) (*MonethaVoucherContractVouchersBoughtIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.FilterLogs(opts, "VouchersBought", userRule)
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContractVouchersBoughtIterator{contract: _MonethaVoucherContract.contract, event: "VouchersBought", logs: logs, sub: sub}, nil
}

// WatchVouchersBought is a free log subscription operation binding the contract event 0xad6ca73eec0e00d9089e39fad2962bb18992dd5feeae519dcb4cef2aabd93e30.
//
// Solidity: e VouchersBought(user indexed address, vouchersBought uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) WatchVouchersBought(opts *bind.WatchOpts, sink chan<- *MonethaVoucherContractVouchersBought, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.WatchLogs(opts, "VouchersBought", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaVoucherContractVouchersBought)
				if err := _MonethaVoucherContract.contract.UnpackLog(event, "VouchersBought", log); err != nil {
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

// MonethaVoucherContractVouchersSoldIterator is returned from FilterVouchersSold and is used to iterate over the raw logs and unpacked data for VouchersSold events raised by the MonethaVoucherContract contract.
type MonethaVoucherContractVouchersSoldIterator struct {
	Event *MonethaVoucherContractVouchersSold // Event containing the contract specifics and raw log

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
func (it *MonethaVoucherContractVouchersSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaVoucherContractVouchersSold)
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
		it.Event = new(MonethaVoucherContractVouchersSold)
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
func (it *MonethaVoucherContractVouchersSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaVoucherContractVouchersSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaVoucherContractVouchersSold represents a VouchersSold event raised by the MonethaVoucherContract contract.
type MonethaVoucherContractVouchersSold struct {
	User                 common.Address
	VouchersSold         *big.Int
	AmountWeiTransferred *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterVouchersSold is a free log retrieval operation binding the contract event 0x4f4ec1e156ff2a5d94d15690727a0c7c94705b9dd53e5e0abf3fa16195b106e2.
//
// Solidity: e VouchersSold(user indexed address, vouchersSold uint256, amountWeiTransferred uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) FilterVouchersSold(opts *bind.FilterOpts, user []common.Address) (*MonethaVoucherContractVouchersSoldIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.FilterLogs(opts, "VouchersSold", userRule)
	if err != nil {
		return nil, err
	}
	return &MonethaVoucherContractVouchersSoldIterator{contract: _MonethaVoucherContract.contract, event: "VouchersSold", logs: logs, sub: sub}, nil
}

// WatchVouchersSold is a free log subscription operation binding the contract event 0x4f4ec1e156ff2a5d94d15690727a0c7c94705b9dd53e5e0abf3fa16195b106e2.
//
// Solidity: e VouchersSold(user indexed address, vouchersSold uint256, amountWeiTransferred uint256)
func (_MonethaVoucherContract *MonethaVoucherContractFilterer) WatchVouchersSold(opts *bind.WatchOpts, sink chan<- *MonethaVoucherContractVouchersSold, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MonethaVoucherContract.contract.WatchLogs(opts, "VouchersSold", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaVoucherContractVouchersSold)
				if err := _MonethaVoucherContract.contract.UnpackLog(event, "VouchersSold", log); err != nil {
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

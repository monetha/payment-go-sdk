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

// MonethaTokenHoldersProgramContractABI is the input ABI used to generate the binding from.
const MonethaTokenHoldersProgramContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalStacked\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"reclaimEtherTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isMonethaAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakedBy\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mthToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reclaimEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"reclaimTokenTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"monethaVoucher\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"participateFromTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_isMonethaAddress\",\"type\":\"bool\"}],\"name\":\"setMonethaAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_mthToken\",\"type\":\"address\"},{\"name\":\"_monethaVoucher\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vouchers\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"weis\",\"type\":\"uint256\"}],\"name\":\"VouchersPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vouchers\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"weis\",\"type\":\"uint256\"}],\"name\":\"VouchersSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"mthTokens\",\"type\":\"uint256\"}],\"name\":\"ParticipationStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"mthTokens\",\"type\":\"uint256\"}],\"name\":\"ParticipationStopped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"vouchers\",\"type\":\"uint256\"}],\"name\":\"VouchersRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReclaimTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReclaimEther\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_isMonethaAddress\",\"type\":\"bool\"}],\"name\":\"MonethaAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"buyVouchers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"sellVouchers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isAllowedToParticipateNow\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"participate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isAllowedToRedeemNow\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"redeem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cancelParticipation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MonethaTokenHoldersProgramContractBin is the compiled bytecode used for deploying new contracts.
const MonethaTokenHoldersProgramContractBin = `0x60806040526002805460ff1916905534801561001a57600080fd5b50604051604080611a5983398101604052805160209091015160008054600160a060020a03191633179055600160a060020a03811615156100bc57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f6d7573742062652076616c696420616464726573730000000000000000000000604482015290519081900360640190fd5b600160a060020a038216151561013357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f6d7573742062652076616c696420616464726573730000000000000000000000604482015290519081900360640190fd5b6002805461010060a860020a031916610100600160a060020a039485160217905560038054600160a060020a031916919092161790556000196004556118db8061017e6000396000f3006080604052600436106101245763ffffffff60e060020a60003504166317623c41811461014457806317ffc3201461016d57806329eae70d1461018e5780632e84a374146101b557806331d41325146101d95780633f4ba83a146101fa5780635c975abb1461020f5780635f56a31a14610224578063715018a61461024557806380bfc5991461025a5780638456cb591461026f5780638758570714610284578063899bf897146102995780638da5cb5b146102ae57806393306c43146102df5780639f727c27146102f45780639fe98bda14610309578063a7abdf0314610333578063b186099714610348578063be040fb01461035d578063c07e339114610372578063d11711a214610398578063dae37abc146103ad578063f2fde38b146103c2575b3360009081526001602052604090205460ff16151561014257600080fd5b005b34801561015057600080fd5b506101596103e3565b604080519115158252519081900360200190f35b34801561017957600080fd5b50610142600160a060020a0360043516610400565b34801561019a57600080fd5b506101a361050e565b60408051918252519081900360200190f35b3480156101c157600080fd5b50610142600160a060020a0360043516602435610514565b3480156101e557600080fd5b50610159600160a060020a0360043516610605565b34801561020657600080fd5b5061014261061a565b34801561021b57600080fd5b50610159610677565b34801561023057600080fd5b506101a3600160a060020a0360043516610680565b34801561025157600080fd5b50610142610692565b34801561026657600080fd5b506101426106fe565b34801561027b57600080fd5b50610142610890565b34801561029057600080fd5b506101426108ef565b3480156102a557600080fd5b50610159610c5f565b3480156102ba57600080fd5b506102c3610c77565b60408051600160a060020a039092168252519081900360200190f35b3480156102eb57600080fd5b506102c3610c86565b34801561030057600080fd5b50610142610c9a565b34801561031557600080fd5b50610142600160a060020a0360043581169060243516604435610d35565b34801561033f57600080fd5b506102c3610e0a565b34801561035457600080fd5b506101a3610e19565b34801561036957600080fd5b50610142610e1f565b34801561037e57600080fd5b50610142600160a060020a036004351660243515156110f4565b3480156103a457600080fd5b5061014261116f565b3480156103b957600080fd5b506101426113e3565b3480156103ce57600080fd5b50610142600160a060020a03600435166113ef565b600060045442101580156103fb57506103fb42611412565b905090565b60008054600160a060020a0316331461041857600080fd5b604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051600160a060020a038416916370a082319160248083019260209291908290030181600087803b15801561047957600080fd5b505af115801561048d573d6000803e3d6000fd5b505050506040513d60208110156104a357600080fd5b50516000549091506104c890600160a060020a0384811691168363ffffffff61142a16565b600054604080518381529051600160a060020a03909216917f355069f20974db323c9dcd100e8bf13fb2acc1884e5ec05b0a89c09e15ce810f9181900360200190a25050565b60065481565b600054600160a060020a0316331461052b57600080fd5b600160a060020a038216151561058b576040805160e560020a62461bcd02815260206004820152601b60248201527f7a65726f2061646472657373206973206e6f7420616c6c6f7765640000000000604482015290519081900360640190fd5b604051600160a060020a0383169082156108fc029083906000818181858888f193505050501580156105c1573d6000803e3d6000fd5b50604080518281529051600160a060020a038416917fb54913b2b58b2e96ea9b4e96ba2353cf13426af9d3f252e0c17899a93c4ce125919081900360200190a25050565b60016020526000908152604090205460ff1681565b600054600160a060020a0316331461063157600080fd5b60025460ff16151561064257600080fd5b6002805460ff191690556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3390600090a1565b60025460ff1681565b60056020526000908152604090205481565b600054600160a060020a031633146106a957600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b33600090815260016020526040812054819060ff16151561071e57600080fd5b6000196004908155600354604080517f8d8c36e9000000000000000000000000000000000000000000000000000000008152309381019390935251600160a060020a0390911691638d8c36e99160248083019260209291908290030181600087803b15801561078c57600080fd5b505af11580156107a0573d6000803e3d6000fd5b505050506040513d60208110156107b657600080fd5b5051600354604080517f6481f081000000000000000000000000000000000000000000000000000000008152600481018490529051929450600160a060020a0390911691636481f081916024808201926020929091908290030181600087803b15801561082257600080fd5b505af1158015610836573d6000803e3d6000fd5b505050506040513d602081101561084c57600080fd5b5051604080518481526020810183905281519293507f707c653f7d61b8660b51074cc8c2d9e99f9c34436ea908d62bea7a5a4cb18d4f929081900390910190a15050565b600054600160a060020a031633146108a757600080fd5b60025460ff16156108b757600080fd5b6002805460ff191660011790556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62590600090a1565b33600090815260016020526040812054819081908190819060ff16151561091557600080fd5b3031945060008511610971576040805160e560020a62461bcd02815260206004820152601760248201527f706f7369746976652062616c616e6365206e6565646564000000000000000000604482015290519081900360640190fd5b600360009054906101000a9004600160a060020a0316600160a060020a031663c3fdbef86040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156109c457600080fd5b505af11580156109d8573d6000803e3d6000fd5b505050506040513d60208110156109ee57600080fd5b5051935060008411610a4a576040805160e560020a62461bcd02815260206004820152601560248201527f6e6f20766f75636865727320617661696c61626c650000000000000000000000604482015290519081900360640190fd5b600354604080517fe65d1522000000000000000000000000000000000000000000000000000000008152303160048201529051600160a060020a039092169163e65d1522916024808201926020929091908290030181600087803b158015610ab157600080fd5b505af1158015610ac5573d6000803e3d6000fd5b505050506040513d6020811015610adb57600080fd5b5051925083831115610aeb578392505b600354604080517f669dafe8000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a039092169163669dafe8916024808201926020929091908290030181600087803b158015610b5257600080fd5b505af1158015610b66573d6000803e3d6000fd5b505050506040513d6020811015610b7c57600080fd5b50519450610b89426114c9565b5091509150610b988282611566565b6004908155600354604080517fb040872100000000000000000000000000000000000000000000000000000000815292830186905251600160a060020a039091169163b040872191889160248082019260009290919082900301818588803b158015610c0357600080fd5b505af1158015610c17573d6000803e3d6000fd5b505060408051878152602081018a905281517f9fca5352dcf7d3de26b4fcc50718ab421678e3c4cb800a589bdc49fbdee51a7995509081900390910192509050a15050505050565b600060045442101580156103fb57506103fb42611592565b600054600160a060020a031681565b6002546101009004600160a060020a031681565b60008054600160a060020a03163314610cb257600080fd5b5060008054604051303192600160a060020a03909216916108fc841502918491818181858888f19350505050158015610cef573d6000803e3d6000fd5b50600054604080518381529051600160a060020a03909216917fb54913b2b58b2e96ea9b4e96ba2353cf13426af9d3f252e0c17899a93c4ce1259181900360200190a250565b600054600160a060020a03163314610d4c57600080fd5b600160a060020a0382161515610dac576040805160e560020a62461bcd02815260206004820152601b60248201527f7a65726f2061646472657373206973206e6f7420616c6c6f7765640000000000604482015290519081900360640190fd5b610dc6600160a060020a038416838363ffffffff61142a16565b604080518281529051600160a060020a038416917f355069f20974db323c9dcd100e8bf13fb2acc1884e5ec05b0a89c09e15ce810f919081900360200190a2505050565b600354600160a060020a031681565b60045481565b6000806000806004544210151515610e81576040805160e560020a62461bcd02815260206004820152601360248201527f746f6f206561726c7920746f2072656465656d00000000000000000000000000604482015290519081900360640190fd5b610e8a42611592565b1515610f06576040805160e560020a62461bcd02815260206004820152602360248201527f72656465656d206973206e6f7420616c6c6f77656420617420746865206d6f6d60448201527f656e740000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b610f0e6115d7565b600354604080517f8d8c36e90000000000000000000000000000000000000000000000000000000081523060048201529051939750919550600160a060020a031691638d8c36e9916024808201926020929091908290030181600087803b158015610f7857600080fd5b505af1158015610f8c573d6000803e3d6000fd5b505050506040513d6020811015610fa257600080fd5b50519150610fc683610fba848763ffffffff6116bc16565b9063ffffffff6116e516565b600354604080517fb2bea9c1000000000000000000000000000000000000000000000000000000008152336004820152602481018490529051929350600160a060020a039091169163b2bea9c1916044808201926020929091908290030181600087803b15801561103657600080fd5b505af115801561104a573d6000803e3d6000fd5b505050506040513d602081101561106057600080fd5b505115156110b8576040805160e560020a62461bcd02815260206004820152601960248201527f766f75636865727320776173206e6f742072656c656173656400000000000000604482015290519081900360640190fd5b60408051828152905133917fb7dd3fa0284dca52a2a1e7985ef6eb7d2bd8a4fb3388f8baf55a64fd5ff4c94a919081900360200190a250505050565b600054600160a060020a0316331461110b57600080fd5b600160a060020a038216600081815260016020908152604091829020805460ff191685151590811790915582519384529083015280517fa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e6929281900390910190a15050565b6004546000904210156111cc576040805160e560020a62461bcd02815260206004820152601860248201527f746f6f206561726c7920746f2070617274696369706174650000000000000000604482015290519081900360640190fd5b6111d542611412565b1515611251576040805160e560020a62461bcd02815260206004820152602960248201527f7061727469636970617465206f6e207468652031737420646179206f6620657660448201527f657279206d6f6e74680000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600254604080517fdd62ed3e0000000000000000000000000000000000000000000000000000000081523360048201523060248201529051610100909204600160a060020a03169163dd62ed3e916044808201926020929091908290030181600087803b1580156112c157600080fd5b505af11580156112d5573d6000803e3d6000fd5b505050506040513d60208110156112eb57600080fd5b5051905060008111611347576040805160e560020a62461bcd02815260206004820152601960248201527f706f73697469766520616c6c6f77616e6365206e656564656400000000000000604482015290519081900360640190fd5b600254611364906101009004600160a060020a03163330846116fa565b33600090815260056020526040902054611384908263ffffffff6117ab16565b336000908152600560205260409020556006546113a7908263ffffffff6117ab16565b60065560408051828152905133917fbd6451f729baeebe302c84f173755114650674de07ed895a77469662f88a3adf919081900360200190a250565b6113eb6115d7565b5050565b600054600160a060020a0316331461140657600080fd5b61140f816117b8565b50565b60008061141e836114c9565b60011495945050505050565b82600160a060020a031663a9059cbb83836040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561148d57600080fd5b505af11580156114a1573d6000803e3d6000fd5b505050506040513d60208110156114b757600080fd5b505115156114c457600080fd5b505050565b60008080808080620151808704965062023ab162018e90600489020104600f019250600483048784010362254381019150611c896109891960148402010495506004860461016d870283030390506177896103e882020494506103e8610259860204601e86028203039350600d851161154e5761126c8603955060018503945061155c565b61126b86039550600d850394505b5050509193909250565b6001016000600c82111561157d5760019283019291505b61158983836001611835565b90505b92915050565b60008060006115a0846114c9565b5091509150836115b0838361188f565b61070801111580156115cf57506107086115ca8383611566565b038411155b949350505050565b336000908152600560205260408120549080821161163f576040805160e560020a62461bcd02815260206004820152601560248201527f6d7573742062652061207061727469636970616e740000000000000000000000604482015290519081900360640190fd5b5060065433600090815260056020526040812055611663818363ffffffff61189d16565b600655600254611682906101009004600160a060020a0316338461142a565b60408051838152905133917fd34c6d0846c2cbde96c7163ff2cc2c1e0ae463dddf965d55d7fa9b21ae99c8f6919081900360200190a29091565b60008215156116cd5750600061158c565b508181028183828115156116dd57fe5b041461158c57fe5b600081838115156116f257fe5b049392505050565b604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a0385811660048301528481166024830152604482018490529151918616916323b872dd916064808201926020929091908290030181600087803b15801561176e57600080fd5b505af1158015611782573d6000803e3d6000fd5b505050506040513d602081101561179857600080fd5b505115156117a557600080fd5b50505050565b8181018281101561158c57fe5b600160a060020a03811615156117cd57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60006002831161184c57600c830192506001840393505b6101908404606485046004860461016d87020103019050816005600185016003020484601e02010181019050620afac98103905062015180810290509392505050565b600061158983836002611835565b6000828211156118a957fe5b509003905600a165627a7a72305820a10ac7aedb1ebc331d3c74efd4d5aeb3471dac3f456f655e52d4f4b226ca9fb30029`

// DeployMonethaTokenHoldersProgramContract deploys a new Ethereum contract, binding an instance of MonethaTokenHoldersProgramContract to it.
func DeployMonethaTokenHoldersProgramContract(auth *bind.TransactOpts, backend bind.ContractBackend, _mthToken common.Address, _monethaVoucher common.Address) (common.Address, *types.Transaction, *MonethaTokenHoldersProgramContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MonethaTokenHoldersProgramContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MonethaTokenHoldersProgramContractBin), backend, _mthToken, _monethaVoucher)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MonethaTokenHoldersProgramContract{MonethaTokenHoldersProgramContractCaller: MonethaTokenHoldersProgramContractCaller{contract: contract}, MonethaTokenHoldersProgramContractTransactor: MonethaTokenHoldersProgramContractTransactor{contract: contract}, MonethaTokenHoldersProgramContractFilterer: MonethaTokenHoldersProgramContractFilterer{contract: contract}}, nil
}

// MonethaTokenHoldersProgramContract is an auto generated Go binding around an Ethereum contract.
type MonethaTokenHoldersProgramContract struct {
	MonethaTokenHoldersProgramContractCaller     // Read-only binding to the contract
	MonethaTokenHoldersProgramContractTransactor // Write-only binding to the contract
	MonethaTokenHoldersProgramContractFilterer   // Log filterer for contract events
}

// MonethaTokenHoldersProgramContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MonethaTokenHoldersProgramContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonethaTokenHoldersProgramContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MonethaTokenHoldersProgramContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonethaTokenHoldersProgramContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MonethaTokenHoldersProgramContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonethaTokenHoldersProgramContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MonethaTokenHoldersProgramContractSession struct {
	Contract     *MonethaTokenHoldersProgramContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                       // Call options to use throughout this session
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// MonethaTokenHoldersProgramContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MonethaTokenHoldersProgramContractCallerSession struct {
	Contract *MonethaTokenHoldersProgramContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                             // Call options to use throughout this session
}

// MonethaTokenHoldersProgramContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MonethaTokenHoldersProgramContractTransactorSession struct {
	Contract     *MonethaTokenHoldersProgramContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                             // Transaction auth options to use throughout this session
}

// MonethaTokenHoldersProgramContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MonethaTokenHoldersProgramContractRaw struct {
	Contract *MonethaTokenHoldersProgramContract // Generic contract binding to access the raw methods on
}

// MonethaTokenHoldersProgramContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MonethaTokenHoldersProgramContractCallerRaw struct {
	Contract *MonethaTokenHoldersProgramContractCaller // Generic read-only contract binding to access the raw methods on
}

// MonethaTokenHoldersProgramContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MonethaTokenHoldersProgramContractTransactorRaw struct {
	Contract *MonethaTokenHoldersProgramContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMonethaTokenHoldersProgramContract creates a new instance of MonethaTokenHoldersProgramContract, bound to a specific deployed contract.
func NewMonethaTokenHoldersProgramContract(address common.Address, backend bind.ContractBackend) (*MonethaTokenHoldersProgramContract, error) {
	contract, err := bindMonethaTokenHoldersProgramContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MonethaTokenHoldersProgramContract{MonethaTokenHoldersProgramContractCaller: MonethaTokenHoldersProgramContractCaller{contract: contract}, MonethaTokenHoldersProgramContractTransactor: MonethaTokenHoldersProgramContractTransactor{contract: contract}, MonethaTokenHoldersProgramContractFilterer: MonethaTokenHoldersProgramContractFilterer{contract: contract}}, nil
}

// NewMonethaTokenHoldersProgramContractCaller creates a new read-only instance of MonethaTokenHoldersProgramContract, bound to a specific deployed contract.
func NewMonethaTokenHoldersProgramContractCaller(address common.Address, caller bind.ContractCaller) (*MonethaTokenHoldersProgramContractCaller, error) {
	contract, err := bindMonethaTokenHoldersProgramContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MonethaTokenHoldersProgramContractCaller{contract: contract}, nil
}

// NewMonethaTokenHoldersProgramContractTransactor creates a new write-only instance of MonethaTokenHoldersProgramContract, bound to a specific deployed contract.
func NewMonethaTokenHoldersProgramContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MonethaTokenHoldersProgramContractTransactor, error) {
	contract, err := bindMonethaTokenHoldersProgramContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MonethaTokenHoldersProgramContractTransactor{contract: contract}, nil
}

// NewMonethaTokenHoldersProgramContractFilterer creates a new log filterer instance of MonethaTokenHoldersProgramContract, bound to a specific deployed contract.
func NewMonethaTokenHoldersProgramContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MonethaTokenHoldersProgramContractFilterer, error) {
	contract, err := bindMonethaTokenHoldersProgramContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MonethaTokenHoldersProgramContractFilterer{contract: contract}, nil
}

// bindMonethaTokenHoldersProgramContract binds a generic wrapper to an already deployed contract.
func bindMonethaTokenHoldersProgramContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MonethaTokenHoldersProgramContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MonethaTokenHoldersProgramContract.Contract.MonethaTokenHoldersProgramContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.MonethaTokenHoldersProgramContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.MonethaTokenHoldersProgramContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MonethaTokenHoldersProgramContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.contract.Transact(opts, method, params...)
}

// IsAllowedToParticipateNow is a free data retrieval call binding the contract method 0x17623c41.
//
// Solidity: function isAllowedToParticipateNow() constant returns(bool)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractCaller) IsAllowedToParticipateNow(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MonethaTokenHoldersProgramContract.contract.Call(opts, out, "isAllowedToParticipateNow")
	return *ret0, err
}

// IsAllowedToParticipateNow is a free data retrieval call binding the contract method 0x17623c41.
//
// Solidity: function isAllowedToParticipateNow() constant returns(bool)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) IsAllowedToParticipateNow() (bool, error) {
	return _MonethaTokenHoldersProgramContract.Contract.IsAllowedToParticipateNow(&_MonethaTokenHoldersProgramContract.CallOpts)
}

// IsAllowedToParticipateNow is a free data retrieval call binding the contract method 0x17623c41.
//
// Solidity: function isAllowedToParticipateNow() constant returns(bool)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractCallerSession) IsAllowedToParticipateNow() (bool, error) {
	return _MonethaTokenHoldersProgramContract.Contract.IsAllowedToParticipateNow(&_MonethaTokenHoldersProgramContract.CallOpts)
}

// IsAllowedToRedeemNow is a free data retrieval call binding the contract method 0x899bf897.
//
// Solidity: function isAllowedToRedeemNow() constant returns(bool)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractCaller) IsAllowedToRedeemNow(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MonethaTokenHoldersProgramContract.contract.Call(opts, out, "isAllowedToRedeemNow")
	return *ret0, err
}

// IsAllowedToRedeemNow is a free data retrieval call binding the contract method 0x899bf897.
//
// Solidity: function isAllowedToRedeemNow() constant returns(bool)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) IsAllowedToRedeemNow() (bool, error) {
	return _MonethaTokenHoldersProgramContract.Contract.IsAllowedToRedeemNow(&_MonethaTokenHoldersProgramContract.CallOpts)
}

// IsAllowedToRedeemNow is a free data retrieval call binding the contract method 0x899bf897.
//
// Solidity: function isAllowedToRedeemNow() constant returns(bool)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractCallerSession) IsAllowedToRedeemNow() (bool, error) {
	return _MonethaTokenHoldersProgramContract.Contract.IsAllowedToRedeemNow(&_MonethaTokenHoldersProgramContract.CallOpts)
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractCaller) IsMonethaAddress(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MonethaTokenHoldersProgramContract.contract.Call(opts, out, "isMonethaAddress", arg0)
	return *ret0, err
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) IsMonethaAddress(arg0 common.Address) (bool, error) {
	return _MonethaTokenHoldersProgramContract.Contract.IsMonethaAddress(&_MonethaTokenHoldersProgramContract.CallOpts, arg0)
}

// IsMonethaAddress is a free data retrieval call binding the contract method 0x31d41325.
//
// Solidity: function isMonethaAddress( address) constant returns(bool)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractCallerSession) IsMonethaAddress(arg0 common.Address) (bool, error) {
	return _MonethaTokenHoldersProgramContract.Contract.IsMonethaAddress(&_MonethaTokenHoldersProgramContract.CallOpts, arg0)
}

// MonethaVoucher is a free data retrieval call binding the contract method 0xa7abdf03.
//
// Solidity: function monethaVoucher() constant returns(address)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractCaller) MonethaVoucher(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonethaTokenHoldersProgramContract.contract.Call(opts, out, "monethaVoucher")
	return *ret0, err
}

// MonethaVoucher is a free data retrieval call binding the contract method 0xa7abdf03.
//
// Solidity: function monethaVoucher() constant returns(address)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) MonethaVoucher() (common.Address, error) {
	return _MonethaTokenHoldersProgramContract.Contract.MonethaVoucher(&_MonethaTokenHoldersProgramContract.CallOpts)
}

// MonethaVoucher is a free data retrieval call binding the contract method 0xa7abdf03.
//
// Solidity: function monethaVoucher() constant returns(address)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractCallerSession) MonethaVoucher() (common.Address, error) {
	return _MonethaTokenHoldersProgramContract.Contract.MonethaVoucher(&_MonethaTokenHoldersProgramContract.CallOpts)
}

// MthToken is a free data retrieval call binding the contract method 0x93306c43.
//
// Solidity: function mthToken() constant returns(address)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractCaller) MthToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonethaTokenHoldersProgramContract.contract.Call(opts, out, "mthToken")
	return *ret0, err
}

// MthToken is a free data retrieval call binding the contract method 0x93306c43.
//
// Solidity: function mthToken() constant returns(address)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) MthToken() (common.Address, error) {
	return _MonethaTokenHoldersProgramContract.Contract.MthToken(&_MonethaTokenHoldersProgramContract.CallOpts)
}

// MthToken is a free data retrieval call binding the contract method 0x93306c43.
//
// Solidity: function mthToken() constant returns(address)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractCallerSession) MthToken() (common.Address, error) {
	return _MonethaTokenHoldersProgramContract.Contract.MthToken(&_MonethaTokenHoldersProgramContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonethaTokenHoldersProgramContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) Owner() (common.Address, error) {
	return _MonethaTokenHoldersProgramContract.Contract.Owner(&_MonethaTokenHoldersProgramContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractCallerSession) Owner() (common.Address, error) {
	return _MonethaTokenHoldersProgramContract.Contract.Owner(&_MonethaTokenHoldersProgramContract.CallOpts)
}

// ParticipateFromTimestamp is a free data retrieval call binding the contract method 0xb1860997.
//
// Solidity: function participateFromTimestamp() constant returns(uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractCaller) ParticipateFromTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaTokenHoldersProgramContract.contract.Call(opts, out, "participateFromTimestamp")
	return *ret0, err
}

// ParticipateFromTimestamp is a free data retrieval call binding the contract method 0xb1860997.
//
// Solidity: function participateFromTimestamp() constant returns(uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) ParticipateFromTimestamp() (*big.Int, error) {
	return _MonethaTokenHoldersProgramContract.Contract.ParticipateFromTimestamp(&_MonethaTokenHoldersProgramContract.CallOpts)
}

// ParticipateFromTimestamp is a free data retrieval call binding the contract method 0xb1860997.
//
// Solidity: function participateFromTimestamp() constant returns(uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractCallerSession) ParticipateFromTimestamp() (*big.Int, error) {
	return _MonethaTokenHoldersProgramContract.Contract.ParticipateFromTimestamp(&_MonethaTokenHoldersProgramContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MonethaTokenHoldersProgramContract.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) Paused() (bool, error) {
	return _MonethaTokenHoldersProgramContract.Contract.Paused(&_MonethaTokenHoldersProgramContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractCallerSession) Paused() (bool, error) {
	return _MonethaTokenHoldersProgramContract.Contract.Paused(&_MonethaTokenHoldersProgramContract.CallOpts)
}

// StakedBy is a free data retrieval call binding the contract method 0x5f56a31a.
//
// Solidity: function stakedBy( address) constant returns(uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractCaller) StakedBy(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaTokenHoldersProgramContract.contract.Call(opts, out, "stakedBy", arg0)
	return *ret0, err
}

// StakedBy is a free data retrieval call binding the contract method 0x5f56a31a.
//
// Solidity: function stakedBy( address) constant returns(uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) StakedBy(arg0 common.Address) (*big.Int, error) {
	return _MonethaTokenHoldersProgramContract.Contract.StakedBy(&_MonethaTokenHoldersProgramContract.CallOpts, arg0)
}

// StakedBy is a free data retrieval call binding the contract method 0x5f56a31a.
//
// Solidity: function stakedBy( address) constant returns(uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractCallerSession) StakedBy(arg0 common.Address) (*big.Int, error) {
	return _MonethaTokenHoldersProgramContract.Contract.StakedBy(&_MonethaTokenHoldersProgramContract.CallOpts, arg0)
}

// TotalStacked is a free data retrieval call binding the contract method 0x29eae70d.
//
// Solidity: function totalStacked() constant returns(uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractCaller) TotalStacked(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonethaTokenHoldersProgramContract.contract.Call(opts, out, "totalStacked")
	return *ret0, err
}

// TotalStacked is a free data retrieval call binding the contract method 0x29eae70d.
//
// Solidity: function totalStacked() constant returns(uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) TotalStacked() (*big.Int, error) {
	return _MonethaTokenHoldersProgramContract.Contract.TotalStacked(&_MonethaTokenHoldersProgramContract.CallOpts)
}

// TotalStacked is a free data retrieval call binding the contract method 0x29eae70d.
//
// Solidity: function totalStacked() constant returns(uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractCallerSession) TotalStacked() (*big.Int, error) {
	return _MonethaTokenHoldersProgramContract.Contract.TotalStacked(&_MonethaTokenHoldersProgramContract.CallOpts)
}

// BuyVouchers is a paid mutator transaction binding the contract method 0x87585707.
//
// Solidity: function buyVouchers() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactor) BuyVouchers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.contract.Transact(opts, "buyVouchers")
}

// BuyVouchers is a paid mutator transaction binding the contract method 0x87585707.
//
// Solidity: function buyVouchers() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) BuyVouchers() (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.BuyVouchers(&_MonethaTokenHoldersProgramContract.TransactOpts)
}

// BuyVouchers is a paid mutator transaction binding the contract method 0x87585707.
//
// Solidity: function buyVouchers() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactorSession) BuyVouchers() (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.BuyVouchers(&_MonethaTokenHoldersProgramContract.TransactOpts)
}

// CancelParticipation is a paid mutator transaction binding the contract method 0xdae37abc.
//
// Solidity: function cancelParticipation() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactor) CancelParticipation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.contract.Transact(opts, "cancelParticipation")
}

// CancelParticipation is a paid mutator transaction binding the contract method 0xdae37abc.
//
// Solidity: function cancelParticipation() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) CancelParticipation() (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.CancelParticipation(&_MonethaTokenHoldersProgramContract.TransactOpts)
}

// CancelParticipation is a paid mutator transaction binding the contract method 0xdae37abc.
//
// Solidity: function cancelParticipation() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactorSession) CancelParticipation() (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.CancelParticipation(&_MonethaTokenHoldersProgramContract.TransactOpts)
}

// Participate is a paid mutator transaction binding the contract method 0xd11711a2.
//
// Solidity: function participate() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactor) Participate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.contract.Transact(opts, "participate")
}

// Participate is a paid mutator transaction binding the contract method 0xd11711a2.
//
// Solidity: function participate() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) Participate() (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.Participate(&_MonethaTokenHoldersProgramContract.TransactOpts)
}

// Participate is a paid mutator transaction binding the contract method 0xd11711a2.
//
// Solidity: function participate() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactorSession) Participate() (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.Participate(&_MonethaTokenHoldersProgramContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) Pause() (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.Pause(&_MonethaTokenHoldersProgramContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactorSession) Pause() (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.Pause(&_MonethaTokenHoldersProgramContract.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactor) ReclaimEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.contract.Transact(opts, "reclaimEther")
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) ReclaimEther() (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.ReclaimEther(&_MonethaTokenHoldersProgramContract.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactorSession) ReclaimEther() (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.ReclaimEther(&_MonethaTokenHoldersProgramContract.TransactOpts)
}

// ReclaimEtherTo is a paid mutator transaction binding the contract method 0x2e84a374.
//
// Solidity: function reclaimEtherTo(_to address, _value uint256) returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactor) ReclaimEtherTo(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.contract.Transact(opts, "reclaimEtherTo", _to, _value)
}

// ReclaimEtherTo is a paid mutator transaction binding the contract method 0x2e84a374.
//
// Solidity: function reclaimEtherTo(_to address, _value uint256) returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) ReclaimEtherTo(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.ReclaimEtherTo(&_MonethaTokenHoldersProgramContract.TransactOpts, _to, _value)
}

// ReclaimEtherTo is a paid mutator transaction binding the contract method 0x2e84a374.
//
// Solidity: function reclaimEtherTo(_to address, _value uint256) returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactorSession) ReclaimEtherTo(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.ReclaimEtherTo(&_MonethaTokenHoldersProgramContract.TransactOpts, _to, _value)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(_token address) returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactor) ReclaimToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.contract.Transact(opts, "reclaimToken", _token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(_token address) returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) ReclaimToken(_token common.Address) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.ReclaimToken(&_MonethaTokenHoldersProgramContract.TransactOpts, _token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(_token address) returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactorSession) ReclaimToken(_token common.Address) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.ReclaimToken(&_MonethaTokenHoldersProgramContract.TransactOpts, _token)
}

// ReclaimTokenTo is a paid mutator transaction binding the contract method 0x9fe98bda.
//
// Solidity: function reclaimTokenTo(_token address, _to address, _value uint256) returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactor) ReclaimTokenTo(opts *bind.TransactOpts, _token common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.contract.Transact(opts, "reclaimTokenTo", _token, _to, _value)
}

// ReclaimTokenTo is a paid mutator transaction binding the contract method 0x9fe98bda.
//
// Solidity: function reclaimTokenTo(_token address, _to address, _value uint256) returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) ReclaimTokenTo(_token common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.ReclaimTokenTo(&_MonethaTokenHoldersProgramContract.TransactOpts, _token, _to, _value)
}

// ReclaimTokenTo is a paid mutator transaction binding the contract method 0x9fe98bda.
//
// Solidity: function reclaimTokenTo(_token address, _to address, _value uint256) returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactorSession) ReclaimTokenTo(_token common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.ReclaimTokenTo(&_MonethaTokenHoldersProgramContract.TransactOpts, _token, _to, _value)
}

// Redeem is a paid mutator transaction binding the contract method 0xbe040fb0.
//
// Solidity: function redeem() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactor) Redeem(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.contract.Transact(opts, "redeem")
}

// Redeem is a paid mutator transaction binding the contract method 0xbe040fb0.
//
// Solidity: function redeem() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) Redeem() (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.Redeem(&_MonethaTokenHoldersProgramContract.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xbe040fb0.
//
// Solidity: function redeem() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactorSession) Redeem() (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.Redeem(&_MonethaTokenHoldersProgramContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.RenounceOwnership(&_MonethaTokenHoldersProgramContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.RenounceOwnership(&_MonethaTokenHoldersProgramContract.TransactOpts)
}

// SellVouchers is a paid mutator transaction binding the contract method 0x80bfc599.
//
// Solidity: function sellVouchers() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactor) SellVouchers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.contract.Transact(opts, "sellVouchers")
}

// SellVouchers is a paid mutator transaction binding the contract method 0x80bfc599.
//
// Solidity: function sellVouchers() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) SellVouchers() (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.SellVouchers(&_MonethaTokenHoldersProgramContract.TransactOpts)
}

// SellVouchers is a paid mutator transaction binding the contract method 0x80bfc599.
//
// Solidity: function sellVouchers() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactorSession) SellVouchers() (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.SellVouchers(&_MonethaTokenHoldersProgramContract.TransactOpts)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactor) SetMonethaAddress(opts *bind.TransactOpts, _address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.contract.Transact(opts, "setMonethaAddress", _address, _isMonethaAddress)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) SetMonethaAddress(_address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.SetMonethaAddress(&_MonethaTokenHoldersProgramContract.TransactOpts, _address, _isMonethaAddress)
}

// SetMonethaAddress is a paid mutator transaction binding the contract method 0xc07e3391.
//
// Solidity: function setMonethaAddress(_address address, _isMonethaAddress bool) returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactorSession) SetMonethaAddress(_address common.Address, _isMonethaAddress bool) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.SetMonethaAddress(&_MonethaTokenHoldersProgramContract.TransactOpts, _address, _isMonethaAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.TransferOwnership(&_MonethaTokenHoldersProgramContract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.TransferOwnership(&_MonethaTokenHoldersProgramContract.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractSession) Unpause() (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.Unpause(&_MonethaTokenHoldersProgramContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _MonethaTokenHoldersProgramContract.Contract.Unpause(&_MonethaTokenHoldersProgramContract.TransactOpts)
}

// MonethaTokenHoldersProgramContractMonethaAddressSetIterator is returned from FilterMonethaAddressSet and is used to iterate over the raw logs and unpacked data for MonethaAddressSet events raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractMonethaAddressSetIterator struct {
	Event *MonethaTokenHoldersProgramContractMonethaAddressSet // Event containing the contract specifics and raw log

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
func (it *MonethaTokenHoldersProgramContractMonethaAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaTokenHoldersProgramContractMonethaAddressSet)
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
		it.Event = new(MonethaTokenHoldersProgramContractMonethaAddressSet)
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
func (it *MonethaTokenHoldersProgramContractMonethaAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaTokenHoldersProgramContractMonethaAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaTokenHoldersProgramContractMonethaAddressSet represents a MonethaAddressSet event raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractMonethaAddressSet struct {
	Address          common.Address
	IsMonethaAddress bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMonethaAddressSet is a free log retrieval operation binding the contract event 0xa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e692.
//
// Solidity: e MonethaAddressSet(_address address, _isMonethaAddress bool)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) FilterMonethaAddressSet(opts *bind.FilterOpts) (*MonethaTokenHoldersProgramContractMonethaAddressSetIterator, error) {

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.FilterLogs(opts, "MonethaAddressSet")
	if err != nil {
		return nil, err
	}
	return &MonethaTokenHoldersProgramContractMonethaAddressSetIterator{contract: _MonethaTokenHoldersProgramContract.contract, event: "MonethaAddressSet", logs: logs, sub: sub}, nil
}

// WatchMonethaAddressSet is a free log subscription operation binding the contract event 0xa551de8741dbb2092ce6bc142fd0ff3af5dfbf87d0aa594619fccddb0141e692.
//
// Solidity: e MonethaAddressSet(_address address, _isMonethaAddress bool)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) WatchMonethaAddressSet(opts *bind.WatchOpts, sink chan<- *MonethaTokenHoldersProgramContractMonethaAddressSet) (event.Subscription, error) {

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.WatchLogs(opts, "MonethaAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaTokenHoldersProgramContractMonethaAddressSet)
				if err := _MonethaTokenHoldersProgramContract.contract.UnpackLog(event, "MonethaAddressSet", log); err != nil {
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

// MonethaTokenHoldersProgramContractOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractOwnershipRenouncedIterator struct {
	Event *MonethaTokenHoldersProgramContractOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *MonethaTokenHoldersProgramContractOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaTokenHoldersProgramContractOwnershipRenounced)
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
		it.Event = new(MonethaTokenHoldersProgramContractOwnershipRenounced)
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
func (it *MonethaTokenHoldersProgramContractOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaTokenHoldersProgramContractOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaTokenHoldersProgramContractOwnershipRenounced represents a OwnershipRenounced event raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*MonethaTokenHoldersProgramContractOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MonethaTokenHoldersProgramContractOwnershipRenouncedIterator{contract: _MonethaTokenHoldersProgramContract.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *MonethaTokenHoldersProgramContractOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaTokenHoldersProgramContractOwnershipRenounced)
				if err := _MonethaTokenHoldersProgramContract.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// MonethaTokenHoldersProgramContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractOwnershipTransferredIterator struct {
	Event *MonethaTokenHoldersProgramContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MonethaTokenHoldersProgramContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaTokenHoldersProgramContractOwnershipTransferred)
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
		it.Event = new(MonethaTokenHoldersProgramContractOwnershipTransferred)
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
func (it *MonethaTokenHoldersProgramContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaTokenHoldersProgramContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaTokenHoldersProgramContractOwnershipTransferred represents a OwnershipTransferred event raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MonethaTokenHoldersProgramContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MonethaTokenHoldersProgramContractOwnershipTransferredIterator{contract: _MonethaTokenHoldersProgramContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MonethaTokenHoldersProgramContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaTokenHoldersProgramContractOwnershipTransferred)
				if err := _MonethaTokenHoldersProgramContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// MonethaTokenHoldersProgramContractParticipationStartedIterator is returned from FilterParticipationStarted and is used to iterate over the raw logs and unpacked data for ParticipationStarted events raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractParticipationStartedIterator struct {
	Event *MonethaTokenHoldersProgramContractParticipationStarted // Event containing the contract specifics and raw log

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
func (it *MonethaTokenHoldersProgramContractParticipationStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaTokenHoldersProgramContractParticipationStarted)
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
		it.Event = new(MonethaTokenHoldersProgramContractParticipationStarted)
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
func (it *MonethaTokenHoldersProgramContractParticipationStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaTokenHoldersProgramContractParticipationStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaTokenHoldersProgramContractParticipationStarted represents a ParticipationStarted event raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractParticipationStarted struct {
	Participant common.Address
	MthTokens   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterParticipationStarted is a free log retrieval operation binding the contract event 0xbd6451f729baeebe302c84f173755114650674de07ed895a77469662f88a3adf.
//
// Solidity: e ParticipationStarted(participant indexed address, mthTokens uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) FilterParticipationStarted(opts *bind.FilterOpts, participant []common.Address) (*MonethaTokenHoldersProgramContractParticipationStartedIterator, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.FilterLogs(opts, "ParticipationStarted", participantRule)
	if err != nil {
		return nil, err
	}
	return &MonethaTokenHoldersProgramContractParticipationStartedIterator{contract: _MonethaTokenHoldersProgramContract.contract, event: "ParticipationStarted", logs: logs, sub: sub}, nil
}

// WatchParticipationStarted is a free log subscription operation binding the contract event 0xbd6451f729baeebe302c84f173755114650674de07ed895a77469662f88a3adf.
//
// Solidity: e ParticipationStarted(participant indexed address, mthTokens uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) WatchParticipationStarted(opts *bind.WatchOpts, sink chan<- *MonethaTokenHoldersProgramContractParticipationStarted, participant []common.Address) (event.Subscription, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.WatchLogs(opts, "ParticipationStarted", participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaTokenHoldersProgramContractParticipationStarted)
				if err := _MonethaTokenHoldersProgramContract.contract.UnpackLog(event, "ParticipationStarted", log); err != nil {
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

// MonethaTokenHoldersProgramContractParticipationStoppedIterator is returned from FilterParticipationStopped and is used to iterate over the raw logs and unpacked data for ParticipationStopped events raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractParticipationStoppedIterator struct {
	Event *MonethaTokenHoldersProgramContractParticipationStopped // Event containing the contract specifics and raw log

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
func (it *MonethaTokenHoldersProgramContractParticipationStoppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaTokenHoldersProgramContractParticipationStopped)
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
		it.Event = new(MonethaTokenHoldersProgramContractParticipationStopped)
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
func (it *MonethaTokenHoldersProgramContractParticipationStoppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaTokenHoldersProgramContractParticipationStoppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaTokenHoldersProgramContractParticipationStopped represents a ParticipationStopped event raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractParticipationStopped struct {
	Participant common.Address
	MthTokens   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterParticipationStopped is a free log retrieval operation binding the contract event 0xd34c6d0846c2cbde96c7163ff2cc2c1e0ae463dddf965d55d7fa9b21ae99c8f6.
//
// Solidity: e ParticipationStopped(participant indexed address, mthTokens uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) FilterParticipationStopped(opts *bind.FilterOpts, participant []common.Address) (*MonethaTokenHoldersProgramContractParticipationStoppedIterator, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.FilterLogs(opts, "ParticipationStopped", participantRule)
	if err != nil {
		return nil, err
	}
	return &MonethaTokenHoldersProgramContractParticipationStoppedIterator{contract: _MonethaTokenHoldersProgramContract.contract, event: "ParticipationStopped", logs: logs, sub: sub}, nil
}

// WatchParticipationStopped is a free log subscription operation binding the contract event 0xd34c6d0846c2cbde96c7163ff2cc2c1e0ae463dddf965d55d7fa9b21ae99c8f6.
//
// Solidity: e ParticipationStopped(participant indexed address, mthTokens uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) WatchParticipationStopped(opts *bind.WatchOpts, sink chan<- *MonethaTokenHoldersProgramContractParticipationStopped, participant []common.Address) (event.Subscription, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.WatchLogs(opts, "ParticipationStopped", participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaTokenHoldersProgramContractParticipationStopped)
				if err := _MonethaTokenHoldersProgramContract.contract.UnpackLog(event, "ParticipationStopped", log); err != nil {
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

// MonethaTokenHoldersProgramContractPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractPauseIterator struct {
	Event *MonethaTokenHoldersProgramContractPause // Event containing the contract specifics and raw log

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
func (it *MonethaTokenHoldersProgramContractPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaTokenHoldersProgramContractPause)
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
		it.Event = new(MonethaTokenHoldersProgramContractPause)
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
func (it *MonethaTokenHoldersProgramContractPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaTokenHoldersProgramContractPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaTokenHoldersProgramContractPause represents a Pause event raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) FilterPause(opts *bind.FilterOpts) (*MonethaTokenHoldersProgramContractPauseIterator, error) {

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &MonethaTokenHoldersProgramContractPauseIterator{contract: _MonethaTokenHoldersProgramContract.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *MonethaTokenHoldersProgramContractPause) (event.Subscription, error) {

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaTokenHoldersProgramContractPause)
				if err := _MonethaTokenHoldersProgramContract.contract.UnpackLog(event, "Pause", log); err != nil {
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

// MonethaTokenHoldersProgramContractReclaimEtherIterator is returned from FilterReclaimEther and is used to iterate over the raw logs and unpacked data for ReclaimEther events raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractReclaimEtherIterator struct {
	Event *MonethaTokenHoldersProgramContractReclaimEther // Event containing the contract specifics and raw log

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
func (it *MonethaTokenHoldersProgramContractReclaimEtherIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaTokenHoldersProgramContractReclaimEther)
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
		it.Event = new(MonethaTokenHoldersProgramContractReclaimEther)
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
func (it *MonethaTokenHoldersProgramContractReclaimEtherIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaTokenHoldersProgramContractReclaimEtherIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaTokenHoldersProgramContractReclaimEther represents a ReclaimEther event raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractReclaimEther struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReclaimEther is a free log retrieval operation binding the contract event 0xb54913b2b58b2e96ea9b4e96ba2353cf13426af9d3f252e0c17899a93c4ce125.
//
// Solidity: e ReclaimEther(to indexed address, amount uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) FilterReclaimEther(opts *bind.FilterOpts, to []common.Address) (*MonethaTokenHoldersProgramContractReclaimEtherIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.FilterLogs(opts, "ReclaimEther", toRule)
	if err != nil {
		return nil, err
	}
	return &MonethaTokenHoldersProgramContractReclaimEtherIterator{contract: _MonethaTokenHoldersProgramContract.contract, event: "ReclaimEther", logs: logs, sub: sub}, nil
}

// WatchReclaimEther is a free log subscription operation binding the contract event 0xb54913b2b58b2e96ea9b4e96ba2353cf13426af9d3f252e0c17899a93c4ce125.
//
// Solidity: e ReclaimEther(to indexed address, amount uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) WatchReclaimEther(opts *bind.WatchOpts, sink chan<- *MonethaTokenHoldersProgramContractReclaimEther, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.WatchLogs(opts, "ReclaimEther", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaTokenHoldersProgramContractReclaimEther)
				if err := _MonethaTokenHoldersProgramContract.contract.UnpackLog(event, "ReclaimEther", log); err != nil {
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

// MonethaTokenHoldersProgramContractReclaimTokensIterator is returned from FilterReclaimTokens and is used to iterate over the raw logs and unpacked data for ReclaimTokens events raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractReclaimTokensIterator struct {
	Event *MonethaTokenHoldersProgramContractReclaimTokens // Event containing the contract specifics and raw log

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
func (it *MonethaTokenHoldersProgramContractReclaimTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaTokenHoldersProgramContractReclaimTokens)
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
		it.Event = new(MonethaTokenHoldersProgramContractReclaimTokens)
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
func (it *MonethaTokenHoldersProgramContractReclaimTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaTokenHoldersProgramContractReclaimTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaTokenHoldersProgramContractReclaimTokens represents a ReclaimTokens event raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractReclaimTokens struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReclaimTokens is a free log retrieval operation binding the contract event 0x355069f20974db323c9dcd100e8bf13fb2acc1884e5ec05b0a89c09e15ce810f.
//
// Solidity: e ReclaimTokens(to indexed address, amount uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) FilterReclaimTokens(opts *bind.FilterOpts, to []common.Address) (*MonethaTokenHoldersProgramContractReclaimTokensIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.FilterLogs(opts, "ReclaimTokens", toRule)
	if err != nil {
		return nil, err
	}
	return &MonethaTokenHoldersProgramContractReclaimTokensIterator{contract: _MonethaTokenHoldersProgramContract.contract, event: "ReclaimTokens", logs: logs, sub: sub}, nil
}

// WatchReclaimTokens is a free log subscription operation binding the contract event 0x355069f20974db323c9dcd100e8bf13fb2acc1884e5ec05b0a89c09e15ce810f.
//
// Solidity: e ReclaimTokens(to indexed address, amount uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) WatchReclaimTokens(opts *bind.WatchOpts, sink chan<- *MonethaTokenHoldersProgramContractReclaimTokens, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.WatchLogs(opts, "ReclaimTokens", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaTokenHoldersProgramContractReclaimTokens)
				if err := _MonethaTokenHoldersProgramContract.contract.UnpackLog(event, "ReclaimTokens", log); err != nil {
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

// MonethaTokenHoldersProgramContractUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractUnpauseIterator struct {
	Event *MonethaTokenHoldersProgramContractUnpause // Event containing the contract specifics and raw log

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
func (it *MonethaTokenHoldersProgramContractUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaTokenHoldersProgramContractUnpause)
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
		it.Event = new(MonethaTokenHoldersProgramContractUnpause)
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
func (it *MonethaTokenHoldersProgramContractUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaTokenHoldersProgramContractUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaTokenHoldersProgramContractUnpause represents a Unpause event raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) FilterUnpause(opts *bind.FilterOpts) (*MonethaTokenHoldersProgramContractUnpauseIterator, error) {

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &MonethaTokenHoldersProgramContractUnpauseIterator{contract: _MonethaTokenHoldersProgramContract.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *MonethaTokenHoldersProgramContractUnpause) (event.Subscription, error) {

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaTokenHoldersProgramContractUnpause)
				if err := _MonethaTokenHoldersProgramContract.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// MonethaTokenHoldersProgramContractVouchersPurchasedIterator is returned from FilterVouchersPurchased and is used to iterate over the raw logs and unpacked data for VouchersPurchased events raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractVouchersPurchasedIterator struct {
	Event *MonethaTokenHoldersProgramContractVouchersPurchased // Event containing the contract specifics and raw log

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
func (it *MonethaTokenHoldersProgramContractVouchersPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaTokenHoldersProgramContractVouchersPurchased)
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
		it.Event = new(MonethaTokenHoldersProgramContractVouchersPurchased)
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
func (it *MonethaTokenHoldersProgramContractVouchersPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaTokenHoldersProgramContractVouchersPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaTokenHoldersProgramContractVouchersPurchased represents a VouchersPurchased event raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractVouchersPurchased struct {
	Vouchers *big.Int
	Weis     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVouchersPurchased is a free log retrieval operation binding the contract event 0x9fca5352dcf7d3de26b4fcc50718ab421678e3c4cb800a589bdc49fbdee51a79.
//
// Solidity: e VouchersPurchased(vouchers uint256, weis uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) FilterVouchersPurchased(opts *bind.FilterOpts) (*MonethaTokenHoldersProgramContractVouchersPurchasedIterator, error) {

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.FilterLogs(opts, "VouchersPurchased")
	if err != nil {
		return nil, err
	}
	return &MonethaTokenHoldersProgramContractVouchersPurchasedIterator{contract: _MonethaTokenHoldersProgramContract.contract, event: "VouchersPurchased", logs: logs, sub: sub}, nil
}

// WatchVouchersPurchased is a free log subscription operation binding the contract event 0x9fca5352dcf7d3de26b4fcc50718ab421678e3c4cb800a589bdc49fbdee51a79.
//
// Solidity: e VouchersPurchased(vouchers uint256, weis uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) WatchVouchersPurchased(opts *bind.WatchOpts, sink chan<- *MonethaTokenHoldersProgramContractVouchersPurchased) (event.Subscription, error) {

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.WatchLogs(opts, "VouchersPurchased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaTokenHoldersProgramContractVouchersPurchased)
				if err := _MonethaTokenHoldersProgramContract.contract.UnpackLog(event, "VouchersPurchased", log); err != nil {
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

// MonethaTokenHoldersProgramContractVouchersRedeemedIterator is returned from FilterVouchersRedeemed and is used to iterate over the raw logs and unpacked data for VouchersRedeemed events raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractVouchersRedeemedIterator struct {
	Event *MonethaTokenHoldersProgramContractVouchersRedeemed // Event containing the contract specifics and raw log

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
func (it *MonethaTokenHoldersProgramContractVouchersRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaTokenHoldersProgramContractVouchersRedeemed)
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
		it.Event = new(MonethaTokenHoldersProgramContractVouchersRedeemed)
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
func (it *MonethaTokenHoldersProgramContractVouchersRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaTokenHoldersProgramContractVouchersRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaTokenHoldersProgramContractVouchersRedeemed represents a VouchersRedeemed event raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractVouchersRedeemed struct {
	Participant common.Address
	Vouchers    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVouchersRedeemed is a free log retrieval operation binding the contract event 0xb7dd3fa0284dca52a2a1e7985ef6eb7d2bd8a4fb3388f8baf55a64fd5ff4c94a.
//
// Solidity: e VouchersRedeemed(participant indexed address, vouchers uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) FilterVouchersRedeemed(opts *bind.FilterOpts, participant []common.Address) (*MonethaTokenHoldersProgramContractVouchersRedeemedIterator, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.FilterLogs(opts, "VouchersRedeemed", participantRule)
	if err != nil {
		return nil, err
	}
	return &MonethaTokenHoldersProgramContractVouchersRedeemedIterator{contract: _MonethaTokenHoldersProgramContract.contract, event: "VouchersRedeemed", logs: logs, sub: sub}, nil
}

// WatchVouchersRedeemed is a free log subscription operation binding the contract event 0xb7dd3fa0284dca52a2a1e7985ef6eb7d2bd8a4fb3388f8baf55a64fd5ff4c94a.
//
// Solidity: e VouchersRedeemed(participant indexed address, vouchers uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) WatchVouchersRedeemed(opts *bind.WatchOpts, sink chan<- *MonethaTokenHoldersProgramContractVouchersRedeemed, participant []common.Address) (event.Subscription, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.WatchLogs(opts, "VouchersRedeemed", participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaTokenHoldersProgramContractVouchersRedeemed)
				if err := _MonethaTokenHoldersProgramContract.contract.UnpackLog(event, "VouchersRedeemed", log); err != nil {
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

// MonethaTokenHoldersProgramContractVouchersSoldIterator is returned from FilterVouchersSold and is used to iterate over the raw logs and unpacked data for VouchersSold events raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractVouchersSoldIterator struct {
	Event *MonethaTokenHoldersProgramContractVouchersSold // Event containing the contract specifics and raw log

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
func (it *MonethaTokenHoldersProgramContractVouchersSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonethaTokenHoldersProgramContractVouchersSold)
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
		it.Event = new(MonethaTokenHoldersProgramContractVouchersSold)
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
func (it *MonethaTokenHoldersProgramContractVouchersSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonethaTokenHoldersProgramContractVouchersSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonethaTokenHoldersProgramContractVouchersSold represents a VouchersSold event raised by the MonethaTokenHoldersProgramContract contract.
type MonethaTokenHoldersProgramContractVouchersSold struct {
	Vouchers *big.Int
	Weis     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVouchersSold is a free log retrieval operation binding the contract event 0x707c653f7d61b8660b51074cc8c2d9e99f9c34436ea908d62bea7a5a4cb18d4f.
//
// Solidity: e VouchersSold(vouchers uint256, weis uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) FilterVouchersSold(opts *bind.FilterOpts) (*MonethaTokenHoldersProgramContractVouchersSoldIterator, error) {

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.FilterLogs(opts, "VouchersSold")
	if err != nil {
		return nil, err
	}
	return &MonethaTokenHoldersProgramContractVouchersSoldIterator{contract: _MonethaTokenHoldersProgramContract.contract, event: "VouchersSold", logs: logs, sub: sub}, nil
}

// WatchVouchersSold is a free log subscription operation binding the contract event 0x707c653f7d61b8660b51074cc8c2d9e99f9c34436ea908d62bea7a5a4cb18d4f.
//
// Solidity: e VouchersSold(vouchers uint256, weis uint256)
func (_MonethaTokenHoldersProgramContract *MonethaTokenHoldersProgramContractFilterer) WatchVouchersSold(opts *bind.WatchOpts, sink chan<- *MonethaTokenHoldersProgramContractVouchersSold) (event.Subscription, error) {

	logs, sub, err := _MonethaTokenHoldersProgramContract.contract.WatchLogs(opts, "VouchersSold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonethaTokenHoldersProgramContractVouchersSold)
				if err := _MonethaTokenHoldersProgramContract.contract.UnpackLog(event, "VouchersSold", log); err != nil {
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

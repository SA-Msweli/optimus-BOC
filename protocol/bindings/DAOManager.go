// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DAOManagerMetaData contains all meta data concerning the DAOManager contract.
var DAOManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DAO_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TREASURY_FUNDER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"goal\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"votingPeriodDays\",\"type\":\"uint64\"}],\"name\":\"createDAO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"daoId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"daoId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"creditTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"finalizeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"daoId\",\"type\":\"uint256\"}],\"name\":\"getBnplTerms\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"numInstallments\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allowedIntervalMinDays\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allowedIntervalMaxDays\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lateFeeBps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gracePeriodDays\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"rescheduleAllowed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"minDownPaymentBps\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"daoId\",\"type\":\"uint256\"}],\"name\":\"getTreasuryBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"daoId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"investmentAmount\",\"type\":\"uint256\"}],\"name\":\"joinDAO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"daoId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"proposalData\",\"type\":\"bytes\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"daoId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numInstallments\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allowedIntervalMinDays\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allowedIntervalMaxDays\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lateFeeBps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gracePeriodDays\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"rescheduleAllowed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"minDownPaymentBps\",\"type\":\"uint256\"}],\"name\":\"setBnplTerms\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenVault\",\"type\":\"address\"}],\"name\":\"setTokenVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"daoId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numInstallments\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedIntervalMinDays\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedIntervalMaxDays\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lateFeeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gracePeriodDays\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"rescheduleAllowed\",\"type\":\"bool\"}],\"name\":\"BnplTermsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"daoId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"goal\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"name\":\"DaoCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"daoId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"investment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"joinedAt\",\"type\":\"uint256\"}],\"name\":\"MemberJoined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executedAt\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalizedAt\",\"type\":\"uint256\"}],\"name\":\"ProposalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"daoId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ProposalOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"daoId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"TreasuryDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"daoId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TreasuryWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"}]",
	Bin: "0x608060405234604c5760017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055600180556001600255603c336050565b5060405161156f90816100d98239f35b5f80fd5b6001600160a01b0381165f9081525f5160206116485f395f51905f52602052604090205460ff1660d3576001600160a01b03165f8181525f5160206116485f395f51905f5260205260408120805460ff191660011790553391907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b505f9056fe6080806040526004361015610012575f80fd5b5f905f3560e01c90816301ffc9a71461128d575080630d61b51914610e3f578063184512cd14610e12578063248a9ca314610ddf5780632f2ff15d14610da157806336568abe14610d5c57806340e8928d14610bc95780635652077c14610af95780635bc789d914610ad05780635f90ebaf146108985780636497a8a0146107e857806391d148541461079f578063a217fddf14610783578063bc4b910114610635578063c9d27afe14610498578063cec4aec914610369578063d14dd2f21461032e578063d3c095a3146102bd578063d547741f14610276578063e3065c20146101425763f859e5dc14610105575f80fd5b3461013f578060031936011261013f5760206040517f41fe02d5181d01d5018a1014fd4a15ed3038aaddfe356993fb28797863b276d58152f35b80fd5b503461013f57606036600319011261013f5760043561015f6112e0565b90604435918184526003602052604084206003810154151580610267575b1561022257600281018054906001820180921161020e576040927fd1c23df765b8db6e8e7b2077d1e985dc2c7b79dcd08cc1a6dc48b566bddeea73949260019255016101ca8682546113b0565b9055838652600660205281862060018060a01b0382165f52602052815f206101f38682546113b0565b905581519485524260208601526001600160a01b031693a380f35b634e487b7160e01b87526011600452602487fd5b60405162461bcd60e51b815260206004820152601a60248201527f44414f5f4e4f545f464f554e445f4f525f444953534f4c5645440000000000006044820152606490fd5b5060ff6004820154161561017d565b503461013f57604036600319011261013f576102b96004356102966112e0565b906102b46102af825f525f602052600160405f20015490565b6113f9565b6114b9565b5080f35b503461013f57602036600319011261013f57604060e091600435815260046020522080549060018101549060028101546003820154600483015491600660ff600586015416940154946040519687526020870152604086015260608501526080840152151560a083015260c0820152f35b503461013f578060031936011261013f5760206040517f9729b62c1a280a968a683045fd5bbd60b97be4a6b7b0a5b624c9bbc54d6f3b598152f35b503461013f57604036600319011261013f5760043560243590808352600360205261039c600360408520015415156113bd565b7f41fe02d5181d01d5018a1014fd4a15ed3038aaddfe356993fb28797863b276d5835260208381526040808520335f908152925290205460ff168015610478575b15610442578083526003602052600160408420016103fc8382546113b0565b90558083526003602052600160408420015460405192835260208301527f79bdb62d9a421ad24bcd8ac4547a138852f10e01346b861c15122348a55d26a460403393a380f35b60405162461bcd60e51b815260206004820152600e60248201526d1393d517d055551213d49256915160921b6044820152606490fd5b5082805260208381526040808520335f908152925290205460ff166103dd565b503461013f57604036600319011261013f57602435600435811515820361063157808352600560205260408320600281015442111580610622575b156105ed57818452600760209081526040808620335f908152925290205460ff166105b85780548452600660209081526040808620335f90815292529020547fcbdf6214089cba887ecbf35a0b6a734589959c9763342c756bb2a80ca2bc9f6e9161059791806105b257506001905b848752600760209081526040808920335f90815292529020805460ff19166001179055851561059d576003016105798282546113b0565b90555b60408051951515865260208601919091523394918291820190565b0390a380f35b6004016105ab8282546113b0565b905561057c565b90610542565b60405162461bcd60e51b815260206004820152600d60248201526c1053149150511657d593d51151609a1b6044820152606490fd5b60405162461bcd60e51b815260206004820152600d60248201526c1593d5125391d7d0d313d4d151609a1b6044820152606490fd5b5060ff600582015416156104d3565b8280fd5b503461013f5761010036600319011261013f5760043560243560443560643560843560a4359160c4359384151580950361077f578688526003602052610683600360408a20015415156113bd565b60405160e0810181811067ffffffffffffffff82111761076b579185939160068a6107448a7f7e4912d5452f297c8c78ebd5196c5d177af96195a7ef704138d2031ebba23b309c9a8f60c09d9b996040528c87526020870188815260408801908a8252604060608a01938d855260808b0195865260a08b0196875260c08b019860e4358a52815260046020522098518955516001890155516002880155516003870155516004860155511515600585019060ff801983541691151516179055565b51910155604051958652602086015260408501526060840152608083015260a0820152a280f35b634e487b7160e01b8a52604160045260248afd5b8780fd5b503461013f578060031936011261013f57602090604051908152f35b503461013f57604036600319011261013f5760406107bb6112e0565b91600435815280602052209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b503461013f57602036600319011261013f576108026112f6565b81805260208281526040808420335f908152925290205460ff1615610880576001600160a01b0316801561084b576bffffffffffffffffffffffff60a01b600854161760085580f35b60405162461bcd60e51b815260206004820152600d60248201526c1253959053125117d590555315609a1b6044820152606490fd5b63e2517d3f60e01b8252336004526024829052604482fd5b503461013f57604036600319011261013f5760243560043567ffffffffffffffff821161063157366023830112156106315781600401359167ffffffffffffffff8311610acc573660248483010111610acc576002548452600560205260408420908282556001820161090b815461130c565b601f8111610a7a575b508486601f8211600114610a12578791610a04575b508560011b905f198760031b1c19161790555b828552600360205267ffffffffffffffff604086205460a81c16620151808102908082046201518014901517156109f05790606060209660026109a07fe329ef3cbffa31df6d4ce1a3434d5e2f1d3c36c623ba785df206407fdf5154c595426113b0565b95019485556002549687956109b48761138e565b6002555491816024604051968795865260408d8701528260408701520185850137828201840152601f01601f19168101030190a3604051908152f35b634e487b7160e01b86526011600452602486fd5b60249150830101355f610929565b82885260208820915086601f198116895b818110610a5c575010610a40575b5050600185811b01905561093c565b8301602401355f19600388901b60f8161c191690555f80610a31565b8684016024013585556001909401936020938401938a935001610a23565b858111156109145781875260208720601f870160051c9060208810610ac4575b81601f9101920160051c0390875b828110610ab6575050610914565b808960019284015501610aa8565b889150610a9a565b8380fd5b503461013f578060031936011261013f576008546040516001600160a01b039091168152602090f35b503461013f57602036600319011261013f576004358082526005602052604082206002810154421180610bba575b15610b7c5760058101805460ff191660011790556003810154600490910154604080519290911082524260208301527fa343a04149c0a6043f138d8b88de27d6b196b0d4c4ed894c888d82194e57ec7c91a280f35b60405162461bcd60e51b81526020600482015260166024820152751393d517d49150511657d3d497d1925390531256915160521b6044820152606490fd5b5060ff60058201541615610b27565b503461013f57606036600319011261013f57610be36112f6565b60243560ff8116809103610631576044359067ffffffffffffffff8216809203610acc5760015492610c148461138e565b60015560405190610100820182811067ffffffffffffffff821117610d4857927fd5e8d9ec2efa3ad9dddddd1463468fe0584e4de99d109b0b0c85ed35a202d88f92600587610d3160209a6040968399885260018060a01b0316998a86528c86019087825288870190815260608701838152608088019184835260a08901934285528b60c08b019680885260e08c0199818b528152602060039052209960018060a01b039051166bffffffffffffffffffffffff60a01b8b5416178a55519089549067ffffffffffffffff60a81b905160a81b169160ff60a01b9060a01b169068ffffffffffffffffff60a01b191617178855516001880155516002870155516003860155511515600485019060ff801983541691151516179055565b5191015581519081524287820152a3604051908152f35b634e487b7160e01b87526041600452602487fd5b503461013f57604036600319011261013f57610d766112e0565b336001600160a01b03821603610d92576102b9906004356114b9565b63334bd91960e11b8252600482fd5b503461013f57604036600319011261013f576102b9600435610dc16112e0565b90610dda6102af825f525f602052600160405f20015490565b611431565b503461013f57602036600319011261013f576020610e0a6004355f525f602052600160405f20015490565b604051908152f35b503461013f57602036600319011261013f5760016040602092600435815260038452200154604051908152f35b503461114e57602036600319011261114e5760043560027f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00541461127e5760027f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055805f52600560205260405f209060058201805460ff8116156112495760ff8160081c16611211576003840154600485015410156111dd576101009061ff0019161790557ff758fc91e01b00ea6b4a6138756f7f28e021f9bf21db6dbf8c36c88eb737257a6020604051428152a26008546001600160a01b031690811515806111c7575b610f51575b8260017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005580f35b6001810190604051915f9083815492610f698461130c565b9283835260208301946001811690815f146111a85750600114611164575b50610f9492500384611344565b60608380518101031261114e57610faa9061137a565b92610fbc60606040850151940161137a565b6001600160a01b039485169416908415158061115b575b80611152575b610fe6575b505050610f29565b803b1561114e5760405163f3fef3a360e01b81526001600160a01b038616600482015260248101859052905f908290604490829084905af180156111435761112e575b5060405163a9059cbb60e01b81526001600160a01b03821660048201526024810184905260208160448189895af19081156111235786916110e4575b50156110ad577f3994603528a4ce7136b8ccd2cc94e047aa196fd8a1f735323028f632ad214dfb92606092549460405192835260208301526040820152a25f80808080610fde565b60405162461bcd60e51b815260206004820152600f60248201526e1514905394d1915497d19052531151608a1b6044820152606490fd5b90506020813d60201161111b575b816110ff60209383611344565b8101031261111757518015158103611117575f611065565b8580fd5b3d91506110f2565b6040513d88823e3d90fd5b61113b9195505f90611344565b5f935f611029565b6040513d5f823e3d90fd5b5f80fd5b50811515610fd9565b50831515610fd3565b90505f9291925260205f20905f915b81831061118c575050906020610f94928201015f610f87565b6020919350806001915483858a01015201910190918592611173565b60ff1916865250610f9493151560051b830160200191505f9050610f87565b5060606111d7600183015461130c565b14610f24565b60405162461bcd60e51b815260206004820152600c60248201526b1393d517d054141493d5915160a21b6044820152606490fd5b60405162461bcd60e51b815260206004820152601060248201526f1053149150511657d1561150d555115160821b6044820152606490fd5b60405162461bcd60e51b815260206004820152600d60248201526c1393d517d19253905312569151609a1b6044820152606490fd5b633ee5aeb560e01b5f5260045ffd5b3461114e57602036600319011261114e576004359063ffffffff60e01b821680920361114e57602091637965db0b60e01b81149081156112cf575b5015158152f35b6301ffc9a760e01b149050836112c8565b602435906001600160a01b038216820361114e57565b600435906001600160a01b038216820361114e57565b90600182811c9216801561133a575b602083101461132657565b634e487b7160e01b5f52602260045260245ffd5b91607f169161131b565b90601f8019910116810190811067ffffffffffffffff82111761136657604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b038216820361114e57565b5f19811461139c5760010190565b634e487b7160e01b5f52601160045260245ffd5b9190820180921161139c57565b156113c457565b60405162461bcd60e51b815260206004820152600d60248201526c111053d7d393d517d193d55391609a1b6044820152606490fd5b5f8181526020818152604080832033845290915290205460ff161561141b5750565b63e2517d3f60e01b5f523360045260245260445ffd5b5f818152602081815260408083206001600160a01b038616845290915290205460ff166114b3575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50505f90565b5f818152602081815260408083206001600160a01b038616845290915290205460ff16156114b3575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a460019056fea2646970667358221220940c3bd3bcd32d4c5dd51a90f0ed3d395ca92d4ff7d6935f5f30120a87302af964736f6c63430008210033ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
}

// DAOManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use DAOManagerMetaData.ABI instead.
var DAOManagerABI = DAOManagerMetaData.ABI

// DAOManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DAOManagerMetaData.Bin instead.
var DAOManagerBin = DAOManagerMetaData.Bin

// DeployDAOManager deploys a new Ethereum contract, binding an instance of DAOManager to it.
func DeployDAOManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DAOManager, error) {
	parsed, err := DAOManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DAOManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DAOManager{DAOManagerCaller: DAOManagerCaller{contract: contract}, DAOManagerTransactor: DAOManagerTransactor{contract: contract}, DAOManagerFilterer: DAOManagerFilterer{contract: contract}}, nil
}

// DAOManager is an auto generated Go binding around an Ethereum contract.
type DAOManager struct {
	DAOManagerCaller     // Read-only binding to the contract
	DAOManagerTransactor // Write-only binding to the contract
	DAOManagerFilterer   // Log filterer for contract events
}

// DAOManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DAOManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAOManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DAOManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAOManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DAOManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAOManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DAOManagerSession struct {
	Contract     *DAOManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DAOManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DAOManagerCallerSession struct {
	Contract *DAOManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DAOManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DAOManagerTransactorSession struct {
	Contract     *DAOManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DAOManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DAOManagerRaw struct {
	Contract *DAOManager // Generic contract binding to access the raw methods on
}

// DAOManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DAOManagerCallerRaw struct {
	Contract *DAOManagerCaller // Generic read-only contract binding to access the raw methods on
}

// DAOManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DAOManagerTransactorRaw struct {
	Contract *DAOManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDAOManager creates a new instance of DAOManager, bound to a specific deployed contract.
func NewDAOManager(address common.Address, backend bind.ContractBackend) (*DAOManager, error) {
	contract, err := bindDAOManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DAOManager{DAOManagerCaller: DAOManagerCaller{contract: contract}, DAOManagerTransactor: DAOManagerTransactor{contract: contract}, DAOManagerFilterer: DAOManagerFilterer{contract: contract}}, nil
}

// NewDAOManagerCaller creates a new read-only instance of DAOManager, bound to a specific deployed contract.
func NewDAOManagerCaller(address common.Address, caller bind.ContractCaller) (*DAOManagerCaller, error) {
	contract, err := bindDAOManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DAOManagerCaller{contract: contract}, nil
}

// NewDAOManagerTransactor creates a new write-only instance of DAOManager, bound to a specific deployed contract.
func NewDAOManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*DAOManagerTransactor, error) {
	contract, err := bindDAOManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DAOManagerTransactor{contract: contract}, nil
}

// NewDAOManagerFilterer creates a new log filterer instance of DAOManager, bound to a specific deployed contract.
func NewDAOManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*DAOManagerFilterer, error) {
	contract, err := bindDAOManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DAOManagerFilterer{contract: contract}, nil
}

// bindDAOManager binds a generic wrapper to an already deployed contract.
func bindDAOManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DAOManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAOManager *DAOManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DAOManager.Contract.DAOManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAOManager *DAOManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAOManager.Contract.DAOManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAOManager *DAOManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAOManager.Contract.DAOManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAOManager *DAOManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DAOManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAOManager *DAOManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAOManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAOManager *DAOManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAOManager.Contract.contract.Transact(opts, method, params...)
}

// DAOADMINROLE is a free data retrieval call binding the contract method 0xd14dd2f2.
//
// Solidity: function DAO_ADMIN_ROLE() view returns(bytes32)
func (_DAOManager *DAOManagerCaller) DAOADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DAOManager.contract.Call(opts, &out, "DAO_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DAOADMINROLE is a free data retrieval call binding the contract method 0xd14dd2f2.
//
// Solidity: function DAO_ADMIN_ROLE() view returns(bytes32)
func (_DAOManager *DAOManagerSession) DAOADMINROLE() ([32]byte, error) {
	return _DAOManager.Contract.DAOADMINROLE(&_DAOManager.CallOpts)
}

// DAOADMINROLE is a free data retrieval call binding the contract method 0xd14dd2f2.
//
// Solidity: function DAO_ADMIN_ROLE() view returns(bytes32)
func (_DAOManager *DAOManagerCallerSession) DAOADMINROLE() ([32]byte, error) {
	return _DAOManager.Contract.DAOADMINROLE(&_DAOManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DAOManager *DAOManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DAOManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DAOManager *DAOManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DAOManager.Contract.DEFAULTADMINROLE(&_DAOManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DAOManager *DAOManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DAOManager.Contract.DEFAULTADMINROLE(&_DAOManager.CallOpts)
}

// TREASURYFUNDERROLE is a free data retrieval call binding the contract method 0xf859e5dc.
//
// Solidity: function TREASURY_FUNDER_ROLE() view returns(bytes32)
func (_DAOManager *DAOManagerCaller) TREASURYFUNDERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DAOManager.contract.Call(opts, &out, "TREASURY_FUNDER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TREASURYFUNDERROLE is a free data retrieval call binding the contract method 0xf859e5dc.
//
// Solidity: function TREASURY_FUNDER_ROLE() view returns(bytes32)
func (_DAOManager *DAOManagerSession) TREASURYFUNDERROLE() ([32]byte, error) {
	return _DAOManager.Contract.TREASURYFUNDERROLE(&_DAOManager.CallOpts)
}

// TREASURYFUNDERROLE is a free data retrieval call binding the contract method 0xf859e5dc.
//
// Solidity: function TREASURY_FUNDER_ROLE() view returns(bytes32)
func (_DAOManager *DAOManagerCallerSession) TREASURYFUNDERROLE() ([32]byte, error) {
	return _DAOManager.Contract.TREASURYFUNDERROLE(&_DAOManager.CallOpts)
}

// GetBnplTerms is a free data retrieval call binding the contract method 0xd3c095a3.
//
// Solidity: function getBnplTerms(uint256 daoId) view returns(uint256 numInstallments, uint256 allowedIntervalMinDays, uint256 allowedIntervalMaxDays, uint256 lateFeeBps, uint256 gracePeriodDays, bool rescheduleAllowed, uint256 minDownPaymentBps)
func (_DAOManager *DAOManagerCaller) GetBnplTerms(opts *bind.CallOpts, daoId *big.Int) (struct {
	NumInstallments        *big.Int
	AllowedIntervalMinDays *big.Int
	AllowedIntervalMaxDays *big.Int
	LateFeeBps             *big.Int
	GracePeriodDays        *big.Int
	RescheduleAllowed      bool
	MinDownPaymentBps      *big.Int
}, error) {
	var out []interface{}
	err := _DAOManager.contract.Call(opts, &out, "getBnplTerms", daoId)

	outstruct := new(struct {
		NumInstallments        *big.Int
		AllowedIntervalMinDays *big.Int
		AllowedIntervalMaxDays *big.Int
		LateFeeBps             *big.Int
		GracePeriodDays        *big.Int
		RescheduleAllowed      bool
		MinDownPaymentBps      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NumInstallments = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AllowedIntervalMinDays = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AllowedIntervalMaxDays = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LateFeeBps = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.GracePeriodDays = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RescheduleAllowed = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.MinDownPaymentBps = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBnplTerms is a free data retrieval call binding the contract method 0xd3c095a3.
//
// Solidity: function getBnplTerms(uint256 daoId) view returns(uint256 numInstallments, uint256 allowedIntervalMinDays, uint256 allowedIntervalMaxDays, uint256 lateFeeBps, uint256 gracePeriodDays, bool rescheduleAllowed, uint256 minDownPaymentBps)
func (_DAOManager *DAOManagerSession) GetBnplTerms(daoId *big.Int) (struct {
	NumInstallments        *big.Int
	AllowedIntervalMinDays *big.Int
	AllowedIntervalMaxDays *big.Int
	LateFeeBps             *big.Int
	GracePeriodDays        *big.Int
	RescheduleAllowed      bool
	MinDownPaymentBps      *big.Int
}, error) {
	return _DAOManager.Contract.GetBnplTerms(&_DAOManager.CallOpts, daoId)
}

// GetBnplTerms is a free data retrieval call binding the contract method 0xd3c095a3.
//
// Solidity: function getBnplTerms(uint256 daoId) view returns(uint256 numInstallments, uint256 allowedIntervalMinDays, uint256 allowedIntervalMaxDays, uint256 lateFeeBps, uint256 gracePeriodDays, bool rescheduleAllowed, uint256 minDownPaymentBps)
func (_DAOManager *DAOManagerCallerSession) GetBnplTerms(daoId *big.Int) (struct {
	NumInstallments        *big.Int
	AllowedIntervalMinDays *big.Int
	AllowedIntervalMaxDays *big.Int
	LateFeeBps             *big.Int
	GracePeriodDays        *big.Int
	RescheduleAllowed      bool
	MinDownPaymentBps      *big.Int
}, error) {
	return _DAOManager.Contract.GetBnplTerms(&_DAOManager.CallOpts, daoId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DAOManager *DAOManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _DAOManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DAOManager *DAOManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DAOManager.Contract.GetRoleAdmin(&_DAOManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DAOManager *DAOManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DAOManager.Contract.GetRoleAdmin(&_DAOManager.CallOpts, role)
}

// GetTreasuryBalance is a free data retrieval call binding the contract method 0x184512cd.
//
// Solidity: function getTreasuryBalance(uint256 daoId) view returns(uint256)
func (_DAOManager *DAOManagerCaller) GetTreasuryBalance(opts *bind.CallOpts, daoId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DAOManager.contract.Call(opts, &out, "getTreasuryBalance", daoId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTreasuryBalance is a free data retrieval call binding the contract method 0x184512cd.
//
// Solidity: function getTreasuryBalance(uint256 daoId) view returns(uint256)
func (_DAOManager *DAOManagerSession) GetTreasuryBalance(daoId *big.Int) (*big.Int, error) {
	return _DAOManager.Contract.GetTreasuryBalance(&_DAOManager.CallOpts, daoId)
}

// GetTreasuryBalance is a free data retrieval call binding the contract method 0x184512cd.
//
// Solidity: function getTreasuryBalance(uint256 daoId) view returns(uint256)
func (_DAOManager *DAOManagerCallerSession) GetTreasuryBalance(daoId *big.Int) (*big.Int, error) {
	return _DAOManager.Contract.GetTreasuryBalance(&_DAOManager.CallOpts, daoId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DAOManager *DAOManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _DAOManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DAOManager *DAOManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DAOManager.Contract.HasRole(&_DAOManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DAOManager *DAOManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DAOManager.Contract.HasRole(&_DAOManager.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DAOManager *DAOManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DAOManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DAOManager *DAOManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DAOManager.Contract.SupportsInterface(&_DAOManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DAOManager *DAOManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DAOManager.Contract.SupportsInterface(&_DAOManager.CallOpts, interfaceId)
}

// TokenVault is a free data retrieval call binding the contract method 0x5bc789d9.
//
// Solidity: function tokenVault() view returns(address)
func (_DAOManager *DAOManagerCaller) TokenVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DAOManager.contract.Call(opts, &out, "tokenVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenVault is a free data retrieval call binding the contract method 0x5bc789d9.
//
// Solidity: function tokenVault() view returns(address)
func (_DAOManager *DAOManagerSession) TokenVault() (common.Address, error) {
	return _DAOManager.Contract.TokenVault(&_DAOManager.CallOpts)
}

// TokenVault is a free data retrieval call binding the contract method 0x5bc789d9.
//
// Solidity: function tokenVault() view returns(address)
func (_DAOManager *DAOManagerCallerSession) TokenVault() (common.Address, error) {
	return _DAOManager.Contract.TokenVault(&_DAOManager.CallOpts)
}

// CreateDAO is a paid mutator transaction binding the contract method 0x40e8928d.
//
// Solidity: function createDAO(address creator, uint8 goal, uint64 votingPeriodDays) returns(uint256 daoId)
func (_DAOManager *DAOManagerTransactor) CreateDAO(opts *bind.TransactOpts, creator common.Address, goal uint8, votingPeriodDays uint64) (*types.Transaction, error) {
	return _DAOManager.contract.Transact(opts, "createDAO", creator, goal, votingPeriodDays)
}

// CreateDAO is a paid mutator transaction binding the contract method 0x40e8928d.
//
// Solidity: function createDAO(address creator, uint8 goal, uint64 votingPeriodDays) returns(uint256 daoId)
func (_DAOManager *DAOManagerSession) CreateDAO(creator common.Address, goal uint8, votingPeriodDays uint64) (*types.Transaction, error) {
	return _DAOManager.Contract.CreateDAO(&_DAOManager.TransactOpts, creator, goal, votingPeriodDays)
}

// CreateDAO is a paid mutator transaction binding the contract method 0x40e8928d.
//
// Solidity: function createDAO(address creator, uint8 goal, uint64 votingPeriodDays) returns(uint256 daoId)
func (_DAOManager *DAOManagerTransactorSession) CreateDAO(creator common.Address, goal uint8, votingPeriodDays uint64) (*types.Transaction, error) {
	return _DAOManager.Contract.CreateDAO(&_DAOManager.TransactOpts, creator, goal, votingPeriodDays)
}

// CreditTreasury is a paid mutator transaction binding the contract method 0xcec4aec9.
//
// Solidity: function creditTreasury(uint256 daoId, uint256 amount) returns()
func (_DAOManager *DAOManagerTransactor) CreditTreasury(opts *bind.TransactOpts, daoId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _DAOManager.contract.Transact(opts, "creditTreasury", daoId, amount)
}

// CreditTreasury is a paid mutator transaction binding the contract method 0xcec4aec9.
//
// Solidity: function creditTreasury(uint256 daoId, uint256 amount) returns()
func (_DAOManager *DAOManagerSession) CreditTreasury(daoId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _DAOManager.Contract.CreditTreasury(&_DAOManager.TransactOpts, daoId, amount)
}

// CreditTreasury is a paid mutator transaction binding the contract method 0xcec4aec9.
//
// Solidity: function creditTreasury(uint256 daoId, uint256 amount) returns()
func (_DAOManager *DAOManagerTransactorSession) CreditTreasury(daoId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _DAOManager.Contract.CreditTreasury(&_DAOManager.TransactOpts, daoId, amount)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0d61b519.
//
// Solidity: function executeProposal(uint256 proposalId) returns()
func (_DAOManager *DAOManagerTransactor) ExecuteProposal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _DAOManager.contract.Transact(opts, "executeProposal", proposalId)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0d61b519.
//
// Solidity: function executeProposal(uint256 proposalId) returns()
func (_DAOManager *DAOManagerSession) ExecuteProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _DAOManager.Contract.ExecuteProposal(&_DAOManager.TransactOpts, proposalId)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0d61b519.
//
// Solidity: function executeProposal(uint256 proposalId) returns()
func (_DAOManager *DAOManagerTransactorSession) ExecuteProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _DAOManager.Contract.ExecuteProposal(&_DAOManager.TransactOpts, proposalId)
}

// FinalizeProposal is a paid mutator transaction binding the contract method 0x5652077c.
//
// Solidity: function finalizeProposal(uint256 proposalId) returns()
func (_DAOManager *DAOManagerTransactor) FinalizeProposal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _DAOManager.contract.Transact(opts, "finalizeProposal", proposalId)
}

// FinalizeProposal is a paid mutator transaction binding the contract method 0x5652077c.
//
// Solidity: function finalizeProposal(uint256 proposalId) returns()
func (_DAOManager *DAOManagerSession) FinalizeProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _DAOManager.Contract.FinalizeProposal(&_DAOManager.TransactOpts, proposalId)
}

// FinalizeProposal is a paid mutator transaction binding the contract method 0x5652077c.
//
// Solidity: function finalizeProposal(uint256 proposalId) returns()
func (_DAOManager *DAOManagerTransactorSession) FinalizeProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _DAOManager.Contract.FinalizeProposal(&_DAOManager.TransactOpts, proposalId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DAOManager *DAOManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DAOManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DAOManager *DAOManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DAOManager.Contract.GrantRole(&_DAOManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DAOManager *DAOManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DAOManager.Contract.GrantRole(&_DAOManager.TransactOpts, role, account)
}

// JoinDAO is a paid mutator transaction binding the contract method 0xe3065c20.
//
// Solidity: function joinDAO(uint256 daoId, address member, uint256 investmentAmount) returns()
func (_DAOManager *DAOManagerTransactor) JoinDAO(opts *bind.TransactOpts, daoId *big.Int, member common.Address, investmentAmount *big.Int) (*types.Transaction, error) {
	return _DAOManager.contract.Transact(opts, "joinDAO", daoId, member, investmentAmount)
}

// JoinDAO is a paid mutator transaction binding the contract method 0xe3065c20.
//
// Solidity: function joinDAO(uint256 daoId, address member, uint256 investmentAmount) returns()
func (_DAOManager *DAOManagerSession) JoinDAO(daoId *big.Int, member common.Address, investmentAmount *big.Int) (*types.Transaction, error) {
	return _DAOManager.Contract.JoinDAO(&_DAOManager.TransactOpts, daoId, member, investmentAmount)
}

// JoinDAO is a paid mutator transaction binding the contract method 0xe3065c20.
//
// Solidity: function joinDAO(uint256 daoId, address member, uint256 investmentAmount) returns()
func (_DAOManager *DAOManagerTransactorSession) JoinDAO(daoId *big.Int, member common.Address, investmentAmount *big.Int) (*types.Transaction, error) {
	return _DAOManager.Contract.JoinDAO(&_DAOManager.TransactOpts, daoId, member, investmentAmount)
}

// Propose is a paid mutator transaction binding the contract method 0x5f90ebaf.
//
// Solidity: function propose(uint256 daoId, bytes proposalData) returns(uint256 proposalId)
func (_DAOManager *DAOManagerTransactor) Propose(opts *bind.TransactOpts, daoId *big.Int, proposalData []byte) (*types.Transaction, error) {
	return _DAOManager.contract.Transact(opts, "propose", daoId, proposalData)
}

// Propose is a paid mutator transaction binding the contract method 0x5f90ebaf.
//
// Solidity: function propose(uint256 daoId, bytes proposalData) returns(uint256 proposalId)
func (_DAOManager *DAOManagerSession) Propose(daoId *big.Int, proposalData []byte) (*types.Transaction, error) {
	return _DAOManager.Contract.Propose(&_DAOManager.TransactOpts, daoId, proposalData)
}

// Propose is a paid mutator transaction binding the contract method 0x5f90ebaf.
//
// Solidity: function propose(uint256 daoId, bytes proposalData) returns(uint256 proposalId)
func (_DAOManager *DAOManagerTransactorSession) Propose(daoId *big.Int, proposalData []byte) (*types.Transaction, error) {
	return _DAOManager.Contract.Propose(&_DAOManager.TransactOpts, daoId, proposalData)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_DAOManager *DAOManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _DAOManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_DAOManager *DAOManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _DAOManager.Contract.RenounceRole(&_DAOManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_DAOManager *DAOManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _DAOManager.Contract.RenounceRole(&_DAOManager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DAOManager *DAOManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DAOManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DAOManager *DAOManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DAOManager.Contract.RevokeRole(&_DAOManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DAOManager *DAOManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DAOManager.Contract.RevokeRole(&_DAOManager.TransactOpts, role, account)
}

// SetBnplTerms is a paid mutator transaction binding the contract method 0xbc4b9101.
//
// Solidity: function setBnplTerms(uint256 daoId, uint256 numInstallments, uint256 allowedIntervalMinDays, uint256 allowedIntervalMaxDays, uint256 lateFeeBps, uint256 gracePeriodDays, bool rescheduleAllowed, uint256 minDownPaymentBps) returns()
func (_DAOManager *DAOManagerTransactor) SetBnplTerms(opts *bind.TransactOpts, daoId *big.Int, numInstallments *big.Int, allowedIntervalMinDays *big.Int, allowedIntervalMaxDays *big.Int, lateFeeBps *big.Int, gracePeriodDays *big.Int, rescheduleAllowed bool, minDownPaymentBps *big.Int) (*types.Transaction, error) {
	return _DAOManager.contract.Transact(opts, "setBnplTerms", daoId, numInstallments, allowedIntervalMinDays, allowedIntervalMaxDays, lateFeeBps, gracePeriodDays, rescheduleAllowed, minDownPaymentBps)
}

// SetBnplTerms is a paid mutator transaction binding the contract method 0xbc4b9101.
//
// Solidity: function setBnplTerms(uint256 daoId, uint256 numInstallments, uint256 allowedIntervalMinDays, uint256 allowedIntervalMaxDays, uint256 lateFeeBps, uint256 gracePeriodDays, bool rescheduleAllowed, uint256 minDownPaymentBps) returns()
func (_DAOManager *DAOManagerSession) SetBnplTerms(daoId *big.Int, numInstallments *big.Int, allowedIntervalMinDays *big.Int, allowedIntervalMaxDays *big.Int, lateFeeBps *big.Int, gracePeriodDays *big.Int, rescheduleAllowed bool, minDownPaymentBps *big.Int) (*types.Transaction, error) {
	return _DAOManager.Contract.SetBnplTerms(&_DAOManager.TransactOpts, daoId, numInstallments, allowedIntervalMinDays, allowedIntervalMaxDays, lateFeeBps, gracePeriodDays, rescheduleAllowed, minDownPaymentBps)
}

// SetBnplTerms is a paid mutator transaction binding the contract method 0xbc4b9101.
//
// Solidity: function setBnplTerms(uint256 daoId, uint256 numInstallments, uint256 allowedIntervalMinDays, uint256 allowedIntervalMaxDays, uint256 lateFeeBps, uint256 gracePeriodDays, bool rescheduleAllowed, uint256 minDownPaymentBps) returns()
func (_DAOManager *DAOManagerTransactorSession) SetBnplTerms(daoId *big.Int, numInstallments *big.Int, allowedIntervalMinDays *big.Int, allowedIntervalMaxDays *big.Int, lateFeeBps *big.Int, gracePeriodDays *big.Int, rescheduleAllowed bool, minDownPaymentBps *big.Int) (*types.Transaction, error) {
	return _DAOManager.Contract.SetBnplTerms(&_DAOManager.TransactOpts, daoId, numInstallments, allowedIntervalMinDays, allowedIntervalMaxDays, lateFeeBps, gracePeriodDays, rescheduleAllowed, minDownPaymentBps)
}

// SetTokenVault is a paid mutator transaction binding the contract method 0x6497a8a0.
//
// Solidity: function setTokenVault(address _tokenVault) returns()
func (_DAOManager *DAOManagerTransactor) SetTokenVault(opts *bind.TransactOpts, _tokenVault common.Address) (*types.Transaction, error) {
	return _DAOManager.contract.Transact(opts, "setTokenVault", _tokenVault)
}

// SetTokenVault is a paid mutator transaction binding the contract method 0x6497a8a0.
//
// Solidity: function setTokenVault(address _tokenVault) returns()
func (_DAOManager *DAOManagerSession) SetTokenVault(_tokenVault common.Address) (*types.Transaction, error) {
	return _DAOManager.Contract.SetTokenVault(&_DAOManager.TransactOpts, _tokenVault)
}

// SetTokenVault is a paid mutator transaction binding the contract method 0x6497a8a0.
//
// Solidity: function setTokenVault(address _tokenVault) returns()
func (_DAOManager *DAOManagerTransactorSession) SetTokenVault(_tokenVault common.Address) (*types.Transaction, error) {
	return _DAOManager.Contract.SetTokenVault(&_DAOManager.TransactOpts, _tokenVault)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 proposalId, bool support) returns()
func (_DAOManager *DAOManagerTransactor) Vote(opts *bind.TransactOpts, proposalId *big.Int, support bool) (*types.Transaction, error) {
	return _DAOManager.contract.Transact(opts, "vote", proposalId, support)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 proposalId, bool support) returns()
func (_DAOManager *DAOManagerSession) Vote(proposalId *big.Int, support bool) (*types.Transaction, error) {
	return _DAOManager.Contract.Vote(&_DAOManager.TransactOpts, proposalId, support)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 proposalId, bool support) returns()
func (_DAOManager *DAOManagerTransactorSession) Vote(proposalId *big.Int, support bool) (*types.Transaction, error) {
	return _DAOManager.Contract.Vote(&_DAOManager.TransactOpts, proposalId, support)
}

// DAOManagerBnplTermsUpdatedIterator is returned from FilterBnplTermsUpdated and is used to iterate over the raw logs and unpacked data for BnplTermsUpdated events raised by the DAOManager contract.
type DAOManagerBnplTermsUpdatedIterator struct {
	Event *DAOManagerBnplTermsUpdated // Event containing the contract specifics and raw log

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
func (it *DAOManagerBnplTermsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOManagerBnplTermsUpdated)
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
		it.Event = new(DAOManagerBnplTermsUpdated)
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
func (it *DAOManagerBnplTermsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOManagerBnplTermsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOManagerBnplTermsUpdated represents a BnplTermsUpdated event raised by the DAOManager contract.
type DAOManagerBnplTermsUpdated struct {
	DaoId                  *big.Int
	NumInstallments        *big.Int
	AllowedIntervalMinDays *big.Int
	AllowedIntervalMaxDays *big.Int
	LateFeeBps             *big.Int
	GracePeriodDays        *big.Int
	RescheduleAllowed      bool
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterBnplTermsUpdated is a free log retrieval operation binding the contract event 0x7e4912d5452f297c8c78ebd5196c5d177af96195a7ef704138d2031ebba23b30.
//
// Solidity: event BnplTermsUpdated(uint256 indexed daoId, uint256 numInstallments, uint256 allowedIntervalMinDays, uint256 allowedIntervalMaxDays, uint256 lateFeeBps, uint256 gracePeriodDays, bool rescheduleAllowed)
func (_DAOManager *DAOManagerFilterer) FilterBnplTermsUpdated(opts *bind.FilterOpts, daoId []*big.Int) (*DAOManagerBnplTermsUpdatedIterator, error) {

	var daoIdRule []interface{}
	for _, daoIdItem := range daoId {
		daoIdRule = append(daoIdRule, daoIdItem)
	}

	logs, sub, err := _DAOManager.contract.FilterLogs(opts, "BnplTermsUpdated", daoIdRule)
	if err != nil {
		return nil, err
	}
	return &DAOManagerBnplTermsUpdatedIterator{contract: _DAOManager.contract, event: "BnplTermsUpdated", logs: logs, sub: sub}, nil
}

// WatchBnplTermsUpdated is a free log subscription operation binding the contract event 0x7e4912d5452f297c8c78ebd5196c5d177af96195a7ef704138d2031ebba23b30.
//
// Solidity: event BnplTermsUpdated(uint256 indexed daoId, uint256 numInstallments, uint256 allowedIntervalMinDays, uint256 allowedIntervalMaxDays, uint256 lateFeeBps, uint256 gracePeriodDays, bool rescheduleAllowed)
func (_DAOManager *DAOManagerFilterer) WatchBnplTermsUpdated(opts *bind.WatchOpts, sink chan<- *DAOManagerBnplTermsUpdated, daoId []*big.Int) (event.Subscription, error) {

	var daoIdRule []interface{}
	for _, daoIdItem := range daoId {
		daoIdRule = append(daoIdRule, daoIdItem)
	}

	logs, sub, err := _DAOManager.contract.WatchLogs(opts, "BnplTermsUpdated", daoIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOManagerBnplTermsUpdated)
				if err := _DAOManager.contract.UnpackLog(event, "BnplTermsUpdated", log); err != nil {
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

// ParseBnplTermsUpdated is a log parse operation binding the contract event 0x7e4912d5452f297c8c78ebd5196c5d177af96195a7ef704138d2031ebba23b30.
//
// Solidity: event BnplTermsUpdated(uint256 indexed daoId, uint256 numInstallments, uint256 allowedIntervalMinDays, uint256 allowedIntervalMaxDays, uint256 lateFeeBps, uint256 gracePeriodDays, bool rescheduleAllowed)
func (_DAOManager *DAOManagerFilterer) ParseBnplTermsUpdated(log types.Log) (*DAOManagerBnplTermsUpdated, error) {
	event := new(DAOManagerBnplTermsUpdated)
	if err := _DAOManager.contract.UnpackLog(event, "BnplTermsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOManagerDaoCreatedIterator is returned from FilterDaoCreated and is used to iterate over the raw logs and unpacked data for DaoCreated events raised by the DAOManager contract.
type DAOManagerDaoCreatedIterator struct {
	Event *DAOManagerDaoCreated // Event containing the contract specifics and raw log

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
func (it *DAOManagerDaoCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOManagerDaoCreated)
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
		it.Event = new(DAOManagerDaoCreated)
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
func (it *DAOManagerDaoCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOManagerDaoCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOManagerDaoCreated represents a DaoCreated event raised by the DAOManager contract.
type DAOManagerDaoCreated struct {
	DaoId     *big.Int
	Creator   common.Address
	Goal      uint8
	CreatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDaoCreated is a free log retrieval operation binding the contract event 0xd5e8d9ec2efa3ad9dddddd1463468fe0584e4de99d109b0b0c85ed35a202d88f.
//
// Solidity: event DaoCreated(uint256 indexed daoId, address indexed creator, uint8 goal, uint256 createdAt)
func (_DAOManager *DAOManagerFilterer) FilterDaoCreated(opts *bind.FilterOpts, daoId []*big.Int, creator []common.Address) (*DAOManagerDaoCreatedIterator, error) {

	var daoIdRule []interface{}
	for _, daoIdItem := range daoId {
		daoIdRule = append(daoIdRule, daoIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _DAOManager.contract.FilterLogs(opts, "DaoCreated", daoIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &DAOManagerDaoCreatedIterator{contract: _DAOManager.contract, event: "DaoCreated", logs: logs, sub: sub}, nil
}

// WatchDaoCreated is a free log subscription operation binding the contract event 0xd5e8d9ec2efa3ad9dddddd1463468fe0584e4de99d109b0b0c85ed35a202d88f.
//
// Solidity: event DaoCreated(uint256 indexed daoId, address indexed creator, uint8 goal, uint256 createdAt)
func (_DAOManager *DAOManagerFilterer) WatchDaoCreated(opts *bind.WatchOpts, sink chan<- *DAOManagerDaoCreated, daoId []*big.Int, creator []common.Address) (event.Subscription, error) {

	var daoIdRule []interface{}
	for _, daoIdItem := range daoId {
		daoIdRule = append(daoIdRule, daoIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _DAOManager.contract.WatchLogs(opts, "DaoCreated", daoIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOManagerDaoCreated)
				if err := _DAOManager.contract.UnpackLog(event, "DaoCreated", log); err != nil {
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

// ParseDaoCreated is a log parse operation binding the contract event 0xd5e8d9ec2efa3ad9dddddd1463468fe0584e4de99d109b0b0c85ed35a202d88f.
//
// Solidity: event DaoCreated(uint256 indexed daoId, address indexed creator, uint8 goal, uint256 createdAt)
func (_DAOManager *DAOManagerFilterer) ParseDaoCreated(log types.Log) (*DAOManagerDaoCreated, error) {
	event := new(DAOManagerDaoCreated)
	if err := _DAOManager.contract.UnpackLog(event, "DaoCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOManagerMemberJoinedIterator is returned from FilterMemberJoined and is used to iterate over the raw logs and unpacked data for MemberJoined events raised by the DAOManager contract.
type DAOManagerMemberJoinedIterator struct {
	Event *DAOManagerMemberJoined // Event containing the contract specifics and raw log

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
func (it *DAOManagerMemberJoinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOManagerMemberJoined)
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
		it.Event = new(DAOManagerMemberJoined)
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
func (it *DAOManagerMemberJoinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOManagerMemberJoinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOManagerMemberJoined represents a MemberJoined event raised by the DAOManager contract.
type DAOManagerMemberJoined struct {
	DaoId      *big.Int
	Member     common.Address
	Investment *big.Int
	JoinedAt   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMemberJoined is a free log retrieval operation binding the contract event 0xd1c23df765b8db6e8e7b2077d1e985dc2c7b79dcd08cc1a6dc48b566bddeea73.
//
// Solidity: event MemberJoined(uint256 indexed daoId, address indexed member, uint256 investment, uint256 joinedAt)
func (_DAOManager *DAOManagerFilterer) FilterMemberJoined(opts *bind.FilterOpts, daoId []*big.Int, member []common.Address) (*DAOManagerMemberJoinedIterator, error) {

	var daoIdRule []interface{}
	for _, daoIdItem := range daoId {
		daoIdRule = append(daoIdRule, daoIdItem)
	}
	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _DAOManager.contract.FilterLogs(opts, "MemberJoined", daoIdRule, memberRule)
	if err != nil {
		return nil, err
	}
	return &DAOManagerMemberJoinedIterator{contract: _DAOManager.contract, event: "MemberJoined", logs: logs, sub: sub}, nil
}

// WatchMemberJoined is a free log subscription operation binding the contract event 0xd1c23df765b8db6e8e7b2077d1e985dc2c7b79dcd08cc1a6dc48b566bddeea73.
//
// Solidity: event MemberJoined(uint256 indexed daoId, address indexed member, uint256 investment, uint256 joinedAt)
func (_DAOManager *DAOManagerFilterer) WatchMemberJoined(opts *bind.WatchOpts, sink chan<- *DAOManagerMemberJoined, daoId []*big.Int, member []common.Address) (event.Subscription, error) {

	var daoIdRule []interface{}
	for _, daoIdItem := range daoId {
		daoIdRule = append(daoIdRule, daoIdItem)
	}
	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _DAOManager.contract.WatchLogs(opts, "MemberJoined", daoIdRule, memberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOManagerMemberJoined)
				if err := _DAOManager.contract.UnpackLog(event, "MemberJoined", log); err != nil {
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

// ParseMemberJoined is a log parse operation binding the contract event 0xd1c23df765b8db6e8e7b2077d1e985dc2c7b79dcd08cc1a6dc48b566bddeea73.
//
// Solidity: event MemberJoined(uint256 indexed daoId, address indexed member, uint256 investment, uint256 joinedAt)
func (_DAOManager *DAOManagerFilterer) ParseMemberJoined(log types.Log) (*DAOManagerMemberJoined, error) {
	event := new(DAOManagerMemberJoined)
	if err := _DAOManager.contract.UnpackLog(event, "MemberJoined", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOManagerProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the DAOManager contract.
type DAOManagerProposalExecutedIterator struct {
	Event *DAOManagerProposalExecuted // Event containing the contract specifics and raw log

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
func (it *DAOManagerProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOManagerProposalExecuted)
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
		it.Event = new(DAOManagerProposalExecuted)
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
func (it *DAOManagerProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOManagerProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOManagerProposalExecuted represents a ProposalExecuted event raised by the DAOManager contract.
type DAOManagerProposalExecuted struct {
	ProposalId *big.Int
	ExecutedAt *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0xf758fc91e01b00ea6b4a6138756f7f28e021f9bf21db6dbf8c36c88eb737257a.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalId, uint256 executedAt)
func (_DAOManager *DAOManagerFilterer) FilterProposalExecuted(opts *bind.FilterOpts, proposalId []*big.Int) (*DAOManagerProposalExecutedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _DAOManager.contract.FilterLogs(opts, "ProposalExecuted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &DAOManagerProposalExecutedIterator{contract: _DAOManager.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0xf758fc91e01b00ea6b4a6138756f7f28e021f9bf21db6dbf8c36c88eb737257a.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalId, uint256 executedAt)
func (_DAOManager *DAOManagerFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *DAOManagerProposalExecuted, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _DAOManager.contract.WatchLogs(opts, "ProposalExecuted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOManagerProposalExecuted)
				if err := _DAOManager.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0xf758fc91e01b00ea6b4a6138756f7f28e021f9bf21db6dbf8c36c88eb737257a.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalId, uint256 executedAt)
func (_DAOManager *DAOManagerFilterer) ParseProposalExecuted(log types.Log) (*DAOManagerProposalExecuted, error) {
	event := new(DAOManagerProposalExecuted)
	if err := _DAOManager.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOManagerProposalFinalizedIterator is returned from FilterProposalFinalized and is used to iterate over the raw logs and unpacked data for ProposalFinalized events raised by the DAOManager contract.
type DAOManagerProposalFinalizedIterator struct {
	Event *DAOManagerProposalFinalized // Event containing the contract specifics and raw log

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
func (it *DAOManagerProposalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOManagerProposalFinalized)
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
		it.Event = new(DAOManagerProposalFinalized)
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
func (it *DAOManagerProposalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOManagerProposalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOManagerProposalFinalized represents a ProposalFinalized event raised by the DAOManager contract.
type DAOManagerProposalFinalized struct {
	ProposalId  *big.Int
	Approved    bool
	FinalizedAt *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalFinalized is a free log retrieval operation binding the contract event 0xa343a04149c0a6043f138d8b88de27d6b196b0d4c4ed894c888d82194e57ec7c.
//
// Solidity: event ProposalFinalized(uint256 indexed proposalId, bool approved, uint256 finalizedAt)
func (_DAOManager *DAOManagerFilterer) FilterProposalFinalized(opts *bind.FilterOpts, proposalId []*big.Int) (*DAOManagerProposalFinalizedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _DAOManager.contract.FilterLogs(opts, "ProposalFinalized", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &DAOManagerProposalFinalizedIterator{contract: _DAOManager.contract, event: "ProposalFinalized", logs: logs, sub: sub}, nil
}

// WatchProposalFinalized is a free log subscription operation binding the contract event 0xa343a04149c0a6043f138d8b88de27d6b196b0d4c4ed894c888d82194e57ec7c.
//
// Solidity: event ProposalFinalized(uint256 indexed proposalId, bool approved, uint256 finalizedAt)
func (_DAOManager *DAOManagerFilterer) WatchProposalFinalized(opts *bind.WatchOpts, sink chan<- *DAOManagerProposalFinalized, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _DAOManager.contract.WatchLogs(opts, "ProposalFinalized", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOManagerProposalFinalized)
				if err := _DAOManager.contract.UnpackLog(event, "ProposalFinalized", log); err != nil {
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

// ParseProposalFinalized is a log parse operation binding the contract event 0xa343a04149c0a6043f138d8b88de27d6b196b0d4c4ed894c888d82194e57ec7c.
//
// Solidity: event ProposalFinalized(uint256 indexed proposalId, bool approved, uint256 finalizedAt)
func (_DAOManager *DAOManagerFilterer) ParseProposalFinalized(log types.Log) (*DAOManagerProposalFinalized, error) {
	event := new(DAOManagerProposalFinalized)
	if err := _DAOManager.contract.UnpackLog(event, "ProposalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOManagerProposalOpenedIterator is returned from FilterProposalOpened and is used to iterate over the raw logs and unpacked data for ProposalOpened events raised by the DAOManager contract.
type DAOManagerProposalOpenedIterator struct {
	Event *DAOManagerProposalOpened // Event containing the contract specifics and raw log

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
func (it *DAOManagerProposalOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOManagerProposalOpened)
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
		it.Event = new(DAOManagerProposalOpened)
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
func (it *DAOManagerProposalOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOManagerProposalOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOManagerProposalOpened represents a ProposalOpened event raised by the DAOManager contract.
type DAOManagerProposalOpened struct {
	ProposalId *big.Int
	DaoId      *big.Int
	Expiry     *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalOpened is a free log retrieval operation binding the contract event 0xe329ef3cbffa31df6d4ce1a3434d5e2f1d3c36c623ba785df206407fdf5154c5.
//
// Solidity: event ProposalOpened(uint256 indexed proposalId, uint256 indexed daoId, uint256 expiry, bytes data)
func (_DAOManager *DAOManagerFilterer) FilterProposalOpened(opts *bind.FilterOpts, proposalId []*big.Int, daoId []*big.Int) (*DAOManagerProposalOpenedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var daoIdRule []interface{}
	for _, daoIdItem := range daoId {
		daoIdRule = append(daoIdRule, daoIdItem)
	}

	logs, sub, err := _DAOManager.contract.FilterLogs(opts, "ProposalOpened", proposalIdRule, daoIdRule)
	if err != nil {
		return nil, err
	}
	return &DAOManagerProposalOpenedIterator{contract: _DAOManager.contract, event: "ProposalOpened", logs: logs, sub: sub}, nil
}

// WatchProposalOpened is a free log subscription operation binding the contract event 0xe329ef3cbffa31df6d4ce1a3434d5e2f1d3c36c623ba785df206407fdf5154c5.
//
// Solidity: event ProposalOpened(uint256 indexed proposalId, uint256 indexed daoId, uint256 expiry, bytes data)
func (_DAOManager *DAOManagerFilterer) WatchProposalOpened(opts *bind.WatchOpts, sink chan<- *DAOManagerProposalOpened, proposalId []*big.Int, daoId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var daoIdRule []interface{}
	for _, daoIdItem := range daoId {
		daoIdRule = append(daoIdRule, daoIdItem)
	}

	logs, sub, err := _DAOManager.contract.WatchLogs(opts, "ProposalOpened", proposalIdRule, daoIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOManagerProposalOpened)
				if err := _DAOManager.contract.UnpackLog(event, "ProposalOpened", log); err != nil {
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

// ParseProposalOpened is a log parse operation binding the contract event 0xe329ef3cbffa31df6d4ce1a3434d5e2f1d3c36c623ba785df206407fdf5154c5.
//
// Solidity: event ProposalOpened(uint256 indexed proposalId, uint256 indexed daoId, uint256 expiry, bytes data)
func (_DAOManager *DAOManagerFilterer) ParseProposalOpened(log types.Log) (*DAOManagerProposalOpened, error) {
	event := new(DAOManagerProposalOpened)
	if err := _DAOManager.contract.UnpackLog(event, "ProposalOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the DAOManager contract.
type DAOManagerRoleAdminChangedIterator struct {
	Event *DAOManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *DAOManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOManagerRoleAdminChanged)
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
		it.Event = new(DAOManagerRoleAdminChanged)
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
func (it *DAOManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOManagerRoleAdminChanged represents a RoleAdminChanged event raised by the DAOManager contract.
type DAOManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DAOManager *DAOManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DAOManagerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _DAOManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DAOManagerRoleAdminChangedIterator{contract: _DAOManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DAOManager *DAOManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DAOManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _DAOManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOManagerRoleAdminChanged)
				if err := _DAOManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DAOManager *DAOManagerFilterer) ParseRoleAdminChanged(log types.Log) (*DAOManagerRoleAdminChanged, error) {
	event := new(DAOManagerRoleAdminChanged)
	if err := _DAOManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the DAOManager contract.
type DAOManagerRoleGrantedIterator struct {
	Event *DAOManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *DAOManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOManagerRoleGranted)
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
		it.Event = new(DAOManagerRoleGranted)
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
func (it *DAOManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOManagerRoleGranted represents a RoleGranted event raised by the DAOManager contract.
type DAOManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DAOManager *DAOManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DAOManagerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DAOManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DAOManagerRoleGrantedIterator{contract: _DAOManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DAOManager *DAOManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DAOManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DAOManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOManagerRoleGranted)
				if err := _DAOManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DAOManager *DAOManagerFilterer) ParseRoleGranted(log types.Log) (*DAOManagerRoleGranted, error) {
	event := new(DAOManagerRoleGranted)
	if err := _DAOManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the DAOManager contract.
type DAOManagerRoleRevokedIterator struct {
	Event *DAOManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *DAOManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOManagerRoleRevoked)
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
		it.Event = new(DAOManagerRoleRevoked)
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
func (it *DAOManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOManagerRoleRevoked represents a RoleRevoked event raised by the DAOManager contract.
type DAOManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DAOManager *DAOManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DAOManagerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DAOManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DAOManagerRoleRevokedIterator{contract: _DAOManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DAOManager *DAOManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DAOManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DAOManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOManagerRoleRevoked)
				if err := _DAOManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DAOManager *DAOManagerFilterer) ParseRoleRevoked(log types.Log) (*DAOManagerRoleRevoked, error) {
	event := new(DAOManagerRoleRevoked)
	if err := _DAOManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOManagerTreasuryDepositedIterator is returned from FilterTreasuryDeposited and is used to iterate over the raw logs and unpacked data for TreasuryDeposited events raised by the DAOManager contract.
type DAOManagerTreasuryDepositedIterator struct {
	Event *DAOManagerTreasuryDeposited // Event containing the contract specifics and raw log

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
func (it *DAOManagerTreasuryDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOManagerTreasuryDeposited)
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
		it.Event = new(DAOManagerTreasuryDeposited)
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
func (it *DAOManagerTreasuryDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOManagerTreasuryDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOManagerTreasuryDeposited represents a TreasuryDeposited event raised by the DAOManager contract.
type DAOManagerTreasuryDeposited struct {
	DaoId      *big.Int
	By         common.Address
	Amount     *big.Int
	NewBalance *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTreasuryDeposited is a free log retrieval operation binding the contract event 0x79bdb62d9a421ad24bcd8ac4547a138852f10e01346b861c15122348a55d26a4.
//
// Solidity: event TreasuryDeposited(uint256 indexed daoId, address indexed by, uint256 amount, uint256 newBalance)
func (_DAOManager *DAOManagerFilterer) FilterTreasuryDeposited(opts *bind.FilterOpts, daoId []*big.Int, by []common.Address) (*DAOManagerTreasuryDepositedIterator, error) {

	var daoIdRule []interface{}
	for _, daoIdItem := range daoId {
		daoIdRule = append(daoIdRule, daoIdItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _DAOManager.contract.FilterLogs(opts, "TreasuryDeposited", daoIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return &DAOManagerTreasuryDepositedIterator{contract: _DAOManager.contract, event: "TreasuryDeposited", logs: logs, sub: sub}, nil
}

// WatchTreasuryDeposited is a free log subscription operation binding the contract event 0x79bdb62d9a421ad24bcd8ac4547a138852f10e01346b861c15122348a55d26a4.
//
// Solidity: event TreasuryDeposited(uint256 indexed daoId, address indexed by, uint256 amount, uint256 newBalance)
func (_DAOManager *DAOManagerFilterer) WatchTreasuryDeposited(opts *bind.WatchOpts, sink chan<- *DAOManagerTreasuryDeposited, daoId []*big.Int, by []common.Address) (event.Subscription, error) {

	var daoIdRule []interface{}
	for _, daoIdItem := range daoId {
		daoIdRule = append(daoIdRule, daoIdItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _DAOManager.contract.WatchLogs(opts, "TreasuryDeposited", daoIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOManagerTreasuryDeposited)
				if err := _DAOManager.contract.UnpackLog(event, "TreasuryDeposited", log); err != nil {
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

// ParseTreasuryDeposited is a log parse operation binding the contract event 0x79bdb62d9a421ad24bcd8ac4547a138852f10e01346b861c15122348a55d26a4.
//
// Solidity: event TreasuryDeposited(uint256 indexed daoId, address indexed by, uint256 amount, uint256 newBalance)
func (_DAOManager *DAOManagerFilterer) ParseTreasuryDeposited(log types.Log) (*DAOManagerTreasuryDeposited, error) {
	event := new(DAOManagerTreasuryDeposited)
	if err := _DAOManager.contract.UnpackLog(event, "TreasuryDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOManagerTreasuryWithdrawnIterator is returned from FilterTreasuryWithdrawn and is used to iterate over the raw logs and unpacked data for TreasuryWithdrawn events raised by the DAOManager contract.
type DAOManagerTreasuryWithdrawnIterator struct {
	Event *DAOManagerTreasuryWithdrawn // Event containing the contract specifics and raw log

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
func (it *DAOManagerTreasuryWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOManagerTreasuryWithdrawn)
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
		it.Event = new(DAOManagerTreasuryWithdrawn)
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
func (it *DAOManagerTreasuryWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOManagerTreasuryWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOManagerTreasuryWithdrawn represents a TreasuryWithdrawn event raised by the DAOManager contract.
type DAOManagerTreasuryWithdrawn struct {
	DaoId  *big.Int
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTreasuryWithdrawn is a free log retrieval operation binding the contract event 0x3994603528a4ce7136b8ccd2cc94e047aa196fd8a1f735323028f632ad214dfb.
//
// Solidity: event TreasuryWithdrawn(uint256 indexed daoId, address to, address token, uint256 amount)
func (_DAOManager *DAOManagerFilterer) FilterTreasuryWithdrawn(opts *bind.FilterOpts, daoId []*big.Int) (*DAOManagerTreasuryWithdrawnIterator, error) {

	var daoIdRule []interface{}
	for _, daoIdItem := range daoId {
		daoIdRule = append(daoIdRule, daoIdItem)
	}

	logs, sub, err := _DAOManager.contract.FilterLogs(opts, "TreasuryWithdrawn", daoIdRule)
	if err != nil {
		return nil, err
	}
	return &DAOManagerTreasuryWithdrawnIterator{contract: _DAOManager.contract, event: "TreasuryWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTreasuryWithdrawn is a free log subscription operation binding the contract event 0x3994603528a4ce7136b8ccd2cc94e047aa196fd8a1f735323028f632ad214dfb.
//
// Solidity: event TreasuryWithdrawn(uint256 indexed daoId, address to, address token, uint256 amount)
func (_DAOManager *DAOManagerFilterer) WatchTreasuryWithdrawn(opts *bind.WatchOpts, sink chan<- *DAOManagerTreasuryWithdrawn, daoId []*big.Int) (event.Subscription, error) {

	var daoIdRule []interface{}
	for _, daoIdItem := range daoId {
		daoIdRule = append(daoIdRule, daoIdItem)
	}

	logs, sub, err := _DAOManager.contract.WatchLogs(opts, "TreasuryWithdrawn", daoIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOManagerTreasuryWithdrawn)
				if err := _DAOManager.contract.UnpackLog(event, "TreasuryWithdrawn", log); err != nil {
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

// ParseTreasuryWithdrawn is a log parse operation binding the contract event 0x3994603528a4ce7136b8ccd2cc94e047aa196fd8a1f735323028f632ad214dfb.
//
// Solidity: event TreasuryWithdrawn(uint256 indexed daoId, address to, address token, uint256 amount)
func (_DAOManager *DAOManagerFilterer) ParseTreasuryWithdrawn(log types.Log) (*DAOManagerTreasuryWithdrawn, error) {
	event := new(DAOManagerTreasuryWithdrawn)
	if err := _DAOManager.contract.UnpackLog(event, "TreasuryWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOManagerVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the DAOManager contract.
type DAOManagerVoteCastIterator struct {
	Event *DAOManagerVoteCast // Event containing the contract specifics and raw log

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
func (it *DAOManagerVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOManagerVoteCast)
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
		it.Event = new(DAOManagerVoteCast)
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
func (it *DAOManagerVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOManagerVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOManagerVoteCast represents a VoteCast event raised by the DAOManager contract.
type DAOManagerVoteCast struct {
	ProposalId *big.Int
	Voter      common.Address
	Support    bool
	Weight     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0xcbdf6214089cba887ecbf35a0b6a734589959c9763342c756bb2a80ca2bc9f6e.
//
// Solidity: event VoteCast(uint256 indexed proposalId, address indexed voter, bool support, uint256 weight)
func (_DAOManager *DAOManagerFilterer) FilterVoteCast(opts *bind.FilterOpts, proposalId []*big.Int, voter []common.Address) (*DAOManagerVoteCastIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _DAOManager.contract.FilterLogs(opts, "VoteCast", proposalIdRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &DAOManagerVoteCastIterator{contract: _DAOManager.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0xcbdf6214089cba887ecbf35a0b6a734589959c9763342c756bb2a80ca2bc9f6e.
//
// Solidity: event VoteCast(uint256 indexed proposalId, address indexed voter, bool support, uint256 weight)
func (_DAOManager *DAOManagerFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *DAOManagerVoteCast, proposalId []*big.Int, voter []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _DAOManager.contract.WatchLogs(opts, "VoteCast", proposalIdRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOManagerVoteCast)
				if err := _DAOManager.contract.UnpackLog(event, "VoteCast", log); err != nil {
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

// ParseVoteCast is a log parse operation binding the contract event 0xcbdf6214089cba887ecbf35a0b6a734589959c9763342c756bb2a80ca2bc9f6e.
//
// Solidity: event VoteCast(uint256 indexed proposalId, address indexed voter, bool support, uint256 weight)
func (_DAOManager *DAOManagerFilterer) ParseVoteCast(log types.Log) (*DAOManagerVoteCast, error) {
	event := new(DAOManagerVoteCast)
	if err := _DAOManager.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

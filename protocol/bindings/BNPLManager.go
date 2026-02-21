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

// BNPLManagerMetaData contains all meta data concerning the BNPLManager contract.
var BNPLManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BNPL_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arrangementId\",\"type\":\"uint256\"}],\"name\":\"activateBNPL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arrangementId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"installmentNumber\",\"type\":\"uint8\"}],\"name\":\"applyLateFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"daoId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"intervalSeconds\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"createBNPL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"arrangementId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arrangementId\",\"type\":\"uint256\"}],\"name\":\"getArrangement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"daoId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numInstallments\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"installmentAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"intervalSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lateFeeBps\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arrangementId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"installmentNumber\",\"type\":\"uint8\"}],\"name\":\"makePayment\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arrangementId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newStartTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newIntervalSeconds\",\"type\":\"uint256\"}],\"name\":\"reschedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_daoManager\",\"type\":\"address\"}],\"name\":\"setDaoManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"arrangementId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activatedAt\",\"type\":\"uint256\"}],\"name\":\"BNPLActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"arrangementId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"completedAt\",\"type\":\"uint256\"}],\"name\":\"BNPLCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"arrangementId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"daoId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numInstallments\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"name\":\"BNPLCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"arrangementId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"installmentNumber\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BNPLLateFeeApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"arrangementId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"installmentNumber\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BNPLPaymentMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"arrangementId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldScheduleHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newScheduleHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BNPLRescheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"arrangementId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"intervalSeconds\",\"type\":\"uint256\"}],\"name\":\"BNPLScheduleChosen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"}]",
	Bin: "0x60806040523460475760017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005560018055603733604b565b506040516113be90816100d48239f35b5f80fd5b6001600160a01b0381165f9081525f5160206114925f395f51905f52602052604090205460ff1660ce576001600160a01b03165f8181525f5160206114925f395f51905f5260205260408120805460ff191660011790553391907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b505f9056fe6080806040526004361015610012575f80fd5b5f905f3560e01c90816301ffc9a714611078575080630446be8a14610de657806317758c2814610c9857806318c60ed614610899578063248a9ca3146108665780632f2ff15d1461082857806336568abe146107e357806347f2c01d146106e957806356710d59146105b95780635a03d98d1461057e57806391d148541461053557806394785d53146104b85780639d820d891461015c578063a217fddf14610140578063d547741f146100f95763f322c755146100ce575f80fd5b346100f657806003193601126100f6576004546040516001600160a01b039091168152602090f35b80fd5b50346100f65760403660031901126100f65761013c6004356101196110eb565b90610137610132825f525f602052600160405f20015490565b611248565b611308565b5080f35b50346100f657806003193601126100f657602090604051908152f35b50610166366110cb565b60027f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0054146104a95760027f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055818352600260205260408320906101cc82541515611101565b60058201805460ff83169081101561046e576101eb8360068601611205565b90549060031b1c928334106104325786908682526003602052604082208383526020526040822061021d86825461122e565b9055600b8601948383528560205260408320600160ff19825416179055877fe8508bd3254b66027ff6d975978f9924a4472a58ad4da3ea70c8bead5ebad6c761028e610269843461123b565b94604051918291339642918460409194939260ff606083019616825260208201520152565b0390a380610399575b5050158061038a575b61034b575b54600191855b828110610322575b5050506102e3575b8260017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005580f35b600a01600260ff198254161790557fe8451c2098cc569ef00cac8bc592206c2ac1a28927a2b722c43fd1d958b631866020604051428152a25f806102bb565b8087528160205260ff6040882054161561033e576001016102ab565b50505050825f80806102b3565b600a8301600160ff19825416179055837fa85c968b84b91afcab116039a1596982b94cac2053615874fc98869fbc03739b6020604051428152a26102a5565b5060ff600a84015416156102a0565b81808092335af13d1561042d573d67ffffffffffffffff811161041957604051906103ce601f8201601f19166020018361117f565b81528760203d92013e5b156103e457855f610297565b60405162461bcd60e51b815260206004820152600d60248201526c14915195539117d19052531151609a1b6044820152606490fd5b634e487b7160e01b88526041600452602488fd5b6103d8565b60405162461bcd60e51b8152602060048201526014602482015273125394d551919250d251539517d410565351539560621b6044820152606490fd5b60405162461bcd60e51b81526020600482015260136024820152721253959053125117d25394d510531313515395606a1b6044820152606490fd5b633ee5aeb560e01b8352600483fd5b50346100f65760203660031901126100f6576004356001600160a01b038116908190036105315781805260208281526040808420335f908152925290205460ff1615610519576bffffffffffffffffffffffff60a01b600454161760045580f35b63e2517d3f60e01b8252336004526024829052604482fd5b5080fd5b50346100f65760403660031901126100f65760406105516110eb565b91600435815280602052209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b50346100f657806003193601126100f65760206040517fc5ec7611e8ff3fd4204773b752a273a89ef71fa4e42a8d3c0cd8a5412ee105578152f35b50346100f65760203660031901126100f65760043581526002602052604081209081549160018101549160018060a01b036002830154169160018060a01b0360038201541691600482015490600583015490600684019560078501549260088601549460ff600a60098901549801541697604051938460208c54918281520190819c88526020882090885b8181106106d3575050508561065a91038661117f565b6040519b6101608d019d8d5260208d015260408c015260608b015260808a015260a089015261016060c08901525180985261018087019590975b8089106106bb57505085965060e08601526101008501526101208401526101408301520390f35b90956020806001928951815201970198019790610694565b8254845260209093019260019283019201610644565b50346100f65760203660031901126100f65760043580825260026020526040822061071681541515611101565b600a81019060ff8254166107b057838052600b01602052604083205460ff161561077257600160ff198254161790557fa85c968b84b91afcab116039a1596982b94cac2053615874fc98869fbc03739b6020604051428152a280f35b60405162461bcd60e51b815260206004820152601660248201527511925494d517d410565351539517d49154555254915160521b6044820152606490fd5b60405162461bcd60e51b815260206004820152600b60248201526a4e4f545f50454e44494e4760a81b6044820152606490fd5b50346100f65760403660031901126100f6576107fd6110eb565b336001600160a01b038216036108195761013c90600435611308565b63334bd91960e11b8252600482fd5b50346100f65760403660031901126100f65761013c6004356108486110eb565b90610861610132825f525f602052600160405f20015490565b611280565b50346100f65760203660031901126100f65760206108916004355f525f602052600160405f20015490565b604051908152f35b50346100f65760c03660031901126100f6576004356108b66110eb565b60443590606435926084359260a43567ffffffffffffffff8111610c945736602382011215610c9457806004013567ffffffffffffffff8111610c905736910160240111610c8c57600454949560249560e0906001600160a01b031661091d81151561113d565b6040519788809263d3c095a3851b82528760048301525afa908115610c805780968182918394610c42575b508815610bfd5762015180810290808204620151801490151715610be9578710610baf5762015180810290808204620151801490151715610b62578611610b7657600154965f198814610b62576001888101815588835260026020819052604084208a8155918201879055810180546001600160a01b031990811633179091556003820180549091166001600160a01b0389161790556004810185905560058101828155600782018b9055600882018990556009820194909455600a8101805460ff191690558185046006610a26610a2085846111f2565b8861123b565b920191825485845580610b3d575b50845b848110610ac357505060208a807f683528e54d1d686403cfb3fde979f90e4d361cd6dfda99bae38fd3b9b9dab45160408f8e8e8e8e8e549086519360018060a01b031684528a84015285830152426060830152867f40b3e5fc479cc74229507a38ae0bd501055d6f96698e28e1c3de987b2b851bae60803394a4825191825286820152a2604051908152f35b5f198501858111610b29578103610b2057610adf825b8461122e565b8454680100000000000000008110156104195790610b0582600180959401885587611205565b819291549060031b91821b915f19901b191617905501610a37565b610adf86610ad9565b634e487b7160e01b87526011600452602487fd5b83865260208620865b828110610b54575050610a34565b808860019284015501610b46565b634e487b7160e01b82526011600452602482fd5b60405162461bcd60e51b8152602060048201526011602482015270494e54455256414c5f544f4f5f4c4f4e4760781b6044820152606490fd5b60405162461bcd60e51b8152602060048201526012602482015271125395115495905317d513d3d7d4d213d49560721b6044820152606490fd5b634e487b7160e01b83526011600452602483fd5b60405162461bcd60e51b815260206004820152601760248201527f44414f5f424e504c5f4e4f545f434f4e464947555245440000000000000000006044820152606490fd5b9298505050610c69915060e03d60e011610c79575b610c61818361117f565b8101906111b5565b505050909297919290925f610948565b503d610c57565b604051903d90823e3d90fd5b8580fd5b8780fd5b8680fd5b5034610dd957610ca7366110cb565b815f52600260205260405f20610cbf81541515611101565b612710610ce3610cd28460068501611205565b905460098501549160031b1c6111f2565b049060048101610cf483825461122e565b90556004546001600160a01b03169081151580610ddd575b610d55575b50506040805160ff939093168352602083019190915242908201527fa771582d9eca55eb62b5ff82bfe51f20afbdb13ab06f2c7ed88ca01c80e322e890606090a280f35b60010154813b15610dd9575f9160448392604051948593849263cec4aec960e01b845260048401528760248401525af18015610dce57610d96575b80610d11565b7fa771582d9eca55eb62b5ff82bfe51f20afbdb13ab06f2c7ed88ca01c80e322e8929194505f610dc59161117f565b5f939091610d90565b6040513d5f823e3d90fd5b5f80fd5b50821515610d0c565b34610dd9576060366003190112610dd957600435602435604435825f52600260205260405f20610e1881541515611101565b6004546001600160a01b0316610e2f81151561113d565b60e0600183015460246040518094819363d3c095a3861b835260048301525afa8015610dce575f905f925f91611050575b50156110125760028301546001600160a01b031633148015610fdb575b15610fa75762015180810290808204620151801490151715610f93578310159081610f75575b5015610f37577f3986b9d87f88ca7357a73fd27b9b6b2d54915820b5c9dcb7cadf98ec84e391aa92826060936007840193836008865492019182546040519060208201928352604082015260408152610efc898261117f565b5190209555556040519060208201928352604082015260408152610f20848261117f565b5190206040519182526020820152426040820152a2005b60405162461bcd60e51b8152602060048201526016602482015275494e54455256414c5f4f55545f4f465f424f554e445360501b6044820152606490fd5b6201518080820292508115918304141715610f935782111585610ea3565b634e487b7160e01b5f52601160045260245ffd5b60405162461bcd60e51b815260206004820152600c60248201526b15539055551213d49256915160a21b6044820152606490fd5b50335f9081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16610e7d565b60405162461bcd60e51b8152602060048201526016602482015275149154d0d2115115531157d393d517d0531313d5d15160521b6044820152606490fd5b91505061106c915060e03d60e011610c7957610c61818361117f565b50945050509187610e60565b34610dd9576020366003190112610dd9576004359063ffffffff60e01b8216809203610dd957602091637965db0b60e01b81149081156110ba575b5015158152f35b6301ffc9a760e01b149050836110b3565b6040906003190112610dd9576004359060243560ff81168103610dd95790565b602435906001600160a01b0382168203610dd957565b1561110857565b60405162461bcd60e51b815260206004820152600d60248201526c10549497d393d517d193d55391609a1b6044820152606490fd5b1561114457565b60405162461bcd60e51b8152602060048201526013602482015272111053d7d350539051d15497d393d517d4d155606a1b6044820152606490fd5b90601f8019910116810190811067ffffffffffffffff8211176111a157604052565b634e487b7160e01b5f52604160045260245ffd5b908160e0910312610dd95780519160208201519160408101519160608201519160808101519160a08201518015158103610dd95760c09092015190565b81810292918115918404141715610f9357565b805482101561121a575f5260205f2001905f90565b634e487b7160e01b5f52603260045260245ffd5b91908201809211610f9357565b91908203918211610f9357565b5f8181526020818152604080832033845290915290205460ff161561126a5750565b63e2517d3f60e01b5f523360045260245260445ffd5b5f818152602081815260408083206001600160a01b038616845290915290205460ff16611302575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50505f90565b5f818152602081815260408083206001600160a01b038616845290915290205460ff1615611302575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a460019056fea2646970667358221220d89b98924f8a115d838f0d0a9657ffa8d98e141de7991d5800d869b178a2903864736f6c63430008210033ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
}

// BNPLManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use BNPLManagerMetaData.ABI instead.
var BNPLManagerABI = BNPLManagerMetaData.ABI

// BNPLManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BNPLManagerMetaData.Bin instead.
var BNPLManagerBin = BNPLManagerMetaData.Bin

// DeployBNPLManager deploys a new Ethereum contract, binding an instance of BNPLManager to it.
func DeployBNPLManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BNPLManager, error) {
	parsed, err := BNPLManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BNPLManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BNPLManager{BNPLManagerCaller: BNPLManagerCaller{contract: contract}, BNPLManagerTransactor: BNPLManagerTransactor{contract: contract}, BNPLManagerFilterer: BNPLManagerFilterer{contract: contract}}, nil
}

// BNPLManager is an auto generated Go binding around an Ethereum contract.
type BNPLManager struct {
	BNPLManagerCaller     // Read-only binding to the contract
	BNPLManagerTransactor // Write-only binding to the contract
	BNPLManagerFilterer   // Log filterer for contract events
}

// BNPLManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BNPLManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BNPLManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BNPLManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BNPLManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BNPLManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BNPLManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BNPLManagerSession struct {
	Contract     *BNPLManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BNPLManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BNPLManagerCallerSession struct {
	Contract *BNPLManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BNPLManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BNPLManagerTransactorSession struct {
	Contract     *BNPLManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BNPLManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BNPLManagerRaw struct {
	Contract *BNPLManager // Generic contract binding to access the raw methods on
}

// BNPLManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BNPLManagerCallerRaw struct {
	Contract *BNPLManagerCaller // Generic read-only contract binding to access the raw methods on
}

// BNPLManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BNPLManagerTransactorRaw struct {
	Contract *BNPLManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBNPLManager creates a new instance of BNPLManager, bound to a specific deployed contract.
func NewBNPLManager(address common.Address, backend bind.ContractBackend) (*BNPLManager, error) {
	contract, err := bindBNPLManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BNPLManager{BNPLManagerCaller: BNPLManagerCaller{contract: contract}, BNPLManagerTransactor: BNPLManagerTransactor{contract: contract}, BNPLManagerFilterer: BNPLManagerFilterer{contract: contract}}, nil
}

// NewBNPLManagerCaller creates a new read-only instance of BNPLManager, bound to a specific deployed contract.
func NewBNPLManagerCaller(address common.Address, caller bind.ContractCaller) (*BNPLManagerCaller, error) {
	contract, err := bindBNPLManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BNPLManagerCaller{contract: contract}, nil
}

// NewBNPLManagerTransactor creates a new write-only instance of BNPLManager, bound to a specific deployed contract.
func NewBNPLManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*BNPLManagerTransactor, error) {
	contract, err := bindBNPLManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BNPLManagerTransactor{contract: contract}, nil
}

// NewBNPLManagerFilterer creates a new log filterer instance of BNPLManager, bound to a specific deployed contract.
func NewBNPLManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*BNPLManagerFilterer, error) {
	contract, err := bindBNPLManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BNPLManagerFilterer{contract: contract}, nil
}

// bindBNPLManager binds a generic wrapper to an already deployed contract.
func bindBNPLManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BNPLManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BNPLManager *BNPLManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BNPLManager.Contract.BNPLManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BNPLManager *BNPLManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BNPLManager.Contract.BNPLManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BNPLManager *BNPLManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BNPLManager.Contract.BNPLManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BNPLManager *BNPLManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BNPLManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BNPLManager *BNPLManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BNPLManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BNPLManager *BNPLManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BNPLManager.Contract.contract.Transact(opts, method, params...)
}

// BNPLADMINROLE is a free data retrieval call binding the contract method 0x5a03d98d.
//
// Solidity: function BNPL_ADMIN_ROLE() view returns(bytes32)
func (_BNPLManager *BNPLManagerCaller) BNPLADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BNPLManager.contract.Call(opts, &out, "BNPL_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BNPLADMINROLE is a free data retrieval call binding the contract method 0x5a03d98d.
//
// Solidity: function BNPL_ADMIN_ROLE() view returns(bytes32)
func (_BNPLManager *BNPLManagerSession) BNPLADMINROLE() ([32]byte, error) {
	return _BNPLManager.Contract.BNPLADMINROLE(&_BNPLManager.CallOpts)
}

// BNPLADMINROLE is a free data retrieval call binding the contract method 0x5a03d98d.
//
// Solidity: function BNPL_ADMIN_ROLE() view returns(bytes32)
func (_BNPLManager *BNPLManagerCallerSession) BNPLADMINROLE() ([32]byte, error) {
	return _BNPLManager.Contract.BNPLADMINROLE(&_BNPLManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BNPLManager *BNPLManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BNPLManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BNPLManager *BNPLManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BNPLManager.Contract.DEFAULTADMINROLE(&_BNPLManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BNPLManager *BNPLManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BNPLManager.Contract.DEFAULTADMINROLE(&_BNPLManager.CallOpts)
}

// DaoManager is a free data retrieval call binding the contract method 0xf322c755.
//
// Solidity: function daoManager() view returns(address)
func (_BNPLManager *BNPLManagerCaller) DaoManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BNPLManager.contract.Call(opts, &out, "daoManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaoManager is a free data retrieval call binding the contract method 0xf322c755.
//
// Solidity: function daoManager() view returns(address)
func (_BNPLManager *BNPLManagerSession) DaoManager() (common.Address, error) {
	return _BNPLManager.Contract.DaoManager(&_BNPLManager.CallOpts)
}

// DaoManager is a free data retrieval call binding the contract method 0xf322c755.
//
// Solidity: function daoManager() view returns(address)
func (_BNPLManager *BNPLManagerCallerSession) DaoManager() (common.Address, error) {
	return _BNPLManager.Contract.DaoManager(&_BNPLManager.CallOpts)
}

// GetArrangement is a free data retrieval call binding the contract method 0x56710d59.
//
// Solidity: function getArrangement(uint256 arrangementId) view returns(uint256 id, uint256 daoId, address payer, address recipient, uint256 totalAmount, uint256 numInstallments, uint256[] installmentAmounts, uint256 startTimestamp, uint256 intervalSeconds, uint256 lateFeeBps, uint8 status)
func (_BNPLManager *BNPLManagerCaller) GetArrangement(opts *bind.CallOpts, arrangementId *big.Int) (struct {
	Id                 *big.Int
	DaoId              *big.Int
	Payer              common.Address
	Recipient          common.Address
	TotalAmount        *big.Int
	NumInstallments    *big.Int
	InstallmentAmounts []*big.Int
	StartTimestamp     *big.Int
	IntervalSeconds    *big.Int
	LateFeeBps         *big.Int
	Status             uint8
}, error) {
	var out []interface{}
	err := _BNPLManager.contract.Call(opts, &out, "getArrangement", arrangementId)

	outstruct := new(struct {
		Id                 *big.Int
		DaoId              *big.Int
		Payer              common.Address
		Recipient          common.Address
		TotalAmount        *big.Int
		NumInstallments    *big.Int
		InstallmentAmounts []*big.Int
		StartTimestamp     *big.Int
		IntervalSeconds    *big.Int
		LateFeeBps         *big.Int
		Status             uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DaoId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Payer = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Recipient = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.TotalAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.NumInstallments = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.InstallmentAmounts = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)
	outstruct.StartTimestamp = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.IntervalSeconds = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.LateFeeBps = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[10], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetArrangement is a free data retrieval call binding the contract method 0x56710d59.
//
// Solidity: function getArrangement(uint256 arrangementId) view returns(uint256 id, uint256 daoId, address payer, address recipient, uint256 totalAmount, uint256 numInstallments, uint256[] installmentAmounts, uint256 startTimestamp, uint256 intervalSeconds, uint256 lateFeeBps, uint8 status)
func (_BNPLManager *BNPLManagerSession) GetArrangement(arrangementId *big.Int) (struct {
	Id                 *big.Int
	DaoId              *big.Int
	Payer              common.Address
	Recipient          common.Address
	TotalAmount        *big.Int
	NumInstallments    *big.Int
	InstallmentAmounts []*big.Int
	StartTimestamp     *big.Int
	IntervalSeconds    *big.Int
	LateFeeBps         *big.Int
	Status             uint8
}, error) {
	return _BNPLManager.Contract.GetArrangement(&_BNPLManager.CallOpts, arrangementId)
}

// GetArrangement is a free data retrieval call binding the contract method 0x56710d59.
//
// Solidity: function getArrangement(uint256 arrangementId) view returns(uint256 id, uint256 daoId, address payer, address recipient, uint256 totalAmount, uint256 numInstallments, uint256[] installmentAmounts, uint256 startTimestamp, uint256 intervalSeconds, uint256 lateFeeBps, uint8 status)
func (_BNPLManager *BNPLManagerCallerSession) GetArrangement(arrangementId *big.Int) (struct {
	Id                 *big.Int
	DaoId              *big.Int
	Payer              common.Address
	Recipient          common.Address
	TotalAmount        *big.Int
	NumInstallments    *big.Int
	InstallmentAmounts []*big.Int
	StartTimestamp     *big.Int
	IntervalSeconds    *big.Int
	LateFeeBps         *big.Int
	Status             uint8
}, error) {
	return _BNPLManager.Contract.GetArrangement(&_BNPLManager.CallOpts, arrangementId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BNPLManager *BNPLManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BNPLManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BNPLManager *BNPLManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BNPLManager.Contract.GetRoleAdmin(&_BNPLManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BNPLManager *BNPLManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BNPLManager.Contract.GetRoleAdmin(&_BNPLManager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BNPLManager *BNPLManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BNPLManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BNPLManager *BNPLManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BNPLManager.Contract.HasRole(&_BNPLManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BNPLManager *BNPLManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BNPLManager.Contract.HasRole(&_BNPLManager.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BNPLManager *BNPLManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BNPLManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BNPLManager *BNPLManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BNPLManager.Contract.SupportsInterface(&_BNPLManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BNPLManager *BNPLManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BNPLManager.Contract.SupportsInterface(&_BNPLManager.CallOpts, interfaceId)
}

// ActivateBNPL is a paid mutator transaction binding the contract method 0x47f2c01d.
//
// Solidity: function activateBNPL(uint256 arrangementId) returns()
func (_BNPLManager *BNPLManagerTransactor) ActivateBNPL(opts *bind.TransactOpts, arrangementId *big.Int) (*types.Transaction, error) {
	return _BNPLManager.contract.Transact(opts, "activateBNPL", arrangementId)
}

// ActivateBNPL is a paid mutator transaction binding the contract method 0x47f2c01d.
//
// Solidity: function activateBNPL(uint256 arrangementId) returns()
func (_BNPLManager *BNPLManagerSession) ActivateBNPL(arrangementId *big.Int) (*types.Transaction, error) {
	return _BNPLManager.Contract.ActivateBNPL(&_BNPLManager.TransactOpts, arrangementId)
}

// ActivateBNPL is a paid mutator transaction binding the contract method 0x47f2c01d.
//
// Solidity: function activateBNPL(uint256 arrangementId) returns()
func (_BNPLManager *BNPLManagerTransactorSession) ActivateBNPL(arrangementId *big.Int) (*types.Transaction, error) {
	return _BNPLManager.Contract.ActivateBNPL(&_BNPLManager.TransactOpts, arrangementId)
}

// ApplyLateFee is a paid mutator transaction binding the contract method 0x17758c28.
//
// Solidity: function applyLateFee(uint256 arrangementId, uint8 installmentNumber) returns()
func (_BNPLManager *BNPLManagerTransactor) ApplyLateFee(opts *bind.TransactOpts, arrangementId *big.Int, installmentNumber uint8) (*types.Transaction, error) {
	return _BNPLManager.contract.Transact(opts, "applyLateFee", arrangementId, installmentNumber)
}

// ApplyLateFee is a paid mutator transaction binding the contract method 0x17758c28.
//
// Solidity: function applyLateFee(uint256 arrangementId, uint8 installmentNumber) returns()
func (_BNPLManager *BNPLManagerSession) ApplyLateFee(arrangementId *big.Int, installmentNumber uint8) (*types.Transaction, error) {
	return _BNPLManager.Contract.ApplyLateFee(&_BNPLManager.TransactOpts, arrangementId, installmentNumber)
}

// ApplyLateFee is a paid mutator transaction binding the contract method 0x17758c28.
//
// Solidity: function applyLateFee(uint256 arrangementId, uint8 installmentNumber) returns()
func (_BNPLManager *BNPLManagerTransactorSession) ApplyLateFee(arrangementId *big.Int, installmentNumber uint8) (*types.Transaction, error) {
	return _BNPLManager.Contract.ApplyLateFee(&_BNPLManager.TransactOpts, arrangementId, installmentNumber)
}

// CreateBNPL is a paid mutator transaction binding the contract method 0x18c60ed6.
//
// Solidity: function createBNPL(uint256 daoId, address recipient, uint256 totalAmount, uint256 startTimestamp, uint256 intervalSeconds, bytes ) returns(uint256 arrangementId)
func (_BNPLManager *BNPLManagerTransactor) CreateBNPL(opts *bind.TransactOpts, daoId *big.Int, recipient common.Address, totalAmount *big.Int, startTimestamp *big.Int, intervalSeconds *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _BNPLManager.contract.Transact(opts, "createBNPL", daoId, recipient, totalAmount, startTimestamp, intervalSeconds, arg5)
}

// CreateBNPL is a paid mutator transaction binding the contract method 0x18c60ed6.
//
// Solidity: function createBNPL(uint256 daoId, address recipient, uint256 totalAmount, uint256 startTimestamp, uint256 intervalSeconds, bytes ) returns(uint256 arrangementId)
func (_BNPLManager *BNPLManagerSession) CreateBNPL(daoId *big.Int, recipient common.Address, totalAmount *big.Int, startTimestamp *big.Int, intervalSeconds *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _BNPLManager.Contract.CreateBNPL(&_BNPLManager.TransactOpts, daoId, recipient, totalAmount, startTimestamp, intervalSeconds, arg5)
}

// CreateBNPL is a paid mutator transaction binding the contract method 0x18c60ed6.
//
// Solidity: function createBNPL(uint256 daoId, address recipient, uint256 totalAmount, uint256 startTimestamp, uint256 intervalSeconds, bytes ) returns(uint256 arrangementId)
func (_BNPLManager *BNPLManagerTransactorSession) CreateBNPL(daoId *big.Int, recipient common.Address, totalAmount *big.Int, startTimestamp *big.Int, intervalSeconds *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _BNPLManager.Contract.CreateBNPL(&_BNPLManager.TransactOpts, daoId, recipient, totalAmount, startTimestamp, intervalSeconds, arg5)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BNPLManager *BNPLManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BNPLManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BNPLManager *BNPLManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BNPLManager.Contract.GrantRole(&_BNPLManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BNPLManager *BNPLManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BNPLManager.Contract.GrantRole(&_BNPLManager.TransactOpts, role, account)
}

// MakePayment is a paid mutator transaction binding the contract method 0x9d820d89.
//
// Solidity: function makePayment(uint256 arrangementId, uint8 installmentNumber) payable returns()
func (_BNPLManager *BNPLManagerTransactor) MakePayment(opts *bind.TransactOpts, arrangementId *big.Int, installmentNumber uint8) (*types.Transaction, error) {
	return _BNPLManager.contract.Transact(opts, "makePayment", arrangementId, installmentNumber)
}

// MakePayment is a paid mutator transaction binding the contract method 0x9d820d89.
//
// Solidity: function makePayment(uint256 arrangementId, uint8 installmentNumber) payable returns()
func (_BNPLManager *BNPLManagerSession) MakePayment(arrangementId *big.Int, installmentNumber uint8) (*types.Transaction, error) {
	return _BNPLManager.Contract.MakePayment(&_BNPLManager.TransactOpts, arrangementId, installmentNumber)
}

// MakePayment is a paid mutator transaction binding the contract method 0x9d820d89.
//
// Solidity: function makePayment(uint256 arrangementId, uint8 installmentNumber) payable returns()
func (_BNPLManager *BNPLManagerTransactorSession) MakePayment(arrangementId *big.Int, installmentNumber uint8) (*types.Transaction, error) {
	return _BNPLManager.Contract.MakePayment(&_BNPLManager.TransactOpts, arrangementId, installmentNumber)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BNPLManager *BNPLManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BNPLManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BNPLManager *BNPLManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BNPLManager.Contract.RenounceRole(&_BNPLManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BNPLManager *BNPLManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BNPLManager.Contract.RenounceRole(&_BNPLManager.TransactOpts, role, callerConfirmation)
}

// Reschedule is a paid mutator transaction binding the contract method 0x0446be8a.
//
// Solidity: function reschedule(uint256 arrangementId, uint256 newStartTimestamp, uint256 newIntervalSeconds) returns()
func (_BNPLManager *BNPLManagerTransactor) Reschedule(opts *bind.TransactOpts, arrangementId *big.Int, newStartTimestamp *big.Int, newIntervalSeconds *big.Int) (*types.Transaction, error) {
	return _BNPLManager.contract.Transact(opts, "reschedule", arrangementId, newStartTimestamp, newIntervalSeconds)
}

// Reschedule is a paid mutator transaction binding the contract method 0x0446be8a.
//
// Solidity: function reschedule(uint256 arrangementId, uint256 newStartTimestamp, uint256 newIntervalSeconds) returns()
func (_BNPLManager *BNPLManagerSession) Reschedule(arrangementId *big.Int, newStartTimestamp *big.Int, newIntervalSeconds *big.Int) (*types.Transaction, error) {
	return _BNPLManager.Contract.Reschedule(&_BNPLManager.TransactOpts, arrangementId, newStartTimestamp, newIntervalSeconds)
}

// Reschedule is a paid mutator transaction binding the contract method 0x0446be8a.
//
// Solidity: function reschedule(uint256 arrangementId, uint256 newStartTimestamp, uint256 newIntervalSeconds) returns()
func (_BNPLManager *BNPLManagerTransactorSession) Reschedule(arrangementId *big.Int, newStartTimestamp *big.Int, newIntervalSeconds *big.Int) (*types.Transaction, error) {
	return _BNPLManager.Contract.Reschedule(&_BNPLManager.TransactOpts, arrangementId, newStartTimestamp, newIntervalSeconds)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BNPLManager *BNPLManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BNPLManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BNPLManager *BNPLManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BNPLManager.Contract.RevokeRole(&_BNPLManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BNPLManager *BNPLManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BNPLManager.Contract.RevokeRole(&_BNPLManager.TransactOpts, role, account)
}

// SetDaoManager is a paid mutator transaction binding the contract method 0x94785d53.
//
// Solidity: function setDaoManager(address _daoManager) returns()
func (_BNPLManager *BNPLManagerTransactor) SetDaoManager(opts *bind.TransactOpts, _daoManager common.Address) (*types.Transaction, error) {
	return _BNPLManager.contract.Transact(opts, "setDaoManager", _daoManager)
}

// SetDaoManager is a paid mutator transaction binding the contract method 0x94785d53.
//
// Solidity: function setDaoManager(address _daoManager) returns()
func (_BNPLManager *BNPLManagerSession) SetDaoManager(_daoManager common.Address) (*types.Transaction, error) {
	return _BNPLManager.Contract.SetDaoManager(&_BNPLManager.TransactOpts, _daoManager)
}

// SetDaoManager is a paid mutator transaction binding the contract method 0x94785d53.
//
// Solidity: function setDaoManager(address _daoManager) returns()
func (_BNPLManager *BNPLManagerTransactorSession) SetDaoManager(_daoManager common.Address) (*types.Transaction, error) {
	return _BNPLManager.Contract.SetDaoManager(&_BNPLManager.TransactOpts, _daoManager)
}

// BNPLManagerBNPLActivatedIterator is returned from FilterBNPLActivated and is used to iterate over the raw logs and unpacked data for BNPLActivated events raised by the BNPLManager contract.
type BNPLManagerBNPLActivatedIterator struct {
	Event *BNPLManagerBNPLActivated // Event containing the contract specifics and raw log

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
func (it *BNPLManagerBNPLActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BNPLManagerBNPLActivated)
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
		it.Event = new(BNPLManagerBNPLActivated)
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
func (it *BNPLManagerBNPLActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BNPLManagerBNPLActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BNPLManagerBNPLActivated represents a BNPLActivated event raised by the BNPLManager contract.
type BNPLManagerBNPLActivated struct {
	ArrangementId *big.Int
	ActivatedAt   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBNPLActivated is a free log retrieval operation binding the contract event 0xa85c968b84b91afcab116039a1596982b94cac2053615874fc98869fbc03739b.
//
// Solidity: event BNPLActivated(uint256 indexed arrangementId, uint256 activatedAt)
func (_BNPLManager *BNPLManagerFilterer) FilterBNPLActivated(opts *bind.FilterOpts, arrangementId []*big.Int) (*BNPLManagerBNPLActivatedIterator, error) {

	var arrangementIdRule []interface{}
	for _, arrangementIdItem := range arrangementId {
		arrangementIdRule = append(arrangementIdRule, arrangementIdItem)
	}

	logs, sub, err := _BNPLManager.contract.FilterLogs(opts, "BNPLActivated", arrangementIdRule)
	if err != nil {
		return nil, err
	}
	return &BNPLManagerBNPLActivatedIterator{contract: _BNPLManager.contract, event: "BNPLActivated", logs: logs, sub: sub}, nil
}

// WatchBNPLActivated is a free log subscription operation binding the contract event 0xa85c968b84b91afcab116039a1596982b94cac2053615874fc98869fbc03739b.
//
// Solidity: event BNPLActivated(uint256 indexed arrangementId, uint256 activatedAt)
func (_BNPLManager *BNPLManagerFilterer) WatchBNPLActivated(opts *bind.WatchOpts, sink chan<- *BNPLManagerBNPLActivated, arrangementId []*big.Int) (event.Subscription, error) {

	var arrangementIdRule []interface{}
	for _, arrangementIdItem := range arrangementId {
		arrangementIdRule = append(arrangementIdRule, arrangementIdItem)
	}

	logs, sub, err := _BNPLManager.contract.WatchLogs(opts, "BNPLActivated", arrangementIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BNPLManagerBNPLActivated)
				if err := _BNPLManager.contract.UnpackLog(event, "BNPLActivated", log); err != nil {
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

// ParseBNPLActivated is a log parse operation binding the contract event 0xa85c968b84b91afcab116039a1596982b94cac2053615874fc98869fbc03739b.
//
// Solidity: event BNPLActivated(uint256 indexed arrangementId, uint256 activatedAt)
func (_BNPLManager *BNPLManagerFilterer) ParseBNPLActivated(log types.Log) (*BNPLManagerBNPLActivated, error) {
	event := new(BNPLManagerBNPLActivated)
	if err := _BNPLManager.contract.UnpackLog(event, "BNPLActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BNPLManagerBNPLCompletedIterator is returned from FilterBNPLCompleted and is used to iterate over the raw logs and unpacked data for BNPLCompleted events raised by the BNPLManager contract.
type BNPLManagerBNPLCompletedIterator struct {
	Event *BNPLManagerBNPLCompleted // Event containing the contract specifics and raw log

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
func (it *BNPLManagerBNPLCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BNPLManagerBNPLCompleted)
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
		it.Event = new(BNPLManagerBNPLCompleted)
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
func (it *BNPLManagerBNPLCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BNPLManagerBNPLCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BNPLManagerBNPLCompleted represents a BNPLCompleted event raised by the BNPLManager contract.
type BNPLManagerBNPLCompleted struct {
	ArrangementId *big.Int
	CompletedAt   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBNPLCompleted is a free log retrieval operation binding the contract event 0xe8451c2098cc569ef00cac8bc592206c2ac1a28927a2b722c43fd1d958b63186.
//
// Solidity: event BNPLCompleted(uint256 indexed arrangementId, uint256 completedAt)
func (_BNPLManager *BNPLManagerFilterer) FilterBNPLCompleted(opts *bind.FilterOpts, arrangementId []*big.Int) (*BNPLManagerBNPLCompletedIterator, error) {

	var arrangementIdRule []interface{}
	for _, arrangementIdItem := range arrangementId {
		arrangementIdRule = append(arrangementIdRule, arrangementIdItem)
	}

	logs, sub, err := _BNPLManager.contract.FilterLogs(opts, "BNPLCompleted", arrangementIdRule)
	if err != nil {
		return nil, err
	}
	return &BNPLManagerBNPLCompletedIterator{contract: _BNPLManager.contract, event: "BNPLCompleted", logs: logs, sub: sub}, nil
}

// WatchBNPLCompleted is a free log subscription operation binding the contract event 0xe8451c2098cc569ef00cac8bc592206c2ac1a28927a2b722c43fd1d958b63186.
//
// Solidity: event BNPLCompleted(uint256 indexed arrangementId, uint256 completedAt)
func (_BNPLManager *BNPLManagerFilterer) WatchBNPLCompleted(opts *bind.WatchOpts, sink chan<- *BNPLManagerBNPLCompleted, arrangementId []*big.Int) (event.Subscription, error) {

	var arrangementIdRule []interface{}
	for _, arrangementIdItem := range arrangementId {
		arrangementIdRule = append(arrangementIdRule, arrangementIdItem)
	}

	logs, sub, err := _BNPLManager.contract.WatchLogs(opts, "BNPLCompleted", arrangementIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BNPLManagerBNPLCompleted)
				if err := _BNPLManager.contract.UnpackLog(event, "BNPLCompleted", log); err != nil {
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

// ParseBNPLCompleted is a log parse operation binding the contract event 0xe8451c2098cc569ef00cac8bc592206c2ac1a28927a2b722c43fd1d958b63186.
//
// Solidity: event BNPLCompleted(uint256 indexed arrangementId, uint256 completedAt)
func (_BNPLManager *BNPLManagerFilterer) ParseBNPLCompleted(log types.Log) (*BNPLManagerBNPLCompleted, error) {
	event := new(BNPLManagerBNPLCompleted)
	if err := _BNPLManager.contract.UnpackLog(event, "BNPLCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BNPLManagerBNPLCreatedIterator is returned from FilterBNPLCreated and is used to iterate over the raw logs and unpacked data for BNPLCreated events raised by the BNPLManager contract.
type BNPLManagerBNPLCreatedIterator struct {
	Event *BNPLManagerBNPLCreated // Event containing the contract specifics and raw log

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
func (it *BNPLManagerBNPLCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BNPLManagerBNPLCreated)
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
		it.Event = new(BNPLManagerBNPLCreated)
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
func (it *BNPLManagerBNPLCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BNPLManagerBNPLCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BNPLManagerBNPLCreated represents a BNPLCreated event raised by the BNPLManager contract.
type BNPLManagerBNPLCreated struct {
	ArrangementId   *big.Int
	DaoId           *big.Int
	Payer           common.Address
	Recipient       common.Address
	TotalAmount     *big.Int
	NumInstallments *big.Int
	CreatedAt       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBNPLCreated is a free log retrieval operation binding the contract event 0x40b3e5fc479cc74229507a38ae0bd501055d6f96698e28e1c3de987b2b851bae.
//
// Solidity: event BNPLCreated(uint256 indexed arrangementId, uint256 indexed daoId, address indexed payer, address recipient, uint256 totalAmount, uint256 numInstallments, uint256 createdAt)
func (_BNPLManager *BNPLManagerFilterer) FilterBNPLCreated(opts *bind.FilterOpts, arrangementId []*big.Int, daoId []*big.Int, payer []common.Address) (*BNPLManagerBNPLCreatedIterator, error) {

	var arrangementIdRule []interface{}
	for _, arrangementIdItem := range arrangementId {
		arrangementIdRule = append(arrangementIdRule, arrangementIdItem)
	}
	var daoIdRule []interface{}
	for _, daoIdItem := range daoId {
		daoIdRule = append(daoIdRule, daoIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _BNPLManager.contract.FilterLogs(opts, "BNPLCreated", arrangementIdRule, daoIdRule, payerRule)
	if err != nil {
		return nil, err
	}
	return &BNPLManagerBNPLCreatedIterator{contract: _BNPLManager.contract, event: "BNPLCreated", logs: logs, sub: sub}, nil
}

// WatchBNPLCreated is a free log subscription operation binding the contract event 0x40b3e5fc479cc74229507a38ae0bd501055d6f96698e28e1c3de987b2b851bae.
//
// Solidity: event BNPLCreated(uint256 indexed arrangementId, uint256 indexed daoId, address indexed payer, address recipient, uint256 totalAmount, uint256 numInstallments, uint256 createdAt)
func (_BNPLManager *BNPLManagerFilterer) WatchBNPLCreated(opts *bind.WatchOpts, sink chan<- *BNPLManagerBNPLCreated, arrangementId []*big.Int, daoId []*big.Int, payer []common.Address) (event.Subscription, error) {

	var arrangementIdRule []interface{}
	for _, arrangementIdItem := range arrangementId {
		arrangementIdRule = append(arrangementIdRule, arrangementIdItem)
	}
	var daoIdRule []interface{}
	for _, daoIdItem := range daoId {
		daoIdRule = append(daoIdRule, daoIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _BNPLManager.contract.WatchLogs(opts, "BNPLCreated", arrangementIdRule, daoIdRule, payerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BNPLManagerBNPLCreated)
				if err := _BNPLManager.contract.UnpackLog(event, "BNPLCreated", log); err != nil {
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

// ParseBNPLCreated is a log parse operation binding the contract event 0x40b3e5fc479cc74229507a38ae0bd501055d6f96698e28e1c3de987b2b851bae.
//
// Solidity: event BNPLCreated(uint256 indexed arrangementId, uint256 indexed daoId, address indexed payer, address recipient, uint256 totalAmount, uint256 numInstallments, uint256 createdAt)
func (_BNPLManager *BNPLManagerFilterer) ParseBNPLCreated(log types.Log) (*BNPLManagerBNPLCreated, error) {
	event := new(BNPLManagerBNPLCreated)
	if err := _BNPLManager.contract.UnpackLog(event, "BNPLCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BNPLManagerBNPLLateFeeAppliedIterator is returned from FilterBNPLLateFeeApplied and is used to iterate over the raw logs and unpacked data for BNPLLateFeeApplied events raised by the BNPLManager contract.
type BNPLManagerBNPLLateFeeAppliedIterator struct {
	Event *BNPLManagerBNPLLateFeeApplied // Event containing the contract specifics and raw log

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
func (it *BNPLManagerBNPLLateFeeAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BNPLManagerBNPLLateFeeApplied)
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
		it.Event = new(BNPLManagerBNPLLateFeeApplied)
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
func (it *BNPLManagerBNPLLateFeeAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BNPLManagerBNPLLateFeeAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BNPLManagerBNPLLateFeeApplied represents a BNPLLateFeeApplied event raised by the BNPLManager contract.
type BNPLManagerBNPLLateFeeApplied struct {
	ArrangementId     *big.Int
	InstallmentNumber uint8
	FeeAmount         *big.Int
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBNPLLateFeeApplied is a free log retrieval operation binding the contract event 0xa771582d9eca55eb62b5ff82bfe51f20afbdb13ab06f2c7ed88ca01c80e322e8.
//
// Solidity: event BNPLLateFeeApplied(uint256 indexed arrangementId, uint8 installmentNumber, uint256 feeAmount, uint256 timestamp)
func (_BNPLManager *BNPLManagerFilterer) FilterBNPLLateFeeApplied(opts *bind.FilterOpts, arrangementId []*big.Int) (*BNPLManagerBNPLLateFeeAppliedIterator, error) {

	var arrangementIdRule []interface{}
	for _, arrangementIdItem := range arrangementId {
		arrangementIdRule = append(arrangementIdRule, arrangementIdItem)
	}

	logs, sub, err := _BNPLManager.contract.FilterLogs(opts, "BNPLLateFeeApplied", arrangementIdRule)
	if err != nil {
		return nil, err
	}
	return &BNPLManagerBNPLLateFeeAppliedIterator{contract: _BNPLManager.contract, event: "BNPLLateFeeApplied", logs: logs, sub: sub}, nil
}

// WatchBNPLLateFeeApplied is a free log subscription operation binding the contract event 0xa771582d9eca55eb62b5ff82bfe51f20afbdb13ab06f2c7ed88ca01c80e322e8.
//
// Solidity: event BNPLLateFeeApplied(uint256 indexed arrangementId, uint8 installmentNumber, uint256 feeAmount, uint256 timestamp)
func (_BNPLManager *BNPLManagerFilterer) WatchBNPLLateFeeApplied(opts *bind.WatchOpts, sink chan<- *BNPLManagerBNPLLateFeeApplied, arrangementId []*big.Int) (event.Subscription, error) {

	var arrangementIdRule []interface{}
	for _, arrangementIdItem := range arrangementId {
		arrangementIdRule = append(arrangementIdRule, arrangementIdItem)
	}

	logs, sub, err := _BNPLManager.contract.WatchLogs(opts, "BNPLLateFeeApplied", arrangementIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BNPLManagerBNPLLateFeeApplied)
				if err := _BNPLManager.contract.UnpackLog(event, "BNPLLateFeeApplied", log); err != nil {
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

// ParseBNPLLateFeeApplied is a log parse operation binding the contract event 0xa771582d9eca55eb62b5ff82bfe51f20afbdb13ab06f2c7ed88ca01c80e322e8.
//
// Solidity: event BNPLLateFeeApplied(uint256 indexed arrangementId, uint8 installmentNumber, uint256 feeAmount, uint256 timestamp)
func (_BNPLManager *BNPLManagerFilterer) ParseBNPLLateFeeApplied(log types.Log) (*BNPLManagerBNPLLateFeeApplied, error) {
	event := new(BNPLManagerBNPLLateFeeApplied)
	if err := _BNPLManager.contract.UnpackLog(event, "BNPLLateFeeApplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BNPLManagerBNPLPaymentMadeIterator is returned from FilterBNPLPaymentMade and is used to iterate over the raw logs and unpacked data for BNPLPaymentMade events raised by the BNPLManager contract.
type BNPLManagerBNPLPaymentMadeIterator struct {
	Event *BNPLManagerBNPLPaymentMade // Event containing the contract specifics and raw log

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
func (it *BNPLManagerBNPLPaymentMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BNPLManagerBNPLPaymentMade)
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
		it.Event = new(BNPLManagerBNPLPaymentMade)
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
func (it *BNPLManagerBNPLPaymentMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BNPLManagerBNPLPaymentMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BNPLManagerBNPLPaymentMade represents a BNPLPaymentMade event raised by the BNPLManager contract.
type BNPLManagerBNPLPaymentMade struct {
	ArrangementId     *big.Int
	InstallmentNumber uint8
	Payer             common.Address
	Amount            *big.Int
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBNPLPaymentMade is a free log retrieval operation binding the contract event 0xe8508bd3254b66027ff6d975978f9924a4472a58ad4da3ea70c8bead5ebad6c7.
//
// Solidity: event BNPLPaymentMade(uint256 indexed arrangementId, uint8 installmentNumber, address indexed payer, uint256 amount, uint256 timestamp)
func (_BNPLManager *BNPLManagerFilterer) FilterBNPLPaymentMade(opts *bind.FilterOpts, arrangementId []*big.Int, payer []common.Address) (*BNPLManagerBNPLPaymentMadeIterator, error) {

	var arrangementIdRule []interface{}
	for _, arrangementIdItem := range arrangementId {
		arrangementIdRule = append(arrangementIdRule, arrangementIdItem)
	}

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _BNPLManager.contract.FilterLogs(opts, "BNPLPaymentMade", arrangementIdRule, payerRule)
	if err != nil {
		return nil, err
	}
	return &BNPLManagerBNPLPaymentMadeIterator{contract: _BNPLManager.contract, event: "BNPLPaymentMade", logs: logs, sub: sub}, nil
}

// WatchBNPLPaymentMade is a free log subscription operation binding the contract event 0xe8508bd3254b66027ff6d975978f9924a4472a58ad4da3ea70c8bead5ebad6c7.
//
// Solidity: event BNPLPaymentMade(uint256 indexed arrangementId, uint8 installmentNumber, address indexed payer, uint256 amount, uint256 timestamp)
func (_BNPLManager *BNPLManagerFilterer) WatchBNPLPaymentMade(opts *bind.WatchOpts, sink chan<- *BNPLManagerBNPLPaymentMade, arrangementId []*big.Int, payer []common.Address) (event.Subscription, error) {

	var arrangementIdRule []interface{}
	for _, arrangementIdItem := range arrangementId {
		arrangementIdRule = append(arrangementIdRule, arrangementIdItem)
	}

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _BNPLManager.contract.WatchLogs(opts, "BNPLPaymentMade", arrangementIdRule, payerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BNPLManagerBNPLPaymentMade)
				if err := _BNPLManager.contract.UnpackLog(event, "BNPLPaymentMade", log); err != nil {
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

// ParseBNPLPaymentMade is a log parse operation binding the contract event 0xe8508bd3254b66027ff6d975978f9924a4472a58ad4da3ea70c8bead5ebad6c7.
//
// Solidity: event BNPLPaymentMade(uint256 indexed arrangementId, uint8 installmentNumber, address indexed payer, uint256 amount, uint256 timestamp)
func (_BNPLManager *BNPLManagerFilterer) ParseBNPLPaymentMade(log types.Log) (*BNPLManagerBNPLPaymentMade, error) {
	event := new(BNPLManagerBNPLPaymentMade)
	if err := _BNPLManager.contract.UnpackLog(event, "BNPLPaymentMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BNPLManagerBNPLRescheduledIterator is returned from FilterBNPLRescheduled and is used to iterate over the raw logs and unpacked data for BNPLRescheduled events raised by the BNPLManager contract.
type BNPLManagerBNPLRescheduledIterator struct {
	Event *BNPLManagerBNPLRescheduled // Event containing the contract specifics and raw log

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
func (it *BNPLManagerBNPLRescheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BNPLManagerBNPLRescheduled)
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
		it.Event = new(BNPLManagerBNPLRescheduled)
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
func (it *BNPLManagerBNPLRescheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BNPLManagerBNPLRescheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BNPLManagerBNPLRescheduled represents a BNPLRescheduled event raised by the BNPLManager contract.
type BNPLManagerBNPLRescheduled struct {
	ArrangementId   *big.Int
	OldScheduleHash [32]byte
	NewScheduleHash [32]byte
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBNPLRescheduled is a free log retrieval operation binding the contract event 0x3986b9d87f88ca7357a73fd27b9b6b2d54915820b5c9dcb7cadf98ec84e391aa.
//
// Solidity: event BNPLRescheduled(uint256 indexed arrangementId, bytes32 oldScheduleHash, bytes32 newScheduleHash, uint256 timestamp)
func (_BNPLManager *BNPLManagerFilterer) FilterBNPLRescheduled(opts *bind.FilterOpts, arrangementId []*big.Int) (*BNPLManagerBNPLRescheduledIterator, error) {

	var arrangementIdRule []interface{}
	for _, arrangementIdItem := range arrangementId {
		arrangementIdRule = append(arrangementIdRule, arrangementIdItem)
	}

	logs, sub, err := _BNPLManager.contract.FilterLogs(opts, "BNPLRescheduled", arrangementIdRule)
	if err != nil {
		return nil, err
	}
	return &BNPLManagerBNPLRescheduledIterator{contract: _BNPLManager.contract, event: "BNPLRescheduled", logs: logs, sub: sub}, nil
}

// WatchBNPLRescheduled is a free log subscription operation binding the contract event 0x3986b9d87f88ca7357a73fd27b9b6b2d54915820b5c9dcb7cadf98ec84e391aa.
//
// Solidity: event BNPLRescheduled(uint256 indexed arrangementId, bytes32 oldScheduleHash, bytes32 newScheduleHash, uint256 timestamp)
func (_BNPLManager *BNPLManagerFilterer) WatchBNPLRescheduled(opts *bind.WatchOpts, sink chan<- *BNPLManagerBNPLRescheduled, arrangementId []*big.Int) (event.Subscription, error) {

	var arrangementIdRule []interface{}
	for _, arrangementIdItem := range arrangementId {
		arrangementIdRule = append(arrangementIdRule, arrangementIdItem)
	}

	logs, sub, err := _BNPLManager.contract.WatchLogs(opts, "BNPLRescheduled", arrangementIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BNPLManagerBNPLRescheduled)
				if err := _BNPLManager.contract.UnpackLog(event, "BNPLRescheduled", log); err != nil {
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

// ParseBNPLRescheduled is a log parse operation binding the contract event 0x3986b9d87f88ca7357a73fd27b9b6b2d54915820b5c9dcb7cadf98ec84e391aa.
//
// Solidity: event BNPLRescheduled(uint256 indexed arrangementId, bytes32 oldScheduleHash, bytes32 newScheduleHash, uint256 timestamp)
func (_BNPLManager *BNPLManagerFilterer) ParseBNPLRescheduled(log types.Log) (*BNPLManagerBNPLRescheduled, error) {
	event := new(BNPLManagerBNPLRescheduled)
	if err := _BNPLManager.contract.UnpackLog(event, "BNPLRescheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BNPLManagerBNPLScheduleChosenIterator is returned from FilterBNPLScheduleChosen and is used to iterate over the raw logs and unpacked data for BNPLScheduleChosen events raised by the BNPLManager contract.
type BNPLManagerBNPLScheduleChosenIterator struct {
	Event *BNPLManagerBNPLScheduleChosen // Event containing the contract specifics and raw log

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
func (it *BNPLManagerBNPLScheduleChosenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BNPLManagerBNPLScheduleChosen)
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
		it.Event = new(BNPLManagerBNPLScheduleChosen)
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
func (it *BNPLManagerBNPLScheduleChosenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BNPLManagerBNPLScheduleChosenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BNPLManagerBNPLScheduleChosen represents a BNPLScheduleChosen event raised by the BNPLManager contract.
type BNPLManagerBNPLScheduleChosen struct {
	ArrangementId   *big.Int
	StartTimestamp  *big.Int
	IntervalSeconds *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBNPLScheduleChosen is a free log retrieval operation binding the contract event 0x683528e54d1d686403cfb3fde979f90e4d361cd6dfda99bae38fd3b9b9dab451.
//
// Solidity: event BNPLScheduleChosen(uint256 indexed arrangementId, uint256 startTimestamp, uint256 intervalSeconds)
func (_BNPLManager *BNPLManagerFilterer) FilterBNPLScheduleChosen(opts *bind.FilterOpts, arrangementId []*big.Int) (*BNPLManagerBNPLScheduleChosenIterator, error) {

	var arrangementIdRule []interface{}
	for _, arrangementIdItem := range arrangementId {
		arrangementIdRule = append(arrangementIdRule, arrangementIdItem)
	}

	logs, sub, err := _BNPLManager.contract.FilterLogs(opts, "BNPLScheduleChosen", arrangementIdRule)
	if err != nil {
		return nil, err
	}
	return &BNPLManagerBNPLScheduleChosenIterator{contract: _BNPLManager.contract, event: "BNPLScheduleChosen", logs: logs, sub: sub}, nil
}

// WatchBNPLScheduleChosen is a free log subscription operation binding the contract event 0x683528e54d1d686403cfb3fde979f90e4d361cd6dfda99bae38fd3b9b9dab451.
//
// Solidity: event BNPLScheduleChosen(uint256 indexed arrangementId, uint256 startTimestamp, uint256 intervalSeconds)
func (_BNPLManager *BNPLManagerFilterer) WatchBNPLScheduleChosen(opts *bind.WatchOpts, sink chan<- *BNPLManagerBNPLScheduleChosen, arrangementId []*big.Int) (event.Subscription, error) {

	var arrangementIdRule []interface{}
	for _, arrangementIdItem := range arrangementId {
		arrangementIdRule = append(arrangementIdRule, arrangementIdItem)
	}

	logs, sub, err := _BNPLManager.contract.WatchLogs(opts, "BNPLScheduleChosen", arrangementIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BNPLManagerBNPLScheduleChosen)
				if err := _BNPLManager.contract.UnpackLog(event, "BNPLScheduleChosen", log); err != nil {
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

// ParseBNPLScheduleChosen is a log parse operation binding the contract event 0x683528e54d1d686403cfb3fde979f90e4d361cd6dfda99bae38fd3b9b9dab451.
//
// Solidity: event BNPLScheduleChosen(uint256 indexed arrangementId, uint256 startTimestamp, uint256 intervalSeconds)
func (_BNPLManager *BNPLManagerFilterer) ParseBNPLScheduleChosen(log types.Log) (*BNPLManagerBNPLScheduleChosen, error) {
	event := new(BNPLManagerBNPLScheduleChosen)
	if err := _BNPLManager.contract.UnpackLog(event, "BNPLScheduleChosen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BNPLManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BNPLManager contract.
type BNPLManagerRoleAdminChangedIterator struct {
	Event *BNPLManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BNPLManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BNPLManagerRoleAdminChanged)
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
		it.Event = new(BNPLManagerRoleAdminChanged)
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
func (it *BNPLManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BNPLManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BNPLManagerRoleAdminChanged represents a RoleAdminChanged event raised by the BNPLManager contract.
type BNPLManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BNPLManager *BNPLManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BNPLManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _BNPLManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BNPLManagerRoleAdminChangedIterator{contract: _BNPLManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BNPLManager *BNPLManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BNPLManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _BNPLManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BNPLManagerRoleAdminChanged)
				if err := _BNPLManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_BNPLManager *BNPLManagerFilterer) ParseRoleAdminChanged(log types.Log) (*BNPLManagerRoleAdminChanged, error) {
	event := new(BNPLManagerRoleAdminChanged)
	if err := _BNPLManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BNPLManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BNPLManager contract.
type BNPLManagerRoleGrantedIterator struct {
	Event *BNPLManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *BNPLManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BNPLManagerRoleGranted)
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
		it.Event = new(BNPLManagerRoleGranted)
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
func (it *BNPLManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BNPLManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BNPLManagerRoleGranted represents a RoleGranted event raised by the BNPLManager contract.
type BNPLManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BNPLManager *BNPLManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BNPLManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _BNPLManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BNPLManagerRoleGrantedIterator{contract: _BNPLManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BNPLManager *BNPLManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BNPLManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BNPLManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BNPLManagerRoleGranted)
				if err := _BNPLManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_BNPLManager *BNPLManagerFilterer) ParseRoleGranted(log types.Log) (*BNPLManagerRoleGranted, error) {
	event := new(BNPLManagerRoleGranted)
	if err := _BNPLManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BNPLManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BNPLManager contract.
type BNPLManagerRoleRevokedIterator struct {
	Event *BNPLManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BNPLManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BNPLManagerRoleRevoked)
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
		it.Event = new(BNPLManagerRoleRevoked)
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
func (it *BNPLManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BNPLManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BNPLManagerRoleRevoked represents a RoleRevoked event raised by the BNPLManager contract.
type BNPLManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BNPLManager *BNPLManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BNPLManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _BNPLManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BNPLManagerRoleRevokedIterator{contract: _BNPLManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BNPLManager *BNPLManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BNPLManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BNPLManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BNPLManagerRoleRevoked)
				if err := _BNPLManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_BNPLManager *BNPLManagerFilterer) ParseRoleRevoked(log types.Log) (*BNPLManagerRoleRevoked, error) {
	event := new(BNPLManagerRoleRevoked)
	if err := _BNPLManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

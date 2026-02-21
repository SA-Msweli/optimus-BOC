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

// LoanManagerMetaData contains all meta data concerning the LoanManager contract.
var LoanManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LOAN_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"approveLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateBps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"durationSeconds\",\"type\":\"uint256\"}],\"name\":\"createLoan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"getAccruedInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"getAmountOwed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"getLoan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateBps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"makePayment\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"markDefaulted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"LoanApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestRateBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"LoanCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"LoanDefaulted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"name\":\"PaymentMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"}]",
	Bin: "0x60806040523460475760017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005560018055603733604b565b50604051610ae590816100d48239f35b5f80fd5b6001600160a01b0381165f9081525f516020610bb95f395f51905f52602052604090205460ff1660ce576001600160a01b03165f8181525f516020610bb95f395f51905f5260205260408120805460ff191660011790553391907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b505f9056fe6080806040526004361015610012575f80fd5b5f3560e01c90816301ffc9a71461080b57508063248a9ca3146107e15780632f2ff15d146107a457806336568abe146107605780633987031314610726578063504006ca146106a45780635114cb52146103f5578063732164501461039557806391d148541461034d5780639339ea6914610324578063a217fddf1461030a578063aadc1ac1146102a8578063afd765ba1461012f578063d139d9b9146101095763d547741f146100c1575f80fd5b34610105576040366003190112610105576101036004356100e061085e565b906100fe6100f9825f525f602052600160405f20015490565b61090b565b6109c5565b005b5f80fd5b346101055760203660031901126101055760206101276004356108cb565b604051908152f35b346101055760a0366003190112610105576004356001600160a01b03811690819003610105576084359060643590604435801561027257831561023d57600154925f19841461022957608084927f2b6e7be0390a80ec9c24c00d1dbf95d0cc27e42970c49fdd9d2d9f8b7a8764669260209760018601600155855f526002895260405f209086825560018201886bffffffffffffffffffffffff60a01b82541617905583600283015582600383015560076101f16004840192428455426108b1565b92600581019384555f60068201550160ff198154169055549054916040519384528984015260408301526060820152a3604051908152f35b634e487b7160e01b5f52601160045260245ffd5b60405162461bcd60e51b815260206004820152600d60248201526c2d22a927afa22aa920aa24a7a760991b6044820152606490fd5b60405162461bcd60e51b815260206004820152600e60248201526d16915493d7d414925390d254105360921b6044820152606490fd5b3461010557602036600319011261010557600435805f526002602052600760405f206102d681541515610874565b01805460ff1916600117905533907f8b6240997bc3af56f5a9aec0408b8286fb78fb0573f2ba96445be3918511cb935f80a3005b34610105575f3660031901126101055760206040515f8152f35b34610105576020366003190112610105576004355f526002602052602061012760405f20610a58565b346101055760403660031901126101055761036661085e565b6004355f525f60205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b3461010557602036600319011261010557600435805f526002602052600760405f206103c381541515610874565b01805460ff191660031790557f0789b7097e8066538cfaa1132488b132e14ba5f0c938c8b7aaf8cf40356aab0b5f80a2005b60203660031901126101055760043560027f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0054146106955760027f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055805f52600260205260405f2061046981541515610874565b6007810191600160ff8454160361065c57341561062857600661049861048e84610a58565b60028501546108b1565b92019081548084115f14610621576104b081856108be565b915b348310610616576104c43480936108b1565b8085556104d183346108be565b93838082111561060d576104e4916108be565b905b604051938452602084015260408301527f2144ed2980d83d1b3c2d4c4ffbaf398ded900d0e5e18fcc085a686569beb967360603393a380610562575b50541015610551575b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055005b805460ff191660021790558061052b565b5f80808093335af13d15610608573d67ffffffffffffffff81116105f45760405190601f8101601f19908116603f0116820167ffffffffffffffff8111838210176105f45760405281525f60203d92013e5b156105bf5783610522565b60405162461bcd60e51b815260206004820152600d60248201526c14915195539117d19052531151609a1b6044820152606490fd5b634e487b7160e01b5f52604160045260245ffd5b6105b4565b50505f906104e6565b6104c48380936108b1565b5f916104b2565b60405162461bcd60e51b815260206004820152600c60248201526b16915493d7d410565351539560a21b6044820152606490fd5b60405162461bcd60e51b81526020600482015260116024820152701313d05397d393d517d054141493d59151607a1b6044820152606490fd5b633ee5aeb560e01b5f5260045ffd5b34610105576020366003190112610105576004355f52600260205261010060405f2080549060018060a01b03600182015416906002810154600382015460048301549060058401549260ff600760068701549601541695604051978852602088015260408701526060860152608085015260a084015260c083015260e0820152f35b34610105575f3660031901126101055760206040517fd8a2202037a0202b1ba54400786c6ec71c73c1d2417d7c029d18db53247b12538152f35b346101055760403660031901126101055761077961085e565b336001600160a01b0382160361079557610103906004356109c5565b63334bd91960e11b5f5260045ffd5b34610105576040366003190112610105576101036004356107c361085e565b906107dc6100f9825f525f602052600160405f20015490565b610943565b346101055760203660031901126101055760206101276004355f525f602052600160405f20015490565b34610105576020366003190112610105576004359063ffffffff60e01b821680920361010557602091637965db0b60e01b811490811561084d575b5015158152f35b6301ffc9a760e01b14905083610846565b602435906001600160a01b038216820361010557565b1561087b57565b60405162461bcd60e51b815260206004820152600e60248201526d1313d05397d393d517d193d5539160921b6044820152606490fd5b9190820180921161022957565b9190820391821161022957565b5f52600260205260405f2060066108ee6108e483610a58565b60028401546108b1565b9101548082111561090557610902916108be565b90565b50505f90565b5f8181526020818152604080832033845290915290205460ff161561092d5750565b63e2517d3f60e01b5f523360045260245260445ffd5b5f818152602081815260408083206001600160a01b038616845290915290205460ff16610905575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b5f818152602081815260408083206001600160a01b038616845290915290205460ff1615610905575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b8181029291811591840414171561022957565b805415610aaa57610a9b81600464496cebb8009301548042115f14610a9f57610a84610a9691426108be565b915b6003600282015491015490610a45565b610a45565b0490565b50610a965f91610a86565b505f9056fea2646970667358221220c2daea65dde36914f04dc290dbab0c066ec6bf54c8f63825f3671e29e32376ac64736f6c63430008210033ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
}

// LoanManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use LoanManagerMetaData.ABI instead.
var LoanManagerABI = LoanManagerMetaData.ABI

// LoanManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LoanManagerMetaData.Bin instead.
var LoanManagerBin = LoanManagerMetaData.Bin

// DeployLoanManager deploys a new Ethereum contract, binding an instance of LoanManager to it.
func DeployLoanManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LoanManager, error) {
	parsed, err := LoanManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LoanManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LoanManager{LoanManagerCaller: LoanManagerCaller{contract: contract}, LoanManagerTransactor: LoanManagerTransactor{contract: contract}, LoanManagerFilterer: LoanManagerFilterer{contract: contract}}, nil
}

// LoanManager is an auto generated Go binding around an Ethereum contract.
type LoanManager struct {
	LoanManagerCaller     // Read-only binding to the contract
	LoanManagerTransactor // Write-only binding to the contract
	LoanManagerFilterer   // Log filterer for contract events
}

// LoanManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LoanManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoanManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LoanManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoanManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoanManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoanManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoanManagerSession struct {
	Contract     *LoanManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoanManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoanManagerCallerSession struct {
	Contract *LoanManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// LoanManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoanManagerTransactorSession struct {
	Contract     *LoanManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// LoanManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LoanManagerRaw struct {
	Contract *LoanManager // Generic contract binding to access the raw methods on
}

// LoanManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoanManagerCallerRaw struct {
	Contract *LoanManagerCaller // Generic read-only contract binding to access the raw methods on
}

// LoanManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoanManagerTransactorRaw struct {
	Contract *LoanManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLoanManager creates a new instance of LoanManager, bound to a specific deployed contract.
func NewLoanManager(address common.Address, backend bind.ContractBackend) (*LoanManager, error) {
	contract, err := bindLoanManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LoanManager{LoanManagerCaller: LoanManagerCaller{contract: contract}, LoanManagerTransactor: LoanManagerTransactor{contract: contract}, LoanManagerFilterer: LoanManagerFilterer{contract: contract}}, nil
}

// NewLoanManagerCaller creates a new read-only instance of LoanManager, bound to a specific deployed contract.
func NewLoanManagerCaller(address common.Address, caller bind.ContractCaller) (*LoanManagerCaller, error) {
	contract, err := bindLoanManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoanManagerCaller{contract: contract}, nil
}

// NewLoanManagerTransactor creates a new write-only instance of LoanManager, bound to a specific deployed contract.
func NewLoanManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*LoanManagerTransactor, error) {
	contract, err := bindLoanManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoanManagerTransactor{contract: contract}, nil
}

// NewLoanManagerFilterer creates a new log filterer instance of LoanManager, bound to a specific deployed contract.
func NewLoanManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*LoanManagerFilterer, error) {
	contract, err := bindLoanManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoanManagerFilterer{contract: contract}, nil
}

// bindLoanManager binds a generic wrapper to an already deployed contract.
func bindLoanManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LoanManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoanManager *LoanManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LoanManager.Contract.LoanManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoanManager *LoanManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoanManager.Contract.LoanManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoanManager *LoanManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoanManager.Contract.LoanManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoanManager *LoanManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LoanManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoanManager *LoanManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoanManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoanManager *LoanManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoanManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LoanManager *LoanManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LoanManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LoanManager *LoanManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LoanManager.Contract.DEFAULTADMINROLE(&_LoanManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LoanManager *LoanManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LoanManager.Contract.DEFAULTADMINROLE(&_LoanManager.CallOpts)
}

// LOANADMINROLE is a free data retrieval call binding the contract method 0x39870313.
//
// Solidity: function LOAN_ADMIN_ROLE() view returns(bytes32)
func (_LoanManager *LoanManagerCaller) LOANADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LoanManager.contract.Call(opts, &out, "LOAN_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LOANADMINROLE is a free data retrieval call binding the contract method 0x39870313.
//
// Solidity: function LOAN_ADMIN_ROLE() view returns(bytes32)
func (_LoanManager *LoanManagerSession) LOANADMINROLE() ([32]byte, error) {
	return _LoanManager.Contract.LOANADMINROLE(&_LoanManager.CallOpts)
}

// LOANADMINROLE is a free data retrieval call binding the contract method 0x39870313.
//
// Solidity: function LOAN_ADMIN_ROLE() view returns(bytes32)
func (_LoanManager *LoanManagerCallerSession) LOANADMINROLE() ([32]byte, error) {
	return _LoanManager.Contract.LOANADMINROLE(&_LoanManager.CallOpts)
}

// GetAccruedInterest is a free data retrieval call binding the contract method 0x9339ea69.
//
// Solidity: function getAccruedInterest(uint256 loanId) view returns(uint256)
func (_LoanManager *LoanManagerCaller) GetAccruedInterest(opts *bind.CallOpts, loanId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LoanManager.contract.Call(opts, &out, "getAccruedInterest", loanId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccruedInterest is a free data retrieval call binding the contract method 0x9339ea69.
//
// Solidity: function getAccruedInterest(uint256 loanId) view returns(uint256)
func (_LoanManager *LoanManagerSession) GetAccruedInterest(loanId *big.Int) (*big.Int, error) {
	return _LoanManager.Contract.GetAccruedInterest(&_LoanManager.CallOpts, loanId)
}

// GetAccruedInterest is a free data retrieval call binding the contract method 0x9339ea69.
//
// Solidity: function getAccruedInterest(uint256 loanId) view returns(uint256)
func (_LoanManager *LoanManagerCallerSession) GetAccruedInterest(loanId *big.Int) (*big.Int, error) {
	return _LoanManager.Contract.GetAccruedInterest(&_LoanManager.CallOpts, loanId)
}

// GetAmountOwed is a free data retrieval call binding the contract method 0xd139d9b9.
//
// Solidity: function getAmountOwed(uint256 loanId) view returns(uint256)
func (_LoanManager *LoanManagerCaller) GetAmountOwed(opts *bind.CallOpts, loanId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LoanManager.contract.Call(opts, &out, "getAmountOwed", loanId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOwed is a free data retrieval call binding the contract method 0xd139d9b9.
//
// Solidity: function getAmountOwed(uint256 loanId) view returns(uint256)
func (_LoanManager *LoanManagerSession) GetAmountOwed(loanId *big.Int) (*big.Int, error) {
	return _LoanManager.Contract.GetAmountOwed(&_LoanManager.CallOpts, loanId)
}

// GetAmountOwed is a free data retrieval call binding the contract method 0xd139d9b9.
//
// Solidity: function getAmountOwed(uint256 loanId) view returns(uint256)
func (_LoanManager *LoanManagerCallerSession) GetAmountOwed(loanId *big.Int) (*big.Int, error) {
	return _LoanManager.Contract.GetAmountOwed(&_LoanManager.CallOpts, loanId)
}

// GetLoan is a free data retrieval call binding the contract method 0x504006ca.
//
// Solidity: function getLoan(uint256 loanId) view returns(uint256 id, address borrower, uint256 principal, uint256 interestRateBps, uint256 startTime, uint256 endTime, uint256 amountPaid, uint8 status)
func (_LoanManager *LoanManagerCaller) GetLoan(opts *bind.CallOpts, loanId *big.Int) (struct {
	Id              *big.Int
	Borrower        common.Address
	Principal       *big.Int
	InterestRateBps *big.Int
	StartTime       *big.Int
	EndTime         *big.Int
	AmountPaid      *big.Int
	Status          uint8
}, error) {
	var out []interface{}
	err := _LoanManager.contract.Call(opts, &out, "getLoan", loanId)

	outstruct := new(struct {
		Id              *big.Int
		Borrower        common.Address
		Principal       *big.Int
		InterestRateBps *big.Int
		StartTime       *big.Int
		EndTime         *big.Int
		AmountPaid      *big.Int
		Status          uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Borrower = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Principal = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.InterestRateBps = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.AmountPaid = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[7], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetLoan is a free data retrieval call binding the contract method 0x504006ca.
//
// Solidity: function getLoan(uint256 loanId) view returns(uint256 id, address borrower, uint256 principal, uint256 interestRateBps, uint256 startTime, uint256 endTime, uint256 amountPaid, uint8 status)
func (_LoanManager *LoanManagerSession) GetLoan(loanId *big.Int) (struct {
	Id              *big.Int
	Borrower        common.Address
	Principal       *big.Int
	InterestRateBps *big.Int
	StartTime       *big.Int
	EndTime         *big.Int
	AmountPaid      *big.Int
	Status          uint8
}, error) {
	return _LoanManager.Contract.GetLoan(&_LoanManager.CallOpts, loanId)
}

// GetLoan is a free data retrieval call binding the contract method 0x504006ca.
//
// Solidity: function getLoan(uint256 loanId) view returns(uint256 id, address borrower, uint256 principal, uint256 interestRateBps, uint256 startTime, uint256 endTime, uint256 amountPaid, uint8 status)
func (_LoanManager *LoanManagerCallerSession) GetLoan(loanId *big.Int) (struct {
	Id              *big.Int
	Borrower        common.Address
	Principal       *big.Int
	InterestRateBps *big.Int
	StartTime       *big.Int
	EndTime         *big.Int
	AmountPaid      *big.Int
	Status          uint8
}, error) {
	return _LoanManager.Contract.GetLoan(&_LoanManager.CallOpts, loanId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LoanManager *LoanManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LoanManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LoanManager *LoanManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LoanManager.Contract.GetRoleAdmin(&_LoanManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LoanManager *LoanManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LoanManager.Contract.GetRoleAdmin(&_LoanManager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LoanManager *LoanManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _LoanManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LoanManager *LoanManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LoanManager.Contract.HasRole(&_LoanManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LoanManager *LoanManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LoanManager.Contract.HasRole(&_LoanManager.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LoanManager *LoanManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LoanManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LoanManager *LoanManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LoanManager.Contract.SupportsInterface(&_LoanManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LoanManager *LoanManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LoanManager.Contract.SupportsInterface(&_LoanManager.CallOpts, interfaceId)
}

// ApproveLoan is a paid mutator transaction binding the contract method 0xaadc1ac1.
//
// Solidity: function approveLoan(uint256 loanId) returns()
func (_LoanManager *LoanManagerTransactor) ApproveLoan(opts *bind.TransactOpts, loanId *big.Int) (*types.Transaction, error) {
	return _LoanManager.contract.Transact(opts, "approveLoan", loanId)
}

// ApproveLoan is a paid mutator transaction binding the contract method 0xaadc1ac1.
//
// Solidity: function approveLoan(uint256 loanId) returns()
func (_LoanManager *LoanManagerSession) ApproveLoan(loanId *big.Int) (*types.Transaction, error) {
	return _LoanManager.Contract.ApproveLoan(&_LoanManager.TransactOpts, loanId)
}

// ApproveLoan is a paid mutator transaction binding the contract method 0xaadc1ac1.
//
// Solidity: function approveLoan(uint256 loanId) returns()
func (_LoanManager *LoanManagerTransactorSession) ApproveLoan(loanId *big.Int) (*types.Transaction, error) {
	return _LoanManager.Contract.ApproveLoan(&_LoanManager.TransactOpts, loanId)
}

// CreateLoan is a paid mutator transaction binding the contract method 0xafd765ba.
//
// Solidity: function createLoan(address borrower, uint256 , uint256 principal, uint256 interestRateBps, uint256 durationSeconds) returns(uint256 loanId)
func (_LoanManager *LoanManagerTransactor) CreateLoan(opts *bind.TransactOpts, borrower common.Address, arg1 *big.Int, principal *big.Int, interestRateBps *big.Int, durationSeconds *big.Int) (*types.Transaction, error) {
	return _LoanManager.contract.Transact(opts, "createLoan", borrower, arg1, principal, interestRateBps, durationSeconds)
}

// CreateLoan is a paid mutator transaction binding the contract method 0xafd765ba.
//
// Solidity: function createLoan(address borrower, uint256 , uint256 principal, uint256 interestRateBps, uint256 durationSeconds) returns(uint256 loanId)
func (_LoanManager *LoanManagerSession) CreateLoan(borrower common.Address, arg1 *big.Int, principal *big.Int, interestRateBps *big.Int, durationSeconds *big.Int) (*types.Transaction, error) {
	return _LoanManager.Contract.CreateLoan(&_LoanManager.TransactOpts, borrower, arg1, principal, interestRateBps, durationSeconds)
}

// CreateLoan is a paid mutator transaction binding the contract method 0xafd765ba.
//
// Solidity: function createLoan(address borrower, uint256 , uint256 principal, uint256 interestRateBps, uint256 durationSeconds) returns(uint256 loanId)
func (_LoanManager *LoanManagerTransactorSession) CreateLoan(borrower common.Address, arg1 *big.Int, principal *big.Int, interestRateBps *big.Int, durationSeconds *big.Int) (*types.Transaction, error) {
	return _LoanManager.Contract.CreateLoan(&_LoanManager.TransactOpts, borrower, arg1, principal, interestRateBps, durationSeconds)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LoanManager *LoanManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LoanManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LoanManager *LoanManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LoanManager.Contract.GrantRole(&_LoanManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LoanManager *LoanManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LoanManager.Contract.GrantRole(&_LoanManager.TransactOpts, role, account)
}

// MakePayment is a paid mutator transaction binding the contract method 0x5114cb52.
//
// Solidity: function makePayment(uint256 loanId) payable returns()
func (_LoanManager *LoanManagerTransactor) MakePayment(opts *bind.TransactOpts, loanId *big.Int) (*types.Transaction, error) {
	return _LoanManager.contract.Transact(opts, "makePayment", loanId)
}

// MakePayment is a paid mutator transaction binding the contract method 0x5114cb52.
//
// Solidity: function makePayment(uint256 loanId) payable returns()
func (_LoanManager *LoanManagerSession) MakePayment(loanId *big.Int) (*types.Transaction, error) {
	return _LoanManager.Contract.MakePayment(&_LoanManager.TransactOpts, loanId)
}

// MakePayment is a paid mutator transaction binding the contract method 0x5114cb52.
//
// Solidity: function makePayment(uint256 loanId) payable returns()
func (_LoanManager *LoanManagerTransactorSession) MakePayment(loanId *big.Int) (*types.Transaction, error) {
	return _LoanManager.Contract.MakePayment(&_LoanManager.TransactOpts, loanId)
}

// MarkDefaulted is a paid mutator transaction binding the contract method 0x73216450.
//
// Solidity: function markDefaulted(uint256 loanId) returns()
func (_LoanManager *LoanManagerTransactor) MarkDefaulted(opts *bind.TransactOpts, loanId *big.Int) (*types.Transaction, error) {
	return _LoanManager.contract.Transact(opts, "markDefaulted", loanId)
}

// MarkDefaulted is a paid mutator transaction binding the contract method 0x73216450.
//
// Solidity: function markDefaulted(uint256 loanId) returns()
func (_LoanManager *LoanManagerSession) MarkDefaulted(loanId *big.Int) (*types.Transaction, error) {
	return _LoanManager.Contract.MarkDefaulted(&_LoanManager.TransactOpts, loanId)
}

// MarkDefaulted is a paid mutator transaction binding the contract method 0x73216450.
//
// Solidity: function markDefaulted(uint256 loanId) returns()
func (_LoanManager *LoanManagerTransactorSession) MarkDefaulted(loanId *big.Int) (*types.Transaction, error) {
	return _LoanManager.Contract.MarkDefaulted(&_LoanManager.TransactOpts, loanId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LoanManager *LoanManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LoanManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LoanManager *LoanManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LoanManager.Contract.RenounceRole(&_LoanManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LoanManager *LoanManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LoanManager.Contract.RenounceRole(&_LoanManager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LoanManager *LoanManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LoanManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LoanManager *LoanManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LoanManager.Contract.RevokeRole(&_LoanManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LoanManager *LoanManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LoanManager.Contract.RevokeRole(&_LoanManager.TransactOpts, role, account)
}

// LoanManagerLoanApprovedIterator is returned from FilterLoanApproved and is used to iterate over the raw logs and unpacked data for LoanApproved events raised by the LoanManager contract.
type LoanManagerLoanApprovedIterator struct {
	Event *LoanManagerLoanApproved // Event containing the contract specifics and raw log

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
func (it *LoanManagerLoanApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerLoanApproved)
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
		it.Event = new(LoanManagerLoanApproved)
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
func (it *LoanManagerLoanApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerLoanApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerLoanApproved represents a LoanApproved event raised by the LoanManager contract.
type LoanManagerLoanApproved struct {
	LoanId   *big.Int
	Approver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLoanApproved is a free log retrieval operation binding the contract event 0x8b6240997bc3af56f5a9aec0408b8286fb78fb0573f2ba96445be3918511cb93.
//
// Solidity: event LoanApproved(uint256 indexed loanId, address indexed approver)
func (_LoanManager *LoanManagerFilterer) FilterLoanApproved(opts *bind.FilterOpts, loanId []*big.Int, approver []common.Address) (*LoanManagerLoanApprovedIterator, error) {

	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}
	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "LoanApproved", loanIdRule, approverRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerLoanApprovedIterator{contract: _LoanManager.contract, event: "LoanApproved", logs: logs, sub: sub}, nil
}

// WatchLoanApproved is a free log subscription operation binding the contract event 0x8b6240997bc3af56f5a9aec0408b8286fb78fb0573f2ba96445be3918511cb93.
//
// Solidity: event LoanApproved(uint256 indexed loanId, address indexed approver)
func (_LoanManager *LoanManagerFilterer) WatchLoanApproved(opts *bind.WatchOpts, sink chan<- *LoanManagerLoanApproved, loanId []*big.Int, approver []common.Address) (event.Subscription, error) {

	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}
	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "LoanApproved", loanIdRule, approverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerLoanApproved)
				if err := _LoanManager.contract.UnpackLog(event, "LoanApproved", log); err != nil {
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

// ParseLoanApproved is a log parse operation binding the contract event 0x8b6240997bc3af56f5a9aec0408b8286fb78fb0573f2ba96445be3918511cb93.
//
// Solidity: event LoanApproved(uint256 indexed loanId, address indexed approver)
func (_LoanManager *LoanManagerFilterer) ParseLoanApproved(log types.Log) (*LoanManagerLoanApproved, error) {
	event := new(LoanManagerLoanApproved)
	if err := _LoanManager.contract.UnpackLog(event, "LoanApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanManagerLoanCreatedIterator is returned from FilterLoanCreated and is used to iterate over the raw logs and unpacked data for LoanCreated events raised by the LoanManager contract.
type LoanManagerLoanCreatedIterator struct {
	Event *LoanManagerLoanCreated // Event containing the contract specifics and raw log

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
func (it *LoanManagerLoanCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerLoanCreated)
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
		it.Event = new(LoanManagerLoanCreated)
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
func (it *LoanManagerLoanCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerLoanCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerLoanCreated represents a LoanCreated event raised by the LoanManager contract.
type LoanManagerLoanCreated struct {
	LoanId          *big.Int
	Borrower        common.Address
	Principal       *big.Int
	InterestRateBps *big.Int
	StartTime       *big.Int
	EndTime         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLoanCreated is a free log retrieval operation binding the contract event 0x2b6e7be0390a80ec9c24c00d1dbf95d0cc27e42970c49fdd9d2d9f8b7a876466.
//
// Solidity: event LoanCreated(uint256 indexed loanId, address indexed borrower, uint256 principal, uint256 interestRateBps, uint256 startTime, uint256 endTime)
func (_LoanManager *LoanManagerFilterer) FilterLoanCreated(opts *bind.FilterOpts, loanId []*big.Int, borrower []common.Address) (*LoanManagerLoanCreatedIterator, error) {

	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "LoanCreated", loanIdRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerLoanCreatedIterator{contract: _LoanManager.contract, event: "LoanCreated", logs: logs, sub: sub}, nil
}

// WatchLoanCreated is a free log subscription operation binding the contract event 0x2b6e7be0390a80ec9c24c00d1dbf95d0cc27e42970c49fdd9d2d9f8b7a876466.
//
// Solidity: event LoanCreated(uint256 indexed loanId, address indexed borrower, uint256 principal, uint256 interestRateBps, uint256 startTime, uint256 endTime)
func (_LoanManager *LoanManagerFilterer) WatchLoanCreated(opts *bind.WatchOpts, sink chan<- *LoanManagerLoanCreated, loanId []*big.Int, borrower []common.Address) (event.Subscription, error) {

	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "LoanCreated", loanIdRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerLoanCreated)
				if err := _LoanManager.contract.UnpackLog(event, "LoanCreated", log); err != nil {
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

// ParseLoanCreated is a log parse operation binding the contract event 0x2b6e7be0390a80ec9c24c00d1dbf95d0cc27e42970c49fdd9d2d9f8b7a876466.
//
// Solidity: event LoanCreated(uint256 indexed loanId, address indexed borrower, uint256 principal, uint256 interestRateBps, uint256 startTime, uint256 endTime)
func (_LoanManager *LoanManagerFilterer) ParseLoanCreated(log types.Log) (*LoanManagerLoanCreated, error) {
	event := new(LoanManagerLoanCreated)
	if err := _LoanManager.contract.UnpackLog(event, "LoanCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanManagerLoanDefaultedIterator is returned from FilterLoanDefaulted and is used to iterate over the raw logs and unpacked data for LoanDefaulted events raised by the LoanManager contract.
type LoanManagerLoanDefaultedIterator struct {
	Event *LoanManagerLoanDefaulted // Event containing the contract specifics and raw log

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
func (it *LoanManagerLoanDefaultedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerLoanDefaulted)
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
		it.Event = new(LoanManagerLoanDefaulted)
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
func (it *LoanManagerLoanDefaultedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerLoanDefaultedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerLoanDefaulted represents a LoanDefaulted event raised by the LoanManager contract.
type LoanManagerLoanDefaulted struct {
	LoanId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLoanDefaulted is a free log retrieval operation binding the contract event 0x0789b7097e8066538cfaa1132488b132e14ba5f0c938c8b7aaf8cf40356aab0b.
//
// Solidity: event LoanDefaulted(uint256 indexed loanId)
func (_LoanManager *LoanManagerFilterer) FilterLoanDefaulted(opts *bind.FilterOpts, loanId []*big.Int) (*LoanManagerLoanDefaultedIterator, error) {

	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "LoanDefaulted", loanIdRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerLoanDefaultedIterator{contract: _LoanManager.contract, event: "LoanDefaulted", logs: logs, sub: sub}, nil
}

// WatchLoanDefaulted is a free log subscription operation binding the contract event 0x0789b7097e8066538cfaa1132488b132e14ba5f0c938c8b7aaf8cf40356aab0b.
//
// Solidity: event LoanDefaulted(uint256 indexed loanId)
func (_LoanManager *LoanManagerFilterer) WatchLoanDefaulted(opts *bind.WatchOpts, sink chan<- *LoanManagerLoanDefaulted, loanId []*big.Int) (event.Subscription, error) {

	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "LoanDefaulted", loanIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerLoanDefaulted)
				if err := _LoanManager.contract.UnpackLog(event, "LoanDefaulted", log); err != nil {
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

// ParseLoanDefaulted is a log parse operation binding the contract event 0x0789b7097e8066538cfaa1132488b132e14ba5f0c938c8b7aaf8cf40356aab0b.
//
// Solidity: event LoanDefaulted(uint256 indexed loanId)
func (_LoanManager *LoanManagerFilterer) ParseLoanDefaulted(log types.Log) (*LoanManagerLoanDefaulted, error) {
	event := new(LoanManagerLoanDefaulted)
	if err := _LoanManager.contract.UnpackLog(event, "LoanDefaulted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanManagerPaymentMadeIterator is returned from FilterPaymentMade and is used to iterate over the raw logs and unpacked data for PaymentMade events raised by the LoanManager contract.
type LoanManagerPaymentMadeIterator struct {
	Event *LoanManagerPaymentMade // Event containing the contract specifics and raw log

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
func (it *LoanManagerPaymentMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerPaymentMade)
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
		it.Event = new(LoanManagerPaymentMade)
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
func (it *LoanManagerPaymentMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerPaymentMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerPaymentMade represents a PaymentMade event raised by the LoanManager contract.
type LoanManagerPaymentMade struct {
	LoanId     *big.Int
	Payer      common.Address
	Amount     *big.Int
	AmountPaid *big.Int
	Remaining  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPaymentMade is a free log retrieval operation binding the contract event 0x2144ed2980d83d1b3c2d4c4ffbaf398ded900d0e5e18fcc085a686569beb9673.
//
// Solidity: event PaymentMade(uint256 indexed loanId, address indexed payer, uint256 amount, uint256 amountPaid, uint256 remaining)
func (_LoanManager *LoanManagerFilterer) FilterPaymentMade(opts *bind.FilterOpts, loanId []*big.Int, payer []common.Address) (*LoanManagerPaymentMadeIterator, error) {

	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "PaymentMade", loanIdRule, payerRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerPaymentMadeIterator{contract: _LoanManager.contract, event: "PaymentMade", logs: logs, sub: sub}, nil
}

// WatchPaymentMade is a free log subscription operation binding the contract event 0x2144ed2980d83d1b3c2d4c4ffbaf398ded900d0e5e18fcc085a686569beb9673.
//
// Solidity: event PaymentMade(uint256 indexed loanId, address indexed payer, uint256 amount, uint256 amountPaid, uint256 remaining)
func (_LoanManager *LoanManagerFilterer) WatchPaymentMade(opts *bind.WatchOpts, sink chan<- *LoanManagerPaymentMade, loanId []*big.Int, payer []common.Address) (event.Subscription, error) {

	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "PaymentMade", loanIdRule, payerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerPaymentMade)
				if err := _LoanManager.contract.UnpackLog(event, "PaymentMade", log); err != nil {
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

// ParsePaymentMade is a log parse operation binding the contract event 0x2144ed2980d83d1b3c2d4c4ffbaf398ded900d0e5e18fcc085a686569beb9673.
//
// Solidity: event PaymentMade(uint256 indexed loanId, address indexed payer, uint256 amount, uint256 amountPaid, uint256 remaining)
func (_LoanManager *LoanManagerFilterer) ParsePaymentMade(log types.Log) (*LoanManagerPaymentMade, error) {
	event := new(LoanManagerPaymentMade)
	if err := _LoanManager.contract.UnpackLog(event, "PaymentMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the LoanManager contract.
type LoanManagerRoleAdminChangedIterator struct {
	Event *LoanManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LoanManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerRoleAdminChanged)
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
		it.Event = new(LoanManagerRoleAdminChanged)
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
func (it *LoanManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerRoleAdminChanged represents a RoleAdminChanged event raised by the LoanManager contract.
type LoanManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LoanManager *LoanManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LoanManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerRoleAdminChangedIterator{contract: _LoanManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LoanManager *LoanManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LoanManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerRoleAdminChanged)
				if err := _LoanManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_LoanManager *LoanManagerFilterer) ParseRoleAdminChanged(log types.Log) (*LoanManagerRoleAdminChanged, error) {
	event := new(LoanManagerRoleAdminChanged)
	if err := _LoanManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the LoanManager contract.
type LoanManagerRoleGrantedIterator struct {
	Event *LoanManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *LoanManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerRoleGranted)
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
		it.Event = new(LoanManagerRoleGranted)
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
func (it *LoanManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerRoleGranted represents a RoleGranted event raised by the LoanManager contract.
type LoanManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LoanManager *LoanManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LoanManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerRoleGrantedIterator{contract: _LoanManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LoanManager *LoanManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LoanManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerRoleGranted)
				if err := _LoanManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_LoanManager *LoanManagerFilterer) ParseRoleGranted(log types.Log) (*LoanManagerRoleGranted, error) {
	event := new(LoanManagerRoleGranted)
	if err := _LoanManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the LoanManager contract.
type LoanManagerRoleRevokedIterator struct {
	Event *LoanManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LoanManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerRoleRevoked)
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
		it.Event = new(LoanManagerRoleRevoked)
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
func (it *LoanManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerRoleRevoked represents a RoleRevoked event raised by the LoanManager contract.
type LoanManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LoanManager *LoanManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LoanManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerRoleRevokedIterator{contract: _LoanManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LoanManager *LoanManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LoanManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerRoleRevoked)
				if err := _LoanManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_LoanManager *LoanManagerFilterer) ParseRoleRevoked(log types.Log) (*LoanManagerRoleRevoked, error) {
	event := new(LoanManagerRoleRevoked)
	if err := _LoanManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

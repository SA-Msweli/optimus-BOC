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

// TokenVaultMetaData contains all meta data concerning the TokenVault contract.
var TokenVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VAULT_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"daoId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TreasuryWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"}]",
	Bin: "0x60806040523460435760017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00556033336047565b506040516108c390816100d08239f35b5f80fd5b6001600160a01b0381165f9081525f5160206109935f395f51905f52602052604090205460ff1660ca576001600160a01b03165f8181525f5160206109935f395f51905f5260205260408120805460ff191660011790553391907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b505f9056fe6080806040526004361015610012575f80fd5b5f3560e01c90816301ffc9a71461068057508063248a9ca31461064e5780632f2ff15d1461061157806336568abe146105cd57806347e7ef241461046d5780635e5a24a41461043357806391d14854146103eb578063a217fddf146103d1578063d547741f1461038d578063f3fef3a3146100d15763f8b2cb4f14610095575f80fd5b346100cd5760203660031901126100cd576001600160a01b036100b66106e9565b165f526001602052602060405f2054604051908152f35b5f80fd5b346100cd5760403660031901126100cd576100ea6106e9565b6024359060027f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00541461037e5760027f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055335f9081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff168015610347575b156103115760018060a01b031690815f5260016020528060405f2054106102d557815f52600160205260405f208054908282039182116102c1575560405163a9059cbb60e01b8152336004820152602481018290526020816044815f875af19081156102b6575f91610287575b501561024a57815f52600160205260405f205460405191825260208201527fc2b4a290c20fb28939d29f102514fbffd2b73c059ffba8b78250c94161d5fcc660403392a360017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055005b60405162461bcd60e51b8152602060048201526015602482015274115490cc8c17d514905394d1915497d19052531151605a1b6044820152606490fd5b6102a9915060203d6020116102af575b6102a181836106ff565b810190610735565b836101e0565b503d610297565b6040513d5f823e3d90fd5b634e487b7160e01b5f52601160045260245ffd5b60405162461bcd60e51b8152602060048201526014602482015273494e53554646494349454e545f42414c414e434560601b6044820152606490fd5b60405162461bcd60e51b815260206004820152600e60248201526d1393d517d055551213d49256915160921b6044820152606490fd5b50335f9081527f892718e7a1a469066675b5f782493ac861b4f98b3c42b6d42b67d853aabe987b602052604090205460ff16610173565b633ee5aeb560e01b5f5260045ffd5b346100cd5760403660031901126100cd576103cf6004356103ac6106d3565b906103ca6103c5825f525f602052600160405f20015490565b61074d565b61080d565b005b346100cd575f3660031901126100cd5760206040515f8152f35b346100cd5760403660031901126100cd576104046106d3565b6004355f525f60205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b346100cd575f3660031901126100cd5760206040517fd1473398bb66596de5d1ea1fc8e303ff2ac23265adc9144b1b52065dc4f0934b8152f35b346100cd5760403660031901126100cd576104866106e9565b60243590811561059a576040516323b872dd60e01b8152336004820152306024820152604481018390526001600160a01b039190911691906020816064815f875af19081156102b6575f9161057b575b501561053657815f52600160205260405f208054908282018092116102c15755815f52600160205260405f205460405191825260208201527fdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d760403392a3005b60405162461bcd60e51b815260206004820152601a60248201527f45524332305f5452414e534645525f46524f4d5f4641494c45440000000000006044820152606490fd5b610594915060203d6020116102af576102a181836106ff565b836104d6565b60405162461bcd60e51b815260206004820152600b60248201526a16915493d7d05353d5539560aa1b6044820152606490fd5b346100cd5760403660031901126100cd576105e66106d3565b336001600160a01b03821603610602576103cf9060043561080d565b63334bd91960e11b5f5260045ffd5b346100cd5760403660031901126100cd576103cf6004356106306106d3565b906106496103c5825f525f602052600160405f20015490565b610785565b346100cd5760203660031901126100cd5760206106786004355f525f602052600160405f20015490565b604051908152f35b346100cd5760203660031901126100cd576004359063ffffffff60e01b82168092036100cd57602091637965db0b60e01b81149081156106c2575b5015158152f35b6301ffc9a760e01b149050836106bb565b602435906001600160a01b03821682036100cd57565b600435906001600160a01b03821682036100cd57565b90601f8019910116810190811067ffffffffffffffff82111761072157604052565b634e487b7160e01b5f52604160045260245ffd5b908160209103126100cd575180151581036100cd5790565b5f8181526020818152604080832033845290915290205460ff161561076f5750565b63e2517d3f60e01b5f523360045260245260445ffd5b5f818152602081815260408083206001600160a01b038616845290915290205460ff16610807575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50505f90565b5f818152602081815260408083206001600160a01b038616845290915290205460ff1615610807575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a460019056fea2646970667358221220104696703a8fb544f2ad189d511ce41b2582e48243cfb6eb880e3f825a20381464736f6c63430008210033ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
}

// TokenVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenVaultMetaData.ABI instead.
var TokenVaultABI = TokenVaultMetaData.ABI

// TokenVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenVaultMetaData.Bin instead.
var TokenVaultBin = TokenVaultMetaData.Bin

// DeployTokenVault deploys a new Ethereum contract, binding an instance of TokenVault to it.
func DeployTokenVault(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenVault, error) {
	parsed, err := TokenVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenVaultBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenVault{TokenVaultCaller: TokenVaultCaller{contract: contract}, TokenVaultTransactor: TokenVaultTransactor{contract: contract}, TokenVaultFilterer: TokenVaultFilterer{contract: contract}}, nil
}

// TokenVault is an auto generated Go binding around an Ethereum contract.
type TokenVault struct {
	TokenVaultCaller     // Read-only binding to the contract
	TokenVaultTransactor // Write-only binding to the contract
	TokenVaultFilterer   // Log filterer for contract events
}

// TokenVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenVaultSession struct {
	Contract     *TokenVault       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenVaultCallerSession struct {
	Contract *TokenVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TokenVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenVaultTransactorSession struct {
	Contract     *TokenVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenVaultRaw struct {
	Contract *TokenVault // Generic contract binding to access the raw methods on
}

// TokenVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenVaultCallerRaw struct {
	Contract *TokenVaultCaller // Generic read-only contract binding to access the raw methods on
}

// TokenVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenVaultTransactorRaw struct {
	Contract *TokenVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenVault creates a new instance of TokenVault, bound to a specific deployed contract.
func NewTokenVault(address common.Address, backend bind.ContractBackend) (*TokenVault, error) {
	contract, err := bindTokenVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenVault{TokenVaultCaller: TokenVaultCaller{contract: contract}, TokenVaultTransactor: TokenVaultTransactor{contract: contract}, TokenVaultFilterer: TokenVaultFilterer{contract: contract}}, nil
}

// NewTokenVaultCaller creates a new read-only instance of TokenVault, bound to a specific deployed contract.
func NewTokenVaultCaller(address common.Address, caller bind.ContractCaller) (*TokenVaultCaller, error) {
	contract, err := bindTokenVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenVaultCaller{contract: contract}, nil
}

// NewTokenVaultTransactor creates a new write-only instance of TokenVault, bound to a specific deployed contract.
func NewTokenVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenVaultTransactor, error) {
	contract, err := bindTokenVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenVaultTransactor{contract: contract}, nil
}

// NewTokenVaultFilterer creates a new log filterer instance of TokenVault, bound to a specific deployed contract.
func NewTokenVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenVaultFilterer, error) {
	contract, err := bindTokenVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenVaultFilterer{contract: contract}, nil
}

// bindTokenVault binds a generic wrapper to an already deployed contract.
func bindTokenVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenVault *TokenVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenVault.Contract.TokenVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenVault *TokenVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenVault.Contract.TokenVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenVault *TokenVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenVault.Contract.TokenVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenVault *TokenVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenVault *TokenVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenVault *TokenVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenVault.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TokenVault *TokenVaultCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenVault.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TokenVault *TokenVaultSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TokenVault.Contract.DEFAULTADMINROLE(&_TokenVault.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TokenVault *TokenVaultCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TokenVault.Contract.DEFAULTADMINROLE(&_TokenVault.CallOpts)
}

// VAULTMANAGERROLE is a free data retrieval call binding the contract method 0x5e5a24a4.
//
// Solidity: function VAULT_MANAGER_ROLE() view returns(bytes32)
func (_TokenVault *TokenVaultCaller) VAULTMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenVault.contract.Call(opts, &out, "VAULT_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VAULTMANAGERROLE is a free data retrieval call binding the contract method 0x5e5a24a4.
//
// Solidity: function VAULT_MANAGER_ROLE() view returns(bytes32)
func (_TokenVault *TokenVaultSession) VAULTMANAGERROLE() ([32]byte, error) {
	return _TokenVault.Contract.VAULTMANAGERROLE(&_TokenVault.CallOpts)
}

// VAULTMANAGERROLE is a free data retrieval call binding the contract method 0x5e5a24a4.
//
// Solidity: function VAULT_MANAGER_ROLE() view returns(bytes32)
func (_TokenVault *TokenVaultCallerSession) VAULTMANAGERROLE() ([32]byte, error) {
	return _TokenVault.Contract.VAULTMANAGERROLE(&_TokenVault.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_TokenVault *TokenVaultCaller) GetBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenVault.contract.Call(opts, &out, "getBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_TokenVault *TokenVaultSession) GetBalance(token common.Address) (*big.Int, error) {
	return _TokenVault.Contract.GetBalance(&_TokenVault.CallOpts, token)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_TokenVault *TokenVaultCallerSession) GetBalance(token common.Address) (*big.Int, error) {
	return _TokenVault.Contract.GetBalance(&_TokenVault.CallOpts, token)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TokenVault *TokenVaultCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TokenVault.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TokenVault *TokenVaultSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TokenVault.Contract.GetRoleAdmin(&_TokenVault.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TokenVault *TokenVaultCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TokenVault.Contract.GetRoleAdmin(&_TokenVault.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TokenVault *TokenVaultCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TokenVault.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TokenVault *TokenVaultSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TokenVault.Contract.HasRole(&_TokenVault.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TokenVault *TokenVaultCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TokenVault.Contract.HasRole(&_TokenVault.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenVault *TokenVaultCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TokenVault.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenVault *TokenVaultSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenVault.Contract.SupportsInterface(&_TokenVault.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenVault *TokenVaultCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenVault.Contract.SupportsInterface(&_TokenVault.CallOpts, interfaceId)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_TokenVault *TokenVaultTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenVault.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_TokenVault *TokenVaultSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenVault.Contract.Deposit(&_TokenVault.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_TokenVault *TokenVaultTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenVault.Contract.Deposit(&_TokenVault.TransactOpts, token, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TokenVault *TokenVaultTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenVault.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TokenVault *TokenVaultSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenVault.Contract.GrantRole(&_TokenVault.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TokenVault *TokenVaultTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenVault.Contract.GrantRole(&_TokenVault.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TokenVault *TokenVaultTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TokenVault.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TokenVault *TokenVaultSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TokenVault.Contract.RenounceRole(&_TokenVault.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TokenVault *TokenVaultTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TokenVault.Contract.RenounceRole(&_TokenVault.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TokenVault *TokenVaultTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenVault.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TokenVault *TokenVaultSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenVault.Contract.RevokeRole(&_TokenVault.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TokenVault *TokenVaultTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenVault.Contract.RevokeRole(&_TokenVault.TransactOpts, role, account)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address token, uint256 amount) returns()
func (_TokenVault *TokenVaultTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenVault.contract.Transact(opts, "withdraw", token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address token, uint256 amount) returns()
func (_TokenVault *TokenVaultSession) Withdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenVault.Contract.Withdraw(&_TokenVault.TransactOpts, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address token, uint256 amount) returns()
func (_TokenVault *TokenVaultTransactorSession) Withdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenVault.Contract.Withdraw(&_TokenVault.TransactOpts, token, amount)
}

// TokenVaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the TokenVault contract.
type TokenVaultDepositIterator struct {
	Event *TokenVaultDeposit // Event containing the contract specifics and raw log

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
func (it *TokenVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVaultDeposit)
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
		it.Event = new(TokenVaultDeposit)
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
func (it *TokenVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVaultDeposit represents a Deposit event raised by the TokenVault contract.
type TokenVaultDeposit struct {
	From       common.Address
	Token      common.Address
	Amount     *big.Int
	NewBalance *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed from, address indexed token, uint256 amount, uint256 newBalance)
func (_TokenVault *TokenVaultFilterer) FilterDeposit(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*TokenVaultDepositIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenVault.contract.FilterLogs(opts, "Deposit", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenVaultDepositIterator{contract: _TokenVault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed from, address indexed token, uint256 amount, uint256 newBalance)
func (_TokenVault *TokenVaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *TokenVaultDeposit, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenVault.contract.WatchLogs(opts, "Deposit", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVaultDeposit)
				if err := _TokenVault.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed from, address indexed token, uint256 amount, uint256 newBalance)
func (_TokenVault *TokenVaultFilterer) ParseDeposit(log types.Log) (*TokenVaultDeposit, error) {
	event := new(TokenVaultDeposit)
	if err := _TokenVault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenVaultRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TokenVault contract.
type TokenVaultRoleAdminChangedIterator struct {
	Event *TokenVaultRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TokenVaultRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVaultRoleAdminChanged)
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
		it.Event = new(TokenVaultRoleAdminChanged)
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
func (it *TokenVaultRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVaultRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVaultRoleAdminChanged represents a RoleAdminChanged event raised by the TokenVault contract.
type TokenVaultRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TokenVault *TokenVaultFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TokenVaultRoleAdminChangedIterator, error) {

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

	logs, sub, err := _TokenVault.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TokenVaultRoleAdminChangedIterator{contract: _TokenVault.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TokenVault *TokenVaultFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TokenVaultRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _TokenVault.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVaultRoleAdminChanged)
				if err := _TokenVault.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_TokenVault *TokenVaultFilterer) ParseRoleAdminChanged(log types.Log) (*TokenVaultRoleAdminChanged, error) {
	event := new(TokenVaultRoleAdminChanged)
	if err := _TokenVault.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenVaultRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TokenVault contract.
type TokenVaultRoleGrantedIterator struct {
	Event *TokenVaultRoleGranted // Event containing the contract specifics and raw log

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
func (it *TokenVaultRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVaultRoleGranted)
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
		it.Event = new(TokenVaultRoleGranted)
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
func (it *TokenVaultRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVaultRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVaultRoleGranted represents a RoleGranted event raised by the TokenVault contract.
type TokenVaultRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenVault *TokenVaultFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TokenVaultRoleGrantedIterator, error) {

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

	logs, sub, err := _TokenVault.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenVaultRoleGrantedIterator{contract: _TokenVault.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenVault *TokenVaultFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TokenVaultRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TokenVault.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVaultRoleGranted)
				if err := _TokenVault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_TokenVault *TokenVaultFilterer) ParseRoleGranted(log types.Log) (*TokenVaultRoleGranted, error) {
	event := new(TokenVaultRoleGranted)
	if err := _TokenVault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenVaultRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TokenVault contract.
type TokenVaultRoleRevokedIterator struct {
	Event *TokenVaultRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TokenVaultRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVaultRoleRevoked)
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
		it.Event = new(TokenVaultRoleRevoked)
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
func (it *TokenVaultRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVaultRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVaultRoleRevoked represents a RoleRevoked event raised by the TokenVault contract.
type TokenVaultRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenVault *TokenVaultFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TokenVaultRoleRevokedIterator, error) {

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

	logs, sub, err := _TokenVault.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenVaultRoleRevokedIterator{contract: _TokenVault.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenVault *TokenVaultFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TokenVaultRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TokenVault.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVaultRoleRevoked)
				if err := _TokenVault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_TokenVault *TokenVaultFilterer) ParseRoleRevoked(log types.Log) (*TokenVaultRoleRevoked, error) {
	event := new(TokenVaultRoleRevoked)
	if err := _TokenVault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenVaultTreasuryWithdrawnIterator is returned from FilterTreasuryWithdrawn and is used to iterate over the raw logs and unpacked data for TreasuryWithdrawn events raised by the TokenVault contract.
type TokenVaultTreasuryWithdrawnIterator struct {
	Event *TokenVaultTreasuryWithdrawn // Event containing the contract specifics and raw log

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
func (it *TokenVaultTreasuryWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVaultTreasuryWithdrawn)
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
		it.Event = new(TokenVaultTreasuryWithdrawn)
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
func (it *TokenVaultTreasuryWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVaultTreasuryWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVaultTreasuryWithdrawn represents a TreasuryWithdrawn event raised by the TokenVault contract.
type TokenVaultTreasuryWithdrawn struct {
	DaoId  *big.Int
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTreasuryWithdrawn is a free log retrieval operation binding the contract event 0x3994603528a4ce7136b8ccd2cc94e047aa196fd8a1f735323028f632ad214dfb.
//
// Solidity: event TreasuryWithdrawn(uint256 indexed daoId, address to, address token, uint256 amount)
func (_TokenVault *TokenVaultFilterer) FilterTreasuryWithdrawn(opts *bind.FilterOpts, daoId []*big.Int) (*TokenVaultTreasuryWithdrawnIterator, error) {

	var daoIdRule []interface{}
	for _, daoIdItem := range daoId {
		daoIdRule = append(daoIdRule, daoIdItem)
	}

	logs, sub, err := _TokenVault.contract.FilterLogs(opts, "TreasuryWithdrawn", daoIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenVaultTreasuryWithdrawnIterator{contract: _TokenVault.contract, event: "TreasuryWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTreasuryWithdrawn is a free log subscription operation binding the contract event 0x3994603528a4ce7136b8ccd2cc94e047aa196fd8a1f735323028f632ad214dfb.
//
// Solidity: event TreasuryWithdrawn(uint256 indexed daoId, address to, address token, uint256 amount)
func (_TokenVault *TokenVaultFilterer) WatchTreasuryWithdrawn(opts *bind.WatchOpts, sink chan<- *TokenVaultTreasuryWithdrawn, daoId []*big.Int) (event.Subscription, error) {

	var daoIdRule []interface{}
	for _, daoIdItem := range daoId {
		daoIdRule = append(daoIdRule, daoIdItem)
	}

	logs, sub, err := _TokenVault.contract.WatchLogs(opts, "TreasuryWithdrawn", daoIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVaultTreasuryWithdrawn)
				if err := _TokenVault.contract.UnpackLog(event, "TreasuryWithdrawn", log); err != nil {
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
func (_TokenVault *TokenVaultFilterer) ParseTreasuryWithdrawn(log types.Log) (*TokenVaultTreasuryWithdrawn, error) {
	event := new(TokenVaultTreasuryWithdrawn)
	if err := _TokenVault.contract.UnpackLog(event, "TreasuryWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenVaultWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the TokenVault contract.
type TokenVaultWithdrawalIterator struct {
	Event *TokenVaultWithdrawal // Event containing the contract specifics and raw log

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
func (it *TokenVaultWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVaultWithdrawal)
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
		it.Event = new(TokenVaultWithdrawal)
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
func (it *TokenVaultWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVaultWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVaultWithdrawal represents a Withdrawal event raised by the TokenVault contract.
type TokenVaultWithdrawal struct {
	To         common.Address
	Token      common.Address
	Amount     *big.Int
	NewBalance *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xc2b4a290c20fb28939d29f102514fbffd2b73c059ffba8b78250c94161d5fcc6.
//
// Solidity: event Withdrawal(address indexed to, address indexed token, uint256 amount, uint256 newBalance)
func (_TokenVault *TokenVaultFilterer) FilterWithdrawal(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*TokenVaultWithdrawalIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenVault.contract.FilterLogs(opts, "Withdrawal", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenVaultWithdrawalIterator{contract: _TokenVault.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xc2b4a290c20fb28939d29f102514fbffd2b73c059ffba8b78250c94161d5fcc6.
//
// Solidity: event Withdrawal(address indexed to, address indexed token, uint256 amount, uint256 newBalance)
func (_TokenVault *TokenVaultFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *TokenVaultWithdrawal, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenVault.contract.WatchLogs(opts, "Withdrawal", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVaultWithdrawal)
				if err := _TokenVault.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0xc2b4a290c20fb28939d29f102514fbffd2b73c059ffba8b78250c94161d5fcc6.
//
// Solidity: event Withdrawal(address indexed to, address indexed token, uint256 amount, uint256 newBalance)
func (_TokenVault *TokenVaultFilterer) ParseWithdrawal(log types.Log) (*TokenVaultWithdrawal, error) {
	event := new(TokenVaultWithdrawal)
	if err := _TokenVault.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

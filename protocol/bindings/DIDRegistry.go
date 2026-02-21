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

// DIDRegistryMetaData contains all meta data concerning the DIDRegistry contract.
var DIDRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRIVY_LINKER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RISK_UPDATER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createDID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getPrivyCredentialHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"privyCredentialHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getRiskProfileScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"privyCredentialHash\",\"type\":\"bytes32\"}],\"name\":\"linkPrivyCredential\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newScore\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"riskProfileHash\",\"type\":\"bytes32\"}],\"name\":\"updateRiskProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"name\":\"DIDCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"credentialHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linkedAt\",\"type\":\"uint256\"}],\"name\":\"PrivyCredentialLinked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newScore\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"profileHash\",\"type\":\"bytes32\"}],\"name\":\"RiskProfileUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DIDAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DIDNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"}],\"name\":\"InvalidScore\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"}]",
	Bin: "0x608060405234601f57600f336023565b5060405161090190816100ac8239f35b5f80fd5b6001600160a01b0381165f9081525f5160206109ad5f395f51905f52602052604090205460ff1660a6576001600160a01b03165f8181525f5160206109ad5f395f51905f5260205260408120805460ff191660011790553391907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b505f9056fe6080806040526004361015610012575f80fd5b5f3560e01c90816301ffc9a71461070c57508063248a9ca3146106da5780632f2ff15d1461069d5780633203625d1461064d57806336568abe1461060957806362f413cf146104c35780636409ae84146103c557806391d148541461037d5780639e4a940714610318578063a217fddf146102fe578063d547741f146102ba578063e0e1253914610179578063e4f424741461013f578063f6a3d24e146100ff5763fa8bd459146100c1575f80fd5b346100fb575f3660031901126100fb5760206040517fbf2465d25933abb1a6f03cf38e39796fd203c6e9e02e7dfaf081f459539067278152f35b5f80fd5b346100fb5760203660031901126100fb576001600160a01b03610120610775565b165f526001602052602060ff600460405f200154166040519015158152f35b346100fb575f3660031901126100fb5760206040517f3ef93660fe76799d6ebfbeffaa0460c881987fe2056f06796753bd21da5d04598152f35b346100fb5760603660031901126100fb57610192610775565b6001600160a01b03165f8181526001602052604090206004015460443591906024359060ff16156102a757335f9081527fc308041186a58d8ecc9630df87f438e5fd0add8f95645e5e947dde3ed7b0526e602052604090205460ff16158061026f575b61026157612710811161024f5760207fa154b3f5a54004ece811c5f85e274ebd0a3bf06f638a85d81bcf9812efed92c991835f526001825280600160405f200155835f526001825284600260405f200155604051908152a3005b632c870e6b60e11b5f5260045260245ffd5b6282b42960e81b5f5260045ffd5b50335f9081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16156101f5565b50638ff3599d60e01b5f5260045260245ffd5b346100fb5760403660031901126100fb576102fc6004356102d961075f565b906102f76102f2825f525f602052600160405f20015490565b61078b565b61084b565b005b346100fb575f3660031901126100fb5760206040515f8152f35b346100fb5760203660031901126100fb576001600160a01b03610339610775565b16805f52600160205260ff600460405f200154161561036b575f5260016020526020600160405f200154604051908152f35b638ff3599d60e01b5f5260045260245ffd5b346100fb5760403660031901126100fb5761039661075f565b6004355f525f60205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b346100fb5760403660031901126100fb576103de610775565b6001600160a01b03165f81815260016020526040902060040154602435919060ff161561036b57803314158061048b575b80610453575b61026157805f5260016020528160405f20557f4596ababddc93caf74264f7c81871c1f082a8c46a6adbe67dd3498ce8292cd4c6020604051428152a3005b50335f9081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff1615610415565b50335f9081527f67696308b45e02560b1246e590b9633c98f796b60d088ddabe5552c17774459f602052604090205460ff161561040f565b346100fb5760203660031901126100fb576001600160a01b036104e4610775565b16805f52600160205260ff600460405f200154166105f75780331415806105bf575b6102615760405160a0810181811067ffffffffffffffff8211176105ab576040525f8152600460208201915f8352604081015f81526060820190428252608083019460018652865f52600160205260405f2093518455516001840155516002830155516003820155019051151560ff801983541691161790557ff4a1669bf0a95363a2cc45302648a0fd62fe025c4ab93f0555f0876926d5f7416020604051428152a2005b634e487b7160e01b5f52604160045260245ffd5b50335f9081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff1615610506565b6351bf827560e11b5f5260045260245ffd5b346100fb5760403660031901126100fb5761062261075f565b336001600160a01b0382160361063e576102fc9060043561084b565b63334bd91960e11b5f5260045ffd5b346100fb5760203660031901126100fb576001600160a01b0361066e610775565b16805f52600160205260ff600460405f200154161561036b575f526001602052602060405f2054604051908152f35b346100fb5760403660031901126100fb576102fc6004356106bc61075f565b906106d56102f2825f525f602052600160405f20015490565b6107c3565b346100fb5760203660031901126100fb5760206107046004355f525f602052600160405f20015490565b604051908152f35b346100fb5760203660031901126100fb576004359063ffffffff60e01b82168092036100fb57602091637965db0b60e01b811490811561074e575b5015158152f35b6301ffc9a760e01b14905083610747565b602435906001600160a01b03821682036100fb57565b600435906001600160a01b03821682036100fb57565b5f8181526020818152604080832033845290915290205460ff16156107ad5750565b63e2517d3f60e01b5f523360045260245260445ffd5b5f818152602081815260408083206001600160a01b038616845290915290205460ff16610845575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50505f90565b5f818152602081815260408083206001600160a01b038616845290915290205460ff1615610845575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a460019056fea264697066735822122054652628235f4b89c4840cc7615ba0bdddddf584c3bc2b9be987c73a062039d564736f6c63430008210033ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
}

// DIDRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use DIDRegistryMetaData.ABI instead.
var DIDRegistryABI = DIDRegistryMetaData.ABI

// DIDRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DIDRegistryMetaData.Bin instead.
var DIDRegistryBin = DIDRegistryMetaData.Bin

// DeployDIDRegistry deploys a new Ethereum contract, binding an instance of DIDRegistry to it.
func DeployDIDRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DIDRegistry, error) {
	parsed, err := DIDRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DIDRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DIDRegistry{DIDRegistryCaller: DIDRegistryCaller{contract: contract}, DIDRegistryTransactor: DIDRegistryTransactor{contract: contract}, DIDRegistryFilterer: DIDRegistryFilterer{contract: contract}}, nil
}

// DIDRegistry is an auto generated Go binding around an Ethereum contract.
type DIDRegistry struct {
	DIDRegistryCaller     // Read-only binding to the contract
	DIDRegistryTransactor // Write-only binding to the contract
	DIDRegistryFilterer   // Log filterer for contract events
}

// DIDRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DIDRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIDRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DIDRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIDRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DIDRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIDRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DIDRegistrySession struct {
	Contract     *DIDRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DIDRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DIDRegistryCallerSession struct {
	Contract *DIDRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DIDRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DIDRegistryTransactorSession struct {
	Contract     *DIDRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DIDRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DIDRegistryRaw struct {
	Contract *DIDRegistry // Generic contract binding to access the raw methods on
}

// DIDRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DIDRegistryCallerRaw struct {
	Contract *DIDRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// DIDRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DIDRegistryTransactorRaw struct {
	Contract *DIDRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDIDRegistry creates a new instance of DIDRegistry, bound to a specific deployed contract.
func NewDIDRegistry(address common.Address, backend bind.ContractBackend) (*DIDRegistry, error) {
	contract, err := bindDIDRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DIDRegistry{DIDRegistryCaller: DIDRegistryCaller{contract: contract}, DIDRegistryTransactor: DIDRegistryTransactor{contract: contract}, DIDRegistryFilterer: DIDRegistryFilterer{contract: contract}}, nil
}

// NewDIDRegistryCaller creates a new read-only instance of DIDRegistry, bound to a specific deployed contract.
func NewDIDRegistryCaller(address common.Address, caller bind.ContractCaller) (*DIDRegistryCaller, error) {
	contract, err := bindDIDRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DIDRegistryCaller{contract: contract}, nil
}

// NewDIDRegistryTransactor creates a new write-only instance of DIDRegistry, bound to a specific deployed contract.
func NewDIDRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*DIDRegistryTransactor, error) {
	contract, err := bindDIDRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DIDRegistryTransactor{contract: contract}, nil
}

// NewDIDRegistryFilterer creates a new log filterer instance of DIDRegistry, bound to a specific deployed contract.
func NewDIDRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*DIDRegistryFilterer, error) {
	contract, err := bindDIDRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DIDRegistryFilterer{contract: contract}, nil
}

// bindDIDRegistry binds a generic wrapper to an already deployed contract.
func bindDIDRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DIDRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIDRegistry *DIDRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DIDRegistry.Contract.DIDRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIDRegistry *DIDRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIDRegistry.Contract.DIDRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIDRegistry *DIDRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIDRegistry.Contract.DIDRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIDRegistry *DIDRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DIDRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIDRegistry *DIDRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIDRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIDRegistry *DIDRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIDRegistry.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DIDRegistry *DIDRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DIDRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DIDRegistry *DIDRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DIDRegistry.Contract.DEFAULTADMINROLE(&_DIDRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DIDRegistry *DIDRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DIDRegistry.Contract.DEFAULTADMINROLE(&_DIDRegistry.CallOpts)
}

// PRIVYLINKERROLE is a free data retrieval call binding the contract method 0xfa8bd459.
//
// Solidity: function PRIVY_LINKER_ROLE() view returns(bytes32)
func (_DIDRegistry *DIDRegistryCaller) PRIVYLINKERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DIDRegistry.contract.Call(opts, &out, "PRIVY_LINKER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PRIVYLINKERROLE is a free data retrieval call binding the contract method 0xfa8bd459.
//
// Solidity: function PRIVY_LINKER_ROLE() view returns(bytes32)
func (_DIDRegistry *DIDRegistrySession) PRIVYLINKERROLE() ([32]byte, error) {
	return _DIDRegistry.Contract.PRIVYLINKERROLE(&_DIDRegistry.CallOpts)
}

// PRIVYLINKERROLE is a free data retrieval call binding the contract method 0xfa8bd459.
//
// Solidity: function PRIVY_LINKER_ROLE() view returns(bytes32)
func (_DIDRegistry *DIDRegistryCallerSession) PRIVYLINKERROLE() ([32]byte, error) {
	return _DIDRegistry.Contract.PRIVYLINKERROLE(&_DIDRegistry.CallOpts)
}

// RISKUPDATERROLE is a free data retrieval call binding the contract method 0xe4f42474.
//
// Solidity: function RISK_UPDATER_ROLE() view returns(bytes32)
func (_DIDRegistry *DIDRegistryCaller) RISKUPDATERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DIDRegistry.contract.Call(opts, &out, "RISK_UPDATER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RISKUPDATERROLE is a free data retrieval call binding the contract method 0xe4f42474.
//
// Solidity: function RISK_UPDATER_ROLE() view returns(bytes32)
func (_DIDRegistry *DIDRegistrySession) RISKUPDATERROLE() ([32]byte, error) {
	return _DIDRegistry.Contract.RISKUPDATERROLE(&_DIDRegistry.CallOpts)
}

// RISKUPDATERROLE is a free data retrieval call binding the contract method 0xe4f42474.
//
// Solidity: function RISK_UPDATER_ROLE() view returns(bytes32)
func (_DIDRegistry *DIDRegistryCallerSession) RISKUPDATERROLE() ([32]byte, error) {
	return _DIDRegistry.Contract.RISKUPDATERROLE(&_DIDRegistry.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address owner) view returns(bool)
func (_DIDRegistry *DIDRegistryCaller) Exists(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _DIDRegistry.contract.Call(opts, &out, "exists", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address owner) view returns(bool)
func (_DIDRegistry *DIDRegistrySession) Exists(owner common.Address) (bool, error) {
	return _DIDRegistry.Contract.Exists(&_DIDRegistry.CallOpts, owner)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address owner) view returns(bool)
func (_DIDRegistry *DIDRegistryCallerSession) Exists(owner common.Address) (bool, error) {
	return _DIDRegistry.Contract.Exists(&_DIDRegistry.CallOpts, owner)
}

// GetPrivyCredentialHash is a free data retrieval call binding the contract method 0x3203625d.
//
// Solidity: function getPrivyCredentialHash(address owner) view returns(bytes32 privyCredentialHash)
func (_DIDRegistry *DIDRegistryCaller) GetPrivyCredentialHash(opts *bind.CallOpts, owner common.Address) ([32]byte, error) {
	var out []interface{}
	err := _DIDRegistry.contract.Call(opts, &out, "getPrivyCredentialHash", owner)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPrivyCredentialHash is a free data retrieval call binding the contract method 0x3203625d.
//
// Solidity: function getPrivyCredentialHash(address owner) view returns(bytes32 privyCredentialHash)
func (_DIDRegistry *DIDRegistrySession) GetPrivyCredentialHash(owner common.Address) ([32]byte, error) {
	return _DIDRegistry.Contract.GetPrivyCredentialHash(&_DIDRegistry.CallOpts, owner)
}

// GetPrivyCredentialHash is a free data retrieval call binding the contract method 0x3203625d.
//
// Solidity: function getPrivyCredentialHash(address owner) view returns(bytes32 privyCredentialHash)
func (_DIDRegistry *DIDRegistryCallerSession) GetPrivyCredentialHash(owner common.Address) ([32]byte, error) {
	return _DIDRegistry.Contract.GetPrivyCredentialHash(&_DIDRegistry.CallOpts, owner)
}

// GetRiskProfileScore is a free data retrieval call binding the contract method 0x9e4a9407.
//
// Solidity: function getRiskProfileScore(address owner) view returns(uint256)
func (_DIDRegistry *DIDRegistryCaller) GetRiskProfileScore(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DIDRegistry.contract.Call(opts, &out, "getRiskProfileScore", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRiskProfileScore is a free data retrieval call binding the contract method 0x9e4a9407.
//
// Solidity: function getRiskProfileScore(address owner) view returns(uint256)
func (_DIDRegistry *DIDRegistrySession) GetRiskProfileScore(owner common.Address) (*big.Int, error) {
	return _DIDRegistry.Contract.GetRiskProfileScore(&_DIDRegistry.CallOpts, owner)
}

// GetRiskProfileScore is a free data retrieval call binding the contract method 0x9e4a9407.
//
// Solidity: function getRiskProfileScore(address owner) view returns(uint256)
func (_DIDRegistry *DIDRegistryCallerSession) GetRiskProfileScore(owner common.Address) (*big.Int, error) {
	return _DIDRegistry.Contract.GetRiskProfileScore(&_DIDRegistry.CallOpts, owner)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DIDRegistry *DIDRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _DIDRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DIDRegistry *DIDRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DIDRegistry.Contract.GetRoleAdmin(&_DIDRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DIDRegistry *DIDRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DIDRegistry.Contract.GetRoleAdmin(&_DIDRegistry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DIDRegistry *DIDRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _DIDRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DIDRegistry *DIDRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DIDRegistry.Contract.HasRole(&_DIDRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DIDRegistry *DIDRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DIDRegistry.Contract.HasRole(&_DIDRegistry.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DIDRegistry *DIDRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DIDRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DIDRegistry *DIDRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DIDRegistry.Contract.SupportsInterface(&_DIDRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DIDRegistry *DIDRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DIDRegistry.Contract.SupportsInterface(&_DIDRegistry.CallOpts, interfaceId)
}

// CreateDID is a paid mutator transaction binding the contract method 0x62f413cf.
//
// Solidity: function createDID(address owner) returns()
func (_DIDRegistry *DIDRegistryTransactor) CreateDID(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _DIDRegistry.contract.Transact(opts, "createDID", owner)
}

// CreateDID is a paid mutator transaction binding the contract method 0x62f413cf.
//
// Solidity: function createDID(address owner) returns()
func (_DIDRegistry *DIDRegistrySession) CreateDID(owner common.Address) (*types.Transaction, error) {
	return _DIDRegistry.Contract.CreateDID(&_DIDRegistry.TransactOpts, owner)
}

// CreateDID is a paid mutator transaction binding the contract method 0x62f413cf.
//
// Solidity: function createDID(address owner) returns()
func (_DIDRegistry *DIDRegistryTransactorSession) CreateDID(owner common.Address) (*types.Transaction, error) {
	return _DIDRegistry.Contract.CreateDID(&_DIDRegistry.TransactOpts, owner)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DIDRegistry *DIDRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DIDRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DIDRegistry *DIDRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DIDRegistry.Contract.GrantRole(&_DIDRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DIDRegistry *DIDRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DIDRegistry.Contract.GrantRole(&_DIDRegistry.TransactOpts, role, account)
}

// LinkPrivyCredential is a paid mutator transaction binding the contract method 0x6409ae84.
//
// Solidity: function linkPrivyCredential(address owner, bytes32 privyCredentialHash) returns()
func (_DIDRegistry *DIDRegistryTransactor) LinkPrivyCredential(opts *bind.TransactOpts, owner common.Address, privyCredentialHash [32]byte) (*types.Transaction, error) {
	return _DIDRegistry.contract.Transact(opts, "linkPrivyCredential", owner, privyCredentialHash)
}

// LinkPrivyCredential is a paid mutator transaction binding the contract method 0x6409ae84.
//
// Solidity: function linkPrivyCredential(address owner, bytes32 privyCredentialHash) returns()
func (_DIDRegistry *DIDRegistrySession) LinkPrivyCredential(owner common.Address, privyCredentialHash [32]byte) (*types.Transaction, error) {
	return _DIDRegistry.Contract.LinkPrivyCredential(&_DIDRegistry.TransactOpts, owner, privyCredentialHash)
}

// LinkPrivyCredential is a paid mutator transaction binding the contract method 0x6409ae84.
//
// Solidity: function linkPrivyCredential(address owner, bytes32 privyCredentialHash) returns()
func (_DIDRegistry *DIDRegistryTransactorSession) LinkPrivyCredential(owner common.Address, privyCredentialHash [32]byte) (*types.Transaction, error) {
	return _DIDRegistry.Contract.LinkPrivyCredential(&_DIDRegistry.TransactOpts, owner, privyCredentialHash)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_DIDRegistry *DIDRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _DIDRegistry.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_DIDRegistry *DIDRegistrySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _DIDRegistry.Contract.RenounceRole(&_DIDRegistry.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_DIDRegistry *DIDRegistryTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _DIDRegistry.Contract.RenounceRole(&_DIDRegistry.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DIDRegistry *DIDRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DIDRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DIDRegistry *DIDRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DIDRegistry.Contract.RevokeRole(&_DIDRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DIDRegistry *DIDRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DIDRegistry.Contract.RevokeRole(&_DIDRegistry.TransactOpts, role, account)
}

// UpdateRiskProfile is a paid mutator transaction binding the contract method 0xe0e12539.
//
// Solidity: function updateRiskProfile(address owner, uint256 newScore, bytes32 riskProfileHash) returns()
func (_DIDRegistry *DIDRegistryTransactor) UpdateRiskProfile(opts *bind.TransactOpts, owner common.Address, newScore *big.Int, riskProfileHash [32]byte) (*types.Transaction, error) {
	return _DIDRegistry.contract.Transact(opts, "updateRiskProfile", owner, newScore, riskProfileHash)
}

// UpdateRiskProfile is a paid mutator transaction binding the contract method 0xe0e12539.
//
// Solidity: function updateRiskProfile(address owner, uint256 newScore, bytes32 riskProfileHash) returns()
func (_DIDRegistry *DIDRegistrySession) UpdateRiskProfile(owner common.Address, newScore *big.Int, riskProfileHash [32]byte) (*types.Transaction, error) {
	return _DIDRegistry.Contract.UpdateRiskProfile(&_DIDRegistry.TransactOpts, owner, newScore, riskProfileHash)
}

// UpdateRiskProfile is a paid mutator transaction binding the contract method 0xe0e12539.
//
// Solidity: function updateRiskProfile(address owner, uint256 newScore, bytes32 riskProfileHash) returns()
func (_DIDRegistry *DIDRegistryTransactorSession) UpdateRiskProfile(owner common.Address, newScore *big.Int, riskProfileHash [32]byte) (*types.Transaction, error) {
	return _DIDRegistry.Contract.UpdateRiskProfile(&_DIDRegistry.TransactOpts, owner, newScore, riskProfileHash)
}

// DIDRegistryDIDCreatedIterator is returned from FilterDIDCreated and is used to iterate over the raw logs and unpacked data for DIDCreated events raised by the DIDRegistry contract.
type DIDRegistryDIDCreatedIterator struct {
	Event *DIDRegistryDIDCreated // Event containing the contract specifics and raw log

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
func (it *DIDRegistryDIDCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIDRegistryDIDCreated)
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
		it.Event = new(DIDRegistryDIDCreated)
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
func (it *DIDRegistryDIDCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIDRegistryDIDCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIDRegistryDIDCreated represents a DIDCreated event raised by the DIDRegistry contract.
type DIDRegistryDIDCreated struct {
	Owner     common.Address
	CreatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDIDCreated is a free log retrieval operation binding the contract event 0xf4a1669bf0a95363a2cc45302648a0fd62fe025c4ab93f0555f0876926d5f741.
//
// Solidity: event DIDCreated(address indexed owner, uint256 createdAt)
func (_DIDRegistry *DIDRegistryFilterer) FilterDIDCreated(opts *bind.FilterOpts, owner []common.Address) (*DIDRegistryDIDCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _DIDRegistry.contract.FilterLogs(opts, "DIDCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &DIDRegistryDIDCreatedIterator{contract: _DIDRegistry.contract, event: "DIDCreated", logs: logs, sub: sub}, nil
}

// WatchDIDCreated is a free log subscription operation binding the contract event 0xf4a1669bf0a95363a2cc45302648a0fd62fe025c4ab93f0555f0876926d5f741.
//
// Solidity: event DIDCreated(address indexed owner, uint256 createdAt)
func (_DIDRegistry *DIDRegistryFilterer) WatchDIDCreated(opts *bind.WatchOpts, sink chan<- *DIDRegistryDIDCreated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _DIDRegistry.contract.WatchLogs(opts, "DIDCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIDRegistryDIDCreated)
				if err := _DIDRegistry.contract.UnpackLog(event, "DIDCreated", log); err != nil {
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

// ParseDIDCreated is a log parse operation binding the contract event 0xf4a1669bf0a95363a2cc45302648a0fd62fe025c4ab93f0555f0876926d5f741.
//
// Solidity: event DIDCreated(address indexed owner, uint256 createdAt)
func (_DIDRegistry *DIDRegistryFilterer) ParseDIDCreated(log types.Log) (*DIDRegistryDIDCreated, error) {
	event := new(DIDRegistryDIDCreated)
	if err := _DIDRegistry.contract.UnpackLog(event, "DIDCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DIDRegistryPrivyCredentialLinkedIterator is returned from FilterPrivyCredentialLinked and is used to iterate over the raw logs and unpacked data for PrivyCredentialLinked events raised by the DIDRegistry contract.
type DIDRegistryPrivyCredentialLinkedIterator struct {
	Event *DIDRegistryPrivyCredentialLinked // Event containing the contract specifics and raw log

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
func (it *DIDRegistryPrivyCredentialLinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIDRegistryPrivyCredentialLinked)
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
		it.Event = new(DIDRegistryPrivyCredentialLinked)
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
func (it *DIDRegistryPrivyCredentialLinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIDRegistryPrivyCredentialLinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIDRegistryPrivyCredentialLinked represents a PrivyCredentialLinked event raised by the DIDRegistry contract.
type DIDRegistryPrivyCredentialLinked struct {
	Owner          common.Address
	CredentialHash [32]byte
	LinkedAt       *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPrivyCredentialLinked is a free log retrieval operation binding the contract event 0x4596ababddc93caf74264f7c81871c1f082a8c46a6adbe67dd3498ce8292cd4c.
//
// Solidity: event PrivyCredentialLinked(address indexed owner, bytes32 indexed credentialHash, uint256 linkedAt)
func (_DIDRegistry *DIDRegistryFilterer) FilterPrivyCredentialLinked(opts *bind.FilterOpts, owner []common.Address, credentialHash [][32]byte) (*DIDRegistryPrivyCredentialLinkedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var credentialHashRule []interface{}
	for _, credentialHashItem := range credentialHash {
		credentialHashRule = append(credentialHashRule, credentialHashItem)
	}

	logs, sub, err := _DIDRegistry.contract.FilterLogs(opts, "PrivyCredentialLinked", ownerRule, credentialHashRule)
	if err != nil {
		return nil, err
	}
	return &DIDRegistryPrivyCredentialLinkedIterator{contract: _DIDRegistry.contract, event: "PrivyCredentialLinked", logs: logs, sub: sub}, nil
}

// WatchPrivyCredentialLinked is a free log subscription operation binding the contract event 0x4596ababddc93caf74264f7c81871c1f082a8c46a6adbe67dd3498ce8292cd4c.
//
// Solidity: event PrivyCredentialLinked(address indexed owner, bytes32 indexed credentialHash, uint256 linkedAt)
func (_DIDRegistry *DIDRegistryFilterer) WatchPrivyCredentialLinked(opts *bind.WatchOpts, sink chan<- *DIDRegistryPrivyCredentialLinked, owner []common.Address, credentialHash [][32]byte) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var credentialHashRule []interface{}
	for _, credentialHashItem := range credentialHash {
		credentialHashRule = append(credentialHashRule, credentialHashItem)
	}

	logs, sub, err := _DIDRegistry.contract.WatchLogs(opts, "PrivyCredentialLinked", ownerRule, credentialHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIDRegistryPrivyCredentialLinked)
				if err := _DIDRegistry.contract.UnpackLog(event, "PrivyCredentialLinked", log); err != nil {
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

// ParsePrivyCredentialLinked is a log parse operation binding the contract event 0x4596ababddc93caf74264f7c81871c1f082a8c46a6adbe67dd3498ce8292cd4c.
//
// Solidity: event PrivyCredentialLinked(address indexed owner, bytes32 indexed credentialHash, uint256 linkedAt)
func (_DIDRegistry *DIDRegistryFilterer) ParsePrivyCredentialLinked(log types.Log) (*DIDRegistryPrivyCredentialLinked, error) {
	event := new(DIDRegistryPrivyCredentialLinked)
	if err := _DIDRegistry.contract.UnpackLog(event, "PrivyCredentialLinked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DIDRegistryRiskProfileUpdatedIterator is returned from FilterRiskProfileUpdated and is used to iterate over the raw logs and unpacked data for RiskProfileUpdated events raised by the DIDRegistry contract.
type DIDRegistryRiskProfileUpdatedIterator struct {
	Event *DIDRegistryRiskProfileUpdated // Event containing the contract specifics and raw log

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
func (it *DIDRegistryRiskProfileUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIDRegistryRiskProfileUpdated)
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
		it.Event = new(DIDRegistryRiskProfileUpdated)
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
func (it *DIDRegistryRiskProfileUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIDRegistryRiskProfileUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIDRegistryRiskProfileUpdated represents a RiskProfileUpdated event raised by the DIDRegistry contract.
type DIDRegistryRiskProfileUpdated struct {
	Owner       common.Address
	NewScore    *big.Int
	ProfileHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRiskProfileUpdated is a free log retrieval operation binding the contract event 0xa154b3f5a54004ece811c5f85e274ebd0a3bf06f638a85d81bcf9812efed92c9.
//
// Solidity: event RiskProfileUpdated(address indexed owner, uint256 newScore, bytes32 indexed profileHash)
func (_DIDRegistry *DIDRegistryFilterer) FilterRiskProfileUpdated(opts *bind.FilterOpts, owner []common.Address, profileHash [][32]byte) (*DIDRegistryRiskProfileUpdatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var profileHashRule []interface{}
	for _, profileHashItem := range profileHash {
		profileHashRule = append(profileHashRule, profileHashItem)
	}

	logs, sub, err := _DIDRegistry.contract.FilterLogs(opts, "RiskProfileUpdated", ownerRule, profileHashRule)
	if err != nil {
		return nil, err
	}
	return &DIDRegistryRiskProfileUpdatedIterator{contract: _DIDRegistry.contract, event: "RiskProfileUpdated", logs: logs, sub: sub}, nil
}

// WatchRiskProfileUpdated is a free log subscription operation binding the contract event 0xa154b3f5a54004ece811c5f85e274ebd0a3bf06f638a85d81bcf9812efed92c9.
//
// Solidity: event RiskProfileUpdated(address indexed owner, uint256 newScore, bytes32 indexed profileHash)
func (_DIDRegistry *DIDRegistryFilterer) WatchRiskProfileUpdated(opts *bind.WatchOpts, sink chan<- *DIDRegistryRiskProfileUpdated, owner []common.Address, profileHash [][32]byte) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var profileHashRule []interface{}
	for _, profileHashItem := range profileHash {
		profileHashRule = append(profileHashRule, profileHashItem)
	}

	logs, sub, err := _DIDRegistry.contract.WatchLogs(opts, "RiskProfileUpdated", ownerRule, profileHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIDRegistryRiskProfileUpdated)
				if err := _DIDRegistry.contract.UnpackLog(event, "RiskProfileUpdated", log); err != nil {
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

// ParseRiskProfileUpdated is a log parse operation binding the contract event 0xa154b3f5a54004ece811c5f85e274ebd0a3bf06f638a85d81bcf9812efed92c9.
//
// Solidity: event RiskProfileUpdated(address indexed owner, uint256 newScore, bytes32 indexed profileHash)
func (_DIDRegistry *DIDRegistryFilterer) ParseRiskProfileUpdated(log types.Log) (*DIDRegistryRiskProfileUpdated, error) {
	event := new(DIDRegistryRiskProfileUpdated)
	if err := _DIDRegistry.contract.UnpackLog(event, "RiskProfileUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DIDRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the DIDRegistry contract.
type DIDRegistryRoleAdminChangedIterator struct {
	Event *DIDRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *DIDRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIDRegistryRoleAdminChanged)
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
		it.Event = new(DIDRegistryRoleAdminChanged)
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
func (it *DIDRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIDRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIDRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the DIDRegistry contract.
type DIDRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DIDRegistry *DIDRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DIDRegistryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _DIDRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DIDRegistryRoleAdminChangedIterator{contract: _DIDRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DIDRegistry *DIDRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DIDRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _DIDRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIDRegistryRoleAdminChanged)
				if err := _DIDRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_DIDRegistry *DIDRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*DIDRegistryRoleAdminChanged, error) {
	event := new(DIDRegistryRoleAdminChanged)
	if err := _DIDRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DIDRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the DIDRegistry contract.
type DIDRegistryRoleGrantedIterator struct {
	Event *DIDRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *DIDRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIDRegistryRoleGranted)
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
		it.Event = new(DIDRegistryRoleGranted)
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
func (it *DIDRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIDRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIDRegistryRoleGranted represents a RoleGranted event raised by the DIDRegistry contract.
type DIDRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DIDRegistry *DIDRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DIDRegistryRoleGrantedIterator, error) {

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

	logs, sub, err := _DIDRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DIDRegistryRoleGrantedIterator{contract: _DIDRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DIDRegistry *DIDRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DIDRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DIDRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIDRegistryRoleGranted)
				if err := _DIDRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_DIDRegistry *DIDRegistryFilterer) ParseRoleGranted(log types.Log) (*DIDRegistryRoleGranted, error) {
	event := new(DIDRegistryRoleGranted)
	if err := _DIDRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DIDRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the DIDRegistry contract.
type DIDRegistryRoleRevokedIterator struct {
	Event *DIDRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *DIDRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIDRegistryRoleRevoked)
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
		it.Event = new(DIDRegistryRoleRevoked)
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
func (it *DIDRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIDRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIDRegistryRoleRevoked represents a RoleRevoked event raised by the DIDRegistry contract.
type DIDRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DIDRegistry *DIDRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DIDRegistryRoleRevokedIterator, error) {

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

	logs, sub, err := _DIDRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DIDRegistryRoleRevokedIterator{contract: _DIDRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DIDRegistry *DIDRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DIDRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DIDRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIDRegistryRoleRevoked)
				if err := _DIDRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_DIDRegistry *DIDRegistryFilterer) ParseRoleRevoked(log types.Log) (*DIDRegistryRoleRevoked, error) {
	event := new(DIDRegistryRoleRevoked)
	if err := _DIDRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

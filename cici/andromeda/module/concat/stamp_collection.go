// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package concat

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

// AddressUtilsABI is the input ABI used to generate the binding from.
const AddressUtilsABI = "[]"

// AddressUtilsBin is the compiled bytecode used for deploying new contracts.
const AddressUtilsBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058205d54facd45004e05085d7382c204c3c12a5366761fefabe6d4e65d6ced3ab1d80029`

// DeployAddressUtils deploys a new Ethereum contract, binding an instance of AddressUtils to it.
func DeployAddressUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressUtils{AddressUtilsCaller: AddressUtilsCaller{contract: contract}, AddressUtilsTransactor: AddressUtilsTransactor{contract: contract}, AddressUtilsFilterer: AddressUtilsFilterer{contract: contract}}, nil
}

// AddressUtils is an auto generated Go binding around an Ethereum contract.
type AddressUtils struct {
	AddressUtilsCaller     // Read-only binding to the contract
	AddressUtilsTransactor // Write-only binding to the contract
	AddressUtilsFilterer   // Log filterer for contract events
}

// AddressUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressUtilsSession struct {
	Contract     *AddressUtils     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressUtilsCallerSession struct {
	Contract *AddressUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AddressUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressUtilsTransactorSession struct {
	Contract     *AddressUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AddressUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressUtilsRaw struct {
	Contract *AddressUtils // Generic contract binding to access the raw methods on
}

// AddressUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressUtilsCallerRaw struct {
	Contract *AddressUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// AddressUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressUtilsTransactorRaw struct {
	Contract *AddressUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressUtils creates a new instance of AddressUtils, bound to a specific deployed contract.
func NewAddressUtils(address common.Address, backend bind.ContractBackend) (*AddressUtils, error) {
	contract, err := bindAddressUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressUtils{AddressUtilsCaller: AddressUtilsCaller{contract: contract}, AddressUtilsTransactor: AddressUtilsTransactor{contract: contract}, AddressUtilsFilterer: AddressUtilsFilterer{contract: contract}}, nil
}

// NewAddressUtilsCaller creates a new read-only instance of AddressUtils, bound to a specific deployed contract.
func NewAddressUtilsCaller(address common.Address, caller bind.ContractCaller) (*AddressUtilsCaller, error) {
	contract, err := bindAddressUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUtilsCaller{contract: contract}, nil
}

// NewAddressUtilsTransactor creates a new write-only instance of AddressUtils, bound to a specific deployed contract.
func NewAddressUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressUtilsTransactor, error) {
	contract, err := bindAddressUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUtilsTransactor{contract: contract}, nil
}

// NewAddressUtilsFilterer creates a new log filterer instance of AddressUtils, bound to a specific deployed contract.
func NewAddressUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressUtilsFilterer, error) {
	contract, err := bindAddressUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressUtilsFilterer{contract: contract}, nil
}

// bindAddressUtils binds a generic wrapper to an already deployed contract.
func bindAddressUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUtils *AddressUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddressUtils.Contract.AddressUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUtils *AddressUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUtils.Contract.AddressUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUtils *AddressUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUtils.Contract.AddressUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUtils *AddressUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddressUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUtils *AddressUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUtils *AddressUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUtils.Contract.contract.Transact(opts, method, params...)
}

// ERC721ABI is the input ABI used to generate the binding from.
const ERC721ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"_exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721Bin is the compiled bytecode used for deploying new contracts.
const ERC721Bin = `0x`

// DeployERC721 deploys a new Ethereum contract, binding an instance of ERC721 to it.
func DeployERC721(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721{ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract}}, nil
}

// ERC721 is an auto generated Go binding around an Ethereum contract.
type ERC721 struct {
	ERC721Caller     // Read-only binding to the contract
	ERC721Transactor // Write-only binding to the contract
	ERC721Filterer   // Log filterer for contract events
}

// ERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721Session struct {
	Contract     *ERC721           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721CallerSession struct {
	Contract *ERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721TransactorSession struct {
	Contract     *ERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721Raw struct {
	Contract *ERC721 // Generic contract binding to access the raw methods on
}

// ERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721CallerRaw struct {
	Contract *ERC721Caller // Generic read-only contract binding to access the raw methods on
}

// ERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721TransactorRaw struct {
	Contract *ERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721 creates a new instance of ERC721, bound to a specific deployed contract.
func NewERC721(address common.Address, backend bind.ContractBackend) (*ERC721, error) {
	contract, err := bindERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721{ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract}}, nil
}

// NewERC721Caller creates a new read-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Caller(address common.Address, caller bind.ContractCaller) (*ERC721Caller, error) {
	contract, err := bindERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Caller{contract: contract}, nil
}

// NewERC721Transactor creates a new write-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC721Transactor, error) {
	contract, err := bindERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Transactor{contract: contract}, nil
}

// NewERC721Filterer creates a new log filterer instance of ERC721, bound to a specific deployed contract.
func NewERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC721Filterer, error) {
	contract, err := bindERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721Filterer{contract: contract}, nil
}

// bindERC721 binds a generic wrapper to an already deployed contract.
func bindERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.ERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721 *ERC721Caller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721 *ERC721Session) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721 *ERC721CallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, _owner)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(_exists bool)
func (_ERC721 *ERC721Caller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(_exists bool)
func (_ERC721 *ERC721Session) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721.Contract.Exists(&_ERC721.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(_exists bool)
func (_ERC721 *ERC721CallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721.Contract.Exists(&_ERC721.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721 *ERC721Caller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721 *ERC721Session) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721 *ERC721CallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721 *ERC721Caller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721 *ERC721Session) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721 *ERC721CallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_ERC721 *ERC721Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_ERC721 *ERC721Session) Name() (string, error) {
	return _ERC721.Contract.Name(&_ERC721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_ERC721 *ERC721CallerSession) Name() (string, error) {
	return _ERC721.Contract.Name(&_ERC721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721 *ERC721Caller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721 *ERC721Session) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721 *ERC721CallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, _tokenId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_ERC721 *ERC721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_ERC721 *ERC721Session) Symbol() (string, error) {
	return _ERC721.Contract.Symbol(&_ERC721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_ERC721 *ERC721CallerSession) Symbol() (string, error) {
	return _ERC721.Contract.Symbol(&_ERC721.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC721 *ERC721Caller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC721 *ERC721Session) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC721.Contract.TokenByIndex(&_ERC721.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC721 *ERC721CallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC721.Contract.TokenByIndex(&_ERC721.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(_tokenId uint256)
func (_ERC721 *ERC721Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(_tokenId uint256)
func (_ERC721 *ERC721Session) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721.Contract.TokenOfOwnerByIndex(&_ERC721.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(_tokenId uint256)
func (_ERC721 *ERC721CallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721.Contract.TokenOfOwnerByIndex(&_ERC721.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_ERC721 *ERC721Caller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_ERC721 *ERC721Session) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC721.Contract.TokenURI(&_ERC721.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_ERC721 *ERC721CallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC721.Contract.TokenURI(&_ERC721.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721 *ERC721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721 *ERC721Session) TotalSupply() (*big.Int, error) {
	return _ERC721.Contract.TotalSupply(&_ERC721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721 *ERC721CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721.Contract.TotalSupply(&_ERC721.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721 *ERC721Transactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721 *ERC721Session) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721 *ERC721TransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721 *ERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721 *ERC721Session) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721 *ERC721TransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721 *ERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721 *ERC721Session) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721 *ERC721TransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, _operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721 *ERC721Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721 *ERC721Session) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721 *ERC721TransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, _from, _to, _tokenId)
}

// ERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721 contract.
type ERC721ApprovalIterator struct {
	Event *ERC721Approval // Event containing the contract specifics and raw log

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
func (it *ERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Approval)
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
		it.Event = new(ERC721Approval)
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
func (it *ERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Approval represents a Approval event raised by the ERC721 contract.
type ERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_ERC721 *ERC721Filterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address) (*ERC721ApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalIterator{contract: _ERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_ERC721 *ERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721Approval, _owner []common.Address, _approved []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Approval)
				if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721 contract.
type ERC721ApprovalForAllIterator struct {
	Event *ERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721ApprovalForAll)
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
		it.Event = new(ERC721ApprovalForAll)
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
func (it *ERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721ApprovalForAll represents a ApprovalForAll event raised by the ERC721 contract.
type ERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721 *ERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ERC721ApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalForAllIterator{contract: _ERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721 *ERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721ApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721ApprovalForAll)
				if err := _ERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721 contract.
type ERC721TransferIterator struct {
	Event *ERC721Transfer // Event containing the contract specifics and raw log

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
func (it *ERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Transfer)
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
		it.Event = new(ERC721Transfer)
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
func (it *ERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Transfer represents a Transfer event raised by the ERC721 contract.
type ERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_ERC721 *ERC721Filterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ERC721TransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TransferIterator{contract: _ERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_ERC721 *ERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721Transfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Transfer)
				if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC721BasicABI is the input ABI used to generate the binding from.
const ERC721BasicABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"_exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721BasicBin is the compiled bytecode used for deploying new contracts.
const ERC721BasicBin = `0x`

// DeployERC721Basic deploys a new Ethereum contract, binding an instance of ERC721Basic to it.
func DeployERC721Basic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721Basic, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721BasicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721BasicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Basic{ERC721BasicCaller: ERC721BasicCaller{contract: contract}, ERC721BasicTransactor: ERC721BasicTransactor{contract: contract}, ERC721BasicFilterer: ERC721BasicFilterer{contract: contract}}, nil
}

// ERC721Basic is an auto generated Go binding around an Ethereum contract.
type ERC721Basic struct {
	ERC721BasicCaller     // Read-only binding to the contract
	ERC721BasicTransactor // Write-only binding to the contract
	ERC721BasicFilterer   // Log filterer for contract events
}

// ERC721BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721BasicSession struct {
	Contract     *ERC721Basic      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721BasicCallerSession struct {
	Contract *ERC721BasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC721BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721BasicTransactorSession struct {
	Contract     *ERC721BasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC721BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721BasicRaw struct {
	Contract *ERC721Basic // Generic contract binding to access the raw methods on
}

// ERC721BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721BasicCallerRaw struct {
	Contract *ERC721BasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721BasicTransactorRaw struct {
	Contract *ERC721BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Basic creates a new instance of ERC721Basic, bound to a specific deployed contract.
func NewERC721Basic(address common.Address, backend bind.ContractBackend) (*ERC721Basic, error) {
	contract, err := bindERC721Basic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Basic{ERC721BasicCaller: ERC721BasicCaller{contract: contract}, ERC721BasicTransactor: ERC721BasicTransactor{contract: contract}, ERC721BasicFilterer: ERC721BasicFilterer{contract: contract}}, nil
}

// NewERC721BasicCaller creates a new read-only instance of ERC721Basic, bound to a specific deployed contract.
func NewERC721BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC721BasicCaller, error) {
	contract, err := bindERC721Basic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicCaller{contract: contract}, nil
}

// NewERC721BasicTransactor creates a new write-only instance of ERC721Basic, bound to a specific deployed contract.
func NewERC721BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721BasicTransactor, error) {
	contract, err := bindERC721Basic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTransactor{contract: contract}, nil
}

// NewERC721BasicFilterer creates a new log filterer instance of ERC721Basic, bound to a specific deployed contract.
func NewERC721BasicFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721BasicFilterer, error) {
	contract, err := bindERC721Basic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicFilterer{contract: contract}, nil
}

// bindERC721Basic binds a generic wrapper to an already deployed contract.
func bindERC721Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Basic *ERC721BasicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Basic.Contract.ERC721BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Basic *ERC721BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Basic.Contract.ERC721BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Basic *ERC721BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Basic.Contract.ERC721BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Basic *ERC721BasicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Basic *ERC721BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Basic *ERC721BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Basic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721Basic *ERC721BasicCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Basic.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721Basic *ERC721BasicSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Basic.Contract.BalanceOf(&_ERC721Basic.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721Basic *ERC721BasicCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Basic.Contract.BalanceOf(&_ERC721Basic.CallOpts, _owner)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(_exists bool)
func (_ERC721Basic *ERC721BasicCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Basic.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(_exists bool)
func (_ERC721Basic *ERC721BasicSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721Basic.Contract.Exists(&_ERC721Basic.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(_exists bool)
func (_ERC721Basic *ERC721BasicCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721Basic.Contract.Exists(&_ERC721Basic.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721Basic *ERC721BasicCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Basic.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721Basic *ERC721BasicSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Basic.Contract.GetApproved(&_ERC721Basic.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721Basic *ERC721BasicCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Basic.Contract.GetApproved(&_ERC721Basic.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721Basic *ERC721BasicCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Basic.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721Basic *ERC721BasicSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Basic.Contract.IsApprovedForAll(&_ERC721Basic.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721Basic *ERC721BasicCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Basic.Contract.IsApprovedForAll(&_ERC721Basic.CallOpts, _owner, _operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721Basic *ERC721BasicCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Basic.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721Basic *ERC721BasicSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Basic.Contract.OwnerOf(&_ERC721Basic.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721Basic *ERC721BasicCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Basic.Contract.OwnerOf(&_ERC721Basic.CallOpts, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721Basic *ERC721BasicTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Basic.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721Basic *ERC721BasicSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Basic.Contract.Approve(&_ERC721Basic.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721Basic *ERC721BasicTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Basic.Contract.Approve(&_ERC721Basic.TransactOpts, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721Basic *ERC721BasicTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Basic.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721Basic *ERC721BasicSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Basic.Contract.SafeTransferFrom(&_ERC721Basic.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721Basic *ERC721BasicTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Basic.Contract.SafeTransferFrom(&_ERC721Basic.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721Basic *ERC721BasicTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Basic.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721Basic *ERC721BasicSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Basic.Contract.SetApprovalForAll(&_ERC721Basic.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721Basic *ERC721BasicTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Basic.Contract.SetApprovalForAll(&_ERC721Basic.TransactOpts, _operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721Basic *ERC721BasicTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Basic.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721Basic *ERC721BasicSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Basic.Contract.TransferFrom(&_ERC721Basic.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721Basic *ERC721BasicTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Basic.Contract.TransferFrom(&_ERC721Basic.TransactOpts, _from, _to, _tokenId)
}

// ERC721BasicApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Basic contract.
type ERC721BasicApprovalIterator struct {
	Event *ERC721BasicApproval // Event containing the contract specifics and raw log

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
func (it *ERC721BasicApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721BasicApproval)
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
		it.Event = new(ERC721BasicApproval)
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
func (it *ERC721BasicApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721BasicApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721BasicApproval represents a Approval event raised by the ERC721Basic contract.
type ERC721BasicApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_ERC721Basic *ERC721BasicFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address) (*ERC721BasicApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _ERC721Basic.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicApprovalIterator{contract: _ERC721Basic.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_ERC721Basic *ERC721BasicFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721BasicApproval, _owner []common.Address, _approved []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _ERC721Basic.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721BasicApproval)
				if err := _ERC721Basic.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC721BasicApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Basic contract.
type ERC721BasicApprovalForAllIterator struct {
	Event *ERC721BasicApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721BasicApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721BasicApprovalForAll)
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
		it.Event = new(ERC721BasicApprovalForAll)
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
func (it *ERC721BasicApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721BasicApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721BasicApprovalForAll represents a ApprovalForAll event raised by the ERC721Basic contract.
type ERC721BasicApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721Basic *ERC721BasicFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ERC721BasicApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Basic.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicApprovalForAllIterator{contract: _ERC721Basic.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721Basic *ERC721BasicFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721BasicApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Basic.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721BasicApprovalForAll)
				if err := _ERC721Basic.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ERC721BasicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Basic contract.
type ERC721BasicTransferIterator struct {
	Event *ERC721BasicTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721BasicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721BasicTransfer)
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
		it.Event = new(ERC721BasicTransfer)
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
func (it *ERC721BasicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721BasicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721BasicTransfer represents a Transfer event raised by the ERC721Basic contract.
type ERC721BasicTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_ERC721Basic *ERC721BasicFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ERC721BasicTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC721Basic.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTransferIterator{contract: _ERC721Basic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_ERC721Basic *ERC721BasicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721BasicTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC721Basic.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721BasicTransfer)
				if err := _ERC721Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC721BasicTokenABI is the input ABI used to generate the binding from.
const ERC721BasicTokenABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"approved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721BasicTokenBin is the compiled bytecode used for deploying new contracts.
const ERC721BasicTokenBin = `0x608060405234801561001057600080fd5b506109f5806100206000396000f3006080604052600436106100a35763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663081812fc81146100a8578063095ea7b3146100dc57806323b872dd1461010257806342842e0e1461012c5780634f558e79146101565780636352211e1461018257806370a082311461019a578063a22cb465146101cd578063b88d4fde146101f3578063e985e9c514610262575b600080fd5b3480156100b457600080fd5b506100c0600435610289565b60408051600160a060020a039092168252519081900360200190f35b3480156100e857600080fd5b50610100600160a060020a03600435166024356102a4565b005b34801561010e57600080fd5b50610100600160a060020a0360043581169060243516604435610396565b34801561013857600080fd5b50610100600160a060020a0360043581169060243516604435610445565b34801561016257600080fd5b5061016e60043561047d565b604080519115158252519081900360200190f35b34801561018e57600080fd5b506100c060043561049a565b3480156101a657600080fd5b506101bb600160a060020a03600435166104c4565b60408051918252519081900360200190f35b3480156101d957600080fd5b50610100600160a060020a036004351660243515156104f7565b3480156101ff57600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261010094600160a060020a03813581169560248035909216956044359536956084940191819084018382808284375094975061057b9650505050505050565b34801561026e57600080fd5b5061016e600160a060020a03600435811690602435166105ba565b600090815260016020526040902054600160a060020a031690565b60006102af8261049a565b9050600160a060020a0383811690821614156102ca57600080fd5b33600160a060020a03821614806102e657506102e681336105ba565b15156102f157600080fd5b60006102fc83610289565b600160a060020a031614158061031a5750600160a060020a03831615155b1561039157600082815260016020908152604091829020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03878116918217909255835186815293519093918516927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35b505050565b806103a133826105e8565b15156103ac57600080fd5b600160a060020a03841615156103c157600080fd5b600160a060020a03831615156103d657600080fd5b6103e08483610647565b6103ea84836106f5565b6103f4838361078a565b82600160a060020a031684600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a350505050565b8061045033826105e8565b151561045b57600080fd5b610477848484602060405190810160405280600081525061057b565b50505050565b600090815260208190526040902054600160a060020a0316151590565b600081815260208190526040812054600160a060020a03168015156104be57600080fd5b92915050565b6000600160a060020a03821615156104db57600080fd5b50600160a060020a031660009081526002602052604090205490565b600160a060020a03821633141561050d57600080fd5b336000818152600360209081526040808320600160a060020a03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b8161058633826105e8565b151561059157600080fd5b61059c858585610396565b6105a885858585610819565b15156105b357600080fd5b5050505050565b600160a060020a03918216600090815260036020908152604080832093909416825291909152205460ff1690565b6000806105f48361049a565b905080600160a060020a031684600160a060020a0316148061062f575083600160a060020a031661062484610289565b600160a060020a0316145b8061063f575061063f81856105ba565b949350505050565b81600160a060020a031661065a8261049a565b600160a060020a03161461066d57600080fd5b600081815260016020526040902054600160a060020a0316156106f1576000818152600160209081526040808320805473ffffffffffffffffffffffffffffffffffffffff1916905580518481529051600160a060020a038616927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35b5050565b81600160a060020a03166107088261049a565b600160a060020a03161461071b57600080fd5b600160a060020a03821660009081526002602052604090205461074590600163ffffffff6109a216565b600160a060020a039092166000908152600260209081526040808320949094559181529081905220805473ffffffffffffffffffffffffffffffffffffffff19169055565b600081815260208190526040902054600160a060020a0316156107ac57600080fd5b600081815260208181526040808320805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038716908117909155835260029091529020546107f99060016109b4565b600160a060020a0390921660009081526002602052604090209190915550565b60008061082e85600160a060020a03166109c1565b151561083d5760019150610999565b84600160a060020a031663f0b9e5ba8786866040518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018084600160a060020a0316600160a060020a0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156108d55781810151838201526020016108bd565b50505050905090810190601f1680156109025780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b15801561092357600080fd5b505af1158015610937573d6000803e3d6000fd5b505050506040513d602081101561094d57600080fd5b50517fffffffff0000000000000000000000000000000000000000000000000000000081167ff0b9e5ba0000000000000000000000000000000000000000000000000000000014925090505b50949350505050565b6000828211156109ae57fe5b50900390565b818101828110156104be57fe5b6000903b11905600a165627a7a72305820bed8986f67e3b490484de6c0715c1ecdc2d50c1ee53d8cd50abb943164cd08610029`

// DeployERC721BasicToken deploys a new Ethereum contract, binding an instance of ERC721BasicToken to it.
func DeployERC721BasicToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721BasicToken, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721BasicTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721BasicTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721BasicToken{ERC721BasicTokenCaller: ERC721BasicTokenCaller{contract: contract}, ERC721BasicTokenTransactor: ERC721BasicTokenTransactor{contract: contract}, ERC721BasicTokenFilterer: ERC721BasicTokenFilterer{contract: contract}}, nil
}

// ERC721BasicToken is an auto generated Go binding around an Ethereum contract.
type ERC721BasicToken struct {
	ERC721BasicTokenCaller     // Read-only binding to the contract
	ERC721BasicTokenTransactor // Write-only binding to the contract
	ERC721BasicTokenFilterer   // Log filterer for contract events
}

// ERC721BasicTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721BasicTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721BasicTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721BasicTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721BasicTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721BasicTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721BasicTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721BasicTokenSession struct {
	Contract     *ERC721BasicToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721BasicTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721BasicTokenCallerSession struct {
	Contract *ERC721BasicTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC721BasicTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721BasicTokenTransactorSession struct {
	Contract     *ERC721BasicTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC721BasicTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721BasicTokenRaw struct {
	Contract *ERC721BasicToken // Generic contract binding to access the raw methods on
}

// ERC721BasicTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721BasicTokenCallerRaw struct {
	Contract *ERC721BasicTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721BasicTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721BasicTokenTransactorRaw struct {
	Contract *ERC721BasicTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721BasicToken creates a new instance of ERC721BasicToken, bound to a specific deployed contract.
func NewERC721BasicToken(address common.Address, backend bind.ContractBackend) (*ERC721BasicToken, error) {
	contract, err := bindERC721BasicToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicToken{ERC721BasicTokenCaller: ERC721BasicTokenCaller{contract: contract}, ERC721BasicTokenTransactor: ERC721BasicTokenTransactor{contract: contract}, ERC721BasicTokenFilterer: ERC721BasicTokenFilterer{contract: contract}}, nil
}

// NewERC721BasicTokenCaller creates a new read-only instance of ERC721BasicToken, bound to a specific deployed contract.
func NewERC721BasicTokenCaller(address common.Address, caller bind.ContractCaller) (*ERC721BasicTokenCaller, error) {
	contract, err := bindERC721BasicToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTokenCaller{contract: contract}, nil
}

// NewERC721BasicTokenTransactor creates a new write-only instance of ERC721BasicToken, bound to a specific deployed contract.
func NewERC721BasicTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721BasicTokenTransactor, error) {
	contract, err := bindERC721BasicToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTokenTransactor{contract: contract}, nil
}

// NewERC721BasicTokenFilterer creates a new log filterer instance of ERC721BasicToken, bound to a specific deployed contract.
func NewERC721BasicTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721BasicTokenFilterer, error) {
	contract, err := bindERC721BasicToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTokenFilterer{contract: contract}, nil
}

// bindERC721BasicToken binds a generic wrapper to an already deployed contract.
func bindERC721BasicToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721BasicTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721BasicToken *ERC721BasicTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721BasicToken.Contract.ERC721BasicTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721BasicToken *ERC721BasicTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.ERC721BasicTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721BasicToken *ERC721BasicTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.ERC721BasicTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721BasicToken *ERC721BasicTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721BasicToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721BasicToken *ERC721BasicTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721BasicToken *ERC721BasicTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ERC721BasicToken *ERC721BasicTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721BasicToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ERC721BasicToken *ERC721BasicTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721BasicToken.Contract.BalanceOf(&_ERC721BasicToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ERC721BasicToken *ERC721BasicTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721BasicToken.Contract.BalanceOf(&_ERC721BasicToken.CallOpts, _owner)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_ERC721BasicToken *ERC721BasicTokenCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721BasicToken.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_ERC721BasicToken *ERC721BasicTokenSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721BasicToken.Contract.Exists(&_ERC721BasicToken.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_ERC721BasicToken *ERC721BasicTokenCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721BasicToken.Contract.Exists(&_ERC721BasicToken.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_ERC721BasicToken *ERC721BasicTokenCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721BasicToken.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_ERC721BasicToken *ERC721BasicTokenSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721BasicToken.Contract.GetApproved(&_ERC721BasicToken.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_ERC721BasicToken *ERC721BasicTokenCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721BasicToken.Contract.GetApproved(&_ERC721BasicToken.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(approved bool)
func (_ERC721BasicToken *ERC721BasicTokenCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721BasicToken.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(approved bool)
func (_ERC721BasicToken *ERC721BasicTokenSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721BasicToken.Contract.IsApprovedForAll(&_ERC721BasicToken.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(approved bool)
func (_ERC721BasicToken *ERC721BasicTokenCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721BasicToken.Contract.IsApprovedForAll(&_ERC721BasicToken.CallOpts, _owner, _operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_ERC721BasicToken *ERC721BasicTokenCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721BasicToken.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_ERC721BasicToken *ERC721BasicTokenSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721BasicToken.Contract.OwnerOf(&_ERC721BasicToken.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_ERC721BasicToken *ERC721BasicTokenCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721BasicToken.Contract.OwnerOf(&_ERC721BasicToken.CallOpts, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721BasicToken.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721BasicToken *ERC721BasicTokenSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.Approve(&_ERC721BasicToken.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.Approve(&_ERC721BasicToken.TransactOpts, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721BasicToken.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721BasicToken *ERC721BasicTokenSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.SafeTransferFrom(&_ERC721BasicToken.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.SafeTransferFrom(&_ERC721BasicToken.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, _to common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721BasicToken.contract.Transact(opts, "setApprovalForAll", _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_ERC721BasicToken *ERC721BasicTokenSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.SetApprovalForAll(&_ERC721BasicToken.TransactOpts, _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactorSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.SetApprovalForAll(&_ERC721BasicToken.TransactOpts, _to, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721BasicToken.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721BasicToken *ERC721BasicTokenSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.TransferFrom(&_ERC721BasicToken.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.TransferFrom(&_ERC721BasicToken.TransactOpts, _from, _to, _tokenId)
}

// ERC721BasicTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721BasicToken contract.
type ERC721BasicTokenApprovalIterator struct {
	Event *ERC721BasicTokenApproval // Event containing the contract specifics and raw log

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
func (it *ERC721BasicTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721BasicTokenApproval)
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
		it.Event = new(ERC721BasicTokenApproval)
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
func (it *ERC721BasicTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721BasicTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721BasicTokenApproval represents a Approval event raised by the ERC721BasicToken contract.
type ERC721BasicTokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_ERC721BasicToken *ERC721BasicTokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address) (*ERC721BasicTokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _ERC721BasicToken.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTokenApprovalIterator{contract: _ERC721BasicToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_ERC721BasicToken *ERC721BasicTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721BasicTokenApproval, _owner []common.Address, _approved []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _ERC721BasicToken.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721BasicTokenApproval)
				if err := _ERC721BasicToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC721BasicTokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721BasicToken contract.
type ERC721BasicTokenApprovalForAllIterator struct {
	Event *ERC721BasicTokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721BasicTokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721BasicTokenApprovalForAll)
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
		it.Event = new(ERC721BasicTokenApprovalForAll)
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
func (it *ERC721BasicTokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721BasicTokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721BasicTokenApprovalForAll represents a ApprovalForAll event raised by the ERC721BasicToken contract.
type ERC721BasicTokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721BasicToken *ERC721BasicTokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ERC721BasicTokenApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721BasicToken.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTokenApprovalForAllIterator{contract: _ERC721BasicToken.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721BasicToken *ERC721BasicTokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721BasicTokenApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721BasicToken.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721BasicTokenApprovalForAll)
				if err := _ERC721BasicToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ERC721BasicTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721BasicToken contract.
type ERC721BasicTokenTransferIterator struct {
	Event *ERC721BasicTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721BasicTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721BasicTokenTransfer)
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
		it.Event = new(ERC721BasicTokenTransfer)
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
func (it *ERC721BasicTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721BasicTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721BasicTokenTransfer represents a Transfer event raised by the ERC721BasicToken contract.
type ERC721BasicTokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_ERC721BasicToken *ERC721BasicTokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ERC721BasicTokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC721BasicToken.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTokenTransferIterator{contract: _ERC721BasicToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_ERC721BasicToken *ERC721BasicTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721BasicTokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC721BasicToken.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721BasicTokenTransfer)
				if err := _ERC721BasicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC721EnumerableABI is the input ABI used to generate the binding from.
const ERC721EnumerableABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"_exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721EnumerableBin is the compiled bytecode used for deploying new contracts.
const ERC721EnumerableBin = `0x`

// DeployERC721Enumerable deploys a new Ethereum contract, binding an instance of ERC721Enumerable to it.
func DeployERC721Enumerable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721Enumerable, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721EnumerableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721EnumerableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Enumerable{ERC721EnumerableCaller: ERC721EnumerableCaller{contract: contract}, ERC721EnumerableTransactor: ERC721EnumerableTransactor{contract: contract}, ERC721EnumerableFilterer: ERC721EnumerableFilterer{contract: contract}}, nil
}

// ERC721Enumerable is an auto generated Go binding around an Ethereum contract.
type ERC721Enumerable struct {
	ERC721EnumerableCaller     // Read-only binding to the contract
	ERC721EnumerableTransactor // Write-only binding to the contract
	ERC721EnumerableFilterer   // Log filterer for contract events
}

// ERC721EnumerableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721EnumerableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721EnumerableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721EnumerableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721EnumerableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721EnumerableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721EnumerableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721EnumerableSession struct {
	Contract     *ERC721Enumerable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721EnumerableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721EnumerableCallerSession struct {
	Contract *ERC721EnumerableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC721EnumerableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721EnumerableTransactorSession struct {
	Contract     *ERC721EnumerableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC721EnumerableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721EnumerableRaw struct {
	Contract *ERC721Enumerable // Generic contract binding to access the raw methods on
}

// ERC721EnumerableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721EnumerableCallerRaw struct {
	Contract *ERC721EnumerableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721EnumerableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721EnumerableTransactorRaw struct {
	Contract *ERC721EnumerableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Enumerable creates a new instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721Enumerable(address common.Address, backend bind.ContractBackend) (*ERC721Enumerable, error) {
	contract, err := bindERC721Enumerable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Enumerable{ERC721EnumerableCaller: ERC721EnumerableCaller{contract: contract}, ERC721EnumerableTransactor: ERC721EnumerableTransactor{contract: contract}, ERC721EnumerableFilterer: ERC721EnumerableFilterer{contract: contract}}, nil
}

// NewERC721EnumerableCaller creates a new read-only instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721EnumerableCaller(address common.Address, caller bind.ContractCaller) (*ERC721EnumerableCaller, error) {
	contract, err := bindERC721Enumerable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableCaller{contract: contract}, nil
}

// NewERC721EnumerableTransactor creates a new write-only instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721EnumerableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721EnumerableTransactor, error) {
	contract, err := bindERC721Enumerable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableTransactor{contract: contract}, nil
}

// NewERC721EnumerableFilterer creates a new log filterer instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721EnumerableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721EnumerableFilterer, error) {
	contract, err := bindERC721Enumerable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableFilterer{contract: contract}, nil
}

// bindERC721Enumerable binds a generic wrapper to an already deployed contract.
func bindERC721Enumerable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721EnumerableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Enumerable *ERC721EnumerableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Enumerable.Contract.ERC721EnumerableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Enumerable *ERC721EnumerableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.ERC721EnumerableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Enumerable *ERC721EnumerableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.ERC721EnumerableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Enumerable *ERC721EnumerableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Enumerable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Enumerable *ERC721EnumerableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Enumerable *ERC721EnumerableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Enumerable.Contract.BalanceOf(&_ERC721Enumerable.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Enumerable.Contract.BalanceOf(&_ERC721Enumerable.CallOpts, _owner)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(_exists bool)
func (_ERC721Enumerable *ERC721EnumerableCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(_exists bool)
func (_ERC721Enumerable *ERC721EnumerableSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721Enumerable.Contract.Exists(&_ERC721Enumerable.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(_exists bool)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721Enumerable.Contract.Exists(&_ERC721Enumerable.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721Enumerable *ERC721EnumerableCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721Enumerable *ERC721EnumerableSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.GetApproved(&_ERC721Enumerable.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.GetApproved(&_ERC721Enumerable.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721Enumerable *ERC721EnumerableSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Enumerable.Contract.IsApprovedForAll(&_ERC721Enumerable.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Enumerable.Contract.IsApprovedForAll(&_ERC721Enumerable.CallOpts, _owner, _operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721Enumerable *ERC721EnumerableCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721Enumerable *ERC721EnumerableSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.OwnerOf(&_ERC721Enumerable.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.OwnerOf(&_ERC721Enumerable.CallOpts, _tokenId)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenByIndex(&_ERC721Enumerable.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenByIndex(&_ERC721Enumerable.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(_tokenId uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(_tokenId uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenOfOwnerByIndex(&_ERC721Enumerable.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(_tokenId uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenOfOwnerByIndex(&_ERC721Enumerable.CallOpts, _owner, _index)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) TotalSupply() (*big.Int, error) {
	return _ERC721Enumerable.Contract.TotalSupply(&_ERC721Enumerable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721Enumerable.Contract.TotalSupply(&_ERC721Enumerable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.Approve(&_ERC721Enumerable.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.Approve(&_ERC721Enumerable.TransactOpts, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SafeTransferFrom(&_ERC721Enumerable.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SafeTransferFrom(&_ERC721Enumerable.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SetApprovalForAll(&_ERC721Enumerable.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SetApprovalForAll(&_ERC721Enumerable.TransactOpts, _operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.TransferFrom(&_ERC721Enumerable.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.TransferFrom(&_ERC721Enumerable.TransactOpts, _from, _to, _tokenId)
}

// ERC721EnumerableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Enumerable contract.
type ERC721EnumerableApprovalIterator struct {
	Event *ERC721EnumerableApproval // Event containing the contract specifics and raw log

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
func (it *ERC721EnumerableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721EnumerableApproval)
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
		it.Event = new(ERC721EnumerableApproval)
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
func (it *ERC721EnumerableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721EnumerableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721EnumerableApproval represents a Approval event raised by the ERC721Enumerable contract.
type ERC721EnumerableApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_ERC721Enumerable *ERC721EnumerableFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address) (*ERC721EnumerableApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableApprovalIterator{contract: _ERC721Enumerable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_ERC721Enumerable *ERC721EnumerableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721EnumerableApproval, _owner []common.Address, _approved []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721EnumerableApproval)
				if err := _ERC721Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC721EnumerableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Enumerable contract.
type ERC721EnumerableApprovalForAllIterator struct {
	Event *ERC721EnumerableApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721EnumerableApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721EnumerableApprovalForAll)
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
		it.Event = new(ERC721EnumerableApprovalForAll)
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
func (it *ERC721EnumerableApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721EnumerableApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721EnumerableApprovalForAll represents a ApprovalForAll event raised by the ERC721Enumerable contract.
type ERC721EnumerableApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721Enumerable *ERC721EnumerableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ERC721EnumerableApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableApprovalForAllIterator{contract: _ERC721Enumerable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721Enumerable *ERC721EnumerableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721EnumerableApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721EnumerableApprovalForAll)
				if err := _ERC721Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ERC721EnumerableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Enumerable contract.
type ERC721EnumerableTransferIterator struct {
	Event *ERC721EnumerableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721EnumerableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721EnumerableTransfer)
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
		it.Event = new(ERC721EnumerableTransfer)
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
func (it *ERC721EnumerableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721EnumerableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721EnumerableTransfer represents a Transfer event raised by the ERC721Enumerable contract.
type ERC721EnumerableTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_ERC721Enumerable *ERC721EnumerableFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ERC721EnumerableTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableTransferIterator{contract: _ERC721Enumerable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_ERC721Enumerable *ERC721EnumerableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721EnumerableTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721EnumerableTransfer)
				if err := _ERC721Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC721MetadataABI is the input ABI used to generate the binding from.
const ERC721MetadataABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"_exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721MetadataBin is the compiled bytecode used for deploying new contracts.
const ERC721MetadataBin = `0x`

// DeployERC721Metadata deploys a new Ethereum contract, binding an instance of ERC721Metadata to it.
func DeployERC721Metadata(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721Metadata, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721MetadataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721MetadataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Metadata{ERC721MetadataCaller: ERC721MetadataCaller{contract: contract}, ERC721MetadataTransactor: ERC721MetadataTransactor{contract: contract}, ERC721MetadataFilterer: ERC721MetadataFilterer{contract: contract}}, nil
}

// ERC721Metadata is an auto generated Go binding around an Ethereum contract.
type ERC721Metadata struct {
	ERC721MetadataCaller     // Read-only binding to the contract
	ERC721MetadataTransactor // Write-only binding to the contract
	ERC721MetadataFilterer   // Log filterer for contract events
}

// ERC721MetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721MetadataSession struct {
	Contract     *ERC721Metadata   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721MetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721MetadataCallerSession struct {
	Contract *ERC721MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC721MetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721MetadataTransactorSession struct {
	Contract     *ERC721MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC721MetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721MetadataRaw struct {
	Contract *ERC721Metadata // Generic contract binding to access the raw methods on
}

// ERC721MetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721MetadataCallerRaw struct {
	Contract *ERC721MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721MetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721MetadataTransactorRaw struct {
	Contract *ERC721MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Metadata creates a new instance of ERC721Metadata, bound to a specific deployed contract.
func NewERC721Metadata(address common.Address, backend bind.ContractBackend) (*ERC721Metadata, error) {
	contract, err := bindERC721Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Metadata{ERC721MetadataCaller: ERC721MetadataCaller{contract: contract}, ERC721MetadataTransactor: ERC721MetadataTransactor{contract: contract}, ERC721MetadataFilterer: ERC721MetadataFilterer{contract: contract}}, nil
}

// NewERC721MetadataCaller creates a new read-only instance of ERC721Metadata, bound to a specific deployed contract.
func NewERC721MetadataCaller(address common.Address, caller bind.ContractCaller) (*ERC721MetadataCaller, error) {
	contract, err := bindERC721Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataCaller{contract: contract}, nil
}

// NewERC721MetadataTransactor creates a new write-only instance of ERC721Metadata, bound to a specific deployed contract.
func NewERC721MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721MetadataTransactor, error) {
	contract, err := bindERC721Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataTransactor{contract: contract}, nil
}

// NewERC721MetadataFilterer creates a new log filterer instance of ERC721Metadata, bound to a specific deployed contract.
func NewERC721MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721MetadataFilterer, error) {
	contract, err := bindERC721Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataFilterer{contract: contract}, nil
}

// bindERC721Metadata binds a generic wrapper to an already deployed contract.
func bindERC721Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721MetadataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Metadata *ERC721MetadataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Metadata.Contract.ERC721MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Metadata *ERC721MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.ERC721MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Metadata *ERC721MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.ERC721MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Metadata *ERC721MetadataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Metadata *ERC721MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Metadata *ERC721MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721Metadata *ERC721MetadataCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721Metadata *ERC721MetadataSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Metadata.Contract.BalanceOf(&_ERC721Metadata.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721Metadata *ERC721MetadataCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Metadata.Contract.BalanceOf(&_ERC721Metadata.CallOpts, _owner)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(_exists bool)
func (_ERC721Metadata *ERC721MetadataCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(_exists bool)
func (_ERC721Metadata *ERC721MetadataSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721Metadata.Contract.Exists(&_ERC721Metadata.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(_exists bool)
func (_ERC721Metadata *ERC721MetadataCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721Metadata.Contract.Exists(&_ERC721Metadata.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721Metadata *ERC721MetadataCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721Metadata *ERC721MetadataSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Metadata.Contract.GetApproved(&_ERC721Metadata.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721Metadata *ERC721MetadataCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Metadata.Contract.GetApproved(&_ERC721Metadata.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721Metadata *ERC721MetadataCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721Metadata *ERC721MetadataSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Metadata.Contract.IsApprovedForAll(&_ERC721Metadata.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721Metadata *ERC721MetadataCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Metadata.Contract.IsApprovedForAll(&_ERC721Metadata.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_ERC721Metadata *ERC721MetadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_ERC721Metadata *ERC721MetadataSession) Name() (string, error) {
	return _ERC721Metadata.Contract.Name(&_ERC721Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_ERC721Metadata *ERC721MetadataCallerSession) Name() (string, error) {
	return _ERC721Metadata.Contract.Name(&_ERC721Metadata.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721Metadata *ERC721MetadataCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721Metadata *ERC721MetadataSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Metadata.Contract.OwnerOf(&_ERC721Metadata.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721Metadata *ERC721MetadataCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Metadata.Contract.OwnerOf(&_ERC721Metadata.CallOpts, _tokenId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_ERC721Metadata *ERC721MetadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_ERC721Metadata *ERC721MetadataSession) Symbol() (string, error) {
	return _ERC721Metadata.Contract.Symbol(&_ERC721Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_ERC721Metadata *ERC721MetadataCallerSession) Symbol() (string, error) {
	return _ERC721Metadata.Contract.Symbol(&_ERC721Metadata.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_ERC721Metadata *ERC721MetadataCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_ERC721Metadata *ERC721MetadataSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC721Metadata.Contract.TokenURI(&_ERC721Metadata.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_ERC721Metadata *ERC721MetadataCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC721Metadata.Contract.TokenURI(&_ERC721Metadata.CallOpts, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721Metadata *ERC721MetadataTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721Metadata *ERC721MetadataSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.Approve(&_ERC721Metadata.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721Metadata *ERC721MetadataTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.Approve(&_ERC721Metadata.TransactOpts, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721Metadata *ERC721MetadataTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Metadata.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721Metadata *ERC721MetadataSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.SafeTransferFrom(&_ERC721Metadata.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721Metadata *ERC721MetadataTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.SafeTransferFrom(&_ERC721Metadata.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721Metadata *ERC721MetadataTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Metadata.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721Metadata *ERC721MetadataSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.SetApprovalForAll(&_ERC721Metadata.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721Metadata *ERC721MetadataTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.SetApprovalForAll(&_ERC721Metadata.TransactOpts, _operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721Metadata *ERC721MetadataTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721Metadata *ERC721MetadataSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.TransferFrom(&_ERC721Metadata.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721Metadata *ERC721MetadataTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.TransferFrom(&_ERC721Metadata.TransactOpts, _from, _to, _tokenId)
}

// ERC721MetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Metadata contract.
type ERC721MetadataApprovalIterator struct {
	Event *ERC721MetadataApproval // Event containing the contract specifics and raw log

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
func (it *ERC721MetadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MetadataApproval)
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
		it.Event = new(ERC721MetadataApproval)
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
func (it *ERC721MetadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MetadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MetadataApproval represents a Approval event raised by the ERC721Metadata contract.
type ERC721MetadataApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_ERC721Metadata *ERC721MetadataFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address) (*ERC721MetadataApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _ERC721Metadata.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataApprovalIterator{contract: _ERC721Metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_ERC721Metadata *ERC721MetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721MetadataApproval, _owner []common.Address, _approved []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _ERC721Metadata.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MetadataApproval)
				if err := _ERC721Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC721MetadataApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Metadata contract.
type ERC721MetadataApprovalForAllIterator struct {
	Event *ERC721MetadataApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721MetadataApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MetadataApprovalForAll)
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
		it.Event = new(ERC721MetadataApprovalForAll)
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
func (it *ERC721MetadataApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MetadataApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MetadataApprovalForAll represents a ApprovalForAll event raised by the ERC721Metadata contract.
type ERC721MetadataApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721Metadata *ERC721MetadataFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ERC721MetadataApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Metadata.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataApprovalForAllIterator{contract: _ERC721Metadata.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721Metadata *ERC721MetadataFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721MetadataApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Metadata.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MetadataApprovalForAll)
				if err := _ERC721Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ERC721MetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Metadata contract.
type ERC721MetadataTransferIterator struct {
	Event *ERC721MetadataTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721MetadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MetadataTransfer)
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
		it.Event = new(ERC721MetadataTransfer)
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
func (it *ERC721MetadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MetadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MetadataTransfer represents a Transfer event raised by the ERC721Metadata contract.
type ERC721MetadataTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_ERC721Metadata *ERC721MetadataFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ERC721MetadataTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC721Metadata.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataTransferIterator{contract: _ERC721Metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_ERC721Metadata *ERC721MetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721MetadataTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC721Metadata.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MetadataTransfer)
				if err := _ERC721Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC721ReceiverABI is the input ABI used to generate the binding from.
const ERC721ReceiverABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC721ReceiverBin is the compiled bytecode used for deploying new contracts.
const ERC721ReceiverBin = `0x`

// DeployERC721Receiver deploys a new Ethereum contract, binding an instance of ERC721Receiver to it.
func DeployERC721Receiver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721Receiver, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ReceiverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721ReceiverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Receiver{ERC721ReceiverCaller: ERC721ReceiverCaller{contract: contract}, ERC721ReceiverTransactor: ERC721ReceiverTransactor{contract: contract}, ERC721ReceiverFilterer: ERC721ReceiverFilterer{contract: contract}}, nil
}

// ERC721Receiver is an auto generated Go binding around an Ethereum contract.
type ERC721Receiver struct {
	ERC721ReceiverCaller     // Read-only binding to the contract
	ERC721ReceiverTransactor // Write-only binding to the contract
	ERC721ReceiverFilterer   // Log filterer for contract events
}

// ERC721ReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721ReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721ReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721ReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721ReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721ReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721ReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721ReceiverSession struct {
	Contract     *ERC721Receiver   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721ReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721ReceiverCallerSession struct {
	Contract *ERC721ReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC721ReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721ReceiverTransactorSession struct {
	Contract     *ERC721ReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC721ReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721ReceiverRaw struct {
	Contract *ERC721Receiver // Generic contract binding to access the raw methods on
}

// ERC721ReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721ReceiverCallerRaw struct {
	Contract *ERC721ReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721ReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721ReceiverTransactorRaw struct {
	Contract *ERC721ReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Receiver creates a new instance of ERC721Receiver, bound to a specific deployed contract.
func NewERC721Receiver(address common.Address, backend bind.ContractBackend) (*ERC721Receiver, error) {
	contract, err := bindERC721Receiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Receiver{ERC721ReceiverCaller: ERC721ReceiverCaller{contract: contract}, ERC721ReceiverTransactor: ERC721ReceiverTransactor{contract: contract}, ERC721ReceiverFilterer: ERC721ReceiverFilterer{contract: contract}}, nil
}

// NewERC721ReceiverCaller creates a new read-only instance of ERC721Receiver, bound to a specific deployed contract.
func NewERC721ReceiverCaller(address common.Address, caller bind.ContractCaller) (*ERC721ReceiverCaller, error) {
	contract, err := bindERC721Receiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721ReceiverCaller{contract: contract}, nil
}

// NewERC721ReceiverTransactor creates a new write-only instance of ERC721Receiver, bound to a specific deployed contract.
func NewERC721ReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721ReceiverTransactor, error) {
	contract, err := bindERC721Receiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721ReceiverTransactor{contract: contract}, nil
}

// NewERC721ReceiverFilterer creates a new log filterer instance of ERC721Receiver, bound to a specific deployed contract.
func NewERC721ReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721ReceiverFilterer, error) {
	contract, err := bindERC721Receiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721ReceiverFilterer{contract: contract}, nil
}

// bindERC721Receiver binds a generic wrapper to an already deployed contract.
func bindERC721Receiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Receiver *ERC721ReceiverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Receiver.Contract.ERC721ReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Receiver *ERC721ReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Receiver.Contract.ERC721ReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Receiver *ERC721ReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Receiver.Contract.ERC721ReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Receiver *ERC721ReceiverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Receiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Receiver *ERC721ReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Receiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Receiver *ERC721ReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Receiver.Contract.contract.Transact(opts, method, params...)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0xf0b9e5ba.
//
// Solidity: function onERC721Received(_from address, _tokenId uint256, _data bytes) returns(bytes4)
func (_ERC721Receiver *ERC721ReceiverTransactor) OnERC721Received(opts *bind.TransactOpts, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Receiver.contract.Transact(opts, "onERC721Received", _from, _tokenId, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0xf0b9e5ba.
//
// Solidity: function onERC721Received(_from address, _tokenId uint256, _data bytes) returns(bytes4)
func (_ERC721Receiver *ERC721ReceiverSession) OnERC721Received(_from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Receiver.Contract.OnERC721Received(&_ERC721Receiver.TransactOpts, _from, _tokenId, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0xf0b9e5ba.
//
// Solidity: function onERC721Received(_from address, _tokenId uint256, _data bytes) returns(bytes4)
func (_ERC721Receiver *ERC721ReceiverTransactorSession) OnERC721Received(_from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Receiver.Contract.OnERC721Received(&_ERC721Receiver.TransactOpts, _from, _tokenId, _data)
}

// ERC721TokenImplementABI is the input ABI used to generate the binding from.
const ERC721TokenImplementABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"approved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721TokenImplementBin is the compiled bytecode used for deploying new contracts.
const ERC721TokenImplementBin = `0x60c0604052600a60808190527f5374616d70546f6b656e0000000000000000000000000000000000000000000060a09081526200004091600791906200009c565b506040805180820190915260028082527f5354000000000000000000000000000000000000000000000000000000000000602090920191825262000087916008916200009c565b503480156200009557600080fd5b5062000141565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000df57805160ff19168380011785556200010f565b828001600101855582156200010f579182015b828111156200010f578251825591602001919060010190620000f2565b506200011d92915062000121565b5090565b6200013e91905b808211156200011d576000815560010162000128565b90565b610f3680620001516000396000f3006080604052600436106100e55763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100ea578063081812fc14610174578063095ea7b3146101a857806318160ddd146101ce57806323b872dd146101f55780632f745c591461021f57806342842e0e146102435780634f558e791461026d5780634f6ccce7146102995780636352211e146102b157806370a08231146102c957806395d89b41146102ea578063a22cb465146102ff578063b88d4fde14610325578063c87b56dd14610394578063e985e9c5146103ac575b600080fd5b3480156100f657600080fd5b506100ff6103d3565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610139578181015183820152602001610121565b50505050905090810190601f1680156101665780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561018057600080fd5b5061018c60043561046a565b60408051600160a060020a039092168252519081900360200190f35b3480156101b457600080fd5b506101cc600160a060020a0360043516602435610485565b005b3480156101da57600080fd5b506101e3610577565b60408051918252519081900360200190f35b34801561020157600080fd5b506101cc600160a060020a036004358116906024351660443561057d565b34801561022b57600080fd5b506101e3600160a060020a036004351660243561062c565b34801561024f57600080fd5b506101cc600160a060020a0360043581169060243516604435610679565b34801561027957600080fd5b506102856004356106b1565b604080519115158252519081900360200190f35b3480156102a557600080fd5b506101e36004356106ce565b3480156102bd57600080fd5b5061018c600435610703565b3480156102d557600080fd5b506101e3600160a060020a036004351661072d565b3480156102f657600080fd5b506100ff610760565b34801561030b57600080fd5b506101cc600160a060020a036004351660243515156107c1565b34801561033157600080fd5b50604080516020601f6064356004818101359283018490048402850184019095528184526101cc94600160a060020a0381358116956024803590921695604435953695608494019181908401838280828437509497506108459650505050505050565b3480156103a057600080fd5b506100ff600435610884565b3480156103b857600080fd5b50610285600160a060020a0360043581169060243516610939565b60078054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561045f5780601f106104345761010080835404028352916020019161045f565b820191906000526020600020905b81548152906001019060200180831161044257829003601f168201915b505050505090505b90565b600090815260046020526040902054600160a060020a031690565b600061049082610703565b9050600160a060020a0383811690821614156104ab57600080fd5b33600160a060020a03821614806104c757506104c78133610939565b15156104d257600080fd5b60006104dd8361046a565b600160a060020a03161415806104fb5750600160a060020a03831615155b1561057257600082815260046020908152604091829020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03878116918217909255835186815293519093918516927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35b505050565b600b5490565b806105883382610967565b151561059357600080fd5b600160a060020a03841615156105a857600080fd5b600160a060020a03831615156105bd57600080fd5b6105c784836109c6565b6105d18483610a74565b6105db8383610bad565b82600160a060020a031684600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a350505050565b60006106378361072d565b821061064257600080fd5b600160a060020a038316600090815260096020526040902080548390811061066657fe5b9060005260206000200154905092915050565b806106843382610967565b151561068f57600080fd5b6106ab8484846020604051908101604052806000815250610845565b50505050565b600090815260036020526040902054600160a060020a0316151590565b60006106d8610577565b82106106e357600080fd5b600b8054839081106106f157fe5b90600052602060002001549050919050565b600081815260036020526040812054600160a060020a031680151561072757600080fd5b92915050565b6000600160a060020a038216151561074457600080fd5b50600160a060020a031660009081526005602052604090205490565b60088054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561045f5780601f106104345761010080835404028352916020019161045f565b600160a060020a0382163314156107d757600080fd5b336000818152600660209081526040808320600160a060020a03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b816108503382610967565b151561085b57600080fd5b61086685858561057d565b61087285858585610bf6565b151561087d57600080fd5b5050505050565b606061088f826106b1565b151561089a57600080fd5b6000828152600d602090815260409182902080548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452909183018282801561092d5780601f106109025761010080835404028352916020019161092d565b820191906000526020600020905b81548152906001019060200180831161091057829003601f168201915b50505050509050919050565b600160a060020a03918216600090815260066020908152604080832093909416825291909152205460ff1690565b60008061097383610703565b905080600160a060020a031684600160a060020a031614806109ae575083600160a060020a03166109a38461046a565b600160a060020a0316145b806109be57506109be8185610939565b949350505050565b81600160a060020a03166109d982610703565b600160a060020a0316146109ec57600080fd5b600081815260046020526040902054600160a060020a031615610a70576000818152600460209081526040808320805473ffffffffffffffffffffffffffffffffffffffff1916905580518481529051600160a060020a038616927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35b5050565b6000806000610a838585610d7f565b6000848152600a6020908152604080832054600160a060020a0389168452600990925290912054909350610abe90600163ffffffff610e1516565b600160a060020a038616600090815260096020526040902080549193509083908110610ae657fe5b90600052602060002001549050806009600087600160a060020a0316600160a060020a0316815260200190815260200160002084815481101515610b2657fe5b6000918252602080832090910192909255600160a060020a0387168152600990915260408120805484908110610b5857fe5b6000918252602080832090910192909255600160a060020a0387168152600990915260409020805490610b8f906000198301610ecd565b506000938452600a6020526040808520859055908452909220555050565b6000610bb98383610e27565b50600160a060020a03909116600090815260096020908152604080832080546001810182559084528284208101859055938352600a909152902055565b600080610c0b85600160a060020a0316610eb8565b1515610c1a5760019150610d76565b84600160a060020a031663f0b9e5ba8786866040518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018084600160a060020a0316600160a060020a0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610cb2578181015183820152602001610c9a565b50505050905090810190601f168015610cdf5780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b158015610d0057600080fd5b505af1158015610d14573d6000803e3d6000fd5b505050506040513d6020811015610d2a57600080fd5b50517fffffffff0000000000000000000000000000000000000000000000000000000081167ff0b9e5ba0000000000000000000000000000000000000000000000000000000014925090505b50949350505050565b81600160a060020a0316610d9282610703565b600160a060020a031614610da557600080fd5b600160a060020a038216600090815260056020526040902054610dcf90600163ffffffff610e1516565b600160a060020a03909216600090815260056020908152604080832094909455918152600390915220805473ffffffffffffffffffffffffffffffffffffffff19169055565b600082821115610e2157fe5b50900390565b600081815260036020526040902054600160a060020a031615610e4957600080fd5b6000818152600360209081526040808320805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03871690811790915583526005909152902054610e98906001610ec0565b600160a060020a0390921660009081526005602052604090209190915550565b6000903b1190565b8181018281101561072757fe5b8154818355818111156105725760008381526020902061057291810190830161046791905b80821115610f065760008155600101610ef2565b50905600a165627a7a723058207f247d2ec03a7fe1d7b4f76f0498c23f00c7fbfe4af953e3f236176bdd0d785f0029`

// DeployERC721TokenImplement deploys a new Ethereum contract, binding an instance of ERC721TokenImplement to it.
func DeployERC721TokenImplement(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721TokenImplement, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721TokenImplementABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721TokenImplementBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721TokenImplement{ERC721TokenImplementCaller: ERC721TokenImplementCaller{contract: contract}, ERC721TokenImplementTransactor: ERC721TokenImplementTransactor{contract: contract}, ERC721TokenImplementFilterer: ERC721TokenImplementFilterer{contract: contract}}, nil
}

// ERC721TokenImplement is an auto generated Go binding around an Ethereum contract.
type ERC721TokenImplement struct {
	ERC721TokenImplementCaller     // Read-only binding to the contract
	ERC721TokenImplementTransactor // Write-only binding to the contract
	ERC721TokenImplementFilterer   // Log filterer for contract events
}

// ERC721TokenImplementCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721TokenImplementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721TokenImplementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721TokenImplementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721TokenImplementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721TokenImplementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721TokenImplementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721TokenImplementSession struct {
	Contract     *ERC721TokenImplement // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC721TokenImplementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721TokenImplementCallerSession struct {
	Contract *ERC721TokenImplementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ERC721TokenImplementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721TokenImplementTransactorSession struct {
	Contract     *ERC721TokenImplementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ERC721TokenImplementRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721TokenImplementRaw struct {
	Contract *ERC721TokenImplement // Generic contract binding to access the raw methods on
}

// ERC721TokenImplementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721TokenImplementCallerRaw struct {
	Contract *ERC721TokenImplementCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721TokenImplementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721TokenImplementTransactorRaw struct {
	Contract *ERC721TokenImplementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721TokenImplement creates a new instance of ERC721TokenImplement, bound to a specific deployed contract.
func NewERC721TokenImplement(address common.Address, backend bind.ContractBackend) (*ERC721TokenImplement, error) {
	contract, err := bindERC721TokenImplement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenImplement{ERC721TokenImplementCaller: ERC721TokenImplementCaller{contract: contract}, ERC721TokenImplementTransactor: ERC721TokenImplementTransactor{contract: contract}, ERC721TokenImplementFilterer: ERC721TokenImplementFilterer{contract: contract}}, nil
}

// NewERC721TokenImplementCaller creates a new read-only instance of ERC721TokenImplement, bound to a specific deployed contract.
func NewERC721TokenImplementCaller(address common.Address, caller bind.ContractCaller) (*ERC721TokenImplementCaller, error) {
	contract, err := bindERC721TokenImplement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenImplementCaller{contract: contract}, nil
}

// NewERC721TokenImplementTransactor creates a new write-only instance of ERC721TokenImplement, bound to a specific deployed contract.
func NewERC721TokenImplementTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721TokenImplementTransactor, error) {
	contract, err := bindERC721TokenImplement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenImplementTransactor{contract: contract}, nil
}

// NewERC721TokenImplementFilterer creates a new log filterer instance of ERC721TokenImplement, bound to a specific deployed contract.
func NewERC721TokenImplementFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721TokenImplementFilterer, error) {
	contract, err := bindERC721TokenImplement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenImplementFilterer{contract: contract}, nil
}

// bindERC721TokenImplement binds a generic wrapper to an already deployed contract.
func bindERC721TokenImplement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721TokenImplementABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721TokenImplement *ERC721TokenImplementRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721TokenImplement.Contract.ERC721TokenImplementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721TokenImplement *ERC721TokenImplementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721TokenImplement.Contract.ERC721TokenImplementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721TokenImplement *ERC721TokenImplementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721TokenImplement.Contract.ERC721TokenImplementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721TokenImplement *ERC721TokenImplementCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721TokenImplement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721TokenImplement *ERC721TokenImplementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721TokenImplement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721TokenImplement *ERC721TokenImplementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721TokenImplement.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ERC721TokenImplement *ERC721TokenImplementCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721TokenImplement.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ERC721TokenImplement *ERC721TokenImplementSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721TokenImplement.Contract.BalanceOf(&_ERC721TokenImplement.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ERC721TokenImplement *ERC721TokenImplementCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721TokenImplement.Contract.BalanceOf(&_ERC721TokenImplement.CallOpts, _owner)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_ERC721TokenImplement *ERC721TokenImplementCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721TokenImplement.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_ERC721TokenImplement *ERC721TokenImplementSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721TokenImplement.Contract.Exists(&_ERC721TokenImplement.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_ERC721TokenImplement *ERC721TokenImplementCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721TokenImplement.Contract.Exists(&_ERC721TokenImplement.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_ERC721TokenImplement *ERC721TokenImplementCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721TokenImplement.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_ERC721TokenImplement *ERC721TokenImplementSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721TokenImplement.Contract.GetApproved(&_ERC721TokenImplement.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_ERC721TokenImplement *ERC721TokenImplementCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721TokenImplement.Contract.GetApproved(&_ERC721TokenImplement.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(approved bool)
func (_ERC721TokenImplement *ERC721TokenImplementCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721TokenImplement.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(approved bool)
func (_ERC721TokenImplement *ERC721TokenImplementSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721TokenImplement.Contract.IsApprovedForAll(&_ERC721TokenImplement.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(approved bool)
func (_ERC721TokenImplement *ERC721TokenImplementCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721TokenImplement.Contract.IsApprovedForAll(&_ERC721TokenImplement.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC721TokenImplement *ERC721TokenImplementCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721TokenImplement.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC721TokenImplement *ERC721TokenImplementSession) Name() (string, error) {
	return _ERC721TokenImplement.Contract.Name(&_ERC721TokenImplement.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC721TokenImplement *ERC721TokenImplementCallerSession) Name() (string, error) {
	return _ERC721TokenImplement.Contract.Name(&_ERC721TokenImplement.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_ERC721TokenImplement *ERC721TokenImplementCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721TokenImplement.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_ERC721TokenImplement *ERC721TokenImplementSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721TokenImplement.Contract.OwnerOf(&_ERC721TokenImplement.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_ERC721TokenImplement *ERC721TokenImplementCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721TokenImplement.Contract.OwnerOf(&_ERC721TokenImplement.CallOpts, _tokenId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC721TokenImplement *ERC721TokenImplementCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721TokenImplement.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC721TokenImplement *ERC721TokenImplementSession) Symbol() (string, error) {
	return _ERC721TokenImplement.Contract.Symbol(&_ERC721TokenImplement.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC721TokenImplement *ERC721TokenImplementCallerSession) Symbol() (string, error) {
	return _ERC721TokenImplement.Contract.Symbol(&_ERC721TokenImplement.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC721TokenImplement *ERC721TokenImplementCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721TokenImplement.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC721TokenImplement *ERC721TokenImplementSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC721TokenImplement.Contract.TokenByIndex(&_ERC721TokenImplement.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC721TokenImplement *ERC721TokenImplementCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC721TokenImplement.Contract.TokenByIndex(&_ERC721TokenImplement.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_ERC721TokenImplement *ERC721TokenImplementCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721TokenImplement.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_ERC721TokenImplement *ERC721TokenImplementSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721TokenImplement.Contract.TokenOfOwnerByIndex(&_ERC721TokenImplement.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_ERC721TokenImplement *ERC721TokenImplementCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721TokenImplement.Contract.TokenOfOwnerByIndex(&_ERC721TokenImplement.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_ERC721TokenImplement *ERC721TokenImplementCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721TokenImplement.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_ERC721TokenImplement *ERC721TokenImplementSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC721TokenImplement.Contract.TokenURI(&_ERC721TokenImplement.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_ERC721TokenImplement *ERC721TokenImplementCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC721TokenImplement.Contract.TokenURI(&_ERC721TokenImplement.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721TokenImplement *ERC721TokenImplementCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721TokenImplement.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721TokenImplement *ERC721TokenImplementSession) TotalSupply() (*big.Int, error) {
	return _ERC721TokenImplement.Contract.TotalSupply(&_ERC721TokenImplement.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721TokenImplement *ERC721TokenImplementCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721TokenImplement.Contract.TotalSupply(&_ERC721TokenImplement.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721TokenImplement *ERC721TokenImplementTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721TokenImplement.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721TokenImplement *ERC721TokenImplementSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721TokenImplement.Contract.Approve(&_ERC721TokenImplement.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721TokenImplement *ERC721TokenImplementTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721TokenImplement.Contract.Approve(&_ERC721TokenImplement.TransactOpts, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721TokenImplement *ERC721TokenImplementTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721TokenImplement.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721TokenImplement *ERC721TokenImplementSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721TokenImplement.Contract.SafeTransferFrom(&_ERC721TokenImplement.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721TokenImplement *ERC721TokenImplementTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721TokenImplement.Contract.SafeTransferFrom(&_ERC721TokenImplement.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_ERC721TokenImplement *ERC721TokenImplementTransactor) SetApprovalForAll(opts *bind.TransactOpts, _to common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721TokenImplement.contract.Transact(opts, "setApprovalForAll", _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_ERC721TokenImplement *ERC721TokenImplementSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721TokenImplement.Contract.SetApprovalForAll(&_ERC721TokenImplement.TransactOpts, _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_ERC721TokenImplement *ERC721TokenImplementTransactorSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721TokenImplement.Contract.SetApprovalForAll(&_ERC721TokenImplement.TransactOpts, _to, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721TokenImplement *ERC721TokenImplementTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721TokenImplement.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721TokenImplement *ERC721TokenImplementSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721TokenImplement.Contract.TransferFrom(&_ERC721TokenImplement.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721TokenImplement *ERC721TokenImplementTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721TokenImplement.Contract.TransferFrom(&_ERC721TokenImplement.TransactOpts, _from, _to, _tokenId)
}

// ERC721TokenImplementApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721TokenImplement contract.
type ERC721TokenImplementApprovalIterator struct {
	Event *ERC721TokenImplementApproval // Event containing the contract specifics and raw log

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
func (it *ERC721TokenImplementApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721TokenImplementApproval)
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
		it.Event = new(ERC721TokenImplementApproval)
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
func (it *ERC721TokenImplementApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TokenImplementApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721TokenImplementApproval represents a Approval event raised by the ERC721TokenImplement contract.
type ERC721TokenImplementApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_ERC721TokenImplement *ERC721TokenImplementFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address) (*ERC721TokenImplementApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _ERC721TokenImplement.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenImplementApprovalIterator{contract: _ERC721TokenImplement.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_ERC721TokenImplement *ERC721TokenImplementFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721TokenImplementApproval, _owner []common.Address, _approved []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _ERC721TokenImplement.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721TokenImplementApproval)
				if err := _ERC721TokenImplement.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC721TokenImplementApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721TokenImplement contract.
type ERC721TokenImplementApprovalForAllIterator struct {
	Event *ERC721TokenImplementApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721TokenImplementApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721TokenImplementApprovalForAll)
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
		it.Event = new(ERC721TokenImplementApprovalForAll)
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
func (it *ERC721TokenImplementApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TokenImplementApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721TokenImplementApprovalForAll represents a ApprovalForAll event raised by the ERC721TokenImplement contract.
type ERC721TokenImplementApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721TokenImplement *ERC721TokenImplementFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ERC721TokenImplementApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721TokenImplement.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenImplementApprovalForAllIterator{contract: _ERC721TokenImplement.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721TokenImplement *ERC721TokenImplementFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721TokenImplementApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721TokenImplement.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721TokenImplementApprovalForAll)
				if err := _ERC721TokenImplement.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ERC721TokenImplementTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721TokenImplement contract.
type ERC721TokenImplementTransferIterator struct {
	Event *ERC721TokenImplementTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721TokenImplementTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721TokenImplementTransfer)
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
		it.Event = new(ERC721TokenImplementTransfer)
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
func (it *ERC721TokenImplementTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TokenImplementTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721TokenImplementTransfer represents a Transfer event raised by the ERC721TokenImplement contract.
type ERC721TokenImplementTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_ERC721TokenImplement *ERC721TokenImplementFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ERC721TokenImplementTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC721TokenImplement.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenImplementTransferIterator{contract: _ERC721TokenImplement.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_ERC721TokenImplement *ERC721TokenImplementFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721TokenImplementTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC721TokenImplement.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721TokenImplementTransfer)
				if err := _ERC721TokenImplement.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a0319163317905561017f806100326000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610050578063f2fde38b1461008e575b600080fd5b34801561005c57600080fd5b506100656100be565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561009a57600080fd5b506100bc73ffffffffffffffffffffffffffffffffffffffff600435166100da565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60005473ffffffffffffffffffffffffffffffffffffffff1633146100fe57600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615610150576000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff83161790555b505600a165627a7a72305820eca7e899644067fbf010495ce7e9b8f3104c2f0a0729d47603a0f64c6c60a4b70029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// PausableBin is the compiled bytecode used for deploying new contracts.
const PausableBin = `0x608060405260008054600160a860020a031916331790556102f4806100256000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633f4ba83a81146100715780635c975abb1461009a5780638456cb59146100af5780638da5cb5b146100c4578063f2fde38b146100f5575b600080fd5b34801561007d57600080fd5b50610086610118565b604080519115158252519081900360200190f35b3480156100a657600080fd5b506100866101a4565b3480156100bb57600080fd5b506100866101c5565b3480156100d057600080fd5b506100d9610267565b60408051600160a060020a039092168252519081900360200190f35b34801561010157600080fd5b50610116600160a060020a0360043516610276565b005b60008054600160a060020a0316331461013057600080fd5b60005474010000000000000000000000000000000000000000900460ff16151561015957600080fd5b6000805474ff0000000000000000000000000000000000000000191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a150600190565b60005474010000000000000000000000000000000000000000900460ff1681565b60008054600160a060020a031633146101dd57600080fd5b60005474010000000000000000000000000000000000000000900460ff161561020557600080fd5b6000805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a150600190565b600054600160a060020a031681565b600054600160a060020a0316331461028d57600080fd5b600160a060020a038116156102c5576000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b505600a165627a7a723058204f3cb55a9c0c03637f65e170e91dfc2ec37b7b51ff65090b557549abc16a4ebd0029`

// DeployPausable deploys a new Ethereum contract, binding an instance of Pausable to it.
func DeployPausable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pausable, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PausableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pausable *PausableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pausable *PausableSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pausable *PausableCallerSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_Pausable *PausableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_Pausable *PausableSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_Pausable *PausableTransactorSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Pausable *PausableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Pausable *PausableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Pausable *PausableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_Pausable *PausableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_Pausable *PausableSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_Pausable *PausableTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// PausablePauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Pausable contract.
type PausablePauseIterator struct {
	Event *PausablePause // Event containing the contract specifics and raw log

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
func (it *PausablePauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePause)
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
		it.Event = new(PausablePause)
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
func (it *PausablePauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePause represents a Pause event raised by the Pausable contract.
type PausablePause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Pausable *PausableFilterer) FilterPause(opts *bind.FilterOpts) (*PausablePauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PausablePauseIterator{contract: _Pausable.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Pausable *PausableFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PausablePause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePause)
				if err := _Pausable.contract.UnpackLog(event, "Pause", log); err != nil {
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

// PausableUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Pausable contract.
type PausableUnpauseIterator struct {
	Event *PausableUnpause // Event containing the contract specifics and raw log

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
func (it *PausableUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpause)
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
		it.Event = new(PausableUnpause)
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
func (it *PausableUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpause represents a Unpause event raised by the Pausable contract.
type PausableUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Pausable *PausableFilterer) FilterUnpause(opts *bind.FilterOpts) (*PausableUnpauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PausableUnpauseIterator{contract: _Pausable.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Pausable *PausableFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PausableUnpause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpause)
				if err := _Pausable.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// RepoTransactionABI is the input ABI used to generate the binding from.
const RepoTransactionABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_repoCount\",\"type\":\"uint256\"}],\"name\":\"repoIngots\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"approved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"repoIngotsInfo\",\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"repoCount\",\"type\":\"uint256\"},{\"name\":\"seller\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"repoCount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"RepoIngotsSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"stampId\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"totalAmount\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"remainingAmount\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"year\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"setId\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"memberOfSetId\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"appearance\",\"type\":\"uint8\"}],\"name\":\"CreateNewStamp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// RepoTransactionBin is the compiled bytecode used for deploying new contracts.
const RepoTransactionBin = `0x6002805460a060020a60ff021916905560c0604052600a60808190527f5374616d70546f6b656e0000000000000000000000000000000000000000000060a09081526200004e919081620000aa565b506040805180820190915260028082527f535400000000000000000000000000000000000000000000000000000000000060209092019182526200009591600b91620000aa565b50348015620000a357600080fd5b506200014f565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000ed57805160ff19168380011785556200011d565b828001600101855582156200011d579182015b828111156200011d57825182559160200191906001019062000100565b506200012b9291506200012f565b5090565b6200014c91905b808211156200012b576000815560010162000136565b90565b611450806200015f6000396000f30060806040526004361061015e5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630519ce79811461016357806306fdde0314610194578063081812fc1461021e578063095ea7b3146102365780630a0f81681461025c57806318160ddd1461027157806323b872dd1461029857806327d7874c146102c25780632ba73c15146102e35780632f745c59146103045780633f4ba83a1461032857806342842e0e1461033d5780634e0a3379146103675780634f558e79146103885780634f6ccce7146103b45780635c975abb146103cc5780636352211e146103e157806370a08231146103f95780638456cb591461041a578063949d354e1461042f57806395d89b411461044a578063a22cb4651461045f578063b047fb5014610485578063b88d4fde1461049a578063c87b56dd14610509578063e985e9c514610521578063faf20cc514610548575b600080fd5b34801561016f57600080fd5b50610178610587565b60408051600160a060020a039092168252519081900360200190f35b3480156101a057600080fd5b506101a9610596565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101e35781810151838201526020016101cb565b50505050905090810190601f1680156102105780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561022a57600080fd5b5061017860043561062d565b34801561024257600080fd5b5061025a600160a060020a0360043516602435610648565b005b34801561026857600080fd5b5061017861072d565b34801561027d57600080fd5b5061028661073c565b60408051918252519081900360200190f35b3480156102a457600080fd5b5061025a600160a060020a0360043581169060243516604435610742565b3480156102ce57600080fd5b5061025a600160a060020a03600435166107f1565b3480156102ef57600080fd5b5061025a600160a060020a036004351661083f565b34801561031057600080fd5b50610286600160a060020a036004351660243561088d565b34801561033457600080fd5b5061025a6108da565b34801561034957600080fd5b5061025a600160a060020a036004358116906024351660443561093a565b34801561037357600080fd5b5061025a600160a060020a0360043516610972565b34801561039457600080fd5b506103a06004356109c0565b604080519115158252519081900360200190f35b3480156103c057600080fd5b506102866004356109dd565b3480156103d857600080fd5b506103a0610a12565b3480156103ed57600080fd5b50610178600435610a33565b34801561040557600080fd5b50610286600160a060020a0360043516610a5d565b34801561042657600080fd5b5061025a610a90565b34801561043b57600080fd5b5061025a600435602435610b32565b34801561045657600080fd5b506101a9610bf4565b34801561046b57600080fd5b5061025a600160a060020a03600435166024351515610c55565b34801561049157600080fd5b50610178610cd9565b3480156104a657600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261025a94600160a060020a038135811695602480359092169560443595369560849401918190840183828082843750949750610ce89650505050505050565b34801561051557600080fd5b506101a9600435610d27565b34801561052d57600080fd5b506103a0600160a060020a0360043581169060243516610ddc565b34801561055457600080fd5b50610560600435610e0a565b604080519384526020840192909252600160a060020a031682820152519081900360600190f35b600154600160a060020a031681565b600a8054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156106225780601f106105f757610100808354040283529160200191610622565b820191906000526020600020905b81548152906001019060200180831161060557829003601f168201915b505050505090505b90565b600090815260076020526040902054600160a060020a031690565b600061065382610a33565b9050600160a060020a03838116908216141561066e57600080fd5b33600160a060020a038216148061068a575061068a8133610ddc565b151561069557600080fd5b60006106a08361062d565b600160a060020a03161415806106be5750600160a060020a03831615155b15610728576000828152600760209081526040918290208054600160a060020a031916600160a060020a03878116918217909255835186815293519093918516927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35b505050565b600054600160a060020a031681565b600e5490565b8061074d3382610e5d565b151561075857600080fd5b600160a060020a038416151561076d57600080fd5b600160a060020a038316151561078257600080fd5b61078c8483610ebc565b6107968483610f5d565b6107a08383611096565b82600160a060020a031684600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a350505050565b600054600160a060020a0316331461080857600080fd5b600160a060020a038116151561081d57600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a0316331461085657600080fd5b600160a060020a038116151561086b57600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b600061089883610a5d565b82106108a357600080fd5b600160a060020a0383166000908152600c602052604090208054839081106108c757fe5b9060005260206000200154905092915050565b600054600160a060020a031633146108f157600080fd5b60025474010000000000000000000000000000000000000000900460ff16151561091a57600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b806109453382610e5d565b151561095057600080fd5b61096c8484846020604051908101604052806000815250610ce8565b50505050565b600054600160a060020a0316331461098957600080fd5b600160a060020a038116151561099e57600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b600090815260066020526040902054600160a060020a0316151590565b60006109e761073c565b82106109f257600080fd5b600e805483908110610a0057fe5b90600052602060002001549050919050565b60025474010000000000000000000000000000000000000000900460ff1681565b600081815260066020526040812054600160a060020a0316801515610a5757600080fd5b92915050565b6000600160a060020a0382161515610a7457600080fd5b50600160a060020a031660009081526008602052604090205490565b600254600160a060020a0316331480610ab35750600054600160a060020a031633145b80610ac85750600154600160a060020a031633145b1515610ad357600080fd5b60025474010000000000000000000000000000000000000000900460ff1615610afb57600080fd5b6002805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b610b3a6113bc565b610b4433846110df565b1515610b4f57600080fd5b610b5a333085610742565b50604080516060808201835233808352602080840187815284860187815260008981526011845287902086518154600160a060020a031916600160a060020a039091161781559151600183015551600290910155845187815290810186905280850191909152925191927f50cd59a56cf7ed9d5b7a9978578d10cc8529c67bead200fdb943ce7e7c9089da929081900390910190a1505050565b600b8054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156106225780601f106105f757610100808354040283529160200191610622565b600160a060020a038216331415610c6b57600080fd5b336000818152600960209081526040808320600160a060020a03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b600254600160a060020a031681565b81610cf33382610e5d565b1515610cfe57600080fd5b610d09858585610742565b610d15858585856110ff565b1515610d2057600080fd5b5050505050565b6060610d32826109c0565b1515610d3d57600080fd5b60008281526010602090815260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015610dd05780601f10610da557610100808354040283529160200191610dd0565b820191906000526020600020905b815481529060010190602001808311610db357829003601f168201915b50505050509050919050565b600160a060020a03918216600090815260096020908152604080832093909416825291909152205460ff1690565b6000806000610e176113bc565b5050506000918252506011602090815260409182902082516060810184528154600160a060020a0316808252600183015493820184905260029092015493018390529092565b600080610e6983610a33565b905080600160a060020a031684600160a060020a03161480610ea4575083600160a060020a0316610e998461062d565b600160a060020a0316145b80610eb45750610eb48185610ddc565b949350505050565b81600160a060020a0316610ecf82610a33565b600160a060020a031614610ee257600080fd5b600081815260076020526040902054600160a060020a031615610f595760008181526007602090815260408083208054600160a060020a031916905580518481529051600160a060020a038616927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35b5050565b6000806000610f6c8585611288565b6000848152600d6020908152604080832054600160a060020a0389168452600c90925290912054909350610fa790600163ffffffff61131116565b600160a060020a0386166000908152600c6020526040902080549193509083908110610fcf57fe5b9060005260206000200154905080600c600087600160a060020a0316600160a060020a031681526020019081526020016000208481548110151561100f57fe5b6000918252602080832090910192909255600160a060020a0387168152600c9091526040812080548490811061104157fe5b6000918252602080832090910192909255600160a060020a0387168152600c909152604090208054906110789060001983016113e7565b506000938452600d6020526040808520859055908452909220555050565b60006110a28383611323565b50600160a060020a039091166000908152600c6020908152604080832080546001810182559084528284208101859055938352600d909152902055565b600090815260066020526040902054600160a060020a0391821691161490565b60008061111485600160a060020a03166113a7565b1515611123576001915061127f565b84600160a060020a031663f0b9e5ba8786866040518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018084600160a060020a0316600160a060020a0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156111bb5781810151838201526020016111a3565b50505050905090810190601f1680156111e85780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b15801561120957600080fd5b505af115801561121d573d6000803e3d6000fd5b505050506040513d602081101561123357600080fd5b50517fffffffff0000000000000000000000000000000000000000000000000000000081167ff0b9e5ba0000000000000000000000000000000000000000000000000000000014925090505b50949350505050565b81600160a060020a031661129b82610a33565b600160a060020a0316146112ae57600080fd5b600160a060020a0382166000908152600860205260409020546112d890600163ffffffff61131116565b600160a060020a039092166000908152600860209081526040808320949094559181526006909152208054600160a060020a0319169055565b60008282111561131d57fe5b50900390565b600081815260066020526040902054600160a060020a03161561134557600080fd5b60008181526006602090815260408083208054600160a060020a031916600160a060020a038716908117909155835260089091529020546113879060016113af565b600160a060020a0390921660009081526008602052604090209190915550565b6000903b1190565b81810182811015610a5757fe5b6060604051908101604052806000600160a060020a0316815260200160008152602001600081525090565b8154818355818111156107285760008381526020902061072891810190830161062a91905b80821115611420576000815560010161140c565b50905600a165627a7a723058204a7491904c38d36510c87c3251992cae279a92efc50fa60d0945cf56f24121e10029`

// DeployRepoTransaction deploys a new Ethereum contract, binding an instance of RepoTransaction to it.
func DeployRepoTransaction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RepoTransaction, error) {
	parsed, err := abi.JSON(strings.NewReader(RepoTransactionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RepoTransactionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RepoTransaction{RepoTransactionCaller: RepoTransactionCaller{contract: contract}, RepoTransactionTransactor: RepoTransactionTransactor{contract: contract}, RepoTransactionFilterer: RepoTransactionFilterer{contract: contract}}, nil
}

// RepoTransaction is an auto generated Go binding around an Ethereum contract.
type RepoTransaction struct {
	RepoTransactionCaller     // Read-only binding to the contract
	RepoTransactionTransactor // Write-only binding to the contract
	RepoTransactionFilterer   // Log filterer for contract events
}

// RepoTransactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type RepoTransactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RepoTransactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RepoTransactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RepoTransactionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RepoTransactionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RepoTransactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RepoTransactionSession struct {
	Contract     *RepoTransaction  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RepoTransactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RepoTransactionCallerSession struct {
	Contract *RepoTransactionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// RepoTransactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RepoTransactionTransactorSession struct {
	Contract     *RepoTransactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// RepoTransactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type RepoTransactionRaw struct {
	Contract *RepoTransaction // Generic contract binding to access the raw methods on
}

// RepoTransactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RepoTransactionCallerRaw struct {
	Contract *RepoTransactionCaller // Generic read-only contract binding to access the raw methods on
}

// RepoTransactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RepoTransactionTransactorRaw struct {
	Contract *RepoTransactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRepoTransaction creates a new instance of RepoTransaction, bound to a specific deployed contract.
func NewRepoTransaction(address common.Address, backend bind.ContractBackend) (*RepoTransaction, error) {
	contract, err := bindRepoTransaction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RepoTransaction{RepoTransactionCaller: RepoTransactionCaller{contract: contract}, RepoTransactionTransactor: RepoTransactionTransactor{contract: contract}, RepoTransactionFilterer: RepoTransactionFilterer{contract: contract}}, nil
}

// NewRepoTransactionCaller creates a new read-only instance of RepoTransaction, bound to a specific deployed contract.
func NewRepoTransactionCaller(address common.Address, caller bind.ContractCaller) (*RepoTransactionCaller, error) {
	contract, err := bindRepoTransaction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RepoTransactionCaller{contract: contract}, nil
}

// NewRepoTransactionTransactor creates a new write-only instance of RepoTransaction, bound to a specific deployed contract.
func NewRepoTransactionTransactor(address common.Address, transactor bind.ContractTransactor) (*RepoTransactionTransactor, error) {
	contract, err := bindRepoTransaction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RepoTransactionTransactor{contract: contract}, nil
}

// NewRepoTransactionFilterer creates a new log filterer instance of RepoTransaction, bound to a specific deployed contract.
func NewRepoTransactionFilterer(address common.Address, filterer bind.ContractFilterer) (*RepoTransactionFilterer, error) {
	contract, err := bindRepoTransaction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RepoTransactionFilterer{contract: contract}, nil
}

// bindRepoTransaction binds a generic wrapper to an already deployed contract.
func bindRepoTransaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RepoTransactionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RepoTransaction *RepoTransactionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RepoTransaction.Contract.RepoTransactionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RepoTransaction *RepoTransactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RepoTransaction.Contract.RepoTransactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RepoTransaction *RepoTransactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RepoTransaction.Contract.RepoTransactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RepoTransaction *RepoTransactionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RepoTransaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RepoTransaction *RepoTransactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RepoTransaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RepoTransaction *RepoTransactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RepoTransaction.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_RepoTransaction *RepoTransactionCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepoTransaction.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_RepoTransaction *RepoTransactionSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _RepoTransaction.Contract.BalanceOf(&_RepoTransaction.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_RepoTransaction *RepoTransactionCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _RepoTransaction.Contract.BalanceOf(&_RepoTransaction.CallOpts, _owner)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_RepoTransaction *RepoTransactionCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RepoTransaction.contract.Call(opts, out, "ceoAddress")
	return *ret0, err
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_RepoTransaction *RepoTransactionSession) CeoAddress() (common.Address, error) {
	return _RepoTransaction.Contract.CeoAddress(&_RepoTransaction.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_RepoTransaction *RepoTransactionCallerSession) CeoAddress() (common.Address, error) {
	return _RepoTransaction.Contract.CeoAddress(&_RepoTransaction.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_RepoTransaction *RepoTransactionCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RepoTransaction.contract.Call(opts, out, "cfoAddress")
	return *ret0, err
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_RepoTransaction *RepoTransactionSession) CfoAddress() (common.Address, error) {
	return _RepoTransaction.Contract.CfoAddress(&_RepoTransaction.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_RepoTransaction *RepoTransactionCallerSession) CfoAddress() (common.Address, error) {
	return _RepoTransaction.Contract.CfoAddress(&_RepoTransaction.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_RepoTransaction *RepoTransactionCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RepoTransaction.contract.Call(opts, out, "cooAddress")
	return *ret0, err
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_RepoTransaction *RepoTransactionSession) CooAddress() (common.Address, error) {
	return _RepoTransaction.Contract.CooAddress(&_RepoTransaction.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_RepoTransaction *RepoTransactionCallerSession) CooAddress() (common.Address, error) {
	return _RepoTransaction.Contract.CooAddress(&_RepoTransaction.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_RepoTransaction *RepoTransactionCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RepoTransaction.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_RepoTransaction *RepoTransactionSession) Exists(_tokenId *big.Int) (bool, error) {
	return _RepoTransaction.Contract.Exists(&_RepoTransaction.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_RepoTransaction *RepoTransactionCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _RepoTransaction.Contract.Exists(&_RepoTransaction.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_RepoTransaction *RepoTransactionCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RepoTransaction.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_RepoTransaction *RepoTransactionSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _RepoTransaction.Contract.GetApproved(&_RepoTransaction.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_RepoTransaction *RepoTransactionCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _RepoTransaction.Contract.GetApproved(&_RepoTransaction.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(approved bool)
func (_RepoTransaction *RepoTransactionCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RepoTransaction.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(approved bool)
func (_RepoTransaction *RepoTransactionSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _RepoTransaction.Contract.IsApprovedForAll(&_RepoTransaction.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(approved bool)
func (_RepoTransaction *RepoTransactionCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _RepoTransaction.Contract.IsApprovedForAll(&_RepoTransaction.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RepoTransaction *RepoTransactionCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RepoTransaction.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RepoTransaction *RepoTransactionSession) Name() (string, error) {
	return _RepoTransaction.Contract.Name(&_RepoTransaction.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RepoTransaction *RepoTransactionCallerSession) Name() (string, error) {
	return _RepoTransaction.Contract.Name(&_RepoTransaction.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_RepoTransaction *RepoTransactionCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RepoTransaction.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_RepoTransaction *RepoTransactionSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _RepoTransaction.Contract.OwnerOf(&_RepoTransaction.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_RepoTransaction *RepoTransactionCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _RepoTransaction.Contract.OwnerOf(&_RepoTransaction.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_RepoTransaction *RepoTransactionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RepoTransaction.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_RepoTransaction *RepoTransactionSession) Paused() (bool, error) {
	return _RepoTransaction.Contract.Paused(&_RepoTransaction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_RepoTransaction *RepoTransactionCallerSession) Paused() (bool, error) {
	return _RepoTransaction.Contract.Paused(&_RepoTransaction.CallOpts)
}

// RepoIngotsInfo is a free data retrieval call binding the contract method 0xfaf20cc5.
//
// Solidity: function repoIngotsInfo(_tokenId uint256) constant returns(tokenId uint256, repoCount uint256, seller address)
func (_RepoTransaction *RepoTransactionCaller) RepoIngotsInfo(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	TokenId   *big.Int
	RepoCount *big.Int
	Seller    common.Address
}, error) {
	ret := new(struct {
		TokenId   *big.Int
		RepoCount *big.Int
		Seller    common.Address
	})
	out := ret
	err := _RepoTransaction.contract.Call(opts, out, "repoIngotsInfo", _tokenId)
	return *ret, err
}

// RepoIngotsInfo is a free data retrieval call binding the contract method 0xfaf20cc5.
//
// Solidity: function repoIngotsInfo(_tokenId uint256) constant returns(tokenId uint256, repoCount uint256, seller address)
func (_RepoTransaction *RepoTransactionSession) RepoIngotsInfo(_tokenId *big.Int) (struct {
	TokenId   *big.Int
	RepoCount *big.Int
	Seller    common.Address
}, error) {
	return _RepoTransaction.Contract.RepoIngotsInfo(&_RepoTransaction.CallOpts, _tokenId)
}

// RepoIngotsInfo is a free data retrieval call binding the contract method 0xfaf20cc5.
//
// Solidity: function repoIngotsInfo(_tokenId uint256) constant returns(tokenId uint256, repoCount uint256, seller address)
func (_RepoTransaction *RepoTransactionCallerSession) RepoIngotsInfo(_tokenId *big.Int) (struct {
	TokenId   *big.Int
	RepoCount *big.Int
	Seller    common.Address
}, error) {
	return _RepoTransaction.Contract.RepoIngotsInfo(&_RepoTransaction.CallOpts, _tokenId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_RepoTransaction *RepoTransactionCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RepoTransaction.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_RepoTransaction *RepoTransactionSession) Symbol() (string, error) {
	return _RepoTransaction.Contract.Symbol(&_RepoTransaction.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_RepoTransaction *RepoTransactionCallerSession) Symbol() (string, error) {
	return _RepoTransaction.Contract.Symbol(&_RepoTransaction.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_RepoTransaction *RepoTransactionCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepoTransaction.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_RepoTransaction *RepoTransactionSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _RepoTransaction.Contract.TokenByIndex(&_RepoTransaction.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_RepoTransaction *RepoTransactionCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _RepoTransaction.Contract.TokenByIndex(&_RepoTransaction.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_RepoTransaction *RepoTransactionCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepoTransaction.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_RepoTransaction *RepoTransactionSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _RepoTransaction.Contract.TokenOfOwnerByIndex(&_RepoTransaction.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_RepoTransaction *RepoTransactionCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _RepoTransaction.Contract.TokenOfOwnerByIndex(&_RepoTransaction.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_RepoTransaction *RepoTransactionCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RepoTransaction.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_RepoTransaction *RepoTransactionSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _RepoTransaction.Contract.TokenURI(&_RepoTransaction.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_RepoTransaction *RepoTransactionCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _RepoTransaction.Contract.TokenURI(&_RepoTransaction.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RepoTransaction *RepoTransactionCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepoTransaction.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RepoTransaction *RepoTransactionSession) TotalSupply() (*big.Int, error) {
	return _RepoTransaction.Contract.TotalSupply(&_RepoTransaction.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RepoTransaction *RepoTransactionCallerSession) TotalSupply() (*big.Int, error) {
	return _RepoTransaction.Contract.TotalSupply(&_RepoTransaction.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_RepoTransaction *RepoTransactionTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _RepoTransaction.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_RepoTransaction *RepoTransactionSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _RepoTransaction.Contract.Approve(&_RepoTransaction.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_RepoTransaction *RepoTransactionTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _RepoTransaction.Contract.Approve(&_RepoTransaction.TransactOpts, _to, _tokenId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RepoTransaction *RepoTransactionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RepoTransaction.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RepoTransaction *RepoTransactionSession) Pause() (*types.Transaction, error) {
	return _RepoTransaction.Contract.Pause(&_RepoTransaction.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RepoTransaction *RepoTransactionTransactorSession) Pause() (*types.Transaction, error) {
	return _RepoTransaction.Contract.Pause(&_RepoTransaction.TransactOpts)
}

// RepoIngots is a paid mutator transaction binding the contract method 0x949d354e.
//
// Solidity: function repoIngots(_tokenId uint256, _repoCount uint256) returns()
func (_RepoTransaction *RepoTransactionTransactor) RepoIngots(opts *bind.TransactOpts, _tokenId *big.Int, _repoCount *big.Int) (*types.Transaction, error) {
	return _RepoTransaction.contract.Transact(opts, "repoIngots", _tokenId, _repoCount)
}

// RepoIngots is a paid mutator transaction binding the contract method 0x949d354e.
//
// Solidity: function repoIngots(_tokenId uint256, _repoCount uint256) returns()
func (_RepoTransaction *RepoTransactionSession) RepoIngots(_tokenId *big.Int, _repoCount *big.Int) (*types.Transaction, error) {
	return _RepoTransaction.Contract.RepoIngots(&_RepoTransaction.TransactOpts, _tokenId, _repoCount)
}

// RepoIngots is a paid mutator transaction binding the contract method 0x949d354e.
//
// Solidity: function repoIngots(_tokenId uint256, _repoCount uint256) returns()
func (_RepoTransaction *RepoTransactionTransactorSession) RepoIngots(_tokenId *big.Int, _repoCount *big.Int) (*types.Transaction, error) {
	return _RepoTransaction.Contract.RepoIngots(&_RepoTransaction.TransactOpts, _tokenId, _repoCount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_RepoTransaction *RepoTransactionTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _RepoTransaction.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_RepoTransaction *RepoTransactionSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _RepoTransaction.Contract.SafeTransferFrom(&_RepoTransaction.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_RepoTransaction *RepoTransactionTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _RepoTransaction.Contract.SafeTransferFrom(&_RepoTransaction.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_RepoTransaction *RepoTransactionTransactor) SetApprovalForAll(opts *bind.TransactOpts, _to common.Address, _approved bool) (*types.Transaction, error) {
	return _RepoTransaction.contract.Transact(opts, "setApprovalForAll", _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_RepoTransaction *RepoTransactionSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _RepoTransaction.Contract.SetApprovalForAll(&_RepoTransaction.TransactOpts, _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_RepoTransaction *RepoTransactionTransactorSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _RepoTransaction.Contract.SetApprovalForAll(&_RepoTransaction.TransactOpts, _to, _approved)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_RepoTransaction *RepoTransactionTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _RepoTransaction.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_RepoTransaction *RepoTransactionSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _RepoTransaction.Contract.SetCEO(&_RepoTransaction.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_RepoTransaction *RepoTransactionTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _RepoTransaction.Contract.SetCEO(&_RepoTransaction.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_RepoTransaction *RepoTransactionTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _RepoTransaction.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_RepoTransaction *RepoTransactionSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _RepoTransaction.Contract.SetCFO(&_RepoTransaction.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_RepoTransaction *RepoTransactionTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _RepoTransaction.Contract.SetCFO(&_RepoTransaction.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_RepoTransaction *RepoTransactionTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _RepoTransaction.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_RepoTransaction *RepoTransactionSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _RepoTransaction.Contract.SetCOO(&_RepoTransaction.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_RepoTransaction *RepoTransactionTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _RepoTransaction.Contract.SetCOO(&_RepoTransaction.TransactOpts, _newCOO)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_RepoTransaction *RepoTransactionTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _RepoTransaction.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_RepoTransaction *RepoTransactionSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _RepoTransaction.Contract.TransferFrom(&_RepoTransaction.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_RepoTransaction *RepoTransactionTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _RepoTransaction.Contract.TransferFrom(&_RepoTransaction.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RepoTransaction *RepoTransactionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RepoTransaction.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RepoTransaction *RepoTransactionSession) Unpause() (*types.Transaction, error) {
	return _RepoTransaction.Contract.Unpause(&_RepoTransaction.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RepoTransaction *RepoTransactionTransactorSession) Unpause() (*types.Transaction, error) {
	return _RepoTransaction.Contract.Unpause(&_RepoTransaction.TransactOpts)
}

// RepoTransactionApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RepoTransaction contract.
type RepoTransactionApprovalIterator struct {
	Event *RepoTransactionApproval // Event containing the contract specifics and raw log

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
func (it *RepoTransactionApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepoTransactionApproval)
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
		it.Event = new(RepoTransactionApproval)
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
func (it *RepoTransactionApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepoTransactionApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepoTransactionApproval represents a Approval event raised by the RepoTransaction contract.
type RepoTransactionApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_RepoTransaction *RepoTransactionFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address) (*RepoTransactionApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _RepoTransaction.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return &RepoTransactionApprovalIterator{contract: _RepoTransaction.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_RepoTransaction *RepoTransactionFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RepoTransactionApproval, _owner []common.Address, _approved []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _RepoTransaction.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepoTransactionApproval)
				if err := _RepoTransaction.contract.UnpackLog(event, "Approval", log); err != nil {
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

// RepoTransactionApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the RepoTransaction contract.
type RepoTransactionApprovalForAllIterator struct {
	Event *RepoTransactionApprovalForAll // Event containing the contract specifics and raw log

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
func (it *RepoTransactionApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepoTransactionApprovalForAll)
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
		it.Event = new(RepoTransactionApprovalForAll)
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
func (it *RepoTransactionApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepoTransactionApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepoTransactionApprovalForAll represents a ApprovalForAll event raised by the RepoTransaction contract.
type RepoTransactionApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_RepoTransaction *RepoTransactionFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*RepoTransactionApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _RepoTransaction.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &RepoTransactionApprovalForAllIterator{contract: _RepoTransaction.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_RepoTransaction *RepoTransactionFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *RepoTransactionApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _RepoTransaction.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepoTransactionApprovalForAll)
				if err := _RepoTransaction.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// RepoTransactionContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the RepoTransaction contract.
type RepoTransactionContractUpgradeIterator struct {
	Event *RepoTransactionContractUpgrade // Event containing the contract specifics and raw log

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
func (it *RepoTransactionContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepoTransactionContractUpgrade)
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
		it.Event = new(RepoTransactionContractUpgrade)
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
func (it *RepoTransactionContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepoTransactionContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepoTransactionContractUpgrade represents a ContractUpgrade event raised by the RepoTransaction contract.
type RepoTransactionContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_RepoTransaction *RepoTransactionFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*RepoTransactionContractUpgradeIterator, error) {

	logs, sub, err := _RepoTransaction.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &RepoTransactionContractUpgradeIterator{contract: _RepoTransaction.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_RepoTransaction *RepoTransactionFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *RepoTransactionContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _RepoTransaction.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepoTransactionContractUpgrade)
				if err := _RepoTransaction.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// RepoTransactionCreateNewStampIterator is returned from FilterCreateNewStamp and is used to iterate over the raw logs and unpacked data for CreateNewStamp events raised by the RepoTransaction contract.
type RepoTransactionCreateNewStampIterator struct {
	Event *RepoTransactionCreateNewStamp // Event containing the contract specifics and raw log

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
func (it *RepoTransactionCreateNewStampIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepoTransactionCreateNewStamp)
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
		it.Event = new(RepoTransactionCreateNewStamp)
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
func (it *RepoTransactionCreateNewStampIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepoTransactionCreateNewStampIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepoTransactionCreateNewStamp represents a CreateNewStamp event raised by the RepoTransaction contract.
type RepoTransactionCreateNewStamp struct {
	StampId         uint32
	TotalAmount     uint32
	RemainingAmount uint32
	Name            [32]byte
	Year            uint16
	SetId           uint16
	MemberOfSetId   uint8
	Appearance      uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCreateNewStamp is a free log retrieval operation binding the contract event 0x4710117ee2ed24df4322c27506325e38cbbb47a788bc132e0c5d7d107f434256.
//
// Solidity: e CreateNewStamp(stampId uint32, totalAmount uint32, remainingAmount uint32, name bytes32, year uint16, setId uint16, memberOfSetId uint8, appearance uint8)
func (_RepoTransaction *RepoTransactionFilterer) FilterCreateNewStamp(opts *bind.FilterOpts) (*RepoTransactionCreateNewStampIterator, error) {

	logs, sub, err := _RepoTransaction.contract.FilterLogs(opts, "CreateNewStamp")
	if err != nil {
		return nil, err
	}
	return &RepoTransactionCreateNewStampIterator{contract: _RepoTransaction.contract, event: "CreateNewStamp", logs: logs, sub: sub}, nil
}

// WatchCreateNewStamp is a free log subscription operation binding the contract event 0x4710117ee2ed24df4322c27506325e38cbbb47a788bc132e0c5d7d107f434256.
//
// Solidity: e CreateNewStamp(stampId uint32, totalAmount uint32, remainingAmount uint32, name bytes32, year uint16, setId uint16, memberOfSetId uint8, appearance uint8)
func (_RepoTransaction *RepoTransactionFilterer) WatchCreateNewStamp(opts *bind.WatchOpts, sink chan<- *RepoTransactionCreateNewStamp) (event.Subscription, error) {

	logs, sub, err := _RepoTransaction.contract.WatchLogs(opts, "CreateNewStamp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepoTransactionCreateNewStamp)
				if err := _RepoTransaction.contract.UnpackLog(event, "CreateNewStamp", log); err != nil {
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

// RepoTransactionRepoIngotsSuccessfulIterator is returned from FilterRepoIngotsSuccessful and is used to iterate over the raw logs and unpacked data for RepoIngotsSuccessful events raised by the RepoTransaction contract.
type RepoTransactionRepoIngotsSuccessfulIterator struct {
	Event *RepoTransactionRepoIngotsSuccessful // Event containing the contract specifics and raw log

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
func (it *RepoTransactionRepoIngotsSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepoTransactionRepoIngotsSuccessful)
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
		it.Event = new(RepoTransactionRepoIngotsSuccessful)
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
func (it *RepoTransactionRepoIngotsSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepoTransactionRepoIngotsSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepoTransactionRepoIngotsSuccessful represents a RepoIngotsSuccessful event raised by the RepoTransaction contract.
type RepoTransactionRepoIngotsSuccessful struct {
	TokenId   *big.Int
	RepoCount *big.Int
	Seller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRepoIngotsSuccessful is a free log retrieval operation binding the contract event 0x50cd59a56cf7ed9d5b7a9978578d10cc8529c67bead200fdb943ce7e7c9089da.
//
// Solidity: e RepoIngotsSuccessful(tokenId uint256, repoCount uint256, seller address)
func (_RepoTransaction *RepoTransactionFilterer) FilterRepoIngotsSuccessful(opts *bind.FilterOpts) (*RepoTransactionRepoIngotsSuccessfulIterator, error) {

	logs, sub, err := _RepoTransaction.contract.FilterLogs(opts, "RepoIngotsSuccessful")
	if err != nil {
		return nil, err
	}
	return &RepoTransactionRepoIngotsSuccessfulIterator{contract: _RepoTransaction.contract, event: "RepoIngotsSuccessful", logs: logs, sub: sub}, nil
}

// WatchRepoIngotsSuccessful is a free log subscription operation binding the contract event 0x50cd59a56cf7ed9d5b7a9978578d10cc8529c67bead200fdb943ce7e7c9089da.
//
// Solidity: e RepoIngotsSuccessful(tokenId uint256, repoCount uint256, seller address)
func (_RepoTransaction *RepoTransactionFilterer) WatchRepoIngotsSuccessful(opts *bind.WatchOpts, sink chan<- *RepoTransactionRepoIngotsSuccessful) (event.Subscription, error) {

	logs, sub, err := _RepoTransaction.contract.WatchLogs(opts, "RepoIngotsSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepoTransactionRepoIngotsSuccessful)
				if err := _RepoTransaction.contract.UnpackLog(event, "RepoIngotsSuccessful", log); err != nil {
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

// RepoTransactionTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RepoTransaction contract.
type RepoTransactionTransferIterator struct {
	Event *RepoTransactionTransfer // Event containing the contract specifics and raw log

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
func (it *RepoTransactionTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepoTransactionTransfer)
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
		it.Event = new(RepoTransactionTransfer)
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
func (it *RepoTransactionTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepoTransactionTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepoTransactionTransfer represents a Transfer event raised by the RepoTransaction contract.
type RepoTransactionTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_RepoTransaction *RepoTransactionFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*RepoTransactionTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _RepoTransaction.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &RepoTransactionTransferIterator{contract: _RepoTransaction.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_RepoTransaction *RepoTransactionFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RepoTransactionTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _RepoTransaction.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepoTransactionTransfer)
				if err := _RepoTransaction.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058208d76b724952b195fc84e0242046f4afc9ef8db7c2c75ba53745a03d9bed5c4a60029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StampAccessControlABI is the input ABI used to generate the binding from.
const StampAccessControlABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// StampAccessControlBin is the compiled bytecode used for deploying new contracts.
const StampAccessControlBin = `0x60806040526002805460a060020a60ff021916905534801561002057600080fd5b5061043d806100306000396000f3006080604052600436106100985763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630519ce79811461009d5780630a0f8168146100ce57806327d7874c146100e35780632ba73c15146101065780633f4ba83a146101275780634e0a33791461013c5780635c975abb1461015d5780638456cb5914610186578063b047fb501461019b575b600080fd5b3480156100a957600080fd5b506100b26101b0565b60408051600160a060020a039092168252519081900360200190f35b3480156100da57600080fd5b506100b26101bf565b3480156100ef57600080fd5b50610104600160a060020a03600435166101ce565b005b34801561011257600080fd5b50610104600160a060020a0360043516610229565b34801561013357600080fd5b50610104610284565b34801561014857600080fd5b50610104600160a060020a03600435166102e4565b34801561016957600080fd5b5061017261033f565b604080519115158252519081900360200190f35b34801561019257600080fd5b50610104610360565b3480156101a757600080fd5b506100b2610402565b600154600160a060020a031681565b600054600160a060020a031681565b600054600160a060020a031633146101e557600080fd5b600160a060020a03811615156101fa57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a0316331461024057600080fd5b600160a060020a038116151561025557600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a0316331461029b57600080fd5b60025474010000000000000000000000000000000000000000900460ff1615156102c457600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b600054600160a060020a031633146102fb57600080fd5b600160a060020a038116151561031057600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60025474010000000000000000000000000000000000000000900460ff1681565b600254600160a060020a03163314806103835750600054600160a060020a031633145b806103985750600154600160a060020a031633145b15156103a357600080fd5b60025474010000000000000000000000000000000000000000900460ff16156103cb57600080fd5b6002805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b600254600160a060020a0316815600a165627a7a72305820522f9f35fd2fa254203cc8d5c2ba1721da0a29e00936d193e5304d8f5253876a0029`

// DeployStampAccessControl deploys a new Ethereum contract, binding an instance of StampAccessControl to it.
func DeployStampAccessControl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StampAccessControl, error) {
	parsed, err := abi.JSON(strings.NewReader(StampAccessControlABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StampAccessControlBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StampAccessControl{StampAccessControlCaller: StampAccessControlCaller{contract: contract}, StampAccessControlTransactor: StampAccessControlTransactor{contract: contract}, StampAccessControlFilterer: StampAccessControlFilterer{contract: contract}}, nil
}

// StampAccessControl is an auto generated Go binding around an Ethereum contract.
type StampAccessControl struct {
	StampAccessControlCaller     // Read-only binding to the contract
	StampAccessControlTransactor // Write-only binding to the contract
	StampAccessControlFilterer   // Log filterer for contract events
}

// StampAccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type StampAccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StampAccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StampAccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StampAccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StampAccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StampAccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StampAccessControlSession struct {
	Contract     *StampAccessControl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StampAccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StampAccessControlCallerSession struct {
	Contract *StampAccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// StampAccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StampAccessControlTransactorSession struct {
	Contract     *StampAccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// StampAccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type StampAccessControlRaw struct {
	Contract *StampAccessControl // Generic contract binding to access the raw methods on
}

// StampAccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StampAccessControlCallerRaw struct {
	Contract *StampAccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// StampAccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StampAccessControlTransactorRaw struct {
	Contract *StampAccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStampAccessControl creates a new instance of StampAccessControl, bound to a specific deployed contract.
func NewStampAccessControl(address common.Address, backend bind.ContractBackend) (*StampAccessControl, error) {
	contract, err := bindStampAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StampAccessControl{StampAccessControlCaller: StampAccessControlCaller{contract: contract}, StampAccessControlTransactor: StampAccessControlTransactor{contract: contract}, StampAccessControlFilterer: StampAccessControlFilterer{contract: contract}}, nil
}

// NewStampAccessControlCaller creates a new read-only instance of StampAccessControl, bound to a specific deployed contract.
func NewStampAccessControlCaller(address common.Address, caller bind.ContractCaller) (*StampAccessControlCaller, error) {
	contract, err := bindStampAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StampAccessControlCaller{contract: contract}, nil
}

// NewStampAccessControlTransactor creates a new write-only instance of StampAccessControl, bound to a specific deployed contract.
func NewStampAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*StampAccessControlTransactor, error) {
	contract, err := bindStampAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StampAccessControlTransactor{contract: contract}, nil
}

// NewStampAccessControlFilterer creates a new log filterer instance of StampAccessControl, bound to a specific deployed contract.
func NewStampAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*StampAccessControlFilterer, error) {
	contract, err := bindStampAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StampAccessControlFilterer{contract: contract}, nil
}

// bindStampAccessControl binds a generic wrapper to an already deployed contract.
func bindStampAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StampAccessControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StampAccessControl *StampAccessControlRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StampAccessControl.Contract.StampAccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StampAccessControl *StampAccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampAccessControl.Contract.StampAccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StampAccessControl *StampAccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StampAccessControl.Contract.StampAccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StampAccessControl *StampAccessControlCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StampAccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StampAccessControl *StampAccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampAccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StampAccessControl *StampAccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StampAccessControl.Contract.contract.Transact(opts, method, params...)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_StampAccessControl *StampAccessControlCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampAccessControl.contract.Call(opts, out, "ceoAddress")
	return *ret0, err
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_StampAccessControl *StampAccessControlSession) CeoAddress() (common.Address, error) {
	return _StampAccessControl.Contract.CeoAddress(&_StampAccessControl.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_StampAccessControl *StampAccessControlCallerSession) CeoAddress() (common.Address, error) {
	return _StampAccessControl.Contract.CeoAddress(&_StampAccessControl.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_StampAccessControl *StampAccessControlCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampAccessControl.contract.Call(opts, out, "cfoAddress")
	return *ret0, err
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_StampAccessControl *StampAccessControlSession) CfoAddress() (common.Address, error) {
	return _StampAccessControl.Contract.CfoAddress(&_StampAccessControl.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_StampAccessControl *StampAccessControlCallerSession) CfoAddress() (common.Address, error) {
	return _StampAccessControl.Contract.CfoAddress(&_StampAccessControl.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_StampAccessControl *StampAccessControlCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampAccessControl.contract.Call(opts, out, "cooAddress")
	return *ret0, err
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_StampAccessControl *StampAccessControlSession) CooAddress() (common.Address, error) {
	return _StampAccessControl.Contract.CooAddress(&_StampAccessControl.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_StampAccessControl *StampAccessControlCallerSession) CooAddress() (common.Address, error) {
	return _StampAccessControl.Contract.CooAddress(&_StampAccessControl.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_StampAccessControl *StampAccessControlCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StampAccessControl.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_StampAccessControl *StampAccessControlSession) Paused() (bool, error) {
	return _StampAccessControl.Contract.Paused(&_StampAccessControl.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_StampAccessControl *StampAccessControlCallerSession) Paused() (bool, error) {
	return _StampAccessControl.Contract.Paused(&_StampAccessControl.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StampAccessControl *StampAccessControlTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampAccessControl.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StampAccessControl *StampAccessControlSession) Pause() (*types.Transaction, error) {
	return _StampAccessControl.Contract.Pause(&_StampAccessControl.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StampAccessControl *StampAccessControlTransactorSession) Pause() (*types.Transaction, error) {
	return _StampAccessControl.Contract.Pause(&_StampAccessControl.TransactOpts)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_StampAccessControl *StampAccessControlTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _StampAccessControl.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_StampAccessControl *StampAccessControlSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _StampAccessControl.Contract.SetCEO(&_StampAccessControl.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_StampAccessControl *StampAccessControlTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _StampAccessControl.Contract.SetCEO(&_StampAccessControl.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_StampAccessControl *StampAccessControlTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _StampAccessControl.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_StampAccessControl *StampAccessControlSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _StampAccessControl.Contract.SetCFO(&_StampAccessControl.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_StampAccessControl *StampAccessControlTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _StampAccessControl.Contract.SetCFO(&_StampAccessControl.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_StampAccessControl *StampAccessControlTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _StampAccessControl.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_StampAccessControl *StampAccessControlSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _StampAccessControl.Contract.SetCOO(&_StampAccessControl.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_StampAccessControl *StampAccessControlTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _StampAccessControl.Contract.SetCOO(&_StampAccessControl.TransactOpts, _newCOO)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StampAccessControl *StampAccessControlTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampAccessControl.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StampAccessControl *StampAccessControlSession) Unpause() (*types.Transaction, error) {
	return _StampAccessControl.Contract.Unpause(&_StampAccessControl.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StampAccessControl *StampAccessControlTransactorSession) Unpause() (*types.Transaction, error) {
	return _StampAccessControl.Contract.Unpause(&_StampAccessControl.TransactOpts)
}

// StampAccessControlContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the StampAccessControl contract.
type StampAccessControlContractUpgradeIterator struct {
	Event *StampAccessControlContractUpgrade // Event containing the contract specifics and raw log

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
func (it *StampAccessControlContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampAccessControlContractUpgrade)
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
		it.Event = new(StampAccessControlContractUpgrade)
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
func (it *StampAccessControlContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampAccessControlContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampAccessControlContractUpgrade represents a ContractUpgrade event raised by the StampAccessControl contract.
type StampAccessControlContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_StampAccessControl *StampAccessControlFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*StampAccessControlContractUpgradeIterator, error) {

	logs, sub, err := _StampAccessControl.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &StampAccessControlContractUpgradeIterator{contract: _StampAccessControl.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_StampAccessControl *StampAccessControlFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *StampAccessControlContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _StampAccessControl.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampAccessControlContractUpgrade)
				if err := _StampAccessControl.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// StampBaseABI is the input ABI used to generate the binding from.
const StampBaseABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"approved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"stampId\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"totalAmount\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"remainingAmount\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"year\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"setId\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"memberOfSetId\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"appearance\",\"type\":\"uint8\"}],\"name\":\"CreateNewStamp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// StampBaseBin is the compiled bytecode used for deploying new contracts.
const StampBaseBin = `0x6002805460a060020a60ff021916905560c0604052600a60808190527f5374616d70546f6b656e0000000000000000000000000000000000000000000060a09081526200004e919081620000aa565b506040805180820190915260028082527f535400000000000000000000000000000000000000000000000000000000000060209092019182526200009591600b91620000aa565b50348015620000a357600080fd5b506200014f565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000ed57805160ff19168380011785556200011d565b828001600101855582156200011d579182015b828111156200011d57825182559160200191906001019062000100565b506200012b9291506200012f565b5090565b6200014c91905b808211156200012b576000815560010162000136565b90565b611280806200015f6000396000f3006080604052600436106101485763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630519ce79811461014d57806306fdde031461017e578063081812fc14610208578063095ea7b3146102205780630a0f81681461024657806318160ddd1461025b57806323b872dd1461028257806327d7874c146102ac5780632ba73c15146102cd5780632f745c59146102ee5780633f4ba83a1461031257806342842e0e146103275780634e0a3379146103515780634f558e79146103725780634f6ccce71461039e5780635c975abb146103b65780636352211e146103cb57806370a08231146103e35780638456cb591461040457806395d89b4114610419578063a22cb4651461042e578063b047fb5014610454578063b88d4fde14610469578063c87b56dd146104d8578063e985e9c5146104f0575b600080fd5b34801561015957600080fd5b50610162610517565b60408051600160a060020a039092168252519081900360200190f35b34801561018a57600080fd5b50610193610526565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101cd5781810151838201526020016101b5565b50505050905090810190601f1680156101fa5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561021457600080fd5b506101626004356105bd565b34801561022c57600080fd5b50610244600160a060020a03600435166024356105d8565b005b34801561025257600080fd5b506101626106bd565b34801561026757600080fd5b506102706106cc565b60408051918252519081900360200190f35b34801561028e57600080fd5b50610244600160a060020a03600435811690602435166044356106d2565b3480156102b857600080fd5b50610244600160a060020a0360043516610781565b3480156102d957600080fd5b50610244600160a060020a03600435166107cf565b3480156102fa57600080fd5b50610270600160a060020a036004351660243561081d565b34801561031e57600080fd5b5061024461086a565b34801561033357600080fd5b50610244600160a060020a03600435811690602435166044356108ca565b34801561035d57600080fd5b50610244600160a060020a0360043516610902565b34801561037e57600080fd5b5061038a600435610950565b604080519115158252519081900360200190f35b3480156103aa57600080fd5b5061027060043561096d565b3480156103c257600080fd5b5061038a6109a2565b3480156103d757600080fd5b506101626004356109c3565b3480156103ef57600080fd5b50610270600160a060020a03600435166109ed565b34801561041057600080fd5b50610244610a20565b34801561042557600080fd5b50610193610ac2565b34801561043a57600080fd5b50610244600160a060020a03600435166024351515610b23565b34801561046057600080fd5b50610162610ba7565b34801561047557600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261024494600160a060020a038135811695602480359092169560443595369560849401918190840183828082843750949750610bb69650505050505050565b3480156104e457600080fd5b50610193600435610bf5565b3480156104fc57600080fd5b5061038a600160a060020a0360043581169060243516610caa565b600154600160a060020a031681565b600a8054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105b25780601f10610587576101008083540402835291602001916105b2565b820191906000526020600020905b81548152906001019060200180831161059557829003601f168201915b505050505090505b90565b600090815260076020526040902054600160a060020a031690565b60006105e3826109c3565b9050600160a060020a0383811690821614156105fe57600080fd5b33600160a060020a038216148061061a575061061a8133610caa565b151561062557600080fd5b6000610630836105bd565b600160a060020a031614158061064e5750600160a060020a03831615155b156106b8576000828152600760209081526040918290208054600160a060020a031916600160a060020a03878116918217909255835186815293519093918516927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35b505050565b600054600160a060020a031681565b600e5490565b806106dd3382610cd8565b15156106e857600080fd5b600160a060020a03841615156106fd57600080fd5b600160a060020a038316151561071257600080fd5b61071c8483610d37565b6107268483610dd8565b6107308383610f11565b82600160a060020a031684600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a350505050565b600054600160a060020a0316331461079857600080fd5b600160a060020a03811615156107ad57600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a031633146107e657600080fd5b600160a060020a03811615156107fb57600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b6000610828836109ed565b821061083357600080fd5b600160a060020a0383166000908152600c6020526040902080548390811061085757fe5b9060005260206000200154905092915050565b600054600160a060020a0316331461088157600080fd5b60025474010000000000000000000000000000000000000000900460ff1615156108aa57600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b806108d53382610cd8565b15156108e057600080fd5b6108fc8484846020604051908101604052806000815250610bb6565b50505050565b600054600160a060020a0316331461091957600080fd5b600160a060020a038116151561092e57600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b600090815260066020526040902054600160a060020a0316151590565b60006109776106cc565b821061098257600080fd5b600e80548390811061099057fe5b90600052602060002001549050919050565b60025474010000000000000000000000000000000000000000900460ff1681565b600081815260066020526040812054600160a060020a03168015156109e757600080fd5b92915050565b6000600160a060020a0382161515610a0457600080fd5b50600160a060020a031660009081526008602052604090205490565b600254600160a060020a0316331480610a435750600054600160a060020a031633145b80610a585750600154600160a060020a031633145b1515610a6357600080fd5b60025474010000000000000000000000000000000000000000900460ff1615610a8b57600080fd5b6002805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b600b8054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105b25780601f10610587576101008083540402835291602001916105b2565b600160a060020a038216331415610b3957600080fd5b336000818152600960209081526040808320600160a060020a03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b600254600160a060020a031681565b81610bc13382610cd8565b1515610bcc57600080fd5b610bd78585856106d2565b610be385858585610f5a565b1515610bee57600080fd5b5050505050565b6060610c0082610950565b1515610c0b57600080fd5b60008281526010602090815260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015610c9e5780601f10610c7357610100808354040283529160200191610c9e565b820191906000526020600020905b815481529060010190602001808311610c8157829003601f168201915b50505050509050919050565b600160a060020a03918216600090815260096020908152604080832093909416825291909152205460ff1690565b600080610ce4836109c3565b905080600160a060020a031684600160a060020a03161480610d1f575083600160a060020a0316610d14846105bd565b600160a060020a0316145b80610d2f5750610d2f8185610caa565b949350505050565b81600160a060020a0316610d4a826109c3565b600160a060020a031614610d5d57600080fd5b600081815260076020526040902054600160a060020a031615610dd45760008181526007602090815260408083208054600160a060020a031916905580518481529051600160a060020a038616927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35b5050565b6000806000610de785856110e3565b6000848152600d6020908152604080832054600160a060020a0389168452600c90925290912054909350610e2290600163ffffffff61116c16565b600160a060020a0386166000908152600c6020526040902080549193509083908110610e4a57fe5b9060005260206000200154905080600c600087600160a060020a0316600160a060020a0316815260200190815260200160002084815481101515610e8a57fe5b6000918252602080832090910192909255600160a060020a0387168152600c90915260408120805484908110610ebc57fe5b6000918252602080832090910192909255600160a060020a0387168152600c90915260409020805490610ef3906000198301611217565b506000938452600d6020526040808520859055908452909220555050565b6000610f1d838361117e565b50600160a060020a039091166000908152600c6020908152604080832080546001810182559084528284208101859055938352600d909152902055565b600080610f6f85600160a060020a0316611202565b1515610f7e57600191506110da565b84600160a060020a031663f0b9e5ba8786866040518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018084600160a060020a0316600160a060020a0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611016578181015183820152602001610ffe565b50505050905090810190601f1680156110435780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b15801561106457600080fd5b505af1158015611078573d6000803e3d6000fd5b505050506040513d602081101561108e57600080fd5b50517fffffffff0000000000000000000000000000000000000000000000000000000081167ff0b9e5ba0000000000000000000000000000000000000000000000000000000014925090505b50949350505050565b81600160a060020a03166110f6826109c3565b600160a060020a03161461110957600080fd5b600160a060020a03821660009081526008602052604090205461113390600163ffffffff61116c16565b600160a060020a039092166000908152600860209081526040808320949094559181526006909152208054600160a060020a0319169055565b60008282111561117857fe5b50900390565b600081815260066020526040902054600160a060020a0316156111a057600080fd5b60008181526006602090815260408083208054600160a060020a031916600160a060020a038716908117909155835260089091529020546111e290600161120a565b600160a060020a0390921660009081526008602052604090209190915550565b6000903b1190565b818101828110156109e757fe5b8154818355818111156106b8576000838152602090206106b89181019083016105ba91905b80821115611250576000815560010161123c565b50905600a165627a7a723058206b8007c8d662c131d592e72c134a15a3441b9eb55e056ff17b2a1d4a3655c9810029`

// DeployStampBase deploys a new Ethereum contract, binding an instance of StampBase to it.
func DeployStampBase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StampBase, error) {
	parsed, err := abi.JSON(strings.NewReader(StampBaseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StampBaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StampBase{StampBaseCaller: StampBaseCaller{contract: contract}, StampBaseTransactor: StampBaseTransactor{contract: contract}, StampBaseFilterer: StampBaseFilterer{contract: contract}}, nil
}

// StampBase is an auto generated Go binding around an Ethereum contract.
type StampBase struct {
	StampBaseCaller     // Read-only binding to the contract
	StampBaseTransactor // Write-only binding to the contract
	StampBaseFilterer   // Log filterer for contract events
}

// StampBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type StampBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StampBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StampBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StampBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StampBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StampBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StampBaseSession struct {
	Contract     *StampBase        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StampBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StampBaseCallerSession struct {
	Contract *StampBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// StampBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StampBaseTransactorSession struct {
	Contract     *StampBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StampBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type StampBaseRaw struct {
	Contract *StampBase // Generic contract binding to access the raw methods on
}

// StampBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StampBaseCallerRaw struct {
	Contract *StampBaseCaller // Generic read-only contract binding to access the raw methods on
}

// StampBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StampBaseTransactorRaw struct {
	Contract *StampBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStampBase creates a new instance of StampBase, bound to a specific deployed contract.
func NewStampBase(address common.Address, backend bind.ContractBackend) (*StampBase, error) {
	contract, err := bindStampBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StampBase{StampBaseCaller: StampBaseCaller{contract: contract}, StampBaseTransactor: StampBaseTransactor{contract: contract}, StampBaseFilterer: StampBaseFilterer{contract: contract}}, nil
}

// NewStampBaseCaller creates a new read-only instance of StampBase, bound to a specific deployed contract.
func NewStampBaseCaller(address common.Address, caller bind.ContractCaller) (*StampBaseCaller, error) {
	contract, err := bindStampBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StampBaseCaller{contract: contract}, nil
}

// NewStampBaseTransactor creates a new write-only instance of StampBase, bound to a specific deployed contract.
func NewStampBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*StampBaseTransactor, error) {
	contract, err := bindStampBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StampBaseTransactor{contract: contract}, nil
}

// NewStampBaseFilterer creates a new log filterer instance of StampBase, bound to a specific deployed contract.
func NewStampBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*StampBaseFilterer, error) {
	contract, err := bindStampBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StampBaseFilterer{contract: contract}, nil
}

// bindStampBase binds a generic wrapper to an already deployed contract.
func bindStampBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StampBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StampBase *StampBaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StampBase.Contract.StampBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StampBase *StampBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampBase.Contract.StampBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StampBase *StampBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StampBase.Contract.StampBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StampBase *StampBaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StampBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StampBase *StampBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StampBase *StampBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StampBase.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StampBase *StampBaseCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampBase.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StampBase *StampBaseSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StampBase.Contract.BalanceOf(&_StampBase.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StampBase *StampBaseCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StampBase.Contract.BalanceOf(&_StampBase.CallOpts, _owner)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_StampBase *StampBaseCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampBase.contract.Call(opts, out, "ceoAddress")
	return *ret0, err
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_StampBase *StampBaseSession) CeoAddress() (common.Address, error) {
	return _StampBase.Contract.CeoAddress(&_StampBase.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_StampBase *StampBaseCallerSession) CeoAddress() (common.Address, error) {
	return _StampBase.Contract.CeoAddress(&_StampBase.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_StampBase *StampBaseCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampBase.contract.Call(opts, out, "cfoAddress")
	return *ret0, err
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_StampBase *StampBaseSession) CfoAddress() (common.Address, error) {
	return _StampBase.Contract.CfoAddress(&_StampBase.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_StampBase *StampBaseCallerSession) CfoAddress() (common.Address, error) {
	return _StampBase.Contract.CfoAddress(&_StampBase.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_StampBase *StampBaseCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampBase.contract.Call(opts, out, "cooAddress")
	return *ret0, err
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_StampBase *StampBaseSession) CooAddress() (common.Address, error) {
	return _StampBase.Contract.CooAddress(&_StampBase.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_StampBase *StampBaseCallerSession) CooAddress() (common.Address, error) {
	return _StampBase.Contract.CooAddress(&_StampBase.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_StampBase *StampBaseCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StampBase.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_StampBase *StampBaseSession) Exists(_tokenId *big.Int) (bool, error) {
	return _StampBase.Contract.Exists(&_StampBase.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_StampBase *StampBaseCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _StampBase.Contract.Exists(&_StampBase.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_StampBase *StampBaseCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampBase.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_StampBase *StampBaseSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _StampBase.Contract.GetApproved(&_StampBase.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_StampBase *StampBaseCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _StampBase.Contract.GetApproved(&_StampBase.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(approved bool)
func (_StampBase *StampBaseCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StampBase.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(approved bool)
func (_StampBase *StampBaseSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _StampBase.Contract.IsApprovedForAll(&_StampBase.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(approved bool)
func (_StampBase *StampBaseCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _StampBase.Contract.IsApprovedForAll(&_StampBase.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StampBase *StampBaseCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StampBase.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StampBase *StampBaseSession) Name() (string, error) {
	return _StampBase.Contract.Name(&_StampBase.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StampBase *StampBaseCallerSession) Name() (string, error) {
	return _StampBase.Contract.Name(&_StampBase.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_StampBase *StampBaseCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampBase.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_StampBase *StampBaseSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _StampBase.Contract.OwnerOf(&_StampBase.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_StampBase *StampBaseCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _StampBase.Contract.OwnerOf(&_StampBase.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_StampBase *StampBaseCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StampBase.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_StampBase *StampBaseSession) Paused() (bool, error) {
	return _StampBase.Contract.Paused(&_StampBase.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_StampBase *StampBaseCallerSession) Paused() (bool, error) {
	return _StampBase.Contract.Paused(&_StampBase.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_StampBase *StampBaseCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StampBase.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_StampBase *StampBaseSession) Symbol() (string, error) {
	return _StampBase.Contract.Symbol(&_StampBase.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_StampBase *StampBaseCallerSession) Symbol() (string, error) {
	return _StampBase.Contract.Symbol(&_StampBase.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_StampBase *StampBaseCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampBase.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_StampBase *StampBaseSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _StampBase.Contract.TokenByIndex(&_StampBase.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_StampBase *StampBaseCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _StampBase.Contract.TokenByIndex(&_StampBase.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_StampBase *StampBaseCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampBase.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_StampBase *StampBaseSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _StampBase.Contract.TokenOfOwnerByIndex(&_StampBase.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_StampBase *StampBaseCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _StampBase.Contract.TokenOfOwnerByIndex(&_StampBase.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_StampBase *StampBaseCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StampBase.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_StampBase *StampBaseSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _StampBase.Contract.TokenURI(&_StampBase.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_StampBase *StampBaseCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _StampBase.Contract.TokenURI(&_StampBase.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StampBase *StampBaseCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampBase.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StampBase *StampBaseSession) TotalSupply() (*big.Int, error) {
	return _StampBase.Contract.TotalSupply(&_StampBase.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StampBase *StampBaseCallerSession) TotalSupply() (*big.Int, error) {
	return _StampBase.Contract.TotalSupply(&_StampBase.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_StampBase *StampBaseTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampBase.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_StampBase *StampBaseSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampBase.Contract.Approve(&_StampBase.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_StampBase *StampBaseTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampBase.Contract.Approve(&_StampBase.TransactOpts, _to, _tokenId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StampBase *StampBaseTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampBase.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StampBase *StampBaseSession) Pause() (*types.Transaction, error) {
	return _StampBase.Contract.Pause(&_StampBase.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StampBase *StampBaseTransactorSession) Pause() (*types.Transaction, error) {
	return _StampBase.Contract.Pause(&_StampBase.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_StampBase *StampBaseTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _StampBase.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_StampBase *StampBaseSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _StampBase.Contract.SafeTransferFrom(&_StampBase.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_StampBase *StampBaseTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _StampBase.Contract.SafeTransferFrom(&_StampBase.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_StampBase *StampBaseTransactor) SetApprovalForAll(opts *bind.TransactOpts, _to common.Address, _approved bool) (*types.Transaction, error) {
	return _StampBase.contract.Transact(opts, "setApprovalForAll", _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_StampBase *StampBaseSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _StampBase.Contract.SetApprovalForAll(&_StampBase.TransactOpts, _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_StampBase *StampBaseTransactorSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _StampBase.Contract.SetApprovalForAll(&_StampBase.TransactOpts, _to, _approved)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_StampBase *StampBaseTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _StampBase.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_StampBase *StampBaseSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _StampBase.Contract.SetCEO(&_StampBase.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_StampBase *StampBaseTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _StampBase.Contract.SetCEO(&_StampBase.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_StampBase *StampBaseTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _StampBase.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_StampBase *StampBaseSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _StampBase.Contract.SetCFO(&_StampBase.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_StampBase *StampBaseTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _StampBase.Contract.SetCFO(&_StampBase.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_StampBase *StampBaseTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _StampBase.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_StampBase *StampBaseSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _StampBase.Contract.SetCOO(&_StampBase.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_StampBase *StampBaseTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _StampBase.Contract.SetCOO(&_StampBase.TransactOpts, _newCOO)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_StampBase *StampBaseTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampBase.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_StampBase *StampBaseSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampBase.Contract.TransferFrom(&_StampBase.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_StampBase *StampBaseTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampBase.Contract.TransferFrom(&_StampBase.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StampBase *StampBaseTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampBase.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StampBase *StampBaseSession) Unpause() (*types.Transaction, error) {
	return _StampBase.Contract.Unpause(&_StampBase.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StampBase *StampBaseTransactorSession) Unpause() (*types.Transaction, error) {
	return _StampBase.Contract.Unpause(&_StampBase.TransactOpts)
}

// StampBaseApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StampBase contract.
type StampBaseApprovalIterator struct {
	Event *StampBaseApproval // Event containing the contract specifics and raw log

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
func (it *StampBaseApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampBaseApproval)
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
		it.Event = new(StampBaseApproval)
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
func (it *StampBaseApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampBaseApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampBaseApproval represents a Approval event raised by the StampBase contract.
type StampBaseApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_StampBase *StampBaseFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address) (*StampBaseApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _StampBase.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return &StampBaseApprovalIterator{contract: _StampBase.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_StampBase *StampBaseFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StampBaseApproval, _owner []common.Address, _approved []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _StampBase.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampBaseApproval)
				if err := _StampBase.contract.UnpackLog(event, "Approval", log); err != nil {
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

// StampBaseApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the StampBase contract.
type StampBaseApprovalForAllIterator struct {
	Event *StampBaseApprovalForAll // Event containing the contract specifics and raw log

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
func (it *StampBaseApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampBaseApprovalForAll)
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
		it.Event = new(StampBaseApprovalForAll)
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
func (it *StampBaseApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampBaseApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampBaseApprovalForAll represents a ApprovalForAll event raised by the StampBase contract.
type StampBaseApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_StampBase *StampBaseFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*StampBaseApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _StampBase.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &StampBaseApprovalForAllIterator{contract: _StampBase.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_StampBase *StampBaseFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *StampBaseApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _StampBase.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampBaseApprovalForAll)
				if err := _StampBase.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// StampBaseContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the StampBase contract.
type StampBaseContractUpgradeIterator struct {
	Event *StampBaseContractUpgrade // Event containing the contract specifics and raw log

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
func (it *StampBaseContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampBaseContractUpgrade)
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
		it.Event = new(StampBaseContractUpgrade)
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
func (it *StampBaseContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampBaseContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampBaseContractUpgrade represents a ContractUpgrade event raised by the StampBase contract.
type StampBaseContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_StampBase *StampBaseFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*StampBaseContractUpgradeIterator, error) {

	logs, sub, err := _StampBase.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &StampBaseContractUpgradeIterator{contract: _StampBase.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_StampBase *StampBaseFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *StampBaseContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _StampBase.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampBaseContractUpgrade)
				if err := _StampBase.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// StampBaseCreateNewStampIterator is returned from FilterCreateNewStamp and is used to iterate over the raw logs and unpacked data for CreateNewStamp events raised by the StampBase contract.
type StampBaseCreateNewStampIterator struct {
	Event *StampBaseCreateNewStamp // Event containing the contract specifics and raw log

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
func (it *StampBaseCreateNewStampIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampBaseCreateNewStamp)
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
		it.Event = new(StampBaseCreateNewStamp)
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
func (it *StampBaseCreateNewStampIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampBaseCreateNewStampIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampBaseCreateNewStamp represents a CreateNewStamp event raised by the StampBase contract.
type StampBaseCreateNewStamp struct {
	StampId         uint32
	TotalAmount     uint32
	RemainingAmount uint32
	Name            [32]byte
	Year            uint16
	SetId           uint16
	MemberOfSetId   uint8
	Appearance      uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCreateNewStamp is a free log retrieval operation binding the contract event 0x4710117ee2ed24df4322c27506325e38cbbb47a788bc132e0c5d7d107f434256.
//
// Solidity: e CreateNewStamp(stampId uint32, totalAmount uint32, remainingAmount uint32, name bytes32, year uint16, setId uint16, memberOfSetId uint8, appearance uint8)
func (_StampBase *StampBaseFilterer) FilterCreateNewStamp(opts *bind.FilterOpts) (*StampBaseCreateNewStampIterator, error) {

	logs, sub, err := _StampBase.contract.FilterLogs(opts, "CreateNewStamp")
	if err != nil {
		return nil, err
	}
	return &StampBaseCreateNewStampIterator{contract: _StampBase.contract, event: "CreateNewStamp", logs: logs, sub: sub}, nil
}

// WatchCreateNewStamp is a free log subscription operation binding the contract event 0x4710117ee2ed24df4322c27506325e38cbbb47a788bc132e0c5d7d107f434256.
//
// Solidity: e CreateNewStamp(stampId uint32, totalAmount uint32, remainingAmount uint32, name bytes32, year uint16, setId uint16, memberOfSetId uint8, appearance uint8)
func (_StampBase *StampBaseFilterer) WatchCreateNewStamp(opts *bind.WatchOpts, sink chan<- *StampBaseCreateNewStamp) (event.Subscription, error) {

	logs, sub, err := _StampBase.contract.WatchLogs(opts, "CreateNewStamp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampBaseCreateNewStamp)
				if err := _StampBase.contract.UnpackLog(event, "CreateNewStamp", log); err != nil {
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

// StampBaseTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StampBase contract.
type StampBaseTransferIterator struct {
	Event *StampBaseTransfer // Event containing the contract specifics and raw log

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
func (it *StampBaseTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampBaseTransfer)
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
		it.Event = new(StampBaseTransfer)
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
func (it *StampBaseTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampBaseTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampBaseTransfer represents a Transfer event raised by the StampBase contract.
type StampBaseTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_StampBase *StampBaseFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*StampBaseTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _StampBase.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &StampBaseTransferIterator{contract: _StampBase.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_StampBase *StampBaseFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StampBaseTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _StampBase.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampBaseTransfer)
				if err := _StampBase.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// StampCollectionABI is the input ABI used to generate the binding from.
const StampCollectionABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"stampInfo\",\"outputs\":[{\"name\":\"stampId\",\"type\":\"uint32\"},{\"name\":\"year\",\"type\":\"uint16\"},{\"name\":\"setId\",\"type\":\"uint16\"},{\"name\":\"memberOfSetId\",\"type\":\"uint8\"},{\"name\":\"totalAmount\",\"type\":\"uint32\"},{\"name\":\"remainingAmount\",\"type\":\"uint32\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"appearance\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STAMP_STARTING_PRICE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"promoCreatedCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stampId\",\"type\":\"uint32\"},{\"name\":\"_totalAmount\",\"type\":\"uint32\"},{\"name\":\"_remainingAmount\",\"type\":\"uint32\"},{\"name\":\"_name\",\"type\":\"bytes32\"},{\"name\":\"_year\",\"type\":\"uint16\"},{\"name\":\"_setId\",\"type\":\"uint16\"},{\"name\":\"_memberOfSetId\",\"type\":\"uint8\"},{\"name\":\"_appearance\",\"type\":\"uint8\"}],\"name\":\"releaseNewStampToTransaction\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STAMP_AUCTION_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_v2Address\",\"type\":\"address\"}],\"name\":\"setNewAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"HISTORY_CREATION_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_repoCount\",\"type\":\"uint256\"}],\"name\":\"repoIngots\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"createTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stampCreatedCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_setId\",\"type\":\"uint16\"},{\"name\":\"_membersId\",\"type\":\"uint8[]\"}],\"name\":\"releaseNewStampSet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stampId\",\"type\":\"uint32\"},{\"name\":\"_year\",\"type\":\"uint16\"},{\"name\":\"_setId\",\"type\":\"uint16\"},{\"name\":\"_memberOfSetId\",\"type\":\"uint8\"},{\"name\":\"_totalAmount\",\"type\":\"uint32\"},{\"name\":\"_remainingAmount\",\"type\":\"uint32\"},{\"name\":\"_name\",\"type\":\"bytes32\"},{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_appearance\",\"type\":\"uint8\"}],\"name\":\"releaseNewPromoStampToAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PROMO_CREATION_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"approved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transactionInfo\",\"outputs\":[{\"name\":\"price\",\"type\":\"uint128\"},{\"name\":\"seller\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"repoIngotsInfo\",\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"repoCount\",\"type\":\"uint256\"},{\"name\":\"seller\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"repoCount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"RepoIngotsSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"TransactionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"TransactionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TransactionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"stampId\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"totalAmount\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"remainingAmount\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"year\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"setId\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"memberOfSetId\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"appearance\",\"type\":\"uint8\"}],\"name\":\"CreateNewStamp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// StampCollectionBin is the compiled bytecode used for deploying new contracts.
const StampCollectionBin = `0x6002805460a060020a60ff021916905560c0604052600a60808190527f5374616d70546f6b656e0000000000000000000000000000000000000000000060a09081526200004e919081620000ef565b506040805180820190915260028082527f535400000000000000000000000000000000000000000000000000000000000060209092019182526200009591600b91620000ef565b50348015620000a357600080fd5b506002805460008054600160a060020a03199081163390811790925560a060020a60ff021990921674010000000000000000000000000000000000000000179190911617905562000194565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200013257805160ff191683800117855562000162565b8280016001018555821562000162579182015b828111156200016257825182559160200191906001019062000145565b506200017092915062000174565b5090565b6200019191905b808211156200017057600081556001016200017b565b90565b6122b380620001a46000396000f30060806040526004361061020e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663025afcc7811461021057806302e043da1461027f5780630519ce79146102a657806305e45546146102d757806306fdde03146102ec578063081812fc14610376578063095ea7b31461038e5780630a0f8168146103b257806318160ddd146103c757806323b872dd146103dc57806327d7874c146104065780632af9cff1146104275780632ba73c151461046d5780632f745c591461048e5780633f4ba83a146104b257806342842e0e146104c75780634e0a3379146104f15780634f558e79146105125780634f6ccce71461053e57806357421909146105565780635c975abb1461056b5780635fd8c710146105805780636352211e146105955780636af04a57146105ad57806370a08231146105c257806371587988146105e3578063794d325c146106045780638456cb5914610619578063949d354e1461062e57806395d89b4114610649578063a22cb4651461065e578063a82f8d5814610684578063ac3fc432146106ab578063b047fb50146106c0578063b88d4fde146106d5578063c33bb60a14610744578063c87b56dd1461076c578063cca2704d14610784578063d96a094a146107df578063defb9584146107ea578063e985e9c5146107ff578063f4fae99514610826578063faf20cc514610872575b005b34801561021c57600080fd5b506102286004356108b1565b6040805163ffffffff998a16815261ffff9889166020820152969097168688015260ff94851660608701529287166080860152951660a084015260c08301949094529290921660e083015251908190036101000190f35b34801561028b57600080fd5b50610294610956565b60408051918252519081900360200190f35b3480156102b257600080fd5b506102bb610961565b60408051600160a060020a039092168252519081900360200190f35b3480156102e357600080fd5b50610294610970565b3480156102f857600080fd5b50610301610976565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561033b578181015183820152602001610323565b50505050905090810190601f1680156103685780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561038257600080fd5b506102bb600435610a0d565b34801561039a57600080fd5b5061020e600160a060020a0360043516602435610a28565b3480156103be57600080fd5b506102bb610b0d565b3480156103d357600080fd5b50610294610b1c565b3480156103e857600080fd5b5061020e600160a060020a0360043581169060243516604435610b22565b34801561041257600080fd5b5061020e600160a060020a0360043516610bd1565b34801561043357600080fd5b5061029463ffffffff6004358116906024358116906044351660643561ffff60843581169060a4351660ff60c43581169060e43516610c1f565b34801561047957600080fd5b5061020e600160a060020a0360043516610ccc565b34801561049a57600080fd5b50610294600160a060020a0360043516602435610d1a565b3480156104be57600080fd5b5061020e610d67565b3480156104d357600080fd5b5061020e600160a060020a0360043581169060243516604435610db6565b3480156104fd57600080fd5b5061020e600160a060020a0360043516610dee565b34801561051e57600080fd5b5061052a600435610e3c565b604080519115158252519081900360200190f35b34801561054a57600080fd5b50610294600435610e59565b34801561056257600080fd5b50610294610e8e565b34801561057757600080fd5b5061052a610e95565b34801561058c57600080fd5b5061020e610ea5565b3480156105a157600080fd5b506102bb600435610eff565b3480156105b957600080fd5b506102bb610f29565b3480156105ce57600080fd5b50610294600160a060020a0360043516610f38565b3480156105ef57600080fd5b5061020e600160a060020a0360043516610f6b565b34801561061057600080fd5b50610294610fee565b34801561062557600080fd5b5061020e610ff5565b34801561063a57600080fd5b5061020e600435602435611075565b34801561065557600080fd5b50610301611137565b34801561066a57600080fd5b5061020e600160a060020a03600435166024351515611198565b34801561069057600080fd5b5061020e600435602435600160a060020a036044351661121c565b3480156106b757600080fd5b506102946112a3565b3480156106cc57600080fd5b506102bb6112a9565b3480156106e157600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261020e94600160a060020a0381358116956024803590921695604435953695608494019181908401838280828437509497506112b89650505050505050565b34801561075057600080fd5b5061020e6004803561ffff1690602480359081019101356112f7565b34801561077857600080fd5b50610301600435611447565b34801561079057600080fd5b5061020e63ffffffff60043581169061ffff6024358116916044359091169060ff60643581169160843581169160a4359091169060c43590600160a060020a0360e435169061010435166114fc565b61020e60043561156c565b3480156107f657600080fd5b506102946115b4565b34801561080b57600080fd5b5061052a600160a060020a03600435811690602435166115ba565b34801561083257600080fd5b5061083e6004356115e8565b604080516fffffffffffffffffffffffffffffffff9093168352600160a060020a0390911660208301528051918290030190f35b34801561087e57600080fd5b5061088a600435611621565b604080519384526020840192909252600160a060020a031682820152519081900360600190f35b600080600080600080600080600060038a8154811015156108ce57fe5b600091825260209091206003909102018054600182015460029092015463ffffffff8083169e61ffff640100000000850481169f5066010000000000008504169d5060ff68010000000000000000850481169d506901000000000000000000850483169c506d0100000000000000000000000000909404909116995092975016945092505050565b662386f26fc1000081565b600154600160a060020a031681565b60135481565b600a8054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610a025780601f106109d757610100808354040283529160200191610a02565b820191906000526020600020905b8154815290600101906020018083116109e557829003601f168201915b505050505090505b90565b600090815260076020526040902054600160a060020a031690565b6000610a3382610eff565b9050600160a060020a038381169082161415610a4e57600080fd5b33600160a060020a0382161480610a6a5750610a6a81336115ba565b1515610a7557600080fd5b6000610a8083610a0d565b600160a060020a0316141580610a9e5750600160a060020a03831615155b15610b08576000828152600760209081526040918290208054600160a060020a031916600160a060020a03878116918217909255835186815293519093918516927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35b505050565b600054600160a060020a031681565b600e5490565b80610b2d3382611674565b1515610b3857600080fd5b600160a060020a0384161515610b4d57600080fd5b600160a060020a0383161515610b6257600080fd5b610b6c84836116d3565b610b768483611773565b610b8083836118ac565b82600160a060020a031684600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a350505050565b600054600160a060020a03163314610be857600080fd5b600160a060020a0381161515610bfd57600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b6000806000610c2c6120e7565b600254600160a060020a03163314610c4357600080fd5b6014546244aa2011610c5457600080fd5b610c648c8c8c8c8c8c8c8c6118f5565b9250610c703084611b36565b610c80888c63ffffffff16611b85565b604080518082019091526fffffffffffffffffffffffffffffffff821681523060208201529092509050610cb48382611baa565b50506014805460010190559998505050505050505050565b600054600160a060020a03163314610ce357600080fd5b600160a060020a0381161515610cf857600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b6000610d2583610f38565b8210610d3057600080fd5b600160a060020a0383166000908152600c60205260409020805483908110610d5457fe5b9060005260206000200154905092915050565b600054600160a060020a03163314610d7e57600080fd5b60025460a060020a900460ff161515610d9657600080fd5b601554600160a060020a031615610dac57600080fd5b610db4611c4d565b565b80610dc13382611674565b1515610dcc57600080fd5b610de884848460206040519081016040528060008152506112b8565b50505050565b600054600160a060020a03163314610e0557600080fd5b600160a060020a0381161515610e1a57600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b600090815260066020526040902054600160a060020a0316151590565b6000610e63610b1c565b8210610e6e57600080fd5b600e805483908110610e7c57fe5b90600052602060002001549050919050565b62093a8081565b60025460a060020a900460ff1681565b600154600090600160a060020a03163314610ebf57600080fd5b50600154604051303191600160a060020a03169082156108fc029083906000818181858888f19350505050158015610efb573d6000803e3d6000fd5b5050565b600081815260066020526040812054600160a060020a0316801515610f2357600080fd5b92915050565b601554600160a060020a031681565b6000600160a060020a0382161515610f4f57600080fd5b50600160a060020a031660009081526008602052604090205490565b600054600160a060020a03163314610f8257600080fd5b60025460a060020a900460ff161515610f9a57600080fd5b60158054600160a060020a038316600160a060020a0319909116811790915560408051918252517f450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa4461993059181900360200190a150565b6244aa2081565b600254600160a060020a03163314806110185750600054600160a060020a031633145b8061102d5750600154600160a060020a031633145b151561103857600080fd5b60025460a060020a900460ff161561104f57600080fd5b6002805474ff0000000000000000000000000000000000000000191660a060020a179055565b61107d6120fe565b6110873384611c9c565b151561109257600080fd5b61109d333085610b22565b50604080516060808201835233808352602080840187815284860187815260008981526012845287902086518154600160a060020a031916600160a060020a039091161781559151600183015551600290910155845187815290810186905280850191909152925191927f50cd59a56cf7ed9d5b7a9978578d10cc8529c67bead200fdb943ce7e7c9089da929081900390910190a1505050565b600b8054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610a025780601f106109d757610100808354040283529160200191610a02565b600160a060020a0382163314156111ae57600080fd5b336000818152600960209081526040808320600160a060020a03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b6112246120e7565b6fffffffffffffffffffffffffffffffff8316831461124257600080fd5b3361124c85610eff565b600160a060020a03161461125f57600080fd5b61126a333086610b22565b50604080518082019091526fffffffffffffffffffffffffffffffff83168152600160a060020a0382166020820152610de88482611baa565b60145481565b600254600160a060020a031681565b816112c33382611674565b15156112ce57600080fd5b6112d9858585610b22565b6112e585858585611cbc565b15156112f057600080fd5b5050505050565b6112ff612129565b600254600090600160a060020a0316331461131957600080fd5b60408051908101604052808661ffff16815260200185858080602002602001604051908101604052809392919081815260200183836020028082843750505092909352509193506000925050505b61ffff81168311156113b3578461ffff166005600086868561ffff16818110151561138e57fe5b6020908102929092013560ff1683525081019190915260400160002055600101611367565b600480546001810180835560009290925283517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b6002909202918201805461ffff191661ffff9092169190911781556020808601518051879461143c937f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19c909101920190612141565b505050505050505050565b606061145282610e3c565b151561145d57600080fd5b60008281526010602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156114f05780601f106114c5576101008083540402835291602001916114f0565b820191906000526020600020905b8154815290600101906020018083116114d357829003601f168201915b50505050509050919050565b600254600090600160a060020a0316331461151657600080fd5b5081600160a060020a03811615156115365750600254600160a060020a03165b60135461c3501161154657600080fd5b6115568a8787878d8d8d896118f5565b5050601380546001019055505050505050505050565b6115768134611e45565b306000908152600960209081526040808320338085529252909120805460ff191660011790556115a69082610a28565b6115b1303383610b22565b50565b61c35081565b600160a060020a03918216600090815260096020908152604080832093909416825291909152205460ff1690565b600090815260116020526040902080546001909101546fffffffffffffffffffffffffffffffff90911691600160a060020a0390911690565b600080600061162e6120fe565b5050506000918252506012602090815260409182902082516060810184528154600160a060020a0316808252600183015493820184905260029092015493018390529092565b60008061168083610eff565b905080600160a060020a031684600160a060020a031614806116bb575083600160a060020a03166116b084610a0d565b600160a060020a0316145b806116cb57506116cb81856115ba565b949350505050565b81600160a060020a03166116e682610eff565b600160a060020a0316146116f957600080fd5b600081815260076020526040902054600160a060020a031615610efb5760008181526007602090815260408083208054600160a060020a031916905580518481529051600160a060020a038616927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35050565b60008060006117828585611f17565b6000848152600d6020908152604080832054600160a060020a0389168452600c909252909120549093506117bd90600163ffffffff611fa016565b600160a060020a0386166000908152600c60205260409020805491935090839081106117e557fe5b9060005260206000200154905080600c600087600160a060020a0316600160a060020a031681526020019081526020016000208481548110151561182557fe5b6000918252602080832090910192909255600160a060020a0387168152600c9091526040812080548490811061185757fe5b6000918252602080832090910192909255600160a060020a0387168152600c9091526040902080549061188e9060001983016121e7565b506000938452600d6020526040808520859055908452909220555050565b60006118b88383611fb2565b50600160a060020a039091166000908152600c6020908152604080832080546001810182559084528284208101859055938352600d909152902055565b60006118ff61220b565b610100604051908101604052808b63ffffffff1681526020018761ffff1681526020018661ffff1681526020018560ff1681526020018a63ffffffff1681526020018963ffffffff168152602001886000191681526020018460ff16815250905060016003829080600181540180825580915050906001820390600052602060002090600302016000909192909190915060008201518160000160006101000a81548163ffffffff021916908363ffffffff16021790555060208201518160000160046101000a81548161ffff021916908361ffff16021790555060408201518160000160066101000a81548161ffff021916908361ffff16021790555060608201518160000160086101000a81548160ff021916908360ff16021790555060808201518160000160096101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600d6101000a81548163ffffffff021916908363ffffffff16021790555060c0820151816001019060001916905560e08201518160020160006101000a81548160ff021916908360ff16021790555050500391508163ffffffff1682141515611ab657600080fd5b6040805163ffffffff808d168252808c1660208301528a16818301526060810189905261ffff8089166080830152871660a082015260ff80871660c0830152851660e082015290517f4710117ee2ed24df4322c27506325e38cbbb47a788bc132e0c5d7d107f434256918190036101000190a15098975050505050505050565b611b408282612036565b600e80546000838152600f60205260408120829055600182018355919091527fbb7b4a454dc3493923482f07822329ed19e8244eff582cc204f8554c3620c3fd015550565b600080826107a7850361ffff16026313ab6680811515611ba157fe5b04949350505050565b600082815260116020908152604091829020835181546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff90911690811782558483015160019092018054600160a060020a031916600160a060020a039093169290921790915582518581529182015281517f6110ba4ca3a2a4b05be8a5160cfb9b4dfbdc1738e1c8e424a1546fc3f2bce843929181900390910190a15050565b600054600160a060020a03163314611c6457600080fd5b60025460a060020a900460ff161515611c7c57600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b600090815260066020526040902054600160a060020a0391821691161490565b600080611cd185600160a060020a031661209a565b1515611ce05760019150611e3c565b84600160a060020a031663f0b9e5ba8786866040518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018084600160a060020a0316600160a060020a0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611d78578181015183820152602001611d60565b50505050905090810190601f168015611da55780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b158015611dc657600080fd5b505af1158015611dda573d6000803e3d6000fd5b505050506040513d6020811015611df057600080fd5b50517fffffffff0000000000000000000000000000000000000000000000000000000081167ff0b9e5ba0000000000000000000000000000000000000000000000000000000014925090505b50949350505050565b6000828152601160205260408120600181015481549192600160a060020a03909116916fffffffffffffffffffffffffffffffff1690848214611e8757600080fd5b611e90866120a2565b50604051303190600160a060020a038416906201869f19870180156108fc02916000818181858888f19350505050158015611ecf573d6000803e3d6000fd5b506040805187815260208101879052338183015290517f7b31057475c88e821f05f6e63ac0a08a794cf369937c11da22272b73ef90187f9181900360600190a1505050505050565b81600160a060020a0316611f2a82610eff565b600160a060020a031614611f3d57600080fd5b600160a060020a038216600090815260086020526040902054611f6790600163ffffffff611fa016565b600160a060020a039092166000908152600860209081526040808320949094559181526006909152208054600160a060020a0319169055565b600082821115611fac57fe5b50900390565b600081815260066020526040902054600160a060020a031615611fd457600080fd5b60008181526006602090815260408083208054600160a060020a031916600160a060020a038716908117909155835260089091529020546120169060016120da565b600160a060020a0390921660009081526008602052604090209190915550565b600160a060020a038216151561204b57600080fd5b61205582826118ac565b604080518281529051600160a060020a038416916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b6000903b1190565b600090815260116020526040902080546fffffffffffffffffffffffffffffffff191681556001018054600160a060020a0319169055565b81810182811015610f2357fe5b604080518082019091526000808252602082015290565b6060604051908101604052806000600160a060020a0316815260200160008152602001600081525090565b60408051808201909152600081526060602082015290565b82805482825590600052602060002090601f016020900481019282156121d75791602002820160005b838211156121a857835183826101000a81548160ff021916908360ff160217905550926020019260010160208160000104928301926001030261216a565b80156121d55782816101000a81549060ff02191690556001016020816000010492830192600103026121a8565b505b506121e392915061224f565b5090565b815481835581811115610b0857600083815260209020610b0891810190830161226d565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081019190915290565b610a0a91905b808211156121e357805460ff19168155600101612255565b610a0a91905b808211156121e357600081556001016122735600a165627a7a723058204686f00babc79154e79e6c7cbe2e0a94c01837a45aa4357b75037b40acaf51980029`

// DeployStampCollection deploys a new Ethereum contract, binding an instance of StampCollection to it.
func DeployStampCollection(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StampCollection, error) {
	parsed, err := abi.JSON(strings.NewReader(StampCollectionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StampCollectionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StampCollection{StampCollectionCaller: StampCollectionCaller{contract: contract}, StampCollectionTransactor: StampCollectionTransactor{contract: contract}, StampCollectionFilterer: StampCollectionFilterer{contract: contract}}, nil
}

// StampCollection is an auto generated Go binding around an Ethereum contract.
type StampCollection struct {
	StampCollectionCaller     // Read-only binding to the contract
	StampCollectionTransactor // Write-only binding to the contract
	StampCollectionFilterer   // Log filterer for contract events
}

// StampCollectionCaller is an auto generated read-only Go binding around an Ethereum contract.
type StampCollectionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StampCollectionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StampCollectionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StampCollectionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StampCollectionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StampCollectionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StampCollectionSession struct {
	Contract     *StampCollection  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StampCollectionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StampCollectionCallerSession struct {
	Contract *StampCollectionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// StampCollectionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StampCollectionTransactorSession struct {
	Contract     *StampCollectionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// StampCollectionRaw is an auto generated low-level Go binding around an Ethereum contract.
type StampCollectionRaw struct {
	Contract *StampCollection // Generic contract binding to access the raw methods on
}

// StampCollectionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StampCollectionCallerRaw struct {
	Contract *StampCollectionCaller // Generic read-only contract binding to access the raw methods on
}

// StampCollectionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StampCollectionTransactorRaw struct {
	Contract *StampCollectionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStampCollection creates a new instance of StampCollection, bound to a specific deployed contract.
func NewStampCollection(address common.Address, backend bind.ContractBackend) (*StampCollection, error) {
	contract, err := bindStampCollection(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StampCollection{StampCollectionCaller: StampCollectionCaller{contract: contract}, StampCollectionTransactor: StampCollectionTransactor{contract: contract}, StampCollectionFilterer: StampCollectionFilterer{contract: contract}}, nil
}

// NewStampCollectionCaller creates a new read-only instance of StampCollection, bound to a specific deployed contract.
func NewStampCollectionCaller(address common.Address, caller bind.ContractCaller) (*StampCollectionCaller, error) {
	contract, err := bindStampCollection(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StampCollectionCaller{contract: contract}, nil
}

// NewStampCollectionTransactor creates a new write-only instance of StampCollection, bound to a specific deployed contract.
func NewStampCollectionTransactor(address common.Address, transactor bind.ContractTransactor) (*StampCollectionTransactor, error) {
	contract, err := bindStampCollection(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StampCollectionTransactor{contract: contract}, nil
}

// NewStampCollectionFilterer creates a new log filterer instance of StampCollection, bound to a specific deployed contract.
func NewStampCollectionFilterer(address common.Address, filterer bind.ContractFilterer) (*StampCollectionFilterer, error) {
	contract, err := bindStampCollection(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StampCollectionFilterer{contract: contract}, nil
}

// bindStampCollection binds a generic wrapper to an already deployed contract.
func bindStampCollection(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StampCollectionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StampCollection *StampCollectionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StampCollection.Contract.StampCollectionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StampCollection *StampCollectionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampCollection.Contract.StampCollectionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StampCollection *StampCollectionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StampCollection.Contract.StampCollectionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StampCollection *StampCollectionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StampCollection.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StampCollection *StampCollectionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampCollection.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StampCollection *StampCollectionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StampCollection.Contract.contract.Transact(opts, method, params...)
}

// HISTORYCREATIONLIMIT is a free data retrieval call binding the contract method 0x794d325c.
//
// Solidity: function HISTORY_CREATION_LIMIT() constant returns(uint256)
func (_StampCollection *StampCollectionCaller) HISTORYCREATIONLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "HISTORY_CREATION_LIMIT")
	return *ret0, err
}

// HISTORYCREATIONLIMIT is a free data retrieval call binding the contract method 0x794d325c.
//
// Solidity: function HISTORY_CREATION_LIMIT() constant returns(uint256)
func (_StampCollection *StampCollectionSession) HISTORYCREATIONLIMIT() (*big.Int, error) {
	return _StampCollection.Contract.HISTORYCREATIONLIMIT(&_StampCollection.CallOpts)
}

// HISTORYCREATIONLIMIT is a free data retrieval call binding the contract method 0x794d325c.
//
// Solidity: function HISTORY_CREATION_LIMIT() constant returns(uint256)
func (_StampCollection *StampCollectionCallerSession) HISTORYCREATIONLIMIT() (*big.Int, error) {
	return _StampCollection.Contract.HISTORYCREATIONLIMIT(&_StampCollection.CallOpts)
}

// PROMOCREATIONLIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() constant returns(uint256)
func (_StampCollection *StampCollectionCaller) PROMOCREATIONLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "PROMO_CREATION_LIMIT")
	return *ret0, err
}

// PROMOCREATIONLIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() constant returns(uint256)
func (_StampCollection *StampCollectionSession) PROMOCREATIONLIMIT() (*big.Int, error) {
	return _StampCollection.Contract.PROMOCREATIONLIMIT(&_StampCollection.CallOpts)
}

// PROMOCREATIONLIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() constant returns(uint256)
func (_StampCollection *StampCollectionCallerSession) PROMOCREATIONLIMIT() (*big.Int, error) {
	return _StampCollection.Contract.PROMOCREATIONLIMIT(&_StampCollection.CallOpts)
}

// STAMPAUCTIONDURATION is a free data retrieval call binding the contract method 0x57421909.
//
// Solidity: function STAMP_AUCTION_DURATION() constant returns(uint256)
func (_StampCollection *StampCollectionCaller) STAMPAUCTIONDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "STAMP_AUCTION_DURATION")
	return *ret0, err
}

// STAMPAUCTIONDURATION is a free data retrieval call binding the contract method 0x57421909.
//
// Solidity: function STAMP_AUCTION_DURATION() constant returns(uint256)
func (_StampCollection *StampCollectionSession) STAMPAUCTIONDURATION() (*big.Int, error) {
	return _StampCollection.Contract.STAMPAUCTIONDURATION(&_StampCollection.CallOpts)
}

// STAMPAUCTIONDURATION is a free data retrieval call binding the contract method 0x57421909.
//
// Solidity: function STAMP_AUCTION_DURATION() constant returns(uint256)
func (_StampCollection *StampCollectionCallerSession) STAMPAUCTIONDURATION() (*big.Int, error) {
	return _StampCollection.Contract.STAMPAUCTIONDURATION(&_StampCollection.CallOpts)
}

// STAMPSTARTINGPRICE is a free data retrieval call binding the contract method 0x02e043da.
//
// Solidity: function STAMP_STARTING_PRICE() constant returns(uint256)
func (_StampCollection *StampCollectionCaller) STAMPSTARTINGPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "STAMP_STARTING_PRICE")
	return *ret0, err
}

// STAMPSTARTINGPRICE is a free data retrieval call binding the contract method 0x02e043da.
//
// Solidity: function STAMP_STARTING_PRICE() constant returns(uint256)
func (_StampCollection *StampCollectionSession) STAMPSTARTINGPRICE() (*big.Int, error) {
	return _StampCollection.Contract.STAMPSTARTINGPRICE(&_StampCollection.CallOpts)
}

// STAMPSTARTINGPRICE is a free data retrieval call binding the contract method 0x02e043da.
//
// Solidity: function STAMP_STARTING_PRICE() constant returns(uint256)
func (_StampCollection *StampCollectionCallerSession) STAMPSTARTINGPRICE() (*big.Int, error) {
	return _StampCollection.Contract.STAMPSTARTINGPRICE(&_StampCollection.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StampCollection *StampCollectionCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StampCollection *StampCollectionSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StampCollection.Contract.BalanceOf(&_StampCollection.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StampCollection *StampCollectionCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StampCollection.Contract.BalanceOf(&_StampCollection.CallOpts, _owner)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_StampCollection *StampCollectionCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "ceoAddress")
	return *ret0, err
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_StampCollection *StampCollectionSession) CeoAddress() (common.Address, error) {
	return _StampCollection.Contract.CeoAddress(&_StampCollection.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_StampCollection *StampCollectionCallerSession) CeoAddress() (common.Address, error) {
	return _StampCollection.Contract.CeoAddress(&_StampCollection.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_StampCollection *StampCollectionCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "cfoAddress")
	return *ret0, err
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_StampCollection *StampCollectionSession) CfoAddress() (common.Address, error) {
	return _StampCollection.Contract.CfoAddress(&_StampCollection.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_StampCollection *StampCollectionCallerSession) CfoAddress() (common.Address, error) {
	return _StampCollection.Contract.CfoAddress(&_StampCollection.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_StampCollection *StampCollectionCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "cooAddress")
	return *ret0, err
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_StampCollection *StampCollectionSession) CooAddress() (common.Address, error) {
	return _StampCollection.Contract.CooAddress(&_StampCollection.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_StampCollection *StampCollectionCallerSession) CooAddress() (common.Address, error) {
	return _StampCollection.Contract.CooAddress(&_StampCollection.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_StampCollection *StampCollectionCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_StampCollection *StampCollectionSession) Exists(_tokenId *big.Int) (bool, error) {
	return _StampCollection.Contract.Exists(&_StampCollection.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_StampCollection *StampCollectionCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _StampCollection.Contract.Exists(&_StampCollection.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_StampCollection *StampCollectionCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_StampCollection *StampCollectionSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _StampCollection.Contract.GetApproved(&_StampCollection.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_StampCollection *StampCollectionCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _StampCollection.Contract.GetApproved(&_StampCollection.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(approved bool)
func (_StampCollection *StampCollectionCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(approved bool)
func (_StampCollection *StampCollectionSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _StampCollection.Contract.IsApprovedForAll(&_StampCollection.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(approved bool)
func (_StampCollection *StampCollectionCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _StampCollection.Contract.IsApprovedForAll(&_StampCollection.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StampCollection *StampCollectionCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StampCollection *StampCollectionSession) Name() (string, error) {
	return _StampCollection.Contract.Name(&_StampCollection.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StampCollection *StampCollectionCallerSession) Name() (string, error) {
	return _StampCollection.Contract.Name(&_StampCollection.CallOpts)
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_StampCollection *StampCollectionCaller) NewContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "newContractAddress")
	return *ret0, err
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_StampCollection *StampCollectionSession) NewContractAddress() (common.Address, error) {
	return _StampCollection.Contract.NewContractAddress(&_StampCollection.CallOpts)
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_StampCollection *StampCollectionCallerSession) NewContractAddress() (common.Address, error) {
	return _StampCollection.Contract.NewContractAddress(&_StampCollection.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_StampCollection *StampCollectionCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_StampCollection *StampCollectionSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _StampCollection.Contract.OwnerOf(&_StampCollection.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_StampCollection *StampCollectionCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _StampCollection.Contract.OwnerOf(&_StampCollection.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_StampCollection *StampCollectionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_StampCollection *StampCollectionSession) Paused() (bool, error) {
	return _StampCollection.Contract.Paused(&_StampCollection.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_StampCollection *StampCollectionCallerSession) Paused() (bool, error) {
	return _StampCollection.Contract.Paused(&_StampCollection.CallOpts)
}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() constant returns(uint256)
func (_StampCollection *StampCollectionCaller) PromoCreatedCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "promoCreatedCount")
	return *ret0, err
}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() constant returns(uint256)
func (_StampCollection *StampCollectionSession) PromoCreatedCount() (*big.Int, error) {
	return _StampCollection.Contract.PromoCreatedCount(&_StampCollection.CallOpts)
}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() constant returns(uint256)
func (_StampCollection *StampCollectionCallerSession) PromoCreatedCount() (*big.Int, error) {
	return _StampCollection.Contract.PromoCreatedCount(&_StampCollection.CallOpts)
}

// RepoIngotsInfo is a free data retrieval call binding the contract method 0xfaf20cc5.
//
// Solidity: function repoIngotsInfo(_tokenId uint256) constant returns(tokenId uint256, repoCount uint256, seller address)
func (_StampCollection *StampCollectionCaller) RepoIngotsInfo(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	TokenId   *big.Int
	RepoCount *big.Int
	Seller    common.Address
}, error) {
	ret := new(struct {
		TokenId   *big.Int
		RepoCount *big.Int
		Seller    common.Address
	})
	out := ret
	err := _StampCollection.contract.Call(opts, out, "repoIngotsInfo", _tokenId)
	return *ret, err
}

// RepoIngotsInfo is a free data retrieval call binding the contract method 0xfaf20cc5.
//
// Solidity: function repoIngotsInfo(_tokenId uint256) constant returns(tokenId uint256, repoCount uint256, seller address)
func (_StampCollection *StampCollectionSession) RepoIngotsInfo(_tokenId *big.Int) (struct {
	TokenId   *big.Int
	RepoCount *big.Int
	Seller    common.Address
}, error) {
	return _StampCollection.Contract.RepoIngotsInfo(&_StampCollection.CallOpts, _tokenId)
}

// RepoIngotsInfo is a free data retrieval call binding the contract method 0xfaf20cc5.
//
// Solidity: function repoIngotsInfo(_tokenId uint256) constant returns(tokenId uint256, repoCount uint256, seller address)
func (_StampCollection *StampCollectionCallerSession) RepoIngotsInfo(_tokenId *big.Int) (struct {
	TokenId   *big.Int
	RepoCount *big.Int
	Seller    common.Address
}, error) {
	return _StampCollection.Contract.RepoIngotsInfo(&_StampCollection.CallOpts, _tokenId)
}

// StampCreatedCount is a free data retrieval call binding the contract method 0xac3fc432.
//
// Solidity: function stampCreatedCount() constant returns(uint256)
func (_StampCollection *StampCollectionCaller) StampCreatedCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "stampCreatedCount")
	return *ret0, err
}

// StampCreatedCount is a free data retrieval call binding the contract method 0xac3fc432.
//
// Solidity: function stampCreatedCount() constant returns(uint256)
func (_StampCollection *StampCollectionSession) StampCreatedCount() (*big.Int, error) {
	return _StampCollection.Contract.StampCreatedCount(&_StampCollection.CallOpts)
}

// StampCreatedCount is a free data retrieval call binding the contract method 0xac3fc432.
//
// Solidity: function stampCreatedCount() constant returns(uint256)
func (_StampCollection *StampCollectionCallerSession) StampCreatedCount() (*big.Int, error) {
	return _StampCollection.Contract.StampCreatedCount(&_StampCollection.CallOpts)
}

// StampInfo is a free data retrieval call binding the contract method 0x025afcc7.
//
// Solidity: function stampInfo(_id uint256) constant returns(stampId uint32, year uint16, setId uint16, memberOfSetId uint8, totalAmount uint32, remainingAmount uint32, name bytes32, appearance uint8)
func (_StampCollection *StampCollectionCaller) StampInfo(opts *bind.CallOpts, _id *big.Int) (struct {
	StampId         uint32
	Year            uint16
	SetId           uint16
	MemberOfSetId   uint8
	TotalAmount     uint32
	RemainingAmount uint32
	Name            [32]byte
	Appearance      uint8
}, error) {
	ret := new(struct {
		StampId         uint32
		Year            uint16
		SetId           uint16
		MemberOfSetId   uint8
		TotalAmount     uint32
		RemainingAmount uint32
		Name            [32]byte
		Appearance      uint8
	})
	out := ret
	err := _StampCollection.contract.Call(opts, out, "stampInfo", _id)
	return *ret, err
}

// StampInfo is a free data retrieval call binding the contract method 0x025afcc7.
//
// Solidity: function stampInfo(_id uint256) constant returns(stampId uint32, year uint16, setId uint16, memberOfSetId uint8, totalAmount uint32, remainingAmount uint32, name bytes32, appearance uint8)
func (_StampCollection *StampCollectionSession) StampInfo(_id *big.Int) (struct {
	StampId         uint32
	Year            uint16
	SetId           uint16
	MemberOfSetId   uint8
	TotalAmount     uint32
	RemainingAmount uint32
	Name            [32]byte
	Appearance      uint8
}, error) {
	return _StampCollection.Contract.StampInfo(&_StampCollection.CallOpts, _id)
}

// StampInfo is a free data retrieval call binding the contract method 0x025afcc7.
//
// Solidity: function stampInfo(_id uint256) constant returns(stampId uint32, year uint16, setId uint16, memberOfSetId uint8, totalAmount uint32, remainingAmount uint32, name bytes32, appearance uint8)
func (_StampCollection *StampCollectionCallerSession) StampInfo(_id *big.Int) (struct {
	StampId         uint32
	Year            uint16
	SetId           uint16
	MemberOfSetId   uint8
	TotalAmount     uint32
	RemainingAmount uint32
	Name            [32]byte
	Appearance      uint8
}, error) {
	return _StampCollection.Contract.StampInfo(&_StampCollection.CallOpts, _id)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_StampCollection *StampCollectionCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_StampCollection *StampCollectionSession) Symbol() (string, error) {
	return _StampCollection.Contract.Symbol(&_StampCollection.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_StampCollection *StampCollectionCallerSession) Symbol() (string, error) {
	return _StampCollection.Contract.Symbol(&_StampCollection.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_StampCollection *StampCollectionCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_StampCollection *StampCollectionSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _StampCollection.Contract.TokenByIndex(&_StampCollection.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_StampCollection *StampCollectionCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _StampCollection.Contract.TokenByIndex(&_StampCollection.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_StampCollection *StampCollectionCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_StampCollection *StampCollectionSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _StampCollection.Contract.TokenOfOwnerByIndex(&_StampCollection.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_StampCollection *StampCollectionCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _StampCollection.Contract.TokenOfOwnerByIndex(&_StampCollection.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_StampCollection *StampCollectionCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_StampCollection *StampCollectionSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _StampCollection.Contract.TokenURI(&_StampCollection.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_StampCollection *StampCollectionCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _StampCollection.Contract.TokenURI(&_StampCollection.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StampCollection *StampCollectionCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampCollection.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StampCollection *StampCollectionSession) TotalSupply() (*big.Int, error) {
	return _StampCollection.Contract.TotalSupply(&_StampCollection.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StampCollection *StampCollectionCallerSession) TotalSupply() (*big.Int, error) {
	return _StampCollection.Contract.TotalSupply(&_StampCollection.CallOpts)
}

// TransactionInfo is a free data retrieval call binding the contract method 0xf4fae995.
//
// Solidity: function transactionInfo(_tokenId uint256) constant returns(price uint128, seller address)
func (_StampCollection *StampCollectionCaller) TransactionInfo(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Price  *big.Int
	Seller common.Address
}, error) {
	ret := new(struct {
		Price  *big.Int
		Seller common.Address
	})
	out := ret
	err := _StampCollection.contract.Call(opts, out, "transactionInfo", _tokenId)
	return *ret, err
}

// TransactionInfo is a free data retrieval call binding the contract method 0xf4fae995.
//
// Solidity: function transactionInfo(_tokenId uint256) constant returns(price uint128, seller address)
func (_StampCollection *StampCollectionSession) TransactionInfo(_tokenId *big.Int) (struct {
	Price  *big.Int
	Seller common.Address
}, error) {
	return _StampCollection.Contract.TransactionInfo(&_StampCollection.CallOpts, _tokenId)
}

// TransactionInfo is a free data retrieval call binding the contract method 0xf4fae995.
//
// Solidity: function transactionInfo(_tokenId uint256) constant returns(price uint128, seller address)
func (_StampCollection *StampCollectionCallerSession) TransactionInfo(_tokenId *big.Int) (struct {
	Price  *big.Int
	Seller common.Address
}, error) {
	return _StampCollection.Contract.TransactionInfo(&_StampCollection.CallOpts, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_StampCollection *StampCollectionTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampCollection.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_StampCollection *StampCollectionSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampCollection.Contract.Approve(&_StampCollection.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_StampCollection *StampCollectionTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampCollection.Contract.Approve(&_StampCollection.TransactOpts, _to, _tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(_tokenId uint256) returns()
func (_StampCollection *StampCollectionTransactor) Buy(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampCollection.contract.Transact(opts, "buy", _tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(_tokenId uint256) returns()
func (_StampCollection *StampCollectionSession) Buy(_tokenId *big.Int) (*types.Transaction, error) {
	return _StampCollection.Contract.Buy(&_StampCollection.TransactOpts, _tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(_tokenId uint256) returns()
func (_StampCollection *StampCollectionTransactorSession) Buy(_tokenId *big.Int) (*types.Transaction, error) {
	return _StampCollection.Contract.Buy(&_StampCollection.TransactOpts, _tokenId)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xa82f8d58.
//
// Solidity: function createTransaction(_tokenId uint256, _price uint256, _seller address) returns()
func (_StampCollection *StampCollectionTransactor) CreateTransaction(opts *bind.TransactOpts, _tokenId *big.Int, _price *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _StampCollection.contract.Transact(opts, "createTransaction", _tokenId, _price, _seller)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xa82f8d58.
//
// Solidity: function createTransaction(_tokenId uint256, _price uint256, _seller address) returns()
func (_StampCollection *StampCollectionSession) CreateTransaction(_tokenId *big.Int, _price *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _StampCollection.Contract.CreateTransaction(&_StampCollection.TransactOpts, _tokenId, _price, _seller)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xa82f8d58.
//
// Solidity: function createTransaction(_tokenId uint256, _price uint256, _seller address) returns()
func (_StampCollection *StampCollectionTransactorSession) CreateTransaction(_tokenId *big.Int, _price *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _StampCollection.Contract.CreateTransaction(&_StampCollection.TransactOpts, _tokenId, _price, _seller)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StampCollection *StampCollectionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampCollection.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StampCollection *StampCollectionSession) Pause() (*types.Transaction, error) {
	return _StampCollection.Contract.Pause(&_StampCollection.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StampCollection *StampCollectionTransactorSession) Pause() (*types.Transaction, error) {
	return _StampCollection.Contract.Pause(&_StampCollection.TransactOpts)
}

// ReleaseNewPromoStampToAuction is a paid mutator transaction binding the contract method 0xcca2704d.
//
// Solidity: function releaseNewPromoStampToAuction(_stampId uint32, _year uint16, _setId uint16, _memberOfSetId uint8, _totalAmount uint32, _remainingAmount uint32, _name bytes32, _owner address, _appearance uint8) returns()
func (_StampCollection *StampCollectionTransactor) ReleaseNewPromoStampToAuction(opts *bind.TransactOpts, _stampId uint32, _year uint16, _setId uint16, _memberOfSetId uint8, _totalAmount uint32, _remainingAmount uint32, _name [32]byte, _owner common.Address, _appearance uint8) (*types.Transaction, error) {
	return _StampCollection.contract.Transact(opts, "releaseNewPromoStampToAuction", _stampId, _year, _setId, _memberOfSetId, _totalAmount, _remainingAmount, _name, _owner, _appearance)
}

// ReleaseNewPromoStampToAuction is a paid mutator transaction binding the contract method 0xcca2704d.
//
// Solidity: function releaseNewPromoStampToAuction(_stampId uint32, _year uint16, _setId uint16, _memberOfSetId uint8, _totalAmount uint32, _remainingAmount uint32, _name bytes32, _owner address, _appearance uint8) returns()
func (_StampCollection *StampCollectionSession) ReleaseNewPromoStampToAuction(_stampId uint32, _year uint16, _setId uint16, _memberOfSetId uint8, _totalAmount uint32, _remainingAmount uint32, _name [32]byte, _owner common.Address, _appearance uint8) (*types.Transaction, error) {
	return _StampCollection.Contract.ReleaseNewPromoStampToAuction(&_StampCollection.TransactOpts, _stampId, _year, _setId, _memberOfSetId, _totalAmount, _remainingAmount, _name, _owner, _appearance)
}

// ReleaseNewPromoStampToAuction is a paid mutator transaction binding the contract method 0xcca2704d.
//
// Solidity: function releaseNewPromoStampToAuction(_stampId uint32, _year uint16, _setId uint16, _memberOfSetId uint8, _totalAmount uint32, _remainingAmount uint32, _name bytes32, _owner address, _appearance uint8) returns()
func (_StampCollection *StampCollectionTransactorSession) ReleaseNewPromoStampToAuction(_stampId uint32, _year uint16, _setId uint16, _memberOfSetId uint8, _totalAmount uint32, _remainingAmount uint32, _name [32]byte, _owner common.Address, _appearance uint8) (*types.Transaction, error) {
	return _StampCollection.Contract.ReleaseNewPromoStampToAuction(&_StampCollection.TransactOpts, _stampId, _year, _setId, _memberOfSetId, _totalAmount, _remainingAmount, _name, _owner, _appearance)
}

// ReleaseNewStampSet is a paid mutator transaction binding the contract method 0xc33bb60a.
//
// Solidity: function releaseNewStampSet(_setId uint16, _membersId uint8[]) returns()
func (_StampCollection *StampCollectionTransactor) ReleaseNewStampSet(opts *bind.TransactOpts, _setId uint16, _membersId []uint8) (*types.Transaction, error) {
	return _StampCollection.contract.Transact(opts, "releaseNewStampSet", _setId, _membersId)
}

// ReleaseNewStampSet is a paid mutator transaction binding the contract method 0xc33bb60a.
//
// Solidity: function releaseNewStampSet(_setId uint16, _membersId uint8[]) returns()
func (_StampCollection *StampCollectionSession) ReleaseNewStampSet(_setId uint16, _membersId []uint8) (*types.Transaction, error) {
	return _StampCollection.Contract.ReleaseNewStampSet(&_StampCollection.TransactOpts, _setId, _membersId)
}

// ReleaseNewStampSet is a paid mutator transaction binding the contract method 0xc33bb60a.
//
// Solidity: function releaseNewStampSet(_setId uint16, _membersId uint8[]) returns()
func (_StampCollection *StampCollectionTransactorSession) ReleaseNewStampSet(_setId uint16, _membersId []uint8) (*types.Transaction, error) {
	return _StampCollection.Contract.ReleaseNewStampSet(&_StampCollection.TransactOpts, _setId, _membersId)
}

// ReleaseNewStampToTransaction is a paid mutator transaction binding the contract method 0x2af9cff1.
//
// Solidity: function releaseNewStampToTransaction(_stampId uint32, _totalAmount uint32, _remainingAmount uint32, _name bytes32, _year uint16, _setId uint16, _memberOfSetId uint8, _appearance uint8) returns(uint256)
func (_StampCollection *StampCollectionTransactor) ReleaseNewStampToTransaction(opts *bind.TransactOpts, _stampId uint32, _totalAmount uint32, _remainingAmount uint32, _name [32]byte, _year uint16, _setId uint16, _memberOfSetId uint8, _appearance uint8) (*types.Transaction, error) {
	return _StampCollection.contract.Transact(opts, "releaseNewStampToTransaction", _stampId, _totalAmount, _remainingAmount, _name, _year, _setId, _memberOfSetId, _appearance)
}

// ReleaseNewStampToTransaction is a paid mutator transaction binding the contract method 0x2af9cff1.
//
// Solidity: function releaseNewStampToTransaction(_stampId uint32, _totalAmount uint32, _remainingAmount uint32, _name bytes32, _year uint16, _setId uint16, _memberOfSetId uint8, _appearance uint8) returns(uint256)
func (_StampCollection *StampCollectionSession) ReleaseNewStampToTransaction(_stampId uint32, _totalAmount uint32, _remainingAmount uint32, _name [32]byte, _year uint16, _setId uint16, _memberOfSetId uint8, _appearance uint8) (*types.Transaction, error) {
	return _StampCollection.Contract.ReleaseNewStampToTransaction(&_StampCollection.TransactOpts, _stampId, _totalAmount, _remainingAmount, _name, _year, _setId, _memberOfSetId, _appearance)
}

// ReleaseNewStampToTransaction is a paid mutator transaction binding the contract method 0x2af9cff1.
//
// Solidity: function releaseNewStampToTransaction(_stampId uint32, _totalAmount uint32, _remainingAmount uint32, _name bytes32, _year uint16, _setId uint16, _memberOfSetId uint8, _appearance uint8) returns(uint256)
func (_StampCollection *StampCollectionTransactorSession) ReleaseNewStampToTransaction(_stampId uint32, _totalAmount uint32, _remainingAmount uint32, _name [32]byte, _year uint16, _setId uint16, _memberOfSetId uint8, _appearance uint8) (*types.Transaction, error) {
	return _StampCollection.Contract.ReleaseNewStampToTransaction(&_StampCollection.TransactOpts, _stampId, _totalAmount, _remainingAmount, _name, _year, _setId, _memberOfSetId, _appearance)
}

// RepoIngots is a paid mutator transaction binding the contract method 0x949d354e.
//
// Solidity: function repoIngots(_tokenId uint256, _repoCount uint256) returns()
func (_StampCollection *StampCollectionTransactor) RepoIngots(opts *bind.TransactOpts, _tokenId *big.Int, _repoCount *big.Int) (*types.Transaction, error) {
	return _StampCollection.contract.Transact(opts, "repoIngots", _tokenId, _repoCount)
}

// RepoIngots is a paid mutator transaction binding the contract method 0x949d354e.
//
// Solidity: function repoIngots(_tokenId uint256, _repoCount uint256) returns()
func (_StampCollection *StampCollectionSession) RepoIngots(_tokenId *big.Int, _repoCount *big.Int) (*types.Transaction, error) {
	return _StampCollection.Contract.RepoIngots(&_StampCollection.TransactOpts, _tokenId, _repoCount)
}

// RepoIngots is a paid mutator transaction binding the contract method 0x949d354e.
//
// Solidity: function repoIngots(_tokenId uint256, _repoCount uint256) returns()
func (_StampCollection *StampCollectionTransactorSession) RepoIngots(_tokenId *big.Int, _repoCount *big.Int) (*types.Transaction, error) {
	return _StampCollection.Contract.RepoIngots(&_StampCollection.TransactOpts, _tokenId, _repoCount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_StampCollection *StampCollectionTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _StampCollection.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_StampCollection *StampCollectionSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _StampCollection.Contract.SafeTransferFrom(&_StampCollection.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_StampCollection *StampCollectionTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _StampCollection.Contract.SafeTransferFrom(&_StampCollection.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_StampCollection *StampCollectionTransactor) SetApprovalForAll(opts *bind.TransactOpts, _to common.Address, _approved bool) (*types.Transaction, error) {
	return _StampCollection.contract.Transact(opts, "setApprovalForAll", _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_StampCollection *StampCollectionSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _StampCollection.Contract.SetApprovalForAll(&_StampCollection.TransactOpts, _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_StampCollection *StampCollectionTransactorSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _StampCollection.Contract.SetApprovalForAll(&_StampCollection.TransactOpts, _to, _approved)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_StampCollection *StampCollectionTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _StampCollection.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_StampCollection *StampCollectionSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _StampCollection.Contract.SetCEO(&_StampCollection.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_StampCollection *StampCollectionTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _StampCollection.Contract.SetCEO(&_StampCollection.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_StampCollection *StampCollectionTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _StampCollection.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_StampCollection *StampCollectionSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _StampCollection.Contract.SetCFO(&_StampCollection.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_StampCollection *StampCollectionTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _StampCollection.Contract.SetCFO(&_StampCollection.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_StampCollection *StampCollectionTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _StampCollection.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_StampCollection *StampCollectionSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _StampCollection.Contract.SetCOO(&_StampCollection.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_StampCollection *StampCollectionTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _StampCollection.Contract.SetCOO(&_StampCollection.TransactOpts, _newCOO)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_v2Address address) returns()
func (_StampCollection *StampCollectionTransactor) SetNewAddress(opts *bind.TransactOpts, _v2Address common.Address) (*types.Transaction, error) {
	return _StampCollection.contract.Transact(opts, "setNewAddress", _v2Address)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_v2Address address) returns()
func (_StampCollection *StampCollectionSession) SetNewAddress(_v2Address common.Address) (*types.Transaction, error) {
	return _StampCollection.Contract.SetNewAddress(&_StampCollection.TransactOpts, _v2Address)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_v2Address address) returns()
func (_StampCollection *StampCollectionTransactorSession) SetNewAddress(_v2Address common.Address) (*types.Transaction, error) {
	return _StampCollection.Contract.SetNewAddress(&_StampCollection.TransactOpts, _v2Address)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_StampCollection *StampCollectionTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampCollection.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_StampCollection *StampCollectionSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampCollection.Contract.TransferFrom(&_StampCollection.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_StampCollection *StampCollectionTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampCollection.Contract.TransferFrom(&_StampCollection.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StampCollection *StampCollectionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampCollection.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StampCollection *StampCollectionSession) Unpause() (*types.Transaction, error) {
	return _StampCollection.Contract.Unpause(&_StampCollection.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StampCollection *StampCollectionTransactorSession) Unpause() (*types.Transaction, error) {
	return _StampCollection.Contract.Unpause(&_StampCollection.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_StampCollection *StampCollectionTransactor) WithdrawBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampCollection.contract.Transact(opts, "withdrawBalance")
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_StampCollection *StampCollectionSession) WithdrawBalance() (*types.Transaction, error) {
	return _StampCollection.Contract.WithdrawBalance(&_StampCollection.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_StampCollection *StampCollectionTransactorSession) WithdrawBalance() (*types.Transaction, error) {
	return _StampCollection.Contract.WithdrawBalance(&_StampCollection.TransactOpts)
}

// StampCollectionApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StampCollection contract.
type StampCollectionApprovalIterator struct {
	Event *StampCollectionApproval // Event containing the contract specifics and raw log

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
func (it *StampCollectionApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampCollectionApproval)
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
		it.Event = new(StampCollectionApproval)
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
func (it *StampCollectionApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampCollectionApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampCollectionApproval represents a Approval event raised by the StampCollection contract.
type StampCollectionApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_StampCollection *StampCollectionFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address) (*StampCollectionApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _StampCollection.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return &StampCollectionApprovalIterator{contract: _StampCollection.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_StampCollection *StampCollectionFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StampCollectionApproval, _owner []common.Address, _approved []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _StampCollection.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampCollectionApproval)
				if err := _StampCollection.contract.UnpackLog(event, "Approval", log); err != nil {
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

// StampCollectionApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the StampCollection contract.
type StampCollectionApprovalForAllIterator struct {
	Event *StampCollectionApprovalForAll // Event containing the contract specifics and raw log

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
func (it *StampCollectionApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampCollectionApprovalForAll)
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
		it.Event = new(StampCollectionApprovalForAll)
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
func (it *StampCollectionApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampCollectionApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampCollectionApprovalForAll represents a ApprovalForAll event raised by the StampCollection contract.
type StampCollectionApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_StampCollection *StampCollectionFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*StampCollectionApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _StampCollection.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &StampCollectionApprovalForAllIterator{contract: _StampCollection.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_StampCollection *StampCollectionFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *StampCollectionApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _StampCollection.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampCollectionApprovalForAll)
				if err := _StampCollection.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// StampCollectionContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the StampCollection contract.
type StampCollectionContractUpgradeIterator struct {
	Event *StampCollectionContractUpgrade // Event containing the contract specifics and raw log

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
func (it *StampCollectionContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampCollectionContractUpgrade)
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
		it.Event = new(StampCollectionContractUpgrade)
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
func (it *StampCollectionContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampCollectionContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampCollectionContractUpgrade represents a ContractUpgrade event raised by the StampCollection contract.
type StampCollectionContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_StampCollection *StampCollectionFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*StampCollectionContractUpgradeIterator, error) {

	logs, sub, err := _StampCollection.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &StampCollectionContractUpgradeIterator{contract: _StampCollection.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_StampCollection *StampCollectionFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *StampCollectionContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _StampCollection.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampCollectionContractUpgrade)
				if err := _StampCollection.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// StampCollectionCreateNewStampIterator is returned from FilterCreateNewStamp and is used to iterate over the raw logs and unpacked data for CreateNewStamp events raised by the StampCollection contract.
type StampCollectionCreateNewStampIterator struct {
	Event *StampCollectionCreateNewStamp // Event containing the contract specifics and raw log

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
func (it *StampCollectionCreateNewStampIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampCollectionCreateNewStamp)
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
		it.Event = new(StampCollectionCreateNewStamp)
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
func (it *StampCollectionCreateNewStampIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampCollectionCreateNewStampIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampCollectionCreateNewStamp represents a CreateNewStamp event raised by the StampCollection contract.
type StampCollectionCreateNewStamp struct {
	StampId         uint32
	TotalAmount     uint32
	RemainingAmount uint32
	Name            [32]byte
	Year            uint16
	SetId           uint16
	MemberOfSetId   uint8
	Appearance      uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCreateNewStamp is a free log retrieval operation binding the contract event 0x4710117ee2ed24df4322c27506325e38cbbb47a788bc132e0c5d7d107f434256.
//
// Solidity: e CreateNewStamp(stampId uint32, totalAmount uint32, remainingAmount uint32, name bytes32, year uint16, setId uint16, memberOfSetId uint8, appearance uint8)
func (_StampCollection *StampCollectionFilterer) FilterCreateNewStamp(opts *bind.FilterOpts) (*StampCollectionCreateNewStampIterator, error) {

	logs, sub, err := _StampCollection.contract.FilterLogs(opts, "CreateNewStamp")
	if err != nil {
		return nil, err
	}
	return &StampCollectionCreateNewStampIterator{contract: _StampCollection.contract, event: "CreateNewStamp", logs: logs, sub: sub}, nil
}

// WatchCreateNewStamp is a free log subscription operation binding the contract event 0x4710117ee2ed24df4322c27506325e38cbbb47a788bc132e0c5d7d107f434256.
//
// Solidity: e CreateNewStamp(stampId uint32, totalAmount uint32, remainingAmount uint32, name bytes32, year uint16, setId uint16, memberOfSetId uint8, appearance uint8)
func (_StampCollection *StampCollectionFilterer) WatchCreateNewStamp(opts *bind.WatchOpts, sink chan<- *StampCollectionCreateNewStamp) (event.Subscription, error) {

	logs, sub, err := _StampCollection.contract.WatchLogs(opts, "CreateNewStamp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampCollectionCreateNewStamp)
				if err := _StampCollection.contract.UnpackLog(event, "CreateNewStamp", log); err != nil {
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

// StampCollectionRepoIngotsSuccessfulIterator is returned from FilterRepoIngotsSuccessful and is used to iterate over the raw logs and unpacked data for RepoIngotsSuccessful events raised by the StampCollection contract.
type StampCollectionRepoIngotsSuccessfulIterator struct {
	Event *StampCollectionRepoIngotsSuccessful // Event containing the contract specifics and raw log

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
func (it *StampCollectionRepoIngotsSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampCollectionRepoIngotsSuccessful)
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
		it.Event = new(StampCollectionRepoIngotsSuccessful)
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
func (it *StampCollectionRepoIngotsSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampCollectionRepoIngotsSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampCollectionRepoIngotsSuccessful represents a RepoIngotsSuccessful event raised by the StampCollection contract.
type StampCollectionRepoIngotsSuccessful struct {
	TokenId   *big.Int
	RepoCount *big.Int
	Seller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRepoIngotsSuccessful is a free log retrieval operation binding the contract event 0x50cd59a56cf7ed9d5b7a9978578d10cc8529c67bead200fdb943ce7e7c9089da.
//
// Solidity: e RepoIngotsSuccessful(tokenId uint256, repoCount uint256, seller address)
func (_StampCollection *StampCollectionFilterer) FilterRepoIngotsSuccessful(opts *bind.FilterOpts) (*StampCollectionRepoIngotsSuccessfulIterator, error) {

	logs, sub, err := _StampCollection.contract.FilterLogs(opts, "RepoIngotsSuccessful")
	if err != nil {
		return nil, err
	}
	return &StampCollectionRepoIngotsSuccessfulIterator{contract: _StampCollection.contract, event: "RepoIngotsSuccessful", logs: logs, sub: sub}, nil
}

// WatchRepoIngotsSuccessful is a free log subscription operation binding the contract event 0x50cd59a56cf7ed9d5b7a9978578d10cc8529c67bead200fdb943ce7e7c9089da.
//
// Solidity: e RepoIngotsSuccessful(tokenId uint256, repoCount uint256, seller address)
func (_StampCollection *StampCollectionFilterer) WatchRepoIngotsSuccessful(opts *bind.WatchOpts, sink chan<- *StampCollectionRepoIngotsSuccessful) (event.Subscription, error) {

	logs, sub, err := _StampCollection.contract.WatchLogs(opts, "RepoIngotsSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampCollectionRepoIngotsSuccessful)
				if err := _StampCollection.contract.UnpackLog(event, "RepoIngotsSuccessful", log); err != nil {
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

// StampCollectionTransactionCancelledIterator is returned from FilterTransactionCancelled and is used to iterate over the raw logs and unpacked data for TransactionCancelled events raised by the StampCollection contract.
type StampCollectionTransactionCancelledIterator struct {
	Event *StampCollectionTransactionCancelled // Event containing the contract specifics and raw log

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
func (it *StampCollectionTransactionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampCollectionTransactionCancelled)
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
		it.Event = new(StampCollectionTransactionCancelled)
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
func (it *StampCollectionTransactionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampCollectionTransactionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampCollectionTransactionCancelled represents a TransactionCancelled event raised by the StampCollection contract.
type StampCollectionTransactionCancelled struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransactionCancelled is a free log retrieval operation binding the contract event 0x956fb32199d8b882b2039a14e1be35ab14f7a80b9089fc223f14b43937173e60.
//
// Solidity: e TransactionCancelled(tokenId uint256)
func (_StampCollection *StampCollectionFilterer) FilterTransactionCancelled(opts *bind.FilterOpts) (*StampCollectionTransactionCancelledIterator, error) {

	logs, sub, err := _StampCollection.contract.FilterLogs(opts, "TransactionCancelled")
	if err != nil {
		return nil, err
	}
	return &StampCollectionTransactionCancelledIterator{contract: _StampCollection.contract, event: "TransactionCancelled", logs: logs, sub: sub}, nil
}

// WatchTransactionCancelled is a free log subscription operation binding the contract event 0x956fb32199d8b882b2039a14e1be35ab14f7a80b9089fc223f14b43937173e60.
//
// Solidity: e TransactionCancelled(tokenId uint256)
func (_StampCollection *StampCollectionFilterer) WatchTransactionCancelled(opts *bind.WatchOpts, sink chan<- *StampCollectionTransactionCancelled) (event.Subscription, error) {

	logs, sub, err := _StampCollection.contract.WatchLogs(opts, "TransactionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampCollectionTransactionCancelled)
				if err := _StampCollection.contract.UnpackLog(event, "TransactionCancelled", log); err != nil {
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

// StampCollectionTransactionCreatedIterator is returned from FilterTransactionCreated and is used to iterate over the raw logs and unpacked data for TransactionCreated events raised by the StampCollection contract.
type StampCollectionTransactionCreatedIterator struct {
	Event *StampCollectionTransactionCreated // Event containing the contract specifics and raw log

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
func (it *StampCollectionTransactionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampCollectionTransactionCreated)
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
		it.Event = new(StampCollectionTransactionCreated)
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
func (it *StampCollectionTransactionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampCollectionTransactionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampCollectionTransactionCreated represents a TransactionCreated event raised by the StampCollection contract.
type StampCollectionTransactionCreated struct {
	TokenId *big.Int
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransactionCreated is a free log retrieval operation binding the contract event 0x6110ba4ca3a2a4b05be8a5160cfb9b4dfbdc1738e1c8e424a1546fc3f2bce843.
//
// Solidity: e TransactionCreated(tokenId uint256, price uint256)
func (_StampCollection *StampCollectionFilterer) FilterTransactionCreated(opts *bind.FilterOpts) (*StampCollectionTransactionCreatedIterator, error) {

	logs, sub, err := _StampCollection.contract.FilterLogs(opts, "TransactionCreated")
	if err != nil {
		return nil, err
	}
	return &StampCollectionTransactionCreatedIterator{contract: _StampCollection.contract, event: "TransactionCreated", logs: logs, sub: sub}, nil
}

// WatchTransactionCreated is a free log subscription operation binding the contract event 0x6110ba4ca3a2a4b05be8a5160cfb9b4dfbdc1738e1c8e424a1546fc3f2bce843.
//
// Solidity: e TransactionCreated(tokenId uint256, price uint256)
func (_StampCollection *StampCollectionFilterer) WatchTransactionCreated(opts *bind.WatchOpts, sink chan<- *StampCollectionTransactionCreated) (event.Subscription, error) {

	logs, sub, err := _StampCollection.contract.WatchLogs(opts, "TransactionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampCollectionTransactionCreated)
				if err := _StampCollection.contract.UnpackLog(event, "TransactionCreated", log); err != nil {
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

// StampCollectionTransactionSuccessfulIterator is returned from FilterTransactionSuccessful and is used to iterate over the raw logs and unpacked data for TransactionSuccessful events raised by the StampCollection contract.
type StampCollectionTransactionSuccessfulIterator struct {
	Event *StampCollectionTransactionSuccessful // Event containing the contract specifics and raw log

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
func (it *StampCollectionTransactionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampCollectionTransactionSuccessful)
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
		it.Event = new(StampCollectionTransactionSuccessful)
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
func (it *StampCollectionTransactionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampCollectionTransactionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampCollectionTransactionSuccessful represents a TransactionSuccessful event raised by the StampCollection contract.
type StampCollectionTransactionSuccessful struct {
	TokenId *big.Int
	Price   *big.Int
	Buyer   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransactionSuccessful is a free log retrieval operation binding the contract event 0x7b31057475c88e821f05f6e63ac0a08a794cf369937c11da22272b73ef90187f.
//
// Solidity: e TransactionSuccessful(tokenId uint256, price uint256, buyer address)
func (_StampCollection *StampCollectionFilterer) FilterTransactionSuccessful(opts *bind.FilterOpts) (*StampCollectionTransactionSuccessfulIterator, error) {

	logs, sub, err := _StampCollection.contract.FilterLogs(opts, "TransactionSuccessful")
	if err != nil {
		return nil, err
	}
	return &StampCollectionTransactionSuccessfulIterator{contract: _StampCollection.contract, event: "TransactionSuccessful", logs: logs, sub: sub}, nil
}

// WatchTransactionSuccessful is a free log subscription operation binding the contract event 0x7b31057475c88e821f05f6e63ac0a08a794cf369937c11da22272b73ef90187f.
//
// Solidity: e TransactionSuccessful(tokenId uint256, price uint256, buyer address)
func (_StampCollection *StampCollectionFilterer) WatchTransactionSuccessful(opts *bind.WatchOpts, sink chan<- *StampCollectionTransactionSuccessful) (event.Subscription, error) {

	logs, sub, err := _StampCollection.contract.WatchLogs(opts, "TransactionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampCollectionTransactionSuccessful)
				if err := _StampCollection.contract.UnpackLog(event, "TransactionSuccessful", log); err != nil {
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

// StampCollectionTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StampCollection contract.
type StampCollectionTransferIterator struct {
	Event *StampCollectionTransfer // Event containing the contract specifics and raw log

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
func (it *StampCollectionTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampCollectionTransfer)
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
		it.Event = new(StampCollectionTransfer)
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
func (it *StampCollectionTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampCollectionTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampCollectionTransfer represents a Transfer event raised by the StampCollection contract.
type StampCollectionTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_StampCollection *StampCollectionFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*StampCollectionTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _StampCollection.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &StampCollectionTransferIterator{contract: _StampCollection.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_StampCollection *StampCollectionFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StampCollectionTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _StampCollection.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampCollectionTransfer)
				if err := _StampCollection.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// StampDataSourceABI is the input ABI used to generate the binding from.
const StampDataSourceABI = "[]"

// StampDataSourceBin is the compiled bytecode used for deploying new contracts.
const StampDataSourceBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a72305820d936172599b7e17b30fba4ed6da56246e6588e9e5d547de608e22fe3a4ab1c940029`

// DeployStampDataSource deploys a new Ethereum contract, binding an instance of StampDataSource to it.
func DeployStampDataSource(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StampDataSource, error) {
	parsed, err := abi.JSON(strings.NewReader(StampDataSourceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StampDataSourceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StampDataSource{StampDataSourceCaller: StampDataSourceCaller{contract: contract}, StampDataSourceTransactor: StampDataSourceTransactor{contract: contract}, StampDataSourceFilterer: StampDataSourceFilterer{contract: contract}}, nil
}

// StampDataSource is an auto generated Go binding around an Ethereum contract.
type StampDataSource struct {
	StampDataSourceCaller     // Read-only binding to the contract
	StampDataSourceTransactor // Write-only binding to the contract
	StampDataSourceFilterer   // Log filterer for contract events
}

// StampDataSourceCaller is an auto generated read-only Go binding around an Ethereum contract.
type StampDataSourceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StampDataSourceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StampDataSourceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StampDataSourceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StampDataSourceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StampDataSourceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StampDataSourceSession struct {
	Contract     *StampDataSource  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StampDataSourceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StampDataSourceCallerSession struct {
	Contract *StampDataSourceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// StampDataSourceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StampDataSourceTransactorSession struct {
	Contract     *StampDataSourceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// StampDataSourceRaw is an auto generated low-level Go binding around an Ethereum contract.
type StampDataSourceRaw struct {
	Contract *StampDataSource // Generic contract binding to access the raw methods on
}

// StampDataSourceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StampDataSourceCallerRaw struct {
	Contract *StampDataSourceCaller // Generic read-only contract binding to access the raw methods on
}

// StampDataSourceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StampDataSourceTransactorRaw struct {
	Contract *StampDataSourceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStampDataSource creates a new instance of StampDataSource, bound to a specific deployed contract.
func NewStampDataSource(address common.Address, backend bind.ContractBackend) (*StampDataSource, error) {
	contract, err := bindStampDataSource(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StampDataSource{StampDataSourceCaller: StampDataSourceCaller{contract: contract}, StampDataSourceTransactor: StampDataSourceTransactor{contract: contract}, StampDataSourceFilterer: StampDataSourceFilterer{contract: contract}}, nil
}

// NewStampDataSourceCaller creates a new read-only instance of StampDataSource, bound to a specific deployed contract.
func NewStampDataSourceCaller(address common.Address, caller bind.ContractCaller) (*StampDataSourceCaller, error) {
	contract, err := bindStampDataSource(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StampDataSourceCaller{contract: contract}, nil
}

// NewStampDataSourceTransactor creates a new write-only instance of StampDataSource, bound to a specific deployed contract.
func NewStampDataSourceTransactor(address common.Address, transactor bind.ContractTransactor) (*StampDataSourceTransactor, error) {
	contract, err := bindStampDataSource(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StampDataSourceTransactor{contract: contract}, nil
}

// NewStampDataSourceFilterer creates a new log filterer instance of StampDataSource, bound to a specific deployed contract.
func NewStampDataSourceFilterer(address common.Address, filterer bind.ContractFilterer) (*StampDataSourceFilterer, error) {
	contract, err := bindStampDataSource(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StampDataSourceFilterer{contract: contract}, nil
}

// bindStampDataSource binds a generic wrapper to an already deployed contract.
func bindStampDataSource(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StampDataSourceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StampDataSource *StampDataSourceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StampDataSource.Contract.StampDataSourceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StampDataSource *StampDataSourceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampDataSource.Contract.StampDataSourceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StampDataSource *StampDataSourceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StampDataSource.Contract.StampDataSourceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StampDataSource *StampDataSourceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StampDataSource.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StampDataSource *StampDataSourceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampDataSource.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StampDataSource *StampDataSourceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StampDataSource.Contract.contract.Transact(opts, method, params...)
}

// StampMintingABI is the input ABI used to generate the binding from.
const StampMintingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"STAMP_STARTING_PRICE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"promoCreatedCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stampId\",\"type\":\"uint32\"},{\"name\":\"_totalAmount\",\"type\":\"uint32\"},{\"name\":\"_remainingAmount\",\"type\":\"uint32\"},{\"name\":\"_name\",\"type\":\"bytes32\"},{\"name\":\"_year\",\"type\":\"uint16\"},{\"name\":\"_setId\",\"type\":\"uint16\"},{\"name\":\"_memberOfSetId\",\"type\":\"uint8\"},{\"name\":\"_appearance\",\"type\":\"uint8\"}],\"name\":\"releaseNewStampToTransaction\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STAMP_AUCTION_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"HISTORY_CREATION_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_repoCount\",\"type\":\"uint256\"}],\"name\":\"repoIngots\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"createTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stampCreatedCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_setId\",\"type\":\"uint16\"},{\"name\":\"_membersId\",\"type\":\"uint8[]\"}],\"name\":\"releaseNewStampSet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stampId\",\"type\":\"uint32\"},{\"name\":\"_year\",\"type\":\"uint16\"},{\"name\":\"_setId\",\"type\":\"uint16\"},{\"name\":\"_memberOfSetId\",\"type\":\"uint8\"},{\"name\":\"_totalAmount\",\"type\":\"uint32\"},{\"name\":\"_remainingAmount\",\"type\":\"uint32\"},{\"name\":\"_name\",\"type\":\"bytes32\"},{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_appearance\",\"type\":\"uint8\"}],\"name\":\"releaseNewPromoStampToAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PROMO_CREATION_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"approved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transactionInfo\",\"outputs\":[{\"name\":\"price\",\"type\":\"uint128\"},{\"name\":\"seller\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"repoIngotsInfo\",\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"repoCount\",\"type\":\"uint256\"},{\"name\":\"seller\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"repoCount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"RepoIngotsSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"TransactionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"TransactionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TransactionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"stampId\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"totalAmount\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"remainingAmount\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"year\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"setId\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"memberOfSetId\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"appearance\",\"type\":\"uint8\"}],\"name\":\"CreateNewStamp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// StampMintingBin is the compiled bytecode used for deploying new contracts.
const StampMintingBin = `0x6002805460a060020a60ff021916905560c0604052600a60808190527f5374616d70546f6b656e0000000000000000000000000000000000000000000060a09081526200004e919081620000aa565b506040805180820190915260028082527f535400000000000000000000000000000000000000000000000000000000000060209092019182526200009591600b91620000aa565b50348015620000a357600080fd5b506200014f565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000ed57805160ff19168380011785556200011d565b828001600101855582156200011d579182015b828111156200011d57825182559160200191906001019062000100565b506200012b9291506200012f565b5090565b6200014c91905b808211156200012b576000815560010162000136565b90565b612037806200015f6000396000f3006080604052600436106101e25763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302e043da81146101e75780630519ce791461020e57806305e455461461023f57806306fdde0314610254578063081812fc146102de578063095ea7b3146102f65780630a0f81681461031c57806318160ddd1461033157806323b872dd1461034657806327d7874c146103705780632af9cff1146103915780632ba73c15146103d75780632f745c59146103f85780633f4ba83a1461041c57806342842e0e146104315780634e0a33791461045b5780634f558e791461047c5780634f6ccce7146104a857806357421909146104c05780635c975abb146104d55780636352211e146104ea57806370a0823114610502578063794d325c146105235780638456cb5914610538578063949d354e1461054d57806395d89b4114610568578063a22cb4651461057d578063a82f8d58146105a3578063ac3fc432146105ca578063b047fb50146105df578063b88d4fde146105f4578063c33bb60a14610663578063c87b56dd1461068b578063cca2704d146106a3578063d96a094a146106fe578063defb958414610709578063e985e9c51461071e578063f4fae99514610745578063faf20cc514610791575b600080fd5b3480156101f357600080fd5b506101fc6107d0565b60408051918252519081900360200190f35b34801561021a57600080fd5b506102236107db565b60408051600160a060020a039092168252519081900360200190f35b34801561024b57600080fd5b506101fc6107ea565b34801561026057600080fd5b506102696107f0565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102a357818101518382015260200161028b565b50505050905090810190601f1680156102d05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102ea57600080fd5b50610223600435610887565b34801561030257600080fd5b5061031a600160a060020a03600435166024356108a2565b005b34801561032857600080fd5b50610223610987565b34801561033d57600080fd5b506101fc610996565b34801561035257600080fd5b5061031a600160a060020a036004358116906024351660443561099c565b34801561037c57600080fd5b5061031a600160a060020a0360043516610a4b565b34801561039d57600080fd5b506101fc63ffffffff6004358116906024358116906044351660643561ffff60843581169060a4351660ff60c43581169060e43516610a99565b3480156103e357600080fd5b5061031a600160a060020a0360043516610b46565b34801561040457600080fd5b506101fc600160a060020a0360043516602435610b94565b34801561042857600080fd5b5061031a610be1565b34801561043d57600080fd5b5061031a600160a060020a0360043581169060243516604435610c41565b34801561046757600080fd5b5061031a600160a060020a0360043516610c79565b34801561048857600080fd5b50610494600435610cc7565b604080519115158252519081900360200190f35b3480156104b457600080fd5b506101fc600435610ce4565b3480156104cc57600080fd5b506101fc610d19565b3480156104e157600080fd5b50610494610d20565b3480156104f657600080fd5b50610223600435610d41565b34801561050e57600080fd5b506101fc600160a060020a0360043516610d6b565b34801561052f57600080fd5b506101fc610d9e565b34801561054457600080fd5b5061031a610da5565b34801561055957600080fd5b5061031a600435602435610e47565b34801561057457600080fd5b50610269610f09565b34801561058957600080fd5b5061031a600160a060020a03600435166024351515610f6a565b3480156105af57600080fd5b5061031a600435602435600160a060020a0360443516610fee565b3480156105d657600080fd5b506101fc611075565b3480156105eb57600080fd5b5061022361107b565b34801561060057600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261031a94600160a060020a03813581169560248035909216956044359536956084940191819084018382808284375094975061108a9650505050505050565b34801561066f57600080fd5b5061031a6004803561ffff1690602480359081019101356110c9565b34801561069757600080fd5b50610269600435611219565b3480156106af57600080fd5b5061031a63ffffffff60043581169061ffff6024358116916044359091169060ff60643581169160843581169160a4359091169060c43590600160a060020a0360e435169061010435166112ce565b61031a60043561133e565b34801561071557600080fd5b506101fc611386565b34801561072a57600080fd5b50610494600160a060020a036004358116906024351661138c565b34801561075157600080fd5b5061075d6004356113ba565b604080516fffffffffffffffffffffffffffffffff9093168352600160a060020a0390911660208301528051918290030190f35b34801561079d57600080fd5b506107a96004356113f3565b604080519384526020840192909252600160a060020a031682820152519081900360600190f35b662386f26fc1000081565b600154600160a060020a031681565b60135481565b600a8054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561087c5780601f106108515761010080835404028352916020019161087c565b820191906000526020600020905b81548152906001019060200180831161085f57829003601f168201915b505050505090505b90565b600090815260076020526040902054600160a060020a031690565b60006108ad82610d41565b9050600160a060020a0383811690821614156108c857600080fd5b33600160a060020a03821614806108e457506108e4813361138c565b15156108ef57600080fd5b60006108fa83610887565b600160a060020a03161415806109185750600160a060020a03831615155b15610982576000828152600760209081526040918290208054600160a060020a031916600160a060020a03878116918217909255835186815293519093918516927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35b505050565b600054600160a060020a031681565b600e5490565b806109a73382611446565b15156109b257600080fd5b600160a060020a03841615156109c757600080fd5b600160a060020a03831615156109dc57600080fd5b6109e684836114a5565b6109f08483611546565b6109fa838361167f565b82600160a060020a031684600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a350505050565b600054600160a060020a03163314610a6257600080fd5b600160a060020a0381161515610a7757600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b6000806000610aa6611e6b565b600254600160a060020a03163314610abd57600080fd5b6014546244aa2011610ace57600080fd5b610ade8c8c8c8c8c8c8c8c6116c8565b9250610aea3084611909565b610afa888c63ffffffff16611958565b604080518082019091526fffffffffffffffffffffffffffffffff821681523060208201529092509050610b2e838261197d565b50506014805460010190559998505050505050505050565b600054600160a060020a03163314610b5d57600080fd5b600160a060020a0381161515610b7257600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b6000610b9f83610d6b565b8210610baa57600080fd5b600160a060020a0383166000908152600c60205260409020805483908110610bce57fe5b9060005260206000200154905092915050565b600054600160a060020a03163314610bf857600080fd5b60025474010000000000000000000000000000000000000000900460ff161515610c2157600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b80610c4c3382611446565b1515610c5757600080fd5b610c73848484602060405190810160405280600081525061108a565b50505050565b600054600160a060020a03163314610c9057600080fd5b600160a060020a0381161515610ca557600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b600090815260066020526040902054600160a060020a0316151590565b6000610cee610996565b8210610cf957600080fd5b600e805483908110610d0757fe5b90600052602060002001549050919050565b62093a8081565b60025474010000000000000000000000000000000000000000900460ff1681565b600081815260066020526040812054600160a060020a0316801515610d6557600080fd5b92915050565b6000600160a060020a0382161515610d8257600080fd5b50600160a060020a031660009081526008602052604090205490565b6244aa2081565b600254600160a060020a0316331480610dc85750600054600160a060020a031633145b80610ddd5750600154600160a060020a031633145b1515610de857600080fd5b60025474010000000000000000000000000000000000000000900460ff1615610e1057600080fd5b6002805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b610e4f611e82565b610e593384611a20565b1515610e6457600080fd5b610e6f33308561099c565b50604080516060808201835233808352602080840187815284860187815260008981526012845287902086518154600160a060020a031916600160a060020a039091161781559151600183015551600290910155845187815290810186905280850191909152925191927f50cd59a56cf7ed9d5b7a9978578d10cc8529c67bead200fdb943ce7e7c9089da929081900390910190a1505050565b600b8054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561087c5780601f106108515761010080835404028352916020019161087c565b600160a060020a038216331415610f8057600080fd5b336000818152600960209081526040808320600160a060020a03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b610ff6611e6b565b6fffffffffffffffffffffffffffffffff8316831461101457600080fd5b3361101e85610d41565b600160a060020a03161461103157600080fd5b61103c33308661099c565b50604080518082019091526fffffffffffffffffffffffffffffffff83168152600160a060020a0382166020820152610c73848261197d565b60145481565b600254600160a060020a031681565b816110953382611446565b15156110a057600080fd5b6110ab85858561099c565b6110b785858585611a40565b15156110c257600080fd5b5050505050565b6110d1611ead565b600254600090600160a060020a031633146110eb57600080fd5b60408051908101604052808661ffff16815260200185858080602002602001604051908101604052809392919081815260200183836020028082843750505092909352509193506000925050505b61ffff8116831115611185578461ffff166005600086868561ffff16818110151561116057fe5b6020908102929092013560ff1683525081019190915260400160002055600101611139565b600480546001810180835560009290925283517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b6002909202918201805461ffff191661ffff9092169190911781556020808601518051879461120e937f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19c909101920190611ec5565b505050505050505050565b606061122482610cc7565b151561122f57600080fd5b60008281526010602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156112c25780601f10611297576101008083540402835291602001916112c2565b820191906000526020600020905b8154815290600101906020018083116112a557829003601f168201915b50505050509050919050565b600254600090600160a060020a031633146112e857600080fd5b5081600160a060020a03811615156113085750600254600160a060020a03165b60135461c3501161131857600080fd5b6113288a8787878d8d8d896116c8565b5050601380546001019055505050505050505050565b6113488134611bc9565b306000908152600960209081526040808320338085529252909120805460ff1916600117905561137890826108a2565b61138330338361099c565b50565b61c35081565b600160a060020a03918216600090815260096020908152604080832093909416825291909152205460ff1690565b600090815260116020526040902080546001909101546fffffffffffffffffffffffffffffffff90911691600160a060020a0390911690565b6000806000611400611e82565b5050506000918252506012602090815260409182902082516060810184528154600160a060020a0316808252600183015493820184905260029092015493018390529092565b60008061145283610d41565b905080600160a060020a031684600160a060020a0316148061148d575083600160a060020a031661148284610887565b600160a060020a0316145b8061149d575061149d818561138c565b949350505050565b81600160a060020a03166114b882610d41565b600160a060020a0316146114cb57600080fd5b600081815260076020526040902054600160a060020a0316156115425760008181526007602090815260408083208054600160a060020a031916905580518481529051600160a060020a038616927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35b5050565b60008060006115558585611c9b565b6000848152600d6020908152604080832054600160a060020a0389168452600c9092529091205490935061159090600163ffffffff611d2416565b600160a060020a0386166000908152600c60205260409020805491935090839081106115b857fe5b9060005260206000200154905080600c600087600160a060020a0316600160a060020a03168152602001908152602001600020848154811015156115f857fe5b6000918252602080832090910192909255600160a060020a0387168152600c9091526040812080548490811061162a57fe5b6000918252602080832090910192909255600160a060020a0387168152600c90915260409020805490611661906000198301611f6b565b506000938452600d6020526040808520859055908452909220555050565b600061168b8383611d36565b50600160a060020a039091166000908152600c6020908152604080832080546001810182559084528284208101859055938352600d909152902055565b60006116d2611f8f565b610100604051908101604052808b63ffffffff1681526020018761ffff1681526020018661ffff1681526020018560ff1681526020018a63ffffffff1681526020018963ffffffff168152602001886000191681526020018460ff16815250905060016003829080600181540180825580915050906001820390600052602060002090600302016000909192909190915060008201518160000160006101000a81548163ffffffff021916908363ffffffff16021790555060208201518160000160046101000a81548161ffff021916908361ffff16021790555060408201518160000160066101000a81548161ffff021916908361ffff16021790555060608201518160000160086101000a81548160ff021916908360ff16021790555060808201518160000160096101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600d6101000a81548163ffffffff021916908363ffffffff16021790555060c0820151816001019060001916905560e08201518160020160006101000a81548160ff021916908360ff16021790555050500391508163ffffffff168214151561188957600080fd5b6040805163ffffffff808d168252808c1660208301528a16818301526060810189905261ffff8089166080830152871660a082015260ff80871660c0830152851660e082015290517f4710117ee2ed24df4322c27506325e38cbbb47a788bc132e0c5d7d107f434256918190036101000190a15098975050505050505050565b6119138282611dba565b600e80546000838152600f60205260408120829055600182018355919091527fbb7b4a454dc3493923482f07822329ed19e8244eff582cc204f8554c3620c3fd015550565b600080826107a7850361ffff16026313ab668081151561197457fe5b04949350505050565b600082815260116020908152604091829020835181546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff90911690811782558483015160019092018054600160a060020a031916600160a060020a039093169290921790915582518581529182015281517f6110ba4ca3a2a4b05be8a5160cfb9b4dfbdc1738e1c8e424a1546fc3f2bce843929181900390910190a15050565b600090815260066020526040902054600160a060020a0391821691161490565b600080611a5585600160a060020a0316611e1e565b1515611a645760019150611bc0565b84600160a060020a031663f0b9e5ba8786866040518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018084600160a060020a0316600160a060020a0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611afc578181015183820152602001611ae4565b50505050905090810190601f168015611b295780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b158015611b4a57600080fd5b505af1158015611b5e573d6000803e3d6000fd5b505050506040513d6020811015611b7457600080fd5b50517fffffffff0000000000000000000000000000000000000000000000000000000081167ff0b9e5ba0000000000000000000000000000000000000000000000000000000014925090505b50949350505050565b6000828152601160205260408120600181015481549192600160a060020a03909116916fffffffffffffffffffffffffffffffff1690848214611c0b57600080fd5b611c1486611e26565b50604051303190600160a060020a038416906201869f19870180156108fc02916000818181858888f19350505050158015611c53573d6000803e3d6000fd5b506040805187815260208101879052338183015290517f7b31057475c88e821f05f6e63ac0a08a794cf369937c11da22272b73ef90187f9181900360600190a1505050505050565b81600160a060020a0316611cae82610d41565b600160a060020a031614611cc157600080fd5b600160a060020a038216600090815260086020526040902054611ceb90600163ffffffff611d2416565b600160a060020a039092166000908152600860209081526040808320949094559181526006909152208054600160a060020a0319169055565b600082821115611d3057fe5b50900390565b600081815260066020526040902054600160a060020a031615611d5857600080fd5b60008181526006602090815260408083208054600160a060020a031916600160a060020a03871690811790915583526008909152902054611d9a906001611e5e565b600160a060020a0390921660009081526008602052604090209190915550565b600160a060020a0382161515611dcf57600080fd5b611dd9828261167f565b604080518281529051600160a060020a038416916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b6000903b1190565b600090815260116020526040902080546fffffffffffffffffffffffffffffffff191681556001018054600160a060020a0319169055565b81810182811015610d6557fe5b604080518082019091526000808252602082015290565b6060604051908101604052806000600160a060020a0316815260200160008152602001600081525090565b60408051808201909152600081526060602082015290565b82805482825590600052602060002090601f01602090048101928215611f5b5791602002820160005b83821115611f2c57835183826101000a81548160ff021916908360ff1602179055509260200192600101602081600001049283019260010302611eee565b8015611f595782816101000a81549060ff0219169055600101602081600001049283019260010302611f2c565b505b50611f67929150611fd3565b5090565b81548183558181111561098257600083815260209020610982918101908301611ff1565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081019190915290565b61088491905b80821115611f6757805460ff19168155600101611fd9565b61088491905b80821115611f675760008155600101611ff75600a165627a7a723058202fc27d11e9e9ed50cbb5820c3c4acaba9a437301700fd1ba0b5e2dbd732b33230029`

// DeployStampMinting deploys a new Ethereum contract, binding an instance of StampMinting to it.
func DeployStampMinting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StampMinting, error) {
	parsed, err := abi.JSON(strings.NewReader(StampMintingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StampMintingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StampMinting{StampMintingCaller: StampMintingCaller{contract: contract}, StampMintingTransactor: StampMintingTransactor{contract: contract}, StampMintingFilterer: StampMintingFilterer{contract: contract}}, nil
}

// StampMinting is an auto generated Go binding around an Ethereum contract.
type StampMinting struct {
	StampMintingCaller     // Read-only binding to the contract
	StampMintingTransactor // Write-only binding to the contract
	StampMintingFilterer   // Log filterer for contract events
}

// StampMintingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StampMintingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StampMintingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StampMintingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StampMintingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StampMintingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StampMintingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StampMintingSession struct {
	Contract     *StampMinting     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StampMintingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StampMintingCallerSession struct {
	Contract *StampMintingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StampMintingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StampMintingTransactorSession struct {
	Contract     *StampMintingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StampMintingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StampMintingRaw struct {
	Contract *StampMinting // Generic contract binding to access the raw methods on
}

// StampMintingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StampMintingCallerRaw struct {
	Contract *StampMintingCaller // Generic read-only contract binding to access the raw methods on
}

// StampMintingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StampMintingTransactorRaw struct {
	Contract *StampMintingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStampMinting creates a new instance of StampMinting, bound to a specific deployed contract.
func NewStampMinting(address common.Address, backend bind.ContractBackend) (*StampMinting, error) {
	contract, err := bindStampMinting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StampMinting{StampMintingCaller: StampMintingCaller{contract: contract}, StampMintingTransactor: StampMintingTransactor{contract: contract}, StampMintingFilterer: StampMintingFilterer{contract: contract}}, nil
}

// NewStampMintingCaller creates a new read-only instance of StampMinting, bound to a specific deployed contract.
func NewStampMintingCaller(address common.Address, caller bind.ContractCaller) (*StampMintingCaller, error) {
	contract, err := bindStampMinting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StampMintingCaller{contract: contract}, nil
}

// NewStampMintingTransactor creates a new write-only instance of StampMinting, bound to a specific deployed contract.
func NewStampMintingTransactor(address common.Address, transactor bind.ContractTransactor) (*StampMintingTransactor, error) {
	contract, err := bindStampMinting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StampMintingTransactor{contract: contract}, nil
}

// NewStampMintingFilterer creates a new log filterer instance of StampMinting, bound to a specific deployed contract.
func NewStampMintingFilterer(address common.Address, filterer bind.ContractFilterer) (*StampMintingFilterer, error) {
	contract, err := bindStampMinting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StampMintingFilterer{contract: contract}, nil
}

// bindStampMinting binds a generic wrapper to an already deployed contract.
func bindStampMinting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StampMintingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StampMinting *StampMintingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StampMinting.Contract.StampMintingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StampMinting *StampMintingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampMinting.Contract.StampMintingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StampMinting *StampMintingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StampMinting.Contract.StampMintingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StampMinting *StampMintingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StampMinting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StampMinting *StampMintingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampMinting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StampMinting *StampMintingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StampMinting.Contract.contract.Transact(opts, method, params...)
}

// HISTORYCREATIONLIMIT is a free data retrieval call binding the contract method 0x794d325c.
//
// Solidity: function HISTORY_CREATION_LIMIT() constant returns(uint256)
func (_StampMinting *StampMintingCaller) HISTORYCREATIONLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampMinting.contract.Call(opts, out, "HISTORY_CREATION_LIMIT")
	return *ret0, err
}

// HISTORYCREATIONLIMIT is a free data retrieval call binding the contract method 0x794d325c.
//
// Solidity: function HISTORY_CREATION_LIMIT() constant returns(uint256)
func (_StampMinting *StampMintingSession) HISTORYCREATIONLIMIT() (*big.Int, error) {
	return _StampMinting.Contract.HISTORYCREATIONLIMIT(&_StampMinting.CallOpts)
}

// HISTORYCREATIONLIMIT is a free data retrieval call binding the contract method 0x794d325c.
//
// Solidity: function HISTORY_CREATION_LIMIT() constant returns(uint256)
func (_StampMinting *StampMintingCallerSession) HISTORYCREATIONLIMIT() (*big.Int, error) {
	return _StampMinting.Contract.HISTORYCREATIONLIMIT(&_StampMinting.CallOpts)
}

// PROMOCREATIONLIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() constant returns(uint256)
func (_StampMinting *StampMintingCaller) PROMOCREATIONLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampMinting.contract.Call(opts, out, "PROMO_CREATION_LIMIT")
	return *ret0, err
}

// PROMOCREATIONLIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() constant returns(uint256)
func (_StampMinting *StampMintingSession) PROMOCREATIONLIMIT() (*big.Int, error) {
	return _StampMinting.Contract.PROMOCREATIONLIMIT(&_StampMinting.CallOpts)
}

// PROMOCREATIONLIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() constant returns(uint256)
func (_StampMinting *StampMintingCallerSession) PROMOCREATIONLIMIT() (*big.Int, error) {
	return _StampMinting.Contract.PROMOCREATIONLIMIT(&_StampMinting.CallOpts)
}

// STAMPAUCTIONDURATION is a free data retrieval call binding the contract method 0x57421909.
//
// Solidity: function STAMP_AUCTION_DURATION() constant returns(uint256)
func (_StampMinting *StampMintingCaller) STAMPAUCTIONDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampMinting.contract.Call(opts, out, "STAMP_AUCTION_DURATION")
	return *ret0, err
}

// STAMPAUCTIONDURATION is a free data retrieval call binding the contract method 0x57421909.
//
// Solidity: function STAMP_AUCTION_DURATION() constant returns(uint256)
func (_StampMinting *StampMintingSession) STAMPAUCTIONDURATION() (*big.Int, error) {
	return _StampMinting.Contract.STAMPAUCTIONDURATION(&_StampMinting.CallOpts)
}

// STAMPAUCTIONDURATION is a free data retrieval call binding the contract method 0x57421909.
//
// Solidity: function STAMP_AUCTION_DURATION() constant returns(uint256)
func (_StampMinting *StampMintingCallerSession) STAMPAUCTIONDURATION() (*big.Int, error) {
	return _StampMinting.Contract.STAMPAUCTIONDURATION(&_StampMinting.CallOpts)
}

// STAMPSTARTINGPRICE is a free data retrieval call binding the contract method 0x02e043da.
//
// Solidity: function STAMP_STARTING_PRICE() constant returns(uint256)
func (_StampMinting *StampMintingCaller) STAMPSTARTINGPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampMinting.contract.Call(opts, out, "STAMP_STARTING_PRICE")
	return *ret0, err
}

// STAMPSTARTINGPRICE is a free data retrieval call binding the contract method 0x02e043da.
//
// Solidity: function STAMP_STARTING_PRICE() constant returns(uint256)
func (_StampMinting *StampMintingSession) STAMPSTARTINGPRICE() (*big.Int, error) {
	return _StampMinting.Contract.STAMPSTARTINGPRICE(&_StampMinting.CallOpts)
}

// STAMPSTARTINGPRICE is a free data retrieval call binding the contract method 0x02e043da.
//
// Solidity: function STAMP_STARTING_PRICE() constant returns(uint256)
func (_StampMinting *StampMintingCallerSession) STAMPSTARTINGPRICE() (*big.Int, error) {
	return _StampMinting.Contract.STAMPSTARTINGPRICE(&_StampMinting.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StampMinting *StampMintingCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampMinting.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StampMinting *StampMintingSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StampMinting.Contract.BalanceOf(&_StampMinting.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StampMinting *StampMintingCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StampMinting.Contract.BalanceOf(&_StampMinting.CallOpts, _owner)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_StampMinting *StampMintingCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampMinting.contract.Call(opts, out, "ceoAddress")
	return *ret0, err
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_StampMinting *StampMintingSession) CeoAddress() (common.Address, error) {
	return _StampMinting.Contract.CeoAddress(&_StampMinting.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_StampMinting *StampMintingCallerSession) CeoAddress() (common.Address, error) {
	return _StampMinting.Contract.CeoAddress(&_StampMinting.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_StampMinting *StampMintingCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampMinting.contract.Call(opts, out, "cfoAddress")
	return *ret0, err
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_StampMinting *StampMintingSession) CfoAddress() (common.Address, error) {
	return _StampMinting.Contract.CfoAddress(&_StampMinting.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_StampMinting *StampMintingCallerSession) CfoAddress() (common.Address, error) {
	return _StampMinting.Contract.CfoAddress(&_StampMinting.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_StampMinting *StampMintingCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampMinting.contract.Call(opts, out, "cooAddress")
	return *ret0, err
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_StampMinting *StampMintingSession) CooAddress() (common.Address, error) {
	return _StampMinting.Contract.CooAddress(&_StampMinting.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_StampMinting *StampMintingCallerSession) CooAddress() (common.Address, error) {
	return _StampMinting.Contract.CooAddress(&_StampMinting.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_StampMinting *StampMintingCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StampMinting.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_StampMinting *StampMintingSession) Exists(_tokenId *big.Int) (bool, error) {
	return _StampMinting.Contract.Exists(&_StampMinting.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_StampMinting *StampMintingCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _StampMinting.Contract.Exists(&_StampMinting.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_StampMinting *StampMintingCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampMinting.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_StampMinting *StampMintingSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _StampMinting.Contract.GetApproved(&_StampMinting.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_StampMinting *StampMintingCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _StampMinting.Contract.GetApproved(&_StampMinting.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(approved bool)
func (_StampMinting *StampMintingCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StampMinting.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(approved bool)
func (_StampMinting *StampMintingSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _StampMinting.Contract.IsApprovedForAll(&_StampMinting.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(approved bool)
func (_StampMinting *StampMintingCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _StampMinting.Contract.IsApprovedForAll(&_StampMinting.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StampMinting *StampMintingCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StampMinting.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StampMinting *StampMintingSession) Name() (string, error) {
	return _StampMinting.Contract.Name(&_StampMinting.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StampMinting *StampMintingCallerSession) Name() (string, error) {
	return _StampMinting.Contract.Name(&_StampMinting.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_StampMinting *StampMintingCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampMinting.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_StampMinting *StampMintingSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _StampMinting.Contract.OwnerOf(&_StampMinting.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_StampMinting *StampMintingCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _StampMinting.Contract.OwnerOf(&_StampMinting.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_StampMinting *StampMintingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StampMinting.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_StampMinting *StampMintingSession) Paused() (bool, error) {
	return _StampMinting.Contract.Paused(&_StampMinting.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_StampMinting *StampMintingCallerSession) Paused() (bool, error) {
	return _StampMinting.Contract.Paused(&_StampMinting.CallOpts)
}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() constant returns(uint256)
func (_StampMinting *StampMintingCaller) PromoCreatedCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampMinting.contract.Call(opts, out, "promoCreatedCount")
	return *ret0, err
}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() constant returns(uint256)
func (_StampMinting *StampMintingSession) PromoCreatedCount() (*big.Int, error) {
	return _StampMinting.Contract.PromoCreatedCount(&_StampMinting.CallOpts)
}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() constant returns(uint256)
func (_StampMinting *StampMintingCallerSession) PromoCreatedCount() (*big.Int, error) {
	return _StampMinting.Contract.PromoCreatedCount(&_StampMinting.CallOpts)
}

// RepoIngotsInfo is a free data retrieval call binding the contract method 0xfaf20cc5.
//
// Solidity: function repoIngotsInfo(_tokenId uint256) constant returns(tokenId uint256, repoCount uint256, seller address)
func (_StampMinting *StampMintingCaller) RepoIngotsInfo(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	TokenId   *big.Int
	RepoCount *big.Int
	Seller    common.Address
}, error) {
	ret := new(struct {
		TokenId   *big.Int
		RepoCount *big.Int
		Seller    common.Address
	})
	out := ret
	err := _StampMinting.contract.Call(opts, out, "repoIngotsInfo", _tokenId)
	return *ret, err
}

// RepoIngotsInfo is a free data retrieval call binding the contract method 0xfaf20cc5.
//
// Solidity: function repoIngotsInfo(_tokenId uint256) constant returns(tokenId uint256, repoCount uint256, seller address)
func (_StampMinting *StampMintingSession) RepoIngotsInfo(_tokenId *big.Int) (struct {
	TokenId   *big.Int
	RepoCount *big.Int
	Seller    common.Address
}, error) {
	return _StampMinting.Contract.RepoIngotsInfo(&_StampMinting.CallOpts, _tokenId)
}

// RepoIngotsInfo is a free data retrieval call binding the contract method 0xfaf20cc5.
//
// Solidity: function repoIngotsInfo(_tokenId uint256) constant returns(tokenId uint256, repoCount uint256, seller address)
func (_StampMinting *StampMintingCallerSession) RepoIngotsInfo(_tokenId *big.Int) (struct {
	TokenId   *big.Int
	RepoCount *big.Int
	Seller    common.Address
}, error) {
	return _StampMinting.Contract.RepoIngotsInfo(&_StampMinting.CallOpts, _tokenId)
}

// StampCreatedCount is a free data retrieval call binding the contract method 0xac3fc432.
//
// Solidity: function stampCreatedCount() constant returns(uint256)
func (_StampMinting *StampMintingCaller) StampCreatedCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampMinting.contract.Call(opts, out, "stampCreatedCount")
	return *ret0, err
}

// StampCreatedCount is a free data retrieval call binding the contract method 0xac3fc432.
//
// Solidity: function stampCreatedCount() constant returns(uint256)
func (_StampMinting *StampMintingSession) StampCreatedCount() (*big.Int, error) {
	return _StampMinting.Contract.StampCreatedCount(&_StampMinting.CallOpts)
}

// StampCreatedCount is a free data retrieval call binding the contract method 0xac3fc432.
//
// Solidity: function stampCreatedCount() constant returns(uint256)
func (_StampMinting *StampMintingCallerSession) StampCreatedCount() (*big.Int, error) {
	return _StampMinting.Contract.StampCreatedCount(&_StampMinting.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_StampMinting *StampMintingCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StampMinting.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_StampMinting *StampMintingSession) Symbol() (string, error) {
	return _StampMinting.Contract.Symbol(&_StampMinting.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_StampMinting *StampMintingCallerSession) Symbol() (string, error) {
	return _StampMinting.Contract.Symbol(&_StampMinting.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_StampMinting *StampMintingCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampMinting.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_StampMinting *StampMintingSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _StampMinting.Contract.TokenByIndex(&_StampMinting.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_StampMinting *StampMintingCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _StampMinting.Contract.TokenByIndex(&_StampMinting.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_StampMinting *StampMintingCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampMinting.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_StampMinting *StampMintingSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _StampMinting.Contract.TokenOfOwnerByIndex(&_StampMinting.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_StampMinting *StampMintingCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _StampMinting.Contract.TokenOfOwnerByIndex(&_StampMinting.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_StampMinting *StampMintingCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StampMinting.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_StampMinting *StampMintingSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _StampMinting.Contract.TokenURI(&_StampMinting.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_StampMinting *StampMintingCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _StampMinting.Contract.TokenURI(&_StampMinting.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StampMinting *StampMintingCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampMinting.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StampMinting *StampMintingSession) TotalSupply() (*big.Int, error) {
	return _StampMinting.Contract.TotalSupply(&_StampMinting.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StampMinting *StampMintingCallerSession) TotalSupply() (*big.Int, error) {
	return _StampMinting.Contract.TotalSupply(&_StampMinting.CallOpts)
}

// TransactionInfo is a free data retrieval call binding the contract method 0xf4fae995.
//
// Solidity: function transactionInfo(_tokenId uint256) constant returns(price uint128, seller address)
func (_StampMinting *StampMintingCaller) TransactionInfo(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Price  *big.Int
	Seller common.Address
}, error) {
	ret := new(struct {
		Price  *big.Int
		Seller common.Address
	})
	out := ret
	err := _StampMinting.contract.Call(opts, out, "transactionInfo", _tokenId)
	return *ret, err
}

// TransactionInfo is a free data retrieval call binding the contract method 0xf4fae995.
//
// Solidity: function transactionInfo(_tokenId uint256) constant returns(price uint128, seller address)
func (_StampMinting *StampMintingSession) TransactionInfo(_tokenId *big.Int) (struct {
	Price  *big.Int
	Seller common.Address
}, error) {
	return _StampMinting.Contract.TransactionInfo(&_StampMinting.CallOpts, _tokenId)
}

// TransactionInfo is a free data retrieval call binding the contract method 0xf4fae995.
//
// Solidity: function transactionInfo(_tokenId uint256) constant returns(price uint128, seller address)
func (_StampMinting *StampMintingCallerSession) TransactionInfo(_tokenId *big.Int) (struct {
	Price  *big.Int
	Seller common.Address
}, error) {
	return _StampMinting.Contract.TransactionInfo(&_StampMinting.CallOpts, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_StampMinting *StampMintingTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampMinting.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_StampMinting *StampMintingSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampMinting.Contract.Approve(&_StampMinting.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_StampMinting *StampMintingTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampMinting.Contract.Approve(&_StampMinting.TransactOpts, _to, _tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(_tokenId uint256) returns()
func (_StampMinting *StampMintingTransactor) Buy(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampMinting.contract.Transact(opts, "buy", _tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(_tokenId uint256) returns()
func (_StampMinting *StampMintingSession) Buy(_tokenId *big.Int) (*types.Transaction, error) {
	return _StampMinting.Contract.Buy(&_StampMinting.TransactOpts, _tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(_tokenId uint256) returns()
func (_StampMinting *StampMintingTransactorSession) Buy(_tokenId *big.Int) (*types.Transaction, error) {
	return _StampMinting.Contract.Buy(&_StampMinting.TransactOpts, _tokenId)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xa82f8d58.
//
// Solidity: function createTransaction(_tokenId uint256, _price uint256, _seller address) returns()
func (_StampMinting *StampMintingTransactor) CreateTransaction(opts *bind.TransactOpts, _tokenId *big.Int, _price *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _StampMinting.contract.Transact(opts, "createTransaction", _tokenId, _price, _seller)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xa82f8d58.
//
// Solidity: function createTransaction(_tokenId uint256, _price uint256, _seller address) returns()
func (_StampMinting *StampMintingSession) CreateTransaction(_tokenId *big.Int, _price *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _StampMinting.Contract.CreateTransaction(&_StampMinting.TransactOpts, _tokenId, _price, _seller)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xa82f8d58.
//
// Solidity: function createTransaction(_tokenId uint256, _price uint256, _seller address) returns()
func (_StampMinting *StampMintingTransactorSession) CreateTransaction(_tokenId *big.Int, _price *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _StampMinting.Contract.CreateTransaction(&_StampMinting.TransactOpts, _tokenId, _price, _seller)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StampMinting *StampMintingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampMinting.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StampMinting *StampMintingSession) Pause() (*types.Transaction, error) {
	return _StampMinting.Contract.Pause(&_StampMinting.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StampMinting *StampMintingTransactorSession) Pause() (*types.Transaction, error) {
	return _StampMinting.Contract.Pause(&_StampMinting.TransactOpts)
}

// ReleaseNewPromoStampToAuction is a paid mutator transaction binding the contract method 0xcca2704d.
//
// Solidity: function releaseNewPromoStampToAuction(_stampId uint32, _year uint16, _setId uint16, _memberOfSetId uint8, _totalAmount uint32, _remainingAmount uint32, _name bytes32, _owner address, _appearance uint8) returns()
func (_StampMinting *StampMintingTransactor) ReleaseNewPromoStampToAuction(opts *bind.TransactOpts, _stampId uint32, _year uint16, _setId uint16, _memberOfSetId uint8, _totalAmount uint32, _remainingAmount uint32, _name [32]byte, _owner common.Address, _appearance uint8) (*types.Transaction, error) {
	return _StampMinting.contract.Transact(opts, "releaseNewPromoStampToAuction", _stampId, _year, _setId, _memberOfSetId, _totalAmount, _remainingAmount, _name, _owner, _appearance)
}

// ReleaseNewPromoStampToAuction is a paid mutator transaction binding the contract method 0xcca2704d.
//
// Solidity: function releaseNewPromoStampToAuction(_stampId uint32, _year uint16, _setId uint16, _memberOfSetId uint8, _totalAmount uint32, _remainingAmount uint32, _name bytes32, _owner address, _appearance uint8) returns()
func (_StampMinting *StampMintingSession) ReleaseNewPromoStampToAuction(_stampId uint32, _year uint16, _setId uint16, _memberOfSetId uint8, _totalAmount uint32, _remainingAmount uint32, _name [32]byte, _owner common.Address, _appearance uint8) (*types.Transaction, error) {
	return _StampMinting.Contract.ReleaseNewPromoStampToAuction(&_StampMinting.TransactOpts, _stampId, _year, _setId, _memberOfSetId, _totalAmount, _remainingAmount, _name, _owner, _appearance)
}

// ReleaseNewPromoStampToAuction is a paid mutator transaction binding the contract method 0xcca2704d.
//
// Solidity: function releaseNewPromoStampToAuction(_stampId uint32, _year uint16, _setId uint16, _memberOfSetId uint8, _totalAmount uint32, _remainingAmount uint32, _name bytes32, _owner address, _appearance uint8) returns()
func (_StampMinting *StampMintingTransactorSession) ReleaseNewPromoStampToAuction(_stampId uint32, _year uint16, _setId uint16, _memberOfSetId uint8, _totalAmount uint32, _remainingAmount uint32, _name [32]byte, _owner common.Address, _appearance uint8) (*types.Transaction, error) {
	return _StampMinting.Contract.ReleaseNewPromoStampToAuction(&_StampMinting.TransactOpts, _stampId, _year, _setId, _memberOfSetId, _totalAmount, _remainingAmount, _name, _owner, _appearance)
}

// ReleaseNewStampSet is a paid mutator transaction binding the contract method 0xc33bb60a.
//
// Solidity: function releaseNewStampSet(_setId uint16, _membersId uint8[]) returns()
func (_StampMinting *StampMintingTransactor) ReleaseNewStampSet(opts *bind.TransactOpts, _setId uint16, _membersId []uint8) (*types.Transaction, error) {
	return _StampMinting.contract.Transact(opts, "releaseNewStampSet", _setId, _membersId)
}

// ReleaseNewStampSet is a paid mutator transaction binding the contract method 0xc33bb60a.
//
// Solidity: function releaseNewStampSet(_setId uint16, _membersId uint8[]) returns()
func (_StampMinting *StampMintingSession) ReleaseNewStampSet(_setId uint16, _membersId []uint8) (*types.Transaction, error) {
	return _StampMinting.Contract.ReleaseNewStampSet(&_StampMinting.TransactOpts, _setId, _membersId)
}

// ReleaseNewStampSet is a paid mutator transaction binding the contract method 0xc33bb60a.
//
// Solidity: function releaseNewStampSet(_setId uint16, _membersId uint8[]) returns()
func (_StampMinting *StampMintingTransactorSession) ReleaseNewStampSet(_setId uint16, _membersId []uint8) (*types.Transaction, error) {
	return _StampMinting.Contract.ReleaseNewStampSet(&_StampMinting.TransactOpts, _setId, _membersId)
}

// ReleaseNewStampToTransaction is a paid mutator transaction binding the contract method 0x2af9cff1.
//
// Solidity: function releaseNewStampToTransaction(_stampId uint32, _totalAmount uint32, _remainingAmount uint32, _name bytes32, _year uint16, _setId uint16, _memberOfSetId uint8, _appearance uint8) returns(uint256)
func (_StampMinting *StampMintingTransactor) ReleaseNewStampToTransaction(opts *bind.TransactOpts, _stampId uint32, _totalAmount uint32, _remainingAmount uint32, _name [32]byte, _year uint16, _setId uint16, _memberOfSetId uint8, _appearance uint8) (*types.Transaction, error) {
	return _StampMinting.contract.Transact(opts, "releaseNewStampToTransaction", _stampId, _totalAmount, _remainingAmount, _name, _year, _setId, _memberOfSetId, _appearance)
}

// ReleaseNewStampToTransaction is a paid mutator transaction binding the contract method 0x2af9cff1.
//
// Solidity: function releaseNewStampToTransaction(_stampId uint32, _totalAmount uint32, _remainingAmount uint32, _name bytes32, _year uint16, _setId uint16, _memberOfSetId uint8, _appearance uint8) returns(uint256)
func (_StampMinting *StampMintingSession) ReleaseNewStampToTransaction(_stampId uint32, _totalAmount uint32, _remainingAmount uint32, _name [32]byte, _year uint16, _setId uint16, _memberOfSetId uint8, _appearance uint8) (*types.Transaction, error) {
	return _StampMinting.Contract.ReleaseNewStampToTransaction(&_StampMinting.TransactOpts, _stampId, _totalAmount, _remainingAmount, _name, _year, _setId, _memberOfSetId, _appearance)
}

// ReleaseNewStampToTransaction is a paid mutator transaction binding the contract method 0x2af9cff1.
//
// Solidity: function releaseNewStampToTransaction(_stampId uint32, _totalAmount uint32, _remainingAmount uint32, _name bytes32, _year uint16, _setId uint16, _memberOfSetId uint8, _appearance uint8) returns(uint256)
func (_StampMinting *StampMintingTransactorSession) ReleaseNewStampToTransaction(_stampId uint32, _totalAmount uint32, _remainingAmount uint32, _name [32]byte, _year uint16, _setId uint16, _memberOfSetId uint8, _appearance uint8) (*types.Transaction, error) {
	return _StampMinting.Contract.ReleaseNewStampToTransaction(&_StampMinting.TransactOpts, _stampId, _totalAmount, _remainingAmount, _name, _year, _setId, _memberOfSetId, _appearance)
}

// RepoIngots is a paid mutator transaction binding the contract method 0x949d354e.
//
// Solidity: function repoIngots(_tokenId uint256, _repoCount uint256) returns()
func (_StampMinting *StampMintingTransactor) RepoIngots(opts *bind.TransactOpts, _tokenId *big.Int, _repoCount *big.Int) (*types.Transaction, error) {
	return _StampMinting.contract.Transact(opts, "repoIngots", _tokenId, _repoCount)
}

// RepoIngots is a paid mutator transaction binding the contract method 0x949d354e.
//
// Solidity: function repoIngots(_tokenId uint256, _repoCount uint256) returns()
func (_StampMinting *StampMintingSession) RepoIngots(_tokenId *big.Int, _repoCount *big.Int) (*types.Transaction, error) {
	return _StampMinting.Contract.RepoIngots(&_StampMinting.TransactOpts, _tokenId, _repoCount)
}

// RepoIngots is a paid mutator transaction binding the contract method 0x949d354e.
//
// Solidity: function repoIngots(_tokenId uint256, _repoCount uint256) returns()
func (_StampMinting *StampMintingTransactorSession) RepoIngots(_tokenId *big.Int, _repoCount *big.Int) (*types.Transaction, error) {
	return _StampMinting.Contract.RepoIngots(&_StampMinting.TransactOpts, _tokenId, _repoCount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_StampMinting *StampMintingTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _StampMinting.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_StampMinting *StampMintingSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _StampMinting.Contract.SafeTransferFrom(&_StampMinting.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_StampMinting *StampMintingTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _StampMinting.Contract.SafeTransferFrom(&_StampMinting.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_StampMinting *StampMintingTransactor) SetApprovalForAll(opts *bind.TransactOpts, _to common.Address, _approved bool) (*types.Transaction, error) {
	return _StampMinting.contract.Transact(opts, "setApprovalForAll", _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_StampMinting *StampMintingSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _StampMinting.Contract.SetApprovalForAll(&_StampMinting.TransactOpts, _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_StampMinting *StampMintingTransactorSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _StampMinting.Contract.SetApprovalForAll(&_StampMinting.TransactOpts, _to, _approved)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_StampMinting *StampMintingTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _StampMinting.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_StampMinting *StampMintingSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _StampMinting.Contract.SetCEO(&_StampMinting.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_StampMinting *StampMintingTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _StampMinting.Contract.SetCEO(&_StampMinting.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_StampMinting *StampMintingTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _StampMinting.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_StampMinting *StampMintingSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _StampMinting.Contract.SetCFO(&_StampMinting.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_StampMinting *StampMintingTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _StampMinting.Contract.SetCFO(&_StampMinting.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_StampMinting *StampMintingTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _StampMinting.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_StampMinting *StampMintingSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _StampMinting.Contract.SetCOO(&_StampMinting.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_StampMinting *StampMintingTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _StampMinting.Contract.SetCOO(&_StampMinting.TransactOpts, _newCOO)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_StampMinting *StampMintingTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampMinting.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_StampMinting *StampMintingSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampMinting.Contract.TransferFrom(&_StampMinting.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_StampMinting *StampMintingTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampMinting.Contract.TransferFrom(&_StampMinting.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StampMinting *StampMintingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampMinting.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StampMinting *StampMintingSession) Unpause() (*types.Transaction, error) {
	return _StampMinting.Contract.Unpause(&_StampMinting.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StampMinting *StampMintingTransactorSession) Unpause() (*types.Transaction, error) {
	return _StampMinting.Contract.Unpause(&_StampMinting.TransactOpts)
}

// StampMintingApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StampMinting contract.
type StampMintingApprovalIterator struct {
	Event *StampMintingApproval // Event containing the contract specifics and raw log

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
func (it *StampMintingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampMintingApproval)
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
		it.Event = new(StampMintingApproval)
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
func (it *StampMintingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampMintingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampMintingApproval represents a Approval event raised by the StampMinting contract.
type StampMintingApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_StampMinting *StampMintingFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address) (*StampMintingApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _StampMinting.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return &StampMintingApprovalIterator{contract: _StampMinting.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_StampMinting *StampMintingFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StampMintingApproval, _owner []common.Address, _approved []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _StampMinting.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampMintingApproval)
				if err := _StampMinting.contract.UnpackLog(event, "Approval", log); err != nil {
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

// StampMintingApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the StampMinting contract.
type StampMintingApprovalForAllIterator struct {
	Event *StampMintingApprovalForAll // Event containing the contract specifics and raw log

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
func (it *StampMintingApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampMintingApprovalForAll)
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
		it.Event = new(StampMintingApprovalForAll)
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
func (it *StampMintingApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampMintingApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampMintingApprovalForAll represents a ApprovalForAll event raised by the StampMinting contract.
type StampMintingApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_StampMinting *StampMintingFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*StampMintingApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _StampMinting.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &StampMintingApprovalForAllIterator{contract: _StampMinting.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_StampMinting *StampMintingFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *StampMintingApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _StampMinting.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampMintingApprovalForAll)
				if err := _StampMinting.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// StampMintingContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the StampMinting contract.
type StampMintingContractUpgradeIterator struct {
	Event *StampMintingContractUpgrade // Event containing the contract specifics and raw log

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
func (it *StampMintingContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampMintingContractUpgrade)
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
		it.Event = new(StampMintingContractUpgrade)
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
func (it *StampMintingContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampMintingContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampMintingContractUpgrade represents a ContractUpgrade event raised by the StampMinting contract.
type StampMintingContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_StampMinting *StampMintingFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*StampMintingContractUpgradeIterator, error) {

	logs, sub, err := _StampMinting.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &StampMintingContractUpgradeIterator{contract: _StampMinting.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_StampMinting *StampMintingFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *StampMintingContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _StampMinting.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampMintingContractUpgrade)
				if err := _StampMinting.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// StampMintingCreateNewStampIterator is returned from FilterCreateNewStamp and is used to iterate over the raw logs and unpacked data for CreateNewStamp events raised by the StampMinting contract.
type StampMintingCreateNewStampIterator struct {
	Event *StampMintingCreateNewStamp // Event containing the contract specifics and raw log

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
func (it *StampMintingCreateNewStampIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampMintingCreateNewStamp)
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
		it.Event = new(StampMintingCreateNewStamp)
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
func (it *StampMintingCreateNewStampIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampMintingCreateNewStampIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampMintingCreateNewStamp represents a CreateNewStamp event raised by the StampMinting contract.
type StampMintingCreateNewStamp struct {
	StampId         uint32
	TotalAmount     uint32
	RemainingAmount uint32
	Name            [32]byte
	Year            uint16
	SetId           uint16
	MemberOfSetId   uint8
	Appearance      uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCreateNewStamp is a free log retrieval operation binding the contract event 0x4710117ee2ed24df4322c27506325e38cbbb47a788bc132e0c5d7d107f434256.
//
// Solidity: e CreateNewStamp(stampId uint32, totalAmount uint32, remainingAmount uint32, name bytes32, year uint16, setId uint16, memberOfSetId uint8, appearance uint8)
func (_StampMinting *StampMintingFilterer) FilterCreateNewStamp(opts *bind.FilterOpts) (*StampMintingCreateNewStampIterator, error) {

	logs, sub, err := _StampMinting.contract.FilterLogs(opts, "CreateNewStamp")
	if err != nil {
		return nil, err
	}
	return &StampMintingCreateNewStampIterator{contract: _StampMinting.contract, event: "CreateNewStamp", logs: logs, sub: sub}, nil
}

// WatchCreateNewStamp is a free log subscription operation binding the contract event 0x4710117ee2ed24df4322c27506325e38cbbb47a788bc132e0c5d7d107f434256.
//
// Solidity: e CreateNewStamp(stampId uint32, totalAmount uint32, remainingAmount uint32, name bytes32, year uint16, setId uint16, memberOfSetId uint8, appearance uint8)
func (_StampMinting *StampMintingFilterer) WatchCreateNewStamp(opts *bind.WatchOpts, sink chan<- *StampMintingCreateNewStamp) (event.Subscription, error) {

	logs, sub, err := _StampMinting.contract.WatchLogs(opts, "CreateNewStamp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampMintingCreateNewStamp)
				if err := _StampMinting.contract.UnpackLog(event, "CreateNewStamp", log); err != nil {
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

// StampMintingRepoIngotsSuccessfulIterator is returned from FilterRepoIngotsSuccessful and is used to iterate over the raw logs and unpacked data for RepoIngotsSuccessful events raised by the StampMinting contract.
type StampMintingRepoIngotsSuccessfulIterator struct {
	Event *StampMintingRepoIngotsSuccessful // Event containing the contract specifics and raw log

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
func (it *StampMintingRepoIngotsSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampMintingRepoIngotsSuccessful)
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
		it.Event = new(StampMintingRepoIngotsSuccessful)
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
func (it *StampMintingRepoIngotsSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampMintingRepoIngotsSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampMintingRepoIngotsSuccessful represents a RepoIngotsSuccessful event raised by the StampMinting contract.
type StampMintingRepoIngotsSuccessful struct {
	TokenId   *big.Int
	RepoCount *big.Int
	Seller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRepoIngotsSuccessful is a free log retrieval operation binding the contract event 0x50cd59a56cf7ed9d5b7a9978578d10cc8529c67bead200fdb943ce7e7c9089da.
//
// Solidity: e RepoIngotsSuccessful(tokenId uint256, repoCount uint256, seller address)
func (_StampMinting *StampMintingFilterer) FilterRepoIngotsSuccessful(opts *bind.FilterOpts) (*StampMintingRepoIngotsSuccessfulIterator, error) {

	logs, sub, err := _StampMinting.contract.FilterLogs(opts, "RepoIngotsSuccessful")
	if err != nil {
		return nil, err
	}
	return &StampMintingRepoIngotsSuccessfulIterator{contract: _StampMinting.contract, event: "RepoIngotsSuccessful", logs: logs, sub: sub}, nil
}

// WatchRepoIngotsSuccessful is a free log subscription operation binding the contract event 0x50cd59a56cf7ed9d5b7a9978578d10cc8529c67bead200fdb943ce7e7c9089da.
//
// Solidity: e RepoIngotsSuccessful(tokenId uint256, repoCount uint256, seller address)
func (_StampMinting *StampMintingFilterer) WatchRepoIngotsSuccessful(opts *bind.WatchOpts, sink chan<- *StampMintingRepoIngotsSuccessful) (event.Subscription, error) {

	logs, sub, err := _StampMinting.contract.WatchLogs(opts, "RepoIngotsSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampMintingRepoIngotsSuccessful)
				if err := _StampMinting.contract.UnpackLog(event, "RepoIngotsSuccessful", log); err != nil {
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

// StampMintingTransactionCancelledIterator is returned from FilterTransactionCancelled and is used to iterate over the raw logs and unpacked data for TransactionCancelled events raised by the StampMinting contract.
type StampMintingTransactionCancelledIterator struct {
	Event *StampMintingTransactionCancelled // Event containing the contract specifics and raw log

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
func (it *StampMintingTransactionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampMintingTransactionCancelled)
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
		it.Event = new(StampMintingTransactionCancelled)
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
func (it *StampMintingTransactionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampMintingTransactionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampMintingTransactionCancelled represents a TransactionCancelled event raised by the StampMinting contract.
type StampMintingTransactionCancelled struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransactionCancelled is a free log retrieval operation binding the contract event 0x956fb32199d8b882b2039a14e1be35ab14f7a80b9089fc223f14b43937173e60.
//
// Solidity: e TransactionCancelled(tokenId uint256)
func (_StampMinting *StampMintingFilterer) FilterTransactionCancelled(opts *bind.FilterOpts) (*StampMintingTransactionCancelledIterator, error) {

	logs, sub, err := _StampMinting.contract.FilterLogs(opts, "TransactionCancelled")
	if err != nil {
		return nil, err
	}
	return &StampMintingTransactionCancelledIterator{contract: _StampMinting.contract, event: "TransactionCancelled", logs: logs, sub: sub}, nil
}

// WatchTransactionCancelled is a free log subscription operation binding the contract event 0x956fb32199d8b882b2039a14e1be35ab14f7a80b9089fc223f14b43937173e60.
//
// Solidity: e TransactionCancelled(tokenId uint256)
func (_StampMinting *StampMintingFilterer) WatchTransactionCancelled(opts *bind.WatchOpts, sink chan<- *StampMintingTransactionCancelled) (event.Subscription, error) {

	logs, sub, err := _StampMinting.contract.WatchLogs(opts, "TransactionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampMintingTransactionCancelled)
				if err := _StampMinting.contract.UnpackLog(event, "TransactionCancelled", log); err != nil {
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

// StampMintingTransactionCreatedIterator is returned from FilterTransactionCreated and is used to iterate over the raw logs and unpacked data for TransactionCreated events raised by the StampMinting contract.
type StampMintingTransactionCreatedIterator struct {
	Event *StampMintingTransactionCreated // Event containing the contract specifics and raw log

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
func (it *StampMintingTransactionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampMintingTransactionCreated)
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
		it.Event = new(StampMintingTransactionCreated)
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
func (it *StampMintingTransactionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampMintingTransactionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampMintingTransactionCreated represents a TransactionCreated event raised by the StampMinting contract.
type StampMintingTransactionCreated struct {
	TokenId *big.Int
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransactionCreated is a free log retrieval operation binding the contract event 0x6110ba4ca3a2a4b05be8a5160cfb9b4dfbdc1738e1c8e424a1546fc3f2bce843.
//
// Solidity: e TransactionCreated(tokenId uint256, price uint256)
func (_StampMinting *StampMintingFilterer) FilterTransactionCreated(opts *bind.FilterOpts) (*StampMintingTransactionCreatedIterator, error) {

	logs, sub, err := _StampMinting.contract.FilterLogs(opts, "TransactionCreated")
	if err != nil {
		return nil, err
	}
	return &StampMintingTransactionCreatedIterator{contract: _StampMinting.contract, event: "TransactionCreated", logs: logs, sub: sub}, nil
}

// WatchTransactionCreated is a free log subscription operation binding the contract event 0x6110ba4ca3a2a4b05be8a5160cfb9b4dfbdc1738e1c8e424a1546fc3f2bce843.
//
// Solidity: e TransactionCreated(tokenId uint256, price uint256)
func (_StampMinting *StampMintingFilterer) WatchTransactionCreated(opts *bind.WatchOpts, sink chan<- *StampMintingTransactionCreated) (event.Subscription, error) {

	logs, sub, err := _StampMinting.contract.WatchLogs(opts, "TransactionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampMintingTransactionCreated)
				if err := _StampMinting.contract.UnpackLog(event, "TransactionCreated", log); err != nil {
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

// StampMintingTransactionSuccessfulIterator is returned from FilterTransactionSuccessful and is used to iterate over the raw logs and unpacked data for TransactionSuccessful events raised by the StampMinting contract.
type StampMintingTransactionSuccessfulIterator struct {
	Event *StampMintingTransactionSuccessful // Event containing the contract specifics and raw log

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
func (it *StampMintingTransactionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampMintingTransactionSuccessful)
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
		it.Event = new(StampMintingTransactionSuccessful)
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
func (it *StampMintingTransactionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampMintingTransactionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampMintingTransactionSuccessful represents a TransactionSuccessful event raised by the StampMinting contract.
type StampMintingTransactionSuccessful struct {
	TokenId *big.Int
	Price   *big.Int
	Buyer   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransactionSuccessful is a free log retrieval operation binding the contract event 0x7b31057475c88e821f05f6e63ac0a08a794cf369937c11da22272b73ef90187f.
//
// Solidity: e TransactionSuccessful(tokenId uint256, price uint256, buyer address)
func (_StampMinting *StampMintingFilterer) FilterTransactionSuccessful(opts *bind.FilterOpts) (*StampMintingTransactionSuccessfulIterator, error) {

	logs, sub, err := _StampMinting.contract.FilterLogs(opts, "TransactionSuccessful")
	if err != nil {
		return nil, err
	}
	return &StampMintingTransactionSuccessfulIterator{contract: _StampMinting.contract, event: "TransactionSuccessful", logs: logs, sub: sub}, nil
}

// WatchTransactionSuccessful is a free log subscription operation binding the contract event 0x7b31057475c88e821f05f6e63ac0a08a794cf369937c11da22272b73ef90187f.
//
// Solidity: e TransactionSuccessful(tokenId uint256, price uint256, buyer address)
func (_StampMinting *StampMintingFilterer) WatchTransactionSuccessful(opts *bind.WatchOpts, sink chan<- *StampMintingTransactionSuccessful) (event.Subscription, error) {

	logs, sub, err := _StampMinting.contract.WatchLogs(opts, "TransactionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampMintingTransactionSuccessful)
				if err := _StampMinting.contract.UnpackLog(event, "TransactionSuccessful", log); err != nil {
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

// StampMintingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StampMinting contract.
type StampMintingTransferIterator struct {
	Event *StampMintingTransfer // Event containing the contract specifics and raw log

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
func (it *StampMintingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampMintingTransfer)
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
		it.Event = new(StampMintingTransfer)
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
func (it *StampMintingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampMintingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampMintingTransfer represents a Transfer event raised by the StampMinting contract.
type StampMintingTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_StampMinting *StampMintingFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*StampMintingTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _StampMinting.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &StampMintingTransferIterator{contract: _StampMinting.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_StampMinting *StampMintingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StampMintingTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _StampMinting.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampMintingTransfer)
				if err := _StampMinting.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// StampTransactionABI is the input ABI used to generate the binding from.
const StampTransactionABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"createTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"approved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transactionInfo\",\"outputs\":[{\"name\":\"price\",\"type\":\"uint128\"},{\"name\":\"seller\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"TransactionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"TransactionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TransactionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"stampId\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"totalAmount\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"remainingAmount\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"year\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"setId\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"memberOfSetId\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"appearance\",\"type\":\"uint8\"}],\"name\":\"CreateNewStamp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// StampTransactionBin is the compiled bytecode used for deploying new contracts.
const StampTransactionBin = `0x6002805460a060020a60ff021916905560c0604052600a60808190527f5374616d70546f6b656e0000000000000000000000000000000000000000000060a09081526200004e919081620000aa565b506040805180820190915260028082527f535400000000000000000000000000000000000000000000000000000000000060209092019182526200009591600b91620000aa565b50348015620000a357600080fd5b506200014f565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000ed57805160ff19168380011785556200011d565b828001600101855582156200011d579182015b828111156200011d57825182559160200191906001019062000100565b506200012b9291506200012f565b5090565b6200014c91905b808211156200012b576000815560010162000136565b90565b6115eb806200015f6000396000f3006080604052600436106101695763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630519ce79811461016e57806306fdde031461019f578063081812fc14610229578063095ea7b3146102415780630a0f81681461026757806318160ddd1461027c57806323b872dd146102a357806327d7874c146102cd5780632ba73c15146102ee5780632f745c591461030f5780633f4ba83a1461033357806342842e0e146103485780634e0a3379146103725780634f558e79146103935780634f6ccce7146103bf5780635c975abb146103d75780636352211e146103ec57806370a08231146104045780638456cb591461042557806395d89b411461043a578063a22cb4651461044f578063a82f8d5814610475578063b047fb501461049c578063b88d4fde146104b1578063c87b56dd14610520578063d96a094a14610538578063e985e9c514610543578063f4fae9951461056a575b600080fd5b34801561017a57600080fd5b506101836105b6565b60408051600160a060020a039092168252519081900360200190f35b3480156101ab57600080fd5b506101b46105c5565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101ee5781810151838201526020016101d6565b50505050905090810190601f16801561021b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561023557600080fd5b5061018360043561065c565b34801561024d57600080fd5b50610265600160a060020a0360043516602435610677565b005b34801561027357600080fd5b5061018361075c565b34801561028857600080fd5b5061029161076b565b60408051918252519081900360200190f35b3480156102af57600080fd5b50610265600160a060020a0360043581169060243516604435610771565b3480156102d957600080fd5b50610265600160a060020a0360043516610820565b3480156102fa57600080fd5b50610265600160a060020a036004351661086e565b34801561031b57600080fd5b50610291600160a060020a03600435166024356108bc565b34801561033f57600080fd5b50610265610909565b34801561035457600080fd5b50610265600160a060020a0360043581169060243516604435610969565b34801561037e57600080fd5b50610265600160a060020a03600435166109a1565b34801561039f57600080fd5b506103ab6004356109ef565b604080519115158252519081900360200190f35b3480156103cb57600080fd5b50610291600435610a0c565b3480156103e357600080fd5b506103ab610a41565b3480156103f857600080fd5b50610183600435610a62565b34801561041057600080fd5b50610291600160a060020a0360043516610a8c565b34801561043157600080fd5b50610265610abf565b34801561044657600080fd5b506101b4610b61565b34801561045b57600080fd5b50610265600160a060020a03600435166024351515610bc2565b34801561048157600080fd5b50610265600435602435600160a060020a0360443516610c46565b3480156104a857600080fd5b50610183610ccd565b3480156104bd57600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261026594600160a060020a038135811695602480359092169560443595369560849401918190840183828082843750949750610cdc9650505050505050565b34801561052c57600080fd5b506101b4600435610d1b565b610265600435610dd0565b34801561054f57600080fd5b506103ab600160a060020a0360043581169060243516610e18565b34801561057657600080fd5b50610582600435610e46565b604080516fffffffffffffffffffffffffffffffff9093168352600160a060020a0390911660208301528051918290030190f35b600154600160a060020a031681565b600a8054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156106515780601f1061062657610100808354040283529160200191610651565b820191906000526020600020905b81548152906001019060200180831161063457829003601f168201915b505050505090505b90565b600090815260076020526040902054600160a060020a031690565b600061068282610a62565b9050600160a060020a03838116908216141561069d57600080fd5b33600160a060020a03821614806106b957506106b98133610e18565b15156106c457600080fd5b60006106cf8361065c565b600160a060020a03161415806106ed5750600160a060020a03831615155b15610757576000828152600760209081526040918290208054600160a060020a031916600160a060020a03878116918217909255835186815293519093918516927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35b505050565b600054600160a060020a031681565b600e5490565b8061077c3382610e7f565b151561078757600080fd5b600160a060020a038416151561079c57600080fd5b600160a060020a03831615156107b157600080fd5b6107bb8483610ede565b6107c58483610f7f565b6107cf83836110b8565b82600160a060020a031684600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a350505050565b600054600160a060020a0316331461083757600080fd5b600160a060020a038116151561084c57600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a0316331461088557600080fd5b600160a060020a038116151561089a57600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b60006108c783610a8c565b82106108d257600080fd5b600160a060020a0383166000908152600c602052604090208054839081106108f657fe5b9060005260206000200154905092915050565b600054600160a060020a0316331461092057600080fd5b60025474010000000000000000000000000000000000000000900460ff16151561094957600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b806109743382610e7f565b151561097f57600080fd5b61099b8484846020604051908101604052806000815250610cdc565b50505050565b600054600160a060020a031633146109b857600080fd5b600160a060020a03811615156109cd57600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b600090815260066020526040902054600160a060020a0316151590565b6000610a1661076b565b8210610a2157600080fd5b600e805483908110610a2f57fe5b90600052602060002001549050919050565b60025474010000000000000000000000000000000000000000900460ff1681565b600081815260066020526040812054600160a060020a0316801515610a8657600080fd5b92915050565b6000600160a060020a0382161515610aa357600080fd5b50600160a060020a031660009081526008602052604090205490565b600254600160a060020a0316331480610ae25750600054600160a060020a031633145b80610af75750600154600160a060020a031633145b1515610b0257600080fd5b60025474010000000000000000000000000000000000000000900460ff1615610b2a57600080fd5b6002805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b600b8054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156106515780601f1061062657610100808354040283529160200191610651565b600160a060020a038216331415610bd857600080fd5b336000818152600960209081526040808320600160a060020a03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b610c4e61156b565b6fffffffffffffffffffffffffffffffff83168314610c6c57600080fd5b33610c7685610a62565b600160a060020a031614610c8957600080fd5b610c94333086610771565b50604080518082019091526fffffffffffffffffffffffffffffffff83168152600160a060020a038216602082015261099b8482611101565b600254600160a060020a031681565b81610ce73382610e7f565b1515610cf257600080fd5b610cfd858585610771565b610d09858585856111a4565b1515610d1457600080fd5b5050505050565b6060610d26826109ef565b1515610d3157600080fd5b60008281526010602090815260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015610dc45780601f10610d9957610100808354040283529160200191610dc4565b820191906000526020600020905b815481529060010190602001808311610da757829003601f168201915b50505050509050919050565b610dda813461132d565b306000908152600960209081526040808320338085529252909120805460ff19166001179055610e0a9082610677565b610e15303383610771565b50565b600160a060020a03918216600090815260096020908152604080832093909416825291909152205460ff1690565b600090815260116020526040902080546001909101546fffffffffffffffffffffffffffffffff90911691600160a060020a0390911690565b600080610e8b83610a62565b905080600160a060020a031684600160a060020a03161480610ec6575083600160a060020a0316610ebb8461065c565b600160a060020a0316145b80610ed65750610ed68185610e18565b949350505050565b81600160a060020a0316610ef182610a62565b600160a060020a031614610f0457600080fd5b600081815260076020526040902054600160a060020a031615610f7b5760008181526007602090815260408083208054600160a060020a031916905580518481529051600160a060020a038616927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35b5050565b6000806000610f8e85856113ff565b6000848152600d6020908152604080832054600160a060020a0389168452600c90925290912054909350610fc990600163ffffffff61148816565b600160a060020a0386166000908152600c6020526040902080549193509083908110610ff157fe5b9060005260206000200154905080600c600087600160a060020a0316600160a060020a031681526020019081526020016000208481548110151561103157fe5b6000918252602080832090910192909255600160a060020a0387168152600c9091526040812080548490811061106357fe5b6000918252602080832090910192909255600160a060020a0387168152600c9091526040902080549061109a906000198301611582565b506000938452600d6020526040808520859055908452909220555050565b60006110c4838361149a565b50600160a060020a039091166000908152600c6020908152604080832080546001810182559084528284208101859055938352600d909152902055565b600082815260116020908152604091829020835181546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff90911690811782558483015160019092018054600160a060020a031916600160a060020a039093169290921790915582518581529182015281517f6110ba4ca3a2a4b05be8a5160cfb9b4dfbdc1738e1c8e424a1546fc3f2bce843929181900390910190a15050565b6000806111b985600160a060020a031661151e565b15156111c85760019150611324565b84600160a060020a031663f0b9e5ba8786866040518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018084600160a060020a0316600160a060020a0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611260578181015183820152602001611248565b50505050905090810190601f16801561128d5780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b1580156112ae57600080fd5b505af11580156112c2573d6000803e3d6000fd5b505050506040513d60208110156112d857600080fd5b50517fffffffff0000000000000000000000000000000000000000000000000000000081167ff0b9e5ba0000000000000000000000000000000000000000000000000000000014925090505b50949350505050565b6000828152601160205260408120600181015481549192600160a060020a03909116916fffffffffffffffffffffffffffffffff169084821461136f57600080fd5b61137886611526565b50604051303190600160a060020a038416906201869f19870180156108fc02916000818181858888f193505050501580156113b7573d6000803e3d6000fd5b506040805187815260208101879052338183015290517f7b31057475c88e821f05f6e63ac0a08a794cf369937c11da22272b73ef90187f9181900360600190a1505050505050565b81600160a060020a031661141282610a62565b600160a060020a03161461142557600080fd5b600160a060020a03821660009081526008602052604090205461144f90600163ffffffff61148816565b600160a060020a039092166000908152600860209081526040808320949094559181526006909152208054600160a060020a0319169055565b60008282111561149457fe5b50900390565b600081815260066020526040902054600160a060020a0316156114bc57600080fd5b60008181526006602090815260408083208054600160a060020a031916600160a060020a038716908117909155835260089091529020546114fe90600161155e565b600160a060020a0390921660009081526008602052604090209190915550565b6000903b1190565b600090815260116020526040902080546fffffffffffffffffffffffffffffffff191681556001018054600160a060020a0319169055565b81810182811015610a8657fe5b604080518082019091526000808252602082015290565b8154818355818111156107575760008381526020902061075791810190830161065991905b808211156115bb57600081556001016115a7565b50905600a165627a7a72305820ffb926c01eada2f24cbbcd6843886dee2f8054efac9892424e0b89a04f23e15a0029`

// DeployStampTransaction deploys a new Ethereum contract, binding an instance of StampTransaction to it.
func DeployStampTransaction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StampTransaction, error) {
	parsed, err := abi.JSON(strings.NewReader(StampTransactionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StampTransactionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StampTransaction{StampTransactionCaller: StampTransactionCaller{contract: contract}, StampTransactionTransactor: StampTransactionTransactor{contract: contract}, StampTransactionFilterer: StampTransactionFilterer{contract: contract}}, nil
}

// StampTransaction is an auto generated Go binding around an Ethereum contract.
type StampTransaction struct {
	StampTransactionCaller     // Read-only binding to the contract
	StampTransactionTransactor // Write-only binding to the contract
	StampTransactionFilterer   // Log filterer for contract events
}

// StampTransactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type StampTransactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StampTransactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StampTransactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StampTransactionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StampTransactionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StampTransactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StampTransactionSession struct {
	Contract     *StampTransaction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StampTransactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StampTransactionCallerSession struct {
	Contract *StampTransactionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// StampTransactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StampTransactionTransactorSession struct {
	Contract     *StampTransactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// StampTransactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type StampTransactionRaw struct {
	Contract *StampTransaction // Generic contract binding to access the raw methods on
}

// StampTransactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StampTransactionCallerRaw struct {
	Contract *StampTransactionCaller // Generic read-only contract binding to access the raw methods on
}

// StampTransactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StampTransactionTransactorRaw struct {
	Contract *StampTransactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStampTransaction creates a new instance of StampTransaction, bound to a specific deployed contract.
func NewStampTransaction(address common.Address, backend bind.ContractBackend) (*StampTransaction, error) {
	contract, err := bindStampTransaction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StampTransaction{StampTransactionCaller: StampTransactionCaller{contract: contract}, StampTransactionTransactor: StampTransactionTransactor{contract: contract}, StampTransactionFilterer: StampTransactionFilterer{contract: contract}}, nil
}

// NewStampTransactionCaller creates a new read-only instance of StampTransaction, bound to a specific deployed contract.
func NewStampTransactionCaller(address common.Address, caller bind.ContractCaller) (*StampTransactionCaller, error) {
	contract, err := bindStampTransaction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StampTransactionCaller{contract: contract}, nil
}

// NewStampTransactionTransactor creates a new write-only instance of StampTransaction, bound to a specific deployed contract.
func NewStampTransactionTransactor(address common.Address, transactor bind.ContractTransactor) (*StampTransactionTransactor, error) {
	contract, err := bindStampTransaction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StampTransactionTransactor{contract: contract}, nil
}

// NewStampTransactionFilterer creates a new log filterer instance of StampTransaction, bound to a specific deployed contract.
func NewStampTransactionFilterer(address common.Address, filterer bind.ContractFilterer) (*StampTransactionFilterer, error) {
	contract, err := bindStampTransaction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StampTransactionFilterer{contract: contract}, nil
}

// bindStampTransaction binds a generic wrapper to an already deployed contract.
func bindStampTransaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StampTransactionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StampTransaction *StampTransactionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StampTransaction.Contract.StampTransactionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StampTransaction *StampTransactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampTransaction.Contract.StampTransactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StampTransaction *StampTransactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StampTransaction.Contract.StampTransactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StampTransaction *StampTransactionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StampTransaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StampTransaction *StampTransactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampTransaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StampTransaction *StampTransactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StampTransaction.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StampTransaction *StampTransactionCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampTransaction.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StampTransaction *StampTransactionSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StampTransaction.Contract.BalanceOf(&_StampTransaction.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StampTransaction *StampTransactionCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StampTransaction.Contract.BalanceOf(&_StampTransaction.CallOpts, _owner)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_StampTransaction *StampTransactionCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampTransaction.contract.Call(opts, out, "ceoAddress")
	return *ret0, err
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_StampTransaction *StampTransactionSession) CeoAddress() (common.Address, error) {
	return _StampTransaction.Contract.CeoAddress(&_StampTransaction.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_StampTransaction *StampTransactionCallerSession) CeoAddress() (common.Address, error) {
	return _StampTransaction.Contract.CeoAddress(&_StampTransaction.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_StampTransaction *StampTransactionCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampTransaction.contract.Call(opts, out, "cfoAddress")
	return *ret0, err
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_StampTransaction *StampTransactionSession) CfoAddress() (common.Address, error) {
	return _StampTransaction.Contract.CfoAddress(&_StampTransaction.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_StampTransaction *StampTransactionCallerSession) CfoAddress() (common.Address, error) {
	return _StampTransaction.Contract.CfoAddress(&_StampTransaction.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_StampTransaction *StampTransactionCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampTransaction.contract.Call(opts, out, "cooAddress")
	return *ret0, err
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_StampTransaction *StampTransactionSession) CooAddress() (common.Address, error) {
	return _StampTransaction.Contract.CooAddress(&_StampTransaction.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_StampTransaction *StampTransactionCallerSession) CooAddress() (common.Address, error) {
	return _StampTransaction.Contract.CooAddress(&_StampTransaction.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_StampTransaction *StampTransactionCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StampTransaction.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_StampTransaction *StampTransactionSession) Exists(_tokenId *big.Int) (bool, error) {
	return _StampTransaction.Contract.Exists(&_StampTransaction.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_StampTransaction *StampTransactionCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _StampTransaction.Contract.Exists(&_StampTransaction.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_StampTransaction *StampTransactionCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampTransaction.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_StampTransaction *StampTransactionSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _StampTransaction.Contract.GetApproved(&_StampTransaction.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_StampTransaction *StampTransactionCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _StampTransaction.Contract.GetApproved(&_StampTransaction.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(approved bool)
func (_StampTransaction *StampTransactionCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StampTransaction.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(approved bool)
func (_StampTransaction *StampTransactionSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _StampTransaction.Contract.IsApprovedForAll(&_StampTransaction.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(approved bool)
func (_StampTransaction *StampTransactionCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _StampTransaction.Contract.IsApprovedForAll(&_StampTransaction.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StampTransaction *StampTransactionCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StampTransaction.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StampTransaction *StampTransactionSession) Name() (string, error) {
	return _StampTransaction.Contract.Name(&_StampTransaction.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StampTransaction *StampTransactionCallerSession) Name() (string, error) {
	return _StampTransaction.Contract.Name(&_StampTransaction.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_StampTransaction *StampTransactionCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StampTransaction.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_StampTransaction *StampTransactionSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _StampTransaction.Contract.OwnerOf(&_StampTransaction.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_StampTransaction *StampTransactionCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _StampTransaction.Contract.OwnerOf(&_StampTransaction.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_StampTransaction *StampTransactionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StampTransaction.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_StampTransaction *StampTransactionSession) Paused() (bool, error) {
	return _StampTransaction.Contract.Paused(&_StampTransaction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_StampTransaction *StampTransactionCallerSession) Paused() (bool, error) {
	return _StampTransaction.Contract.Paused(&_StampTransaction.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_StampTransaction *StampTransactionCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StampTransaction.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_StampTransaction *StampTransactionSession) Symbol() (string, error) {
	return _StampTransaction.Contract.Symbol(&_StampTransaction.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_StampTransaction *StampTransactionCallerSession) Symbol() (string, error) {
	return _StampTransaction.Contract.Symbol(&_StampTransaction.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_StampTransaction *StampTransactionCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampTransaction.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_StampTransaction *StampTransactionSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _StampTransaction.Contract.TokenByIndex(&_StampTransaction.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_StampTransaction *StampTransactionCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _StampTransaction.Contract.TokenByIndex(&_StampTransaction.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_StampTransaction *StampTransactionCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampTransaction.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_StampTransaction *StampTransactionSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _StampTransaction.Contract.TokenOfOwnerByIndex(&_StampTransaction.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_StampTransaction *StampTransactionCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _StampTransaction.Contract.TokenOfOwnerByIndex(&_StampTransaction.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_StampTransaction *StampTransactionCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StampTransaction.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_StampTransaction *StampTransactionSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _StampTransaction.Contract.TokenURI(&_StampTransaction.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_StampTransaction *StampTransactionCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _StampTransaction.Contract.TokenURI(&_StampTransaction.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StampTransaction *StampTransactionCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StampTransaction.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StampTransaction *StampTransactionSession) TotalSupply() (*big.Int, error) {
	return _StampTransaction.Contract.TotalSupply(&_StampTransaction.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StampTransaction *StampTransactionCallerSession) TotalSupply() (*big.Int, error) {
	return _StampTransaction.Contract.TotalSupply(&_StampTransaction.CallOpts)
}

// TransactionInfo is a free data retrieval call binding the contract method 0xf4fae995.
//
// Solidity: function transactionInfo(_tokenId uint256) constant returns(price uint128, seller address)
func (_StampTransaction *StampTransactionCaller) TransactionInfo(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Price  *big.Int
	Seller common.Address
}, error) {
	ret := new(struct {
		Price  *big.Int
		Seller common.Address
	})
	out := ret
	err := _StampTransaction.contract.Call(opts, out, "transactionInfo", _tokenId)
	return *ret, err
}

// TransactionInfo is a free data retrieval call binding the contract method 0xf4fae995.
//
// Solidity: function transactionInfo(_tokenId uint256) constant returns(price uint128, seller address)
func (_StampTransaction *StampTransactionSession) TransactionInfo(_tokenId *big.Int) (struct {
	Price  *big.Int
	Seller common.Address
}, error) {
	return _StampTransaction.Contract.TransactionInfo(&_StampTransaction.CallOpts, _tokenId)
}

// TransactionInfo is a free data retrieval call binding the contract method 0xf4fae995.
//
// Solidity: function transactionInfo(_tokenId uint256) constant returns(price uint128, seller address)
func (_StampTransaction *StampTransactionCallerSession) TransactionInfo(_tokenId *big.Int) (struct {
	Price  *big.Int
	Seller common.Address
}, error) {
	return _StampTransaction.Contract.TransactionInfo(&_StampTransaction.CallOpts, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_StampTransaction *StampTransactionTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampTransaction.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_StampTransaction *StampTransactionSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampTransaction.Contract.Approve(&_StampTransaction.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_StampTransaction *StampTransactionTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampTransaction.Contract.Approve(&_StampTransaction.TransactOpts, _to, _tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(_tokenId uint256) returns()
func (_StampTransaction *StampTransactionTransactor) Buy(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampTransaction.contract.Transact(opts, "buy", _tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(_tokenId uint256) returns()
func (_StampTransaction *StampTransactionSession) Buy(_tokenId *big.Int) (*types.Transaction, error) {
	return _StampTransaction.Contract.Buy(&_StampTransaction.TransactOpts, _tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(_tokenId uint256) returns()
func (_StampTransaction *StampTransactionTransactorSession) Buy(_tokenId *big.Int) (*types.Transaction, error) {
	return _StampTransaction.Contract.Buy(&_StampTransaction.TransactOpts, _tokenId)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xa82f8d58.
//
// Solidity: function createTransaction(_tokenId uint256, _price uint256, _seller address) returns()
func (_StampTransaction *StampTransactionTransactor) CreateTransaction(opts *bind.TransactOpts, _tokenId *big.Int, _price *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _StampTransaction.contract.Transact(opts, "createTransaction", _tokenId, _price, _seller)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xa82f8d58.
//
// Solidity: function createTransaction(_tokenId uint256, _price uint256, _seller address) returns()
func (_StampTransaction *StampTransactionSession) CreateTransaction(_tokenId *big.Int, _price *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _StampTransaction.Contract.CreateTransaction(&_StampTransaction.TransactOpts, _tokenId, _price, _seller)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xa82f8d58.
//
// Solidity: function createTransaction(_tokenId uint256, _price uint256, _seller address) returns()
func (_StampTransaction *StampTransactionTransactorSession) CreateTransaction(_tokenId *big.Int, _price *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _StampTransaction.Contract.CreateTransaction(&_StampTransaction.TransactOpts, _tokenId, _price, _seller)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StampTransaction *StampTransactionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampTransaction.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StampTransaction *StampTransactionSession) Pause() (*types.Transaction, error) {
	return _StampTransaction.Contract.Pause(&_StampTransaction.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StampTransaction *StampTransactionTransactorSession) Pause() (*types.Transaction, error) {
	return _StampTransaction.Contract.Pause(&_StampTransaction.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_StampTransaction *StampTransactionTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _StampTransaction.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_StampTransaction *StampTransactionSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _StampTransaction.Contract.SafeTransferFrom(&_StampTransaction.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_StampTransaction *StampTransactionTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _StampTransaction.Contract.SafeTransferFrom(&_StampTransaction.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_StampTransaction *StampTransactionTransactor) SetApprovalForAll(opts *bind.TransactOpts, _to common.Address, _approved bool) (*types.Transaction, error) {
	return _StampTransaction.contract.Transact(opts, "setApprovalForAll", _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_StampTransaction *StampTransactionSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _StampTransaction.Contract.SetApprovalForAll(&_StampTransaction.TransactOpts, _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_StampTransaction *StampTransactionTransactorSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _StampTransaction.Contract.SetApprovalForAll(&_StampTransaction.TransactOpts, _to, _approved)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_StampTransaction *StampTransactionTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _StampTransaction.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_StampTransaction *StampTransactionSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _StampTransaction.Contract.SetCEO(&_StampTransaction.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_StampTransaction *StampTransactionTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _StampTransaction.Contract.SetCEO(&_StampTransaction.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_StampTransaction *StampTransactionTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _StampTransaction.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_StampTransaction *StampTransactionSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _StampTransaction.Contract.SetCFO(&_StampTransaction.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_StampTransaction *StampTransactionTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _StampTransaction.Contract.SetCFO(&_StampTransaction.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_StampTransaction *StampTransactionTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _StampTransaction.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_StampTransaction *StampTransactionSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _StampTransaction.Contract.SetCOO(&_StampTransaction.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_StampTransaction *StampTransactionTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _StampTransaction.Contract.SetCOO(&_StampTransaction.TransactOpts, _newCOO)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_StampTransaction *StampTransactionTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampTransaction.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_StampTransaction *StampTransactionSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampTransaction.Contract.TransferFrom(&_StampTransaction.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_StampTransaction *StampTransactionTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _StampTransaction.Contract.TransferFrom(&_StampTransaction.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StampTransaction *StampTransactionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StampTransaction.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StampTransaction *StampTransactionSession) Unpause() (*types.Transaction, error) {
	return _StampTransaction.Contract.Unpause(&_StampTransaction.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StampTransaction *StampTransactionTransactorSession) Unpause() (*types.Transaction, error) {
	return _StampTransaction.Contract.Unpause(&_StampTransaction.TransactOpts)
}

// StampTransactionApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StampTransaction contract.
type StampTransactionApprovalIterator struct {
	Event *StampTransactionApproval // Event containing the contract specifics and raw log

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
func (it *StampTransactionApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampTransactionApproval)
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
		it.Event = new(StampTransactionApproval)
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
func (it *StampTransactionApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampTransactionApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampTransactionApproval represents a Approval event raised by the StampTransaction contract.
type StampTransactionApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_StampTransaction *StampTransactionFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address) (*StampTransactionApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _StampTransaction.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return &StampTransactionApprovalIterator{contract: _StampTransaction.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_StampTransaction *StampTransactionFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StampTransactionApproval, _owner []common.Address, _approved []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _StampTransaction.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampTransactionApproval)
				if err := _StampTransaction.contract.UnpackLog(event, "Approval", log); err != nil {
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

// StampTransactionApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the StampTransaction contract.
type StampTransactionApprovalForAllIterator struct {
	Event *StampTransactionApprovalForAll // Event containing the contract specifics and raw log

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
func (it *StampTransactionApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampTransactionApprovalForAll)
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
		it.Event = new(StampTransactionApprovalForAll)
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
func (it *StampTransactionApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampTransactionApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampTransactionApprovalForAll represents a ApprovalForAll event raised by the StampTransaction contract.
type StampTransactionApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_StampTransaction *StampTransactionFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*StampTransactionApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _StampTransaction.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &StampTransactionApprovalForAllIterator{contract: _StampTransaction.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_StampTransaction *StampTransactionFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *StampTransactionApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _StampTransaction.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampTransactionApprovalForAll)
				if err := _StampTransaction.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// StampTransactionContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the StampTransaction contract.
type StampTransactionContractUpgradeIterator struct {
	Event *StampTransactionContractUpgrade // Event containing the contract specifics and raw log

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
func (it *StampTransactionContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampTransactionContractUpgrade)
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
		it.Event = new(StampTransactionContractUpgrade)
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
func (it *StampTransactionContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampTransactionContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampTransactionContractUpgrade represents a ContractUpgrade event raised by the StampTransaction contract.
type StampTransactionContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_StampTransaction *StampTransactionFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*StampTransactionContractUpgradeIterator, error) {

	logs, sub, err := _StampTransaction.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &StampTransactionContractUpgradeIterator{contract: _StampTransaction.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_StampTransaction *StampTransactionFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *StampTransactionContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _StampTransaction.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampTransactionContractUpgrade)
				if err := _StampTransaction.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// StampTransactionCreateNewStampIterator is returned from FilterCreateNewStamp and is used to iterate over the raw logs and unpacked data for CreateNewStamp events raised by the StampTransaction contract.
type StampTransactionCreateNewStampIterator struct {
	Event *StampTransactionCreateNewStamp // Event containing the contract specifics and raw log

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
func (it *StampTransactionCreateNewStampIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampTransactionCreateNewStamp)
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
		it.Event = new(StampTransactionCreateNewStamp)
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
func (it *StampTransactionCreateNewStampIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampTransactionCreateNewStampIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampTransactionCreateNewStamp represents a CreateNewStamp event raised by the StampTransaction contract.
type StampTransactionCreateNewStamp struct {
	StampId         uint32
	TotalAmount     uint32
	RemainingAmount uint32
	Name            [32]byte
	Year            uint16
	SetId           uint16
	MemberOfSetId   uint8
	Appearance      uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCreateNewStamp is a free log retrieval operation binding the contract event 0x4710117ee2ed24df4322c27506325e38cbbb47a788bc132e0c5d7d107f434256.
//
// Solidity: e CreateNewStamp(stampId uint32, totalAmount uint32, remainingAmount uint32, name bytes32, year uint16, setId uint16, memberOfSetId uint8, appearance uint8)
func (_StampTransaction *StampTransactionFilterer) FilterCreateNewStamp(opts *bind.FilterOpts) (*StampTransactionCreateNewStampIterator, error) {

	logs, sub, err := _StampTransaction.contract.FilterLogs(opts, "CreateNewStamp")
	if err != nil {
		return nil, err
	}
	return &StampTransactionCreateNewStampIterator{contract: _StampTransaction.contract, event: "CreateNewStamp", logs: logs, sub: sub}, nil
}

// WatchCreateNewStamp is a free log subscription operation binding the contract event 0x4710117ee2ed24df4322c27506325e38cbbb47a788bc132e0c5d7d107f434256.
//
// Solidity: e CreateNewStamp(stampId uint32, totalAmount uint32, remainingAmount uint32, name bytes32, year uint16, setId uint16, memberOfSetId uint8, appearance uint8)
func (_StampTransaction *StampTransactionFilterer) WatchCreateNewStamp(opts *bind.WatchOpts, sink chan<- *StampTransactionCreateNewStamp) (event.Subscription, error) {

	logs, sub, err := _StampTransaction.contract.WatchLogs(opts, "CreateNewStamp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampTransactionCreateNewStamp)
				if err := _StampTransaction.contract.UnpackLog(event, "CreateNewStamp", log); err != nil {
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

// StampTransactionTransactionCancelledIterator is returned from FilterTransactionCancelled and is used to iterate over the raw logs and unpacked data for TransactionCancelled events raised by the StampTransaction contract.
type StampTransactionTransactionCancelledIterator struct {
	Event *StampTransactionTransactionCancelled // Event containing the contract specifics and raw log

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
func (it *StampTransactionTransactionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampTransactionTransactionCancelled)
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
		it.Event = new(StampTransactionTransactionCancelled)
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
func (it *StampTransactionTransactionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampTransactionTransactionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampTransactionTransactionCancelled represents a TransactionCancelled event raised by the StampTransaction contract.
type StampTransactionTransactionCancelled struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransactionCancelled is a free log retrieval operation binding the contract event 0x956fb32199d8b882b2039a14e1be35ab14f7a80b9089fc223f14b43937173e60.
//
// Solidity: e TransactionCancelled(tokenId uint256)
func (_StampTransaction *StampTransactionFilterer) FilterTransactionCancelled(opts *bind.FilterOpts) (*StampTransactionTransactionCancelledIterator, error) {

	logs, sub, err := _StampTransaction.contract.FilterLogs(opts, "TransactionCancelled")
	if err != nil {
		return nil, err
	}
	return &StampTransactionTransactionCancelledIterator{contract: _StampTransaction.contract, event: "TransactionCancelled", logs: logs, sub: sub}, nil
}

// WatchTransactionCancelled is a free log subscription operation binding the contract event 0x956fb32199d8b882b2039a14e1be35ab14f7a80b9089fc223f14b43937173e60.
//
// Solidity: e TransactionCancelled(tokenId uint256)
func (_StampTransaction *StampTransactionFilterer) WatchTransactionCancelled(opts *bind.WatchOpts, sink chan<- *StampTransactionTransactionCancelled) (event.Subscription, error) {

	logs, sub, err := _StampTransaction.contract.WatchLogs(opts, "TransactionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampTransactionTransactionCancelled)
				if err := _StampTransaction.contract.UnpackLog(event, "TransactionCancelled", log); err != nil {
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

// StampTransactionTransactionCreatedIterator is returned from FilterTransactionCreated and is used to iterate over the raw logs and unpacked data for TransactionCreated events raised by the StampTransaction contract.
type StampTransactionTransactionCreatedIterator struct {
	Event *StampTransactionTransactionCreated // Event containing the contract specifics and raw log

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
func (it *StampTransactionTransactionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampTransactionTransactionCreated)
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
		it.Event = new(StampTransactionTransactionCreated)
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
func (it *StampTransactionTransactionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampTransactionTransactionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampTransactionTransactionCreated represents a TransactionCreated event raised by the StampTransaction contract.
type StampTransactionTransactionCreated struct {
	TokenId *big.Int
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransactionCreated is a free log retrieval operation binding the contract event 0x6110ba4ca3a2a4b05be8a5160cfb9b4dfbdc1738e1c8e424a1546fc3f2bce843.
//
// Solidity: e TransactionCreated(tokenId uint256, price uint256)
func (_StampTransaction *StampTransactionFilterer) FilterTransactionCreated(opts *bind.FilterOpts) (*StampTransactionTransactionCreatedIterator, error) {

	logs, sub, err := _StampTransaction.contract.FilterLogs(opts, "TransactionCreated")
	if err != nil {
		return nil, err
	}
	return &StampTransactionTransactionCreatedIterator{contract: _StampTransaction.contract, event: "TransactionCreated", logs: logs, sub: sub}, nil
}

// WatchTransactionCreated is a free log subscription operation binding the contract event 0x6110ba4ca3a2a4b05be8a5160cfb9b4dfbdc1738e1c8e424a1546fc3f2bce843.
//
// Solidity: e TransactionCreated(tokenId uint256, price uint256)
func (_StampTransaction *StampTransactionFilterer) WatchTransactionCreated(opts *bind.WatchOpts, sink chan<- *StampTransactionTransactionCreated) (event.Subscription, error) {

	logs, sub, err := _StampTransaction.contract.WatchLogs(opts, "TransactionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampTransactionTransactionCreated)
				if err := _StampTransaction.contract.UnpackLog(event, "TransactionCreated", log); err != nil {
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

// StampTransactionTransactionSuccessfulIterator is returned from FilterTransactionSuccessful and is used to iterate over the raw logs and unpacked data for TransactionSuccessful events raised by the StampTransaction contract.
type StampTransactionTransactionSuccessfulIterator struct {
	Event *StampTransactionTransactionSuccessful // Event containing the contract specifics and raw log

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
func (it *StampTransactionTransactionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampTransactionTransactionSuccessful)
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
		it.Event = new(StampTransactionTransactionSuccessful)
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
func (it *StampTransactionTransactionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampTransactionTransactionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampTransactionTransactionSuccessful represents a TransactionSuccessful event raised by the StampTransaction contract.
type StampTransactionTransactionSuccessful struct {
	TokenId *big.Int
	Price   *big.Int
	Buyer   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransactionSuccessful is a free log retrieval operation binding the contract event 0x7b31057475c88e821f05f6e63ac0a08a794cf369937c11da22272b73ef90187f.
//
// Solidity: e TransactionSuccessful(tokenId uint256, price uint256, buyer address)
func (_StampTransaction *StampTransactionFilterer) FilterTransactionSuccessful(opts *bind.FilterOpts) (*StampTransactionTransactionSuccessfulIterator, error) {

	logs, sub, err := _StampTransaction.contract.FilterLogs(opts, "TransactionSuccessful")
	if err != nil {
		return nil, err
	}
	return &StampTransactionTransactionSuccessfulIterator{contract: _StampTransaction.contract, event: "TransactionSuccessful", logs: logs, sub: sub}, nil
}

// WatchTransactionSuccessful is a free log subscription operation binding the contract event 0x7b31057475c88e821f05f6e63ac0a08a794cf369937c11da22272b73ef90187f.
//
// Solidity: e TransactionSuccessful(tokenId uint256, price uint256, buyer address)
func (_StampTransaction *StampTransactionFilterer) WatchTransactionSuccessful(opts *bind.WatchOpts, sink chan<- *StampTransactionTransactionSuccessful) (event.Subscription, error) {

	logs, sub, err := _StampTransaction.contract.WatchLogs(opts, "TransactionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampTransactionTransactionSuccessful)
				if err := _StampTransaction.contract.UnpackLog(event, "TransactionSuccessful", log); err != nil {
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

// StampTransactionTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StampTransaction contract.
type StampTransactionTransferIterator struct {
	Event *StampTransactionTransfer // Event containing the contract specifics and raw log

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
func (it *StampTransactionTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampTransactionTransfer)
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
		it.Event = new(StampTransactionTransfer)
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
func (it *StampTransactionTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampTransactionTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampTransactionTransfer represents a Transfer event raised by the StampTransaction contract.
type StampTransactionTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_StampTransaction *StampTransactionFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*StampTransactionTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _StampTransaction.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &StampTransactionTransferIterator{contract: _StampTransaction.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_StampTransaction *StampTransactionFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StampTransactionTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _StampTransaction.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampTransactionTransfer)
				if err := _StampTransaction.contract.UnpackLog(event, "Transfer", log); err != nil {
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

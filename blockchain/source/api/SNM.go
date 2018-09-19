// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// SNMABI is the input ABI used to generate the binding from.
const SNMABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ico\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokensAreFrozen\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_holder\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SNMBin is the compiled bytecode used for deploying new contracts.
const SNMBin = `0x60c0604052600960808190527f4e414320546f6b656e000000000000000000000000000000000000000000000060a090815261003e91600391906100c8565b506040805180820190915260038082527f4e414300000000000000000000000000000000000000000000000000000000006020909201918252610083916004916100c8565b5060126005556006805460a060020a60ff02191690553480156100a557600080fd5b503360009081526001602052604090206a21165458500521280000009055610163565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061010957805160ff1916838001178555610136565b82800160010185558215610136579182015b8281111561013657825182559160200191906001019061011b565b50610142929150610146565b5090565b61016091905b80821115610142576000815560010161014c565b90565b61079c806101726000396000f3006080604052600436106100b95763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100be578063095ea7b31461014857806318160ddd1461016e57806323b872dd14610195578063313ce567146101d357806340c10f19146101e85780635d4522011461020c57806370a082311461023d57806395d89b411461025e578063a9059cbb14610273578063ca67065f14610297578063dd62ed3e146102ac575b600080fd5b3480156100ca57600080fd5b506100d36102d3565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561010d5781810151838201526020016100f5565b50505050905090810190601f16801561013a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561015457600080fd5b5061016c600160a060020a0360043516602435610361565b005b34801561017a57600080fd5b5061018361036f565b60408051918252519081900360200190f35b3480156101a157600080fd5b506101bf600160a060020a0360043581169060243516604435610375565b604080519115158252519081900360200190f35b3480156101df57600080fd5b5061018361038d565b3480156101f457600080fd5b5061016c600160a060020a0360043516602435610393565b34801561021857600080fd5b5061022161041a565b60408051600160a060020a039092168252519081900360200190f35b34801561024957600080fd5b50610183600160a060020a0360043516610429565b34801561026a57600080fd5b506100d3610444565b34801561027f57600080fd5b506101bf600160a060020a036004351660243561049f565b3480156102a357600080fd5b506101bf6104b5565b3480156102b857600080fd5b50610183600160a060020a03600435811690602435166104d6565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103595780601f1061032e57610100808354040283529160200191610359565b820191906000526020600020905b81548152906001019060200180831161033c57829003601f168201915b505050505081565b61036b8282610501565b5050565b60005481565b6000610382848484610563565b506001949350505050565b60055481565b80151561039f57600080fd5b6000546b016f44a83aab6c233c00000090820111156103bd57600080fd5b600160a060020a0382166000818152600160209081526040808320805486019055825485018355805185815290517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35050565b600654600160a060020a031681565b600160a060020a031660009081526001602052604090205490565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103595780601f1061032e57610100808354040283529160200191610359565b60006104ab8383610672565b5060019392505050565b60065474010000000000000000000000000000000000000000900460ff1681565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b336000818152600260209081526040808320600160a060020a03871680855290835292819020859055805185815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35050565b600160a060020a038084166000908152600260209081526040808320338452825280832054938616835260019091528120549091906105a8908463ffffffff61073516565b600160a060020a0380861660009081526001602052604080822093909355908716815220546105dd908463ffffffff61074d16565b600160a060020a038616600090815260016020526040902055610606818463ffffffff61074d16565b600160a060020a03808716600081815260026020908152604080832033845282529182902094909455805187815290519288169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3506001949350505050565b60006040604436101561068457600080fd5b336000908152600160205260409020546106a4908463ffffffff61074d16565b3360009081526001602052604080822092909255600160a060020a038616815220546106d6908463ffffffff61073516565b600160a060020a0385166000818152600160209081526040918290209390935580518681529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35060019392505050565b600082820161074684821015610761565b9392505050565b600061075b83831115610761565b50900390565b80151561076d57600080fd5b505600a165627a7a7230582059d6ad639aacca4eefe6418af1baf721f6d655829ecaf51a652bc720f8bd7b030029`

// DeploySNM deploys a new Ethereum contract, binding an instance of SNM to it.
func DeploySNM(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SNM, error) {
	parsed, err := abi.JSON(strings.NewReader(SNMABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SNMBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SNM{SNMCaller: SNMCaller{contract: contract}, SNMTransactor: SNMTransactor{contract: contract}, SNMFilterer: SNMFilterer{contract: contract}}, nil
}

// SNM is an auto generated Go binding around an Ethereum contract.
type SNM struct {
	SNMCaller     // Read-only binding to the contract
	SNMTransactor // Write-only binding to the contract
	SNMFilterer   // Log filterer for contract events
}

// SNMCaller is an auto generated read-only Go binding around an Ethereum contract.
type SNMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SNMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SNMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SNMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SNMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SNMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SNMSession struct {
	Contract     *SNM              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SNMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SNMCallerSession struct {
	Contract *SNMCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SNMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SNMTransactorSession struct {
	Contract     *SNMTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SNMRaw is an auto generated low-level Go binding around an Ethereum contract.
type SNMRaw struct {
	Contract *SNM // Generic contract binding to access the raw methods on
}

// SNMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SNMCallerRaw struct {
	Contract *SNMCaller // Generic read-only contract binding to access the raw methods on
}

// SNMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SNMTransactorRaw struct {
	Contract *SNMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSNM creates a new instance of SNM, bound to a specific deployed contract.
func NewSNM(address common.Address, backend bind.ContractBackend) (*SNM, error) {
	contract, err := bindSNM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SNM{SNMCaller: SNMCaller{contract: contract}, SNMTransactor: SNMTransactor{contract: contract}, SNMFilterer: SNMFilterer{contract: contract}}, nil
}

// NewSNMCaller creates a new read-only instance of SNM, bound to a specific deployed contract.
func NewSNMCaller(address common.Address, caller bind.ContractCaller) (*SNMCaller, error) {
	contract, err := bindSNM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SNMCaller{contract: contract}, nil
}

// NewSNMTransactor creates a new write-only instance of SNM, bound to a specific deployed contract.
func NewSNMTransactor(address common.Address, transactor bind.ContractTransactor) (*SNMTransactor, error) {
	contract, err := bindSNM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SNMTransactor{contract: contract}, nil
}

// NewSNMFilterer creates a new log filterer instance of SNM, bound to a specific deployed contract.
func NewSNMFilterer(address common.Address, filterer bind.ContractFilterer) (*SNMFilterer, error) {
	contract, err := bindSNM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SNMFilterer{contract: contract}, nil
}

// bindSNM binds a generic wrapper to an already deployed contract.
func bindSNM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SNMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SNM *SNMRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SNM.Contract.SNMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SNM *SNMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SNM.Contract.SNMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SNM *SNMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SNM.Contract.SNMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SNM *SNMCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SNM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SNM *SNMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SNM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SNM *SNMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SNM.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_SNM *SNMCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SNM.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_SNM *SNMSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _SNM.Contract.Allowance(&_SNM.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_SNM *SNMCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _SNM.Contract.Allowance(&_SNM.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SNM *SNMCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SNM.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SNM *SNMSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _SNM.Contract.BalanceOf(&_SNM.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SNM *SNMCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _SNM.Contract.BalanceOf(&_SNM.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_SNM *SNMCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SNM.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_SNM *SNMSession) Decimals() (*big.Int, error) {
	return _SNM.Contract.Decimals(&_SNM.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_SNM *SNMCallerSession) Decimals() (*big.Int, error) {
	return _SNM.Contract.Decimals(&_SNM.CallOpts)
}

// Ico is a free data retrieval call binding the contract method 0x5d452201.
//
// Solidity: function ico() constant returns(address)
func (_SNM *SNMCaller) Ico(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SNM.contract.Call(opts, out, "ico")
	return *ret0, err
}

// Ico is a free data retrieval call binding the contract method 0x5d452201.
//
// Solidity: function ico() constant returns(address)
func (_SNM *SNMSession) Ico() (common.Address, error) {
	return _SNM.Contract.Ico(&_SNM.CallOpts)
}

// Ico is a free data retrieval call binding the contract method 0x5d452201.
//
// Solidity: function ico() constant returns(address)
func (_SNM *SNMCallerSession) Ico() (common.Address, error) {
	return _SNM.Contract.Ico(&_SNM.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SNM *SNMCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SNM.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SNM *SNMSession) Name() (string, error) {
	return _SNM.Contract.Name(&_SNM.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SNM *SNMCallerSession) Name() (string, error) {
	return _SNM.Contract.Name(&_SNM.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SNM *SNMCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SNM.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SNM *SNMSession) Symbol() (string, error) {
	return _SNM.Contract.Symbol(&_SNM.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SNM *SNMCallerSession) Symbol() (string, error) {
	return _SNM.Contract.Symbol(&_SNM.CallOpts)
}

// TokensAreFrozen is a free data retrieval call binding the contract method 0xca67065f.
//
// Solidity: function tokensAreFrozen() constant returns(bool)
func (_SNM *SNMCaller) TokensAreFrozen(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SNM.contract.Call(opts, out, "tokensAreFrozen")
	return *ret0, err
}

// TokensAreFrozen is a free data retrieval call binding the contract method 0xca67065f.
//
// Solidity: function tokensAreFrozen() constant returns(bool)
func (_SNM *SNMSession) TokensAreFrozen() (bool, error) {
	return _SNM.Contract.TokensAreFrozen(&_SNM.CallOpts)
}

// TokensAreFrozen is a free data retrieval call binding the contract method 0xca67065f.
//
// Solidity: function tokensAreFrozen() constant returns(bool)
func (_SNM *SNMCallerSession) TokensAreFrozen() (bool, error) {
	return _SNM.Contract.TokensAreFrozen(&_SNM.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SNM *SNMCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SNM.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SNM *SNMSession) TotalSupply() (*big.Int, error) {
	return _SNM.Contract.TotalSupply(&_SNM.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SNM *SNMCallerSession) TotalSupply() (*big.Int, error) {
	return _SNM.Contract.TotalSupply(&_SNM.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns()
func (_SNM *SNMTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SNM.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns()
func (_SNM *SNMSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SNM.Contract.Approve(&_SNM.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns()
func (_SNM *SNMTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SNM.Contract.Approve(&_SNM.TransactOpts, _spender, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_holder address, _value uint256) returns()
func (_SNM *SNMTransactor) Mint(opts *bind.TransactOpts, _holder common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SNM.contract.Transact(opts, "mint", _holder, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_holder address, _value uint256) returns()
func (_SNM *SNMSession) Mint(_holder common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SNM.Contract.Mint(&_SNM.TransactOpts, _holder, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_holder address, _value uint256) returns()
func (_SNM *SNMTransactorSession) Mint(_holder common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SNM.Contract.Mint(&_SNM.TransactOpts, _holder, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_SNM *SNMTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SNM.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_SNM *SNMSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SNM.Contract.Transfer(&_SNM.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_SNM *SNMTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SNM.Contract.Transfer(&_SNM.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_SNM *SNMTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SNM.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_SNM *SNMSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SNM.Contract.TransferFrom(&_SNM.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_SNM *SNMTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SNM.Contract.TransferFrom(&_SNM.TransactOpts, _from, _to, _value)
}

// SNMApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SNM contract.
type SNMApprovalIterator struct {
	Event *SNMApproval // Event containing the contract specifics and raw log

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
func (it *SNMApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SNMApproval)
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
		it.Event = new(SNMApproval)
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
func (it *SNMApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SNMApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SNMApproval represents a Approval event raised by the SNM contract.
type SNMApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_SNM *SNMFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SNMApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SNM.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SNMApprovalIterator{contract: _SNM.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_SNM *SNMFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SNMApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SNM.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SNMApproval)
				if err := _SNM.contract.UnpackLog(event, "Approval", log); err != nil {
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

// SNMTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SNM contract.
type SNMTransferIterator struct {
	Event *SNMTransfer // Event containing the contract specifics and raw log

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
func (it *SNMTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SNMTransfer)
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
		it.Event = new(SNMTransfer)
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
func (it *SNMTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SNMTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SNMTransfer represents a Transfer event raised by the SNM contract.
type SNMTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_SNM *SNMFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SNMTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SNM.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SNMTransferIterator{contract: _SNM.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_SNM *SNMFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SNMTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SNM.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SNMTransfer)
				if err := _SNM.contract.UnpackLog(event, "Transfer", log); err != nil {
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

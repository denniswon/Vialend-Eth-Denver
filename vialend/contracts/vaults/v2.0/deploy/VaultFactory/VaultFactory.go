// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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
)

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_team\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_this\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"CreatedVault\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stat\",\"type\":\"uint256\"}],\"name\":\"changeStat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sORv\",\"type\":\"address\"}],\"name\":\"checkActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[10]\",\"name\":\"contracts\",\"type\":\"address[10]\"},{\"internalType\":\"uint256\",\"name\":\"_vaultCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_individualCap\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_uniPortion\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_compPortion\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_protocolFee\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"_feetier\",\"type\":\"uint24\"},{\"internalType\":\"uint128\",\"name\":\"_quoteAmount\",\"type\":\"uint128\"}],\"name\":\"createVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"getStat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTeam\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pairStrategy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pairVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"repair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_team\",\"type\":\"address\"}],\"name\":\"setTeam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vaults\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610b79380380610b7983398101604081905261002f91610062565b60008054336001600160a01b031991821617909155600180549091166001600160a01b0392909216919091179055610092565b60006020828403121561007457600080fd5b81516001600160a01b038116811461008b57600080fd5b9392505050565b610ad8806100a16000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80638bce6edd1161008c578063ab04505811610066578063ab045058146101d2578063de8d8b431461020b578063e23cc6c714610234578063ed7dd44e1461024757600080fd5b80638bce6edd146101865780638c64ea4a14610197578063a87d942c146101ca57600080fd5b806306e7cf18146100d4578063095cf5c6146100e957806316279055146100fc578063395701db146101215780635eebb986146101625780636e9960c314610175575b600080fd5b6100e76100e2366004610800565b61026a565b005b6100e76100f7366004610841565b6102f1565b61010e61010a366004610841565b3b90565b6040519081526020015b60405180910390f35b61014a61012f366004610841565b6004602052600090815260409020546001600160a01b031681565b6040516001600160a01b039091168152602001610118565b6100e76101703660046108b6565b61032a565b6000546001600160a01b031661014a565b6001546001600160a01b031661014a565b6101aa6101a53660046109a3565b610626565b604080516001600160a01b03938416815292909116602083015201610118565b60055461010e565b61010e6101e03660046109bc565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b61014a610219366004610841565b6003602052600090815260409020546001600160a01b031681565b6100e76102423660046109bc565b61065f565b61025a610255366004610841565b610684565b6040519015158152602001610118565b6000546001600160a01b0316331461028157600080fd5b600081116102c55760405162461bcd60e51b815260206004820152600c60248201526b1a5b9d985b1a59081cdd185d60a21b60448201526064015b60405180910390fd5b6001600160a01b0392831660009081526002602090815260408083209490951682529290925291902055565b6000546001600160a01b0316331461030857600080fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6101208801516000805460c08b015160e08c0151604051632a3c85d760e01b81523060048201526001600160a01b039384166024820152918316604483015282166064820152608481018b905260a481018a905291921690632a3c85d79060c4016020604051808303816000875af11580156103aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103ce9190610a0b565b90506001600160a01b03811661040b5760405162461bcd60e51b8152602060048201526002602482015261076360f41b60448201526064016102bc565b60006040518061014001604052808b6000600a811061042c5761042c6109f5565b602090810291909101516001600160a01b039081168352339183019190915260408d8101518216818401526060808f01518316908401526080808f015183169084015260a0808f015183169084015260c0808f015183169084015260e0808f015183169084015285821661010080850191909152600080548416610120909501949094528e01519051631e0c270560e01b81529394509192911690631e0c2705906104e59085908c908c908c908c908c90600401610a28565b6020604051808303816000875af1158015610504573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105289190610a0b565b90506001600160a01b0381166105655760405162461bcd60e51b8152602060048201526002602482015261073360f41b60448201526064016102bc565b60405163351a456f60e21b81526001600160a01b03828116600483015284169063d46915bc90602401600060405180830381600087803b1580156105a857600080fd5b505af11580156105bc573d6000803e3d6000fd5b505050506105ca81846106f9565b604080513081523360208201526001600160a01b03838116828401528516606082015290517f341f822f1e2a166284695f5b30ad55eb8e9932e740b7eed3ac9a1871f24a19a79181900360800190a15050505050505050505050565b6005818154811061063657600080fd5b6000918252602090912060029091020180546001909101546001600160a01b0391821692501682565b6000546001600160a01b0316331461067657600080fd5b61068082826106f9565b5050565b6001600160a01b0380821660009081526002602090815260408083206003835281842054909416835292905290812054600114806106f357506001600160a01b038083166000818152600460209081526040808320549094168252600281528382209282529190915220546001145b92915050565b6001600160a01b0391821660008181526002602081815260408084209587168085529582528084208390558484526003825280842080546001600160a01b0319908116881790915586855260048352818520805482168717905581518083019092529481529081019485526005805460018101825593525191027f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0810180549286169284169290921790915591517f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db19092018054929093169116179055565b6001600160a01b03811681146107ed57600080fd5b50565b80356107fb816107d8565b919050565b60008060006060848603121561081557600080fd5b8335610820816107d8565b92506020840135610830816107d8565b929592945050506040919091013590565b60006020828403121561085357600080fd5b813561085e816107d8565b9392505050565b634e487b7160e01b600052604160045260246000fd5b803560ff811681146107fb57600080fd5b803562ffffff811681146107fb57600080fd5b80356001600160801b03811681146107fb57600080fd5b600080600080600080600080610220898b0312156108d357600080fd5b89601f8a01126108e257600080fd5b60405161014080820182811067ffffffffffffffff8211171561090757610907610865565b6040528a01818c82111561091a57600080fd5b8b5b8281101561093b5761092d816107f0565b82526020918201910161091c565b509199503597505061016089013595506109586101808a0161087b565b94506109676101a08a0161087b565b93506109766101c08a0161087b565b92506109856101e08a0161088c565b91506109946102008a0161089f565b90509295985092959890939650565b6000602082840312156109b557600080fd5b5035919050565b600080604083850312156109cf57600080fd5b82356109da816107d8565b915060208301356109ea816107d8565b809150509250929050565b634e487b7160e01b600052603260045260246000fd5b600060208284031215610a1d57600080fd5b815161085e816107d8565b6101e08101818860005b600a811015610a5a5781516001600160a01b0316835260209283019290910190600101610a32565b50505060ff871661014083015260ff861661016083015260ff851661018083015262ffffff84166101a08301526001600160801b0383166101c083015297965050505050505056fea26469706673582212206bbb7619d0cf7fa131402eb38b1940a75c99948f50f8316831f7c5aaeff19c9a64736f6c634300080a0033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, _team common.Address) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend, _team)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// CheckActive is a free data retrieval call binding the contract method 0xed7dd44e.
//
// Solidity: function checkActive(address sORv) view returns(bool)
func (_Api *ApiCaller) CheckActive(opts *bind.CallOpts, sORv common.Address) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "checkActive", sORv)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckActive is a free data retrieval call binding the contract method 0xed7dd44e.
//
// Solidity: function checkActive(address sORv) view returns(bool)
func (_Api *ApiSession) CheckActive(sORv common.Address) (bool, error) {
	return _Api.Contract.CheckActive(&_Api.CallOpts, sORv)
}

// CheckActive is a free data retrieval call binding the contract method 0xed7dd44e.
//
// Solidity: function checkActive(address sORv) view returns(bool)
func (_Api *ApiCallerSession) CheckActive(sORv common.Address) (bool, error) {
	return _Api.Contract.CheckActive(&_Api.CallOpts, sORv)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_Api *ApiCaller) GetAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_Api *ApiSession) GetAdmin() (common.Address, error) {
	return _Api.Contract.GetAdmin(&_Api.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_Api *ApiCallerSession) GetAdmin() (common.Address, error) {
	return _Api.Contract.GetAdmin(&_Api.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Api *ApiCaller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Api *ApiSession) GetCount() (*big.Int, error) {
	return _Api.Contract.GetCount(&_Api.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Api *ApiCallerSession) GetCount() (*big.Int, error) {
	return _Api.Contract.GetCount(&_Api.CallOpts)
}

// GetStat is a free data retrieval call binding the contract method 0xab045058.
//
// Solidity: function getStat(address _strategy, address _vault) view returns(uint256)
func (_Api *ApiCaller) GetStat(opts *bind.CallOpts, _strategy common.Address, _vault common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStat", _strategy, _vault)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStat is a free data retrieval call binding the contract method 0xab045058.
//
// Solidity: function getStat(address _strategy, address _vault) view returns(uint256)
func (_Api *ApiSession) GetStat(_strategy common.Address, _vault common.Address) (*big.Int, error) {
	return _Api.Contract.GetStat(&_Api.CallOpts, _strategy, _vault)
}

// GetStat is a free data retrieval call binding the contract method 0xab045058.
//
// Solidity: function getStat(address _strategy, address _vault) view returns(uint256)
func (_Api *ApiCallerSession) GetStat(_strategy common.Address, _vault common.Address) (*big.Int, error) {
	return _Api.Contract.GetStat(&_Api.CallOpts, _strategy, _vault)
}

// GetTeam is a free data retrieval call binding the contract method 0x8bce6edd.
//
// Solidity: function getTeam() view returns(address)
func (_Api *ApiCaller) GetTeam(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getTeam")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTeam is a free data retrieval call binding the contract method 0x8bce6edd.
//
// Solidity: function getTeam() view returns(address)
func (_Api *ApiSession) GetTeam() (common.Address, error) {
	return _Api.Contract.GetTeam(&_Api.CallOpts)
}

// GetTeam is a free data retrieval call binding the contract method 0x8bce6edd.
//
// Solidity: function getTeam() view returns(address)
func (_Api *ApiCallerSession) GetTeam() (common.Address, error) {
	return _Api.Contract.GetTeam(&_Api.CallOpts)
}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(address addr) view returns(uint256)
func (_Api *ApiCaller) IsContract(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "isContract", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(address addr) view returns(uint256)
func (_Api *ApiSession) IsContract(addr common.Address) (*big.Int, error) {
	return _Api.Contract.IsContract(&_Api.CallOpts, addr)
}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(address addr) view returns(uint256)
func (_Api *ApiCallerSession) IsContract(addr common.Address) (*big.Int, error) {
	return _Api.Contract.IsContract(&_Api.CallOpts, addr)
}

// PairStrategy is a free data retrieval call binding the contract method 0x395701db.
//
// Solidity: function pairStrategy(address ) view returns(address)
func (_Api *ApiCaller) PairStrategy(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "pairStrategy", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairStrategy is a free data retrieval call binding the contract method 0x395701db.
//
// Solidity: function pairStrategy(address ) view returns(address)
func (_Api *ApiSession) PairStrategy(arg0 common.Address) (common.Address, error) {
	return _Api.Contract.PairStrategy(&_Api.CallOpts, arg0)
}

// PairStrategy is a free data retrieval call binding the contract method 0x395701db.
//
// Solidity: function pairStrategy(address ) view returns(address)
func (_Api *ApiCallerSession) PairStrategy(arg0 common.Address) (common.Address, error) {
	return _Api.Contract.PairStrategy(&_Api.CallOpts, arg0)
}

// PairVault is a free data retrieval call binding the contract method 0xde8d8b43.
//
// Solidity: function pairVault(address ) view returns(address)
func (_Api *ApiCaller) PairVault(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "pairVault", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairVault is a free data retrieval call binding the contract method 0xde8d8b43.
//
// Solidity: function pairVault(address ) view returns(address)
func (_Api *ApiSession) PairVault(arg0 common.Address) (common.Address, error) {
	return _Api.Contract.PairVault(&_Api.CallOpts, arg0)
}

// PairVault is a free data retrieval call binding the contract method 0xde8d8b43.
//
// Solidity: function pairVault(address ) view returns(address)
func (_Api *ApiCallerSession) PairVault(arg0 common.Address) (common.Address, error) {
	return _Api.Contract.PairVault(&_Api.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(address strategy, address vault)
func (_Api *ApiCaller) Vaults(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Strategy common.Address
	Vault    common.Address
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "vaults", arg0)

	outstruct := new(struct {
		Strategy common.Address
		Vault    common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Strategy = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Vault = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(address strategy, address vault)
func (_Api *ApiSession) Vaults(arg0 *big.Int) (struct {
	Strategy common.Address
	Vault    common.Address
}, error) {
	return _Api.Contract.Vaults(&_Api.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(address strategy, address vault)
func (_Api *ApiCallerSession) Vaults(arg0 *big.Int) (struct {
	Strategy common.Address
	Vault    common.Address
}, error) {
	return _Api.Contract.Vaults(&_Api.CallOpts, arg0)
}

// ChangeStat is a paid mutator transaction binding the contract method 0x06e7cf18.
//
// Solidity: function changeStat(address _strategy, address _vault, uint256 _stat) returns()
func (_Api *ApiTransactor) ChangeStat(opts *bind.TransactOpts, _strategy common.Address, _vault common.Address, _stat *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "changeStat", _strategy, _vault, _stat)
}

// ChangeStat is a paid mutator transaction binding the contract method 0x06e7cf18.
//
// Solidity: function changeStat(address _strategy, address _vault, uint256 _stat) returns()
func (_Api *ApiSession) ChangeStat(_strategy common.Address, _vault common.Address, _stat *big.Int) (*types.Transaction, error) {
	return _Api.Contract.ChangeStat(&_Api.TransactOpts, _strategy, _vault, _stat)
}

// ChangeStat is a paid mutator transaction binding the contract method 0x06e7cf18.
//
// Solidity: function changeStat(address _strategy, address _vault, uint256 _stat) returns()
func (_Api *ApiTransactorSession) ChangeStat(_strategy common.Address, _vault common.Address, _stat *big.Int) (*types.Transaction, error) {
	return _Api.Contract.ChangeStat(&_Api.TransactOpts, _strategy, _vault, _stat)
}

// CreateVault is a paid mutator transaction binding the contract method 0x5eebb986.
//
// Solidity: function createVault(address[10] contracts, uint256 _vaultCap, uint256 _individualCap, uint8 _uniPortion, uint8 _compPortion, uint8 _protocolFee, uint24 _feetier, uint128 _quoteAmount) returns()
func (_Api *ApiTransactor) CreateVault(opts *bind.TransactOpts, contracts [10]common.Address, _vaultCap *big.Int, _individualCap *big.Int, _uniPortion uint8, _compPortion uint8, _protocolFee uint8, _feetier *big.Int, _quoteAmount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "createVault", contracts, _vaultCap, _individualCap, _uniPortion, _compPortion, _protocolFee, _feetier, _quoteAmount)
}

// CreateVault is a paid mutator transaction binding the contract method 0x5eebb986.
//
// Solidity: function createVault(address[10] contracts, uint256 _vaultCap, uint256 _individualCap, uint8 _uniPortion, uint8 _compPortion, uint8 _protocolFee, uint24 _feetier, uint128 _quoteAmount) returns()
func (_Api *ApiSession) CreateVault(contracts [10]common.Address, _vaultCap *big.Int, _individualCap *big.Int, _uniPortion uint8, _compPortion uint8, _protocolFee uint8, _feetier *big.Int, _quoteAmount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.CreateVault(&_Api.TransactOpts, contracts, _vaultCap, _individualCap, _uniPortion, _compPortion, _protocolFee, _feetier, _quoteAmount)
}

// CreateVault is a paid mutator transaction binding the contract method 0x5eebb986.
//
// Solidity: function createVault(address[10] contracts, uint256 _vaultCap, uint256 _individualCap, uint8 _uniPortion, uint8 _compPortion, uint8 _protocolFee, uint24 _feetier, uint128 _quoteAmount) returns()
func (_Api *ApiTransactorSession) CreateVault(contracts [10]common.Address, _vaultCap *big.Int, _individualCap *big.Int, _uniPortion uint8, _compPortion uint8, _protocolFee uint8, _feetier *big.Int, _quoteAmount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.CreateVault(&_Api.TransactOpts, contracts, _vaultCap, _individualCap, _uniPortion, _compPortion, _protocolFee, _feetier, _quoteAmount)
}

// Repair is a paid mutator transaction binding the contract method 0xe23cc6c7.
//
// Solidity: function repair(address _strategy, address _vault) returns()
func (_Api *ApiTransactor) Repair(opts *bind.TransactOpts, _strategy common.Address, _vault common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "repair", _strategy, _vault)
}

// Repair is a paid mutator transaction binding the contract method 0xe23cc6c7.
//
// Solidity: function repair(address _strategy, address _vault) returns()
func (_Api *ApiSession) Repair(_strategy common.Address, _vault common.Address) (*types.Transaction, error) {
	return _Api.Contract.Repair(&_Api.TransactOpts, _strategy, _vault)
}

// Repair is a paid mutator transaction binding the contract method 0xe23cc6c7.
//
// Solidity: function repair(address _strategy, address _vault) returns()
func (_Api *ApiTransactorSession) Repair(_strategy common.Address, _vault common.Address) (*types.Transaction, error) {
	return _Api.Contract.Repair(&_Api.TransactOpts, _strategy, _vault)
}

// SetTeam is a paid mutator transaction binding the contract method 0x095cf5c6.
//
// Solidity: function setTeam(address _team) returns()
func (_Api *ApiTransactor) SetTeam(opts *bind.TransactOpts, _team common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setTeam", _team)
}

// SetTeam is a paid mutator transaction binding the contract method 0x095cf5c6.
//
// Solidity: function setTeam(address _team) returns()
func (_Api *ApiSession) SetTeam(_team common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetTeam(&_Api.TransactOpts, _team)
}

// SetTeam is a paid mutator transaction binding the contract method 0x095cf5c6.
//
// Solidity: function setTeam(address _team) returns()
func (_Api *ApiTransactorSession) SetTeam(_team common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetTeam(&_Api.TransactOpts, _team)
}

// ApiCreatedVaultIterator is returned from FilterCreatedVault and is used to iterate over the raw logs and unpacked data for CreatedVault events raised by the Api contract.
type ApiCreatedVaultIterator struct {
	Event *ApiCreatedVault // Event containing the contract specifics and raw log

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
func (it *ApiCreatedVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiCreatedVault)
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
		it.Event = new(ApiCreatedVault)
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
func (it *ApiCreatedVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiCreatedVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiCreatedVault represents a CreatedVault event raised by the Api contract.
type ApiCreatedVault struct {
	This     common.Address
	Sender   common.Address
	Strategy common.Address
	Vault    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCreatedVault is a free log retrieval operation binding the contract event 0x341f822f1e2a166284695f5b30ad55eb8e9932e740b7eed3ac9a1871f24a19a7.
//
// Solidity: event CreatedVault(address _this, address sender, address strategy, address vault)
func (_Api *ApiFilterer) FilterCreatedVault(opts *bind.FilterOpts) (*ApiCreatedVaultIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "CreatedVault")
	if err != nil {
		return nil, err
	}
	return &ApiCreatedVaultIterator{contract: _Api.contract, event: "CreatedVault", logs: logs, sub: sub}, nil
}

// WatchCreatedVault is a free log subscription operation binding the contract event 0x341f822f1e2a166284695f5b30ad55eb8e9932e740b7eed3ac9a1871f24a19a7.
//
// Solidity: event CreatedVault(address _this, address sender, address strategy, address vault)
func (_Api *ApiFilterer) WatchCreatedVault(opts *bind.WatchOpts, sink chan<- *ApiCreatedVault) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "CreatedVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiCreatedVault)
				if err := _Api.contract.UnpackLog(event, "CreatedVault", log); err != nil {
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

// ParseCreatedVault is a log parse operation binding the contract event 0x341f822f1e2a166284695f5b30ad55eb8e9932e740b7eed3ac9a1871f24a19a7.
//
// Solidity: event CreatedVault(address _this, address sender, address strategy, address vault)
func (_Api *ApiFilterer) ParseCreatedVault(log types.Log) (*ApiCreatedVault, error) {
	event := new(ApiCreatedVault)
	if err := _Api.contract.UnpackLog(event, "CreatedVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ApiABI is the input ABI used to generate the binding from.
const ApiABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee1\",\"type\":\"uint256\"}],\"name\":\"FlashCallback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Owed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Owed\",\"type\":\"uint256\"}],\"name\":\"MintCallback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"}],\"name\":\"SwapCallback\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pay0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pay1\",\"type\":\"uint256\"}],\"name\":\"flash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"amountSpecified\",\"type\":\"int256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"swap0ForExact1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"swap1ForExact0\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"swapExact0For1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"swapExact1For0\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"swapToHigherSqrtPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"swapToLowerSqrtPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3FlashCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Owed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Owed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ApiBin is the compiled bytecode used for deploying new contracts.
var ApiBin = "0x608060405234801561001057600080fd5b5061136f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063bac7bf7811610071578063bac7bf7814610250578063d34879971461028c578063e2be910914610308578063e9cbafb014610344578063f603482c146103c0578063fa461e33146103fc576100b4565b8063034b0f8f146100b95780632ec20bf9146101035780635b71a46e1461013b5780636dfc0ddb146101885780637b4f5327146101c45780639e77b80514610218575b600080fd5b610101600480360360c08110156100cf57600080fd5b506001600160a01b03813581169160208101359091169060408101359060608101359060808101359060a00135610478565b005b6101016004803603606081101561011957600080fd5b506001600160a01b038135811691602081013582169160409091013516610599565b61016f6004803603606081101561015157600080fd5b506001600160a01b03813516906020810135151590604001356106d5565b6040805192835260208301919091528051918290030190f35b6101016004803603608081101561019e57600080fd5b506001600160a01b0381358116916020810135916040820135811691606001351661083f565b610101600480360360a08110156101da57600080fd5b5080356001600160a01b03908116916020810135909116906040810135600290810b91606081013590910b90608001356001600160801b031661097d565b6101016004803603606081101561022e57600080fd5b506001600160a01b038135811691602081013582169160409091013516610ab7565b6101016004803603608081101561026657600080fd5b506001600160a01b03813581169160208101359160408201358116916060013516610b72565b610101600480360360608110156102a257600080fd5b8135916020810135918101906060810160408201356401000000008111156102c957600080fd5b8201836020820111156102db57600080fd5b803590602001918460018302840111640100000000831117156102fd57600080fd5b509092509050610c31565b6101016004803603608081101561031e57600080fd5b506001600160a01b03813581169160208101359160408201358116916060013516610e6b565b6101016004803603606081101561035a57600080fd5b81359160208101359181019060608101604082013564010000000081111561038157600080fd5b82018360208201111561039357600080fd5b803590602001918460018302840111640100000000831117156103b557600080fd5b509092509050610e86565b610101600480360360808110156103d657600080fd5b506001600160a01b038135811691602081013591604082013581169160600135166110db565b6101016004803603606081101561041257600080fd5b81359160208101359181019060608101604082013564010000000081111561043957600080fd5b82018360208201111561044b57600080fd5b8035906020019184600183028401116401000000008311171561046d57600080fd5b5090925090506110f6565b856001600160a01b031663490e6cbc86868633878760405160200180846001600160a01b0316815260200183815260200182815260200193505050506040516020818303038152906040526040518563ffffffff1660e01b815260040180856001600160a01b0316815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561052a578181015183820152602001610512565b50505050905090810190601f1680156105575780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b15801561057957600080fd5b505af115801561058d573d6000803e3d6000fd5b50505050505050505050565b826001600160a01b031663128acb088260016001600160ff1b03863360405160200180826001600160a01b031681526020019150506040516020818303038152906040526040518663ffffffff1660e01b815260040180866001600160a01b031681526020018515158152602001848152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561065557818101518382015260200161063d565b50505050905090810190601f1680156106825780820380516001836020036101000a031916815260200191505b5096505050505050506040805180830381600087803b1580156106a457600080fd5b505af11580156106b8573d6000803e3d6000fd5b505050506040513d60408110156106ce57600080fd5b5050505050565b600080846001600160a01b031663128acb08338686886107095773fffd8963efd1fc6a506488495d951d5263988d25610710565b6401000276a45b3360405160200180826001600160a01b031681526020019150506040516020818303038152906040526040518663ffffffff1660e01b815260040180866001600160a01b031681526020018515158152602001848152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156107b1578181015183820152602001610799565b50505050905090810190601f1680156107de5780820380516001836020036101000a031916815260200191505b5096505050505050506040805180830381600087803b15801561080057600080fd5b505af1158015610814573d6000803e3d6000fd5b505050506040513d604081101561082a57600080fd5b50805160209091015190969095509350505050565b836001600160a01b031663128acb0883600161085a87611323565b853360405160200180826001600160a01b031681526020019150506040516020818303038152906040526040518663ffffffff1660e01b815260040180866001600160a01b031681526020018515158152602001848152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156108fc5781810151838201526020016108e4565b50505050905090810190601f1680156109295780820380516001836020036101000a031916815260200191505b5096505050505050506040805180830381600087803b15801561094b57600080fd5b505af115801561095f573d6000803e3d6000fd5b505050506040513d604081101561097557600080fd5b505050505050565b846001600160a01b0316633c8a7d8d858585853360405160200180826001600160a01b031681526020019150506040516020818303038152906040526040518663ffffffff1660e01b815260040180866001600160a01b031681526020018560020b81526020018460020b8152602001836001600160801b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610a35578181015183820152602001610a1d565b50505050905090810190601f168015610a625780820380516001836020036101000a031916815260200191505b5096505050505050506040805180830381600087803b158015610a8457600080fd5b505af1158015610a98573d6000803e3d6000fd5b505050506040513d6040811015610aae57600080fd5b50505050505050565b826001600160a01b031663128acb088260006001600160ff1b03863360405160200180826001600160a01b031681526020019150506040516020818303038152906040526040518663ffffffff1660e01b815260040180866001600160a01b031681526020018515158152602001848152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360008381101561065557818101518382015260200161063d565b836001600160a01b031663128acb08836001610b8d87611323565b600003853360405160200180826001600160a01b031681526020019150506040516020818303038152906040526040518663ffffffff1660e01b815260040180866001600160a01b031681526020018515158152602001848152602001836001600160a01b031681526020018060200182810382528381815181526020019150805190602001908083836000838110156108fc5781810151838201526020016108e4565b600082826020811015610c4357600080fd5b506040805187815260208101879052815192356001600160a01b031693507fa0968be00566083701c9ef671c169d7fb05ac8907de4ca17185de74ccbb694d4929081900390910190a18415610d7d57336001600160a01b0316630dfe16816040518163ffffffff1660e01b815260040160206040518083038186803b158015610ccb57600080fd5b505afa158015610cdf573d6000803e3d6000fd5b505050506040513d6020811015610cf557600080fd5b5051604080516323b872dd60e01b81526001600160a01b03848116600483015233602483015260448201899052915191909216916323b872dd9160648083019260209291908290030181600087803b158015610d5057600080fd5b505af1158015610d64573d6000803e3d6000fd5b505050506040513d6020811015610d7a57600080fd5b50505b83156106ce57336001600160a01b031663d21220a76040518163ffffffff1660e01b815260040160206040518083038186803b158015610dbc57600080fd5b505afa158015610dd0573d6000803e3d6000fd5b505050506040513d6020811015610de657600080fd5b5051604080516323b872dd60e01b81526001600160a01b03848116600483015233602483015260448201889052915191909216916323b872dd9160648083019260209291908290030181600087803b158015610e4157600080fd5b505af1158015610e55573d6000803e3d6000fd5b505050506040513d6020811015610aae57600080fd5b836001600160a01b031663128acb0883600061085a87611323565b604080518581526020810185905281517f2b0391b4fa408cfe47abd1977d72985695b2e5ebd3175f55be25f2c68c5df21b929181900390910190a1600080600084846060811015610ed657600080fd5b506001600160a01b038135169350602081013592506040013590508115610fe257336001600160a01b0316630dfe16816040518163ffffffff1660e01b815260040160206040518083038186803b158015610f3057600080fd5b505afa158015610f44573d6000803e3d6000fd5b505050506040513d6020811015610f5a57600080fd5b5051604080516323b872dd60e01b81526001600160a01b03868116600483015233602483015260448201869052915191909216916323b872dd9160648083019260209291908290030181600087803b158015610fb557600080fd5b505af1158015610fc9573d6000803e3d6000fd5b505050506040513d6020811015610fdf57600080fd5b50505b8015610aae57336001600160a01b031663d21220a76040518163ffffffff1660e01b815260040160206040518083038186803b15801561102157600080fd5b505afa158015611035573d6000803e3d6000fd5b505050506040513d602081101561104b57600080fd5b5051604080516323b872dd60e01b81526001600160a01b03868116600483015233602483015260448201859052915191909216916323b872dd9160648083019260209291908290030181600087803b1580156110a657600080fd5b505af11580156110ba573d6000803e3d6000fd5b505050506040513d60208110156110d057600080fd5b505050505050505050565b836001600160a01b031663128acb08836000610b8d87611323565b60008282602081101561110857600080fd5b506040805187815260208101879052815192356001600160a01b031693507fd48241df4a75e663b29e55f9506b31f77ed0f48cfe7e7612d1163144995dc1ca929081900390910190a1600085131561124a57336001600160a01b0316630dfe16816040518163ffffffff1660e01b815260040160206040518083038186803b15801561119357600080fd5b505afa1580156111a7573d6000803e3d6000fd5b505050506040513d60208110156111bd57600080fd5b5051604080516323b872dd60e01b81526001600160a01b03848116600483015233602483015260448201899052915191909216916323b872dd9160648083019260209291908290030181600087803b15801561121857600080fd5b505af115801561122c573d6000803e3d6000fd5b505050506040513d602081101561124257600080fd5b506106ce9050565b600084131561131157336001600160a01b031663d21220a76040518163ffffffff1660e01b815260040160206040518083038186803b15801561128c57600080fd5b505afa1580156112a0573d6000803e3d6000fd5b505050506040513d60208110156112b657600080fd5b5051604080516323b872dd60e01b81526001600160a01b03848116600483015233602483015260448201889052915191909216916323b872dd9160648083019260209291908290030181600087803b15801561121857600080fd5b8415801561131d575083155b6106ce57fe5b6000600160ff1b821061133557600080fd5b509056fea26469706673582212206962888a7e30f4de31e2d0d83e7d23709f8f344a81262c05375734be1cf492f464736f6c63430007060033"

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ApiBin), backend)
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

// Flash is a paid mutator transaction binding the contract method 0x034b0f8f.
//
// Solidity: function flash(address pool, address recipient, uint256 amount0, uint256 amount1, uint256 pay0, uint256 pay1) returns()
func (_Api *ApiTransactor) Flash(opts *bind.TransactOpts, pool common.Address, recipient common.Address, amount0 *big.Int, amount1 *big.Int, pay0 *big.Int, pay1 *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "flash", pool, recipient, amount0, amount1, pay0, pay1)
}

// Flash is a paid mutator transaction binding the contract method 0x034b0f8f.
//
// Solidity: function flash(address pool, address recipient, uint256 amount0, uint256 amount1, uint256 pay0, uint256 pay1) returns()
func (_Api *ApiSession) Flash(pool common.Address, recipient common.Address, amount0 *big.Int, amount1 *big.Int, pay0 *big.Int, pay1 *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Flash(&_Api.TransactOpts, pool, recipient, amount0, amount1, pay0, pay1)
}

// Flash is a paid mutator transaction binding the contract method 0x034b0f8f.
//
// Solidity: function flash(address pool, address recipient, uint256 amount0, uint256 amount1, uint256 pay0, uint256 pay1) returns()
func (_Api *ApiTransactorSession) Flash(pool common.Address, recipient common.Address, amount0 *big.Int, amount1 *big.Int, pay0 *big.Int, pay1 *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Flash(&_Api.TransactOpts, pool, recipient, amount0, amount1, pay0, pay1)
}

// Mint is a paid mutator transaction binding the contract method 0x7b4f5327.
//
// Solidity: function mint(address pool, address recipient, int24 tickLower, int24 tickUpper, uint128 amount) returns()
func (_Api *ApiTransactor) Mint(opts *bind.TransactOpts, pool common.Address, recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "mint", pool, recipient, tickLower, tickUpper, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x7b4f5327.
//
// Solidity: function mint(address pool, address recipient, int24 tickLower, int24 tickUpper, uint128 amount) returns()
func (_Api *ApiSession) Mint(pool common.Address, recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Mint(&_Api.TransactOpts, pool, recipient, tickLower, tickUpper, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x7b4f5327.
//
// Solidity: function mint(address pool, address recipient, int24 tickLower, int24 tickUpper, uint128 amount) returns()
func (_Api *ApiTransactorSession) Mint(pool common.Address, recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Mint(&_Api.TransactOpts, pool, recipient, tickLower, tickUpper, amount)
}

// Swap is a paid mutator transaction binding the contract method 0x5b71a46e.
//
// Solidity: function swap(address pool, bool zeroForOne, int256 amountSpecified) returns(int256 amount0, int256 amount1)
func (_Api *ApiTransactor) Swap(opts *bind.TransactOpts, pool common.Address, zeroForOne bool, amountSpecified *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "swap", pool, zeroForOne, amountSpecified)
}

// Swap is a paid mutator transaction binding the contract method 0x5b71a46e.
//
// Solidity: function swap(address pool, bool zeroForOne, int256 amountSpecified) returns(int256 amount0, int256 amount1)
func (_Api *ApiSession) Swap(pool common.Address, zeroForOne bool, amountSpecified *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Swap(&_Api.TransactOpts, pool, zeroForOne, amountSpecified)
}

// Swap is a paid mutator transaction binding the contract method 0x5b71a46e.
//
// Solidity: function swap(address pool, bool zeroForOne, int256 amountSpecified) returns(int256 amount0, int256 amount1)
func (_Api *ApiTransactorSession) Swap(pool common.Address, zeroForOne bool, amountSpecified *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Swap(&_Api.TransactOpts, pool, zeroForOne, amountSpecified)
}

// Swap0ForExact1 is a paid mutator transaction binding the contract method 0xbac7bf78.
//
// Solidity: function swap0ForExact1(address pool, uint256 amount1Out, address recipient, uint160 sqrtPriceLimitX96) returns()
func (_Api *ApiTransactor) Swap0ForExact1(opts *bind.TransactOpts, pool common.Address, amount1Out *big.Int, recipient common.Address, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "swap0ForExact1", pool, amount1Out, recipient, sqrtPriceLimitX96)
}

// Swap0ForExact1 is a paid mutator transaction binding the contract method 0xbac7bf78.
//
// Solidity: function swap0ForExact1(address pool, uint256 amount1Out, address recipient, uint160 sqrtPriceLimitX96) returns()
func (_Api *ApiSession) Swap0ForExact1(pool common.Address, amount1Out *big.Int, recipient common.Address, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Swap0ForExact1(&_Api.TransactOpts, pool, amount1Out, recipient, sqrtPriceLimitX96)
}

// Swap0ForExact1 is a paid mutator transaction binding the contract method 0xbac7bf78.
//
// Solidity: function swap0ForExact1(address pool, uint256 amount1Out, address recipient, uint160 sqrtPriceLimitX96) returns()
func (_Api *ApiTransactorSession) Swap0ForExact1(pool common.Address, amount1Out *big.Int, recipient common.Address, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Swap0ForExact1(&_Api.TransactOpts, pool, amount1Out, recipient, sqrtPriceLimitX96)
}

// Swap1ForExact0 is a paid mutator transaction binding the contract method 0xf603482c.
//
// Solidity: function swap1ForExact0(address pool, uint256 amount0Out, address recipient, uint160 sqrtPriceLimitX96) returns()
func (_Api *ApiTransactor) Swap1ForExact0(opts *bind.TransactOpts, pool common.Address, amount0Out *big.Int, recipient common.Address, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "swap1ForExact0", pool, amount0Out, recipient, sqrtPriceLimitX96)
}

// Swap1ForExact0 is a paid mutator transaction binding the contract method 0xf603482c.
//
// Solidity: function swap1ForExact0(address pool, uint256 amount0Out, address recipient, uint160 sqrtPriceLimitX96) returns()
func (_Api *ApiSession) Swap1ForExact0(pool common.Address, amount0Out *big.Int, recipient common.Address, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Swap1ForExact0(&_Api.TransactOpts, pool, amount0Out, recipient, sqrtPriceLimitX96)
}

// Swap1ForExact0 is a paid mutator transaction binding the contract method 0xf603482c.
//
// Solidity: function swap1ForExact0(address pool, uint256 amount0Out, address recipient, uint160 sqrtPriceLimitX96) returns()
func (_Api *ApiTransactorSession) Swap1ForExact0(pool common.Address, amount0Out *big.Int, recipient common.Address, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Swap1ForExact0(&_Api.TransactOpts, pool, amount0Out, recipient, sqrtPriceLimitX96)
}

// SwapExact0For1 is a paid mutator transaction binding the contract method 0x6dfc0ddb.
//
// Solidity: function swapExact0For1(address pool, uint256 amount0In, address recipient, uint160 sqrtPriceLimitX96) returns()
func (_Api *ApiTransactor) SwapExact0For1(opts *bind.TransactOpts, pool common.Address, amount0In *big.Int, recipient common.Address, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "swapExact0For1", pool, amount0In, recipient, sqrtPriceLimitX96)
}

// SwapExact0For1 is a paid mutator transaction binding the contract method 0x6dfc0ddb.
//
// Solidity: function swapExact0For1(address pool, uint256 amount0In, address recipient, uint160 sqrtPriceLimitX96) returns()
func (_Api *ApiSession) SwapExact0For1(pool common.Address, amount0In *big.Int, recipient common.Address, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SwapExact0For1(&_Api.TransactOpts, pool, amount0In, recipient, sqrtPriceLimitX96)
}

// SwapExact0For1 is a paid mutator transaction binding the contract method 0x6dfc0ddb.
//
// Solidity: function swapExact0For1(address pool, uint256 amount0In, address recipient, uint160 sqrtPriceLimitX96) returns()
func (_Api *ApiTransactorSession) SwapExact0For1(pool common.Address, amount0In *big.Int, recipient common.Address, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SwapExact0For1(&_Api.TransactOpts, pool, amount0In, recipient, sqrtPriceLimitX96)
}

// SwapExact1For0 is a paid mutator transaction binding the contract method 0xe2be9109.
//
// Solidity: function swapExact1For0(address pool, uint256 amount1In, address recipient, uint160 sqrtPriceLimitX96) returns()
func (_Api *ApiTransactor) SwapExact1For0(opts *bind.TransactOpts, pool common.Address, amount1In *big.Int, recipient common.Address, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "swapExact1For0", pool, amount1In, recipient, sqrtPriceLimitX96)
}

// SwapExact1For0 is a paid mutator transaction binding the contract method 0xe2be9109.
//
// Solidity: function swapExact1For0(address pool, uint256 amount1In, address recipient, uint160 sqrtPriceLimitX96) returns()
func (_Api *ApiSession) SwapExact1For0(pool common.Address, amount1In *big.Int, recipient common.Address, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SwapExact1For0(&_Api.TransactOpts, pool, amount1In, recipient, sqrtPriceLimitX96)
}

// SwapExact1For0 is a paid mutator transaction binding the contract method 0xe2be9109.
//
// Solidity: function swapExact1For0(address pool, uint256 amount1In, address recipient, uint160 sqrtPriceLimitX96) returns()
func (_Api *ApiTransactorSession) SwapExact1For0(pool common.Address, amount1In *big.Int, recipient common.Address, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SwapExact1For0(&_Api.TransactOpts, pool, amount1In, recipient, sqrtPriceLimitX96)
}

// SwapToHigherSqrtPrice is a paid mutator transaction binding the contract method 0x9e77b805.
//
// Solidity: function swapToHigherSqrtPrice(address pool, uint160 sqrtPriceX96, address recipient) returns()
func (_Api *ApiTransactor) SwapToHigherSqrtPrice(opts *bind.TransactOpts, pool common.Address, sqrtPriceX96 *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "swapToHigherSqrtPrice", pool, sqrtPriceX96, recipient)
}

// SwapToHigherSqrtPrice is a paid mutator transaction binding the contract method 0x9e77b805.
//
// Solidity: function swapToHigherSqrtPrice(address pool, uint160 sqrtPriceX96, address recipient) returns()
func (_Api *ApiSession) SwapToHigherSqrtPrice(pool common.Address, sqrtPriceX96 *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Api.Contract.SwapToHigherSqrtPrice(&_Api.TransactOpts, pool, sqrtPriceX96, recipient)
}

// SwapToHigherSqrtPrice is a paid mutator transaction binding the contract method 0x9e77b805.
//
// Solidity: function swapToHigherSqrtPrice(address pool, uint160 sqrtPriceX96, address recipient) returns()
func (_Api *ApiTransactorSession) SwapToHigherSqrtPrice(pool common.Address, sqrtPriceX96 *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Api.Contract.SwapToHigherSqrtPrice(&_Api.TransactOpts, pool, sqrtPriceX96, recipient)
}

// SwapToLowerSqrtPrice is a paid mutator transaction binding the contract method 0x2ec20bf9.
//
// Solidity: function swapToLowerSqrtPrice(address pool, uint160 sqrtPriceX96, address recipient) returns()
func (_Api *ApiTransactor) SwapToLowerSqrtPrice(opts *bind.TransactOpts, pool common.Address, sqrtPriceX96 *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "swapToLowerSqrtPrice", pool, sqrtPriceX96, recipient)
}

// SwapToLowerSqrtPrice is a paid mutator transaction binding the contract method 0x2ec20bf9.
//
// Solidity: function swapToLowerSqrtPrice(address pool, uint160 sqrtPriceX96, address recipient) returns()
func (_Api *ApiSession) SwapToLowerSqrtPrice(pool common.Address, sqrtPriceX96 *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Api.Contract.SwapToLowerSqrtPrice(&_Api.TransactOpts, pool, sqrtPriceX96, recipient)
}

// SwapToLowerSqrtPrice is a paid mutator transaction binding the contract method 0x2ec20bf9.
//
// Solidity: function swapToLowerSqrtPrice(address pool, uint160 sqrtPriceX96, address recipient) returns()
func (_Api *ApiTransactorSession) SwapToLowerSqrtPrice(pool common.Address, sqrtPriceX96 *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Api.Contract.SwapToLowerSqrtPrice(&_Api.TransactOpts, pool, sqrtPriceX96, recipient)
}

// UniswapV3FlashCallback is a paid mutator transaction binding the contract method 0xe9cbafb0.
//
// Solidity: function uniswapV3FlashCallback(uint256 fee0, uint256 fee1, bytes data) returns()
func (_Api *ApiTransactor) UniswapV3FlashCallback(opts *bind.TransactOpts, fee0 *big.Int, fee1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "uniswapV3FlashCallback", fee0, fee1, data)
}

// UniswapV3FlashCallback is a paid mutator transaction binding the contract method 0xe9cbafb0.
//
// Solidity: function uniswapV3FlashCallback(uint256 fee0, uint256 fee1, bytes data) returns()
func (_Api *ApiSession) UniswapV3FlashCallback(fee0 *big.Int, fee1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.Contract.UniswapV3FlashCallback(&_Api.TransactOpts, fee0, fee1, data)
}

// UniswapV3FlashCallback is a paid mutator transaction binding the contract method 0xe9cbafb0.
//
// Solidity: function uniswapV3FlashCallback(uint256 fee0, uint256 fee1, bytes data) returns()
func (_Api *ApiTransactorSession) UniswapV3FlashCallback(fee0 *big.Int, fee1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.Contract.UniswapV3FlashCallback(&_Api.TransactOpts, fee0, fee1, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes data) returns()
func (_Api *ApiTransactor) UniswapV3MintCallback(opts *bind.TransactOpts, amount0Owed *big.Int, amount1Owed *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "uniswapV3MintCallback", amount0Owed, amount1Owed, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes data) returns()
func (_Api *ApiSession) UniswapV3MintCallback(amount0Owed *big.Int, amount1Owed *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.Contract.UniswapV3MintCallback(&_Api.TransactOpts, amount0Owed, amount1Owed, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes data) returns()
func (_Api *ApiTransactorSession) UniswapV3MintCallback(amount0Owed *big.Int, amount1Owed *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.Contract.UniswapV3MintCallback(&_Api.TransactOpts, amount0Owed, amount1Owed, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_Api *ApiTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_Api *ApiSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.Contract.UniswapV3SwapCallback(&_Api.TransactOpts, amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_Api *ApiTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.Contract.UniswapV3SwapCallback(&_Api.TransactOpts, amount0Delta, amount1Delta, data)
}

// ApiFlashCallbackIterator is returned from FilterFlashCallback and is used to iterate over the raw logs and unpacked data for FlashCallback events raised by the Api contract.
type ApiFlashCallbackIterator struct {
	Event *ApiFlashCallback // Event containing the contract specifics and raw log

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
func (it *ApiFlashCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiFlashCallback)
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
		it.Event = new(ApiFlashCallback)
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
func (it *ApiFlashCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiFlashCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiFlashCallback represents a FlashCallback event raised by the Api contract.
type ApiFlashCallback struct {
	Fee0 *big.Int
	Fee1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFlashCallback is a free log retrieval operation binding the contract event 0x2b0391b4fa408cfe47abd1977d72985695b2e5ebd3175f55be25f2c68c5df21b.
//
// Solidity: event FlashCallback(uint256 fee0, uint256 fee1)
func (_Api *ApiFilterer) FilterFlashCallback(opts *bind.FilterOpts) (*ApiFlashCallbackIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "FlashCallback")
	if err != nil {
		return nil, err
	}
	return &ApiFlashCallbackIterator{contract: _Api.contract, event: "FlashCallback", logs: logs, sub: sub}, nil
}

// WatchFlashCallback is a free log subscription operation binding the contract event 0x2b0391b4fa408cfe47abd1977d72985695b2e5ebd3175f55be25f2c68c5df21b.
//
// Solidity: event FlashCallback(uint256 fee0, uint256 fee1)
func (_Api *ApiFilterer) WatchFlashCallback(opts *bind.WatchOpts, sink chan<- *ApiFlashCallback) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "FlashCallback")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiFlashCallback)
				if err := _Api.contract.UnpackLog(event, "FlashCallback", log); err != nil {
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

// ParseFlashCallback is a log parse operation binding the contract event 0x2b0391b4fa408cfe47abd1977d72985695b2e5ebd3175f55be25f2c68c5df21b.
//
// Solidity: event FlashCallback(uint256 fee0, uint256 fee1)
func (_Api *ApiFilterer) ParseFlashCallback(log types.Log) (*ApiFlashCallback, error) {
	event := new(ApiFlashCallback)
	if err := _Api.contract.UnpackLog(event, "FlashCallback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiMintCallbackIterator is returned from FilterMintCallback and is used to iterate over the raw logs and unpacked data for MintCallback events raised by the Api contract.
type ApiMintCallbackIterator struct {
	Event *ApiMintCallback // Event containing the contract specifics and raw log

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
func (it *ApiMintCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiMintCallback)
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
		it.Event = new(ApiMintCallback)
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
func (it *ApiMintCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiMintCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiMintCallback represents a MintCallback event raised by the Api contract.
type ApiMintCallback struct {
	Amount0Owed *big.Int
	Amount1Owed *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMintCallback is a free log retrieval operation binding the contract event 0xa0968be00566083701c9ef671c169d7fb05ac8907de4ca17185de74ccbb694d4.
//
// Solidity: event MintCallback(uint256 amount0Owed, uint256 amount1Owed)
func (_Api *ApiFilterer) FilterMintCallback(opts *bind.FilterOpts) (*ApiMintCallbackIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "MintCallback")
	if err != nil {
		return nil, err
	}
	return &ApiMintCallbackIterator{contract: _Api.contract, event: "MintCallback", logs: logs, sub: sub}, nil
}

// WatchMintCallback is a free log subscription operation binding the contract event 0xa0968be00566083701c9ef671c169d7fb05ac8907de4ca17185de74ccbb694d4.
//
// Solidity: event MintCallback(uint256 amount0Owed, uint256 amount1Owed)
func (_Api *ApiFilterer) WatchMintCallback(opts *bind.WatchOpts, sink chan<- *ApiMintCallback) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "MintCallback")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiMintCallback)
				if err := _Api.contract.UnpackLog(event, "MintCallback", log); err != nil {
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

// ParseMintCallback is a log parse operation binding the contract event 0xa0968be00566083701c9ef671c169d7fb05ac8907de4ca17185de74ccbb694d4.
//
// Solidity: event MintCallback(uint256 amount0Owed, uint256 amount1Owed)
func (_Api *ApiFilterer) ParseMintCallback(log types.Log) (*ApiMintCallback, error) {
	event := new(ApiMintCallback)
	if err := _Api.contract.UnpackLog(event, "MintCallback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiSwapCallbackIterator is returned from FilterSwapCallback and is used to iterate over the raw logs and unpacked data for SwapCallback events raised by the Api contract.
type ApiSwapCallbackIterator struct {
	Event *ApiSwapCallback // Event containing the contract specifics and raw log

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
func (it *ApiSwapCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiSwapCallback)
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
		it.Event = new(ApiSwapCallback)
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
func (it *ApiSwapCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiSwapCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiSwapCallback represents a SwapCallback event raised by the Api contract.
type ApiSwapCallback struct {
	Amount0Delta *big.Int
	Amount1Delta *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwapCallback is a free log retrieval operation binding the contract event 0xd48241df4a75e663b29e55f9506b31f77ed0f48cfe7e7612d1163144995dc1ca.
//
// Solidity: event SwapCallback(int256 amount0Delta, int256 amount1Delta)
func (_Api *ApiFilterer) FilterSwapCallback(opts *bind.FilterOpts) (*ApiSwapCallbackIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "SwapCallback")
	if err != nil {
		return nil, err
	}
	return &ApiSwapCallbackIterator{contract: _Api.contract, event: "SwapCallback", logs: logs, sub: sub}, nil
}

// WatchSwapCallback is a free log subscription operation binding the contract event 0xd48241df4a75e663b29e55f9506b31f77ed0f48cfe7e7612d1163144995dc1ca.
//
// Solidity: event SwapCallback(int256 amount0Delta, int256 amount1Delta)
func (_Api *ApiFilterer) WatchSwapCallback(opts *bind.WatchOpts, sink chan<- *ApiSwapCallback) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "SwapCallback")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiSwapCallback)
				if err := _Api.contract.UnpackLog(event, "SwapCallback", log); err != nil {
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

// ParseSwapCallback is a log parse operation binding the contract event 0xd48241df4a75e663b29e55f9506b31f77ed0f48cfe7e7612d1163144995dc1ca.
//
// Solidity: event SwapCallback(int256 amount0Delta, int256 amount1Delta)
func (_Api *ApiFilterer) ParseSwapCallback(log types.Log) (*ApiSwapCallback, error) {
	event := new(ApiSwapCallback)
	if err := _Api.contract.UnpackLog(event, "SwapCallback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

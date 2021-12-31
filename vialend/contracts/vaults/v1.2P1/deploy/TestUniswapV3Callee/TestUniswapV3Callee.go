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
const ApiABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee1\",\"type\":\"uint256\"}],\"name\":\"FlashCallback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Owed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Owed\",\"type\":\"uint256\"}],\"name\":\"MintCallback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"}],\"name\":\"SwapCallback\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pay0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pay1\",\"type\":\"uint256\"}],\"name\":\"flash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"baseAmount\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"}],\"name\":\"getQuoteAtTick\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"name\":\"getTwap\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"amountSpecified\",\"type\":\"int256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"swap0ForExact1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"swap1ForExact0\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"swapExact0For1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"swapExact1For0\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"swapToHigherSqrtPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"swapToLowerSqrtPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3FlashCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Owed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Owed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ApiBin is the compiled bytecode used for deploying new contracts.
var ApiBin = "0x608060405234801561001057600080fd5b50611cbd806100206000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80639e77b8051161008c578063e2be910911610066578063e2be9109146103dd578063e9cbafb014610419578063f603482c14610493578063fa461e33146104cf576100ea565b80639e77b805146102ef578063bac7bf7814610327578063d348799714610363576100ea565b80635b71a46e116100c85780635b71a46e146101c957806366d7505b146102165780636dfc0ddb1461025f5780637b4f53271461029b576100ea565b8063034b0f8f146100ef5780632ec20bf91461013957806343c57a2714610171575b600080fd5b610137600480360360c081101561010557600080fd5b506001600160a01b03813581169160208101359091169060408101359060608101359060808101359060a00135610549565b005b6101376004803603606081101561014f57600080fd5b506001600160a01b03813581169160208101358216916040909101351661066a565b6101b76004803603608081101561018757600080fd5b50803560020b906001600160801b03602082013516906001600160a01b03604082013581169160600135166107a6565b60408051918252519081900360200190f35b6101fd600480360360608110156101df57600080fd5b506001600160a01b038135169060208101351515906040013561089d565b6040805192835260208301919091528051918290030190f35b6102486004803603604081101561022c57600080fd5b5080356001600160a01b0316906020013563ffffffff16610a07565b6040805160029290920b8252519081900360200190f35b6101376004803603608081101561027557600080fd5b506001600160a01b03813581169160208101359160408201358116916060013516610d04565b610137600480360360a08110156102b157600080fd5b5080356001600160a01b03908116916020810135909116906040810135600290810b91606081013590910b90608001356001600160801b0316610e42565b6101376004803603606081101561030557600080fd5b506001600160a01b038135811691602081013582169160409091013516610f7c565b6101376004803603608081101561033d57600080fd5b506001600160a01b03813581169160208101359160408201358116916060013516611037565b6101376004803603606081101561037957600080fd5b813591602081013591810190606081016040820135600160201b81111561039f57600080fd5b8201836020820111156103b157600080fd5b803590602001918460018302840111600160201b831117156103d257600080fd5b5090925090506110f6565b610137600480360360808110156103f357600080fd5b506001600160a01b03813581169160208101359160408201358116916060013516611330565b6101376004803603606081101561042f57600080fd5b813591602081013591810190606081016040820135600160201b81111561045557600080fd5b82018360208201111561046757600080fd5b803590602001918460018302840111600160201b8311171561048857600080fd5b50909250905061134b565b610137600480360360808110156104a957600080fd5b506001600160a01b038135811691602081013591604082013581169160600135166115a0565b610137600480360360608110156104e557600080fd5b813591602081013591810190606081016040820135600160201b81111561050b57600080fd5b82018360208201111561051d57600080fd5b803590602001918460018302840111600160201b8311171561053e57600080fd5b5090925090506115bb565b856001600160a01b031663490e6cbc86868633878760405160200180846001600160a01b0316815260200183815260200182815260200193505050506040516020818303038152906040526040518563ffffffff1660e01b815260040180856001600160a01b0316815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156105fb5781810151838201526020016105e3565b50505050905090810190601f1680156106285780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b15801561064a57600080fd5b505af115801561065e573d6000803e3d6000fd5b50505050505050505050565b826001600160a01b031663128acb088260016001600160ff1b03863360405160200180826001600160a01b031681526020019150506040516020818303038152906040526040518663ffffffff1660e01b815260040180866001600160a01b031681526020018515158152602001848152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561072657818101518382015260200161070e565b50505050905090810190601f1680156107535780820380516001836020036101000a031916815260200191505b5096505050505050506040805180830381600087803b15801561077557600080fd5b505af1158015610789573d6000803e3d6000fd5b505050506040513d604081101561079f57600080fd5b5050505050565b6000806107b2866117e8565b90506001600160801b036001600160a01b03821611610821576001600160a01b0380821680029084811690861610610801576107fc600160c01b876001600160801b031683611b19565b610819565b61081981876001600160801b0316600160c01b611b19565b925050610894565b60006108406001600160a01b0383168068010000000000000000611b19565b9050836001600160a01b0316856001600160a01b03161061087857610873600160801b876001600160801b031683611b19565b610890565b61089081876001600160801b0316600160801b611b19565b9250505b50949350505050565b600080846001600160a01b031663128acb08338686886108d15773fffd8963efd1fc6a506488495d951d5263988d256108d8565b6401000276a45b3360405160200180826001600160a01b031681526020019150506040516020818303038152906040526040518663ffffffff1660e01b815260040180866001600160a01b031681526020018515158152602001848152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610979578181015183820152602001610961565b50505050905090810190601f1680156109a65780820380516001836020036101000a031916815260200191505b5096505050505050506040805180830381600087803b1580156109c857600080fd5b505af11580156109dc573d6000803e3d6000fd5b505050506040513d60408110156109f257600080fd5b50805160209091015190969095509350505050565b600063ffffffff8216610a47576040805162461bcd60e51b815260206004820152600360248201526207842560ec1b604482015290519081900360640190fd5b6040805160028082526060820183526000926020830190803683370190505090508281600081518110610a7657fe5b602002602001019063ffffffff16908163ffffffff1681525050600081600181518110610a9f57fe5b63ffffffff90921660209283029190910182015260405163883bdbfd60e01b8152600481018281528351602483015283516000936001600160a01b0389169363883bdbfd938793909283926044019185820191028083838b5b83811015610b10578181015183820152602001610af8565b505050509050019250505060006040518083038186803b158015610b3357600080fd5b505afa158015610b47573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040908152811015610b7057600080fd5b8101908080516040519392919084600160201b821115610b8f57600080fd5b908301906020820185811115610ba457600080fd5b82518660208202830111600160201b82111715610bc057600080fd5b82525081516020918201928201910280838360005b83811015610bed578181015183820152602001610bd5565b5050505090500160405260200180516040519392919084600160201b821115610c1557600080fd5b908301906020820185811115610c2a57600080fd5b82518660208202830111600160201b82111715610c4657600080fd5b82525081516020918201928201910280838360005b83811015610c73578181015183820152602001610c5b565b50505050905001604052505050509050600081600081518110610c9257fe5b602002602001015182600181518110610ca757fe5b60200260200101510390508463ffffffff168160060b81610cc457fe5b05935060008160060b128015610cee57508463ffffffff168160060b81610ce757fe5b0760060b15155b15610cfb57600019909301925b50505092915050565b836001600160a01b031663128acb08836001610d1f87611bd4565b853360405160200180826001600160a01b031681526020019150506040516020818303038152906040526040518663ffffffff1660e01b815260040180866001600160a01b031681526020018515158152602001848152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610dc1578181015183820152602001610da9565b50505050905090810190601f168015610dee5780820380516001836020036101000a031916815260200191505b5096505050505050506040805180830381600087803b158015610e1057600080fd5b505af1158015610e24573d6000803e3d6000fd5b505050506040513d6040811015610e3a57600080fd5b505050505050565b846001600160a01b0316633c8a7d8d858585853360405160200180826001600160a01b031681526020019150506040516020818303038152906040526040518663ffffffff1660e01b815260040180866001600160a01b031681526020018560020b81526020018460020b8152602001836001600160801b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610efa578181015183820152602001610ee2565b50505050905090810190601f168015610f275780820380516001836020036101000a031916815260200191505b5096505050505050506040805180830381600087803b158015610f4957600080fd5b505af1158015610f5d573d6000803e3d6000fd5b505050506040513d6040811015610f7357600080fd5b50505050505050565b826001600160a01b031663128acb088260006001600160ff1b03863360405160200180826001600160a01b031681526020019150506040516020818303038152906040526040518663ffffffff1660e01b815260040180866001600160a01b031681526020018515158152602001848152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360008381101561072657818101518382015260200161070e565b836001600160a01b031663128acb0883600161105287611bd4565b600003853360405160200180826001600160a01b031681526020019150506040516020818303038152906040526040518663ffffffff1660e01b815260040180866001600160a01b031681526020018515158152602001848152602001836001600160a01b03168152602001806020018281038252838181518152602001915080519060200190808383600083811015610dc1578181015183820152602001610da9565b60008282602081101561110857600080fd5b506040805187815260208101879052815192356001600160a01b031693507fa0968be00566083701c9ef671c169d7fb05ac8907de4ca17185de74ccbb694d4929081900390910190a1841561124257336001600160a01b0316630dfe16816040518163ffffffff1660e01b815260040160206040518083038186803b15801561119057600080fd5b505afa1580156111a4573d6000803e3d6000fd5b505050506040513d60208110156111ba57600080fd5b5051604080516323b872dd60e01b81526001600160a01b03848116600483015233602483015260448201899052915191909216916323b872dd9160648083019260209291908290030181600087803b15801561121557600080fd5b505af1158015611229573d6000803e3d6000fd5b505050506040513d602081101561123f57600080fd5b50505b831561079f57336001600160a01b031663d21220a76040518163ffffffff1660e01b815260040160206040518083038186803b15801561128157600080fd5b505afa158015611295573d6000803e3d6000fd5b505050506040513d60208110156112ab57600080fd5b5051604080516323b872dd60e01b81526001600160a01b03848116600483015233602483015260448201889052915191909216916323b872dd9160648083019260209291908290030181600087803b15801561130657600080fd5b505af115801561131a573d6000803e3d6000fd5b505050506040513d6020811015610f7357600080fd5b836001600160a01b031663128acb08836000610d1f87611bd4565b604080518581526020810185905281517f2b0391b4fa408cfe47abd1977d72985695b2e5ebd3175f55be25f2c68c5df21b929181900390910190a160008060008484606081101561139b57600080fd5b506001600160a01b0381351693506020810135925060400135905081156114a757336001600160a01b0316630dfe16816040518163ffffffff1660e01b815260040160206040518083038186803b1580156113f557600080fd5b505afa158015611409573d6000803e3d6000fd5b505050506040513d602081101561141f57600080fd5b5051604080516323b872dd60e01b81526001600160a01b03868116600483015233602483015260448201869052915191909216916323b872dd9160648083019260209291908290030181600087803b15801561147a57600080fd5b505af115801561148e573d6000803e3d6000fd5b505050506040513d60208110156114a457600080fd5b50505b8015610f7357336001600160a01b031663d21220a76040518163ffffffff1660e01b815260040160206040518083038186803b1580156114e657600080fd5b505afa1580156114fa573d6000803e3d6000fd5b505050506040513d602081101561151057600080fd5b5051604080516323b872dd60e01b81526001600160a01b03868116600483015233602483015260448201859052915191909216916323b872dd9160648083019260209291908290030181600087803b15801561156b57600080fd5b505af115801561157f573d6000803e3d6000fd5b505050506040513d602081101561159557600080fd5b505050505050505050565b836001600160a01b031663128acb0883600061105287611bd4565b6000828260208110156115cd57600080fd5b506040805187815260208101879052815192356001600160a01b031693507fd48241df4a75e663b29e55f9506b31f77ed0f48cfe7e7612d1163144995dc1ca929081900390910190a1600085131561170f57336001600160a01b0316630dfe16816040518163ffffffff1660e01b815260040160206040518083038186803b15801561165857600080fd5b505afa15801561166c573d6000803e3d6000fd5b505050506040513d602081101561168257600080fd5b5051604080516323b872dd60e01b81526001600160a01b03848116600483015233602483015260448201899052915191909216916323b872dd9160648083019260209291908290030181600087803b1580156116dd57600080fd5b505af11580156116f1573d6000803e3d6000fd5b505050506040513d602081101561170757600080fd5b5061079f9050565b60008413156117d657336001600160a01b031663d21220a76040518163ffffffff1660e01b815260040160206040518083038186803b15801561175157600080fd5b505afa158015611765573d6000803e3d6000fd5b505050506040513d602081101561177b57600080fd5b5051604080516323b872dd60e01b81526001600160a01b03848116600483015233602483015260448201889052915191909216916323b872dd9160648083019260209291908290030181600087803b1580156116dd57600080fd5b841580156117e2575083155b61079f57fe5b60008060008360020b126117ff578260020b611807565b8260020b6000035b9050620d89e8811115611845576040805162461bcd60e51b81526020600482015260016024820152601560fa1b604482015290519081900360640190fd5b60006001821661185957600160801b61186b565b6ffffcb933bd6fad37aa2d162d1a5940015b70ffffffffffffffffffffffffffffffffff169050600282161561189f576ffff97272373d413259a46990580e213a0260801c5b60048216156118be576ffff2e50f5f656932ef12357cf3c7fdcc0260801c5b60088216156118dd576fffe5caca7e10e4e61c3624eaa0941cd00260801c5b60108216156118fc576fffcb9843d60f6159c9db58835c9266440260801c5b602082161561191b576fff973b41fa98c081472e6896dfb254c00260801c5b604082161561193a576fff2ea16466c96a3843ec78b326b528610260801c5b6080821615611959576ffe5dee046a99a2a811c461f1969c30530260801c5b610100821615611979576ffcbe86c7900a88aedcffc83b479aa3a40260801c5b610200821615611999576ff987a7253ac413176f2b074cf7815e540260801c5b6104008216156119b9576ff3392b0822b70005940c7a398e4b70f30260801c5b6108008216156119d9576fe7159475a2c29b7443b29c7fa6e889d90260801c5b6110008216156119f9576fd097f3bdfd2022b8845ad8f792aa58250260801c5b612000821615611a19576fa9f746462d870fdf8a65dc1f90e061e50260801c5b614000821615611a39576f70d869a156d2a1b890bb3df62baf32f70260801c5b618000821615611a59576f31be135f97d08fd981231505542fcfa60260801c5b62010000821615611a7a576f09aa508b5b7a84e1c677de54f3e99bc90260801c5b62020000821615611a9a576e5d6af8dedb81196699c329225ee6040260801c5b62040000821615611ab9576d2216e584f5fa1ea926041bedfe980260801c5b62080000821615611ad6576b048a170391f7dc42444e8fa20260801c5b60008460020b1315611af1578060001981611aed57fe5b0490505b600160201b810615611b04576001611b07565b60005b60ff16602082901c0192505050919050565b6000806000611b288686611bea565b9150915060008480611b3657fe5b868809905082811115611b4a576001820391505b918290039181611b6857848381611b5d57fe5b049350505050611bcd565b848210611bbc576040805162461bcd60e51b815260206004820152601a60248201527f46756c6c4d6174683a2046554c4c4449565f4f564552464c4f57000000000000604482015290519081900360640190fd5b611bc7838387611c17565b93505050505b9392505050565b6000600160ff1b8210611be657600080fd5b5090565b6000808060001984860990508385029250828103915082811015611c0f576001820391505b509250929050565b60008181038216808381611c2757fe5b049250808581611c3357fe5b049450808160000381611c4257fe5b6002858103808702820302808702820302808702820302808702820302808702820302808702820302958602900390940293046001019390930293909301029291505056fea26469706673582212206c44bf403491a8fe150c45d54f8ae9176c7a4daeb618fd30c1fa50f64a1871a764736f6c63430007060033"

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

// GetQuoteAtTick is a free data retrieval call binding the contract method 0x43c57a27.
//
// Solidity: function getQuoteAtTick(int24 tick, uint128 baseAmount, address baseToken, address quoteToken) pure returns(uint256 quoteAmount)
func (_Api *ApiCaller) GetQuoteAtTick(opts *bind.CallOpts, tick *big.Int, baseAmount *big.Int, baseToken common.Address, quoteToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getQuoteAtTick", tick, baseAmount, baseToken, quoteToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuoteAtTick is a free data retrieval call binding the contract method 0x43c57a27.
//
// Solidity: function getQuoteAtTick(int24 tick, uint128 baseAmount, address baseToken, address quoteToken) pure returns(uint256 quoteAmount)
func (_Api *ApiSession) GetQuoteAtTick(tick *big.Int, baseAmount *big.Int, baseToken common.Address, quoteToken common.Address) (*big.Int, error) {
	return _Api.Contract.GetQuoteAtTick(&_Api.CallOpts, tick, baseAmount, baseToken, quoteToken)
}

// GetQuoteAtTick is a free data retrieval call binding the contract method 0x43c57a27.
//
// Solidity: function getQuoteAtTick(int24 tick, uint128 baseAmount, address baseToken, address quoteToken) pure returns(uint256 quoteAmount)
func (_Api *ApiCallerSession) GetQuoteAtTick(tick *big.Int, baseAmount *big.Int, baseToken common.Address, quoteToken common.Address) (*big.Int, error) {
	return _Api.Contract.GetQuoteAtTick(&_Api.CallOpts, tick, baseAmount, baseToken, quoteToken)
}

// GetTwap is a free data retrieval call binding the contract method 0x66d7505b.
//
// Solidity: function getTwap(address pool, uint32 period) view returns(int24 tick)
func (_Api *ApiCaller) GetTwap(opts *bind.CallOpts, pool common.Address, period uint32) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getTwap", pool, period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTwap is a free data retrieval call binding the contract method 0x66d7505b.
//
// Solidity: function getTwap(address pool, uint32 period) view returns(int24 tick)
func (_Api *ApiSession) GetTwap(pool common.Address, period uint32) (*big.Int, error) {
	return _Api.Contract.GetTwap(&_Api.CallOpts, pool, period)
}

// GetTwap is a free data retrieval call binding the contract method 0x66d7505b.
//
// Solidity: function getTwap(address pool, uint32 period) view returns(int24 tick)
func (_Api *ApiCallerSession) GetTwap(pool common.Address, period uint32) (*big.Int, error) {
	return _Api.Contract.GetTwap(&_Api.CallOpts, pool, period)
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

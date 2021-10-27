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
const ApiABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cToken0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cToken1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cEth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_protocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxTotalSupply\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"_maxTwapDeviation\",\"type\":\"int24\"},{\"internalType\":\"uint32\",\"name\":\"_twapDuration\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_uniPortion\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_team\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"uFees0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"uFees1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lFees0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lFees1\",\"type\":\"uint256\"}],\"name\":\"CollectFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"ErrorLogging\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"subject\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"GeneralA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"subject\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"GeneralB\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"subject\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"GeneralS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MyLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MyLog2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MyLog4\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance1\",\"type\":\"uint256\"}],\"name\":\"RebalanceLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"Snapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Assetholder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"deposit0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"current0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"current1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Fees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"U3Fees0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"U3Fees1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"LcFees0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"LcFees1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"AccruedProtocolFees0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"AccruedProtocolFees1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alloc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cHigh\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cLow\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"collectProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curComp0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curComp1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToken1\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"doRebalence\",\"type\":\"bool\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getETHBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"getPositionAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"baseAmount\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"}],\"name\":\"getQuoteAtTick\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"getSSLiquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"getStoredAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"name\":\"getTwap\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTwapDeviation\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeCTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"setGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMax\",\"type\":\"uint256\"}],\"name\":\"setMaxTotalSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_maxTwapDeviation\",\"type\":\"int24\"}],\"name\":\"setMaxTwapDeviation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_protocolFee\",\"type\":\"uint256\"}],\"name\":\"setProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_team\",\"type\":\"address\"}],\"name\":\"setTeam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_twapDuration\",\"type\":\"uint32\"}],\"name\":\"setTwapDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"ratio\",\"type\":\"uint8\"}],\"name\":\"setUniPortionRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"newLow\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"newHigh\",\"type\":\"int24\"}],\"name\":\"strategy1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20Contract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cErc20Contract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_numTokensToSupply\",\"type\":\"uint256\"}],\"name\":\"supplyErc20ToCompound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_cEtherContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ctoken\",\"type\":\"address\"}],\"name\":\"supplyEthToCompound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"swapAmount\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"team\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"twapDuration\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniPortion\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ApiBin is the compiled bytecode used for deploying new contracts.
var ApiBin = "0x60c06040523480156200001157600080fd5b50604051620064003803806200640083398181016040526101608110156200003857600080fd5b5080516020808301516040808501516060860151608087015160a088015160c089015160e08a01516101008b01516101208c0151610140909c01518851808a018a52600e81526d05669614c656e6420546f6b656e360941b818d019081528a51808c01909b5260048b52630564c54360e41b9c8b019c909c52600160005580519c9d9a9c989b979a9699959894979396929592949193909291620000df91601c91620004a0565b508051620000f590601d906020840190620004a0565b5050601e805460ff1916601217905550600180546001600160a01b03199081163317909155600280546001600160a01b038481169184169190911790915560058054918e1691909216811790915560408051630dfe168160e01b81529051630dfe168191600480820192602092909190829003018186803b1580156200017a57600080fd5b505afa1580156200018f573d6000803e3d6000fd5b505050506040513d6020811015620001a657600080fd5b505160601b6001600160601b0319166080526040805163d21220a760e01b815290516001600160a01b038d169163d21220a7916004808301926020929190829003018186803b158015620001f957600080fd5b505afa1580156200020e573d6000803e3d6000fd5b505050506040513d60208110156200022557600080fd5b505160601b6001600160601b03191660a05260138054600160401b600160e01b031916680100000000000000006001600160a01b038c81169190910291909117909155601480546001600160a01b03199081168b841617909155601580549091168983161790558a16620002c9576040805162461bcd60e51b815260206004808301919091526024820152630ae8aa8960e31b604482015290519081900360640190fd5b89601660006101000a8154816001600160a01b0302191690836001600160a01b03160217905550856004819055508a6001600160a01b031663d0c93a7c6040518163ffffffff1660e01b815260040160206040518083038186803b1580156200033157600080fd5b505afa15801562000346573d6000803e3d6000fd5b505050506040513d60208110156200035d57600080fd5b50516005805460029290920b62ffffff16600160a01b0262ffffff60a01b1990921691909117905560648610620003c1576040805162461bcd60e51b815260206004820152600360248201526228232960e91b604482015290519081900360640190fd5b60128590556013805460ff84166701000000000000000260ff60381b1963ffffffff871663ffffffff1960028a900b62ffffff81166401000000000262ffffff60201b1990961695909517161716179091556000126200044e576040805162461bcd60e51b815260206004820152600360248201526213551160ea1b604482015290519081900360640190fd5b60008363ffffffff16116200048f576040805162461bcd60e51b8152602060048201526002602482015261151160f21b604482015290519081900360640190fd5b50505050505050505050506200054c565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282620004d8576000855562000523565b82601f10620004f357805160ff191683800117855562000523565b8280016001018555821562000523579182015b828111156200052357825182559160200191906001019062000506565b506200053192915062000535565b5090565b5b8082111562000531576000815560010162000536565b60805160601c60a05160601c615df46200060c60003980610f5452806115f45280611d9a5280611e345280612cc05280612d735280612ecb528061314752806132f252806133ba5280613439528061356c5280613ac75280613e42528061432952806149dd525080610f0a52806111b652806115515280611d795280611df95280612d395280612e3252806130f25280613180528061324852806132c7528061352f5280613a275280613da9528061430852806149bc5250615df46000f3fe60806040526004361061036f5760003560e01c80636e947298116101c6578063c433c80a116100f7578063d368867f11610095578063e7c7cb911161006f578063e7c7cb9114610d76578063f2a40db814610d8b578063f39c38a014610db5578063fa461e3314610dca57610376565b8063d368867f14610d11578063d8d7f96f14610d26578063dd62ed3e14610d3b57610376565b8063cda206b0116100d1578063cda206b014610c0d578063d17cfebd14610c60578063d21220a714610c75578063d348799714610c8a57610376565b8063c433c80a14610bb3578063c8a746a614610be3578063ccac7d8b14610bf857610376565b8063a457c2d711610164578063a91ef6eb1161013e578063a91ef6eb14610b1f578063ab033ea914610b56578063b0e21e8a14610b89578063bbc001c314610b9e57610376565b8063a457c2d714610a54578063a4ee141b14610a8d578063a9059cbb14610ae657610376565b8063787dce3d116101a0578063787dce3d146109eb57806385f2aef214610a155780638e843c6c14610a2a57806395d89b4114610a3f57610376565b80636e9472981461097557806370a082311461098a57806373232c60146109bd57610376565b806331dc5b95116102a057806343c57a271161023e578063624a177a11610218578063624a177a1461087a5780636253c718146108a757806366d7505b146108de57806368a5787a1461093457610376565b806343c57a27146107b45780635aa6e675146108075780635bb6aa851461081c57610376565b80633b7ba25f1161027a5780633b7ba25f146106e25780633cbff3fe146107255780633f3e4c111461075257806343a0d0661461077c57610376565b806331dc5b951461064c57806339509351146106615780633aaa36e61461069a57610376565b806318160ddd1161030d57806326d89545116102e757806326d89545146105b45780632ab4d052146105e25780632e1a7d4d146105f7578063313ce5671461062157610376565b806318160ddd14610535578063238efcbc1461055c57806323b872dd1461057157610376565b8063095cf5c611610349578063095cf5c61461046f578063095ea7b3146104a25780630dfe1681146104ef5780631755ff211461052057610376565b80630430c1301461037857806306b0b656146103b757806306fdde03146103e557610376565b3661037657005b005b34801561038457600080fd5b506103766004803603606081101561039b57600080fd5b50803590602081013590604001356001600160a01b0316610e51565b3480156103c357600080fd5b506103cc610f80565b6040805192835260208301919091528051918290030190f35b3480156103f157600080fd5b506103fa611091565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561043457818101518382015260200161041c565b50505050905090810190601f1680156104615780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561047b57600080fd5b506103766004803603602081101561049257600080fd5b50356001600160a01b0316611128565b3480156104ae57600080fd5b506104db600480360360408110156104c557600080fd5b506001600160a01b038135169060200135611196565b604080519115158252519081900360200190f35b3480156104fb57600080fd5b506105046111b4565b604080516001600160a01b039092168252519081900360200190f35b34801561052c57600080fd5b506105046111d8565b34801561054157600080fd5b5061054a6111e7565b60408051918252519081900360200190f35b34801561056857600080fd5b506103766111ed565b34801561057d57600080fd5b506104db6004803603606081101561059457600080fd5b506001600160a01b03813581169160208101359091169060400135611254565b3480156105c057600080fd5b506105c96112dc565b6040805163ffffffff9092168252519081900360200190f35b3480156105ee57600080fd5b5061054a6112e8565b34801561060357600080fd5b506103cc6004803603602081101561061a57600080fd5b50356112ee565b34801561062d57600080fd5b50610636611711565b6040805160ff9092168252519081900360200190f35b34801561065857600080fd5b5061063661171a565b34801561066d57600080fd5b506104db6004803603604081101561068457600080fd5b506001600160a01b03813516906020013561172a565b3480156106a657600080fd5b506106af611778565b604080519687526020870195909552858501939093526060850191909152608084015260a0830152519081900360c00190f35b3480156106ee57600080fd5b5061054a6004803603606081101561070557600080fd5b506001600160a01b0381358116916020810135909116906040013561178d565b34801561073157600080fd5b506103766004803603602081101561074857600080fd5b503560020b611ada565b34801561075e57600080fd5b506103766004803603602081101561077557600080fd5b5035611b9a565b34801561078857600080fd5b506103766004803603606081101561079f57600080fd5b50803590602081013590604001351515611c3b565b3480156107c057600080fd5b5061054a600480360360808110156107d757600080fd5b50803560020b906001600160801b03602082013516906001600160a01b0360408201358116916060013516611f51565b34801561081357600080fd5b50610504612043565b34801561082857600080fd5b5061084f6004803603602081101561083f57600080fd5b50356001600160a01b0316612052565b6040805195865260208601949094528484019290925260608401526080830152519081900360a00190f35b34801561088657600080fd5b506103766004803603602081101561089d57600080fd5b503560ff16612081565b3480156108b357600080fd5b50610376600480360360408110156108ca57600080fd5b508035600290810b9160200135900b612132565b3480156108ea57600080fd5b5061091d6004803603604081101561090157600080fd5b5080356001600160a01b0316906020013563ffffffff16612226565b6040805160029290920b8252519081900360200190f35b34801561094057600080fd5b506103cc6004803603606081101561095757600080fd5b508035906020810135151590604001356001600160a01b03166124ea565b34801561098157600080fd5b5061054a61261f565b34801561099657600080fd5b5061054a600480360360208110156109ad57600080fd5b50356001600160a01b0316612623565b6104db600480360360408110156109d357600080fd5b506001600160a01b0381358116916020013516612642565b3480156109f757600080fd5b5061037660048036036020811015610a0e57600080fd5b5035612822565b348015610a2157600080fd5b506105046128b8565b348015610a3657600080fd5b5061091d6128c7565b348015610a4b57600080fd5b506103fa6128d7565b348015610a6057600080fd5b506104db60048036036040811015610a7757600080fd5b506001600160a01b038135169060200135612938565b348015610a9957600080fd5b50610ac060048036036020811015610ab057600080fd5b50356001600160a01b03166129a0565b604080519485526020850193909352838301919091526060830152519081900360800190f35b348015610af257600080fd5b506104db60048036036040811015610b0957600080fd5b506001600160a01b0381351690602001356129d1565b348015610b2b57600080fd5b506103cc60048036036040811015610b4257600080fd5b508035600290810b9160200135900b6129e5565b348015610b6257600080fd5b5061037660048036036020811015610b7957600080fd5b50356001600160a01b0316612a7c565b348015610b9557600080fd5b5061054a612aea565b348015610baa57600080fd5b50610376612af0565b348015610bbf57600080fd5b5061037660048036036020811015610bd657600080fd5b503563ffffffff16612b47565b348015610bef57600080fd5b5061054a612bf9565b348015610c0457600080fd5b5061054a612bff565b348015610c1957600080fd5b50610c4460048036036040811015610c3057600080fd5b508035600290810b9160200135900b612c05565b604080516001600160801b039092168252519081900360200190f35b348015610c6c57600080fd5b50610376612c1d565b348015610c8157600080fd5b50610504612cbe565b348015610c9657600080fd5b5061037660048036036060811015610cad57600080fd5b813591602081013591810190606081016040820135600160201b811115610cd357600080fd5b820183602082011115610ce557600080fd5b803590602001918460018302840111600160201b83111715610d0657600080fd5b509092509050612ce2565b348015610d1d57600080fd5b5061091d612da0565b348015610d3257600080fd5b50610376612db0565b348015610d4757600080fd5b5061054a60048036036040811015610d5e57600080fd5b506001600160a01b0381358116916020013516613460565b348015610d8257600080fd5b5061091d61348b565b348015610d9757600080fd5b5061050460048036036020811015610dae57600080fd5b503561349b565b348015610dc157600080fd5b506105046134c5565b348015610dd657600080fd5b5061037660048036036060811015610ded57600080fd5b813591602081013591810190606081016040820135600160201b811115610e1357600080fd5b820183602082011115610e2557600080fd5b803590602001918460018302840111600160201b83111715610e4657600080fd5b5090925090506134d4565b6001546001600160a01b03163314610e9d576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b6010548311801590610eb157506011548211155b610ee7576040805162461bcd60e51b8152602060048201526002602482015261043560f41b604482015290519081900360640190fd5b8215610f3157601054610efa9084613593565b601055610f316001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001682856135f0565b8115610f7b57601154610f449083613593565b601155610f7b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001682846135f0565b505050565b600080601360089054906101000a90046001600160a01b03166001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610fe557600080fd5b505afa158015610ff9573d6000803e3d6000fd5b505050506040513d602081101561100f57600080fd5b5051601454604080516370a0823160e01b815230600482015290519294506001600160a01b03909116916370a0823191602480820192602092909190829003018186803b15801561105f57600080fd5b505afa158015611073573d6000803e3d6000fd5b505050506040513d602081101561108957600080fd5b505191929050565b601c8054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561111d5780601f106110f25761010080835404028352916020019161111d565b820191906000526020600020905b81548152906001019060200180831161110057829003601f168201915b505050505090505b90565b6001546001600160a01b03163314611174576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b600280546001600160a01b0319166001600160a01b0392909216919091179055565b60006111aa6111a3613642565b8484613646565b5060015b92915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6005546001600160a01b031690565b601b5490565b6003546001600160a01b03163314611240576040805162461bcd60e51b815260206004820152601160248201527070656e64696e67476f7665726e616e636560781b604482015290519081900360640190fd5b600180546001600160a01b03191633179055565b6000611261848484613732565b6112d18461126d613642565b6112cc85604051806060016040528060288152602001615cbe602891396001600160a01b038a166000908152601a60205260408120906112ab613642565b6001600160a01b03168152602081019190915260400160002054919061388f565b613646565b5060015b9392505050565b60135463ffffffff1681565b60125481565b60008060026000541415611337576040805162461bcd60e51b815260206004820152601f6024820152600080516020615ba6833981519152604482015290519081900360640190fd5b600260005533801580159061135557506001600160a01b0381163014155b61138c576040805162461bcd60e51b8152602060048201526003602482015262746f5760e81b604482015290519081900360640190fd5b60008411801561139d575060648411155b6113d7576040805162461bcd60e51b815260206004808301919091526024820152631418db9d60e21b604482015290519081900360640190fd5b6113df613926565b50600061140060646113fa876113f486612623565b90613963565b906139bc565b9050600061140c6111e7565b905060008111611449576040805162461bcd60e51b815260206004820152600360248201526207473360ec1b604482015290519081900360640190fd5b80821115611484576040805162461bcd60e51b815260206004820152600360248201526273683160e81b604482015290519081900360640190fd5b611494816113fa846113f4613a23565b94506114a6816113fa846113f4613ac3565b93506114b28383613b32565b600c546114d2906114c99083906113fa9086613963565b600c5490613593565b600c55600d546114f5906114ec9083906113fa9086613963565b600d5490613593565b600d55600e546115189061150f9083906113fa9086613963565b600e5490613593565b600e55600f5461153b906115329083906113fa9086613963565b600f5490613593565b600f5584156115e1576115786001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001684876135f0565b6001600160a01b0383166000908152600b602052604090206002015485106115bb576001600160a01b0383166000908152600b60205260409020600201546115bd565b845b6001600160a01b0384166000908152600b6020526040902060020180549190910390555b83156116845761161b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001684866135f0565b6001600160a01b0383166000908152600b6020526040902060030154841061165e576001600160a01b0383166000908152600b6020526040902060030154611660565b835b6001600160a01b0384166000908152600b6020526040902060030180549190910390555b61168d83612623565b6116bf576001600160a01b0383166000908152600b602052604081208181556001810182905560028101829055600301555b6040805183815260208101879052808201869052905133917f02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94919081900360600190a250506001600055509092909150565b601e5460ff1690565b601354600160381b900460ff1681565b60006111aa611737613642565b846112cc85601a6000611748613642565b6001600160a01b03908116825260208083019390935260409182016000908120918c168152925290205490613c2e565b600c54600d54600e54600f5460105460115486565b6000836001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156117dc57600080fd5b505afa1580156117f0573d6000803e3d6000fd5b505050506040513d602081101561180657600080fd5b5051821115611846576040805162461bcd60e51b815260206004820152600760248201526662616c616e636560c81b604482015290519081900360640190fd5b600084905060008490506000816001600160a01b031663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561188d57600080fd5b505af11580156118a1573d6000803e3d6000fd5b505050506040513d60208110156118b757600080fd5b50516040805160208101839052818152601b818301527f45786368616e6765205261746520287363616c6564207570293a20000000000060608201529051919250600080516020615d50833981519152919081900360800190a16000826001600160a01b031663ae9d70b06040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561194e57600080fd5b505af1158015611962573d6000803e3d6000fd5b505050506040513d602081101561197857600080fd5b505160408051602081018390528181526018818301527f537570706c7920526174653a20287363616c656420757029000000000000000060608201529051919250600080516020615d50833981519152919081900360800190a1836001600160a01b031663095ea7b388886040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b158015611a2957600080fd5b505af1158015611a3d573d6000803e3d6000fd5b505050506040513d6020811015611a5357600080fd5b50506040805163140e25ad60e31b81526004810188905290516000916001600160a01b0386169163a0712d689160248082019260209290919082900301818787803b158015611aa157600080fd5b505af1158015611ab5573d6000803e3d6000fd5b505050506040513d6020811015611acb57600080fd5b50519998505050505050505050565b6001546001600160a01b03163314611b26576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b60008160020b13611b71576040805162461bcd60e51b815260206004820152601060248201526f36b0bc2a3bb0b82232bb34b0ba34b7b760811b604482015290519081900360640190fd5b6013805460029290920b62ffffff16600160201b0266ffffff0000000019909216919091179055565b60026000541415611be0576040805162461bcd60e51b815260206004820152601f6024820152600080516020615ba6833981519152604482015290519081900360640190fd5b60026000556001546001600160a01b03163314611c31576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b6012556001600055565b60026000541415611c81576040805162461bcd60e51b815260206004820152601f6024820152600080516020615ba6833981519152604482015290519081900360640190fd5b60026000553383151580611c955750600083115b611ccf576040805162461bcd60e51b815260206004808301919091526024820152630616d74360e41b604482015290519081900360640190fd5b6001600160a01b03811615801590611cf057506001600160a01b0381163014155b8015611d0a57506002546001600160a01b03828116911614155b611d40576040805162461bcd60e51b8152602060048201526002602482015261746f60f01b604482015290519081900360640190fd5b611d4981613c88565b600554601354611dbe91611d6e916001600160a01b039091169063ffffffff16612226565b670de0b6b3a76400007f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611f51565b60215560008080611dcf8787613d09565b925092509250611ddd613926565b50611de6613d69565b8115611e2157611e216001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016333085613ffc565b8015611e5c57611e5c6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016333084613ffc565b611e668484614056565b601254611e716111e7565b1115611eaa576040805162461bcd60e51b815260206004820152600360248201526204341560ec1b604482015290519081900360640190fd5b6001600160a01b0384166000908152600b6020526040902080548301815560018101805483019055600281018054840190556003018054820190558415611ef657611ef6600080614148565b604080518481526020810184905280820183905290516001600160a01b0386169133917f4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f69181900360600190a3505060016000555050505050565b600080611f5d86614393565b90506001600160801b036001600160a01b03821611611fcc576001600160a01b0380821680029084811690861610611fac57611fa7600160c01b876001600160801b0316836146c4565b611fc4565b611fc481876001600160801b0316600160c01b6146c4565b92505061203a565b6000611fe66001600160a01b03831680600160401b6146c4565b9050836001600160a01b0316856001600160a01b03161061201e57612019600160801b876001600160801b0316836146c4565b612036565b61203681876001600160801b0316600160801b6146c4565b9250505b50949350505050565b6001546001600160a01b031681565b600b60205260009081526040902080546001820154600283015460038401546004909401549293919290919085565b6001546001600160a01b031633146120cd576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b60648160ff16111561210e576040805162461bcd60e51b8152602060048201526005602482015264726174696f60d81b604482015290519081900360640190fd5b6013805460ff909216600160381b0267ff0000000000000019909216919091179055565b60026000541415612178576040805162461bcd60e51b815260206004820152601f6024820152600080516020615ba6833981519152604482015290519081900360640190fd5b60026000556001546001600160a01b031633146121c9576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b60006121d36111e7565b1161220a576040805162461bcd60e51b8152602060048201526002602482015261074360f41b604482015290519081900360640190fd5b612212613926565b5061221d8282614148565b50506001600055565b604080516002808252606082018352600092849284929091602083019080368337019050509050818160008151811061225b57fe5b602002602001019063ffffffff16908163ffffffff168152505060008160018151811061228457fe5b63ffffffff90921660209283029190910182015260405163883bdbfd60e01b8152600481018281528351602483015283516000936001600160a01b038a169363883bdbfd938793909283926044019185820191028083838b5b838110156122f55781810151838201526020016122dd565b505050509050019250505060006040518083038186803b15801561231857600080fd5b505afa15801561232c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604090815281101561235557600080fd5b8101908080516040519392919084600160201b82111561237457600080fd5b90830190602082018581111561238957600080fd5b82518660208202830111600160201b821117156123a557600080fd5b82525081516020918201928201910280838360005b838110156123d25781810151838201526020016123ba565b5050505090500160405260200180516040519392919084600160201b8211156123fa57600080fd5b90830190602082018581111561240f57600080fd5b82518660208202830111600160201b8211171561242b57600080fd5b82525081516020918201928201910280838360005b83811015612458578181015183820152602001612440565b5050505090500160405250505050905060008160008151811061247757fe5b60200260200101518260018151811061248c57fe5b60200260200101510390508363ffffffff168160060b816124a957fe5b05945060008160060b1280156124d357508363ffffffff168160060b816124cc57fe5b0760060b15155b156124e057600019909401935b5050505092915050565b600554604080513360208083019190915282518083038201815282840193849052630251596160e31b90935230604483018181528715156064850152608484018990526001600160a01b0387811660a486015260a060c48601908152865160e48701528651600098899893169663128acb08968c958e958d9590949193909261010490910191908501908083838f5b83811015612591578181015183820152602001612579565b50505050905090810190601f1680156125be5780820380516001836020036101000a031916815260200191505b5096505050505050506040805180830381600087803b1580156125e057600080fd5b505af11580156125f4573d6000803e3d6000fd5b505050506040513d604081101561260a57600080fd5b50805160209091015190969095509350505050565b4790565b6001600160a01b0381166000908152601960205260409020545b919050565b60008083905060008390506000816001600160a01b031663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561268a57600080fd5b505af115801561269e573d6000803e3d6000fd5b505050506040513d60208110156126b457600080fd5b505160408051602081018390528181526023918101829052919250600080516020615d5083398151915291839181906060820190615c2e82396040019250505060405180910390a16000826001600160a01b031663ae9d70b06040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561273957600080fd5b505af115801561274d573d6000803e3d6000fd5b505050506040513d602081101561276357600080fd5b5051604080516020818101849052828252818301527f537570706c7920526174653a20287363616c656420757020627920316531382960608201529051919250600080516020615d50833981519152919081900360800190a1836001600160a01b0316631249c58b6203d090476040518363ffffffff1660e01b81526004016000604051808303818589803b1580156127fb57600080fd5b5088f115801561280f573d6000803e3d6000fd5b5060019c9b505050505050505050505050565b6001546001600160a01b0316331461286e576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b620f424081106128b3576040805162461bcd60e51b815260206004820152600b60248201526a70726f746f636f6c46656560a81b604482015290519081900360640190fd5b600455565b6002546001600160a01b031681565b600554600160d01b900460020b81565b601d8054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561111d5780601f106110f25761010080835404028352916020019161111d565b60006111aa612945613642565b846112cc85604051806060016040528060258152602001615d9a60259139601a600061296f613642565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919061388f565b6001600160a01b03166000908152600b60205260409020805460018201546002830154600390930154919390929190565b60006111aa6129de613642565b8484613732565b60008060008060006129f78787614773565b9450945050509250612a0a87878561482b565b6004549196509450600090612a2390620f424090613593565b9050612a49612a42620f42406113fa6001600160801b03871685613963565b8790613c2e565b9550612a6f612a68620f42406113fa6001600160801b03861685613963565b8690613c2e565b9450505050509250929050565b6001546001600160a01b03163314612ac8576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b600380546001600160a01b0319166001600160a01b0392909216919091179055565b60045481565b6001546001600160a01b03163314612b3c576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b612b44613926565b50565b6001546001600160a01b03163314612b93576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b60008163ffffffff1611612bdd576040805162461bcd60e51b815260206004820152600c60248201526b3a3bb0b8223ab930ba34b7b760a11b604482015290519081900360640190fd5b6013805463ffffffff191663ffffffff92909216919091179055565b60175481565b60185481565b6000612c118383614773565b50929695505050505050565b6001546001600160a01b03163314612c69576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b600080612c74610f80565b90925090508115612c9d57601354612c9d908390600160401b90046001600160a01b03166148d5565b8015612cba57601454612cba9082906001600160a01b03166148d5565b5050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6005546001600160a01b03163314612d26576040805162461bcd60e51b8152602060048201526002602482015261505360f01b604482015290519081900360640190fd5b8315612d6057612d606001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001633866135f0565b8215612d9a57612d9a6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001633856135f0565b50505050565b600554600160b81b900460020b81565b6001546001600160a01b03163314612dfc576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b612e04614994565b612e0c612c1d565b600a54604080516370a0823160e01b815230600482015290516000916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016916370a0823191602480820192602092909190829003018186803b158015612e7957600080fd5b505afa158015612e8d573d6000803e3d6000fd5b505050506040513d6020811015612ea357600080fd5b5051604080516370a0823160e01b815230600482015290519192506000916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016916370a08231916024808301926020929190829003018186803b158015612f1157600080fd5b505afa158015612f25573d6000803e3d6000fd5b505050506040513d6020811015612f3b57600080fd5b505190506000612f496111e7565b905060005b8481101561317b576000612f82600a8381548110612f6857fe5b6000918252602090912001546001600160a01b0316612623565b9050612faf600a8381548110612f9457fe5b6000918252602090912001546001600160a01b031682613b32565b6000600b6000600a8581548110612fc257fe5b60009182526020808320909101546001600160a01b03168352820192909252604001812091909155600a8054600b91839186908110612ffd57fe5b60009182526020808320909101546001600160a01b03168352820192909252604001812060010191909155600a8054600b9183918690811061303b57fe5b60009182526020808320909101546001600160a01b03168352820192909252604001812060020191909155600a8054600b9183918690811061307957fe5b60009182526020808320909101546001600160a01b031683528201929092526040018120600301919091556130b2846113fa8885613963565b905060006130c4856113fa8886613963565b9050811561311b5761311b600a85815481106130dc57fe5b6000918252602090912001546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169116846135f0565b801561317057613170600a858154811061313157fe5b6000918252602090912001546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169116836135f0565b505050600101612f4e565b5060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156131eb57600080fd5b505afa1580156131ff573d6000803e3d6000fd5b505050506040513d602081101561321557600080fd5b505111156132ee57600154604080516370a0823160e01b815230600482015290516132ee926001600160a01b03908116927f0000000000000000000000000000000000000000000000000000000000000000909116916370a0823191602480820192602092909190829003018186803b15801561329157600080fd5b505afa1580156132a5573d6000803e3d6000fd5b505050506040513d60208110156132bb57600080fd5b50516001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691906135f0565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561335d57600080fd5b505afa158015613371573d6000803e3d6000fd5b505050506040513d602081101561338757600080fd5b50511115612d9a57600154604080516370a0823160e01b81523060048201529051612d9a926001600160a01b03908116927f0000000000000000000000000000000000000000000000000000000000000000909116916370a0823191602480820192602092909190829003018186803b15801561340357600080fd5b505afa158015613417573d6000803e3d6000fd5b505050506040513d602081101561342d57600080fd5b50516001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691906135f0565b6001600160a01b039182166000908152601a6020908152604080832093909416825291909152205490565b601354600160201b900460020b81565b600a81815481106134ab57600080fd5b6000918252602090912001546001600160a01b0316905081565b6003546001600160a01b031681565b6005546001600160a01b03163314613519576040805162461bcd60e51b815260206004820152600360248201526228299960e91b604482015290519081900360640190fd5b6000841315613556576135566001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001633866135f0565b6000831315612d9a57612d9a6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001633856135f0565b6000828211156135ea576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052610f7b908490614a0b565b3390565b6001600160a01b03831661368b5760405162461bcd60e51b8152600401808060200182810382526024815260200180615d2c6024913960400191505060405180910390fd5b6001600160a01b0382166136d05760405162461bcd60e51b8152600401808060200182810382526022815260200180615be86022913960400191505060405180910390fd5b6001600160a01b038084166000818152601a6020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166137775760405162461bcd60e51b8152600401808060200182810382526025815260200180615d076025913960400191505060405180910390fd5b6001600160a01b0382166137bc5760405162461bcd60e51b8152600401808060200182810382526023815260200180615b836023913960400191505060405180910390fd5b6137c7838383610f7b565b61380481604051806060016040528060268152602001615c51602691396001600160a01b038616600090815260196020526040902054919061388f565b6001600160a01b0380851660009081526019602052604080822093909355908416815220546138339082613c2e565b6001600160a01b0380841660008181526019602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000818484111561391e5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156138e35781810151838201526020016138cb565b50505050905090810190601f1680156139105780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6000806139316111e7565b905060008111613945576000915050611125565b61394d614994565b801561395b5761395b614abc565b600191505090565b600082613972575060006111ae565b8282028284828161397f57fe5b04146112d55760405162461bcd60e51b8152600401808060200182810382526021815260200180615c9d6021913960400191505060405180910390fd5b6000808211613a12576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b818381613a1b57fe5b049392505050565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015613a9257600080fd5b505afa158015613aa6573d6000803e3d6000fd5b505050506040513d6020811015613abc57600080fd5b5051905090565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015613a9257600080fd5b6001600160a01b038216613b775760405162461bcd60e51b8152600401808060200182810382526021815260200180615ce66021913960400191505060405180910390fd5b613b8382600083610f7b565b613bc081604051806060016040528060228152602001615bc6602291396001600160a01b038516600090815260196020526040902054919061388f565b6001600160a01b038316600090815260196020526040902055601b54613be69082613593565b601b556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b6000828201838110156112d5576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6001600160a01b0381166000908152600b6020526040902060040154612b4457600a8054600181019091557fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a80180546001600160a01b0319166001600160a01b039290921691821790556000908152600b6020526040902043600490910155565b60215460009083908390613d48576040805162461bcd60e51b81526020600482015260016024820152600560fc1b604482015290519081900360640190fd5b60215481670de0b6b3a76400000281613d5d57fe5b04820192509250925092565b6000613d736111e7565b905060008111613d835750613ffa565b600a54604080516370a0823160e01b815230600482015290516000916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016916370a0823191602480820192602092909190829003018186803b158015613df057600080fd5b505afa158015613e04573d6000803e3d6000fd5b505050506040513d6020811015613e1a57600080fd5b5051604080516370a0823160e01b815230600482015290519192506000916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016916370a08231916024808301926020929190829003018186803b158015613e8857600080fd5b505afa158015613e9c573d6000803e3d6000fd5b505050506040513d6020811015613eb257600080fd5b5051905060005b83811015613ff4576000613ed3600a8381548110612f6857fe5b90508015613feb576000613eeb876113fa8785613963565b90506000613efd886113fa8786613963565b90506000613f0b8383613d09565b50509050838114613fe757613f41600a8681548110613f2657fe5b6000918252602090912001546001600160a01b031685613b32565b613f6c600a8681548110613f5157fe5b6000918252602090912001546001600160a01b031682614056565b82600b6000600a8881548110613f7e57fe5b60009182526020808320909101546001600160a01b03168352820192909252604001812060020191909155600a80548492600b92909189908110613fbe57fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020600301555b5050505b50600101613eb9565b50505050505b565b604080516001600160a01b0380861660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052612d9a908590614a0b565b6001600160a01b0382166140b1576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b6140bd60008383610f7b565b601b546140ca9082613c2e565b601b556001600160a01b0382166000908152601960205260409020546140f09082613c2e565b6001600160a01b03831660008181526019602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b60055460408051633850c7bd60e01b815290516000926001600160a01b031691633850c7bd9160048083019260e0929190829003018186803b15801561418d57600080fd5b505afa1580156141a1573d6000803e3d6000fd5b505050506040513d60e08110156141b757600080fd5b50602001519050600283900b1580156141d357508160020b6000145b1561427657600554600160d01b9004600290810b900b1580156142045750600554600160b81b9004600290810b900b155b1561420f5750612cba565b6005546002600160b81b8204810b600160d01b8304820b03810b81900591600160a01b9004810b9081810b90848401900b8161424757fe5b600554919005919091029350600160a01b9004600290810b9081810b90838503900b8161427057fe5b05029350505b614291838383600560149054906101000a900460020b614c4b565b6013546000906142b4906064906113fa90600160381b900460ff166113f4613a23565b905060006142db60646113fa601360079054906101000a900460ff1660ff166113f4613ac3565b90506142e985858484614e1f565b60006142f3613a23565b905060006142ff613ac3565b9050600061434f7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000008585614f9f565b905080614389576040805162461bcd60e51b815260206004820152600360248201526204d43760ec1b604482015290519081900360640190fd5b5050505050505050565b60008060008360020b126143aa578260020b6143b2565b8260020b6000035b9050620d89e88111156143f0576040805162461bcd60e51b81526020600482015260016024820152601560fa1b604482015290519081900360640190fd5b60006001821661440457600160801b614416565b6ffffcb933bd6fad37aa2d162d1a5940015b70ffffffffffffffffffffffffffffffffff169050600282161561444a576ffff97272373d413259a46990580e213a0260801c5b6004821615614469576ffff2e50f5f656932ef12357cf3c7fdcc0260801c5b6008821615614488576fffe5caca7e10e4e61c3624eaa0941cd00260801c5b60108216156144a7576fffcb9843d60f6159c9db58835c9266440260801c5b60208216156144c6576fff973b41fa98c081472e6896dfb254c00260801c5b60408216156144e5576fff2ea16466c96a3843ec78b326b528610260801c5b6080821615614504576ffe5dee046a99a2a811c461f1969c30530260801c5b610100821615614524576ffcbe86c7900a88aedcffc83b479aa3a40260801c5b610200821615614544576ff987a7253ac413176f2b074cf7815e540260801c5b610400821615614564576ff3392b0822b70005940c7a398e4b70f30260801c5b610800821615614584576fe7159475a2c29b7443b29c7fa6e889d90260801c5b6110008216156145a4576fd097f3bdfd2022b8845ad8f792aa58250260801c5b6120008216156145c4576fa9f746462d870fdf8a65dc1f90e061e50260801c5b6140008216156145e4576f70d869a156d2a1b890bb3df62baf32f70260801c5b618000821615614604576f31be135f97d08fd981231505542fcfa60260801c5b62010000821615614625576f09aa508b5b7a84e1c677de54f3e99bc90260801c5b62020000821615614645576e5d6af8dedb81196699c329225ee6040260801c5b62040000821615614664576d2216e584f5fa1ea926041bedfe980260801c5b62080000821615614681576b048a170391f7dc42444e8fa20260801c5b60008460020b131561469c57806000198161469857fe5b0490505b600160201b8106156146af5760016146b2565b60005b60ff16602082901c0192505050919050565b60008080600019858709868602925082811090839003039050806146fa57600084116146ef57600080fd5b5082900490506112d5565b80841161470657600080fd5b6000848688096000868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b600080600080600080614787308989615074565b6005546040805163514ea4bf60e01b81526004810184905290519293506001600160a01b039091169163514ea4bf9160248082019260a092909190829003018186803b1580156147d657600080fd5b505afa1580156147ea573d6000803e3d6000fd5b505050506040513d60a081101561480057600080fd5b508051602082015160408301516060840151608090940151929c919b50995091975095509350505050565b6000806000600560009054906101000a90046001600160a01b03166001600160a01b0316633850c7bd6040518163ffffffff1660e01b815260040160e06040518083038186803b15801561487e57600080fd5b505afa158015614892573d6000803e3d6000fd5b505050506040513d60e08110156148a857600080fd5b505190506148c8816148b988614393565b6148c288614393565b876150ca565b9250925050935093915050565b6000816001600160a01b031663db006a75846040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b15801561491d57600080fd5b505af1158015614931573d6000803e3d6000fd5b505050506040513d602081101561494757600080fd5b505160408051602081018390528181526024918101829052919250600080516020615d5083398151915291839181906060820190615c0a82396040019250505060405180910390a1505050565b60008061499f610f80565b9150915060008211806149b25750600081115b15614a0357614a037f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000008484615166565b612cba615491565b6000614a60826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661555a9092919063ffffffff16565b805190915015610f7b57808060200190516020811015614a7f57600080fd5b5051610f7b5760405162461bcd60e51b815260040180806020018281038252602a815260200180615d70602a913960400191505060405180910390fd5b6004546008546006546000918291614ad391613c2e565b600954600754614ae291613c2e565b6006546007546008546009546040805194855260208501939093528383019190915260608301525192945090925030917fb9a5bd9c6fc68ddbb2525a58437661292930f5a4adc18d5fa410ac494a15d0c39181900360800190a2600654600c54614b4b91613c2e565b600c55600754600d54614b5d91613c2e565b600d55600854600e54614b6f91613c2e565b600e55600954600f54614b8191613c2e565b600f5560006009819055600881905560078190556006819055808415613ff4578315614bca57614bb660646113fa8688613963565b601054909250614bc69083613c2e565b6010555b8215614bf357614bdf60646113fa8588613963565b601154909150614bef9082613c2e565b6011555b6000821180614c025750600081115b15613ff4576000614c138383613d09565b5050600254909150614c2e906001600160a01b031682614056565b600254614c43906001600160a01b0316613c88565b505050505050565b8260020b8460020b12614c8a576040805162461bcd60e51b8152602060048201526002602482015261563160f01b604482015290519081900360640190fd5b8160020b8460020b12614cc9576040805162461bcd60e51b81526020600482015260026024820152612b1960f11b604482015290519081900360640190fd5b8160020b8360020b13614d08576040805162461bcd60e51b8152602060048201526002602482015261563360f01b604482015290519081900360640190fd5b620d89e719600285900b1215614d4a576040805162461bcd60e51b8152602060048201526002602482015261158d60f21b604482015290519081900360640190fd5b620d89e8600284900b1315614d8b576040805162461bcd60e51b8152602060048201526002602482015261563560f01b604482015290519081900360640190fd5b8060020b8460020b81614d9a57fe5b0760020b15614dd5576040805162461bcd60e51b81526020600482015260026024820152612b1b60f11b604482015290519081900360640190fd5b8060020b8360020b81614de457fe5b0760020b15612d9a576040805162461bcd60e51b8152602060048201526002602482015261563760f01b604482015290519081900360640190fd5b60055460408051633850c7bd60e01b815290516000926001600160a01b031691633850c7bd9160048083019260e0929190829003018186803b158015614e6457600080fd5b505afa158015614e78573d6000803e3d6000fd5b505050506040513d60e0811015614e8e57600080fd5b505190506000614eb182614ea188614393565b614eaa88614393565b8787615571565b60055460408051633c8a7d8d60e01b815230600482015260028a810b602483015289900b60448201526001600160801b038416606482015260a06084820152600060a4820181905282519495506001600160a01b0390931693633c8a7d8d9360c480840194938390030190829087803b158015614f2d57600080fd5b505af1158015614f41573d6000803e3d6000fd5b505050506040513d6040811015614f5757600080fd5b505060058054600297880b62ffffff908116600160b81b0262ffffff60b81b199890990b16600160d01b0262ffffff60d01b1990911617959095169590951790935550505050565b601654601354601454601785905560188490556000926001600160a01b0390811692600160401b9004811691811690881683141561500857614fe086615635565b601554614ff6906001600160a01b031682612642565b5061500287828761178d565b50615066565b826001600160a01b0316876001600160a01b0316141561504d5761502b85615635565b601554615041906001600160a01b031683612642565b5061500288838861178d565b61505888838861178d565b5061506487828761178d565b505b506001979650505050505050565b6040805160609490941b6bffffffffffffffffffffffff1916602080860191909152600293840b60e890811b60348701529290930b90911b60378401528051808403601a018152603a9093019052815191012090565b600080836001600160a01b0316856001600160a01b031611156150eb579293925b846001600160a01b0316866001600160a01b0316116151165761510f85858561569c565b915061515d565b836001600160a01b0316866001600160a01b0316101561514f5761513b86858561569c565b9150615148858785615705565b905061515d565b61515a858585615705565b90505b94509492505050565b601654601354601454604080516370a0823160e01b815230600482015290516001600160a01b0394851694600160401b90940484169392831692600092908a16916370a0823191602480820192602092909190829003018186803b1580156151cd57600080fd5b505afa1580156151e1573d6000803e3d6000fd5b505050506040513d60208110156151f757600080fd5b5051604080516370a0823160e01b815230600482015290519192506000916001600160a01b038a16916370a08231916024808301926020929190829003018186803b15801561524557600080fd5b505afa158015615259573d6000803e3d6000fd5b505050506040513d602081101561526f57600080fd5b505190506001600160a01b0389811690861614156152a85761529187856148d5565b615299615748565b6152a386846148d5565b6152f2565b846001600160a01b0316886001600160a01b031614156152de576152cc86846148d5565b6152d4615748565b6152a387856148d5565b6152e886846148d5565b6152f287856148d5565b6000615377838b6001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561534557600080fd5b505afa158015615359573d6000803e3d6000fd5b505050506040513d602081101561536f57600080fd5b505190613593565b905060006153cc838b6001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561534557600080fd5b90506000821180156153df575060175482115b156153f6576017546153f2908390613593565b6008555b600081118015615407575060185481115b1561541e5760185461541a908290613593565b6009555b600854600954604080516020810193909352828101919091526060808352600e908301526d3932b6b7bb32b632b73234b7339d60911b6080830152517fb626a5b8d37253edd986272eff7a2541a1538cad44624e413d51713f0b2a414d9181900360a00190a15050505050505050505050565b600554600160b81b9004600290810b900b1580156154bd5750600554600160d01b9004600290810b900b155b156154c757613ffa565b6005546000906154ec90600160b81b8104600290810b91600160d01b9004900b614773565b5050600554929350600092839250829150819061551f90600160b81b8104600290810b91600160d01b9004900b8761579e565b93509350935093506000856001600160801b03161115613ff4576155438285613593565b6006556155508184613593565b6007555050505050565b60606155698484600085615910565b949350505050565b6000836001600160a01b0316856001600160a01b03161115615591579293925b846001600160a01b0316866001600160a01b0316116155bc576155b5858585615a60565b905061562c565b836001600160a01b0316866001600160a01b0316101561561e5760006155e3878686615a60565b905060006155f2878986615ac3565b9050806001600160801b0316826001600160801b0316106156135780615615565b815b9250505061562c565b615629858584615ac3565b90505b95945050505050565b8015612b445760165460408051632e1a7d4d60e01b81526004810184905290516001600160a01b0390921691632e1a7d4d9160248082019260009290919082900301818387803b15801561568857600080fd5b505af1158015613ff4573d6000803e3d6000fd5b6000826001600160a01b0316846001600160a01b031611156156bc579192915b836001600160a01b03166156f5606060ff16846001600160801b0316901b8686036001600160a01b0316866001600160a01b03166146c4565b816156fc57fe5b04949350505050565b6000826001600160a01b0316846001600160a01b03161115615725579192915b615569826001600160801b03168585036001600160a01b0316600160601b6146c4565b4715613ffa57601660009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0476040518263ffffffff1660e01b81526004016000604051808303818588803b15801561568857600080fd5b60008080806001600160801b03851615615850576005546040805163a34123a760e01b815260028a810b600483015289900b60248201526001600160801b038816604482015281516001600160a01b039093169263a34123a7926064808401939192918290030181600087803b15801561581757600080fd5b505af115801561582b573d6000803e3d6000fd5b505050506040513d604081101561584157600080fd5b50805160209091015190945092505b600554604080516309e3d67b60e31b815230600482015260028a810b602483015289900b60448201526001600160801b0360648201819052608482015281516001600160a01b0390931692634f1eb3d89260a4808401939192918290030181600087803b1580156158c057600080fd5b505af11580156158d4573d6000803e3d6000fd5b505050506040513d60408110156158ea57600080fd5b50805160209091015194989397506001600160801b039081169650909316935090915050565b6060824710156159515760405162461bcd60e51b8152600401808060200182810382526026815260200180615c776026913960400191505060405180910390fd5b61595a85615b00565b6159ab576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b600080866001600160a01b031685876040518082805190602001908083835b602083106159e95780518252601f1990920191602091820191016159ca565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114615a4b576040519150601f19603f3d011682016040523d82523d6000602084013e615a50565b606091505b5091509150612036828286615b06565b6000826001600160a01b0316846001600160a01b03161115615a80579192915b6000615aa3856001600160a01b0316856001600160a01b0316600160601b6146c4565b905061562c615abe84838888036001600160a01b03166146c4565b615b6c565b6000826001600160a01b0316846001600160a01b03161115615ae3579192915b615569615abe83600160601b8787036001600160a01b03166146c4565b3b151590565b60608315615b155750816112d5565b825115615b255782518084602001fd5b60405162461bcd60e51b81526020600482018181528451602484015284518593919283926044019190850190808383600083156138e35781810151838201526020016138cb565b806001600160801b038116811461263d57600080fdfe45524332303a207472616e7366657220746f20746865207a65726f20616464726573735265656e7472616e637947756172643a207265656e7472616e742063616c6c0045524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332303a20617070726f766520746f20746865207a65726f206164647265737349662074686973206973206e6f7420302c2074686572652077617320616e206572726f7245786368616e6765205261746520287363616c65642075702062792031653138293a2045524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e6365416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f7745524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f20616464726573738d1cced004452bd270777a8c670f9f7e7c4fdde56f2db331fe289d39dc2624ad5361666545524332303a204552433230206f7065726174696f6e20646964206e6f74207375636365656445524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220f7ac1a21dd7bd74b4f461c36c312235426f33889ebd49a8178ee19b5f660176764736f6c63430007060033"

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, _pool common.Address, _weth common.Address, _cToken0 common.Address, _cToken1 common.Address, _cEth common.Address, _protocolFee *big.Int, _maxTotalSupply *big.Int, _maxTwapDeviation *big.Int, _twapDuration uint32, _uniPortion uint8, _team common.Address) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ApiBin), backend, _pool, _weth, _cToken0, _cToken1, _cEth, _protocolFee, _maxTotalSupply, _maxTwapDeviation, _twapDuration, _uniPortion, _team)
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

// Assetholder is a free data retrieval call binding the contract method 0x5bb6aa85.
//
// Solidity: function Assetholder(address ) view returns(uint256 deposit0, uint256 deposit1, uint256 current0, uint256 current1, uint256 block)
func (_Api *ApiCaller) Assetholder(opts *bind.CallOpts, arg0 common.Address) (struct {
	Deposit0 *big.Int
	Deposit1 *big.Int
	Current0 *big.Int
	Current1 *big.Int
	Block    *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "Assetholder", arg0)

	outstruct := new(struct {
		Deposit0 *big.Int
		Deposit1 *big.Int
		Current0 *big.Int
		Current1 *big.Int
		Block    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Deposit0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Deposit1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Current0 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Current1 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Assetholder is a free data retrieval call binding the contract method 0x5bb6aa85.
//
// Solidity: function Assetholder(address ) view returns(uint256 deposit0, uint256 deposit1, uint256 current0, uint256 current1, uint256 block)
func (_Api *ApiSession) Assetholder(arg0 common.Address) (struct {
	Deposit0 *big.Int
	Deposit1 *big.Int
	Current0 *big.Int
	Current1 *big.Int
	Block    *big.Int
}, error) {
	return _Api.Contract.Assetholder(&_Api.CallOpts, arg0)
}

// Assetholder is a free data retrieval call binding the contract method 0x5bb6aa85.
//
// Solidity: function Assetholder(address ) view returns(uint256 deposit0, uint256 deposit1, uint256 current0, uint256 current1, uint256 block)
func (_Api *ApiCallerSession) Assetholder(arg0 common.Address) (struct {
	Deposit0 *big.Int
	Deposit1 *big.Int
	Current0 *big.Int
	Current1 *big.Int
	Block    *big.Int
}, error) {
	return _Api.Contract.Assetholder(&_Api.CallOpts, arg0)
}

// Fees is a free data retrieval call binding the contract method 0x3aaa36e6.
//
// Solidity: function Fees() view returns(uint256 U3Fees0, uint256 U3Fees1, uint256 LcFees0, uint256 LcFees1, uint256 AccruedProtocolFees0, uint256 AccruedProtocolFees1)
func (_Api *ApiCaller) Fees(opts *bind.CallOpts) (struct {
	U3Fees0              *big.Int
	U3Fees1              *big.Int
	LcFees0              *big.Int
	LcFees1              *big.Int
	AccruedProtocolFees0 *big.Int
	AccruedProtocolFees1 *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "Fees")

	outstruct := new(struct {
		U3Fees0              *big.Int
		U3Fees1              *big.Int
		LcFees0              *big.Int
		LcFees1              *big.Int
		AccruedProtocolFees0 *big.Int
		AccruedProtocolFees1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.U3Fees0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.U3Fees1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LcFees0 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LcFees1 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AccruedProtocolFees0 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AccruedProtocolFees1 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Fees is a free data retrieval call binding the contract method 0x3aaa36e6.
//
// Solidity: function Fees() view returns(uint256 U3Fees0, uint256 U3Fees1, uint256 LcFees0, uint256 LcFees1, uint256 AccruedProtocolFees0, uint256 AccruedProtocolFees1)
func (_Api *ApiSession) Fees() (struct {
	U3Fees0              *big.Int
	U3Fees1              *big.Int
	LcFees0              *big.Int
	LcFees1              *big.Int
	AccruedProtocolFees0 *big.Int
	AccruedProtocolFees1 *big.Int
}, error) {
	return _Api.Contract.Fees(&_Api.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x3aaa36e6.
//
// Solidity: function Fees() view returns(uint256 U3Fees0, uint256 U3Fees1, uint256 LcFees0, uint256 LcFees1, uint256 AccruedProtocolFees0, uint256 AccruedProtocolFees1)
func (_Api *ApiCallerSession) Fees() (struct {
	U3Fees0              *big.Int
	U3Fees1              *big.Int
	LcFees0              *big.Int
	LcFees1              *big.Int
	AccruedProtocolFees0 *big.Int
	AccruedProtocolFees1 *big.Int
}, error) {
	return _Api.Contract.Fees(&_Api.CallOpts)
}

// Accounts is a free data retrieval call binding the contract method 0xf2a40db8.
//
// Solidity: function accounts(uint256 ) view returns(address)
func (_Api *ApiCaller) Accounts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "accounts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Accounts is a free data retrieval call binding the contract method 0xf2a40db8.
//
// Solidity: function accounts(uint256 ) view returns(address)
func (_Api *ApiSession) Accounts(arg0 *big.Int) (common.Address, error) {
	return _Api.Contract.Accounts(&_Api.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xf2a40db8.
//
// Solidity: function accounts(uint256 ) view returns(address)
func (_Api *ApiCallerSession) Accounts(arg0 *big.Int) (common.Address, error) {
	return _Api.Contract.Accounts(&_Api.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Api *ApiCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Api *ApiSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Api.Contract.Allowance(&_Api.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Api *ApiCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Api.Contract.Allowance(&_Api.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Api *ApiCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Api *ApiSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Api.Contract.BalanceOf(&_Api.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Api *ApiCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Api.Contract.BalanceOf(&_Api.CallOpts, account)
}

// CHigh is a free data retrieval call binding the contract method 0x8e843c6c.
//
// Solidity: function cHigh() view returns(int24)
func (_Api *ApiCaller) CHigh(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "cHigh")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHigh is a free data retrieval call binding the contract method 0x8e843c6c.
//
// Solidity: function cHigh() view returns(int24)
func (_Api *ApiSession) CHigh() (*big.Int, error) {
	return _Api.Contract.CHigh(&_Api.CallOpts)
}

// CHigh is a free data retrieval call binding the contract method 0x8e843c6c.
//
// Solidity: function cHigh() view returns(int24)
func (_Api *ApiCallerSession) CHigh() (*big.Int, error) {
	return _Api.Contract.CHigh(&_Api.CallOpts)
}

// CLow is a free data retrieval call binding the contract method 0xd368867f.
//
// Solidity: function cLow() view returns(int24)
func (_Api *ApiCaller) CLow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "cLow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CLow is a free data retrieval call binding the contract method 0xd368867f.
//
// Solidity: function cLow() view returns(int24)
func (_Api *ApiSession) CLow() (*big.Int, error) {
	return _Api.Contract.CLow(&_Api.CallOpts)
}

// CLow is a free data retrieval call binding the contract method 0xd368867f.
//
// Solidity: function cLow() view returns(int24)
func (_Api *ApiCallerSession) CLow() (*big.Int, error) {
	return _Api.Contract.CLow(&_Api.CallOpts)
}

// CurComp0 is a free data retrieval call binding the contract method 0xc8a746a6.
//
// Solidity: function curComp0() view returns(uint256)
func (_Api *ApiCaller) CurComp0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "curComp0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurComp0 is a free data retrieval call binding the contract method 0xc8a746a6.
//
// Solidity: function curComp0() view returns(uint256)
func (_Api *ApiSession) CurComp0() (*big.Int, error) {
	return _Api.Contract.CurComp0(&_Api.CallOpts)
}

// CurComp0 is a free data retrieval call binding the contract method 0xc8a746a6.
//
// Solidity: function curComp0() view returns(uint256)
func (_Api *ApiCallerSession) CurComp0() (*big.Int, error) {
	return _Api.Contract.CurComp0(&_Api.CallOpts)
}

// CurComp1 is a free data retrieval call binding the contract method 0xccac7d8b.
//
// Solidity: function curComp1() view returns(uint256)
func (_Api *ApiCaller) CurComp1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "curComp1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurComp1 is a free data retrieval call binding the contract method 0xccac7d8b.
//
// Solidity: function curComp1() view returns(uint256)
func (_Api *ApiSession) CurComp1() (*big.Int, error) {
	return _Api.Contract.CurComp1(&_Api.CallOpts)
}

// CurComp1 is a free data retrieval call binding the contract method 0xccac7d8b.
//
// Solidity: function curComp1() view returns(uint256)
func (_Api *ApiCallerSession) CurComp1() (*big.Int, error) {
	return _Api.Contract.CurComp1(&_Api.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Api *ApiCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Api *ApiSession) Decimals() (uint8, error) {
	return _Api.Contract.Decimals(&_Api.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Api *ApiCallerSession) Decimals() (uint8, error) {
	return _Api.Contract.Decimals(&_Api.CallOpts)
}

// GetCAmounts is a free data retrieval call binding the contract method 0x06b0b656.
//
// Solidity: function getCAmounts() view returns(uint256 amountA, uint256 amountB)
func (_Api *ApiCaller) GetCAmounts(opts *bind.CallOpts) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getCAmounts")

	outstruct := new(struct {
		AmountA *big.Int
		AmountB *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCAmounts is a free data retrieval call binding the contract method 0x06b0b656.
//
// Solidity: function getCAmounts() view returns(uint256 amountA, uint256 amountB)
func (_Api *ApiSession) GetCAmounts() (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	return _Api.Contract.GetCAmounts(&_Api.CallOpts)
}

// GetCAmounts is a free data retrieval call binding the contract method 0x06b0b656.
//
// Solidity: function getCAmounts() view returns(uint256 amountA, uint256 amountB)
func (_Api *ApiCallerSession) GetCAmounts() (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	return _Api.Contract.GetCAmounts(&_Api.CallOpts)
}

// GetETHBalance is a free data retrieval call binding the contract method 0x6e947298.
//
// Solidity: function getETHBalance() view returns(uint256)
func (_Api *ApiCaller) GetETHBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getETHBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetETHBalance is a free data retrieval call binding the contract method 0x6e947298.
//
// Solidity: function getETHBalance() view returns(uint256)
func (_Api *ApiSession) GetETHBalance() (*big.Int, error) {
	return _Api.Contract.GetETHBalance(&_Api.CallOpts)
}

// GetETHBalance is a free data retrieval call binding the contract method 0x6e947298.
//
// Solidity: function getETHBalance() view returns(uint256)
func (_Api *ApiCallerSession) GetETHBalance() (*big.Int, error) {
	return _Api.Contract.GetETHBalance(&_Api.CallOpts)
}

// GetPositionAmounts is a free data retrieval call binding the contract method 0xa91ef6eb.
//
// Solidity: function getPositionAmounts(int24 tickLower, int24 tickUpper) view returns(uint256 amount0, uint256 amount1)
func (_Api *ApiCaller) GetPositionAmounts(opts *bind.CallOpts, tickLower *big.Int, tickUpper *big.Int) (struct {
	Amount0 *big.Int
	Amount1 *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getPositionAmounts", tickLower, tickUpper)

	outstruct := new(struct {
		Amount0 *big.Int
		Amount1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPositionAmounts is a free data retrieval call binding the contract method 0xa91ef6eb.
//
// Solidity: function getPositionAmounts(int24 tickLower, int24 tickUpper) view returns(uint256 amount0, uint256 amount1)
func (_Api *ApiSession) GetPositionAmounts(tickLower *big.Int, tickUpper *big.Int) (struct {
	Amount0 *big.Int
	Amount1 *big.Int
}, error) {
	return _Api.Contract.GetPositionAmounts(&_Api.CallOpts, tickLower, tickUpper)
}

// GetPositionAmounts is a free data retrieval call binding the contract method 0xa91ef6eb.
//
// Solidity: function getPositionAmounts(int24 tickLower, int24 tickUpper) view returns(uint256 amount0, uint256 amount1)
func (_Api *ApiCallerSession) GetPositionAmounts(tickLower *big.Int, tickUpper *big.Int) (struct {
	Amount0 *big.Int
	Amount1 *big.Int
}, error) {
	return _Api.Contract.GetPositionAmounts(&_Api.CallOpts, tickLower, tickUpper)
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

// GetSSLiquidity is a free data retrieval call binding the contract method 0xcda206b0.
//
// Solidity: function getSSLiquidity(int24 tickLower, int24 tickUpper) view returns(uint128 liquidity)
func (_Api *ApiCaller) GetSSLiquidity(opts *bind.CallOpts, tickLower *big.Int, tickUpper *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getSSLiquidity", tickLower, tickUpper)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSSLiquidity is a free data retrieval call binding the contract method 0xcda206b0.
//
// Solidity: function getSSLiquidity(int24 tickLower, int24 tickUpper) view returns(uint128 liquidity)
func (_Api *ApiSession) GetSSLiquidity(tickLower *big.Int, tickUpper *big.Int) (*big.Int, error) {
	return _Api.Contract.GetSSLiquidity(&_Api.CallOpts, tickLower, tickUpper)
}

// GetSSLiquidity is a free data retrieval call binding the contract method 0xcda206b0.
//
// Solidity: function getSSLiquidity(int24 tickLower, int24 tickUpper) view returns(uint128 liquidity)
func (_Api *ApiCallerSession) GetSSLiquidity(tickLower *big.Int, tickUpper *big.Int) (*big.Int, error) {
	return _Api.Contract.GetSSLiquidity(&_Api.CallOpts, tickLower, tickUpper)
}

// GetStoredAssets is a free data retrieval call binding the contract method 0xa4ee141b.
//
// Solidity: function getStoredAssets(address who) view returns(uint256, uint256, uint256, uint256)
func (_Api *ApiCaller) GetStoredAssets(opts *bind.CallOpts, who common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStoredAssets", who)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetStoredAssets is a free data retrieval call binding the contract method 0xa4ee141b.
//
// Solidity: function getStoredAssets(address who) view returns(uint256, uint256, uint256, uint256)
func (_Api *ApiSession) GetStoredAssets(who common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Api.Contract.GetStoredAssets(&_Api.CallOpts, who)
}

// GetStoredAssets is a free data retrieval call binding the contract method 0xa4ee141b.
//
// Solidity: function getStoredAssets(address who) view returns(uint256, uint256, uint256, uint256)
func (_Api *ApiCallerSession) GetStoredAssets(who common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Api.Contract.GetStoredAssets(&_Api.CallOpts, who)
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

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Api *ApiCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Api *ApiSession) Governance() (common.Address, error) {
	return _Api.Contract.Governance(&_Api.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Api *ApiCallerSession) Governance() (common.Address, error) {
	return _Api.Contract.Governance(&_Api.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_Api *ApiCaller) MaxTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "maxTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_Api *ApiSession) MaxTotalSupply() (*big.Int, error) {
	return _Api.Contract.MaxTotalSupply(&_Api.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_Api *ApiCallerSession) MaxTotalSupply() (*big.Int, error) {
	return _Api.Contract.MaxTotalSupply(&_Api.CallOpts)
}

// MaxTwapDeviation is a free data retrieval call binding the contract method 0xe7c7cb91.
//
// Solidity: function maxTwapDeviation() view returns(int24)
func (_Api *ApiCaller) MaxTwapDeviation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "maxTwapDeviation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTwapDeviation is a free data retrieval call binding the contract method 0xe7c7cb91.
//
// Solidity: function maxTwapDeviation() view returns(int24)
func (_Api *ApiSession) MaxTwapDeviation() (*big.Int, error) {
	return _Api.Contract.MaxTwapDeviation(&_Api.CallOpts)
}

// MaxTwapDeviation is a free data retrieval call binding the contract method 0xe7c7cb91.
//
// Solidity: function maxTwapDeviation() view returns(int24)
func (_Api *ApiCallerSession) MaxTwapDeviation() (*big.Int, error) {
	return _Api.Contract.MaxTwapDeviation(&_Api.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Api *ApiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Api *ApiSession) Name() (string, error) {
	return _Api.Contract.Name(&_Api.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Api *ApiCallerSession) Name() (string, error) {
	return _Api.Contract.Name(&_Api.CallOpts)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_Api *ApiCaller) PendingGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "pendingGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_Api *ApiSession) PendingGovernance() (common.Address, error) {
	return _Api.Contract.PendingGovernance(&_Api.CallOpts)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_Api *ApiCallerSession) PendingGovernance() (common.Address, error) {
	return _Api.Contract.PendingGovernance(&_Api.CallOpts)
}

// PoolAddress is a free data retrieval call binding the contract method 0x1755ff21.
//
// Solidity: function poolAddress() view returns(address)
func (_Api *ApiCaller) PoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "poolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolAddress is a free data retrieval call binding the contract method 0x1755ff21.
//
// Solidity: function poolAddress() view returns(address)
func (_Api *ApiSession) PoolAddress() (common.Address, error) {
	return _Api.Contract.PoolAddress(&_Api.CallOpts)
}

// PoolAddress is a free data retrieval call binding the contract method 0x1755ff21.
//
// Solidity: function poolAddress() view returns(address)
func (_Api *ApiCallerSession) PoolAddress() (common.Address, error) {
	return _Api.Contract.PoolAddress(&_Api.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_Api *ApiCaller) ProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "protocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_Api *ApiSession) ProtocolFee() (*big.Int, error) {
	return _Api.Contract.ProtocolFee(&_Api.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_Api *ApiCallerSession) ProtocolFee() (*big.Int, error) {
	return _Api.Contract.ProtocolFee(&_Api.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Api *ApiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Api *ApiSession) Symbol() (string, error) {
	return _Api.Contract.Symbol(&_Api.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Api *ApiCallerSession) Symbol() (string, error) {
	return _Api.Contract.Symbol(&_Api.CallOpts)
}

// Team is a free data retrieval call binding the contract method 0x85f2aef2.
//
// Solidity: function team() view returns(address)
func (_Api *ApiCaller) Team(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "team")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Team is a free data retrieval call binding the contract method 0x85f2aef2.
//
// Solidity: function team() view returns(address)
func (_Api *ApiSession) Team() (common.Address, error) {
	return _Api.Contract.Team(&_Api.CallOpts)
}

// Team is a free data retrieval call binding the contract method 0x85f2aef2.
//
// Solidity: function team() view returns(address)
func (_Api *ApiCallerSession) Team() (common.Address, error) {
	return _Api.Contract.Team(&_Api.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Api *ApiCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Api *ApiSession) Token0() (common.Address, error) {
	return _Api.Contract.Token0(&_Api.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Api *ApiCallerSession) Token0() (common.Address, error) {
	return _Api.Contract.Token0(&_Api.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Api *ApiCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Api *ApiSession) Token1() (common.Address, error) {
	return _Api.Contract.Token1(&_Api.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Api *ApiCallerSession) Token1() (common.Address, error) {
	return _Api.Contract.Token1(&_Api.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Api *ApiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Api *ApiSession) TotalSupply() (*big.Int, error) {
	return _Api.Contract.TotalSupply(&_Api.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Api *ApiCallerSession) TotalSupply() (*big.Int, error) {
	return _Api.Contract.TotalSupply(&_Api.CallOpts)
}

// TwapDuration is a free data retrieval call binding the contract method 0x26d89545.
//
// Solidity: function twapDuration() view returns(uint32)
func (_Api *ApiCaller) TwapDuration(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "twapDuration")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TwapDuration is a free data retrieval call binding the contract method 0x26d89545.
//
// Solidity: function twapDuration() view returns(uint32)
func (_Api *ApiSession) TwapDuration() (uint32, error) {
	return _Api.Contract.TwapDuration(&_Api.CallOpts)
}

// TwapDuration is a free data retrieval call binding the contract method 0x26d89545.
//
// Solidity: function twapDuration() view returns(uint32)
func (_Api *ApiCallerSession) TwapDuration() (uint32, error) {
	return _Api.Contract.TwapDuration(&_Api.CallOpts)
}

// UniPortion is a free data retrieval call binding the contract method 0x31dc5b95.
//
// Solidity: function uniPortion() view returns(uint8)
func (_Api *ApiCaller) UniPortion(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "uniPortion")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// UniPortion is a free data retrieval call binding the contract method 0x31dc5b95.
//
// Solidity: function uniPortion() view returns(uint8)
func (_Api *ApiSession) UniPortion() (uint8, error) {
	return _Api.Contract.UniPortion(&_Api.CallOpts)
}

// UniPortion is a free data retrieval call binding the contract method 0x31dc5b95.
//
// Solidity: function uniPortion() view returns(uint8)
func (_Api *ApiCallerSession) UniPortion() (uint8, error) {
	return _Api.Contract.UniPortion(&_Api.CallOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Api *ApiTransactor) AcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "acceptGovernance")
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Api *ApiSession) AcceptGovernance() (*types.Transaction, error) {
	return _Api.Contract.AcceptGovernance(&_Api.TransactOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Api *ApiTransactorSession) AcceptGovernance() (*types.Transaction, error) {
	return _Api.Contract.AcceptGovernance(&_Api.TransactOpts)
}

// Alloc is a paid mutator transaction binding the contract method 0xbbc001c3.
//
// Solidity: function alloc() returns()
func (_Api *ApiTransactor) Alloc(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "alloc")
}

// Alloc is a paid mutator transaction binding the contract method 0xbbc001c3.
//
// Solidity: function alloc() returns()
func (_Api *ApiSession) Alloc() (*types.Transaction, error) {
	return _Api.Contract.Alloc(&_Api.TransactOpts)
}

// Alloc is a paid mutator transaction binding the contract method 0xbbc001c3.
//
// Solidity: function alloc() returns()
func (_Api *ApiTransactorSession) Alloc() (*types.Transaction, error) {
	return _Api.Contract.Alloc(&_Api.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Api *ApiTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Api *ApiSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Approve(&_Api.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Api *ApiTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Approve(&_Api.TransactOpts, spender, amount)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x0430c130.
//
// Solidity: function collectProtocol(uint256 amount0, uint256 amount1, address to) returns()
func (_Api *ApiTransactor) CollectProtocol(opts *bind.TransactOpts, amount0 *big.Int, amount1 *big.Int, to common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "collectProtocol", amount0, amount1, to)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x0430c130.
//
// Solidity: function collectProtocol(uint256 amount0, uint256 amount1, address to) returns()
func (_Api *ApiSession) CollectProtocol(amount0 *big.Int, amount1 *big.Int, to common.Address) (*types.Transaction, error) {
	return _Api.Contract.CollectProtocol(&_Api.TransactOpts, amount0, amount1, to)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x0430c130.
//
// Solidity: function collectProtocol(uint256 amount0, uint256 amount1, address to) returns()
func (_Api *ApiTransactorSession) CollectProtocol(amount0 *big.Int, amount1 *big.Int, to common.Address) (*types.Transaction, error) {
	return _Api.Contract.CollectProtocol(&_Api.TransactOpts, amount0, amount1, to)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Api *ApiTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Api *ApiSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Api.Contract.DecreaseAllowance(&_Api.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Api *ApiTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Api.Contract.DecreaseAllowance(&_Api.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 amountToken0, uint256 amountToken1, bool doRebalence) returns()
func (_Api *ApiTransactor) Deposit(opts *bind.TransactOpts, amountToken0 *big.Int, amountToken1 *big.Int, doRebalence bool) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "deposit", amountToken0, amountToken1, doRebalence)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 amountToken0, uint256 amountToken1, bool doRebalence) returns()
func (_Api *ApiSession) Deposit(amountToken0 *big.Int, amountToken1 *big.Int, doRebalence bool) (*types.Transaction, error) {
	return _Api.Contract.Deposit(&_Api.TransactOpts, amountToken0, amountToken1, doRebalence)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 amountToken0, uint256 amountToken1, bool doRebalence) returns()
func (_Api *ApiTransactorSession) Deposit(amountToken0 *big.Int, amountToken1 *big.Int, doRebalence bool) (*types.Transaction, error) {
	return _Api.Contract.Deposit(&_Api.TransactOpts, amountToken0, amountToken1, doRebalence)
}

// EmergencyBurn is a paid mutator transaction binding the contract method 0xd8d7f96f.
//
// Solidity: function emergencyBurn() returns()
func (_Api *ApiTransactor) EmergencyBurn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "emergencyBurn")
}

// EmergencyBurn is a paid mutator transaction binding the contract method 0xd8d7f96f.
//
// Solidity: function emergencyBurn() returns()
func (_Api *ApiSession) EmergencyBurn() (*types.Transaction, error) {
	return _Api.Contract.EmergencyBurn(&_Api.TransactOpts)
}

// EmergencyBurn is a paid mutator transaction binding the contract method 0xd8d7f96f.
//
// Solidity: function emergencyBurn() returns()
func (_Api *ApiTransactorSession) EmergencyBurn() (*types.Transaction, error) {
	return _Api.Contract.EmergencyBurn(&_Api.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Api *ApiTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Api *ApiSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Api.Contract.IncreaseAllowance(&_Api.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Api *ApiTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Api.Contract.IncreaseAllowance(&_Api.TransactOpts, spender, addedValue)
}

// RemoveCTokens is a paid mutator transaction binding the contract method 0xd17cfebd.
//
// Solidity: function removeCTokens() returns()
func (_Api *ApiTransactor) RemoveCTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "removeCTokens")
}

// RemoveCTokens is a paid mutator transaction binding the contract method 0xd17cfebd.
//
// Solidity: function removeCTokens() returns()
func (_Api *ApiSession) RemoveCTokens() (*types.Transaction, error) {
	return _Api.Contract.RemoveCTokens(&_Api.TransactOpts)
}

// RemoveCTokens is a paid mutator transaction binding the contract method 0xd17cfebd.
//
// Solidity: function removeCTokens() returns()
func (_Api *ApiTransactorSession) RemoveCTokens() (*types.Transaction, error) {
	return _Api.Contract.RemoveCTokens(&_Api.TransactOpts)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_Api *ApiTransactor) SetGovernance(opts *bind.TransactOpts, _governance common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setGovernance", _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_Api *ApiSession) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetGovernance(&_Api.TransactOpts, _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_Api *ApiTransactorSession) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetGovernance(&_Api.TransactOpts, _governance)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x3f3e4c11.
//
// Solidity: function setMaxTotalSupply(uint256 newMax) returns()
func (_Api *ApiTransactor) SetMaxTotalSupply(opts *bind.TransactOpts, newMax *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setMaxTotalSupply", newMax)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x3f3e4c11.
//
// Solidity: function setMaxTotalSupply(uint256 newMax) returns()
func (_Api *ApiSession) SetMaxTotalSupply(newMax *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetMaxTotalSupply(&_Api.TransactOpts, newMax)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x3f3e4c11.
//
// Solidity: function setMaxTotalSupply(uint256 newMax) returns()
func (_Api *ApiTransactorSession) SetMaxTotalSupply(newMax *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetMaxTotalSupply(&_Api.TransactOpts, newMax)
}

// SetMaxTwapDeviation is a paid mutator transaction binding the contract method 0x3cbff3fe.
//
// Solidity: function setMaxTwapDeviation(int24 _maxTwapDeviation) returns()
func (_Api *ApiTransactor) SetMaxTwapDeviation(opts *bind.TransactOpts, _maxTwapDeviation *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setMaxTwapDeviation", _maxTwapDeviation)
}

// SetMaxTwapDeviation is a paid mutator transaction binding the contract method 0x3cbff3fe.
//
// Solidity: function setMaxTwapDeviation(int24 _maxTwapDeviation) returns()
func (_Api *ApiSession) SetMaxTwapDeviation(_maxTwapDeviation *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetMaxTwapDeviation(&_Api.TransactOpts, _maxTwapDeviation)
}

// SetMaxTwapDeviation is a paid mutator transaction binding the contract method 0x3cbff3fe.
//
// Solidity: function setMaxTwapDeviation(int24 _maxTwapDeviation) returns()
func (_Api *ApiTransactorSession) SetMaxTwapDeviation(_maxTwapDeviation *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetMaxTwapDeviation(&_Api.TransactOpts, _maxTwapDeviation)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 _protocolFee) returns()
func (_Api *ApiTransactor) SetProtocolFee(opts *bind.TransactOpts, _protocolFee *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setProtocolFee", _protocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 _protocolFee) returns()
func (_Api *ApiSession) SetProtocolFee(_protocolFee *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetProtocolFee(&_Api.TransactOpts, _protocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 _protocolFee) returns()
func (_Api *ApiTransactorSession) SetProtocolFee(_protocolFee *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetProtocolFee(&_Api.TransactOpts, _protocolFee)
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

// SetTwapDuration is a paid mutator transaction binding the contract method 0xc433c80a.
//
// Solidity: function setTwapDuration(uint32 _twapDuration) returns()
func (_Api *ApiTransactor) SetTwapDuration(opts *bind.TransactOpts, _twapDuration uint32) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setTwapDuration", _twapDuration)
}

// SetTwapDuration is a paid mutator transaction binding the contract method 0xc433c80a.
//
// Solidity: function setTwapDuration(uint32 _twapDuration) returns()
func (_Api *ApiSession) SetTwapDuration(_twapDuration uint32) (*types.Transaction, error) {
	return _Api.Contract.SetTwapDuration(&_Api.TransactOpts, _twapDuration)
}

// SetTwapDuration is a paid mutator transaction binding the contract method 0xc433c80a.
//
// Solidity: function setTwapDuration(uint32 _twapDuration) returns()
func (_Api *ApiTransactorSession) SetTwapDuration(_twapDuration uint32) (*types.Transaction, error) {
	return _Api.Contract.SetTwapDuration(&_Api.TransactOpts, _twapDuration)
}

// SetUniPortionRatio is a paid mutator transaction binding the contract method 0x624a177a.
//
// Solidity: function setUniPortionRatio(uint8 ratio) returns()
func (_Api *ApiTransactor) SetUniPortionRatio(opts *bind.TransactOpts, ratio uint8) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setUniPortionRatio", ratio)
}

// SetUniPortionRatio is a paid mutator transaction binding the contract method 0x624a177a.
//
// Solidity: function setUniPortionRatio(uint8 ratio) returns()
func (_Api *ApiSession) SetUniPortionRatio(ratio uint8) (*types.Transaction, error) {
	return _Api.Contract.SetUniPortionRatio(&_Api.TransactOpts, ratio)
}

// SetUniPortionRatio is a paid mutator transaction binding the contract method 0x624a177a.
//
// Solidity: function setUniPortionRatio(uint8 ratio) returns()
func (_Api *ApiTransactorSession) SetUniPortionRatio(ratio uint8) (*types.Transaction, error) {
	return _Api.Contract.SetUniPortionRatio(&_Api.TransactOpts, ratio)
}

// Strategy1 is a paid mutator transaction binding the contract method 0x6253c718.
//
// Solidity: function strategy1(int24 newLow, int24 newHigh) returns()
func (_Api *ApiTransactor) Strategy1(opts *bind.TransactOpts, newLow *big.Int, newHigh *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "strategy1", newLow, newHigh)
}

// Strategy1 is a paid mutator transaction binding the contract method 0x6253c718.
//
// Solidity: function strategy1(int24 newLow, int24 newHigh) returns()
func (_Api *ApiSession) Strategy1(newLow *big.Int, newHigh *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Strategy1(&_Api.TransactOpts, newLow, newHigh)
}

// Strategy1 is a paid mutator transaction binding the contract method 0x6253c718.
//
// Solidity: function strategy1(int24 newLow, int24 newHigh) returns()
func (_Api *ApiTransactorSession) Strategy1(newLow *big.Int, newHigh *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Strategy1(&_Api.TransactOpts, newLow, newHigh)
}

// SupplyErc20ToCompound is a paid mutator transaction binding the contract method 0x3b7ba25f.
//
// Solidity: function supplyErc20ToCompound(address _erc20Contract, address _cErc20Contract, uint256 _numTokensToSupply) returns(uint256)
func (_Api *ApiTransactor) SupplyErc20ToCompound(opts *bind.TransactOpts, _erc20Contract common.Address, _cErc20Contract common.Address, _numTokensToSupply *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "supplyErc20ToCompound", _erc20Contract, _cErc20Contract, _numTokensToSupply)
}

// SupplyErc20ToCompound is a paid mutator transaction binding the contract method 0x3b7ba25f.
//
// Solidity: function supplyErc20ToCompound(address _erc20Contract, address _cErc20Contract, uint256 _numTokensToSupply) returns(uint256)
func (_Api *ApiSession) SupplyErc20ToCompound(_erc20Contract common.Address, _cErc20Contract common.Address, _numTokensToSupply *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SupplyErc20ToCompound(&_Api.TransactOpts, _erc20Contract, _cErc20Contract, _numTokensToSupply)
}

// SupplyErc20ToCompound is a paid mutator transaction binding the contract method 0x3b7ba25f.
//
// Solidity: function supplyErc20ToCompound(address _erc20Contract, address _cErc20Contract, uint256 _numTokensToSupply) returns(uint256)
func (_Api *ApiTransactorSession) SupplyErc20ToCompound(_erc20Contract common.Address, _cErc20Contract common.Address, _numTokensToSupply *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SupplyErc20ToCompound(&_Api.TransactOpts, _erc20Contract, _cErc20Contract, _numTokensToSupply)
}

// SupplyEthToCompound is a paid mutator transaction binding the contract method 0x73232c60.
//
// Solidity: function supplyEthToCompound(address _cEtherContract, address _ctoken) payable returns(bool)
func (_Api *ApiTransactor) SupplyEthToCompound(opts *bind.TransactOpts, _cEtherContract common.Address, _ctoken common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "supplyEthToCompound", _cEtherContract, _ctoken)
}

// SupplyEthToCompound is a paid mutator transaction binding the contract method 0x73232c60.
//
// Solidity: function supplyEthToCompound(address _cEtherContract, address _ctoken) payable returns(bool)
func (_Api *ApiSession) SupplyEthToCompound(_cEtherContract common.Address, _ctoken common.Address) (*types.Transaction, error) {
	return _Api.Contract.SupplyEthToCompound(&_Api.TransactOpts, _cEtherContract, _ctoken)
}

// SupplyEthToCompound is a paid mutator transaction binding the contract method 0x73232c60.
//
// Solidity: function supplyEthToCompound(address _cEtherContract, address _ctoken) payable returns(bool)
func (_Api *ApiTransactorSession) SupplyEthToCompound(_cEtherContract common.Address, _ctoken common.Address) (*types.Transaction, error) {
	return _Api.Contract.SupplyEthToCompound(&_Api.TransactOpts, _cEtherContract, _ctoken)
}

// Swap is a paid mutator transaction binding the contract method 0x68a5787a.
//
// Solidity: function swap(int256 swapAmount, bool zeroForOne, uint160 sqrtPriceLimitX96) returns(int256, int256)
func (_Api *ApiTransactor) Swap(opts *bind.TransactOpts, swapAmount *big.Int, zeroForOne bool, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "swap", swapAmount, zeroForOne, sqrtPriceLimitX96)
}

// Swap is a paid mutator transaction binding the contract method 0x68a5787a.
//
// Solidity: function swap(int256 swapAmount, bool zeroForOne, uint160 sqrtPriceLimitX96) returns(int256, int256)
func (_Api *ApiSession) Swap(swapAmount *big.Int, zeroForOne bool, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Swap(&_Api.TransactOpts, swapAmount, zeroForOne, sqrtPriceLimitX96)
}

// Swap is a paid mutator transaction binding the contract method 0x68a5787a.
//
// Solidity: function swap(int256 swapAmount, bool zeroForOne, uint160 sqrtPriceLimitX96) returns(int256, int256)
func (_Api *ApiTransactorSession) Swap(swapAmount *big.Int, zeroForOne bool, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Swap(&_Api.TransactOpts, swapAmount, zeroForOne, sqrtPriceLimitX96)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Api *ApiTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Api *ApiSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Transfer(&_Api.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Api *ApiTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Transfer(&_Api.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Api *ApiTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Api *ApiSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.TransferFrom(&_Api.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Api *ApiTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.TransferFrom(&_Api.TransactOpts, sender, recipient, amount)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_Api *ApiTransactor) UniswapV3MintCallback(opts *bind.TransactOpts, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "uniswapV3MintCallback", amount0, amount1, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_Api *ApiSession) UniswapV3MintCallback(amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.Contract.UniswapV3MintCallback(&_Api.TransactOpts, amount0, amount1, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_Api *ApiTransactorSession) UniswapV3MintCallback(amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.Contract.UniswapV3MintCallback(&_Api.TransactOpts, amount0, amount1, data)
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

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 percent) returns(uint256 amount0, uint256 amount1)
func (_Api *ApiTransactor) Withdraw(opts *bind.TransactOpts, percent *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "withdraw", percent)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 percent) returns(uint256 amount0, uint256 amount1)
func (_Api *ApiSession) Withdraw(percent *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Withdraw(&_Api.TransactOpts, percent)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 percent) returns(uint256 amount0, uint256 amount1)
func (_Api *ApiTransactorSession) Withdraw(percent *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Withdraw(&_Api.TransactOpts, percent)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Api *ApiTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Api.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Api *ApiSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Api.Contract.Fallback(&_Api.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Api *ApiTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Api.Contract.Fallback(&_Api.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Api *ApiTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Api *ApiSession) Receive() (*types.Transaction, error) {
	return _Api.Contract.Receive(&_Api.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Api *ApiTransactorSession) Receive() (*types.Transaction, error) {
	return _Api.Contract.Receive(&_Api.TransactOpts)
}

// ApiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Api contract.
type ApiApprovalIterator struct {
	Event *ApiApproval // Event containing the contract specifics and raw log

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
func (it *ApiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiApproval)
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
		it.Event = new(ApiApproval)
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
func (it *ApiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiApproval represents a Approval event raised by the Api contract.
type ApiApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Api *ApiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ApiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ApiApprovalIterator{contract: _Api.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Api *ApiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ApiApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiApproval)
				if err := _Api.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Api *ApiFilterer) ParseApproval(log types.Log) (*ApiApproval, error) {
	event := new(ApiApproval)
	if err := _Api.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiCollectFeesIterator is returned from FilterCollectFees and is used to iterate over the raw logs and unpacked data for CollectFees events raised by the Api contract.
type ApiCollectFeesIterator struct {
	Event *ApiCollectFees // Event containing the contract specifics and raw log

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
func (it *ApiCollectFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiCollectFees)
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
		it.Event = new(ApiCollectFees)
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
func (it *ApiCollectFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiCollectFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiCollectFees represents a CollectFees event raised by the Api contract.
type ApiCollectFees struct {
	Maker  common.Address
	UFees0 *big.Int
	UFees1 *big.Int
	LFees0 *big.Int
	LFees1 *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCollectFees is a free log retrieval operation binding the contract event 0xb9a5bd9c6fc68ddbb2525a58437661292930f5a4adc18d5fa410ac494a15d0c3.
//
// Solidity: event CollectFees(address indexed maker, uint256 uFees0, uint256 uFees1, uint256 lFees0, uint256 lFees1)
func (_Api *ApiFilterer) FilterCollectFees(opts *bind.FilterOpts, maker []common.Address) (*ApiCollectFeesIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "CollectFees", makerRule)
	if err != nil {
		return nil, err
	}
	return &ApiCollectFeesIterator{contract: _Api.contract, event: "CollectFees", logs: logs, sub: sub}, nil
}

// WatchCollectFees is a free log subscription operation binding the contract event 0xb9a5bd9c6fc68ddbb2525a58437661292930f5a4adc18d5fa410ac494a15d0c3.
//
// Solidity: event CollectFees(address indexed maker, uint256 uFees0, uint256 uFees1, uint256 lFees0, uint256 lFees1)
func (_Api *ApiFilterer) WatchCollectFees(opts *bind.WatchOpts, sink chan<- *ApiCollectFees, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "CollectFees", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiCollectFees)
				if err := _Api.contract.UnpackLog(event, "CollectFees", log); err != nil {
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

// ParseCollectFees is a log parse operation binding the contract event 0xb9a5bd9c6fc68ddbb2525a58437661292930f5a4adc18d5fa410ac494a15d0c3.
//
// Solidity: event CollectFees(address indexed maker, uint256 uFees0, uint256 uFees1, uint256 lFees0, uint256 lFees1)
func (_Api *ApiFilterer) ParseCollectFees(log types.Log) (*ApiCollectFees, error) {
	event := new(ApiCollectFees)
	if err := _Api.contract.UnpackLog(event, "CollectFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Api contract.
type ApiDepositIterator struct {
	Event *ApiDeposit // Event containing the contract specifics and raw log

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
func (it *ApiDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiDeposit)
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
		it.Event = new(ApiDeposit)
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
func (it *ApiDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiDeposit represents a Deposit event raised by the Api contract.
type ApiDeposit struct {
	Sender  common.Address
	To      common.Address
	Shares  *big.Int
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*ApiDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ApiDepositIterator{contract: _Api.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ApiDeposit, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiDeposit)
				if err := _Api.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiFilterer) ParseDeposit(log types.Log) (*ApiDeposit, error) {
	event := new(ApiDeposit)
	if err := _Api.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiErrorLoggingIterator is returned from FilterErrorLogging and is used to iterate over the raw logs and unpacked data for ErrorLogging events raised by the Api contract.
type ApiErrorLoggingIterator struct {
	Event *ApiErrorLogging // Event containing the contract specifics and raw log

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
func (it *ApiErrorLoggingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiErrorLogging)
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
		it.Event = new(ApiErrorLogging)
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
func (it *ApiErrorLoggingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiErrorLoggingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiErrorLogging represents a ErrorLogging event raised by the Api contract.
type ApiErrorLogging struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterErrorLogging is a free log retrieval operation binding the contract event 0xdcbe72357183be3854887085e70fd914c4c1733a94106878f48e8ba069d8fabc.
//
// Solidity: event ErrorLogging(string reason)
func (_Api *ApiFilterer) FilterErrorLogging(opts *bind.FilterOpts) (*ApiErrorLoggingIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "ErrorLogging")
	if err != nil {
		return nil, err
	}
	return &ApiErrorLoggingIterator{contract: _Api.contract, event: "ErrorLogging", logs: logs, sub: sub}, nil
}

// WatchErrorLogging is a free log subscription operation binding the contract event 0xdcbe72357183be3854887085e70fd914c4c1733a94106878f48e8ba069d8fabc.
//
// Solidity: event ErrorLogging(string reason)
func (_Api *ApiFilterer) WatchErrorLogging(opts *bind.WatchOpts, sink chan<- *ApiErrorLogging) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "ErrorLogging")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiErrorLogging)
				if err := _Api.contract.UnpackLog(event, "ErrorLogging", log); err != nil {
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

// ParseErrorLogging is a log parse operation binding the contract event 0xdcbe72357183be3854887085e70fd914c4c1733a94106878f48e8ba069d8fabc.
//
// Solidity: event ErrorLogging(string reason)
func (_Api *ApiFilterer) ParseErrorLogging(log types.Log) (*ApiErrorLogging, error) {
	event := new(ApiErrorLogging)
	if err := _Api.contract.UnpackLog(event, "ErrorLogging", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiGeneralAIterator is returned from FilterGeneralA and is used to iterate over the raw logs and unpacked data for GeneralA events raised by the Api contract.
type ApiGeneralAIterator struct {
	Event *ApiGeneralA // Event containing the contract specifics and raw log

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
func (it *ApiGeneralAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiGeneralA)
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
		it.Event = new(ApiGeneralA)
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
func (it *ApiGeneralAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiGeneralAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiGeneralA represents a GeneralA event raised by the Api contract.
type ApiGeneralA struct {
	Subject string
	Value   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGeneralA is a free log retrieval operation binding the contract event 0x3ab002a7eebd5665644b0637a05a350d75a661095d978c8094688b8913cb9c39.
//
// Solidity: event GeneralA(string subject, address value)
func (_Api *ApiFilterer) FilterGeneralA(opts *bind.FilterOpts) (*ApiGeneralAIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "GeneralA")
	if err != nil {
		return nil, err
	}
	return &ApiGeneralAIterator{contract: _Api.contract, event: "GeneralA", logs: logs, sub: sub}, nil
}

// WatchGeneralA is a free log subscription operation binding the contract event 0x3ab002a7eebd5665644b0637a05a350d75a661095d978c8094688b8913cb9c39.
//
// Solidity: event GeneralA(string subject, address value)
func (_Api *ApiFilterer) WatchGeneralA(opts *bind.WatchOpts, sink chan<- *ApiGeneralA) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "GeneralA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiGeneralA)
				if err := _Api.contract.UnpackLog(event, "GeneralA", log); err != nil {
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

// ParseGeneralA is a log parse operation binding the contract event 0x3ab002a7eebd5665644b0637a05a350d75a661095d978c8094688b8913cb9c39.
//
// Solidity: event GeneralA(string subject, address value)
func (_Api *ApiFilterer) ParseGeneralA(log types.Log) (*ApiGeneralA, error) {
	event := new(ApiGeneralA)
	if err := _Api.contract.UnpackLog(event, "GeneralA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiGeneralBIterator is returned from FilterGeneralB and is used to iterate over the raw logs and unpacked data for GeneralB events raised by the Api contract.
type ApiGeneralBIterator struct {
	Event *ApiGeneralB // Event containing the contract specifics and raw log

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
func (it *ApiGeneralBIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiGeneralB)
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
		it.Event = new(ApiGeneralB)
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
func (it *ApiGeneralBIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiGeneralBIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiGeneralB represents a GeneralB event raised by the Api contract.
type ApiGeneralB struct {
	Subject string
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGeneralB is a free log retrieval operation binding the contract event 0xb6bc63617433f8cf440bcc140feaaee6ef7d5796bef15eddbb581271d94754b2.
//
// Solidity: event GeneralB(string subject, uint256 value)
func (_Api *ApiFilterer) FilterGeneralB(opts *bind.FilterOpts) (*ApiGeneralBIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "GeneralB")
	if err != nil {
		return nil, err
	}
	return &ApiGeneralBIterator{contract: _Api.contract, event: "GeneralB", logs: logs, sub: sub}, nil
}

// WatchGeneralB is a free log subscription operation binding the contract event 0xb6bc63617433f8cf440bcc140feaaee6ef7d5796bef15eddbb581271d94754b2.
//
// Solidity: event GeneralB(string subject, uint256 value)
func (_Api *ApiFilterer) WatchGeneralB(opts *bind.WatchOpts, sink chan<- *ApiGeneralB) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "GeneralB")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiGeneralB)
				if err := _Api.contract.UnpackLog(event, "GeneralB", log); err != nil {
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

// ParseGeneralB is a log parse operation binding the contract event 0xb6bc63617433f8cf440bcc140feaaee6ef7d5796bef15eddbb581271d94754b2.
//
// Solidity: event GeneralB(string subject, uint256 value)
func (_Api *ApiFilterer) ParseGeneralB(log types.Log) (*ApiGeneralB, error) {
	event := new(ApiGeneralB)
	if err := _Api.contract.UnpackLog(event, "GeneralB", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiGeneralSIterator is returned from FilterGeneralS and is used to iterate over the raw logs and unpacked data for GeneralS events raised by the Api contract.
type ApiGeneralSIterator struct {
	Event *ApiGeneralS // Event containing the contract specifics and raw log

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
func (it *ApiGeneralSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiGeneralS)
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
		it.Event = new(ApiGeneralS)
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
func (it *ApiGeneralSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiGeneralSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiGeneralS represents a GeneralS event raised by the Api contract.
type ApiGeneralS struct {
	Subject string
	Value   string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGeneralS is a free log retrieval operation binding the contract event 0x461a711dabbb89f674949025eaa8bc10b8f56e252fe8cdbd4d8eeb3069f018b8.
//
// Solidity: event GeneralS(string subject, string value)
func (_Api *ApiFilterer) FilterGeneralS(opts *bind.FilterOpts) (*ApiGeneralSIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "GeneralS")
	if err != nil {
		return nil, err
	}
	return &ApiGeneralSIterator{contract: _Api.contract, event: "GeneralS", logs: logs, sub: sub}, nil
}

// WatchGeneralS is a free log subscription operation binding the contract event 0x461a711dabbb89f674949025eaa8bc10b8f56e252fe8cdbd4d8eeb3069f018b8.
//
// Solidity: event GeneralS(string subject, string value)
func (_Api *ApiFilterer) WatchGeneralS(opts *bind.WatchOpts, sink chan<- *ApiGeneralS) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "GeneralS")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiGeneralS)
				if err := _Api.contract.UnpackLog(event, "GeneralS", log); err != nil {
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

// ParseGeneralS is a log parse operation binding the contract event 0x461a711dabbb89f674949025eaa8bc10b8f56e252fe8cdbd4d8eeb3069f018b8.
//
// Solidity: event GeneralS(string subject, string value)
func (_Api *ApiFilterer) ParseGeneralS(log types.Log) (*ApiGeneralS, error) {
	event := new(ApiGeneralS)
	if err := _Api.contract.UnpackLog(event, "GeneralS", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiMyLogIterator is returned from FilterMyLog and is used to iterate over the raw logs and unpacked data for MyLog events raised by the Api contract.
type ApiMyLogIterator struct {
	Event *ApiMyLog // Event containing the contract specifics and raw log

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
func (it *ApiMyLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiMyLog)
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
		it.Event = new(ApiMyLog)
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
func (it *ApiMyLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiMyLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiMyLog represents a MyLog event raised by the Api contract.
type ApiMyLog struct {
	Arg0 string
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMyLog is a free log retrieval operation binding the contract event 0x8d1cced004452bd270777a8c670f9f7e7c4fdde56f2db331fe289d39dc2624ad.
//
// Solidity: event MyLog(string arg0, uint256 arg1)
func (_Api *ApiFilterer) FilterMyLog(opts *bind.FilterOpts) (*ApiMyLogIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "MyLog")
	if err != nil {
		return nil, err
	}
	return &ApiMyLogIterator{contract: _Api.contract, event: "MyLog", logs: logs, sub: sub}, nil
}

// WatchMyLog is a free log subscription operation binding the contract event 0x8d1cced004452bd270777a8c670f9f7e7c4fdde56f2db331fe289d39dc2624ad.
//
// Solidity: event MyLog(string arg0, uint256 arg1)
func (_Api *ApiFilterer) WatchMyLog(opts *bind.WatchOpts, sink chan<- *ApiMyLog) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "MyLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiMyLog)
				if err := _Api.contract.UnpackLog(event, "MyLog", log); err != nil {
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

// ParseMyLog is a log parse operation binding the contract event 0x8d1cced004452bd270777a8c670f9f7e7c4fdde56f2db331fe289d39dc2624ad.
//
// Solidity: event MyLog(string arg0, uint256 arg1)
func (_Api *ApiFilterer) ParseMyLog(log types.Log) (*ApiMyLog, error) {
	event := new(ApiMyLog)
	if err := _Api.contract.UnpackLog(event, "MyLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiMyLog2Iterator is returned from FilterMyLog2 and is used to iterate over the raw logs and unpacked data for MyLog2 events raised by the Api contract.
type ApiMyLog2Iterator struct {
	Event *ApiMyLog2 // Event containing the contract specifics and raw log

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
func (it *ApiMyLog2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiMyLog2)
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
		it.Event = new(ApiMyLog2)
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
func (it *ApiMyLog2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiMyLog2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiMyLog2 represents a MyLog2 event raised by the Api contract.
type ApiMyLog2 struct {
	Arg0 string
	Arg1 *big.Int
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMyLog2 is a free log retrieval operation binding the contract event 0xb626a5b8d37253edd986272eff7a2541a1538cad44624e413d51713f0b2a414d.
//
// Solidity: event MyLog2(string arg0, uint256 arg1, uint256 arg2)
func (_Api *ApiFilterer) FilterMyLog2(opts *bind.FilterOpts) (*ApiMyLog2Iterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "MyLog2")
	if err != nil {
		return nil, err
	}
	return &ApiMyLog2Iterator{contract: _Api.contract, event: "MyLog2", logs: logs, sub: sub}, nil
}

// WatchMyLog2 is a free log subscription operation binding the contract event 0xb626a5b8d37253edd986272eff7a2541a1538cad44624e413d51713f0b2a414d.
//
// Solidity: event MyLog2(string arg0, uint256 arg1, uint256 arg2)
func (_Api *ApiFilterer) WatchMyLog2(opts *bind.WatchOpts, sink chan<- *ApiMyLog2) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "MyLog2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiMyLog2)
				if err := _Api.contract.UnpackLog(event, "MyLog2", log); err != nil {
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

// ParseMyLog2 is a log parse operation binding the contract event 0xb626a5b8d37253edd986272eff7a2541a1538cad44624e413d51713f0b2a414d.
//
// Solidity: event MyLog2(string arg0, uint256 arg1, uint256 arg2)
func (_Api *ApiFilterer) ParseMyLog2(log types.Log) (*ApiMyLog2, error) {
	event := new(ApiMyLog2)
	if err := _Api.contract.UnpackLog(event, "MyLog2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiMyLog4Iterator is returned from FilterMyLog4 and is used to iterate over the raw logs and unpacked data for MyLog4 events raised by the Api contract.
type ApiMyLog4Iterator struct {
	Event *ApiMyLog4 // Event containing the contract specifics and raw log

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
func (it *ApiMyLog4Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiMyLog4)
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
		it.Event = new(ApiMyLog4)
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
func (it *ApiMyLog4Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiMyLog4Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiMyLog4 represents a MyLog4 event raised by the Api contract.
type ApiMyLog4 struct {
	Arg0 string
	Arg1 *big.Int
	Arg2 *big.Int
	Arg3 *big.Int
	Arg4 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMyLog4 is a free log retrieval operation binding the contract event 0xa20514a72e3ef3eedfb172365769e68cd0eb83df0ef135feb5394a31045f7528.
//
// Solidity: event MyLog4(string arg0, uint256 arg1, uint256 arg2, uint256 arg3, uint256 arg4)
func (_Api *ApiFilterer) FilterMyLog4(opts *bind.FilterOpts) (*ApiMyLog4Iterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "MyLog4")
	if err != nil {
		return nil, err
	}
	return &ApiMyLog4Iterator{contract: _Api.contract, event: "MyLog4", logs: logs, sub: sub}, nil
}

// WatchMyLog4 is a free log subscription operation binding the contract event 0xa20514a72e3ef3eedfb172365769e68cd0eb83df0ef135feb5394a31045f7528.
//
// Solidity: event MyLog4(string arg0, uint256 arg1, uint256 arg2, uint256 arg3, uint256 arg4)
func (_Api *ApiFilterer) WatchMyLog4(opts *bind.WatchOpts, sink chan<- *ApiMyLog4) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "MyLog4")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiMyLog4)
				if err := _Api.contract.UnpackLog(event, "MyLog4", log); err != nil {
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

// ParseMyLog4 is a log parse operation binding the contract event 0xa20514a72e3ef3eedfb172365769e68cd0eb83df0ef135feb5394a31045f7528.
//
// Solidity: event MyLog4(string arg0, uint256 arg1, uint256 arg2, uint256 arg3, uint256 arg4)
func (_Api *ApiFilterer) ParseMyLog4(log types.Log) (*ApiMyLog4, error) {
	event := new(ApiMyLog4)
	if err := _Api.contract.UnpackLog(event, "MyLog4", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiRebalanceLogIterator is returned from FilterRebalanceLog and is used to iterate over the raw logs and unpacked data for RebalanceLog events raised by the Api contract.
type ApiRebalanceLogIterator struct {
	Event *ApiRebalanceLog // Event containing the contract specifics and raw log

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
func (it *ApiRebalanceLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiRebalanceLog)
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
		it.Event = new(ApiRebalanceLog)
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
func (it *ApiRebalanceLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiRebalanceLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiRebalanceLog represents a RebalanceLog event raised by the Api contract.
type ApiRebalanceLog struct {
	Liquidity   *big.Int
	NewBalance0 *big.Int
	NewBalance1 *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRebalanceLog is a free log retrieval operation binding the contract event 0xf290cad35129b75da3cfb1a6dfcb58954eabf3dd09731a5ea198c49a4ab2c032.
//
// Solidity: event RebalanceLog(uint128 liquidity, uint256 newBalance0, uint256 newBalance1)
func (_Api *ApiFilterer) FilterRebalanceLog(opts *bind.FilterOpts) (*ApiRebalanceLogIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "RebalanceLog")
	if err != nil {
		return nil, err
	}
	return &ApiRebalanceLogIterator{contract: _Api.contract, event: "RebalanceLog", logs: logs, sub: sub}, nil
}

// WatchRebalanceLog is a free log subscription operation binding the contract event 0xf290cad35129b75da3cfb1a6dfcb58954eabf3dd09731a5ea198c49a4ab2c032.
//
// Solidity: event RebalanceLog(uint128 liquidity, uint256 newBalance0, uint256 newBalance1)
func (_Api *ApiFilterer) WatchRebalanceLog(opts *bind.WatchOpts, sink chan<- *ApiRebalanceLog) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "RebalanceLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiRebalanceLog)
				if err := _Api.contract.UnpackLog(event, "RebalanceLog", log); err != nil {
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

// ParseRebalanceLog is a log parse operation binding the contract event 0xf290cad35129b75da3cfb1a6dfcb58954eabf3dd09731a5ea198c49a4ab2c032.
//
// Solidity: event RebalanceLog(uint128 liquidity, uint256 newBalance0, uint256 newBalance1)
func (_Api *ApiFilterer) ParseRebalanceLog(log types.Log) (*ApiRebalanceLog, error) {
	event := new(ApiRebalanceLog)
	if err := _Api.contract.UnpackLog(event, "RebalanceLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiSnapshotIterator is returned from FilterSnapshot and is used to iterate over the raw logs and unpacked data for Snapshot events raised by the Api contract.
type ApiSnapshotIterator struct {
	Event *ApiSnapshot // Event containing the contract specifics and raw log

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
func (it *ApiSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiSnapshot)
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
		it.Event = new(ApiSnapshot)
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
func (it *ApiSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiSnapshot represents a Snapshot event raised by the Api contract.
type ApiSnapshot struct {
	Tick         *big.Int
	TotalAmount0 *big.Int
	TotalAmount1 *big.Int
	TotalSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSnapshot is a free log retrieval operation binding the contract event 0x210f60adf1db7a02e9db9a49ec7c2eb2060c516cbcfd01a0c05288144738ee5d.
//
// Solidity: event Snapshot(int24 tick, uint256 totalAmount0, uint256 totalAmount1, uint256 totalSupply)
func (_Api *ApiFilterer) FilterSnapshot(opts *bind.FilterOpts) (*ApiSnapshotIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return &ApiSnapshotIterator{contract: _Api.contract, event: "Snapshot", logs: logs, sub: sub}, nil
}

// WatchSnapshot is a free log subscription operation binding the contract event 0x210f60adf1db7a02e9db9a49ec7c2eb2060c516cbcfd01a0c05288144738ee5d.
//
// Solidity: event Snapshot(int24 tick, uint256 totalAmount0, uint256 totalAmount1, uint256 totalSupply)
func (_Api *ApiFilterer) WatchSnapshot(opts *bind.WatchOpts, sink chan<- *ApiSnapshot) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiSnapshot)
				if err := _Api.contract.UnpackLog(event, "Snapshot", log); err != nil {
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

// ParseSnapshot is a log parse operation binding the contract event 0x210f60adf1db7a02e9db9a49ec7c2eb2060c516cbcfd01a0c05288144738ee5d.
//
// Solidity: event Snapshot(int24 tick, uint256 totalAmount0, uint256 totalAmount1, uint256 totalSupply)
func (_Api *ApiFilterer) ParseSnapshot(log types.Log) (*ApiSnapshot, error) {
	event := new(ApiSnapshot)
	if err := _Api.contract.UnpackLog(event, "Snapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Api contract.
type ApiTransferIterator struct {
	Event *ApiTransfer // Event containing the contract specifics and raw log

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
func (it *ApiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiTransfer)
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
		it.Event = new(ApiTransfer)
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
func (it *ApiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiTransfer represents a Transfer event raised by the Api contract.
type ApiTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Api *ApiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ApiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ApiTransferIterator{contract: _Api.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Api *ApiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ApiTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiTransfer)
				if err := _Api.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Api *ApiFilterer) ParseTransfer(log types.Log) (*ApiTransfer, error) {
	event := new(ApiTransfer)
	if err := _Api.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Api contract.
type ApiWithdrawIterator struct {
	Event *ApiWithdraw // Event containing the contract specifics and raw log

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
func (it *ApiWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiWithdraw)
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
		it.Event = new(ApiWithdraw)
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
func (it *ApiWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiWithdraw represents a Withdraw event raised by the Api contract.
type ApiWithdraw struct {
	Sender  common.Address
	Shares  *big.Int
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94.
//
// Solidity: event Withdraw(address indexed sender, uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address) (*ApiWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Withdraw", senderRule)
	if err != nil {
		return nil, err
	}
	return &ApiWithdrawIterator{contract: _Api.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94.
//
// Solidity: event Withdraw(address indexed sender, uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ApiWithdraw, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Withdraw", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiWithdraw)
				if err := _Api.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94.
//
// Solidity: event Withdraw(address indexed sender, uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiFilterer) ParseWithdraw(log types.Log) (*ApiWithdraw, error) {
	event := new(ApiWithdraw)
	if err := _Api.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

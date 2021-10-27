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
const ApiABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cToken0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cToken1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cEth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_protocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxTotalSupply\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"_maxTwapDeviation\",\"type\":\"int24\"},{\"internalType\":\"uint32\",\"name\":\"_twapDuration\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_uniPortion\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_team\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burned0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burned1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collect0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collect1\",\"type\":\"uint256\"}],\"name\":\"CollectFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"ErrorLogging\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"subject\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"GeneralA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"subject\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"GeneralB\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"subject\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"GeneralS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MyLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MyLog2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MyLog4\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance1\",\"type\":\"uint256\"}],\"name\":\"RebalanceLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"Snapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"AccumulateUniswapFees0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AccumulateUniswapFees1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Assetholder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"deposit0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fees0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fees1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedProtocolFees0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedProtocolFees1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alloc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cHigh\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cLow\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"collectProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curComp0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curComp1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curUni0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curUni1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToken1\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"doRebalence\",\"type\":\"bool\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getETHBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"getPositionAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"baseAmount\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"}],\"name\":\"getQuoteAtTick\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"getSSLiquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"getStoredAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"name\":\"getTwap\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lFees0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lFees1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRebalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTick\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTwapDeviation\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeCTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"setGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMax\",\"type\":\"uint256\"}],\"name\":\"setMaxTotalSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_maxTwapDeviation\",\"type\":\"int24\"}],\"name\":\"setMaxTwapDeviation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_protocolFee\",\"type\":\"uint256\"}],\"name\":\"setProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_team\",\"type\":\"address\"}],\"name\":\"setTeam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_twapDuration\",\"type\":\"uint32\"}],\"name\":\"setTwapDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"ratio\",\"type\":\"uint8\"}],\"name\":\"setUniPortionRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"newLow\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"newHigh\",\"type\":\"int24\"}],\"name\":\"strategy1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20Contract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cErc20Contract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_numTokensToSupply\",\"type\":\"uint256\"}],\"name\":\"supplyErc20ToCompound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_cEtherContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ctoken\",\"type\":\"address\"}],\"name\":\"supplyEthToCompound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"swapAmount\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"team\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"twapDuration\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uFees0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uFees1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniPortion\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ApiBin is the compiled bytecode used for deploying new contracts.
var ApiBin = "0x60c06040523480156200001157600080fd5b50604051620065523803806200655283398181016040526101608110156200003857600080fd5b5080516020808301516040808501516060860151608087015160a088015160c089015160e08a01516101008b01516101208c0151610140909c01518851808a018a52600e81526d05669614c656e6420546f6b656e360941b818d019081528a51808c01909b5260048b52630564c54360e41b9c8b019c909c52600160005580519c9d9a9c989b979a9699959894979396929592949193909291620000df91601e91620004e9565b508051620000f590601f906020840190620004e9565b506012602060006101000a81548160ff021916908360ff160217905550505033600160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600260006101000a8154816001600160a01b0302191690836001600160a01b031602179055508a600560006101000a8154816001600160a01b0302191690836001600160a01b031602179055508a6001600160a01b0316630dfe16816040518163ffffffff1660e01b815260040160206040518083038186803b158015620001c357600080fd5b505afa158015620001d8573d6000803e3d6000fd5b505050506040513d6020811015620001ef57600080fd5b505160601b6001600160601b0319166080526040805163d21220a760e01b815290516001600160a01b038d169163d21220a7916004808301926020929190829003018186803b1580156200024257600080fd5b505afa15801562000257573d6000803e3d6000fd5b505050506040513d60208110156200026e57600080fd5b505160601b6001600160601b03191660a05260138054600160401b600160e01b031916680100000000000000006001600160a01b038c81169190910291909117909155601480546001600160a01b03199081168b841617909155601580549091168983161790558a1662000312576040805162461bcd60e51b815260206004808301919091526024820152630ae8aa8960e31b604482015290519081900360640190fd5b89601660006101000a8154816001600160a01b0302191690836001600160a01b03160217905550856004819055508a6001600160a01b031663d0c93a7c6040518163ffffffff1660e01b815260040160206040518083038186803b1580156200037a57600080fd5b505afa1580156200038f573d6000803e3d6000fd5b505050506040513d6020811015620003a657600080fd5b50516005805460029290920b62ffffff16600160a01b0262ffffff60a01b19909216919091179055606486106200040a576040805162461bcd60e51b815260206004820152600360248201526228232960e91b604482015290519081900360640190fd5b60128590556013805460ff84166701000000000000000260ff60381b1963ffffffff871663ffffffff1960028a900b62ffffff81166401000000000262ffffff60201b19909616959095171617161790915560001262000497576040805162461bcd60e51b815260206004820152600360248201526213551160ea1b604482015290519081900360640190fd5b60008363ffffffff1611620004d8576040805162461bcd60e51b8152602060048201526002602482015261151160f21b604482015290519081900360640190fd5b505050505050505050505062000595565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826200052157600085556200056c565b82601f106200053c57805160ff19168380011785556200056c565b828001600101855582156200056c579182015b828111156200056c5782518255916020019190600101906200054f565b506200057a9291506200057e565b5090565b5b808211156200057a57600081556001016200057f565b60805160601c60a05160601c615efd620006556000398061108152806116925280611df15280611e8b5280612d2c5280612ddf5280612f3d528061313d52806132e852806133b0528061342f528061356e5280613bd55280613f5052806144345280614ae852508061103752806112e352806115f85280611dd05280611e505280612da55280612ea452806130e85280613176528061323e52806132bd52806135315280613b355280613eb752806144135280614ac75250615efd6000f3fe6080604052600436106103e85760003560e01c806375bd711111610208578063c433c80a11610118578063d595a12b116100ab578063eae989a21161007a578063eae989a214610e8e578063f10d8fc814610ea3578063f2a40db814610eb8578063f39c38a014610ee2578063fa461e3314610ef7576103ef565b8063d595a12b14610e14578063d8d7f96f14610e29578063dd62ed3e14610e3e578063e7c7cb9114610e79576103ef565b8063d17cfebd116100e7578063d17cfebd14610d4e578063d21220a714610d63578063d348799714610d78578063d368867f14610dff576103ef565b8063c433c80a14610ca1578063c8a746a614610cd1578063ccac7d8b14610ce6578063cda206b014610cfb576103ef565b8063a4ee141b1161019b578063aa489e061161016a578063aa489e0614610c1a578063ab033ea914610c2f578063b0e21e8a14610c62578063bbc001c314610c77578063c29188bc14610c8c576103ef565b8063a4ee141b14610b3c578063a833d7ca14610b95578063a9059cbb14610baa578063a91ef6eb14610be3576103ef565b80638e843c6c116101d75780638e843c6c14610ac457806395d89b4114610ad9578063a00fa77f14610aee578063a457c2d714610b03576103ef565b806375bd711114610a5b578063787dce3d14610a7057806382fc254314610a9a57806385f2aef214610aaf576103ef565b806339509351116103035780635aa6e6751161029657806366d7505b1161026557806366d7505b1461096557806368a5787a146109a45780636e947298146109e557806370a08231146109fa57806373232c6014610a2d576103ef565b80635aa6e6751461088e5780635bb6aa85146108a3578063624a177a146109015780636253c7181461092e576103ef565b80633f3e4c11116102d25780633f3e4c11146107c457806343a0d066146107ee57806343c57a2714610826578063580b82e314610879576103ef565b806339509351146106ef5780633b7ba25f146107285780633cbff3fe1461076b5780633dfa5d8714610798576103ef565b806318160ddd1161037b5780632ab4d0521161034a5780632ab4d052146106705780632e1a7d4d14610685578063313ce567146106af57806331dc5b95146106da576103ef565b806318160ddd146105d5578063238efcbc146105ea57806323b872dd146105ff57806326d8954514610642576103ef565b8063095ea7b3116103b7578063095ea7b31461051b5780630dfe168114610568578063106b9ca1146105995780631755ff21146105c0576103ef565b80630430c130146103f157806306b0b6561461043057806306fdde031461045e578063095cf5c6146104e8576103ef565b366103ef57005b005b3480156103fd57600080fd5b506103ef6004803603606081101561041457600080fd5b50803590602081013590604001356001600160a01b0316610f7e565b34801561043c57600080fd5b506104456110ad565b6040805192835260208301919091528051918290030190f35b34801561046a57600080fd5b506104736111be565b6040805160208082528351818301528351919283929083019185019080838360005b838110156104ad578181015183820152602001610495565b50505050905090810190601f1680156104da5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156104f457600080fd5b506103ef6004803603602081101561050b57600080fd5b50356001600160a01b0316611255565b34801561052757600080fd5b506105546004803603604081101561053e57600080fd5b506001600160a01b0381351690602001356112c3565b604080519115158252519081900360200190f35b34801561057457600080fd5b5061057d6112e1565b604080516001600160a01b039092168252519081900360200190f35b3480156105a557600080fd5b506105ae611305565b60408051918252519081900360200190f35b3480156105cc57600080fd5b5061057d61130b565b3480156105e157600080fd5b506105ae61131a565b3480156105f657600080fd5b506103ef611320565b34801561060b57600080fd5b506105546004803603606081101561062257600080fd5b506001600160a01b03813581169160208101359091169060400135611387565b34801561064e57600080fd5b5061065761140f565b6040805163ffffffff9092168252519081900360200190f35b34801561067c57600080fd5b506105ae61141b565b34801561069157600080fd5b50610445600480360360208110156106a857600080fd5b5035611421565b3480156106bb57600080fd5b506106c4611774565b6040805160ff9092168252519081900360200190f35b3480156106e657600080fd5b506106c461177d565b3480156106fb57600080fd5b506105546004803603604081101561071257600080fd5b506001600160a01b03813516906020013561178d565b34801561073457600080fd5b506105ae6004803603606081101561074b57600080fd5b506001600160a01b038135811691602081013590911690604001356117db565b34801561077757600080fd5b506103ef6004803603602081101561078e57600080fd5b503560020b611b28565b3480156107a457600080fd5b506107ad611be8565b6040805160029290920b8252519081900360200190f35b3480156107d057600080fd5b506103ef600480360360208110156107e757600080fd5b5035611bf1565b3480156107fa57600080fd5b506103ef6004803603606081101561081157600080fd5b50803590602081013590604001351515611c92565b34801561083257600080fd5b506105ae6004803603608081101561084957600080fd5b50803560020b906001600160801b03602082013516906001600160a01b0360408201358116916060013516611f94565b34801561088557600080fd5b506105ae612086565b34801561089a57600080fd5b5061057d61208c565b3480156108af57600080fd5b506108d6600480360360208110156108c657600080fd5b50356001600160a01b031661209b565b6040805195865260208601949094528484019290925260608401526080830152519081900360a00190f35b34801561090d57600080fd5b506103ef6004803603602081101561092457600080fd5b503560ff166120ca565b34801561093a57600080fd5b506103ef6004803603604081101561095157600080fd5b508035600290810b9160200135900b61217b565b34801561097157600080fd5b506107ad6004803603604081101561098857600080fd5b5080356001600160a01b0316906020013563ffffffff1661226f565b3480156109b057600080fd5b50610445600480360360608110156109c757600080fd5b508035906020810135151590604001356001600160a01b0316612533565b3480156109f157600080fd5b506105ae612668565b348015610a0657600080fd5b506105ae60048036036020811015610a1d57600080fd5b50356001600160a01b031661266c565b61055460048036036040811015610a4357600080fd5b506001600160a01b038135811691602001351661268b565b348015610a6757600080fd5b506105ae61286b565b348015610a7c57600080fd5b506103ef60048036036020811015610a9357600080fd5b5035612871565b348015610aa657600080fd5b506105ae612907565b348015610abb57600080fd5b5061057d61290d565b348015610ad057600080fd5b506107ad61291c565b348015610ae557600080fd5b5061047361292c565b348015610afa57600080fd5b506105ae61298c565b348015610b0f57600080fd5b5061055460048036036040811015610b2657600080fd5b506001600160a01b038135169060200135612992565b348015610b4857600080fd5b50610b6f60048036036020811015610b5f57600080fd5b50356001600160a01b03166129fa565b604080519485526020850193909352838301919091526060830152519081900360800190f35b348015610ba157600080fd5b506105ae612a2b565b348015610bb657600080fd5b5061055460048036036040811015610bcd57600080fd5b506001600160a01b038135169060200135612a31565b348015610bef57600080fd5b5061044560048036036040811015610c0657600080fd5b508035600290810b9160200135900b612a45565b348015610c2657600080fd5b506105ae612adc565b348015610c3b57600080fd5b506103ef60048036036020811015610c5257600080fd5b50356001600160a01b0316612ae2565b348015610c6e57600080fd5b506105ae612b50565b348015610c8357600080fd5b506103ef612b56565b348015610c9857600080fd5b506105ae612bad565b348015610cad57600080fd5b506103ef60048036036020811015610cc457600080fd5b503563ffffffff16612bb3565b348015610cdd57600080fd5b506105ae612c65565b348015610cf257600080fd5b506105ae612c6b565b348015610d0757600080fd5b50610d3260048036036040811015610d1e57600080fd5b508035600290810b9160200135900b612c71565b604080516001600160801b039092168252519081900360200190f35b348015610d5a57600080fd5b506103ef612c89565b348015610d6f57600080fd5b5061057d612d2a565b348015610d8457600080fd5b506103ef60048036036060811015610d9b57600080fd5b813591602081013591810190606081016040820135600160201b811115610dc157600080fd5b820183602082011115610dd357600080fd5b803590602001918460018302840111600160201b83111715610df457600080fd5b509092509050612d4e565b348015610e0b57600080fd5b506107ad612e0c565b348015610e2057600080fd5b506105ae612e1c565b348015610e3557600080fd5b506103ef612e22565b348015610e4a57600080fd5b506105ae60048036036040811015610e6157600080fd5b506001600160a01b0381358116916020013516613456565b348015610e8557600080fd5b506107ad613481565b348015610e9a57600080fd5b506105ae613491565b348015610eaf57600080fd5b506105ae613497565b348015610ec457600080fd5b5061057d60048036036020811015610edb57600080fd5b503561349d565b348015610eee57600080fd5b5061057d6134c7565b348015610f0357600080fd5b506103ef60048036036060811015610f1a57600080fd5b813591602081013591810190606081016040820135600160201b811115610f4057600080fd5b820183602082011115610f5257600080fd5b803590602001918460018302840111600160201b83111715610f7357600080fd5b5090925090506134d6565b6001546001600160a01b03163314610fca576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b82600a5410158015610fde575081600b5410155b611014576040805162461bcd60e51b8152602060048201526002602482015261043560f41b604482015290519081900360640190fd5b821561105e57600a546110279084613595565b600a5561105e6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001682856135f2565b81156110a857600b546110719083613595565b600b556110a86001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001682846135f2565b505050565b600080601360089054906101000a90046001600160a01b03166001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561111257600080fd5b505afa158015611126573d6000803e3d6000fd5b505050506040513d602081101561113c57600080fd5b5051601454604080516370a0823160e01b815230600482015290519294506001600160a01b03909116916370a0823191602480820192602092909190829003018186803b15801561118c57600080fd5b505afa1580156111a0573d6000803e3d6000fd5b505050506040513d60208110156111b657600080fd5b505191929050565b601e8054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561124a5780601f1061121f5761010080835404028352916020019161124a565b820191906000526020600020905b81548152906001019060200180831161122d57829003601f168201915b505050505090505b90565b6001546001600160a01b031633146112a1576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b600280546001600160a01b0319166001600160a01b0392909216919091179055565b60006112d76112d0613644565b8484613648565b5060015b92915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b600e5481565b6005546001600160a01b031690565b601d5490565b6003546001600160a01b03163314611373576040805162461bcd60e51b815260206004820152601160248201527070656e64696e67476f7665726e616e636560781b604482015290519081900360640190fd5b600180546001600160a01b03191633179055565b6000611394848484613734565b611404846113a0613644565b6113ff85604051806060016040528060288152602001615dc7602891396001600160a01b038a166000908152601c60205260408120906113de613644565b6001600160a01b031681526020810191909152604001600020549190613891565b613648565b5060015b9392505050565b60135463ffffffff1681565b60125481565b6000806002600054141561146a576040805162461bcd60e51b815260206004820152601f6024820152600080516020615caf833981519152604482015290519081900360640190fd5b600260005533801580159061148857506001600160a01b0381163014155b6114bf576040805162461bcd60e51b8152602060048201526003602482015262746f5760e81b604482015290519081900360640190fd5b6000841180156114d0575060648411155b61150a576040805162461bcd60e51b815260206004808301919091526024820152631418db9d60e21b604482015290519081900360640190fd5b611512613928565b506000611533606461152d876115278661266c565b90613a71565b90613aca565b9050600061153f61131a565b90506000811161157c576040805162461bcd60e51b815260206004820152600360248201526207473360ec1b604482015290519081900360640190fd5b808211156115b7576040805162461bcd60e51b815260206004820152600360248201526273683160e81b604482015290519081900360640190fd5b6115c78161152d84611527613b31565b94506115d98161152d84611527613bd1565b93506115e58383613c40565b841561167f5761161f6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001684876135f2565b6001600160a01b038316600090815260116020526040902054851061165c576001600160a01b03831660009081526011602052604090205461165e565b845b6001600160a01b038416600090815260116020526040902080549190910390555b8315611722576116b96001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001684866135f2565b6001600160a01b03831660009081526011602052604090206001015484106116fc576001600160a01b0383166000908152601160205260409020600101546116fe565b835b6001600160a01b038416600090815260116020526040902060010180549190910390555b6040805183815260208101879052808201869052905133917f02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94919081900360600190a250506001600055509092909150565b60205460ff1690565b601354600160381b900460ff1681565b60006112d761179a613644565b846113ff85601c60006117ab613644565b6001600160a01b03908116825260208083019390935260409182016000908120918c168152925290205490613d3c565b6000836001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561182a57600080fd5b505afa15801561183e573d6000803e3d6000fd5b505050506040513d602081101561185457600080fd5b5051821115611894576040805162461bcd60e51b815260206004820152600760248201526662616c616e636560c81b604482015290519081900360640190fd5b600084905060008490506000816001600160a01b031663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156118db57600080fd5b505af11580156118ef573d6000803e3d6000fd5b505050506040513d602081101561190557600080fd5b50516040805160208101839052818152601b818301527f45786368616e6765205261746520287363616c6564207570293a20000000000060608201529051919250600080516020615e59833981519152919081900360800190a16000826001600160a01b031663ae9d70b06040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561199c57600080fd5b505af11580156119b0573d6000803e3d6000fd5b505050506040513d60208110156119c657600080fd5b505160408051602081018390528181526018818301527f537570706c7920526174653a20287363616c656420757029000000000000000060608201529051919250600080516020615e59833981519152919081900360800190a1836001600160a01b031663095ea7b388886040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b158015611a7757600080fd5b505af1158015611a8b573d6000803e3d6000fd5b505050506040513d6020811015611aa157600080fd5b50506040805163140e25ad60e31b81526004810188905290516000916001600160a01b0386169163a0712d689160248082019260209290919082900301818787803b158015611aef57600080fd5b505af1158015611b03573d6000803e3d6000fd5b505050506040513d6020811015611b1957600080fd5b50519998505050505050505050565b6001546001600160a01b03163314611b74576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b60008160020b13611bbf576040805162461bcd60e51b815260206004820152601060248201526f36b0bc2a3bb0b82232bb34b0ba34b7b760811b604482015290519081900360640190fd5b6013805460029290920b62ffffff16600160201b0266ffffff0000000019909216919091179055565b600f5460020b81565b60026000541415611c37576040805162461bcd60e51b815260206004820152601f6024820152600080516020615caf833981519152604482015290519081900360640190fd5b60026000556001546001600160a01b03163314611c88576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b6012556001600055565b60026000541415611cd8576040805162461bcd60e51b815260206004820152601f6024820152600080516020615caf833981519152604482015290519081900360640190fd5b60026000553383151580611cec5750600083115b611d26576040805162461bcd60e51b815260206004808301919091526024820152630616d74360e41b604482015290519081900360640190fd5b6001600160a01b03811615801590611d4757506001600160a01b0381163014155b8015611d6157506002546001600160a01b03828116911614155b611d97576040805162461bcd60e51b8152602060048201526002602482015261746f60f01b604482015290519081900360640190fd5b611da081613d96565b600554601354611e1591611dc5916001600160a01b039091169063ffffffff1661226f565b670de0b6b3a76400007f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611f94565b60235560008080611e268787613e17565b925092509250611e34613928565b50611e3d613e77565b8115611e7857611e786001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016333085614107565b8015611eb357611eb36001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016333084614107565b611ebd8484614161565b601254611ec861131a565b1115611f01576040805162461bcd60e51b815260206004820152600360248201526204341560ec1b604482015290519081900360640190fd5b6001600160a01b03841660009081526011602052604090208054830181556001018054820190558415611f3957611f39600080614253565b604080518481526020810184905280820183905290516001600160a01b0386169133917f4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f69181900360600190a3505060016000555050505050565b600080611fa08661449e565b90506001600160801b036001600160a01b0382161161200f576001600160a01b0380821680029084811690861610611fef57611fea600160c01b876001600160801b0316836147cf565b612007565b61200781876001600160801b0316600160c01b6147cf565b92505061207d565b60006120296001600160a01b03831680600160401b6147cf565b9050836001600160a01b0316856001600160a01b0316106120615761205c600160801b876001600160801b0316836147cf565b612079565b61207981876001600160801b0316600160801b6147cf565b9250505b50949350505050565b60175481565b6001546001600160a01b031681565b601160205260009081526040902080546001820154600283015460038401546004909401549293919290919085565b6001546001600160a01b03163314612116576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b60648160ff161115612157576040805162461bcd60e51b8152602060048201526005602482015264726174696f60d81b604482015290519081900360640190fd5b6013805460ff909216600160381b0267ff0000000000000019909216919091179055565b600260005414156121c1576040805162461bcd60e51b815260206004820152601f6024820152600080516020615caf833981519152604482015290519081900360640190fd5b60026000556001546001600160a01b03163314612212576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b600061221c61131a565b11612253576040805162461bcd60e51b8152602060048201526002602482015261074360f41b604482015290519081900360640190fd5b61225b613928565b506122668282614253565b50506001600055565b60408051600280825260608201835260009284928492909160208301908036833701905050905081816000815181106122a457fe5b602002602001019063ffffffff16908163ffffffff16815250506000816001815181106122cd57fe5b63ffffffff90921660209283029190910182015260405163883bdbfd60e01b8152600481018281528351602483015283516000936001600160a01b038a169363883bdbfd938793909283926044019185820191028083838b5b8381101561233e578181015183820152602001612326565b505050509050019250505060006040518083038186803b15801561236157600080fd5b505afa158015612375573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604090815281101561239e57600080fd5b8101908080516040519392919084600160201b8211156123bd57600080fd5b9083019060208201858111156123d257600080fd5b82518660208202830111600160201b821117156123ee57600080fd5b82525081516020918201928201910280838360005b8381101561241b578181015183820152602001612403565b5050505090500160405260200180516040519392919084600160201b82111561244357600080fd5b90830190602082018581111561245857600080fd5b82518660208202830111600160201b8211171561247457600080fd5b82525081516020918201928201910280838360005b838110156124a1578181015183820152602001612489565b505050509050016040525050505090506000816000815181106124c057fe5b6020026020010151826001815181106124d557fe5b60200260200101510390508363ffffffff168160060b816124f257fe5b05945060008160060b12801561251c57508363ffffffff168160060b8161251557fe5b0760060b15155b1561252957600019909401935b5050505092915050565b600554604080513360208083019190915282518083038201815282840193849052630251596160e31b90935230604483018181528715156064850152608484018990526001600160a01b0387811660a486015260a060c48601908152865160e48701528651600098899893169663128acb08968c958e958d9590949193909261010490910191908501908083838f5b838110156125da5781810151838201526020016125c2565b50505050905090810190601f1680156126075780820380516001836020036101000a031916815260200191505b5096505050505050506040805180830381600087803b15801561262957600080fd5b505af115801561263d573d6000803e3d6000fd5b505050506040513d604081101561265357600080fd5b50805160209091015190969095509350505050565b4790565b6001600160a01b0381166000908152601b60205260409020545b919050565b60008083905060008390506000816001600160a01b031663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156126d357600080fd5b505af11580156126e7573d6000803e3d6000fd5b505050506040513d60208110156126fd57600080fd5b505160408051602081018390528181526023918101829052919250600080516020615e5983398151915291839181906060820190615d3782396040019250505060405180910390a16000826001600160a01b031663ae9d70b06040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561278257600080fd5b505af1158015612796573d6000803e3d6000fd5b505050506040513d60208110156127ac57600080fd5b5051604080516020818101849052828252818301527f537570706c7920526174653a20287363616c656420757020627920316531382960608201529051919250600080516020615e59833981519152919081900360800190a1836001600160a01b0316631249c58b6203d090476040518363ffffffff1660e01b81526004016000604051808303818589803b15801561284457600080fd5b5088f1158015612858573d6000803e3d6000fd5b5060019c9b505050505050505050505050565b60185481565b6001546001600160a01b031633146128bd576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b620f42408110612902576040805162461bcd60e51b815260206004820152600b60248201526a70726f746f636f6c46656560a81b604482015290519081900360640190fd5b600455565b60075481565b6002546001600160a01b031681565b600554600160d01b900460020b81565b601f805460408051602060026000196101006001871615020190941693909304808501849004840282018401909252818152606093909290919083018282801561124a5780601f1061121f5761010080835404028352916020019161124a565b600b5481565b60006112d761299f613644565b846113ff85604051806060016040528060258152602001615ea360259139601c60006129c9613644565b6001600160a01b03908116825260208083019390935260409182016000908120918d16815292529020549190613891565b6001600160a01b03166000908152601160205260409020805460018201546002830154600390930154919390929190565b60085481565b60006112d7612a3e613644565b8484613734565b6000806000806000612a57878761487e565b9450945050509250612a6a878785614936565b6004549196509450600090612a8390620f424090613595565b9050612aa9612aa2620f424061152d6001600160801b03871685613a71565b8790613d3c565b9550612acf612ac8620f424061152d6001600160801b03861685613a71565b8690613d3c565b9450505050509250929050565b600c5481565b6001546001600160a01b03163314612b2e576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b600380546001600160a01b0319166001600160a01b0392909216919091179055565b60045481565b6001546001600160a01b03163314612ba2576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b612baa613928565b50565b600d5481565b6001546001600160a01b03163314612bff576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b60008163ffffffff1611612c49576040805162461bcd60e51b815260206004820152600c60248201526b3a3bb0b8223ab930ba34b7b760a11b604482015290519081900360640190fd5b6013805463ffffffff191663ffffffff92909216919091179055565b60195481565b601a5481565b6000612c7d838361487e565b50929695505050505050565b6001546001600160a01b03163314612cd5576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b600080612ce06110ad565b90925090508115612d0957601354612d09908390600160401b90046001600160a01b03166149e0565b8015612d2657601454612d269082906001600160a01b03166149e0565b5050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6005546001600160a01b03163314612d92576040805162461bcd60e51b8152602060048201526002602482015261505360f01b604482015290519081900360640190fd5b8315612dcc57612dcc6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001633866135f2565b8215612e0657612e066001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001633856135f2565b50505050565b600554600160b81b900460020b81565b60065481565b6001546001600160a01b03163314612e6e576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b612e76614a9f565b612e7e612c89565b601054604080516370a0823160e01b815230600482015290516000916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016916370a0823191602480820192602092909190829003018186803b158015612eeb57600080fd5b505afa158015612eff573d6000803e3d6000fd5b505050506040513d6020811015612f1557600080fd5b5051604080516370a0823160e01b815230600482015290519192506000916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016916370a08231916024808301926020929190829003018186803b158015612f8357600080fd5b505afa158015612f97573d6000803e3d6000fd5b505050506040513d6020811015612fad57600080fd5b505190506000612fbb61131a565b905060005b84811015613171576000612ff460108381548110612fda57fe5b6000918252602090912001546001600160a01b031661266c565b90506130216010838154811061300657fe5b6000918252602090912001546001600160a01b031682613c40565b6000601160006010858154811061303457fe5b60009182526020808320909101546001600160a01b031683528201929092526040018120919091556010805460119183918690811061306f57fe5b60009182526020808320909101546001600160a01b031683528201929092526040018120600101919091556130a88461152d8885613a71565b905060006130ba8561152d8886613a71565b9050811561311157613111601085815481106130d257fe5b6000918252602090912001546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169116846135f2565b8015613166576131666010858154811061312757fe5b6000918252602090912001546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169116836135f2565b505050600101612fc0565b5060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156131e157600080fd5b505afa1580156131f5573d6000803e3d6000fd5b505050506040513d602081101561320b57600080fd5b505111156132e457600154604080516370a0823160e01b815230600482015290516132e4926001600160a01b03908116927f0000000000000000000000000000000000000000000000000000000000000000909116916370a0823191602480820192602092909190829003018186803b15801561328757600080fd5b505afa15801561329b573d6000803e3d6000fd5b505050506040513d60208110156132b157600080fd5b50516001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691906135f2565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561335357600080fd5b505afa158015613367573d6000803e3d6000fd5b505050506040513d602081101561337d57600080fd5b50511115612e0657600154604080516370a0823160e01b81523060048201529051612e06926001600160a01b03908116927f0000000000000000000000000000000000000000000000000000000000000000909116916370a0823191602480820192602092909190829003018186803b1580156133f957600080fd5b505afa15801561340d573d6000803e3d6000fd5b505050506040513d602081101561342357600080fd5b50516001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691906135f2565b6001600160a01b039182166000908152601c6020908152604080832093909416825291909152205490565b601354600160201b900460020b81565b600a5481565b60095481565b601081815481106134ad57600080fd5b6000918252602090912001546001600160a01b0316905081565b6003546001600160a01b031681565b6005546001600160a01b0316331461351b576040805162461bcd60e51b815260206004820152600360248201526228299960e91b604482015290519081900360640190fd5b6000841315613558576135586001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001633866135f2565b6000831315612e0657612e066001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001633856135f2565b6000828211156135ec576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b1790526110a8908490614b16565b3390565b6001600160a01b03831661368d5760405162461bcd60e51b8152600401808060200182810382526024815260200180615e356024913960400191505060405180910390fd5b6001600160a01b0382166136d25760405162461bcd60e51b8152600401808060200182810382526022815260200180615cf16022913960400191505060405180910390fd5b6001600160a01b038084166000818152601c6020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166137795760405162461bcd60e51b8152600401808060200182810382526025815260200180615e106025913960400191505060405180910390fd5b6001600160a01b0382166137be5760405162461bcd60e51b8152600401808060200182810382526023815260200180615c8c6023913960400191505060405180910390fd5b6137c98383836110a8565b61380681604051806060016040528060268152602001615d5a602691396001600160a01b0386166000908152601b60205260409020549190613891565b6001600160a01b038085166000908152601b602052604080822093909355908416815220546138359082613d3c565b6001600160a01b038084166000818152601b602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600081848411156139205760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156138e55781810151838201526020016138cd565b50505050905090810190601f1680156139125780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6000613932614a9f565b600061393c61131a565b905060008111613950576000915050611252565b8015613a6957600080600080613964614bc7565b935093509350935083600014801561397a575082155b1561398d57600095505050505050611252565b60105460005b81811015613a625760006139ad60108381548110612fda57fe5b905060006139bf8961152d8885613a71565b905060006139d18a61152d8886613a71565b90508160116000601087815481106139e557fe5b60009182526020808320909101546001600160a01b0316835282019290925260400181206002018054909201909155601080548392601192909188908110613a2957fe5b60009182526020808320909101546001600160a01b03168352820192909252604001902060030180549091019055505050600101613993565b5050505050505b600191505090565b600082613a80575060006112db565b82820282848281613a8d57fe5b04146114085760405162461bcd60e51b8152600401808060200182810382526021815260200180615da66021913960400191505060405180910390fd5b6000808211613b20576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b818381613b2957fe5b049392505050565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015613ba057600080fd5b505afa158015613bb4573d6000803e3d6000fd5b505050506040513d6020811015613bca57600080fd5b5051905090565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015613ba057600080fd5b6001600160a01b038216613c855760405162461bcd60e51b8152600401808060200182810382526021815260200180615def6021913960400191505060405180910390fd5b613c91826000836110a8565b613cce81604051806060016040528060228152602001615ccf602291396001600160a01b0385166000908152601b60205260409020549190613891565b6001600160a01b0383166000908152601b6020526040902055601d54613cf49082613595565b601d556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b600082820183811015611408576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6001600160a01b038116600090815260116020526040902060040154612baa5760108054600181019091557f1b6847dc741a1b0cd08d278845f9d819d87b734759afb55fe2de5cb82a9ae6720180546001600160a01b0319166001600160a01b03929092169182179055600090815260116020526040902043600490910155565b60235460009083908390613e56576040805162461bcd60e51b81526020600482015260016024820152600560fc1b604482015290519081900360640190fd5b60235481670de0b6b3a76400000281613e6b57fe5b04820192509250925092565b6000613e8161131a565b905060008111613e915750614105565b601054604080516370a0823160e01b815230600482015290516000916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016916370a0823191602480820192602092909190829003018186803b158015613efe57600080fd5b505afa158015613f12573d6000803e3d6000fd5b505050506040513d6020811015613f2857600080fd5b5051604080516370a0823160e01b815230600482015290519192506000916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016916370a08231916024808301926020929190829003018186803b158015613f9657600080fd5b505afa158015613faa573d6000803e3d6000fd5b505050506040513d6020811015613fc057600080fd5b5051905060005b838110156140ff576000613fe160108381548110612fda57fe5b905080156140f6576000613ff98761152d8785613a71565b9050600061400b8861152d8786613a71565b905060006140198383613e17565b505090508381146140f25761404f6010868154811061403457fe5b6000918252602090912001546001600160a01b031685613c40565b61407a6010868154811061405f57fe5b6000918252602090912001546001600160a01b031682614161565b82601160006010888154811061408c57fe5b60009182526020808320909101546001600160a01b031683528201929092526040018120919091556010805484926011929091899081106140c957fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020600101555b5050505b50600101613fc7565b50505050505b565b604080516001600160a01b0380861660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052612e06908590614b16565b6001600160a01b0382166141bc576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b6141c8600083836110a8565b601d546141d59082613d3c565b601d556001600160a01b0382166000908152601b60205260409020546141fb9082613d3c565b6001600160a01b0383166000818152601b602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b60055460408051633850c7bd60e01b815290516000926001600160a01b031691633850c7bd9160048083019260e0929190829003018186803b15801561429857600080fd5b505afa1580156142ac573d6000803e3d6000fd5b505050506040513d60e08110156142c257600080fd5b50602001519050600283900b1580156142de57508160020b6000145b1561438157600554600160d01b9004600290810b900b15801561430f5750600554600160b81b9004600290810b900b155b1561431a5750612d26565b6005546002600160b81b8204810b600160d01b8304820b03810b81900591600160a01b9004810b9081810b90848401900b8161435257fe5b600554919005919091029350600160a01b9004600290810b9081810b90838503900b8161437b57fe5b05029350505b61439c838383600560149054906101000a900460020b614ce0565b6013546000906143bf9060649061152d90600160381b900460ff16611527613b31565b905060006143e6606461152d601360079054906101000a900460ff1660ff16611527613bd1565b90506143f485858484614eb4565b60006143fe613b31565b9050600061440a613bd1565b9050600061445a7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000008585615041565b905080614494576040805162461bcd60e51b815260206004820152600360248201526204d43760ec1b604482015290519081900360640190fd5b5050505050505050565b60008060008360020b126144b5578260020b6144bd565b8260020b6000035b9050620d89e88111156144fb576040805162461bcd60e51b81526020600482015260016024820152601560fa1b604482015290519081900360640190fd5b60006001821661450f57600160801b614521565b6ffffcb933bd6fad37aa2d162d1a5940015b70ffffffffffffffffffffffffffffffffff1690506002821615614555576ffff97272373d413259a46990580e213a0260801c5b6004821615614574576ffff2e50f5f656932ef12357cf3c7fdcc0260801c5b6008821615614593576fffe5caca7e10e4e61c3624eaa0941cd00260801c5b60108216156145b2576fffcb9843d60f6159c9db58835c9266440260801c5b60208216156145d1576fff973b41fa98c081472e6896dfb254c00260801c5b60408216156145f0576fff2ea16466c96a3843ec78b326b528610260801c5b608082161561460f576ffe5dee046a99a2a811c461f1969c30530260801c5b61010082161561462f576ffcbe86c7900a88aedcffc83b479aa3a40260801c5b61020082161561464f576ff987a7253ac413176f2b074cf7815e540260801c5b61040082161561466f576ff3392b0822b70005940c7a398e4b70f30260801c5b61080082161561468f576fe7159475a2c29b7443b29c7fa6e889d90260801c5b6110008216156146af576fd097f3bdfd2022b8845ad8f792aa58250260801c5b6120008216156146cf576fa9f746462d870fdf8a65dc1f90e061e50260801c5b6140008216156146ef576f70d869a156d2a1b890bb3df62baf32f70260801c5b61800082161561470f576f31be135f97d08fd981231505542fcfa60260801c5b62010000821615614730576f09aa508b5b7a84e1c677de54f3e99bc90260801c5b62020000821615614750576e5d6af8dedb81196699c329225ee6040260801c5b6204000082161561476f576d2216e584f5fa1ea926041bedfe980260801c5b6208000082161561478c576b048a170391f7dc42444e8fa20260801c5b60008460020b13156147a75780600019816147a357fe5b0490505b600160201b8106156147ba5760016147bd565b60005b60ff16602082901c0192505050919050565b600080806000198587098686029250828110908390030390508061480557600084116147fa57600080fd5b508290049050611408565b80841161481157600080fd5b6000848688096000868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b600080600080600080614892308989615116565b6005546040805163514ea4bf60e01b81526004810184905290519293506001600160a01b039091169163514ea4bf9160248082019260a092909190829003018186803b1580156148e157600080fd5b505afa1580156148f5573d6000803e3d6000fd5b505050506040513d60a081101561490b57600080fd5b508051602082015160408301516060840151608090940151929c919b50995091975095509350505050565b6000806000600560009054906101000a90046001600160a01b03166001600160a01b0316633850c7bd6040518163ffffffff1660e01b815260040160e06040518083038186803b15801561498957600080fd5b505afa15801561499d573d6000803e3d6000fd5b505050506040513d60e08110156149b357600080fd5b505190506149d3816149c48861449e565b6149cd8861449e565b8761516c565b9250925050935093915050565b6000816001600160a01b031663db006a75846040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b158015614a2857600080fd5b505af1158015614a3c573d6000803e3d6000fd5b505050506040513d6020811015614a5257600080fd5b505160408051602081018390528181526024918101829052919250600080516020615e5983398151915291839181906060820190615d1382396040019250505060405180910390a1505050565b600080614aaa6110ad565b915091506000821180614abd5750600081115b15614b0e57614b0e7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000008484615208565b612d26615533565b6000614b6b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166156139092919063ffffffff16565b8051909150156110a857808060200190516020811015614b8a57600080fd5b50516110a85760405162461bcd60e51b815260040180806020018281038252602a815260200180615e79602a913960400191505060405180910390fd5b60008060008060006004549050614beb601754600654613d3c90919063ffffffff16565b9450614c04601854600754613d3c90919063ffffffff16565b93506000808215614cd0578615614c4457614c24606461152d8986613a71565b9150614c308783613595565b600a54909550614c409083613d3c565b600a555b8515614c7957614c59606461152d8886613a71565b9050614c658682613595565b600b54909450614c759082613d3c565b600b555b6000821180614c885750600081115b15614ccb576000614c998383613e17565b5050600254909150614cb4906001600160a01b031682614161565b600254614cc9906001600160a01b0316613d96565b505b614cd7565b8694508593505b50505090919293565b8260020b8460020b12614d1f576040805162461bcd60e51b8152602060048201526002602482015261563160f01b604482015290519081900360640190fd5b8160020b8460020b12614d5e576040805162461bcd60e51b81526020600482015260026024820152612b1960f11b604482015290519081900360640190fd5b8160020b8360020b13614d9d576040805162461bcd60e51b8152602060048201526002602482015261563360f01b604482015290519081900360640190fd5b620d89e719600285900b1215614ddf576040805162461bcd60e51b8152602060048201526002602482015261158d60f21b604482015290519081900360640190fd5b620d89e8600284900b1315614e20576040805162461bcd60e51b8152602060048201526002602482015261563560f01b604482015290519081900360640190fd5b8060020b8460020b81614e2f57fe5b0760020b15614e6a576040805162461bcd60e51b81526020600482015260026024820152612b1b60f11b604482015290519081900360640190fd5b8060020b8360020b81614e7957fe5b0760020b15612e06576040805162461bcd60e51b8152602060048201526002602482015261563760f01b604482015290519081900360640190fd5b60055460408051633850c7bd60e01b815290516000926001600160a01b031691633850c7bd9160048083019260e0929190829003018186803b158015614ef957600080fd5b505afa158015614f0d573d6000803e3d6000fd5b505050506040513d60e0811015614f2357600080fd5b505190506000614f4682614f368861449e565b614f3f8861449e565b878761562a565b60055460408051633c8a7d8d60e01b815230600482015260028a810b602483015289900b60448201526001600160801b038416606482015260a06084820152600060a4820181905282519495506001600160a01b0390931693633c8a7d8d9360c480840194938390030190829087803b158015614fc257600080fd5b505af1158015614fd6573d6000803e3d6000fd5b505050506040513d6040811015614fec57600080fd5b508051602090910151600955600855505060058054600295860b62ffffff908116600160b81b0262ffffff60b81b199690970b16600160d01b0262ffffff60d01b199091161793909316939093179091555050565b6016546013546014546019859055601a8490556000926001600160a01b0390811692600160401b900481169181169088168314156150aa57615082866156ee565b601554615098906001600160a01b03168261268b565b506150a48782876117db565b50615108565b826001600160a01b0316876001600160a01b031614156150ef576150cd856156ee565b6015546150e3906001600160a01b03168361268b565b506150a48883886117db565b6150fa8883886117db565b506151068782876117db565b505b506001979650505050505050565b6040805160609490941b6bffffffffffffffffffffffff1916602080860191909152600293840b60e890811b60348701529290930b90911b60378401528051808403601a018152603a9093019052815191012090565b600080836001600160a01b0316856001600160a01b0316111561518d579293925b846001600160a01b0316866001600160a01b0316116151b8576151b1858585615755565b91506151ff565b836001600160a01b0316866001600160a01b031610156151f1576151dd868585615755565b91506151ea8587856157be565b90506151ff565b6151fc8585856157be565b90505b94509492505050565b601654601354601454604080516370a0823160e01b815230600482015290516001600160a01b0394851694600160401b90940484169392831692600092908a16916370a0823191602480820192602092909190829003018186803b15801561526f57600080fd5b505afa158015615283573d6000803e3d6000fd5b505050506040513d602081101561529957600080fd5b5051604080516370a0823160e01b815230600482015290519192506000916001600160a01b038a16916370a08231916024808301926020929190829003018186803b1580156152e757600080fd5b505afa1580156152fb573d6000803e3d6000fd5b505050506040513d602081101561531157600080fd5b505190506001600160a01b03898116908616141561534a5761533387856149e0565b61533b615801565b61534586846149e0565b615394565b846001600160a01b0316886001600160a01b031614156153805761536e86846149e0565b615376615801565b61534587856149e0565b61538a86846149e0565b61539487856149e0565b6000615419838b6001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156153e757600080fd5b505afa1580156153fb573d6000803e3d6000fd5b505050506040513d602081101561541157600080fd5b505190613595565b9050600061546e838b6001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156153e757600080fd5b9050600082118015615481575060195482115b1561549857601954615494908390613595565b6017555b6000811180156154a95750601a5481115b156154c057601a546154bc908290613595565b6018555b601754601854604080516020810193909352828101919091526060808352600e908301526d3932b6b7bb32b632b73234b7339d60911b6080830152517fb626a5b8d37253edd986272eff7a2541a1538cad44624e413d51713f0b2a414d9181900360a00190a15050505050505050505050565b600554600160b81b9004600290810b900b15801561555f5750600554600160d01b9004600290810b900b155b1561556957614105565b60055460009061558e90600160b81b8104600290810b91600160d01b9004900b61487e565b505060055492935060009283925082915081906155c190600160b81b8104600290810b91600160d01b9004900b87615857565b93509350935093506000856001600160801b031611156140ff576155e58285613595565b6006556155f28184613595565b6007819055600654600c80549091019055600d805490910190555050505050565b60606156228484600085615a19565b949350505050565b6000836001600160a01b0316856001600160a01b0316111561564a579293925b846001600160a01b0316866001600160a01b0316116156755761566e858585615b69565b90506156e5565b836001600160a01b0316866001600160a01b031610156156d757600061569c878686615b69565b905060006156ab878986615bcc565b9050806001600160801b0316826001600160801b0316106156cc57806156ce565b815b925050506156e5565b6156e2858584615bcc565b90505b95945050505050565b8015612baa5760165460408051632e1a7d4d60e01b81526004810184905290516001600160a01b0390921691632e1a7d4d9160248082019260009290919082900301818387803b15801561574157600080fd5b505af11580156140ff573d6000803e3d6000fd5b6000826001600160a01b0316846001600160a01b03161115615775579192915b836001600160a01b03166157ae606060ff16846001600160801b0316901b8686036001600160a01b0316866001600160a01b03166147cf565b816157b557fe5b04949350505050565b6000826001600160a01b0316846001600160a01b031611156157de579192915b615622826001600160801b03168585036001600160a01b0316600160601b6147cf565b471561410557601660009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0476040518263ffffffff1660e01b81526004016000604051808303818588803b15801561574157600080fd5b60008080806001600160801b03851615615909576005546040805163a34123a760e01b815260028a810b600483015289900b60248201526001600160801b038816604482015281516001600160a01b039093169263a34123a7926064808401939192918290030181600087803b1580156158d057600080fd5b505af11580156158e4573d6000803e3d6000fd5b505050506040513d60408110156158fa57600080fd5b50805160209091015190945092505b600554604080516309e3d67b60e31b815230600482015260028a810b602483015289900b60448201526001600160801b0360648201819052608482015281516001600160a01b0390931692634f1eb3d89260a4808401939192918290030181600087803b15801561597957600080fd5b505af115801561598d573d6000803e3d6000fd5b505050506040513d60408110156159a357600080fd5b508051602091820151604080516001600160801b038a81168252948101899052808201889052928416606084018190529390911660808301819052905192945092507fe1b50fe06c4de1d84efe4528f0b6548f3c89b0be3f36f36deca670048545c559919081900360a00190a193509350935093565b606082471015615a5a5760405162461bcd60e51b8152600401808060200182810382526026815260200180615d806026913960400191505060405180910390fd5b615a6385615c09565b615ab4576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b600080866001600160a01b031685876040518082805190602001908083835b60208310615af25780518252601f199092019160209182019101615ad3565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114615b54576040519150601f19603f3d011682016040523d82523d6000602084013e615b59565b606091505b5091509150612079828286615c0f565b6000826001600160a01b0316846001600160a01b03161115615b89579192915b6000615bac856001600160a01b0316856001600160a01b0316600160601b6147cf565b90506156e5615bc784838888036001600160a01b03166147cf565b615c75565b6000826001600160a01b0316846001600160a01b03161115615bec579192915b615622615bc783600160601b8787036001600160a01b03166147cf565b3b151590565b60608315615c1e575081611408565b825115615c2e5782518084602001fd5b60405162461bcd60e51b81526020600482018181528451602484015284518593919283926044019190850190808383600083156138e55781810151838201526020016138cd565b806001600160801b038116811461268657600080fdfe45524332303a207472616e7366657220746f20746865207a65726f20616464726573735265656e7472616e637947756172643a207265656e7472616e742063616c6c0045524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332303a20617070726f766520746f20746865207a65726f206164647265737349662074686973206973206e6f7420302c2074686572652077617320616e206572726f7245786368616e6765205261746520287363616c65642075702062792031653138293a2045524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e6365416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f7745524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f20616464726573738d1cced004452bd270777a8c670f9f7e7c4fdde56f2db331fe289d39dc2624ad5361666545524332303a204552433230206f7065726174696f6e20646964206e6f74207375636365656445524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa26469706673582212207739279586c8fce3b72da9c53a6dacc681c936a6d1f6bccdedccff8fcda1e6d664736f6c63430007060033"

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

// AccumulateUniswapFees0 is a free data retrieval call binding the contract method 0xaa489e06.
//
// Solidity: function AccumulateUniswapFees0() view returns(uint256)
func (_Api *ApiCaller) AccumulateUniswapFees0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "AccumulateUniswapFees0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulateUniswapFees0 is a free data retrieval call binding the contract method 0xaa489e06.
//
// Solidity: function AccumulateUniswapFees0() view returns(uint256)
func (_Api *ApiSession) AccumulateUniswapFees0() (*big.Int, error) {
	return _Api.Contract.AccumulateUniswapFees0(&_Api.CallOpts)
}

// AccumulateUniswapFees0 is a free data retrieval call binding the contract method 0xaa489e06.
//
// Solidity: function AccumulateUniswapFees0() view returns(uint256)
func (_Api *ApiCallerSession) AccumulateUniswapFees0() (*big.Int, error) {
	return _Api.Contract.AccumulateUniswapFees0(&_Api.CallOpts)
}

// AccumulateUniswapFees1 is a free data retrieval call binding the contract method 0xc29188bc.
//
// Solidity: function AccumulateUniswapFees1() view returns(uint256)
func (_Api *ApiCaller) AccumulateUniswapFees1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "AccumulateUniswapFees1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulateUniswapFees1 is a free data retrieval call binding the contract method 0xc29188bc.
//
// Solidity: function AccumulateUniswapFees1() view returns(uint256)
func (_Api *ApiSession) AccumulateUniswapFees1() (*big.Int, error) {
	return _Api.Contract.AccumulateUniswapFees1(&_Api.CallOpts)
}

// AccumulateUniswapFees1 is a free data retrieval call binding the contract method 0xc29188bc.
//
// Solidity: function AccumulateUniswapFees1() view returns(uint256)
func (_Api *ApiCallerSession) AccumulateUniswapFees1() (*big.Int, error) {
	return _Api.Contract.AccumulateUniswapFees1(&_Api.CallOpts)
}

// Assetholder is a free data retrieval call binding the contract method 0x5bb6aa85.
//
// Solidity: function Assetholder(address ) view returns(uint256 deposit0, uint256 deposit1, uint256 fees0, uint256 fees1, uint256 block)
func (_Api *ApiCaller) Assetholder(opts *bind.CallOpts, arg0 common.Address) (struct {
	Deposit0 *big.Int
	Deposit1 *big.Int
	Fees0    *big.Int
	Fees1    *big.Int
	Block    *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "Assetholder", arg0)

	outstruct := new(struct {
		Deposit0 *big.Int
		Deposit1 *big.Int
		Fees0    *big.Int
		Fees1    *big.Int
		Block    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Deposit0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Deposit1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Fees0 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Fees1 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Assetholder is a free data retrieval call binding the contract method 0x5bb6aa85.
//
// Solidity: function Assetholder(address ) view returns(uint256 deposit0, uint256 deposit1, uint256 fees0, uint256 fees1, uint256 block)
func (_Api *ApiSession) Assetholder(arg0 common.Address) (struct {
	Deposit0 *big.Int
	Deposit1 *big.Int
	Fees0    *big.Int
	Fees1    *big.Int
	Block    *big.Int
}, error) {
	return _Api.Contract.Assetholder(&_Api.CallOpts, arg0)
}

// Assetholder is a free data retrieval call binding the contract method 0x5bb6aa85.
//
// Solidity: function Assetholder(address ) view returns(uint256 deposit0, uint256 deposit1, uint256 fees0, uint256 fees1, uint256 block)
func (_Api *ApiCallerSession) Assetholder(arg0 common.Address) (struct {
	Deposit0 *big.Int
	Deposit1 *big.Int
	Fees0    *big.Int
	Fees1    *big.Int
	Block    *big.Int
}, error) {
	return _Api.Contract.Assetholder(&_Api.CallOpts, arg0)
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

// AccruedProtocolFees0 is a free data retrieval call binding the contract method 0xeae989a2.
//
// Solidity: function accruedProtocolFees0() view returns(uint256)
func (_Api *ApiCaller) AccruedProtocolFees0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "accruedProtocolFees0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedProtocolFees0 is a free data retrieval call binding the contract method 0xeae989a2.
//
// Solidity: function accruedProtocolFees0() view returns(uint256)
func (_Api *ApiSession) AccruedProtocolFees0() (*big.Int, error) {
	return _Api.Contract.AccruedProtocolFees0(&_Api.CallOpts)
}

// AccruedProtocolFees0 is a free data retrieval call binding the contract method 0xeae989a2.
//
// Solidity: function accruedProtocolFees0() view returns(uint256)
func (_Api *ApiCallerSession) AccruedProtocolFees0() (*big.Int, error) {
	return _Api.Contract.AccruedProtocolFees0(&_Api.CallOpts)
}

// AccruedProtocolFees1 is a free data retrieval call binding the contract method 0xa00fa77f.
//
// Solidity: function accruedProtocolFees1() view returns(uint256)
func (_Api *ApiCaller) AccruedProtocolFees1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "accruedProtocolFees1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedProtocolFees1 is a free data retrieval call binding the contract method 0xa00fa77f.
//
// Solidity: function accruedProtocolFees1() view returns(uint256)
func (_Api *ApiSession) AccruedProtocolFees1() (*big.Int, error) {
	return _Api.Contract.AccruedProtocolFees1(&_Api.CallOpts)
}

// AccruedProtocolFees1 is a free data retrieval call binding the contract method 0xa00fa77f.
//
// Solidity: function accruedProtocolFees1() view returns(uint256)
func (_Api *ApiCallerSession) AccruedProtocolFees1() (*big.Int, error) {
	return _Api.Contract.AccruedProtocolFees1(&_Api.CallOpts)
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

// CurUni0 is a free data retrieval call binding the contract method 0xa833d7ca.
//
// Solidity: function curUni0() view returns(uint256)
func (_Api *ApiCaller) CurUni0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "curUni0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurUni0 is a free data retrieval call binding the contract method 0xa833d7ca.
//
// Solidity: function curUni0() view returns(uint256)
func (_Api *ApiSession) CurUni0() (*big.Int, error) {
	return _Api.Contract.CurUni0(&_Api.CallOpts)
}

// CurUni0 is a free data retrieval call binding the contract method 0xa833d7ca.
//
// Solidity: function curUni0() view returns(uint256)
func (_Api *ApiCallerSession) CurUni0() (*big.Int, error) {
	return _Api.Contract.CurUni0(&_Api.CallOpts)
}

// CurUni1 is a free data retrieval call binding the contract method 0xf10d8fc8.
//
// Solidity: function curUni1() view returns(uint256)
func (_Api *ApiCaller) CurUni1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "curUni1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurUni1 is a free data retrieval call binding the contract method 0xf10d8fc8.
//
// Solidity: function curUni1() view returns(uint256)
func (_Api *ApiSession) CurUni1() (*big.Int, error) {
	return _Api.Contract.CurUni1(&_Api.CallOpts)
}

// CurUni1 is a free data retrieval call binding the contract method 0xf10d8fc8.
//
// Solidity: function curUni1() view returns(uint256)
func (_Api *ApiCallerSession) CurUni1() (*big.Int, error) {
	return _Api.Contract.CurUni1(&_Api.CallOpts)
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

// LFees0 is a free data retrieval call binding the contract method 0x580b82e3.
//
// Solidity: function lFees0() view returns(uint256)
func (_Api *ApiCaller) LFees0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "lFees0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LFees0 is a free data retrieval call binding the contract method 0x580b82e3.
//
// Solidity: function lFees0() view returns(uint256)
func (_Api *ApiSession) LFees0() (*big.Int, error) {
	return _Api.Contract.LFees0(&_Api.CallOpts)
}

// LFees0 is a free data retrieval call binding the contract method 0x580b82e3.
//
// Solidity: function lFees0() view returns(uint256)
func (_Api *ApiCallerSession) LFees0() (*big.Int, error) {
	return _Api.Contract.LFees0(&_Api.CallOpts)
}

// LFees1 is a free data retrieval call binding the contract method 0x75bd7111.
//
// Solidity: function lFees1() view returns(uint256)
func (_Api *ApiCaller) LFees1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "lFees1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LFees1 is a free data retrieval call binding the contract method 0x75bd7111.
//
// Solidity: function lFees1() view returns(uint256)
func (_Api *ApiSession) LFees1() (*big.Int, error) {
	return _Api.Contract.LFees1(&_Api.CallOpts)
}

// LFees1 is a free data retrieval call binding the contract method 0x75bd7111.
//
// Solidity: function lFees1() view returns(uint256)
func (_Api *ApiCallerSession) LFees1() (*big.Int, error) {
	return _Api.Contract.LFees1(&_Api.CallOpts)
}

// LastRebalance is a free data retrieval call binding the contract method 0x106b9ca1.
//
// Solidity: function lastRebalance() view returns(uint256)
func (_Api *ApiCaller) LastRebalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "lastRebalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRebalance is a free data retrieval call binding the contract method 0x106b9ca1.
//
// Solidity: function lastRebalance() view returns(uint256)
func (_Api *ApiSession) LastRebalance() (*big.Int, error) {
	return _Api.Contract.LastRebalance(&_Api.CallOpts)
}

// LastRebalance is a free data retrieval call binding the contract method 0x106b9ca1.
//
// Solidity: function lastRebalance() view returns(uint256)
func (_Api *ApiCallerSession) LastRebalance() (*big.Int, error) {
	return _Api.Contract.LastRebalance(&_Api.CallOpts)
}

// LastTick is a free data retrieval call binding the contract method 0x3dfa5d87.
//
// Solidity: function lastTick() view returns(int24)
func (_Api *ApiCaller) LastTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "lastTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTick is a free data retrieval call binding the contract method 0x3dfa5d87.
//
// Solidity: function lastTick() view returns(int24)
func (_Api *ApiSession) LastTick() (*big.Int, error) {
	return _Api.Contract.LastTick(&_Api.CallOpts)
}

// LastTick is a free data retrieval call binding the contract method 0x3dfa5d87.
//
// Solidity: function lastTick() view returns(int24)
func (_Api *ApiCallerSession) LastTick() (*big.Int, error) {
	return _Api.Contract.LastTick(&_Api.CallOpts)
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

// UFees0 is a free data retrieval call binding the contract method 0xd595a12b.
//
// Solidity: function uFees0() view returns(uint256)
func (_Api *ApiCaller) UFees0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "uFees0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UFees0 is a free data retrieval call binding the contract method 0xd595a12b.
//
// Solidity: function uFees0() view returns(uint256)
func (_Api *ApiSession) UFees0() (*big.Int, error) {
	return _Api.Contract.UFees0(&_Api.CallOpts)
}

// UFees0 is a free data retrieval call binding the contract method 0xd595a12b.
//
// Solidity: function uFees0() view returns(uint256)
func (_Api *ApiCallerSession) UFees0() (*big.Int, error) {
	return _Api.Contract.UFees0(&_Api.CallOpts)
}

// UFees1 is a free data retrieval call binding the contract method 0x82fc2543.
//
// Solidity: function uFees1() view returns(uint256)
func (_Api *ApiCaller) UFees1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "uFees1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UFees1 is a free data retrieval call binding the contract method 0x82fc2543.
//
// Solidity: function uFees1() view returns(uint256)
func (_Api *ApiSession) UFees1() (*big.Int, error) {
	return _Api.Contract.UFees1(&_Api.CallOpts)
}

// UFees1 is a free data retrieval call binding the contract method 0x82fc2543.
//
// Solidity: function uFees1() view returns(uint256)
func (_Api *ApiCallerSession) UFees1() (*big.Int, error) {
	return _Api.Contract.UFees1(&_Api.CallOpts)
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
	Liquidity *big.Int
	Burned0   *big.Int
	Burned1   *big.Int
	Collect0  *big.Int
	Collect1  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectFees is a free log retrieval operation binding the contract event 0xe1b50fe06c4de1d84efe4528f0b6548f3c89b0be3f36f36deca670048545c559.
//
// Solidity: event CollectFees(uint128 liquidity, uint256 burned0, uint256 burned1, uint256 collect0, uint256 collect1)
func (_Api *ApiFilterer) FilterCollectFees(opts *bind.FilterOpts) (*ApiCollectFeesIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "CollectFees")
	if err != nil {
		return nil, err
	}
	return &ApiCollectFeesIterator{contract: _Api.contract, event: "CollectFees", logs: logs, sub: sub}, nil
}

// WatchCollectFees is a free log subscription operation binding the contract event 0xe1b50fe06c4de1d84efe4528f0b6548f3c89b0be3f36f36deca670048545c559.
//
// Solidity: event CollectFees(uint128 liquidity, uint256 burned0, uint256 burned1, uint256 collect0, uint256 collect1)
func (_Api *ApiFilterer) WatchCollectFees(opts *bind.WatchOpts, sink chan<- *ApiCollectFees) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "CollectFees")
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

// ParseCollectFees is a log parse operation binding the contract event 0xe1b50fe06c4de1d84efe4528f0b6548f3c89b0be3f36f36deca670048545c559.
//
// Solidity: event CollectFees(uint128 liquidity, uint256 burned0, uint256 burned1, uint256 collect0, uint256 collect1)
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

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
const ApiABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cToken0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cToken1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cEth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_protocolFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxTotalSupply\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"_maxTwapDeviation\",\"type\":\"int24\"},{\"internalType\":\"uint32\",\"name\":\"_twapDuration\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_uniPortionRate\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesToVault0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesToVault1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesToProtocol0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesToProtocol1\",\"type\":\"uint256\"}],\"name\":\"CollectFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"ErrorLogging\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"subject\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"GeneralA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"subject\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"GeneralB\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"subject\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"GeneralS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MyLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MyLog2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance1\",\"type\":\"uint256\"}],\"name\":\"RebalanceLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"Snapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"nameToken0\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"nameToken1\",\"type\":\"string\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"AccumulateUniswapFees0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AccumulateUniswapFees1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CEther\",\"outputs\":[{\"internalType\":\"contractIcEther\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CToken0\",\"outputs\":[{\"internalType\":\"contractIcErc20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CToken1\",\"outputs\":[{\"internalType\":\"contractIcErc20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedProtocolFees0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedProtocolFees1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cHigh\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cLow\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"collectProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToken1\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20Contract\",\"type\":\"address\"}],\"name\":\"getCTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getETHBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"getPositionAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"getSSLiquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTVL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTwap\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUniswapPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRebalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTick\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTwapDeviation\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"redeemType\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_cErc20Contract\",\"type\":\"address\"}],\"name\":\"redeemCErc20Tokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"redeemType\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_cEtherContract\",\"type\":\"address\"}],\"name\":\"redeemCEth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"setGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMax\",\"type\":\"uint256\"}],\"name\":\"setMaxTotalSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_maxTwapDeviation\",\"type\":\"int24\"}],\"name\":\"setMaxTwapDeviation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_team\",\"type\":\"address\"}],\"name\":\"setTeam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_twapDuration\",\"type\":\"uint32\"}],\"name\":\"setTwapDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"ratio\",\"type\":\"uint8\"}],\"name\":\"setUniPortionRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"newLow\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"newHigh\",\"type\":\"int24\"},{\"internalType\":\"int256\",\"name\":\"swapAmount\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"}],\"name\":\"strategy0\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"newLow\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"newHigh\",\"type\":\"int24\"}],\"name\":\"strategy1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20Contract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cErc20Contract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_numTokensToSupply\",\"type\":\"uint256\"}],\"name\":\"supplyErc20ToCompound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_cEtherContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ctoken\",\"type\":\"address\"}],\"name\":\"supplyEthToCompound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"swapAmount\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"team\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"twapDuration\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whiteHacker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ApiBin is the compiled bytecode used for deploying new contracts.
var ApiBin = "0x6101606040523480156200001257600080fd5b50604051620064eb380380620064eb83398181016040526101408110156200003957600080fd5b5080516020808301516040808501516060860151608087015160a088015160c089015160e08a01516101008b0151610120909b015187518089018952600d81526c2b34b0a632b732102a37b5b2b760991b818c019081528951808b01909a526003808b526215931560ea1b9c8b019c909c5281519c9d9a9c989b979a96999598949793969592949193620000d09290919062000482565b508051620000e690600490602084019062000482565b50506005805460ff19166012179055506001600655600780546001600160a01b0319908116339081179092556008805490911690911790556001600160601b031960608b901b1660805260408051630dfe168160e01b815290516001600160a01b038c1691630dfe1681916004808301926020929190829003018186803b1580156200017157600080fd5b505afa15801562000186573d6000803e3d6000fd5b505050506040513d60208110156200019d57600080fd5b505160601b6001600160601b03191660a0526040805163d21220a760e01b815290516001600160a01b038c169163d21220a7916004808301926020929190829003018186803b158015620001f057600080fd5b505afa15801562000205573d6000803e3d6000fd5b505050506040513d60208110156200021c57600080fd5b50516001600160601b0319606091821b811660c05289821b81166101005288821b8116610120529087901b16610140526001600160a01b03891662000291576040805162461bcd60e51b815260206004808301919091526024820152630ae8aa8960e31b604482015290519081900360640190fd5b88600a60006101000a8154816001600160a01b0302191690836001600160a01b0316021790555084600d81905550896001600160a01b031663d0c93a7c6040518163ffffffff1660e01b815260040160206040518083038186803b158015620002f957600080fd5b505afa1580156200030e573d6000803e3d6000fd5b505050506040513d60208110156200032557600080fd5b5051600290810b900b60e81b60e052620f424085106200037e576040805162461bcd60e51b815260206004820152600f60248201526e70726f746f636f6c4665655261746560881b604482015290519081900360640190fd5b600b8490556013805460ff83166a01000000000000000000000260ff60501b1963ffffffff861663010000000266ffffffff00000019600289900b62ffffff81166701000000000000000262ffffff60381b19909616959095171617161790915560001262000427576040805162461bcd60e51b815260206004820152601060248201526f36b0bc2a3bb0b82232bb34b0ba34b7b760811b604482015290519081900360640190fd5b60008263ffffffff161162000472576040805162461bcd60e51b815260206004820152600c60248201526b3a3bb0b8223ab930ba34b7b760a11b604482015290519081900360640190fd5b505050505050505050506200052e565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282620004ba576000855562000505565b82601f10620004d557805160ff191683800117855562000505565b8280016001018555821562000505579182015b8281111562000505578251825591602001919060010190620004e8565b506200051392915062000517565b5090565b5b8082111562000513576000815560010162000518565b60805160601c60a05160601c60c05160601c60e05160e81c6101005160601c6101205160601c6101405160601c615e55620006966000398061251452806147fe528061485d5250806111425280611c195280611c8852806147c7528061504c5250806110885280611b75528061369752806147a5528061502a525080612a885280612af952806136d352806144c2528061453d525080611057528061175d528061190252806121655280612206528061247052806136f7528061379b52806138b25280613a785280613c2152806147775280614ffc52508061100d52806112df528061172352806117be528061209752806121385280612cf8528061376152806138755280613a3d5280613be452806147555280614fda525080611309528061143e52806126a152806129ec5280612d6852806137245280613b6c528061463952806148cd528061498b5280614a5f5280614d995280614e6152806151415250615e556000f3fe6080604052600436106103e85760003560e01c8063629d940511610208578063c311d04911610118578063d3487997116100ab578063dd62ed3e1161007a578063dd62ed3e14610e10578063e2bbb15814610e4b578063e7c7cb9114610e99578063eae989a214610eae578063fa461e3314610ec3576103ef565b8063d348799714610d1c578063d368867f14610da3578063d8d7f96f14610db8578063dc2c256f14610dcd576103ef565b8063cda206b0116100e7578063cda206b014610c9f578063d0ace48e14610c49578063d0c93a7c14610cf2578063d21220a714610d07576103ef565b8063c311d04914610bef578063c433c80a14610c19578063c6fb501414610c49578063cbcf94aa14610c8a576103ef565b806397b3fcaa1161019b578063a91ef6eb1161016a578063a91ef6eb14610b28578063aa489e0614610b5f578063ab033ea914610b74578063bb47dc1c14610ba7578063c29188bc14610bda576103ef565b806397b3fcaa14610a8c578063a00fa77f14610aa1578063a457c2d714610ab6578063a9059cbb14610aef576103ef565b806373232c60116101d757806373232c6014610a1f57806385f2aef214610a4d5780638e843c6c14610a6257806395d89b4114610a77576103ef565b8063629d94051461098157806368a5787a146109965780636e947298146109d757806370a08231146109ec576103ef565b8063313ce567116103035780633f3e4c11116102965780635aa6e675116102655780635aa6e675146108b25780635d29dab4146108c75780635d752a9a14610908578063624a177a1461091d5780636253c7181461094a576103ef565b80633f3e4c111461084957806341aec5381461087357806353257e001461088857806358f858801461089d576103ef565b80633bdd61b7116102d25780633bdd61b7146107955780633cbff3fe146107aa5780633dfa5d87146107d75780633e5aae6614610803576103ef565b8063313ce567146106d9578063323807171461070457806339509351146107195780633b7ba25f14610752576103ef565b806318160ddd1161037b5780632ab4d0521161034a5780632ab4d052146106705780632b6e1923146106855780632e1a7d4d1461069a5780632e4b8f60146106c4576103ef565b806318160ddd146105d5578063238efcbc146105ea57806323b872dd146105ff57806326d8954514610642576103ef565b8063095ea7b3116103b7578063095ea7b31461051b5780630dfe168114610568578063106b9ca11461059957806316f0115b146105c0576103ef565b80630430c130146103f157806306b0b6561461043057806306fdde031461045e578063095cf5c6146104e8576103ef565b366103ef57005b005b3480156103fd57600080fd5b506103ef6004803603606081101561041457600080fd5b50803590602081013590604001356001600160a01b0316610f4a565b34801561043c57600080fd5b50610445611083565b6040805192835260208301919091528051918290030190f35b34801561046a57600080fd5b506104736111bb565b6040805160208082528351818301528351919283929083019185019080838360005b838110156104ad578181015183820152602001610495565b50505050905090810190601f1680156104da5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156104f457600080fd5b506103ef6004803603602081101561050b57600080fd5b50356001600160a01b0316611251565b34801561052757600080fd5b506105546004803603604081101561053e57600080fd5b506001600160a01b0381351690602001356112bf565b604080519115158252519081900360200190f35b34801561057457600080fd5b5061057d6112dd565b604080516001600160a01b039092168252519081900360200190f35b3480156105a557600080fd5b506105ae611301565b60408051918252519081900360200190f35b3480156105cc57600080fd5b5061057d611307565b3480156105e157600080fd5b506105ae61132b565b3480156105f657600080fd5b506103ef611331565b34801561060b57600080fd5b506105546004803603606081101561062257600080fd5b506001600160a01b03813581169160208101359091169060400135611398565b34801561064e57600080fd5b50610657611420565b6040805163ffffffff9092168252519081900360200190f35b34801561067c57600080fd5b506105ae611433565b34801561069157600080fd5b506105ae611439565b3480156106a657600080fd5b50610445600480360360208110156106bd57600080fd5b50356114f2565b3480156106d057600080fd5b50610445611b4c565b3480156106e557600080fd5b506106ee611c7d565b6040805160ff9092168252519081900360200190f35b34801561071057600080fd5b5061057d611c86565b34801561072557600080fd5b506105546004803603604081101561073c57600080fd5b506001600160a01b038135169060200135611caa565b34801561075e57600080fd5b506105ae6004803603606081101561077557600080fd5b506001600160a01b03813581169160208101359091169060400135611cf8565b3480156107a157600080fd5b506103ef612045565b3480156107b657600080fd5b506103ef600480360360208110156107cd57600080fd5b503560020b61225f565b3480156107e357600080fd5b506107ec61231e565b6040805160029290920b8252519081900360200190f35b34801561080f57600080fd5b506103ef6004803603608081101561082657600080fd5b508035600290810b91602081013590910b90604081013590606001351515612327565b34801561085557600080fd5b506103ef6004803603602081101561086c57600080fd5b50356123c5565b34801561087f57600080fd5b506105ae612466565b34801561089457600080fd5b5061057d612512565b3480156108a957600080fd5b506105ae612536565b3480156108be57600080fd5b5061057d61253c565b3480156108d357600080fd5b506103ef600480360360608110156108ea57600080fd5b508035906001600160a01b036020820135811691604001351661254b565b34801561091457600080fd5b506107ec6125f6565b34801561092957600080fd5b506103ef6004803603602081101561094057600080fd5b503560ff166128ac565b34801561095657600080fd5b506103ef6004803603604081101561096d57600080fd5b508035600290810b9160200135900b612955565b34801561098d57600080fd5b506105ae612cee565b3480156109a257600080fd5b50610445600480360360608110156109b957600080fd5b508035906020810135151590604001356001600160a01b0316612d63565b3480156109e357600080fd5b506105ae612ec8565b3480156109f857600080fd5b506105ae60048036036020811015610a0f57600080fd5b50356001600160a01b0316612ecc565b61055460048036036040811015610a3557600080fd5b506001600160a01b0381358116916020013516612eeb565b348015610a5957600080fd5b5061057d6130cb565b348015610a6e57600080fd5b506107ec6130da565b348015610a8357600080fd5b506104736130ea565b348015610a9857600080fd5b5061044561314b565b348015610aad57600080fd5b506105ae6131ae565b348015610ac257600080fd5b5061055460048036036040811015610ad957600080fd5b506001600160a01b0381351690602001356131b4565b348015610afb57600080fd5b5061055460048036036040811015610b1257600080fd5b506001600160a01b03813516906020013561321c565b348015610b3457600080fd5b5061044560048036036040811015610b4b57600080fd5b508035600290810b9160200135900b613230565b348015610b6b57600080fd5b506105ae6132c7565b348015610b8057600080fd5b506103ef60048036036020811015610b9757600080fd5b50356001600160a01b03166132cd565b348015610bb357600080fd5b506105ae60048036036020811015610bca57600080fd5b50356001600160a01b031661333b565b348015610be657600080fd5b506105ae6133bc565b348015610bfb57600080fd5b506103ef60048036036020811015610c1257600080fd5b50356133c2565b348015610c2557600080fd5b506103ef60048036036020811015610c3c57600080fd5b503563ffffffff16613493565b348015610c5557600080fd5b5061055460048036036060811015610c6c57600080fd5b508035906020810135151590604001356001600160a01b031661354b565b348015610c9657600080fd5b5061057d613695565b348015610cab57600080fd5b50610cd660048036036040811015610cc257600080fd5b508035600290810b9160200135900b6136b9565b604080516001600160801b039092168252519081900360200190f35b348015610cfe57600080fd5b506107ec6136d1565b348015610d1357600080fd5b5061057d6136f5565b348015610d2857600080fd5b506103ef60048036036060811015610d3f57600080fd5b813591602081013591810190606081016040820135600160201b811115610d6557600080fd5b820183602082011115610d7757600080fd5b803590602001918460018302840111600160201b83111715610d9857600080fd5b509092509050613719565b348015610daf57600080fd5b506107ec6137c8565b348015610dc457600080fd5b506103ef6137d1565b348015610dd957600080fd5b506103ef60048036036060811015610df057600080fd5b506001600160a01b03813581169160208101359160409091013516613827565b348015610e1c57600080fd5b506105ae60048036036040811015610e3357600080fd5b506001600160a01b0381358116916020013516613934565b348015610e5757600080fd5b50610e7b60048036036040811015610e6e57600080fd5b508035906020013561395f565b60408051938452602084019290925282820152519081900360600190f35b348015610ea557600080fd5b506107ec613b4b565b348015610eba57600080fd5b506105ae613b5b565b348015610ecf57600080fd5b506103ef60048036036060811015610ee657600080fd5b813591602081013591810190606081016040820135600160201b811115610f0c57600080fd5b820183602082011115610f1e57600080fd5b803590602001918460018302840111600160201b83111715610f3f57600080fd5b509092509050613b61565b6007546001600160a01b03163314610f96576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b82600e5410158015610faa575081600f5410155b610fea576040805162461bcd60e51b815260206004820152600c60248201526b70726f746f636f6c6665657360a01b604482015290519081900360640190fd5b821561103457600e54610ffd9084613c48565b600e556110346001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168285613ca5565b811561107e57600f546110479083613c48565b600f5561107e6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168284613ca5565b505050565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156110f357600080fd5b505afa158015611107573d6000803e3d6000fd5b505050506040513d602081101561111d57600080fd5b5051604080516370a0823160e01b815230600482015290519193506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016916370a0823191602480820192602092909190829003018186803b15801561118957600080fd5b505afa15801561119d573d6000803e3d6000fd5b505050506040513d60208110156111b357600080fd5b505191929050565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156112475780601f1061121c57610100808354040283529160200191611247565b820191906000526020600020905b81548152906001019060200180831161122a57829003601f168201915b5050505050905090565b6007546001600160a01b0316331461129d576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b600880546001600160a01b0319166001600160a01b0392909216919091179055565b60006112d36112cc613cf7565b8484613cfb565b5060015b92915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60125481565b7f000000000000000000000000000000000000000000000000000000000000000081565b60025490565b6009546001600160a01b03163314611384576040805162461bcd60e51b815260206004820152601160248201527070656e64696e67476f7665726e616e636560781b604482015290519081900360640190fd5b600780546001600160a01b03191633179055565b60006113a5848484613de7565b611415846113b1613cf7565b61141085604051806060016040528060288152602001615d1f602891396001600160a01b038a166000908152600160205260408120906113ef613cf7565b6001600160a01b031681526020810191909152604001600020549190613f42565b613cfb565b5060015b9392505050565b6013546301000000900463ffffffff1681565b600b5481565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633850c7bd6040518163ffffffff1660e01b815260040160e06040518083038186803b15801561149557600080fd5b505afa1580156114a9573d6000803e3d6000fd5b505050506040513d60e08110156114bf57600080fd5b5051905060c06114ea670de0b6b3a76400006114e46001600160a01b03851680613fd9565b90613fd9565b901c91505090565b6000806002600654141561153b576040805162461bcd60e51b815260206004820152601f6024820152600080516020615c07833981519152604482015290519081900360640190fd5b600260065533801580159061155957506001600160a01b0381163014155b61158f576040805162461bcd60e51b8152602060048201526002602482015261746f60f01b604482015290519081900360640190fd5b6000841180156115a0575060648411155b6115db576040805162461bcd60e51b81526020600482015260076024820152661c195c98d95b9d60ca1b604482015290519081900360640190fd5b60006115f560646115ef876114e433612ecc565b90614032565b9050600061160161132b565b90506000811161163e576040805162461bcd60e51b815260206004820152600360248201526207473360ec1b604482015290519081900360640190fd5b8082111561167f576040805162461bcd60e51b8152602060048201526009602482015268736861726573202d3160b81b604482015290519081900360640190fd5b6116893383614099565b600061169b826115ef856114e4612cee565b905060006116af836115ef866114e4612466565b600c5490915060009081906116d490600281810b9163010000009004900b8888614195565b915091506000806116e5888861423b565b90925090506116fe826116f888876142ec565b906142ec565b9a5061170e816116f887866142ec565b99508a1561174a5761174a6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168a8d613ca5565b8915611784576117846001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168a8c613ca5565b886001600160a01b0316336001600160a01b03167f9167547f76aeba7bf8001d943eb4573822255e96f6ee40278381b71fad6af5ad8a8e8e7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166306fdde036040518163ffffffff1660e01b815260040160006040518083038186803b15801561181557600080fd5b505afa158015611829573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561185257600080fd5b8101908080516040519392919084600160201b82111561187157600080fd5b90830190602082018581111561188657600080fd5b8251600160201b81118282018810171561189f57600080fd5b82525081516020918201929091019080838360005b838110156118cc5781810151838201526020016118b4565b50505050905090810190601f1680156118f95780820380516001836020036101000a031916815260200191505b506040525050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166306fdde036040518163ffffffff1660e01b815260040160006040518083038186803b15801561195957600080fd5b505afa15801561196d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561199657600080fd5b8101908080516040519392919084600160201b8211156119b557600080fd5b9083019060208201858111156119ca57600080fd5b8251600160201b8111828201881017156119e357600080fd5b82525081516020918201929091019080838360005b83811015611a105781810151838201526020016119f8565b50505050905090810190601f168015611a3d5780820380516001836020036101000a031916815260200191505b50604052505050604051808681526020018581526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015611a97578181015183820152602001611a7f565b50505050905090810190601f168015611ac45780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015611af7578181015183820152602001611adf565b50505050905090810190601f168015611b245780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a35050505050505050506001600681905550915091565b600080600080611b5a611083565b915091506000611c05611bfe6b204fce5e3e250261100000007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663182df0f56040518163ffffffff1660e01b815260040160206040518083038186803b158015611bcc57600080fd5b505afa158015611be0573d6000803e3d6000fd5b505050506040513d6020811015611bf657600080fd5b505190614032565b8490613fd9565b90506000611c70611bfe662386f26fc100007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663182df0f56040518163ffffffff1660e01b815260040160206040518083038186803b158015611bcc57600080fd5b9195509093505050509091565b60055460ff1690565b7f000000000000000000000000000000000000000000000000000000000000000081565b60006112d3611cb7613cf7565b846114108560016000611cc8613cf7565b6001600160a01b03908116825260208083019390935260409182016000908120918c1681529252902054906142ec565b6000836001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015611d4757600080fd5b505afa158015611d5b573d6000803e3d6000fd5b505050506040513d6020811015611d7157600080fd5b5051821115611db1576040805162461bcd60e51b815260206004820152600760248201526662616c616e636560c81b604482015290519081900360640190fd5b600084905060008490506000816001600160a01b031663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015611df857600080fd5b505af1158015611e0c573d6000803e3d6000fd5b505050506040513d6020811015611e2257600080fd5b50516040805160208101839052818152601b818301527f45786368616e6765205261746520287363616c6564207570293a20000000000060608201529051919250600080516020615db1833981519152919081900360800190a16000826001600160a01b031663ae9d70b06040518163ffffffff1660e01b8152600401602060405180830381600087803b158015611eb957600080fd5b505af1158015611ecd573d6000803e3d6000fd5b505050506040513d6020811015611ee357600080fd5b505160408051602081018390528181526018818301527f537570706c7920526174653a20287363616c656420757029000000000000000060608201529051919250600080516020615db1833981519152919081900360800190a1836001600160a01b031663095ea7b388886040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b158015611f9457600080fd5b505af1158015611fa8573d6000803e3d6000fd5b505050506040513d6020811015611fbe57600080fd5b50506040805163140e25ad60e31b81526004810188905290516000916001600160a01b0386169163a0712d689160248082019260209290919082900301818787803b15801561200c57600080fd5b505af1158015612020573d6000803e3d6000fd5b505050506040513d602081101561203657600080fd5b50519998505050505050505050565b6007546001600160a01b03163314612091576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b61215f337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561210257600080fd5b505afa158015612116573d6000803e3d6000fd5b505050506040513d602081101561212c57600080fd5b50516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169190613ca5565b61222d337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156121d057600080fd5b505afa1580156121e4573d6000803e3d6000fd5b505050506040513d60208110156121fa57600080fd5b50516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169190613ca5565b604051339081904780156108fc02916000818181858888f1935050505015801561225b573d6000803e3d6000fd5b5050565b6008546001600160a01b031633146122a7576040805162461bcd60e51b815260206004808301919091526024820152637465616d60e01b604482015290519081900360640190fd5b60008160020b136122f2576040805162461bcd60e51b815260206004820152601060248201526f36b0bc2a3bb0b82232bb34b0ba34b7b760811b604482015290519081900360640190fd5b6013805460029290920b62ffffff16600160381b0269ffffff0000000000000019909216919091179055565b60135460020b81565b6002600654141561236d576040805162461bcd60e51b815260206004820152601f6024820152600080516020615c07833981519152604482015290519081900360640190fd5b60026006556008546001600160a01b031633146123ba576040805162461bcd60e51b815260206004808301919091526024820152637465616d60e01b604482015290519081900360640190fd5b505060016006555050565b6002600654141561240b576040805162461bcd60e51b815260206004820152601f6024820152600080516020615c07833981519152604482015290519081900360640190fd5b60026006556007546001600160a01b0316331461245c576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b600b556001600655565b600061250d600f547f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156124db57600080fd5b505afa1580156124ef573d6000803e3d6000fd5b505050506040513d602081101561250557600080fd5b505190613c48565b905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b600d5481565b6007546001600160a01b031681565b6007546001600160a01b03163314612597576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b6125ab6001600160a01b0383168285613ca5565b6040805160208101859052818152600f818301526e2bb4ba34323930bb9022b93199181d60891b60608201529051600080516020615db18339815191529181900360800190a1505050565b6013546040805160028082526060820183526000936301000000900463ffffffff169284929190602083019080368337019050509050818160008151811061263a57fe5b602002602001019063ffffffff16908163ffffffff168152505060008160018151811061266357fe5b63ffffffff90921660209283029190910182015260405163883bdbfd60e01b8152600481018281528351602483015283516000936001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169363883bdbfd938793909283926044019185820191028083838b5b838110156126f45781810151838201526020016126dc565b505050509050019250505060006040518083038186803b15801561271757600080fd5b505afa15801561272b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604090815281101561275457600080fd5b8101908080516040519392919084600160201b82111561277357600080fd5b90830190602082018581111561278857600080fd5b82518660208202830111600160201b821117156127a457600080fd5b82525081516020918201928201910280838360005b838110156127d15781810151838201526020016127b9565b5050505090500160405260200180516040519392919084600160201b8211156127f957600080fd5b90830190602082018581111561280e57600080fd5b82518660208202830111600160201b8211171561282a57600080fd5b82525081516020918201928201910280838360005b8381101561285757818101518382015260200161283f565b505050509050016040525050505090508263ffffffff168160008151811061287b57fe5b60200260200101518260018151811061289057fe5b60200260200101510360060b816128a357fe5b05935050505090565b6008546001600160a01b031633146128f4576040805162461bcd60e51b815260206004808301919091526024820152637465616d60e01b604482015290519081900360640190fd5b60648160ff161115612935576040805162461bcd60e51b8152602060048201526005602482015264726174696f60d81b604482015290519081900360640190fd5b6013805460ff909216600160501b0260ff60501b19909216919091179055565b6002600654141561299b576040805162461bcd60e51b815260206004820152601f6024820152600080516020615c07833981519152604482015290519081900360640190fd5b60026006556008546001600160a01b031633146129e8576040805162461bcd60e51b815260206004808301919091526024820152637465616d60e01b604482015290519081900360640190fd5b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633850c7bd6040518163ffffffff1660e01b815260040160e06040518083038186803b158015612a4357600080fd5b505afa158015612a57573d6000803e3d6000fd5b505050506040513d60e0811015612a6d57600080fd5b50602001519050612a7f838383614346565b600083830390507f000000000000000000000000000000000000000000000000000000000000000081620d89e719010160020b8260020b13612af7576040805162461bcd60e51b815260206004820152600c60248201526b7469636b20746f6f206c6f7760a01b604482015290519081900360640190fd5b7f000000000000000000000000000000000000000000000000000000000000000081620d89e719600003030360020b8260020b12612b6c576040805162461bcd60e51b815260206004820152600d60248201526c0e8d2c6d640e8dede40d0d2ced609b1b604482015290519081900360640190fd5b6000612b766125f6565b905060008160020b8460020b13612b8f57838203612b93565b8184035b9050601360079054906101000a900460020b60020b8160020b1315612beb576040805162461bcd60e51b81526020600482015260096024820152683232bb34b0ba34b7b760b91b604482015290519081900360640190fd5b612bf36145b6565b601354600090612c16906064906115ef90600160501b900460ff166114e4612cee565b90506000612c3d60646115ef6013600a9054906101000a900460ff1660ff166114e4612466565b9050612c4b888884846145d7565b6000612c55612cee565b90506000612c61612466565b90506000612c6f838361474d565b905080612cbb576040805162461bcd60e51b81526020600482015260156024820152741b195b991a5b99c81cdd5c1c1b1e4819985a5b1959605a1b604482015290519081900360640190fd5b50504260125550506013805460029690960b62ffffff1662ffffff19909616959095179094555050600160065550505050565b600061250d600e547f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156124db57600080fd5b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663128acb08308688873360405160200180826001600160a01b031681526020019150506040516020818303038152906040526040518663ffffffff1660e01b815260040180866001600160a01b031681526020018515158152602001848152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015612e3a578181015183820152602001612e22565b50505050905090810190601f168015612e675780820380516001836020036101000a031916815260200191505b5096505050505050506040805180830381600087803b158015612e8957600080fd5b505af1158015612e9d573d6000803e3d6000fd5b505050506040513d6040811015612eb357600080fd5b50805160209091015190969095509350505050565b4790565b6001600160a01b0381166000908152602081905260409020545b919050565b60008083905060008390506000816001600160a01b031663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015612f3357600080fd5b505af1158015612f47573d6000803e3d6000fd5b505050506040513d6020811015612f5d57600080fd5b505160408051602081018390528181526023918101829052919250600080516020615db183398151915291839181906060820190615c8f82396040019250505060405180910390a16000826001600160a01b031663ae9d70b06040518163ffffffff1660e01b8152600401602060405180830381600087803b158015612fe257600080fd5b505af1158015612ff6573d6000803e3d6000fd5b505050506040513d602081101561300c57600080fd5b5051604080516020818101849052828252818301527f537570706c7920526174653a20287363616c656420757020627920316531382960608201529051919250600080516020615db1833981519152919081900360800190a1836001600160a01b0316631249c58b6203d090476040518363ffffffff1660e01b81526004016000604051808303818589803b1580156130a457600080fd5b5088f11580156130b8573d6000803e3d6000fd5b5060019c9b505050505050505050505050565b6008546001600160a01b031681565b600c546301000000900460020b81565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156112475780601f1061121c57610100808354040283529160200191611247565b600c5460009081908190819061316f90600281810b9163010000009004900b613230565b9150915060008061317e611b4c565b91509150613192826116f8866116f8612cee565b95506131a4816116f8856116f8612466565b9450505050509091565b600f5481565b60006112d36131c1613cf7565b8461141085604051806060016040528060258152602001615dfb60259139600160006131eb613cf7565b6001600160a01b03908116825260208083019390935260409182016000908120918d16815292529020549190613f42565b60006112d3613229613cf7565b8484613de7565b600080600080600061324287876148b5565b9450945050509250613255878785614984565b600d54919650945060009061326e90620f424090613c48565b905061329461328d620f42406115ef6001600160801b03871685613fd9565b87906142ec565b95506132ba6132b3620f42406115ef6001600160801b03861685613fd9565b86906142ec565b9450505050509250929050565b60105481565b6007546001600160a01b03163314613319576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b600980546001600160a01b0319166001600160a01b0392909216919091179055565b6000816001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561338a57600080fd5b505afa15801561339e573d6000803e3d6000fd5b505050506040513d60208110156133b457600080fd5b505192915050565b60115481565b6007546001600160a01b0316331461340e576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b604051339082156108fc029083906000818181858888f1935050505015801561343b573d6000803e3d6000fd5b5060408051602081018390528181526017818301527f5769746864726177457468206d73672e73656e6465723a00000000000000000060608201529051600080516020615db18339815191529181900360800190a150565b6008546001600160a01b031633146134db576040805162461bcd60e51b815260206004808301919091526024820152637465616d60e01b604482015290519081900360640190fd5b60008163ffffffff1611613525576040805162461bcd60e51b815260206004820152600c60248201526b3a3bb0b8223ab930ba34b7b760a11b604482015290519081900360640190fd5b6013805463ffffffff90921663010000000266ffffffff00000019909216919091179055565b60008181600185151514156135d357816001600160a01b031663db006a75876040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b1580156135a057600080fd5b505af11580156135b4573d6000803e3d6000fd5b505050506040513d60208110156135ca57600080fd5b50519050613648565b816001600160a01b031663852a12e3876040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b15801561361957600080fd5b505af115801561362d573d6000803e3d6000fd5b505050506040513d602081101561364357600080fd5b505190505b600080516020615db183398151915281604051808060200183815260200182810382526024815260200180615c6b602491396040019250505060405180910390a150600195945050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60006136c583836148b5565b50929695505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461374e57600080fd5b8315613788576137886001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163386613ca5565b82156137c2576137c26001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163385613ca5565b50505050565b600c5460020b81565b6007546001600160a01b0316331461381d576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b6138256145b6565b565b6007546001600160a01b03163314613873576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316836001600160a01b0316141580156138e757507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316836001600160a01b031614155b613920576040805162461bcd60e51b81526020600482015260056024820152643a37b5b2b760d91b604482015290519081900360640190fd5b61107e6001600160a01b0384168284613ca5565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6000806000600260065414156139aa576040805162461bcd60e51b815260206004820152601f6024820152600080516020615c07833981519152604482015290519081900360640190fd5b60026006553380158015906139c857506001600160a01b0381163014155b6139fe576040805162461bcd60e51b8152602060048201526002602482015261746f60f01b604482015290519081900360640190fd5b600c54613a1990600281810b9163010000009004900b614a39565b613a238686614b0a565b919550935091508215613a6557613a656001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016333086614b38565b8115613aa057613aa06001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016333085614b38565b613aaa8185614b92565b604080518581526020810185905280820184905290516001600160a01b0383169133917f4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f69181900360600190a3600b54613b0261132b565b1115613b3b576040805162461bcd60e51b815260206004820152600360248201526204341560ec1b604482015290519081900360640190fd5b5060016006819055509250925092565b601354600160381b900460020b81565b600e5481565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614613bce576040805162461bcd60e51b815260206004820152600d60248201526c1cd95b99195c880f481c1bdbdb609a1b604482015290519081900360640190fd5b6000841315613c0b57613c0b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163386613ca5565b60008313156137c2576137c26001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163385613ca5565b600082821115613c9f576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b17905261107e908490614c82565b3390565b6001600160a01b038316613d405760405162461bcd60e51b8152600401808060200182810382526024815260200180615d8d6024913960400191505060405180910390fd5b6001600160a01b038216613d855760405162461bcd60e51b8152600401808060200182810382526022815260200180615c496022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b038316613e2c5760405162461bcd60e51b8152600401808060200182810382526025815260200180615d686025913960400191505060405180910390fd5b6001600160a01b038216613e715760405162461bcd60e51b8152600401808060200182810382526023815260200180615be46023913960400191505060405180910390fd5b613e7c83838361107e565b613eb981604051806060016040528060268152602001615cb2602691396001600160a01b0386166000908152602081905260409020549190613f42565b6001600160a01b038085166000908152602081905260408082209390935590841681522054613ee890826142ec565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008184841115613fd15760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015613f96578181015183820152602001613f7e565b50505050905090810190601f168015613fc35780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600082613fe8575060006112d7565b82820282848281613ff557fe5b04146114195760405162461bcd60e51b8152600401808060200182810382526021815260200180615cfe6021913960400191505060405180910390fd5b6000808211614088576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b81838161409157fe5b049392505050565b6001600160a01b0382166140de5760405162461bcd60e51b8152600401808060200182810382526021815260200180615d476021913960400191505060405180910390fd5b6140ea8260008361107e565b61412781604051806060016040528060228152602001615c27602291396001600160a01b0385166000908152602081905260409020549190613f42565b6001600160a01b03831660009081526020819052604090205560025461414d9082613c48565b6002556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b60008060006141a487876148b5565b50505050905060006141cc856115ef88856001600160801b0316613fd990919063ffffffff16565b90508015614230576000806000806141ed8c8c6141e888614d33565b614d4a565b9296509094509250905061420f6142088a6115ef858e613fd9565b85906142ec565b97506142296142228a6115ef848e613fd9565b84906142ec565b9650505050505b505094509492505050565b600080600080614249611083565b9092509050600061425e866115ef858a613fd9565b90506000614270876115ef858b613fd9565b905061427c8282614fd5565b604080516020810184905280820183905260608082526012908201527102fb13ab9372632b73234b733a9b430b932960751b608082015290517fb626a5b8d37253edd986272eff7a2541a1538cad44624e413d51713f0b2a414d9181900360a00190a19097909650945050505050565b600082820183811015611419576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b8160020b8360020b12614391576040805162461bcd60e51b815260206004820152600e60248201526d3637bbb2b9101e102fbab83832b960911b604482015290519081900360640190fd5b8060020b8360020b126143db576040805162461bcd60e51b815260206004820152600d60248201526c6c6f776572203e205f7469636b60981b604482015290519081900360640190fd5b8060020b8260020b13614426576040805162461bcd60e51b815260206004820152600e60248201526d5f7570706572203e205f7469636b60901b604482015290519081900360640190fd5b620d89e719600284900b1215614473576040805162461bcd60e51b815260206004820152600d60248201526c4c6f77657220746f6f206c6f7760981b604482015290519081900360640190fd5b620d89e8600283900b13156144c0576040805162461bcd60e51b815260206004820152600e60248201526d0aae0e0cae440e8dede40d0d2ced60931b604482015290519081900360640190fd5b7f000000000000000000000000000000000000000000000000000000000000000060020b8360020b816144ef57fe5b0760020b1561453b576040805162461bcd60e51b81526020600482015260136024820152724c6f7765722025207469636b53706163696e6760681b604482015290519081900360640190fd5b7f000000000000000000000000000000000000000000000000000000000000000060020b8260020b8161456a57fe5b0760020b1561107e576040805162461bcd60e51b815260206004820152601360248201527255707065722025207469636b53706163696e6760681b604482015290519081900360640190fd5b6000806145c1611083565b915091506145cf8282614fd5565b61225b6150f9565b60006145e58585858561513c565b60408051633c8a7d8d60e01b8152306004820152600288810b602483015287900b60448201526001600160801b038316606482015260a06084820152600060a4820181905282519394506001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001693633c8a7d8d9360c4808501949193918390030190829087803b15801561467f57600080fd5b505af1158015614693573d6000803e3d6000fd5b505050506040513d60408110156146a957600080fd5b50600090506146b6612cee565b905060006146c2612466565b604080516001600160801b03861681526020810185905280820183905290519192507ff290cad35129b75da3cfb1a6dfcb58954eabf3dd09731a5ea198c49a4ab2c032919081900360600190a15050600c8054600296870b62ffffff90811662ffffff199790980b1663010000000265ffffff00000019909116179490941694909417909255505050565b600a546000907f0000000000000000000000000000000000000000000000000000000000000000907f0000000000000000000000000000000000000000000000000000000000000000906001600160a01b03908116907f0000000000000000000000000000000000000000000000000000000000000000907f0000000000000000000000000000000000000000000000000000000000000000908516831415614835576147f9886151ed565b6148237f000000000000000000000000000000000000000000000000000000000000000082612eeb565b5061482f848289611cf8565b506148a7565b826001600160a01b0316846001600160a01b0316141561488e57614858876151ed565b6148827f000000000000000000000000000000000000000000000000000000000000000083612eeb565b5061482f85838a611cf8565b61489985838a611cf8565b506148a5848289611cf8565b505b506001979650505050505050565b6000806000806000806148c93089896152a7565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663514ea4bf826040518263ffffffff1660e01b81526004018082815260200191505060a06040518083038186803b15801561492f57600080fd5b505afa158015614943573d6000803e3d6000fd5b505050506040513d60a081101561495957600080fd5b508051602082015160408301516060840151608090940151929c919b50995091975095509350505050565b60008060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633850c7bd6040518163ffffffff1660e01b815260040160e06040518083038186803b1580156149e257600080fd5b505afa1580156149f6573d6000803e3d6000fd5b505050506040513d60e0811015614a0c57600080fd5b50519050614a2c81614a1d886152fd565b614a26886152fd565b8761562e565b9250925050935093915050565b6000614a4583836148b5565b5050505090506000816001600160801b0316111561107e577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a34123a7848460006040518463ffffffff1660e01b8152600401808460020b81526020018360020b815260200182815260200193505050506040805180830381600087803b158015614ad957600080fd5b505af1158015614aed573d6000803e3d6000fd5b505050506040513d6040811015614b0357600080fd5b5050505050565b600080600080614b18611439565b8693508592509050614b2e826116f88584613fd9565b9350509250925092565b604080516001600160a01b0380861660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b1790526137c2908590614c82565b6001600160a01b038216614bed576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b614bf96000838361107e565b600254614c0690826142ec565b6002556001600160a01b038216600090815260208190526040902054614c2c90826142ec565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b6000614cd7826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166156ca9092919063ffffffff16565b80519091501561107e57808060200190516020811015614cf657600080fd5b505161107e5760405162461bcd60e51b815260040180806020018281038252602a815260200180615dd1602a913960400191505060405180910390fd5b60006001600160801b03821115614d4657fe5b5090565b60008080806001600160801b03851615614e16576040805163a34123a760e01b8152600289810b600483015288900b60248201526001600160801b038716604482015281516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169263a34123a792606480820193918290030181600087803b158015614ddd57600080fd5b505af1158015614df1573d6000803e3d6000fd5b505050506040513d6040811015614e0757600080fd5b50805160209091015190945092505b604080516309e3d67b60e31b8152306004820152600289810b602483015288900b60448201526001600160801b03606482018190526084820152815160009283926001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001692634f1eb3d89260a48084019391929182900301818787803b158015614ea657600080fd5b505af1158015614eba573d6000803e3d6000fd5b505050506040513d6040811015614ed057600080fd5b5080516020909101516001600160801b039182169350169050614ef38287613c48565b9350614eff8186613c48565b60108054860190556011805482019055600d5490935060009081908015614f8057614f31620f42406115ef8984613fd9565b9250614f44620f42406115ef8884613fd9565b9150614f508784613c48565b9650614f5c8683613c48565b600e54909650614f6c90846142ec565b600e55600f54614f7c90836142ec565b600f555b60408051888152602081018890528082018590526060810184905290517f1ac56d7e866e3f5ea9aa92aa11758ead39a0a5f013f3fefb0f47cb9d008edd279181900360800190a1505050505093509350935093565b600a547f0000000000000000000000000000000000000000000000000000000000000000907f0000000000000000000000000000000000000000000000000000000000000000906001600160a01b03908116907f0000000000000000000000000000000000000000000000000000000000000000907f000000000000000000000000000000000000000000000000000000000000000090600190861684141561509d5761508388828561354b565b5061508c6156e1565b61509787828461354b565b506150ef565b836001600160a01b0316856001600160a01b031614156150d6576150c287828461354b565b506150cb6156e1565b61509788828561354b565b6150e187828461354b565b506150ed88828561354b565b505b5050505050505050565b600c5460009061511790600281810b9163010000009004900b6148b5565b5050600c54929350614b0392600281810b93506301000000909104900b905083614d4a565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633850c7bd6040518163ffffffff1660e01b815260040160e06040518083038186803b15801561519857600080fd5b505afa1580156151ac573d6000803e3d6000fd5b505050506040513d60e08110156151c257600080fd5b505190506151e3816151d3886152fd565b6151dc886152fd565b878761574b565b9695505050505050565b801561525957600a5460408051632e1a7d4d60e01b81526004810184905290516001600160a01b0390921691632e1a7d4d9160248082019260009290919082900301818387803b15801561524057600080fd5b505af1158015615254573d6000803e3d6000fd5b505050505b60408051476020820152818152601581830152743ab73bb930b83832b21032ba341030b6b7bab73a1d60591b60608201529051600080516020615db18339815191529181900360800190a150565b6040805160609490941b6bffffffffffffffffffffffff1916602080860191909152600293840b60e890811b60348701529290930b90911b60378401528051808403601a018152603a9093019052815191012090565b60008060008360020b12615314578260020b61531c565b8260020b6000035b9050620d89e881111561535a576040805162461bcd60e51b81526020600482015260016024820152601560fa1b604482015290519081900360640190fd5b60006001821661536e57600160801b615380565b6ffffcb933bd6fad37aa2d162d1a5940015b70ffffffffffffffffffffffffffffffffff16905060028216156153b4576ffff97272373d413259a46990580e213a0260801c5b60048216156153d3576ffff2e50f5f656932ef12357cf3c7fdcc0260801c5b60088216156153f2576fffe5caca7e10e4e61c3624eaa0941cd00260801c5b6010821615615411576fffcb9843d60f6159c9db58835c9266440260801c5b6020821615615430576fff973b41fa98c081472e6896dfb254c00260801c5b604082161561544f576fff2ea16466c96a3843ec78b326b528610260801c5b608082161561546e576ffe5dee046a99a2a811c461f1969c30530260801c5b61010082161561548e576ffcbe86c7900a88aedcffc83b479aa3a40260801c5b6102008216156154ae576ff987a7253ac413176f2b074cf7815e540260801c5b6104008216156154ce576ff3392b0822b70005940c7a398e4b70f30260801c5b6108008216156154ee576fe7159475a2c29b7443b29c7fa6e889d90260801c5b61100082161561550e576fd097f3bdfd2022b8845ad8f792aa58250260801c5b61200082161561552e576fa9f746462d870fdf8a65dc1f90e061e50260801c5b61400082161561554e576f70d869a156d2a1b890bb3df62baf32f70260801c5b61800082161561556e576f31be135f97d08fd981231505542fcfa60260801c5b6201000082161561558f576f09aa508b5b7a84e1c677de54f3e99bc90260801c5b620200008216156155af576e5d6af8dedb81196699c329225ee6040260801c5b620400008216156155ce576d2216e584f5fa1ea926041bedfe980260801c5b620800008216156155eb576b048a170391f7dc42444e8fa20260801c5b60008460020b131561560657806000198161560257fe5b0490505b600160201b81061561561957600161561c565b60005b60ff16602082901c0192505050919050565b600080836001600160a01b0316856001600160a01b0316111561564f579293925b846001600160a01b0316866001600160a01b03161161567a5761567385858561580f565b91506156c1565b836001600160a01b0316866001600160a01b031610156156b35761569f86858561580f565b91506156ac858785615878565b90506156c1565b6156be858585615878565b90505b94509492505050565b60606156d984846000856158bb565b949350505050565b471561382557600a60009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0476040518263ffffffff1660e01b81526004016000604051808303818588803b15801561573757600080fd5b505af1158015614b03573d6000803e3d6000fd5b6000836001600160a01b0316856001600160a01b0316111561576b579293925b846001600160a01b0316866001600160a01b0316116157965761578f858585615a16565b9050615806565b836001600160a01b0316866001600160a01b031610156157f85760006157bd878686615a16565b905060006157cc878986615a79565b9050806001600160801b0316826001600160801b0316106157ed57806157ef565b815b92505050615806565b615803858584615a79565b90505b95945050505050565b6000826001600160a01b0316846001600160a01b0316111561582f579192915b836001600160a01b0316615868606060ff16846001600160801b0316901b8686036001600160a01b0316866001600160a01b0316615ab2565b8161586f57fe5b04949350505050565b6000826001600160a01b0316846001600160a01b03161115615898579192915b6156d9826001600160801b03168585036001600160a01b0316600160601b615ab2565b6060824710156158fc5760405162461bcd60e51b8152600401808060200182810382526026815260200180615cd86026913960400191505060405180910390fd5b61590585615b61565b615956576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b600080866001600160a01b031685876040518082805190602001908083835b602083106159945780518252601f199092019160209182019101615975565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d80600081146159f6576040519150601f19603f3d011682016040523d82523d6000602084013e6159fb565b606091505b5091509150615a0b828286615b67565b979650505050505050565b6000826001600160a01b0316846001600160a01b03161115615a36579192915b6000615a59856001600160a01b0316856001600160a01b0316600160601b615ab2565b9050615806615a7484838888036001600160a01b0316615ab2565b615bcd565b6000826001600160a01b0316846001600160a01b03161115615a99579192915b6156d9615a7483600160601b8787036001600160a01b03165b6000808060001985870986860292508281109083900303905080615ae85760008411615add57600080fd5b508290049050611419565b808411615af457600080fd5b6000848688096000868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b3b151590565b60608315615b76575081611419565b825115615b865782518084602001fd5b60405162461bcd60e51b8152602060048201818152845160248401528451859391928392604401919085019080838360008315613f96578181015183820152602001613f7e565b806001600160801b0381168114612ee657600080fdfe45524332303a207472616e7366657220746f20746865207a65726f20616464726573735265656e7472616e637947756172643a207265656e7472616e742063616c6c0045524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332303a20617070726f766520746f20746865207a65726f206164647265737349662074686973206973206e6f7420302c2074686572652077617320616e206572726f7245786368616e6765205261746520287363616c65642075702062792031653138293a2045524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e6365416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f7745524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f20616464726573738d1cced004452bd270777a8c670f9f7e7c4fdde56f2db331fe289d39dc2624ad5361666545524332303a204552433230206f7065726174696f6e20646964206e6f74207375636365656445524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa26469706673582212208a5cc0dab829194a879c31f7bad7404d130e21ea0d2f9d67c4eb9e8e40504eda64736f6c63430007060033"

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, _pool common.Address, _weth common.Address, _cToken0 common.Address, _cToken1 common.Address, _cEth common.Address, _protocolFeeRate *big.Int, _maxTotalSupply *big.Int, _maxTwapDeviation *big.Int, _twapDuration uint32, _uniPortionRate uint8) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ApiBin), backend, _pool, _weth, _cToken0, _cToken1, _cEth, _protocolFeeRate, _maxTotalSupply, _maxTwapDeviation, _twapDuration, _uniPortionRate)
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

// CEther is a free data retrieval call binding the contract method 0x53257e00.
//
// Solidity: function CEther() view returns(address)
func (_Api *ApiCaller) CEther(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "CEther")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CEther is a free data retrieval call binding the contract method 0x53257e00.
//
// Solidity: function CEther() view returns(address)
func (_Api *ApiSession) CEther() (common.Address, error) {
	return _Api.Contract.CEther(&_Api.CallOpts)
}

// CEther is a free data retrieval call binding the contract method 0x53257e00.
//
// Solidity: function CEther() view returns(address)
func (_Api *ApiCallerSession) CEther() (common.Address, error) {
	return _Api.Contract.CEther(&_Api.CallOpts)
}

// CToken0 is a free data retrieval call binding the contract method 0xcbcf94aa.
//
// Solidity: function CToken0() view returns(address)
func (_Api *ApiCaller) CToken0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "CToken0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CToken0 is a free data retrieval call binding the contract method 0xcbcf94aa.
//
// Solidity: function CToken0() view returns(address)
func (_Api *ApiSession) CToken0() (common.Address, error) {
	return _Api.Contract.CToken0(&_Api.CallOpts)
}

// CToken0 is a free data retrieval call binding the contract method 0xcbcf94aa.
//
// Solidity: function CToken0() view returns(address)
func (_Api *ApiCallerSession) CToken0() (common.Address, error) {
	return _Api.Contract.CToken0(&_Api.CallOpts)
}

// CToken1 is a free data retrieval call binding the contract method 0x32380717.
//
// Solidity: function CToken1() view returns(address)
func (_Api *ApiCaller) CToken1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "CToken1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CToken1 is a free data retrieval call binding the contract method 0x32380717.
//
// Solidity: function CToken1() view returns(address)
func (_Api *ApiSession) CToken1() (common.Address, error) {
	return _Api.Contract.CToken1(&_Api.CallOpts)
}

// CToken1 is a free data retrieval call binding the contract method 0x32380717.
//
// Solidity: function CToken1() view returns(address)
func (_Api *ApiCallerSession) CToken1() (common.Address, error) {
	return _Api.Contract.CToken1(&_Api.CallOpts)
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

// GetBalance0 is a free data retrieval call binding the contract method 0x629d9405.
//
// Solidity: function getBalance0() view returns(uint256)
func (_Api *ApiCaller) GetBalance0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getBalance0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance0 is a free data retrieval call binding the contract method 0x629d9405.
//
// Solidity: function getBalance0() view returns(uint256)
func (_Api *ApiSession) GetBalance0() (*big.Int, error) {
	return _Api.Contract.GetBalance0(&_Api.CallOpts)
}

// GetBalance0 is a free data retrieval call binding the contract method 0x629d9405.
//
// Solidity: function getBalance0() view returns(uint256)
func (_Api *ApiCallerSession) GetBalance0() (*big.Int, error) {
	return _Api.Contract.GetBalance0(&_Api.CallOpts)
}

// GetBalance1 is a free data retrieval call binding the contract method 0x41aec538.
//
// Solidity: function getBalance1() view returns(uint256)
func (_Api *ApiCaller) GetBalance1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getBalance1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance1 is a free data retrieval call binding the contract method 0x41aec538.
//
// Solidity: function getBalance1() view returns(uint256)
func (_Api *ApiSession) GetBalance1() (*big.Int, error) {
	return _Api.Contract.GetBalance1(&_Api.CallOpts)
}

// GetBalance1 is a free data retrieval call binding the contract method 0x41aec538.
//
// Solidity: function getBalance1() view returns(uint256)
func (_Api *ApiCallerSession) GetBalance1() (*big.Int, error) {
	return _Api.Contract.GetBalance1(&_Api.CallOpts)
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

// GetCTokenBalance is a free data retrieval call binding the contract method 0xbb47dc1c.
//
// Solidity: function getCTokenBalance(address _erc20Contract) view returns(uint256)
func (_Api *ApiCaller) GetCTokenBalance(opts *bind.CallOpts, _erc20Contract common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getCTokenBalance", _erc20Contract)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCTokenBalance is a free data retrieval call binding the contract method 0xbb47dc1c.
//
// Solidity: function getCTokenBalance(address _erc20Contract) view returns(uint256)
func (_Api *ApiSession) GetCTokenBalance(_erc20Contract common.Address) (*big.Int, error) {
	return _Api.Contract.GetCTokenBalance(&_Api.CallOpts, _erc20Contract)
}

// GetCTokenBalance is a free data retrieval call binding the contract method 0xbb47dc1c.
//
// Solidity: function getCTokenBalance(address _erc20Contract) view returns(uint256)
func (_Api *ApiCallerSession) GetCTokenBalance(_erc20Contract common.Address) (*big.Int, error) {
	return _Api.Contract.GetCTokenBalance(&_Api.CallOpts, _erc20Contract)
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

// GetLendingAmounts is a free data retrieval call binding the contract method 0x2e4b8f60.
//
// Solidity: function getLendingAmounts() view returns(uint256, uint256)
func (_Api *ApiCaller) GetLendingAmounts(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getLendingAmounts")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetLendingAmounts is a free data retrieval call binding the contract method 0x2e4b8f60.
//
// Solidity: function getLendingAmounts() view returns(uint256, uint256)
func (_Api *ApiSession) GetLendingAmounts() (*big.Int, *big.Int, error) {
	return _Api.Contract.GetLendingAmounts(&_Api.CallOpts)
}

// GetLendingAmounts is a free data retrieval call binding the contract method 0x2e4b8f60.
//
// Solidity: function getLendingAmounts() view returns(uint256, uint256)
func (_Api *ApiCallerSession) GetLendingAmounts() (*big.Int, *big.Int, error) {
	return _Api.Contract.GetLendingAmounts(&_Api.CallOpts)
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

// GetTVL is a free data retrieval call binding the contract method 0x97b3fcaa.
//
// Solidity: function getTVL() view returns(uint256 total0, uint256 total1)
func (_Api *ApiCaller) GetTVL(opts *bind.CallOpts) (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getTVL")

	outstruct := new(struct {
		Total0 *big.Int
		Total1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Total0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Total1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTVL is a free data retrieval call binding the contract method 0x97b3fcaa.
//
// Solidity: function getTVL() view returns(uint256 total0, uint256 total1)
func (_Api *ApiSession) GetTVL() (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	return _Api.Contract.GetTVL(&_Api.CallOpts)
}

// GetTVL is a free data retrieval call binding the contract method 0x97b3fcaa.
//
// Solidity: function getTVL() view returns(uint256 total0, uint256 total1)
func (_Api *ApiCallerSession) GetTVL() (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	return _Api.Contract.GetTVL(&_Api.CallOpts)
}

// GetTwap is a free data retrieval call binding the contract method 0x5d752a9a.
//
// Solidity: function getTwap() view returns(int24)
func (_Api *ApiCaller) GetTwap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getTwap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTwap is a free data retrieval call binding the contract method 0x5d752a9a.
//
// Solidity: function getTwap() view returns(int24)
func (_Api *ApiSession) GetTwap() (*big.Int, error) {
	return _Api.Contract.GetTwap(&_Api.CallOpts)
}

// GetTwap is a free data retrieval call binding the contract method 0x5d752a9a.
//
// Solidity: function getTwap() view returns(int24)
func (_Api *ApiCallerSession) GetTwap() (*big.Int, error) {
	return _Api.Contract.GetTwap(&_Api.CallOpts)
}

// GetUniswapPrice is a free data retrieval call binding the contract method 0x2b6e1923.
//
// Solidity: function getUniswapPrice() view returns(uint256 price)
func (_Api *ApiCaller) GetUniswapPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getUniswapPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUniswapPrice is a free data retrieval call binding the contract method 0x2b6e1923.
//
// Solidity: function getUniswapPrice() view returns(uint256 price)
func (_Api *ApiSession) GetUniswapPrice() (*big.Int, error) {
	return _Api.Contract.GetUniswapPrice(&_Api.CallOpts)
}

// GetUniswapPrice is a free data retrieval call binding the contract method 0x2b6e1923.
//
// Solidity: function getUniswapPrice() view returns(uint256 price)
func (_Api *ApiCallerSession) GetUniswapPrice() (*big.Int, error) {
	return _Api.Contract.GetUniswapPrice(&_Api.CallOpts)
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

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Api *ApiCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Api *ApiSession) Pool() (common.Address, error) {
	return _Api.Contract.Pool(&_Api.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Api *ApiCallerSession) Pool() (common.Address, error) {
	return _Api.Contract.Pool(&_Api.CallOpts)
}

// ProtocolFeeRate is a free data retrieval call binding the contract method 0x58f85880.
//
// Solidity: function protocolFeeRate() view returns(uint256)
func (_Api *ApiCaller) ProtocolFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "protocolFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeRate is a free data retrieval call binding the contract method 0x58f85880.
//
// Solidity: function protocolFeeRate() view returns(uint256)
func (_Api *ApiSession) ProtocolFeeRate() (*big.Int, error) {
	return _Api.Contract.ProtocolFeeRate(&_Api.CallOpts)
}

// ProtocolFeeRate is a free data retrieval call binding the contract method 0x58f85880.
//
// Solidity: function protocolFeeRate() view returns(uint256)
func (_Api *ApiCallerSession) ProtocolFeeRate() (*big.Int, error) {
	return _Api.Contract.ProtocolFeeRate(&_Api.CallOpts)
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

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_Api *ApiCaller) TickSpacing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "tickSpacing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_Api *ApiSession) TickSpacing() (*big.Int, error) {
	return _Api.Contract.TickSpacing(&_Api.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_Api *ApiCallerSession) TickSpacing() (*big.Int, error) {
	return _Api.Contract.TickSpacing(&_Api.CallOpts)
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

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 amountToken0, uint256 amountToken1) returns(uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiTransactor) Deposit(opts *bind.TransactOpts, amountToken0 *big.Int, amountToken1 *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "deposit", amountToken0, amountToken1)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 amountToken0, uint256 amountToken1) returns(uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiSession) Deposit(amountToken0 *big.Int, amountToken1 *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Deposit(&_Api.TransactOpts, amountToken0, amountToken1)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 amountToken0, uint256 amountToken1) returns(uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiTransactorSession) Deposit(amountToken0 *big.Int, amountToken1 *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Deposit(&_Api.TransactOpts, amountToken0, amountToken1)
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

// RedeemCErc20Tokens is a paid mutator transaction binding the contract method 0xd0ace48e.
//
// Solidity: function redeemCErc20Tokens(uint256 amount, bool redeemType, address _cErc20Contract) returns(bool)
func (_Api *ApiTransactor) RedeemCErc20Tokens(opts *bind.TransactOpts, amount *big.Int, redeemType bool, _cErc20Contract common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "redeemCErc20Tokens", amount, redeemType, _cErc20Contract)
}

// RedeemCErc20Tokens is a paid mutator transaction binding the contract method 0xd0ace48e.
//
// Solidity: function redeemCErc20Tokens(uint256 amount, bool redeemType, address _cErc20Contract) returns(bool)
func (_Api *ApiSession) RedeemCErc20Tokens(amount *big.Int, redeemType bool, _cErc20Contract common.Address) (*types.Transaction, error) {
	return _Api.Contract.RedeemCErc20Tokens(&_Api.TransactOpts, amount, redeemType, _cErc20Contract)
}

// RedeemCErc20Tokens is a paid mutator transaction binding the contract method 0xd0ace48e.
//
// Solidity: function redeemCErc20Tokens(uint256 amount, bool redeemType, address _cErc20Contract) returns(bool)
func (_Api *ApiTransactorSession) RedeemCErc20Tokens(amount *big.Int, redeemType bool, _cErc20Contract common.Address) (*types.Transaction, error) {
	return _Api.Contract.RedeemCErc20Tokens(&_Api.TransactOpts, amount, redeemType, _cErc20Contract)
}

// RedeemCEth is a paid mutator transaction binding the contract method 0xc6fb5014.
//
// Solidity: function redeemCEth(uint256 amount, bool redeemType, address _cEtherContract) returns(bool)
func (_Api *ApiTransactor) RedeemCEth(opts *bind.TransactOpts, amount *big.Int, redeemType bool, _cEtherContract common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "redeemCEth", amount, redeemType, _cEtherContract)
}

// RedeemCEth is a paid mutator transaction binding the contract method 0xc6fb5014.
//
// Solidity: function redeemCEth(uint256 amount, bool redeemType, address _cEtherContract) returns(bool)
func (_Api *ApiSession) RedeemCEth(amount *big.Int, redeemType bool, _cEtherContract common.Address) (*types.Transaction, error) {
	return _Api.Contract.RedeemCEth(&_Api.TransactOpts, amount, redeemType, _cEtherContract)
}

// RedeemCEth is a paid mutator transaction binding the contract method 0xc6fb5014.
//
// Solidity: function redeemCEth(uint256 amount, bool redeemType, address _cEtherContract) returns(bool)
func (_Api *ApiTransactorSession) RedeemCEth(amount *big.Int, redeemType bool, _cEtherContract common.Address) (*types.Transaction, error) {
	return _Api.Contract.RedeemCEth(&_Api.TransactOpts, amount, redeemType, _cEtherContract)
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

// Strategy0 is a paid mutator transaction binding the contract method 0x3e5aae66.
//
// Solidity: function strategy0(int24 newLow, int24 newHigh, int256 swapAmount, bool zeroForOne) returns()
func (_Api *ApiTransactor) Strategy0(opts *bind.TransactOpts, newLow *big.Int, newHigh *big.Int, swapAmount *big.Int, zeroForOne bool) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "strategy0", newLow, newHigh, swapAmount, zeroForOne)
}

// Strategy0 is a paid mutator transaction binding the contract method 0x3e5aae66.
//
// Solidity: function strategy0(int24 newLow, int24 newHigh, int256 swapAmount, bool zeroForOne) returns()
func (_Api *ApiSession) Strategy0(newLow *big.Int, newHigh *big.Int, swapAmount *big.Int, zeroForOne bool) (*types.Transaction, error) {
	return _Api.Contract.Strategy0(&_Api.TransactOpts, newLow, newHigh, swapAmount, zeroForOne)
}

// Strategy0 is a paid mutator transaction binding the contract method 0x3e5aae66.
//
// Solidity: function strategy0(int24 newLow, int24 newHigh, int256 swapAmount, bool zeroForOne) returns()
func (_Api *ApiTransactorSession) Strategy0(newLow *big.Int, newHigh *big.Int, swapAmount *big.Int, zeroForOne bool) (*types.Transaction, error) {
	return _Api.Contract.Strategy0(&_Api.TransactOpts, newLow, newHigh, swapAmount, zeroForOne)
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

// Sweep is a paid mutator transaction binding the contract method 0xdc2c256f.
//
// Solidity: function sweep(address token, uint256 amount, address to) returns()
func (_Api *ApiTransactor) Sweep(opts *bind.TransactOpts, token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "sweep", token, amount, to)
}

// Sweep is a paid mutator transaction binding the contract method 0xdc2c256f.
//
// Solidity: function sweep(address token, uint256 amount, address to) returns()
func (_Api *ApiSession) Sweep(token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Api.Contract.Sweep(&_Api.TransactOpts, token, amount, to)
}

// Sweep is a paid mutator transaction binding the contract method 0xdc2c256f.
//
// Solidity: function sweep(address token, uint256 amount, address to) returns()
func (_Api *ApiTransactorSession) Sweep(token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Api.Contract.Sweep(&_Api.TransactOpts, token, amount, to)
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

// WhiteHacker is a paid mutator transaction binding the contract method 0x3bdd61b7.
//
// Solidity: function whiteHacker() returns()
func (_Api *ApiTransactor) WhiteHacker(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "whiteHacker")
}

// WhiteHacker is a paid mutator transaction binding the contract method 0x3bdd61b7.
//
// Solidity: function whiteHacker() returns()
func (_Api *ApiSession) WhiteHacker() (*types.Transaction, error) {
	return _Api.Contract.WhiteHacker(&_Api.TransactOpts)
}

// WhiteHacker is a paid mutator transaction binding the contract method 0x3bdd61b7.
//
// Solidity: function whiteHacker() returns()
func (_Api *ApiTransactorSession) WhiteHacker() (*types.Transaction, error) {
	return _Api.Contract.WhiteHacker(&_Api.TransactOpts)
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

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x5d29dab4.
//
// Solidity: function withdrawERC20(uint256 amount, address erc20, address to) returns()
func (_Api *ApiTransactor) WithdrawERC20(opts *bind.TransactOpts, amount *big.Int, erc20 common.Address, to common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "withdrawERC20", amount, erc20, to)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x5d29dab4.
//
// Solidity: function withdrawERC20(uint256 amount, address erc20, address to) returns()
func (_Api *ApiSession) WithdrawERC20(amount *big.Int, erc20 common.Address, to common.Address) (*types.Transaction, error) {
	return _Api.Contract.WithdrawERC20(&_Api.TransactOpts, amount, erc20, to)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x5d29dab4.
//
// Solidity: function withdrawERC20(uint256 amount, address erc20, address to) returns()
func (_Api *ApiTransactorSession) WithdrawERC20(amount *big.Int, erc20 common.Address, to common.Address) (*types.Transaction, error) {
	return _Api.Contract.WithdrawERC20(&_Api.TransactOpts, amount, erc20, to)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xc311d049.
//
// Solidity: function withdrawEth(uint256 amount) returns()
func (_Api *ApiTransactor) WithdrawEth(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "withdrawEth", amount)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xc311d049.
//
// Solidity: function withdrawEth(uint256 amount) returns()
func (_Api *ApiSession) WithdrawEth(amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.WithdrawEth(&_Api.TransactOpts, amount)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xc311d049.
//
// Solidity: function withdrawEth(uint256 amount) returns()
func (_Api *ApiTransactorSession) WithdrawEth(amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.WithdrawEth(&_Api.TransactOpts, amount)
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
	FeesToVault0    *big.Int
	FeesToVault1    *big.Int
	FeesToProtocol0 *big.Int
	FeesToProtocol1 *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCollectFees is a free log retrieval operation binding the contract event 0x1ac56d7e866e3f5ea9aa92aa11758ead39a0a5f013f3fefb0f47cb9d008edd27.
//
// Solidity: event CollectFees(uint256 feesToVault0, uint256 feesToVault1, uint256 feesToProtocol0, uint256 feesToProtocol1)
func (_Api *ApiFilterer) FilterCollectFees(opts *bind.FilterOpts) (*ApiCollectFeesIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "CollectFees")
	if err != nil {
		return nil, err
	}
	return &ApiCollectFeesIterator{contract: _Api.contract, event: "CollectFees", logs: logs, sub: sub}, nil
}

// WatchCollectFees is a free log subscription operation binding the contract event 0x1ac56d7e866e3f5ea9aa92aa11758ead39a0a5f013f3fefb0f47cb9d008edd27.
//
// Solidity: event CollectFees(uint256 feesToVault0, uint256 feesToVault1, uint256 feesToProtocol0, uint256 feesToProtocol1)
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

// ParseCollectFees is a log parse operation binding the contract event 0x1ac56d7e866e3f5ea9aa92aa11758ead39a0a5f013f3fefb0f47cb9d008edd27.
//
// Solidity: event CollectFees(uint256 feesToVault0, uint256 feesToVault1, uint256 feesToProtocol0, uint256 feesToProtocol1)
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
	Sender     common.Address
	To         common.Address
	Shares     *big.Int
	Amount0    *big.Int
	Amount1    *big.Int
	NameToken0 string
	NameToken1 string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9167547f76aeba7bf8001d943eb4573822255e96f6ee40278381b71fad6af5ad.
//
// Solidity: event Withdraw(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1, string nameToken0, string nameToken1)
func (_Api *ApiFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*ApiWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Withdraw", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ApiWithdrawIterator{contract: _Api.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9167547f76aeba7bf8001d943eb4573822255e96f6ee40278381b71fad6af5ad.
//
// Solidity: event Withdraw(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1, string nameToken0, string nameToken1)
func (_Api *ApiFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ApiWithdraw, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Withdraw", senderRule, toRule)
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

// ParseWithdraw is a log parse operation binding the contract event 0x9167547f76aeba7bf8001d943eb4573822255e96f6ee40278381b71fad6af5ad.
//
// Solidity: event Withdraw(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1, string nameToken0, string nameToken1)
func (_Api *ApiFilterer) ParseWithdraw(log types.Log) (*ApiWithdraw, error) {
	event := new(ApiWithdraw)
	if err := _Api.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

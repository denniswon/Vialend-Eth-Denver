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
const ApiABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cToken0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cToken1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cEth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_protocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxTotalSupply\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"_maxTwapDeviation\",\"type\":\"int24\"},{\"internalType\":\"uint32\",\"name\":\"_twapDuration\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"_quoteAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"_uniPortion\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_team\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"uFees0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"uFees1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lFees0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lFees1\",\"type\":\"uint256\"}],\"name\":\"CollectFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"ErrorLogging\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"subject\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"GeneralA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"subject\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"GeneralB\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"subject\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"GeneralS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MyLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MyLog2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MyLog4\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance1\",\"type\":\"uint256\"}],\"name\":\"RebalanceLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"Snapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Assetholder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"deposit0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"current0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"current1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CEther\",\"outputs\":[{\"internalType\":\"contractIcEther\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CToken0\",\"outputs\":[{\"internalType\":\"contractIcErc20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CToken1\",\"outputs\":[{\"internalType\":\"contractIcErc20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Fees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"U3Fees0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"U3Fees1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"LcFees0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"LcFees1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"AccruedProtocolFees0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"AccruedProtocolFees1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"contractIWETH9\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alloc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cHigh\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cLow\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curComp0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curComp1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToken1\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"doRebalance\",\"type\":\"bool\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"getPositionAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"baseAmount\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"}],\"name\":\"getQuoteAtTick\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"getSSLiquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"name\":\"getTwap\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteAmount\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"setGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMax\",\"type\":\"uint256\"}],\"name\":\"setMaxTotalSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_maxTwapDeviation\",\"type\":\"int24\"}],\"name\":\"setMaxTwapDeviation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_protocolFee\",\"type\":\"uint256\"}],\"name\":\"setProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_team\",\"type\":\"address\"}],\"name\":\"setTeam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_twapDuration\",\"type\":\"uint32\"}],\"name\":\"setTwapDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"ratio\",\"type\":\"uint8\"}],\"name\":\"setUniPortionRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"newLow\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"newHigh\",\"type\":\"int24\"}],\"name\":\"strategy1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20Contract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cErc20Contract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_numTokensToSupply\",\"type\":\"uint256\"}],\"name\":\"supplyErc20ToCompound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_cEtherContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ctoken\",\"type\":\"address\"}],\"name\":\"supplyEthToCompound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"team\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniPortion\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"percent\",\"type\":\"uint8\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ApiBin is the compiled bytecode used for deploying new contracts.
var ApiBin = "0x60c0604052600180546001600160801b031916670de0b6b3a76400001790553480156200002b57600080fd5b50604051620065473803806200654783398181016040526101808110156200005257600080fd5b5080516020808301516040808501516060860151608087015160a088015160c089015160e08a01516101008b01516101208c01516101408d0151610160909d01518951808b018b52600e81526d05669614c656e6420546f6b656e360941b818e019081528b51808d01909c5260048c52630564c54360e41b9d8c019d909d52600160005580519d9e9b9d999c989b979a96999598949793969295939491939092916200010291601e91906200052a565b5080516200011890601f9060208401906200052a565b506012602060006101000a81548160ff021916908360ff160217905550505033600260006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600360006101000a8154816001600160a01b0302191690836001600160a01b031602179055508b600660006101000a8154816001600160a01b0302191690836001600160a01b031602179055508b6001600160a01b0316630dfe16816040518163ffffffff1660e01b815260040160206040518083038186803b158015620001e657600080fd5b505afa158015620001fb573d6000803e3d6000fd5b505050506040513d60208110156200021257600080fd5b505160601b6001600160601b0319166080526040805163d21220a760e01b815290516001600160a01b038e169163d21220a7916004808301926020929190829003018186803b1580156200026557600080fd5b505afa1580156200027a573d6000803e3d6000fd5b505050506040513d60208110156200029157600080fd5b505160601b6001600160601b03191660a05260158054600160401b600160e01b031916680100000000000000006001600160a01b038d81169190910291909117909155601680546001600160a01b03199081168c841617909155601780549091168a83161790558b1662000335576040805162461bcd60e51b815260206004808301919091526024820152630ae8aa8960e31b604482015290519081900360640190fd5b8a601860006101000a8154816001600160a01b0302191690836001600160a01b03160217905550866005819055508b6001600160a01b031663d0c93a7c6040518163ffffffff1660e01b815260040160206040518083038186803b1580156200039d57600080fd5b505afa158015620003b2573d6000803e3d6000fd5b505050506040513d6020811015620003c957600080fd5b50516006805460029290920b62ffffff16600160a01b0262ffffff60a01b19909216919091179055606487106200042d576040805162461bcd60e51b815260206004820152600360248201526228232960e91b604482015290519081900360640190fd5b601486905560158054600180546001600160801b0387166001600160801b031990911617905560ff84166701000000000000000260ff60381b1963ffffffff881663ffffffff1960028b900b62ffffff81166401000000000262ffffff60201b199096169590951716171617909155600012620004d7576040805162461bcd60e51b815260206004820152600360248201526213551160ea1b604482015290519081900360640190fd5b60008463ffffffff161162000518576040805162461bcd60e51b8152602060048201526002602482015261151160f21b604482015290519081900360640190fd5b505050505050505050505050620005d6565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282620005625760008555620005ad565b82601f106200057d57805160ff1916838001178555620005ad565b82800160010185558215620005ad579182015b82811115620005ad57825182559160200191906001019062000590565b50620005bb929150620005bf565b5090565b5b80821115620005bb5760008155600101620005c0565b60805160601c60a05160601c615eb162000696600039806117c2528061185c5280611fe052806128965280612a045280612ab75280612c2b5280612edb5280613086528061314e52806131cd52806132f9528061384f5280613d6d528061452752806147f052508061102852806117a152806118215280611fa352806127f35280612a7d5280612b925280612e865280612f145280612fdc528061305b52806132bc52806137b65280613d4c528061448752806147cf5250615eb16000f3fe60806040526004361061036f5760003560e01c806370a08231116101c6578063c433c80a116100f7578063d348799711610095578063dd62ed3e1161006f578063dd62ed3e14610ce2578063f2a40db814610d1d578063f39c38a014610d47578063fa461e3314610d5c57610376565b8063d348799714610c31578063d368867f14610cb8578063d8d7f96f14610ccd57610376565b8063cbcf94aa116100d1578063cbcf94aa14610bbb578063ccac7d8b14610bd0578063cda206b014610be5578063d21220a714610c1c57610376565b8063c433c80a14610b49578063c6ab5d9014610b79578063c8a746a614610ba657610376565b8063a9059cbb11610164578063ab033ea91161013e578063ab033ea914610ad7578063ad5c464814610b0a578063b0e21e8a14610b1f578063bbc001c314610b3457610376565b8063a9059cbb14610a34578063a91ef6eb14610a6d578063aabfd57214610aa457610376565b806385f2aef2116101a057806385f2aef2146109bc5780638e843c6c146109d157806395d89b41146109e6578063a457c2d7146109fb57610376565b806370a082311461093157806373232c6014610964578063787dce3d1461099257610376565b806339509351116102a057806353257e001161023e578063624a177a11610218578063624a177a1461083e5780636253c7181461086b57806366d7505b146108a25780636ea056a9146108f857610376565b806353257e00146107b65780635aa6e675146107cb5780635bb6aa85146107e057610376565b80633cbff3fe1161027a5780633cbff3fe146106d45780633f3e4c111461070157806343a0d0661461072b57806343c57a271461076357610376565b806339509351146106105780633aaa36e6146106495780633b7ba25f1461069157610376565b806318160ddd1161030d5780632ab4d052116102e75780632ab4d052146105a6578063313ce567146105bb57806331dc5b95146105e657806332380717146105fb57610376565b806318160ddd14610527578063238efcbc1461054e57806323b872dd1461056357610376565b8063095ea7b311610349578063095ea7b3146104635780630d40886d146104b05780630dfe1681146104e15780631755ff211461051257610376565b806306b0b6561461037857806306fdde03146103a6578063095cf5c61461043057610376565b3661037657005b005b34801561038457600080fd5b5061038d610de3565b6040805192835260208301919091528051918290030190f35b3480156103b257600080fd5b506103bb610ef4565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103f55781810151838201526020016103dd565b50505050905090810190601f1680156104225780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561043c57600080fd5b506103766004803603602081101561045357600080fd5b50356001600160a01b0316610f8b565b34801561046f57600080fd5b5061049c6004803603604081101561048657600080fd5b506001600160a01b038135169060200135610ff9565b604080519115158252519081900360200190f35b3480156104bc57600080fd5b506104c5611017565b604080516001600160801b039092168252519081900360200190f35b3480156104ed57600080fd5b506104f6611026565b604080516001600160a01b039092168252519081900360200190f35b34801561051e57600080fd5b506104f661104a565b34801561053357600080fd5b5061053c611059565b60408051918252519081900360200190f35b34801561055a57600080fd5b5061037661105f565b34801561056f57600080fd5b5061049c6004803603606081101561058657600080fd5b506001600160a01b038135811691602081013590911690604001356110c6565b3480156105b257600080fd5b5061053c61114e565b3480156105c757600080fd5b506105d0611154565b6040805160ff9092168252519081900360200190f35b3480156105f257600080fd5b506105d061115d565b34801561060757600080fd5b506104f661116d565b34801561061c57600080fd5b5061049c6004803603604081101561063357600080fd5b506001600160a01b03813516906020013561117c565b34801561065557600080fd5b5061065e6111ca565b604080519687526020870195909552858501939093526060850191909152608084015260a0830152519081900360c00190f35b34801561069d57600080fd5b5061053c600480360360608110156106b457600080fd5b506001600160a01b038135811691602081013590911690604001356111df565b3480156106e057600080fd5b50610376600480360360208110156106f757600080fd5b503560020b61152c565b34801561070d57600080fd5b506103766004803603602081101561072457600080fd5b50356115ec565b34801561073757600080fd5b506103766004803603606081101561074e57600080fd5b5080359060208101359060400135151561168d565b34801561076f57600080fd5b5061053c6004803603608081101561078657600080fd5b50803560020b906001600160801b03602082013516906001600160a01b036040820135811691606001351661196d565b3480156107c257600080fd5b506104f6611a5f565b3480156107d757600080fd5b506104f6611a6e565b3480156107ec57600080fd5b506108136004803603602081101561080357600080fd5b50356001600160a01b0316611a7d565b6040805195865260208601949094528484019290925260608401526080830152519081900360a00190f35b34801561084a57600080fd5b506103766004803603602081101561086157600080fd5b503560ff16611aac565b34801561087757600080fd5b506103766004803603604081101561088e57600080fd5b508035600290810b9160200135900b611b5d565b3480156108ae57600080fd5b506108e1600480360360408110156108c557600080fd5b5080356001600160a01b0316906020013563ffffffff16611c58565b6040805160029290920b8252519081900360200190f35b34801561090457600080fd5b506103766004803603604081101561091b57600080fd5b506001600160a01b038135169060200135611f55565b34801561093d57600080fd5b5061053c6004803603602081101561095457600080fd5b50356001600160a01b0316612066565b61049c6004803603604081101561097a57600080fd5b506001600160a01b0381358116916020013516612085565b34801561099e57600080fd5b50610376600480360360208110156109b557600080fd5b5035612265565b3480156109c857600080fd5b506104f66122fa565b3480156109dd57600080fd5b506108e1612309565b3480156109f257600080fd5b506103bb612319565b348015610a0757600080fd5b5061049c60048036036040811015610a1e57600080fd5b506001600160a01b038135169060200135612379565b348015610a4057600080fd5b5061049c60048036036040811015610a5757600080fd5b506001600160a01b0381351690602001356123e1565b348015610a7957600080fd5b5061038d60048036036040811015610a9057600080fd5b508035600290810b9160200135900b6123f5565b348015610ab057600080fd5b5061053c60048036036020811015610ac757600080fd5b50356001600160a01b031661248c565b348015610ae357600080fd5b5061037660048036036020811015610afa57600080fd5b50356001600160a01b031661249e565b348015610b1657600080fd5b506104f661250c565b348015610b2b57600080fd5b5061053c61251b565b348015610b4057600080fd5b50610376612521565b348015610b5557600080fd5b5061037660048036036020811015610b6c57600080fd5b503563ffffffff16612578565b348015610b8557600080fd5b5061038d60048036036020811015610b9c57600080fd5b503560ff1661262a565b348015610bb257600080fd5b5061053c6129c8565b348015610bc757600080fd5b506104f66129ce565b348015610bdc57600080fd5b5061053c6129e4565b348015610bf157600080fd5b506104c560048036036040811015610c0857600080fd5b508035600290810b9160200135900b6129ea565b348015610c2857600080fd5b506104f6612a02565b348015610c3d57600080fd5b5061037660048036036060811015610c5457600080fd5b813591602081013591810190606081016040820135600160201b811115610c7a57600080fd5b820183602082011115610c8c57600080fd5b803590602001918460018302840111600160201b83111715610cad57600080fd5b509092509050612a26565b348015610cc457600080fd5b506108e1612ae4565b348015610cd957600080fd5b50610376612af4565b348015610cee57600080fd5b5061053c60048036036040811015610d0557600080fd5b506001600160a01b03813581169160200135166131fd565b348015610d2957600080fd5b506104f660048036036020811015610d4057600080fd5b5035613228565b348015610d5357600080fd5b506104f6613252565b348015610d6857600080fd5b5061037660048036036060811015610d7f57600080fd5b813591602081013591810190606081016040820135600160201b811115610da557600080fd5b820183602082011115610db757600080fd5b803590602001918460018302840111600160201b83111715610dd857600080fd5b509092509050613261565b600080601560089054906101000a90046001600160a01b03166001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610e4857600080fd5b505afa158015610e5c573d6000803e3d6000fd5b505050506040513d6020811015610e7257600080fd5b5051601654604080516370a0823160e01b815230600482015290519294506001600160a01b03909116916370a0823191602480820192602092909190829003018186803b158015610ec257600080fd5b505afa158015610ed6573d6000803e3d6000fd5b505050506040513d6020811015610eec57600080fd5b505191929050565b601e8054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610f805780601f10610f5557610100808354040283529160200191610f80565b820191906000526020600020905b815481529060010190602001808311610f6357829003601f168201915b505050505090505b90565b6002546001600160a01b03163314610fd7576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b600380546001600160a01b0319166001600160a01b0392909216919091179055565b600061100d611006613320565b8484613324565b5060015b92915050565b6001546001600160801b031681565b7f000000000000000000000000000000000000000000000000000000000000000081565b6006546001600160a01b031690565b601d5490565b6004546001600160a01b031633146110b2576040805162461bcd60e51b815260206004820152601160248201527070656e64696e67476f7665726e616e636560781b604482015290519081900360640190fd5b600280546001600160a01b03191633179055565b60006110d3848484613410565b611143846110df613320565b61113e85604051806060016040528060288152602001615d7b602891396001600160a01b038a166000908152601c602052604081209061111d613320565b6001600160a01b03168152602081019190915260400160002054919061356d565b613324565b5060015b9392505050565b60145481565b60205460ff1690565b601554600160381b900460ff1681565b6016546001600160a01b031681565b600061100d611189613320565b8461113e85601c600061119a613320565b6001600160a01b03908116825260208083019390935260409182016000908120918c168152925290205490613604565b600e54600f5460105460115460125460135486565b6000836001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561122e57600080fd5b505afa158015611242573d6000803e3d6000fd5b505050506040513d602081101561125857600080fd5b5051821115611298576040805162461bcd60e51b815260206004820152600760248201526662616c616e636560c81b604482015290519081900360640190fd5b600084905060008490506000816001600160a01b031663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156112df57600080fd5b505af11580156112f3573d6000803e3d6000fd5b505050506040513d602081101561130957600080fd5b50516040805160208101839052818152601b818301527f45786368616e6765205261746520287363616c6564207570293a20000000000060608201529051919250600080516020615e0d833981519152919081900360800190a16000826001600160a01b031663ae9d70b06040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156113a057600080fd5b505af11580156113b4573d6000803e3d6000fd5b505050506040513d60208110156113ca57600080fd5b505160408051602081018390528181526018818301527f537570706c7920526174653a20287363616c656420757029000000000000000060608201529051919250600080516020615e0d833981519152919081900360800190a1836001600160a01b031663095ea7b388886040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b15801561147b57600080fd5b505af115801561148f573d6000803e3d6000fd5b505050506040513d60208110156114a557600080fd5b50506040805163140e25ad60e31b81526004810188905290516000916001600160a01b0386169163a0712d689160248082019260209290919082900301818787803b1580156114f357600080fd5b505af1158015611507573d6000803e3d6000fd5b505050506040513d602081101561151d57600080fd5b50519998505050505050505050565b6002546001600160a01b03163314611578576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b60008160020b136115c3576040805162461bcd60e51b815260206004820152601060248201526f36b0bc2a3bb0b82232bb34b0ba34b7b760811b604482015290519081900360640190fd5b6015805460029290920b62ffffff16600160201b0266ffffff0000000019909216919091179055565b60026000541415611632576040805162461bcd60e51b815260206004820152601f6024820152600080516020615c53833981519152604482015290519081900360640190fd5b60026000819055546001600160a01b03163314611683576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b6014556001600055565b600260005414156116d3576040805162461bcd60e51b815260206004820152601f6024820152600080516020615c53833981519152604482015290519081900360640190fd5b6002600055821515806116e65750600082115b611720576040805162461bcd60e51b815260206004808301919091526024820152630616d74360e41b604482015290519081900360640190fd5b6003546001600160a01b0316331415611765576040805162461bcd60e51b8152602060048201526002602482015261746f60f01b604482015290519081900360640190fd5b61176e3361365e565b6006546015546117e691611793916001600160a01b039091169063ffffffff16611c58565b6001546001600160801b03167f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000061196d565b602155600080806117f786866136d9565b92509250925061180561373c565b5061180e613779565b8115611849576118496001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016333085613a09565b8015611884576118846001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016333084613a09565b61188e3384613a63565b601454611899611059565b11156118d2576040805162461bcd60e51b815260206004820152600360248201526204341560ec1b604482015290519081900360640190fd5b336000908152600d602052604090208054830181556001810180548301905560028101805484019055600301805482019055831561191d5761191d600080611918611059565b613b55565b6040805184815260208101849052808201839052905133917f36af321ec8d3c75236829c5317affd40ddb308863a1236d2d277a4025cccee1e919081900360600190a25050600160005550505050565b60008061197986613dd2565b90506001600160801b036001600160a01b038216116119e8576001600160a01b03808216800290848116908616106119c8576119c3600160c01b876001600160801b031683614103565b6119e0565b6119e081876001600160801b0316600160c01b614103565b925050611a56565b6000611a026001600160a01b03831680600160401b614103565b9050836001600160a01b0316856001600160a01b031610611a3a57611a35600160801b876001600160801b031683614103565b611a52565b611a5281876001600160801b0316600160801b614103565b9250505b50949350505050565b6017546001600160a01b031681565b6002546001600160a01b031681565b600d60205260009081526040902080546001820154600283015460038401546004909401549293919290919085565b6002546001600160a01b03163314611af8576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b60648160ff161115611b39576040805162461bcd60e51b8152602060048201526005602482015264726174696f60d81b604482015290519081900360640190fd5b6015805460ff909216600160381b0267ff0000000000000019909216919091179055565b60026000541415611ba3576040805162461bcd60e51b815260206004820152601f6024820152600080516020615c53833981519152604482015290519081900360640190fd5b60026000819055546001600160a01b03163314611bf4576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b6000611bfe611059565b905060008111611c3a576040805162461bcd60e51b8152602060048201526002602482015261074360f41b604482015290519081900360640190fd5b611c4261373c565b50611c4e838383613b55565b5050600160005550565b600063ffffffff8216611c98576040805162461bcd60e51b815260206004820152600360248201526207842560ec1b604482015290519081900360640190fd5b6040805160028082526060820183526000926020830190803683370190505090508281600081518110611cc757fe5b602002602001019063ffffffff16908163ffffffff1681525050600081600181518110611cf057fe5b63ffffffff90921660209283029190910182015260405163883bdbfd60e01b8152600481018281528351602483015283516000936001600160a01b0389169363883bdbfd938793909283926044019185820191028083838b5b83811015611d61578181015183820152602001611d49565b505050509050019250505060006040518083038186803b158015611d8457600080fd5b505afa158015611d98573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040908152811015611dc157600080fd5b8101908080516040519392919084600160201b821115611de057600080fd5b908301906020820185811115611df557600080fd5b82518660208202830111600160201b82111715611e1157600080fd5b82525081516020918201928201910280838360005b83811015611e3e578181015183820152602001611e26565b5050505090500160405260200180516040519392919084600160201b821115611e6657600080fd5b908301906020820185811115611e7b57600080fd5b82518660208202830111600160201b82111715611e9757600080fd5b82525081516020918201928201910280838360005b83811015611ec4578181015183820152602001611eac565b50505050905001604052505050509050600081600081518110611ee357fe5b602002602001015182600181518110611ef857fe5b60200260200101510390508463ffffffff168160060b81611f1557fe5b05935060008160060b128015611f3f57508463ffffffff168160060b81611f3857fe5b0760060b15155b15611f4c57600019909301925b50505092915050565b6002546001600160a01b03163314611fa1576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b03161415801561201557507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b031614155b61204e576040805162461bcd60e51b81526020600482015260056024820152643a37b5b2b760d91b604482015290519081900360640190fd5b6120626001600160a01b03831633836141b2565b5050565b6001600160a01b0381166000908152601b60205260409020545b919050565b60008083905060008390506000816001600160a01b031663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156120cd57600080fd5b505af11580156120e1573d6000803e3d6000fd5b505050506040513d60208110156120f757600080fd5b505160408051602081018390528181526023918101829052919250600080516020615e0d83398151915291839181906060820190615cb782396040019250505060405180910390a16000826001600160a01b031663ae9d70b06040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561217c57600080fd5b505af1158015612190573d6000803e3d6000fd5b505050506040513d60208110156121a657600080fd5b5051604080516020818101849052828252818301527f537570706c7920526174653a20287363616c656420757020627920316531382960608201529051919250600080516020615e0d833981519152919081900360800190a1836001600160a01b0316631249c58b6203d090476040518363ffffffff1660e01b81526004016000604051808303818589803b15801561223e57600080fd5b5088f1158015612252573d6000803e3d6000fd5b5060019c9b505050505050505050505050565b6002546001600160a01b031633146122b1576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b60148111156122f5576040805162461bcd60e51b815260206004820152600b60248201526a70726f746f636f6c46656560a81b604482015290519081900360640190fd5b600555565b6003546001600160a01b031681565b600654600160d01b900460020b81565b601f8054604080516020600260001961010060018716150201909416939093048085018490048402820184019092528181526060939092909190830182828015610f805780601f10610f5557610100808354040283529160200191610f80565b600061100d612386613320565b8461113e85604051806060016040528060258152602001615e5760259139601c60006123b0613320565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919061356d565b600061100d6123ee613320565b8484613410565b60008060008060006124078787614204565b945094505050925061241a8787856142bc565b600554919650945060009061243190606490614366565b905061245b612454606461244e6001600160801b038716856143c3565b9061441c565b8790613604565b955061247f612478606461244e6001600160801b038616856143c3565b8690613604565b9450505050509250929050565b600c6020526000908152604090205481565b6002546001600160a01b031633146124ea576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b600480546001600160a01b0319166001600160a01b0392909216919091179055565b6018546001600160a01b031681565b60055481565b6002546001600160a01b0316331461256d576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b61257561373c565b50565b6002546001600160a01b031633146125c4576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b60008163ffffffff161161260e576040805162461bcd60e51b815260206004820152600c60248201526b3a3bb0b8223ab930ba34b7b760a11b604482015290519081900360640190fd5b6015805463ffffffff191663ffffffff92909216919091179055565b60008060026000541415612673576040805162461bcd60e51b815260206004820152601f6024820152600080516020615c53833981519152604482015290519081900360640190fd5b600260005533606460ff851611156126b7576040805162461bcd60e51b8152602060048201526002602482015261706360f01b604482015290519081900360640190fd5b60006126c1611059565b905060006126e0606461244e8860ff166126da87612066565b906143c3565b90508181111561271d576040805162461bcd60e51b815260206004820152600360248201526273683160e81b604482015290519081900360640190fd5b61272561373c565b506127368261244e836126da614483565b94506127488261244e836126da614523565b93506127548382614592565b600e546127749061276b90849061244e90856143c3565b600e5490614366565b600e55600f546127979061278e90849061244e90856143c3565b600f5490614366565b600f556010546127ba906127b190849061244e90856143c3565b60105490614366565b6010556011546127dd906127d490849061244e90856143c3565b60115490614366565b60115584156128835761281a6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001684876141b2565b6001600160a01b0383166000908152600d6020526040902060020154851061285d576001600160a01b0383166000908152600d602052604090206002015461285f565b845b6001600160a01b0384166000908152600d6020526040902060020180549190910390555b8315612926576128bd6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001684866141b2565b6001600160a01b0383166000908152600d60205260409020600301548410612900576001600160a01b0383166000908152600d6020526040902060030154612902565b835b6001600160a01b0384166000908152600d6020526040902060030180549190910390555b61292f83612066565b61296a5761293c8361468e565b6001600160a01b0383166000908152600d602052604081208181556001810182905560028101829055600301555b61297660008084613b55565b6040805182815260208101879052808201869052905133917f02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94919081900360600190a250506001600055509092909150565b60195481565b601554600160401b90046001600160a01b031681565b601a5481565b60006129f68383614204565b50929695505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6006546001600160a01b03163314612a6a576040805162461bcd60e51b8152602060048201526002602482015261505360f01b604482015290519081900360640190fd5b8315612aa457612aa46001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001633866141b2565b8215612ade57612ade6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001633856141b2565b50505050565b600654600160b81b900460020b81565b6002546001600160a01b03163314612b40576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b612b486147a7565b600080612b53610de3565b91509150600080831180612b675750600082115b600b54604080516370a0823160e01b8152306004820152905192935090916000916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016916370a0823191602480820192602092909190829003018186803b158015612bd957600080fd5b505afa158015612bed573d6000803e3d6000fd5b505050506040513d6020811015612c0357600080fd5b5051604080516370a0823160e01b815230600482015290519192506000916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016916370a08231916024808301926020929190829003018186803b158015612c7157600080fd5b505afa158015612c85573d6000803e3d6000fd5b505050506040513d6020811015612c9b57600080fd5b505190506000612ca9611059565b905060005b84811015612f0f576000612ce2600b8381548110612cc857fe5b6000918252602090912001546001600160a01b0316612066565b90508615612d1857612d18600b8381548110612cfa57fe5b6000918252602090912001546001600160a01b031682858c8c61481e565b612d43600b8381548110612d2857fe5b6000918252602090912001546001600160a01b031682614592565b6000600d6000600b8581548110612d5657fe5b60009182526020808320909101546001600160a01b03168352820192909252604001812091909155600b8054600d91839186908110612d9157fe5b60009182526020808320909101546001600160a01b03168352820192909252604001812060010191909155600b8054600d91839186908110612dcf57fe5b60009182526020808320909101546001600160a01b03168352820192909252604001812060020191909155600b8054600d91839186908110612e0d57fe5b60009182526020808320909101546001600160a01b03168352820192909252604001812060030191909155612e468461244e88856143c3565b90506000612e588561244e88866143c3565b90508115612eaf57612eaf600b8581548110612e7057fe5b6000918252602090912001546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169116846141b2565b8015612f0457612f04600b8581548110612ec557fe5b6000918252602090912001546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169116836141b2565b505050600101612cae565b5060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015612f7f57600080fd5b505afa158015612f93573d6000803e3d6000fd5b505050506040513d6020811015612fa957600080fd5b5051111561308257600254604080516370a0823160e01b81523060048201529051613082926001600160a01b03908116927f0000000000000000000000000000000000000000000000000000000000000000909116916370a0823191602480820192602092909190829003018186803b15801561302557600080fd5b505afa158015613039573d6000803e3d6000fd5b505050506040513d602081101561304f57600080fd5b50516001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691906141b2565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156130f157600080fd5b505afa158015613105573d6000803e3d6000fd5b505050506040513d602081101561311b57600080fd5b505111156131f457600254604080516370a0823160e01b815230600482015290516131f4926001600160a01b03908116927f0000000000000000000000000000000000000000000000000000000000000000909116916370a0823191602480820192602092909190829003018186803b15801561319757600080fd5b505afa1580156131ab573d6000803e3d6000fd5b505050506040513d60208110156131c157600080fd5b50516001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691906141b2565b50505050505050565b6001600160a01b039182166000908152601c6020908152604080832093909416825291909152205490565b600b818154811061323857600080fd5b6000918252602090912001546001600160a01b0316905081565b6004546001600160a01b031681565b6006546001600160a01b031633146132a6576040805162461bcd60e51b815260206004820152600360248201526228299960e91b604482015290519081900360640190fd5b60008413156132e3576132e36001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001633866141b2565b6000831315612ade57612ade6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001633856141b2565b3390565b6001600160a01b0383166133695760405162461bcd60e51b8152600401808060200182810382526024815260200180615de96024913960400191505060405180910390fd5b6001600160a01b0382166133ae5760405162461bcd60e51b8152600401808060200182810382526022815260200180615c956022913960400191505060405180910390fd5b6001600160a01b038084166000818152601c6020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166134555760405162461bcd60e51b8152600401808060200182810382526025815260200180615dc46025913960400191505060405180910390fd5b6001600160a01b03821661349a5760405162461bcd60e51b8152600401808060200182810382526023815260200180615c306023913960400191505060405180910390fd5b6134a5838383613dcd565b6134e281604051806060016040528060268152602001615cda602691396001600160a01b0386166000908152601b6020526040902054919061356d565b6001600160a01b038085166000908152601b602052604080822093909355908416815220546135119082613604565b6001600160a01b038084166000818152601b602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600081848411156135fc5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156135c15781810151838201526020016135a9565b50505050905090810190601f1680156135ee5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600082820183811015611147576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b61366781614885565b61257557600b80546001810182557f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db90180546001600160a01b0319166001600160a01b03939093169283179055546000918252600c6020908152604080842092909255600d9052902043600490910155565b60215460009083908390613718576040805162461bcd60e51b81526020600482015260016024820152600560fc1b604482015290519081900360640190fd5b6021546001546001600160801b031682028161373057fe5b04820192509250925092565b600080613747611059565b90506000811161375b576000915050610f88565b6137636147a7565b8015613771576137716148a2565b600191505090565b6000613783611059565b9050806137905750613a07565b600b54604080516370a0823160e01b815230600482015290516000916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016916370a0823191602480820192602092909190829003018186803b1580156137fd57600080fd5b505afa158015613811573d6000803e3d6000fd5b505050506040513d602081101561382757600080fd5b5051604080516370a0823160e01b815230600482015290519192506000916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016916370a08231916024808301926020929190829003018186803b15801561389557600080fd5b505afa1580156138a9573d6000803e3d6000fd5b505050506040513d60208110156138bf57600080fd5b5051905060005b83811015613a015760006138e0600b8381548110612cc857fe5b905080156139f85760006138f88761244e87856143c3565b9050600061390a8861244e87866143c3565b9050600061391883836136d9565b505090508381146139f45761394e600b868154811061393357fe5b6000918252602090912001546001600160a01b031685614592565b613979600b868154811061395e57fe5b6000918252602090912001546001600160a01b031682613a63565b82600d6000600b888154811061398b57fe5b60009182526020808320909101546001600160a01b03168352820192909252604001812060020191909155600b80548492600d929091899081106139cb57fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020600301555b5050505b506001016138c6565b50505050505b565b604080516001600160a01b0380861660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052612ade908590614a31565b6001600160a01b038216613abe576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b613aca60008383613dcd565b601d54613ad79082613604565b601d556001600160a01b0382166000908152601b6020526040902054613afd9082613604565b6001600160a01b0383166000818152601b602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b60065460408051633850c7bd60e01b815290516000926001600160a01b031691633850c7bd9160048083019260e0929190829003018186803b158015613b9a57600080fd5b505afa158015613bae573d6000803e3d6000fd5b505050506040513d60e0811015613bc457600080fd5b5060200151905081613bd65750613dcd565b8360020b6000148015613bec57508260020b6000145b15613c8f57600654600160d01b9004600290810b900b158015613c1d5750600654600160b81b9004600290810b900b155b15613c285750613dcd565b6006546002600160b81b8204810b600160d01b8304820b03810b81900591600160a01b9004810b9081810b90848401900b81613c6057fe5b600654919005919091029450600160a01b9004600290810b9081810b90838503900b81613c8957fe5b05029450505b613caa848483600660149054906101000a900460020b614ae2565b601554600160381b900460ff1615613d1757601554600090613cdf9060649061244e90600160381b900460ff166126da614483565b90506000613d06606461244e601560079054906101000a900460ff1660ff166126da614523565b9050613d1486868484614cb6565b50505b6015546064600160381b90910460ff161015612ade576000613d37614483565b90506000613d43614523565b90506000613d937f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000008585614e36565b9050806131f4576040805162461bcd60e51b815260206004820152600360248201526204d43760ec1b604482015290519081900360640190fd5b505050565b60008060008360020b12613de9578260020b613df1565b8260020b6000035b9050620d89e8811115613e2f576040805162461bcd60e51b81526020600482015260016024820152601560fa1b604482015290519081900360640190fd5b600060018216613e4357600160801b613e55565b6ffffcb933bd6fad37aa2d162d1a5940015b70ffffffffffffffffffffffffffffffffff1690506002821615613e89576ffff97272373d413259a46990580e213a0260801c5b6004821615613ea8576ffff2e50f5f656932ef12357cf3c7fdcc0260801c5b6008821615613ec7576fffe5caca7e10e4e61c3624eaa0941cd00260801c5b6010821615613ee6576fffcb9843d60f6159c9db58835c9266440260801c5b6020821615613f05576fff973b41fa98c081472e6896dfb254c00260801c5b6040821615613f24576fff2ea16466c96a3843ec78b326b528610260801c5b6080821615613f43576ffe5dee046a99a2a811c461f1969c30530260801c5b610100821615613f63576ffcbe86c7900a88aedcffc83b479aa3a40260801c5b610200821615613f83576ff987a7253ac413176f2b074cf7815e540260801c5b610400821615613fa3576ff3392b0822b70005940c7a398e4b70f30260801c5b610800821615613fc3576fe7159475a2c29b7443b29c7fa6e889d90260801c5b611000821615613fe3576fd097f3bdfd2022b8845ad8f792aa58250260801c5b612000821615614003576fa9f746462d870fdf8a65dc1f90e061e50260801c5b614000821615614023576f70d869a156d2a1b890bb3df62baf32f70260801c5b618000821615614043576f31be135f97d08fd981231505542fcfa60260801c5b62010000821615614064576f09aa508b5b7a84e1c677de54f3e99bc90260801c5b62020000821615614084576e5d6af8dedb81196699c329225ee6040260801c5b620400008216156140a3576d2216e584f5fa1ea926041bedfe980260801c5b620800008216156140c0576b048a170391f7dc42444e8fa20260801c5b60008460020b13156140db5780600019816140d757fe5b0490505b600160201b8106156140ee5760016140f1565b60005b60ff16602082901c0192505050919050565b6000808060001985870986860292508281109083900303905080614139576000841161412e57600080fd5b508290049050611147565b80841161414557600080fd5b6000848688096000868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052613dcd908490614a31565b600080600080600080614218308989614f0b565b6006546040805163514ea4bf60e01b81526004810184905290519293506001600160a01b039091169163514ea4bf9160248082019260a092909190829003018186803b15801561426757600080fd5b505afa15801561427b573d6000803e3d6000fd5b505050506040513d60a081101561429157600080fd5b508051602082015160408301516060840151608090940151929c919b50995091975095509350505050565b6000806000600660009054906101000a90046001600160a01b03166001600160a01b0316633850c7bd6040518163ffffffff1660e01b815260040160e06040518083038186803b15801561430f57600080fd5b505afa158015614323573d6000803e3d6000fd5b505050506040513d60e081101561433957600080fd5b505190506143598161434a88613dd2565b61435388613dd2565b87614f61565b9250925050935093915050565b6000828211156143bd576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6000826143d257506000611011565b828202828482816143df57fe5b04146111475760405162461bcd60e51b8152600401808060200182810382526021815260200180615d5a6021913960400191505060405180910390fd5b6000808211614472576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b81838161447b57fe5b049392505050565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156144f257600080fd5b505afa158015614506573d6000803e3d6000fd5b505050506040513d602081101561451c57600080fd5b5051905090565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156144f257600080fd5b6001600160a01b0382166145d75760405162461bcd60e51b8152600401808060200182810382526021815260200180615da36021913960400191505060405180910390fd5b6145e382600083613dcd565b61462081604051806060016040528060228152602001615c73602291396001600160a01b0385166000908152601b6020526040902054919061356d565b6001600160a01b0383166000908152601b6020526040902055601d546146469082614366565b601d556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b61469781614885565b6146a057612575565b6001600160a01b0381166000908152600c6020526040902054600b80546000199283019281019081106146cf57fe5b600091825260209091200154600b80546001600160a01b0390921691839081106146f557fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600101600c6000600b848154811061473857fe5b60009182526020808320909101546001600160a01b03168352820192909252604001902055600b80548061476857fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b0384168252600c905260408120555050565b6000806147b2610de3565b9150915060008211806147c55750600081115b15614816576148167f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000008484614ffd565b612062615328565b600061482e8461244e85886143c3565b905060006148408561244e85896143c3565b9050811561486757601554614867908890600160401b90046001600160a01b031684615467565b80156131f4576016546131f49088906001600160a01b031683615467565b6001600160a01b03166000908152600c6020526040902054151590565b60055460095460075460009182916148b991613604565b600a546008546148c891613604565b600754600854600954600a546040805194855260208501939093528383019190915260608301525192945090925030917fb9a5bd9c6fc68ddbb2525a58437661292930f5a4adc18d5fa410ac494a15d0c39181900360800190a2600754600e5461493191613604565b600e55600854600f5461494391613604565b600f5560095460105461495591613604565b601055600a5460115461496791613604565b6011556000600a819055600981905560088190556007819055808415613a015783156149b05761499c606461244e86886143c3565b6012549092506149ac9083613604565b6012555b82156149d9576149c5606461244e85886143c3565b6013549091506149d59082613604565b6013555b60008211806149e85750600081115b15613a015760006149f983836136d9565b5050600354909150614a14906001600160a01b031682613a63565b600354614a29906001600160a01b031661365e565b505050505050565b6000614a86826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166155899092919063ffffffff16565b805190915015613dcd57808060200190516020811015614aa557600080fd5b5051613dcd5760405162461bcd60e51b815260040180806020018281038252602a815260200180615e2d602a913960400191505060405180910390fd5b8260020b8460020b12614b21576040805162461bcd60e51b8152602060048201526002602482015261563160f01b604482015290519081900360640190fd5b8160020b8460020b12614b60576040805162461bcd60e51b81526020600482015260026024820152612b1960f11b604482015290519081900360640190fd5b8160020b8360020b13614b9f576040805162461bcd60e51b8152602060048201526002602482015261563360f01b604482015290519081900360640190fd5b620d89e719600285900b1215614be1576040805162461bcd60e51b8152602060048201526002602482015261158d60f21b604482015290519081900360640190fd5b620d89e8600284900b1315614c22576040805162461bcd60e51b8152602060048201526002602482015261563560f01b604482015290519081900360640190fd5b8060020b8460020b81614c3157fe5b0760020b15614c6c576040805162461bcd60e51b81526020600482015260026024820152612b1b60f11b604482015290519081900360640190fd5b8060020b8360020b81614c7b57fe5b0760020b15612ade576040805162461bcd60e51b8152602060048201526002602482015261563760f01b604482015290519081900360640190fd5b60065460408051633850c7bd60e01b815290516000926001600160a01b031691633850c7bd9160048083019260e0929190829003018186803b158015614cfb57600080fd5b505afa158015614d0f573d6000803e3d6000fd5b505050506040513d60e0811015614d2557600080fd5b505190506000614d4882614d3888613dd2565b614d4188613dd2565b87876155a0565b60065460408051633c8a7d8d60e01b815230600482015260028a810b602483015289900b60448201526001600160801b038416606482015260a06084820152600060a4820181905282519495506001600160a01b0390931693633c8a7d8d9360c480840194938390030190829087803b158015614dc457600080fd5b505af1158015614dd8573d6000803e3d6000fd5b505050506040513d6040811015614dee57600080fd5b505060068054600297880b62ffffff908116600160b81b0262ffffff60b81b199890990b16600160d01b0262ffffff60d01b1990911617959095169590951790935550505050565b6018546015546016546019859055601a8490556000926001600160a01b0390811692600160401b90048116918116908816831415614e9f57614e7786615664565b601754614e8d906001600160a01b031682612085565b50614e998782876111df565b50614efd565b826001600160a01b0316876001600160a01b03161415614ee457614ec285615664565b601754614ed8906001600160a01b031683612085565b50614e998883886111df565b614eef8883886111df565b50614efb8782876111df565b505b506001979650505050505050565b6040805160609490941b6bffffffffffffffffffffffff1916602080860191909152600293840b60e890811b60348701529290930b90911b60378401528051808403601a018152603a9093019052815191012090565b600080836001600160a01b0316856001600160a01b03161115614f82579293925b846001600160a01b0316866001600160a01b031611614fad57614fa68585856156cb565b9150614ff4565b836001600160a01b0316866001600160a01b03161015614fe657614fd28685856156cb565b9150614fdf858785615734565b9050614ff4565b614ff1858585615734565b90505b94509492505050565b601854601554601654604080516370a0823160e01b815230600482015290516001600160a01b0394851694600160401b90940484169392831692600092908a16916370a0823191602480820192602092909190829003018186803b15801561506457600080fd5b505afa158015615078573d6000803e3d6000fd5b505050506040513d602081101561508e57600080fd5b5051604080516370a0823160e01b815230600482015290519192506000916001600160a01b038a16916370a08231916024808301926020929190829003018186803b1580156150dc57600080fd5b505afa1580156150f0573d6000803e3d6000fd5b505050506040513d602081101561510657600080fd5b505190506001600160a01b03898116908616141561513f576151288785615777565b6151306157f5565b61513a8684615777565b615189565b846001600160a01b0316886001600160a01b03161415615175576151638684615777565b61516b6157f5565b61513a8785615777565b61517f8684615777565b6151898785615777565b600061520e838b6001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156151dc57600080fd5b505afa1580156151f0573d6000803e3d6000fd5b505050506040513d602081101561520657600080fd5b505190614366565b90506000615263838b6001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156151dc57600080fd5b9050600082118015615276575060195482115b1561528d57601954615289908390614366565b6009555b60008111801561529e5750601a5481115b156152b557601a546152b1908290614366565b600a555b600954600a54604080516020810193909352828101919091526060808352600e908301526d3932b6b7bb32b632b73234b7339d60911b6080830152517fb626a5b8d37253edd986272eff7a2541a1538cad44624e413d51713f0b2a414d9181900360a00190a15050505050505050505050565b600654600160b81b9004600290810b900b1580156153545750600654600160d01b9004600290810b900b155b1561535e57613a07565b60065460009061538390600160b81b8104600290810b91600160d01b9004900b614204565b505060065492935060009283925082915081906153b690600160b81b8104600290810b91600160d01b9004900b8761584b565b93509350935093506000856001600160801b031611156153eb576153da8285614366565b6007556153e78184614366565b6008555b60065461540d90600160b81b8104600290810b91600160d01b9004900b614204565b5050604080516001600160801b03851660208201528181526034918101829052939850600080516020615e0d83398151915293899350915081906060820190615d2682396040019250505060405180910390a15050505050565b6017546001600160a01b0390811690831681141561550857826001600160a01b031663a9059cbb85846040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b1580156154d657600080fd5b505af11580156154ea573d6000803e3d6000fd5b505050506040513d602081101561550057600080fd5b50612ade9050565b826001600160a01b031663a9059cbb85846040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b15801561555f57600080fd5b505af1158015615573573d6000803e3d6000fd5b505050506040513d6020811015614a2957600080fd5b606061559884846000856159bd565b949350505050565b6000836001600160a01b0316856001600160a01b031611156155c0579293925b846001600160a01b0316866001600160a01b0316116155eb576155e4858585615b0d565b905061565b565b836001600160a01b0316866001600160a01b0316101561564d576000615612878686615b0d565b90506000615621878986615b70565b9050806001600160801b0316826001600160801b0316106156425780615644565b815b9250505061565b565b615658858584615b70565b90505b95945050505050565b80156125755760185460408051632e1a7d4d60e01b81526004810184905290516001600160a01b0390921691632e1a7d4d9160248082019260009290919082900301818387803b1580156156b757600080fd5b505af1158015613a01573d6000803e3d6000fd5b6000826001600160a01b0316846001600160a01b031611156156eb579192915b836001600160a01b0316615724606060ff16846001600160801b0316901b8686036001600160a01b0316866001600160a01b0316614103565b8161572b57fe5b04949350505050565b6000826001600160a01b0316846001600160a01b03161115615754579192915b615598826001600160801b03168585036001600160a01b0316600160601b614103565b6000816001600160a01b031663db006a75846040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b1580156157bf57600080fd5b505af11580156157d3573d6000803e3d6000fd5b505050506040513d60208110156157e957600080fd5b505190508015613dcd57fe5b4715613a0757601860009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0476040518263ffffffff1660e01b81526004016000604051808303818588803b1580156156b757600080fd5b60008080806001600160801b038516156158fd576006546040805163a34123a760e01b815260028a810b600483015289900b60248201526001600160801b038816604482015281516001600160a01b039093169263a34123a7926064808401939192918290030181600087803b1580156158c457600080fd5b505af11580156158d8573d6000803e3d6000fd5b505050506040513d60408110156158ee57600080fd5b50805160209091015190945092505b600654604080516309e3d67b60e31b815230600482015260028a810b602483015289900b60448201526001600160801b0360648201819052608482015281516001600160a01b0390931692634f1eb3d89260a4808401939192918290030181600087803b15801561596d57600080fd5b505af1158015615981573d6000803e3d6000fd5b505050506040513d604081101561599757600080fd5b50805160209091015194989397506001600160801b039081169650909316935090915050565b6060824710156159fe5760405162461bcd60e51b8152600401808060200182810382526026815260200180615d006026913960400191505060405180910390fd5b615a0785615bad565b615a58576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b600080866001600160a01b031685876040518082805190602001908083835b60208310615a965780518252601f199092019160209182019101615a77565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114615af8576040519150601f19603f3d011682016040523d82523d6000602084013e615afd565b606091505b5091509150611a52828286615bb3565b6000826001600160a01b0316846001600160a01b03161115615b2d579192915b6000615b50856001600160a01b0316856001600160a01b0316600160601b614103565b905061565b615b6b84838888036001600160a01b0316614103565b615c19565b6000826001600160a01b0316846001600160a01b03161115615b90579192915b615598615b6b83600160601b8787036001600160a01b0316614103565b3b151590565b60608315615bc2575081611147565b825115615bd25782518084602001fd5b60405162461bcd60e51b81526020600482018181528451602484015284518593919283926044019190850190808383600083156135c15781810151838201526020016135a9565b806001600160801b038116811461208057600080fdfe45524332303a207472616e7366657220746f20746865207a65726f20616464726573735265656e7472616e637947756172643a207265656e7472616e742063616c6c0045524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332303a20617070726f766520746f20746865207a65726f206164647265737345786368616e6765205261746520287363616c65642075702062792031653138293a2045524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e6365416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c756e69206c6971756964697479203d2049662074686973206973206e6f7420302c2074686572652077617320616e206572726f72536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f7745524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f20616464726573738d1cced004452bd270777a8c670f9f7e7c4fdde56f2db331fe289d39dc2624ad5361666545524332303a204552433230206f7065726174696f6e20646964206e6f74207375636365656445524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220f21ef9837a64343bda41113cb68ebf884d3c11ddb189ebecd664d04abb60192f64736f6c63430007060033"

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, _pool common.Address, _weth common.Address, _cToken0 common.Address, _cToken1 common.Address, _cEth common.Address, _protocolFee *big.Int, _maxTotalSupply *big.Int, _maxTwapDeviation *big.Int, _twapDuration uint32, _quoteAmount *big.Int, _uniPortion uint8, _team common.Address) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ApiBin), backend, _pool, _weth, _cToken0, _cToken1, _cEth, _protocolFee, _maxTotalSupply, _maxTwapDeviation, _twapDuration, _quoteAmount, _uniPortion, _team)
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

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Api *ApiCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Api *ApiSession) WETH() (common.Address, error) {
	return _Api.Contract.WETH(&_Api.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Api *ApiCallerSession) WETH() (common.Address, error) {
	return _Api.Contract.WETH(&_Api.CallOpts)
}

// AccId is a free data retrieval call binding the contract method 0xaabfd572.
//
// Solidity: function accId(address ) view returns(uint256)
func (_Api *ApiCaller) AccId(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "accId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccId is a free data retrieval call binding the contract method 0xaabfd572.
//
// Solidity: function accId(address ) view returns(uint256)
func (_Api *ApiSession) AccId(arg0 common.Address) (*big.Int, error) {
	return _Api.Contract.AccId(&_Api.CallOpts, arg0)
}

// AccId is a free data retrieval call binding the contract method 0xaabfd572.
//
// Solidity: function accId(address ) view returns(uint256)
func (_Api *ApiCallerSession) AccId(arg0 common.Address) (*big.Int, error) {
	return _Api.Contract.AccId(&_Api.CallOpts, arg0)
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

// QuoteAmount is a free data retrieval call binding the contract method 0x0d40886d.
//
// Solidity: function quoteAmount() view returns(uint128)
func (_Api *ApiCaller) QuoteAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "quoteAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteAmount is a free data retrieval call binding the contract method 0x0d40886d.
//
// Solidity: function quoteAmount() view returns(uint128)
func (_Api *ApiSession) QuoteAmount() (*big.Int, error) {
	return _Api.Contract.QuoteAmount(&_Api.CallOpts)
}

// QuoteAmount is a free data retrieval call binding the contract method 0x0d40886d.
//
// Solidity: function quoteAmount() view returns(uint128)
func (_Api *ApiCallerSession) QuoteAmount() (*big.Int, error) {
	return _Api.Contract.QuoteAmount(&_Api.CallOpts)
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
// Solidity: function deposit(uint256 amountToken0, uint256 amountToken1, bool doRebalance) returns()
func (_Api *ApiTransactor) Deposit(opts *bind.TransactOpts, amountToken0 *big.Int, amountToken1 *big.Int, doRebalance bool) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "deposit", amountToken0, amountToken1, doRebalance)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 amountToken0, uint256 amountToken1, bool doRebalance) returns()
func (_Api *ApiSession) Deposit(amountToken0 *big.Int, amountToken1 *big.Int, doRebalance bool) (*types.Transaction, error) {
	return _Api.Contract.Deposit(&_Api.TransactOpts, amountToken0, amountToken1, doRebalance)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 amountToken0, uint256 amountToken1, bool doRebalance) returns()
func (_Api *ApiTransactorSession) Deposit(amountToken0 *big.Int, amountToken1 *big.Int, doRebalance bool) (*types.Transaction, error) {
	return _Api.Contract.Deposit(&_Api.TransactOpts, amountToken0, amountToken1, doRebalance)
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

// Sweep is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address token, uint256 amount) returns()
func (_Api *ApiTransactor) Sweep(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "sweep", token, amount)
}

// Sweep is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address token, uint256 amount) returns()
func (_Api *ApiSession) Sweep(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Sweep(&_Api.TransactOpts, token, amount)
}

// Sweep is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address token, uint256 amount) returns()
func (_Api *ApiTransactorSession) Sweep(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Sweep(&_Api.TransactOpts, token, amount)
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

// Withdraw is a paid mutator transaction binding the contract method 0xc6ab5d90.
//
// Solidity: function withdraw(uint8 percent) returns(uint256 amount0, uint256 amount1)
func (_Api *ApiTransactor) Withdraw(opts *bind.TransactOpts, percent uint8) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "withdraw", percent)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc6ab5d90.
//
// Solidity: function withdraw(uint8 percent) returns(uint256 amount0, uint256 amount1)
func (_Api *ApiSession) Withdraw(percent uint8) (*types.Transaction, error) {
	return _Api.Contract.Withdraw(&_Api.TransactOpts, percent)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc6ab5d90.
//
// Solidity: function withdraw(uint8 percent) returns(uint256 amount0, uint256 amount1)
func (_Api *ApiTransactorSession) Withdraw(percent uint8) (*types.Transaction, error) {
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
	Shares  *big.Int
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x36af321ec8d3c75236829c5317affd40ddb308863a1236d2d277a4025cccee1e.
//
// Solidity: event Deposit(address indexed sender, uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*ApiDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &ApiDepositIterator{contract: _Api.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x36af321ec8d3c75236829c5317affd40ddb308863a1236d2d277a4025cccee1e.
//
// Solidity: event Deposit(address indexed sender, uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ApiDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Deposit", senderRule)
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

// ParseDeposit is a log parse operation binding the contract event 0x36af321ec8d3c75236829c5317affd40ddb308863a1236d2d277a4025cccee1e.
//
// Solidity: event Deposit(address indexed sender, uint256 shares, uint256 amount0, uint256 amount1)
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

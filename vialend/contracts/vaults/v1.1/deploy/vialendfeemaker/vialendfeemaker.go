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
const ApiABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cToken0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cToken1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cEth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_protocolFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxTotalSupply\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"_maxTwapDeviation\",\"type\":\"int24\"},{\"internalType\":\"uint32\",\"name\":\"_twapDuration\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_uniPortionRate\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesToVault0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesToVault1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesToProtocol0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesToProtocol1\",\"type\":\"uint256\"}],\"name\":\"CollectFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"ErrorLogging\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"subject\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"GeneralA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"subject\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"GeneralB\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"subject\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"GeneralS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MyLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MyLog2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance1\",\"type\":\"uint256\"}],\"name\":\"RebalanceLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"Snapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolamount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolamount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lendingamount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lendingamount1\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"AccumulateUniswapFees0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AccumulateUniswapFees1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Assetholder\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"capital0\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"capital1\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"fees0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fees1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"Amount\",\"type\":\"uint256\"}],\"name\":\"_unwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_wrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedProtocolFees0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedProtocolFees1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cHigh\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cLow\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"collectProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToken1\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"getCapital\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getETHBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"getPositionAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"getSSLiquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTwap\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUniswapPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRebalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTick\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTwapDeviation\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_cErc20Contract\",\"type\":\"address\"}],\"name\":\"redeemCErc20Tokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_cEtherContract\",\"type\":\"address\"}],\"name\":\"redeemCEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"setGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMax\",\"type\":\"uint256\"}],\"name\":\"setMaxTotalSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_maxTwapDeviation\",\"type\":\"int24\"}],\"name\":\"setMaxTwapDeviation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_team\",\"type\":\"address\"}],\"name\":\"setTeam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_twapDuration\",\"type\":\"uint32\"}],\"name\":\"setTwapDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"ratio\",\"type\":\"uint8\"}],\"name\":\"setUniPortionRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"newLow\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"newHigh\",\"type\":\"int24\"},{\"internalType\":\"int256\",\"name\":\"swapAmount\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"}],\"name\":\"strategy0\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"newLow\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"newHigh\",\"type\":\"int24\"}],\"name\":\"strategy1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20Contract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cErc20Contract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_numTokensToSupply\",\"type\":\"uint256\"}],\"name\":\"supplyErc20ToCompound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_cEtherContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ctoken\",\"type\":\"address\"}],\"name\":\"supplyEthToCompound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"swapAmount\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"team\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"twapDuration\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniPortionRate\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ApiBin is the compiled bytecode used for deploying new contracts.
var ApiBin = "0x60c06040523480156200001157600080fd5b50604051620063423803806200634283398181016040526101408110156200003857600080fd5b5080516020808301516040808501516060860151608087015160a088015160c089015160e08a01516101008b0151610120909b015187518089018952600e81526d05669614c656e6420546f6b656e360941b818c019081528951808b01909a5260048a52630564c54360e41b9b8a019b909b52600160005580519b9c999b979a969995989497939692959491939092620000d6916016919062000491565b508051620000ec90601790602084019062000491565b50506018805460ff1916601217905550600180546001600160a01b031990811633908117909255600280548216909217909155600b80546001600160a01b038d1692168217905560408051630dfe168160e01b81529051630dfe168191600480820192602092909190829003018186803b1580156200016a57600080fd5b505afa1580156200017f573d6000803e3d6000fd5b505050506040513d60208110156200019657600080fd5b505160601b6001600160601b0319166080526040805163d21220a760e01b815290516001600160a01b038c169163d21220a7916004808301926020929190829003018186803b158015620001e957600080fd5b505afa158015620001fe573d6000803e3d6000fd5b505050506040513d60208110156200021557600080fd5b505160601b6001600160601b03191660a05260078054600160401b600160e01b031916680100000000000000006001600160a01b038b81169190910291909117909155600880546001600160a01b03199081168a841617909155600980549091168883161790558916620002b9576040805162461bcd60e51b815260206004808301919091526024820152630ae8aa8960e31b604482015290519081900360640190fd5b88600a60006101000a8154816001600160a01b0302191690836001600160a01b0316021790555084600c81905550896001600160a01b031663d0c93a7c6040518163ffffffff1660e01b815260040160206040518083038186803b1580156200032157600080fd5b505afa15801562000336573d6000803e3d6000fd5b505050506040513d60208110156200034d57600080fd5b5051600b805460029290920b62ffffff16600160a01b0262ffffff60a01b19909216919091179055620f42408510620003b3576040805162461bcd60e51b815260206004820152600360248201526228232960e91b604482015290519081900360640190fd5b60068490556007805460ff83166701000000000000000260ff60381b1963ffffffff861663ffffffff19600289900b62ffffff81166401000000000262ffffff60201b19909616959095171617161790915560001262000440576040805162461bcd60e51b815260206004820152600360248201526213551160ea1b604482015290519081900360640190fd5b60008263ffffffff161162000481576040805162461bcd60e51b8152602060048201526002602482015261151160f21b604482015290519081900360640190fd5b505050505050505050506200053d565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282620004c9576000855562000514565b82601f10620004e457805160ff191683800117855562000514565b8280016001018555821562000514579182015b8281111562000514578251825591602001919060010190620004f7565b506200052292915062000526565b5090565b5b8082111562000522576000815560010162000527565b60805160601c60a05160601c615d3b6200060760003980610fd452806116805280611f1252806125a452806126bf5280612ee85280612f9b52806130ee52806132b6528061342952806134a852806136aa52806138555280613ece52806142d552806143c852806150e0525080610f8a52806112355280611626528061250f528061269e52806127535280612f6152806130565280613222528061335852806133d7528061366f52806138185280613ead52806142b4528061430b52806150a05250615d3b6000f3fe6080604052600436106103bc5760003560e01c80636253c718116101f2578063aa489e061161010d578063d8d7f96f116100a0578063e7c7cb911161006f578063e7c7cb9114610de6578063eae989a214610dfb578063f2a40db814610e10578063fa461e3314610e3a576103c3565b8063d8d7f96f14610d15578063dd62ed3e14610d2a578063e1c220bf14610d65578063e2bbb15814610d98576103c3565b8063cda206b0116100dc578063cda206b014610c11578063d21220a714610c64578063d348799714610c79578063d368867f14610d00576103c3565b8063aa489e0614610b84578063ab033ea914610b99578063c29188bc14610bcc578063c433c80a14610be1576103c3565b806385f2aef211610185578063a457c2d711610154578063a457c2d714610adb578063a5541f2814610a4e578063a9059cbb14610b14578063a91ef6eb14610b4d576103c3565b806385f2aef214610a875780638e843c6c14610a9c57806395d89b4114610ab1578063a00fa77f14610ac6576103c3565b806370a08231116101c157806370a08231146109d857806373232c6014610a0b5780637c23b92114610a3957806384fea66714610a4e576103c3565b80636253c71814610936578063629d94051461096d57806368a5787a146109825780636e947298146109c3576103c3565b80632e1a7d4d116102e25780633f3e4c11116102755780635aa6e675116102445780635aa6e675146108ac5780635bb6aa85146108c15780635d752a9a146108f4578063624a177a14610909576103c3565b80633f3e4c111461082e57806341aec53814610858578063457077cc1461086d57806358f8588014610897576103c3565b80633b7ba25f116102b15780633b7ba25f1461074c5780633cbff3fe1461078f5780633dfa5d87146107bc5780633e5aae66146107e8576103c3565b80632e1a7d4d146106995780632e4b8f60146106c3578063313ce567146106fe5780633950935114610713576103c3565b8063106b9ca11161035a57806323b872dd1161032957806323b872dd146105fe57806326d89545146106415780632ab4d0521461066f5780632b6e192314610684576103c3565b8063106b9ca1146105985780631755ff21146105bf57806318160ddd146105d4578063238efcbc146105e9576103c3565b806306fdde031161039657806306fdde031461045d578063095cf5c6146104e7578063095ea7b31461051a5780630dfe168114610567576103c3565b806302cfafc8146103c55780630430c130146103f057806306b0b6561461042f576103c3565b366103c357005b005b3480156103d157600080fd5b506103da610ec1565b6040805160ff9092168252519081900360200190f35b3480156103fc57600080fd5b506103c36004803603606081101561041357600080fd5b50803590602081013590604001356001600160a01b0316610ed1565b34801561043b57600080fd5b50610444611000565b6040805192835260208301919091528051918290030190f35b34801561046957600080fd5b50610472611111565b6040805160208082528351818301528351919283929083019185019080838360005b838110156104ac578181015183820152602001610494565b50505050905090810190601f1680156104d95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156104f357600080fd5b506103c36004803603602081101561050a57600080fd5b50356001600160a01b03166111a7565b34801561052657600080fd5b506105536004803603604081101561053d57600080fd5b506001600160a01b038135169060200135611215565b604080519115158252519081900360200190f35b34801561057357600080fd5b5061057c611233565b604080516001600160a01b039092168252519081900360200190f35b3480156105a457600080fd5b506105ad611257565b60408051918252519081900360200190f35b3480156105cb57600080fd5b5061057c61125d565b3480156105e057600080fd5b506105ad61126c565b3480156105f557600080fd5b506103c3611272565b34801561060a57600080fd5b506105536004803603606081101561062157600080fd5b506001600160a01b038135811691602081013590911690604001356112d9565b34801561064d57600080fd5b50610656611361565b6040805163ffffffff9092168252519081900360200190f35b34801561067b57600080fd5b506105ad61136d565b34801561069057600080fd5b506105ad611373565b3480156106a557600080fd5b50610444600480360360208110156106bc57600080fd5b5035611421565b3480156106cf57600080fd5b506106d8611747565b604080519485526020850193909352838301919091526060830152519081900360800190f35b34801561070a57600080fd5b506103da611966565b34801561071f57600080fd5b506105536004803603604081101561073657600080fd5b506001600160a01b03813516906020013561196f565b34801561075857600080fd5b506105ad6004803603606081101561076f57600080fd5b506001600160a01b038135811691602081013590911690604001356119bd565b34801561079b57600080fd5b506103c3600480360360208110156107b257600080fd5b503560020b611d0a565b3480156107c857600080fd5b506107d1611dc6565b6040805160029290920b8252519081900360200190f35b3480156107f457600080fd5b506103c36004803603608081101561080b57600080fd5b508035600290810b91602081013590910b90604081013590606001351515611dcf565b34801561083a57600080fd5b506103c36004803603602081101561085157600080fd5b5035611e6d565b34801561086457600080fd5b506105ad611f0e565b34801561087957600080fd5b506103c36004803603602081101561089057600080fd5b5035611fae565b3480156108a357600080fd5b506105ad61201d565b3480156108b857600080fd5b5061057c612023565b3480156108cd57600080fd5b506106d8600480360360208110156108e457600080fd5b50356001600160a01b0316612032565b34801561090057600080fd5b506107d1612059565b34801561091557600080fd5b506103c36004803603602081101561092c57600080fd5b503560ff166122ed565b34801561094257600080fd5b506103c36004803603604081101561095957600080fd5b508035600290810b9160200135900b61239a565b34801561097957600080fd5b506105ad61274f565b34801561098e57600080fd5b50610444600480360360608110156109a557600080fd5b508035906020810135151590604001356001600160a01b03166127be565b3480156109cf57600080fd5b506105ad6128f3565b3480156109e457600080fd5b506105ad600480360360208110156109fb57600080fd5b50356001600160a01b03166128f7565b61055360048036036040811015610a2157600080fd5b506001600160a01b0381358116916020013516612916565b348015610a4557600080fd5b506103c3612af6565b348015610a5a57600080fd5b506103c360048036036040811015610a7157600080fd5b50803590602001356001600160a01b0316612b4e565b348015610a9357600080fd5b5061057c612c0d565b348015610aa857600080fd5b506107d1612c1c565b348015610abd57600080fd5b50610472612c2c565b348015610ad257600080fd5b506105ad612c8d565b348015610ae757600080fd5b5061055360048036036040811015610afe57600080fd5b506001600160a01b038135169060200135612c93565b348015610b2057600080fd5b5061055360048036036040811015610b3757600080fd5b506001600160a01b038135169060200135612cfb565b348015610b5957600080fd5b5061044460048036036040811015610b7057600080fd5b508035600290810b9160200135900b612d0f565b348015610b9057600080fd5b506105ad612da6565b348015610ba557600080fd5b506103c360048036036020811015610bbc57600080fd5b50356001600160a01b0316612dac565b348015610bd857600080fd5b506105ad612e1a565b348015610bed57600080fd5b506103c360048036036020811015610c0457600080fd5b503563ffffffff16612e20565b348015610c1d57600080fd5b50610c4860048036036040811015610c3457600080fd5b508035600290810b9160200135900b612ece565b604080516001600160801b039092168252519081900360200190f35b348015610c7057600080fd5b5061057c612ee6565b348015610c8557600080fd5b506103c360048036036060811015610c9c57600080fd5b813591602081013591810190606081016040820135600160201b811115610cc257600080fd5b820183602082011115610cd457600080fd5b803590602001918460018302840111600160201b83111715610cf557600080fd5b509092509050612f0a565b348015610d0c57600080fd5b506107d1612fc8565b348015610d2157600080fd5b506103c3612fd8565b348015610d3657600080fd5b506105ad60048036036040811015610d4d57600080fd5b506001600160a01b03813581169160200135166134cf565b348015610d7157600080fd5b506106d860048036036020811015610d8857600080fd5b50356001600160a01b03166134fa565b348015610da457600080fd5b50610dc860048036036040811015610dbb57600080fd5b508035906020013561352b565b60408051938452602084019290925282820152519081900360600190f35b348015610df257600080fd5b506107d161377d565b348015610e0757600080fd5b506105ad61378d565b348015610e1c57600080fd5b5061057c60048036036020811015610e3357600080fd5b5035613793565b348015610e4657600080fd5b506103c360048036036060811015610e5d57600080fd5b813591602081013591810190606081016040820135600160201b811115610e8357600080fd5b820183602082011115610e9557600080fd5b803590602001918460018302840111600160201b83111715610eb657600080fd5b5090925090506137bd565b600754600160381b900460ff1681565b6001546001600160a01b03163314610f1d576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b82600d5410158015610f31575081600e5410155b610f67576040805162461bcd60e51b8152602060048201526002602482015261043560f41b604482015290519081900360640190fd5b8215610fb157600d54610f7a908461387c565b600d55610fb16001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001682856138d9565b8115610ffb57600e54610fc4908361387c565b600e55610ffb6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001682846138d9565b505050565b600080600760089054906101000a90046001600160a01b03166001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561106557600080fd5b505afa158015611079573d6000803e3d6000fd5b505050506040513d602081101561108f57600080fd5b5051600854604080516370a0823160e01b815230600482015290519294506001600160a01b03909116916370a0823191602480820192602092909190829003018186803b1580156110df57600080fd5b505afa1580156110f3573d6000803e3d6000fd5b505050506040513d602081101561110957600080fd5b505191929050565b60168054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561119d5780601f106111725761010080835404028352916020019161119d565b820191906000526020600020905b81548152906001019060200180831161118057829003601f168201915b5050505050905090565b6001546001600160a01b031633146111f3576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b600280546001600160a01b0319166001600160a01b0392909216919091179055565b600061122961122261392b565b848461392f565b5060015b92915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60115481565b600b546001600160a01b031690565b60155490565b6003546001600160a01b031633146112c5576040805162461bcd60e51b815260206004820152601160248201527070656e64696e67476f7665726e616e636560781b604482015290519081900360640190fd5b600180546001600160a01b03191633179055565b60006112e6848484613a1b565b611356846112f261392b565b61135185604051806060016040528060288152602001615c05602891396001600160a01b038a1660009081526014602052604081209061133061392b565b6001600160a01b031681526020810191909152604001600020549190613b78565b61392f565b5060015b9392505050565b60075463ffffffff1681565b60065481565b600080600b60009054906101000a90046001600160a01b03166001600160a01b0316633850c7bd6040518163ffffffff1660e01b815260040160e06040518083038186803b1580156113c457600080fd5b505afa1580156113d8573d6000803e3d6000fd5b505050506040513d60e08110156113ee57600080fd5b5051905060c0611419670de0b6b3a76400006114136001600160a01b03851680613c0f565b90613c0f565b901c91505090565b6000806002600054141561146a576040805162461bcd60e51b815260206004820152601f6024820152600080516020615aed833981519152604482015290519081900360640190fd5b600260005533801580159061148857506001600160a01b0381163014155b6114bf576040805162461bcd60e51b81526020600482015260036024820152623a379960e91b604482015290519081900360640190fd5b6000841180156114d0575060648411155b611509576040805162461bcd60e51b81526020600482015260056024820152641c18d95b9d60da1b604482015290519081900360640190fd5b6000611523606461151d87611413866128f7565b90613c68565b9050600061152f61126c565b90506000811161156c576040805162461bcd60e51b815260206004820152600360248201526207473360ec1b604482015290519081900360640190fd5b808211156115a7576040805162461bcd60e51b815260206004820152600360248201526273683160e81b604482015290519081900360640190fd5b6115b18383613ccf565b600b5460009081906115da90600160b81b8104600290810b91600160d01b9004900b8686613dcb565b915091506000806115eb8686613e71565b915091506115ff8561151d8861141361274f565b98506116118561151d88611413611f0e565b9750881561166d5761164d6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016888b6138d9565b6001600160a01b038716600090815260056020526040902080548a900390555b87156116ca576116a76001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016888a6138d9565b6001600160a01b0387166000908152600560205260409020600101805489900390555b60408051878152602081018b90528082018a9052606081018690526080810185905260a0810184905260c0810183905290516001600160a01b0389169133917f9465fbc710f5c7ebb85aec935f56a888087125eaf287d09f2d3b59c640b3eafd9181900360e00190a3505050505050506001600081905550915091565b600080600080600080611758611000565b9150915060006012905060006006905060006117fb8363ffffffff16600a0a61151d600760089054906101000a90046001600160a01b03166001600160a01b031663182df0f56040518163ffffffff1660e01b815260040160206040518083038186803b1580156117c857600080fd5b505afa1580156117dc573d6000803e3d6000fd5b505050506040513d60208110156117f257600080fd5b50518890613c0f565b905060006118636008846012010363ffffffff16600a0a61151d600860009054906101000a90046001600160a01b03166001600160a01b031663182df0f56040518163ffffffff1660e01b815260040160206040518083038186803b1580156117c857600080fd5b90508181600760089054906101000a90046001600160a01b03166001600160a01b031663182df0f56040518163ffffffff1660e01b815260040160206040518083038186803b1580156118b557600080fd5b505afa1580156118c9573d6000803e3d6000fd5b505050506040513d60208110156118df57600080fd5b50516008546040805163182df0f560e01b815290516001600160a01b039092169163182df0f591600480820192602092909190829003018186803b15801561192657600080fd5b505afa15801561193a573d6000803e3d6000fd5b505050506040513d602081101561195057600080fd5b5051929d919c509a509098509650505050505050565b60185460ff1690565b600061122961197c61392b565b84611351856014600061198d61392b565b6001600160a01b03908116825260208083019390935260409182016000908120918c168152925290205490614058565b6000836001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015611a0c57600080fd5b505afa158015611a20573d6000803e3d6000fd5b505050506040513d6020811015611a3657600080fd5b5051821115611a76576040805162461bcd60e51b815260206004820152600760248201526662616c616e636560c81b604482015290519081900360640190fd5b600084905060008490506000816001600160a01b031663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015611abd57600080fd5b505af1158015611ad1573d6000803e3d6000fd5b505050506040513d6020811015611ae757600080fd5b50516040805160208101839052818152601b818301527f45786368616e6765205261746520287363616c6564207570293a20000000000060608201529051919250600080516020615c97833981519152919081900360800190a16000826001600160a01b031663ae9d70b06040518163ffffffff1660e01b8152600401602060405180830381600087803b158015611b7e57600080fd5b505af1158015611b92573d6000803e3d6000fd5b505050506040513d6020811015611ba857600080fd5b505160408051602081018390528181526018818301527f537570706c7920526174653a20287363616c656420757029000000000000000060608201529051919250600080516020615c97833981519152919081900360800190a1836001600160a01b031663095ea7b388886040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b158015611c5957600080fd5b505af1158015611c6d573d6000803e3d6000fd5b505050506040513d6020811015611c8357600080fd5b50506040805163140e25ad60e31b81526004810188905290516000916001600160a01b0386169163a0712d689160248082019260209290919082900301818787803b158015611cd157600080fd5b505af1158015611ce5573d6000803e3d6000fd5b505050506040513d6020811015611cfb57600080fd5b50519998505050505050505050565b6002546001600160a01b03163314611d52576040805162461bcd60e51b815260206004808301919091526024820152637465616d60e01b604482015290519081900360640190fd5b60008160020b13611d9d576040805162461bcd60e51b815260206004820152601060248201526f36b0bc2a3bb0b82232bb34b0ba34b7b760811b604482015290519081900360640190fd5b6007805460029290920b62ffffff16600160201b0266ffffff0000000019909216919091179055565b60125460020b81565b60026000541415611e15576040805162461bcd60e51b815260206004820152601f6024820152600080516020615aed833981519152604482015290519081900360640190fd5b60026000819055546001600160a01b03163314611e62576040805162461bcd60e51b815260206004808301919091526024820152637465616d60e01b604482015290519081900360640190fd5b505060016000555050565b60026000541415611eb3576040805162461bcd60e51b815260206004820152601f6024820152600080516020615aed833981519152604482015290519081900360640190fd5b60026000556001546001600160a01b03163314611f04576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b6006556001600055565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015611f7d57600080fd5b505afa158015611f91573d6000803e3d6000fd5b505050506040513d6020811015611fa757600080fd5b5051905090565b801561201a57600a5460408051632e1a7d4d60e01b81526004810184905290516001600160a01b0390921691632e1a7d4d9160248082019260009290919082900301818387803b15801561200157600080fd5b505af1158015612015573d6000803e3d6000fd5b505050505b50565b600c5481565b6001546001600160a01b031681565b60056020526000908152604090208054600182015460028301546003909301549192909184565b60075460408051600280825260608201835260009363ffffffff169284929190602083019080368337019050509050818160008151811061209657fe5b602002602001019063ffffffff16908163ffffffff16815250506000816001815181106120bf57fe5b63ffffffff909216602092830291909101820152600b5460405163883bdbfd60e01b8152600481018381528451602483015284516000946001600160a01b039094169363883bdbfd93879392839260449092019185810191028083838b5b8381101561213557818101518382015260200161211d565b505050509050019250505060006040518083038186803b15801561215857600080fd5b505afa15801561216c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604090815281101561219557600080fd5b8101908080516040519392919084600160201b8211156121b457600080fd5b9083019060208201858111156121c957600080fd5b82518660208202830111600160201b821117156121e557600080fd5b82525081516020918201928201910280838360005b838110156122125781810151838201526020016121fa565b5050505090500160405260200180516040519392919084600160201b82111561223a57600080fd5b90830190602082018581111561224f57600080fd5b82518660208202830111600160201b8211171561226b57600080fd5b82525081516020918201928201910280838360005b83811015612298578181015183820152602001612280565b505050509050016040525050505090508263ffffffff16816000815181106122bc57fe5b6020026020010151826001815181106122d157fe5b60200260200101510360060b816122e457fe5b05935050505090565b6002546001600160a01b03163314612335576040805162461bcd60e51b815260206004808301919091526024820152637465616d60e01b604482015290519081900360640190fd5b60648160ff161115612376576040805162461bcd60e51b8152602060048201526005602482015264726174696f60d81b604482015290519081900360640190fd5b6007805460ff909216600160381b0267ff0000000000000019909216919091179055565b600260005414156123e0576040805162461bcd60e51b815260206004820152601f6024820152600080516020615aed833981519152604482015290519081900360640190fd5b60026000819055546001600160a01b0316331461242d576040805162461bcd60e51b815260206004808301919091526024820152637465616d60e01b604482015290519081900360640190fd5b600061243761126c565b1161246e576040805162461bcd60e51b8152602060048201526002602482015261074360f41b604482015290519081900360640190fd5b600b5460408051633850c7bd60e01b815290516000926001600160a01b031691633850c7bd9160048083019260e0929190829003018186803b1580156124b357600080fd5b505afa1580156124c7573d6000803e3d6000fd5b505050506040513d60e08110156124dd57600080fd5b506020015190506124ef8383836140b2565b604080516370a0823160e01b815230600482015290516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016916370a08231916024808301926020929190829003018186803b15801561255557600080fd5b505afa158015612569573d6000803e3d6000fd5b505050506040513d602081101561257f57600080fd5b5051601955604080516370a0823160e01b815230600482015290516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016916370a08231916024808301926020929190829003018186803b1580156125ea57600080fd5b505afa1580156125fe573d6000803e3d6000fd5b505050506040513d602081101561261457600080fd5b5051601a556126216142a0565b612629614307565b600061264c606461151d60078054906101000a900460ff1660ff1661141361274f565b90506000612671606461151d60078054906101000a900460ff1660ff16611413611f0e565b905061267f858584846146cd565b600061268961274f565b90506000612695611f0e565b905060006126e57f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000085856147c8565b90508061271f576040805162461bcd60e51b815260206004820152600360248201526226223360e91b604482015290519081900360640190fd5b50504260115550506012805460029390930b62ffffff1662ffffff19909316929092179091555050600160005550565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015611f7d57600080fd5b600b54604080513360208083019190915282518083038201815282840193849052630251596160e31b90935230604483018181528715156064850152608484018990526001600160a01b0387811660a486015260a060c48601908152865160e48701528651600098899893169663128acb08968c958e958d9590949193909261010490910191908501908083838f5b8381101561286557818101518382015260200161284d565b50505050905090810190601f1680156128925780820380516001836020036101000a031916815260200191505b5096505050505050506040805180830381600087803b1580156128b457600080fd5b505af11580156128c8573d6000803e3d6000fd5b505050506040513d60408110156128de57600080fd5b50805160209091015190969095509350505050565b4790565b6001600160a01b0381166000908152601360205260409020545b919050565b60008083905060008390506000816001600160a01b031663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561295e57600080fd5b505af1158015612972573d6000803e3d6000fd5b505050506040513d602081101561298857600080fd5b505160408051602081018390528181526023918101829052919250600080516020615c9783398151915291839181906060820190615b7582396040019250505060405180910390a16000826001600160a01b031663ae9d70b06040518163ffffffff1660e01b8152600401602060405180830381600087803b158015612a0d57600080fd5b505af1158015612a21573d6000803e3d6000fd5b505050506040513d6020811015612a3757600080fd5b5051604080516020818101849052828252818301527f537570706c7920526174653a20287363616c656420757020627920316531382960608201529051919250600080516020615c97833981519152919081900360800190a1836001600160a01b0316631249c58b6203d090476040518363ffffffff1660e01b81526004016000604051808303818589803b158015612acf57600080fd5b5088f1158015612ae3573d6000803e3d6000fd5b5060019c9b505050505050505050505050565b4715612b4c57600a60009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0476040518263ffffffff1660e01b81526004016000604051808303818588803b15801561200157600080fd5b565b6000816001600160a01b031663db006a75846040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b158015612b9657600080fd5b505af1158015612baa573d6000803e3d6000fd5b505050506040513d6020811015612bc057600080fd5b505160408051602081018390528181526024918101829052919250600080516020615c9783398151915291839181906060820190615b5182396040019250505060405180910390a1505050565b6002546001600160a01b031681565b600b54600160d01b900460020b81565b60178054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561119d5780601f106111725761010080835404028352916020019161119d565b600e5481565b6000611229612ca061392b565b8461135185604051806060016040528060258152602001615ce16025913960146000612cca61392b565b6001600160a01b03908116825260208083019390935260409182016000908120918d16815292529020549190613b78565b6000611229612d0861392b565b8484613a1b565b6000806000806000612d218787614893565b9450945050509250612d3487878561494b565b600c549196509450600090612d4d90620f42409061387c565b9050612d73612d6c620f424061151d6001600160801b03871685613c0f565b8790614058565b9550612d99612d92620f424061151d6001600160801b03861685613c0f565b8690614058565b9450505050509250929050565b600f5481565b6001546001600160a01b03163314612df8576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b600380546001600160a01b0319166001600160a01b0392909216919091179055565b60105481565b6002546001600160a01b03163314612e68576040805162461bcd60e51b815260206004808301919091526024820152637465616d60e01b604482015290519081900360640190fd5b60008163ffffffff1611612eb2576040805162461bcd60e51b815260206004820152600c60248201526b3a3bb0b8223ab930ba34b7b760a11b604482015290519081900360640190fd5b6007805463ffffffff191663ffffffff92909216919091179055565b6000612eda8383614893565b50929695505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b600b546001600160a01b03163314612f4e576040805162461bcd60e51b8152602060048201526002602482015261505360f01b604482015290519081900360640190fd5b8315612f8857612f886001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001633866138d9565b8215612fc257612fc26001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001633856138d9565b50505050565b600b54600160b81b900460020b81565b6001546001600160a01b03163314613024576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b61302c6142a0565b60048054604080516370a0823160e01b815230938101939093525190916000916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016916370a08231916024808301926020929190829003018186803b15801561309c57600080fd5b505afa1580156130b0573d6000803e3d6000fd5b505050506040513d60208110156130c657600080fd5b5051604080516370a0823160e01b815230600482015290519192506000916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016916370a08231916024808301926020929190829003018186803b15801561313457600080fd5b505afa158015613148573d6000803e3d6000fd5b505050506040513d602081101561315e57600080fd5b50519050600061316c61126c565b905060005b8481101561332c5760006131a56004838154811061318b57fe5b6000918252602090912001546001600160a01b03166128f7565b90506131d2600483815481106131b757fe5b6000918252602090912001546001600160a01b031682613ccf565b60006131e28461151d8885613c0f565b905060006131f48561151d8886613c0f565b9050811561328a5761324b6004858154811061320c57fe5b6000918252602090912001546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169116846138d9565b81600560006004878154811061325d57fe5b60009182526020808320909101546001600160a01b03168352820192909252604001902080549190910390555b8015613321576132df600485815481106132a057fe5b6000918252602090912001546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169116836138d9565b8060056000600487815481106132f157fe5b60009182526020808320909101546001600160a01b03168352820192909252604001902060010180549190910390555b505050600101613171565b50600154604080516370a0823160e01b815230600482015290516133fe926001600160a01b03908116927f0000000000000000000000000000000000000000000000000000000000000000909116916370a0823191602480820192602092909190829003018186803b1580156133a157600080fd5b505afa1580156133b5573d6000803e3d6000fd5b505050506040513d60208110156133cb57600080fd5b50516001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691906138d9565b600154604080516370a0823160e01b81523060048201529051612fc2926001600160a01b03908116927f0000000000000000000000000000000000000000000000000000000000000000909116916370a0823191602480820192602092909190829003018186803b15801561347257600080fd5b505afa158015613486573d6000803e3d6000fd5b505050506040513d602081101561349c57600080fd5b50516001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691906138d9565b6001600160a01b03918216600090815260146020908152604080832093909416825291909152205490565b6001600160a01b03166000908152600560205260409020805460018201546002830154600390930154919390929190565b600080600060026000541415613576576040805162461bcd60e51b815260206004820152601f6024820152600080516020615aed833981519152604482015290519081900360640190fd5b6002600055338515158061358a5750600085115b6135c4576040805162461bcd60e51b815260206004808301919091526024820152630616d74360e41b604482015290519081900360640190fd5b6001600160a01b038116158015906135e557506001600160a01b0381163014155b61361b576040805162461bcd60e51b8152602060048201526002602482015261746f60f01b604482015290519081900360640190fd5b6001600160a01b038116600090815260056020526040902080548701815560010180548601905561364b816149f5565b6136558686614aa0565b919550935091508215613697576136976001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016333086614abf565b81156136d2576136d26001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016333085614abf565b6136dc8185614b19565b604080518581526020810185905280820184905290516001600160a01b0383169133917f4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f69181900360600190a360065461373461126c565b111561376d576040805162461bcd60e51b815260206004820152600360248201526204341560ec1b604482015290519081900360640190fd5b5060016000819055509250925092565b600754600160201b900460020b81565b600d5481565b600481815481106137a357600080fd5b6000918252602090912001546001600160a01b0316905081565b600b546001600160a01b03163314613802576040805162461bcd60e51b815260206004820152600360248201526228299960e91b604482015290519081900360640190fd5b600084131561383f5761383f6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001633866138d9565b6000831315612fc257612fc26001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001633856138d9565b6000828211156138d3576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052610ffb908490614c0b565b3390565b6001600160a01b0383166139745760405162461bcd60e51b8152600401808060200182810382526024815260200180615c736024913960400191505060405180910390fd5b6001600160a01b0382166139b95760405162461bcd60e51b8152600401808060200182810382526022815260200180615b2f6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260146020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b038316613a605760405162461bcd60e51b8152600401808060200182810382526025815260200180615c4e6025913960400191505060405180910390fd5b6001600160a01b038216613aa55760405162461bcd60e51b8152600401808060200182810382526023815260200180615aca6023913960400191505060405180910390fd5b613ab0838383610ffb565b613aed81604051806060016040528060268152602001615b98602691396001600160a01b0386166000908152601360205260409020549190613b78565b6001600160a01b038085166000908152601360205260408082209390935590841681522054613b1c9082614058565b6001600160a01b0380841660008181526013602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008184841115613c075760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015613bcc578181015183820152602001613bb4565b50505050905090810190601f168015613bf95780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600082613c1e5750600061122d565b82820282848281613c2b57fe5b041461135a5760405162461bcd60e51b8152600401808060200182810382526021815260200180615be46021913960400191505060405180910390fd5b6000808211613cbe576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b818381613cc757fe5b049392505050565b6001600160a01b038216613d145760405162461bcd60e51b8152600401808060200182810382526021815260200180615c2d6021913960400191505060405180910390fd5b613d2082600083610ffb565b613d5d81604051806060016040528060228152602001615b0d602291396001600160a01b0385166000908152601360205260409020549190613b78565b6001600160a01b038316600090815260136020526040902055601554613d83908261387c565b6015556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b6000806000613dda8787614893565b5050505090506000613e028561151d88856001600160801b0316613c0f90919063ffffffff16565b90508015613e6657600080600080613e238c8c613e1e88614cbc565b614cd3565b92965090945092509050613e45613e3e8a61151d858e613c0f565b8590614058565b9750613e5f613e588a61151d848e613c0f565b8490614058565b9650505050505b505094509492505050565b600080600080613e7f611000565b90925090506000613e948661151d858a613c0f565b90506000613ea68761151d858b613c0f565b9050613ef47f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000008484614f26565b6000601290506000600690506000613f768363ffffffff16600a0a61151d600760089054906101000a90046001600160a01b03166001600160a01b031663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015613f6257600080fd5b505af11580156117dc573d6000803e3d6000fd5b90506000613fe06008846012010363ffffffff16600a0a61151d600860009054906101000a90046001600160a01b03166001600160a01b031663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015613f6257600080fd5b604080516020810185905280820183905260608082526012908201527102fb13ab9372632b73234b733a9b430b932960751b608082015290519192507fb626a5b8d37253edd986272eff7a2541a1538cad44624e413d51713f0b2a414d919081900360a00190a1909b909a5098505050505050505050565b60008282018381101561135a576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b8160020b8360020b126140f1576040805162461bcd60e51b8152602060048201526002602482015261563160f01b604482015290519081900360640190fd5b8060020b8360020b12614130576040805162461bcd60e51b81526020600482015260026024820152612b1960f11b604482015290519081900360640190fd5b8060020b8260020b1361416f576040805162461bcd60e51b8152602060048201526002602482015261563360f01b604482015290519081900360640190fd5b620d89e719600284900b12156141b1576040805162461bcd60e51b8152602060048201526002602482015261158d60f21b604482015290519081900360640190fd5b620d89e8600283900b13156141f2576040805162461bcd60e51b8152602060048201526002602482015261563560f01b604482015290519081900360640190fd5b600b54600160a01b9004600290810b810b9084900b8161420e57fe5b0760020b15614249576040805162461bcd60e51b81526020600482015260026024820152612b1b60f11b604482015290519081900360640190fd5b600b54600160a01b9004600290810b810b9083900b8161426557fe5b0760020b15610ffb576040805162461bcd60e51b8152602060048201526002602482015261563760f01b604482015290519081900360640190fd5b6000806142ab611000565b915091506142fb7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000008484614f26565b614303614fc7565b5050565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561437657600080fd5b505afa15801561438a573d6000803e3d6000fd5b505050506040513d60208110156143a057600080fd5b5051604080516370a0823160e01b815230600482015290519192506000916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016916370a08231916024808301926020929190829003018186803b15801561440e57600080fd5b505afa158015614422573d6000803e3d6000fd5b505050506040513d602081101561443857600080fd5b5051905081158015614448575080155b15614454575050612b4c565b6000601954116144645760198290555b6000601a541161447457601a8190555b601954601a5460408051928503602084018190529184038382018190526060808552600490850152636665657360e01b60808501529051919290917fb626a5b8d37253edd986272eff7a2541a1538cad44624e413d51713f0b2a414d9181900360a00190a160008211806144e85750600081115b15612fc2576000806144fa8484615018565b6004549193509150600061450c61126c565b905080156146c35760005b828110156146c15760006145316004838154811061318b57fe5b905060006001600160a01b03166004838154811061454b57fe5b6000918252602090912001546001600160a01b0316148015906145985750306001600160a01b03166004838154811061458057fe5b6000918252602090912001546001600160a01b031614155b6145ce576040805162461bcd60e51b8152602060048201526002602482015261041360f41b604482015290519081900360640190fd5b60006145de8461151d8985613c0f565b905060006145f08561151d8986613c0f565b905060006145fe8383614aa0565b5050905061462d6004868154811061461257fe5b6000918252602090912001546001600160a01b031682614b19565b82600560006004888154811061463f57fe5b60009182526020808320909101546001600160a01b031683528201929092526040018120600201805490920190915560048054849260059290918990811061468357fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020600301805490910190555050506001919091019050614517565b505b5050505050505050565b60006146db8585858561511f565b600b5460408051633c8a7d8d60e01b8152306004820152600289810b602483015288900b60448201526001600160801b038416606482015260a06084820152600060a4820181905282519495506001600160a01b0390931693633c8a7d8d9360c480840194938390030190829087803b15801561475757600080fd5b505af115801561476b573d6000803e3d6000fd5b505050506040513d604081101561478157600080fd5b5050600b8054600296870b62ffffff908116600160b81b0262ffffff60b81b199790980b16600160d01b0262ffffff60d01b19909116179490941694909417909255505050565b600a546007546008546000926001600160a01b0390811692600160401b90048116918116908816831415614827576147ff86611fae565b600954614815906001600160a01b031682612916565b506148218782876119bd565b50614885565b826001600160a01b0316876001600160a01b0316141561486c5761484a85611fae565b600954614860906001600160a01b031683612916565b506148218883886119bd565b6148778883886119bd565b506148838782876119bd565b505b506001979650505050505050565b6000806000806000806148a73089896151c5565b600b546040805163514ea4bf60e01b81526004810184905290519293506001600160a01b039091169163514ea4bf9160248082019260a092909190829003018186803b1580156148f657600080fd5b505afa15801561490a573d6000803e3d6000fd5b505050506040513d60a081101561492057600080fd5b508051602082015160408301516060840151608090940151929c919b50995091975095509350505050565b6000806000600b60009054906101000a90046001600160a01b03166001600160a01b0316633850c7bd6040518163ffffffff1660e01b815260040160e06040518083038186803b15801561499e57600080fd5b505afa1580156149b2573d6000803e3d6000fd5b505050506040513d60e08110156149c857600080fd5b505190506149e8816149d98861521b565b6149e28861521b565b8761554c565b9250925050935093915050565b600454600090815b81811015614a4857836001600160a01b031660048281548110614a1c57fe5b6000918252602090912001546001600160a01b03161415614a405760019250614a48565b6001016149fd565b5081610ffb57600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0180546001600160a01b0385166001600160a01b0319909116179055505050565b60008282614ab6614ab18383613c0f565b6155e8565b92509250925092565b604080516001600160a01b0380861660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052612fc2908590614c0b565b6001600160a01b038216614b74576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b614b8060008383610ffb565b601554614b8d9082614058565b6015556001600160a01b038216600090815260136020526040902054614bb39082614058565b6001600160a01b03831660008181526013602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b6000614c60826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661561a9092919063ffffffff16565b805190915015610ffb57808060200190516020811015614c7f57600080fd5b5051610ffb5760405162461bcd60e51b815260040180806020018281038252602a815260200180615cb7602a913960400191505060405180910390fd5b60006001600160801b03821115614ccf57fe5b5090565b60008080806001600160801b03851615614d8557600b546040805163a34123a760e01b815260028a810b600483015289900b60248201526001600160801b038816604482015281516001600160a01b039093169263a34123a7926064808401939192918290030181600087803b158015614d4c57600080fd5b505af1158015614d60573d6000803e3d6000fd5b505050506040513d6040811015614d7657600080fd5b50805160209091015190945092505b600b54604080516309e3d67b60e31b815230600482015260028a810b602483015289900b60448201526001600160801b03606482018190526084820152815160009384936001600160a01b0390911692634f1eb3d89260a4808301939282900301818787803b158015614df757600080fd5b505af1158015614e0b573d6000803e3d6000fd5b505050506040513d6040811015614e2157600080fd5b5080516020909101516001600160801b039182169350169050614e44828761387c565b9350614e50818661387c565b600f8054860190556010805482019055600c5490935060009081908015614ed157614e82620f424061151d8984613c0f565b9250614e95620f424061151d8884613c0f565b9150614ea1878461387c565b9650614ead868361387c565b600d54909650614ebd9084614058565b600d55600e54614ecd9083614058565b600e555b60408051888152602081018890528082018590526060810184905290517f1ac56d7e866e3f5ea9aa92aa11758ead39a0a5f013f3fefb0f47cb9d008edd279181900360800190a1505050505093509350935093565b600a546007546008546001600160a01b0392831692600160401b909204821691908116908716831415614f7457614f5d8583612b4e565b614f65612af6565b614f6f8482612b4e565b614fbe565b826001600160a01b0316866001600160a01b03161415614faa57614f988482612b4e565b614fa0612af6565b614f6f8583612b4e565b614fb48482612b4e565b614fbe8583612b4e565b50505050505050565b600b54600090614fec90600160b81b8104600290810b91600160d01b9004900b614893565b5050600b5492935061201592600160b81b8104600290810b9350600160d01b909104900b905083614cd3565b600c5460009081908015615110576000615039620f424061151d8885613c0f565b9050600061504e620f424061151d8886613c0f565b905061505a878361387c565b9450615066868261387c565b600d549094506150769083614058565b600d55600e546150869082614058565b600e5581156150c9576002546150c9906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169116846138d9565b801561510957600254615109906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169116836138d9565b5050615117565b8492508391505b509250929050565b600080600b60009054906101000a90046001600160a01b03166001600160a01b0316633850c7bd6040518163ffffffff1660e01b815260040160e06040518083038186803b15801561517057600080fd5b505afa158015615184573d6000803e3d6000fd5b505050506040513d60e081101561519a57600080fd5b505190506151bb816151ab8861521b565b6151b48861521b565b8787615631565b9695505050505050565b6040805160609490941b6bffffffffffffffffffffffff1916602080860191909152600293840b60e890811b60348701529290930b90911b60378401528051808403601a018152603a9093019052815191012090565b60008060008360020b12615232578260020b61523a565b8260020b6000035b9050620d89e8811115615278576040805162461bcd60e51b81526020600482015260016024820152601560fa1b604482015290519081900360640190fd5b60006001821661528c57600160801b61529e565b6ffffcb933bd6fad37aa2d162d1a5940015b70ffffffffffffffffffffffffffffffffff16905060028216156152d2576ffff97272373d413259a46990580e213a0260801c5b60048216156152f1576ffff2e50f5f656932ef12357cf3c7fdcc0260801c5b6008821615615310576fffe5caca7e10e4e61c3624eaa0941cd00260801c5b601082161561532f576fffcb9843d60f6159c9db58835c9266440260801c5b602082161561534e576fff973b41fa98c081472e6896dfb254c00260801c5b604082161561536d576fff2ea16466c96a3843ec78b326b528610260801c5b608082161561538c576ffe5dee046a99a2a811c461f1969c30530260801c5b6101008216156153ac576ffcbe86c7900a88aedcffc83b479aa3a40260801c5b6102008216156153cc576ff987a7253ac413176f2b074cf7815e540260801c5b6104008216156153ec576ff3392b0822b70005940c7a398e4b70f30260801c5b61080082161561540c576fe7159475a2c29b7443b29c7fa6e889d90260801c5b61100082161561542c576fd097f3bdfd2022b8845ad8f792aa58250260801c5b61200082161561544c576fa9f746462d870fdf8a65dc1f90e061e50260801c5b61400082161561546c576f70d869a156d2a1b890bb3df62baf32f70260801c5b61800082161561548c576f31be135f97d08fd981231505542fcfa60260801c5b620100008216156154ad576f09aa508b5b7a84e1c677de54f3e99bc90260801c5b620200008216156154cd576e5d6af8dedb81196699c329225ee6040260801c5b620400008216156154ec576d2216e584f5fa1ea926041bedfe980260801c5b62080000821615615509576b048a170391f7dc42444e8fa20260801c5b60008460020b131561552457806000198161552057fe5b0490505b600160201b81061561553757600161553a565b60005b60ff16602082901c0192505050919050565b600080836001600160a01b0316856001600160a01b0316111561556d579293925b846001600160a01b0316866001600160a01b031611615598576155918585856156f5565b91506155df565b836001600160a01b0316866001600160a01b031610156155d1576155bd8685856156f5565b91506155ca85878561575e565b90506155df565b6155dc85858561575e565b90505b94509492505050565b600060026001830104825b8082101561135a57508060028180868161560957fe5b04018161561257fe5b0491506155f3565b606061562984846000856157a1565b949350505050565b6000836001600160a01b0316856001600160a01b03161115615651579293925b846001600160a01b0316866001600160a01b03161161567c576156758585856158fc565b90506156ec565b836001600160a01b0316866001600160a01b031610156156de5760006156a38786866158fc565b905060006156b287898661595f565b9050806001600160801b0316826001600160801b0316106156d357806156d5565b815b925050506156ec565b6156e985858461595f565b90505b95945050505050565b6000826001600160a01b0316846001600160a01b03161115615715579192915b836001600160a01b031661574e606060ff16846001600160801b0316901b8686036001600160a01b0316866001600160a01b0316615998565b8161575557fe5b04949350505050565b6000826001600160a01b0316846001600160a01b0316111561577e579192915b615629826001600160801b03168585036001600160a01b0316600160601b615998565b6060824710156157e25760405162461bcd60e51b8152600401808060200182810382526026815260200180615bbe6026913960400191505060405180910390fd5b6157eb85615a47565b61583c576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b600080866001600160a01b031685876040518082805190602001908083835b6020831061587a5780518252601f19909201916020918201910161585b565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d80600081146158dc576040519150601f19603f3d011682016040523d82523d6000602084013e6158e1565b606091505b50915091506158f1828286615a4d565b979650505050505050565b6000826001600160a01b0316846001600160a01b0316111561591c579192915b600061593f856001600160a01b0316856001600160a01b0316600160601b615998565b90506156ec61595a84838888036001600160a01b0316615998565b615ab3565b6000826001600160a01b0316846001600160a01b0316111561597f579192915b61562961595a83600160601b8787036001600160a01b03165b60008080600019858709868602925082811090839003039050806159ce57600084116159c357600080fd5b50829004905061135a565b8084116159da57600080fd5b6000848688096000868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b3b151590565b60608315615a5c57508161135a565b825115615a6c5782518084602001fd5b60405162461bcd60e51b8152602060048201818152845160248401528451859391928392604401919085019080838360008315613bcc578181015183820152602001613bb4565b806001600160801b038116811461291157600080fdfe45524332303a207472616e7366657220746f20746865207a65726f20616464726573735265656e7472616e637947756172643a207265656e7472616e742063616c6c0045524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332303a20617070726f766520746f20746865207a65726f206164647265737349662074686973206973206e6f7420302c2074686572652077617320616e206572726f7245786368616e6765205261746520287363616c65642075702062792031653138293a2045524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e6365416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f7745524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f20616464726573738d1cced004452bd270777a8c670f9f7e7c4fdde56f2db331fe289d39dc2624ad5361666545524332303a204552433230206f7065726174696f6e20646964206e6f74207375636365656445524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa26469706673582212201c1502a118c6be638718fff7add228b4c8a3f3206387000f5aac5a41811464fc64736f6c63430007060033"

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

// Assetholder is a free data retrieval call binding the contract method 0x5bb6aa85.
//
// Solidity: function Assetholder(address ) view returns(int256 capital0, int256 capital1, uint256 fees0, uint256 fees1)
func (_Api *ApiCaller) Assetholder(opts *bind.CallOpts, arg0 common.Address) (struct {
	Capital0 *big.Int
	Capital1 *big.Int
	Fees0    *big.Int
	Fees1    *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "Assetholder", arg0)

	outstruct := new(struct {
		Capital0 *big.Int
		Capital1 *big.Int
		Fees0    *big.Int
		Fees1    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Capital0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Capital1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Fees0 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Fees1 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Assetholder is a free data retrieval call binding the contract method 0x5bb6aa85.
//
// Solidity: function Assetholder(address ) view returns(int256 capital0, int256 capital1, uint256 fees0, uint256 fees1)
func (_Api *ApiSession) Assetholder(arg0 common.Address) (struct {
	Capital0 *big.Int
	Capital1 *big.Int
	Fees0    *big.Int
	Fees1    *big.Int
}, error) {
	return _Api.Contract.Assetholder(&_Api.CallOpts, arg0)
}

// Assetholder is a free data retrieval call binding the contract method 0x5bb6aa85.
//
// Solidity: function Assetholder(address ) view returns(int256 capital0, int256 capital1, uint256 fees0, uint256 fees1)
func (_Api *ApiCallerSession) Assetholder(arg0 common.Address) (struct {
	Capital0 *big.Int
	Capital1 *big.Int
	Fees0    *big.Int
	Fees1    *big.Int
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

// GetCapital is a free data retrieval call binding the contract method 0xe1c220bf.
//
// Solidity: function getCapital(address who) view returns(int256, int256, uint256, uint256)
func (_Api *ApiCaller) GetCapital(opts *bind.CallOpts, who common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getCapital", who)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetCapital is a free data retrieval call binding the contract method 0xe1c220bf.
//
// Solidity: function getCapital(address who) view returns(int256, int256, uint256, uint256)
func (_Api *ApiSession) GetCapital(who common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Api.Contract.GetCapital(&_Api.CallOpts, who)
}

// GetCapital is a free data retrieval call binding the contract method 0xe1c220bf.
//
// Solidity: function getCapital(address who) view returns(int256, int256, uint256, uint256)
func (_Api *ApiCallerSession) GetCapital(who common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Api.Contract.GetCapital(&_Api.CallOpts, who)
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
// Solidity: function getLendingAmounts() view returns(uint256, uint256, uint256, uint256)
func (_Api *ApiCaller) GetLendingAmounts(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getLendingAmounts")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetLendingAmounts is a free data retrieval call binding the contract method 0x2e4b8f60.
//
// Solidity: function getLendingAmounts() view returns(uint256, uint256, uint256, uint256)
func (_Api *ApiSession) GetLendingAmounts() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Api.Contract.GetLendingAmounts(&_Api.CallOpts)
}

// GetLendingAmounts is a free data retrieval call binding the contract method 0x2e4b8f60.
//
// Solidity: function getLendingAmounts() view returns(uint256, uint256, uint256, uint256)
func (_Api *ApiCallerSession) GetLendingAmounts() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
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

// UniPortionRate is a free data retrieval call binding the contract method 0x02cfafc8.
//
// Solidity: function uniPortionRate() view returns(uint8)
func (_Api *ApiCaller) UniPortionRate(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "uniPortionRate")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// UniPortionRate is a free data retrieval call binding the contract method 0x02cfafc8.
//
// Solidity: function uniPortionRate() view returns(uint8)
func (_Api *ApiSession) UniPortionRate() (uint8, error) {
	return _Api.Contract.UniPortionRate(&_Api.CallOpts)
}

// UniPortionRate is a free data retrieval call binding the contract method 0x02cfafc8.
//
// Solidity: function uniPortionRate() view returns(uint8)
func (_Api *ApiCallerSession) UniPortionRate() (uint8, error) {
	return _Api.Contract.UniPortionRate(&_Api.CallOpts)
}

// Unwrap is a paid mutator transaction binding the contract method 0x457077cc.
//
// Solidity: function _unwrap(uint256 Amount) returns()
func (_Api *ApiTransactor) Unwrap(opts *bind.TransactOpts, Amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "_unwrap", Amount)
}

// Unwrap is a paid mutator transaction binding the contract method 0x457077cc.
//
// Solidity: function _unwrap(uint256 Amount) returns()
func (_Api *ApiSession) Unwrap(Amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Unwrap(&_Api.TransactOpts, Amount)
}

// Unwrap is a paid mutator transaction binding the contract method 0x457077cc.
//
// Solidity: function _unwrap(uint256 Amount) returns()
func (_Api *ApiTransactorSession) Unwrap(Amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Unwrap(&_Api.TransactOpts, Amount)
}

// Wrap is a paid mutator transaction binding the contract method 0x7c23b921.
//
// Solidity: function _wrap() returns()
func (_Api *ApiTransactor) Wrap(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "_wrap")
}

// Wrap is a paid mutator transaction binding the contract method 0x7c23b921.
//
// Solidity: function _wrap() returns()
func (_Api *ApiSession) Wrap() (*types.Transaction, error) {
	return _Api.Contract.Wrap(&_Api.TransactOpts)
}

// Wrap is a paid mutator transaction binding the contract method 0x7c23b921.
//
// Solidity: function _wrap() returns()
func (_Api *ApiTransactorSession) Wrap() (*types.Transaction, error) {
	return _Api.Contract.Wrap(&_Api.TransactOpts)
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

// RedeemCErc20Tokens is a paid mutator transaction binding the contract method 0x84fea667.
//
// Solidity: function redeemCErc20Tokens(uint256 amount, address _cErc20Contract) returns()
func (_Api *ApiTransactor) RedeemCErc20Tokens(opts *bind.TransactOpts, amount *big.Int, _cErc20Contract common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "redeemCErc20Tokens", amount, _cErc20Contract)
}

// RedeemCErc20Tokens is a paid mutator transaction binding the contract method 0x84fea667.
//
// Solidity: function redeemCErc20Tokens(uint256 amount, address _cErc20Contract) returns()
func (_Api *ApiSession) RedeemCErc20Tokens(amount *big.Int, _cErc20Contract common.Address) (*types.Transaction, error) {
	return _Api.Contract.RedeemCErc20Tokens(&_Api.TransactOpts, amount, _cErc20Contract)
}

// RedeemCErc20Tokens is a paid mutator transaction binding the contract method 0x84fea667.
//
// Solidity: function redeemCErc20Tokens(uint256 amount, address _cErc20Contract) returns()
func (_Api *ApiTransactorSession) RedeemCErc20Tokens(amount *big.Int, _cErc20Contract common.Address) (*types.Transaction, error) {
	return _Api.Contract.RedeemCErc20Tokens(&_Api.TransactOpts, amount, _cErc20Contract)
}

// RedeemCEth is a paid mutator transaction binding the contract method 0xa5541f28.
//
// Solidity: function redeemCEth(uint256 amount, address _cEtherContract) returns()
func (_Api *ApiTransactor) RedeemCEth(opts *bind.TransactOpts, amount *big.Int, _cEtherContract common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "redeemCEth", amount, _cEtherContract)
}

// RedeemCEth is a paid mutator transaction binding the contract method 0xa5541f28.
//
// Solidity: function redeemCEth(uint256 amount, address _cEtherContract) returns()
func (_Api *ApiSession) RedeemCEth(amount *big.Int, _cEtherContract common.Address) (*types.Transaction, error) {
	return _Api.Contract.RedeemCEth(&_Api.TransactOpts, amount, _cEtherContract)
}

// RedeemCEth is a paid mutator transaction binding the contract method 0xa5541f28.
//
// Solidity: function redeemCEth(uint256 amount, address _cEtherContract) returns()
func (_Api *ApiTransactorSession) RedeemCEth(amount *big.Int, _cEtherContract common.Address) (*types.Transaction, error) {
	return _Api.Contract.RedeemCEth(&_Api.TransactOpts, amount, _cEtherContract)
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
	Sender         common.Address
	To             common.Address
	Shares         *big.Int
	Amount0        *big.Int
	Amount1        *big.Int
	Poolamount0    *big.Int
	Poolamount1    *big.Int
	Lendingamount0 *big.Int
	Lendingamount1 *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9465fbc710f5c7ebb85aec935f56a888087125eaf287d09f2d3b59c640b3eafd.
//
// Solidity: event Withdraw(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1, uint256 poolamount0, uint256 poolamount1, uint256 lendingamount0, uint256 lendingamount1)
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

// WatchWithdraw is a free log subscription operation binding the contract event 0x9465fbc710f5c7ebb85aec935f56a888087125eaf287d09f2d3b59c640b3eafd.
//
// Solidity: event Withdraw(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1, uint256 poolamount0, uint256 poolamount1, uint256 lendingamount0, uint256 lendingamount1)
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

// ParseWithdraw is a log parse operation binding the contract event 0x9465fbc710f5c7ebb85aec935f56a888087125eaf287d09f2d3b59c640b3eafd.
//
// Solidity: event Withdraw(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1, uint256 poolamount0, uint256 poolamount1, uint256 lendingamount0, uint256 lendingamount1)
func (_Api *ApiFilterer) ParseWithdraw(log types.Log) (*ApiWithdraw, error) {
	event := new(ApiWithdraw)
	if err := _Api.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

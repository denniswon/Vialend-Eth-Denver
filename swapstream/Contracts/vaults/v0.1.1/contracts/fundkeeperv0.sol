// SPDX-License-Identifier: Unlicense

/*
0x7E83D2768B6A96FD753AB307fac88b08E2B7debA
*/

pragma solidity >=0.4.21 <= 0.8.6;

import "./@openzeppelin/contracts/math/Math.sol";
import "./@openzeppelin/contracts/math/SafeMath.sol";
import "./@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "./@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
import "./@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "./@uniswap/v3-core/contracts/interfaces/callback/IUniswapV3MintCallback.sol";
import "./@uniswap/v3-core/contracts/interfaces/callback/IUniswapV3SwapCallback.sol";
import "./@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";
import "./@uniswap/v3-core/contracts/libraries/TickMath.sol";
import "./@uniswap/v3-periphery/contracts/libraries/LiquidityAmounts.sol";
import "./@uniswap/v3-periphery/contracts/libraries/PositionKey.sol";

import "../interfaces/IFundKeeper.sol";

/// @author  Swapstream
/// @title   FundKeeper
/// @notice  A Smart Contract that helps liquidity providers managing their funds on Uniswap V3.

contract FundKeeper is 
	IFundKeeper,
    ReentrancyGuard

{
    using SafeMath for uint256;

    event Deposit(
        address indexed sender,
        address indexed to,
        uint256 shares,
        uint256 amount0,
        uint256 amount1
    );

    event Withdraw(
        address indexed sender,
        address indexed to,
        uint256 shares,
        uint256 amount0,
        uint256 amount1
    );

    event CollectFees(
        uint256 feesToVault0,
        uint256 feesToVault1,
        uint256 feesToProtocol0,
        uint256 feesToProtocol1
    );

    //event Snapshot(int24 tick, uint256 totalAmount0, uint256 totalAmount1, uint256 totalSupply);


    uint256 public protocolFee;
    uint256 public maxTotalSupply;
    address public strategy;
    address public governance;
    address public pendingGovernance;

    int24 public baseLower;
    int24 public baseUpper;
    int24 public limitLower;
    int24 public limitUpper;
    uint256 public accruedProtocolFees0;
    uint256 public accruedProtocolFees1;

    /// @dev 
    /// @param _pool Uniswap V3 pool address
    /// @param _protocolFee Protocol fee expressed as multiple of 1e-6
    /// @param _maxTotalSupply Cap on total supply
    
    constructor(
        address _pool,
        uint256 _protocolFee,
        uint256 _maxTotalSupply
    )  {

		require (_pool != address(0));
        protocolFee = _protocolFee;
        maxTotalSupply = _maxTotalSupply;
        governance = msg.sender;

        require(_protocolFee < 1e6, "protocolFee");
    }

    
    /// @notice Deposits tokens .
    /// @dev These tokens sit in the vault and are not used for liquidity on
    /// Uniswap until the next rebalance. Also note it's not necessary to check
    /// if user manipulated price to deposit cheaper, as the value of range
    /// orders can only by manipulated higher.
    /// @param amount0Desired Max amount of token0 to deposit
    /// @param amount1Desired Max amount of token1 to deposit
    /// @param amount0Min Revert if resulting `amount0` is less than this
    /// @param amount1Min Revert if resulting `amount1` is less than this
    /// @param to Recipient of shares
    /// @return shares Number of shares minted
    /// @return amount0 Amount of token0 deposited
    /// @return amount1 Amount of token1 deposited
     
    function deposit(
        uint256 amount0Desired,
        uint256 amount1Desired,
        uint256 amount0Min,
        uint256 amount1Min,
        address to
    )
        external
        override
        returns (
            uint256 shares,
            uint256 amount0,
            uint256 amount1
        )
    {
        require(amount0Desired > 0 || amount1Desired > 0, "amount0Desired or amount1Desired");
       
        require(to != address(0) && to != address(this), "to");

        shares = (amount0Desired * amount1Desired) / 10;
        
        amount0 = amount0Desired;
        
        amount1 = amount1Desired;
        
        require(amount0 >= amount0Min, "amount0Min");

        require(amount1 >= amount1Min, "amount1Min");

        
		
        emit Deposit(msg.sender, to, shares, amount0, amount1);
    }


    
    /// @notice Withdraws tokens in proportion to the vault's holdings.
    /// @param shares Shares burned by sender
    /// @param amount0Min Revert if resulting `amount0` is smaller than this
    /// @param amount1Min Revert if resulting `amount1` is smaller than this
    /// @param to Recipient of tokens
    /// @return amount0 Amount of token0 sent to recipient
    /// @return amount1 Amount of token1 sent to recipient
     
    function withdraw(
        uint256 shares,
        uint256 amount0Min,
        uint256 amount1Min,
        address to
    ) external override returns (uint256 amount0, uint256 amount1) {
    	
        require(shares > 0, "shares");
        require(to != address(0) && to != address(this), "to");


        amount0 = amount0Min + 1;
        amount1 = amount1Min + 1;

        require(amount0 >= amount0Min, "amount0Min");

        require(amount1 >= amount1Min, "amount1Min");
		
		
        emit Withdraw(msg.sender, to, shares, amount0, amount1);
    }

    
    function getTotalAmounts() public pure override returns (uint256 total0, uint256 total1) {
        // (uint256 baseAmount0, uint256 baseAmount1) = getPositionAmounts(baseLower, baseUpper);
        // (uint256 limitAmount0, uint256 limitAmount1) =  getPositionAmounts(limitLower, limitUpper);
        
        //total0 = getBalance0().add(baseAmount0).add(limitAmount0);
        //total1 = getBalance1().add(baseAmount1).add(limitAmount1);
        
        total0 = 100;
        total1 = 10;
    }




    modifier onlyGovernance {
        require(msg.sender == governance, "governance");
        _;
    }
    
    ///
	/// @notice 3 functiosn for demo
	/// @dev 
	///		Hello(), Greet(), plus(x,y)
	///
	function Hello() public pure returns (string memory) {
        return "Hello World";
    }
    
    function Greet(string memory str) public pure returns (string memory) {
        return str;
    }
    
    function plus(uint256 x, uint256 y) public pure returns (uint256) {
        return x.add(y);
    }    
}
// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.5.0;

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

//import "../interfaces/IFundKeeper.sol";
import "../interfaces/IFundKeeperEvents.sol";


/// @author  Swapstream
/// @title   FundKeeper
/// @notice  A Smart Contract that helps liquidity providers managing their funds on Uniswap V3.

contract FundKeeper is 
    ERC20,
    IFundKeeperEvents, 
    ReentrancyGuard
    

{
    using SafeERC20 for IERC20;
    using SafeMath for uint256;
	

    IUniswapV3Pool public immutable pool;

	IERC20 public immutable token0;
    IERC20 public immutable token1;
    IERC20 public immutable ttoken;
    
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
    ///	@param _ttoken SWSM team token
    /// @param _protocolFee Protocol fee expressed as multiple of 1e-6
    /// @param _maxTotalSupply Cap on total supply
    
    constructor(
        address _pool,
        address _ttoken,
        uint256 _protocolFee,
        uint256 _maxTotalSupply
    ) ERC20("SWAP STREAM","SWSM") {

		pool = IUniswapV3Pool(_pool);	
		emit GeneralA("pool", _pool);
		
        token0 = IERC20(IUniswapV3Pool(_pool).token0());
        emit GeneralS("token0", "token0");
        
        token1 = IERC20(IUniswapV3Pool(_pool).token1());
        emit GeneralS("token1", "token1");
        
        ttoken = IERC20(_ttoken);
        emit GeneralS("ttoken", "ttoken");

        protocolFee = _protocolFee;
        maxTotalSupply = _maxTotalSupply;
        governance = msg.sender;

        require(_protocolFee < 1e6, "protocolFee");
    }

    
    /// - Deposit tokens 
    /// @notice tokens get deposited in this smart contract will be held until next rebalance. 
    /// @param amountToken0 amount of token0 to deposit
    /// @param amountToken0 amount of token1 to deposit
    /// @param amount0Min Revert if resulting `amount0` is less than this
    /// @param amount1Min Revert if resulting `amount1` is less than this
    /// @param user  The user's address of the tokens 
    /// @return shares Number of shares minted
    /// @return amount0 Amount of token0 deposited
    /// @return amount1 Amount of token1 deposited
     
    function deposit(
        uint256 amountToken0,
        uint256 amountToken1,
        uint256 amount0Min,
        uint256 amount1Min,
        address user
    )
        external
        nonReentrant
        
        returns (
            uint256 shares,
            uint256 amount0,
            uint256 amount1
        )
    {
        require(amountToken0 > 0 || amountToken1 > 0, "amountToken0 or amountToken1");
       
        require(user != address(0) && user != address(this), "user address invalid");

  		/// #debug will be replaced. Poke positions so to get uniswap v3 fees up to date. 
        //_poke(baseLower, baseUpper);
        //_poke(limitLower, limitUpper);

		
		require(amountToken0 > 0,"amountToken0");  //#debug remove later
		require(amountToken1 > 0,"amountToken1");  //#debug remove later

        (shares, amount0, amount1) = _shares(amountToken0, amountToken1); //_sharesByToken(amountToken0, amountToken1);
        require(amount0 >= amount0Min, "Deposit amount0");
        require(amount1 >= amount1Min, "Deposit amount1");
        require(shares > 0, uint2str(shares));

        // Pull in tokens from sender
        if (amount0 > 0) token0.safeTransferFrom(msg.sender, address(this), amount0);
        if (amount1 > 0) token1.safeTransferFrom(msg.sender, address(this), amount1);

        // Mint shares to recipient
        _mint(user, shares);

		// send user ttoken        
		uint giveOutRate = 3;
		
        ttoken.safeTransfer(user, giveOutRate.div(10000).mul(shares));

        emit Deposit(msg.sender, user, shares, amount0, amount1,token0.name(),token1.name());

        //require(totalSupply() <= maxTotalSupply, "maxTotalSupply");
    }

    function _shares(uint256 amountToken0, uint256 amountToken1)
        internal
        pure
        returns (
            uint256 shares,
            uint256 amount0,
            uint256 amount1
        )
        {
			amount0 = amountToken0;
			amount1 = amountToken1;
        	shares = amount0+amount1;
        }

    
    
    /// - user Withdraw 
    /// @param user address of Recipient of tokens
    /// @param shares number of Shares owned by sender to be burned
    /// @return amount0 Amount of token0 sent to user 
    /// @return amount1 Amount of token1 sent to user
    /// @param amount0Min Revert if resulting `amount0` is smaller than this
    /// @param amount1Min Revert if resulting `amount1` is smaller than this
    
    function withdraw(
        uint256 shares,
        uint256 amount0Min,
        uint256 amount1Min,
        address user
    ) external  nonReentrant returns (uint256 amount0, uint256 amount1) {
        require(shares > 0, "shares");

        require(user != address(0) && user != address(this), "user");


        // Burn shares
        _burn(msg.sender, shares);

/**
        uint256 totalSupply = totalSupply();

        // Calculate token amounts proportional to unused balances
        uint256 unusedAmount0 = getBalance0().mul(shares).div(totalSupply);
        uint256 unusedAmount1 = getBalance1().mul(shares).div(totalSupply);

        // Withdraw proportion of liquidity from Uniswap pool
        (uint256 baseAmount0, uint256 baseAmount1) =
            _burnLiquidityShare(baseLower, baseUpper, shares, totalSupply);
        (uint256 limitAmount0, uint256 limitAmount1) =
            _burnLiquidityShare(limitLower, limitUpper, shares, totalSupply);

        // Sum up total amounts owed to recipient
        amount0 = unusedAmount0.add(baseAmount0).add(limitAmount0);
        amount1 = unusedAmount1.add(baseAmount1).add(limitAmount1);
*/

        //require(amount0 >= amount0Min, "withdraw amount0>min");
        //require(amount1 >= amount1Min, "withdraw amount1>min");
        
        //#debug remove later
        amount0 = amount0Min;
        amount1 = amount1Min;

        // Push tokens to recipient
        if (amount0 > 0) token0.safeTransfer(user, amount0);
        if (amount1 > 0) token1.safeTransfer(user, amount1);

        emit Withdraw(msg.sender, user, shares, amount0, amount1,token0.name(),token1.name());
    }

    function _sharesByToken(uint256 amountToken0, uint256 amountToken1)
        internal
        view
        returns (
            uint256 shares,
            uint256 amount0,
            uint256 amount1
        )
    {
    		// #debug remove it later
			require(amountToken0 > 0 );
			require(amountToken1 > 0 );
			
    	// total Amount of the shares in vault
        uint256 totalSupply = totalSupply();
        
        //get total amount from current balance + position liquidity for each token
        (uint256 total0, uint256 total1) = getTotalAmounts();

        // If total supply > 0, vault can't be empty
        assert(totalSupply == 0 || total0 > 0 || total1 > 0);
        
        (shares,amount0,amount1) = (amount0+amount1, amount0,amount1);

/*
        if (totalSupply == 0) {
            // For first deposit, just use the amounts desired
            amount0 = amountToken0;
            amount1 = amountToken1;
            shares = Math.max(amount0, amount1);
        } else if (total0 == 0) { 
            amount1 = amountToken1;
            shares = amount1.mul(totalSupply).div(total1);
        } else if (total1 == 0) {
            amount0 = amountToken0;
            shares = amount0.mul(totalSupply).div(total0);
        } else {
            uint256 cross = Math.min(amountToken0.mul(total1), amountToken1.mul(total0));
            require(cross > 0, "cross");

            // Round up amounts
            amount0 = cross.sub(1).div(total1).add(1);
            amount1 = cross.sub(1).div(total0).add(1);
            shares = cross.mul(totalSupply).div(total0).div(total1);
        }
        
     */
      
    }    
    
    function getTotalAmounts() 
    	public 
    	view  
    	returns (uint256 total0, uint256 total1) {
         
        (uint256 Amount0, uint256 Amount1) = (0,0) ;  // getPositionAmounts(baseLower, baseUpper);
        
        total0 = getBalance0().add(Amount0);
        total1 = getBalance1().add(Amount1);

    }


	/**
     * @notice Amounts of token0 and token1 held in vault's position. Includes
     * owed fees but excludes the proportion of fees that will be paid to the
     * protocol. Doesn't include fees accrued since last poke.
    /// @param tickLower lower line price from last rebalance, 
    /// @param tickUpper upper line price from last rebalance
     */
    function getPositionAmounts(int24 tickLower, int24 tickUpper)
        public
        view
        returns (uint256 amount0, uint256 amount1)
    {
        (uint128 liquidity, , , uint128 tokensOwed0, uint128 tokensOwed1) =
            _position(tickLower, tickUpper);
        (amount0, amount1) = _amountsForLiquidity(tickLower,tickUpper, liquidity);

        // Subtract protocol fees
        uint256 oneMinusFee = uint256(1e6).sub(protocolFee);
        amount0 = amount0.add(uint256(tokensOwed0).mul(oneMinusFee).div(1e6));
        amount1 = amount1.add(uint256(tokensOwed1).mul(oneMinusFee).div(1e6));
    }


	/// @dev Wrapper around `IUniswapV3Pool.positions()`.
    function _position(int24 tickLower, int24 tickUpper)
        internal
        view
        returns (
            uint128,
            uint256,
            uint256,
            uint128,
            uint128
        )
    {
        bytes32 positionKey = PositionKey.compute(address(this), tickLower, tickUpper);
        return pool.positions(positionKey);
    }
    
    /// @dev Wrapper around `LiquidityAmounts.getAmountsForLiquidity()`.
    function _amountsForLiquidity(
        int24 tickLower,
        int24 tickUpper,
        uint128 liquidity
    ) internal view returns (uint256, uint256) {
        (uint160 sqrtRatioX96, , , , , , ) = pool.slot0();
        return
            LiquidityAmounts.getAmountsForLiquidity(
                sqrtRatioX96,
                TickMath.getSqrtRatioAtTick(tickLower),
                TickMath.getSqrtRatioAtTick(tickUpper),
                liquidity
            );
    }


    /// @notice return Balance of available token0.
     
    function getBalance0() public view returns (uint256) {
        return token0.balanceOf(address(this)).sub(accruedProtocolFees0);
    }

    /// @notice return Balance of available token1.
    
    function getBalance1() public view returns (uint256) {
        return token1.balanceOf(address(this)).sub(accruedProtocolFees1);
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


function uint2str(uint _i) internal pure returns (string memory _uintAsString) {
        if (_i == 0) {
            return "0";
        }
        uint j = _i;
        uint len;
        while (j != 0) {
            len++;
            j /= 10;
        }
        bytes memory bstr = new bytes(len);
        uint k = len - 1;
        while (_i != 0) {
            bstr[k--] = byte(uint8(48 + _i % 10));
            _i /= 10;
        }
        return string(bstr);
    }
}


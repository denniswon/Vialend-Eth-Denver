// SPDX-License-Identifier: MIT
pragma solidity >=0.5.0;

import "../interfaces/IWETH9.sol";
import "../interfaces/IcEth.sol";
import "../interfaces/IcERC20.sol";
import "../@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../@openzeppelin/contracts/token/ERC20/ERC20.sol";

import "../vialendfeemaker.sol.go";


contract Strategy1 {

 	ViaLendFeeMaker public immutable feemaker;
	
	

	constructor(
        address _feemaker,
        address _team,
        address payable _weth,
		address _cToken0,
        address _cToken1,
        address _cEth,
        int24 _maxTwapDeviation,
        uint32 _twapDuration
       ){
    	
		
		
    }


	function strategy0(
		int24 newLow,
        int24 newHigh,
        int256 swapAmount,
        bool zeroForOne
        
		) external nonReentrant onlyTeam  {
		
        (	uint160 sqrtPriceX96, 	int24 tick, , , , , ) 	= pool.slot0();

  		_validRange(newLow, newHigh, tick);  // passed 1200 , 2100, 18382
        
        // Check price is not too close to min/max allowed by Uniswap. Price
        // shouldn't be this extreme unless something was wrong with the pool.

        int24 range = newHigh - newLow ;
            
        require(tick > TickMath.MIN_TICK + range  + tickSpacing, "tick too low");
        require(tick < TickMath.MAX_TICK - range  - tickSpacing, "tick too high");

        // Check price has not moved a lot recently. This mitigates price
        // manipulation during rebalance and also prevents placing orders
        // when it's too volatile.
        int24 twap = getTwap();
        int24 deviation = tick > twap ? tick - twap : twap - tick;
        require(deviation <= maxTwapDeviation, _hint2("deviation,maxTwapDeviation", uint256(deviation),uint256(maxTwapDeviation), 0,""));

        
        (uint256 balance0, uint256 balance1 ) = feemaker.removeUniswap();
        
    //    uint256 balance0 = getBalance0();
     //   uint256 balance1 = getBalance1();
        
		// to be optimized outside
		uint160 sqrtPriceLimitX96 = zeroForOne ? TickMath.MIN_SQRT_RATIO + 1 : TickMath.MAX_SQRT_RATIO - 1;

		if (swapAmount > 0) {
	        swap( 
	        	 swapAmount, 
				 zeroForOne ,
				sqrtPriceLimitX96 );
		}
        
        

//        int24 tickFloor = _floor(tick);
//        int24 tickCeil = tickFloor + tickSpacing;

		//add position
		feemaker.rebalance(
	        newLow,
	        newHigh,
			tick
			);

        lastRebalance = block.timestamp;
        lastTick = tick;
        		
	}
	

function strategy1(
		int24 newLow,
        int24 newHigh
		) external nonReentrant   {
		
		require(msg.sender == team, "strategy1 team");
		
        ( uint160 sqrtPriceX96, 	int24 tick, , , , , ) 	= pool.slot0();

  		_validRange(newLow, newHigh, tick);  // passed 1200 , 2100, 18382
        
        // Check price is not too close to min/max allowed by Uniswap. Price
        // shouldn't be this extreme unless something was wrong with the pool.

        int24 range = newHigh - newLow ;
            
        require(tick > TickMath.MIN_TICK + range  + tickSpacing, "tick too low");
        require(tick < TickMath.MAX_TICK - range  - tickSpacing, "tick too high");

        int24 twap = getTwap();
        int24 deviation = tick > twap ? tick - twap : twap - tick;
        
        // avoid high slipage
        require(deviation <= maxTwapDeviation, _hint2("deviation,maxTwapDeviation", uint256(deviation),uint256(maxTwapDeviation), 0,""));

        // remove positions from uniswap and lending pool get back to vault
        (uint256 unUsedbalance0, uint256 unUsedbalance1) = feemaker.removePositions();
        
        // rebalance, 90% lending, 10% liquidity


        //add 90% assets to lending
        uint8 lendingPortion = 90;
        uint256 lendingBalance0 = unUsedbalance0.mul(lendingPortion).div(100);
        uint256 lendingBalance1 = unUsedbalance1.mul(lendingPortion).div(100);
        
        bool result = feemaker.lendingSupply( lendingBalance0,  lendingBalance1);
        
        require(result, "lending supply failed");
       

		//add rest portion to uniswap
		feemaker.rebalance(
	        newLow,
	        newHigh,
			tick
			);

        lastRebalance = block.timestamp;
        lastTick = tick;
        		
	}

 
	    /// @dev Fetches time-weighted average price in ticks from Uniswap pool.
    function getTwap() public view returns (int24) {
        uint32 _twapDuration = twapDuration;
        uint32[] memory secondsAgo = new uint32[](2);
        secondsAgo[0] = _twapDuration;
        secondsAgo[1] = 0;

        (int56[] memory tickCumulatives, ) = pool.observe(secondsAgo);
        return int24((tickCumulatives[1] - tickCumulatives[0]) / _twapDuration);
    }

    /// @dev Rounds tick down towards negative infinity so that it's a multiple
    /// of `tickSpacing`.
    function _floor(int24 tick) internal view returns (int24) {
        int24 compressed = tick / tickSpacing;
        if (tick < 0 && tick % tickSpacing != 0) compressed--;
        return compressed * tickSpacing;
    }
   

    function setMaxTwapDeviation(int24 _maxTwapDeviation) external onlyTeam {
        require(_maxTwapDeviation > 0, "maxTwapDeviation");
        maxTwapDeviation = _maxTwapDeviation;
    }

    function setTwapDuration(uint32 _twapDuration) external onlyTeam {
        require(_twapDuration > 0, "twapDuration");
        twapDuration = _twapDuration;
    }
	

    /// @dev Uses same governance as underlying vault.
    modifier onlyGovernance {
        require(msg.sender == feemaker.governance(), "governance");
        _;
    }

	modifier onlyTeam {
        require(msg.sender == feemaker.team(), "team");
        _;
    }   

}
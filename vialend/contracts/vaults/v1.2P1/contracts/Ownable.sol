// SPDX-License-Identifier: MIT

pragma solidity >0.5.0;

import "./@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "./@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";


/**
* @title Ownable
* @dev The Ownable contract has an owner address, and provides basic authorization control
* functions.
*/
contract Ownable is ReentrancyGuard {


    uint128 public quoteAmount = 1e18;	// for uniswap v3 oracle  if token0 is weth 1e18, if first token0 is usdc 1e6


    address public governance;
    address public team;
    address public pendingGovernance;

    uint256 public protocolFee;
    

    IUniswapV3Pool internal  pool;
    int24 internal  tickSpacing;

    int24 public cLow;
    int24 public cHigh;

	uint256 internal uFees0;	// uniswap fee of token0 for the current position
	uint256 internal uFees1;	// uniswap fee of token1 for the current position
	uint256 internal lFees0;
	uint256 internal lFees1;

	struct FeeStruct {
		uint256 U3Fees0;				// uni v3 fees 0
		uint256 U3Fees1;				// uni v3 fees 1
		uint256 LcFees0;				// compound fees0
		uint256 LcFees1;				// compound fees1
	    uint256 AccruedProtocolFees0;		// for view
	    uint256 AccruedProtocolFees1;		// for view
	}

     // Asset of each user.
	struct Assets  {
    	uint256 deposit0;   	// user's accumulative deposits for token0
		uint256 deposit1; 	// user's accumulative deposits for token1
		uint256 current0;		// log current locked value of token0
		uint256 current1;		// log current locked value of token1
		uint256 block; 		//  last block number that a user made deposit
    }
    
    address[] public accounts;
    
    mapping(address => uint )  public accId; 	// index of address in the accounts array
    
    mapping(address => Assets)  public Assetholder;
    
    FeeStruct public Fees;

    
    uint256 public maxTotalSupply;
    
    uint32 internal twapDuration;
    
    
	int24 internal maxTwapDeviation;    
    
	uint8 public uniPortion ;

	
	event MyLog(string, uint256);
	event MyLog2(string, uint256,uint256);
	event MyLog4(string, uint256,uint256,uint256,uint256);



	/// maintain a user address array
    function _push(address _addr ) internal {
    	
		if (! _exist( _addr)  ) {
			
			accounts.push(_addr);
		 	
			accId[_addr] = accounts.length;
			
			Assetholder[_addr].block = block.number;
		}
		
		// if (Assetholder[_address].block == 0 ) {
        	
		// 	accounts.push(_address);
			
		// 	accId[_address] = accounts.length-1;
			
		// 	Assetholder[_address].block = block.number;	
			
		// } else {
		// 	Assetholder[_address].block = block.number;	// update block number
		// }

    }

	function _exist(address _addr ) internal view returns (bool){
		return ( accId[_addr] > 0 );
	}
	
	///@notice set new maxTotalSupply
	function setMaxTotalSupply(uint256 newMax) external nonReentrant onlyGovernance {
			maxTotalSupply = newMax;
	}

	function setGovernance(address _governance) external onlyGovernance {
        pendingGovernance = _governance;
    }
    /**
     * @notice `setGovernance()` should be called by the existing governance
     * address prior to calling this function.
     */
    function acceptGovernance() external {
         require(msg.sender == pendingGovernance, "pendingGovernance");
        governance = msg.sender;
    }

	function setTeam(address _team) external onlyGovernance {
        team = _team;
    }
    
	
    modifier onlyGovernance {
         require(msg.sender == governance, "governance");
        _;
    }
    

    function setMaxTwapDeviation(int24 _maxTwapDeviation) external onlyGovernance {
         require(_maxTwapDeviation > 0, "maxTwapDeviation");
        maxTwapDeviation = _maxTwapDeviation;
    }

    function setTwapDuration(uint32 _twapDuration) external onlyGovernance {
         require(_twapDuration > 0, "twapDuration");
        twapDuration = _twapDuration;
    }
    
    function setUniPortionRatio(uint8 ratio) external onlyGovernance {
    	require (ratio <= 100,"ratio");
		uniPortion = ratio;
    }

    function setProtocolFee(uint256 _protocolFee) external onlyGovernance {
        require(_protocolFee <= 20, "protocolFee");
        protocolFee = _protocolFee;
    }
    


}

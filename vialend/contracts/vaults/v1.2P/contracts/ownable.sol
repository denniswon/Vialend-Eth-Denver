// SPDX-License-Identifier: MIT

pragma solidity >0.5.0;

import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";


/**
* @title Ownable
* @dev The Ownable contract has an owner address, and provides basic authorization control
* functions.
*/
contract Ownable is ReentrancyGuard {

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
		uint256 U3Fees0;
		uint256 U3Fees1;
		uint256 LcFees0;
		uint256 LcFees1;
	    uint256 AccruedProtocolFees0;
	    uint256 AccruedProtocolFees1;
	}

     // Asset of each user.
	struct Assets  {
    	uint256 deposit0;   	// user's accumulative deposits for token0
		uint256 deposit1; 	// user's accumulative deposits for token1
		uint256 current0;		// log current value of token0
		uint256 current1;		// log current value of token1
		uint256 block; 		// the block number that added an account
    }
    
    address[] public accounts;
    
    
    mapping(address => Assets)  public Assetholder;
    
    FeeStruct public Fees;

	
    //Asset storage c = Assetholder[accounts[i]]
  	//Asset({capital0:amountToken0, capital1:0,fees0:0,fees1:0});
    
    uint256 public maxTotalSupply;
    uint32 public twapDuration;
	int24 public maxTwapDeviation;    
    
	uint8 public uniPortion ;


	
	event MyLog(string, uint256);
	event MyLog2(string, uint256,uint256);
	event MyLog4(string, uint256,uint256,uint256,uint256);



	/// maintain a user address array
    function _push(address _address ) internal {
    	
		if (Assetholder[_address].block == 0 ) {
        	accounts.push(_address);
			Assetholder[_address].block = block.number;
		}

    }

	// commen out because code size exceeds 24576bytes    
    // function list() external view onlyGovernance returns (address[]  memory ) {
    //     return accounts;
    // }


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
        require(_protocolFee < 1e6, "protocolFee");
        protocolFee = _protocolFee;
    }
    
	function getETHBalance() public view returns (uint256) {
         return address(this).balance;
    }

	// /// return capital amount
	function getStoredAssets(address who) public view returns( uint256 , uint256, uint256, uint256 ) {
	   	return (
			Assetholder[who].deposit0,
			Assetholder[who].deposit1 , 
			Assetholder[who].current0 , 
			Assetholder[who].current1 );
	}



}

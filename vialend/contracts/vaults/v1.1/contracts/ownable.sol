// SPDX-License-Identifier: MIT

pragma solidity >0.5.0;

import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";


/**
* @title Ownable
* @dev The Ownable contract has an owner address, and provides basic authorization control
* functions, this simplifies the implementation of "user permissions".
*/
contract Ownable is ReentrancyGuard {
    address public governance;
    address public team;
    address pendingGovernance;
    
    address[] public accounts;
 
    struct Asset  {
    	int256 capital0;
		int256 capital1;
		uint256 fees0;
		uint256 fees1;
//		uint256[] deposittime;  // {1624136400, 1624136405... }
    }
    
    mapping(address => Asset)  Assetholder;
	
    //Asset storage c = Assetholder[accounts[i]]
  	//Asset({capital0:amountToken0, capital1:0,fees0:0,fees1:0});
    
    uint256 public maxTotalSupply;
    uint32 public twapDuration;
	int24 public maxTwapDeviation;    
    
	uint8 public uniPortionRate ;


    function push(address _address ) internal {
    	
		bool _exists= false;
		uint cnt = accounts.length;
		
		for (uint i=0; i< cnt; i++) {
			if (accounts[i] == _address) {
				_exists = true;
				break;
			}
        }
        
        if (!_exists) accounts.push(_address);

    }

	// commen out because code size exceeds 24576bytes    
    // function list() external view onlyTeam returns (address[]  memory ) {
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
    
    modifier onlyTeam {
         require(msg.sender == team, "team");
        _;
    }

	
    modifier onlyGovernance {
         require(msg.sender == governance, "governance");
        _;
    }
    
    

    function setMaxTwapDeviation(int24 _maxTwapDeviation) external onlyTeam {
         require(_maxTwapDeviation > 0, "maxTwapDeviation");
        maxTwapDeviation = _maxTwapDeviation;
    }

    function setTwapDuration(uint32 _twapDuration) external onlyTeam {
         require(_twapDuration > 0, "twapDuration");
        twapDuration = _twapDuration;
    }
    
    function setUniPortionRatio(uint8 ratio) external onlyTeam {
    	require (ratio <= 100,"ratio");
		uniPortionRate = ratio;
    }


	function getETHBalance() public view returns (uint256) {
         return address(this).balance;
    }

	/// return capital amount
	function getCapital(address who) public view returns( int256 , int256,uint256, uint256 ) {
	  	return (Assetholder[who].capital0,Assetholder[who].capital1 , Assetholder[who].fees0 , Assetholder[who].fees1 );
	}

	


}

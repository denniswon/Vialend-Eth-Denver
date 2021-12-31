// SPDX-License-Identifier: MIT

pragma solidity >=0.8.8 <0.9.0;

import "./@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
import "./@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "./@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "./interfaces/IVaultFactory.sol";
import "./Ownable.sol";
import "./interfaces/IStrategy.sol";

/// @author  ViaLend Team
/// @title   Vault for Strategy Uni + Compound
/// @notice  A Smart Contract that helps liquidity providers managing their funds on Uniswap V3 and Compound.

contract ViaVault is 
	ERC20,
	ReentrancyGuard,
	Ownable
{
    using SafeERC20 for IERC20;
    
    address public immutable factory;  
    address public  strategy;  
    address public immutable admin; 
    address public immutable token0;   // token0
	address public immutable token1;   // token1

	uint256 public vaultCap;
	uint256 public individualCap;
	bool public isUsed = false;
	
    constructor(
    	address _factory,
		address _admin,
        address _token0,
		address _token1,
		uint256 _vaultCap,
		uint256 _individualCap
    ) ERC20("ViaLend Uni Compound Token","VUC0") {
    	
		factory = _factory; //msg.sender;
		admin = _admin;
        token0 = _token0;
        token1 = _token1;
        vaultCap = _vaultCap;
		individualCap = _individualCap;
    }

	event Deposit(
        address indexed sender,
        uint256 shares,
        uint256 amount0,
        uint256 amount1
    );

    event Withdraw(
        address indexed sender,
        uint256 shares,
        uint256 amount0,
        uint256 amount1
     
    );

 	modifier onlyAdmin {
        require (msg.sender == admin, 'admin');  
        _;
    }
    
    modifier onlyStrategy {
    	//require (IVaultFactory(factory).pairStrategy(address(this)) == msg.sender, "strat" );
        require (msg.sender == strategy, 'strat');
        _;
    }
    
    modifier onlyActive {
        require (IVaultFactory(factory).checkActive( address(this) ), 'not active');
        _;
    }
    
    
    
    modifier onlyFactory {
    	require(msg.sender == factory, 'factory');
		_;
    }
    
	/// fallback function has been split into receive() and fallback(). It is a new change of the compiler.
	receive() external payable {}
	fallback() external payable {}

	

	function setSrategy(address _strat) external onlyFactory {
		strategy = _strat;
	}

	function setVaultCap(uint256 newMax) external onlyAdmin {
			vaultCap = newMax;
	}

	function setIndividualCap(uint256 newMax) external onlyAdmin {
			individualCap = newMax;
	}
	
	function checkCap(uint256 amount0,uint256 amount1) public view returns(uint256){
		uint256 price = IStrategy(strategy).getPrice();
		uint256 currentCap = (amount0 * price ) + amount1;
		return( currentCap );
	}

    function emergencyBurn() external onlyAdmin {
    	IStrategy(strategy).callFunds();
    }
		
 	function sweep( address _token, uint256 amount) external  onlyAdmin {
		require (msg.sender != address(0), "s0x");
		require(_token != token0 && _token != token1, 'tok');
        IERC20(_token).safeTransfer(msg.sender, amount);
    }	

    function deposit( uint256 amountIn0, uint256 amountIn1) external onlyActive {
    	
		require(amountIn0>0 || amountIn1>0, 'amt0');
		require(msg.sender != address(0), 'd0x');
       	(uint256 shares, uint256 amount0, uint256 amount1) = calcPositionShares(amountIn0, amountIn1);
		require (shares > 0, 'share<=0'); //affirm the shares
		
        _mint( msg.sender, shares);
        if (amount0 > 0) IERC20(token0).safeTransferFrom(msg.sender, address(this), amount0);
        if (amount1 > 0) IERC20(token1).safeTransferFrom(msg.sender, address(this), amount1);
        
        isUsed = false;

        emit Deposit(msg.sender, shares, amount0, amount1);
        
    }     

	///@notice withdraw function
	///@dev withdraw function 
    function withdraw( uint8 percent ) external 
    	returns (uint256 amount0, uint256 amount1) 
	{
		require (msg.sender != address(0), "w0x");
		require ( percent <= 100, 'pc');
        uint256 shares = balanceOf(msg.sender) * percent / 100;
		uint256 _totalSupply = totalSupply();
		require(shares > 0 && shares <= _totalSupply, 'shares');
        _burn(msg.sender, shares);	// burn user's share
        
        (uint256 stratTVL0, uint256 stratTVL1) = IStrategy(strategy).getTotalAmounts();
        (uint256 vaultTVL0, uint256 vaultTVL1) = ( getbalance0(), getbalance1());
        (uint256 total0, uint256 total1) = (stratTVL0 + vaultTVL0, stratTVL1 + vaultTVL1);
        amount0 = total0 * shares / _totalSupply; 
        amount1 = total1 * shares / _totalSupply; 
        
        bool success = false;
        if (vaultTVL0 < amount0 || vaultTVL1 < amount1 )  {   // fund in this vault is not enough to withdraw
        	success = IStrategy(strategy).vaultWithdraw(msg.sender, amount0, amount1);
	        (uint256 newBalance0, uint256 newBalance1) = ( getbalance0(), getbalance1());
			require(newBalance0 >= amount0 && newBalance1 >= amount1, "new balance");
        } else {
        	success = true;
        }
        
        require(success, "strat fund");
        
		IERC20(token0).safeTransfer(msg.sender, amount0);
		IERC20(token1).safeTransfer(msg.sender, amount1);

        // get tvl in stratVault 
        // get tvl sitting in this vault as they are already in 
        // total0, total1 
        // calc withdrawal amount0 , amount1 based on shares, totalsupply and total0, total1 and best ratio for vault
        // best ratio is r(todo)
        // 	amount0 = total0 * shares / totalsupply fun(r)
        // 	amount1 = total1 * shares / totalsupply fun(r)
        //  check if balance0 here is enough for withdraw, otherwise, call strategy to do re-balance or remove some position and send to this vault
        // **check if balance here in this vault is enough again. 
        // send fund to withdrawer.

        
		//log
        emit Withdraw(msg.sender, shares, amount0, amount1);
    }
	
	function calcPositionShares( uint256 amountIn0, uint256 amountIn1) 
		public 
		returns(uint256 shares, uint256 amount0, uint256 amount1)
	{
		(amount0, amount1) = (amountIn0, amountIn1);
		uint256 p = IStrategy(strategy).getPrice(); 
		require(p>0,'p0');

		uint256 totalSupply = totalSupply();
		if (totalSupply == 0 ) {
			shares = calcShare(p, amount0, amount1);
			
		} else {
			( uint256 total0, uint256 total1 ) = IStrategy(strategy).getTotalAmounts();

			total0 += getbalance0();
			total1 += getbalance1();
			
			uint256 estTotalSupply = calcShare(p, total0, total1);
			uint256 estShare =   calcShare(p, amountIn0, amountIn1); 
			shares = estShare * totalSupply / estTotalSupply ;
		}
	}
	
	function calcShare(uint256 p, uint256 a0, uint256 a1) public pure returns(uint256){
		return(a0 + a1/p); 
	}
	
    
    //move funds to strategy
    function moveFunds() external onlyStrategy {
    	//idea: factory.getStrategy() instead of onlyStrategy
		uint256 a0 = getbalance0();
		uint256 a1 = getbalance1();
		if (a0 > 0) IERC20(token0).safeTransfer(strategy, a0);
        if (a1 > 0) IERC20(token1).safeTransfer(strategy, a1);
        isUsed = true;
    }
    
    function getbalance0() public view returns(uint256) {
    	return(IERC20(token0).balanceOf(address(this)));
    }
  
    function getbalance1() public view returns(uint256) {
    	return(IERC20(token1).balanceOf(address(this)));
    }
  

}
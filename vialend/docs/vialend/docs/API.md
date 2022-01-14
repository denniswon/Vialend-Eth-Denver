
### Deposit
	`contract` 
		 ViaVault.sol 

	`method ` 
		deposit( uint256 amountIn0, uint256 amountIn1) external onlyActive 

	`input params`	
		uint256 amountIn0  --  amount in token0
		uint256 amountIn1  --  amount in token1
	
	`event` 
		event Deposit( address indexed sender, uint256 shares,  uint256 amount0, uint256 amount1  );
	
	`go example` 
		/scripts/project/testVault.go   func Deposit()
		
### Withdraw
	`contract` 
		 ViaVault.sol 

	`method ` 
		withdraw( uint8 percent) external

	`input params`	
		uint8 percent  --  percentage of user's total shares. 1 - 100 
	
	`event` 
		event PendingWithdraw(msg.sender, pending)   
			pending = true  withdraw is put in pending list until next rebalance.
			pending = false withdraw has been processed. 
	
		if pending = false, another event will be emitted:
		event Withdraw( address indexed sender, uint256 shares,  uint256 amount0, uint256 amount1);
     
    );

	`go example` 
		/scripts/project/testVault.go   func Withdraw()
		
		
		
### Rebalance 
	`contract` 
		 VaultStrategy.sol 

	`method ` 
		function rebalance(int24 newLow,  int24 newHigh) external onlyActive

	`input params`	
		int24 newLow  	-- new tick low
		int24 newHigh  --  new tick high
	
	`event` 
		event Rebalance(address indexed,uint256 u0,  uint256 u1, uint256 c0,  uint256 c1);
     
    );

	`go example` 
		/scripts/project/testVault.go   func Rebalance()
		
		
		

		
		
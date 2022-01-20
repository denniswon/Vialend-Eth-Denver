# v2.0
## ----- Quick Note ---- 

Changes from old methods and properties:
	
	`maxTotalSupply`  - public `vaultCap`   - ViaVault
	`governance`   	- function getAdmin()  - VaultFactory 
	`team`		- function  getTeam()  - VaultFactory  
	`uniPortion`	- public uniPortion  - VaultStrategy 
	`protocolFee`     - protocolFeeRate - VaultStrategy 
	
	new：
	
	public compPortion  - VaultStrategy  
	     compound potion rate， 可以从0 - 100
	
	uint256 public individualCap;   - ViaVault  
		用户个人最大可存入的数额 
	
	function getPrice() public method.   -- VaultStrategy 
		得到 两个token 之间的价格， 比如 weth/ usdc ，  getPrice() 返回4000，  即为 1 weth = 4000 usdc 

## ----- Pseudo Scripts ----- 

### Current Deposit  (user side)
	i.e. weth/usdc
	
	price = strategy.getPrice()
	amount0 = weth . balanceOf(user address)
	amount1= usdc.balanceOf(user address)
	totalInUSDC = amount0 * price + amount1
	current deposit display = totalInUSDC / 10^usdc.decimals

### GetTVL  
	get balanceOf vault
	get balanceOf strategy
	get position in uni v3   
	get position in lending
	get position in squeeth
	get position in dydx
	
	GetTVL = accumulate above 
	
	referal: func GetTVL() in testVault.go
	
		
## ----- Contracts ----- 
### ViaVault.sol
	. funds vault. 
	. manage funds deposit/withdraw and calculate shares
	 
### VaultStrategy.sol
	. manage strategies
	. manage rebalance 
	 
### VaultFactory.sol
	. administration of vaults and strategies
	. vault and strategy registration
	. set/get admin and team 
	. change vault status
	. call deployer to create vault in air
	
## ----- Public features ----- 
### VaultFactory
#### public properties

	mapping (address => mapping(address =>uint )) public stat;   // 1: active ; 2 pending;  3 emergency ; 4 abandoned     
    mapping (address => address) public pairs;	// strategy <->  vault   
    VaultReg[] public vaults;
	    struct VaultReg {
	    	address strategy;
			address vault;
	    }

    

#### public methods
    function getAdmin() external view returns(address);
    function getTeam() external view returns(address);
    function setTeam(address _team) external;
  	function changeStat(address _strategy, address _vault, uint _stat) external;   // change status of vault and strategy
	function getCount() external; 	 // get stored vaults array size
	function getStat(address) external view returns(uint);		// get status of vault and strategy. 1 = active
	function getPair0(address _addr) external view returns(address);		// get vault or strategy address by given strategy or vault address


### ViaVault 
#### public properties
   
    address public immutable factory;  // vaultFactory 
    address public immutable token0;   // token0
	address public immutable token1;   // token1

	uint256 public vaultCap;		// maxTotalSupply
	uint256 public individualCap;	// max Individual supply : todo

	Withdrawal[] public wds; 		// pending withdrawal array
		struct Withdrawal {
			address recipient;
			uint8 percent;
		}

	
#### public methods

	function deposit( uint256 amountIn0, uint256 amountIn1) external
	function withdraw( uint8 percent ) external nonReentrant returns (uint256 amount0, uint256 amount1)
	function setVaultCap(uint256 newMax) external onlyAdmin
	function setIndividualCap(uint256 newMax) external onlyAdmin
	function checkCap(uint256 amount0,uint256 amount1) public view returns(uint256)
	function sweep( address _token, uint256 amount) external  onlyAdmin
	
### VaultStrategy 
#### public properties
	address public creator;			// the strategy creator. if it's deployed by admin team, creator is admin  team
	address public immutable factory;	// the vault factory
	address public immutable _WETH;
    address public immutable token0;         // underlying token0
    address public immutable token1;         // underlying token1
	address public protocol;			// where fee cuts to protocol 
	ICEth public immutable _CETH;
    IUniswapV3Pool public immutable pool;        // get by uni factory, token0, token1, feetier
    uint128 public immutable quoteamount;  		// base amount for calculating price, based on token0 decimal, ie: 1e18 for eth, 1e8 for wbtc
    uint8 public uniPortion;       // uniswap portion ratio
    uint8 public compPortion;       // compound portion ratio
    uint8 public protocolFeeRate;		// 0 - 20% of profit
	uint8 public motivatorFeeRate;		// 0- 10% from profit for keeping system running by press buttons
	uint32 public twapDuration;        // oracle twap durantion
    int24 public tickSpacing;
    int24 public cLow;
    int24 public cHigh;
    int24 public maxTwapDeviation;     // for twap     

    mapping (address => address) public _CTOKEN;
	mapping (address => uint8) public Decimals;

	// to be working
    mapping (uint => address[] ) public motivator;		// who helped to triggering buttons e.g. motivator[1].push( address )  1=rebalance
	address[] public motivators;			
	
#### public methods
	 function rebalance(int24 newLow, int24 newHigh) external nonReentrant
	 function setMaxTwapDeviation(int24 _maxTwapDeviation) external onlyCreator
	 function setTwapDuration(uint32 _twapDuration) external onlyCreator
     function setPortionRatio(uint8 uni, uint8 comp) external onlyCreator {
	 function alloc() public returns ( bool )
	 function calcShares( uint256 totalSupply, uint256 amountIn0, uint256 amountIn1) public returns(uint256 shares, uint256 amount0, uint256 amount1)
	 function getUniAmounts(int24 tickLower, int24 tickUpper)  public view returns (uint256 amount0, uint256 amount1) 
	 function getPrice() public view returns(uint256) -- get price at twap




## ----- Events ----- 
### 1
### 2
### 3


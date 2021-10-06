// SPDX-License-Identifier: MIT
pragma solidity >=0.5.0;

import "../@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../@openzeppelin/contracts/token/ERC20/ERC20.sol";

import "../interfaces/IWETH9.sol";
import "../interfaces/IcEth.sol";
import "../interfaces/IcERC20.sol";



import "./ownable.sol";

 contract ViaCompound is Ownable {
  	
	
    using SafeMath for uint256;
    
	event MyLog(string, uint256);
	event MyLog2(string, uint256,uint256);

	IcErc20 internal  CToken0; 
 	IcErc20 internal  CToken1; 
 	IcEther internal  CEther; 
  	IWETH9 internal WETH;

///compound procedures

function supplyEthToCompound(address payable _cEtherContract, address _ctoken)
        public
        
        payable
        returns (bool)
    {
        // Create a reference to the corresponding cToken contract
        IcEther cEth = IcEther(_cEtherContract);
        IcErc20 ctoken = IcErc20(_ctoken);

        // Amount of current exchange rate from cToken to underlying
         // exchange rate (how much ETH one cETH is worth) 
        uint256 exchangeRateMantissa = ctoken.exchangeRateCurrent();
        emit MyLog("Exchange Rate (scaled up by 1e18): ", exchangeRateMantissa);

        // Amount added to you supply balance this block
        uint256 supplyRateMantissa = ctoken.supplyRatePerBlock();
        emit MyLog("Supply Rate: (scaled up by 1e18)", supplyRateMantissa);

        cEth.mint{gas:250000,value:address(this).balance}();
        return true;
    }

    function supplyErc20ToCompound(
        address _erc20Contract,
        address _cErc20Contract,
        uint256 _numTokensToSupply
    ) public  returns (uint) {
    	
		
		 require(_numTokensToSupply <=  IERC20(_erc20Contract).balanceOf(address(this)) ,"balance");
        // Create a reference to the underlying asset contract, like DAI.
        ERC20 underlying = ERC20(_erc20Contract);

        // Create a reference to the corresponding cToken contract, like cDAI
        IcErc20 cToken = IcErc20(_cErc20Contract);

        // Amount of current exchange rate from cToken to underlying
        // exchange rate (how much underlying token one cToken is worth) 
        uint256 exchangeRateMantissa = cToken.exchangeRateCurrent();
        emit MyLog("Exchange Rate (scaled up): ", exchangeRateMantissa);

        // Amount added to you supply balance this block
        uint256 supplyRateMantissa = cToken.supplyRatePerBlock();
        emit MyLog("Supply Rate: (scaled up)", supplyRateMantissa);

        // Approve transfer on the ERC20 contract
        underlying.approve(_cErc20Contract, _numTokensToSupply);

        // Mint cTokens
        uint mintResult = cToken.mint(_numTokensToSupply);
        return mintResult;
    }
/*
yes, I suggest you keep track of the initial deposit amount and the cToken amount
you can do this using a mapping
// user -> underlying -> amount
mapping(address => mapping(address => uint)) public underlyingBalances;

// user -> cToken -> amount
mapping(address => mapping(address => uint)) public cTokenBalances;
you can use balanceOf before and after the mint / redeem transactions to do your accounting
so the contract holds the cTokens and does all the accounting for the users

*/
    function redeemCErc20Tokens(
        uint256 amount,
        address _cErc20Contract
    ) public   {

        // `amount` is scaled up, see decimal table here:
        // https://compound.finance/docs#protocol-math

        uint256 redeemResult;

        redeemResult = IcErc20(_cErc20Contract).redeem(amount);
        
        emit MyLog("If this is not 0, there was an error", redeemResult);

        
    }

    function redeemCEth(
        uint256 amount,
        address _cEtherContract
    ) public   {

        // `amount` is scaled up by 1e18 to avoid decimals
        uint256 redeemResult;

        redeemResult = IcEther(_cEtherContract).redeem(amount);
        
        emit MyLog("If this is not 0, there was an error", redeemResult);


    }

	function _wrap() public {
	
	    if (address(this).balance != 0) {
	        WETH.deposit{value:address(this).balance}();
	    }   
		
	}


	function _unwrap(uint256 Amount) public
	{
	   // address payable sender= msg.sender;
	
	    if (Amount != 0) {
	        WETH.withdraw(Amount);
	    }
	}   
    
    
 	function getCAmounts() public view returns (uint256 amountA, uint256 amountB) {
		
		amountA = CToken0.balanceOf(address(this) ) ;
		amountB = CToken1.balanceOf(address(this) ) ;

	}
	
	
   function getLendingAmounts() public view returns(uint256 , uint256, uint256, uint256 ){

    	(uint256 cAmount0, uint256 cAmount1) = getCAmounts();
		
		
        //uint8 decimals0 = EIP20Interface(CToken0.underlying()).decimals();
        
        //uint8 decimals1 = EIP20Interface(CToken1.underlying()).decimals();
            
		//require(token0.decimals() ==       underlyingDecimals = EIP20Interface(cErc20.underlying()).decimals();
		
		//note: exchangeRateStored is   exchangeRate * 10 ** underlying decimals
		uint32 decimal0 = 18; // to do
		uint32 decimal1 = 6;  // to do 
		
		//oneCTokenInUnderlying = exchangeRateCurrent / (1 * 10 ^ (18 + underlyingDecimals - cTokenDecimals))
        uint256 amount0 = cAmount0.mul(CToken0.exchangeRateStored()).div(10**decimal0) ;
        uint256 amount1 = cAmount1.mul(CToken1.exchangeRateStored()).div(10** (18 + decimal1 - 8) ) ;

		
		return(amount0,amount1,CToken0.exchangeRateStored(),CToken1.exchangeRateStored());
    }
	

function lendingSupply(address underlying0, address underlying1,uint256 amount0, uint256 amount1) internal returns(bool) {
		
		address weth = address(WETH);
		address cToken0 = address(CToken0);
		address cToken1 = address(CToken1);
		
        if (underlying0 == weth ) {
			
			_unwrap(amount0);
	        supplyEthToCompound( payable(address(CEther)), cToken1 );
	        supplyErc20ToCompound( underlying1,   cToken1, amount1);


        } else if (underlying1 == weth ) {
			_unwrap(amount1);
	        supplyEthToCompound( payable(address(CEther)), cToken0 );
	        supplyErc20ToCompound( underlying0,   cToken0, amount0);
        } else {

	        supplyErc20ToCompound( underlying0,  cToken0, amount0);
	        supplyErc20ToCompound( underlying1,  cToken1, amount1);
        }

		return true;
	}


	function removeLending(address underlying0, address underlying1, uint256 amount0, uint256 amount1) internal {
        
		address weth = address(WETH);
		address cToken0 = address(CToken0);
		address cToken1 = address(CToken1);
		//address cEther = address(CEther);

        // Withdraw all current supply from lending pool
        
		
        if (underlying0 == weth ) {
        	redeemCEth(amount0,cToken0);
			_wrap();
	        redeemCErc20Tokens( amount1, cToken1 );


        } else if (underlying1 == weth ) {
        	redeemCEth(amount1,cToken1);
			_wrap();
	        redeemCErc20Tokens( amount0,  cToken0 );
        } else {
	        redeemCErc20Tokens( amount1,  cToken1 );
	        redeemCErc20Tokens( amount0,  cToken0 );
        }
        
  		

	}


}
// SPDX-License-Identifier: MIT
pragma solidity 0.7.6;

import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

import "interfaces/IWETH9.sol";
import "interfaces/IcEth.sol";
import "interfaces/IcERC20.sol";



contract VialendComp {
    event MyLog(string, uint256);
    
    using SafeERC20 for IERC20;

   	WETH9_ internal WETH;
	
	address public owner;
 
    constructor (address payable WETHAddress ) {
        owner = msg.sender;
        
        require(WETHAddress != address(0), "WETH is the zero address");
	
	    WETH = WETH9_(WETHAddress);
    }
    
    function supplyEthToCompound(address payable _cEtherContract)
        public
        payable
        returns (bool)
    {
        // Create a reference to the corresponding cToken contract
        IcEth cToken = IcEth(_cEtherContract);

        // Amount of current exchange rate from cToken to underlying
        uint256 exchangeRateMantissa = cToken.exchangeRateCurrent();
        emit MyLog("Exchange Rate (scaled up by 1e18): ", exchangeRateMantissa);

        // Amount added to you supply balance this block
        uint256 supplyRateMantissa = cToken.supplyRatePerBlock();
        emit MyLog("Supply Rate: (scaled up by 1e18)", supplyRateMantissa);

        cToken.mint{gas:250000,value:msg.value}();
        return true;
    }

    function supplyErc20ToCompound(
        address _erc20Contract,
        address _cErc20Contract,
        uint256 _numTokensToSupply
    ) public returns (uint) {
    	
		
		require(_numTokensToSupply <=  IERC20(_erc20Contract).balanceOf(address(this)) ,"balance");
        // Create a reference to the underlying asset contract, like DAI.
        ERC20 underlying = ERC20(_erc20Contract);

        // Create a reference to the corresponding cToken contract, like cDAI
        IcERC20 cToken = IcERC20(_cErc20Contract);

        // Amount of current exchange rate from cToken to underlying
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

    function redeemCErc20Tokens(
        uint256 amount,
        bool redeemType,
        address _cErc20Contract
    ) public returns (bool) {
        // Create a reference to the corresponding cToken contract, like cDAI
        IcERC20 cToken = IcERC20(_cErc20Contract);

        // `amount` is scaled up, see decimal table here:
        // https://compound.finance/docs#protocol-math

        uint256 redeemResult;

        if (redeemType == true) {
            // Retrieve your asset based on a cToken amount
            redeemResult = cToken.redeem(amount);
        } else {
            // Retrieve your asset based on an amount of the asset
            redeemResult = cToken.redeemUnderlying(amount);
        }

        // Error codes are listed here:
        // https://compound.finance/developers/ctokens#ctoken-error-codes
        emit MyLog("If this is not 0, there was an error", redeemResult);

        return true;
    }

    function redeemCEth(
        uint256 amount,
        bool redeemType,
        address _cEtherContract
    ) public returns (bool) {
        // Create a reference to the corresponding cToken contract
        IcEth cToken = IcEth(_cEtherContract);

        // `amount` is scaled up by 1e18 to avoid decimals

        uint256 redeemResult;

        if (redeemType == true) {
            // Retrieve your asset based on a cToken amount
            redeemResult = cToken.redeem(amount);
        } else {
            // Retrieve your asset based on an amount of the asset
            redeemResult = cToken.redeemUnderlying(amount);
        }

        // Error codes are listed here:
        // https://compound.finance/docs/ctokens#ctoken-error-codes
        emit MyLog("If this is not 0, there was an error", redeemResult);

        return true;
    }

	// to do : onlyGovernance
 	function withdrawERC20(
        uint256 amount,
        address erc20,
        address to
    ) external  onlyOwner {
        
        require(amount > 0, "amount");

        require(to != address(this) && to !=erc20 ,"to");
        
        
        
        IERC20(erc20).safeTransfer(to, amount);
        

        emit MyLog("Withdraw Erc20:", amount);
        
        
    }


	// to do : onlyGovernance
 	function withdrawEth(   uint256 amount  ) external  onlyOwner {
        
         
        require(amount <= getETHBalance(), "amount");

        msg.sender.transfer(amount);

        emit MyLog("WithdrawEth msg.sender:", amount);
        
    }

    
    function getETHBalance() public view returns (uint256) {
         return address(this).balance;
    }


 

	function wrap() payable public  {
	    uint256 ETHAmount =msg.value;
	
	    //create WETH from ETH
	    if (msg.value != 0) {
	        WETH.deposit{value:msg.value}();
	    }   
	    require(WETH.balanceOf(address(this))>=ETHAmount,"Ethereum not deposited");
	}


	function unwrap(uint256 Amount) public 
	{
	    address payable sender= msg.sender;
	
	    if (Amount != 0) {
	        WETH.withdraw(Amount);
	        sender.transfer(address(this).balance);
	    }
	}   
    


   modifier onlyOwner () {
       require(msg.sender == owner, "This can only be called by the contract owner!");
       _;
     }

    // This is needed to receive ETH when calling `redeemCEth`
//    function() external payable {}
	fallback() external payable {}
	 receive() external payable {}
	
	
}
// SPDX-License-Identifier: MIT

pragma solidity >=0.8.8;

import "@aave/interfaces/Ipool.sol";
import "@aave/flashloan/base/FlashLoanReceiverBase.sol";
import "@uniswap/smart-router-contracts/contracts/interfaces/ISwapRouter02.sol";
import "@uniswap/v3-periphery/contracts/libraries/TransferHelper.sol";
import "interfaces/IAaveShorter.sol";

/**
 * @notice Contract for creating a short position using Aave for borrowing and Uniswap for swap.
 */
contract AaveShorter is FlashLoanReceiverBase, IAaveShorter {
    address SWAP_ROUTER = "0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45";

    address colateral;
    address shorting;
    address sender;
    uint256 amount;
    IPool aavePool;
    
    function short(
        address _aavePool,
		address _colateral,
        address _shorting,
		uint256 _amount,
        unit256 _shortSize

    ) external returns(bool) {
        colateral = _colateral;
        shorting = _shorting;
        amount = _amount;
        aavePool = IPool(_aavePool);
        sender = msg.sender;

        bytes params;

        // flash loan the amount times the size of the short, swap for more colateral and deposit.
        aavePool.flashLoan(
            address(this),
            [shorting],
            [_ammount * _shortSize],     // for a 2x short the borrowed asset must be doubled with respect to the colateral ammount size
            [1],                         // 1 stable, 2 variable
            sender,
            params,
            0
        );
    }
    
    /**
    * @notice Executes an operation after receiving the flash-borrowed assets
    * @dev Ensure that the contract can return the debt + premium, e.g., has
    *      enough funds to repay and has approved the Pool to pull the total amount
    * @param assets The addresses of the flash-borrowed assets
    * @param amounts The amounts of the flash-borrowed assets
    * @param premiums The fee of each flash-borrowed asset
    * @param initiator The address of the flashloan initiator
    * @param params The byte-encoded params passed when initiating the flashloan
    * @return True if the execution of the operation succeeds, false otherwise
    */
    function executeOperation(
        address[] calldata assets,
        uint256[] calldata amounts,
        uint256[] calldata premiums,
        address initiator,
        bytes calldata params
    ) external returns (bool) {
        ISwapRouter02 router = ISwapRouter02(SWAP_ROUTER);
        
        // approve
        TransferHelper.safeApprove(shorting, address(SWAP_ROUTER), amounts);

        // swap
        ISwapRouter02.ExactInputSingleParams memory params =
            ISwapRouter.ExactInputSingleParams({
                tokenIn: shorting,
                tokenOut: colateral,
                fee: 3000,
                recipient: msg.sender,
                deadline: block.timestamp,
                amountIn: amountIn,
                amountOutMinimum: 0,   // TODO: Does this value need to be set.
                sqrtPriceLimitX96: 0
            });

        uint256 out = router.exactInputSingle(params);

        // deposit
        aavePool.deposit(colateral, amount + out, sender, 0);

        return true;
    }   	
}

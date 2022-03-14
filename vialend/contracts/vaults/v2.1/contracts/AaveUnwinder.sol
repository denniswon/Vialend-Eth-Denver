// SPDX-License-Identifier: MIT

pragma solidity >=0.8.8;

import "@aave/interfaces/Ipool.sol";
import "@aave/flashloan/base/FlashLoanReceiverBase.sol";
import "@uniswap/smart-router-contracts/contracts/interfaces/ISwapRouter02.sol";
import "@uniswap/v3-periphery/contracts/libraries/TransferHelper.sol";
import "interfaces/IAaveUnwinder.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";


/**
 * @notice Contract for creating a short position using Aave for borrowing and Uniswap for swap.
 */
contract AaveUnwinder is FlashLoanReceiverBase, IAaveUnwinder {
    address SWAP_ROUTER = "0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45";

    address colateral;
    address unwinding;
    address sender;
    uint256 amount;
    IPool aavePool;
    
    function unwind(
        address _aavePool,
		address _colateral,
        address _unwinding,
		uint256 _amount,
        unit256 _unwindSize

    ) external returns(bool) {
        colateral = _colateral;
        unwinding = _unwinding;
        amount = _amount;
        aavePool = IPool(_aavePool);
        sender = msg.sender;

        bytes params;

        // flash loan the amount, swap to pay back debt.
        aavePool.flashLoan(
            address(this),
            [colateral],
            [_unwindSize],      // unwind by the amount specified 
            [0],                // 0 needs to be repaied
            sender,
            params,
            0
        );

        return true;
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
                tokenIn: assets[0],
                tokenOut: unwinding,
                fee: 3000,
                recipient: initiator,
                deadline: block.timestamp,
                amountIn: amounts[0],
                amountOutMinimum: 0,   // TODO: Does this value need to be set.
                sqrtPriceLimitX96: 0
            });

        uint256 out = router.exactInputSingle(params);

        // repay debts 
        uint256 rateMode = 2;          // Stable: 1, Variable: 2
        aavePool.repay(unwinding, out, rateMode, sender);

        // withdraw colateral to repay flash
        uint flashRepay = ammount[0].add(premiums[0]);
        aavePool.withdraw(colateral, flashRepay, sender);

        // approve contract to pull funds to repay flash
        IERC20(assets[0]).approve(address(aavePool), flashRepay);

        return true;
    } 
}

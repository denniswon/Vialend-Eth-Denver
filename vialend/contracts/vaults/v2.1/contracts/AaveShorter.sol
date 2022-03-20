// SPDX-License-Identifier: MIT

pragma solidity 0.8.10;

import {TransferHelper} from "@uniswap/v3-periphery/contracts/libraries/TransferHelper.sol";

import {IPoolAddressesProvider} from "https://github.com/aave/aave-v3-core/blob/master/contracts/interfaces/IPoolAddressesProvider.sol";
import {IPool} from "https://github.com/aave/aave-v3-core/blob/master/contracts/interfaces/IPool.sol";
import {DataTypes} from "https://github.com/aave/aave-v3-core/blob/master/contracts/protocol/libraries/types/DataTypes.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";


contract AaveHelper {

    /**
     * Create a short on aave with the given collateral and short addresses.
     */
    function short(
        address _aavePoolProvider,
        address _callback,
        address _collateral,
        address _shorting,
        uint256 _collateralSize,  // 1,000,000 usdc // 2x short requres: collateral / ethprice * 2 = shortsize
        uint256 _shortSize       // 666.66 eth
    ) internal returns(bool) {

        // Approve call back to transfer collateral
        TransferHelper.safeApprove(_collateral, _callback, _collateralSize);

        // set up flash loan
        bytes memory params = abi.encode(_collateral, _collateralSize);
        address[] memory assets = new address[](1);
        assets[0] = _shorting;
        uint256[] memory sizes = new uint256[](1);
        sizes[0] = _shortSize;
        uint256[] memory modes = new uint256[](1);
        modes[0] = 2;           // 0 repay, 1 stable, 2 variable

        // flash loan the amount times the size of the short, swap for more collateral and deposit.
        IPool(IPoolAddressesProvider(_aavePoolProvider).getPool()).flashLoan(
            _callback,
            assets,
            sizes,
            modes,              
            address(this),
            params,
            0
        );

        return true;
    }

    /**
     * Unwind the entire short using the balances of aToken and debtTokens at the given addresses.
     */
    function unwind(                 // short (eth), total colat (usdc) => short unwind, usdc = aUsdc * index
        address _aavePoolProvider,
        address _callback,
        address _collateral,
        address _shorted,
        uint256 _ethPerUsdc  // 10^18 / (10^6) -> 10^12 > 1
    ) internal returns(bool) {

        // Get the atoken addresses
        IPool pool = IPool(IPoolAddressesProvider(_aavePoolProvider).getPool());
        DataTypes.ReserveData memory collateralReserves = pool.getReserveData(_collateral);
        address aToken = collateralReserves.aTokenAddress;

        // Approve transfer of aTokens to callback
        TransferHelper.safeApprove(aToken, _callback, IERC20(aToken).balanceOf(address(this)));

        uint256 debt = getDebt(pool, _shorted);

        // (address _collateral, uint256 _collateralSizeOfDebt) = abi.decode(params, (address, uint256));
        bytes memory params = abi.encode(_collateral, debt / _ethPerUsdc);  // 10^18 / (10^12) -> 10^6

        address[] memory assets = new address[](1);
        assets[0] = _shorted;
        uint256[] memory sizes = new uint256[](1);
        sizes[0] = debt;
        uint256[] memory modes = new uint256[](1);
        modes[0] = 0;   // 0 repay, 1 stable, 2 variable

        // flash loan the amount times the size of the short, swap for more collateral and deposit.
        pool.flashLoan(
            _callback,
            assets,
            sizes,
            modes,              
            address(this),
            params,
            0
        );

        return true;
    }

    /**
     * Get debt for given address on Aave.
     */
    function getDebt(IPool pool, address _shorted) internal returns(uint256) {
        // get debt amount to flash loan
        DataTypes.ReserveData memory debtReserves = pool.getReserveData(_shorted);
        address debtTokenAddress = debtReserves.variableDebtTokenAddress;
        return IERC20(debtTokenAddress).balanceOf(address(this)); // denominated in underlying asset not aToken.
    }
}

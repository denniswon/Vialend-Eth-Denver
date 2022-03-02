// SPDX-License-Identifier: MIT

pragma solidity 0.8.10;

import {IPoolAddressesProvider} from 'https://github.com/aave/aave-v3-core/blob/master/contracts/interfaces/IPoolAddressesProvider.sol';
import {IPool} from 'https://github.com/aave/aave-v3-core/blob/master/contracts/interfaces/IPool.sol';

contract AaveHelper {
        // option 2 this method goes to vault. 
    function short(
        address _aavePoolProvider,
        address _callback,
        address _colateral,
        address _shorting,
        uint256 _colateralSize,  // 1,000,000 usdc // 2x short requres: colateral / ethprice * 2 = shortsize 
        uint256 _shortSize       // 666.66 eth
    ) internal returns(bool) {

        bytes memory params = abi.encode(_colateral, _colateralSize);
        address[] memory assets = new address[](1);
        assets[0] = _shorting;
        uint256[] memory sizes = new uint256[](1);
        sizes[0] = _shortSize;
        uint256[] memory modes = new uint256[](1);
        modes[0] = 1;


        // flash loan the amount times the size of the short, swap for more colateral and deposit.
        IPool(IPoolAddressesProvider(_aavePoolProvider).getPool()).flashLoan(
            _callback,          // pass call back contract address
            assets,
            sizes,
            modes,                // 0 repay, 1 stable, 2 variable
            msg.sender,
            params,
            0
        );

        return true;
    }
}
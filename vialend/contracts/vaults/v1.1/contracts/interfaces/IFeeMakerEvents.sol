// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.5.6;

/// @title Events emitted by feemaker
/// @notice Contains all events emitted by feemaker
interface IFeeMakerEvents {
	
	event GeneralS(string subject, string value);
	
	event GeneralB(string subject, uint256 value);

	event GeneralA(string subject, address value);
	
    event Deposit(
        address indexed sender,
        address indexed to,
        uint256 shares,
        uint256 amount0,
        uint256 amount1
    );

    event Withdraw(
        address indexed sender,
        address indexed to,
        uint256 shares,
        uint256 amount0,
        uint256 amount1,
        uint256 poolamount0,
        uint256 poolamount1,
        uint256 lendingamount0,
        uint256 lendingamount1
    );

    event CollectFees(
        uint256 feesToVault0,
        uint256 feesToVault1,
        uint256 feesToProtocol0,
        uint256 feesToProtocol1
    );

	event RebalanceLog(uint128 liquidity, uint256 newBalance0, uint256 newBalance1);

    event Snapshot(int24 tick, uint256 totalAmount0, uint256 totalAmount1, uint256 totalSupply);
    
    event ErrorLogging(string reason);

}
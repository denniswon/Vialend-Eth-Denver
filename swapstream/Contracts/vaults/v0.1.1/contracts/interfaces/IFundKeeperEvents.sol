// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.4.21 <= 0.8.6;

/// @title Events emitted by fundkeeper
/// @notice Contains all events emitted by fundkeeper
interface IFundKeeperEvents {
	
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
        string nameToken0,
        string nameToken1
    );

    event CollectFees(
        uint256 feesToVault0,
        uint256 feesToVault1,
        uint256 feesToProtocol0,
        uint256 feesToProtocol1
    );

    event Snapshot(int24 tick, uint256 totalAmount0, uint256 totalAmount1, uint256 totalSupply);

}
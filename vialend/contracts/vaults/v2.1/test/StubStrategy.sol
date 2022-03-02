// SPDX-License-Identifier: MIT

pragma solidity 0.8.10;


import {AaveHelper} from "../contracts/AaveHelper.sol";
import "@uniswap/v3-periphery/contracts/libraries/TransferHelper.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
* Stub strategy uses the AaveHelper in the expected way the actual strategy will use it. 
*/
contract StubStrategy is AaveHelper {

    address public aavePoolProvider = 0xA55125A90d75a95EC00130E8E8C197dB5641Eb19;
    address public shortCallback = 0xBA9410f84791A185Ee7088e0B2486BDa38c6515F;
    address public aaveUSDC = 0x5B8B635c2665791cf62fe429cB149EaB42A3cEd8;
    address public aaveEth = 0x98a5F1520f7F7fb1e83Fe3398f9aBd151f8C65ed;
    uint256 public usdcPerEth = 2600;
    address public owner;

    constructor (address _owner) {
        owner = _owner;
    }
    
    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    function testHelper(uint256 amount) public onlyOwner {
        TransferHelper.safeTransferFrom(aaveUSDC, msg.sender, address(this), amount);
        TransferHelper.safeApprove(aaveUSDC, shortCallback, amount + IERC20(aaveUSDC).allowance(address(this), shortCallback));
        short(aavePoolProvider, shortCallback, aaveUSDC, aaveEth, amount, amount / usdcPerEth * 2);
    }

    function setUsdcPerEth(uint256 _usdcPerEth) public onlyOwner {
        usdcPerEth = _usdcPerEth;
    }

    function setShortCallback(address _callback) public onlyOwner {
        shortCallback = _callback;
    }
}
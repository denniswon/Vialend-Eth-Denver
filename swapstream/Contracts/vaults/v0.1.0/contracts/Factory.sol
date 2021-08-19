// SPDX-License-Identifier: Unlicense

pragma solidity 0.8.6;

import "./ownable.sol";


/**
 * @title   Alpha Vault
 * @notice  A vault that provides liquidity on Uniswap V3.

	Factory cannot created new pools until 

 */

contract SwapStreamFactory is Ownable {

    event NewVault(uint vaultId, string name, uint dna);

    uint dnaDigits = 16;
    uint dnaModulus = 10 ** dnaDigits;
    // 1. Define `cooldownTime` here
    uint cooldownTime = 1 days;



    struct Vault {
        address
        address token0;
        address token1;
        uint dna;
        uint32 level;
        uint32 readyTime;
    }

    Vault[] public vaults;

    mapping (uint => address) public vaultToOwner;
    mapping (address => uint) ownerVaultCount;

    function _createVault(string memory _name, uint _dna) internal {
        // 2. Update the following line:
        uint id = vaults.push(Vault(_name, _dna)) - 1;
        vaultToOwner[id] = msg.sender;
        ownerVaultCount[msg.sender]++;
        emit NewVault(id, _name, _dna);
    }

    function _generateRandomDna(string memory _str) private view returns (uint) {
        uint rand = uint(keccak256(abi.encodePacked(_str)));
        return rand % dnaModulus;
    }

    function createRandomVault(string memory _name) public {
        require(ownerVaultCount[msg.sender] == 0);
        uint randDna = _generateRandomDna(_name);
        randDna = randDna - randDna % 100;
        _createVault(_name, randDna);
    }

}

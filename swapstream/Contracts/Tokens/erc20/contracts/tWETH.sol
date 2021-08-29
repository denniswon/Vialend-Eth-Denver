// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;

import "./@OpenZeppelin/contracts/token/ERC20/ERC20.sol";


/**
INFO:
tokens total supply
weth       6,718,113.782417289750857735
USDC  26,512,386,160.20955 

*** DEPLOY NEW Token

1. Change token name and symbol in token.sol 
C:\Users\xdotk\go\src\goblockchain\Practice\Tokens\erc20\contracts\token.sol

2. compile token.sol in wsl   , use solc version 0.7.6
 /// solc-select use 0.7.6
 /// cd /mnt/c/Users/xdotk/go/src/goblockchain/Practice/tokens/erc20/contracts 
	
  solc solc --optimize --overwrite --bin token.sol -o ../build
  solc solc --optimize --overwrite --bin token.sol -o ../build

3. back to windows 
/// C:\Users\xdotk\go\src\goblockchain\Practice\Tokens\erc20\contracts>
abigen --abi=../build/token.abi --bin=../build/token.bin --pkg=api --out=../deploy/token/token.go

4. modify token owner and token address in token_deploy.go 
// C:\Users\xdotk\go\src\goblockchain\Practice\Tokens\erc20\deploy\token_deploy.go

5. run token_deploy.go

** 手动添加token 到uniswap 并创建pool
** pool 创建完成后， 在etherscan里找到pool 地址
** 修改fundkeeper_deploy.go 里的pool地址。 并运行发布新的fundkeeper 合约,  得到新的fundkeeper合约地址
** 修改testFundKeeper.go 里的fundkeeperContract 地址。并运行 


*/
contract tWETH is ERC20 {

	/// SAMPLE 
	//    constructor () ERC20("sample token name", "shortName") {
        //_mint(msg.sender, 1000000000 * (10 ** uint256(decimals())));
        
        /// tWETH
        //_mint(msg.sender, 6718113 * (10 ** 18));

        /// tUSDC
        // _mint(msg.sender, 26512386160 * (10 ** 18));


    constructor () ERC20("TTO wETH", "tWETH") {
		_mint(msg.sender, 6718113 * (10 ** 18));

    }
}
// SPDX-License-Identifier: MIT

pragma solidity >=0.8.8 <0.9.0;

import "./libraries/Errors.sol";
import "./interfaces/IViaVault.sol";
import "./interfaces/Istrategy.sol";
import "./interfaces/IDeployer.sol";


/// @author Vialend Team
/// @notice Vault Manager for vault and it's strategy to be registered and managed. 

contract VaultFactory {
			    
    address private viaAdmin;
    address private team;  // team account which to collect fees separately, default viaAdmin

    mapping (address => mapping(address =>uint )) private stat;   // pending 2, active 1, abandoned 3
    mapping (address => address) public pairVault;	//give strategy get vault
    mapping (address => address) public pairStrategy;	//give vault get strategy

    struct VaultReg {
    	address strategy;
		address vault;
    }
    
    VaultReg[] public vaults;
    
    constructor(address _team ) {
        viaAdmin = msg.sender;
        team = _team;
    }
    
    
    event CreatedVault(address _this, address sender, address strategy, address vault);

	 modifier onlyAdmin() {
        require(msg.sender == viaAdmin);
        _;
    }
    
    function getCount() public view returns(uint) {
	    return vaults.length;
	}
	
    function getAdmin() public view returns(address){
    	return viaAdmin;
    }

    function getTeam() public view returns(address){
    	return team;
    }

    function setTeam(address _team) public onlyAdmin {
    	team = _team;
    }
    
    function createVault(
    	address[10] memory contracts,
		uint	_vaultCap,
		uint	_individualCap,
		uint8  _uniPortion,
		uint8  _compPortion, 
		uint8  _protocolFee,
		uint24	_feetier,
		uint128  _quoteAmount
	
		) external {
    	
	 
	/*	
		address[0] protocolAddr
		address[1] _reserve
		address[2] _weth
		address[3] ceth
		address[4] ctoken0
		address[5] ctoken1
		address[6] _token0 
		address[7] _token1 
		address[8] _strategyD
		address[9] _vaultD
		uint	_vaultCap,
		uint	_individualCap
		uint8  _uniPortion,
		uint8  _compPortion, 
		uint8  _protocolFee,
		uint24	_feetier,
		uint128  _quoteAmount
	
*/  

		address v = IDeployer(contracts[9]).deployVault(address(this),viaAdmin, contracts[6], contracts[7],  _vaultCap, _individualCap);

		require(v != address(0),"v0");
		
		// msg.sender is the creator
		address[10] memory addrs = [contracts[0],msg.sender,contracts[2],contracts[3],contracts[4],contracts[5],contracts[6],contracts[7], v, viaAdmin];
		
		address s = IDeployer(contracts[8]).deployStrategy(addrs,  _uniPortion,  _compPortion,   _protocolFee, _feetier, _quoteAmount);
		
		require(s != address(0),"s0");
		
		IViaVault(v).setSrategy(s);
		
		pair(s, v);

	    emit CreatedVault(address(this), msg.sender, s, v);

    }
    

	///@notice change vault stat
	/// 2: pending (defult), 1: active, 3: abandoned, 0: invalid
  	function changeStat(address _strategy, address _vault, uint _stat) external onlyAdmin {
		require (_stat > 0, "invalid stat");
        stat[_strategy][_vault] = _stat;
    }
  

	function isContract(address addr) public view returns (uint256) {
        // This method relies on extcodesize, which returns size of the contract in
        // construction, since the code is only stored at the end of the
        // constructor execution.

        uint256 size;
        // solhint-disable-next-line no-inline-assembly
        assembly { size := extcodesize(addr) }
        return size;
    }

    function getStat(address _strategy, address _vault ) public view returns(uint) {
    		return(stat[_strategy][_vault]);
    }
    
    function checkActive(address sORv ) public view returns(bool) {
   	    return ( stat[sORv][pairVault[sORv]] == 1 || 
				stat[pairStrategy[sORv]][sORv] == 1 );

    }

    function pair(address _strategy, address _vault ) internal {
    	stat[_strategy][_vault] = 2;  
		pairVault[_strategy] = _vault;
		pairStrategy[_vault] = _strategy;
		vaults.push(VaultReg(_strategy, _vault));
    }


    function repair(address _strategy, address _vault ) public onlyAdmin {
    		pair(_strategy, _vault);
    }
	  
  
}

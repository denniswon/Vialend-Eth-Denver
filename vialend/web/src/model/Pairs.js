/* eslint-disable indent */
/* eslint-disable no-tabs */
import Token from './Token'
function Pairs () {
	this.index = 0
	this.id = 0
	this.token0 = new Token()
	this.token1 = new Token()
	this.feeTier = '0.3%'
	this.currentAPR = ''
	this.capacity = ''
	this.tvl = 0.00
	this.maxTVL = 0.00
	this.ofCapUsed = 0.00
	this.vaultRange = ''
	this.token0Approved = false
	this.token1Approved = false

	this.vaultLending = 0
	this.tvlTotal0 = 0.00
	this.tvlTotal1 = 0.00
	this.tvlTotal0USD = 0
	this.tvlTotal1USD = 0
	this.lendingRatio = 0
	this.uniswapRatio = 0
	this.uniToken0Rate = 0
	this.uniToken1Rate = 0
	this.lendingToken0Rate = 0
	this.lendingToken1Rate = 0
	this.cLow = 0
	this.cHigh = 0
	this.token0BalanceInWallet = 0
	this.token1BalanceInWallet = 0
	this.token0BalanceInVault = 0
	this.token1BalanceInVault = 0
	// Vault contract
	this.vaultAddress = ''
	this.token0Address = ''
	this.token1Address = ''
	this.token0LendingAddress = ''
	this.token1LendingAddress = ''
	this.poolAddress = ''
	this.keeperContract = null
	// Token contract
	this.token0Contract = null
	this.token1Contract = null
	// Lending Info
	this.token0LendingContract = null
	this.token1LendingContract = null
	// myvaluelocked
	this.myShare = 0
	this.totalShares = 0
	this.myValueToken0Locked = 0.00
	this.myValueToken1Locked = 0.00
	this.myValueToken0USDLocked = 0
	this.myValueToken1USDLocked = 0
}

Pairs.prototype.Empty = function () {
	this.feeTier = '0'
	this.currentAPR = ''
	this.capacity = ''
	this.tvl = 0.00
	this.maxTVL = 0.00
	this.ofCapUsed = 0.00
	this.vaultRange = ''
	this.vaultLending = 0
	this.tvlTotal0 = 0.00
	this.tvlTotal1 = 0.00
	this.tvlTotal0USD = 0
	this.tvlTotal1USD = 0
	this.lendingRatio = 0
	this.uniswapRatio = 0
	this.uniToken0Rate = 0
	this.uniToken1Rate = 0
	this.lendingToken0Rate = 0
	this.lendingToken1Rate = 0
	this.myShare = 0
	this.totalShares = 0
	this.myValueToken0Locked = 0.00
	this.myValueToken1Locked = 0.00
	this.myValueToken0USDLocked = 0
	this.myValueToken1USDLocked = 0
}
export default Pairs

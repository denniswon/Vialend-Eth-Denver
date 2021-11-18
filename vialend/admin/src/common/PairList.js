import Web3 from 'web3'
import contractABI from '../ABI/ViaLendFeeMakerABI.json'
import ViaLendTokenABI from '../ABI/tokenABI.json'
import Token from '../model/Token'
import Pairs from '../model/Pairs'
import ArrayList from '../common/ArrayList'

var web3
if (typeof web3 !== 'undefined') {
	web3 = new Web3(web3.currentProvider)
	console.log('web3 provider:web3.currentProvider')
} else {
	// set the provider you want from Web3.providers
	web3 = new Web3(new Web3.providers.HttpProvider('https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6'))
	console.log('web3 provider:goerli')
}

export function createPairs () {
	var pairsObj = {
		name: 'testname',
		pairsList: new ArrayList(),
		contractInstance: function (abi, address) {
			return new web3.eth.Contract(abi, address)
		},
		getPairList: async function () {
			this.pairsList.clear()
			var pair1 = new Pairs()
			pair1.index = 0
			pair1.id = 1
			pair1.vaultAddress = '0x4aaE0bc3052aD3AB125Ae654f0f2C55Dbd9D6e17'
			this.vaultAddress = pair1.vaultAddress
			pair1 = await this.getPairBasicInfo(pair1)
			this.pairsList.add(pair1)
			var pair2 = new Pairs()
			pair2.index = 1
			pair2.id = 2
			pair2.vaultAddress = '0x35938d9b221238BBcE1F9b5196FFeE0f87E22D26'
			pair2 = await this.getPairBasicInfo(pair2)
			this.pairsList.add(pair2)
			return this.pairsList
		},
		getPairBasicInfo: async function (pair) {
			var token0 = new Token()
			var token1 = new Token()
			var keeperContract = this.contractInstance(contractABI, pair.vaultAddress)
			// if (this.selectedPairIndex === pair.index) this.keeperContract = keeperContract
			pair.poolAddress = await keeperContract.methods.poolAddress().call()
			token0.tokenAddress = await keeperContract.methods.token0().call()
			token1.tokenAddress = await keeperContract.methods.token1().call()
			token0.token0LendingAddress = await keeperContract.methods.CToken0().call()
			token1.token1LendingAddress = await keeperContract.methods.CToken1().call()
			console.log('token0.tokenAddress=', token0.tokenAddress)
			console.log('token1.tokenAddress=', token1.tokenAddress)
			// load contract instance
			var token0Contract = this.contractInstance(ViaLendTokenABI, token0.tokenAddress)
			var token1Contract = this.contractInstance(ViaLendTokenABI, token1.tokenAddress)
			token0.iconLink = 'images/weth.png'
			token0.name = await token0Contract.methods.name().call()
			token0.symbol = await token0Contract.methods.symbol().call()
			token0.decimals = await token0Contract.methods.decimals().call()
			token1.iconLink = 'images/usdc.png'
			token1.name = await token1Contract.methods.name().call()
			token1.symbol = await token1Contract.methods.symbol().call()
			token1.decimals = await token1Contract.methods.decimals().call()
			pair.cLow = await keeperContract.methods.cLow().call()
			pair.cHigh = await keeperContract.methods.cHigh().call()
			console.log('clow===========', pair.cLow)
			pair.token0 = token0
			pair.token1 = token1
			return pair
		}
	}
	return pairsObj
}

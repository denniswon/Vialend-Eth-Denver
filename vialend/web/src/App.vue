<template>
  <div class="super_container"
       v-loading="pairsInfoLoading"
       element-loading-text="Loading"
       element-loading-spinner="el-icon-loading"
       element-loading-background="rgba(0, 0, 0, 0.8)">
    <router-view />
    <footer class="footer">
      <div class="container">
        <div class="row">
          <div class="col">
            <div class="footer_content d-flex flex-md-row flex-column align-items-center align-items-start justify-content-start">
              <div class="copyright">
              </div>
            </div>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script>
import Web3 from 'web3'
import contractABI from './ABI/ViaLendFeeMakerABI.json'
import uniswapV3PoolABI from './ABI/UniswapV3PoolABI.json'
import ViaLendTokenABI from './ABI/tokenABI.json'
import VaultBridgeABI from './ABI/VaultBridge.json'
import Token from './model/Token'
import Pairs from './model/Pairs'
import ArrayList from './common/ArrayList'
import axios from 'axios'

if (typeof web3 !== 'undefined') {
  web3 = new Web3(web3.currentProvider)
  console.log('web3 provider:web3.currentProvider')
} else {
  // set the provider you want from Web3.providers
  web3 = new Web3(new Web3.providers.HttpProvider('https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6'))
  console.log('web3 provider:goerli')
}

export default {
  name: 'App',
  data () {
    return {
      pairsList: new ArrayList(),
      pairsInfoLoading: false,
      pairsLoadComplete: false,
      bridgeAddress: '0x033F3C5eAd18496BA462783fe9396CFE751a2342'
    }
  },
  created: function () {
    // load the basic information of the token pair
    // this.loadPairsInfo()
    this.getTokensRateOfUSD()
  },
  methods: {
    getName () {
      console.log('name=' + this.$store.state.name)
    },
    contractInstance (abi, address) {
      return new web3.eth.Contract(abi, address)
    },
    async loadPairsInfo () {
      this.pairsInfoLoading = true
      this.pairsLoadComplete = false
      var bridgeContract = await this.contractInstance(VaultBridgeABI, this.bridgeAddress)
      var vault0AddressInContract = await bridgeContract.methods.getAddress(0).call()
      var vault1AddressInContract = await bridgeContract.methods.getAddress(1).call()
      // console.log('vaultAddressInContract=', vaultAddressInContract)
      var pair1 = new Pairs()
      pair1.index = 0
      pair1.id = 1
      pair1.vaultAddress = vault0AddressInContract
      // pair1.vaultAddress = '0x31C048503Bf4e15720025fb27D774DDc1829D925'
      // pair1.token0LendingAddress = '0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF'
      // pair1.token1LendingAddress = '0xCEC4a43eBB02f9B80916F1c718338169d6d5C1F0'
      pair1 = await this.loadTokensInfo(pair1)
      // pair1 = await this.getPairPublicInfo(pair1)
      this.pairsList.add(pair1)
      console.log('Pair1 loading completed')
      var pair2 = new Pairs()
      pair2.index = 1
      pair1.id = 2
      pair2.vaultAddress = vault1AddressInContract
      // pair2.vaultAddress = '0xf231F818a111FE5d2EFf006451689eCBbf5ef159'
      // pair2.token0LendingAddress = '0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF'
      // pair2.token1LendingAddress = '0xCEC4a43eBB02f9B80916F1c718338169d6d5C1F0'
      pair2 = await this.loadTokensInfo(pair2)
      pair2.token1.iconLink = 'images/dai.png'
      // pair2 = await this.getPairPublicInfo(pair2)
      this.pairsList.add(pair2)
      console.log('Pair2 loading completed')
      // console.log('pair1 token0 name is', pair1.token0.iconLink)
      // console.log('pair1 token1 name is', pair1.token1.iconLink)
      // console.log('pairsList.length=', this.pairsList.size())
      this.pairsLoadComplete = true
      this.pairsInfoLoading = false
    },
    async loadTokensInfo (pair) {
      var token0 = new Token()
      var token1 = new Token()
      var keeperContract = await this.contractInstance(contractABI, pair.vaultAddress)
      pair.poolAddress = await keeperContract.methods.poolAddress().call()
      token0.tokenAddress = await keeperContract.methods.token0().call()
      token1.tokenAddress = await keeperContract.methods.token1().call()
      token0.token0LendingAddress = await keeperContract.methods.CToken0().call()
      token1.token1LendingAddress = await keeperContract.methods.CToken1().call()
      var poolContract = await this.contractInstance(uniswapV3PoolABI, pair.poolAddress)
      var token0Contract = await this.contractInstance(ViaLendTokenABI, token0.tokenAddress)
      var token1Contract = await this.contractInstance(ViaLendTokenABI, token1.tokenAddress)
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
      pair.feeTier = await poolContract.methods.fee().call()
      pair.token0 = token0
      pair.token1 = token1
      return pair
    },
    async getTokensRateOfUSD () {
      // var response
      // console.log('this.$store.state.tokenExchangeTable=', this.$store.state.tokenExchangeTable)
      // for (var i = 0; i < this.$store.state.tokenExchangeTable.length; i++) {
      //   // console.log('token' + i + '=', this.$store.state.tokenExchangeTable[i].symbol)
      //   response = await axios.get('https://api.coinlore.net/api/ticker/?id=' + this.$store.state.tokenExchangeTable[i].id)
      //   if (response !== undefined && response !== null) {
      //     console.log('response=', response.data.length)
      //     if (response.data.length > 0) {
      //       this.$store.state.tokenExchangeTable[i].price_usd = response.data[0].price_usd
      //     }
      //   }
      // }
      // var response = await axios.get('https://ipfs.io/ipfs/QmXKQpiZ1eNGiAky141p8Q7Rfaunt1qy9M9Zi3MP7U3YZx/')
      // console.log('response=', response)
      // if (response.data !== undefined && response.data !== null) {
      //   for (var i = 0; i < response.data.length; i++) {
      //     console.log('vault address[', i, ']=', response.data[i])
      //   }
      // }
      this.$store.state.tokenExchangeRateLoaded = true
    }
  }
}
</script>

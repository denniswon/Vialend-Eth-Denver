<template>
  <div class="super_container">
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
import contractABI from './ABI/contractABI.json'
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
      vaultAddress: '0x58b008535DC1d06dbAe22201D3F10b79c80f9dd4',
      keeperContract: null
    }
  },
  created: function () {
    this.keeperContract = new web3.eth.Contract(
      contractABI,
      this.vaultAddress
    )
    this.getTokensRateOfUSD()
  },
  methods: {
    getName () {
      console.log('name=' + this.$store.state.name)
    },
    getTokensRateOfUSD () {
      axios.get('https://api.coinlore.net/api/tickers/').then((response) => {
        // console.log('data=', JSON.stringify(response.data.data))
        var allTokensList = response.data.data
        if (allTokensList !== null && allTokensList !== undefined) {
          this.$store.state.allTokensList = response.data.data
          console.log('allTokensList count:', allTokensList.length)
          for (var i = 0; i < allTokensList.length; i++) {
            if (allTokensList[i].symbol === 'ETH') {
              this.$store.state.token0RateOfUSD = allTokensList[i].price_usd
              console.log('Get token0RateOfUSD is ', this.$store.state.token0RateOfUSD)
            } else if (allTokensList[i].symbol === 'USDC') {
              this.$store.state.token1RateOfUSD = allTokensList[i].price_usd
              console.log('Get token1RateOfUSD is ', this.$store.state.token1RateOfUSD)
            }
            if (this.$store.state.token0RateOfUSD > 0 && this.$store.state.token1RateOfUSD > 0) {
              this.$store.state.tokenExchangeRateLoaded = true
              break
            }
          }
        }
      })
    }
  }
}
</script>

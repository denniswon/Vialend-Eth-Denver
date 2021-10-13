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
      vaultAddress: '0xAda1AC9D1dBFFF1270f0009d8f20bd0332F48113',
      poolAddress: '',
      keeperContract: null,
      allTokensList: null
    }
  },
  created: function () {
    this.keeperContract = new web3.eth.Contract(
      contractABI,
      this.vaultAddress
    )
    this.getPoolAddress()
    axios.get('https://api.coinlore.net/api/tickers/').then((response) => {
      // console.log('data=', JSON.stringify(response.data.data))
      // var obj = JSON.parse(response.data.data)
      this.$store.state.allTokensList = response.data.data
      if (this.$store.state.allTokensList !== null && this.$store.state.allTokensList !== undefined) { console.log('the tokens count:', this.$store.state.allTokensList.length) }
    })
  },
  methods: {
    async getPoolAddress () {
      this.poolAddress = await this.keeperContract.methods.poolAddress().call()
    },
    getName () {
      console.log('name=' + this.$store.state.name)
    }
  }
}
</script>

<style>
/* @import "./assets/styles/bootstrap-4.1.2/bootstrap.min.css";
@import "./assets/plugins/font-awesome-4.7.0/css/font-awesome.min.css";
@import "./assets/plugins/OwlCarousel2-2.2.1/owl.carousel.css";
@import "./assets/plugins/OwlCarousel2-2.2.1/owl.theme.default.css";
@import "./assets/plugins/OwlCarousel2-2.2.1/animate.css";
@import "./assets/plugins/colorbox/colorbox.css";
@import "./assets/styles/main_styles.css";
@import "./assets/styles/responsive.css"; */
</style>

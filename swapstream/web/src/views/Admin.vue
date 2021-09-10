<template>
  <div>
    <Header />
    <div class="home">
      <div class="parallax_background parallax-window"
           data-parallax="scroll"
           data-speed="0.8"></div>
      <div class="home_container">
        <div class="container">
          <div class="row align-items-center my-financial">
            <div class="col-md-3 text-center">
              <ul>
                <li class="title">My Liquidity</li>
                <li class="value">{{myLiquidity}}</li>
              </ul>
            </div>
            <div class="col-md-6 text-center apy-container">
              <div class="apy-style">
                <ul>
                  <li class="apy-title">Net APY</li>
                  <li class="apy-content">...</li>
                </ul>
              </div>
            </div>
            <div class="col-md-3 text-center">
              <ul>
                <li class="title">My Revenue</li>
                <li class="value">$0.00000000000000</li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import Header from '@/components/Header.vue'

const _this = this

export default {
  components: { Header },
  data () {
    return {
      vaultAddress: '0xaa16E934A327D500fdE1493302CeB394Ff6Ff0b2',
      currentAccount: null,
      keeperContract: null,
      isConnected: false,
      walletButtonClass: 'walletButton',
      connectClass: 'wallet_connected',
      disConnectClass: 'wallet_disconnected',
      StatusButtonText: 'Connect Wallet',
      treatyDialogVisible: false,
      myLiquidity: 0
    }
  },
  created: function () {
    // this.showTreaty()
    console.log('load create function')
    this.keeperContract = new web3.eth.Contract(
      contractABI,
      this.vaultAddress
    )
  },
  mounted () {
    window.connectWallet = this.connectWallet
  },
  methods: {
    getMyLiquidity () {
      if (this.keeperContract != null) {
        var myLiq = this.keeperContract.methods
          .balanceOf(this.currentAccount)
          .call()
          .then(val => {
            this.myLiquidity = val
            console.log('BalanceOf=' + val)
          })
        // console.log('BalanceOf=' + myLiq + '|' + JSON.stringify(myLiq))
      } else {
        console.log('keeperContract is null')
      }
    }
  }
}
</script>

<style scoped>
.my-financial ul li {
  line-height: 30px;
}
.my-financial .title {
  font-size: 16px;
}
.my-financial .value {
  font-size: 20px;
  color: white;
}
.my-financial .apy-title {
  margin-top: 50px;
  font-size: 20px;
  color: white;
}
.my-financial .apy-content {
  font-size: 20px;
  color: white;
}
.apy-container {
  position: relative;
}
.apy-style {
  position: absolute;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  margin: auto;
  width: 200px;
  height: 200px;
  border: 5px solid #9900ff;
  border-radius: 50%;
}
.dialog-content {
  line-height: 25px;
}
.dialog-checkbox {
  line-height: 50px;
}
.dialog-footer {
  line-height: 50px;
}
.wallet_connected {
  width: 76px !important;
  padding-left: 10px;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 1;
  -webkit-box-orient: vertical;
  word-break: break-all;
}
.wallet_disconnected {
  width: 151px;
}
.walletButton {
  width: 151px;
  height: 37px;
  background: #2e3f61;
  text-align: center;
  border-radius: 6px;
}
.walletButton:hover {
  background: #637496;
}
.walletButton a {
  display: block;
  font-size: 16px;
  font-weight: 500;
  color: #ffffff;
  line-height: 37px;
}
</style>

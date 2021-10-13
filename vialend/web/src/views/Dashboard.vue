<template>
  <div>
    <Header ref="headerComp" />
    <!-- Menu -->
    <div class="menu">
      <div class="background_image"></div>
      <div class="
          menu_content
          d-flex
          flex-column
          align-items-center
          justify-content-center
        ">
        <ul class="menu_nav_list text-center">
          <li><a href="index.html">Docs</a></li>
          <li><a href="about.html">Vote</a></li>
        </ul>
        <div class="menu_review"><a href="#">Launch App</a></div>
      </div>
    </div>
    <div class="home">
      <div class="parallax_background parallax-window"
           data-parallax="scroll"
           data-speed="0.8"></div>
      <div class="home_container">
        <div class="container">
          <div class="row align-items-center my-financial">
            <div class="col-md-3 text-center"
                 v-loading="myValueLockLoading"
                 element-loading-spinner="el-icon-loading"
                 element-loading-background="rgba(0, 0, 0, 0.8)">
              <ul>
                <li class="title">My Value Locked</li>
                <li class="value">{{token0Symbol + ':' + myValueToken0Locked.toFixed(2)}}</li>
                <li class="value">{{token1Symbol + ':' + myValueToken1Locked.toFixed(2)}}</li>
              </ul>
            </div>
            <div class="col-md-6 text-center apy-container">
              <div class="apy-style"
                   v-loading="netAPYLoading"
                   element-loading-spinner="el-icon-loading"
                   element-loading-background="rgba(0, 0, 0, 0.8)">
                <ul>
                  <li class="apy-title">Net APY</li>
                  <li class="apy-content">...</li>
                </ul>
              </div>
            </div>
            <div class="col-md-3 text-center"
                 v-loading="myEarnedValueLoading"
                 element-loading-spinner="el-icon-loading"
                 element-loading-background="rgba(0, 0, 0, 0.8)">
              <ul>
                <li class="title">My Earned Value</li>
                <li class="value">-</li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
    <MyPositions />

    <el-dialog title="SMART VAULTS"
               :visible.sync="treatyDialogVisible"
               :append-to-body="true"
               width="30%"
               center>
      <span class="dialog-content">SwapStream's Smart Vaults are a new experimental DeFi application that is open-source software maintained by a distributed autonomous organization.
        The software is recorded as a set of smart contracts on Ethereum with no promise to return profits or guarantee the security of funds.
        The application is not regulated and may not be subjected to the same laws and regulations that are applicable to other types of businesses,organizations,firms,individuals,partnerships,or solo practitioners.
        It is important to recognize that you can lose up to 100% of your funds,with no recourse to compensation.
        Please ensure you do your own research and understand the risks before proceeding to use these vaults.
        For more information you can read and review the contracts at the following location</span>
      <span class="dialog-checkbox">
        <el-checkbox>I understand and agree to these conditions</el-checkbox>
      </span>
      <span slot="footer"
            class="dialog-footer">
        <el-button type="primary"
                   @click="goSign">SIGN</el-button>
        <router-view />
      </span>
    </el-dialog>
  </div>
</template>

<script>
import Header from '@/components/Header.vue'
import MyPositions from '@/components/MyPositions.vue'

const _this = this

export default {
  components: { Header, MyPositions },
  data () {
    return {
      myValueLockLoading: false,
      myEarnedValueLoading: false,
      netAPYLoading: false,
      vaultAddress: this.$parent.vaultAddress,
      keeperContract: this.$parent.keeperContract,
      currentAccount: null,
      isConnected: false,
      treatyDialogVisible: false,
      myLiquidity: 0,
      myShare: 0,
      totalSupply: 0,
      myValueLocked: 0,
      myValueToken0Locked: 0,
      myValueToken1Locked: 0,
      myEarnedValue: 0,
      tvl: 0,
      token0Symbol: '',
      token1Symbol: ''
    }
  },
  created: function () {
    console.log('store name=', this.$store.state.name)
  },
  mounted () {
    window.connectWallet = this.connectWallet
  },
  watch: {
    '$store.state.isConnected': function () {
      this.isConnected = this.$store.state.isConnected
      if (!this.isConnected) {
        this.myValueToken0Locked = 0
        this.myValueToken1Locked = 0
      }
    }
  },
  methods: {
    showTreaty () {
      this.treatyDialogVisible = true
      // console.log(this.$store.state.name)
    },
    goSign () {

    },
    setWalletStatus () {
      this.$refs.headerComp.setWalletStatus()
    },
    async getShares () {
      this.myShare = await this.keeperContract.methods.balanceOf(ethereum.selectedAddress).call()
      console.log('user address=', ethereum.selectedAddress, ';shares=', this.shares)
      this.keeperContract.methods
        .totalSupply()
        .call()
        .then(val => {
          console.log('totalSupply= ' + val)
          this.totalSupply = val
        })
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

.adminButton {
  width: 100px;
  height: 37px;
  background: #2a4988;
  text-align: center;
  border-radius: 6px;
}
.adminButton:hover {
  background: #637496;
}
.adminButton a {
  display: block;
  font-size: 16px;
  font-weight: 500;
  color: #ffffff;
  line-height: 37px;
}
</style>

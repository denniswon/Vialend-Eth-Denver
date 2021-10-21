<template>
  <!-- Header -->
  <header class="header">
    <div class="container">
      <div class="row">
        <div class="col">
          <div class="
                header_content
                d-flex
                flex-row
                align-items-center
                justify-content-start
              ">
            <div class="logo">
              <a href="#">
                <div style="vertical-align: middle;">
                  <img src="images/logo.png"
                       width="30"
                       height="33" />&nbsp;&nbsp;Vialend
                </div>
                <div></div>
              </a>
            </div>
            <nav class="main_nav">
              <ul class="d-flex flex-row align-items-center justify-content-start">
                <li>
                  <router-link to='/dashboard'>
                    Dashboard
                  </router-link>
                  <!-- <a href="/dashboard">Dashboard</a> -->
                </li>
                <li><a href="#">Vote</a></li>
              </ul>
            </nav>
            <div class="
      header_extra
      d-flex
      flex-row
      align-items-center
      justify-content-start
      ml-auto
    ">

              <div class="adminButton">
                <router-link to='/admin'>
                  Admin
                </router-link>
              </div>&nbsp;&nbsp;
              <div :class="[walletButtonClass,isConnected ? connectClass:disConnectClass]">
                <a href="#"
                   @click="setWalletStatus">{{ StatusButtonText }}</a>
              </div>
            </div>
            <div class="hamburger ml-auto">
              <i class="fa fa-bars"
                 aria-hidden="true"></i>
            </div>
          </div>
        </div>
      </div>
    </div>
  </header>
</template>

<script>
export default {
  data () {
    return {
      isConnected: this.$store.state.isConnected,
      walletButtonClass: 'walletButton',
      connectClass: 'wallet_connected',
      disConnectClass: 'wallet_disconnected',
      StatusButtonText: 'Connect Wallet',
      centerDialogVisible: false
    }
  },
  created: function () {
    this.getChainId()
    if (this.$store.state.StatusButtonText !== '') {
      this.StatusButtonText = this.$store.state.StatusButtonText
    }
  },
  mounted () {
    window.connectWallet = this.connectWallet
    this.fn()
  },
  watch: {
    isConnected (newStatus, oldStatus) {
      console.log('newStatus=', newStatus, ';oldStatus=', oldStatus)
      // this.$parent.isConnected = newStatus
      this.$store.state.isConnected = newStatus
      console.log('this.$store.state.isConnected header watch=', this.$store.state.isConnected)
    }
  },
  methods: {
    fn () {
      window.ethereum.autoRefreshOnNetworkChange = false
      window.ethereum.on('accountsChanged', () => {
        console.log('accountsChanged')
        this.connectWallet()
      })
      window.ethereum.on('networkChanged', () => {
        console.log('networkChanged')
        this.getChainId()
        this.networkChanged()
      })
    },
    getChainId () {
      var _this = this
      web3.eth.getChainId().then(function (val) {
        _this.$store.state.chainId = val
        console.log('ChainId:', val)
      })
    },
    checkChain () {
      if (this.$store.state.chainId === 5) { return true } else { return false }
    },
    networkChanged () {
      this.StatusButtonText = 'Connect Wallet'
      this.$store.state.StatusButtonText = 'Connect Wallet'
      this.isConnected = false
    },
    lanuchApp () {
      this.centerDialogVisible = false
      this.$router.push({ path: '/dashboard' })
    },
    // connect wallet or disconnect wallet
    setWalletStatus () {
      if (this.isConnected) {
        console.log('call disconnect')
        this.isConnected = false
        // ethereum.on('disconnect', error => { console.log(error) })
        this.StatusButtonText = 'Connect Wallet'
        this.$store.state.StatusButtonText = 'Connect Wallet'
      } else {
        if (ethereum.isConnected() && this.currentAccount != null) {
          this.isConnected = true
          console.log('this.currentAccount=' + this.currentAccount)
          this.StatusButtonText = this.currentAccount
          this.$store.state.StatusButtonText = this.currentAccount
          this.$store.state.currentAccount = this.currentAccount
        } else {
          this.connectWallet()
        }
      }
      if (!this.checkChain()) {
        this.isConnected = false
        this.StatusButtonText = 'Wrong network'
        this.$store.state.StatusButtonText = 'Wrong network'
        this.$message({
          message: 'Please select Goerli Test Network.',
          type: 'warning'
        })
      }
    },
    connectWallet () {
      if (!this.checkChain()) {
        this.StatusButtonText = 'Wrong network'
        this.$store.state.StatusButtonText = 'Wrong network'
        return
      }
      console.log('call connectWallet')
      ethereum
        .request({ method: 'eth_requestAccounts' })
        .then(this.handleAccountsChanged)
        .catch((err) => {
          // Some unexpected error.
          // For backwards compatibility reasons, if no accounts are available,
          // eth_accounts will return an empty array.
          console.error(err)
        })
    },
    handleAccountsChanged (accounts) {
      if (accounts.length === 0) {
        // MetaMask is locked or the user has not connected any accounts
        console.log('Please connect to MetaMask.')
      } else if (accounts[0] !== this.currentAccount) {
        this.currentAccount = accounts[0]
        // Do any other work!
        this.isConnected = true
        this.StatusButtonText = this.currentAccount
        this.$store.state.StatusButtonText = this.currentAccount
        this.$store.state.currentAccount = this.currentAccount
        console.log('account status:' + ethereum.isConnected())
        // this.$refs.supplyliq.checkConnectionStatus()
      }
    }
  }
}
</script>

<style>
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
  font-size: 16px;
  font-weight: 500;
  color: #ffffff;
  line-height: 37px;
}
</style>

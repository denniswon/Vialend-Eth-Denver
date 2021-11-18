<template>
  <!-- Header -->
  <header class="header">
    <div :class="containerClass">
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

              <div class="adminButton"
                   v-show="isAdmin">
                <router-link to='/admin'>
                  Admin
                </router-link>
              </div>&nbsp;&nbsp;
              <div v-if="$store.state.isConnected"
                   :class="[walletButtonClass,connectClass]">
                <a href="#"
                   @click="disconnectWallet">{{ $store.state.currentAccount }}</a>
              </div>
              <div v-else-if="validNetwork"
                   :class="[walletButtonClass,disConnectClass]">
                <a href="#"
                   @click="connectWallet">Wrong network</a>
              </div>
              <div v-else
                   :class="[walletButtonClass,disConnectClass]">
                <a href="#"
                   @click="connectWallet">Connect Wallet</a>
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
import VaultAdminABI from '../ABI/VaultAdmin.json'

export default {
  data () {
    return {
      isAdmin: false,
      validNetwork: false,
      isConnected: this.$store.state.isConnected,
      walletButtonClass: 'walletButton',
      connectClass: 'wallet_connected',
      disConnectClass: 'wallet_disconnected',
      containerClass: 'container',
      centerDialogVisible: false
    }
  },
  created: async function () {
    if (await this.checkChain()) {
      if (ethereum.selectedAddress !== null && ethereum.selectedAddress !== undefined) {
        console.log('currentAccount is connected.')
        this.$store.state.currentAccount = ethereum.selectedAddress
        this.$store.state.isConnected = true
        this.isConnected = true
        this.checkValidAdmin()
      } else {
        this.$store.state.currentAccount = ''
        this.$store.state.isConnected = false
        this.isConnected = false
      }
    } else {
      this.$store.state.currentAccount = ''
      this.$store.state.isConnected = false
      this.isConnected = false
    }
  },
  mounted () {
    window.connectWallet = this.connectWallet
    this.fn()
  },
  watch: {
    isConnected (newStatus, oldStatus) {
      console.log('newStatus=', newStatus, ';oldStatus=', oldStatus)
      this.$store.state.isConnected = newStatus
      console.log('this.$store.state.isConnected header watch=', this.$store.state.isConnected)
    }
  },
  methods: {
    contractInstance (abi, address) {
      return new web3.eth.Contract(abi, address)
    },
    async checkValidAdmin () {
      if (ethereum.selectedAddress !== null && ethereum.selectedAddress !== undefined) {
        var vaultAdminContract = await this.contractInstance(VaultAdminABI, '0xb6F0049e37D32dED0ED2FAEeE7b69930FA49A879')
        this.isAdmin = await vaultAdminContract.methods.authAdmin(ethereum.selectedAddress).call()
        console.log('isAdmin=', this.isAdmin)
        this.$store.state.isAdmin = this.isAdmin
        localStorage.setItem('isAdmin', this.isAdmin)
        // var vaultAdminList = await vaultAdminContract.methods.getAdmin().call()
        // console.log('admin list=', vaultAdminList)
      }
    },
    fn () {
      window.ethereum.autoRefreshOnNetworkChange = false
      window.ethereum.on('accountsChanged', async () => {
        console.log('accountsChanged')
        if (await this.checkChain()) {
          this.changeAccount()
          this.checkValidAdmin()
        }
      })
      window.ethereum.on('networkChanged', async () => {
        console.log('networkChanged')
        if (await this.checkChain()) {
          this.changeAccount()
          this.checkValidAdmin()
        }
      })
      window.ethereum.on('disconnect', () => {
        console.log('metamask disconnect')
        this.$store.state.currentAccount = ''
        this.$store.state.isConnected = false
        this.isConnected = false
      })
    },
    async checkChain () {
      this.$store.state.chainId = await web3.eth.getChainId()
      console.log('check chain id=', this.$store.state.chainId)
      if (this.$store.state.availableChainId.includes(this.$store.state.chainId)) {
        this.$store.state.isConnected = true
        this.$store.state.validNetwork = true
        return true
      } else {
        this.isAdmin = false
        this.$store.state.isConnected = false
        this.$store.state.validNetwork = false
        return false
      }
    },
    connectWallet () {
      if (ethereum.selectedAddress !== null && ethereum.selectedAddress !== undefined) {
        if (!this.$store.state.validNetwork) {
          this.$message({
            message: 'Please select Goerli Test Network.',
            type: 'warning'
          })
        } else {
          this.$store.state.currentAccount = ethereum.selectedAddress
          this.$store.state.isConnected = true
          this.$store.state.validNetwork = true
        }
      } else {
        console.log('call connectWallet')
        ethereum
          .request({ method: 'eth_requestAccounts' })
          .then(this.requestAccountsCallBack)
          .catch((err) => {
            // Some unexpected error.
            // For backwards compatibility reasons, if no accounts are available,
            // eth_accounts will return an empty array.
            console.error(err)
          })
      }
    },
    requestAccountsCallBack (accounts) {
      if (accounts.length === 0) {
        // MetaMask is locked or the user has not connected any accounts
        console.log('Please connect to MetaMask.')
      } else if (accounts[0] !== this.$store.state.currentAccount) {
        this.$store.state.currentAccount = accounts[0]
        this.$store.state.isConnected = true
        // console.log('account status:' + ethereum.isConnected())
      } else {
        this.$store.state.isConnected = true
        console.log('accounts[0]:' + accounts[0])
      }
    },
    changeAccount () {
      if (ethereum.selectedAddress !== null && ethereum.selectedAddress !== undefined) {
        this.$store.state.isConnected = true
        this.$store.state.currentAccount = ethereum.selectedAddress
      } else {
        this.$store.state.currentAccount = ''
        this.$store.state.isConnected = false
        this.isConnected = false
        this.isAdmin = false
      }
    },
    disconnectWallet () {
      this.$store.state.isConnected = false
      this.$store.state.currentAccount = ''
      console.log('Disconnect wallet!')
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

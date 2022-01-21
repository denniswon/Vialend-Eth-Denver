<template>
  <div class="navbar-container">
    <hamburger id="hamburger-container"
               :is-active="sidebar.opened"
               class="hamburger-container"
               @toggle-click="toggleSideBar" />
    <div class="btn-nav">
      <router-link to="/">
        <el-button :type="$route.meta.title === 'products' ? 'primary' : ''"
                   size="small">Products</el-button>
      </router-link>&nbsp;&nbsp;
      <router-link to="/portfolio">
        <el-button :type="$route.meta.title === 'portfolio' ? 'primary' : ''"
                   size="small">Portfolio</el-button>
      </router-link>
    </div>
    <div class="right-menu">
      <div class="newmsgicon">
        <el-button type="text">Unerified</el-button>
      </div>
      <div class="newmsgicon">
        <svg-icon name="newmsg"
                  width="26px"
                  height="26px"></svg-icon>
      </div>
      <!-- <button @click="logout">Logout</button> -->
      <div class="btnnetwork">
        <div v-if="$store.state.isConnected === false" :class="[walletButtonClass, disConnectClass]">
          <a href="#" @click="connectWallet">Connect Wallet</a>
        </div>
        <div v-else-if="$store.state.isConnected && $store.state.validNetwork"
             :class="[walletButtonClass, connectClass]">
          <!-- <a href="#" @click="disconnectWallet">{{ $store.state.currentAccount }}</a> -->
          <el-popover placement="bottom"
                      width="180"
                      @show="showWalletMenu"
                      @hide="hideWalletMenu"
                      trigger="click">
            <div class="connect-menu">
              <ul>
                <li>CHANGE WALLET</li>
                <li>COPY ADDRESS</li>
                <li><a :href="'https://goerli.etherscan.io/address/' + $store.state.currentAccount" target="_blank">OPEN IN ETHERSCAN</a></li>
                <li @click="disconnectWallet">DISCONNECT</li>
              </ul>
            </div>
            <el-link slot="reference" :underline="false">{{ formatWalletAddress($store.state.currentAccount)}}<i :class="['el-select__caret','el-icon-arrow-up',walletMenuVisible ? 'is-reverse' : '']"></i></el-link>
          </el-popover>
        </div>
        <div v-else
             :class="[walletButtonClass, disConnectClass]">
          <a href="#"
             @click="connectWallet">Wrong network</a>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator'
import { getWeb3Instance, contractInstance } from '../../../common/Web3'
import { AppModule } from '@/store/modules/app'
import { UserModule } from '@/store/modules/user'
import Hamburger from '@/components/Hamburger/index.vue'
// import { AbiItem } from 'web3-utils'
const VaultBridgeABI = require('../../../abi/VaultBridge.json')

@Component({
  name: 'Navbar',
  components: {
    Hamburger
  }
})
export default class extends Vue {
  web3 = getWeb3Instance()
  isAdmin = '0'
  validNetwork = false
  isConnected = this.$store.state.isConnected
  walletButtonClass = 'walletButton'
  connectClass = 'wallet_connected'
  disConnectClass = 'wallet_disconnected'
  containerClass = 'container'
  centerDialogVisible = false
  walletMenuVisible = false
  doDisconnect = false

  get sidebar() {
    return AppModule.sidebar
  }

  get device() {
    return AppModule.device.toString()
  }

  get avatar() {
    return UserModule.avatar
  }

  private toggleSideBar() {
    AppModule.ToggleSideBar(false)
  }

  logout() {
    UserModule.LogOut()
    this.$router
      .push(`/login?redirect=${this.$route.fullPath}`)
      .catch((err) => {
        console.warn(err)
      })
  }

  @Watch('$store.state.isConnected')
  walletConnected(newStatus: boolean, oldStatus: boolean) {
    console.log(
      '$store.state.isConnected newStatus=',
      newStatus,
      ';oldStatus=',
      oldStatus
    )
    this.$store.state.isConnected = newStatus
    console.log(
      'this. header isConnected watch=',
      this.$store.state.isConnected
    )
  }

  showWalletMenu(): void {
    this.walletMenuVisible = true
  }

  hideWalletMenu(): void {
    this.walletMenuVisible = false
  }

  formatWalletAddress(addr: string) {
    if (addr !== null && addr !== undefined && addr !== '') {
      return addr
        .toUpperCase()
        .substr(0, 6)
        .concat('...', addr.toUpperCase().substr(-4, 4))
    } else {
      return ''
    }
  }

  mounted() {
    (window as any).connectWallet = this.connectWallet
    this.fn()
  }

  async checkValidAdmin() {
    console.log('(window as any).ethereum.selectedAddress =', (window as any).ethereum.selectedAddress)
    if ((window as any).ethereum.selectedAddress !== null && (window as any).ethereum.selectedAddress !== undefined) {
      const vaultAdminContract = await contractInstance(VaultBridgeABI, this.$store.state.bridgeAddress)
      this.isAdmin = await vaultAdminContract.methods.getPermit((window as any).ethereum.selectedAddress).call()
      if (this.isAdmin === '1') {
        this.$store.state.isAdmin = true
        console.log('isAdmin=true')
      } else {
        this.$store.state.isAdmin = false
        console.log('isAdmin=false')
      }
      UserModule.ChangeRoles(this.$store.state.isAdmin ? 'admin' : 'user').then(() => {
        this.$emit('role changed')
      })
      localStorage.setItem('isAdmin', this.isAdmin.toString())
      // var vaultAdminList = await vaultAdminContract.methods.getAdmin().call()
      // console.log('admin list=', vaultAdminList)
    }
  }

  fn() {
    (window as any).ethereum.autoRefreshOnNetworkChange = false;
    (window as any).ethereum.on('accountsChanged', async() => {
      console.log('accountsChanged')
      if (await this.checkChain()) {
        this.changeAccount()
        this.checkValidAdmin()
      }
    });
    (window as any).ethereum.on('networkChanged', async() => {
      console.log('networkChanged')
      if (await this.checkChain()) {
        this.changeAccount()
        this.checkValidAdmin()
      }
    })
    ;(window as any).ethereum.on('disconnect', () => {
      console.log('metamask disconnect')
      this.$store.state.currentAccount = ''
      this.$store.state.isConnected = false
      this.isConnected = false
    })
  }

  handleCommand(command: any) {
    this.$message('click on item ' + command)
  }

  async checkChain() {
    this.$store.state.chainId = await this.web3.eth.getChainId()
    console.log('check chain id=', this.$store.state.chainId)
    if (this.$store.state.availableChainId.includes(this.$store.state.chainId)) {
      // this.$store.state.isConnected = true
      this.$store.state.validNetwork = true
      return true
    } else {
      this.$store.state.isConnected = false
      this.$store.state.validNetwork = false
      return false
    }
  }

  async connectWallet() {
    this.$store.state.doDisconnect = false
    if ((window as any).ethereum.selectedAddress !== null && (window as any).ethereum.selectedAddress !== undefined) {
      if (!this.$store.state.validNetwork) {
        this.$message({
          message: 'Please select Goerli Test Network.',
          type: 'warning'
        })
      } else {
        if (await this.checkChain()) {
          this.changeAccount()
          this.checkValidAdmin()
        }
      }
    } else {
      console.log('call connectWallet')
      ;(window as any).ethereum
        .request({ method: 'eth_requestAccounts' })
        .then(this.requestAccountsCallBack)
        .catch((err: any) => {
          // Some unexpected error.
          // For backwards compatibility reasons, if no accounts are available,
          // eth_accounts will return an empty array.
          console.error(err)
        })
    }
  }

  requestAccountsCallBack(accounts: string) {
    if (accounts.length === 0) {
      // MetaMask is locked or the user has not connected any accounts
      console.log('Please connect to MetaMask.')
    } else if (accounts[0] !== this.$store.state.currentAccount) {
      this.$store.state.currentAccount = accounts[0]
      this.$store.state.isConnected = true
      console.log('account status:' + (window as any).ethereum.isConnected())
    } else {
      this.$store.state.isConnected = true
      console.log('accounts[0]:' + accounts[0])
    }
  }

  changeAccount() {
    console.log('doChangeAccount')
    if ((window as any).ethereum.selectedAddress !== null && (window as any).ethereum.selectedAddress !== undefined) {
      this.$store.state.isConnected = true
      this.$store.state.currentAccount = (window as any).ethereum.selectedAddress
    } else {
      this.$store.state.isConnected = false
      this.$store.state.currentAccount = ''
      this.isConnected = false
    }
  }

  disconnectWallet() {
    this.$store.state.isConnected = false
    this.$store.state.currentAccount = ''
    this.$store.state.isAdmin = false
    this.$store.state.doDisconnect = true
    UserModule.ChangeRoles('user').then(() => {
      this.$emit('role changed')
    })
    console.log('Disconnect wallet!')
  }

  async created() {
    if (await this.checkChain()) {
      if ((window as any).ethereum.selectedAddress !== null && (window as any).ethereum.selectedAddress !== undefined) {
        if (!this.$store.state.doDisconnect) {
          console.log('currentAccount is connected.')
          this.$store.state.currentAccount = (window as any).ethereum.selectedAddress
          this.$store.state.isConnected = true
          this.isConnected = true
          this.checkValidAdmin()
        }
      } else {
        console.log('ethereum.selectedAddress is null.')
        this.$store.state.currentAccount = ''
        this.$store.state.isConnected = false
        this.isConnected = false
      }
    } else {
      console.log('there is invalid network.')
      this.$store.state.currentAccount = ''
      this.$store.state.isConnected = false
      this.isConnected = false
    }
  }
}
</script>

<style>
.el-popover {
  background-color: #121218;
  border-radius: 12px;
}
.el-popper[x-placement^='bottom'] .popper__arrow::after {
  border-bottom-color: #121218;
}
</style>

<style lang="scss" scoped>
.navbar-container {
  height: 50px;
  overflow: hidden;
  position: relative;
  background: #fff;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);

  .hamburger-container {
    line-height: 46px;
    height: 100%;
    float: left;
    padding: 0 15px;
    cursor: pointer;
    transition: background 0.3s;
    -webkit-tap-highlight-color: transparent;

    &:hover {
      background: rgba(0, 0, 0, 0.025);
    }
  }

  .btn-nav {
    line-height: 46px;
    width: 900px;
    height: 100%;
    float: left;
    padding: 0 15px;
    margin: 0 auto;
  }

  .breadcrumb-container {
    float: left;
  }

  .errLog-container {
    display: inline-block;
    vertical-align: top;
  }

  .right-menu {
    display: inline-block;
    position: absolute;
    right: 0;
    height: 100%;
    line-height: 50px;

    &:focus {
      outline: none;
    }

    .right-menu-item {
      display: inline-block;
      padding: 0 8px;
      height: 100%;
      font-size: 18px;
      color: #5a5e66;
      vertical-align: text-bottom;

      &.hover-effect {
        cursor: pointer;
        transition: background 0.3s;

        &:hover {
          background: rgba(0, 0, 0, 0.025);
        }
      }
    }

    .avatar-container {
      margin-right: 30px;

      .avatar-wrapper {
        margin-top: 5px;
        position: relative;

        .user-avatar {
          cursor: pointer;
          width: 40px;
          height: 40px;
          border-radius: 10px;
        }

        .el-icon-caret-bottom {
          cursor: pointer;
          position: absolute;
          right: -20px;
          top: 25px;
          font-size: 12px;
        }
      }
    }
  }
}

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
  width: 136px !important;
  text-align: left !important;
  padding-left: 14px;
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
  margin: 6px;
}
// .walletButton:hover {
//   background: #637496;
// }
.walletButton a {
  display: block;
  font-size: 14px;
  font-weight: 500;
  color: #ffffff;
  line-height: 36px;
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
.newmsgicon {
  display: inline-block;
  vertical-align: middle;
  margin-top: 5px;
  margin-right: 30px;
  line-height: 20px;
}
.btnnetwork {
  display: inline-block;
  vertical-align: middle;
  margin-right: 10px;
}
.svg-icon {
  vertical-align: 0px !important;
}
.connect-menu {
}
.connect-menu ul {
  margin-block-start: 0.5em !important;
  margin-block-end: 0.5em !important;
  padding-inline-start: 10px !important;
}
.connect-menu li {
  color: #ffffff;
  padding: 0;
  margin: 0;
  list-style: none;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  font-size: 14px;
  line-height: 36px;
  cursor:pointer;
}
.connect-menu li:hover {
  color: #f56c6c;
}

.el-select__caret {
  transition: transform 0.3s;
  transform: rotate(180deg);
  cursor: pointer;
  margin-left: 10px;
}
.is-reverse {
  -webkit-transform: rotate(0deg);
  transform: rotate(0deg);
}
.el-link.el-link--default:hover {
  color: #ffffff !important;
}
</style>

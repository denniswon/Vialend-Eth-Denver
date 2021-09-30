<template>
  <div>
    <!-- Header -->
    <Header ref="headerComp" />

    <div>
      <div class="text item">
        <p><b>Pool Address:</b> {{resultPoolAddress}}</p>
        <p><b>Fee Tier:</b> {{feeTier}}</p>
        <p><b>Shares:</b> {{shares}}</p>
        <p><b>Total Shares:</b> {{totalShares}}</p>
        <p><b>TVL:</b> {{tvlTotal0}} + {{tvlTotal1}}</p>
        <p><b>Of cap used:</b> {{ofCapUsed}}</p>
        <p><b>Max TVL:</b> {{maxTVL}}</p>
        <p><b>Range:</b> {{vaultRange}}</p>
        <p><b>Lending:</b> {{vaultLending}}</p>
        <p><b>Current APR:</b> {{currentAPR}}</p>
      </div>
      <el-card class="box-card">
        <div class="text item">
          <b>Tokens Information:</b>
        </div>
        <table class="table">
          <thead>
            <th>Name</th>
            <th>Address</th>
            <th>Symbol</th>
            <th>Decimal</th>
          </thead>
          <tbody>
            <tr>
              <td>{{token0Name}}</td>
              <td>{{token0Address}}</td>
              <td>{{token0Symbol}}</td>
              <td>{{token0Decimal}}</td>
            </tr>
            <tr>
              <td>{{token1Name}}</td>
              <td>{{token1Address}}</td>
              <td>{{token1Symbol}}</td>
              <td>{{token1Decimal}}</td>
            </tr>
          </tbody>
        </table>
      </el-card>
    </div>
  </div>
</template>

<script>
import Web3 from 'web3'
import ViaLendFeeMakerABI from '../ABI/ViaLendFeeMakerABI.json'
import ViaLendTokenABI from '../ABI/tokenABI.json'
import ViaLendPoolABI from '../ABI/UniswapV3PoolABI.json'
import { TickMath } from '../utils/tickMath'
import { encodeSqrtRatioX96 } from '../utils/encodeSqrtRatioX96'
import Header from '@/components/Header.vue'

const _this = this
const POOL_SQRT_RATIO_START = encodeSqrtRatioX96(100e6, 100e18)
const POOL_TICK_CURRENT = TickMath.getTickAtSqrtRatio(POOL_SQRT_RATIO_START)

if (typeof web3 !== 'undefined') {
  web3 = new Web3(web3.currentProvider)
  console.log('web3 provider:web3.currentProvider')
} else {
  // set the provider you want from Web3.providers
  web3 = new Web3(new Web3.providers.HttpProvider('https://goerli.infura.io'))
  console.log('web3 provider:web3.HttpProvider')
}

window.ethereum.autoRefreshOnNetworkChange = false
window.ethereum.on('accountsChanged', () => {
  console.log('accountsChanged')
  connectWallet()
})
window.ethereum.on('networkChanged', () => {
  console.log('networkChanged')
  connectWallet()
})

export default {
  components: { Header },
  data () {
    return {
      vaultAddress: '0x31E84D42aB6DEf5Dac84b761b0E5004179e07778',
      currentAccount: null,
      keeperContract: null,
      token0Contract: null,
      token1Contract: null,
      poolContract: null,
      isConnected: false,
      walletButtonClass: 'walletButton',
      connectClass: 'wallet_connected',
      disConnectClass: 'wallet_disconnected',
      StatusButtonText: 'Connect Wallet',
      resultContent: '',
      resultPoolAddress: '',
      resultToken0: '',
      resultToken1: '',
      feeTier: '',
      token0Address: '',
      token0Name: '',
      token0Symbol: '',
      token0Decimal: 0,
      token1Address: '',
      token1Name: '',
      token1Symbol: '',
      token1Decimal: 0,
      shares: 0,
      totalShares: 0,
      tvl: 0,
      tvlTotal0: 0,
      tvlTotal1: 0,
      ofCapUsed: '',
      maxTVL: '',
      vaultRange: '',
      vaultLending: '',
      currentAPR: ''
    }
  },
  created: function () {
    console.log('load create function')
    this.keeperContract = new web3.eth.Contract(
      ViaLendFeeMakerABI,
      this.vaultAddress
    )
    this.token0Contract = new web3.eth.Contract(
      ViaLendTokenABI,
      '0x48FCb48bb7F70F399E35d9eC95fd2A614960Dcf8'
    )
    this.token1Contract = new web3.eth.Contract(
      ViaLendTokenABI,
      '0x6f38602e142D0Bd3BC162f5912535f543D3B73d7'
    )
    this.poolContract = new web3.eth.Contract(
      ViaLendPoolABI,
      '0xc4C92691f69fadDd684257E9f5A8d6f9D2c79a93'
    )
  },
  mounted () {
    window.connectWallet = this.connectWallet
    this.getTokenInfo()
    this.getData()
    this.getPoolInfo()
    this.getShares()
  },
  methods: {
    getData () {
      var _this = this
      this.keeperContract.methods
        .pool()
        .call()
        .then(val => {
          console.log('pool address=' + val)
          _this.resultPoolAddress = val
        })
      this.keeperContract.methods
        .token0()
        .call()
        .then(val => {
          console.log('Token0: ' + val)
          _this.token0Address = val
        })
      this.keeperContract.methods
        .token1()
        .call()
        .then(val => {
          console.log('Token1: ' + val)
          _this.token1Address = val
        })
      this.keeperContract.methods
        .decimals()
        .call()
        .then(val => {
          console.log('decimals: ' + val)
        })
      this.keeperContract.methods
        .protocolFee()
        .call()
        .then(val => {
          console.log('protocolFee: ' + val)
        })
      // Get TVL
      this.keeperContract.methods
        .getTVL()
        .call()
        .then(val => {
          console.log('getTVL: ' + JSON.stringify(val))
          console.log('decimal=', this.token0Decimal)
          _this.tvlTotal0 = val.total0
          _this.tvlTotal1 = val.total1
        })
      // Get Max TVL
      this.keeperContract.methods
        .maxTotalSupply()
        .call()
        .then(val => {
          console.log('maxTVL: ' + val)
          _this.maxTVL = val
          // _this.ofCapUsed =
        })
    },
    async getTokenInfo () {
      // Get Token0 Information
      this.token0Name = await this.token0Contract.methods.name().call()
      console.log('token0Name=', this.token0Name)
      this.token0Symbol = await this.token0Contract.methods.symbol().call()
      console.log('token0Symbol=', this.token0Symbol)
      this.token0Decimal = await this.token0Contract.methods.decimals().call()
      console.log('token0Decimal=', this.token0Decimal)
      // Get Token1 Information
      this.token1Name = await this.token1Contract.methods.name().call()
      console.log('token1Name=', this.token1Name)
      this.token1Symbol = await this.token1Contract.methods.symbol().call()
      console.log('token1Symbol=', this.token1Symbol)
      this.token1Decimal = await this.token1Contract.methods.decimals().call()
      console.log('token1Decimal=', this.token1Decimal)
    },
    async getPoolInfo () {
      this.feeTier = await this.poolContract.methods.fee().call()
      console.log('feeTier=', this.feeTier)
    },
    async getShares () {
      this.shares = await this.keeperContract.methods.balanceOf(ethereum.selectedAddress).call()
      console.log('user address=', ethereum.selectedAddress, ';shares=', this.shares)
      this.keeperContract.methods
        .totalSupply()
        .call()
        .then(val => {
          console.log('totalSupply= ' + val)
          this.totalShares = val
        })
    },
    setWalletStatus () {
      if (this.isConnected) {
        console.log('call disconnect')
        this.isConnected = false
        // ethereum.on('disconnect', error => { console.log(error) })
        this.StatusButtonText = 'Connect Wallet'
      } else {
        if (ethereum.isConnected() && this.currentAccount != null) {
          this.isConnected = true
          console.log('this.currentAccount=' + this.currentAccount)
          this.StatusButtonText = this.currentAccount
          // this.$refs.supplyliq.checkConnectionStatus()
        } else {
          this.connectWallet()
        }
      }
    },
    connectWallet () {
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
        console.log('account status:' + ethereum.isConnected())
        // this.$refs.supplyliq.checkConnectionStatus()
        // this.getMyLiquidity()
      }
    }
  }
}
</script>

<style scoped>
.text {
  font-size: 14px;
}

.item {
  padding: 18px 0;
}

.box-card {
  width: 100%;
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

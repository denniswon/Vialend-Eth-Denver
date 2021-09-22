<template>
  <div>
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
                         height="33" />&nbsp;&nbsp;SwapStream1
                  </div>
                  <div></div>
                </a>
              </div>
              <nav class="main_nav">
                <ul class="d-flex flex-row align-items-center justify-content-start">
                  <li><a href="#">Dashboard</a></li>
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
      <button type="button"
              class="btn btn-primary btn-sm"
              @click="getToken0">GetToken0</button>
      &nbsp;
      <button type="button"
              class="btn btn-primary btn-sm"
              @click="getTickLower">GetTickLower</button>
      &nbsp;
      <button type="button"
              class="btn btn-primary btn-sm"
              @click="getPositions">GetPositions</button>
      &nbsp;
      <button type="button"
              class="btn btn-primary btn-sm"
              @click="priceToTick">PriceToTick</button>
      &nbsp;
      <button type="button"
              class="btn btn-primary btn-sm"
              @click="toPrice">toPrice</button>

    </div>
  </div>
</template>

<script>
import Web3 from 'web3'
// import { ethers } from 'ethers'
import contractABI from '../ABI/UniswapV3PoolABI.json'
import trackerABI from '../ABI/nonfungiblePositionManagerABI.json'
import invariant from 'tiny-invariant'
import { TickMath } from '../utils/tickMath'
import { encodeSqrtRatioX96 } from '../utils/encodeSqrtRatioX96'
import { nearestUsableTick } from '../utils/nearestUsableTick'
import { Price, Token, Fraction } from '@uniswap/sdk-core'
import { tickToPrice, priceToClosestTick } from '../utils/priceTickConversions'
import JSBI from 'jsbi'
import { Q192 } from '../internalConstants'

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
  // switch (_this.$route.path) {
  //   case '/':
  //     this.connectMetaMask()
  //     break
  //   default:
  // }
})
window.ethereum.on('networkChanged', () => {
  console.log('networkChanged')
  connectWallet()
})

export default {
  data () {
    return {
      vaultAddress: '0x04b1560f4f58612a24cf13531f4706c817e8a5fe',
      trackerAddress: '0xC36442b4a4522E871399CD717aBDD847Ab11FE88',
      currentAccount: null,
      keeperContract: null,
      trackerContract: null,
      isConnected: false,
      walletButtonClass: 'walletButton',
      connectClass: 'wallet_connected',
      disConnectClass: 'wallet_disconnected',
      StatusButtonText: 'Connect Wallet'
    }
  },
  created: function () {
    // this.showTreaty()
    console.log('load create function')
    // this.keeperContract = new ethers.Contract(
    //   this.vaultAddress,
    //   contractABI,
    //   ethers.provider
    // )
    this.keeperContract = new web3.eth.Contract(
      contractABI,
      this.vaultAddress
    )
    this.trackerContract = new web3.eth.Contract(
      trackerABI,
      this.trackerAddress
    )
  },
  mounted () {
    window.connectWallet = this.connectWallet
  },
  methods: {
    getToken (num) {
      var sortOrder, decimals, chainId
      if (num === 0) {
        sortOrder = 0
        decimals = 18
        chainId = 5
        return new Token(
          chainId,
          `0x88177e1a55c6Ca956A738abbF6d87148217a8Cb0`,
          decimals,
          `T0`,
          `token0`
        )
      } else {
        sortOrder = 1
        decimals = 0
        chainId = 5
        return new Token(
          chainId,
          `0x495A3648AfDeb15Ce7B0cDDff44EAeB5E014cEAD`,
          decimals,
          `T1`,
          `token1`
        )
      }

      // if (sortOrder > 9 || sortOrder % 1 !== 0) throw new Error('invalid sort order')
    },
    priceToTick () {
      const token0 = this.getToken(0)
      const token1 = this.getToken(1)
      console.log('tick=', priceToClosestTick(new Price(token1, token0, 2799, 3200)))
    },
    toPrice () {
      const token0 = this.getToken(0)
      const token1 = this.getToken(1)
      const sqrtRatioX96 = TickMath.getSqrtRatioAtTick(-196950)

      const ratioX192 = JSBI.multiply(sqrtRatioX96, sqrtRatioX96)
      console.log('sqrtRatioX96=', sqrtRatioX96)
      console.log('ratioX192=', ratioX192)
      console.log('price111=', new Fraction(ratioX192, Q192).toSignificant(5))

      console.log(tickToPrice(token1, token0, -196950))
      console.log(tickToPrice(token1, token0, -195610))
      console.log(tickToPrice(token1, token0, -196950).toSignificant(5))
      console.log(tickToPrice(token1, token0, -195610).toSignificant(5))
    },
    getPositions () {
      this.trackerContract.methods
        .positions('0x04b1560f4f58612a24cf13531f4706c817e8a5fe')
        .call()
        .then(val => {
          console.log('val=' + val)
        })
    },
    getTickLower () {
      console.log('tickLowerA = ', Math.ceil(-887272 / 60) * 60)
      console.log('tickLowerB = ', nearestUsableTick(-200333, 60) - 120)
      console.log('price = ', TickMath.getSqrtRatioAtTick(-200333))

      this.getUpperLower()
    },
    getUpperLower () {
      console.log('tickUpper = ', nearestUsableTick(-200321, 60) + 120)
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
          this.$refs.supplyliq.checkConnectionStatus()
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
        this.$refs.supplyliq.checkConnectionStatus()
        this.getMyLiquidity()
      }
    },
    async getToken0 () {
      if (this.keeperContract != null) {
        // await this.keeperContract.fee()
        // console.log('token0=', token0)
        // console.log('BalanceOf=' + myLiq + '|' + JSON.stringify(myLiq))
        this.keeperContract.methods
          .factory()
          .call()
          .then(val => {
            console.log('factory=' + val)
          })
        this.keeperContract.methods
          .token0()
          .call()
          .then(val => {
            console.log('token0=' + val)
          })
        this.keeperContract.methods
          .token1()
          .call()
          .then(val => {
            console.log('token1=' + val)
          })
        this.keeperContract.methods
          .fee()
          .call()
          .then(val => {
            console.log('fee=' + val)
          })
        this.keeperContract.methods
          .tickSpacing()
          .call()
          .then(val => {
            console.log('tickSpacing=' + val)
          })
        this.keeperContract.methods
          .maxLiquidityPerTick()
          .call()
          .then(val => {
            console.log('maxLiquidityPerTick=' + val)
          })
        this.keeperContract.methods
          .liquidity()
          .call()
          .then(val => {
            console.log('liquidity=' + val)
          })
        this.keeperContract.methods
          .slot0()
          .call()
          .then(val => {
            console.log('slot0=' + JSON.stringify(val))
          })
        // this.keeperContract.methods
        //   .tickCurrent()
        //   .call()
        //   .then(val => {
        //     console.log('tickCurrent=' + JSON.stringify(val))
        //   })
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

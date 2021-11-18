import Vue from 'vue'
import Vuex from 'vuex'
import getters from './getters'
import app from './modules/app'
import settings from './modules/settings'
import user from './modules/user'
import ArrayList from '../common/ArrayList'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    name: 'Vialend',
    pairsList: new ArrayList(),
    availableChainId: [5],
    tokenExchangeTable: [
      {
        id: '32610',
        symbol: 'WETH',
        name: 'WETH',
        nameid: 'weth',
        rank: 2059,
        price_usd: '0',
        percent_change_24h: '0',
        percent_change_1h: '0',
        percent_change_7d: '0',
        price_btc: '0',
        market_cap_usd: '0',
        volume24: 0,
        volume24a: 0,
        csupply: '',
        tsupply: '',
        msupply: ''
      },
      {
        id: '33285',
        symbol: 'USDC',
        name: 'USD Coin',
        nameid: 'usd-coin',
        rank: 10,
        price_usd: '0',
        percent_change_24h: '0',
        percent_change_1h: '0',
        percent_change_7d: '0',
        price_btc: '0',
        market_cap_usd: '0',
        volume24: 0,
        volume24a: 0,
        csupply: '',
        tsupply: '',
        msupply: ''
      },
      {
        id: '32417',
        symbol: 'DAI',
        name: 'Dai',
        nameid: 'dai',
        rank: 287,
        price_usd: '0',
        percent_change_24h: '0',
        percent_change_1h: '0',
        percent_change_7d: '0',
        price_btc: '0',
        market_cap_usd: '0',
        volume24: 0,
        volume24a: 0,
        csupply: '',
        tsupply: '',
        msupply: ''
      }
    ],
    currentAccount: '',
    isConnected: false,
    StatusButtonText: '',
    chainId: 0,
    allTokensList: null,
    token0RateOfUSD: 0,
    token1RateOfUSD: 0,
    tokenExchangeRateLoaded: false
  },
  modules: {
    app,
    settings,
    user
  },
  getters,
  actions: {
    setApproveStatus (state, dt) {
      // console.log('token=' + dt.token + ';status=' + dt.status)
      var key = dt.token + '_ApproveStatus'
      var value = dt.status
      localStorage.setItem(key, value)
      console.log('setItem,key=' + key + ';value=' + value)
      console.log('setItem after:val=' + localStorage.getItem(key))
    },
    getApproveStatus (obj, dt) {
      var key = dt.token + '_ApproveStatus'
      var value
      if (localStorage.getItem(key) == null || localStorage.getItem(key) === undefined) {
        return false
      } else {
        value = localStorage.getItem(key)
        console.log('getItem,key=' + key + ';value=' + value)
        return value
      }
    },
    setPairsList (state, dt) {
      this.state.pairsList = dt
      localStorage.setItem('pairslist', dt)
    },
    getPairsList (obj, dt) {
      var key = 'pairslist'
      if (localStorage.getItem(key) == null || localStorage.getItem(key) === undefined) {
        return null
      } else {
        return localStorage.getItem(key)
      }
    }
  }
})

export default store

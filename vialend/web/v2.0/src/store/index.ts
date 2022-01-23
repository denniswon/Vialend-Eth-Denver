import Vue from 'vue'
import Vuex from 'vuex'
import { getWeb3Instance } from '../common/Web3'

Vue.use(Vuex)

// Declare empty store first, dynamically register all modules later.
export default new Vuex.Store({
  state: {
    name: 'Vialend',
    availableChainId: [5],
    pairsData: undefined,
    bridgeAddress: '0x428EeA0B87f8E0f5653155057f58aaaBb667A3ec',
    factoryAddress: '',
    validNetwork: false,
    isAdmin: false,
    doDisconnect: false,
    tokenExchangeTable: [
      {
        id: '32610',
        symbol: 'WETH',
        name: 'WETH',
        nameid: 'weth',
        rank: 2059,
        price_usd: '4706.15',
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
        price_usd: '0.999044',
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
        price_usd: '1.00',
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
        id: '33422',
        symbol: 'WBTC',
        name: 'Wrapped Bitcoin',
        nameid: 'wrapped-bitcoin',
        rank: 17,
        price_usd: '54284.08',
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
  getters: {
    getPriceUSDBySymbol: (state) => (symbol: string) => {
      return state.tokenExchangeTable.find(item => item.symbol === symbol)
    }
  },
  mutations: {
  },
  actions: {
    setApproveStatus(state, dt) {
      // console.log('token=' + dt.token + ';status=' + dt.status)
      const key = dt.token + '_ApproveStatus'
      const value = dt.status
      localStorage.setItem(key, value)
      console.log('setItem,key=' + key + ';value=' + value)
      console.log('setItem after:val=' + localStorage.getItem(key))
    },
    getApproveStatus(obj, dt) {
      const key = dt.token + '_ApproveStatus'
      let value
      if (localStorage.getItem(key) == null || localStorage.getItem(key) === undefined) {
        return false
      } else {
        value = localStorage.getItem(key)
        console.log('getItem,key=' + key + ';value=' + value)
        return value
      }
    },
    setSessionData(state, dt) {
      sessionStorage.setItem(dt.key, dt.value)
    },
    getSessionData(state, dt) {
      return sessionStorage.getItem(dt.key)
    },
    getChainID() {
      const web3 = getWeb3Instance()
      return web3.eth.getChainId()
    },
    async checkChain() {
      const web3 = getWeb3Instance()
      this.state.chainId = await web3.eth.getChainId()
      console.log('check chain id=', this.state.chainId)
      if (this.state.availableChainId.includes(this.state.chainId)) {
        this.state.validNetwork = true
        return true
      } else {
        this.state.isAdmin = false
        this.state.isConnected = false
        this.state.validNetwork = false
        return false
      }
    }
  }
})

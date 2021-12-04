import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    name: 'Vialend',
    availableChainId: [5],
    validNetwork: false,
    isAdmin: false,
    tokenExchangeTable: [
      { id: '32610',
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
        msupply: '' },
      { id: '33285',
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
        msupply: '' },
      { id: '32417',
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
        msupply: '' },
      { id: '33422',
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
        msupply: '' }
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
    }
  }
})

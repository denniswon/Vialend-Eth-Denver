import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    name: 'SwapStream.',
    isConnected: false,
    StatusButtonText: '',
    chainId: 0,
    allTokensList: null,
    token0RateOfUSD: 0,
    token1RateOfUSD: 0,
    tokenExchangeRateLoaded: false
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

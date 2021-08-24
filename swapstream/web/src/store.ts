import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {

  },
  mutations: {

  },
  actions: {
    setApproveStatus (state, dt) {
      // console.log('token=' + dt.token + ';status=' + dt.status)
      var key = dt.token + '_ApproveStatus'
      var value = dt.status
      sessionStorage.setItem(key, value)
      console.log('setItem,key=' + key + ';value=' + value)
      console.log('setItem after:val=' + sessionStorage.getItem(key))
    },
    getApproveStatus (obj, dt) {
      var key = dt.token + '_ApproveStatus'
      var value
      if (sessionStorage.getItem(key) == null || sessionStorage.getItem(key) === undefined) {
        return false
      } else {
        value = sessionStorage.getItem(key)
        console.log('getItem,key=' + key + ';value=' + value)
        return value
      }
    }
  }
})

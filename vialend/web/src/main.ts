import Vue from 'vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import App from './App.vue'
import store from './store'
import router from './router'
import Web3 from 'web3'
import axios from 'axios'
import VueAxios from 'vue-axios'

Vue.config.productionTip = false
Vue.use(ElementUI, VueAxios, axios)
Vue.prototype.Web3 = Web3

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')

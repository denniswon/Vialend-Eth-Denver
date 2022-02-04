<template>
  <div id="app" element-loading-text="Loading" element-loading-spinner="el-icon-loading" element-loading-background="rgba(0, 0, 0, 0.8)">
    <router-view />
  </div>
</template>

<script lang="ts">
// import Vue from 'vue'
import axios from 'axios'
import { Component, Vue } from 'vue-property-decorator'
import { UserModule } from '@/store/modules/user'
import { checkChain } from '@/constants/chains'
import { changeAccount, checkValidAdmin } from '@/common/Account'

const ethereum = (window as any).ethereum

@Component({
  name: 'App',
  components: {}
})
export default class extends Vue {
  async getTokensRateOfUSD() {
    let response
    console.log('this.$store.state.tokenExchangeTable=', this.$store.state.tokenExchangeTable)
    for (let i = 0; i < this.$store.state.tokenExchangeTable.length; i++) {
      // console.log('token' + i + '=', this.$store.state.tokenExchangeTable[i].symbol)
      response = await axios.get('https://api.coinlore.net/api/ticker/?id=' + this.$store.state.tokenExchangeTable[i].id)
      if (response !== undefined && response !== null) {
        console.log('response=', response.data.length)
        if (response.data.length > 0) {
          this.$store.state.tokenExchangeTable[i].price_usd = response.data[0].price_usd
        }
      }
    }
    this.$store.state.tokenExchangeRateLoaded = true
  }

  clearSessionData() {
    this.$store.dispatch('removeSessionData', { key: 'pairBaseInfo' })
    this.$store.dispatch('removeSessionData', { key: 'factoryAddress' })
  }

  async created() {
    ethereum.on('accountsChanged', async () => {
      console.log('accountsChanged')
      if (await checkChain()) {
        changeAccount()
        checkValidAdmin()
      }
    })

    ethereum.on('networkChanged', async () => {
      console.log('networkChanged')
      this.clearSessionData()
      if (await checkChain()) {
        changeAccount()
        checkValidAdmin()
      } else {
        this.$store.state.isConnected = false
        this.$store.state.currentAccount = ''
        this.$store.state.isAdmin = false
        this.$store.state.doDisconnect = true
        UserModule.ChangeRoles('user').then(() => {
          this.$emit('role changed')
        })
      }
    })

    ethereum.on('disconnect', () => {
      console.log('metamask disconnect')
      this.$store.state.currentAccount = ''
      this.$store.state.isConnected = false
    })

    if (await checkChain()) {
      if ((window as any).ethereum.selectedAddress !== null && (window as any).ethereum.selectedAddress !== undefined) {
        if (!this.$store.state.doDisconnect) {
          console.log('currentAccount is connected.')
          this.$store.state.currentAccount = (window as any).ethereum.selectedAddress
          this.$store.state.isConnected = true
          checkValidAdmin()
        }
      } else {
        console.log('ethereum.selectedAddress is null.')
        this.$store.state.currentAccount = ''
        this.$store.state.isConnected = false
      }
    } else {
      console.log('there is invalid network.')
      this.$store.state.currentAccount = ''
      this.$store.state.isConnected = false
    }
    console.log('App load chainId = ', this.$store.state.chainId)
    this.getTokensRateOfUSD()
  }
}
</script>

<style>
a {
  text-decoration: none !important;
}
@media only screen and (max-width: 767px) {
  .menu_products {
    display: block !important;
  }
}

@media only screen and (min-width: 768px) {
  .menu_products {
    display: block !important;
  }
}

@media only screen and (min-width: 992px) {
  .menu_products {
    display: none !important;
  }
}
@media only screen and (min-width: 1200px) {
  .menu_products {
    display: none !important;
  }
}
@media only screen and (min-width: 1920px) {
  .menu_products {
    display: none !important;
  }
}
</style>

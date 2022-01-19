<template>
  <div id="app"
       element-loading-text="Loading"
       element-loading-spinner="el-icon-loading"
       element-loading-background="rgba(0, 0, 0, 0.8)">
    <router-view />
  </div>
</template>

<script lang="ts">
// import Vue from 'vue'
import { Component, Vue } from 'vue-property-decorator'
import axios from 'axios'

@Component({
  name: 'App',
  components: { }
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

  created() {
    this.getTokensRateOfUSD()
  }
}
</script>

<style>
a{
    text-decoration:none !important;
}
@media only screen and (max-width: 767px) {
    .menu_products{
        display:block !important;
    }
}

@media only screen and (min-width: 768px){
    .menu_products{
        display:block !important;
    }
}

@media only screen and (min-width: 992px){
    .menu_products{
        display:none !important;
    }
}
@media only screen and (min-width: 1200px){
    .menu_products{
        display:none !important;
    }
}
@media only screen and (min-width: 1920px){
    .menu_products{
        display:none !important;
    }
}
</style>

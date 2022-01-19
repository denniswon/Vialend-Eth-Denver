<template>
    <div id="setting_container" v-loading="loadTickerLoading">
          <div class="token_exchange_form">
            <el-card class="box-card-currency">
              <div slot="header"
                   class="clearfix">
                <span>Currency Converter</span>
              </div>
              <div class="text item">
                <el-form :inline="true"
                         class="demo-form-inline">
                  <el-form-item>
                    <el-input v-model="currTokenVal"
                              placeholder="1"></el-input>
                  </el-form-item>
                  <el-form-item>
                    <el-select v-model="currTokenId"
                               placeholder="Please select"
                               @change="exchangeTokenChange">
                      <el-option v-for="item in tokensList"
                                 :key="item.nameid"
                                 :label="item.symbol"
                                 :value="item.symbol">
                      </el-option>
                    </el-select>
                    =
                  </el-form-item>
                  <el-form-item>
                    <el-input v-model="usdTokenVal"></el-input>
                  </el-form-item>
                  <el-form-item>
                    USD
                  </el-form-item>
                </el-form>
              </div>
            </el-card>
          </div>
        </div>
</template>

<script>
import { Component, Vue, Watch } from 'vue-property-decorator'
import axios from 'axios'

@Component({
  name: 'Settings',
  components: { }
})
export default class extends Vue {
    loadTickerLoading = false
    currTokenVal = 1
    currTokenId = 'BTC'
    tokensList = null
    usdTokenVal = 0
    priceUSD = 0

    @Watch('currTokenVal')
    watchCurrTokenVal(val) {
      if (val !== '') {
        if (!isNaN(val)) {
          console.log('currTokenVal=', val)
          this.usdTokenVal = val * this.priceUSD
        }
      }
    }

    exchangeTokenChange(val) {
      console.log('token change=', val)
      for (let i = 0; i < this.tokensList.length; i++) {
        if (this.tokensList[i].symbol === val) {
          this.usdTokenVal = this.tokensList[i].price_usd
          this.priceUSD = this.tokensList[i].price_usd
          this.currTokenVal = 1
          break
        }
      }
    }

    async created() {
      this.loadTickerLoading = true
      const response = await axios.get('https://api.coinlore.net/api/tickers/')
      if (response !== undefined && response !== null) {
        console.log('response=', response.data.data)
        if (response.data.data.length > 0) {
          this.tokensList = response.data.data
          this.priceUSD = this.tokensList[0].price_usd
          this.usdTokenVal = this.priceUSD
          console.log('this.tokensList=', this.tokensList)
        }
      }
      this.loadTickerLoading = false
    }
}
</script>

<style scoped>
.token_exchange_form {
  width: 100%;
  font-size: 20px;
  color: black;
  padding: 10px;
}
.box-card-currency {
  width: 900px;
  margin: 10px;
}
.clearfix:before,
.clearfix:after {
  display: table;
  content: '';
}
.clearfix:after {
  clear: both;
}
.el-card /deep/ .el-card__header {
  background-color: #102130 !important;
  color: #ffffff;
  height: 80px;
}
.el-card /deep/ .el-card__body {
  background-color: #121218 !important;
  padding: 0px;
  color: #ffffff;
  padding:30px;
}
</style>

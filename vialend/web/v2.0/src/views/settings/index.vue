<template>
    <div id="setting_container" v-loading="loadTickerLoading">
          <div class="token_exchange_form">
            <el-tabs type="border-card">
  <el-tab-pane label="CurrencyConverter">
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
  </el-tab-pane>
  <el-tab-pane label="Emergency">
    <div class="pairs-selector">
      <el-select
        v-model="pairSelectedIndex"
        placeholder="Please select pairs"
        v-if="pairsData.pairsList && pairsData.pairsList.size() > 0"
      >
        <el-option
          v-for="pair in pairsData.pairsList._getData()"
          :key="pair.index"
          :label="pair.token0.symbol + '/' + pair.token1.symbol"
          :value="pair.index"
        ></el-option>
      </el-select>
    </div>
    <div class="clear"></div>
    <div class="stat-container" v-loading="statDataLoading">
    <el-input v-model="stat"  style="width:300px;"></el-input>&nbsp;&nbsp;
    <el-button type="primary" @click="setStat" v-loading="changeStatLoading">Change Stat</el-button>
    </div>
  </el-tab-pane>
</el-tabs>
          </div>
        </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator'
import { contractInstance } from '../../common/Web3'
import PairsData from '../../common/PairsData'
import Pairs from '../../model/Pairs'
import axios from 'axios'

const VaultFactoryABI = require('../../abi/VaultFactory.json')
const ethereum = (window as any).ethereum

@Component({
  name: 'Settings',
  components: { }
})
export default class extends Vue {
    pairsData = new PairsData()
    pairsList = this.pairsData.pairsList
    currentPair = new Pairs()
    pairSelectedIndex = ''
    loadTickerLoading = false
    currTokenVal = 1
    currTokenId = 'BTC'
    tokensList = null
    usdTokenVal = 0
    priceUSD = 0
    factoryAddress = ''
    statDataLoading = false
    stat = '0'
    changeStatLoading = false

    @Watch('currTokenVal')
    watchCurrTokenVal(val:number) {
      if (val >= 0) {
        if (!isNaN(val)) {
          console.log('currTokenVal=', val)
          this.usdTokenVal = val * this.priceUSD
        }
      }
    }

    @Watch('pairSelectedIndex')
    async watchPairSelectedIndex(newVal: number, oldVal: number) {
      console.log('pairSelectedIndex:', newVal, ',old value:', oldVal)
      this.statDataLoading = true
      this.factoryAddress = await this.$store.dispatch('getSessionData', { key: 'factoryAddress' })
      this.currentPair = this.pairsData.pairsList.get(newVal)
      if (this.factoryAddress !== undefined && this.factoryAddress !== '') {
        const factoryContract = await contractInstance(VaultFactoryABI, this.factoryAddress)
        console.log('strategy address = ', this.currentPair.strategyAddress)
        console.log('vault address = ', this.currentPair.vaultAddress)
        this.stat = await factoryContract.methods.stat(this.currentPair.strategyAddress, this.currentPair.vaultAddress).call()
      }
      this.statDataLoading = false
      console.log('Emergency stat=', this.stat)
    }

    exchangeTokenChange(val:number) {
      console.log('token change=', val)
      for (let i = 0; i < (this.tokensList as any).length; i++) {
        if ((this.tokensList as any)[i].symbol === val) {
          this.usdTokenVal = (this.tokensList as any)[i].price_usd
          this.priceUSD = (this.tokensList as any)[i].price_usd
          this.currTokenVal = 1
          break
        }
      }
    }

    async setStat() {
      console.log('set stat value:', this.stat)
      const _this = this
      const factoryContract = await contractInstance(VaultFactoryABI, this.factoryAddress)
      // await factoryContract.methods.stat(this.currentPair.strategyAddress, this.currentPair.vaultAddress, this.stat).call()
      if (factoryContract !== null) {
        this.changeStatLoading = true
        factoryContract.methods
          .changeStat(
            this.currentPair.strategyAddress,
            this.currentPair.vaultAddress,
            this.stat
          )
          .send({
            from: ethereum.selectedAddress,
            // gasPrice: '80000000000',
            // gas: 600000,
            value: 0
          })
          .on('confirmation', function(confirmationNumber:number, receipt:any) {
            console.log((new Date()).toLocaleString(), ':{changeStat confirm number:', confirmationNumber, ',receipt:', receipt.status, '}')
          })
          .on('receipt', function(receipt:any) {
            if (_this.changeStatLoading === true) {
              _this.changeStatLoading = false
              if (receipt.status) {
                // _this.goToEtherscan = 'https://goerli.etherscan.io/tx/' + receipt.transactionHash
                console.log('changeStat receipt:', receipt)
                _this.$message('changeStat submitted!')
              } else {
                _this.$message('changeStat failed!')
              }
              console.log((new Date()).toLocaleString(), 'changeStat receipt status:', receipt)
            }
          })
          .on('error', function(error:any) {
            _this.$message.error(error)
            _this.changeStatLoading = false
            console.log('changeStat error:', error)
          })
      }
      this.changeStatLoading = false
    }

    async created() {
      this.loadTickerLoading = true
      if (this.$store.state.validNetwork && this.$store.state.isConnected && this.pairsData.pairsList.size() === 0) {
        console.log('pair loading')
        await this.pairsData.loadPairsInfo()
      }
      const response = await axios.get('https://api.coinlore.net/api/tickers/')
      if (response !== undefined && response !== null) {
        console.log('response=', response.data.data)
        if (response.data.data.length > 0) {
          this.tokensList = response.data.data
          this.priceUSD = (this.tokensList as any)[0].price_usd
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
.pairs-selector {
  width: 100%;
  margin-bottom: 20px;
}
</style>

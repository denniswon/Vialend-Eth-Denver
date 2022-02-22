<template>
  <div class="portfolio-component">
    <div class="portfolio-container">
      <div class="pf-title">Portfolio Summary</div>
      <div class="part1">
        <div class="title">Balances</div>
        <div class="bal-val">
          <div class="amount" style="float:left">$6558.45</div>
          <div class="unit" style="float:left">USD</div>
          <div class="filter" style="float:right">
            <ul>
              <li>1W</li>
              <li>1M</li>
              <li>ALL</li>
            </ul>
          </div>
        </div>
        <div style="clear:both"></div>
        <div class="chart">
          <apexchart width="100%" height="200" type="area" :options="chartOptions" :series="series"></apexchart>
        </div>
        <div class="yieldEamed">
          <el-row :gutter="10">
            <el-col class="col1" :xs="24" :sm="8" :md="8" :lg="8" :xl="8">
              <span class="title">Yield Earned</span>
              <br />
              <span class="amount">+$793.90</span>
            </el-col>
            <el-col class="col2" :xs="24" :sm="8" :md="8" :lg="8" :xl="8">
              <span class="title">ROI</span>
              <br />
              <span class="amount">+0.18%</span>
            </el-col>
            <el-col class="col3" :xs="24" :sm="8" :md="8" :lg="8" :xl="8">
              <span class="title">$VIA Balance</span>
              <br />
              <span class="amount">0.00</span>
            </el-col>
          </el-row>
        </div>
      </div>
      <div style="clear:both;height:30px;"></div>
      <div class="pf-title">Positions</div>
      <div class="part2">
        <table class="table" v-loading="pairsData.pairsBaseInfoLoading" element-loading-text="Loading" element-loading-spinner="el-icon-loading" element-loading-background="rgba(0, 0, 0, 0.8)">
          <tbody v-for="pair in pairsList._getData()" :key="pair.id">
            <tr v-if="pair.disabled === false">
              <td class="symbol">
                <cryptoicon :symbol="pair.token1.symbol" size="60" />
              </td>
              <td>
                <span class="token-title">{{ pair.token0.symbol }}-{{ pair.token1.symbol }}</span>
                <br />
                <span class="token-desc">Uniswap V3 Compound Yield Generator</span>
              </td>
              <td>
                <span class="perc">{{ Number(Number(pair.currentAPR).toFixed(2)) }}%</span>
              </td>
              <td>
                <span class="amount">
                  {{ Number(pair.tvlTotal0).toFixed(2) }} / {{ Number(pair.tvlTotal1).toFixed(2) }}
                  <br />
                  ${{ (Number(pair.tvlTotal0USD) + Number(pair.tvlTotal1USD)).toFixed(2) }}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator'
import PairsData from '../../common/PairsData'

@Component({
  name: 'Portfolio',
  components: {}
})
export default class extends Vue {
  exchangeTimer: any
  pairsData = new PairsData()
  pairsList = this.pairsData.pairsList

  chartOptions = {
    chart: {
      id: 'vuechart-example',
      type: 'area',
      toolbar: {
        show: false
      },
      sparkline: {
        enabled: false
      }
    },
    fill: {
      type: 'gradient',
      gradient: {
        shadeIntensity: 1,
        opacityFrom: 0.7,
        opacityTo: 0.9,
        stops: [0, 90, 100]
      }
    },
    xaxis: {
      categories: [1991, 1992, 1993, 1994, 1995, 1996, 1997, 1998]
    },
    grid: {
      show: false
    }
  }

  series = [
    {
      name: 'series-1',
      data: [30, 40, 35, 50, 49, 60, 70, 91]
    }
  ]

  get pairsListCount() {
    return this.pairsData.pairsList.size()
  }

  @Watch('pairsListCount')
  watchPairsListChange(newVal: number, oldVal: number) {
    console.log('Change of PairsList:', newVal, ',old value:', oldVal)
    this.pairsData.calculateAPR = true
    for (let i = 0; i < newVal; i++) {
      console.log('pair id=', this.pairsData.pairsList.get(i).id)
      if (!this.pairsData.pairsList.get(i).loadDataCompleted) {
        this.pairsData.pairsList.get(i).gettingData = true
        this.pairsData.getPairPublicData(this.pairsData.pairsList.get(i))
      }
    }
  }

  @Watch('$store.state.currentAccount')
  watchCurrentAccount(newVal: string, oldVal: string) {
    console.log('currentAccount:', newVal, ';previousAccount:', oldVal)
    if (newVal !== '' && this.$store.state.validNetwork) {
      console.log('currentAccount changed,pairlist size:', this.pairsData.pairsList.size())
      if (this.pairsData.pairsList.size() === 0 && !this.pairsData.pairsBaseInfoLoading) {
        this.pairsData.loadPairsInfo()
      } else {
        // this.pairsData.pairsListComplete = false
        // this.loadMyPairsList()
      }
    }
  }

  getTokenRateOfUSD(symbol: string) {
    // console.log('symbol=', symbol)
    if (this.$store.getters.getPriceUSDBySymbol(symbol) !== undefined) {
      // console.log('eth exchange table of ETH:', this.$store.getters.getPriceUSDBySymbol(symbol).price_usd)
      return this.$store.getters.getPriceUSDBySymbol(symbol).price_usd
    } else {
      return -1
    }
  }

  exchangeTokensIntoUSD() {
    if (this.$store.state.tokenExchangeRateLoaded) {
      // this.myValueLocked_Token0Sum = 0
      // this.myValueLocked_Token1Sum = 0
      let token0RateOfUSD = 0
      let token1RateOfUSD = 0
      for (let i = 0; i < this.pairsList.size(); i++) {
        token0RateOfUSD = this.getTokenRateOfUSD(this.pairsList.get(i).token0.symbol)
        token1RateOfUSD = this.getTokenRateOfUSD(this.pairsList.get(i).token1.symbol)
        console.log('token0 RateOfUSD=', token0RateOfUSD)
        console.log('token1 RateOfUSD=', token1RateOfUSD)
        if (this.pairsList.get(i).tvlTotal0 >= 0) this.pairsList.get(i).tvlTotal0USD = Number(this.pairsList.get(i).tvlTotal0) * Number(token0RateOfUSD)
        if (this.pairsList.get(i).tvlTotal1 >= 0) this.pairsList.get(i).tvlTotal1USD = Number(this.pairsList.get(i).tvlTotal1) * Number(token1RateOfUSD)
        this.pairsList.get(i).tvl = this.pairsList.get(i).tvlTotal0USD + this.pairsList.get(i).tvlTotal1USD
        //  if (this.pairsList.get(i).myValueToken0Locked >= 0) {
        //    this.myValueLocked_Token0Sum += this.pairsList.get(i).myValueToken0Locked
        //    this.pairsList.get(i).myValueToken0USDLocked = Number(this.pairsList.get(i).myValueToken0Locked) * Number(token0RateOfUSD)
        //  }
        //  if (this.pairsList.get(i).myValueToken1Locked >= 0) {
        //    this.myValueLocked_Token1Sum += this.pairsList.get(i).myValueToken1Locked
        //    this.pairsList.get(i).myValueToken1USDLocked = Number(this.pairsList.get(i).myValueToken1Locked) * Number(token1RateOfUSD)
        //  }
      }
      //    if (this.myValueLocked_Token0Sum >= 0 && this.myValueLocked_Token1Sum >= 0) {
      //      this.myValueLocked = this.myValueLocked_Token0Sum * Number(token0RateOfUSD) + this.myValueLocked_Token1Sum * Number(token1RateOfUSD)
      //    }
      // console.log('myValueLocked_Token0Sum=', this.myValueLocked_Token0Sum, 'myValueLocked_Token1Sum=', this.myValueLocked_Token1Sum)
      // this.myEarnedValue = this.fees0Total * Number(this.$store.state.token0RateOfUSD) + this.fees1Total * Number(this.$store.state.token1RateOfUSD)
      // this.myEarnedValue = Number(this.userFee0Total) * Number(token0RateOfUSD) + Number(this.userFee1Total) * Number(token1RateOfUSD)
    }
    console.log('tokenExchangeRateLoaded1:', this.$store.state.tokenExchangeRateLoaded, 'this.pairsList.size()=', this.pairsList.size())
  }

  created() {
    this.showDebugLog()
    if (this.$store.state.validNetwork && this.$store.state.isConnected) {
      if (this.pairsData.pairsList.size() > 0) {
        for (let i = 0; i < this.pairsList.size(); i++) {
          this.pairsData.pairsList.get(i).gettingData = true
          this.pairsData.calculateAPR = true
          this.pairsData.getPairPublicData(this.pairsData.pairsList.get(i))
        }
      } else {
        this.pairsData.loadPairsInfo()
      }
    }
  }

  mounted() {
    // this.exchangeTimer = setInterval(this.exchangeTokensIntoUSD, 1000)
  }

  beforeDestroy() {
    clearInterval(this.exchangeTimer)
  }

  showDebugLog() {
    console.log('this.$store.state.validNetwork=', this.$store.state.validNetwork)
    console.log('this.$store.state.isConnected=', this.$store.state.isConnected)
    console.log('this.pairsData.pairsList.size()=', this.pairsData.pairsList.size())
  }

  // mounted() {
  //   const ctx = document.getElementById('myChart')
  //   const myChart = new Chart(ctx as any, {
  //     type: 'line',
  //     data: data,
  //     options: {}
  //   })
  //   myChart.resize(200, 50)
  // }
}
</script>

<style scoped>
.portfolio-component {
  position: absolute;
  width: 100%;
  height: 100%;
  padding: 0;
  margin: 0 auto;
  float: left;
}

.portfolio-container {
  padding: 30px;
  margin: 0 auto;
}

.pf-title {
  padding: 20px 10px 20px 10px;
  font-size: 26px;
  font-family: Arial, Helvetica, sans-serif;
  margin-top: 20px;
}
.part1,
.part2 {
  height: auto;
  border: 2px solid #000000;
  background-color: #121218;
  color: #707074;
}
.part1 .title {
  font-family: Arial, Helvetica, sans-serif;
  font-size: 16px;
  padding: 20px;
  color: #acacab;
}
.part1 .bal-val {
  padding: 0px;
}
.part1 .bal-val .amount {
  font-size: 26px;
  margin-left: 20px;
  color: #ffffff;
}
.part1 .bal-val .unit {
  margin-left: 10px;
  font-size: 16px;
  color: #acacab;
}
.part1 .bal-val .filter li {
  float: left;
  list-style: none;
  width: 50px;
}
.part1 .chart {
  position: relative;
  width: 100%;
  height: auto;
}
.part1 .yieldEamed {
  width: 100%;
  height: auto;
}
.part1 .yieldEamed .title {
  font-family: Arial, Helvetica, sans-serif;
  font-size: 16px;
  padding: 20px;
  color: #acacab;
}
.part1 .yieldEamed .amount {
  font-size: 16px;
  margin-left: 20px;
  color: #ffffff;
}
.el-row {
  height: auto;
  margin-left: 0px !important;
  margin-right: 0px !important;
}
.el-col {
  height: 100%;
}

.pf-positions {
  padding: 10px;
  font-size: 16px;
  font-family: Arial, Helvetica, sans-serif;
}
.table > :not(:first-child) {
  border-top: none;
}
.part2 {
  height: auto;
  border: 0px solid #000000;
  border-radius: 16px;
}
.part2 /deep/ .el-loading-mask {
  border-radius: 16px !important;
}
.part2 .table {
  width: 100%;
  min-height: 83px;
}
.part2 .symbol {
  padding: 10px;
}
.part2 .token-title {
  font-size: 20px;
  font-family: Arial, Helvetica, sans-serif;
  font-weight: bold;
  color: #ffffff;
}
.part2 .token-desc {
  font-size: 16px;
  font-family: Arial, Helvetica, sans-serif;
  color: #acacab;
}
.part2 .perc {
  font-size: 16px;
  font-family: Arial, Helvetica, sans-serif;
  font-weight: bold;
  color: #ffffff;
}
.part2 .amount {
  font-size: 16px;
  font-family: Arial, Helvetica, sans-serif;
  font-weight: bold;
  color: #ffffff;
}
.el-row .col1,
.el-row .col2,
.el-row .col3 {
  height: 100px;
  padding: 20px 20px 20px 20px !important;
  line-height: 26px;
}
@media only screen and (max-width: 767px) {
  .portfolio-container {
    width: 100%;
  }
  .el-row .col1 {
    border-top: 2px solid #1c1c22;
  }
  .el-row .col2 {
    border-top: 2px solid #1c1c22;
  }
  .el-row .col3 {
    border-top: 2px solid #1c1c22;
  }
}

@media only screen and (min-width: 768px) {
  .portfolio-container {
    width: 630px;
  }
  .el-row .col1 {
    border-top: 2px solid #1c1c22;
    border-right: 2px solid #1c1c22;
  }
  .el-row .col2 {
    border-top: 2px solid #1c1c22;
    border-right: 2px solid #1c1c22;
  }
  .el-row .col3 {
    border-top: 2px solid #1c1c22;
  }
}

@media only screen and (min-width: 992px) {
  .portfolio-container {
    width: 1100px;
  }
  .el-row .col1 {
    border-top: 2px solid #1c1c22;
    border-right: 2px solid #1c1c22;
  }
  .el-row .col2 {
    border-top: 2px solid #1c1c22;
    border-right: 2px solid #1c1c22;
  }
  .el-row .col3 {
    border-top: 2px solid #1c1c22;
  }
}
@media only screen and (min-width: 1200px) {
  .portfolio-container {
    width: 1100px;
  }
  .el-row .col1 {
    border-top: 2px solid #1c1c22;
    border-right: 2px solid #1c1c22;
  }
  .el-row .col2 {
    border-top: 2px solid #1c1c22;
    border-right: 2px solid #1c1c22;
  }
  .el-row .col3 {
    border-top: 2px solid #1c1c22;
  }
}
@media only screen and (min-width: 1920px) {
  .portfolio-container {
    width: 1100px;
  }
  .el-row .col1 {
    border-top: 2px solid #1c1c22;
    border-right: 2px solid #1c1c22;
  }
  .el-row .col2 {
    border-top: 2px solid #1c1c22;
    border-right: 2px solid #1c1c22;
  }
  .el-row .col3 {
    border-top: 2px solid #1c1c22;
  }
}
</style>

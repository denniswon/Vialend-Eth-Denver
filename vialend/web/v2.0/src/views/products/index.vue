<template>
  <div class="pairs-list" v-loading="pairsData.pairsBaseInfoLoading">
    <!-- <MyPairsList v-bind:pairsData="pairsData"></MyPairsList> -->
    <div id="pairs_container" class="pairs-container">
      <div style="margin:15px;">
        <el-select
          v-model="listQuery.network"
          placeholder="Network"
          clearable
          style="width: 120px"
          class="filter-item"
          size="small"
        >
          <el-option
            v-for="item in networkOptions"
            :key="item"
            :label="item"
            :value="item"
          /> </el-select
        >&nbsp;&nbsp;
        <el-select
          v-model="listQuery.strategy"
          placeholder="Strategy"
          clearable
          style="width: 120px"
          class="filter-item"
          size="small"
        >
          <el-option
            v-for="item in strategyOptions"
            :key="item"
            :label="item"
            :value="item"
          />
        </el-select>
        <el-select
          v-model="listQuery.depositasset"
          placeholder="Deposit Asset"
          clearable
          style="width: 120px"
          class="filter-item"
          size="small"
        >
          <el-option
            v-for="item in depositassetOptions"
            :key="item"
            :label="item"
            :value="item"
          />
        </el-select>
        <el-select
          v-model="listQuery.sortby"
          placeholder="Sort By"
          clearable
          style="width: 120px"
          class="filter-item"
          size="small"
        >
          <el-option
            v-for="item in sortbyOptions"
            :key="item"
            :label="item"
            :value="item"
          />
        </el-select>
        <div class="phone">
          <svg-icon name="phone" width="26px" height="26px"></svg-icon>
        </div>
      </div>
      <div class="clear"></div>

      <el-row
        :gutter="20"
        class="panel-group"
        justify="center"
        v-if="pairsListCount > 0"
      >
        <el-col
          :xs="23"
          :sm="12"
          :md="7"
          :lg="7"
          class="card-panel-col"
          v-for="pair in pairsData.pairsList._getData()"
          :key="pair.id"
        >
            <vue-flip
              width="320px"
              horizontal="true"
              height="500px"
              v-model="pair.flipped"
            >
              <template v-slot:front class="front">
                <el-card
                  class="box-card"
                  v-loading="pair.gettingData"
                  element-loading-text="Loading pair data"
                  element-loading-spinner="el-icon-loading"
                  element-loading-background="rgba(0, 0, 0, 0.8)"
                >
                  <div slot="header" class="clearfix card-title">
                    <div class="block">
                      <el-tag size="small" hit>UNI-COMP</el-tag>&nbsp;&nbsp;
                      <el-tag
                        class="tag-item"
                        type="info"
                        effect="plain"
                        size="small"
                        >V1</el-tag
                      >
                    </div>

                    <span
                      style="float: right; padding: 3px 0;cursor: pointer;"
                      type="text"
                      @click="pair.flipped = !pair.flipped"
                    >
                      <i class="el-icon-menu"></i>
                    </span>
                  </div>
                  <div class="token-icon">
                    <cryptoicon :symbol="pair.token1.symbol" size="60" />
                  </div>
                  <div class="card-container">
                    <router-link :to="'products/' + pair.token0.symbol.concat('-', pair.token1.symbol)">
                    <div class="card-pairinfo">
                      <div class="pair-name text item">
                        <span
                          >{{ pair.token0.symbol }} -
                          {{ pair.token1.symbol }}</span
                        >
                      </div>
                      <div class="pair-intro text item">
                        <span>
                          Generates yield by providing concentrated liquidity on
                          Uniswap V3 and lending pools.
                        </span>
                      </div>
                      <div class="pair-apy text item">
                        <span class="f12">CURRENT PROJECTED YIELD(APY)</span>
                        <br />
                        <span class="f24"
                          >{{ Number(pair.netAPY).toFixed(2) }}%</span
                        >
                      </div>
                      <div class="text item">
                        <span class="pair-current-deposits"
                          >Current Deposits</span
                        ><br>
                        <span class="pair-deposits-val">
                          {{
                            pair.currentDeposits
                          }}&nbsp;{{ pair.token1.symbol }}
                        </span>
                      </div>
                      <div style="clear:both;"></div>
                      <div class="text item">
                        <el-slider
                          v-model="maxCapacity"
                          :show-tooltip="false"
                        ></el-slider>
                      </div>
                      <div style="clear:both;"></div>
                      <div class="text item">
                        <span class="pair-max-capacity">Max Capacity</span>
                        <span class="pair-capacity-val"
                          >100M&nbsp;{{ pair.token1.symbol }}</span
                        >
                      </div>
                    </div>
                    </router-link>
                    <div class="card-your-position">
                      <span class="text">Your Position</span>
                      <span class="pos-val">---</span>
                    </div>
                  </div>
                </el-card>
              </template>
              <template v-slot:back class="back">
                <el-card
                  class="box-card"
                  v-loading="pair.gettingData"
                  element-loading-text="Loading pair data"
                  element-loading-spinner="el-icon-loading"
                  element-loading-background="rgba(0, 0, 0, 0.8)"
                >
                  <div slot="header" class="clearfix card-title">
                    <span style="font-size:20px;"
                      >{{ pair.token0.symbol }} / {{ pair.token1.symbol }}</span
                    >
                    <span
                      style="float: right; padding: 3px 0"
                      type="text"
                      @click="pair.flipped = !pair.flipped"
                    >
                      <i class="el-icon-s-data"></i>
                    </span>
                  </div>
                  <router-link :to="'products/' + pair.token0.symbol.concat('-', pair.token1.symbol)">
                  <div class="text item">
                    <span>Token Detail</span>
                  </div>
                  <div class="text item">
                    <br />
                    <span>&nbsp;</span>
                  </div>
                  <div class="text item">
                    <br />
                    <span>&nbsp;</span>
                  </div>
                  <div class="text item">
                    <br />
                    <span>&nbsp;</span>
                  </div>
                  </router-link>
                </el-card>
              </template>
            </vue-flip>
        </el-col>
      </el-row>
    </div>
  </div>
  <!-- element-loading-text="Loading pairs list"
     element-loading-spinner="el-icon-loading"
     element-loading-background="rgba(0, 0, 0, 0.1)" -->
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator'
import PairsData from '../../common/PairsData'
// import MyPairsList from '../../components/PairsList/MyPairsList.vue'
import VueFlip from 'vue-flip'

@Component({
  name: 'Proudcts',
  components: { 'vue-flip': VueFlip }
})
export default class extends Vue {
  pairsData = new PairsData();
  // pairsList = this.pairsData.pairsList;
  maxCapacity = 60;
  reloadPairs = false;

  private listQuery = {
    page: 1,
    limit: 20,
    network: undefined,
    strategy: undefined,
    depositasset: undefined,
    sortby: undefined
  };

  private networkOptions = ['Ethereum', 'Polygon', 'Optimism'];
  private strategyOptions = [
    'UniV3+Compound',
    'UniV3+Aave',
    'UniV3 Charm-Style',
    'UniV3 Visor-STYLE',
    'Bancor3'
  ];

  private depositassetOptions = ['All', 'WETH', 'USDC', 'WBTC', 'DAI', 'USDT'];
  private sortbyOptions = ['None', 'APY', 'Current Deposits'];

  get pairsListCount() {
    return this.pairsData.pairsList.size()
  }

  async created() {
    this.pairsData.calculateAPY = true
    console.log('bridgeAddress value123=', this.pairsData.bridgeAddress)
    console.log('pairsList.size=', this.pairsData.pairsList.size())
    console.log('validNetwork=', this.$store.state.validNetwork)
    const validChain = await this.$store.dispatch('checkChain')
    console.log('product chainId=', validChain)
    if (validChain && this.pairsData.pairsList.size() === 0) {
      this.pairsData.loadPairsInfo()
    }
  }

  loadPairPublicData() {
    for (let i = 0; i < this.pairsData.pairsList.size(); i++) {
      if (!this.pairsData.pairsList.get(i).loadDataCompleted || this.reloadPairs) {
        console.log('pair id=', this.pairsData.pairsList.get(i).id)
        this.pairsData.pairsList.get(i).gettingData = true
        this.pairsData.getPairPublicData(this.pairsData.pairsList.get(i))
      }
    }
    if (this.reloadPairs) this.reloadPairs = false
  }

  @Watch('pairsListCount')
  watchPairsListChange(newVal: number, oldVal: number) {
    console.log('Change of PairsList:', newVal, ',old value:', oldVal)
    this.loadPairPublicData()
  }

  @Watch('$store.state.currentAccount')
  watchCurrentAccount(newVal: string, oldVal: string) {
    console.log('currentAccount:', newVal, ';previousAccount:', oldVal)
    if (newVal !== '' && this.$store.state.validNetwork) {
      console.log('currentAccount changed,pairlist size:', this.pairsData.pairsList.size())
      if (this.pairsData.pairsList.size() === 0 && !this.pairsData.pairsBaseInfoLoading) {
        this.pairsData.loadPairsInfo()
      } else {
        this.reloadPairs = true
        this.loadPairPublicData()
      }
    }
  }

  @Watch('$store.state.isConnected')
  watchConnectionStatus() {
    console.log(
      'Dashboard $store.state.isConnected:',
      this.$store.state.isConnected,
      'this.$store.state.validNetwork=',
      this.$store.state.validNetwork
    )
    if (!this.$store.state.validNetwork || !this.$store.state.isConnected) {
      this.clearMyData()
    }
    // if (this.$store.state.validNetwork && this.$store.state.isConnected && !this.pairsData.pairsBaseInfoLoading) {
    //   this.pairsData.loadPairsInfo()
    // } else {
    //   this.clearMyData()
    // }
  }

  @Watch('$store.state.chainId')
  watchNetworkChainId(newVal: number, oldVal: number) {
    console.log('chainid newVal=', newVal, 'oldVal=', oldVal)
    if (oldVal > 0 && this.$store.state.validNetwork) {
      console.log('network chainId changed,pairlist size:', this.pairsData.pairsList.size())
      if (this.pairsData.pairsList.size() === 0 && !this.pairsData.pairsBaseInfoLoading) {
        this.pairsData.loadPairsInfo()
      }
    } else if (!this.$store.state.validNetwork) {
      this.clearMyData()
    }
  }

  clearMyData() {
    // this.pairsData.pairsList.clear()
    // sessionStorage.removeItem('pairBaseInfo')
    console.log('clearMyData:pairsList.size=', this.pairsData.pairsList.size())
    for (let i = 0; i < this.pairsData.pairsList.size(); i++) {
      this.pairsData.pairsList.get(i).gettingData = false
      this.pairsData.pairsList.get(i).currentDeposits = 0
      this.pairsData.pairsList.get(i).netAPY = 0
    }
    // this.pairsLoadComplete = false
    // this.pairsListComplete = false
    // this.myValueLocked = 0.00
    // this.netAPYTotal = 0.00
    // this.userFee0Total = 0
    // this.userFee1Total = 0
    // this.myEarnedValue = 0.00
  }
}
</script>

<style scoped>
.pairs-list {
  position: absolute;
  width: 100%;
  height: 100%;
  padding: 0;
  margin: 0;
  float: left;
  /* background-color: cornflowerblue; */
}
.pairs-container {
  margin: 50px;
  margin: 0 auto;
}
.card-panel-col {
  height: 515px;
}
.box-card {
  width: 300px;
  height: 500px;
  margin: 10px;
  padding-top: 0px;
  margin-top: 20px;
  /* background-color: #121218; */
  border-radius: 16px;
}
.box-card:hover {
  border-style: solid;
  border-color: #03a9f4;
  box-shadow: 0 0 15px #03a9f4;
}
.el-card /deep/ .el-card__header {
  background-color: #102130 !important;
  color: #ffffff;
  height: 106px;
}
.el-card /deep/ .el-card__body {
  background-color: #121218 !important;
  padding: 0px;
  color: #ffffff;
}
.box-card-currency {
  width: 900px;
  margin: 10px;
}
.el-form-item__label {
  line-height: 20px !important;
}
html,
body {
  margin: 0;
  padding: 0;
  height: 100%;
}
#app,
.box {
  height: 100%;
}
.el-container {
  height: 100%;
}
.el-input-number.is-without-controls .el-input__inner {
  width: 200px;
  line-height: 30px;
  height: 28px;
}
.block {
  float: left;
}
.card-container {
  width: 100%;
  height: 600px;
  margin: 0px;
  line-height: 22px;
  background-color: #1c1c22;
}
.card-pairinfo {
  height: 330px;
  width: 100%;
  margin: 0;
  padding: 32px 20px 10px 20px;
  background-color: #121218;
  border-radius: 16px;
}
.card-your-position {
  width: 100%;
  position: absolute;
  margin: 20px;
}
.card-your-position .pos-val {
  margin-left: 135px;
}

.f12 {
  color: #ffffff;
  font-family: sans-serif;
  font-size: 12px;
}
.f24 {
  color: #ffffff;
  font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
  font-size: 24px;
  text-transform: uppercase;
}
.token-icon {
  position: absolute;
  margin-top: -30px;
  margin-left: 26px;
}
.pair-name {
  font-size: 20px;
  font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
  padding-bottom: 20px;
}
.pair-intro {
  color: #ffffff;
  font-family: Inter, sans-serif;
  font-size: 14px;
  padding-bottom: 10px;
}
.pair-apy {
  color: #ffffff;
  font-family: Inter, sans-serif;
  padding-bottom: 10px;
}
.pair-current-deposits,
.pair-max-capacity {
  float: left;
  color: #ffffff;
  font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
  font-size: 12px;
}
.pair-deposits-val,
.pair-capacity-val {
  float: right;
  color: #ffffff;
  font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
  font-size: 12px;
  font-weight: bold;
}
.filter-item {
  padding-right: 10px;
}
.phone {
  display: inline;
  padding-top: 10px;
}
.svg-icon {
  vertical-align: -8px !important;
}
.front {
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #673ab7;
  color: white;
  width: 100%;
  height: 100%;
}

.back {
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #ffc107;
  color: white;
  width: 100%;
  height: 100%;
}

@media only screen and (max-width: 767px) {
  .pairs-container {
    width: 320px;
  }
}

@media only screen and (min-width: 768px) {
  .pairs-container {
    width: 630px;
  }
}

@media only screen and (min-width: 992px) {
  .pairs-container {
    width: 1100px;
  }
}
@media only screen and (min-width: 1200px) {
  .pairs-container {
    width: 1100px;
  }
}
@media only screen and (min-width: 1920px) {
  .pairs-container {
    width: 1100px;
  }
}
</style>

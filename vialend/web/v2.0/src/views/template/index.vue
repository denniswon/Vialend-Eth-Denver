<template>
    <div class="pairs-list" v-loading="pairsData.pairsBaseInfoLoading">
        <!-- <MyPairsList v-bind:pairsData="pairsData"></MyPairsList> -->
         <div id="pairs_container" class="pairs-container">

    <el-row :gutter="20" class="panel-group" justify="center" v-if="pairsList.size() > 0">
      <el-col
        :xs="23"
        :sm="12"
        :md="7"
        :lg="7"
        class="card-panel-col"
        v-for="pair in pairsList._getData()"
        :key="pair.id"
      >
        <vue-flip width="320px" horizontal="true" height="500px" v-model="pair.flipped">
          <template v-slot:front  class="front">
            <el-card
              class="box-card"
              v-loading="pair.gettingData"
              element-loading-text="Loading pair data"
              element-loading-spinner="el-icon-loading"
              element-loading-background="rgba(0, 0, 0, 0.8)"
            >
              <div slot="header" class="clearfix card-title">
                <div class="block">
                  <div class="pair-name text item">
                    <span>{{ pair.token0.symbol }} - {{ pair.token1.symbol }}</span>
                  </div>
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
                <div class="card-pairinfo">
                  <div class="text item">
                      Fee Tier<br><span><input class="btn btn-default btn-sm"
                               type="button"
                               :value="Number(pair.feeTier / 10000) + '%'"></span>
                    </div>
                    <div class="text item">
                      Current APR<br><span>{{Number(pair.currentAPR).toFixed(2)}}%</span>
                    </div>
                    <div class="text item">
                      TVL<br><span>{{Number(pair.tvlTotal0).toFixed(2)}} / {{Number(pair.tvlTotal1).toFixed(2)}}<br>
                        <!-- ${{(Number(pair.tvlTotal0USD) + Number(pair.tvlTotal1USD)).toFixed(2)}} -->
                      </span>
                    </div>
                    <div class="text item">
                      Assets ratio<br><span>{{Number(pair.lendingRatio).toFixed(2)}}% Compound<br>
                        {{Number(pair.uniswapRatio).toFixed(2)}}% Uniswap V3</span>
                    </div>
                </div>
              </div>
            </el-card>
          </template>
          <template v-slot:back class="back">
            <el-card class="box-card" v-loading="pair.gettingData"
                element-loading-text="Loading pair data"
                element-loading-spinner="el-icon-loading"
                element-loading-background="rgba(0, 0, 0, 0.8)">
              <div slot="header" class="clearfix card-title">
                <span style="font-size:20px;">{{ pair.token0.symbol }} / {{ pair.token1.symbol }}</span>
                <span
                  style="float: right; padding: 3px 0"
                  type="text"
                  @click="pair.flipped = !pair.flipped"
                >
                  <i class="el-icon-s-data"></i>
                </span>
              </div>

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
  name: 'Templates',
  components: { 'vue-flip': VueFlip }
})
export default class extends Vue {
    pairsData = new PairsData();
    pairsList = this.pairsData.pairsList;
    maxCapacity = 60

    get pairsListCount() {
      return this.pairsData.pairsList.size()
    }

    created() {
      this.pairsData.calculateAPR = true
      console.log('bridgeAddress value123=', this.pairsData.bridgeAddress)
      console.log('pairsList.size=', this.pairsData.pairsList.size())
      if (this.$store.state.validNetwork && this.$store.state.isConnected && this.pairsData.pairsList.size() === 0) {
        this.pairsData.loadPairsInfo()
      }
    }

    @Watch('pairsListCount')
    watchPairsListChange(newVal: number, oldVal: number) {
      console.log('Change of PairsList:', newVal, ',old value:', oldVal)
      for (let i = 0; i < newVal; i++) {
        console.log('pair id=', this.pairsData.pairsList.get(i).id)
        if (!this.pairsData.pairsList.get(i).loadDataCompleted) {
          this.pairsData.pairsList.get(i).gettingData = true
          this.pairsData.getPairPublicData(this.pairsData.pairsList.get(i))
        }
      }
    }

    @Watch('$store.state.currentAccount')
    watchCurrentAccount(newVal:string, oldVal:string) {
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

    @Watch('$store.state.isConnected')
    watchConnectionStatus(newVal:boolean) {
      console.log('Dashboard $store.state.isConnected:', this.$store.state.isConnected, 'this.$store.state.validNetwork=', this.$store.state.validNetwork)
      // if (newVal && this.$store.state.validNetwork && this.pairsData.pairsLoadComplete && !this.myPairsListLoading) {
      // // this.isConnected = this.$store.state.isConnected
      //   // this.pairsListComplete = false
      //   // this.loadMyPairsList()
      //   console.log('isConnected changed,load loadMyPairsList')
      // } else {
      //   this.clearMyData()
      // }
    }

    @Watch('$store.state.chainId')
    watchNetworkChainId(newVal:number, oldVal:number) {
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
      this.pairsData.pairsList.clear()
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
.pairs-list{
    position:absolute;
    width:100%;
    height: 100%;
    padding:0;
    margin:0;
    float:left;
    /* background-color: cornflowerblue; */
}
.pairs-container{
    margin:50px;
    margin:0 auto;
}
.card-panel-col{
    height:515px;
}
.box-card {
  width: 300px;
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
  height: 90px;
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
  margin: 0px;
  line-height: 22px;
  background-color: #1c1c22;
}
.card-pairinfo {
  height: 300px;
  width: 100%;
  margin: 0;
  padding: 38px 20px 10px 20px;
  background-color: #121218;
  border-radius: 16px;
  line-height: 26px;
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
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
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
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
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
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  font-size: 12px;
}
.pair-deposits-val,
.pair-capacity-val {
  float: right;
  color: #ffffff;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
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
    background-color: #673AB7;
    color: white;
    width: 100%;
    height: 100%;
  }

  .back {
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: #FFC107;
    color: white;
    width: 100%;
    height: 100%;
  }

    @media only screen and (max-width: 767px) {
        .pairs-container{
            width:320px;
        }
    }

    @media only screen and (min-width: 768px){
        .pairs-container{
            width:630px;
        }
    }

    @media only screen and (min-width: 992px){
        .pairs-container{
            width:1100px;
        }
    }
    @media only screen and (min-width: 1200px){
        .pairs-container{
            width:1100px;
        }
    }
    @media only screen and (min-width: 1920px){
        .pairs-container{
            width:1100px;
        }
    }
</style>

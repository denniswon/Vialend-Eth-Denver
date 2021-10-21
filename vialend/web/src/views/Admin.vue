<template>
  <div>
    <Header />
    <el-tabs type="border-card">
      <el-tab-pane>
        <span slot="label"><i class="el-icon-date"></i> Rebalance</span>
        <div style="width:65%"
             v-loading="loadingTokensInfoStatus">
          <table class="table table-striped">
            <thead>
              <tr>
                <th>Token balance</th>
                <th>Wallet</th>
                <th>Vault</th>
                <th>Uni-V3</th>
                <th>Compound</th>
                <!-- <th>fees</th> -->
              </tr>
            </thead>
            <tbody>
              <tr>
                <th>{{token0Symbol}}</th>
                <td>{{Number(token0BalanceInWallet).toFixed(2) + ' / $' + token0BalanceUSDInWallet.toFixed(2)}}</td>
                <td>{{Number(token0BalanceInVault).toFixed(2) + ' / $' + token0BalanceUSDInVault.toFixed(2)}}</td>
                <td>{{Number(token0BalanceInPool).toFixed(2) + ' / $' + token0BalanceUSDInPool.toFixed(2)}}</td>
                <td>{{Number(token0BalanceInLending).toFixed(2) + ' / $' + token0BalanceUSDInLending.toFixed(2)}}</td>
                <!-- <td>{{accruedProtocolFees0}}</td> -->
              </tr>
              <tr>
                <th>{{token1Symbol}}</th>
                <td>{{Number(token1BalanceInWallet).toFixed(2) + ' / $' + token1BalanceUSDInWallet.toFixed(2)}}</td>
                <td>{{Number(token1BalanceInVault).toFixed(2) + ' / $' + token1BalanceUSDInVault.toFixed(2)}}</td>
                <td>{{Number(token1BalanceInPool).toFixed(2) + ' / $' + token1BalanceUSDInPool.toFixed(2)}}</td>
                <td>{{Number(token1BalanceInLending).toFixed(2) + ' / $' + token1BalanceUSDInLending.toFixed(2)}}</td>
                <!-- <td>{{accruedProtocolFees1}}</td> -->
              </tr>
              <tr>
                <th>USD</th>
                <td>${{(Number(token0BalanceUSDInWallet) + Number(token1BalanceUSDInWallet)).toFixed(2)}}</td>
                <td>${{(Number(token0BalanceUSDInVault) + Number(token1BalanceUSDInVault)).toFixed(2)}}</td>
                <td>${{(Number(token0BalanceUSDInPool) + Number(token1BalanceUSDInPool)).toFixed(2)}}</td>
                <td>${{(Number(token0BalanceUSDInLending) + Number(token1BalanceUSDInLending)).toFixed(2)}}</td>
                <!-- <td>{{accruedProtocolFees1}}</td> -->
              </tr>
            </tbody>
          </table>
        </div>
        <hr>
        Min Tick:{{tickLower}}<br>
        Max Tick:{{tickUpper}}<br>
        Current Tick:{{currentTick}}<br>
        <span style="color:#ff0000;margin-right:10px;">Current Price:{{currentPrice}}</span>
        <el-button size="mini"
                   :style="{display:rangeStatusDisplay}">
          <span :class="[dotStyle,rangeStatusStyle]"></span>
          <span class="status">{{rangeStatus}}</span>
        </el-button><br>
        <div class="block"
             style="width:65%;margin-top:20px;">
          <!-- <el-slider v-model="tickRange"
                     range
                     :marks="getMarks">
          </el-slider> -->
          <input type="text"
                 class="js-range-slider"
                 name="my_range"
                 value="" />
        </div>

        <hr>
        <div class="range_title">Set Price Range</div>
        <el-form :inline="true"
                 :model="rangeForm"
                 class="demo-form-inline">
          <el-form-item label="Min Price">
            <el-input-number v-model="rangeForm.minPrice"
                             @change="minPriceChange"
                             :precision="1"
                             :step="0.1"></el-input-number><br>
            {{token1Symbol}} per {{token0Symbol}}
          </el-form-item>
          <el-form-item label="Max Price">
            <el-input-number v-model="rangeForm.maxPrice"
                             @change="maxPriceChange"
                             :precision="1"
                             :step="0.1"></el-input-number><br>
            {{token1Symbol}} per {{token0Symbol}}
          </el-form-item>
          <el-form-item>
            <el-button type="primary"
                       :loading="rebalanceLoading"
                       @click="doRebalance">Rebalance</el-button>
            <!-- <el-button type="primary"
                       @click="getY">getY</el-button> -->
            &nbsp;&nbsp;
            <a :href="goToEtherscan"
               target="_blank"
               :class="['btn','btn-primary',viewOnEtherscanDisable?'disabled':'']"
               role="button">View on etherscan</a>
          </el-form-item>
        </el-form>
        {{errorRebalance}}
      </el-tab-pane>
      <el-tab-pane>
        <span slot="label"><i class="el-icon-date"></i> Security</span>
        remove all liquidity from uniswap
      </el-tab-pane>
      <el-tab-pane>
        <span slot="label"><i class="el-icon-date"></i> Settings</span>
        <div class="token_exchange_form">
          <div style="margin-bottom:20px;">Currency Converter</div>
          <el-form :inline="true"
                   class="demo-form-inline">
            <el-form-item>
              <el-input v-model="currTokenVal"
                        placeholder="1"></el-input>
            </el-form-item>
            <el-form-item>
              <el-select v-model="currTokenId"
                         placeholder="请选择"
                         @change="exchangeTokenChange">
                <el-option v-for="item in tokensList"
                           :key="item.nameid"
                           :label="item.symbol"
                           :value="item.symbol">
                </el-option>
              </el-select>
            </el-form-item>
            =
            <el-form-item>
              <el-input v-model="usdTokenVal"></el-input>
            </el-form-item>USD
          </el-form>
        </div>
        <hr>
        <el-form :inline="true"
                 class="demo-form-inline">
          <el-form-item>UniPortionRatio:</el-form-item>
          <el-form-item>
            <el-input v-model="uniPortionRatio"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary"
                       @click="SetUniPortionRatio">SetUniPortionRatio</el-button>
          </el-form-item>
          <br>
          <el-form-item>MaxTotalSupply:</el-form-item>
          <el-form-item>
            <el-input v-model="maxTotalSupply"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary"
                       @click="SetMaxTotalSupply">SetMaxTotalSupply</el-button>
          </el-form-item>
          <br>
          <el-form-item>Governance:</el-form-item>
          <el-form-item>
            <el-input v-model="governanceAddress"
                      style="width:500px;"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary"
                       @click="SetGovernance">SetGovernance</el-button>
            <el-button type="primary"
                       @click="acceptGovernance">AccecptGovernance</el-button>
          </el-form-item>
          <br>
          <el-form-item>Team:</el-form-item>
          <el-form-item>
            <el-input v-model="teamAddress"
                      style="width:500px;"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary"
                       @click="SetTeam">SetTeam</el-button>
          </el-form-item>
          <br>
          <el-form-item>
            <el-button type="warning"
                       @click="EmergencyBurn">EmergencyBurn</el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>
    </el-tabs>

  </div>
</template>

<script>
import Web3 from 'web3'
import Header from '@/components/Header.vue'
import uniswapV3PoolABI from '../ABI/UniswapV3PoolABI.json'
import ViaLendTokenABI from '../ABI/tokenABI.json'
import axios from 'axios'

if (typeof web3 !== 'undefined') {
  web3 = new Web3(web3.currentProvider)
  console.log('web3 provider:web3.currentProvider')
} else {
  // set the provider you want from Web3.providers
  web3 = new Web3(new Web3.providers.HttpProvider('https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6'))
  console.log('web3 provider:goerli')
}

export default {
  components: { Header },
  data () {
    return {
      isConnected: this.$store.state.isConnected,
      keeperUniswapV3Contract: null,
      poolAddress: '',
      tickLower: 0,
      tickUpper: 0,
      currentTick: 0,
      currentPrice: 0,
      tickRange: [20, 80],
      currentPercent: 0,
      marks: {
      },
      rangeForm: {
        minPrice: 0.0,
        maxPrice: 0.0
      },
      errorRebalance: '',
      rebalanceLoading: false,
      loadingTokensInfoStatus: false,
      token0Contract: null,
      token1Contract: null,
      token0LendingContract: null,
      token1LendingContract: null,
      token0Address: '',
      token0Name: '',
      token0Symbol: '',
      token0Decimal: 0,
      token0Balance: 0,
      token1Address: '',
      token1Name: '',
      token1Symbol: '',
      token1Decimal: 0,
      token0BalanceInWallet: 0.00,
      token1BalanceInWallet: 0.00,
      token0BalanceUSDInWallet: 0.00,
      token1BalanceUSDInWallet: 0.00,
      token0BalanceInVault: 0.00,
      token1BalanceInVault: 0.00,
      token0BalanceUSDInVault: 0.00,
      token1BalanceUSDInVault: 0.00,
      token0BalanceInPool: 0,
      token1BalanceInPool: 0,
      token0BalanceUSDInPool: 0,
      token1BalanceUSDInPool: 0,
      token0BalanceInLending: 0,
      token1BalanceInLending: 0,
      token0BalanceUSDInLending: 0,
      token1BalanceUSDInLending: 0,
      dotStyle: 'dot',
      rangeStatusStyle: '',
      rangeStatusDisplay: 'none',
      rangeStatus: '',
      tokensList: null,
      currTokenId: 'BTC',
      currTokenVal: 1,
      usdTokenVal: 0,
      priceUSD: 0,
      maxTotalSupply: 0,
      uniPortionRatio: 0,
      governanceAddress: '',
      teamAddress: '',
      viewOnEtherscanDisable: true,
      transactionHash: '',
      goToEtherscan: '',
      accruedProtocolFees0: 0,
      accruedProtocolFees1: 0
    }
  },
  created: function () {
    this.initData()
  },
  mounted () {
  },
  computed: {
    newMinPrice () {
      return this.rangeForm.minPrice
    },
    newMaxPrice () {
      return this.rangeForm.maxPrice
    },
    getMarks () {
      var marks = {}
      if (this.currentPercent !== 0) {
        marks[this.currentPercent] = this.currentTick
      }
      return marks
    }
  },
  watch: {
    newMinPrice (price) {
      console.log('new min price=', price)
      if (!isNaN(price)) { this.tickLower = this.priceToTick(price) } else { this.tickLower = 0 }
      this.drawTickRangeChart()
    },
    newMaxPrice (price) {
      console.log('new max price=', price)
      this.tickUpper = this.priceToTick(price)
      if (!isNaN(price)) { this.tickUpper = this.priceToTick(price) } else { this.tickUpper = 0 }
      this.drawTickRangeChart()
    },
    '$store.state.isConnected': function () {
      this.isConnected = this.$store.state.isConnected
      if (this.isConnected) {
        this.getTokensBalanceInWallet()
      } else {
        this.token0BalanceInWallet = 0.00
        this.token1BalanceInWallet = 0.00
        this.token0BalanceUSDInWallet = 0.00
        this.token1BalanceUSDInWallet = 0.00
      }
    },
    '$store.state.currentAccount': function () {
      this.getTokensBalanceInWallet()
    },
    '$store.state.chainId': function () {
      this.initData()
    },
    currTokenVal (val) {
      if (val !== '') {
        if (!isNaN(val)) {
          console.log('currTokenVal=', val)
          this.usdTokenVal = val * this.priceUSD
        }
      }
    },
    currentTick (val) {
      this.drawTickRangeChart()
    }
  },
  methods: {
    async initData () {
      if (await this.$parent.checkCurrentChain() === false) {
        this.token0BalanceInWallet = 0.00
        this.token1BalanceInWallet = 0.00
        this.token0BalanceUSDInWallet = 0.00
        this.token1BalanceUSDInWallet = 0.00
        this.token0BalanceInVault = 0.00
        this.token1BalanceInVault = 0.00
        this.token0BalanceUSDInVault = 0.00
        this.token1BalanceUSDInVault = 0.00
        this.token0BalanceInPool = 0.00
        this.token1BalanceInPool = 0.00
        this.token0BalanceUSDInPool = 0.00
        this.token1BalanceUSDInPool = 0.00
        this.token0BalanceInLending = 0.00
        this.token1BalanceInLending = 0.00
        this.token0BalanceUSDInLending = 0.00
        this.token1BalanceUSDInLending = 0.00
        return
      }
      this.getTokensList()
      this.poolAddress = await this.$parent.keeperContract.methods.poolAddress().call()
      this.token0Address = await this.$parent.keeperContract.methods.token0().call()
      this.token1Address = await this.$parent.keeperContract.methods.token1().call()
      console.log('poolAddress=', this.poolAddress)
      console.log('token0Address=', this.token0Address)
      console.log('token1Address=', this.token1Address)
      // var lendingAmounts = await this.$parent.keeperContract.methods.getLendingAmounts().call()
      this.token0LendingContract = new web3.eth.Contract(
        ViaLendTokenABI,
        '0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF'
      )
      this.token1LendingContract = new web3.eth.Contract(
        ViaLendTokenABI,
        '0xCEC4a43eBB02f9B80916F1c718338169d6d5C1F0'
      )
      this.token0Contract = new web3.eth.Contract(
        ViaLendTokenABI,
        this.token0Address
      )
      this.token1Contract = new web3.eth.Contract(
        ViaLendTokenABI,
        this.token1Address
      )
      this.keeperUniswapV3Contract = new web3.eth.Contract(
        uniswapV3PoolABI,
        this.poolAddress
      )
      await this.getSlot0()
      this.getTokensInfo()
      this.loadPriceRange()
      this.getSettingData()
      // this.loadFees()
    },
    async getTokensList () {
      axios.get('https://api.coinlore.net/api/tickers/').then((response) => {
        this.tokensList = response.data.data
        for (var i = 0; i < this.tokensList.length; i++) {
          if (this.tokensList[i].symbol === this.currTokenId) {
            this.usdTokenVal = this.tokensList[i].price_usd
            this.priceUSD = this.tokensList[i].price_usd
            break
          }
        }
        console.log('tokensList length=', this.tokensList.length)
      })
    },
    async getTokensInfo () {
      this.loadingTokensInfoStatus = true
      // Get Token0 Information
      this.token0Name = await this.token0Contract.methods.name().call()
      console.log('token0Name=', this.token0Name)
      this.token0Symbol = await this.token0Contract.methods.symbol().call()
      console.log('token0Symbol=', this.token0Symbol)
      this.token0Decimal = await this.token0Contract.methods.decimals().call()
      console.log('token0Decimal=', this.token0Decimal)
      // Get Token1 Information
      this.token1Name = await this.token1Contract.methods.name().call()
      console.log('token1Name=', this.token1Name)
      this.token1Symbol = await this.token1Contract.methods.symbol().call()
      console.log('token1Symbol=', this.token1Symbol)
      this.token1Decimal = await this.token1Contract.methods.decimals().call()
      console.log('token1Decimal=', this.token1Decimal)
      // ----------------------Tokens balance ------------------------------
      // Token0 balance in wallet and vault
      // console.log('ethereum.selectedAddress=', ethereum.selectedAddress)
      this.getTokensBalanceInWallet()
      // Token0 and token1 balance in pool
      var tmpTickLower = await this.$parent.keeperContract.methods.cLow().call()
      var tmpTickUpper = await this.$parent.keeperContract.methods.cHigh().call()
      var result = await this.$parent.keeperContract.methods.getPositionAmounts(BigInt(tmpTickLower), BigInt(tmpTickUpper)).call()
      if (result !== undefined && result !== null) {
        this.token0BalanceInPool = result.amount0 / Number(Math.pow(10, this.token0Decimal))
        this.token0BalanceUSDInPool = Number(this.token0BalanceInPool) * Number(this.$store.state.token0RateOfUSD)
        this.token1BalanceInPool = result.amount1 / Number(Math.pow(10, this.token1Decimal))
        this.token1BalanceUSDInPool = Number(this.token1BalanceInPool) * Number(this.$store.state.token1RateOfUSD)
      }
      // Token0 and token1 balance in Compound
      var exchangeRate0 = await this.token0LendingContract.methods.exchangeRateStored().call()
      var exchangeRate1 = await this.token1LendingContract.methods.exchangeRateStored().call()
      var CAmount0 = await this.token0LendingContract.methods.balanceOf(this.$parent.vaultAddress).call()
      var CAmount1 = await this.token1LendingContract.methods.balanceOf(this.$parent.vaultAddress).call()
      var underlying0 = CAmount0 * exchangeRate0 / Math.pow(10, 18)
      var underlying1 = CAmount1 * exchangeRate1 / Math.pow(10, 18)
      console.log('underlying0=', underlying0, 'underlying1=', underlying1)
      if (!isNaN(underlying0) && !isNaN(underlying1)) {
        this.token0BalanceInLending = underlying0 / Number(Math.pow(10, this.token0Decimal))
        this.token0BalanceUSDInLending = Number(this.token0BalanceInLending) * Number(this.$store.state.token0RateOfUSD)
        this.token1BalanceInLending = underlying1 / Number(Math.pow(10, this.token1Decimal))
        this.token1BalanceUSDInLending = Number(this.token1BalanceInLending) * Number(this.$store.state.token1RateOfUSD)
      }
      this.loadingTokensInfoStatus = false
    },
    async getTokensBalanceInWallet () {
      if (!this.isConnected) {
        // this.$message('Please connect wallet!')
      } else {
        var token0BalanceWeiInWallet = await this.token0Contract.methods.balanceOf(ethereum.selectedAddress).call()
        this.token0BalanceInWallet = (Number(token0BalanceWeiInWallet) / Number(Math.pow(10, this.token0Decimal))).toFixed(2)
        this.token0BalanceUSDInWallet = Number(this.token0BalanceInWallet) * Number(this.$store.state.token0RateOfUSD)
        var token0BalanceWeiInVault = await this.$parent.keeperContract.methods.getBalance0().call()
        this.token0BalanceInVault = (Number(token0BalanceWeiInVault) / Number(Math.pow(10, this.token0Decimal))).toFixed(2)
        this.token0BalanceUSDInVault = Number(this.token0BalanceInVault) * Number(this.$store.state.token0RateOfUSD)
        // Token1 balance in wallet and vault
        var token1BalanceWeiInWallet = await this.token1Contract.methods.balanceOf(ethereum.selectedAddress).call()
        this.token1BalanceInWallet = (Number(token1BalanceWeiInWallet) / Number(Math.pow(10, this.token1Decimal))).toFixed(2)
        this.token1BalanceUSDInWallet = Number(this.token1BalanceInWallet) * Number(this.$store.state.token1RateOfUSD)
        var token1BalanceWeiInVault = await this.$parent.keeperContract.methods.getBalance1().call()
        this.token1BalanceInVault = (Number(token1BalanceWeiInVault) / Number(Math.pow(10, this.token1Decimal))).toFixed(2)
        this.token1BalanceUSDInVault = Number(this.token1BalanceInVault) * Number(this.$store.state.token1RateOfUSD)
      }
    },
    async loadFees () {
      var accumulateUniswapFees0 = await this.$parent.keeperContract.methods.AccumulateUniswapFees0().call()
      var accumulateUniswapFees1 = await this.$parent.keeperContract.methods.AccumulateUniswapFees1().call()
      this.accruedProtocolFees0 = await this.$parent.keeperContract.methods.accruedProtocolFees0().call()
      this.accruedProtocolFees1 = await this.$parent.keeperContract.methods.accruedProtocolFees1().call()
      console.log('accumulateUniswapFees0=', accumulateUniswapFees0)
      console.log('accumulateUniswapFees1=', accumulateUniswapFees1)
      console.log('accruedProtocolFees0=', this.accruedProtocolFees0)
      console.log('accruedProtocolFees1=', this.accruedProtocolFees1)
    },
    async loadPriceRange () {
      var _this = this
      var tmpTickLower, tmpTickUpper
      tmpTickLower = await this.$parent.keeperContract.methods.cLow().call()
      tmpTickUpper = await this.$parent.keeperContract.methods.cHigh().call()
      if (tmpTickLower !== '0' && tmpTickUpper !== '0') {
        this.tickLower = tmpTickLower
        this.tickUpper = tmpTickUpper
        console.log('tickLower_price=', this.tickLower)
        console.log('tickUpper_price=', this.tickUpper)
      }
      this.rangeForm.minPrice = this.tickToPrice(this.tickLower)
      this.rangeForm.maxPrice = this.tickToPrice(this.tickUpper)

      console.log('maxPrice123=', this.rangeForm.maxPrice)
      var maxPriceVal = parseFloat(this.rangeForm.maxPrice) + 800
      $('.js-range-slider').ionRangeSlider({
        skin: 'big',
        type: 'double',
        grid: true,
        min: (parseFloat(this.rangeForm.minPrice) - 230).toFixed(1),
        max: (parseFloat(this.rangeForm.maxPrice) + 230).toFixed(1),
        from: this.rangeForm.minPrice,
        to: this.rangeForm.maxPrice,
        prefix: '$',
        step: 0.1,
        onChange: function (data) {
          console.log('from=', data.from)
          _this.rangeForm.minPrice = data.from
          _this.rangeForm.maxPrice = data.to
        }
      })
    },
    minPriceChange (val) {
      // console.log('min price:', val)
      $('.js-range-slider').data('ionRangeSlider').update({ from: val })
    },
    maxPriceChange (val) {
      // console.log('max price:', val)
      $('.js-range-slider').data('ionRangeSlider').update({ to: val })
    },
    drawTickRangeChart () {
      if (this.tickLower !== 0 && this.tickUpper !== 0) {
        var leftMargin = this.tickLower - 10000
        var rightMargin = this.tickUpper + 10000
        var leftPercent = parseInt(((this.tickLower - leftMargin) / (rightMargin - leftMargin)) * 100)
        var rightPercent = parseInt(((this.tickUpper - leftMargin) / (rightMargin - leftMargin)) * 100)
        this.currentPercent = parseInt(((this.currentTick - leftMargin) / (rightMargin - leftMargin)) * 100)
        this.tickRange = [leftPercent, rightPercent]
        console.log('currentPercent=', this.currentPercent)
        if (this.tickLower > this.currentTick || this.tickUpper < this.currentTick) {
          this.rangeStatusStyle = 'outOfRange'
          this.rangeStatus = 'Out of range'
        } else {
          this.rangeStatusStyle = 'inRange'
          this.rangeStatus = 'In range'
        }
        this.rangeStatusDisplay = ''
      }
    },
    async getSlot0 () {
      if (this.keeperUniswapV3Contract !== null) {
        var _this = this
        this.keeperUniswapV3Contract.methods
          .slot0()
          .call()
          .then(slot => {
            if (slot !== undefined && slot !== null && slot !== '') {
              console.log('slot0=' + JSON.stringify(slot))
              // this.tickLower = Math.round(parseInt(this.rangeForm.tickLower) / 60) * 60
              // this.tickUpper = Math.round(parseInt(this.rangeForm.tickUpper) / 60) * 60
              _this.currentTick = slot['tick']
              _this.currentPrice = this.tickToPrice(_this.currentTick)
              console.log('currentTick0=', _this.currentTick)
              if (_this.tickLower === 0 || _this.tickUpper === 0) {
                console.log('currentTick1=', _this.currentTick)
                _this.tickLower = parseInt(_this.currentTick) - 1000
                _this.tickUpper = parseInt(_this.currentTick) + 1000
                console.log('tickLower0=', _this.tickLower, 'tickUpper0=', _this.tickUpper)
              } else {
                console.log('tickLower1=', _this.tickLower, 'tickUpper1=', _this.tickUpper)
              }
            }
          })
      }
    },
    async SetUniPortionRatio () {
      var _this = this
      var showMessage = false
      if (this.$parent.keeperContract != null) {
        this.$parent.keeperContract.methods
          .setUniPortionRatio(
            BigInt(this.uniPortionRatio)
          )
          .send({
            from: ethereum.selectedAddress,
            gasPrice: '80000000000',
            gas: 600000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            if (showMessage === false) {
              _this.$message('Set UniPortionRatio Successful!')
              showMessage = true
            }
          })
          .on('receipt', function (receipt) {
            if (showMessage === false) {
              _this.$message('Set UniPortionRatio Successful!')
              showMessage = true
            }
          })
          .on('error', function (error) {
            if (showMessage === false) {
              _this.$message.error(error)
              showMessage = true
              console.log(error)
            }
          })
      }
    },
    async SetMaxTotalSupply () {
      var _this = this
      var showMessage = false
      if (this.$parent.keeperContract != null) {
        this.$parent.keeperContract.methods
          .setMaxTotalSupply(
            BigInt(this.maxTotalSupply)
          )
          .send({
            from: ethereum.selectedAddress,
            gasPrice: '80000000000',
            gas: 600000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            if (showMessage === false) {
              _this.$message('Set MaxTotalSupply Successful!')
              showMessage = true
            }
          })
          .on('receipt', function (receipt) {
            if (showMessage === false) {
              _this.$message('Set MaxTotalSupply Successful!')
              showMessage = true
            }
          })
          .on('error', function (error) {
            if (showMessage === false) {
              _this.$message.error(error)
              showMessage = true
              console.log(error)
            }
          })
      }
    },
    async SetGovernance () {
      var _this = this
      var showMessage = false
      if (this.$parent.keeperContract != null) {
        this.$parent.keeperContract.methods
          .setGovernance(
            this.governanceAddress
          )
          .send({
            from: ethereum.selectedAddress,
            gasPrice: '80000000000',
            gas: 600000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            if (showMessage === false) {
              _this.$message('Set Governance Successful!')
              showMessage = true
            }
          })
          .on('receipt', function (receipt) {
            if (showMessage === false) {
              _this.$message('Set Governance Successful!')
              showMessage = true
            }
          })
          .on('error', function (error) {
            if (showMessage === false) {
              _this.$message.error(error)
              showMessage = true
              console.log(error)
            }
          })
      }
    },
    async acceptGovernance () {
      var _this = this
      var showMessage = false
      if (this.$parent.keeperContract != null) {
        this.$parent.keeperContract.methods
          .acceptGovernance()
          .send({
            from: ethereum.selectedAddress,
            gasPrice: '80000000000',
            gas: 600000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            if (showMessage === false) {
              _this.$message('Accept Governance Successful!')
              showMessage = true
            }
          })
          .on('receipt', function (receipt) {
            if (showMessage === false) {
              _this.$message('Accept Governance Successful!')
              showMessage = true
            }
          })
          .on('error', function (error) {
            if (showMessage === false) {
              _this.$message.error(error)
              showMessage = true
              console.log(error)
            }
          })
      }
    },
    async SetTeam () {
      var _this = this
      var showMessage = false
      if (this.$parent.keeperContract != null) {
        this.$parent.keeperContract.methods
          .setTeam(
            this.teamAddress
          )
          .send({
            from: ethereum.selectedAddress,
            gasPrice: '80000000000',
            gas: 600000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            if (showMessage === false) {
              _this.$message('SetTeam Successful!')
              showMessage = true
            }
          })
          .on('receipt', function (receipt) {
            if (showMessage === false) {
              _this.$message('SetTeam Successful!')
              showMessage = true
            }
          })
          .on('error', function (error) {
            if (showMessage === false) {
              _this.$message.error(error)
              showMessage = true
              console.log(error)
            }
          })
      }
    },
    async EmergencyBurn () {
      var _this = this
      var showMessage = false
      if (this.$parent.keeperContract != null) {
        this.$parent.keeperContract.methods
          .emergencyBurn()
          .send({
            from: ethereum.selectedAddress,
            gasPrice: '80000000000',
            gas: 600000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            if (showMessage === false) {
              _this.$message('Emergency Burn Successful!')
              showMessage = true
            }
          })
          .on('receipt', function (receipt) {
            if (showMessage === false) {
              _this.$message('Emergency Burn Successful!')
              showMessage = true
            }
          })
          .on('error', function (error) {
            if (showMessage === false) {
              _this.$message.error(error)
              showMessage = true
              console.log(error)
            }
          })
      }
    },
    exchangeTokenChange (val) {
      console.log('token change=', val)
      for (var i = 0; i < this.tokensList.length; i++) {
        if (this.tokensList[i].symbol === val) {
          this.usdTokenVal = this.tokensList[i].price_usd
          this.priceUSD = this.tokensList[i].price_usd
          this.currTokenVal = 1
          break
        }
      }
    },
    doRebalance () {
      if (!this.isConnected) {
        this.$message('Please connect wallet!')
        return
      }
      var _this = this
      this.rebalanceLoading = true
      if (isNaN(this.rangeForm.minPrice) || this.rangeForm.minPrice <= 0) {
        this.$message({
          message: 'Please input positive number greater than 0 in Min Price.',
          type: 'warning'
        })
      }
      if (isNaN(this.rangeForm.maxPrice) || this.rangeForm.maxPrice <= 0) {
        this.$message({
          message: 'Please input positive number greater than 0 in Max Price.',
          type: 'warning'
        })
      }
      if (this.$parent.keeperContract != null) {
        console.log('account address is ' + ethereum.selectedAddress)
        console.log('this.rangeForm.tickLower=', Math.round(parseInt(this.tickLower) / 60) * 60)
        console.log('this.rangeForm.tickUpper=', Math.round(parseInt(this.tickUpper) / 60) * 60)

        this.$parent.keeperContract.methods
          .strategy1(
            BigInt(Math.round(parseInt(this.tickLower) / 60) * 60),
            BigInt(Math.round(parseInt(this.tickUpper) / 60) * 60)
          )
          .send({
            from: ethereum.selectedAddress,
            // gasPrice: '1000000000',
            // gas: 900000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            if (_this.rebalanceLoading) {
              _this.rebalanceLoading = false
              _this.$message('rebalance confirmed!')
              console.log('confirmation,confirmationNumber:' + confirmationNumber + ',receipt:' + JSON.stringify(receipt))
              if (receipt != null) {
                _this.goToEtherscan = 'https://goerli.etherscan.io/tx/' + receipt.transactionHash
                _this.viewOnEtherscanDisable = false
              }
              // _this.errorRebalance = receipt
              _this.initData()
            }
          })
          .on('receipt', function (receipt) {
            if (_this.rebalanceLoading) {
              _this.rebalanceLoading = false
              _this.$message('rebalance receipt!')
              console.log('receipt:' + JSON.stringify(receipt))
              // _this.errorRebalance = receipt
              if (receipt != null) {
                _this.goToEtherscan = 'https://goerli.etherscan.io/tx/' + receipt.transactionHash
                _this.viewOnEtherscanDisable = false
              }
              _this.initData()
            }
          })
          .on('error', function (error) {
            if (_this.rebalanceLoading) {
              _this.rebalanceLoading = false
              _this.$message.error(error)
            }
            // _this.errorRebalance = error
          })
      }
    },
    priceToTick (price) {
      return parseInt(Math.log(price / 1000000000000) / Math.log(1.0001))
    },
    tickToPrice (tick) {
      return (Math.pow(1.0001, tick) * Math.pow(10, 12)).toFixed(1)
    },
    onSubmit () {
      console.log('submit!')
    },
    get_liquidity_0 (x, sa, sb) {
      return x * sa * sb / (sb - sa)
    },
    calculate_y (L, sp, sa) {
      return L * (sp - sa)
    },
    calculate_c (p, d, x, y) {
      return y / ((d - 1) * p * x + y)
    },
    calculate_d (p, c, x, y) {
      return 1 + y * (1 - c) / (c * p * x)
    },
    getY () {
      var p = 3879.97
      var a = 300
      var b = 5000
      var x = 1
      var sp = Math.sqrt(p)
      var sa = Math.sqrt(a)
      var sb = Math.sqrt(b)
      console.log('Math.pow(p,0.5)=', Math.pow(p, 0.5), ';sp=', sp)

      var L = this.get_liquidity_0(x, sp, sb)
      var y = this.calculate_y(L, sp, sa)
      console.log('amount of USDC y=', y)
      var c = sb / sp
      var d = sa / sp
      var ic = (p, d, x, y)
      console.log('ic=,', ic)

      var id = this.calculate_d(p, c, x, y)
      console.log('id=', id)
      var C = Math.pow(ic, 2)
      console.log('C=', C)
      var D = Math.pow(id, 2)
      console.log('D=', D)
      console.log('p_a=', D * p, ',(', D * 100, ' of P),p_b=', C * p, ',(', C * 100, ' of P')
    },
    async getSettingData () {
      this.maxTotalSupply = await this.$parent.keeperContract.methods.maxTotalSupply().call()
      this.governanceAddress = await this.$parent.keeperContract.methods.governance().call()
      // this.teamAddress = await this.$parent.keeperContract.methods.team().call()
    }

  }
}
</script>

<style scoped>
.range_title {
  font-size: 20px;
  color: black;
  padding: 10px;
}
.token_exchange_form {
  width: 100%;
  font-size: 20px;
  color: black;
  padding: 10px;
}
.dot {
  position: absolute;
  width: 10px;
  height: 10px;
  border-radius: 100%;
  box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.2);
}
.status {
  margin-left: 15px;
}
.inRange {
  background: #0d7404;
}
.outOfRange {
  background: brown;
}
</style>

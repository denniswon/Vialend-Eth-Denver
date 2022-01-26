<template>
  <div id="pairs_container" class="pairs-container" v-loading="pairsData.pairsBaseInfoLoading">
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
    <el-tabs type="border-card">
      <el-tab-pane>
        <span slot="label">
          <i class="el-icon-menu"></i> Rebalance
        </span>
        <div style="width:100%">
          <table class="table table-striped">
            <thead>
              <tr>
                <th>Token balance</th>
                <th>Wallet</th>
                <th>Vault</th>
                <th>Uni-V3</th>
                <th>Compound</th>
              </tr>
            </thead>
            <tbody
              v-if="pairSelectedIndex !== '' && Number(pairSelectedIndex) >= 0"
              v-loading="pairsData.pairsBalanceLoading"
            >
              <tr>
                <th>{{currentPair.token0.symbol}}</th>
                <td>{{Number(currentPair.token0.balanceInWallet).toFixed(2)}}</td>
                <!-- + ' / $' + token0BalanceUSDInWallet.toFixed(2) -->
                <td>{{Math.floor(Number(currentPair.token0.balanceInVault))}}</td>
                <!-- + ' / $' + token0BalanceUSDInVault.toFixed(2) -->
                <td>{{Math.floor(Number(currentPair.token0.balanceInPool))}}</td>
                <!-- + ' / $' + token0BalanceUSDInPool.toFixed(2) -->
                <td>{{Math.floor(Number(currentPair.token0.balanceInLending))}}</td>
                <!-- + ' / $' + token0BalanceUSDInLending.toFixed(2) -->
              </tr>
              <tr>
                <th>{{currentPair.token1.symbol}}</th>
                <td>{{Number(currentPair.token1.balanceInWallet).toFixed(2)}}</td>
                <!-- + ' / $' + token1BalanceUSDInWallet.toFixed(2) -->
                <td>{{Math.floor(Number(currentPair.token1.balanceInVault))}}</td>
                <!-- + ' / $' + token1BalanceUSDInVault.toFixed(2) -->
                <td>{{Math.floor(Number(currentPair.token1.balanceInPool))}}</td>
                <!-- + ' / $' + token1BalanceUSDInPool.toFixed(2) -->
                <td>{{Math.floor(Number(currentPair.token1.balanceInLending))}}</td>
                <!-- + ' / $' + token1BalanceUSDInLending.toFixed(2) -->
                <!-- <td>{{accruedProtocolFees1}}</td> -->
              </tr>
            </tbody>
          </table>

          <el-card class="box-card">
            <div slot="header" class="clearfix">
              <div class="range-title">
                <div class="range-text">Set Price Range</div>
                <div class="range-status">
                <el-button size="mini" :style="{display: rangeStatusDisplay}">
                  <span :class="['dot',rangeStatusStyle]"></span>
                  <span class="status">{{rangeStatus}}</span>
                </el-button>
              </div>
              </div>
              <div class="tick-info">
                <ul>
                  <li>
                    <el-tag type="info">Min Tick:{{currentPair.tickLower}}</el-tag>
                  </li>
                  <li>
                    <el-tag type="info">Max Tick:{{currentPair.tickUpper}}</el-tag>
                  </li>
                  <li>
                    <el-tag type="info">Current Tick:{{currentPair.currentTick}}</el-tag>
                  </li>
                  <li>
                    <el-tag type="danger">Current Price:{{currentPair.currentPrice}}</el-tag>
                  </li>
                </ul>
              </div>
            </div>
            <div class="text item">

              <div class="clearfix"></div>
              <div class="block ranger-slider">
                <ion-range-slider ref="priceRangeSlider" v-model="dataRange"></ion-range-slider>
              </div>
              <div class="clearfix"></div>
              <div class="price-setting">
                <el-form :inline="true" :model="rangeForm" class="demo-form-inline">
                  <el-form-item label="Min Price">
                    <el-input-number
                      v-model="rangeForm.minPrice"
                      size="medium"
                      @change="minPriceChange"
                      :precision="precision"
                      :step="step"
                    ></el-input-number>
                    <br />
                    {{currentPair.token1.symbol}} per {{currentPair.token0.symbol}}
                  </el-form-item>
                  <el-form-item label="Max Price">
                    <el-input-number
                      v-model="rangeForm.maxPrice"
                      size="medium"
                      @change="maxPriceChange"
                      :precision="precision"
                      :step="step"
                    ></el-input-number>
                    <br />
                    {{currentPair.token1.symbol}} per {{currentPair.token0.symbol}}
                  </el-form-item>
                <br>
                <div class="rebalance_action">
                    <el-button-group>
                      <el-button :loading="rebalanceLoading" type="primary" :disabled="currentPair.vaultAddress !== '' ? false : true" @click="doRebalance">Rebalance</el-button>
                      <el-button type="primary" :disabled="doRebalanceEtherscanDisable" class="godetail">
                        <a v-if="rebalanceTransHashLink !== ''" :href="rebalanceTransHashLink" target="_blank">
                          <svg-icon name="goto-detail" width="12" height="12" style="transform:scale(2);" />
                        </a>
                        <svg-icon v-else name="goto-detail" width="12" height="12" style="transform:scale(2);" />
                      </el-button>
                    </el-button-group>
                    &nbsp;&nbsp;
                    <el-button-group>
                      <el-button :loading="removePositionsLoading" type="primary" :disabled="currentPair.vaultAddress !== '' ? false : true" @click="doRemovePositions">RemovePositions</el-button>
                      <el-button type="primary" :disabled="removePositionsEtherscanDisable" class="godetail">
                        <a v-if="removePositionTransHashLink !== ''" :href="removePositionTransHashLink" target="_blank">
                          <svg-icon name="goto-detail" width="12" height="12" style="transform:scale(2);" />
                        </a>
                        <svg-icon v-else name="goto-detail" width="12" height="12" style="transform:scale(2);" />
                      </el-button>
                    </el-button-group>
                  </div>
                </el-form>
                {{errorRebalance}}
              </div>
            </div>
          </el-card>
        </div>
      </el-tab-pane>
      <el-tab-pane>
        <span slot="label">
          <svg-icon name="security" /> Security
        </span>
        <div width="100%">
          <el-row :gutter="20" class="panel-group" justify="center">
            <el-col :xs="23" :sm="12" :md="12" :lg="12">
              <div class="security-col">
                <span class="sec-title">Pool Address:</span>&nbsp;&nbsp;
                <span class="sec-val">{{currentPair.poolAddress}}</span>
                <a
                  :href="'https://goerli.etherscan.io/address/'+currentPair.poolAddress"
                  style="cursor:hand"
                  target="_black"
                >
                  <svg-icon name="goto-detail" width="30" height="30" />
                </a>
              </div>
            </el-col>
            <el-col :xs="23" :sm="12" :md="12" :lg="12">
              <div class="security-col">
                <span class="sec-title">Vault Address:</span>&nbsp;&nbsp;
                <span class="sec-val">{{currentPair.vaultAddress}}</span>
                <a
                  :href="'https://goerli.etherscan.io/address/'+currentPair.vaultAddress"
                  style="cursor:hand"
                  target="_black"
                >
                  <svg-icon name="goto-detail" width="30" height="30" />
                </a>
              </div>
            </el-col>
            <el-col :xs="23" :sm="12" :md="12" :lg="12">
              <div class="security-col">
                <span class="sec-title">Factory Address:</span>&nbsp;&nbsp;
                <span class="sec-val">{{factoryAddress}}</span>
                <a
                  :href="'https://goerli.etherscan.io/address/'+factoryAddress"
                  style="cursor:hand"
                  target="_black"
                >
                  <svg-icon name="goto-detail" width="30" height="30" />
                </a>
              </div>
            </el-col>
            <el-col :xs="23" :sm="12" :md="12" :lg="12">
              <div class="security-col">
                <span class="sec-title">Strategy Address:</span>&nbsp;&nbsp;
                <span class="sec-val">{{currentPair.strategyAddress}}</span>
                <a
                  :href="'https://goerli.etherscan.io/address/'+currentPair.strategyAddress"
                  style="cursor:hand"
                  target="_black"
                >
                  <svg-icon name="goto-detail" width="30" height="30" />
                </a>
              </div>
            </el-col>
          </el-row>
          <div class="emergency-container" v-show="showEmergencyButton">
            <el-button
              type="warning"
              :loading="emergencyLoading"
              @click="emergencyBurn"
            >EmergencyBurn</el-button>&nbsp;&nbsp; &nbsp;&nbsp; &nbsp;&nbsp;
            <a
              :href="goToEtherscan"
              target="_blank"
              :class="['btn','btn-primary',emergencyEtherscanDisable?'disabled':'']"
              role="button"
            >View on etherscan</a>
          </div>
        </div>
      </el-tab-pane>
      <el-tab-pane>
        <span slot="label">
          <i class="el-icon-setting"></i> Settings
        </span>
        <div class="table-responsive">
        <table class="table-setting">
          <tr>
            <td class="setting-title">UniPortionRatio:</td>
            <td>
              <el-input v-model="currentPair.uniPortionRatio" style="width:500px;"></el-input>
            </td>
            <td>
              <el-button type="primary" :loading="uniPortionRatioLoading" @click="setUniPortionRatio">SetUniPortionRatio</el-button>&nbsp;&nbsp;
              <a
                :href="goToEtherscan"
                target="_blank"
                :class="['btn','btn-primary',uniPortionRatioEtherscanDisable?'disabled':'']"
                role="button"
              >View on etherscan</a>
            </td>
          </tr>
          <tr>
            <td class="setting-title">MaxTotalSupply:</td>
            <td>
              <el-input v-model="currentPair.maxTotalSupply" style="width:500px;"></el-input>
            </td>
            <td>
              <el-button
                type="primary"
                :loading="maxTotalSupplyLoading"
                @click="setMaxTotalSupply"
              >SetMaxTotalSupply</el-button>&nbsp;&nbsp;
              <a
                :href="goToEtherscan"
                target="_blank"
                :class="['btn','btn-primary',maxTotalSupplyEtherscanDisable?'disabled':'']"
                role="button"
              >View on etherscan</a>
            </td>
          </tr>
          <tr>
            <td class="setting-title">ProtocolFee:</td>
            <td>
              <el-input v-model="currentPair.protocolFee" style="width:500px;"></el-input>
            </td>
            <td>
              <el-button
                type="primary"
                :loading="protocolFeeLoading"
                @click="setProtocolFee"
              >SetProtocolFee</el-button>&nbsp;&nbsp;
              <a
                :href="goToEtherscan"
                target="_blank"
                :class="['btn','btn-primary',protocolFeeEtherscanDisable?'disabled':'']"
                role="button"
              >View on etherscan</a>
            </td>
          </tr>
          <tr>
            <td class="setting-title">Governance:</td>
            <td>
              <el-input v-model="currentPair.governanceAddress" style="width:500px;"></el-input>&nbsp;&nbsp;
              <a
                  :href="'https://goerli.etherscan.io/address/'+currentPair.governanceAddress"
                  style="cursor:hand"
                  target="_black"
                >
                  <svg-icon name="goto-detail" width="30" height="30" />
                </a>
            </td>
            <td>
              <el-button
                type="primary"
                :loading="governanceLoading"
                @click="setGovernance"
              >SetGovernance</el-button>
              <el-button
                type="primary"
                :loading="acceptGovernanceLoading"
                @click="acceptGovernance"
              >AccecptGovernance</el-button>&nbsp;&nbsp;
              <a
                :href="goToEtherscan"
                target="_blank"
                :class="['btn','btn-primary',governanceEtherscanDisable?'disabled':'']"
                role="button"
              >View on etherscan</a>
            </td>
          </tr>
          <tr>
            <td class="setting-title">Team:</td>
            <td>
              <el-input v-model="currentPair.teamAddress" style="width:500px;"></el-input>&nbsp;&nbsp;
              <a
                  :href="'https://goerli.etherscan.io/address/'+currentPair.teamAddress"
                  style="cursor:hand"
                  target="_black"
                >
                  <svg-icon name="goto-detail" width="30" height="30" />
                </a>
            </td>
            <td>
              <el-button type="primary" :loading="teamLoading" @click="setTeam" size="medium">SetTeam</el-button>&nbsp;&nbsp;
              <a
                :href="goToEtherscan"
                target="_blank"
                :class="['btn','btn-primary',teamEtherscanDisable?'disabled':'']"
                role="button"
              >View on etherscan</a>
            </td>
          </tr>
        </table>
        </div>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator'
import PairsData from '../../common/PairsData'
import { contractInstance } from '../../common/Web3'
import { getEtherscanTx, getEtherscanAddress } from '../../common/Etherscan'
import Pairs from '../../model/Pairs'
import IonRangeSlider from '@/components/IonRangeSlider/index.vue'
import { priceToTick, tickToPrice } from '../../common/Tools'
import $ from 'jquery'
import 'bootstrap'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.min.js'

const ViaLendFeeMakerABI = require('../../abi/ViaLendFeeMakerABI.json')
const VaultFactoryABI = require('../../abi/VaultFactory.json')
Vue.prototype.$ = $
const ethereum = (window as any).ethereum

@Component({
  name: 'Proudcts',
  components: { IonRangeSlider }
})
export default class extends Vue {
  pairsData = new PairsData()
  pairsList = this.pairsData.pairsList
  pairSelectedIndex = ''
  factoryAddress = ''
  showEmergencyButton = false
  currentPair = new Pairs()
  rangeStatusDisplay = 'none'
  errorRebalance = ''
  rebalanceLoading = false
  removePositionsLoading = false
  emergencyLoading = false
  protocolFeeLoading = false
  teamLoading = false
  uniPortionRatioLoading = false
  maxTotalSupplyLoading = false
  governanceLoading = false
  acceptGovernanceLoading = false
  deployLoading = false
  doRebalanceEtherscanDisable = true
  removePositionsEtherscanDisable = true
  emergencyEtherscanDisable = true
  protocolFeeEtherscanDisable = true
  uniPortionRatioEtherscanDisable = true
  maxTotalSupplyEtherscanDisable = true
  governanceEtherscanDisable = true
  teamEtherscanDisable = true
  loadingTokensInfoStatus = false
  currentPercent = 0
  rangeStatusStyle = ''
  tickRange = [20, 80]
  step = 1
  maxTotalSupply = 0
  uniPortionRatio = 0
  governanceAddress = ''
  teamAddress = ''
  uniPortion = 0
  protocolFee = 0
  viewOnEtherscanDisable = true
  transactionHash = ''
  goToEtherscan = ''
  // priceRangeSlider = {}
  dataRange = { from: 0, to: 100 }
  rangeStatus = ''
  precision = 18
  rangeForm = {
    minPrice: 0.0,
    maxPrice: 0.0
  }

  rebalanceTransHashLink = ''
  removePositionTransHashLink = ''

  get pairsListCount() {
    return this.pairsData.pairsList.size()
  }

  get newMinPrice() {
    return this.rangeForm.minPrice
  }

  get newMaxPrice() {
    return this.rangeForm.maxPrice
  }

  async doRebalance() {
    if (!this.$store.state.isConnected) {
      this.$message('Please connect wallet!')
      return
    }
    const _this = this
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
    const keeperContract = await contractInstance(ViaLendFeeMakerABI, this.currentPair.vaultAddress)
    if (keeperContract != null) {
      console.log('account address is ' + ethereum.selectedAddress, ';ticklower:', this.currentPair.tickLower)
      console.log('this.rangeForm.tickLower=', Math.round(parseInt(this.currentPair.tickLower.toString()) / 60) * 60)
      console.log('this.rangeForm.tickUpper=', Math.round(parseInt(this.currentPair.tickUpper.toString()) / 60) * 60)

      keeperContract.methods
        .strategy1(
          BigInt(Math.round(parseInt(this.currentPair.tickLower.toString()) / 60) * 60),
          BigInt(Math.round(parseInt(this.currentPair.tickUpper.toString()) / 60) * 60)
        )
        .send({
          from: ethereum.selectedAddress,
          // gasPrice: '1000000000',
          // gas: 900000,
          value: 0
        })
        .on('confirmation', function(confirmationNumber:number, receipt:any) {
          console.log((new Date()).toLocaleString(), ':{deposit confirm number:', confirmationNumber, ',receipt:', receipt.status, '}')
        })
        .on('receipt', function(receipt:any) {
          if (_this.rebalanceLoading) {
            _this.rebalanceLoading = false
          }
          if (receipt.status) {
            _this.rebalanceTransHashLink = getEtherscanTx(receipt.transactionHash)
            _this.doRebalanceEtherscanDisable = false
            _this.$message('Rebalance submitted!')
            // _this.loadTokensBalance()
          } else { _this.$message('Rebalance failed!') }
        })
        .on('error', function(error:any) {
          _this.rebalanceLoading = false
          _this.$message.error(error)
        })
    }
  }

  dotStyle() {
    console.log('')
  }

  async doRemovePositions() {
    const _this = this
    const keeperContract = await contractInstance(ViaLendFeeMakerABI, this.currentPair.vaultAddress)
    this.removePositionsLoading = true
    if (keeperContract != null) {
      keeperContract.methods
        .alloc()
        .send({
          from: ethereum.selectedAddress,
          // gasPrice: '1000000000',
          // gas: 900000,
          value: 0
        })
        .on('confirmation', function(confirmationNumber:number, receipt:any) {
          console.log((new Date()).toLocaleString(), ':{deposit confirm number:', confirmationNumber, ',receipt:', receipt.status, '}')
        })
        .on('receipt', function(receipt:any) {
          if (_this.removePositionsLoading) {
            _this.removePositionsLoading = false
          }
          if (receipt.status) {
            _this.removePositionTransHashLink = getEtherscanTx(receipt.transactionHash)
            _this.removePositionsEtherscanDisable = false
            _this.$message('RemovePositions submitted!')
            // _this.loadTokensBalance()
          } else { _this.$message('RemovePositions failed!') }
        })
        .on('error', function(error:any) {
          _this.removePositionsLoading = false
          _this.$message.error(error)
        })
    }
  }

  emergencyBurn() {
    console.log()
  }

  setUniPortionRatio() {
    console.log('')
  }

  setMaxTotalSupply() {
    console.log('')
  }

  setProtocolFee() {
    console.log('')
  }

  setGovernance() {
    console.log('')
  }

  acceptGovernance() {
    console.log('')
  }

  setTeam() {
    console.log('')
  }

  calcRangeStatus(pair: Pairs) {
    if (pair.tickLower !== 0 && pair.tickUpper !== 0) {
      const leftMargin = pair.tickLower - 10000
      const rightMargin = pair.tickUpper + 10000
      const leftPercent = parseInt(
        (
          ((pair.tickLower - leftMargin) / (rightMargin - leftMargin)) *
          100
        ).toString()
      )
      const rightPercent = parseInt(
        (
          ((pair.tickUpper - leftMargin) / (rightMargin - leftMargin)) *
          100
        ).toString()
      )
      this.currentPercent = parseInt(
        (
          ((pair.currentTick - leftMargin) / (rightMargin - leftMargin)) *
          100
        ).toString()
      )
      this.tickRange = [leftPercent, rightPercent]
      console.log('currentPercent=', this.currentPercent)
      if (
        pair.tickLower > pair.currentTick ||
        pair.tickUpper < pair.currentTick
      ) {
        this.rangeStatusStyle = 'outOfRange'
        this.rangeStatus = 'Out of range'
      } else {
        this.rangeStatusStyle = 'inRange'
        this.rangeStatus = 'In range'
      }
      this.rangeStatusDisplay = ''
    }
  }

  async created() {
    console.log('bridgeAddress value123=', this.pairsData.bridgeAddress)
    console.log('pairsList.size=', this.pairsData.pairsList.size())
    if (this.$store.state.validNetwork && this.$store.state.isConnected && this.pairsData.pairsList.size() === 0) {
      console.log('pair loading')

      await this.pairsData.loadPairsInfo()
    }
    this.factoryAddress = await this.$store.dispatch('getSessionData', { key: 'factoryAddress' })
    console.log('factory address = ', this.factoryAddress)
    console.log('strategy address = ', this.currentPair.strategyAddress)
    console.log('vault address = ', this.currentPair.vaultAddress)
  }

  @Watch('pairSelectedIndex')
  async watchPairSelectedIndex(newVal: number, oldVal: number) {
    console.log('pairSelectedIndex:', newVal, ',old value:', oldVal)
    console.log(
      'pairsData.pairsList.get(pairSelectedIndex).token1.balanceInWallet=',
      this.pairsData.pairsList.get(0).token1.balanceInWallet
    )
    this.currentPair = this.pairsData.pairsList.get(newVal)
    this.pairsData.getTokensBalance(this.currentPair)
    this.loadPriceRange(this.currentPair)
    this.calculateRangeStatus()
    // this.pairsData.getPairsSettingData(this.currentPair)
    let stat = '-1'
    if (this.factoryAddress !== undefined && this.factoryAddress !== '') {
      const factoryContract = await contractInstance(VaultFactoryABI, this.factoryAddress)
      stat = await factoryContract.methods.stat(this.currentPair.strategyAddress, this.currentPair.vaultAddress).call()
    }
    if (stat === '3') {
      this.showEmergencyButton = true
    }
    console.log('Emergency stat=', stat)
  }

  @Watch('newMinPrice')
  watchNewMinPrice(newVal: number, oldVal: number) {
    console.log('new min price=', newVal, 'old min price', oldVal)
    if (!isNaN(newVal)) {
      this.currentPair.tickLower = priceToTick(newVal, this.currentPair.token0.decimals, this.currentPair.token1.decimals)
      console.log('watchNewMinPrice->tickLower:', this.currentPair.tickLower)
    } else { this.currentPair.tickLower = 0 }
    this.calculateRangeStatus()
  }

  @Watch('newMaxPrice')
  watchNewMaxPrice(newVal: number, oldVal: number) {
    console.log('new max price=', newVal, 'old min price', oldVal)
    if (!isNaN(newVal)) {
      console.log('tickUpper->price:', newVal)
      this.currentPair.tickUpper = priceToTick(newVal, this.currentPair.token0.decimals, this.currentPair.token1.decimals)
    } else { this.currentPair.tickUpper = 0 }
    this.calculateRangeStatus()
  }

  calculateRangeStatus() {
    if (this.currentPair.tickLower !== 0 && this.currentPair.tickUpper !== 0) {
      const leftMargin = this.currentPair.tickLower - 10000
      const rightMargin = this.currentPair.tickUpper + 10000
      const leftPercent = parseInt((((this.currentPair.tickLower - leftMargin) / (rightMargin - leftMargin)) * 100).toString())
      const rightPercent = parseInt((((this.currentPair.tickUpper - leftMargin) / (rightMargin - leftMargin)) * 100).toString())
      this.currentPercent = parseInt((((this.currentPair.currentTick - leftMargin) / (rightMargin - leftMargin)) * 100).toString())
      this.tickRange = [leftPercent, rightPercent]
      if (this.currentPair.tickLower > this.currentPair.currentTick || this.currentPair.tickUpper < this.currentPair.currentTick) {
        this.rangeStatusStyle = 'outOfRange'
        this.rangeStatus = 'Out of range'
      } else {
        this.rangeStatusStyle = 'inRange'
        this.rangeStatus = 'In range'
      }
      this.rangeStatusDisplay = ''
    }
  }

  loadPriceRange(pair: Pairs) {
    this.rangeForm.minPrice = tickToPrice(
      pair.tickLower,
      pair.token0.decimals,
      pair.token1.decimals
    )
    this.rangeForm.maxPrice = tickToPrice(
      pair.tickUpper,
      pair.token0.decimals,
      pair.token1.decimals
    )
    console.log(
      'rangeForm.minPrice=',
      this.rangeForm.minPrice,
      'rangeForm.maxPrice=',
      this.rangeForm.maxPrice
    )
    // const _this = this;
    // (<any>$('.js-range-slider')).ionRangeSlider({
    //   skin: 'big',
    //   type: 'double',
    //   grid: true,
    //   prefix: '$',
    //   step: this.step,
    //   onChange: function(data: { from: number, to: number }) {
    //     console.log('from=', data.from)
    //     _this.rangeForm.minPrice = data.from
    //     _this.rangeForm.maxPrice = data.to
    //   }
    // })
    // this.priceRangeSlider = $('.js-range-slider').data('ionRangeSlider');

    ;(this.$refs.priceRangeSlider as any).sayHello()
    // (this.$refs.priceRangeSlider as any).doUpdate(
    //   20,
    //   500,
    //   60,
    //   200
    // )
    ;(this.$refs.priceRangeSlider as any).doUpdate(
      parseFloat(this.rangeForm.minPrice.toString()) - 230 < 0
        ? 0
        : (parseFloat(this.rangeForm.minPrice.toString()) - 230).toFixed(1),
      (parseFloat(this.rangeForm.maxPrice.toString()) + 230).toFixed(1),
      this.rangeForm.minPrice,
      this.rangeForm.maxPrice
    )
    // (this.priceRangeSlider as any).update({
    //   min: (parseFloat(this.rangeForm.minPrice.toString()) - 230) < 0 ? 0 : (parseFloat(this.rangeForm.minPrice.toString()) - 230).toFixed(1),
    //   max: (parseFloat(this.rangeForm.maxPrice.toString()) + 230).toFixed(1),
    //   from: this.rangeForm.minPrice,
    //   to: this.rangeForm.maxPrice
    // })
  }

  //       @Watch('pairsListCount')
  //   watchPairsListChange(newVal: number, oldVal: number) {
  //     console.log('Change of PairsList:', newVal, ',old value:', oldVal)
  //     for (let i = 0; i < newVal; i++) {
  //       console.log('pair id=', this.pairsData.pairsList.get(i).id)
  //       if (!this.pairsData.pairsList.get(i).loadDataCompleted && !this.pairsData.pairsList.get(i).gettingData) {
  //         this.pairsData.pairsList.get(i).gettingData = true
  //         this.pairsData.getPairPublicData(this.pairsData.pairsList.get(i))
  //       }
  //     }
  //   }

  @Watch('$store.state.currentAccount')
  watchCurrentAccount(newVal: string, oldVal: string) {
    console.log('currentAccount:', newVal, ';previousAccount:', oldVal)
    if (newVal !== '' && this.$store.state.validNetwork) {
      console.log(
        'currentAccount changed,pairlist size:',
        this.pairsData.pairsList.size()
      )
      if (
        this.pairsData.pairsList.size() === 0 &&
        !this.pairsData.pairsBaseInfoLoading
      ) {
        this.pairsData.loadPairsInfo()
      } else {
        // this.pairsData.pairsListComplete = false
        // this.loadMyPairsList()
      }
    }
  }

  @Watch('$store.state.isConnected')
  watchConnectionStatus(newVal: boolean) {
    console.log(
      'Dashboard $store.state.isConnected:',
      this.$store.state.isConnected,
      'this.$store.state.validNetwork=',
      this.$store.state.validNetwork
    )
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
  watchNetworkChainId(newVal: number, oldVal: number) {
    console.log('chainid newVal=', newVal, 'oldVal=', oldVal)
    if (oldVal > 0 && this.$store.state.validNetwork) {
      console.log(
        'network chainId changed,pairlist size:',
        this.pairsData.pairsList.size()
      )
      if (
        this.pairsData.pairsList.size() === 0 &&
        !this.pairsData.pairsBaseInfoLoading
      ) {
        this.pairsData.loadPairsInfo()
      }
    } else if (!this.$store.state.validNetwork) {
      this.clearMyData()
    }
  }

  minPriceChange(from: number) {
    // console.log('min price:', val)
    (this.$refs.priceRangeSlider as any).doUpdate(
      (parseFloat(this.rangeForm.minPrice.toString()) - 230 < 0) ? 0 : (parseFloat(this.rangeForm.minPrice.toString()) - 230).toFixed(1),
      (parseFloat(this.rangeForm.maxPrice.toString()) + 230).toFixed(1),
      from,
      this.rangeForm.maxPrice
    )
  }

  maxPriceChange(to: number) {
    // console.log('max price:', val)
    // $('.js-range-slider').data('ionRangeSlider').update({ to: val })
    (this.$refs.priceRangeSlider as any).doUpdate(
      (parseFloat(this.rangeForm.minPrice.toString()) - 230 < 0) ? 0 : (parseFloat(this.rangeForm.minPrice.toString()) - 230).toFixed(1),
      (parseFloat(this.rangeForm.maxPrice.toString()) + 230).toFixed(1),
      this.rangeForm.minPrice,
      to
    )
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
.rebalance_action{
  font-family: monospace,'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
  width:100%;
  margin:0 auto;
  text-align: center;

}

.remove_position_action{
  width:180px;
    font-family: monospace,'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
    border: 1px solid #1890ff;
    background-color: #409eff;
    color:#ffffff;
    font-size:16px;
    margin-top:10px;
    padding-left: 10px;
    border-radius: 4px;
}

.godetail{
  padding:10px 10px !important;
}
.godetail a{
  color:#ffffff;
}
.table>tbody{
  border:none !important;
}
.table>:not(caption)>*>*{
  border: none !important;
}
.el-link.el-link--default{
  color:#ffffff !important;
}
.rebalance_action a{
  color:#ffffff;
}
.range-text{
  float:left;
}
.range-status{
  float:left;
  margin-left:10px;
}
.ranger-slider{
width:100%;
padding:30px;
}
.security-col {
  border-radius: 4px;
  border: 1px solid #ebeef5;
  background-color: #fff;
  overflow: hidden;
  color: #303133;
  box-shadow: 0 2px 12px 0 rgb(0 0 0 / 10%);
  transition: 0.3s;
  line-height: 20px;
  padding: 12px;
  margin-bottom: 10px;
}
.security-col .sec-title {
  font-family: Georgia, 'Times New Roman', Times, serif;
  font-size: 16px;
  font-weight: bold;
}
.security-col .sec-val {
  font-family: sans-serif;
  font-size: 14px;
}
.emergency-container {
  margin: 0 auto;
  text-align: center;
  padding-top: 30px;
}
.pairs-selector {
  width: 100%;
  margin-bottom: 20px;
  border: 1px solid #e7e7e7;
  background-color: #f8f8f8;
}
.el-input-number--medium{
  width:300px !important;
}

.table-setting{
    width:100%;
}
.table-setting tr{
    height:80px;
    border-bottom: 1px dashed #ddd;
}
.table-setting .setting-title {
  font-family: Georgia, 'Times New Roman', Times, serif;
  font-size: 16px;
  font-weight: bold;
}
.el-select {
  padding: 10px;
}
.el-header {
  padding: 0px !important;
}
.el-main {
  padding: 0px !important;
}
.el-menu {
  border-right: 0px !important;
}
.el-aside {
  background-color: #304156 !important;
}

.breadcrumb {
  padding: 20px;
  margin-bottom: 0px !important;
}
.block {
  float: left;
}
/deep/.el-tabs .el-tabs__item{
   font-size: 16px;
}
.el-divider--horizontal {
  margin: 0px !important;
}
.range-title {
  display: inline;
  font-size: 16px;
  font-weight: bold;
  color: black;
}
.range-status {
  padding-bottom: 30px;
}
.el-tag {
  font-size: 14px !important;
}
.tick-info {
  display: inline;
  float: right;
}
.tick-info ul {
  list-style: none;
  margin-left: 50px;
}
.tick-info li {
  display: inline;
  margin: 5px;
  font-family: serif, Georgia, 'Times New Roman', Times;
  font-size: 14px;
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
.list-group li {
  padding: 20px;
  color: white;
}
.list-group a {
  color: white;
}
.list-group-item {
  background-color: #545c64 !important;
}
.list-group-item.active,
.list-group-item.active:focus,
.list-group-item.active:hover {
  background-color: #337ab7 !important;
  border-color: #337ab7 !important;
}
.menu_text a {
  color: #fff;
}
.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}
.range-text
.clearfix:before,
.clearfix:after {
  display: table;
  content: '';
}
.clearfix:after {
  clear: both;
}

.box-card {
  width: 100%;
  border-radius: 16px;
}
.box-card:hover {
  border-style: solid;
}
.el-card /deep/ .el-card__header {
  color: #000000;
  padding: 18px 20px 5px 20px !important;
}
.el-card /deep/ .el-card__body {
  color: #ffffff;
}
.box-card-currency {
  width: 900px;
  margin: 10px;
}
.price-setting {
  padding-top: 50px;
  text-align: center;
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
.pairs-container {
  padding: 20px;
}
.el-container {
  height: 100%;
}
.el-input-number.is-without-controls .el-input__inner {
  width: 200px;
  line-height: 30px;
  height: 28px;
}
.el-row .col1,
.el-row .col2,
.el-row .col3,
.el-row .col4 {
  height: 30px;
  padding: 0px !important;
  line-height: 26px;
  border: 0px solid #dcdfe6;
}
@media only screen and (max-width: 767px) {
  .portfolio-container {
    width: 100%;
  }
  .el-row .col1 {
    border-top: 1px solid #1c1c22;
  }
  .el-row .col2 {
    border-top: 1px solid #1c1c22;
  }
  .el-row .col3 {
    border-top: 1px solid #1c1c22;
  }
  .el-row .col4 {
    border-top: 1px solid #1c1c22;
  }
}

@media only screen and (min-width: 768px) {
  .portfolio-container {
    width: 630px;
  }
  .el-row .col1 {
    border-right: none;
  }
  .el-row .col2 {
    border-right: none;
  }
  .el-row .col3 {
    border-right: none;
  }
}

@media only screen and (min-width: 992px) {
  .portfolio-container {
    width: 1100px;
  }
  .el-row .col1 {
    border-right: none;
  }
  .el-row .col2 {
    border-right: none;
  }
  .el-row .col3 {
    border-right: none;
  }
}
@media only screen and (min-width: 1200px) {
  .portfolio-container {
    width: 1100px;
  }
  .el-row .col1 {
    border-right: none;
  }
  .el-row .col2 {
    border-right: none;
  }
  .el-row .col3 {
    border-right: none;
  }
}
@media only screen and (min-width: 1920px) {
  .portfolio-container {
    width: 1100px;
  }
  .el-row .col1 {
    border-right: none;
  }
  .el-row .col2 {
    border-right: none;
  }
  .el-row .col3 {
    border-right: none;
  }
}
</style>

<template>
  <div class="myliquidity_box">
    <div class="container">
      <div class="row">
        <div class="col">
          <div class="myliquidity_box_container d-flex flex-row align-items-center justify-content-start">
            <div class="my_liquidity_container">
              <div class="my-liquidity-title">Portfolio</div>
              <div class="add-liquidity">
                <el-button type="next"
                           icon="el-icon-plus"
                           @click="showNewPositionDialog">New Position</el-button>
              </div>
              <table class="table">
                <thead>
                  <th>Smart Vaults</th>
                  <th>Fee Tier</th>
                  <th>Current APR</th>
                  <th>TVL</th>
                  <th>Assets ratio</th>
                  <th>&nbsp;</th>
                </thead>
                <tbody>
                  <tr v-for="item in pairsData"
                      :key="item.id">
                    <td>
                      <span>
                        <img :src="item.smartVaults[0].iconLink"
                             width="30"
                             height="30" />
                        <img :src="item.smartVaults[1].iconLink"
                             width="30"
                             height="30" />
                      </span>&nbsp;&nbsp;
                      <span>{{ item.smartVaults[0].symbol }} / {{ item.smartVaults[1].symbol }}</span>
                    </td>
                    <td><input class="btn btn-default btn-sm"
                             type="button"
                             value="0.30%"></td>
                    <td>-</td>
                    <td>
                      <span v-loading="tvlDataLoading"
                            element-loading-spinner="el-icon-loading"
                            element-loading-background="rgba(0, 0, 0, 0)">{{Number(myValueToken0Locked).toFixed(2)}}/{{Number(myValueToken1Locked).toFixed(2)}}
                      </span>
                    </td>
                    <td>
                      <span v-loading="assetsRatioLoading"
                            element-loading-spinner="el-icon-loading"
                            element-loading-background="rgba(0, 0, 0, 0)">
                        {{Number(lendingRatio).toFixed(2)}}% Lending<br>
                        {{Number(uniswapRatio).toFixed(2)}}% Uniswap V3</span>
                    </td>
                    <td>
                      <el-dropdown @command="handleCommand">
                        <span class="el-dropdown-link">
                          <i class="el-icon-arrow-down el-icon--right"></i>
                        </span>
                        <el-dropdown-menu slot="dropdown">
                          <el-dropdown-item command="withdraw">Withdraw</el-dropdown-item>
                        </el-dropdown-menu>
                      </el-dropdown>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
    <el-dialog :visible.sync="newPositionDialogVisible"
               :append-to-body="true"
               width="380px"
               :close-on-click-modal="false"
               center>
      <span class="dialog_title"
            slot="title">{{newPositionDialogTitle}}</span>
      <!--Step1 content-->
      <div v-loading="vaultInfoLoading"
           :style="
           {display:dialogStep1Display}">
        <el-select v-model="selectedPair"
                   placeholder="Please select pairs"
                   size="medium"
                   style="width:300px;"
                   @change="selectedPairChanage">
          <el-option v-for="item in pairsData"
                     :key="item.id"
                     :label="item.smartVaults[0].symbol + '/' + item.smartVaults[1].symbol + '(' + item.feeTier + ')'"
                     :value="item.id">
            <span style="float: right">{{ item.smartVaults[0].symbol }} / {{ item.smartVaults[1].symbol }}({{item.feeTier}})</span>
            <span style="float: left; color: #8492a6; font-size: 13px"><img :src="item.smartVaults[0].iconLink"
                   width="20" />&nbsp;&nbsp;<img :src="item.smartVaults[1].iconLink"
                   width="20" /></span>
          </el-option>
        </el-select>
        <div class="add_position_intro">
          <p>This vault will concentrate a percentage of your funds as likquidity in Uniswap V3 and use the remaining funds to earn interest in lending pools(Compound,AAVE). The vault rebalances its inventory between lending pools and Uniswap to avoid paying high swap fees & permanent loss.</p>
        </div>
        <el-card class="vault_title">
          <div slot="header"
               class="clearfix">
            <span>Vault Information</span>
          </div>
          <div class="text item">
            <table class="vault_info">
              <tr>
                <td class="c1"
                    width="50%">TVL</td>
                <td class="c2"
                    width="50%">${{Number(tvl).toFixed(2)}}</td>
              </tr>
              <tr>
                <td class="c1">% of cap used</td>
                <td class="c2">{{ofCapUsed}}%</td>
              </tr>
              <tr>
                <td class="c1">Max TVL</td>
                <td class="c2">${{maxTVL}}</td>
              </tr>
              <tr>
                <td class="c1">Range</td>
                <td class="c2">{{vaultRange}}</td>
              </tr>
              <tr>
                <td class="c1">Lending</td>
                <td class="c2">{{vaultLending}}</td>
              </tr>
              <tr>
                <td class="c1">Current APR</td>
                <td class="c2">{{currentAPR}}</td>
              </tr>
            </table>
          </div>
        </el-card>
        <div class="action_list">
          <el-button type="info"
                     style="width:45%;"
                     @click="closeNewPositionDialog">Cancel</el-button>
          <el-button type="next"
                     style="width:45%;float: right;"
                     @click="toStep2">Next</el-button>
        </div>
      </div>
      <!--Step2 content-->
      <div :style="{display:dialogStep2Display}">
        <div class="step2_pair">
          <span style=" color: #8492a6; font-size: 13px"><img :src="this.pairsData[selectedPair-1].smartVaults[0].iconLink"
                 width="30" />&nbsp;&nbsp;<img :src="this.pairsData[selectedPair-1].smartVaults[1].iconLink"
                 width="30" /></span>&nbsp;&nbsp;
          <span>{{this.pairsData[selectedPair-1].smartVaults[0].name}} / {{this.pairsData[selectedPair-1].smartVaults[1].name}}({{this.pairsData[selectedPair-1].feeTier}})</span>
        </div>
        <div class="step2_intro">
          <p>To supply or withdraw liquidity for USDC / WETH to the Vialend protocol,you need to enable it first.</p>
          <p>All funds are compounded during rebalance. 10% of fees earned from Uniswap and Compound will go to a Vialend treasury to be controlled by the Vialend governance community.</p>
        </div>
        <div class="step_token">
          <div style="float: left;margin:8px;width:40px;">
            <img :src="this.pairsData[selectedPair-1].smartVaults[0].iconLink"
                 width="38" />
          </div>
          <div style="float:left;width:236px;">
            <span class="token_balance">Balance:{{token0BalanceInWallet}}(<el-link type="primary"
                       class="token_balance_max"
                       @click="newLiqudityToken0 = token0BalanceInWallet">Max</el-link>)</span>
            <span class="token_textbox">
              <el-input v-model="newLiqudityToken0"
                        placeholder="0.00"
                        dir="rtl"
                        type="text"
                        pattern="^[0-9]*[.,]?[0-9]*$"
                        autocorrect="off"
                        autocomplete="off"
                        inputmode="decimal"></el-input>
            </span>
          </div>
        </div>

        <div class="step_token">
          <div style="float: left;margin:8px;width:40px;">
            <img :src="this.pairsData[selectedPair-1].smartVaults[1].iconLink"
                 width="38" />
          </div>
          <div style="float:left;width:236px;">
            <span class="token_balance">Balance:{{token1BalanceInWallet}}(<el-link type="primary"
                       class="token_balance_max"
                       @click="newLiqudityToken1 = token1BalanceInWallet">Max</el-link>)</span>
            <span class="token_textbox">
              <el-input v-model="newLiqudityToken1"
                        placeholder="0.00"
                        dir="rtl"
                        type="text"
                        pattern="^[0-9]*[.,]?[0-9]*$"
                        autocorrect="off"
                        autocomplete="off"
                        inputmode="decimal"></el-input>
            </span>
          </div>
        </div>
        <div class="action_list">
          <el-button type="primary"
                     :disabled="token0Approved == true ? true : false"
                     :loading="approve0Loading"
                     @click="approveToken(0)"
                     :style="{width:'100%',display:(token0Approved ? 'none':'')}">Approve {{token0Name}}</el-button>
        </div>
        <div class="action_list">
          <el-button type="primary"
                     :disabled="token1Approved == true ? true : false"
                     :loading="approve1Loading"
                     @click="approveToken(1)"
                     :style="{width:'100%',display:(token1Approved ? 'none':'')}">Approve {{token1Name}}</el-button>
        </div>
        <div class="action_list">
          <el-button type="info"
                     style="width:45%;"
                     @click="backToStep1">Back</el-button>
          <el-button type="next"
                     @click="deposit"
                     :loading="depositLoading"
                     :disabled="btnDepositDisabled"
                     :style="{width:'45%',float: 'right',display:isConnected?'':'none'}">Next</el-button>
          <el-button type="
                     next"
                     style="width:100%;"
                     @click="connectWallet"
                     :style="{display:isConnected?'none':''}">Connect Wallet</el-button>
        </div>
      </div>
      <!--Step3 content -->
      <div :style="{display:dialogStep3Display}">
        <el-card class="step3_intro">
          <p>
            Vialend Vaults are an experimental open-source protocol created and maintained by an internet governance community.
            The software is recorded as a set of smart contracts that integrate with existing Ethereum applications with no promise to return profits or guarantee the security of funds.
            It is important to recognize that you can lose up to 100% of your funds,with no recourse for compensation.Please ensure you do your own research and understand the risks before proceeding to use these automated vaults.
            We recommend that you review the contracts below and only submit funds you can afford to lose.
          </p>
        </el-card>
        <el-card class="vault_title">
          <div slot="header"
               class="clearfix">
            <span>Contracts</span>
          </div>
          <div class="text item">
            <table class="vault_info">
              <tr>
                <td class="c1"
                    width="50%">Vaulit</td>
                <td class="c2"
                    width="50%">0xBD7C6D...&nbsp;&nbsp;<img src="images/todetail.png"
                       width="20" /></td>
              </tr>
              <tr>
                <td class="c1">Strategy</td>
                <td class="c2">0xBD7C6D...&nbsp;&nbsp;<img src="images/todetail.png"
                       width="20" /></td>
              </tr>
              <tr>
                <td class="c1">Uniswap V3 Pool</td>
                <td class="c2">ETH / USDC&nbsp;&nbsp;<img src="images/todetail.png"
                       width="20" /></td>
              </tr>
            </table>
          </div>
        </el-card>
        <div class="sign_agreement">
          <el-checkbox v-model="signAgreement"
                       class="sign_agreement_check"></el-checkbox>I understand the risks and agree to the conditions
        </div>
        <div class="action_list">
          <el-button type="info"
                     style="width:45%;"
                     @click="toStep2">Back</el-button>
          <el-button type="next"
                     style="width:45%;float: right;"
                     @click="toResult">Sign & Add Funds</el-button>
        </div>
      </div>
      <!--Result content-->
      <div :style="{display:dialogResultDisplay}">
        <div class="result_intro">
          <div class="result_icon"><img src="images/success.png" /></div>
          <div class="result_text">
            Your liquidity has been provided to Uniswap and funds loaded in Compound.
          </div>
        </div>
        <div class="result_tokens">
          <table class="result_tokens_list">
            <thead>
              <tr>
                <td class="result_tokens_title"
                    colspan="2">Total Value</td>
              </tr>
            </thead>
            <tr>
              <td class="c1"
                  width="32%"><img :src="this.pairsData[selectedPair-1].smartVaults[0].iconLink"
                     width="20" />&nbsp;&nbsp;{{this.pairsData[selectedPair-1].smartVaults[0].symbol}}</td>
              <td class="c2"
                  width="68%">{{(Number(token0BalanceInVault) / Number(Math.pow(10, token0Decimal))).toFixed(2) }}<br>({{(uniToken0Rate * 100).toFixed(2)}}% in Uniswap {{(lendingToken0Rate * 100).toFixed(2)}}% in Compound)</td>
            </tr>
            <tr>
              <td class="c1"><img :src="this.pairsData[selectedPair-1].smartVaults[1].iconLink"
                     width="20" />&nbsp;&nbsp;{{this.pairsData[selectedPair-1].smartVaults[1].symbol}}</td>
              <td class="c2">{{(Number(token1BalanceInVault) / Number(Math.pow(10, token1Decimal))).toFixed(2) }}<br>({{(uniToken1Rate * 100).toFixed(2)}}% in Uniswap {{(lendingToken1Rate * 100).toFixed(2)}}% in Compound)</td>
            </tr>
          </table>
        </div>
        <div class="action_list">
          <el-button type="next"
                     style="width:100%;"
                     @click="closeNewPositionDialog">Done</el-button>
        </div>
      </div>
    </el-dialog>
    <el-dialog :visible.sync="withdrawDialogVisible"
               :append-to-body="true"
               width="380px"
               :close-on-click-modal="false"
               center>
      <span class="dialog_title"
            slot="title">Withdraw</span>
      <table class="table">
        <tr>
          <td>Vault shares:{{shareValue}}
          </td>
        </tr>
        <tr>
          <td style="text-align:right;">
            <el-button type="danger"
                       size="mini"
                       @click="setSharePercent(25)">25%</el-button>
            <el-button type="danger"
                       size="mini"
                       @click="setSharePercent(50)">50%</el-button>
            <el-button type="danger"
                       size="mini"
                       @click="setSharePercent(75)">75%</el-button>
            <el-button type="danger"
                       size="mini"
                       @click="setSharePercent(100)">Max</el-button>
          </td>
        </tr>
        <tr>
          <td><span class="share_percent">{{sharePercent}}%</span></td>
        </tr>
        <tr>
          <td>
            <el-slider v-model="sharePercent"></el-slider>
            <!-- <el-input placeholder="0.0"
                            type="text"
                            v-model="shareValue"
                            style="text-align:right;"></el-input> -->
          </td>
        </tr>
        <tr>
          <td>
            <el-button type="primary"
                       :disabled="!isConnected"
                       @click="withdraw"
                       :loading="withdrawLoading"
                       style="width:100%;">Withdraw</el-button>
          </td>
        </tr>
      </table>
    </el-dialog>
  </div>
</template>

<script>
import PairsData from '../data/PairsData'
import ViaLendTokenABI from '../ABI/tokenABI.json'
import ViaLendPoolABI from '../ABI/UniswapV3PoolABI.json'

export default {
  name: 'MyPositions',
  data () {
    return {
      selectedPair: 1,
      selectedText: '',
      isConnected: this.$store.state.isConnected,
      newPositionButtonDisable: true,
      vaultInfoLoading: true,
      tvlDataLoading: false,
      assetsRatioLoading: false,
      pairsData: PairsData,
      token0Contract: null,
      token1Contract: null,
      poolContract: null,
      poolAddress: '',
      withdrawDialogVisible: false,
      withdrawLoading: false,
      shareValue: 0,
      sharePercent: 25,
      newPositionDialogVisible: false,
      newPositionDialogTitle: 'Step1:Choose a Vault',
      dialogStep1Title: 'Step1:Choose a Vault',
      dialogStep2Title: 'Step2:Add Liquidity',
      dialogStep3Title: 'Step3:Sign Agreement',
      dialogResultTitle: 'Successful!',
      dialogStep1Display: '',
      dialogStep2Display: 'none',
      dialogStep3Display: 'none',
      dialogResultDisplay: 'none',
      newLiqudityToken0: null,
      newLiqudityToken1: null,
      signAgreement: false,
      feeTier: '',
      token0Address: '',
      token0Name: '',
      token0Symbol: '',
      token0Decimal: 0,
      token0BalanceInWallet: 0,
      token0BalanceInVault: 0,
      token1Address: '',
      token1Name: '',
      token1Symbol: '',
      token1Decimal: 0,
      token1BalanceInWallet: 0,
      token1BalanceInVault: 0,
      shares: 0,
      totalShares: 0,
      tvl: 0,
      tvlTotal0: 0,
      tvlTotal1: 0,
      ofCapUsed: '',
      maxTVL: '',
      vaultRange: '',
      vaultLending: '',
      currentAPR: 0,
      myValueLocked: 0,
      myValueToken0Locked: 0,
      myValueToken1Locked: 0,
      myEarnedValue: 0,
      // deposit variable
      depositToken0: '',
      depositToken1: '',
      token0Approved: false,
      token1Approved: false,
      btnDepositDisabled: false,
      approve0Loading: false,
      approve1Loading: false,
      depositLoading: false,
      allTokensList: null,
      // position data
      cLow: 0,
      cHigh: 0,
      uniswapRatio: 0.0,
      lendingRatio: 0.0,
      uniToken0Rate: 0,
      uniToken1Rate: 0,
      lendingToken0Rate: 0,
      lendingToken1Rate: 0
    }
  },
  created: async function () {
    var _this = this
    this.poolAddress = this.$parent.poolAddress
    this.token0Address = await this.$parent.keeperContract.methods.token0().call()
    this.token1Address = await this.$parent.keeperContract.methods.token1().call()
    // console.log('this.token0Address=', this.token0Address)
    this.poolContract = new web3.eth.Contract(
      ViaLendPoolABI,
      this.poolAddress
    )
    this.token0Contract = new web3.eth.Contract(
      ViaLendTokenABI,
      this.token0Address
    )
    this.token1Contract = new web3.eth.Contract(
      ViaLendTokenABI,
      this.token1Address
    )
    this.cLow = await this.$parent.keeperContract.methods.cLow().call()
    this.cHigh = await this.$parent.keeperContract.methods.cHigh().call()
    console.log('cLow=', this.cLow, ';cHigh=', this.cHigh)
    this.getTokensInfo()
    this.getVaultInfo()
  },
  watch: {
    '$store.state.isConnected': function () {
      this.isConnected = this.$store.state.isConnected
      if (this.isConnected) {
        this.getVaultInfo()
      }
    },
    '$store.state.allTokensList': function () {
      // console.log('pairstoke=', this.$store.state.allTokensList)
      this.allTokensList = this.$store.state.allTokensList
      for (var i = 0; i < this.allTokensList.length; i++) {
        if (this.allTokensList[i].symbol === 'ETH') {
          this.pairsData[0].smartVaults[0].rateOfUSD = this.allTokensList[i].price_usd
        } else if (this.allTokensList[i].symbol === 'USDC') {
          this.pairsData[0].smartVaults[1].rateOfUSD = this.allTokensList[i].price_usd
        }
        if (this.pairsData[0].smartVaults[0].rateOfUSD > 0 && this.pairsData[0].smartVaults[1].rateOfUSD > 0) {
          break
        }
      }
      if (this.allTokensList !== null && this.allTokensList.length > 0) {
        this.getVaultInfo()
        console.log('allTokensList length:', this.allTokensList.length)
        console.log('ETH rateofusd:', this.pairsData[0].smartVaults[0].rateOfUSD)
        console.log('USDC rateofusd:', this.pairsData[0].smartVaults[1].rateOfUSD)
      }
    },
    newPositionDialogVisible (val) {
      if (val === false) {
        console.log('call myvaluelocked function')
        this.getVaultInfo()
      }
    },
    sharePercent (val) {
      this.getShares(val)
    }
  },
  mounted () {
  },
  methods: {
    connectWallet () {
      this.$parent.setWalletStatus()
      console.log('wallet connection status:', this.isConnected)
    },
    async getTokensInfo () {
      // Get Token0 Information
      this.token0Name = await this.token0Contract.methods.name().call()
      console.log('token0Name=', this.token0Name)
      this.token0Symbol = await this.token0Contract.methods.symbol().call()
      console.log('token0Symbol=', this.token0Symbol)
      this.token0Decimal = await this.token0Contract.methods.decimals().call()
      console.log('token0Decimal=', this.token0Decimal)
      this.$parent.token0Symbol = this.token0Symbol
      // Get Token1 Information
      this.token1Name = await this.token1Contract.methods.name().call()
      console.log('token1Name=', this.token1Name)
      this.token1Symbol = await this.token1Contract.methods.symbol().call()
      console.log('token1Symbol=', this.token1Symbol)
      this.token1Decimal = await this.token1Contract.methods.decimals().call()
      console.log('token1Decimal01=', this.token1Decimal)
      this.$parent.token1Symbol = this.token1Symbol
      this.pairsData = [{
        'id': 1,
        'smartVaults': [
          { 'iconLink': 'images/weth.png', 'name': this.token0Name, 'abi': 'ABI/contractABI.js', 'tokenAddress': this.token0Address, 'decimals': this.token0Decimal, 'symbol': this.token0Symbol, 'rateOfUSD': 0 },
          { 'iconLink': 'images/usdc.png', 'name': this.token1Name, 'abi': 'ABI/contractABI.js', 'tokenAddress': this.token1Address, 'decimals': this.token1Decimal, 'symbol': this.token1Symbol, 'rateOfUSD': 0 }
        ],
        'feeTier': '0.30%',
        'currentAPR': '505.66%',
        'capacity': '35.1%',
        'TVL': '$2,366,149',
        'disabled': false
      }]
      // console.log('token0_reteOfUSD=', this.pairsData[0].smartVaults[0].rateOfUSD)
      // console.log('token1_reteOfUSD=', this.pairsData[0].smartVaults[1].rateOfUSD)
    },
    async getVaultInfo () {
      if (!this.isConnected) {
        return
      }
      // Set loading status
      this.vaultInfoLoading = true
      this.$parent.myValueLockLoading = true
      this.tvlDataLoading = true
      this.assetsRatioLoading = true
      // Get Shares
      this.shares = await this.$parent.keeperContract.methods.balanceOf(ethereum.selectedAddress).call()
      console.log('user address=', ethereum.selectedAddress, ';shares=', this.shares)
      this.totalShares = await this.$parent.keeperContract.methods.totalSupply().call()

      // ---------- Get TVL Start-----------------
      console.log('getVaultInfo_cLow=', this.cLow)
      console.log('getVaultInfo_cHigh=', this.cHigh)
      var uniliqs = await this.$parent.keeperContract.methods.getPositionAmounts(BigInt(this.cLow), BigInt(this.cHigh)).call()
      console.log('balance in uniswap:', uniliqs)
      // --Get Lending
      var lendingAmounts = await this.$parent.keeperContract.methods.getLendingAmounts().call()
      console.log('lendingAmounts:', lendingAmounts)
      this.vaultLending = Number(lendingAmounts[0]) + Number(lendingAmounts[1]) // Not yet converted to USD
      var cLending = await this.$parent.keeperContract.methods.getCAmounts().call()
      this.token0BalanceInVault = await this.token0Contract.methods.balanceOf(this.$parent.vaultAddress).call()
      console.log('this.token0BalanceInVault=', this.token0BalanceInVault)
      this.token1BalanceInVault = await this.token1Contract.methods.balanceOf(this.$parent.vaultAddress).call()
      console.log('balance1=', this.token1BalanceInVault)
      console.log('rateOfUSD', this.pairsData[this.selectedPair - 1].smartVaults[0].rateOfUSD)
      // tokenDecimal temporary solution
      var t0Decimal = await this.token0Contract.methods.decimals().call()
      var t1Decimal = await this.token1Contract.methods.decimals().call()
      this.tvlTotal0 = (Number(this.token0BalanceInVault) + Number(uniliqs.amount0) + Number(lendingAmounts[0])) / Number(Math.pow(10, t0Decimal))
      var rateofusd = this.pairsData[this.selectedPair - 1].smartVaults[0].rateOfUSD
      this.tvlTotal1 = (Number(this.token1BalanceInVault) + Number(uniliqs.amount1) + Number(lendingAmounts[1])) / Number(Math.pow(10, t1Decimal))
      console.log('tvlTotal0=', this.tvlTotal0, 'tvlTotal1=', this.tvlTotal1)
      console.log('token0USD=', this.pairsData[0].smartVaults[0].rateOfUSD)
      console.log('token1USD=', this.pairsData[0].smartVaults[1].rateOfUSD)
      this.tvl = this.tvlTotal0 * this.pairsData[0].smartVaults[0].rateOfUSD +
        this.tvlTotal1 * this.pairsData[0].smartVaults[1].rateOfUSD
      console.log('TVL=', this.tvl)
      // ---------- Get TVL End-----------------
      // Get Max TVL
      this.maxTVL = await this.$parent.keeperContract.methods.maxTotalSupply().call()
      console.log('getVaultInfo_maxTVL=', this.maxTVL)
      // Get % of cap used
      this.ofCapUsed = (this.totalShares / this.maxTVL * 100).toFixed(2)
      // Get myValueLocked and myEarnedValue
      if (Number(this.shares) === 0 || Number(this.totalShares) === 0) {
        this.myValueLocked = 0
        this.myEarnedValue = 0
        this.myValueToken0Locked = this.tvlTotal0
        this.myValueToken1Locked = this.tvlTotal1
        this.$parent.myValueLocked = this.myValueLocked
        this.$parent.myValueToken0Locked = 0
        this.$parent.myValueToken1Locked = 0
        this.$parent.myEarnedValue = this.myEarnedValue
      } else {
        this.myValueLocked = this.shares / this.totalShares * (this.tvlTotal0 + this.tvlTotal1)
        // this.myEarnedValue = this.shares / this.totalShares * (this.tvlTotal0 + this.tvlTotal1) - this.shares
        this.myValueToken0Locked = this.tvlTotal0
        this.myValueToken1Locked = this.tvlTotal1
        this.myEarnedValue = 0
        this.$parent.myValueLocked = this.myValueLocked
        this.$parent.myValueToken0Locked = this.shares / this.totalShares * this.tvlTotal0
        this.$parent.myValueToken1Locked = this.shares / this.totalShares * this.tvlTotal1
        this.$parent.myEarnedValue = this.myEarnedValue
        // Assets ratio
        var tmpTickLower = await this.$parent.keeperContract.methods.cLow().call()
        var tmpTickUpper = await this.$parent.keeperContract.methods.cHigh().call()
        var result = await this.$parent.keeperContract.methods.getPositionAmounts(BigInt(tmpTickLower), BigInt(tmpTickUpper)).call()
        var token0BalanceInPool, token1BalanceInPool
        if (result !== undefined && result !== null) {
          token0BalanceInPool = result.amount0
          token1BalanceInPool = result.amount1
        }
        var token0BalanceInLending, token1BalanceInLending
        if (lendingAmounts !== null) {
          token0BalanceInLending = lendingAmounts[0]
          token1BalanceInLending = lendingAmounts[1]
        }
        var totalUniswap = Number(token0BalanceInPool) * 300 + Number(token1BalanceInPool)
        var totalLending = Number(token0BalanceInLending) * 300 + Number(token1BalanceInLending)
        var totalUsdc = totalUniswap + totalLending
        this.uniswapRatio = totalUniswap / totalUsdc * 100
        this.lendingRatio = totalLending / totalUsdc * 100
        this.uniToken0Rate = Number(token0BalanceInPool) / (Number(token0BalanceInPool) + Number(token0BalanceInLending))
        this.uniToken1Rate = Number(token1BalanceInPool) / (Number(token1BalanceInPool) + Number(token1BalanceInLending))
        this.lendingToken0Rate = Number(token0BalanceInLending) / (Number(token0BalanceInPool) + Number(token0BalanceInLending))
        this.lendingToken1Rate = Number(token1BalanceInLending) / (Number(token1BalanceInPool) + Number(token1BalanceInLending))

        console.log('totalUniswap=', totalUniswap)
        console.log('totalLending=', totalLending)
        console.log('total_usdc=', totalUsdc)
        console.log('uniswapRatio=', this.uniswapRatio)
        console.log('lendingRatio=', this.lendingRatio)
      }
      // Cancel loading status
      this.$parent.myValueLockLoading = false
      this.tvlDataLoading = false
      this.assetsRatioLoading = false

      console.log('this.shares=', this.shares, ';this.totalShares=', this.totalShares, ';this.tvl=', this.tvl)
      // Get APR
      if (Number(this.shares) === 0 || Number(this.totalShares) === 0) {
        this.currentAPR = 0
      } else {
        console.log('share=', this.shares, ';totalshares=', this.totalShares, ';tvl=', (this.tvlTotal0 + this.tvlTotal1))
        this.currentAPR = this.shares / this.totalShares * (this.tvlTotal0 + this.tvlTotal1)
      }
      // Get TVL
      // var tvlObj = await this.keeperContract.methods.getTVL().call()
      // if (tvlObj.total0 !== undefined && tvlObj.total1 !== undefined) {
      //   this.tvlTotal0 = tvlObj.total0 * this.pairsData[this.selectedPair - 1].smartVaults[0].rateOfUSD
      //   this.tvlTotal1 = tvlObj.total1 * this.pairsData[this.selectedPair - 1].smartVaults[1].rateOfUSD
      //   console.log('rateToken0=', this.pairsData[this.selectedPair - 1].smartVaults[0].rateOfUSD)
      //   console.log('rateToken1=', this.pairsData[this.selectedPair - 1].smartVaults[1].rateOfUSD)
      //   console.log('tvlTotal0=', this.tvlTotal0)
      //   console.log('tvlTotal1=', this.tvlTotal1)
      //   console.log('maxTVL=', this.maxTVL)
      //   this.tvl = parseInt(this.tvlTotal0 / Math.pow(10, this.token0Decimal) + this.tvlTotal1 / Math.pow(10, this.token1Decimal))

      // }
      // Get Range Data
      // var rangeObj = await this.keeperContract.methods.getPositionAmounts(BigInt(this.cLow), BigInt(this.cHigh)).call()
      // if (rangeObj !== undefined && rangeObj !== null) {
      //   console.log('range object=', JSON.stringify(rangeObj))
      //   this.vaultRange = rangeObj.amount0 + '-' + rangeObj.amount1
      // }
      this.vaultRange = (Math.pow(1.0001, this.cLow) * Math.pow(10, this.token0Decimal - this.token1Decimal)).toFixed(1) + '-' + (Math.pow(1.0001, this.cHigh) * Math.pow(10, this.token0Decimal - this.token1Decimal)).toFixed(1)
      console.log('cLow1=', this.cLow, ';cHigh1=', this.cHigh)

      this.vaultInfoLoading = false
    },
    async getPoolInfo () {
      // this.feeTier = await this.poolContract.methods.fee().call()
      console.log('feeTier=', this.feeTier)
    },
    async getTokensBalanceInWallet () {
      // token0 balance in wallet
      var balanceWei = await this.token0Contract.methods.balanceOf(ethereum.selectedAddress).call()
      this.token0BalanceInWallet = (parseInt(balanceWei / (Math.pow(10, this.token0Decimal)) * 1000) / 1000).toFixed(3)
      // token1 balance in wallet
      balanceWei = await this.token1Contract.methods.balanceOf(ethereum.selectedAddress).call()
      this.token1BalanceInWallet = (parseInt(balanceWei / (Math.pow(10, this.token1Decimal)) * 1000) / 1000).toFixed(3)
      // console.log('token1BalanceInWallet=', this.token1Balance)
    },
    setSharePercent (percent) {
      this.sharePercent = percent
    },
    async getShares (percent) {
      if (this.$parent.keeperContract != null) {
        this.ashares = await this.$parent.keeperContract.methods.balanceOf(ethereum.selectedAddress).call()
        console.log('ashares1=', this.ashares)
        this.shareValue = BigInt(parseInt(this.ashares * percent / 100))
        console.log('shares=', this.shareValue)
      }
    },
    enableDepositFeature () {
      if (this.token0Approved && this.token1Approved) {
        this.btnDepositDisabled = false
      } else {
        this.btnDepositDisabled = true
      }
    },
    approveToken (index) {
      var _this = this
      var tokenContract, tokenName
      if (index === 0) {
        this.token0Approved = false
        tokenContract = this.token0Contract
        tokenName = this.token0Name
      } else {
        this.token1Approved = false
        tokenContract = this.token1Contract
        tokenName = this.token1Name
      }
      console.log('vaultAddress=' + this.$parent.vaultAddress)
      if (tokenContract != null) {
        if (index === 0) { _this.approve0Loading = true } else { _this.approve1Loading = true }
        tokenContract.methods
          .approve(this.$parent.vaultAddress, BigInt(90000000000000000000000))
          .send({
            from: ethereum.selectedAddress,
            gasPrice: '10000000000',
            gas: 200000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            _this.$store.dispatch('setApproveStatus', {
              'token': tokenName, 'status': true
            })
            if (index === 0) {
              _this.token0Approved = true
              if (_this.approve0Loading) {
                _this.approve0Loading = false
                _this.$message(tokenName + ' approved!')
              }
            } else {
              _this.token1Approved = true
              if (_this.approve1Loading) {
                _this.approve1Loading = false
                _this.$message(tokenName + ' approved!')
              }
            }
            if (_this.token0Approved && _this.token1Approved) {
              _this.btnDepositDisabled = false
            }
            console.log('token' + index + ' confirmation')
          })
          .on('receipt', function (receipt) {
            // $('#depositresult').text('Successfully deposited!');
            _this.$store.dispatch('setApproveStatus', {
              'token': tokenName, 'status': true
            })
            if (index === 0) {
              _this.token0Approved = true
              if (_this.approve0Loading) {
                _this.approve0Loading = false
                _this.$message(tokenName + ' approved!')
              }
            } else {
              _this.token1Approved = true
              if (_this.approve1Loading) {
                _this.approve1Loading = false
                _this.$message(tokenName + ' approved!')
              }
            }
            if (_this.token0Approved && _this.token1Approved) {
              _this.btnDepositDisabled = false
            }
            console.log('token' + index + ' receipt')
          })
          .on('error', function (error) {
            if (index === 0) {
              _this.token0Approved = false
            } else {
              _this.token1Approved = false
            }
            if (index === 0) { _this.approve0Loading = false } else { _this.approve1Loading = false }
            _this.$message.error('token' + index + ' error' + error)
            console.log('token' + index + ' error' + error)
          })
      }
    },
    deposit () {
      var _this = this
      console.log('t0=', BigInt(this.newLiqudityToken0 * Math.pow(10, this.token0Decimal)))
      console.log('t1=', BigInt(this.newLiqudityToken0 * Math.pow(10, this.token1Decimal)))
      // return
      if (this.$parent.keeperContract != null) {
        console.log('account address is ' + ethereum.selectedAddress)
        this.depositLoading = true
        this.$parent.keeperContract.methods
          .deposit(
            BigInt(this.newLiqudityToken0 * Math.pow(10, this.token0Decimal)),
            BigInt(this.newLiqudityToken1 * Math.pow(10, this.token1Decimal))
          )
          .send({
            from: ethereum.selectedAddress,
            gasPrice: '80000000000',
            gas: 600000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            if (_this.depositLoading === true) {
              _this.depositLoading = false
              _this.$message('Successfully deposited!')
              _this.supplyDialogVisible = false
              console.log('confirmation')
              _this.getVaultInfo()
              _this.toResult()
            }
          })
          .on('receipt', function (receipt) {
            if (_this.depositLoading === true) {
              _this.depositLoading = false
              _this.$message('Successfully deposited!')
              _this.supplyDialogVisible = false
              console.log('receipt')
              _this.getVaultInfo()
              _this.toResult()
            }
          })
          .on('error', function (error) {
            _this.$message.error(error)
            _this.depositLoading = false
          })
      }
    },
    withdraw () {
      var _this = this
      if (this.$parent.keeperContract != null) {
        console.log('account address is ' + ethereum.selectedAddress)
        this.withdrawLoading = true
        this.$parent.keeperContract.methods
          .withdraw(
            BigInt(this.sharePercent)
          )
          .send({
            from: ethereum.selectedAddress,
            gasPrice: '80000000000',
            gas: 600000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            if (_this.withdrawLoading === true) {
              _this.withdrawLoading = false
              _this.$message('Successfully withdrawed!')
              _this.withdrawDialogVisible = false
              console.log('confirmation')
            }
          })
          .on('receipt', function (receipt) {
            if (_this.withdrawLoading === true) {
              _this.withdrawLoading = false
              _this.$message('Successfully withdrawed!')
              _this.withdrawDialogVisible = false
              console.log('receipt')
            }
          })
          .on('error', function (error) {
            _this.$message.error(error)
            _this.withdrawLoading = false
          })
      }
    },
    showNewPositionDialog () {
      if (!this.isConnected) {
        this.$message('Please connect wallet!')
        return
      }
      this.getVaultInfo()
      this.newPositionDialogTitle = this.dialogStep1Title
      this.newPositionDialogVisible = true
      this.dialogStep1Display = ''
      this.dialogStep2Display = 'none'
      this.dialogStep3Display = 'none'
      this.dialogResultDisplay = 'none'
    },
    showWithdrawDialog () {
      if (!this.isConnected) {
        this.$message('Please connect wallet!')
        return
      }
      this.getShares(this.sharePercent)
      this.withdrawDialogVisible = true
    },
    selectedPairChanage (val) {
      console.log('select value=', val)
      console.log('selectedPair=', this.selectedPair)
    },
    closeNewPositionDialog () {
      this.newPositionDialogVisible = false
    },
    backToStep1 () {
      this.newPositionDialogTitle = this.dialogStep1Title
      this.dialogStep1Display = ''
      this.dialogStep2Display = 'none'
      this.dialogStep3Display = 'none'
      this.dialogResultDisplay = 'none'
    },
    toStep2 () {
      this.newPositionDialogTitle = this.dialogStep2Title
      this.dialogStep1Display = 'none'
      this.dialogStep2Display = ''
      this.dialogStep3Display = 'none'
      this.dialogResultDisplay = 'none'
      this.getTokensBalanceInWallet()
      this.getTokenApproveStatus()
    },
    getTokenApproveStatus () {
      this.$store.dispatch('getApproveStatus', { 'token': this.token0Name }).then(res => {
        console.log(this.token0Name + 'Approved status is:' + JSON.stringify(res))
        if (res === 'true') {
          this.token0Approved = true
          console.log('this.btnToken0ApproveDisabled = true ')
        } else {
          this.token0Approved = false
          console.log('this.btnToken0ApproveDisabled = false ')
        }
        this.enableDepositFeature()
      })
      this.$store.dispatch('getApproveStatus', { 'token': this.token1Name }).then(res => {
        console.log(this.token1Name + 'Approved status is:' + JSON.stringify(res))
        if (res === 'true') {
          this.token1Approved = true
          console.log('this.btnToken1ApproveDisabled = true ')
        } else {
          this.token1Approved = false
          console.log('this.btnToken1ApproveDisabled = false ')
        }
        this.enableDepositFeature()
      })
    },
    // backToStep2 () {
    //   this.newPositionDialogTitle = this.dialogStep2Title
    //   this.dialogStep1Display = 'none'
    //   this.dialogStep2Display = ''
    //   this.dialogStep3Display = 'none'
    //   this.dialogResultDisplay = 'none'
    //   this.getTokensBalanceInWallet()
    // },
    toStep3 () {
      this.newPositionDialogTitle = this.dialogStep3Title
      this.dialogStep1Display = 'none'
      this.dialogStep2Display = 'none'
      this.dialogStep3Display = ''
      this.dialogResultDisplay = 'none'
    },
    toResult () {
      this.getBalanceInVault()
      this.newPositionDialogTitle = this.dialogResultTitle
      this.dialogStep1Display = 'none'
      this.dialogStep2Display = 'none'
      this.dialogStep3Display = 'none'
      this.dialogResultDisplay = ''
    },
    async getBalanceInVault () {
      this.token0BalanceInVault = await this.token0Contract.methods.balanceOf(this.$parent.vaultAddress).call()
      console.log('token0BalanceInVault=', this.token0BalanceInVault)
      this.token1BalanceInVault = await this.token1Contract.methods.balanceOf(this.$parent.vaultAddress).call()
      console.log('token0BalanceInVault=', this.token0BalanceInVault)
    },
    handleCommand (command) {
      // this.$message('click on item ' + command)
      if (command === 'withdraw') {
        this.showWithdrawDialog()
      }
    }
  }
}
</script>

<style scoped>
.dialog_title {
  font-size: 26px;
  color: #000000;
  font-weight: bold;
}
.vault_title .el-card__header {
  padding: 0px !important;
}
.el-button--next {
  background: #9900ff;
  border-color: #9900ff;
  color: #ffffff;
}
.el-button--next.is-disabled {
  background: #ba6af0;
  border-color: #ba6af0;
  color: #ffffff;
}
.el-button--next:hover {
  background: #ba6af0;
  border-color: #ba6af0;
  color: #ffffff;
}
.add_position_intro {
  color: #000000;
  margin-bottom: 10px;
  text-align: justify;
  word-wrap: break-word;
  word-break: keep-all;
  text-justify: inter-ideograph;
}
.add_position_intro p {
  text-align: justify;
  color: #000000;
  line-height: 20px;
  font-size: 14px;
  margin-top: 10px;
  word-wrap: break-word;
  word-break: keep-all;
}
.vault_title {
  font-size: 16px;
  font-weight: bold;
}
.vault_info {
  width: 100%;
  font-size: 12px;
  line-height: 32px;
}
.vault_info .c1 {
  font-size: 14px;
  font-weight: bold;
}
.vault_info .c2 {
  font-size: 14px;
  font-weight: bold;
  color: #9900ff;
}
.action_list {
  width: 100%;
  margin-top: 10px;
}
.step2_pair {
  font-size: 14px;
  padding: 0 0 10px 30px;
  font-weight: bold;
  color: #000000;
}
.step2_intro,
.step3_intro p {
  font-size: 12px;
  margin-top: 0px;
  color: #000000;
}
.step2_intro p {
  font-size: 12px;
  line-height: 20px;
  color: #000000;
}
.step3_intro {
  margin-bottom: 10px;
}
.sign_agreement {
  margin: 5px;
  font-size: 12px;
}
.sign_agreement_check {
  margin: 5px;
  font-size: 12px;
  color: #000000;
}
.step_token {
  width: 100%;
  height: 55px;
  border: 1px solid #cccccc;
  margin-top: 10px;
}
.step_token .token_balance {
  float: right;
  font-size: 12px;
  padding: 3px 3px 0px 3px;
}
.step_token .token_textbox {
  float: right;
  font-size: 12px;
  padding: 3px;
}
.step_token .token_balance .token_balance_max {
  font-size: 12px;
}
.step_token >>> .el-input__inner {
  border: 0px !important;
  height: 26px;
  width: 230px;
}
.result_intro {
  border: 2px solid #9900ff;
  text-align: center;
  width: 100%;
  height: 100%;
  background: #b4a7d6;
}
.result_intro .result_icon {
  padding: 20px;
}
.result_intro .result_text {
  width: 100%;
  text-align: center;
  padding: 0 10px 10px 10px;
  font-weight: bold;
  color: #9900ff;
}
.result_tokens {
  border: 0;
  text-align: center;
  width: 100%;
  height: 100%;
}
.result_tokens_list {
  border: 0;
  text-align: left;
  line-height: 26px;
  width: 100%;
  height: 100%;
  margin-top: 10px;
}
.result_tokens_list .c1 {
  font-size: 14px;
  font-weight: bold;
  vertical-align: top;
}
.result_tokens_list .c2 {
  font-size: 12px;
  font-weight: bold;
}
.result_tokens_title {
  border: 0;
  text-align: left;
  font-size: 14px;
  font-weight: bold;
  line-height: 50px;
}
.myliquidity_box {
  width: 100%;
  background: transparent;
}
.my-liquidity-title {
  display: inline;
  font-size: 20px;
  font-weight: bold;
  color: #000000;
  margin-bottom: 20px;
}
.add-liquidity {
  display: inline;
  float: right;
}
.myliquidity_box_container {
  width: calc(100% + 1px);
  left: 0px;
  background: #ffffff;
  margin-top: -100px;
  z-index: 1;
  padding: 50px 66px 0 83px;
  border: 2px solid #000000;
  border-radius: 10px;
}
.my_liquidity_container {
  width: 100%;
}
.my_liquidity_container .table {
  margin-top: 20px;
}
.my_liquidity_container .table th {
  font-size: 14px;
  font-weight: bold;
  color: #000000;
  text-align: left;
}
.my_liquidity_container .table td {
  /* line-height: 50px; */
  font-size: 12px;
  font-weight: bold;
  color: #000000;
  text-align: left;
}
.table td,
.table th {
  border: 0px solid transparent;
}
.el-dropdown-link {
  cursor: pointer;
  color: #409eff;
}
.el-icon-arrow-down {
  font-size: 12px;
}
</style>

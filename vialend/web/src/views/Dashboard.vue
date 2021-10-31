<template>
  <div v-loading="initDataLoading"
       element-loading-text="Loading"
       element-loading-spinner="el-icon-loading"
       element-loading-background="rgba(0, 0, 0, 0.8)">
    <Header ref="headerComponents" />
    <!-- Menu -->
    <div class="menu">
      <div class="background_image"></div>
      <div class="
          menu_content
          d-flex
          flex-column
          align-items-center
          justify-content-center
        ">
        <ul class="menu_nav_list text-center">
          <li><a href="index.html">Docs</a></li>
          <li><a href="about.html">Vote</a></li>
        </ul>
        <div class="menu_review"><a href="#">Launch App</a></div>
      </div>
    </div>
    <div class="home">
      <div class="parallax_background parallax-window"
           data-parallax="scroll"
           data-speed="0.8"></div>
      <div class="home_container">
        <div class="container">
          <div class="row align-items-center my-financial">
            <div class="col-md-4 text-center"
                 v-loading="myValueLockLoading"
                 element-loading-spinner="el-icon-loading"
                 element-loading-background="rgba(0, 0, 0, 0.8)">
              <ul>
                <li class="title">My Value Locked</li>
                <!-- <li class="value"
                    v-if="pairsList.size() > 0">
                  <ul>
                    <li v-for="pair in pairsList._getData()"
                        :key="pair.id">
                      {{pair.token0.symbol}} / {{pair.token1.symbol}} : {{pair.myValueToken0Locked.toFixed(2)}} / {{pair.myValueToken1Locked.toFixed(2)}} <br>
                      USD:${{ (Number(pair.myValueToken0USDLocked) + Number(pair.myValueToken1USDLocked)).toFixed(2)}}
                    </li>
                  </ul>
                </li> -->
                <li class="value">USD {{Number(myValueLocked).toFixed(2)}}</li>
              </ul>
            </div>
            <div class="col-md-4 text-center apy-container">
              <div class="apy-style"
                   v-loading="myValueLockLoading"
                   element-loading-spinner="el-icon-loading"
                   element-loading-background="rgba(0, 0, 0, 0.8)">
                <ul>
                  <li class="apy-title">Net APY</li>
                  <li class="apy-content">{{netAPYTotal.toFixed(2)}}%</li>
                </ul>
              </div>
            </div>
            <div class="col-md-4 text-center"
                 v-loading="myValueLockLoading"
                 element-loading-spinner="el-icon-loading"
                 element-loading-background="rgba(0, 0, 0, 0.8)">
              <ul>
                <li class="title">My Earned Value</li>
                <li class="value">WETH {{Number(userFee0Total).toFixed(2)}}</li>
                <li class="value">USDC {{Number(userFee1Total).toFixed(2)}}</li>
                <li class="value">USD {{Number(myEarnedValue).toFixed(2)}}</li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!--Portfolio Begin-->
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
                  <tbody v-if="pairsList.size() > 0">
                    <tr v-for="pair in pairsList._getData()"
                        :key="pair.id">
                      <td>
                        <span>
                          <img :src="pair.token0.iconLink"
                               width="30"
                               height="30" />
                          <img :src="pair.token1.iconLink"
                               width="30"
                               height="30" />
                        </span>&nbsp;&nbsp;
                        <span>{{ pair.token0.symbol }} / {{ pair.token1.symbol }}</span>
                      </td>
                      <td><input class="btn btn-default btn-sm"
                               type="button"
                               :value="pair.feeTier"></td>
                      <td>{{Number(pair.currentAPR).toFixed(2)}}%</td>
                      <td>
                        <span v-loading="tvlDataLoading"
                              element-loading-spinner="el-icon-loading"
                              element-loading-background="rgba(0, 0, 0, 0)">
                          {{Number(pair.tvlTotal0).toFixed(2)}} / {{Number(pair.tvlTotal1).toFixed(2)}}<br>
                          ${{(Number(pair.tvlTotal0USD) + Number(pair.tvlTotal1USD)).toFixed(2)}}
                        </span>
                      </td>
                      <td>
                        <span v-loading="assetsRatioLoading"
                              element-loading-spinner="el-icon-loading"
                              element-loading-background="rgba(0, 0, 0, 0)">
                          {{Number(pair.lendingRatio).toFixed(2)}}% Compound<br>
                          {{Number(pair.uniswapRatio).toFixed(2)}}% Uniswap V3</span>
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
                  <tbody v-else>
                    <tr>
                      <td colspan="5"
                          style="text-align:center;">no data</td>
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
          <el-select v-model="selectedPairIndex"
                     placeholder="Please select pairs"
                     size="medium"
                     style="width:300px;"
                     @change="selectedPairChanage">
            <el-option v-for="pair in pairsList._getData()"
                       :key="pair.index"
                       :label="pair.token0.symbol + '/' + pair.token1.symbol + '(' + pair.feeTier + ')'"
                       :value="pair.index">
              <span style="float: right">{{ pair.token0.symbol }} / {{ pair.token1.symbol }}({{pair.feeTier}})</span>
              <span style="float: left; color: #8492a6; font-size: 13px"><img :src="pair.token0.iconLink"
                     width="20" />&nbsp;&nbsp;<img :src="pair.token1.iconLink"
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
                      width="50%"
                      :loading="newPositionTVLLoading"
                      v-if="pairsList.size() > 0">${{Number(pairsList.get(selectedPairIndex).tvl).toFixed(2)}}</td>
                </tr>
                <tr>
                  <td class="c1">% of cap used</td>
                  <td class="c2"
                      v-if="pairsList.size() > 0">{{pairsList.get(selectedPairIndex).ofCapUsed}}%</td>
                </tr>
                <tr>
                  <td class="c1">Max TVL</td>
                  <td class="c2"
                      v-if="pairsList.size() > 0">${{pairsList.get(selectedPairIndex).maxTVL}}</td>
                </tr>
                <tr>
                  <td class="c1">Range</td>
                  <td class="c2"
                      v-if="pairsList.size() > 0">{{pairsList.get(selectedPairIndex).vaultRange}}</td>
                </tr>
                <!-- <tr>
                  <td class="c1">Lending</td>
                  <td class="c2">{{vaultLending}}</td>
                </tr> -->
                <tr>
                  <td class="c1">Current APR</td>
                  <td class="c2"
                      v-if="pairsList.size() > 0">-</td>
                  <!--{{Number(pairsList.get(selectedPairIndex).currentAPR).toFixed(2)}} -->
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
          <div class="step2_pair"
               v-if="pairsList.size() > 0">
            <span style=" color: #8492a6; font-size: 13px"><img :src="pairsList.get(selectedPairIndex).token0.iconLink"
                   width="30" />&nbsp;&nbsp;<img :src="pairsList.get(selectedPairIndex).token1.iconLink"
                   width="30" /></span>&nbsp;&nbsp;
            <span>{{pairsList.get(selectedPairIndex).token0.symbol}} / {{pairsList.get(selectedPairIndex).token1.symbol}}({{pairsList.get(selectedPairIndex).feeTier}})</span>
          </div>
          <div class="step2_intro">
            <p>To supply or withdraw liquidity for USDC / WETH to the Vialend protocol,you need to enable it first.</p>
            <p>All funds are compounded during rebalance. 10% of fees earned from Uniswap and Compound will go to a Vialend treasury to be controlled by the Vialend governance community.</p>
          </div>
          <div class="step_token"
               v-if="pairsList.size() > 0">
            <div style="float: left;margin:8px;width:40px;">
              <img :src="pairsList.get(selectedPairIndex).token0.iconLink"
                   width="38" />
            </div>
            <div style="float:left;width:236px;">
              <span class="token_balance">Balance:{{pairsList.get(selectedPairIndex).token0BalanceInWallet}}(<el-link type="primary"
                         class="token_balance_max"
                         @click="newLiqudityToken0 = pairsList.get(selectedPairIndex).token0BalanceInWallet">Max</el-link>)</span>
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

          <div class="step_token"
               v-if="pairsList.size() > 0">
            <div style="float: left;margin:8px;width:40px;">
              <img :src="pairsList.get(selectedPairIndex).token1.iconLink"
                   width="38" />
            </div>
            <div style="float:left;width:236px;">
              <span class="token_balance">Balance:{{pairsList.get(selectedPairIndex).token1BalanceInWallet}}(<el-link type="primary"
                         class="token_balance_max"
                         @click="newLiqudityToken1 = pairsList.get(selectedPairIndex).token1BalanceInWallet">Max</el-link>)</span>
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
          <div class="action_list"
               v-if="pairsList.size() > 0">
            <el-button type="primary"
                       :disabled="pairsList.get(selectedPairIndex).token0Approved == true ? true : false"
                       :loading="approve0Loading"
                       @click="approveToken(0)"
                       :style="{width:'100%',display:(pairsList.get(selectedPairIndex).token0Approved ? 'none':'')}">Approve {{pairsList.get(selectedPairIndex).token0.symbol}}</el-button>
          </div>
          <div class="action_list"
               v-if="pairsList.size() > 0">
            <el-button type="primary"
                       :disabled="pairsList.get(selectedPairIndex).token1Approved == true ? true : false"
                       :loading="approve1Loading"
                       @click="approveToken(1)"
                       :style="{width:'100%',display:(pairsList.get(selectedPairIndex).token1Approved ? 'none':'')}">Approve {{pairsList.get(selectedPairIndex).token1.symbol}}</el-button>
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
            <table class="result_tokens_list"
                   v-if="pairsList.size() > 0">
              <thead>
                <tr>
                  <td class="result_tokens_title"
                      colspan="2">Total Value</td>
                </tr>
              </thead>
              <tr>
                <td class="c1"
                    width="25%"><img :src="pairsList.get(selectedPairIndex).token0.iconLink"
                       width="20" />&nbsp;&nbsp;{{pairsList.get(selectedPairIndex).token0.symbol}}</td>
                <td class="c2"
                    width="75%">{{(Number(pairsList.get(selectedPairIndex).token0BalanceInVault) / Number(Math.pow(10, pairsList.get(selectedPairIndex).token0.decimals))).toFixed(2) }}<br>({{(pairsList.get(selectedPairIndex).uniToken0Rate * 100).toFixed(2)}}% in Uniswap {{(pairsList.get(selectedPairIndex).lendingToken0Rate * 100).toFixed(2)}}% in Compound)</td>
              </tr>
              <tr>
                <td class="c1"><img :src="pairsList.get(selectedPairIndex).token1.iconLink"
                       width="20" />&nbsp;&nbsp;{{pairsList.get(selectedPairIndex).token1.symbol}}</td>
                <td class="c2">{{(Number(pairsList.get(selectedPairIndex).token1BalanceInVault) / Number(Math.pow(10, pairsList.get(selectedPairIndex).token1.decimals))).toFixed(2) }}<br>({{(pairsList.get(selectedPairIndex).uniToken1Rate * 100).toFixed(2)}}% in Uniswap {{(pairsList.get(selectedPairIndex).lendingToken1Rate * 100).toFixed(2)}}% in Compound)</td>
              </tr>
            </table>
          </div>
          <div class="action_list">
            <a :href="goToEtherscan"
               target="_blank"
               style="width:100%"
               :class="['btn','btn-view-etherscan',depositedEtherscanDisable?'disabled':'']"
               role="button">View on etherscan</a>&nbsp;&nbsp;
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
              <a :href="goToEtherscan"
                 target="_blank"
                 style="width:100%"
                 :class="['btn','btn-view-etherscan',withdrawedEtherscanDisable?'disabled':'']"
                 role="button">View on etherscan</a>&nbsp;&nbsp;
              <el-button type="primary"
                         :disabled="btnWithdrawDisabled"
                         @click="withdraw"
                         :loading="withdrawLoading"
                         style="width:100%;">Withdraw</el-button>
            </td>
          </tr>
        </table>
      </el-dialog>
    </div>
    <!--Portfolio End-->
    <el-dialog title="SMART VAULTS"
               :visible.sync="treatyDialogVisible"
               :append-to-body="true"
               width="30%"
               center>
      <span class="dialog-content">SwapStream's Smart Vaults are a new experimental DeFi application that is open-source software maintained by a distributed autonomous organization.
        The software is recorded as a set of smart contracts on Ethereum with no promise to return profits or guarantee the security of funds.
        The application is not regulated and may not be subjected to the same laws and regulations that are applicable to other types of businesses,organizations,firms,individuals,partnerships,or solo practitioners.
        It is important to recognize that you can lose up to 100% of your funds,with no recourse to compensation.
        Please ensure you do your own research and understand the risks before proceeding to use these vaults.
        For more information you can read and review the contracts at the following location</span>
      <span class="dialog-checkbox">
        <el-checkbox>I understand and agree to these conditions</el-checkbox>
      </span>
      <span slot="footer"
            class="dialog-footer">
        <el-button type="primary"
                   @click="goSign">SIGN</el-button>
        <router-view />
      </span>
    </el-dialog>
  </div>
</template>

<script>
import Header from '@/components/Header.vue'
import contractABI from '../ABI/ViaLendFeeMakerABI.json'
import ViaLendTokenABI from '../ABI/tokenABI.json'
import Token from '../model/Token'
import Pairs from '../model/Pairs'
import ArrayList from '../common/ArrayList'
import PairsData from '../data/PairsData'
import ViaLendPoolABI from '../ABI/UniswapV3PoolABI.json'
import { createContext } from 'react'

export default {
  components: { Header },
  data () {
    return {
      exchangeTimer: '',
      currentAccount: '',
      isConnected: this.$store.state.isConnected,
      pairsList: new ArrayList(),
      selectedPair: 1,
      selectedPairIndex: 0,
      // control's load state
      initDataLoading: false,
      myValueLockLoading: false,
      myEarnedValueLoading: false,
      netAPYLoading: false,
      vaultInfoLoading: true,
      tvlDataLoading: false,
      newPositionTVLLoading: false,
      assetsRatioLoading: false,
      withdrawLoading: false,
      approve0Loading: false,
      approve1Loading: false,
      depositLoading: false,
      // Whether the control is displayed
      treatyDialogVisible: false,
      withdrawDialogVisible: false,
      newPositionDialogVisible: false,
      dialogStep1Display: '',
      dialogStep2Display: 'none',
      dialogStep3Display: 'none',
      dialogResultDisplay: 'none',
      // Whether the control is available
      btnDepositDisabled: false,
      btnWithdrawDisabled: false,
      depositedEtherscanDisable: true,
      withdrawedEtherscanDisable: true,
      // Dialog data information
      newPositionDialogTitle: 'Step1:Choose a Vault',
      dialogStep1Title: 'Step1:Choose a Vault',
      dialogStep2Title: 'Step2:Add Liquidity',
      dialogStep3Title: 'Step3:Sign Agreement',
      dialogResultTitle: 'Successful!',
      // Page-level variables
      signAgreement: false,
      shareValue: 0,
      sharePercent: 25,
      newLiqudityToken0: null,
      newLiqudityToken1: null,
      myValueToken0Locked: 0,
      myValueToken0USDLocked: 0,
      myValueToken1Locked: 0,
      myValueToken1USDLocked: 0,
      myValueLocked_Token0Sum: 0,
      myValueLocked_Token1Sum: 0,
      fees0Total: 0,
      fees1Total: 0,
      userFee0Total: 0,
      userFee1Total: 0,
      myValueLocked: 0.00,
      myEarnedValue: 0.00,
      myEarnedValue1: 0.00,
      netAPYTotal: 0.00,
      goToEtherscan: ''
    }
  },
  created: async function () {
    this.initPageData()
  },
  mounted () {
    window.connectWallet = this.connectWallet
    this.exchangeTimer = setInterval(this.exchangeTokensIntoUSD, 1000)
  },
  computed: {
    rateOfUSDStatus () {
      return this.$store.state.token0RateOfUSD > 0 && this.$store.state.token1RateOfUSD > 0
    }
  },
  beforeDestroy () {
    clearInterval(this.exchangeTimer)
  },
  watch: {
    '$store.state.isConnected': function () {
      this.isConnected = this.$store.state.isConnected
      console.log('Dashboard $store.state.isConnected:', this.isConnected)
      if (this.isConnected) {
        this.loadMyData()
      } else {
        this.myValueToken0Locked = 0
        this.myValueToken1Locked = 0
        this.myValueToken0USDLocked = 0
        this.myValueToken1USDLocked = 0
      }
    },
    // '$store.state.tokenExchangeRateLoaded': function () {
    //   if (this.$store.state.tokenExchangeRateLoaded) {
    //     if (this.tvlTotal0 > 0) this.tvlTotal0USD = Number(this.tvlTotal0USD) * Number(this.$store.state.token0RateOfUSD)
    //     if (this.tvlTotal1 > 0) this.tvlTotal1USD = Number(this.tvlTotal1USD) * Number(this.$store.state.token1RateOfUSD)
    //     if (this.myValueToken0Locked > 0) this.myValueToken0USDLocked = Number(this.myValueToken0Locked) * Number(this.$store.state.token0RateOfUSD)
    //     if (this.myValueToken1Locked > 0) this.myValueToken1USDLocked = Number(this.myValueToken1Locked) * Number(this.$store.state.token1RateOfUSD)
    //   }
    // },
    '$store.state.currentAccount': function () {
      console.log('$store.state.currentAccount:', this.$store.state.currentAccount)
      this.loadMyData()
    },
    '$store.state.chainId': function (newVal, oldVal) {
      if (oldVal > 0) {
        this.initPageData()
      }
      console.log('current chainid changed,id:', this.$store.state.chainId)
      console.log('newVal=', newVal, 'oldVal=', oldVal)
    },
    async newPositionDialogVisible (val) {
      if (val === false) {
        console.log('New Position Dialog closed,get new myvaluelocked data.')
        // this.loadMyData()
      }
    },
    async withdrawDialogVisible (val) {
      if (val === false) {
        console.log('Withdraw Dialog closed,get new myvaluelocked data.')
      }
    },
    sharePercent (val) {
      this.getShares(val)
    }
    // tvlTotal0 (newval) {
    //   if (Number(newval) > 0) {
    //     if (this.$store.state.tokenExchangeRateLoaded) {
    //       this.tvlTotal0USD = Number(this.tvlTotal0) * Number(this.$store.state.token0RateOfUSD)
    //       console.log('tvlTotal0USD=', this.tvlTotal0USD)
    //     }
    //   }
    // },
    // tvlTotal1 (newval) {
    //   if (Number(newval) > 0) {
    //     if (this.$store.state.tokenExchangeRateLoaded) {
    //       this.tvlTotal1USD = Number(this.tvlTotal1) * Number(this.$store.state.token1RateOfUSD)
    //     }
    //   }
    // },
    // myValueToken0Locked (newval) {
    //   if (Number(newval) >= 0) {
    //     if (this.$store.state.tokenExchangeRateLoaded) {
    //       this.myValueToken0USDLocked = Number(this.myValueToken0Locked) * Number(this.$store.state.token0RateOfUSD)
    //       console.log('myValueToken0USDLocked=', this.myValueToken0USDLocked)
    //     }
    //   }
    // },
    // myValueToken1Locked (newval) {
    //   if (Number(newval) >= 0) {
    //     if (this.$store.state.tokenExchangeRateLoaded) {
    //       this.myValueToken1USDLocked = Number(this.myValueToken1Locked) * Number(this.$store.state.token1RateOfUSD)
    //       console.log('myValueToken1USDLocked=', this.myValueToken1USDLocked)
    //     }
    //   }
    // }
  },
  methods: {
    connectWallet () {
      this.$parent.setWalletStatus()
      console.log('wallet connection status:', this.isConnected)
    },
    contractInstance (abi, address) {
      return new web3.eth.Contract(abi, address)
    },
    exchangeTokensIntoUSD () {
      if (this.$store.state.tokenExchangeRateLoaded) {
        this.myValueLocked_Token0Sum = 0
        this.myValueLocked_Token1Sum = 0
        for (var i = 0; i < this.pairsList.size(); i++) {
          if (this.pairsList.get(i).tvlTotal0 >= 0) this.pairsList.get(i).tvlTotal0USD = Number(this.pairsList.get(i).tvlTotal0) * Number(this.$store.state.token0RateOfUSD)
          if (this.pairsList.get(i).tvlTotal1 >= 0) this.pairsList.get(i).tvlTotal1USD = Number(this.pairsList.get(i).tvlTotal1) * Number(this.$store.state.token1RateOfUSD)
          if (this.pairsList.get(i).myValueToken0Locked >= 0) {
            this.myValueLocked_Token0Sum += this.pairsList.get(i).myValueToken0Locked
            this.pairsList.get(i).myValueToken0USDLocked = Number(this.pairsList.get(i).myValueToken0Locked) * Number(this.$store.state.token0RateOfUSD)
          }
          if (this.pairsList.get(i).myValueToken1Locked >= 0) {
            this.myValueLocked_Token1Sum += this.pairsList.get(i).myValueToken1Locked
            this.pairsList.get(i).myValueToken1USDLocked = Number(this.pairsList.get(i).myValueToken1Locked) * Number(this.$store.state.token1RateOfUSD)
          }
        }
        if (this.myValueLocked_Token0Sum >= 0 && this.myValueLocked_Token1Sum >= 0) {
          this.myValueLocked = this.myValueLocked_Token0Sum * Number(this.$store.state.token0RateOfUSD) + this.myValueLocked_Token1Sum * Number(this.$store.state.token1RateOfUSD)
        }
        // this.myEarnedValue = this.fees0Total * Number(this.$store.state.token0RateOfUSD) + this.fees1Total * Number(this.$store.state.token1RateOfUSD)
        this.myEarnedValue = Number(this.userFee0Total) * Number(this.$store.state.token0RateOfUSD) + Number(this.userFee1Total) * Number(this.$store.state.token1RateOfUSD)
      }
      console.log('tokenExchangeRateLoaded:', this.$store.state.tokenExchangeRateLoaded)
    },
    async initPageData () {
      if (await this.$parent.currentChainIsAvailable() === false) {
        for (var i = 0; i < this.pairsList.size(); i++) {
          this.pairsList.get(i).Empty()
          // console.log('pair feeTier=', this.pairsList.get(i).feeTier)
        }
        console.log('Chain is unavailable,clear pairsList:', this.pairsList.size())
        // to be used
        this.myValueToken0Locked = 0.00
        this.myValueToken1Locked = 0.00
        this.myValueToken0USDLocked = 0
        this.myValueToken1USDLocked = 0
      } else {
        this.initDataLoading = true
        this.pairsList.clear()
        var pair1 = new Pairs()
        pair1.index = 0
        pair1.id = 1
        pair1.vaultAddress = '0x6F520a253EC8f4d0B745649a5C02bB7a5201d70b'
        pair1.token0LendingAddress = '0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF'
        pair1.token1LendingAddress = '0xCEC4a43eBB02f9B80916F1c718338169d6d5C1F0'
        pair1 = await this.getPairBasicInfo(pair1)
        pair1 = await this.getPairPublicInfo(pair1)
        this.pairsList.add(pair1)
        var pair2 = new Pairs()
        pair2.index = 1
        pair1.id = 2
        pair2.vaultAddress = '0x6F520a253EC8f4d0B745649a5C02bB7a5201d70b'
        pair2.token0LendingAddress = '0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF'
        pair2.token1LendingAddress = '0xCEC4a43eBB02f9B80916F1c718338169d6d5C1F0'
        pair2 = await this.getPairBasicInfo(pair2)
        pair2 = await this.getPairPublicInfo(pair2)
        this.pairsList.add(pair2)
        console.log('pair1 token0 name is', pair1.token0.iconLink)
        console.log('pair1 token1 name is', pair1.token1.iconLink)
        console.log('pairsList.length=', this.pairsList.size())
        if (this.$store.state.isConnected) {
          this.loadMyData()
        }
        this.initDataLoading = false
      }
    },
    async updatePageData () {
      await this.getPairPublicInfo(this.pairsList.get(this.selectedPairIndex))
      await this.loadMyData()
    },
    async getPairBasicInfo (pair) {
      var token0 = new Token()
      var token1 = new Token()
      var keeperContract = this.contractInstance(contractABI, pair.vaultAddress)
      // poolAddress = await vaultContract.methods.poolAddress().call()
      token0.tokenAddress = await keeperContract.methods.token0().call()
      token1.tokenAddress = await keeperContract.methods.token1().call()
      // poolContract = new web3.eth.Contract(
      //   ViaLendPoolABI,
      //   poolAddress
      // )
      var token0Contract = this.contractInstance(ViaLendTokenABI, token0.tokenAddress)
      var token1Contract = this.contractInstance(ViaLendTokenABI, token1.tokenAddress)
      token0.iconLink = 'images/weth.png'
      token0.name = await token0Contract.methods.name().call()
      token0.symbol = await token0Contract.methods.symbol().call()
      token0.decimals = await token0Contract.methods.decimals().call()
      token1.iconLink = 'images/usdc.png'
      token1.name = await token1Contract.methods.name().call()
      token1.symbol = await token1Contract.methods.symbol().call()
      token1.decimals = await token1Contract.methods.decimals().call()
      pair.cLow = await keeperContract.methods.cLow().call()
      pair.cHigh = await keeperContract.methods.cHigh().call()
      console.log('clow===========', pair.cLow)
      pair.token0 = token0
      pair.token1 = token1
      return pair
    },
    async getPairPublicInfo (pair) {
      this.tvlDataLoading = true
      this.assetsRatioLoading = true
      console.log('pair.vaultAddress=', pair.vaultAddress)
      console.log('pair.token0LendingAddress=', pair.token0LendingAddress)
      console.log('pair.token1LendingAddress=', pair.token1LendingAddress)
      console.log('pair.token0.tokenAddress=', pair.token0.tokenAddress)
      console.log('pair.token1.tokenAddress=', pair.token1.tokenAddress)
      var keeperContract = this.contractInstance(contractABI, pair.vaultAddress)
      var token0LendingContract = this.contractInstance(ViaLendTokenABI, pair.token0LendingAddress)
      var token1LendingContract = this.contractInstance(ViaLendTokenABI, pair.token1LendingAddress)
      var token0Contract = this.contractInstance(ViaLendTokenABI, pair.token0.tokenAddress)
      var token1Contract = this.contractInstance(ViaLendTokenABI, pair.token1.tokenAddress)
      // ---------- Get TVL Begin-----------------
      var uniliqs = await keeperContract.methods.getPositionAmounts(BigInt(pair.cLow), BigInt(pair.cHigh)).call()
      console.log('balance in uniswap:', uniliqs, 'getVaultInfo_cLow=', pair.cLow, 'getVaultInfo_cHigh=', pair.cHigh)
      // -->Get Lending Amounts begin
      var exchangeRate0 = await token0LendingContract.methods.exchangeRateStored().call()
      var exchangeRate1 = await token1LendingContract.methods.exchangeRateStored().call()
      var CAmount0 = await token0LendingContract.methods.balanceOf(pair.vaultAddress).call()
      var CAmount1 = await token1LendingContract.methods.balanceOf(pair.vaultAddress).call()
      var underlying0 = CAmount0 * exchangeRate0 / Math.pow(10, 18)
      var underlying1 = CAmount1 * exchangeRate1 / Math.pow(10, 18)
      console.log('underlying0=', underlying0, 'underlying1=', underlying1)
      // -->Get Lending Amounts end
      // var lendingAmounts = await this.$parent.keeperContract.methods.getLendingAmounts().call()
      pair.vaultLending = Number(underlying0) + Number(underlying1) // Not yet converted to USD
      // var cLending = await this.$parent.keeperContract.methods.getCAmounts().call()
      pair.token0BalanceInVault = await token0Contract.methods.balanceOf(pair.vaultAddress).call()
      pair.token1BalanceInVault = await token1Contract.methods.balanceOf(pair.vaultAddress).call()
      console.log('token0BalanceInVault=', pair.token0BalanceInVault, ';token1BalanceInVault=', pair.token1BalanceInVault)
      var t0Decimal = await token0Contract.methods.decimals().call()
      var t1Decimal = await token1Contract.methods.decimals().call()
      pair.tvlTotal0 = (Number(pair.token0BalanceInVault) + Number(uniliqs.amount0) + Number(underlying0)) / Number(Math.pow(10, t0Decimal))
      pair.tvlTotal1 = (Number(pair.token1BalanceInVault) + Number(uniliqs.amount1) + Number(underlying1)) / Number(Math.pow(10, t1Decimal))
      pair.tvl = pair.tvlTotal0 * 3768.67 + pair.tvlTotal1 * 0.998117
      console.log('tvlTotal0=', pair.tvlTotal0, 'tvlTotal1=', pair.tvlTotal1)
      // console.log('TVL=', this.tvl, ';token0RateOfUSD=', this.pairsList[0].smartVaults[0].rateOfUSD, ';token1RateOfUSD=', this.pairsList[0].smartVaults[1].rateOfUSD)
      // ---------- Get TVL End--------------------------
      // ---------- Get Assets ratio Begin --------------
      var result = await keeperContract.methods.getPositionAmounts(BigInt(pair.cLow), BigInt(pair.cHigh)).call()
      var token0BalanceInPool, token1BalanceInPool, token0BalanceInLending, token1BalanceInLending
      if (result !== undefined && result !== null) {
        token0BalanceInPool = result.amount0
        token1BalanceInPool = result.amount1
      }
      if (!isNaN(underlying0) && !isNaN(underlying1)) {
        token0BalanceInLending = underlying0
        token1BalanceInLending = underlying1
      }
      var totalUniswap = Number(token0BalanceInPool) * 300 + Number(token1BalanceInPool)
      var totalLending = Number(token0BalanceInLending) * 300 + Number(token1BalanceInLending)
      var totalUsdc = totalUniswap + totalLending
      if (totalUsdc === 0) {
        pair.uniswapRatio = 0
        pair.lendingRatio = 0
      } else {
        pair.uniswapRatio = totalUniswap / totalUsdc * 100
        pair.lendingRatio = totalLending / totalUsdc * 100
      }
      if ((Number(token0BalanceInPool) + Number(token0BalanceInLending)) === 0) {
        pair.uniToken0Rate = 0
      } else {
        pair.uniToken0Rate = Number(token0BalanceInPool) / (Number(token0BalanceInPool) + Number(token0BalanceInLending))
      }
      if ((Number(token1BalanceInPool) + Number(token1BalanceInLending)) === 0) {
        pair.uniToken1Rate = 0
      } else {
        pair.uniToken1Rate = Number(token1BalanceInPool) / (Number(token1BalanceInPool) + Number(token1BalanceInLending))
      }
      if ((Number(token0BalanceInPool) + Number(token0BalanceInLending)) === 0) {
        pair.lendingToken0Rate = 0
      } else {
        pair.lendingToken0Rate = Number(token0BalanceInLending) / (Number(token0BalanceInPool) + Number(token0BalanceInLending))
      }
      if ((Number(token1BalanceInPool) + Number(token1BalanceInLending)) === 0) {
        pair.lendingToken1Rate = 0
      } else {
        pair.lendingToken1Rate = Number(token1BalanceInLending) / (Number(token1BalanceInPool) + Number(token1BalanceInLending))
      }
      console.log('totalUniswap=', pair.totalUniswap, 'totalLending=', pair.totalLending, 'total_usdc=', pair.totalUsdc, 'uniswapRatio=', pair.uniswapRatio, 'lendingRatio=', pair.lendingRatio)
      // ---------- Get Assets ratio End --------------
      // Cancel loading status
      this.tvlDataLoading = false
      this.assetsRatioLoading = false
      return pair
    },
    async getVaultInfo () {
      if (!this.isConnected) { return }
      this.vaultInfoLoading = true
      this.myValueLockLoading = true
      var pair = this.pairsList.get(this.selectedPairIndex)
      var keeperContract = this.contractInstance(contractABI, pair.vaultAddress)
      // Get Max TVL
      pair.maxTVL = await keeperContract.methods.maxTotalSupply().call()
      console.log('getVaultInfo_maxTVL=', pair.maxTVL)
      // Get % of cap used
      pair.ofCapUsed = (pair.totalShares / pair.maxTVL * 100).toFixed(2)
      // Get APR
      // if (Number(pair.myShare) === 0 || Number(pair.totalShares) === 0) {
      //   pair.currentAPR = 0
      // } else {
      //   pair.currentAPR = pair.myShare / pair.totalShares * (pair.tvlTotal0 + pair.tvlTotal1)
      // }
      pair.vaultRange = (Math.pow(1.0001, pair.cLow) * Math.pow(10, pair.token0.decimals - pair.token1.decimals)).toFixed(1) + '-' + (Math.pow(1.0001, pair.cHigh) * Math.pow(10, pair.token0.decimals - pair.token1.decimals)).toFixed(1)
      this.vaultInfoLoading = false
      this.myValueLockLoading = false
    },
    async getTokensBalanceInWallet () {
      var pair = this.pairsList.get(this.selectedPairIndex)
      var token0Contract = this.contractInstance(ViaLendTokenABI, pair.token0.tokenAddress)
      var token1Contract = this.contractInstance(ViaLendTokenABI, pair.token1.tokenAddress)
      // token0 balance in wallet
      var balanceWei = await token0Contract.methods.balanceOf(ethereum.selectedAddress).call()
      pair.token0BalanceInWallet = (parseInt(balanceWei / (Math.pow(10, pair.token0.decimals)) * 1000) / 1000).toFixed(3)
      // token1 balance in wallet
      balanceWei = await token1Contract.methods.balanceOf(ethereum.selectedAddress).call()
      pair.token1BalanceInWallet = (parseInt(balanceWei / (Math.pow(10, pair.token1.decimals)) * 1000) / 1000).toFixed(3)
      // console.log('token1BalanceInWallet=', this.token1Balance)
    },
    setSharePercent (percent) {
      this.sharePercent = percent
    },
    async getShares (percent) {
      var keeperContract = this.contractInstance(contractABI, this.pairsList.get(this.selectedPairIndex).vaultAddress)
      if (keeperContract != null) {
        var ashares = await keeperContract.methods.balanceOf(ethereum.selectedAddress).call()
        console.log('ashares=', ashares)
        this.shareValue = BigInt(parseInt(ashares * percent / 100))
        console.log('shares=', this.shareValue)
        if (this.shareValue > 0) this.btnWithdrawDisabled = false
        else this.btnWithdrawDisabled = true
      }
    },
    enableDepositFeature () {
      if (this.pairsList.get(this.selectedPairIndex).token0Approved && this.pairsList.get(this.selectedPairIndex).token1Approved) {
        this.btnDepositDisabled = false
      } else {
        this.btnDepositDisabled = true
      }
    },
    approveToken (index) {
      var _this = this
      var tokenContract, tokenName
      if (index === 0) {
        this.pairsList.get(this.selectedPairIndex).token0Approved = false
        tokenContract = this.contractInstance(ViaLendTokenABI, this.pairsList.get(this.selectedPairIndex).token0.tokenAddress)
        tokenName = this.pairsList.get(this.selectedPairIndex).token0.name
      } else {
        this.pairsList.get(this.selectedPairIndex).token1Approved = false
        tokenContract = this.contractInstance(ViaLendTokenABI, this.pairsList.get(this.selectedPairIndex).token1.tokenAddress)
        tokenName = this.pairsList.get(this.selectedPairIndex).token1.name
      }
      console.log('vaultAddress=' + this.$parent.vaultAddress)
      if (tokenContract != null) {
        if (index === 0) { _this.approve0Loading = true } else { _this.approve1Loading = true }
        tokenContract.methods
          .approve(_this.pairsList.get(this.selectedPairIndex).vaultAddress, BigInt(90000000000000000000000))
          .send({
            from: ethereum.selectedAddress,
            // gasPrice: '10000000000',
            // gas: 200000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            console.log((new Date()).toLocaleString(), ':{approve confirm number:', confirmationNumber, ',receipt:', receipt.status, '}')
          })
          .on('receipt', function (receipt) {
            if (receipt.status) {
              _this.$store.dispatch('setApproveStatus', {
                'token': tokenName, 'status': true
              })
              if (index === 0) {
                _this.pairsList.get(_this.selectedPairIndex).token0Approved = true
                if (_this.approve0Loading) {
                  _this.approve0Loading = false
                  _this.$message(tokenName + ' approved!')
                }
              } else {
                _this.pairsList.get(_this.selectedPairIndex).token1Approved = true
                if (_this.approve1Loading) {
                  _this.approve1Loading = false
                  _this.$message(tokenName + ' approved!')
                }
              }
              if (_this.pairsList.get(_this.selectedPairIndex).token0Approved && _this.pairsList.get(_this.selectedPairIndex).token1Approved) {
                _this.btnDepositDisabled = false
              }
              console.log('token' + index + ' receipt')
            } else {
              _this.$message(tokenName, ' Approve failed!')
            }
          })
          .on('error', function (error) {
            if (index === 0) {
              _this.pairsList.get(_this.selectedPairIndex).token0Approved = false
            } else {
              _this.pairsList.get(_this.selectedPairIndex).token1Approved = false
            }
            if (index === 0) { _this.approve0Loading = false } else { _this.approve1Loading = false }
            _this.$message.error('token' + index + ' error' + error)
            console.log('token' + index + ' error' + error)
          })
      }
    },
    showPairsInfoWithLog () {
      console.log('log_vaultAddress=', this.pairsList.get(this.selectedPairIndex).vaultAddress)
      console.log('log_token0_decimals=', this.pairsList.get(this.selectedPairIndex).token0.decimals)
      console.log('log_token1_decimals=', this.pairsList.get(this.selectedPairIndex).token1.decimals)
      console.log('this.newLiqudityToken0=', this.newLiqudityToken0)
      console.log('this.newLiqudityToken1=', this.newLiqudityToken1)
    },
    deposit () {
      var _this = this
      var keeperContract = this.contractInstance(contractABI, this.pairsList.get(this.selectedPairIndex).vaultAddress)
      var token0Decimal = this.pairsList.get(this.selectedPairIndex).token0.decimals
      var token1Decimal = this.pairsList.get(this.selectedPairIndex).token1.decimals
      this.showPairsInfoWithLog()
      if (keeperContract != null) {
        this.depositLoading = true
        keeperContract.methods
          .deposit(
            BigInt(this.newLiqudityToken0 * Math.pow(10, token0Decimal)),
            BigInt(this.newLiqudityToken1 * Math.pow(10, token1Decimal)),
            true
          )
          .send({
            from: ethereum.selectedAddress,
            // gasPrice: '80000000000',
            // gas: 600000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            // if (_this.depositLoading === true) {
            // _this.depositLoading = false
            // _this.$message('Successfully deposited!')
            // console.log('deposited confirmation')
            // _this.getVaultInfo()
            // _this.toResult()
            // }
            console.log((new Date()).toLocaleString(), ':{deposit confirm number:', confirmationNumber, ',receipt:', receipt.status, '}')
          })
          .on('receipt', function (receipt) {
            if (_this.depositLoading === true) {
              _this.depositLoading = false
              if (receipt.status) {
                _this.newLiqudityToken0 = null
                _this.newLiqudityToken1 = null
                _this.goToEtherscan = 'https://goerli.etherscan.io/tx/' + receipt.transactionHash
                _this.depositedEtherscanDisable = false
                _this.$message('Deposit submitted!')
                _this.updatePageData()
              } else { _this.$message('Deposit failed!') }
              // _this.getVaultInfo()
              _this.toResult()
            }
            console.log((new Date()).toLocaleString(), 'deposit receipt status:', receipt.status)
          })
          .on('error', function (error) {
            _this.$message.error(error)
            _this.depositLoading = false
            console.log('Withdraw error:', error)
          })
      }
    },
    withdraw () {
      var _this = this
      var keeperContract = this.contractInstance(contractABI, this.pairsList.get(this.selectedPairIndex).vaultAddress)
      if (keeperContract != null) {
        this.withdrawLoading = true
        keeperContract.methods
          .withdraw(
            BigInt(this.sharePercent)
          )
          .send({
            from: ethereum.selectedAddress,
            // gasPrice: '80000000000',
            // gas: 600000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            // if (_this.withdrawLoading === true) {
            //   _this.withdrawLoading = false
            //   _this.$message('Successfully withdrawed!')
            //   _this.withdrawDialogVisible = false
            //   _this.getVaultInfo()
            //   console.log('confirmation')
            // }
            console.log((new Date()).toLocaleString(), ':{withdraw confirm number:', confirmationNumber, ',receipt:', receipt.status, '}')
          })
          .on('receipt', function (receipt) {
            if (_this.withdrawLoading === true) {
              _this.withdrawLoading = false
              if (receipt.status) {
                _this.goToEtherscan = 'https://goerli.etherscan.io/tx/' + receipt.transactionHash
                _this.withdrawedEtherscanDisable = false
                _this.getShares(_this.sharePercent)
                _this.$message('Withdraw submitted!')
                _this.updatePageData()
              } else { _this.$message('Withdraw failed!') }
              // _this.withdrawDialogVisible = false
              // _this.getVaultInfo()
              console.log((new Date()).toLocaleString(), 'withdraw receipt status:', receipt.status)
            }
          })
          .on('error', function (error) {
            _this.$message.error(error)
            _this.withdrawLoading = false
            console.log('Withdraw error:', error)
          })
      }
    },
    showNewPositionDialog () {
      if (!this.isConnected) {
        this.$message('Please connect wallet!')
        return
      }
      this.newLiqudityToken0 = null
      this.newLiqudityToken1 = null
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
      this.selectedPairIndex = val
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
      this.$store.dispatch('getApproveStatus', { 'token': this.pairsList.get(this.selectedPairIndex).token0.name }).then(res => {
        if (res === 'true') {
          this.pairsList.get(this.selectedPairIndex).token0Approved = true
          console.log('this.btnToken0ApproveDisabled = true ')
        } else {
          this.pairsList.get(this.selectedPairIndex).token0Approved = false
          console.log('this.btnToken0ApproveDisabled = false ')
        }
        this.enableDepositFeature()
      })
      this.$store.dispatch('getApproveStatus', { 'token': this.pairsList.get(this.selectedPairIndex).token1.name }).then(res => {
        if (res === 'true') {
          this.pairsList.get(this.selectedPairIndex).token1Approved = true
          console.log('this.btnToken1ApproveDisabled = true ')
        } else {
          this.pairsList.get(this.selectedPairIndex).token1Approved = false
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
      var token0Contract = this.contractInstance(ViaLendTokenABI, this.pairsList.get(this.selectedPairIndex).token0.tokenAddress)
      var token1Contract = this.contractInstance(ViaLendTokenABI, this.pairsList.get(this.selectedPairIndex).token1.tokenAddress)
      this.pairsList.get(this.selectedPairIndex).token0BalanceInVault = await token0Contract.methods.balanceOf(this.pairsList.get(this.selectedPairIndex).vaultAddress).call()
      this.pairsList.get(this.selectedPairIndex).token1BalanceInVault = await token1Contract.methods.balanceOf(this.pairsList.get(this.selectedPairIndex).vaultAddress).call()
      console.log('token0BalanceInVault=', this.pairsList.get(this.selectedPairIndex).token0BalanceInVault)
      console.log('token0BalanceInVault=', this.pairsList.get(this.selectedPairIndex).token1BalanceInVault)
    },
    handleCommand (command) {
      // this.$message('click on item ' + command)
      if (command === 'withdraw') {
        this.sharePercent = 25
        this.showWithdrawDialog()
      }
    },
    async loadMyData () {
      if (this.isConnected) {
        this.myValueLockLoading = true
        var myShare, totalShares, tvlTotal0, tvlTotal1
        var ufee0, lfee0, ufee1, lfee1
        var fees
        var keeperContract
        this.fees0Total = 0
        this.fees1Total = 0
        this.userFee0Total = 0
        this.userFee1Total = 0
        this.netAPYTotal = 0
        for (var i = 0; i < this.pairsList.size(); i++) {
          keeperContract = this.contractInstance(contractABI, this.pairsList.get(i).vaultAddress)
          // ---------- Get Shares -------------------
          myShare = await keeperContract.methods.balanceOf(ethereum.selectedAddress).call()
          totalShares = await keeperContract.methods.totalSupply().call()
          // fees = await keeperContract.methods.Fees().call()
          // if (fees != null) {
          //   console.log('fees=', JSON.stringify(fees))
          //   this.fees0Total += Number(assetholder.fees0) / Math.pow(10, Number(this.pairsList.get(i).token0.decimals))
          //   this.fees1Total += Number(assetholder.fees1) / Math.pow(10, Number(this.pairsList.get(i).token1.decimals))
          // }
          // console.log('assetholder=', assetholder)
          tvlTotal0 = this.pairsList.get(i).tvlTotal0
          tvlTotal1 = this.pairsList.get(i).tvlTotal1
          console.log('this.shares=', myShare, ';this.totalShares=', totalShares)
          if (totalShares * tvlTotal0 === 0) {
            this.pairsList.get(i).myValueToken0Locked = 0
          } else {
            this.pairsList.get(i).myValueToken0Locked = myShare * tvlTotal0 / totalShares
          }
          if (totalShares * tvlTotal1 === 0) {
            this.pairsList.get(i).myValueToken1Locked = 0
          } else {
            this.pairsList.get(i).myValueToken1Locked = myShare * tvlTotal1 / totalShares
          }
          this.pairsList.get(i).myShare = myShare
          this.pairsList.get(i).totalShares = totalShares
          fees = await keeperContract.methods.Fees().call()
          if (fees != null) {
            console.log('fees=', JSON.stringify(fees))
            if (Number(totalShares) !== 0) {
              this.userFee0Total += (Number(fees.U3Fees0) + Number(fees.LcFees0)) * myShare / totalShares / Math.pow(10, Number(this.pairsList.get(i).token0.decimals))
              this.userFee1Total += (Number(fees.U3Fees1) + Number(fees.LcFees1)) * myShare / totalShares / Math.pow(10, Number(this.pairsList.get(i).token1.decimals))
            }
          }
          // ufee0 = await keeperContract.methods.uFees0().call()
          // lfee0 = await keeperContract.methods.lFees0().call()
          // ufee1 = await keeperContract.methods.uFees1().call()
          // lfee1 = await keeperContract.methods.lFees1().call()
          // this.userFee0Total += (Number(ufee0) + Number(lfee0)) * myShare / totalShares / Math.pow(10, Number(this.pairsList.get(i).token0.decimals))
          // this.userFee1Total += (Number(ufee1) + Number(lfee1)) * myShare / totalShares / Math.pow(10, Number(this.pairsList.get(i).token1.decimals))
          var Assets = await keeperContract.methods.Assetholder(ethereum.selectedAddress).call()
          console.log('Assets=', JSON.stringify(Assets))
          if (Assets !== null) {
            // calc oraclePriceTwap
            var oraclePriceTwap
            var poolAddress = await keeperContract.methods.poolAddress().call()
            var twapDuration = 2
            var twap = await keeperContract.methods.getTwap(poolAddress, twapDuration).call()
            console.log('twap=', twap)
            if (twap !== null) {
              oraclePriceTwap = await keeperContract.methods.getQuoteAtTick(Number(twap), BigInt(Math.pow(10, 18)), this.pairsList.get(i).token0.tokenAddress, this.pairsList.get(i).token1.tokenAddress).call()
              console.log('oraclePriceTwap=', oraclePriceTwap)
            }
            // cacl myTVL
            var Total0, Total1
            var token0LendingContract = this.contractInstance(ViaLendTokenABI, this.pairsList.get(i).token0LendingAddress)
            var token1LendingContract = this.contractInstance(ViaLendTokenABI, this.pairsList.get(i).token1LendingAddress)
            var token0Contract = this.contractInstance(ViaLendTokenABI, this.pairsList.get(i).token0.tokenAddress)
            var token1Contract = this.contractInstance(ViaLendTokenABI, this.pairsList.get(i).token1.tokenAddress)

            var uniliqs = await keeperContract.methods.getPositionAmounts(BigInt(this.pairsList.get(i).cLow), BigInt(this.pairsList.get(i).cHigh)).call()
            console.log('balance in uniswap:', uniliqs, 'getVaultInfo_cLow=', this.pairsList.get(i).cLow, 'getVaultInfo_cHigh=', this.pairsList.get(i).cHigh)
            // -->Get Lending Amounts begin
            var exchangeRate0 = await token0LendingContract.methods.exchangeRateStored().call()
            var exchangeRate1 = await token1LendingContract.methods.exchangeRateStored().call()
            var CAmount0 = await token0LendingContract.methods.balanceOf(this.pairsList.get(i).vaultAddress).call()
            var CAmount1 = await token1LendingContract.methods.balanceOf(this.pairsList.get(i).vaultAddress).call()
            var underlying0 = CAmount0 * exchangeRate0 / Math.pow(10, 18)
            var underlying1 = CAmount1 * exchangeRate1 / Math.pow(10, 18)
            // console.log('underlying0=', underlying0, 'underlying1=', underlying1)
            // -->Get Lending Amounts end
            // var lendingAmounts = await keeperContract.methods.getLendingAmounts().call()
            // var cLending = await this.$parent.keeperContract.methods.getCAmounts().call()
            var balance0 = await token0Contract.methods.balanceOf(this.pairsList.get(i).vaultAddress).call()
            var balance1 = await token1Contract.methods.balanceOf(this.pairsList.get(i).vaultAddress).call()
            // console.log('token0BalanceInVault=', pair.token0BalanceInVault, ';token1BalanceInVault=', pair.token1BalanceInVault)
            // var t0Decimal = await token0Contract.methods.decimals().call()
            // var t1Decimal = await token1Contract.methods.decimals().call()
            Total0 = Number(balance0) + Number(uniliqs.amount0) + Number(underlying0)
            Total1 = Number(balance1) + Number(uniliqs.amount1) + Number(underlying1)
            var mytvl0, mytvl1
            if (Number(totalShares) === 0) {
              mytvl0 = 0
              mytvl1 = 0
            } else {
              mytvl0 = Total0 * Number(myShare) / Number(totalShares)
              mytvl1 = Total1 * Number(myShare) / Number(totalShares)
            }

            // calc APY
            var oneyearINsec = 365 * 24 * 60 * 60
            var block = await web3.eth.getBlock(Assets.block)
            console.log('block timestamp=', block.timestamp)
            var timestamp = block.timestamp
            var headerNumber = await web3.eth.getBlockNumber()
            var headerBlock = await web3.eth.getBlock(headerNumber)
            // console.log('header=', JSON.stringify(headerBlock))
            var htimestamp = headerBlock.timestamp
            console.log('htimestamp=', htimestamp, ';timestamp=', timestamp)
            var timediff = Number(htimestamp) - Number(timestamp)
            var deposit0 = Assets.deposit0
            var deposit1 = Assets.deposit1
            var fd0 = Number(deposit0)
            var fd1, fm1
            if (oraclePriceTwap === 0) {
              fd1 = 0
              fm1 = 0
            } else {
              fd1 = Number(deposit1) * Math.pow(10, Number(this.pairsList.get(i).token0.decimals)) / oraclePriceTwap
              fm1 = Number(mytvl1) * Math.pow(10, Number(this.pairsList.get(i).token0.decimals)) / oraclePriceTwap
            }
            var fm0 = Number(mytvl0)
            var fdd = fd0 + fd1
            var fmm = fm0 + fm1
            console.log('fm0=', fm0, 'fm1=', fm1)
            console.log('decimal=', this.pairsList.get(i).token0.decimals)
            console.log('fmm=', fmm, 'fdd=', fdd, 'timediff=', timediff, 'oneyearInsec=', oneyearINsec)
            if (fdd === 0 || timediff === 0) {
              this.pairsList.get(i).currentAPR = 0
            } else {
              this.pairsList.get(i).currentAPR = (fmm - fdd) / Number(timediff) * oneyearINsec / fdd
            }
            this.netAPYTotal += Number(this.pairsList.get(i).currentAPR)
            console.log('netAPYTotal=', this.netAPYTotal)
          }
        }
        if (this.pairsList.size() > 0) this.netAPYTotal = Number(this.netAPYTotal) / Number(this.pairsList.size())
        // userfee0 = (ufees0+lfees0) * user share/ totalsupply
        // userfee1 = (ufees1+lfees1) * user share/ totalsupply
        // console.log('netAPYTotal=', this.netAPYTotal)
        console.log('fee0Total=', this.fees0Total, ';fee1Total=', this.fees1Total)
        console.log('userFee0Total=', this.userFee0Total, ';userFee1Total=', this.userFee1Total)
        this.myValueLockLoading = false
      } else {
        this.myValueToken0Locked = 0
        this.myValueToken1Locked = 0
      }
    },
    getOraclePriceTwap () {

    },
    getAPY () {

    },
    showTreaty () {
      this.treatyDialogVisible = true
      // console.log(this.$store.state.name)
    },
    goSign () {
    },
    setWalletStatus () {
      this.$refs.headerComp.setWalletStatus()
    }
  }
}
</script>

<style scoped>
.el-loading-parent--relative {
  height: 100% !important;
}
.my-financial ul li {
  line-height: 30px;
}
.my-financial .title {
  font-size: 16px;
}
.my-financial .value {
  font-size: 20px;
  color: white;
}
.my-financial .apy-title {
  margin-top: 50px;
  font-size: 20px;
  color: white;
}
.my-financial .apy-content {
  font-size: 20px;
  color: white;
}
.apy-container {
  position: relative;
}
.apy-style {
  position: absolute;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  margin: auto;
  width: 200px;
  height: 200px;
  border: 5px solid #9900ff;
  border-radius: 50%;
}
.dialog-content {
  line-height: 25px;
}
.dialog-checkbox {
  line-height: 50px;
}
.dialog-footer {
  line-height: 50px;
}
.wallet_connected {
  width: 76px !important;
  padding-left: 10px;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 1;
  -webkit-box-orient: vertical;
  word-break: break-all;
}
.wallet_disconnected {
  width: 151px;
}
.walletButton {
  width: 151px;
  height: 37px;
  background: #2e3f61;
  text-align: center;
  border-radius: 6px;
}
.walletButton:hover {
  background: #637496;
}
.walletButton a {
  display: block;
  font-size: 16px;
  font-weight: 500;
  color: #ffffff;
  line-height: 37px;
}

.adminButton {
  width: 100px;
  height: 37px;
  background: #2a4988;
  text-align: center;
  border-radius: 6px;
}
.adminButton:hover {
  background: #637496;
}
.adminButton a {
  display: block;
  font-size: 16px;
  font-weight: 500;
  color: #ffffff;
  line-height: 37px;
}
.dialog_title {
  font-size: 26px;
  color: #000000;
  font-weight: bold;
}
.vault_title .el-card__header {
  padding: 0px !important;
}
.btn-view-etherscan {
  font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;
  font-size: 14px;
  color: #000000;
  background: #e6e6e6;
  border-color: #adadad;
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
  margin-top: -60px;
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

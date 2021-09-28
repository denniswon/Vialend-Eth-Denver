<template>
  <div class="myliquidity_box">
    <div class="container">
      <div class="row">
        <div class="col">
          <div class="myliquidity_box_container d-flex flex-row align-items-center justify-content-start">
            <div class="my_liquidity_container">
              <div class="my-liquidity-title">My Value Locked</div>
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
                  <th>Current Value</th>
                  <th>Earned Value</th>
                  <th>&nbsp;</th>
                </thead>
                <tbody>
                  <tr>
                    <td>
                      <img src="images/usdc.png"
                           width="30"
                           height="30" />
                      <img src="images/weth.png"
                           width="30"
                           height="30" />
                      <span>USDC / WETH</span>
                    </td>
                    <td><input class="btn btn-default btn-sm"
                             type="button"
                             value="0.30%"></td>
                    <td>505.66%</td>
                    <td>$33,500</td>
                    <td>$3,546.00</td>
                    <td>&nbsp;</td>
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
               width="350px"
               :close-on-click-modal="false"
               center>
      <span class="dialog_title"
            slot="title">{{newPositionDialogTitle}}</span>
      <!--Step1 content-->
      <div :style="{display:dialogStep1Display}">
        <el-select v-model="selectedPair"
                   placeholder="请选择"
                   size="medium"
                   style="width:300px;"
                   @change="selectedPairChanage">
          <el-option v-for="item in pairsData"
                     :key="item.id"
                     :label="item.smartVaults[0].name + '/' + item.smartVaults[1].name + '(' + item.feeTier + ')'"
                     :value="item.id">
            <span style="float: right">{{ item.smartVaults[0].name }} / {{ item.smartVaults[1].name }}({{item.feeTier}})</span>
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
                    width="50%">$111111111111111</td>
              </tr>
              <tr>
                <td class="c1">% of cap used</td>
                <td class="c2">99.8%</td>
              </tr>
              <tr>
                <td class="c1">Max TVL</td>
                <td class="c2">$3,000,000,00</td>
              </tr>
              <tr>
                <td class="c1">Range</td>
                <td class="c2">33485-69230</td>
              </tr>
              <tr>
                <td class="c1">Lending</td>
                <td class="c2">$2,390,141.00</td>
              </tr>
              <tr>
                <td class="c1">Current APR</td>
                <td class="c2">40.0%</td>
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
            <span class="token_balance">Balance:0(<el-link type="primary"
                       class="token_balance_max">Max</el-link>)</span>
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
            <span class="token_balance">Balance:0(<el-link type="primary"
                       class="token_balance_max">Max</el-link>)</span>
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
        <div class="action_list">
          <el-button type="info"
                     style="width:45%;"
                     @click="backToStep1">Back</el-button>
          <el-button type="next"
                     style="width:45%;float: right;"
                     @click="toStep3">Next</el-button>
        </div>
      </div>
      <!--Step3 content-->
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
                     @click="backToStep2">Back</el-button>
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
                     width="20" />&nbsp;&nbsp;{{this.pairsData[selectedPair-1].smartVaults[0].name}}</td>
              <td class="c2"
                  width="68%">59.1672<br>(10% in Uniswap 45% in Compound)</td>
            </tr>
            <tr>
              <td class="c1"><img :src="this.pairsData[selectedPair-1].smartVaults[1].iconLink"
                     width="20" />&nbsp;&nbsp;{{this.pairsData[selectedPair-1].smartVaults[1].name}}</td>
              <td class="c2">1500<br>(10% in Uniswap 45% in Compound)</td>
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

  </div>
</template>

<script>
import PairsData from '../data/PairsData'

export default {
  name: 'MyPositions',
  data () {
    return {
      selectedPair: 1,
      selectedText: '',
      pairsData: PairsData,
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
      newLiqudityToken0: 0,
      newLiqudityToken1: 0,
      signAgreement: false
    }
  },
  methods: {
    showNewPositionDialog () {
      this.newPositionDialogTitle = this.dialogStep1Title
      this.newPositionDialogVisible = true
      this.dialogStep1Display = ''
      this.dialogStep2Display = 'none'
      this.dialogStep3Display = 'none'
      this.dialogResultDisplay = 'none'
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
    },
    backToStep2 () {
      this.newPositionDialogTitle = this.dialogStep2Title
      this.dialogStep1Display = 'none'
      this.dialogStep2Display = ''
      this.dialogStep3Display = 'none'
      this.dialogResultDisplay = 'none'
    },
    toStep3 () {
      this.newPositionDialogTitle = this.dialogStep3Title
      this.dialogStep1Display = 'none'
      this.dialogStep2Display = 'none'
      this.dialogStep3Display = ''
      this.dialogResultDisplay = 'none'
    },
    toResult () {
      this.newPositionDialogTitle = this.dialogResultTitle
      this.dialogStep1Display = 'none'
      this.dialogStep2Display = 'none'
      this.dialogStep3Display = 'none'
      this.dialogResultDisplay = ''
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
</style>

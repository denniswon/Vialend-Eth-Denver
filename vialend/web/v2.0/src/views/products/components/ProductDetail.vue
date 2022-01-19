<template>
  <div class="pairs-container">
    <div class="pair-intro">
      <div class="pair-title-left">
        <div slot="header" class="clearfix card-title">
          <div class="block">
            <el-tag size="small" hit>UNI-COMP</el-tag>&nbsp;&nbsp;
            <el-tag class="tag-item" type="info" effect="plain" size="small">V1</el-tag>
          </div>
        </div>
        <div class="card-container">
          <div class="card-pairinfo">
            <div class="pair-name text item">
              <span>{{ pairInfo.token0.symbol }} - {{ pairInfo.token1.symbol }}</span>
            </div>
            <div class="text item">
              <span class="pair-current-deposits">Current Deposits</span>
              <span class="pair-deposits-val">
                {{ Number(pairInfo.currentDeposits / Math.pow(10, Number(pairInfo.token1.decimals))).toExponential(5) }}&nbsp;{{
                pairInfo.token1.symbol
                }}
              </span>
            </div>
            <div style="clear:both;"></div>
            <div class="text item">
              <el-slider v-model="maxCapacity" :show-tooltip="false"></el-slider>
            </div>
            <div style="clear:both;"></div>
            <div class="text item">
              <span class="pair-max-capacity">Max Capacity</span>
              <span class="pair-capacity-val">100M&nbsp;{{ pairInfo.token1.symbol }}</span>
            </div>
          </div>
        </div>
      </div>
      <div class="pair-title-right">
        <div class="token-icon">
          <cryptoicon :symbol="pairInfo.token1.symbol" size="260" />
        </div>
      </div>
    </div>
    <hr />
    <div class="vault-info">
      <div class="vault-title">Vault Strategy</div>
      <div class="vault-info-left">
        <div
          class="vault-intro"
        >Lorem ipsum dolor sit amet,consectetur adipiscing elit.Praesent tellus leo,accumsan ut enim sed,aliquam aliquam nibh.In ultrices ex eget odio facilisis,vitae laoreet dui efficitur.Nam fringilla enim ac tincidunt facilisis.Nam sagittis mi ligula,in posuere arcu interdum in.Quisque sodales eros elit, a ultrices sapien dapibus vitae.Donec tempor,magna eget pharetra sodales,nisi libero</div>
        <div class="slidershow">
          <el-carousel :interval="4000" type="card" height="200px">
            <el-carousel-item v-for="item in 6" :key="item">
              <h3 class="medium">{{ item }}</h3>
            </el-carousel-item>
          </el-carousel>
        </div>
      </div>
      <div class="vault-info-right">
        <el-tabs type="border-card" :stretch="true">
          <el-tab-pane label="Deposit">
            <div class="step_token">
                <span class="token_balance">
                  Amount:({{pairInfo.token0.symbol}})
                </span>
                <span class="token_textbox">
                  <el-input
                    v-model="newLiqudityToken0"
                    placeholder="0.00"
                    type="text"
                    class="tokeninput"
                    onkeypress="return /^[0-9]*[.,]?[0-9]*$/.test(this.value.concat(event.key))"
                  >
                  <i slot="suffix" class="el-input__icon"  @click="newLiqudityToken0 = pairInfo.token0.balanceInWallet">Max</i>
                  </el-input>
                </span>
            </div>

            <div class="step_token">
                <span class="token_balance">
                  Amount:({{pairInfo.token1.symbol}})
                </span>
                <span class="token_textbox">
                  <el-input
                    v-model="newLiqudityToken1"
                    placeholder="0.00"
                    type="text"
                    class="tokeninput"
                    onkeypress="return /^[0-9]*[.,]?[0-9]*$/.test(this.value.concat(event.key))"
                  >
                  <i slot="suffix" class="el-input__icon"  @click="newLiqudityToken1 = pairInfo.token1.balanceInWallet">Max</i>
                  </el-input>
                </span>
            </div>
            <div class="action_list" :style="{display: (pairInfo.token0.tokenApproved ? 'none':'')}">
              <el-button
                type="primary"
                :loading="approve0Loading"
                @click="approveToken(pairInfo.token0)"
                :style="{width: '100%'}"
              >Approve {{pairInfo.token0.symbol}}</el-button>
            </div>
            <div class="action_list" :style="{display: (pairInfo.token1.tokenApproved ? 'none':'')}">
              <el-button
                type="primary"
                :loading="approve1Loading"
                @click="approveToken(pairInfo.token1)"
                :style="{width: '100%'}"
              >Approve {{pairInfo.token1.symbol}}</el-button>
            </div>
            <div class="action_list deposit_class">
              <el-button type="primary"
                       @click="deposit"
                       :loading="depositLoading"
                       :disabled="(pairInfo.token0.tokenApproved || pairInfo.token1.tokenApproved) ? false:true"
                       :style="{width: '100%',display: $store.state.isConnected?'':'none'}">Deposit</el-button>
            </div>
            <div class="action_list deposit_class" :style="{display: $store.state.isConnected?'none':''}">
              <el-button
                type="primary"
                style="width:100%;"
                @click="connectWallet"
              >Connect Wallet</el-button>
            </div>
            <div class="contract_address">
                CONTRACT:{{walletAddress}}
                <a :href="'https://goerli.etherscan.io/address/'+pairInfo.vaultAddress" style="cursor:hand" target="_black"><svg-icon
                    name="goto-detail"
                    width="30"
                    height="30"
                /></a>
                </div>
          </el-tab-pane>
          <el-tab-pane label="Withdraw">
              <table class="withdraw-table">
          <tr>
            <td><span class="vault_title">Vault shares:</span>{{shareValue}}
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
            </td>
          </tr>
          <tr>
            <td class="withdraw_class">
              <el-button type="primary"
                         :disabled="btnWithdrawDisabled"
                         @click="withdraw"
                         :loading="withdrawLoading"
                         style="width:100%;">Withdraw</el-button>
            </td>
          </tr>
          <tr>
              <td>
              <div class="contract_address">
                CONTRACT:{{walletAddress}}
                <a :href="'https://goerli.etherscan.io/address/'+pairInfo.vaultAddress" style="cursor:hand" target="_black"><svg-icon
                    name="goto-detail"
                    width="30"
                    height="30"
                /></a>
                </div>
              </td>
              </tr>
        </table>
          </el-tab-pane>
        </el-tabs>
      </div>
    </div>
    <div style="clear:both;"></div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator'
import { contractInstance } from '../../../common/Web3'
import Pairs from '../../../model/Pairs'
import Token from '../../../model/Token'
import { formatWalletAddress } from '../../../common/Tools'

const ViaLendTokenABI = require('../../../abi/VialendTokenABI.json')
const ViaLendFeeMakerABI = require('../../../abi/ViaLendFeeMakerABI.json')
const ethereum = (window as any).ethereum

@Component({
  name: 'ProductDetail',
  components: { }
})
export default class extends Vue {
  // @Prop({ default: false }) private isEdit!: boolean

  private pairInfo = null
  private maxCapacity = 60
  newLiqudityToken0: null | number = null
  newLiqudityToken1: null | number = null
  withdrawLoading = false
  approve0Loading = false
  approve1Loading = false
  depositLoading = false
  btnDepositDisabled = true
  goToEtherscan = ''
  depositedEtherscanDisable = false
  shareValue = 0n
  sharePercent = 25
  btnWithdrawDisabled = false
  withdrawedEtherscanDisable = true

  get walletAddress() {
    return formatWalletAddress((this.pairInfo as any).vaultAddress)
  }

  connectWallet() {
    // this.$parent.connectWallet()
    console.log('wallet connection status:', this.$store.state.isConnected)
  }

  async updatePageData() {
    //   await this.getPairPublicInfo(this.pairsList.get(this.selectedPairIndex))
    //   await this.loadMyData()
  }

  async approveToken(token:Token) {
    // eslint-disable-next-line @typescript-eslint/no-this-alias
    const _this = this
    const tokenContract = await contractInstance(ViaLendTokenABI, token.tokenAddress)
    const tokenName = token.name
    const tokenAddress = token.tokenAddress
    const vaultAddress = (this.pairInfo as any).vaultAddress
    token.tokenApproved = false
    console.log('token.tokenAddress=', token.tokenAddress)
    console.log('vaultAddress=', (this.pairInfo as any).vaultAddress)

    if (tokenContract != null) {
      token.approveLoading = true
      tokenContract.methods
        .approve(vaultAddress, BigInt(90000000000000000000000))
        .send({
          from: ethereum.selectedAddress,
          //   gasPrice: '10000000000',
          //   gas: 200000,
          value: 0
        })
        .on('confirmation', function(confirmationNumber:number, receipt:any) {
          console.log((new Date()).toLocaleString(), ':{approve confirm number:', confirmationNumber, ',receipt:', receipt.status, '}')
        })
        .on('receipt', (receipt:any) => {
          if (receipt.status) {
            _this.$store.dispatch('setApproveStatus', {
              token: tokenAddress, status: true
            })
            token.tokenApproved = true
            if (token.approveLoading) {
              token.approveLoading = false
              _this.$message(tokenName + ' approved!')
            }
            if ((this.pairInfo as any).token0.tokenApproved || (this.pairInfo as any).token1.tokenApproved) {
              _this.btnDepositDisabled = false
            }
          } else {
            _this.$message(tokenName.concat(' Approve failed!'))
          }
        })
        .on('error', function(error:any) {
          token.tokenApproved = false
          token.approveLoading = false
          _this.$message.error(error.message === undefined ? error : error.message)
          console.log('approve error=', error)
        })
    }
  }

  async deposit() {
    const pair:Pairs = this.pairInfo as any
    if (this.newLiqudityToken0 === null || this.newLiqudityToken1 === null || isNaN(this.newLiqudityToken0) || isNaN(this.newLiqudityToken1) || Number(this.newLiqudityToken0) < 0 || Number(this.newLiqudityToken1) < 0 || (Number(this.newLiqudityToken0) === 0 && Number(this.newLiqudityToken1) === 0)) {
      this.$message('Please enter a positive number greater than 0!')
      return
    } else if (Number(this.newLiqudityToken0) > 0 && !pair.token0.tokenApproved) {
      this.$message('Please approve token0 first!')
      return
    } else if (Number(this.newLiqudityToken1) > 0 && !pair.token1.tokenApproved) {
      this.$message('Please approve token1 first!')
      return
    }
    let token0Amount = 0; let token1Amount = 0
    if (this.newLiqudityToken0 === null) { token0Amount = 0 } else { token0Amount = this.newLiqudityToken0 }
    if (this.newLiqudityToken1 === null) { token1Amount = 0 } else { token1Amount = this.newLiqudityToken1 }
    // eslint-disable-next-line @typescript-eslint/no-this-alias
    const _this = this
    const keeperContract = await contractInstance(ViaLendFeeMakerABI, pair.vaultAddress)
    // this.depositToken0 = this.newLiqudityToken0
    // this.depositToken1 = this.newLiqudityToken1
    if (keeperContract != null) {
      this.depositLoading = true
      keeperContract.methods
        .deposit(
          BigInt(token0Amount * Math.pow(10, pair.token0.decimals)),
          BigInt(token1Amount * Math.pow(10, pair.token1.decimals)),
          true
        )
        .send({
          from: ethereum.selectedAddress,
          // gasPrice: '80000000000',
          // gas: 600000,
          value: 0
        })
        .on('confirmation', function(confirmationNumber:number, receipt:any) {
          console.log((new Date()).toLocaleString(), ':{deposit confirm number:', confirmationNumber, ',receipt:', receipt.status, '}')
        })
        .on('receipt', function(receipt:any) {
          console.log('receipt:', receipt)

          if (_this.depositLoading === true) {
            _this.depositLoading = false
            if (receipt.status) {
              _this.newLiqudityToken0 = null
              _this.newLiqudityToken1 = null
              _this.goToEtherscan = 'https://goerli.etherscan.io/tx/' + receipt.transactionHash
              _this.depositedEtherscanDisable = false
              _this.$message('Deposit submitted!')
              _this.updatePageData()
            } else {
              _this.$message('Deposit failed!')
            }
          }
          console.log((new Date()).toLocaleString(), 'deposit receipt status:', receipt.status)
        })
        .on('error', function(error:any) {
          _this.$message.error(error)
          _this.depositLoading = false
          console.log('Deposit error:', error)
        })
    }
  }

  withdraw() {
    const pair:Pairs = this.pairInfo as any
    // eslint-disable-next-line @typescript-eslint/no-this-alias
    const _this = this
    const keeperContract = contractInstance(ViaLendFeeMakerABI, pair.vaultAddress)
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
        .on('confirmation', function(confirmationNumber:number, receipt:any) {
          console.log((new Date()).toLocaleString(), ':{withdraw confirm number:', confirmationNumber, ',receipt:', receipt.status, '}')
        })
        .on('receipt', function(receipt:any) {
          if (_this.withdrawLoading === true) {
            _this.withdrawLoading = false
            if (receipt.status) {
              _this.goToEtherscan = 'https://goerli.etherscan.io/tx/' + receipt.transactionHash
              _this.withdrawedEtherscanDisable = false
              _this.getShares(_this.sharePercent)
              _this.$message('Withdraw submitted!')
              _this.updatePageData()
            } else {
              _this.$message('Withdraw failed!')
            }
            console.log((new Date()).toLocaleString(), 'withdraw receipt status:', receipt.status)
          }
        })
        .on('error', function(error:any) {
          _this.$message.error(error)
          _this.withdrawLoading = false
          console.log('Withdraw error:', error)
        })
    }
  }

  @Watch('sharePercent')
  watchSharePercent(newVal:number) {
    this.getShares(newVal)
  }

  setSharePercent(percent:number) {
    this.sharePercent = percent
  }

  getShares(percent:number) {
    this.shareValue = BigInt(parseInt(((this.pairInfo as any).myShare * percent / 100).toString()))
    console.log('myShare=', (this.pairInfo as any).myShare, 'shares=', this.shareValue)
    if (this.shareValue > 0) this.btnWithdrawDisabled = false
    else this.btnWithdrawDisabled = true
  }

  created() {
    const pairsymbol = this.$route.params && this.$route.params.pair
    const jsonPairInfo: string | null = sessionStorage.getItem(pairsymbol)

    if (jsonPairInfo !== null) {
      this.pairInfo = JSON.parse(jsonPairInfo)
      if (this.pairInfo !== null) {
        this.getShares(this.sharePercent)
        console.log('pair vaultAddress=', (this.pairInfo as any).vaultAddress)
        console.log('pairInfo.token0BalanceInWallet=', (this.pairInfo as any).token0.balanceInWallet)
      } else {
        console.log('pairinfo is null')
      }
    }
  }
}
</script>

<style lang="scss">
</style>

<style scoped>
.pairs-container {
  padding: 30px;
  margin: 0 auto;
}
.pair-title-left {
  width: 60%;
  display: inline-block;
}
.pair-title-right {
  width: 40%;
  display: inline-block;
  vertical-align: top;
  text-align: center;
}
.vault-info-left {
    float: left;
  width: 60%;
}
.vault-info-right {
    float:right;
  width: 30%;
  height:100%;
}
.el-input /deep/ .el-input__inner{
  height:50px !important;
  font-size: 26px !important;
}
.el-input__icon{
    cursor: pointer;
    margin:5px;
}
.el-tabs--card{
    height:360px;
}
.el-tab-pane{
    min-height:310px;
    overflow-x: hidden;
    overflow-y:hidden;
}
.vault-title {
  width: 100%;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  font-size: 20px;
  font-weight: bold;
}
.vault-intro{
    font-family: 'Times New Roman', Times, serif;
    font-size:16px;
    padding-top:10px;
    padding-bottom: 10px;
    line-height: 20px;
}
.card-panel-col {
  height: 515px;
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
.block {
  float: left;
}
.card-container {
  width: 100%;
  margin: 0px;
  line-height: 22px;
}
.card-pairinfo {
  height: 300px;
  width: 100%;
  margin: 0;
  padding: 38px 20px 10px 20px;
  border-radius: 16px;
}
.slidershow{
    padding-top: 20px;
}
.f12 {
  color: #000000;
  font-family: sans-serif;
  font-size: 12px;
}
.f24 {
  color: #000000;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  font-size: 24px;
  text-transform: uppercase;
}
.token-icon {
  margin-left: 26px;
}
.pair-name {
  font-size: 20px;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  padding-bottom: 20px;
}
.pair-intro {
  color: #000000;
  font-family: Inter, sans-serif;
  font-size: 14px;
}
.pair-apy {
  color: #000000;
  font-family: Inter, sans-serif;
  padding-bottom: 10px;
}
.pair-current-deposits,
.pair-max-capacity {
  float: left;
  color: #000000;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  font-size: 12px;
}
.pair-deposits-val,
.pair-capacity-val {
  float: right;
  color: #000000;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  font-size: 12px;
  font-weight: bold;
}
.svg-icon {
  vertical-align: -8px !important;
}
.el-carousel__item h3 {
  color: #475669;
  font-size: 14px;
  opacity: 0.75;
  line-height: 200px;
  margin: 0;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}

.el-input__inner {
  height:50px !important;
  font-size: 26px !important;
}

.step_token{
    padding:5px;
    margin-top: 10px;
}
.step_token .token_balance{
    display: block;
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
    font-size: 14px;
    font-weight: bold;
    padding-bottom:10px;
    margin:0px;
}
.action_list{
    padding-top: 10px;
}
.vault_title{
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
    font-size: 14px;
    font-weight: bold;
}
.deposit_class /deep/ .el-button,
.withdraw_class /deep/ .el-button{
    width: 100%;
    height: 50px !important;
    font-size:21px !important;
}
.withdraw-table {
    width:100%;
  margin-top: 20px;
  border-spacing: 0px !important;
  padding:0px !important;
}
.withdraw-table td,
.withdraw-table th {
  border: 0px solid transparent;
  font-size: 14px;
  color: #000000;
  text-align: left;
  height: 45px;
}

.contract_address{
    font-family: monospace,'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
    border: 1px solid #475669;
    background-color: #2e3f61;
    color:#ffffff;
    font-size:16px;
    margin-top:10px;
    padding-left: 10px;
    border-radius: 4px;
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

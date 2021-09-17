<template>
  <div class="supply_liquidity">
    <div class="container">
      <div class="row">
        <div class="col">
          <div class="supply_liquidity_container d-flex flex-row align-items-center justify-content-start">
            <div class="supply_liquidity_form_container">
              <div class="my-liquidity-title">Supply Liquidity</div>
              <table class="table">
                <thead>
                  <th>Smart Vaults</th>
                  <th>Fee Tier</th>
                  <th>Current APR</th>
                  <th>Capacity</th>
                  <th>TVL</th>
                  <th>&nbsp;</th>
                </thead>
                <tbody>
                  <tr v-for="dt in supplyLiquidityList"
                      :key="dt.number">
                    <td>
                      <img :src="dt.smartVaults[0].iconLink"
                           width="30"
                           height="30" />&nbsp;
                      <img :src="dt.smartVaults[1].iconLink"
                           width="30"
                           height="30" />&nbsp;
                      <span>{{dt.smartVaults[0].name}} / {{dt.smartVaults[1].name}}</span>
                    </td>
                    <td><input class="btn btn-default btn-sm"
                             type="button"
                             value="0.30%"></td>
                    <td>505.66%</td>
                    <td><button type="button"
                              class="btn btn-default btn-sm">35.1%</button></td>
                    <td>$2,366,149</td>
                    <td><button type="button"
                              class="btn btn-primary btn-sm"
                              @click="showSupplyDialog(dt)"
                              :disabled="dt.disabled">Supply</button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
    <el-dialog :title="supplyDialogTitle"
               :visible.sync="supplyDialogVisible"
               :append-to-body="true"
               width="380px"
               center>
      <span>To supply or withdraw liquidity for {{supplyTokens}} to the SwapStream protocol,you need to enable it first.</span>
      <span>
        <el-tabs v-model="activeName"
                 :stretch="true"
                 @tab-click="handleTabClick">
          <el-tab-pane label="Deposit"
                       name="first">
            <table class="table">
              <tr>
                <td><img src="images/usdc.png"
                       width="30px"
                       height="30px" /></td>
                <td>
                  <span class="lblBalance">Balance:{{token0Balance}}(<el-link type="primary"
                             @click="depositToken0 = token0Balance">Max</el-link>)</span>
                  <el-input placeholder="0.0"
                            type="text"
                            v-model="depositToken0"
                            dir="rtl"
                            style="text-align:right;"
                            pattern="^[0-9]*[.,]?[0-9]*$"
                            autocorrect="off"
                            autocomplete="off"
                            inputmode="decimal"></el-input>
                </td>
              </tr>
              <tr>
                <td><img src="images/weth.png"
                       width="30px"
                       height="30px" /></td>
                <td>
                  <span class="lblBalance">Balance:{{token1Balance}}(<el-link type="primary"
                             @click="depositToken1 = token1Balance">Max</el-link>)</span>
                  <el-input placeholder="0.0"
                            type="text"
                            v-model="depositToken1"
                            :model="depositToken1"
                            dir="rtl"
                            style="text-align:right;"
                            pattern="^[0-9]*[.,]?[0-9]*$"
                            autocorrect="off"
                            autocomplete="off"
                            inputmode="decimal"></el-input>
                </td>
              </tr>
              <tr>
                <td colspan="2"
                    style="text-align:center;">
                  <el-button type="primary"
                             :disabled="token0Approved == true ? true : false"
                             @click="approveToken(0)"
                             :style="{display:isConnected?'':'none'}">Approve {{supplyToken0}}</el-button>
                  <el-button type="primary"
                             :disabled="token1Approved == true ? true : false"
                             @click="approveToken(1)"
                             :style="{display:isConnected?'':'none'}">Approve {{supplyToken1}}</el-button>
                  <el-button type="primary"
                             style="width:100%;"
                             @click="connectWallet"
                             :style="{display:isConnected?'none':''}">Connect Wallet</el-button>

                </td>
              </tr>
              <tr>
                <td colspan="2">
                  <el-button type="primary"
                             :disabled="btnDepositDisabled"
                             @click="deposit"
                             :loading="depositLoading"
                             style="width:100%;"
                             :style="{display:isConnected?'':'none'}">Deposit</el-button>
                </td>
              </tr>
            </table>
          </el-tab-pane>
          <el-tab-pane label="Withdraw"
                       name="second">
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
                             :disabled="btnWithdrawDisabled"
                             @click="withdraw"
                             :loading="withdrawLoading"
                             style="width:100%;">Withdraw</el-button>
                </td>
              </tr>
            </table>
          </el-tab-pane>

        </el-tabs>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import Web3 from 'web3'
import contractABI from '../ABI/contractABI.json'
import tokenABI from '../ABI/tokenABI.json'
import SupplyLiquidityList from '../data/SupplyLiqudityList'

if (typeof web3 !== 'undefined') {
  web3 = new Web3(web3.currentProvider)
  console.log('web3 provider:web3.currentProvider')
} else {
  // set the provider you want from Web3.providers
  web3 = new Web3(new Web3.providers.HttpProvider('https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6'))
  console.log('web3 provider:goerli')
}

export default {
  name: 'SupplyLiquidity',
  data () {
    return {
      supplyLiquidityList: SupplyLiquidityList,
      isConnected: false,
      activeName: 'first',
      supplyDialogTitle: '',
      supplyTokens: '',
      supplyToken0: '',
      supplyToken1: '',
      supplyDialogVisible: false,
      depositToken0: 0,
      depositToken1: 0,
      token0Approved: false,
      token1Approved: false,
      btnDepositDisabled: false,
      btnWithdrawDisabled: false,
      testData: '123',
      // contractObj: new web3.eth.Contract(contractABI, contractAddress)
      token0Contract: null,
      token1Contract: null,
      tokenContract: null,
      currentItem: null,
      depositLoading: false,
      withdrawLoading: false,
      shareValue: 0,
      sharePercent: 25,
      token0Balance: 1,
      token1Balance: 2,
      token0Decimals: 0,
      token1Decimals: 0,
      myAccount: '',
      ashares: 0
    }
  },
  created: async function () {
    var decimals
    for (var i = 0; i < this.supplyLiquidityList.length; i++) {
      decimals = await this.getDecimalOfToken(this.supplyLiquidityList[i].smartVaults[0].tokenAddress)
      this.supplyLiquidityList[i].smartVaults[0].decimals = decimals
      // console.log('decimal0', i, '=', decimals)
      decimals = await this.getDecimalOfToken(this.supplyLiquidityList[i].smartVaults[1].tokenAddress)
      this.supplyLiquidityList[i].smartVaults[1].decimals = decimals
      // console.log('decimal1', i, '=', decimals)
    }

    // console.log('SupplyLiquidityList=', JSON.stringify(SupplyLiquidityList))
  },
  mounted: function () {

  },
  watch: {
    sharePercent (val) {
      this.getShares(val)
    }
  },
  methods: {
    getTokenContract (index) {
      this.tokenContract = new web3.eth.Contract(
        contractABI,
        this.currentItem.smartVaults[index].tokenAddress
      )
      console.log('token contract address is:' + this.currentItem.smartVaults[index].tokenAddress)
    },
    connectWallet () {
      this.$parent.setWalletStatus()
      this.isConnected = this.$parent.getConnectionStatus()
      console.log('wallet connection status:', this.isConnected)
    },
    checkConnectionStatus () {
      if (ethereum.isConnected() && this.$parent.isConnected) {
        console.log('wallet is connected')
        this.isConnected = true
      } else {
        console.log('wallet is disconnected')
        this.isConnected = false
      }
    },
    showSupplyDialog (item) {
      var _this = this
      this.checkConnectionStatus()
      this.getAShares()
      this.getShares(this.sharePercent)
      this.currentItem = item
      this.supplyDialogTitle = item.smartVaults[0].name + ' / ' + item.smartVaults[1].name
      this.supplyToken0 = item.smartVaults[0].name
      this.supplyToken1 = item.smartVaults[1].name
      this.token0Decimals = item.smartVaults[0].decimals
      this.token1Decimals = item.smartVaults[1].decimals
      console.log('token0Decimals=', this.token0Decimals)
      console.log('token1Decimals=', this.token1Decimals)
      // Get approve status for token0
      // this.$store.dispatch('setApproveStatus', {
      //   'token': item.smartVaults[0].name, 'status': false
      // })
      // this.token0Approved = this.$store.commit('getApproveStatus', {
      //   'token': 'TestToken'
      // })

      this.$store.dispatch('getApproveStatus', { 'token': item.smartVaults[0].name }).then(res => {
        console.log(item.smartVaults[0].name + 'Approved status is:' + JSON.stringify(res))
        if (res === 'true') {
          this.token0Approved = true
          console.log('this.btnToken0ApproveDisabled = true ')
        } else {
          this.token0Approved = false
          console.log('this.btnToken0ApproveDisabled = false ')
        }
        this.enableDepositFeature()
      })
      this.$store.dispatch('getApproveStatus', { 'token': item.smartVaults[1].name }).then(res => {
        console.log(item.smartVaults[1].name + 'Approved status is:' + JSON.stringify(res))
        if (res === 'true') {
          this.token1Approved = true
          console.log('this.btnToken1ApproveDisabled = true ')
        } else {
          this.token1Approved = false
          console.log('this.btnToken1ApproveDisabled = false ')
        }
        this.enableDepositFeature()
      })
      // Get the balance of the account of given address
      this.getTokensBalanceInWallet(item)
      // console.log('addr=', this.currentItem.smartVaults[0].tokenAddress)
      // web3.eth.getBalance('0xFA5dF5372c03D4968d128D624e3Afeb61031a777').then(function (balance) {
      //   // _this.token0Balance = balance
      //   console.log('balance=', balance)
      // })

      // web3.eth.getBalance(ethereum.selectedAddress).then(console.log)

      // coinContract.methods.balanceOf('0xFdA9705FdB20E9A633D4283AfbFB4a0518418Af8').call({ from: '0x6dd19aEB91d1f43C46f0DD74C9E8A92BFe2a3Cd0' }, function (error, result) {
      //   if (!error) {
      //     console.log('balance result=', result)
      //   } else {
      //     console.log(error)
      //   }
      // })
      // // 查看某个账号的代币余额
      // coinContract.methods.balanceOf('0xFA5dF5372c03D4968d128D624e3Afeb61031a777').call()
      //   .then(val => {
      //     console.log('Coin BalanceOf=' + val)
      //   })

      // coinContract.methods.name().call()
      //   .then(val => {
      //     console.log('name=' + val)
      //   })

      this.supplyDialogVisible = true
    },
    async getDecimalOfToken (address) {
      var coinContract
      coinContract = new web3.eth.Contract(
        tokenABI,
        address)
      var decimal = await coinContract.methods.decimals().call()
      // console.log('decimal_', address, ' = ', decimal)
      return decimal
    },
    async getTokensBalanceInWallet (item) {
      var coinContract
      var balanceWei
      var decimals
      // token0 balance in wallet
      coinContract = new web3.eth.Contract(
        tokenABI,
        item.smartVaults[0].tokenAddress)
      balanceWei = await coinContract.methods.balanceOf(ethereum.selectedAddress).call() // 29803630997051883414242659
      // decimal = await this.getDecimalOfToken(item.smartVaults[0].tokenAddress)
      decimals = item.smartVaults[0].decimals
      console.log('token0BalanceInWallet=', balanceWei, ';token0 decimal=', decimals)
      this.token0Balance = (parseInt(balanceWei / (Math.pow(10, decimals)) * 1000) / 1000).toFixed(3)

      // token1 balance in wallet
      decimals = item.smartVaults[1].decimals
      coinContract = new web3.eth.Contract(
        tokenABI,
        item.smartVaults[1].tokenAddress)
      balanceWei = await coinContract.methods.balanceOf(ethereum.selectedAddress).call() // 29803630997051883414242659
      // this.token1Balance = (parseInt(web3.utils.fromWei(balanceWei) * 1000) / 1000).toFixed(3)
      this.token1Balance = (parseInt(balanceWei / (Math.pow(10, decimals)) * 1000) / 1000).toFixed(3)
      console.log('token1BalanceInWallet=', this.token1Balance, ';token1 decimal=', decimals)
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
      this.getTokenContract(index)
      if (index === 0) {
        this.token0Approved = true
      } else {
        this.token1Approved = true
      }
      console.log('vaultAddress=' + this.$parent.vaultAddress)
      if (this.tokenContract != null) {
        this.tokenContract.methods
          .approve(this.$parent.vaultAddress, BigInt(90000000000000000000000))
          .send({
            from: ethereum.selectedAddress,
            gasPrice: '10000000000',
            gas: 200000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            _this.$store.dispatch('setApproveStatus', {
              'token': _this.currentItem.smartVaults[index].name, 'status': true
            })
            if (_this.token0Approved && _this.token1Approved) {
              _this.btnDepositDisabled = false
            }
            console.log('token' + index + ' confirmation')
          })
          .on('receipt', function (receipt) {
            // $('#depositresult').text('Successfully deposited!');
            _this.$store.dispatch('setApproveStatus', {
              'token': _this.currentItem.smartVaults[index].name, 'status': true
            })
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
            _this.$message.error('token' + index + ' error' + error)
            console.log('token' + index + ' error' + error)
          })
      }
    },
    deposit () {
      var _this = this
      // console.log('t0', BigInt(this.depositToken0 * Math.pow(10, this.token0Decimals)))
      // console.log('t1', BigInt(this.depositToken1 * Math.pow(10, this.token1Decimals)))
      // return
      if (this.$parent.keeperContract != null) {
        console.log('account address is ' + ethereum.selectedAddress)
        this.depositLoading = true
        this.$parent.keeperContract.methods
          .deposit(
            BigInt(this.depositToken0 * Math.pow(10, this.token0Decimals)),
            BigInt(this.depositToken1 * Math.pow(10, this.token1Decimals)),
            ethereum.selectedAddress
          )
          .send({
            from: ethereum.selectedAddress,
            gasPrice: '10000000000',
            gas: 200000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            if (_this.depositLoading === true) {
              _this.depositLoading = false
              _this.$message('Successfully deposited!')
              _this.supplyDialogVisible = false
              console.log('confirmation')
            }
          })
          .on('receipt', function (receipt) {
            if (_this.depositLoading === true) {
              _this.depositLoading = false
              _this.$message('Successfully deposited!')
              _this.supplyDialogVisible = false
              console.log('receipt')
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
            BigInt(this.shareValue),
            BigInt(0),
            BigInt(0),
            ethereum.selectedAddress
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
              _this.supplyDialogVisible = false
              console.log('confirmation')
            }
          })
          .on('receipt', function (receipt) {
            if (_this.withdrawLoading === true) {
              _this.withdrawLoading = false
              _this.$message('Successfully withdrawed!')
              _this.supplyDialogVisible = false
              console.log('receipt')
            }
          })
          .on('error', function (error) {
            _this.$message.error(error)
            _this.withdrawLoading = false
          })
      }
    },
    async getShares (percent) {
      if (this.myAccount !== ethereum.selectedAddress) {
        this.myAccount = ethereum.selectedAddress
        this.getAShares()
      } else if (this.ashare === 0) {
        this.getAShares()
      }
      this.shareValue = BigInt(parseInt(this.ashares * percent / 100))
      console.log('shares=', this.shareValue)
    },
    async getAShares () {
      if (this.$parent.keeperContract != null) {
        this.ashares = await this.$parent.keeperContract.methods.balanceOf(ethereum.selectedAddress).call()
        console.log('ashares1=', this.ashares)
      }
    },
    setSharePercent (percent) {
      this.sharePercent = percent
    },
    handleTabClick (tab, event) {
      console.log(tab, event)
    },
    clearInputInitData (obj) {
      console.log('obj value=' + obj)
      if (obj === 0) {
        // obj = ''
      }
    }
  }
}
</script>

<style scoped>
.table td,
.table th {
  border: 0px solid transparent;
}
.lblBalance {
  float: right;
}
.share_percent {
  font-size: 30px;
  margin: 10px;
}
</style>

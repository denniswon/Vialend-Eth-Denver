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
               width="360px"
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
                  <el-input placeholder="0.0"
                            type="text"
                            v-model="depositToken0"
                            style="text-align:right;"></el-input>
                </td>
              </tr>
              <tr>
                <td><img src="images/weth.png"
                       width="30px"
                       height="30px" /></td>
                <td>
                  <el-input placeholder="0.0"
                            type="text"
                            v-model="depositToken1"
                            :model="depositToken1"></el-input>
                </td>
              </tr>
              <tr>
                <td colspan="2"
                    style="text-align:center;">
                  <el-button type="primary"
                             :disabled="token0Approved == true ? true : false"
                             @click="approveToken(0)">Approve {{supplyToken0}}</el-button>
                  <el-button type="primary"
                             :disabled="token1Approved == true ? true : false"
                             @click="approveToken(1)">Approve {{supplyToken1}}</el-button>

                </td>
              </tr>
              <tr>
                <td colspan="2">
                  <el-button type="primary"
                             :disabled="btnDepositDisabled"
                             @click="deposit"
                             :loading="depositLoading"
                             style="width:100%;">Deposit</el-button>
                </td>
              </tr>
            </table>
          </el-tab-pane>
          <el-tab-pane label="Withdraw"
                       name="second">
            <table class="table">
              <tr>
                <td>Vault shares</td>
                <td>
                  <el-input placeholder="0.0"
                            type="text"
                            v-model="shareValue"
                            style="text-align:right;"></el-input>
                </td>
              </tr>
              <tr>
                <td colspan="2">
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

const vaultAddress = '0xeDaC99A7AE93F6EA3bc23b985553D77eEF7C0009'

if (typeof web3 !== 'undefined') {
  web3 = new Web3(web3.currentProvider)
  console.log('web3 provider:web3.currentProvider')
} else {
  // set the provider you want from Web3.providers
  web3 = new Web3(new Web3.providers.HttpProvider('https://goerli.infura.io'))
  console.log('web3 provider:web3.HttpProvider')
}

export default {
  name: 'SupplyLiquidity',
  data () {
    return {
      supplyLiquidityList: [
        { 'number': 1, 'smartVaults': [{ 'iconLink': 'images/usdc.png', 'name': 'USDC', 'abi': 'ABI/contractABI.js', 'tokenAddress': '0xFA5dF5372c03D4968d128D624e3Afeb61031a777' }, { 'iconLink': 'images/weth.png', 'name': 'WETH', 'abi': 'ABI/contractABI.js', 'tokenAddress': '0x3fF5E22B4be645EF1CCc8C6e32EDe6b35c569AE4' }], 'feeTier': '0.30%', 'currentAPR': '505.66%', 'capacity': '35.1%', 'TVL': '$2,366,149', 'disabled': false },
        { 'number': 2, 'smartVaults': [{ 'iconLink': 'images/usdc.png', 'name': 'USDC', 'abi': '', 'tokenAddress': '' }, { 'iconLink': 'images/usdt.png', 'name': 'USDT', 'abi': '', 'tokenAddress': '' }], 'feeTier': '0.30%', 'currentAPR': '505.66%', 'capacity': '35.1%', 'TVL': '$2,366,149', 'disabled': true },
        { 'number': 3, 'smartVaults': [{ 'iconLink': 'images/wbtc.png', 'name': 'WBTC', 'abi': '', 'tokenAddress': '' }, { 'iconLink': 'images/usdt.png', 'name': 'USDT', 'abi': '', 'tokenAddress': '' }], 'feeTier': '0.30%', 'currentAPR': '505.66%', 'capacity': '35.1%', 'TVL': '$2,366,149', 'disabled': true }
      ],
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
      keeperContract: null,
      currentItem: null,
      depositLoading: false,
      withdrawLoading: false,
      shareValue: 0
    }
  },
  created: function () {
    this.keeperContract = new web3.eth.Contract(
      contractABI,
      vaultAddress
    )
  },
  methods: {
    getTokenContract (index) {
      this.tokenContract = new web3.eth.Contract(
        contractABI,
        this.currentItem.smartVaults[index].tokenAddress
      )
      console.log('token contract address is:' + this.currentItem.smartVaults[index].tokenAddress)
    },
    showSupplyDialog (item) {
      this.currentItem = item
      this.supplyDialogTitle = item.smartVaults[0].name + ' / ' + item.smartVaults[1].name
      this.supplyToken0 = item.smartVaults[0].name
      this.supplyToken1 = item.smartVaults[1].name
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

      this.supplyDialogVisible = true
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
      if (this.tokenContract != null) {
        this.tokenContract.methods
          .approve(vaultAddress, BigInt(90000000000000000000000))
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
      if (this.keeperContract != null) {
        console.log('account address is ' + ethereum.selectedAddress)
        this.depositLoading = true
        this.keeperContract.methods
          .deposit(
            BigInt(this.depositToken0 * 1000000000000000000),
            BigInt(this.depositToken1 * 1000000000000000000),
            BigInt(1),
            BigInt(1),
            ethereum.selectedAddress
          )
          .send({
            from: ethereum.selectedAddress,
            gasPrice: '10000000000',
            gas: 200000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            // GetBalance(accountAddress)
            if (_this.depositLoading === true) {
              _this.depositLoading = false
              _this.$message('Successfully deposited!')
              console.log('confirmation')
            }
          })
          .on('receipt', function (receipt) {
            if (_this.depositLoading === true) {
              _this.depositLoading = false
              _this.$message('Successfully deposited!')
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
      if (this.keeperContract != null) {
        console.log('account address is ' + ethereum.selectedAddress)
        this.withdrawLoading = true
        this.keeperContract.methods
          .withdraw(
            BigInt(this.shareValue * 1000000000000000000),
            BigInt(1000000000000000000),
            BigInt(1000000000000000000),
            ethereum.selectedAddress
          )
          .send({
            from: ethereum.selectedAddress,
            gasPrice: '10000000000',
            gas: 200000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            // GetBalance(accountAddress)
            if (_this.withdrawLoading === true) {
              _this.withdrawLoading = false
              _this.$message('Successfully withdrawed!')
              console.log('confirmation')
            }
          })
          .on('receipt', function (receipt) {
            if (_this.withdrawLoading === true) {
              _this.withdrawLoading = false
              _this.$message('Successfully withdrawed!')
              console.log('receipt')
            }
          })
          .on('error', function (error) {
            _this.$message.error(error)
            _this.withdrawLoading = false
          })
      }
    },
    handleTabClick (tab, event) {
      console.log(tab, event)
    }
  }
}
</script>

<style scoped>
.table td,
.table th {
  border: 0px solid transparent;
}
</style>

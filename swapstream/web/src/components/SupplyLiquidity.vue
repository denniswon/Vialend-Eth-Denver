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
                  <!-- <span class="lblBalance">Balance:{{token0Balance}}(Max)</span> -->
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
                  <!-- <span class="lblBalance">Balance:{{token1Balance}}(Max)</span> -->
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
import coinABI from '../ABI/ERC20.json'

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
      supplyLiquidityList: [
        { 'number': 1, 'smartVaults': [{ 'iconLink': 'images/usdc.png', 'name': 'eUSDC', 'abi': 'ABI/contractABI.js', 'tokenAddress': '0xFdA9705FdB20E9A633D4283AfbFB4a0518418Af8' }, { 'iconLink': 'images/weth.png', 'name': 'eWETH', 'abi': 'ABI/contractABI.js', 'tokenAddress': '0x48FCb48bb7F70F399E35d9eC95fd2A614960Dcf8' }], 'feeTier': '0.30%', 'currentAPR': '505.66%', 'capacity': '35.1%', 'TVL': '$2,366,149', 'disabled': false },
        { 'number': 2, 'smartVaults': [{ 'iconLink': 'images/usdc.png', 'name': 'dUSDC', 'abi': 'ABI/contractABI.js', 'tokenAddress': '0x88177e1a55c6Ca956A738abbF6d87148217a8Cb0' }, { 'iconLink': 'images/weth.png', 'name': 'dWETH', 'abi': 'ABI/contractABI.js', 'tokenAddress': '0x495A3648AfDeb15Ce7B0cDDff44EAeB5E014cEAD' }], 'feeTier': '0.30%', 'currentAPR': '505.66%', 'capacity': '35.1%', 'TVL': '$2,366,149', 'disabled': false },
        { 'number': 3, 'smartVaults': [{ 'iconLink': 'images/wbtc.png', 'name': 'WBTC', 'abi': '', 'tokenAddress': '' }, { 'iconLink': 'images/usdt.png', 'name': 'USDT', 'abi': '', 'tokenAddress': '' }], 'feeTier': '0.30%', 'currentAPR': '505.66%', 'capacity': '35.1%', 'TVL': '$2,366,149', 'disabled': true }
      ],
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
      token0Balance: 1,
      token1Balance: 2
    }
  },
  created: function () {
  },
  mounted: function () {

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
        // console.log('this.isConnected=' + this.isConnected)
        // console.log('isConnected=' + this.$parent.isConnected)
        console.log('wallet is disconnected')
        this.isConnected = false
      }
    },
    showSupplyDialog (item) {
      var _this = this
      this.checkConnectionStatus()
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
      // Get the balance of the account of given address
      console.log('addr=', this.currentItem.smartVaults[0].tokenAddress)
      web3.eth.getBalance('0xFA5dF5372c03D4968d128D624e3Afeb61031a777').then(function (balance) {
        // _this.token0Balance = balance
        console.log('balance=', balance)
      })

      var coinContract = new web3.eth.Contract(
        coinABI,
        '0xaa16E934A327D500fdE1493302CeB394Ff6Ff0b2', { from: '0x6dd19aEB91d1f43C46f0DD74C9E8A92BFe2a3Cd0' })

      web3.eth.getBalance(ethereum.selectedAddress).then(console.log)

      coinContract.methods.balanceOf('0xFdA9705FdB20E9A633D4283AfbFB4a0518418Af8').call({ from: '0x6dd19aEB91d1f43C46f0DD74C9E8A92BFe2a3Cd0' }, function (error, result) {
        if (!error) {
          console.log('result=', result)
        } else {
          console.log(error)
        }
      })
      this.getBal()
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
    async getBal () {
      const minABI = [
        // balanceOf
        {
          constant: true,
          inputs: [{ name: '_owner', type: 'address' }],
          name: 'balanceOf',
          outputs: [{ name: 'balance', type: 'uint256' }],
          type: 'function'
        }
      ]
      var coinContract = new web3.eth.Contract(
        minABI,
        '0xFdA9705FdB20E9A633D4283AfbFB4a0518418Af8')
      const result = await coinContract.methods.balanceOf(ethereum.selectedAddress).call() // 29803630997051883414242659
      // const format = Web3Client.utils.fromWei(result) // 29803630.997051883414242659
      console.log('balance000=', result)
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
      if (this.$parent.keeperContract != null) {
        console.log('account address is ' + ethereum.selectedAddress)
        this.depositLoading = true
        this.$parent.keeperContract.methods
          .deposit(
            BigInt(this.depositToken0 * 1000000000000000000),
            BigInt(this.depositToken1 * 1000000000000000000),
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
            BigInt(this.shareValue * 1000000000000000000),
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
            // GetBalance(accountAddress)
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
.lblBalance {
  float: right;
}
</style>

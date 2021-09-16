<template>
  <div>
    <Header />
    <el-tabs type="border-card">
      <el-tab-pane>
        <span slot="label"><i class="el-icon-date"></i> Rebalance</span>
        Token0 balance in wallet:{{token0BalanceInWallet}}<br>
        Token1 balance in wallet:{{token1BalanceInWallet}}<br>
        Token0 balance in vault:{{token0BalanceInVault}}<br>
        Token1 balance in vault:{{token1BalanceInVault}}<br>
        Token0 balance in pool:{{token0BalanceInPool}}<br>
        Token1 balance in pool:{{token1BalanceInPool}}<br>
        <hr>
        Min Tick:{{tickLower}}<br>
        Max Tick:{{tickUpper}}<br>
        Current Tick:{{currentTick}}<br>
        <hr>
        <div class="range_title">Set Price Range</div>
        <el-form :inline="true"
                 :model="rangeForm"
                 class="demo-form-inline">
          <el-form-item label="Min Price">
            <el-input-number v-model="rangeForm.minPrice"
                             :precision="1"
                             :step="1"></el-input-number><br>
            eUSDC per eWETH
          </el-form-item>
          <el-form-item label="Max Price">
            <el-input-number v-model="rangeForm.maxPrice"
                             :precision="1"
                             :step="1"></el-input-number><br>
            eUSDC per eWETH
          </el-form-item>
          <el-form-item>
            <el-button type="primary"
                       @click="doRebalance">Rebalance</el-button>
          </el-form-item>
        </el-form>
        {{errorRebalance}}
      </el-tab-pane>
      <el-tab-pane>
        <span slot="label"><i class="el-icon-date"></i> Security</span>
        remove all liquidity from uniswap
      </el-tab-pane>
    </el-tabs>

  </div>
</template>

<script>
import Web3 from 'web3'
import Header from '@/components/Header.vue'
import uniswapV3PoolABI from '../ABI/UniswapV3PoolABI.json'
import contractABI from '../ABI/contractABI.json'
import tokenABI from '../ABI/tokenABI.json'

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
      keeperUniswapV3Contract: null,
      keeperContract: null,
      tickLower: 0,
      tickUpper: 0,
      currentTick: 0,
      rangeForm: {
        minPrice: 0.0,
        maxPrice: 0.0
      },
      errorRebalance: '',
      rebalanceLoading: false,
      token0BalanceInWallet: 0,
      token1BalanceInWallet: 0,
      token0BalanceInVault: 0,
      token1BalanceInVault: 0,
      token0BalanceInPool: 0,
      token1BalanceInPool: 0
    }
  },
  created: function () {
    console.log('this.$parent.vaultAddress=', this.$parent.vaultAddress)
    this.keeperUniswapV3Contract = new web3.eth.Contract(
      uniswapV3PoolABI,
      '0x3c7fADe1921Bf9D8308D76d7B09cA54839cfF033'
    )
    this.keeperContract = new web3.eth.Contract(
      contractABI,
      this.$parent.vaultAddress
    )
    this.getSlot0()
    this.getTokensBalanceInWallet()
    this.getTokensBalanceInVaultAndPool()
  },
  mounted () {
  },
  computed: {
    newMinPrice () {
      return this.rangeForm.minPrice
    },
    newMaxPrice () {
      return this.rangeForm.maxPrice
    }
  },
  watch: {
    newMinPrice (price) {
      console.log('new min price=', price)
      if (!isNaN(price)) { this.tickLower = this.priceToTick(price) } else { this.tickLower = 0 }
    },
    newMaxPrice (price) {
      console.log('new max price=', price)
      this.tickUpper = this.priceToTick(price)
      if (!isNaN(price)) { this.tickUpper = this.priceToTick(price) } else { this.tickUpper = 0 }
    }
  },
  methods: {
    async getSlot0 () {
      if (this.keeperUniswapV3Contract !== null) {
        this.keeperUniswapV3Contract.methods
          .slot0()
          .call()
          .then(slot => {
            if (slot !== undefined && slot !== null && slot !== '') {
              console.log('slot0=' + JSON.stringify(slot))
              // var slot0 = JSON.parse(val)
              console.log('tick=', parseInt(slot['tick']) - 1000)
              // this.rangeForm.tickLower = parseInt(slot['tick']) - 1000
              // this.rangeForm.tickUpper = parseInt(slot['tick']) + 1000
              // this.tickLower = Math.round(parseInt(this.rangeForm.tickLower) / 60) * 60
              // this.tickUpper = Math.round(parseInt(this.rangeForm.tickUpper) / 60) * 60
              this.currentTick = slot['tick']
            }
          })
      }
    },
    async getTokensBalanceInWallet () {
      var coinContract
      coinContract = new web3.eth.Contract(
        tokenABI,
        '0x48FCb48bb7F70F399E35d9eC95fd2A614960Dcf8')
      this.token0BalanceInWallet = await coinContract.methods.balanceOf(ethereum.selectedAddress).call() // 29803630997051883414242659
      // const format = Web3Client.utils.fromWei(result) // 29803630.997051883414242659
      console.log('token0BalanceInWallet=', this.token0BalanceInWallet)
      var tokenDecimal = await coinContract.methods.decimals().call()
      console.log('decimal_0x48=', tokenDecimal)
      // token1BalanceInWallet
      coinContract = new web3.eth.Contract(
        tokenABI,
        '0xFdA9705FdB20E9A633D4283AfbFB4a0518418Af8')
      this.token1BalanceInWallet = await coinContract.methods.balanceOf(ethereum.selectedAddress).call() // 29803630997051883414242659
      // const format = Web3Client.utils.fromWei(result) // 29803630.997051883414242659
      console.log('token1BalanceInWallet=', this.token1BalanceInWallet)
      tokenDecimal = await coinContract.methods.decimals().call()
      console.log('decimal_0xFd=', tokenDecimal)
    },
    async getTokensBalanceInVaultAndPool () {
      if (this.$parent.keeperContract != null) {
        this.token0BalanceInVault = await this.$parent.keeperContract.methods.getBalance0().call()
        console.log('token0BalanceInVault=', this.token0BalanceInVault)
        this.token1BalanceInVault = await this.$parent.keeperContract.methods.getBalance1().call()
        console.log('token1BalanceInVault=', this.token1BalanceInVault)
        var result = await this.$parent.keeperContract.methods.getPositionAmounts(-201360, -198910).call()
        if (result !== undefined && result !== null) {
          this.token0BalanceInPool = result.amount0
          this.token1BalanceInPool = result.amount1
          console.log('result000=', result)
        }
        var decimal = await this.$parent.keeperContract.methods.decimals().call()
        console.log('decimal=', decimal)
      }
    },
    doRebalance () {
      var _this = this
      this.rebalanceLoading = true
      if (isNaN(this.rangeForm.minPrice) || this.rangeForm.minPrice <= 0) {
        this.$message({
          message: 'Please input positive number greater than 0 in Min Price.',
          type: 'warning'
        })
        return
      }
      if (isNaN(this.rangeForm.maxPrice) || this.rangeForm.maxPrice <= 0) {
        this.$message({
          message: 'Please input positive number greater than 0 in Max Price.',
          type: 'warning'
        })
        return
      }
      if (this.keeperContract != null) {
        console.log('account address is ' + ethereum.selectedAddress)
        console.log('this.rangeForm.tickLower=', Math.round(parseInt(this.tickLower) / 60) * 60)
        console.log('this.rangeForm.tickUpper=', Math.round(parseInt(this.tickUpper) / 60) * 60)

        this.keeperContract.methods
          .rebalance(
            BigInt(Math.round(parseInt(this.tickLower) / 60) * 60),
            BigInt(Math.round(parseInt(this.tickUpper) / 60) * 60),
            0
          )
          .send({
            from: ethereum.selectedAddress,
            gasPrice: '80000000000',
            gas: 600000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            if (_this.rebalanceLoading) {
              _this.rebalanceLoading = false
              _this.$message('rebalance confirmed!')
              console.log('confirmation')
              // _this.errorRebalance = receipt
            }
          })
          .on('receipt', function (receipt) {
            if (_this.rebalanceLoading) {
              _this.rebalanceLoading = false
              _this.$message('rebalance receipt!')
              console.log('receipt')
              // _this.errorRebalance = receipt
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
    onSubmit () {
      console.log('submit!')
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
</style>

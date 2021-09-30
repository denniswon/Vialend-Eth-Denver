<template>
  <div class="about">
    <el-form :inline="true">
      <el-form-item label="AccountAddress">
        <el-input v-model="walletAddress"
                  maxlength="300"
                  placeholder="AccountAddress"></el-input>
      </el-form-item>

      <el-form-item>
        <el-button type="primary"
                   @click="connectWallet">Connect Wallet</el-button>
        <el-button type="primary"
                   @click="getTestData">Get TestData</el-button>
      </el-form-item>
    </el-form>

    <hr />
    <div>
      Amount0Desired:<input type="text"
             id="txtAmount0Desired"
             name="txtAmount0Desired"
             size="30"
             value="10" />
    </div>
    <div>
      Amount1Desired:<input type="text"
             id="txtAmount1Desired"
             name="txtAmount1Desired"
             size="30"
             value="5" />
    </div>
    <div>
      Amount0Min:<input type="text"
             id="txtAmount0Min_D"
             name="txtAmount0Min_D"
             size="30"
             value="5" />
    </div>
    <div>
      Amount1Min:<input type="text"
             id="txtAmount1Min_D"
             name="txtAmount1Min_D"
             size="30"
             value="2" />
    </div>
    <div>
      <el-button @click="deposit">Deposit</el-button>

      <label id="depositresult"></label>
    </div>
    <hr />
    <div>Shares:<input type="text"
             name="txtShares"
             size="30"
             value="" /></div>
    <div>
      Amount0Min:<input type="text"
             name="txtAmount0Min_W"
             size="30"
             value="" />
    </div>
    <div>
      Amount1Min:<input type="text"
             name="txtAmount1Min_W"
             size="30"
             value="" />
    </div>
    <div>
      <input type="button"
             name="btnWithdraw"
             id="btnWithdraw"
             value="Withdraw" />
    </div>
    <div>TotalAmount:</div>
    <div>
      total0:<input type="text"
             id="txtTotal0"
             name="txtTotal0"
             value="" />
    </div>
    <div>
      total1:<input type="text"
             id="txtTotal1"
             name="txtTotal1"
             value="" />
    </div>
    <div>
      <el-input v-model="testData"
                placeholder="Test Data"></el-input>
    </div>
  </div>
</template>

<script>
// https://blog.csdn.net/qq_36838406/article/details/118386159?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_title~default-0.control&spm=1001.2101.3001.4242
import Web3 from 'web3'
import JSBI from 'jsbi'
import contractABI from '../ABI/contractABI.json'
import uniswapV3PoolABI from '../ABI/UniswapV3PoolABI.json'
import v3FactoryABI from '../ABI/V3FactoryABI.json'
import SqrtPriceMath from '../utils/sqrtPriceMath'
import { TickMath } from '../utils/tickMath'

const privateKey = '014bebbe9b56c23c7c2913e0800a85d83e05e6551769ae15d420b464fafc7c48'
const contractAddress = '0xeDaC99A7AE93F6EA3bc23b985553D77eEF7C0009'
const _this = this

const Tx = require('ethereumjs-tx') // .Transaction

if (typeof web3 !== 'undefined') {
  web3 = new Web3(web3.currentProvider)
  console.log('web3 provider:web3.currentProvider')
} else {
  // set the provider you want from Web3.providers
  web3 = new Web3(new Web3.providers.HttpProvider('https://goerli.infura.io'))
  console.log('web3 provider:web3.HttpProvider')
}

var web3Provider, web3js
var contractObj = new web3.eth.Contract(contractABI, contractAddress)
window.ethereum.autoRefreshOnNetworkChange = false
let currentAccount = null

window.ethereum.on('accountsChanged', () => {
  console.log('accountsChanged')
  connectWallet()
  // switch (_this.$route.path) {
  //   case '/':
  //     this.connectMetaMask()
  //     break
  //   default:
  // }
})
window.ethereum.on('chainChanged', () => {
  console.log('chainChanged')
  connectWallet()
})

function toUint8Arr (str) {
  const buffer = []
  for (let i of str) {
    const _code = i.charCodeAt(0)
    if (_code < 0x80) {
      buffer.push(_code)
    } else if (_code < 0x800) {
      buffer.push(0xc0 + (_code >> 6))
      buffer.push(0x80 + (_code & 0x3f))
    } else if (_code < 0x10000) {
      buffer.push(0xe0 + (_code >> 12))
      buffer.push(0x80 + (_code >> 6 & 0x3f))
      buffer.push(0x80 + (_code & 0x3f))
    }
  }
  return Uint8Array.from(buffer)
}

export default {
  components: {},
  data () {
    return {
      testData: '',
      isConnected: false,
      walletAddress: '',
      amount0Desired: 10,
      amount1Desired: 5,
      amount0Min: 5,
      amount1Min: 2
    }
  },
  created: function () {
    // this.showTreaty()
    // ethereum
    //   .request({ method: 'eth_accounts' })
    //   .then(this.handleAccountsChanged)
    //   .catch((err) => {
    //     // Some unexpected error.
    //     // For backwards compatibility reasons, if no accounts are available,
    //     // eth_accounts will return an empty array.
    //     console.error(err)
    //   })
    console.log('load create function')
  },
  mounted () {
    window.connectWallet = this.connectWallet
  },
  methods: {
    showTreaty () {
      this.treatyDialogVisible = true
      // console.log(this.$store.state.name)
    },
    async getTestData () {
      // this.testData = await contractObj.methods.Hello().call()
      var factoryV3 = new web3.eth.Contract(v3FactoryABI, '0x1F98431c8aD98523631AE4a59f267346ea31F984')
      // eslint-disable-next-line camelcase
      var pool_address = await factoryV3.methods.getPool('0x31E84D42aB6DEf5Dac84b761b0E5004179e07778', '0x48FCb48bb7F70F399E35d9eC95fd2A614960Dcf8', 3000).call()
      // eslint-disable-next-line camelcase
      var pool_1 = new web3.eth.Contract(uniswapV3PoolABI, pool_address)
      // eslint-disable-next-line camelcase
      var pool_balance = await pool_1.methods.slot0.call().call()
      var sqrtPriceX96 = pool_balance[0]
      console.log('pool_address=', pool_address)
      console.log('pool_balance=', pool_balance)
      console.log('varsqrtPriceX96=', sqrtPriceX96)

      // eslint-disable-next-line camelcase
      var number_1 = sqrtPriceX96 * sqrtPriceX96 * (1e18) / (1e6) / JSBI.BigInt(2) ** (JSBI.BigInt(192))

      console.log('number_1=', number_1)
      // var amount1 = SqrtPriceMath.getAmount1Delta(
      //   TickMath.getSqrtRatioAtTick(this.tickLower),
      //   5000,
      //   5.226576398620404512480257589377e-4,
      //   true
      // )
      // console.log('amount1=', amount1)
    },
    connectWallet () {
      if (window.ethereum) {
        web3Provider = window.ethereum
        try {
          // 请求用户授权
          window.ethereum.enable()
        } catch (error) {
          // 用户不授权时
          console.error('User denied account access')
        }
      }
      web3js = new Web3(web3Provider)// web3js就是你需要的web3实例

      web3js.eth.getAccounts(function (error, result) {
        if (!error) {
          currentAccount = result[0]
          this.walletAddress = result[0]
          console.log('this.walletAddress=' + this.walletAddress)
        }// 授权成功后result能正常获取到账号了
      })
    },
    handleAccountsChanged (accounts) {
      if (accounts.length === 0) {
        // MetaMask is locked or the user has not connected any accounts
        console.log('Please connect to MetaMask.')
      } else if (accounts[0] !== currentAccount) {
        currentAccount = accounts[0]
        // Do any other work!
        // if (currentAccount != null && currentAccount.result != null) {
        this.isConnected = true
        // this.walletAddress = currentAccount
        console.log('accounts=', currentAccount)
        // }
      }
    },
    async deposit () {
      // 导入ethereumjs-tx库，通过npm install安装

      // _from为发起交易的地址
      var _from = '0xa5f8dE976675F5241502FF5b142B281fA95647A4'
      // nonce随机数，这里取该账号的交易数量
      // var number = web3.eth.getTransactionCount(_from).toString(16)
      console.log('currentAccount=' + currentAccount)
      var functionEncode = await contractObj.methods.deposit(this.amount0Desired, this.amount1Desired, this.amount0Min, this.amount1Min, currentAccount).encodeABI()
      // console.log('nonce=' + JSON.stringify(number))
      // var rawTx = {
      //   nonce: '0x' + number, // 随机数
      //   // gasPrice和gasLimit如果不知道怎么填，可以参考etherscan上的任意一笔交易的值
      //   gasPrice: '0x77359400',
      //   gasLimit: '0x295f05',
      //   to: contractAddress, // 接受方地址或者合约地址
      //   value: '0x100', // 发送的金额，这里是16进制，实际表示发送256个wei
      //   data: functionEncode
      // }
      // var priKey = Buffer.from('0x014bebbe9b56c23c7c2913e0800a85d83e05e6551769ae15d420b464fafc7c48', 'hex')
      // let nonce = await web3.eth.getTransactionCount(contractAddress)
      // let gasPrice = await web3.eth.getGasPrice()

      // var rawTx = {
      //   nonce: nonce,
      //   from: _from,
      //   gasPrice: gasPrice,
      //   to: contractAddress,
      //   value: '0x00',
      //   data: functionEncode
      // }

      const transactionParameters = {
        nonce: '0x00', // ignored by MetaMask
        gasPrice: '0x520800', // customizable by user during MetaMask confirmation.
        gas: '0x520800', // customizable by user during MetaMask confirmation.
        to: contractAddress, // Required except during contract publications.
        from: ethereum.selectedAddress, // must match user's active address.
        value: '0x00', // Only required to send ether to the recipient from the initiating external account.
        data: functionEncode // Optional, but used for defining smart contract creation and interaction.
        // chainId: '0x3' // Used to prevent transaction reuse across blockchains. Auto-filled by MetaMask.
      }

      // txHash is a hex string
      // As with any RPC call, it may throw an error
      const txHash = await ethereum.request({
        method: 'eth_sendTransaction',
        params: [transactionParameters]
      })
      console.log('txHash=' + txHash)
    }
  }
}
</script>

<style scoped>
.about {
  margin-left: 20px;
}
</style>

<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <button id="btn-wallet" v-on:click="linkWallet()">连接wallet</button>
    <button @click="getAccount()">获取account</button>
    <input id="input-account" readonly="readonly" type="text" value="" />
  </div>
</template>

<script>
import Web3 from "../assets/util/web3.min.js";

export default {
  name: "Home",
  props: {
    msg: String,
  },
  data() {
    return {
      web3Provider: null,
      web3js: null,
    };
  },
  methods: {
    linkWallet() {
      console.log("linking");
      if (window.ethereum) {
        this.web3Provider = window.ethereum;
        // 新版需要请求用户授权
        try {
          window.ethereum.enable();
        } catch (error) {
          alert("用户取消授权");
          return;
        }
      } else if (Web3.web3) {
        // MetaMask Legacy dapp browsers...
        this.web3Provider = Web3.web3.currentProvider;
        console.log("web3.currentProvider:");
        console.log(Web3.web3.currentProvider);
      } else {
        this.web3Provider = new Web3.providers.HttpProvider(
          "http://localhost:8545"
        );
        console.log("https://http-testnet.hecochain.com");
      }
      this.web3js = new Web3(this.web3Provider);

      document.getElementById("btn-wallet").innerText = "重新连接";
    },
    getAccount() {
      if (!this.web3js) {
        document.getElementById("input-account").value = "请先连接 wallet";
        return;
      }
      this.web3js.eth.getAccounts(function (error, result) {
        if (!error) {
          document.getElementById("input-account").value = result;
        } else {
          document.getElementById("input-account").value = "获取地址失败";
        }
      });
      console.log("获取链 ID", this.web3js.eth.net.getId());
      console.log("ChainId获取链 ID", this.web3js.eth.getChainId()),
        this.web3js.eth.net.getId((err, netID) => {
          // Main Network: 1   表示主网
          // Ropsten Test Network: 3  //测试网Ropsten
          // Kovan Test Network: 42  //测试网Kovan
          // Rinkeby Test Network: 4  //测试网Rinkeby
          console.log(netID);
        });
      console.log("查询当前区块高度", this.web3js.eth.getBlockNumber());
      this.web3js.eth.getBlockNumber().then(console.log);
      console.log(
        "当前地址余额：",
        this.web3js.eth
          .getBalance("0x2F63F0B229749147A9eaF813EF62e5EE2b8Db21f")
          .then(console.log)
      );
      //获取当前metamask账户地址的eth余额
      this.web3js.eth.getCoinbase((err, coinbase) => {
        this.web3js.eth.getBalance(coinbase).then(console.log);
      });
      //通过hash查询交易
      console.log(
        "查询交易",
        this.web3js.eth.getTransaction(
          "0x239f038ac2bfe5e456f1480c1ba01564d40719132f78f9244858297251d2fded"
        )
      );
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>

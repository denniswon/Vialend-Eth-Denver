<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
// import Web3 from 'web3'
//https://blog.csdn.net/weixin_33815613/article/details/88803600
// if (typeof web3 !== 'undefined') {
//   web3 = new Web3(web3.currentProvider)
//   console.log('web3 provider:web3.currentProvider')
// } else {
//   // set the provider you want from Web3.providers
//   web3 = new Web3(new Web3.providers.HttpProvider('https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6'))
//   console.log('web3 provider:goerli')
// }

export default {
  name: 'App'
}
</script>

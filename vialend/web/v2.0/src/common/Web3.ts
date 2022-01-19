import Web3 from 'web3'
import { AbiItem } from 'web3-utils'

let web3:any

export function getWeb3Instance() {
  let web3Provider
  if ((window as any).web3) {
    web3Provider = (window as any).web3.currentProvider
  } else {
    web3Provider = new Web3.providers.HttpProvider('https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6')
    console.log('Non-Ethereum browser detected. You should consider trying MetaMask!')
  }
  web3 = new Web3(web3Provider)
  return web3
}

// ((window as any).ethereum) {
//   web3Provider = (window as any).ethereum
//   // try {
//   //   (window as any).ethereum.enable()
//   // } catch (error) {
//   //   console.log(error)
//   // }
// } else if

export function contractInstance(abi:AbiItem, address:string) {
  if (web3 === undefined || web3 === null) {
    web3 = getWeb3Instance()
    console.log('create web3 instance')
  }
  return new web3.eth.Contract(abi, address)
}

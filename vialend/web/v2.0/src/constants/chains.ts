import { getWeb3Instance } from '../common/Web3'
import store from '@/store'

export enum SupportedChainId {
  // MAINNET = 1,
  ROPSTEN = 3,
  // RINKEBY = 4,
  GOERLI = 5
  //, KOVAN = 42
}

interface ChainInfo{
readonly bridgeAddress : string
readonly etherscanTxLink: string
readonly etherscanAddressLink: string
}

type ChainObject = { readonly [chainId: number]: ChainInfo } & { readonly [chainId in SupportedChainId]: ChainInfo}

export const CHAININFO:ChainObject = {
  // [SupportedChainId.MAINNET]: {
  //   bridgeAddress: ''
  // },
  [SupportedChainId.ROPSTEN]: {
    bridgeAddress: '0xE1FE39c9dc4004C4625499063A73fCA6012C1192',
    etherscanTxLink: 'https://ropsten.etherscan.io/tx/',
    etherscanAddressLink: 'https://ropsten.etherscan.io/address/'
  },
  // [SupportedChainId.RINKEBY]: {
  //   bridgeAddress: ''
  // },
  [SupportedChainId.GOERLI]: {
    bridgeAddress: '0x428EeA0B87f8E0f5653155057f58aaaBb667A3ec',
    etherscanTxLink: 'https://goerli.etherscan.io/tx/',
    etherscanAddressLink: 'https://goerli.etherscan.io/address/'
  }
  // ,[SupportedChainId.KOVAN]: {
  //   bridgeAddress: ''
  // }
}

export async function checkChain() {
  let vaildChain
  const web3 = getWeb3Instance()
  const chainId = await web3.eth.getChainId()
  if (chainId in SupportedChainId) {
    console.log('checkChain[id:', chainId, '] -> current ChainId is vaild')
    store.state.validNetwork = true
    vaildChain = true
  } else {
    console.log('checkChain[id:', chainId, ']  -> current ChainId is invaild')
    store.state.isAdmin = false
    store.state.isConnected = false
    store.state.validNetwork = false
    vaildChain = false
  }
  store.state.chainId = chainId
  return vaildChain
}

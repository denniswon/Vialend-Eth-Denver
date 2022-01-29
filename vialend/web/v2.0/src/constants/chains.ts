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
}

type ChainObject = { readonly [chainId: number]: ChainInfo } & { readonly [chainId in SupportedChainId]: ChainInfo}

export const CHAININFO:ChainObject = {
  // [SupportedChainId.MAINNET]: {
  //   bridgeAddress: ''
  // },
  [SupportedChainId.ROPSTEN]: {
    bridgeAddress: '0x033F3C5eAd18496BA462783fe9396CFE751a2342'
  },
  // [SupportedChainId.RINKEBY]: {
  //   bridgeAddress: ''
  // },
  [SupportedChainId.GOERLI]: {
    bridgeAddress: '0x428EeA0B87f8E0f5653155057f58aaaBb667A3ec'
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

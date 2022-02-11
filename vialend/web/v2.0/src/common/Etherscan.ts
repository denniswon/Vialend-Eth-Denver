import { CHAININFO } from '@/constants/chains'
import store from '@/store'

export function getEtherscanTx(txHash:string) {
  const link = CHAININFO[store.state.chainId].etherscanTxLink
  return link.concat(txHash)
}

export function getEtherscanAddress(address:string) {
  console.log('Etherscan store.state.chainId:', store.state.chainId)
  if (store.state.chainId > 0) {
    const link = CHAININFO[store.state.chainId].etherscanAddressLink
    return link.concat(address)
  } else {
    return ''
  }
}

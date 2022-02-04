import { UserModule } from '@/store/modules/user'
import { CHAININFO } from '@/constants/chains'
import { contractInstance } from '@/common/Web3'
import store from '@/store'

const VaultBridgeABI = require('@/abi/VaultBridge.json')
const ethereum = (window as any).ethereum

export function changeAccount() {
  console.log('doChangeAccount')
  if (ethereum.selectedAddress !== null && ethereum.selectedAddress !== undefined) {
    store.state.isConnected = true
    store.state.currentAccount = ethereum.selectedAddress
  } else {
    store.state.isConnected = false
    store.state.currentAccount = ''
  }
}

export async function checkValidAdmin() {
  console.log('Account->bridgeAddress=', CHAININFO[store.state.chainId].bridgeAddress)
  if (ethereum.selectedAddress !== null && ethereum.selectedAddress !== undefined) {
    const vaultAdminContract = await contractInstance(VaultBridgeABI, CHAININFO[store.state.chainId].bridgeAddress)
    const isAdmin = await vaultAdminContract.methods.getPermit((window as any).ethereum.selectedAddress).call()
    if (isAdmin === '1') {
      store.state.isAdmin = true
      console.log('isAdmin=true')
    } else {
      store.state.isAdmin = false
      console.log('isAdmin=false')
    }
    UserModule.ChangeRoles(store.state.isAdmin ? 'admin' : 'user')
  }
}

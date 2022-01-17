import Web3 from './components/web3/web3.min.js'
import Token from './model/Token'
import Pairs from './model/Pairs'

let web3:any

function getWeb3Instance() {
  let web3Provider
  if ((window as any).ethereum) {
    web3Provider = (window as any).ethereum
    // try {
    //   (window as any).ethereum.enable()
    // } catch (error) {
    //   console.log(error)
    // }
  } else if ((window as any).web3) {
    web3Provider = (window as any).web3.currentProvider
  } else {
    web3Provider = new Web3.providers.HttpProvider('https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6')
    console.log('Non-Ethereum browser detected. You should consider trying MetaMask!')
  }
  web3 = new Web3(web3Provider)
  return web3
}

function contractInstance(abi:any, address:string) {
  if (web3 === undefined || web3 === null) {
    web3 = getWeb3Instance()
    console.log('create web3 instance')
  }
  return new web3.eth.Contract(abi, address)
}

//main code
const bridgeContract = await contractInstance(VaultBridgeABI, this.bridgeAddress)
let iNum = 0
let vaultAddressInContract
while (true) {
	try {
		vaultAddressInContract = await bridgeContract.methods.getAddress(iNum).call()
	} catch (err) {
		console.log('getting vault Address in contract occurs error:', err)
	}
	console.log('vaultAddressInContract', iNum, '=', vaultAddressInContract)
	if (vaultAddressInContract === null || vaultAddressInContract === undefined || Number(vaultAddressInContract) === 0) {
		console.log('vaultAddressInContract', iNum, ' is null,so break;')
		break
	}
	const pair = new Pairs()
	pair.index = iNum
	pair.id = iNum + 1
	pair.vaultAddress = vaultAddressInContract
	console.log('vault address(', iNum, ')=', vaultAddressInContract)
	// pair = await this.getPairsBaseInfo(pair)
	this.pairsList.add(pair)
	iNum++
}
let str:string = "hello ts"

console.log('str=',str)
export function getEtherscanTx(txHash:string) {
  const link = 'https://goerli.etherscan.io/tx/'
  return link.concat(txHash)
}

export function getEtherscanAddress(address:string) {
  const link = 'https://goerli.etherscan.io/address/'
  return link.concat(address)
}

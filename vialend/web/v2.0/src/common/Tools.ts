export function formatWalletAddress(addr:string) {
  if (addr !== null && addr !== undefined && addr !== '') {
    return addr.toUpperCase().substring(0, 6).concat('...', addr.toUpperCase().substring(-4, 4))
  } else {
    return ''
  }
}

export function priceToTick(price:number, decimal0:number, decimal1:number) {
  console.log('price=', price, 'decimal0=', decimal0, 'decimal1=', decimal1)
  // return parseInt(Math.log(price / Math.pow(10, (decimal0 - decimal1))) / Math.log(1.0001))
  return Math.log(price / Math.pow(10, (decimal0 - decimal1))) / Math.log(1.0001)
}

export function tickToPrice(tick:number, decimal0:number, decimal1:number) {
  console.log('decimal0=', decimal0, 'decimal1=', decimal1)
  console.log('Math.pow(1.0001, tick)=', Math.pow(1.0001, tick), 'decimal=', decimal0 - decimal1)
  return Math.pow(1.0001, tick) * Math.pow(10, (decimal0 - decimal1))
}

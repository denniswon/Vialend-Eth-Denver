"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.tickToPrice = exports.priceToTick = exports.formatWalletAddress = void 0;
function formatWalletAddress(addr) {
    if (addr !== null && addr !== undefined && addr !== '') {
        return addr.toUpperCase().substring(0, 6).concat('...', addr.toUpperCase().substring(-4, 4));
    }
    else {
        return '';
    }
}
exports.formatWalletAddress = formatWalletAddress;
function priceToTick(price, decimal0, decimal1) {
    console.log('price=', price, 'decimal0=', decimal0, 'decimal1=', decimal1);
    // return parseInt(Math.log(price / Math.pow(10, (decimal0 - decimal1))) / Math.log(1.0001))
    return Math.log(price / Math.pow(10, (decimal0 - decimal1))) / Math.log(1.0001);
}
exports.priceToTick = priceToTick;
function tickToPrice(tick, decimal0, decimal1) {
    console.log('decimal0=', decimal0, 'decimal1=', decimal1);
    console.log('Math.pow(1.0001, tick)=', Math.pow(1.0001, tick), 'decimal=', decimal0 - decimal1);
    return Math.pow(1.0001, tick) * Math.pow(10, (decimal0 - decimal1));
}
exports.tickToPrice = tickToPrice;

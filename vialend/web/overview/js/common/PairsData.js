"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
// import { Vue } from 'vue-property-decorator'
const Web3_1 = require("./Web3");
const Token_1 = __importDefault(require("../model/Token"));
const Pairs_1 = __importDefault(require("../model/Pairs"));
const ArrayList_1 = __importDefault(require("./ArrayList"));
const VaultBridgeABI = require('../abi/VaultBridge.json');
const contractABI = require('../abi/ViaLendFeeMakerABI.json');
const uniswapV3PoolABI = require('../abi/UniswapV3PoolABI.json');
const ViaLendTokenABI = require('../abi/VialendTokenABI.json');
class PairsData {
    constructor() {
        this.pairsList = new ArrayList_1.default();
        this.pairsBaseInfoLoading = false;
        this.pairsBalanceLoading = false;
        this.pairsLoadComplete = false;
        this.bridgeAddress = '0x033F3C5eAd18496BA462783fe9396CFE751a2342';
        this.calculateAPY = false;
        this.calculateAPR = false;
        this.ethereum = window.ethereum;
        this.web3 = (0, Web3_1.getWeb3Instance)();
    }
    loadPairsInfo() {
        return __awaiter(this, void 0, void 0, function* () {
            const pairsBaseData = sessionStorage.getItem('pairBaseInfo');
            // console.log('pairsBaseData=', pairsBaseData)
            if (pairsBaseData !== undefined && pairsBaseData !== null) {
                this.pairsList.elementData = JSON.parse(pairsBaseData).elementData;
                this.pairsLoadComplete = true;
            }
            else {
                this.pairsBaseInfoLoading = true;
                this.pairsLoadComplete = false;
                const bridgeContract = yield (0, Web3_1.contractInstance)(VaultBridgeABI, this.bridgeAddress);
                let iNum = 0;
                let vaultAddressInContract;
                while (true) {
                    try {
                        vaultAddressInContract = yield bridgeContract.methods.getAddress(iNum).call();
                    }
                    catch (err) {
                        console.log('getting vault Address in contract occurs error:', err);
                    }
                    console.log('vaultAddressInContract', iNum, '=', vaultAddressInContract);
                    if (vaultAddressInContract === null || vaultAddressInContract === undefined || Number(vaultAddressInContract) === 0) {
                        console.log('vaultAddressInContract', iNum, ' is null,so break;');
                        break;
                    }
                    const pair = new Pairs_1.default();
                    pair.index = iNum;
                    pair.id = iNum + 1;
                    pair.vaultAddress = vaultAddressInContract;
                    console.log('vault address(', iNum, ')=', vaultAddressInContract);
                    // pair = await this.getPairsBaseInfo(pair)
                    this.pairsList.add(pair);
                    iNum++;
                }
                if (this.pairsList.size() > 0)
                    sessionStorage.setItem('pairBaseInfo', JSON.stringify(this.pairsList));
                console.log('pairsLoadComplete!!!,Pairs count:', this.pairsList.size());
                // this.pairsLoadComplete = true
                this.pairsBaseInfoLoading = false;
            }
        });
    }
    getPairsBaseInfo(pair) {
        return __awaiter(this, void 0, void 0, function* () {
            const token0 = new Token_1.default();
            const token1 = new Token_1.default();
            console.log('loadTokensInfo->vaultAddress:', pair.vaultAddress);
            // const keeperContract = await contractInstance(contractABI, pair.vaultAddress)
            // pair.poolAddress = await keeperContract.methods.poolAddress().call()
            // console.log('loadTokensInfo->poolAddress:', pair.poolAddress)
            // token0.tokenAddress = await keeperContract.methods.token0().call()
            // token1.tokenAddress = await keeperContract.methods.token1().call()
            // console.log('loadTokensInfo->tokenAddress:', token0.tokenAddress)
            // token0.tokenLendingAddress = await keeperContract.methods.CToken0().call()
            // token1.tokenLendingAddress = await keeperContract.methods.CToken1().call()
            // console.log('loadTokensInfo-token0->tokenLendingAddress:', token0.tokenLendingAddress)
            // const poolContract = await contractInstance(uniswapV3PoolABI, pair.poolAddress)
            // const token0Contract = await contractInstance(ViaLendTokenABI, token0.tokenAddress)
            // const token1Contract = await contractInstance(ViaLendTokenABI, token1.tokenAddress)
            // token0.name = await token0Contract.methods.name().call()
            // token0.symbol = await token0Contract.methods.symbol().call()
            // console.log('loadTokensInfo->symbol:', token0.symbol)
            // token0.decimals = await token0Contract.methods.decimals().call()
            // token0.iconLink = this.getIconLink(token0.symbol)
            // console.log('loadTokensInfo->token0.iconLink:', token0.iconLink)
            // token1.name = await token1Contract.methods.name().call()
            // token1.symbol = await token1Contract.methods.symbol().call()
            // token1.decimals = await token1Contract.methods.decimals().call()
            // token1.iconLink = this.getIconLink(token1.symbol)
            // console.log('loadTokensInfo->token1.iconLink:', token1.iconLink)
            // pair.tickLower = await keeperContract.methods.cLow().call()
            // console.log('loadTokensInfo->pair.cLow:', pair.tickLower)
            // pair.tickUpper = await keeperContract.methods.cHigh().call()
            // console.log('loadTokensInfo->pair.cHigh:', pair.tickUpper)
            // const slot0 = await poolContract.methods.slot0().call()
            // if (slot0 !== null && slot0 !== undefined) {
            //   pair.currentTick = slot0.tick
            //   pair.currentPrice = tickToPrice(pair.currentTick, token0.decimals, token1.decimals)
            //   if (pair.tickLower === 0 || pair.tickUpper === 0) {
            //     pair.tickLower = parseInt(pair.currentTick.toString()) - 500
            //     pair.tickUpper = parseInt(pair.currentTick.toString()) + 500
            //     console.log('slot0_tickLower0=', pair.tickLower, 'slot0_tickUpper0=', pair.tickUpper)
            //   } else {
            //     console.log('slot0_tickLower1=', pair.tickLower, ';slot0_tickUpper1=', pair.tickUpper)
            //   }
            // }
            // pair.feeTier = await poolContract.methods.fee().call()
            // console.log('loadTokensInfo->pair.feeTier:', pair.feeTier)
            pair.token0 = token0;
            pair.token1 = token1;
            return pair;
        });
    }
    getPairsSettingData(pair) {
        return __awaiter(this, void 0, void 0, function* () {
            const keeperContract = yield (0, Web3_1.contractInstance)(contractABI, pair.vaultAddress);
            pair.maxTotalSupply = yield keeperContract.methods.maxTotalSupply().call();
            pair.governanceAddress = yield keeperContract.methods.governance().call();
            pair.teamAddress = yield keeperContract.methods.team().call();
            pair.uniPortionRatio = yield keeperContract.methods.uniPortion().call();
            pair.protocolFee = yield keeperContract.methods.protocolFee().call();
        });
    }
    getIconLink(symbol) {
        let iconLink = '';
        switch (symbol.toLowerCase()) {
            case 'weth':
                iconLink = 'images/weth.png';
                break;
            case 'usdc':
                iconLink = 'images/usdc.png';
                break;
            case 'dai':
                iconLink = 'images/dai.png';
                break;
            case 'usdt':
                iconLink = 'images/usdt.png';
                break;
            case 'wbtc':
                iconLink = 'images/wbtc.png';
                break;
        }
        return iconLink;
    }
    getTokensBalance(pair) {
        return __awaiter(this, void 0, void 0, function* () {
            this.pairsBalanceLoading = true;
            const keeperContract = yield (0, Web3_1.contractInstance)(contractABI, pair.vaultAddress);
            const token0LendingContract = yield (0, Web3_1.contractInstance)(ViaLendTokenABI, pair.token0.tokenLendingAddress);
            const token1LendingContract = yield (0, Web3_1.contractInstance)(ViaLendTokenABI, pair.token1.tokenLendingAddress);
            const token0Contract = yield (0, Web3_1.contractInstance)(ViaLendTokenABI, pair.token0.tokenAddress);
            const token1Contract = yield (0, Web3_1.contractInstance)(ViaLendTokenABI, pair.token1.tokenAddress);
            // ---------- Get TVL Begin-----------------
            pair.uniliqs = yield keeperContract.methods.getPositionAmounts(BigInt(pair.tickLower), BigInt(pair.tickUpper)).call();
            if (pair.uniliqs !== undefined && pair.uniliqs !== null) {
                console.log('balance in uniswap:', pair.uniliqs, 'getVaultInfo_cLow=', pair.tickLower, 'getVaultInfo_cHigh=', pair.tickUpper);
                pair.token0.balanceInPool = pair.uniliqs.amount0;
                pair.token1.balanceInPool = pair.uniliqs.amount1;
            }
            else {
                pair.token0.balanceInPool = 0;
                pair.token1.balanceInPool = 0;
            }
            console.log('token0.balanceInPool=', pair.token0.balanceInPool);
            console.log('token1.balanceInPool=', pair.token1.balanceInPool);
            // -->Get Lending Amounts begin
            const exchangeRate0 = yield token0LendingContract.methods.exchangeRateStored().call();
            const exchangeRate1 = yield token1LendingContract.methods.exchangeRateStored().call();
            const CAmount0 = yield token0LendingContract.methods.balanceOf(pair.vaultAddress).call();
            const CAmount1 = yield token1LendingContract.methods.balanceOf(pair.vaultAddress).call();
            pair.token0.balanceInLending = (CAmount0 * exchangeRate0) / Math.pow(10, 18);
            pair.token1.balanceInLending = (CAmount1 * exchangeRate1) / Math.pow(10, 18);
            console.log('token0.balanceInLending=', pair.token0.balanceInLending);
            console.log('token1.balanceInLending=', pair.token1.balanceInLending);
            // -->Get Lending Amounts end
            pair.vaultLending = Number(pair.token0.underlying) + Number(pair.token1.underlying); // Not yet converted to USD
            pair.token0.balanceInVault = yield token0Contract.methods.balanceOf(pair.vaultAddress).call();
            pair.token1.balanceInVault = yield token1Contract.methods.balanceOf(pair.vaultAddress).call();
            console.log('token0.balanceInVault=', pair.token0.balanceInVault);
            console.log('token1.balanceInVault=', pair.token1.balanceInVault);
            // ---------- Get Tokens Balance In Wallet ---------------------
            // token0 balance in wallet
            if (this.ethereum.selectedAddress !== null && this.ethereum.selectedAddress !== undefined) {
                let balanceWei = yield token0Contract.methods.balanceOf(this.ethereum.selectedAddress).call();
                pair.token0.balanceInWallet = parseInt((balanceWei / Math.pow(10, pair.token0.decimals) * 1000).toString()) / 1000;
                // token1 balance in wallet
                balanceWei = yield token1Contract.methods.balanceOf(this.ethereum.selectedAddress).call();
                pair.token1.balanceInWallet = parseInt((balanceWei / Math.pow(10, pair.token1.decimals) * 1000).toString()) / 1000;
            }
            this.pairsBalanceLoading = false;
        });
    }
    getPairPublicData(pair) {
        return __awaiter(this, void 0, void 0, function* () {
            console.log('pair.vaultAddress=', pair.vaultAddress);
            console.log('pair.token0->tokenLendingAddress=', pair.token0.tokenLendingAddress);
            console.log('pair.token1->tokenLendingAddress=', pair.token1.tokenLendingAddress);
            console.log('pair.token0.tokenAddress=', pair.token0.tokenAddress);
            console.log('pair.token1.tokenAddress=', pair.token1.tokenAddress);
            const keeperContract = yield (0, Web3_1.contractInstance)(contractABI, pair.vaultAddress);
            const token0Contract = yield (0, Web3_1.contractInstance)(ViaLendTokenABI, pair.token0.tokenAddress);
            const token1Contract = yield (0, Web3_1.contractInstance)(ViaLendTokenABI, pair.token1.tokenAddress);
            yield this.getTokensBalance(pair);
            pair.token0.decimals = yield token0Contract.methods.decimals().call();
            pair.token1.decimals = yield token1Contract.methods.decimals().call();
            pair.tvlTotal0 = (Number(pair.token0.balanceInVault) + Number(pair.token0.balanceInPool) + Number(pair.token0.balanceInLending)) / Number(Math.pow(10, pair.token0.decimals));
            pair.tvlTotal1 = (Number(pair.token1.balanceInVault) + Number(pair.token1.balanceInPool) + Number(pair.token1.balanceInLending)) / Number(Math.pow(10, pair.token1.decimals));
            console.log('tvlTotal0=', pair.tvlTotal0, 'tvlTotal1=', pair.tvlTotal1);
            const totalUniswap = Number(pair.token0.balanceInPool) * 300 + Number(pair.token1.balanceInPool);
            const totalLending = Number(pair.token0.balanceInLending) * 300 + Number(pair.token1.balanceInLending);
            const totalUsdc = totalUniswap + totalLending;
            if (totalUsdc === 0) {
                pair.uniswapRatio = 0;
                pair.lendingRatio = 0;
            }
            else {
                pair.uniswapRatio = (totalUniswap / totalUsdc) * 100;
                pair.lendingRatio = (totalLending / totalUsdc) * 100;
            }
            if (Number(pair.token0.balanceInPool) + Number(pair.token0.balanceInLending) === 0) {
                pair.uniToken0Rate = 0;
            }
            else {
                pair.uniToken0Rate = Number(pair.token0.balanceInPool) / (Number(pair.token0.balanceInPool) + Number(pair.token0.balanceInLending));
            }
            if (Number(pair.token1.balanceInPool) + Number(pair.token1.balanceInLending) === 0) {
                pair.uniToken1Rate = 0;
            }
            else {
                pair.uniToken1Rate = Number(pair.token1.balanceInPool) / (Number(pair.token1.balanceInPool) + Number(pair.token1.balanceInLending));
            }
            if (Number(pair.token0.balanceInPool) + Number(pair.token0.balanceInLending) === 0) {
                pair.lendingToken0Rate = 0;
            }
            else {
                pair.lendingToken0Rate = Number(pair.token0.balanceInLending) / (Number(pair.token0.balanceInPool) + Number(pair.token0.balanceInLending));
            }
            if (Number(pair.token1.balanceInPool) + Number(pair.token1.balanceInLending) === 0) {
                pair.lendingToken1Rate = 0;
            }
            else {
                pair.lendingToken1Rate = Number(pair.token1.balanceInLending) / (Number(pair.token1.balanceInPool) + Number(pair.token1.balanceInLending));
            }
            // console.log('totalUniswap=', pair.totalUniswap, 'totalLending=', pair.totalLending, 'total_usdc=', pair.totalUsdc, 'uniswapRatio=', pair.uniswapRatio, 'lendingRatio=', pair.lendingRatio)
            // ---------- Get Assets ratio End --------------
            console.log('ethereum.selectedAddress=', this.ethereum.selectedAddress);
            if (this.ethereum.selectedAddress !== null && this.ethereum.selectedAddress !== undefined) {
                pair.myShare = yield keeperContract.methods.balanceOf(this.ethereum.selectedAddress).call();
                console.log('pair.myShare=', pair.myShare);
                pair.totalShares = yield keeperContract.methods.totalSupply().call();
                pair.assets = yield keeperContract.methods.Assetholder(this.ethereum.selectedAddress).call();
                if (Number(pair.totalShares) === 0) {
                    pair.token0.myTVL = 0;
                    pair.token1.myTVL = 0;
                }
                else {
                    pair.token0.myTVL = (pair.tvlTotal0 * Number(pair.myShare)) / Number(pair.totalShares);
                    pair.token1.myTVL = (pair.tvlTotal1 * Number(pair.myShare)) / Number(pair.totalShares);
                }
                // Get token approve status
                const allowA = yield token0Contract.methods.allowance(this.ethereum.selectedAddress, pair.vaultAddress).call();
                const allowB = yield token1Contract.methods.allowance(this.ethereum.selectedAddress, pair.vaultAddress).call();
                console.log('allowA=', allowA, 'allowB=', allowB);
                if (allowA > 0) {
                    pair.token0.tokenApproved = true;
                }
                else {
                    pair.token0.tokenApproved = false;
                }
                if (allowB > 0) {
                    pair.token1.tokenApproved = true;
                }
                else {
                    pair.token1.tokenApproved = false;
                }
                if (this.calculateAPY)
                    yield this.calcAPY(keeperContract, pair);
                if (this.calculateAPR)
                    yield this.calcAPR(keeperContract, pair);
            }
            pair.gettingData = false;
            pair.loadDataCompleted = true;
            sessionStorage.setItem(pair.token0.symbol.concat('-', pair.token1.symbol), JSON.stringify(pair));
            return pair;
        });
    }
    // Calculate APY
    calcAPY(keeperContract, pair) {
        return __awaiter(this, void 0, void 0, function* () {
            if (pair.assets !== null) {
                // get oraclePriceTwap
                let oraclePriceTwap;
                const poolAddress = yield keeperContract.methods.poolAddress().call();
                console.log('poolAddress=', poolAddress);
                // const twapDuration = 2
                // this.sleep(3000)
                const keeperUniswapV3Contract = (0, Web3_1.contractInstance)(uniswapV3PoolABI, poolAddress);
                const slot0 = yield keeperUniswapV3Contract.methods.slot0().call();
                if (slot0 !== null && slot0 !== undefined) {
                    const twap = slot0.tick;
                    oraclePriceTwap = yield keeperContract.methods.getQuoteAtTick(Number(twap), BigInt(Math.pow(10, 18)), pair.token0.tokenAddress, pair.token1.tokenAddress).call();
                    console.log('twap=', twap, ';oraclePriceTwap=', oraclePriceTwap);
                }
                const oneyearINsec = 365 * 24 * 60 * 60;
                const block = yield this.web3.eth.getBlock(pair.assets.block);
                console.log('block timestamp=', block.timestamp);
                const timestamp = block.timestamp;
                const headerNumber = yield this.web3.eth.getBlockNumber();
                const headerBlock = yield this.web3.eth.getBlock(headerNumber);
                // console.log('header=', JSON.stringify(headerBlock))
                const htimestamp = headerBlock.timestamp;
                console.log('htimestamp=', htimestamp, ';timestamp=', timestamp);
                const timediff = Number(htimestamp) - Number(timestamp);
                const deposit0 = pair.assets.deposit0;
                const deposit1 = pair.assets.deposit1;
                const fd0 = Number(deposit0);
                let fd1, fm1;
                if (oraclePriceTwap === 0) {
                    fd1 = 0;
                    fm1 = 0;
                }
                else {
                    fd1 = (Number(deposit1) * Math.pow(10, Number(pair.token0.decimals))) / oraclePriceTwap;
                    fm1 = (Number(pair.token1.myTVL) * Math.pow(10, Number(pair.token0.decimals))) / oraclePriceTwap;
                }
                const fm0 = Number(pair.token0.myTVL);
                const fdd = fd0 + fd1;
                const fmm = fm0 + fm1;
                console.log('fm0=', fm0, 'fm1=', fm1);
                console.log('decimal=', pair.token0.decimals);
                console.log('fmm=', fmm, 'fdd=', fdd, 'timediff=', timediff, 'oneyearInsec=', oneyearINsec);
                // if (fdd === 0 || timediff === 0) {
                //   pair.currentAPR = 0
                // } else {
                //   pair.currentAPR = (fmm - fdd) / Number(timediff) * oneyearINsec / fdd
                // }
                const myDepositInToken1 = Number(deposit0) * Number(oraclePriceTwap) + Number(deposit1);
                console.log('deposit0=', deposit0, 'deposit1=', deposit1, 'myDepositInToken1=', myDepositInToken1);
                pair.currentDeposits = myDepositInToken1;
                const fees = yield keeperContract.methods.Fees().call();
                let myFeesInToken1 = 0;
                if (fees != null) {
                    let myFees0 = 0;
                    let myFees1 = 0;
                    console.log('fees=', JSON.stringify(fees));
                    console.log('totalShares=', pair.totalShares);
                    if (Number(pair.totalShares) !== 0) {
                        // this.userFee0Total += (Number(fees.U3Fees0) + Number(fees.LcFees0)) * pair.myShare / pair.totalShares / Math.pow(10, Number(pair.token0.decimals))
                        // this.userFee1Total += (Number(fees.U3Fees1) + Number(fees.LcFees1)) * pair.myShare / pair.totalShares / Math.pow(10, Number(pair.token1.decimals))
                        // myFee0 <-> this.userFee0Total
                        myFees0 = ((Number(fees.U3Fees0) + Number(fees.LcFees0)) * pair.myShare) / pair.totalShares;
                        myFees1 = ((Number(fees.U3Fees1) + Number(fees.LcFees1)) * pair.myShare) / pair.totalShares;
                        myFeesInToken1 = Number(myFees0) * Number(oraclePriceTwap) + Number(myFees1);
                    }
                }
                // this.netAPYTotal += Number(this.pairsList.get(i).currentAPR)
                if (Number(timediff) !== 0 && Number(myDepositInToken1) !== 0) {
                    pair.netAPY += (((Number(myFeesInToken1) / Number(timediff)) * oneyearINsec) / myDepositInToken1) * 100;
                }
            }
        });
    }
    // Calculate APR
    calcAPR(keeperContract, pair) {
        return __awaiter(this, void 0, void 0, function* () {
            if (pair.assets !== null) {
                // calc APR
                let oraclePriceTwap;
                const poolAddress = yield keeperContract.methods.poolAddress().call();
                const keeperUniswapV3Contract = (0, Web3_1.contractInstance)(uniswapV3PoolABI, poolAddress);
                const slot0 = yield keeperUniswapV3Contract.methods.slot0().call();
                if (slot0 !== null && slot0 !== undefined) {
                    const twap = slot0.tick;
                    oraclePriceTwap = yield keeperContract.methods.getQuoteAtTick(Number(twap), BigInt(Math.pow(10, 18)), pair.token0.tokenAddress, pair.token1.tokenAddress).call();
                    console.log('twap=', twap, ';oraclePriceTwap=', oraclePriceTwap);
                }
                const Total0 = Number(pair.token0.balanceInVault) +
                    Number(pair.uniliqs.amount0) +
                    Number(pair.token0.underlying);
                const Total1 = Number(pair.token1.balanceInVault) +
                    Number(pair.uniliqs.amount1) +
                    Number(pair.token1.underlying);
                let mytvl0, mytvl1;
                if (Number(pair.totalShares) === 0) {
                    mytvl0 = 0;
                    mytvl1 = 0;
                }
                else {
                    mytvl0 = (Total0 * Number(pair.myShare)) / Number(pair.totalShares);
                    mytvl1 = (Total1 * Number(pair.myShare)) / Number(pair.totalShares);
                }
                const oneyearINsec = 365 * 24 * 60 * 60;
                const block = yield this.web3.eth.getBlock(pair.assets.block);
                console.log('block timestamp=', block.timestamp);
                const timestamp = block.timestamp;
                const headerNumber = yield this.web3.eth.getBlockNumber();
                const headerBlock = yield this.web3.eth.getBlock(headerNumber);
                // console.log('header=', JSON.stringify(headerBlock))
                const htimestamp = headerBlock.timestamp;
                console.log('htimestamp=', htimestamp, ';timestamp=', timestamp);
                const timediff = Number(htimestamp) - Number(timestamp);
                const deposit0 = pair.assets.deposit0;
                const deposit1 = pair.assets.deposit1;
                const fd0 = Number(deposit0);
                let fd1, fm1;
                if (oraclePriceTwap === 0) {
                    fd1 = 0;
                    fm1 = 0;
                }
                else {
                    fd1 =
                        (Number(deposit1) * Math.pow(10, Number(pair.token0.decimals))) /
                            oraclePriceTwap;
                    fm1 =
                        (Number(mytvl1) * Math.pow(10, Number(pair.token0.decimals))) /
                            oraclePriceTwap;
                }
                const fm0 = Number(mytvl0);
                const fdd = fd0 + fd1;
                const fmm = fm0 + fm1;
                console.log('fm0=', fm0, 'fm1=', fm1);
                console.log('decimal=', pair.token0.decimals);
                console.log('fmm=', fmm, 'fdd=', fdd, 'timediff=', timediff, 'oneyearInsec=', oneyearINsec);
                if (fdd === 0 || timediff === 0) {
                    pair.currentAPR = 0;
                }
                else {
                    pair.currentAPR = (((fmm - fdd) / Number(timediff)) * oneyearINsec) / fdd;
                }
            }
        });
    }
}
exports.default = PairsData;

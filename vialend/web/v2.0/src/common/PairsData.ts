// import { Vue } from 'vue-property-decorator'
import { contractInstance, getWeb3Instance } from './Web3'
import Pairs from '../model/Pairs'
import ArrayList from './ArrayList'
import store from '@/store'
// import axios from 'axios'
import { priceToTick, tickToPrice } from './Tools'
import { CHAININFO, checkChain } from '@/constants/chains'

const VaultBridgeABI = require('../abi/VaultBridge.json')
const contractABI = require('../abi/ViaLendFeeMakerABI.json')
const uniswapV3PoolABI = require('../abi/UniswapV3PoolABI.json')
const ViaLendTokenABI = require('../abi/VialendTokenABI.json')
const VaultFactoryABI = require('../abi/VaultFactory.json')
const VaultStrategyABI = require('../abi/VaultStrategy.json')

class PairsData {
    pairsList:ArrayList<Pairs>;
    pairsBaseInfoLoading:boolean;
    pairsBalanceLoading:boolean;
    pairsLoadComplete:boolean;
    calculateAPY:boolean;
    calculateAPR:boolean;
    ethereum:any;
    web3:any;
    factoryContract:object;
    factoryAddress:string;
    error:string;

    constructor() {
      this.pairsList = new ArrayList<Pairs>()
      this.pairsBaseInfoLoading = false
      this.pairsBalanceLoading = false
      this.pairsLoadComplete = false
      this.calculateAPY = false
      this.calculateAPR = false
      this.ethereum = (window as any).ethereum
      this.web3 = getWeb3Instance()
      this.factoryAddress = ''
      this.factoryContract = {}
      this.error = ''
    }

    async loadPairsInfo() {
      console.log('CHAININFO chainId=', store.state.chainId)
      console.log('CHAININFO bridgeAddress=', CHAININFO[store.state.chainId].bridgeAddress)
      console.log('CheckChainIDValid=', await checkChain())
      if (!this.pairsBaseInfoLoading) {
        this.pairsBaseInfoLoading = true
        const pairsBaseData = await store.dispatch('getSessionData', { key: 'pairBaseInfo' })
        // console.log('pairsBaseData=', pairsBaseData)
        if (pairsBaseData !== undefined && pairsBaseData !== null) {
          this.pairsList.elementData = JSON.parse(pairsBaseData).elementData
          let pair = new Pairs()
          for (let i = 0; i < this.pairsList.size(); i++) {
            pair = this.pairsList.get(i)
            pair = await this.getPairStatus(pair)
            this.pairsList.set(i, pair)
          }
          this.pairsBaseInfoLoading = false
          this.pairsLoadComplete = true
          // console.log('pairsBaseData has value,pairsBaseData=', JSON.stringify(pairsBaseData))
        } else {
          this.pairsLoadComplete = false
          console.log('pairsBaseData is null')
          const bridgeContract = await contractInstance(VaultBridgeABI, CHAININFO[store.state.chainId].bridgeAddress)
          let iNum = 0
          let vaultAddressInContract
          while (true) {
            try {
              if (iNum === 0) {
                this.factoryAddress = await bridgeContract.methods.getAddress(iNum).call()
                await store.dispatch('setSessionData', { key: 'factoryAddress', value: this.factoryAddress })
                console.log('factoryAddress=', this.factoryAddress)
              } else {
                vaultAddressInContract = await bridgeContract.methods.getAddress(iNum).call()
                console.log('vaultAddressInContract', iNum, '=', vaultAddressInContract)
                if (vaultAddressInContract === null || vaultAddressInContract === undefined || Number(vaultAddressInContract) === 0) {
                  console.log('vaultAddressInContract', iNum, ' is null,so break;')
                  break
                }
              }
              if (iNum > 0) {
                let pair = new Pairs()
                pair.index = iNum - 1
                pair.id = iNum
                pair.vaultAddress = vaultAddressInContract
                console.log('vault address(', iNum, ')=', vaultAddressInContract)
                pair = await this.getPairsBaseInfo(pair)
                pair = await this.getPairStatus(pair)
                this.pairsList.add(pair)
              }
              iNum++
            } catch (err) {
              console.log('getting vault Address in contract occurs error:', err)
              this.error = 'An error occurred while loading pairs information.'
              break
            }
          }
          // if (this.pairsList.size() > 0) sessionStorage.setItem('pairBaseInfo', JSON.stringify(this.pairsList))
          if (this.pairsList.size() > 0 && this.error === '') await store.dispatch('setSessionData', { key: 'pairBaseInfo', value: JSON.stringify(this.pairsList) })
          console.log('pairsLoadComplete!!!,Pairs count:', this.pairsList.size())
          this.pairsLoadComplete = true
          this.pairsBaseInfoLoading = false
        }
      }
    }

    async getPairStatus(pair:Pairs) {
      if (this.factoryAddress !== undefined && this.factoryAddress !== '') {
        const factoryContract = await contractInstance(VaultFactoryABI, this.factoryAddress)
        pair.stat = await factoryContract.methods.stat(pair.strategyAddress, pair.vaultAddress).call()
        if (Number(pair.stat) === 1) {
          pair.disabled = false
        } else {
          pair.disabled = true
        }
      }
      return pair
    }

    async getPairsBaseInfo(pair:Pairs) {
      console.log('loadTokensInfo->vaultAddress:', pair.vaultAddress)
      if (this.factoryAddress !== undefined && this.factoryAddress !== '') {
        this.factoryContract = await contractInstance(VaultFactoryABI, this.factoryAddress)
        console.log('pair.vaultAddress=:', pair.vaultAddress)
        pair.strategyAddress = await (this.factoryContract as any).methods.getPair0(pair.vaultAddress).call()
        console.log('strategyAddress=', pair.strategyAddress)
        const strategyContract = await contractInstance(VaultStrategyABI, pair.strategyAddress)
        pair.poolAddress = await strategyContract.methods.pool().call()
        console.log('loadTokensInfo->poolAddress:', pair.poolAddress)
        pair.token0.tokenAddress = await strategyContract.methods.token0().call()
        pair.token1.tokenAddress = await strategyContract.methods.token1().call()
        console.log('loadTokensInfo->token0Address:', pair.token0.tokenAddress)
        console.log('loadTokensInfo->token1Address:', pair.token1.tokenAddress)
        pair.token0.tokenLendingAddress = await strategyContract.methods._CTOKEN(pair.token0.tokenAddress).call()
        pair.token1.tokenLendingAddress = await strategyContract.methods._CTOKEN(pair.token1.tokenAddress).call()
        console.log('loadTokensInfo-token0->tokenLendingAddress:', pair.token0.tokenLendingAddress)
        console.log('loadTokensInfo-token1->tokenLendingAddress:', pair.token1.tokenLendingAddress)
        pair.tickLower = await strategyContract.methods.cLow().call()
        console.log('loadTokensInfo->pair.cLow:', pair.tickLower)
        pair.tickUpper = await strategyContract.methods.cHigh().call()
        console.log('loadTokensInfo->pair.cHigh:', pair.tickUpper)
        // const keeperContract = await contractInstance(contractABI, pair.vaultAddress)
        const poolContract = await contractInstance(uniswapV3PoolABI, pair.poolAddress)
        const token0Contract = await contractInstance(ViaLendTokenABI, pair.token0.tokenAddress)
        const token1Contract = await contractInstance(ViaLendTokenABI, pair.token1.tokenAddress)
        pair.token0.name = await token0Contract.methods.name().call()
        pair.token0.symbol = await token0Contract.methods.symbol().call()
        console.log('loadTokensInfo->symbol:', pair.token0.symbol)
        pair.token0.decimals = await token0Contract.methods.decimals().call()
        pair.token0.iconLink = this.getIconLink(pair.token0.symbol)
        console.log('loadTokensInfo->token0.iconLink:', pair.token0.iconLink)
        pair.token1.name = await token1Contract.methods.name().call()
        pair.token1.symbol = await token1Contract.methods.symbol().call()
        pair.token1.decimals = await token1Contract.methods.decimals().call()
        pair.token1.iconLink = this.getIconLink(pair.token1.symbol)
        console.log('loadTokensInfo->token1.iconLink:', pair.token1.iconLink)
        // Get pair tickLower and tickUpper
        pair = await this.getTickInfo(pair)
        // Get token approve status
        pair = await this.getTokenApproveStatus(pair)
        pair.feeTier = await poolContract.methods.fee().call()
        console.log('loadTokensInfo->pair.feeTier:', pair.feeTier)
      }
      return pair
    }

    async getTickInfo(pair: Pairs) {
      const poolContract = await contractInstance(uniswapV3PoolABI, pair.poolAddress)
      const slot0 = await poolContract.methods.slot0().call()
      if (slot0 !== null && slot0 !== undefined) {
        pair.currentTick = slot0.tick
        pair.currentPrice = tickToPrice(pair.currentTick, pair.token0.decimals, pair.token1.decimals)
        if (pair.tickLower === 0 || pair.tickUpper === 0) {
          pair.tickLower = parseInt(pair.currentTick.toString()) - 500
          pair.tickUpper = parseInt(pair.currentTick.toString()) + 500
          console.log('slot0_tickLower0=', pair.tickLower, 'slot0_tickUpper0=', pair.tickUpper)
        } else {
          console.log('slot0_tickLower1=', pair.tickLower, ';slot0_tickUpper1=', pair.tickUpper)
        }
      } else {
        console.log('slot0 is null')
      }
      return pair
    }

    async getTokenApproveStatus(pair: Pairs) {
      if (this.ethereum.selectedAddress !== null && this.ethereum.selectedAddress !== undefined) {
        const token0Contract = await contractInstance(ViaLendTokenABI, pair.token0.tokenAddress)
        const token1Contract = await contractInstance(ViaLendTokenABI, pair.token1.tokenAddress)
        const allowA = await token0Contract.methods.allowance(this.ethereum.selectedAddress, pair.vaultAddress).call()
        const allowB = await token1Contract.methods.allowance(this.ethereum.selectedAddress, pair.vaultAddress).call()
        console.log('allowA=', allowA, 'allowB=', allowB)
        if (allowA > 0) {
          pair.token0.tokenApproved = true
        } else {
          pair.token0.tokenApproved = false
        }
        if (allowB > 0) {
          pair.token1.tokenApproved = true
        } else {
          pair.token1.tokenApproved = false
        }
      }
      return pair
    }

    async getPairsSettingData(pair: Pairs) {
      const keeperContract = await contractInstance(contractABI, pair.vaultAddress)
      pair.maxTotalSupply = await keeperContract.methods.maxTotalSupply().call()
      pair.governanceAddress = await keeperContract.methods.governance().call()
      pair.teamAddress = await keeperContract.methods.team().call()
      pair.uniPortionRatio = await keeperContract.methods.uniPortion().call()
      pair.protocolFee = await keeperContract.methods.protocolFee().call()
    }

    getIconLink(symbol:string) {
      let iconLink = ''
      switch (symbol.toLowerCase()) {
        case 'weth':
          iconLink = 'images/weth.png'
          break
        case 'usdc':
          iconLink = 'images/usdc.png'
          break
        case 'dai':
          iconLink = 'images/dai.png'
          break
        case 'usdt':
          iconLink = 'images/usdt.png'
          break
        case 'wbtc':
          iconLink = 'images/wbtc.png'
          break
      }
      return iconLink
    }

    async getTokensBalance(pair: Pairs) {
      this.pairsBalanceLoading = true
      // console.log('contractABI=', contractABI)
      console.log('vault Address:', pair.vaultAddress)
      const strategyContract = await contractInstance(VaultStrategyABI, pair.strategyAddress)
      const token0LendingContract = await contractInstance(ViaLendTokenABI, pair.token0.tokenLendingAddress)
      const token1LendingContract = await contractInstance(ViaLendTokenABI, pair.token1.tokenLendingAddress)
      const token0Contract = await contractInstance(ViaLendTokenABI, pair.token0.tokenAddress)
      const token1Contract = await contractInstance(ViaLendTokenABI, pair.token1.tokenAddress)
      // ---------- Get TVL Begin-----------------
      console.log('getTokensBalance->ticklower:', pair.tickLower, ';tickupper:', pair.tickUpper)
      pair.uniliqs = await strategyContract.methods.getUniAmounts(BigInt(pair.tickLower), BigInt(pair.tickUpper)).call()
      console.log('pair.uniliqs=', pair.uniliqs)

      if (pair.uniliqs !== undefined && pair.uniliqs !== null) {
        console.log('balance in uniswap:', pair.uniliqs, 'getVaultInfo_cLow=', pair.tickLower, 'getVaultInfo_cHigh=', pair.tickUpper)
        pair.token0.balanceInPool = (pair.uniliqs as any).amount0
        pair.token1.balanceInPool = (pair.uniliqs as any).amount1
      } else {
        pair.token0.balanceInPool = 0
        pair.token1.balanceInPool = 0
      }
      console.log('token0.balanceInPool=', pair.token0.balanceInPool)
      console.log('token1.balanceInPool=', pair.token1.balanceInPool)
      // -->Get Lending Amounts begin
      const exchangeRate0 = await token0LendingContract.methods.exchangeRateStored().call()
      const exchangeRate1 = await token1LendingContract.methods.exchangeRateStored().call()
      const CAmount0 = await token0LendingContract.methods.balanceOf(pair.vaultAddress).call()
      const CAmount1 = await token1LendingContract.methods.balanceOf(pair.vaultAddress).call()
      pair.token0.balanceInLending = (CAmount0 * exchangeRate0) / Math.pow(10, 18)
      pair.token1.balanceInLending = (CAmount1 * exchangeRate1) / Math.pow(10, 18)
      console.log('token0.balanceInLending=', pair.token0.balanceInLending)
      console.log('token1.balanceInLending=', pair.token1.balanceInLending)
      // -->Get Lending Amounts end
      pair.vaultLending = Number(pair.token0.underlying) + Number(pair.token1.underlying) // Not yet converted to USD
      pair.token0.balanceInVault = await token0Contract.methods.balanceOf(pair.vaultAddress).call()
      pair.token1.balanceInVault = await token1Contract.methods.balanceOf(pair.vaultAddress).call()
      console.log('token0.balanceInVault=', pair.token0.balanceInVault)
      console.log('token1.balanceInVault=', pair.token1.balanceInVault)

      if (this.ethereum.selectedAddress !== null && this.ethereum.selectedAddress !== undefined) {
        // ---------- Get Tokens Balance In Wallet ---------------------
      // token0 balance in wallet
        let balanceWei:number = await token0Contract.methods.balanceOf(this.ethereum.selectedAddress).call()
        pair.token0.balanceInWallet = parseInt((balanceWei / Math.pow(10, pair.token0.decimals) * 1000).toString()) / 1000
        // token1 balance in wallet
        balanceWei = await token1Contract.methods.balanceOf(this.ethereum.selectedAddress).call()
        pair.token1.balanceInWallet = parseInt((balanceWei / Math.pow(10, pair.token1.decimals) * 1000).toString()) / 1000
      }
      this.pairsBalanceLoading = false
    }

    async getPairPublicData(pair: Pairs) {
      pair.gettingData = true
      console.log('pair.vaultAddress=', pair.vaultAddress)
      console.log('pair.token0->tokenLendingAddress=', pair.token0.tokenLendingAddress)
      console.log('pair.token1->tokenLendingAddress=', pair.token1.tokenLendingAddress)
      console.log('pair.token0.tokenAddress=', pair.token0.tokenAddress)
      console.log('pair.token1.tokenAddress=', pair.token1.tokenAddress)
      const strategyContract = await contractInstance(VaultStrategyABI, pair.strategyAddress)
      const keeperContract = await contractInstance(contractABI, pair.vaultAddress)
      const token0Contract = await contractInstance(ViaLendTokenABI, pair.token0.tokenAddress)
      const token1Contract = await contractInstance(ViaLendTokenABI, pair.token1.tokenAddress)
      await this.getTokensBalance(pair)
      pair.token0.decimals = await token0Contract.methods.decimals().call()
      pair.token1.decimals = await token1Contract.methods.decimals().call()
      pair.tvlTotal0 = (Number(pair.token0.balanceInVault) + Number(pair.token0.balanceInPool) + Number(pair.token0.balanceInLending)) / Number(Math.pow(10, pair.token0.decimals))
      pair.tvlTotal1 = (Number(pair.token1.balanceInVault) + Number(pair.token1.balanceInPool) + Number(pair.token1.balanceInLending)) / Number(Math.pow(10, pair.token1.decimals))
      console.log('tvlTotal0=', pair.tvlTotal0, 'tvlTotal1=', pair.tvlTotal1)
      const totalUniswap = Number(pair.token0.balanceInPool) * 300 + Number(pair.token1.balanceInPool)
      const totalLending = Number(pair.token0.balanceInLending) * 300 + Number(pair.token1.balanceInLending)
      const totalUsdc = totalUniswap + totalLending
      if (totalUsdc === 0) {
        pair.uniswapRatio = 0
        pair.lendingRatio = 0
      } else {
        pair.uniswapRatio = (totalUniswap / totalUsdc) * 100
        pair.lendingRatio = (totalLending / totalUsdc) * 100
      }
      if (Number(pair.token0.balanceInPool) + Number(pair.token0.balanceInLending) === 0) {
        pair.uniToken0Rate = 0
      } else {
        pair.uniToken0Rate = Number(pair.token0.balanceInPool) / (Number(pair.token0.balanceInPool) + Number(pair.token0.balanceInLending))
      }
      if (Number(pair.token1.balanceInPool) + Number(pair.token1.balanceInLending) === 0) {
        pair.uniToken1Rate = 0
      } else {
        pair.uniToken1Rate = Number(pair.token1.balanceInPool) / (Number(pair.token1.balanceInPool) + Number(pair.token1.balanceInLending))
      }
      if (Number(pair.token0.balanceInPool) + Number(pair.token0.balanceInLending) === 0) {
        pair.lendingToken0Rate = 0
      } else {
        pair.lendingToken0Rate = Number(pair.token0.balanceInLending) / (Number(pair.token0.balanceInPool) + Number(pair.token0.balanceInLending))
      }
      if (Number(pair.token1.balanceInPool) + Number(pair.token1.balanceInLending) === 0) {
        pair.lendingToken1Rate = 0
      } else {
        pair.lendingToken1Rate = Number(pair.token1.balanceInLending) / (Number(pair.token1.balanceInPool) + Number(pair.token1.balanceInLending))
      }
      // console.log('totalUniswap=', pair.totalUniswap, 'totalLending=', pair.totalLending, 'total_usdc=', pair.totalUsdc, 'uniswapRatio=', pair.uniswapRatio, 'lendingRatio=', pair.lendingRatio)
      // ---------- Get Assets ratio End --------------
      console.log('ethereum.selectedAddress=', this.ethereum.selectedAddress)
      if (this.ethereum.selectedAddress !== null && this.ethereum.selectedAddress !== undefined) {
        pair.myShare = await keeperContract.methods.balanceOf(this.ethereum.selectedAddress).call()
        console.log('pair.myShare=', pair.myShare)
        pair.totalShares = await keeperContract.methods.totalSupply().call()
        pair.assets = await keeperContract.methods.Assetholder(this.ethereum.selectedAddress).call()
        if (Number(pair.totalShares) === 0) {
          pair.token0.myTVL = 0
          pair.token1.myTVL = 0
        } else {
          pair.token0.myTVL = (pair.tvlTotal0 * Number(pair.myShare)) / Number(pair.totalShares)
          pair.token1.myTVL = (pair.tvlTotal1 * Number(pair.myShare)) / Number(pair.totalShares)
        }

        console.log('calculateAPY:', this.calculateAPY, ';calculateAPR:', this.calculateAPR)

        if (this.calculateAPY) await this.calcAPY(strategyContract, pair)
        if (this.calculateAPR) await this.calcAPR(strategyContract, pair)
      }
      pair.gettingData = false
      pair.loadDataCompleted = true
      console.log('setSessionData->pairsSymbol:start')
      await store.dispatch('setSessionData', { key: pair.token0.symbol.concat('-', pair.token1.symbol), value: JSON.stringify(pair) })
      // sessionStorage.setItem(pair.token0.symbol.concat('-', pair.token1.symbol), JSON.stringify(pair))
      console.log('setSessionData->pairsSymbol:done')

      return pair
    }

    // Calculate APY
    async calcAPY(strategyContract: any, pair: Pairs) {
      if (pair.assets !== null) {
        // get oraclePriceTwap
        let oraclePriceTwap = 0
        console.log('pair.assets=', pair.assets)
        const poolAddress = await strategyContract.methods.pool().call()
        console.log('poolAddress=', poolAddress)
        // const twapDuration = 2
        if (pair.currentTick !== null && pair.currentTick !== undefined) {
          console.log('twap=', pair.currentTick)
          oraclePriceTwap = await strategyContract.methods.getQuoteAtTick(Number(pair.currentTick), BigInt(Math.pow(10, 18)), pair.token0.tokenAddress, pair.token1.tokenAddress).call()
          console.log(';oraclePriceTwap=', oraclePriceTwap)
        }
        const oneyearINsec = 365 * 24 * 60 * 60
        const block = await this.web3.eth.getBlock((pair.assets as any).block)
        console.log('block timestamp=', block.timestamp)
        const timestamp = block.timestamp
        const headerNumber = await this.web3.eth.getBlockNumber()
        const headerBlock = await this.web3.eth.getBlock(headerNumber)
        // console.log('header=', JSON.stringify(headerBlock))
        const htimestamp = headerBlock.timestamp
        console.log('htimestamp=', htimestamp, ';timestamp=', timestamp)
        const timediff = Number(htimestamp) - Number(timestamp)
        const deposit0 = (pair.assets as any).deposit0
        const deposit1 = (pair.assets as any).deposit1
        const fd0 = Number(deposit0)
        let fd1, fm1
        if (oraclePriceTwap === 0) {
          fd1 = 0
          fm1 = 0
        } else {
          fd1 = (Number(deposit1) * Math.pow(10, Number(pair.token0.decimals))) / oraclePriceTwap
          fm1 = (Number(pair.token1.myTVL) * Math.pow(10, Number(pair.token0.decimals))) / oraclePriceTwap
        }
        const fm0 = Number(pair.token0.myTVL)
        const fdd = fd0 + fd1
        const fmm = fm0 + fm1
        console.log('fm0=', fm0, 'fm1=', fm1)
        console.log('decimal=', pair.token0.decimals)
        console.log('fmm=', fmm, 'fdd=', fdd, 'timediff=', timediff, 'oneyearInsec=', oneyearINsec)
        // if (fdd === 0 || timediff === 0) {
        //   pair.currentAPR = 0
        // } else {
        //   pair.currentAPR = (fmm - fdd) / Number(timediff) * oneyearINsec / fdd
        // }
        const myDepositInToken1 = Number(deposit0) * Number(oraclePriceTwap) + Number(deposit1)
        console.log('deposit0=', deposit0, 'deposit1=', deposit1, 'myDepositInToken1=', myDepositInToken1)
        if (this.ethereum.selectedAddress !== null && this.ethereum.selectedAddress !== undefined) {
          const price = await strategyContract.methods.getPrice().call()

          const token0Contract = await contractInstance(ViaLendTokenABI, pair.token0.tokenAddress)
          const token1Contract = await contractInstance(ViaLendTokenABI, pair.token1.tokenAddress)
          const token0Amount = await token0Contract.methods.balanceOf(this.ethereum.selectedAddress).call()
          const token1Amount = await token1Contract.methods.balanceOf(this.ethereum.selectedAddress).call()
          console.log('amount0=', token0Amount)
          console.log('amount1=', token1Amount)
          console.log('price=', price)
          console.log('amount0 * price', token0Amount * price)
          console.log('amount0+amount1=', token0Amount + token1Amount)
          console.log('with decimals', token0Amount / Math.pow(10, Number(pair.token0.decimals)) + token1Amount / Math.pow(10, Number(pair.token1.decimals)))

          const totalAmount = token0Amount * price / Math.pow(10, Number(pair.token0.decimals)) + token1Amount
          console.log('totalAmount=', totalAmount)
          const currDeposits = totalAmount / Math.pow(10, Number(pair.token1.decimals))
          console.log('currDeposits=', currDeposits, ';decimal=', Math.pow(10, Number(pair.token1.decimals)))

          // amount0 = weth . balanceOf(user address)
          // amount1= usdc.balanceOf(user address)
          // totalInUSDC = amount0 * price + amount1
          // current deposit display = totalInUSDC / 10^(decimals(usdc))
          pair.currentDeposits = currDeposits
        }
        //* ****************Fee function is to be developed ***********************/
        // const fees = await keeperContract.methods.Fees().call()
        // let myFeesInToken1 = 0
        // if (fees != null) {
        //   let myFees0 = 0
        //   let myFees1 = 0
        //   console.log('fees=', JSON.stringify(fees))
        //   console.log('totalShares=', pair.totalShares)
        //   if (Number(pair.totalShares) !== 0) {
        //     // this.userFee0Total += (Number(fees.U3Fees0) + Number(fees.LcFees0)) * pair.myShare / pair.totalShares / Math.pow(10, Number(pair.token0.decimals))
        //     // this.userFee1Total += (Number(fees.U3Fees1) + Number(fees.LcFees1)) * pair.myShare / pair.totalShares / Math.pow(10, Number(pair.token1.decimals))
        //     // myFee0 <-> this.userFee0Total
        //     myFees0 = ((Number(fees.U3Fees0) + Number(fees.LcFees0)) * pair.myShare) / pair.totalShares
        //     myFees1 = ((Number(fees.U3Fees1) + Number(fees.LcFees1)) * pair.myShare) / pair.totalShares
        //     myFeesInToken1 = Number(myFees0) * Number(oraclePriceTwap) + Number(myFees1)
        //   }
        // }
        // // this.netAPYTotal += Number(this.pairsList.get(i).currentAPR)
        // if (Number(timediff) !== 0 && Number(myDepositInToken1) !== 0) {
        //   pair.netAPY += (((Number(myFeesInToken1) / Number(timediff)) * oneyearINsec) / myDepositInToken1) * 100
        // }
      }
    }

    // Calculate APR
    async calcAPR(strategyContract: any, pair: Pairs): Promise<void> {
      if (pair.assets !== null) {
        // calc APR
        let oraclePriceTwap
        const poolAddress = await strategyContract.methods.pool().call()
        console.log('poolAddress=', poolAddress)
        // const twapDuration = 2
        if (pair.currentTick !== null && pair.currentTick !== undefined) {
          console.log('twap=', pair.currentTick)
          oraclePriceTwap = await strategyContract.methods.getQuoteAtTick(Number(pair.currentTick), BigInt(Math.pow(10, 18)), pair.token0.tokenAddress, pair.token1.tokenAddress).call()
          console.log(';oraclePriceTwap=', oraclePriceTwap)
        }
        const Total0 =
            Number(pair.token0.balanceInVault) +
            Number((pair.uniliqs as any).amount0) +
            Number(pair.token0.underlying)
        const Total1 =
            Number(pair.token1.balanceInVault) +
            Number((pair.uniliqs as any).amount1) +
            Number(pair.token1.underlying)
        let mytvl0, mytvl1
        if (Number(pair.totalShares) === 0) {
          mytvl0 = 0
          mytvl1 = 0
        } else {
          mytvl0 = (Total0 * Number(pair.myShare)) / Number(pair.totalShares)
          mytvl1 = (Total1 * Number(pair.myShare)) / Number(pair.totalShares)
        }
        const oneyearINsec = 365 * 24 * 60 * 60
        const block = await this.web3.eth.getBlock((pair.assets as any).block)
        console.log('block timestamp=', block.timestamp)
        const timestamp = block.timestamp
        const headerNumber = await this.web3.eth.getBlockNumber()
        const headerBlock = await this.web3.eth.getBlock(headerNumber)
        // console.log('header=', JSON.stringify(headerBlock))
        const htimestamp = headerBlock.timestamp
        console.log('htimestamp=', htimestamp, ';timestamp=', timestamp)
        const timediff = Number(htimestamp) - Number(timestamp)
        const deposit0 = (pair.assets as any).deposit0
        const deposit1 = (pair.assets as any).deposit1
        const fd0 = Number(deposit0)
        let fd1, fm1
        if (oraclePriceTwap === 0) {
          fd1 = 0
          fm1 = 0
        } else {
          fd1 =
              (Number(deposit1) * Math.pow(10, Number(pair.token0.decimals))) /
              oraclePriceTwap
          fm1 =
              (Number(mytvl1) * Math.pow(10, Number(pair.token0.decimals))) /
              oraclePriceTwap
        }
        const fm0 = Number(mytvl0)
        const fdd = fd0 + fd1
        const fmm = fm0 + fm1
        console.log('fm0=', fm0, 'fm1=', fm1)
        console.log('decimal=', pair.token0.decimals)
        console.log('fmm=', fmm, 'fdd=', fdd, 'timediff=', timediff, 'oneyearInsec=', oneyearINsec)
        if (fdd === 0 || timediff === 0) {
          pair.currentAPR = 0
        } else {
          pair.currentAPR = (((fmm - fdd) / Number(timediff)) * oneyearINsec) / fdd
        }
      }
    }
}

export default PairsData

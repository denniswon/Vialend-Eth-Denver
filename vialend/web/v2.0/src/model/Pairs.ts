/* eslint-disable no-mixed-spaces-and-tabs */
/* eslint-disable no-tabs */
import Token from './Token'

export default class Pairs {
    public index:number;
    public id:number;
    public token0:Token;
    public token1:Token;
    public feeTier:string;
    public currentAPR:number;
    public capacity:string;
    public tvl:number;
    public maxTVL:number;
    public ofCapUsed:number;
    public vaultRange:number;
    public vaultLending:number;
    public tvlTotal0:number;
    public tvlTotal1:number;
    public tvlTotal0USD:number;
    public tvlTotal1USD:number;
    public lendingRatio:number;
	public uniswapRatio:number;
	public uniToken0Rate:number;
	public uniToken1Rate:number;
	public lendingToken0Rate:number;
	public lendingToken1Rate:number;
	public tickLower:number;
	public tickUpper:number;
    public currentTick:number;
    public currentPrice:number;
    public vaultAddress:string;
		public strategyAddress:string;

	public token0Address:string;
	public token1Address:string;
	public poolAddress:string;
    public keeperContract:object;
	public token0Contract:object;
	public token1Contract:object;
	public token0LendingContract:object;
	public token1LendingContract:object;
    public myShare:number;
	public totalShares:number;
	public myValueToken0Locked:number;
	public myValueToken1Locked:number;
	public myValueToken0USDLocked:number;
	public myValueToken1USDLocked:number;
	public flipped:boolean;
    public gettingData:boolean;
    public loadDataCompleted:boolean;
    public uniliqs:object;
    public assets:object;
    public netAPY:number;
    public currentDeposits:number;
    // Pair Setting Data
    public maxTotalSupply:string;
    public governanceAddress:string;
    public teamAddress:string;
    public uniPortionRatio:string;
    public protocolFee:string;

    constructor() {
	  this.index = 0
	  this.id = 0
	  this.token0 = new Token()
	  this.token1 = new Token()
	  this.feeTier = ''
	  this.currentAPR = 0
	  this.capacity = ''
	  this.tvl = 0
	  this.maxTVL = 0
	  this.ofCapUsed = 0
	  this.vaultRange = 0
	  this.vaultLending = 0
	  this.tvlTotal0 = 0
	  this.tvlTotal1 = 0
	  this.tvlTotal0USD = 0
	  this.tvlTotal1USD = 0
	  this.lendingRatio = 0
	  this.uniswapRatio = 0
	  this.uniToken0Rate = 0
	  this.uniToken1Rate = 0
	  this.lendingToken0Rate = 0
	  this.lendingToken1Rate = 0
	  this.tickLower = 0
	  this.tickUpper = 0
      this.currentTick = 0
      this.currentPrice = 0
	  this.vaultAddress = ''
      this.strategyAddress = ''
	  this.token0Address = ''
	  this.token1Address = ''
	  this.poolAddress = ''
	  this.keeperContract = {}
	  this.token0Contract = {}
	  this.token1Contract = {}
	  this.token0LendingContract = {}
	  this.token1LendingContract = {}
	  this.myShare = 0
	  this.totalShares = 0
	  this.myValueToken0Locked = 0
	  this.myValueToken1Locked = 0
	  this.myValueToken0USDLocked = 0
	  this.myValueToken1USDLocked = 0
	  this.flipped = false
      this.gettingData = false
      this.loadDataCompleted = false
      this.uniliqs = {}
      this.assets = {}
      this.netAPY = 0
      this.currentDeposits = 0
      this.maxTotalSupply = ''
      this.governanceAddress = ''
      this.teamAddress = ''
      this.uniPortionRatio = ''
      this.protocolFee = ''
    }

  // constructor(index:number, id:number, token0:Token, token1:Token, feeTier:string, currentAPR:number, capacity:string, tvl:number, maxTVL:number, ofCapUsed:number, vaultRange:number, token0Approved:boolean, token1Approved:boolean,
  //     vaultLending:number, tvlTotal0:number, tvlTotal1:number, tvlTotal0USD:number, tvlTotal1USD:number, lendingRatio:number, uniswapRatio:number, uniToken0Rate:number, uniToken1Rate:number, lendingToken0Rate:number, lendingToken1Rate:number,
  //     cLow:number, cHigh:number, token0BalanceInWallet:number, token1BalanceInWallet:number, token0BalanceInVault:number, token1BalanceInVault:number, vaultAddress:string, token0Address:string, token1Address:string,
  //     poolAddress:string, keeperContract:object, token0Contract:object, token1Contract:object, token0LendingContract:object, token1LendingContract:object,
  //     myShare:number, totalShares:number, myValueToken0Locked:number, myValueToken1Locked:number, myValueToken0USDLocked:number, myValueToken1USDLocked:number, flipped:boolean)

  // constructor(index?:any, id?:any, token0?:any, token1?:any, feeTier?:any, currentAPR?:any, capacity?:any, tvl?:any, maxTVL?:any, ofCapUsed?:any, vaultRange?:any, token0Approved?:any, token1Approved?:any,
  //   vaultLending?:any, tvlTotal0?:any, tvlTotal1?:any, tvlTotal0USD?:any, tvlTotal1USD?:any, lendingRatio?:any, uniswapRatio?:any, uniToken0Rate?:any, uniToken1Rate?:any, lendingToken0Rate?:any, lendingToken1Rate?:any,
  //   cLow?:any, cHigh?:any, token0BalanceInWallet?:any, token1BalanceInWallet?:any, token0BalanceInVault?:any, token1BalanceInVault?:any, vaultAddress?:any, token0Address?:any, token1Address?:any,
  //   poolAddress?:any, keeperContract?:any, token0Contract?:any, token1Contract?:any, token0LendingContract?:any, token1LendingContract?:any,
  //   myShare?:any, totalShares?:any, myValueToken0Locked?:any, myValueToken1Locked?:any, myValueToken0USDLocked?:any, myValueToken1USDLocked?:any, flipped?:any)
}

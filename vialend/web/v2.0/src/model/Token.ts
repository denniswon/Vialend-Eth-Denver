
export default class Token {
    public name:string;
    public iconLink:string;
    public abi:string;
    public tokenAddress:string;
    public decimals:number;
    public symbol:string;
    public tokenLendingAddress:string;
    public underlying:number;
    public myTVL:number;
    // balance of token
    public balanceInWallet:number;
    public balanceInVault:number;
    public balanceInPool:number;
    public balanceInLending:number;
    public tokenApproved:boolean;
    public approveLoading:boolean;

    // constructor(name:string, iconLink:string, abi:string, tokenAddress:string, decimals:number, symbol:string)
    // constructor(name?:any, iconLink?:any, abi?:any, tokenAddress?:any, decimals?:any, symbol?:any)
    constructor() {
      this.name = ''
      this.iconLink = ''
      this.abi = ''
      this.tokenAddress = ''
      this.decimals = 0
      this.symbol = ''
      this.tokenLendingAddress = ''
      this.underlying = 0
      this.myTVL = 0
      this.balanceInWallet = 0
      this.balanceInVault = 0
      this.balanceInPool = 0
      this.balanceInLending = 0
      this.tokenApproved = false
      this.approveLoading = false
    }
}

"use strict";
/* eslint-disable no-mixed-spaces-and-tabs */
/* eslint-disable no-tabs */
///<reference path = "Token.ts" />
var VialendModule;
(function (VialendModule) {
    class Pairs {
        constructor() {
            this.index = 0;
            this.id = 0;
            this.token0 = new VialendModule.Token();
            this.token1 = new VialendModule.Token();
            this.feeTier = '';
            this.currentAPR = 0;
            this.capacity = '';
            this.tvl = 0;
            this.maxTVL = 0;
            this.ofCapUsed = 0;
            this.vaultRange = 0;
            this.token0Approved = false;
            this.token1Approved = false;
            this.vaultLending = 0;
            this.tvlTotal0 = 0;
            this.tvlTotal1 = 0;
            this.tvlTotal0USD = 0;
            this.tvlTotal1USD = 0;
            this.lendingRatio = 0;
            this.uniswapRatio = 0;
            this.uniToken0Rate = 0;
            this.uniToken1Rate = 0;
            this.lendingToken0Rate = 0;
            this.lendingToken1Rate = 0;
            this.tickLower = 0;
            this.tickUpper = 0;
            this.currentTick = 0;
            this.currentPrice = 0;
            this.vaultAddress = '';
            this.token0Address = '';
            this.token1Address = '';
            this.poolAddress = '';
            this.keeperContract = {};
            this.token0Contract = {};
            this.token1Contract = {};
            this.token0LendingContract = {};
            this.token1LendingContract = {};
            this.myShare = 0;
            this.totalShares = 0;
            this.myValueToken0Locked = 0;
            this.myValueToken1Locked = 0;
            this.myValueToken0USDLocked = 0;
            this.myValueToken1USDLocked = 0;
            this.flipped = false;
            this.gettingData = false;
            this.loadDataCompleted = false;
            this.uniliqs = {};
            this.assets = {};
            this.netAPY = 0;
            this.currentDeposits = 0;
            this.maxTotalSupply = '';
            this.governanceAddress = '';
            this.teamAddress = '';
            this.uniPortionRatio = '';
            this.protocolFee = '';
        }
    }
    VialendModule.Pairs = Pairs;
})(VialendModule || (VialendModule = {}));

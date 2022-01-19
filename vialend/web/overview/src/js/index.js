"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const web3_min_js_1 = __importDefault(require("./components/web3/web3.min.js"));
// import Token from './model/Token'
// import Pairs from './model/Pairs'
const VaultBridge_V1_json_1 = __importDefault(require("./abi/VaultBridge_V1.json"));
///<reference path = "./model/Token.ts" />
///<reference path = "./model/Pairs.ts" />
var VialendModule;
(function (VialendModule) {
    let web3;
    function getWeb3Instance() {
        let web3Provider;
        if (window.ethereum) {
            web3Provider = window.ethereum;
            // try {
            //   (window as any).ethereum.enable()
            // } catch (error) {
            //   console.log(error)
            // }
        }
        else if (window.web3) {
            web3Provider = window.web3.currentProvider;
        }
        else {
            web3Provider = new web3_min_js_1.default.providers.HttpProvider('https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6');
            console.log('Non-Ethereum browser detected. You should consider trying MetaMask!');
        }
        web3 = new web3_min_js_1.default(web3Provider);
        return web3;
    }
    function contractInstance(abi, address) {
        if (web3 === undefined || web3 === null) {
            web3 = getWeb3Instance();
            console.log('create web3 instance');
        }
        return new web3.eth.Contract(abi, address);
    }
    const bridgeAddress = "0x033F3C5eAd18496BA462783fe9396CFE751a2342";
    //main code
    const bridgeContract = contractInstance(VaultBridge_V1_json_1.default, bridgeAddress);
    let iNum = 0;
    let vaultAddressInContract;
    while (true) {
        try {
            vaultAddressInContract = bridgeContract.methods.getAddress(iNum).call();
        }
        catch (err) {
            console.log('getting vault Address in contract occurs error:', err);
        }
        console.log('vaultAddressInContract', iNum, '=', vaultAddressInContract);
        if (vaultAddressInContract === null || vaultAddressInContract === undefined || Number(vaultAddressInContract) === 0) {
            console.log('vaultAddressInContract', iNum, ' is null,so break;');
            break;
        }
        const pair = new Pairs();
        pair.index = iNum;
        pair.id = iNum + 1;
        pair.vaultAddress = vaultAddressInContract;
        console.log('vault address(', iNum, ')=', vaultAddressInContract);
        // pair = await this.getPairsBaseInfo(pair)
        //this.pairsList.add(pair)
        iNum++;
    }
})(VialendModule || (VialendModule = {}));

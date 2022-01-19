"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.contractInstance = exports.getWeb3Instance = void 0;
const web3_1 = __importDefault(require("web3"));
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
        web3Provider = new web3_1.default.providers.HttpProvider('https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6');
        console.log('Non-Ethereum browser detected. You should consider trying MetaMask!');
    }
    web3 = new web3_1.default(web3Provider);
    return web3;
}
exports.getWeb3Instance = getWeb3Instance;
function contractInstance(abi, address) {
    if (web3 === undefined || web3 === null) {
        web3 = getWeb3Instance();
        console.log('create web3 instance');
    }
    return new web3.eth.Contract(abi, address);
}
exports.contractInstance = contractInstance;

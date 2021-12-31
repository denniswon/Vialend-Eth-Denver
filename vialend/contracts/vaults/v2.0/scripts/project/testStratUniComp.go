package include

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
	"time"

	ViaVault "viaroot/deploy/ViaVault"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

/*
VaultFactory publics

	      mapping (address => uint ) public vaultStatus;  // 1: active, 2: pending, 3: abandoned 0: empty
	      address[] public  vaultRegistry;

	   	function getAdmin() public view returns(address)
	      function getTeam() public view returns(address)
	      function setTeam(address _team) public onlyAdmin
	    	function changeVaultStat(address _vault, uint status) external onlyAdmin
	      function registerVault(address _vault) external
*/

func ViewVaults() {

	facotryInstance, _, _ := GetInstance3()

	myPrintln("View Vault Restry")

	myPrintln("VaultFactory Address:", Network.VaultFactory)

	count, _ := facotryInstance.GetCount(&bind.CallOpts{})
	i := count.Int64()
	for i > 0 {
		i--
		vaults, _ := facotryInstance.Vaults(&bind.CallOpts{}, big.NewInt(i))

		myPrintln(i+1, " of ", count)
		myPrintln("strategy:", vaults.Strategy)
		myPrintln("vault:", vaults.Vault)

		stat, _ := facotryInstance.GetStat(&bind.CallOpts{}, vaults.Strategy, vaults.Vault)
		myPrintln("stat:", stat)
	}

}

func GetStat(s string, v string) {
	facotryInstance, _, _ := GetInstance3()

	myPrintln(".... Check Vault Status.....")

	stat, _ := facotryInstance.GetStat(&bind.CallOpts{}, common.HexToAddress(s), common.HexToAddress(v))
	myPrintln("stat:", s)
	myPrintln("vault:", v)
	myPrintln("stat:", stat)

}
func CheckActive(addr string) {
	facotryInstance, _, _ := GetInstance3()

	myPrintln(".... Check Vault Status.....")

	result, _ := facotryInstance.CheckActive(&bind.CallOpts{}, common.HexToAddress(addr))
	myPrintln("result:", result)

}

func ChangeStat(_strategy string, _vault string, _stat int) {

	factoryInst, _, _ := GetInstance3()

	myPrintln(">>>>>>>> Change Stat <<<<<<<<<<<<< ")

	if tx, err := factoryInst.ChangeStat(Auth,
		common.HexToAddress(_strategy),
		common.HexToAddress(_vault),
		big.NewInt(int64(_stat))); err == nil {

		TxConfirm(tx.Hash())
	} else {
		log.Fatal("changestat", err)
	}

	GetStat(_strategy, _vault)

}

func Repair(_strategy string, _vault string) {

	factoryInst, _, _ := GetInstance3()

	myPrintln(">>>>>>>> Repair <<<<<<<<<<<<< ")

	if tx, err := factoryInst.Repair(Auth,
		common.HexToAddress(_strategy),
		common.HexToAddress(_vault)); err == nil {

		TxConfirm(tx.Hash())
	} else {
		log.Fatal("repair", err)
	}

	GetStat(_strategy, _vault)

}

func ViaVaultPublicList() {

	_, _, viaVaultInstance := GetInstance3()

	myPrintln("Vault Address:", Network.Vault)

	p1, _ := viaVaultInstance.Strategy(&bind.CallOpts{})
	p2, _ := viaVaultInstance.Admin(&bind.CallOpts{})
	p3, _ := viaVaultInstance.Token0(&bind.CallOpts{})
	p4, _ := viaVaultInstance.Token1(&bind.CallOpts{})
	p5, _ := viaVaultInstance.VaultCap(&bind.CallOpts{})
	p6, _ := viaVaultInstance.IndividualCap(&bind.CallOpts{})

	myPrintln()
	myPrintln("\n", ">>>>>>via Vault  public attributes <<<<<<<<<<<", "\n")
	myPrintln("Strategy", p1)
	myPrintln("Admin", p2)
	myPrintln("Token0", p3)
	myPrintln("Token1", p4)
	myPrintln("VaultCap", p5)
	myPrintln("IndividualCap", p6)

	myPrintln("\n", ">>>>>>via Vault  public functions <<<<<<<<<<<", "\n")
	myPrintln("1 function deposit( uint256 amountIn0, uint256 amountIn1) external")
	myPrintln("2 function withdraw( uint8 percent ) external nonReentrant returns (uint256 amount0, uint256 amount1)")
	myPrintln("3 function setVaultCap(uint256 newMax) external onlyAdmin")
	myPrintln("4 function setIndividualCap(uint256 newMax) external onlyAdmin")
	myPrintln("5 function checkCap(uint256 amount0,uint256 amount1) public view returns(uint256)")
	myPrintln("6 function sweep( address _token, uint256 amount) external  onlyAdmin")

}

func ViaStratUniCompPublicList() {
	_, viaStratInstance, _ := GetInstance3()

	myPrintln()
	myPrintln("--Strat Uni Comp address:", Network.VaultStrat)

	p1, _ := viaStratInstance.Admin(&bind.CallOpts{})
	p2, _ := viaStratInstance.Protocol(&bind.CallOpts{})
	p3, _ := viaStratInstance.Pool(&bind.CallOpts{})
	p4, _ := viaStratInstance.Token0(&bind.CallOpts{})
	p5, _ := viaStratInstance.Token1(&bind.CallOpts{})
	p7, _ := viaStratInstance.TickSpacing(&bind.CallOpts{})
	p8, _ := viaStratInstance.CLow(&bind.CallOpts{})
	p9, _ := viaStratInstance.CHigh(&bind.CallOpts{})
	p10, _ := viaStratInstance.UniPortion(&bind.CallOpts{})
	p11, _ := viaStratInstance.CompPortion(&bind.CallOpts{})
	p12, _ := viaStratInstance.CurComp0(&bind.CallOpts{})
	p13, _ := viaStratInstance.CurComp1(&bind.CallOpts{})
	p15, _ := viaStratInstance.ProtocolFee(&bind.CallOpts{})
	p6, _ := viaStratInstance.CTOKEN(&bind.CallOpts{}, common.HexToAddress(Network.TokenA))
	p17, _ := viaStratInstance.CTOKEN(&bind.CallOpts{}, common.HexToAddress(Network.TokenB))

	myPrintln("\n", ">>>>>>via Strat  public attributes <<<<<<<<<<<", "\n")
	myPrintln("Admin ", p1)
	myPrintln("Team", p2)
	myPrintln("Pool", p3)
	myPrintln("Token0", p4)
	myPrintln("Token1", p5)
	myPrintln("Ctoken0", p6)
	myPrintln("Ctoken1", p17)
	myPrintln("TickSpacing", p7)
	myPrintln("CLow", p8)
	myPrintln("CHigh", p9)
	myPrintln("UniPortion", p10)
	myPrintln("CompPortion", p11)
	myPrintln("CurComp0", p12)
	myPrintln("IndividualCap", p13)
	myPrintln("ProtocolFee", p15)
	//myPrintln("TriggerFee", p17)

	myPrintln("\n", ">>>>>>via Strat  public functions <<<<<<<<<<<", "\n")

	myPrintln("1 function emergency(int24 tickLower,int24 tickUpper,uint128 liquidity,uint256 amount0,uint256 amount1,bool redeemType) external onlyAdmin")
	myPrintln("2 function rebalance(int24 newLow, int24 newHigh) external nonReentrant")
	myPrintln("3 function getPrice() public view returns(uint256)")
	myPrintln("4 function getTwap() public view returns (int24 tick)")
	myPrintln("5 function getQuoteAtTick( int24 tick, uint128 baseAmount,  address baseToken, address quoteToken ) public pure returns (uint256 quote)")
	myPrintln("6 function vaultWithdraw(address to, uint256 shares, uint256 totalSupply) external returns(uint256 amount0, uint256 amount1)")
	myPrintln("7 function setMaxTwapDeviation(int24 _maxTwapDeviation) external onlyCreator")
	myPrintln("8 function setTwapDuration(uint32 _twapDuration) external onlyCreator")
	myPrintln("9 function setUniPortionRatio(uint8 ratio) external onlyCreator")
	myPrintln("10 function setCompPortionRatio(uint8 ratio) external onlyCreator")
	myPrintln("11 function setTeam(address _team) external onlyAdmin")
	myPrintln("12 function alloc() public returns ( bool )")
	myPrintln("13 function mintUniV3( int24 newLow, int24 newHigh, uint256 amount0, uint256 amount1 ) public")
	myPrintln("14 function mintCompound(address _underlying, uint256 amount) public")
	myPrintln("15 function calcShares( uint256 totalSupply, uint256 amountIn0, uint256 amountIn1) public returns(uint256 shares, uint256 amount0, uint256 amount1)")
	myPrintln("16 function getUniAmounts(int24 tickLower, int24 tickUpper)  public view returns (uint256 amount0, uint256 amount1)")

}

func ViaVaultPublicFunc(funcId int, params ...interface{}) {

	if funcId == 1 { // deposit
		doDeposit(params[0], params[1])
	}

	if funcId == 2 { // withdraw
		doWithdraw(params[0])
	}

}

func doDeposit(params ...interface{}) {

	amount0, Ok := params[0].(int64)
	amount1, Ok := params[1].(int64)
	if !Ok {
		log.Println("invalid parameter", params)
	}

	_, _, viaVaultInstance := GetInstance3()

	var maxToken0 = PowX(99999, int(Token[0].Decimals)) //new(big.Int).SetString("900000000000000000000000000000", 10)
	var maxToken1 = PowX(99999, int(Token[1].Decimals)) //new(big.Int).SetString("900000000000000000000000000000", 10)

	ApproveToken(common.HexToAddress(Network.TokenA), maxToken0, Network.Vault)
	Sleep(100)
	ApproveToken(common.HexToAddress(Network.TokenB), maxToken1, Network.Vault)
	Sleep(100)

	NonceGen()
	tx, err := viaVaultInstance.Deposit(Auth, big.NewInt(amount0), big.NewInt(amount1))
	if err != nil {
		log.Println("deposit err: ", err)
	}
	TxConfirm(tx.Hash())

}

func doWithdraw(params ...interface{}) {

	percent, Ok := params[0].(uint8)
	if !Ok {
		log.Println("invalid parameter", params)
	}

	_, _, viaVaultInstance := GetInstance3()

	NonceGen()
	tx, err := viaVaultInstance.Withdraw(Auth, uint8(percent))
	if err != nil {
		log.Println("deposit err: ", err)
	}
	TxConfirm(tx.Hash())

}

// via stratUniComp function

func ViaStratUniCompPublicFunc(funcId int, params ...interface{}) {

	if funcId == 1 { // emergency
		//doDeposit(params[0], params[1])
		myPrintln("todo")
	}

	if funcId == 2 { // rebalance
		myPrintln("todo")
	}
	if funcId == 3 {
		GetPriceStratCall()
	}

}

// func GetViaVaultAddress() {
// 	_, viaStratInstance, _ := GetInstance3()
// 	addr, _ := viaStratInstance.Vault(&bind.CallOpts{})
// 	myPrintln(addr)
// }

func GetPriceStratCall() {

	_, viaStratInstance, _ := GetInstance3()

	NonceGen()
	price, err := viaStratInstance.GetPrice(&bind.CallOpts{})
	if err != nil {
		log.Println("getprice err: ", err)
	}
	myPrintln("fun getPrice:", price)

}

func CalcPositionShares(a0 int, a1 int) {
	_, _, vaultInstance := GetInstance3()

	NonceGen()
	tx, err := vaultInstance.CalcPositionShares(Auth, big.NewInt(int64(a0)), big.NewInt(int64(a1)))
	if err != nil {
		log.Println("calcPositionshares err: ", err)
	}
	TxConfirm(tx.Hash())

}

func CheckCap(a0 int, a1 int) {
	_, _, vaultInstance := GetInstance3()
	p, err := vaultInstance.CheckCap(&bind.CallOpts{}, big.NewInt(int64(a0)), big.NewInt(int64(a1)))
	if err != nil {
		log.Println("checkCap err: ", err)
	}
	myPrintln("cap:", p)

}

func MoveFunds() {

	_, _, vaultInstance := GetInstance3()

	myPrintln(">>>>>>>> MoveFunds <<<<<<<<<<<<< ")

	if tx, err := vaultInstance.MoveFunds(Auth); err == nil {
		TxConfirm(tx.Hash())
	} else {
		log.Fatal("movefunds", err)
	}

}

func ViaCommand(cmd string, params []string) {

	if cmd == "lookup" {
		fmt.Println(cmd, ", ", params)
	}

	CmdHandler()

	a1, a2, a3 := GetInstance3()

	_, _, _ = a1, a2, a3

}

func CmdHandler() {

}

func Event(contractAddr string, eventNameString string, fromBlock int, toBlock int) {

	var eventNames []string
	if eventNameString != "" {
		eventNames = strings.Split(eventNameString, ",")
	}
	if toBlock == 0 { //listen
		ViaVaultEventListen(contractAddr, eventNames)
	} else {
		ViaVaultEventFiltered(contractAddr, fromBlock, toBlock, eventNameString)
	}
}

func ViaVaultEventListen(contractAddr string, eventNames []string) {

	contractAddress := common.HexToAddress(contractAddr)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)

	sub, err := ClientWS.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Println("event subcribe:", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(ViaVault.ApiABI)))
	//contractAbi, err := ViaVault.ApiMetaData.GetAbi()
	if err != nil {
		log.Println(err)
	}

	myPrintln("listening event:", eventNames)

	for {
		select {
		case err := <-sub.Err():
			log.Println("sel1:", err)
		case vLog := <-logs:
			for _, eventName := range eventNames {

				event, err := contractAbi.Unpack(eventName, vLog.Data)
				if err != nil {
					log.Println("event:", err, "-> ", eventName)
					myPrintln(vLog)
				} else {
					myPrintln(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\n")
					myPrintln(event)
					myPrintln("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<\n")

					eventBlockInfo(eventName, int64(vLog.BlockNumber))

					for j, v := range event {
						fmt.Println("event[", j, "]=", v)
					}
					myPrintln()

					myPrintln("Topics: (", len(vLog.Topics), ")")
					parseTopic(vLog.Topics)

				}

			}

		}

	}

}

// fileter e.g.
func ViaVaultEventFiltered(contract string, _from int, _to int, eventName string) {

	contractAddress := common.HexToAddress(contract)

	fromBlock := big.NewInt(int64(_from))
	toBlock := big.NewInt(int64(_to))
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := ClientWS.FilterLogs(context.Background(), query)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(" Block from ", fromBlock, " to:", toBlock)

	//contractAbi, err := abi.JSON(strings.NewReader(string(ViaVault.ApiABI)))
	contractAbi, err := ViaVault.ApiMetaData.GetAbi()

	if err != nil {
		log.Println("contractAbi->", err)
	}

	for _, vLog := range logs {

		event, err := contractAbi.Unpack(eventName, vLog.Data)
		if err != nil {
			log.Println("event:", err, "-> ", eventName)
		}

		eventBlockInfo(eventName, int64(vLog.BlockNumber))

		for j, v := range event {
			fmt.Println("event[", j, "]=", v)
		}
		fmt.Println()

		myPrintln("Topics: (", len(vLog.Topics), ")")
		parseTopic(vLog.Topics)
		//		typeof(vLog.Data)
		//		typeof(event)

	}

}

func typeof(v interface{}) {
	fmt.Println(fmt.Sprintf("%T", v))
}

func eventBlockInfo(eventName string, blockNumber int64) {
	block, err := Client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if err != nil {
		log.Fatal("block ", err)
	}
	myPrintln("----- event: ", eventName, " | event time:", time.Unix(int64(block.Time()), 0).Format("2006.01.02 15:04:05"))

}

func parseTopic(topic []common.Hash) {

	for k := 0; k < len(topic); k++ {

		hint := hexToInt(topic[k].String())
		fmt.Println("indexd topic ", k, " ", topic[k])
		fmt.Println("hexToInt:", hint)
	}

}

func parseLog(_eventname string, vLogData []uint8) {

	//contractAbi, err := abi.JSON(strings.NewReader(string(ViaVault.ApiABI)))
	contractAbi, err := ViaVault.ApiMetaData.GetAbi()

	if err != nil {
		log.Println(err)
	}

	event, err := contractAbi.Unpack(_eventname, vLogData)
	if err != nil {
		log.Println("event:", err, "-> ", _eventname)
	}

	fmt.Println("event : ", _eventname, ", (", len(event), ")")
	fmt.Println(time.Now().Format("2006.01.02 15:04:05"))

	for j := 0; j < len(event); j++ {
		fmt.Println("event[", j, "]=", event[j])
	}

}

func hexToInt(hexaString string) int64 {

	// replace 0x or 0X with empty String
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)

	output, err := strconv.ParseInt(numberStr, 16, 64)
	if err != nil {
		log.Println("WARNING: hexToInt error. can be ignored")
	}

	//fmt.Printf("Output %d", output)
	return output

}

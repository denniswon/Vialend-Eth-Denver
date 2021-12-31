package include

import (
	"fmt"
	"log"
	"math/big"
)

func SetProtocolFee(rate *big.Int) {

	vaultInstance := GetVaultInstance()

	tx, err := vaultInstance.SetProtocolFee(Auth, rate)

	if err != nil {
		log.Fatal("setprotocolfee err ", err)
	}

	fmt.Println("setProtocolFee rate=", rate, " tx sent: ", tx.Hash().Hex())

	TxConfirm(tx.Hash())

}
func SetUniswapPortionRatio(ratio uint8) {

	vaultInstance := GetVaultInstance()

	tx, err := vaultInstance.SetUniPortionRatio(Auth, ratio)

	if err != nil {
		log.Fatal("vaultInstance. setuniportion err ", err)
	}

	fmt.Println("setuniportionp tx: ", tx.Hash().Hex())

	//Readstring("tx sent.....  wait for pending..next .. ")
	TxConfirm(tx.Hash())

}

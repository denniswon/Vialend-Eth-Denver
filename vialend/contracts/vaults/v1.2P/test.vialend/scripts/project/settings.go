package include

import (
	"fmt"
	"log"
	"math/big"

	"../config"
)

func SetProtocolFee(rate *big.Int) {

	vaultInstance := GetVaultInstance()

	tx, err := vaultInstance.SetProtocolFee(config.Auth, rate)

	if err != nil {
		log.Fatal("setprotocolfee err ", err)
	}

	fmt.Println("setProtocolFee rate=", rate, " tx sent: ", tx.Hash().Hex())

	config.TxConfirm(tx.Hash())

}
func SetUniswapPortionRatio(ratio uint8) {

	vaultInstance := GetVaultInstance()

	tx, err := vaultInstance.SetUniPortionRatio(config.Auth, ratio)

	if err != nil {
		log.Fatal("vaultInstance. setuniportion err ", err)
	}

	fmt.Println("setuniportionp tx: ", tx.Hash().Hex())

	//config.Readstring("tx sent.....  wait for pending..next .. ")
	config.TxConfirm(tx.Hash())

}

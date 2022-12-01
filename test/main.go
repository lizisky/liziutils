package main

import (
	"fmt"
	"math/big"

	"github.com/liziblockchain/liziutils/utils"
)

func main() {
	tryBigIntJson()
}

func generalTesting() {
	fn := utils.GetAppBaseDir()
	tmp := utils.ToJSON(fn)

	fmt.Println("info:", fn, tmp)

	fmt.Println("now:", utils.GetTimestamp13())

	email := "abc@gmail.com"
	fmt.Println("is valid email:", utils.IsValidEmailAddr(email))
}

func tryBigIntJson() {

	type token struct {
		Time    int64
		Balance *big.Int
	}

	btc := token{
		Time:    utils.GetTimestamp13(),
		Balance: big.NewInt(1234),
	}

	fmt.Println("try big int:", btc)
}

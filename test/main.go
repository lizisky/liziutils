package main

import (
	"fmt"

	"github.com/liziblockchain/liziutils/utils"
)

func main() {
	fn := utils.GetAppBaseDir()
	tmp := utils.ToJSON(fn)

	fmt.Println("info:", fn, tmp)

	fmt.Println("now:", utils.GetTimestamp13())

	email := "abc@gmail.com"
	fmt.Println("is valid email:", utils.IsValidEmailAddr(email))

}

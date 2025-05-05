package main

import (
	"fmt"

	"github.com/bidhan948/go-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("FIRST ELEM AFTER GENESIS")
	chain.AddBlock("SECOND ELEM AFTER GENESIS")
	chain.AddBlock("THIRD ELEM AFTER GENESIS")

	for _, block := range chain.Blocks {
		fmt.Printf("PREV. HASH IS : %x\n", block.PrevHash)
		fmt.Printf("CURR DATA IS  : %s\n", block.Data)
		fmt.Printf("HASH IS  : %x\n", block.Hash)
	}
}

package main

import (
	"app/core"
	"fmt"
	"time"
)

func main() {
	var timeStart = time.Now()

	// initialize the chain with the genesis block
	fmt.Print("Initializing the blockchain")
	var blockChain = core.InitBlockChain()
	fmt.Println(" --> time taken:", time.Since(timeStart))

	// add blocks to the chain
	fmt.Print("Adding the second block")
	blockChain.AddBlock("Send 1 ETH to Vitalik Buterin")
	fmt.Println(" --> time taken:", time.Since(timeStart))

	fmt.Print("Adding the third block")
	blockChain.AddBlock("Send 3 ETHs to Vitalik Buterin")
	fmt.Println(" --> time taken:", time.Since(timeStart))

	// print the number of blocks in the chain
	fmt.Println()
	fmt.Println("Number of Blocks in Blockchain:", len(blockChain.Blocks))
	fmt.Println("The", len(blockChain.Blocks), "blocks are:")
	fmt.Println()

	// print the blocks
	for i, block := range blockChain.Blocks {
		fmt.Println("----------------------------")
		fmt.Println("Block number:", i)
		fmt.Println("Data:", string(block.Data))
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("----------------------------")
		fmt.Println()
	}
}

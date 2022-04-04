package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

func (chain *BlockChain) AddBlock(data string) {
	var block = &Block{}

	block.Timestamp = time.Now().Unix()
	block.Data = []byte(data)

	// if it's the first block, set the previous block hash to nil
	if len(chain.Blocks) == 0 {
		block.PrevBlockHash = nil
	} else {
		var prevBlock = chain.Blocks[len(chain.Blocks)-1]
		var prevBlockHash = prevBlock.Hash
		block.PrevBlockHash = prevBlockHash
	}

	// First, create block's header by merging the bytes array of Timestamp, Data, and PrevBlockHash
	var blockHeader = bytes.Join([][]byte{[]byte(strconv.FormatInt(block.Timestamp, 10)), block.Data, block.PrevBlockHash}, []byte{})

	// Use the SHA-256 hash function to generate the hash of the block's header
	var blockHash = sha256.Sum256(blockHeader)
	block.Hash = blockHash[:]

	chain.Blocks = append(chain.Blocks, block)
}

// to initialize the chain with a genesis block
func InitBlockChain() *BlockChain {
	var blockChain = &BlockChain{} // create a new blockchain*
	blockChain.AddBlock("Genesis") // add the genesis block to the chain*
	return blockChain
}

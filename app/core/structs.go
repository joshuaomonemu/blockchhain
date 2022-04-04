package core

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// chain is just a slice of blocks
type BlockChain struct {
	Blocks []*Block
}

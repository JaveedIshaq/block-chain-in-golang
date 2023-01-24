package main

import (
	"fmt"
	"log"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timeStamp    int64
	transactions []string
}

func newBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timeStamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash

	return b

}

func init() {
	log.SetPrefix("Blockchain: ")
}

func (b *Block) Print() {
	fmt.Printf("nonce %d\n", b.timeStamp)
	fmt.Printf("previousHash %s\n", b.previousHash)
	fmt.Printf("timeStamp %d\n", b.timeStamp)
	fmt.Printf("transactions %s\n", b.transactions)
}

func main() {

	aBlock := newBlock(0, "Init Hash")

	aBlock.Print()

}

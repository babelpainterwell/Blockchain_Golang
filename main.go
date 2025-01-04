package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash 		[]byte // derive the hash based on the data and prev hash 
	Data 		[]byte 
	PrevHash	[]byte 
}

type BlockChain struct {
	blocks 		[]*Block 
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	
	newBlock := Block{
		Data: 			[]byte(data),
		PrevHash: 		prevHash,
	}
	newBlock.DeriveHash()

	return &newBlock
}

func (chain *BlockChain) AddBlock(data string) {

	// get the last hash
	lastBlock := chain.blocks[len(chain.blocks)-1]
	prevHash := lastBlock.Hash

	// create the new block 
	newBlock := CreateBlock(data, prevHash)

	// attach it to the chain 
	chain.blocks = append(chain.blocks, newBlock)
}


func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}



func main() {
	
	chain := InitBlockChain()

	chain.AddBlock("First Block After Genesis")
	chain.AddBlock("Second Block After Genesis")
	chain.AddBlock("Third Block After Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("")
	}
}


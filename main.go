package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	Blocks []*Block
}
type Block struct {
	Hash     []byte
	Data     []byte
	prevHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.prevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]

}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

//Main Program starts here
func main() {
	chain := InitBlockChain()
	chain.AddBlock("This is First Chain")
	chain.AddBlock("This is second Block")
	chain.AddBlock("This is Third Block")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.prevHash)
		fmt.Printf("Data in the Block: %s \n", block.Data)
		fmt.Printf("Hash: %x \n", block.Hash)
	}
}

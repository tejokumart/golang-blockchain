package main

import "fmt"

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

package main

import(
	"fmt"
	"strconv"

	"github.com/tensor-programming/golang-blockchain/blockchain"
)

func main(){
	chain := blockChain.InitBlockChain()

	
	chain.AddBlock("First block after genesis")
	chain.AddBlock("Second block after genesis")
	chain.AddBlock("third block after genesis")

	for _, block := range chain.Blocks{

		fmt.Printf("=========== \n",)
		fmt.Printf("Previus hash: %x\n", block.PrevHash)
		fmt.Printf("datain block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockChain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Printf("=========== \n",)
		fmt.Println()
	}
}
package blockChain

import (
	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./tmp/blocks"
)

type BlockChain struct{
	LastHash []byte
	Database *badger.DB
}

type BlockChainIterator  struct{
	CurrentHash []byte
	Database *badger.DB
}

func InitBlockChain() *BlockChain{
	var lastHash []byte 
	
	opts := badger.DefaultOptions
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	Handle(err)

	err := db.Update(func(txn *badger.Txn) error{
		if _, err := txn.Get([]byte("lh"); err == badger.ErrKeyNotFound){
			fmt.Println("No existing blocchain found")

			genesis := Genesis()
			fmt.Println("Genesis proved")
	
			err = txn.Set([]byte("lh"), genesis.Hash)
			lashHash = genesis.Hash
	
			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			lashHash, err = item.Value()
			return err
		}
	})

	Handle(err)

	blockChain := BlockChain{lashHash, db}
	return &blockChain


}

func (chain *BlockChain) AddBlock(data string){

	var lashHash []byte

	err := chain.Database.View(func(txn *badger.Txn) error{
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lastHash, err = item.Value()
		return err
	})

	Handle(err)

	newBlock := CreateBLock(data, lastHash)
	err = chain.Database.Update(func(txn *badger.Txn) error{
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), newBLock.Hash)
		chain.LastHash = newBlock.Hash
		return err
	})
	Handle(err)
}
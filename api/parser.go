package api

import (
	"encoding/json"
	"eth_tracker/eth"
	"eth_tracker/storage"
	"fmt"
	"strconv"
)

type parser interface {
	GetCurrentBlock() int
	Subscribe(address string) bool
	GetTransactions(address string) []eth.EthereumTransaction
}

type Parser struct{}

func (parser *Parser) GetCurrentBlock() uint64 {
	block, err := strconv.ParseUint(storage.Get("latest_block"), 0, 64)
	if err != nil {
		panic("Block ID conversion failed with " + err.Error())
	}
	return block
}

func (parser *Parser) Subscribe(address string) bool {
	return storage.StoreWatchedAddress(address)
}

func (parser *Parser) GetTransactions(address string) (ethTrans []eth.EthereumTransaction) {
	data := storage.Get(address)
	if data != "" {
		err := json.Unmarshal([]byte(data), &ethTrans)
		if err != nil {
			fmt.Println("Parser: Error decoding JSON:", err, data)
			return nil
		}
	}

	return
}

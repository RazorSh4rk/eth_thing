package storage

import (
	"encoding/json"
	"eth_tracker/eth"
)

// add another address to watch
func StoreWatchedAddress(address string) bool {
	addresses := Get("watched_addresses")

	var arr []string
	err := json.Unmarshal([]byte(addresses), &arr)
	if err != nil {
		return false
	}

	arr = append(arr, addresses)

	modified, err := json.Marshal(arr)
	if err != nil {
		return false
	}

	Set("watched_addresses", string(modified))
	return true
}

func StoreTransaction(addr string, trans eth.EthereumTransaction) {
	transactions := Get(addr)

	var arr []eth.EthereumTransaction
	err := json.Unmarshal([]byte(transactions), &arr)
	if err != nil {
		panic("Error while deserializing stored transactions")
	}

	arr = append(arr, trans)

	modified, err := json.Marshal(arr)
	if err != nil {
		panic("Error while serializing stored transactions")
	}

	Set(addr, string(modified))
}

func GetWatchedAddresses() (arr []string) {
	addresses := Get("watched_addresses")

	err := json.Unmarshal([]byte(addresses), &arr)
	if err != nil {
		panic("Error while deserializing stored addresses")
	}

	return
}

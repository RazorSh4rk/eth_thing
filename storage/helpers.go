package storage

import "encoding/json"

// add another address to watch
func StoreWatchedAddress(address string) {
	addresses := Get("watched_addresses")

	var arr []string
	err := json.Unmarshal([]byte(addresses), &arr)
	if err != nil {
		panic("Error while deserializing stored addresses")
	}

	arr = append(arr, addresses)

	modified, err := json.Marshal(arr)
	if err != nil {
		panic("Error while serializing stored addresses")
	}

	Set("watched_addresses", string(modified))
}

func StoreTransaction(addr string, trans string) {
	transactions := Get(addr)

	var arr []string
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

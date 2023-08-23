package main

import (
	"eth_tracker/api"
	"eth_tracker/storage"
	"time"
)

// time_updated: last time in hex
// latest_block: last block num in hex

func main() {
	storage.Set("time_updated", "0x0")

	// 2 random addresses
	storage.Set("watched_addresses", `["0x388C818CA8B9251b393131C08a736A67ccB19297", "0xaa1a3e85282623e5f4228070f7781c06219ca56ceb224b276b9a25d83d09c46a"]`)

	// wait for a second for storage to initialize
	time.Sleep(time.Second)

	go storage.QueryTimestamp()

	parser := api.Parser{}
	api.Run(&parser)

	for {
	}
}

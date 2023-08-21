package main

import (
	"eth_tracker/storage"
	"time"
)

// time_updated: last time in hex
// latest_block: last block num in hex

func main() {
	storage.Set("time_updated", "0x0")
	storage.Set("watched_addresses", "[]")

	// wait for a second for storage to initialize
	time.Sleep(time.Second)

	go storage.QueryTimestamp()

	for {
	}
}

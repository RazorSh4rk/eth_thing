package storage

import (
	eth "eth_tracker/ETH"
	"fmt"
	"strconv"
	"time"
)

// querying every 2 seconds, if i have time i want to get the difference
// between the 2 timestamps and sleep for like half of that
func QueryTimestamp() {
	for {
		latestBlock := eth.GetLatestBlock()
		latestTime := latestBlock.Result.Timestamp
		storedTime := Get("time_updated")
		latestBlockNum := latestBlock.Result.Number

		if parseTimestamp(latestTime) > parseTimestamp(storedTime) {
			Set("time_updated", latestTime)
			Set("latest_block", latestBlockNum)
			fmt.Println(latestTime)

			// check transactions if we have a new block
			queryTransactions()
		}

		time.Sleep(2 * time.Second)
	}
}

// iterate through transactions in blocks and store them for watched users
// where the key is their address
// TODO: make in/out separate to make it more useful
// i will not do this rn, but would definitely do it for prod
func queryTransactions() {
	all_transactions := eth.GetLatestBlock().Result.Transactions
	for _, hash := range all_transactions {
		transaction := eth.GetTransactionByHash(hash)
		watched := GetWatchedAddresses()

		for _, addr := range watched {
			if transaction.Result.From == addr || transaction.Result.To == addr {
				StoreTransaction(addr, transaction.Result.Hash)
			}
		}
	}
}

func parseTimestamp(timestamp string) uint64 {
	unix_time, err := strconv.ParseUint(timestamp, 0, 64)
	if err != nil {
		panic("Unix time conversion failed with " + err.Error())
	}
	return unix_time
}

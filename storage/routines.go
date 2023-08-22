package storage

import (
	"eth_tracker/eth"
	"fmt"
	"strconv"
	"time"
)

// querying every 2 seconds
func QueryTimestamp() {
	for {
		latestBlock := eth.GetLatestBlock()
		latestTime := latestBlock.Result.Timestamp
		storedTime := Get("time_updated")
		latestBlockNum := latestBlock.Result.Number

		if parseTimestamp(latestTime) > parseTimestamp(storedTime) {
			Set("time_updated", latestTime)
			Set("latest_block", latestBlockNum)
			fmt.Println("Updating timestamp to ", latestTime)

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
	fmt.Println("Queried transactions: ", len(all_transactions))

	watched := GetWatchedAddresses()
	fmt.Println("Watched addresses: ", len(watched))

	for _, hash := range all_transactions {
		transaction := eth.GetTransactionByHash(hash)

		for _, addr := range watched {
			if transaction.Result.From == addr || transaction.Result.To == addr {
				fmt.Println("New transaction for ", addr)
				StoreTransaction(addr, transaction)
			}
		}
	}
}

func parseTimestamp(timestamp string) uint64 {
	// sometimes the http request fails
	// to return so i just skip an iteration
	if timestamp == "" {
		timestamp = Get("time_updated")
	}
	unix_time, err := strconv.ParseUint(timestamp, 0, 64)
	if err != nil {
		panic("Unix time conversion failed with " + err.Error() + " timestamp: " + timestamp)
	}
	return unix_time
}

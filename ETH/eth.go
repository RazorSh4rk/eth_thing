// Thingy for interacting with the ETH json RPC

package eth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Only including the fields I will actually be
// using to not spam up the code
type EthereumBlock struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Hash         string   `json:"hash"`
		Number       string   `json:"number"`
		Timestamp    string   `json:"timestamp"`
		Transactions []string `json:"transactions"`
	} `json:"result"`
}

// same here
type EthereumTransaction struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		From string `json:"from"`
		To   string `json:"to"`
		Hash string `json:"hash"`
	} `json:"result"`
}

var url string = "https://cloudflare-eth.com/"

func GetLatestBlock() (ethBlock EthereumBlock) {
	payload := []byte(`{
		"jsonrpc": "2.0",
		"method": "eth_getBlockByNumber",
		"params": ["latest", false],
		"id": 1
	}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&ethBlock)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	return
}

func GetTransactionByHash(hash string) (ethTrans EthereumTransaction) {
	raw := `{
		"jsonrpc": "2.0",
		"method": "eth_getTransactionByHash",
		"params": [
			"%s"
		],
		"id": 1
	}`
	raw = fmt.Sprintf(raw, hash)
	payload := []byte(raw)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&ethTrans)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	return
}

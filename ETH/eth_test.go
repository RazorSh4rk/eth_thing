package eth

import (
	"testing"
)

func TestGetLatestBlock(t *testing.T) {
	result := GetLatestBlock()

	if result.Jsonrpc != "2.0" {
		t.Errorf("Expected Jsonrpc: 2.0, but got: %s", result.Jsonrpc)
	}

	if result.Result.Hash == "" {
		t.Error("Hash is empty")
	}
}

func TestGetTransactionByHash(t *testing.T) {
	hash := "0xfdff3ee9a24ababd19064f8eea23b335cdb37fe5225974a4bc7984338fe1a039"
	result := GetTransactionByHash(hash)

	if result.Jsonrpc != "2.0" {
		t.Errorf("Expected Jsonrpc: 2.0, but got: %s", result.Jsonrpc)
	}

	if result.Result.Hash != hash {
		t.Errorf("Expected hash: %s, but got: %s", hash, result.Result.Hash)
	}
}

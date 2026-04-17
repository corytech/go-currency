package currency

import (
	"slices"
	"testing"
)

var AllBlockchainNetwork = []string{
	"erc20",
	"trc20",
	"bep20",
	"bep2",
	"sol",
	"base",
	"optimism",
	"arbitrum",
	"stellar",
	"cardano",
	"eip155:10",
	"eip155:137",
	"eip155:8453",
	"eip155:42161",
}

var BlockChainTagMemoRequires = map[string]bool{
	"erc20":        false,
	"trc20":        false,
	"bep20":        false,
	"bep2":         false,
	"sol":          false,
	"base":         false,
	"optimism":     false,
	"arbitrum":     false,
	"stellar":      true,
	"cardano":      true,
	"eip155:10":    false,
	"eip155:137":   false,
	"eip155:8453":  false,
	"eip155:42161": false,
}

var BlockChainIsValid = map[string]bool{
	"erc20":        true,
	"trc20":        true,
	"bep20":        true,
	"bep2":         true,
	"sol":          true,
	"base":         true,
	"optimism":     true,
	"arbitrum":     true,
	"stellar":      true,
	"cardano":      true,
	"eip155:10":    true,
	"eip155:137":   true,
	"eip155:8453":  true,
	"eip155:42161": true,
	"trcNotValid":  false,
}

func TestGetAllBlockchainNetwork(t *testing.T) {
	blockchainNetworks := GetAllBlockchainNetwork()
	for _, blockchainNetwork := range AllBlockchainNetwork {
		if !slices.Contains(blockchainNetworks, BlockchainNetwork(blockchainNetwork)) {
			t.Fatalf("Expected '%v' in list '%v'", blockchainNetwork, blockchainNetworks)
		}
	}
}

func TestIsTagOrMemoRequired(t *testing.T) {
	for blockchainNetwork, isMemoRequired := range BlockChainTagMemoRequires {
		if BlockchainNetwork(blockchainNetwork).isTagOrMemoRequired() != isMemoRequired {
			t.Fatalf("Expected '%v' isTagOrMemoRequired '%v' get '%v'", blockchainNetwork, isMemoRequired, BlockchainNetwork(blockchainNetwork).isTagOrMemoRequired())
		}
	}
}

func TestBlockchainNetworkIsValid(t *testing.T) {
	for key, value := range BlockChainIsValid {
		if BlockchainNetwork(key).IsValid() != value {
			t.Fatalf("Expected '%v' isValid '%v' get '%v'", key, value, BlockchainNetwork(key).IsValid())
		}
	}
}

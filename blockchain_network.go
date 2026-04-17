package currency

import "slices"

type BlockchainNetwork string

const (
	Erc20        BlockchainNetwork = "erc20"
	Trc20        BlockchainNetwork = "trc20"
	Bep20        BlockchainNetwork = "bep20"
	Bep2         BlockchainNetwork = "bep2"
	Sol          BlockchainNetwork = "sol"
	Base         BlockchainNetwork = "base"
	Optimism     BlockchainNetwork = "optimism"
	Arbitrum     BlockchainNetwork = "arbitrum"
	Stellar      BlockchainNetwork = "stellar"
	Cardano      BlockchainNetwork = "cardano"
	Eip155_10    BlockchainNetwork = "eip155:10"
	Eip155_137   BlockchainNetwork = "eip155:137"
	Eip155_8453  BlockchainNetwork = "eip155:8453"
	Eip155_42161 BlockchainNetwork = "eip155:42161"
)

func (n BlockchainNetwork) isTagOrMemoRequired() bool {
	switch n {
	case Stellar,
		Cardano:
		return true
	}
	return false
}

func (n BlockchainNetwork) IsValid() bool {
	return slices.Contains(GetAllBlockchainNetwork(), n)
}

func GetAllBlockchainNetwork() []BlockchainNetwork {
	return []BlockchainNetwork{
		Erc20,
		Trc20,
		Bep20,
		Bep2,
		Sol,
		Base,
		Optimism,
		Arbitrum,
		Stellar,
		Cardano,
		Eip155_10,
		Eip155_137,
		Eip155_8453,
		Eip155_42161,
	}
}

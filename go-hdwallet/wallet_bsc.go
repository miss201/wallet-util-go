package hdwallet

import (
	"github.com/ethereum/go-ethereum/crypto"
)

func init() {
	coins[BSC] = newBSC
}

type bsc struct {
	name   string
	symbol string
	key    *Key

	contract string
}

func newBSC(key *Key) Wallet {
	return &bsc{
		name:   "Binance Smart Chain",
		symbol: "BSC",
		key:    key,
	}
}

func (c *bsc) GetType() uint32 {
	return c.key.opt.CoinType
}

func (c *bsc) GetName() string {
	return c.name
}

func (c *bsc) GetSymbol() string {
	return c.symbol
}

func (c *bsc) GetKey() *Key {
	return c.key
}

func (c *bsc) GetAddress() (string, error) {
	return crypto.PubkeyToAddress(*c.key.PublicECDSA).Hex(), nil
}

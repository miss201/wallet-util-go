package hdwallet

import (
	"github.com/ethereum/go-ethereum/crypto"
)

func init() {
	coins[MATIC] = newMATIC
}

type matic struct {
	name   string
	symbol string
	key    *Key

	contract string
}

func newMATIC(key *Key) Wallet {
	return &matic{
		name:   "Polygon",
		symbol: "MATIC",
		key:    key,
	}
}

func (c *matic) GetType() uint32 {
	return c.key.opt.CoinType
}

func (c *matic) GetName() string {
	return c.name
}

func (c *matic) GetSymbol() string {
	return c.symbol
}

func (c *matic) GetKey() *Key {
	return c.key
}

func (c *matic) GetAddress() (string, error) {
	return crypto.PubkeyToAddress(*c.key.PublicECDSA).Hex(), nil
}

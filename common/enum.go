package common

const (
	BTCSYM  = "BTC"
	ETHSYM  = "ETH"
	USDTSYM = "USDT"
)

const (
	ETHADDRESSLEN    = 40
	ETHPRIVATEKEYLEN = 64
)

// mnemonic language
const (
	English            = "english"
	ChineseSimplified  = "chinese_simplified"
	ChineseTraditional = "chinese_traditional"
	JAPANESE           = "Japanese"
	ITALIAN            = "Italian"
	KOREAN             = "Korean"
	SPANISH            = "Spanish"
	FRENCH             = "French"
)

// zero is deafult of uint32
const (
	Zero      uint32 = 0
	ZeroQuote uint32 = 0x80000000
	BTCToken  uint32 = 0x10000000
	ETHToken  uint32 = 0x20000000
)

// wallet type from bip44
const (
	// https://github.com/satoshilabs/slips/blob/master/slip-0044.md#registered-coin-types
	BTC         = ZeroQuote
	BTC_TESTNET = ZeroQuote + 1
	DASH        = ZeroQuote + 5
	ETH         = ZeroQuote + 60
	BSC         = ZeroQuote + 60

	// btc token
	USDT = BTCToken + 1

	// eth token
	IOST = ETHToken + 1
	USDC = ETHToken + 2
)

var coinTypes = map[uint32]uint32{
	USDT: BTC,
	IOST: ETH,
	USDC: ETH,
}

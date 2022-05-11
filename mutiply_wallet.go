//go:build ignore
// +build ignore

package main

//账户结构体
//type multiplyAccount struct {
//	ErrorCode    string
//	ErrorMessage string
//	Address      string //地址
//	PrivateKey   string //私钥
//	PublicKey    string //公钥
//	Mnemonic     string //助记词，12个字组成，字与字之间使用空格间隔
//}

//账户结构体
type multiplyAccountGo struct {
	ErrorCode    string
	ErrorMessage string
	Address      string //地址
	PrivateKey   string //私钥
	PublicKey    string //公钥
	Mnemonic     string //助记词，12个字组成，字与字之间使用空格间隔
}

/**
*根据coinType创建对应币种的账户
*@coinType: BTC ETH BNB MATIC SOL ...
 */
func CreateAccount(coinType string) multiplyAccountGo {
	account := multiplyAccountGo{}
	switch coinType {
	case "BTC":
		account = BTCW.createAccount()
	case "ETH":
		account = ETHW.createAccount()
	case "BNB":
		account = BNBW.createAccount()
	case "MATIC":
		account = MATICW.createAccount()
	case "SOL":
		account = SOLW.createAccount()
	case "LUNA":
		account = TERRA.createAccount()
	}
	return account
}

/**
通过助记词恢复账户
@coinType:BTC ETH BNB MATIC SOL ...
*/
func MnemonicToAccount(coinType string, mnemonic string) multiplyAccountGo {
	account := multiplyAccountGo{}
	switch coinType {
	case "BTC":
		account = BTCW.createAccountByMnemonic(mnemonic)
	case "ETH":
		account = ETHW.createAccountByMnemonic(mnemonic)
	case "BNB":
		account = BNBW.createAccountByMnemonic(mnemonic)
	case "MATIC":
		account = MATICW.createAccountByMnemonic(mnemonic)
	case "SOL":
		account = SOLW.createAccountByMnemonic(mnemonic)
	case "LUNA":
		account = TERRA.createAccountByMnemonic(mnemonic)
	}
	return account
}

/**
*通过私钥恢复账户
*@coinType:BTC ETH BNB MATIC SOL ...
 */
func GetAccountByPrivateKey(coinType string, privateKey string) multiplyAccountGo {
	account := multiplyAccountGo{}
	switch coinType {
	case "BTC":
		account = BTCW.createAccountByPrivateKey(privateKey)
	case "ETH":
		account = ETHW.createAccountByPrivateKey(privateKey)
	case "BNB":
		account = BNBW.createAccountByPrivateKey(privateKey)
	case "MATIC":
		account = MATICW.createAccountByPrivateKey(privateKey)
	case "SOL":
		account = SOLW.createAccountByPrivateKey(privateKey)
	case "LUNA":
		account = TERRA.createAccountByPrivateKey(privateKey)
	}
	return account
}

//go:build ignore
// +build ignore

package main

//账户结构体
type multiplyAccount struct {
	ErrorCode    string
	ErrorMessage string
	Address      string //地址
	PrivateKey   string //私钥
	PublicKey    string //公钥
	Mnemonic     string //助记词，12个字组成，字与字之间使用空格间隔
}

/**
*根据coinType创建对应币种的账户
*@coinType: BTC ETH BSC MATIC
 */
func CreateAccount(coinType string) *multiplyAccount {
	account := &multiplyAccount{}
	switch coinType {
	case "BTC":
		account = BTCW.createAccount()
	case "ETH":
		account = ETHW.createAccount()
	case "BSC":
		account = BSCW.createAccount()
	case "MATIC":
		account = MATICW.createAccount()
	}
	return account
}

/**
通过助记词恢复账户
@coinType:BTC ETH BSC MATIC ...
*/
func MnemonicToAccount(coinType string, mnemonic string) *multiplyAccount {
	account := &multiplyAccount{}
	switch coinType {
	case "BTC":
		account = BTCW.createAccountByMenmonic(mnemonic)
	case "ETH":
		account = ETHW.createAccountByMenmonic(mnemonic)
	case "BSC":
		account = BSCW.createAccountByMenmonic(mnemonic)
	case "MATIC":
		account = MATICW.createAccountByMenmonic(mnemonic)
	}
	return account
}

/**
*通过私钥恢复账户
*@coinType:BTC ETH BSC MATIC ...
 */
func GetAccountByPrivateKey(coinType string, privateKey string) *multiplyAccount {
	account := &multiplyAccount{}
	switch coinType {
	case "BTC":
		account = BTCW.createAccountByPrivateKey(privateKey)
	case "ETH":
		account = ETHW.createAccountByMenmonic(privateKey)
	case "BSC":
		account = BSCW.createAccountByMenmonic(privateKey)
	case "MATIC":
		account = MATICW.createAccountByMenmonic(privateKey)
	}
	return account
}
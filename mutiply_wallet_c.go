package main

/*
#include <stdlib.h>

typedef struct CBaseString_Flag {
	char *s;
	int len;
} CBaseString;

typedef struct MultiplyAccount_Flag {
	CBaseString ErrorCode;
	CBaseString ErrorMessage;
	CBaseString Address;
	CBaseString PrivateKey;
	CBaseString PublicKey;
	CBaseString Mnemonic;
} multiplyAccount;
*/
import "C"
import (
	"log"
	"unsafe"
)

//账户结构体
type multiplyAccountGo struct {
	ErrorCode    string
	ErrorMessage string
	Address      string //地址
	PrivateKey   string //私钥
	PublicKey    string //公钥
	Mnemonic     string //助记词，12个字组成，字与字之间使用空格间隔
}

type multiplyAccount C.multiplyAccount

//转换为CBaseString
func toCBaseString(str string) C.CBaseString {
	return C.CBaseString{
		s:   (*C.char)(C.CString(str)),
		len: C.int(len(str)),
	}
}

//转换为C的multiplyAccount
func go2CAccount(account multiplyAccountGo) multiplyAccount {
	return multiplyAccount{
		ErrorCode:    toCBaseString(account.ErrorCode),
		ErrorMessage: toCBaseString(account.ErrorMessage),
		Address:      toCBaseString(account.Address),
		PrivateKey:   toCBaseString(account.PrivateKey),
		PublicKey:    toCBaseString(account.PublicKey),
		Mnemonic:     toCBaseString(account.Mnemonic),
	}
}

//export CreateAccount
// 根据coinType创建对应币种的账户
// @coinType: BTC ETH BSC MATIC SQL LUNA...
func CreateAccount(coinType *C.char) multiplyAccount {
	aCoinType := C.GoString(coinType)
	account := multiplyAccountGo{}
	switch aCoinType {
	case "BTC":
		account = BTCW.createAccount()
	case "ETH":
		account = ETHW.createAccount()
	case "BSC": //币安
		account = BSCW.createAccount()
	case "MATIC":
		account = MATICW.createAccount()
	case "SQL":
		account = SQLW.createAccount()
	case "LUNA":
		account = TERRA.createAccount()
	}
	return go2CAccount(account)
}

//export MnemonicToAccount
// 通过助记词恢复账户
// @coinType:BTC ETH BSC MATIC SQL LUNA...
func MnemonicToAccount(coinType *C.char, mnemonic *C.char) multiplyAccount {
	aCoinType := C.GoString(coinType)
	aMnemonic := C.GoString(mnemonic)
	account := multiplyAccountGo{}
	switch aCoinType {
	case "BTC":
		account = BTCW.createAccountByMnemonic(aMnemonic)
	case "ETH":
		account = ETHW.createAccountByMnemonic(aMnemonic)
	case "BSC": //币安
		account = BSCW.createAccountByMnemonic(aMnemonic)
	case "MATIC":
		account = MATICW.createAccountByMnemonic(aMnemonic)
	case "SQL":
		account = SQLW.createAccountByMnemonic(aMnemonic)
	case "LUNA":
		account = TERRA.createAccountByMnemonic(aMnemonic)
	}
	return go2CAccount(account)
}

//export GetAccountByPrivateKey
// 通过私钥恢复账户
// @coinType:BTC ETH BSC MATIC  SQL LUNA...
func GetAccountByPrivateKey(coinType *C.char, privateKey *C.char) multiplyAccount {
	aCoinType := C.GoString(coinType)
	aPrivateKey := C.GoString(privateKey)
	account := multiplyAccountGo{}
	switch aCoinType {
	case "BTC":
		account = BTCW.createAccountByPrivateKey(aPrivateKey)
	case "ETH":
		account = ETHW.createAccountByPrivateKey(aPrivateKey)
	case "BSC": //币安
		account = BSCW.createAccountByPrivateKey(aPrivateKey)
	case "MATIC":
		account = MATICW.createAccountByPrivateKey(aPrivateKey)
	case "SQL":
		account = SQLW.createAccountByPrivateKey(aPrivateKey)
	case "LUNA":
		account = TERRA.createAccountByPrivateKey(aPrivateKey)
	}
	return go2CAccount(account)
}

//export FreeMultiplyAccount
func FreeMultiplyAccount(result multiplyAccount) {
	if result.ErrorCode.s != nil {
		log.Printf("释放ErrorCode")
		C.free(unsafe.Pointer(result.ErrorCode.s))
		log.Printf("释放ErrorCode后：%v", result.ErrorCode.s)
	}
	if result.ErrorMessage.s != nil {
		log.Printf("释放ErrorMessage\n")
		C.free(unsafe.Pointer(result.ErrorMessage.s))
		log.Printf("释放ErrorMessage后：%v\n", result.ErrorMessage.s)
	}
	if result.Address.s != nil {
		log.Printf("释放Address\n")
		C.free(unsafe.Pointer(result.Address.s))
		log.Printf("释放Address后：%v\n", result.Address.s)
	}
	if result.PrivateKey.s != nil {
		log.Printf("释放PrivateKey\n")
		C.free(unsafe.Pointer(result.PrivateKey.s))
		log.Printf("释放PrivateKey后：%v\n", result.PrivateKey.s)
	}
	if result.PublicKey.s != nil {
		log.Printf("释放PublicKey\n")
		C.free(unsafe.Pointer(result.PublicKey.s))
		log.Printf("释放PublicKey后：%v\n", result.PublicKey.s)
	}
	if result.Mnemonic.s != nil {
		log.Printf("释放Mnemonic\n")
		C.free(unsafe.Pointer(result.Mnemonic.s))
		log.Printf("释放Mnemonic后：%v\n", result.Mnemonic.s)
	}
}

func main() {

}

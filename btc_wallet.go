package main

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcutil"
	"github.com/miss201/wallet-util-go/common"
	"log"
)

var (
	BTCMainnet       = "BTCMainnet"
	BTCSegwitMainnet = "BTCSegwitMainnet"
	BTCTestnet       = "BTCTestnet"
	BTCSegwitTestnet = "BTCSegwitTestnet"

	BTCW              *BTCWallet
	BTCSegwitW        *BTCWallet
	BTCTestnetW       *BTCWallet
	BTCTestnetSegwitW *BTCWallet
)

func init() {
	BTCW = NewBTCWallet(BTCMainnet)
	BTCSegwitW = NewBTCWallet(BTCSegwitMainnet)
	BTCTestnetW = NewBTCWallet(BTCTestnet)
	BTCTestnetSegwitW = NewBTCWallet(BTCSegwitTestnet)
}

type BTCWallet struct {
	wallet      *common.Wallet
	mnemonicLen int
}

func NewBTCWallet(wc string) *BTCWallet {
	newWallet := BTCWallet{}
	switch wc {
	case BTCMainnet:
		newWallet = BTCWallet{wallet: common.NewWallet(common.BTC, false, false, &common.BTCParams), mnemonicLen: 12}
	case BTCSegwitMainnet:
		newWallet = BTCWallet{wallet: common.NewWallet(common.BTC, true, false, &common.BTCParams), mnemonicLen: 12}
	case BTCTestnet:
		newWallet = BTCWallet{wallet: common.NewWallet(common.BTC_TESTNET, false, false, &common.BTCTestnetParams), mnemonicLen: 12}
	case BTCSegwitTestnet:
		newWallet = BTCWallet{wallet: common.NewWallet(common.BTC_TESTNET, true, false, &common.BTCTestnetParams), mnemonicLen: 12}
	default:
		newWallet = BTCWallet{wallet: common.NewWallet(common.BTC, false, false, &common.BTCParams), mnemonicLen: 12}
	}
	return &newWallet
}

func (BTCw *BTCWallet) GenerateAddressFromMnemonic(mnemonic, language string) (string, error) {
	return BTCw.wallet.GenerateAddressFromMnemonic(mnemonic, language)
}

func (BTCw *BTCWallet) GenerateAddressFromPrivateKey(privateKey string) (string, error) {
	return BTCw.wallet.GenerateAddressFromPrivateKey(privateKey)
}

func (BTCw *BTCWallet) ExportPrivateKeyFromMnemonic(mnemonic, language string) (string, error) {
	return BTCw.wallet.ExportPrivateKeyFromMnemonic(mnemonic, language)
}

func (BTCw *BTCWallet) CheckAddress(address string) (bool, error) {
	return common.CheckAddress(address, BTCw.wallet.NetParam)
}

func (BTCw *BTCWallet) CheckPrivateKey(privateKey string) (bool, error) {
	return common.CheckPrivateKey(privateKey, BTCw.wallet.NetParam)
}

// get publickey from privatekey
func (BTCw *BTCWallet) GetPubKeyFromPrivateKey(privateKey string) (string, error) {

	isValid, err := BTCw.CheckPrivateKey(privateKey)
	if isValid == false || err != nil {
		return "", err
	}

	wifKey, err := btcutil.DecodeWIF(privateKey)
	if err != nil {
		return "", common.ERR_INVALID_PRIVATEKEY
	}
	pubHex := hex.EncodeToString(wifKey.SerializePubKey())
	return pubHex, nil
}

func (BTCw *BTCWallet) createAccount() multiplyAccountGo {
	mnemonic, err := BTCW.wallet.GenerateMnemonic(12)
	if err != nil {
		log.Printf("生成用户账户助记词出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "A0001",
			ErrorMessage: fmt.Sprintf("生成用户出错:%v", err),
		}
	}
	privateKey, err := BTCW.ExportPrivateKeyFromMnemonic(mnemonic, common.English)
	if err != nil {
		log.Printf("生成用户账户私钥出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "A0001",
			ErrorMessage: fmt.Sprintf("生成用户出错:%v", err),
		}
	}
	publicKey, err := BTCW.GetPubKeyFromPrivateKey(privateKey)
	if err != nil {
		log.Printf("生成用户账户公钥出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "A0001",
			ErrorMessage: fmt.Sprintf("生成用户出错:%v", err),
		}
	}
	address, err := BTCW.GenerateAddressFromPrivateKey(privateKey)
	if err != nil {
		log.Printf("生成用户账户地址出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "A0001",
			ErrorMessage: fmt.Sprintf("生成用户出错:%v", err),
		}
	}
	return multiplyAccountGo{
		Address:    address,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Mnemonic:   mnemonic,
	}
}

func (BTCw *BTCWallet) createAccountByMnemonic(mnemonic string) multiplyAccountGo {
	privateKey, err := BTCw.ExportPrivateKeyFromMnemonic(mnemonic, common.English)
	if err != nil {
		log.Printf("通过助记词获取账户私钥信息出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "M0001",
			ErrorMessage: fmt.Sprintf("通过助记词获取账户信息出错:%v", err),
		}
	}
	publicKey, err := BTCw.GetPubKeyFromPrivateKey(privateKey)
	if err != nil {
		log.Printf("通过助记词获取账户公钥信息出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "M0001",
			ErrorMessage: fmt.Sprintf("通过助记词获取账户信息出错:%v", err),
		}
	}
	address, err := BTCw.GenerateAddressFromPrivateKey(privateKey)
	if err != nil {
		log.Printf("通过助记词获取账户地址信息出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "M0001",
			ErrorMessage: fmt.Sprintf("通过助记词获取账户信息出错:%v", err),
		}
	}
	return multiplyAccountGo{
		Address:    address,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Mnemonic:   mnemonic,
	}
}

func (BTCw *BTCWallet) createAccountByPrivateKey(privateKey string) multiplyAccountGo {
	publicKey, err := BTCw.GetPubKeyFromPrivateKey(privateKey)
	if err != nil {
		log.Printf("通过私钥获取账号公钥信息出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "M0001",
			ErrorMessage: fmt.Sprintf("通过私钥获取账号信息出错:%v", err),
		}
	}
	address, err := BTCw.GenerateAddressFromPrivateKey(privateKey)
	if err != nil {
		log.Printf("通过私钥获取账号地址信息出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "M0001",
			ErrorMessage: fmt.Sprintf("通过私钥获取账号信息出错:%v", err),
		}
	}
	return multiplyAccountGo{
		Address:    address,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
}

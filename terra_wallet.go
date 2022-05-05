package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/FactomProject/go-bip32"
	"github.com/btcsuite/btcutil/bech32"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/ripemd160"
	"log"
)

const (
	terra_prefix string = "terra"
	pubKey_size  int    = 33
	TypeTerra    uint32 = 0x8000014A
	Bip44Purpose uint32 = 0x8000002C
)

var TERRA *TERRAWallet

func init() {
	TERRA = NewTerraWallet()
}

type TERRAWallet struct{}

/**
实例化solana钱包
*/
func NewTerraWallet() *TERRAWallet {
	return &TERRAWallet{}
}

func sha256Hash(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

//从33字节公钥获取地址
func getTerraAddress(pubKey []byte) (string, error) {
	if len(pubKey) != pubKey_size {
		return "", fmt.Errorf("len(pubKey) != 33,len = %d", len(pubKey))
	}
	sha := sha256Hash(pubKey)
	hasherRIPEMD160 := ripemd160.New()
	hasherRIPEMD160.Write(sha[:]) // does not error
	rip160Hash := hasherRIPEMD160.Sum(nil)

	dataBytes, err := bech32.ConvertBits(rip160Hash, 8, 5, true)
	if err != nil {
		return "", fmt.Errorf("ConvertBits error.err = %s", err.Error())
	}
	return bech32.Encode(terra_prefix, dataBytes)
}
func (t *TERRAWallet) createAccount() multiplyAccountGo {
	acc := t.createTerAccount()
	if acc.ErrorMessage != "" {
		acc.ErrorCode = "A0001"
		acc.ErrorMessage = fmt.Sprintf("生成用户账户出错:%s", acc.ErrorMessage)
		log.Printf("%s\n", acc.ErrorMessage)
	} else {
		acc.ErrorCode = ""
	}
	return *acc
}

func (t *TERRAWallet) createTerAccount() *multiplyAccountGo {
	res := &multiplyAccountGo{
		ErrorCode:    "succ",
		ErrorMessage: "",
		Address:      "",
		PrivateKey:   "",
		PublicKey:    "",
		Mnemonic:     "",
	}
	entroy, err := bip39.NewEntropy(256)
	if err != nil {
		res.ErrorCode = "err"
		res.ErrorMessage = fmt.Sprintf("NewEntropy error.err = %s", err.Error())
		return res
	}
	strMnemonic := ""
	strMnemonic, err = bip39.NewMnemonic(entroy)
	if err != nil {
		res.ErrorCode = "err"
		res.ErrorMessage = fmt.Sprintf("NewMnemonic error. err = %s", err.Error())
		return res
	}

	var seeds []byte
	seeds, err = bip39.NewSeedWithErrorChecking(strMnemonic, "")
	if err != nil {
		res.ErrorCode = "err"
		res.ErrorMessage = fmt.Sprintf("NewSeedWithErrorChecking error. err = %s", err.Error())
		return res
	}
	//fmt.Println("seed len = ", len(seeds))
	extKey, errC := createBip44Addr(seeds, TypeTerra, 0x80000000, 0, 0)
	if errC != nil {
		res.ErrorCode = "err"
		res.ErrorMessage = fmt.Sprintf("createBip44Addr error. err = %s", errC.Error())
		return res
	}
	res.PrivateKey = hex.EncodeToString(extKey.Key)
	pubKeyBytes := extKey.PublicKey().Key
	res.PublicKey = hex.EncodeToString(pubKeyBytes)
	res.Mnemonic = strMnemonic
	if res.Address, err = getTerraAddress(pubKeyBytes); err != nil {
		res.ErrorCode = "err"
		res.ErrorMessage = fmt.Sprintf("getTerraAddress error. err = %s", err.Error())
	}
	return res

}

func (t *TERRAWallet) createAccountByMnemonic(strmenmonic string) multiplyAccountGo {
	acc := t.createTerAccountByMnemonic(strmenmonic)
	if acc.ErrorMessage != "" {
		if "M0001" != acc.ErrorCode {
			acc.ErrorCode = "A0001"
			acc.ErrorMessage = fmt.Sprintf("生成用户账户出错:%s", acc.ErrorMessage)
		} else {
			acc.ErrorMessage = fmt.Sprintf("通过助记词获取账户信息出错:%s", acc.ErrorMessage)
		}
		log.Printf("%s\n", acc.ErrorMessage)
	} else {
		acc.ErrorCode = ""
	}
	return *acc

}

func (t *TERRAWallet) createTerAccountByMnemonic(menmonic string) *multiplyAccountGo {
	res := &multiplyAccountGo{
		ErrorCode:    "succ",
		ErrorMessage: "",
		Address:      "",
		PrivateKey:   "",
		PublicKey:    "",
		Mnemonic:     menmonic,
	}
	seeds, err := bip39.NewSeedWithErrorChecking(menmonic, "")
	if err != nil {
		res.ErrorCode = "M0001"
		res.ErrorMessage = fmt.Sprintf("NewSeedWithErrorChecking error.err = %s", err.Error())
		return res
	}
	fmt.Println("seed from menmonic,len = ", len(seeds))
	extKey, errC := createBip44Addr(seeds, TypeTerra, 0x80000000, 0, 0)
	if errC != nil {
		res.ErrorCode = "err"
		res.ErrorMessage = fmt.Sprintf("createBip44Addr error.err = %s", errC.Error())
		return res
	}
	res.PrivateKey = hex.EncodeToString(extKey.Key)
	pubKeyBytes := extKey.PublicKey().Key
	res.PublicKey = hex.EncodeToString(pubKeyBytes)

	if res.Address, err = getTerraAddress(pubKeyBytes); err != nil {
		res.ErrorCode = "err"
		res.ErrorMessage = fmt.Sprintf("getTerraAddress error. err = %s", err.Error())
	}
	return res
}
func (t *TERRAWallet) createAccountByPrivateKey(strPrivateKey string) multiplyAccountGo {
	acc := t.createTerAccountByPrivateKey(strPrivateKey)
	if acc.ErrorMessage != "" {
		acc.ErrorCode = "A0001"
		acc.ErrorMessage = fmt.Sprintf("生成用户账户出错:%s", acc.ErrorMessage)
		log.Printf("%s\n", acc.ErrorMessage)
	} else {
		acc.ErrorCode = ""
	}
	return *acc

}
func (t *TERRAWallet) createTerAccountByPrivateKey(privateKey string) *multiplyAccountGo {
	res := &multiplyAccountGo{
		ErrorCode:    "succ",
		ErrorMessage: "",
		Address:      "",
		PrivateKey:   privateKey,
		PublicKey:    "",
		Mnemonic:     "",
	}
	priKey, err := hex.DecodeString(privateKey)
	if err != nil {
		res.ErrorCode = "err"
		res.ErrorMessage = err.Error()
		return res
	}
	extKey := new(bip32.Key)
	extKey.IsPrivate = true
	extKey.Key = priKey
	pubKeyBytes := extKey.PublicKey().Key
	res.PublicKey = hex.EncodeToString(pubKeyBytes)
	if res.Address, err = getTerraAddress(pubKeyBytes); err != nil {
		res.ErrorCode = "err"
		res.ErrorMessage = fmt.Sprintf("getTerraAddress error. err = %s", err.Error())
	}
	return res
}

func createBip44Addr(seeds []byte, coinType uint32, account uint32, change, index uint32) (*bip32.Key, error) {
	masterKey, errP := bip32.NewMasterKey(seeds) //使用hmac.New(sha512.New, []byte("Bitcoin seed"))来创建
	if errP != nil {
		return nil, fmt.Errorf(" bip32.NewMasterKey error.err = %s", errP.Error())
	}
	child, err := masterKey.NewChildKey(Bip44Purpose)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(coinType)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(account)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(change)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(index)
	if err != nil {
		return nil, err
	}

	return child, nil
}
func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

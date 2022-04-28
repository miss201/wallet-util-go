/**
 * @Author: He
 * @Description:
 * @Date: 2022/04/27 10:48
 */

package main

import (
	"bytes"
	"fmt"
	"github.com/miss201/wallet-util-go/common/base58"
	"github.com/miss201/wallet-util-go/go-hdwallet"
	"github.com/miss201/wallet-util-go/solana"
	"github.com/tyler-smith/go-bip39"
	"log"
)

var SQLW *SQLWallet

func init() {
	SQLW = NewSqlWallet()
}

//solana钱包结构体，暂时为空，以后添加转账 签名等方法时候再来填充
type SQLWallet struct{}

/**
实例化solana钱包
*/
func NewSqlWallet() *SQLWallet {
	return &SQLWallet{}
}

/**
创建solana钱包账户
*/
func (SQLw *SQLWallet) createAccount() multiplyAccountGo {
	mnemonic, err := hdwallet.NewMnemonic(12)
	if err != nil {
		log.Printf("生成用户账户出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "A0001",
			ErrorMessage: fmt.Sprintf("生成用户账户出错:%v", err),
		}
	}
	return SQLw.createAccountByMnemonic(mnemonic)
}

/**
通过助记词返回账户
*/
func (SQLw *SQLWallet) createAccountByMnemonic(mnemonic string) multiplyAccountGo {
	var passPhrase []byte
	seed := bip39.NewSeed(mnemonic, string(passPhrase))
	derivationPath := "m/44'/501'/0'/0'"
	prvKey, err := derivation.DeriveForPath(derivationPath, seed)
	if err != nil {
		log.Printf("通过助记词获取账户信息出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "M0001",
			ErrorMessage: fmt.Sprintf("通过助记词获取账户信息出错:%v", err),
		}
	}
	pubKey, err := prvKey.PublicKey()
	if err != nil {
		log.Printf("通过助记词获取账户信息出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "M0001",
			ErrorMessage: fmt.Sprintf("通过助记词获取账户信息出错:%v", err),
		}
	}
	outPrv := base58.Encode(bytes.Join([][]byte{prvKey.Key, pubKey}, nil))
	outPub := base58.Encode(pubKey)

	//fmt.Println("mne==", mnemonic)
	//fmt.Println("prv==", outPrv)
	//fmt.Println("pub==", outPub)
	return multiplyAccountGo{
		Address:    outPub,
		PrivateKey: outPrv,
		PublicKey:  outPub,
		Mnemonic:   mnemonic,
	}
}

/**
通过私钥返回账户
*/
func (SQLw *SQLWallet) createAccountByPrivateKey(private string) multiplyAccountGo {
	//privateKey []byte通过base58转换出来的 ，因此这里要转换回去
	outPrv := base58.Decode(private)

	publicKey := make([]byte, 32)
	copy(publicKey, outPrv[32:])
	outPub := base58.Encode(publicKey)

	return multiplyAccountGo{
		Address:    outPub,
		PublicKey:  outPub,
		PrivateKey: private,
	}
}

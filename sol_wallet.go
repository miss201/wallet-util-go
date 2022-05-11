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

var SOLW *SOLWallet

func init() {
	SOLW = NewSolWallet()
}

//solana钱包结构体，暂时为空，以后添加转账 签名等方法时候再来填充
type SOLWallet struct{}

/**
实例化solana钱包
*/
func NewSolWallet() *SOLWallet {
	return &SOLWallet{}
}

/**
创建solana钱包账户
*/
func (SOLw *SOLWallet) createAccount() multiplyAccountGo {
	mnemonic, err := hdwallet.NewMnemonic(12)
	if err != nil {
		log.Printf("生成用户账户出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "A0001",
			ErrorMessage: fmt.Sprintf("生成用户账户出错:%v", err),
		}
	}
	return SOLw.createAccountByMnemonic(mnemonic)
}

/**
通过助记词返回账户
*/
func (SOLw *SOLWallet) createAccountByMnemonic(mnemonic string) multiplyAccountGo {
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
func (SQLw *SOLWallet) createAccountByPrivateKey(private string) multiplyAccountGo {
	if private == "" || len(private) != 88 {
		log.Printf("通过私钥获取账户信息出错，私钥长度不是88或者私钥为空")
		return multiplyAccountGo{
			ErrorCode:    "M0001",
			ErrorMessage: fmt.Sprintf("通过私钥获取账户信息出错，私钥长度异常"),
		}
	}
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

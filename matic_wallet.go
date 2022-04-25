/**
 * @Author: He
 * @Description:
 * @Date: 2022/04/20 14:51
 */

package main

import "C"
import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	w_common "github.com/miss201/wallet-util-go/common"
	"github.com/miss201/wallet-util-go/common/ec"
	"log"
)

var (
	MATIC = "MATIC"

	MATICW *MATICWallet
)

func init() {
	MATICW = NewMATICWallet(MATIC)
}

type MATICWallet struct {
	wallet      *w_common.Wallet
	mnemonicLen int
}

func NewMATICWallet(wc string) *MATICWallet {
	newWallet := MATICWallet{}
	switch wc {
	case MATIC:
		newWallet = MATICWallet{
			wallet:      w_common.NewWallet(w_common.MATIC, false, false, nil),
			mnemonicLen: 12,
		}
	default:
		newWallet = MATICWallet{
			wallet:      w_common.NewWallet(w_common.MATIC, false, false, nil),
			mnemonicLen: 12,
		}
	}
	return &newWallet
}

func (MATICw *MATICWallet) GenerateAddressFromMnemonic(mnemonic, language string) (string, error) {
	return MATICW.wallet.GenerateAddressFromMnemonic(mnemonic, language)
}

func (MATICw *MATICWallet) GenerateAddressFromPrivateKey(privateKey string) (string, error) {
	privateKeyBytes, _ := hex.DecodeString(privateKey)
	// (1) new  *PrivateKey、 *PublicKey
	_, publicKey := ec.PrivKeyFromBytes(privateKeyBytes)
	if publicKey == nil {
		return "", nil
	}
	/*publicKey Compressed 33 bytes
	publicKeyStr := hex.EncodeToString(publicKey.SerializeCompressed())*/

	// (2) pubBytes为04 开头的65字节公钥,去掉04后剩下64字节进行Keccak256运算
	pubBytes := crypto.Keccak256(publicKey.SerializeUnCompressed()[1:])
	// (3) 经过Keccak256运算后变成32字节，最终取这32字节的后20字节作为真正的地址
	address := common.BytesToAddress(pubBytes[12:])

	return address.Hex(), nil
}

func (MATICw *MATICWallet) ExportPrivateKeyFromMnemonic(mnemonic, language string) (string, error) {
	return MATICw.wallet.ExportPrivateKeyFromMnemonic(mnemonic, language)
}

func (MATICw *MATICWallet) CheckPrivateKey(privateKey string) (bool, error) {
	//去掉0x（如有）
	rm0xaddr := w_common.RemoveOxFromHex(privateKey)
	//判断长度
	if len(rm0xaddr) != w_common.ETHPRIVATEKEYLEN {
		return false, w_common.ERR_INVALID_PRIVATEKEY_LEN
	}
	//判断stringTohex是否成功
	_, err := hex.DecodeString(rm0xaddr)
	if err != nil {
		return false, w_common.ERR_INVALID_PRIVATEKEY
	}
	return true, nil
}

func (MATICw *MATICWallet) GetPubKeyFromPrivateKey(privateKey string) (string, error) {
	isValid, err := MATICw.CheckPrivateKey(privateKey)
	if isValid == false || err != nil {
		return "", err
	}

	privateKeyBytes, _ := hex.DecodeString(privateKey)
	// (1) new  *PrivateKey、 *PublicKey
	_, publicKey := ec.PrivKeyFromBytes(privateKeyBytes)
	if publicKey == nil {
		return "", nil
	}
	//publicKey Compressed 33 bytes
	return hex.EncodeToString(publicKey.SerializeCompressed()), nil
}

func (MATICw *MATICWallet) createAccount() multiplyAccountGo {
	menmonic, err := MATICw.wallet.GenerateMnemonic(12)
	privateKey, err := MATICw.ExportPrivateKeyFromMnemonic(menmonic, w_common.English)
	publicKey, err := MATICw.GetPubKeyFromPrivateKey(privateKey)
	address, err := MATICw.GenerateAddressFromPrivateKey(privateKey)
	if err != nil {
		log.Printf("生成用户账户出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "A0001",
			ErrorMessage: fmt.Sprintf("生成用户出错:%v", err),
		}
	}
	return multiplyAccountGo{
		Address:    address,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Mnemonic:   menmonic,
	}
}

func (MATICw *MATICWallet) createAccountByMenmonic(menmonic string) multiplyAccountGo {
	privateKey, err := MATICw.ExportPrivateKeyFromMnemonic(menmonic, w_common.English)
	publicKey, err := MATICw.GetPubKeyFromPrivateKey(privateKey)
	address, err := MATICw.GenerateAddressFromPrivateKey(privateKey)
	if err != nil {
		log.Printf("通过助记词获取账户信息出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "M0001",
			ErrorMessage: fmt.Sprintf("通过助记词获取账户信息出错:%v", err),
		}
	}
	return multiplyAccountGo{
		Address:    address,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Mnemonic:   menmonic,
	}
}

func (MATICw *MATICWallet) createAccountByPrivateKey(privateKey string) multiplyAccountGo {
	publicKey, err := MATICw.GetPubKeyFromPrivateKey(privateKey)
	address, err := MATICw.GenerateAddressFromPrivateKey(privateKey)
	if err != nil {
		log.Printf("通过私钥获取账号信息出错：%v\n", err)
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

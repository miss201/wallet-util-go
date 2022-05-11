package main

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
	BNB = "BNB"

	BNBW *BNBWallet
)

func init() {
	BNBW = NewBNBWallet(BNB)
}

//ETHWallet
type BNBWallet struct {
	wallet      *w_common.Wallet
	mnemonicLen int
}

func NewBNBWallet(wc string) *BNBWallet {
	newWallet := BNBWallet{}
	switch wc {
	case BNB:
		newWallet = BNBWallet{
			wallet:      w_common.NewWallet(w_common.BNB, false, false, nil),
			mnemonicLen: 12,
		}
	default:
		newWallet = BNBWallet{
			wallet:      w_common.NewWallet(w_common.BNB, false, false, nil),
			mnemonicLen: 12,
		}
	}
	return &newWallet
}

func (BNBw *BNBWallet) GenerateAddressFromMnemonic(mnemonic, language string) (string, error) {
	return BNBw.wallet.GenerateAddressFromMnemonic(mnemonic, language)
}

func (BNBw *BNBWallet) GenerateAddressFromPrivateKey(privateKey string) (string, error) {
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

func (BNBw *BNBWallet) ExportPrivateKeyFromMnemonic(mnemonic, language string) (string, error) {
	return BNBw.wallet.ExportPrivateKeyFromMnemonic(mnemonic, language)
}

func (BNBw *BNBWallet) CheckPrivateKey(privateKey string) (bool, error) {
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

func (BNBw *BNBWallet) GetPubKeyFromPrivateKey(privateKey string) (string, error) {

	isValid, err := BNBw.CheckPrivateKey(privateKey)
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

func (BNBw *BNBWallet) createAccount() multiplyAccountGo {
	mnemonic, err := BNBw.wallet.GenerateMnemonic(12)
	if err != nil {
		log.Printf("生成用户账户出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "A0001",
			ErrorMessage: fmt.Sprintf("生成用户出错:%v", err),
		}
	}
	privateKey, err := BNBw.ExportPrivateKeyFromMnemonic(mnemonic, w_common.English)
	if err != nil {
		log.Printf("生成用户账户出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "A0001",
			ErrorMessage: fmt.Sprintf("生成用户出错:%v", err),
		}
	}
	publicKey, err := BNBw.GetPubKeyFromPrivateKey(privateKey)
	if err != nil {
		log.Printf("生成用户账户出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "A0001",
			ErrorMessage: fmt.Sprintf("生成用户出错:%v", err),
		}
	}
	address, err := BNBw.GenerateAddressFromPrivateKey(privateKey)
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
		Mnemonic:   mnemonic,
	}
}

func (BNBw *BNBWallet) createAccountByMnemonic(mnemonic string) multiplyAccountGo {
	privateKey, err := BNBw.ExportPrivateKeyFromMnemonic(mnemonic, w_common.English)
	if err != nil {
		log.Printf("通过助记词获取账户信息出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "M0001",
			ErrorMessage: fmt.Sprintf("通过助记词获取账户信息出错:%v", err),
		}
	}
	publicKey, err := BNBw.GetPubKeyFromPrivateKey(privateKey)
	if err != nil {
		log.Printf("通过助记词获取账户信息出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "M0001",
			ErrorMessage: fmt.Sprintf("通过助记词获取账户信息出错:%v", err),
		}
	}
	address, err := BNBw.GenerateAddressFromPrivateKey(privateKey)
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
		Mnemonic:   mnemonic,
	}
}

func (BNBw *BNBWallet) createAccountByPrivateKey(privateKey string) multiplyAccountGo {
	publicKey, err := BNBw.GetPubKeyFromPrivateKey(privateKey)
	if err != nil {
		log.Printf("通过私钥获取账号信息出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "M0001",
			ErrorMessage: fmt.Sprintf("通过私钥获取账号信息出错:%v", err),
		}
	}
	address, err := BNBw.GenerateAddressFromPrivateKey(privateKey)
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

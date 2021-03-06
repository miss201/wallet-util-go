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
	ETH       = "ETH"
	ETHLedger = "ETHLedger"

	ETHW       *ETHWallet
	ETHLedgerW *ETHWallet
)

func init() {
	ETHW = NewETHWallet(ETH)
	ETHLedgerW = NewETHWallet(ETHLedger)
}

//ETHWallet
type ETHWallet struct {
	wallet      *w_common.Wallet
	mnemonicLen int
}

func NewETHWallet(wc string) *ETHWallet {

	newWallet := ETHWallet{}
	switch wc {
	case ETH:
		newWallet = ETHWallet{
			wallet:      w_common.NewWallet(w_common.ETH, false, false, nil),
			mnemonicLen: 12,
		}
	case ETHLedger:
		newWallet = ETHWallet{
			wallet:      w_common.NewWallet(w_common.ETH, false, false, nil),
			mnemonicLen: 12,
		}
	default:
		newWallet = ETHWallet{
			wallet:      w_common.NewWallet(w_common.ETH, false, false, nil),
			mnemonicLen: 12,
		}
	}
	return &newWallet
}

func (ETHw *ETHWallet) GenerateAddressFromMnemonic(mnemonic, language string) (string, error) {
	return ETHw.wallet.GenerateAddressFromMnemonic(mnemonic, language)
}

func (ETHw *ETHWallet) GenerateAddressFromPrivateKey(privateKey string) (string, error) {

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

func (ETHw *ETHWallet) ExportPrivateKeyFromMnemonic(mnemonic, language string) (string, error) {
	return ETHw.wallet.ExportPrivateKeyFromMnemonic(mnemonic, language)
}

func (ETHw *ETHWallet) CheckAddress(address string) (bool, error) {
	//去掉0x（如有）
	rm0xaddr := w_common.RemoveOxFromHex(address)
	//判断长度
	if len(rm0xaddr) != w_common.ETHADDRESSLEN {
		return false, w_common.ERR_ADDRESS_LEN
	}
	//判断stringTohex是否成功
	_, err := hex.DecodeString(rm0xaddr)
	if err != nil {
		return false, w_common.ERR_INVALID_ADDRESS
	}
	return true, nil
}

func (ETHw *ETHWallet) CheckPrivateKey(privateKey string) (bool, error) {
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

func (ETHw *ETHWallet) GetPubKeyFromPrivateKey(privateKey string) (string, error) {

	isValid, err := ETHw.CheckPrivateKey(privateKey)
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

func (ETHw *ETHWallet) createAccount() multiplyAccountGo {
	mnemonic, err := ETHw.wallet.GenerateMnemonic(12)
	if err != nil {
		log.Printf("生成用户账户出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "A0001",
			ErrorMessage: fmt.Sprintf("生成用户助记词出错:%v", err),
		}
	}
	privateKey, err := ETHw.ExportPrivateKeyFromMnemonic(mnemonic, w_common.English)
	if err != nil {
		log.Printf("生成用户账户私钥出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "A0001",
			ErrorMessage: fmt.Sprintf("生成用户出错:%v", err),
		}
	}
	publicKey, err := ETHw.GetPubKeyFromPrivateKey(privateKey)
	if err != nil {
		log.Printf("生成用户账户公钥出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "A0001",
			ErrorMessage: fmt.Sprintf("生成用户出错:%v", err),
		}
	}
	address, err := ETHw.GenerateAddressFromPrivateKey(privateKey)
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

func (ETHw *ETHWallet) createAccountByMnemonic(mnemonic string) multiplyAccountGo {
	privateKey, err := ETHw.ExportPrivateKeyFromMnemonic(mnemonic, w_common.English)
	if err != nil {
		log.Printf("通过助记词获取账户私钥信息出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "M0001",
			ErrorMessage: fmt.Sprintf("通过助记词获取账户信息出错:%v", err),
		}
	}
	publicKey, err := ETHw.GetPubKeyFromPrivateKey(privateKey)
	if err != nil {
		log.Printf("通过助记词获取账户公钥信息出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "M0001",
			ErrorMessage: fmt.Sprintf("通过助记词获取账户信息出错:%v", err),
		}
	}
	address, err := ETHw.GenerateAddressFromPrivateKey(privateKey)
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

func (ETHw *ETHWallet) createAccountByPrivateKey(privateKey string) multiplyAccountGo {
	publicKey, err := ETHw.GetPubKeyFromPrivateKey(privateKey)
	if err != nil {
		log.Printf("通过私钥获取账号公钥信息出错：%v\n", err)
		return multiplyAccountGo{
			ErrorCode:    "M0001",
			ErrorMessage: fmt.Sprintf("通过私钥获取账号信息出错:%v", err),
		}
	}
	address, err := ETHw.GenerateAddressFromPrivateKey(privateKey)
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

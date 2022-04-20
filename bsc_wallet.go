package wallet_util_go

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	w_common "github.com/miss201/wallet-util-go/common"
	"github.com/miss201/wallet-util-go/common/ec"
)

var (
	BSC = "BSC"

	BSCW *BSCWallet
)

func init() {
	BSCW = NewBSCWallet(BSC)
}

//ETHWallet
type BSCWallet struct {
	wallet      *w_common.Wallet
	mnemonicLen int
}

func NewBSCWallet(wc string) *BSCWallet {

	newWallet := BSCWallet{}
	switch wc {
	case BSC:
		newWallet = BSCWallet{
			wallet:      w_common.NewWallet(w_common.BSC, false, false, nil),
			mnemonicLen: 12,
		}
	default:
		newWallet = BSCWallet{
			wallet:      w_common.NewWallet(w_common.BSC, false, false, nil),
			mnemonicLen: 12,
		}
	}
	return &newWallet
}

func (BSCw *BSCWallet) GenerateAddressFromMnemonic(mnemonic, language string) (string, error) {
	return BSCw.wallet.GenerateAddressFromMnemonic(mnemonic, language)
}

func (BSCw *BSCWallet) GenerateAddressFromPrivateKey(privateKey string) (string, error) {

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

func (BSCw *BSCWallet) ExportPrivateKeyFromMnemonic(mnemonic, language string) (string, error) {
	return BSCw.wallet.ExportPrivateKeyFromMnemonic(mnemonic, language)
}

func (BSCw *BSCWallet) CheckPrivateKey(privateKey string) (bool, error) {
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

func (BSCw *BSCWallet) GetPubKeyFromPrivateKey(privateKey string) (string, error) {

	isValid, err := BSCw.CheckPrivateKey(privateKey)
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

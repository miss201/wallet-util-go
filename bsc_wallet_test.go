package wallet_util_go

import (
	"github.com/miss201/wallet-util-go/common"
	"testing"
)

func TestBSCGenerateAddressFromMnemonic(t *testing.T) {

	mnemonic := "wreck bullet carpet nerve belt border often trust exchange believe defense rebel"

	BSCAddress, err := BSCW.GenerateAddressFromMnemonic(mnemonic, common.English)
	if err != nil {
		t.Errorf("Failed to TestGenerateAddressFromMnemonic: %v", err)
	}
	//0x81FD1F7aE91041aAc5fCF7d8Ed3e1dd88Cc1359a
	t.Log("TestImportWalletFromMnemonic , BSCAddress=", BSCAddress)

	BSCPrivateKey, err := BSCW.ExportPrivateKeyFromMnemonic(mnemonic, common.English)
	if err != nil {
		t.Errorf("Failed to ExportPrivateKeyFromMnemonic：%v", err)
	}
	t.Log("ExportPrivateKeyFromMnemonic,BSCPrivateKey=", BSCPrivateKey)

	BSCPublicKey, err := BSCW.GetPubKeyFromPrivateKey(BSCPrivateKey)
	if err != nil {
		t.Errorf("Failed to GetPubKeyFromPrivateKey：%v", err)
	}
	t.Log("GetPubKeyFromPrivateKey,BSCPublicKey=", BSCPublicKey)
}

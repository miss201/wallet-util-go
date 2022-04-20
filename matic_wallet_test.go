package wallet_util_go

import (
	"github.com/miss201/wallet-util-go/common"
	"testing"
)

func TestMATICGenerateAddressFromMnemonic(t *testing.T) {

	mnemonic := "wreck bullet carpet nerve belt border often trust exchange believe defense rebel"

	MATICAddress, err := MATICW.GenerateAddressFromMnemonic(mnemonic, common.English)
	if err != nil {
		t.Errorf("Failed to TestGenerateAddressFromMnemonic: %v", err)
	}
	//0x81FD1F7aE91041aAc5fCF7d8Ed3e1dd88Cc1359a
	t.Log("TestImportWalletFromMnemonic , MATICAddress=", MATICAddress)

	MATICPrivateKey, err := MATICW.ExportPrivateKeyFromMnemonic(mnemonic, common.English)
	if err != nil {
		t.Errorf("Failed to ExportPrivateKeyFromMnemonic：%v", err)
	}
	t.Log("ExportPrivateKeyFromMnemonic,MATICPrivateKey=", MATICPrivateKey)

	MATICPublicKey, err := MATICW.GetPubKeyFromPrivateKey(MATICPrivateKey)
	if err != nil {
		t.Errorf("Failed to GetPubKeyFromPrivateKey：%v", err)
	}
	t.Log("GetPubKeyFromPrivateKey,MATICPublicKey=", MATICPublicKey)
}

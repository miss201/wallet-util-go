package main

import (
	"github.com/miss201/wallet-util-go/common"
	"testing"
)

func TestBNBGenerateAddressFromMnemonic(t *testing.T) {

	mnemonic := "tail merge cousin betray engage yard slab tube hobby shove settle legal"

	BNBAddress, err := BNBW.GenerateAddressFromMnemonic(mnemonic, common.English)
	if err != nil {
		t.Errorf("Failed to TestGenerateAddressFromMnemonic: %v", err)
	}
	//0x81FD1F7aE91041aAc5fCF7d8Ed3e1dd88Cc1359a
	t.Log("TestImportWalletFromMnemonic , BNBAddress=", BNBAddress)

	BNBPrivateKey, err := BNBW.ExportPrivateKeyFromMnemonic(mnemonic, common.English)
	if err != nil {
		t.Errorf("Failed to ExportPrivateKeyFromMnemonic：%v", err)
	}
	t.Log("ExportPrivateKeyFromMnemonic,BNBPrivateKey=", BNBPrivateKey)

	BNBPublicKey, err := BNBW.GetPubKeyFromPrivateKey(BNBPrivateKey)
	if err != nil {
		t.Errorf("Failed to GetPubKeyFromPrivateKey：%v", err)
	}
	t.Log("GetPubKeyFromPrivateKey,BNBPublicKey=", BNBPublicKey)
}

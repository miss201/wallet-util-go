package wallet_util_go

import (
	"github.com/miss201/wallet-util-go/common"
	"testing"
)

func TestETHGenerateAddressFromMnemonic(t *testing.T) {
	mnemonic := "tail merge cousin betray engage yard slab tube hobby shove settle legal"

	ETHAddress, err := ETHW.GenerateAddressFromMnemonic(mnemonic, common.English)
	if err != nil {
		t.Errorf("Failed to TestGenerateAddressFromMnemonic: %v", err)
	}
	ETHAddressLedger, err := ETHLedgerW.GenerateAddressFromMnemonic(mnemonic, common.English)
	if err != nil {
		t.Errorf("Failed to TestGenerateAddressFromMnemonic: %v", err)
	}

	//0x4fB085B8478584B10aF0b1b7401c6DcE38bFcA7d
	t.Log("TestImportWalletFromMnemonic , ETHAddress=", ETHAddress)
	//0x4fB085B8478584B10aF0b1b7401c6DcE38bFcA7d
	t.Log("TestImportWalletFromMnemonic , ETHAddressLedger=", ETHAddressLedger)
}

func TestETHGenerateAddressFromPrivateKey(t *testing.T) {

	privkey := "683a28a354be513d8e808e0f9d07ce217aaf18dad0ae855f51532a05b653ad0b"

	//0x4fB085B8478584B10aF0b1b7401c6DcE38bFcA7d
	address, err := ETHW.GenerateAddressFromPrivateKey(privkey)
	if err != nil {
		t.Error("GenerateAddressFromPrivateKey err:", err)
	}
	t.Log("address :", address)
}

func TestETHExportPrivateKeyFromMnemonic(t *testing.T) {
	mnemonic := "tail merge cousin betray engage yard slab tube hobby shove settle legal"
	ETHPrivateKey, err := ETHW.ExportPrivateKeyFromMnemonic(mnemonic, common.English)
	if err != nil {
		t.Errorf("Failed to TestExportPrivateKeyFromMnemonic: %v", err)
	}
	ETHPrivateKeyLedger, err := ETHLedgerW.ExportPrivateKeyFromMnemonic(mnemonic, common.English)
	if err != nil {
		t.Errorf("Failed to TestExportPrivateKeyFromMnemonic: %v", err)
	}
	//494f8228ae5b6fda6bee1f44eb2c4ed120f210e06acaa8053763efb65638b315
	t.Log("ETHPrivateKey=", ETHPrivateKey)
	//0b98e389e449fa5f388f94bf702066e9ad373e19c2119076f0c276cdd50d776a
	t.Log("ETHPrivateKeyLedger=", ETHPrivateKeyLedger)

}

func TestCheckETHAddress(t *testing.T) {
	ETHAddress := "0x4fB085B8478584B10aF0b1b7401c6DcE38bFcA7d"
	isValid, err := ETHW.CheckAddress(ETHAddress)
	if err != nil {
		t.Error("CheckETHAddress err:", err)
	}
	t.Log("CheckETHAddress :", isValid)
}

func TestCheckETHPrivateKey(t *testing.T) {
	ETHPrivateKey := "6B93D965D9981F9066CCC44B9DBF32B50F411C0DCEDF4A41CA4E7424ABDB6112"
	isValid, err := ETHW.CheckPrivateKey(ETHPrivateKey)
	if err != nil {
		t.Errorf("Failed to Check ETHPrivateKey: %v", err)
	}
	t.Log("TestCheckAddress: ETHPrivateKey=", isValid)
}

func TestETHGetPubKeyFromPrivateKey(t *testing.T) {
	ETHPrivateKey := "0b98e389e449fa5f388f94bf702066e9ad373e19c2119076f0c276cdd50d776a"
	publicKey, err := ETHW.GetPubKeyFromPrivateKey(ETHPrivateKey)
	if err != nil {
		t.Errorf("Failed to GetPubKeyFromPrivateKey ETHPrivateKey: %v", err)
	}
	t.Log("GetPubKeyFromPrivateKey ETHPrivateKey=", publicKey)
}

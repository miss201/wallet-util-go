/**
 * @Author: He
 * @Description:
 * @Date: 2022/04/20 16:49
 */

package main

import (
	"fmt"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	account := CreateAccount("ETH")
	fmt.Println("Mnemonic==", account.Mnemonic)
	fmt.Println("PublicKey==", account.PublicKey)
	fmt.Println("PrivateKey==", account.PrivateKey)
	fmt.Println("Address==", account.Address)
}

func TestMnemonicToAccount(t *testing.T) {
	mnemonic := "badge trick champion uphold fragile midnight conduct adjust order subject distance define"
	account := MnemonicToAccount("BTC", mnemonic)
	fmt.Println("Mnemonic==", account.Mnemonic)
	fmt.Println("PublicKey==", account.PublicKey)
	fmt.Println("PrivateKey==", account.PrivateKey)
	fmt.Println("Address==", account.Address)
}

func TestGetAccountByPrivateKey(t *testing.T) {
	privateKey := "KzMUwZ2rymHNK8os2ApFQQCbG5Vg5DYoJD16GPE88pJMhxuhQjHT"
	account := GetAccountByPrivateKey("ETH", privateKey)
	fmt.Println("PublicKey==", account.PublicKey)
	fmt.Println("PrivateKey==", account.PrivateKey)
	fmt.Println("Address==", account.Address)
}

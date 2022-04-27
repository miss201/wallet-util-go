//go:build ignore
//+build ignore

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
	fmt.Println("ErrorMessage==", account.ErrorMessage)
	fmt.Println("ErrorCode==", account.ErrorCode)
	fmt.Println("Mnemonic==", account.Mnemonic)
	fmt.Println("PublicKey==", account.PublicKey)
	fmt.Println("PrivateKey==", account.PrivateKey)
	fmt.Println("Address==", account.Address)
}

func TestGetAccountByPrivateKey(t *testing.T) {
	privateKey := "683a28a354be513d8e808e0f9d07ce217aaf18dad0ae855f51532a05b653ad0b"
	account := GetAccountByPrivateKey("ETH", privateKey)
	fmt.Println("PublicKey==", account.PublicKey)
	fmt.Println("PrivateKey==", account.PrivateKey)
	fmt.Println("Address==", account.Address)
}

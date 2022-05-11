//go:build ignore
// +build ignore

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
	account := CreateAccount("LUNA")
	fmt.Println("ErrorMessage==", account.ErrorMessage)
	fmt.Println("ErrorCode==", account.ErrorCode)
	fmt.Println("Mnemonic==", account.Mnemonic)     //worth crop scan company resemble latin dune potato ketchup enforce burden flavor
	fmt.Println("PublicKey==", account.PublicKey)   //CyVcHkqYs6oogSoFZdWRmVYkufBjeVa86WxEux9YBgxU
	fmt.Println("PrivateKey==", account.PrivateKey) //2MSLtHbsTTRbb8e3YxLseKkaCwhbEyu8Y5ZFJQAYzSJ2Ro7DgX1D2wd7eGdXPZ4dsENw4f41GXPNytWraD8Ftf6U
	fmt.Println("Address==", account.Address)
}

func TestMnemonicToAccount(t *testing.T) {
	mnemonic := "worth crop scan company resemble latin dune potato ketchup enforce burden flavor"
	account := MnemonicToAccount("SOL", mnemonic)
	fmt.Println("ErrorMessage==", account.ErrorMessage)
	fmt.Println("ErrorCode==", account.ErrorCode)
	fmt.Println("Mnemonic==", account.Mnemonic)
	fmt.Println("PublicKey==", account.PublicKey)   //CyVcHkqYs6oogSoFZdWRmVYkufBjeVa86WxEux9YBgxU
	fmt.Println("PrivateKey==", account.PrivateKey) //2MSLtHbsTTRbb8e3YxLseKkaCwhbEyu8Y5ZFJQAYzSJ2Ro7DgX1D2wd7eGdXPZ4dsENw4f41GXPNytWraD8Ftf6U
	fmt.Println("Address==", account.Address)
}

func TestGetAccountByPrivateKey(t *testing.T) {
	privateKey := "2MSLtHbsTTRbb8e3YxLseKkaCwhbEyu8Y5ZFJQAYzSJ2Ro7DgX1D2wd7eGdXPZ4dsENw4f41GXPNytWraD8Ftf6U"
	account := GetAccountByPrivateKey("SOL", privateKey)
	fmt.Println("PublicKey==", account.PublicKey)
	fmt.Println("PrivateKey==", account.PrivateKey)
	fmt.Println("Address==", account.Address)
}

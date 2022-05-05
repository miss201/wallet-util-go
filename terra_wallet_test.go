/**
 * @Author: He
 * @Description:
 * @Date: 2022/04/28 21:24
 */

package main

import (
	"fmt"
	"testing"
)

func TestNewTerraWallet(t *testing.T) {
	account := TERRA.createAccount()
	fmt.Println(account.PrivateKey)
	fmt.Println(account.Address)
	fmt.Println(account.Mnemonic)
}

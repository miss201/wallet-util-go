/**
 * @Author: He
 * @Description:
 * @Date: 2022/04/27 10:48
 */

package main

import (
	"fmt"
	"testing"
)

func TestNewSolWallet(t *testing.T) {
	solW := NewSolWallet()
	fmt.Println(solW)
}

func TestNewSolWallet4(t *testing.T) {
	account := SOLW.createAccount()
	fmt.Println("account.errorCode===", account.ErrorCode)
	fmt.Println("account.errorMsg===", account.ErrorMessage)
	fmt.Println("account.pri===", account.PrivateKey)    //535GqwYnWkaRLpBSWmGo164x1erjtc1UMc6crKSMjDvzgGDBvKqturmZECC36ex3CcD4uqe452dK2gahzzfVRtGi
	fmt.Println("account.pub===", account.PublicKey)     //5oLYrUp5AxoSTCV4TXgVQJ4g45CBMzkZVy4saRYuatUi
	fmt.Println("account.mnemonic===", account.Mnemonic) //swear army cement maze puppy want mystery bottom feed humble float oval
}

func TestNewSolWallet2(t *testing.T) {
	account := SOLW.createAccountByMnemonic("swear army cement maze puppy want mystery bottom feed humble float oval")
	fmt.Println("account.errorCode===", account.ErrorCode)
	fmt.Println("account.errorMsg===", account.ErrorMessage)
	fmt.Println("account.pri===", account.PrivateKey) //535GqwYnWkaRLpBSWmGo164x1erjtc1UMc6crKSMjDvzgGDBvKqturmZECC36ex3CcD4uqe452dK2gahzzfVRtGi
	fmt.Println("account.pub===", account.PublicKey)  //5oLYrUp5AxoSTCV4TXgVQJ4g45CBMzkZVy4saRYuatUi
	fmt.Println("account.mnemonic===", account.Mnemonic)
}

func TestNewSolWallet3(t *testing.T) {
	account := SOLW.createAccountByPrivateKey("535GqwYnWkaRLpBSWmGo164x1erjtc1UMc6crKSMjDvzgGDBvKqturmZECC36ex3CcD4uqe452dK2gahzzfVRtGi")
	fmt.Println("account.errorCode===", account.ErrorCode)
	fmt.Println("account.errorMsg===", account.ErrorMessage)
	fmt.Println("account.pri===", account.PrivateKey)
	fmt.Println("account.pub===", account.PublicKey) ////HgnmjgxnNc27bQrhHBskN2LDhheSAqgaY3ksydgg4gwd
	fmt.Println("account.mnemonic===", account.Mnemonic)
}

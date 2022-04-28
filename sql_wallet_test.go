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

func TestNewSqlWallet(t *testing.T) {
	sqlW := NewSqlWallet()
	fmt.Println(sqlW)
}

func TestNewSqlWallet4(t *testing.T) {
	account := SQLW.createAccount()
	fmt.Println("account.errorCode===", account.ErrorCode)
	fmt.Println("account.errorMsg===", account.ErrorMessage)
	fmt.Println("account.pri===", account.PrivateKey)    //535GqwYnWkaRLpBSWmGo164x1erjtc1UMc6crKSMjDvzgGDBvKqturmZECC36ex3CcD4uqe452dK2gahzzfVRtGi
	fmt.Println("account.pub===", account.PublicKey)     //5oLYrUp5AxoSTCV4TXgVQJ4g45CBMzkZVy4saRYuatUi
	fmt.Println("account.mnemonic===", account.Mnemonic) //swear army cement maze puppy want mystery bottom feed humble float oval
}

func TestNewSqlWallet2(t *testing.T) {
	account := SQLW.createAccountByMnemonic("swear army cement maze puppy want mystery bottom feed humble float oval")
	fmt.Println("account.errorCode===", account.ErrorCode)
	fmt.Println("account.errorMsg===", account.ErrorMessage)
	fmt.Println("account.pri===", account.PrivateKey) //535GqwYnWkaRLpBSWmGo164x1erjtc1UMc6crKSMjDvzgGDBvKqturmZECC36ex3CcD4uqe452dK2gahzzfVRtGi
	fmt.Println("account.pub===", account.PublicKey)  //5oLYrUp5AxoSTCV4TXgVQJ4g45CBMzkZVy4saRYuatUi
	fmt.Println("account.mnemonic===", account.Mnemonic)
}

func TestNewSqlWallet3(t *testing.T) {
	account := SQLW.createAccountByPrivateKey("535GqwYnWkaRLpBSWmGo164x1erjtc1UMc6crKSMjDvzgGDBvKqturmZECC36ex3CcD4uqe452dK2gahzzfVRtGi")
	fmt.Println("account.errorCode===", account.ErrorCode)
	fmt.Println("account.errorMsg===", account.ErrorMessage)
	fmt.Println("account.pri===", account.PrivateKey)
	fmt.Println("account.pub===", account.PublicKey) ////HgnmjgxnNc27bQrhHBskN2LDhheSAqgaY3ksydgg4gwd
	fmt.Println("account.mnemonic===", account.Mnemonic)
}

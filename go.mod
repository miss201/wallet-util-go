module github.com/miss201/wallet-util-go

go 1.12

replace (
	golang.org/x/crypto v0.0.0-20191227163750-53104e6ec876 => github.com/golang/crypto v0.0.0-20191227163750-53104e6ec876
	golang.org/x/crypto v0.0.0-20200115085410-6d4e4cb37c7d => github.com/golang/crypto v0.0.0-20200115085410-6d4e4cb37c7d
	golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/sync v0.0.0-20181221193216-37e7f081c4d4 => github.com/golang/sync v0.0.0-20181221193216-37e7f081c4d4
	golang.org/x/sys v0.3.0 => github.com/golang/sys v0.3.0
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)

require (
	github.com/FactomProject/basen v0.0.0-20150613233007-fe3947df716e // indirect
	github.com/FactomProject/btcutilecc v0.0.0-20130527213604-d3a63a5752ec // indirect
	github.com/FactomProject/go-bip32 v0.3.5
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/btcsuite/btcutil v1.0.1
	github.com/cmars/basen v0.0.0-20150613233007-fe3947df716e // indirect
	github.com/cpacia/bchutil v0.0.0-20181003130114-b126f6a35b6c
	github.com/ethereum/go-ethereum v1.9.10
	github.com/tyler-smith/go-bip39 v1.0.2
	golang.org/x/crypto v0.0.0-20200115085410-6d4e4cb37c7d
	launchpad.net/gocheck v0.0.0-20140225173054-000000000087 // indirect
)

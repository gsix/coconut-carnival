package main

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
)

func main() {
	//////BIP32 Root Key Mnemonic service////////
	var xprivnohard = "yprvAL1BKLLMXFYsoptMJ7P9Bq8rpbs5dDqdKQj3StAs2WptbsjRqSfimpfErPVPoQAYbUt2MC14NNXVHvuEehzXrDwYJhnzRCM5YCWLCTSAkiZ"

	masterprivkey, _ := hdkeychain.NewKeyFromString(xprivnohard)
	//m/49'
	acct, _ := masterprivkey.Child(hdkeychain.HardenedKeyStart + 49)
	//m/49'/0'
	acctExt, _ := acct.Child(hdkeychain.HardenedKeyStart + 0)
	//m/49'/0'/0'
	acctExtChild, _ := acctExt.Child(hdkeychain.HardenedKeyStart + 1)
	//m/49'/0'/0'/0
	acctExtChildExt, _ := acctExtChild.Child(0)
	//m/49'/0'/0'/0/0
	acctExtKey, _ := acctExtChildExt.Child(0)

	pubKey, err := acctExtKey.ECPubKey()
	if err != nil {
		panic(err)
	}

	pubKeyCompress, err := btcutil.NewAddressPubKey(pubKey.SerializeCompressed(), &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println(err)
		return
	}

	keyHash := btcutil.Hash160(pubKey.SerializeCompressed())
	scriptSig, err := txscript.NewScriptBuilder().AddOp(txscript.OP_0).AddData(keyHash).Script()
	if err != nil {
		panic(err)
	}
	acct0ExtAddr0, err := btcutil.NewAddressScriptHash(scriptSig, &chaincfg.MainNetParams)
	if err != nil {
		panic(err)
	}

	fmt.Println("Public key ", pubKeyCompress.String())
	fmt.Println("Segwit address ", acct0ExtAddr0)
}

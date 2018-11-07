package main

import (
	"fmt"

	"github.com/ltcsuite/ltcd/chaincfg"
	"github.com/ltcsuite/ltcd/txscript"
	"github.com/ltcsuite/ltcutil"
	"github.com/ltcsuite/ltcutil/hdkeychain"
)

//btc
/**/

func main() {
	var xprivnohard = "Ltpv71G8qDifUiNesUNyn2d7oSUMVhDrdcJ1ETiZyVxAvy8MD9C8A4PEKvk3CgnMBnm4FBEwWBgSJ9v7Y2q2ASGA22WeCHStyScj7EhUocQ1TJr"

	masterprivkey, _ := hdkeychain.NewKeyFromString(xprivnohard)

	acct10, _ := masterprivkey.Child(hdkeychain.HardenedKeyStart + 49)

	acct10Ext, _ := acct10.Child(hdkeychain.HardenedKeyStart + 2)

	acct10Ext33, _ := acct10Ext.Child(hdkeychain.HardenedKeyStart + 1)

	acct10Ext333, _ := acct10Ext33.Child(0)

	acct10Ext3334, _ := acct10Ext333.Child(0)

	pubKey, err := acct10Ext3334.ECPubKey()
	if err != nil {
		panic(err)
	}

	keyHash := ltcutil.Hash160(pubKey.SerializeCompressed())
	scriptSig, err := txscript.NewScriptBuilder().AddOp(txscript.OP_0).AddData(keyHash).Script()
	if err != nil {
		panic(err)
	}
	acct0ExtAddr0, err := ltcutil.NewAddressScriptHash(scriptSig, &chaincfg.MainNetParams)
	if err != nil {
		panic(err)
	}
	fmt.Println(acct0ExtAddr0)
}

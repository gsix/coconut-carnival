package main

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

func main() {
	// This comes from
	serializedPubKey := []byte{
		0x04, 0x11, 0xdb, 0x93, 0xe1, 0xdc, 0xdb, 0x8a,
		0x01, 0x6b, 0x49, 0x84, 0x0f, 0x8c, 0x53, 0xbc, 0x1e,
		0xb6, 0x8a, 0x38, 0x2e, 0x97, 0xb1, 0x48, 0x2e, 0xca,
		0xd7, 0xb1, 0x48, 0xa6, 0x90, 0x9a, 0x5c, 0xb2, 0xe0,
		0xea, 0xdd, 0xfb, 0x84, 0xcc, 0xf9, 0x74, 0x44, 0x64,
		0xf8, 0x2e, 0x16, 0x0b, 0xfa, 0x9b, 0x8b, 0x64, 0xf9,
		0xd4, 0xc0, 0x3f, 0x99, 0x9b, 0x86, 0x43, 0xf6, 0x56,
		0xb4, 0x12, 0xa3,
	}
	// Create a new btcutil.Address from the raw public key for the main
	// Bitcoin network.
	mainNetAddr, err := btcutil.NewAddressPubKey(serializedPubKey, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print address details about the original public key.
	fmt.Println("Original Pubkey:")
	fmt.Println("Bitcoin Address:", mainNetAddr.EncodeAddress())
	fmt.Println("Hex:", mainNetAddr.String())
	fmt.Printf("What goes in the transaction script: %x\n", mainNetAddr.ScriptAddress())

	// Set the format to compressed so the uncompressed public key is conveted
	// to a compressed public key and print details about it.
	fmt.Println("\nCompressed Form:")
	mainNetAddr.SetFormat(btcutil.PKFCompressed)
	fmt.Println("Bitcoin Address:", mainNetAddr.EncodeAddress())
	fmt.Println("Hex:", mainNetAddr.String())
	fmt.Printf("What goes in the transaction script: %x\n", mainNetAddr.ScriptAddress())
}

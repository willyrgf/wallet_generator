package main

import (
	"fmt"
	"log"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"golang.org/x/crypto/sha3"
)

func main() {
	// generate a new private key for bitcoin and ethereum.
	privateKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		log.Fatalf("error generating private key: %v", err)
	}

	// bitcoin address generation
	// derive the public key from the private key.
	btcPubKey := privateKey.PubKey()

	// generate a bitcoin wallet address from the public key.
	btcAddress, err := btcutil.NewAddressPubKey(btcPubKey.SerializeCompressed(), &chaincfg.MainNetParams)
	if err != nil {
		log.Fatalf("error generating bitcoin wallet address: %v", err)
	}

	// ethereum address generation
	// ethereum uses the keccak-256 hash of the uncompressed public key minus the first byte.
	ethPubKey := btcPubKey.SerializeUncompressed()[1:]
	hash := sha3.NewLegacyKeccak256()
	hash.Write(ethPubKey)
	ethAddress := hash.Sum(nil)[12:]

	fmt.Printf("bitcoin\n")
	fmt.Printf("private key (hex): %x\n", privateKey.Serialize())
	fmt.Printf("public address: %s\n\n", btcAddress.EncodeAddress())

	fmt.Printf("ethereum\n")
	fmt.Printf("private key (hex): %x\n", privateKey.Serialize())
	fmt.Printf("public address: 0x%x\n", ethAddress)
}

package main

import (
	"fmt"
	"github.com/flashbots/go-boost-utils/bls"
	boostTypes "github.com/flashbots/go-boost-utils/types"
)

func main() {
	sk, pubkey, err := bls.GenerateNewKeypair()
	if err != nil {
		fmt.Println(err)
	}
	//
	fmt.Println(sk.String())
	fmt.Println(pubkey.String())

	pk, err := boostTypes.BlsPublicKeyToPublicKey(pubkey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pk)
}

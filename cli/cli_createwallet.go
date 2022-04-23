package cli

import (
	"blockchain_go/wallet"
	"fmt"
)

func (cli *CLI) createWallet() {
	w, _ := wallet.NewWallet()
	address := w.GetAddress()

	fmt.Printf("Your TSS private key should be kept a secret. Whomever you share the private key with has access to spend all the bitcoins associated with that address.\n Your new address: %s\n Your private key: %s\n", address, wallet.EncodePrivKey(w.PrivateKey))
}

func (cli *CLI) AddWallet(priKey []byte) {

	wallet := wallet.DecodePrivKey(priKey)

	fmt.Printf("Your addres: %s\n", wallet.GetAddress())

}
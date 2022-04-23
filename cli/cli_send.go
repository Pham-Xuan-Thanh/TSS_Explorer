package cli

import (
	"blockchain_go/blockchain"
	"blockchain_go/wallet"
	"fmt"
	"log"
)

func (cli *CLI) send(prkFrom, to string, amount int, allowuser []string, iHash string, mineNow bool) {
	// if !wallet.ValidateAddress(prkFrom) {
	// 	log.Panic("ERROR: Sender address is not valid")
	// }
	if !wallet.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := blockchain.NewBlockchain("3000")
	UTXOSet := blockchain.UTXOSet{bc}
	FTX := blockchain.FTXset{bc}
	defer bc.DB.Close()

	// wallets, err :=
	// if err != nil {
	// 	log.Panic(err)
	// }
	w := wallet.DecodePrivKey([]byte(prkFrom))
	tx := blockchain.NewUTXOTransaction(w, to, amount, allowuser, iHash, &UTXOSet)

	if mineNow {
		cbTx := blockchain.NewCoinbaseTX(string(w.GetAddress()), "")
		txs := []*blockchain.Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
		FTX.UpdateFTX(newBlock)
	} else {
		sendTx(knownNodes[0], tx)
	}

	fmt.Println("Success!")
}
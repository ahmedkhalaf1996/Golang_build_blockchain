package main

import (
	"fmt"
	"log"

	"github.com/ahmedkhalaf1996/Golang_build_blockchain/block"
	"github.com/ahmedkhalaf1996/Golang_build_blockchain/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}
func main() {
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	// Wallet
	t := wallet.NewTransaction(walletM.PrivateKey(), walletM.PublicKey(), "sender", "recipient", 1.0)

	// Blockchain
	blockchain := block.NewBlockchain(walletM.BlockchainAddress())
	isAdded := blockchain.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0,
		walletA.PublicKey(), t.GenerateSignature())

	fmt.Println("Added? ", isAdded)

	// blockchain.Mining()
	// blockchain.Print()

	// fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress()))
	// fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockchainAddress()))
	// fmt.Printf("M %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockchainAddress()))

	// walletM := wallet.NewWallet()
	// t := wallet.NewTransaction(walletM.PrivateKey(), walletM.PublicKey(), "sender", "recipient", 1.0)
	// signature := t.GenerateSignature()

	// valid := ecdsa.Verify(walletM.PublicKey(), []byte("sample message"), signature.R, signature.S)
	// fmt.Println("Signature valid:", valid)

	// walletM := wallet.NewWallet()

	// // Create a transaction
	// t := wallet.NewTransaction(walletM.PrivateKey(), walletM.PublicKey(), "sender", "recipient", 1.0)

	// // Generate a signature
	// signature := t.GenerateSignature()

	// // Use the same data to verify the signature
	// transactionData, _ := json.Marshal(t)  // Convert the transaction to JSON for consistent hashing
	// hash := sha256.Sum256(transactionData) // Hash the same transaction data used during signature generation

	// // Convert fixed-size array to slice for ECDSA verification
	// hashSlice := hash[:]

	// // Verify the signature with correct data
	// isValid := ecdsa.Verify(walletM.PublicKey(), hashSlice, signature.R, signature.S)

	// // Output the result of signature verification
	// fmt.Println("Signature valid:", isValid) // Should return true if valid

}

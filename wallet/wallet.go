package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ahmedkhalaf1996/Golang_build_blockchain/utils"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	privateKey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	blockchainAddress string
}

func NewWallet() *Wallet {
	w := new(Wallet)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &w.privateKey.PublicKey
	// 2, Perform Sha-256 hashing on the public key (32 bytes).
	h2 := sha256.New()
	h2.Write(w.publicKey.X.Bytes())
	h2.Write(w.publicKey.Y.Bytes())
	digest2 := h2.Sum(nil)
	// 3, Perform RIPEMD-160 hashing on the result of SHA-256 (20 bytes).
	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)
	// 4, Add version byte in front of RIPEMD-160 hash (0x00 for main Network).
	vd4 := make([]byte, 21)
	vd4[0] = 0x00
	copy(vd4[1:], digest3[:])
	// 5, preform Sha-256 hash on the extended RIPEMD-160 result
	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)
	// 6, Perform SHA-256 hash on the result of the previous Sha-256 hash.
	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)
	// 7 take the first 4 bytes from the second SHA-256 hash for checksum,
	chsum := digest6[:4]
	// 8, Add the 4 checksum bytest from 7 at the end of extended RIPEMD-160 hash from 4 (25 bytes).
	dc8 := make([]byte, 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[21:], chsum[:])
	// 9, Convert the result from a byte string into base58,
	address := base58.Encode(dc8)
	w.blockchainAddress = address
	return w
}

func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

func (w *Wallet) PublickKeyStr() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.privateKey.Y.Bytes())
}

func (w *Wallet) BlockchainAddress() string {
	return w.blockchainAddress
}

type Transaction struct {
	senderPrivateKey           *ecdsa.PrivateKey
	senderPublicKey            *ecdsa.PublicKey
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, sender string, recipient string, value float32) *Transaction {
	return &Transaction{privateKey, publicKey, sender, recipient, value}
}

// func (t *Transaction) GenerateSignature() *utils.Signature {
// 	m, _ := json.Marshal(t)
// 	h := sha256.Sum256([]byte(m))
// 	r, s, _ := ecdsa.Sign(rand.Reader, t.senderPrivateKey, h[:])
// 	return &utils.Signature{r, s}
// }

// func (t *Transaction) GenerateSignature() *utils.Signature {
// 	m, _ := json.Marshal(t)
// 	h := sha256.Sum256(m)
// 	r, s, err := ecdsa.Sign(rand.Reader, t.senderPrivateKey, h[:])
// 	if err != nil {
// 		log.Println("Error generating signature:", err)
// 		return nil
// 	}
// 	return &utils.Signature{R: r, S: s}
// }

func (t *Transaction) GenerateSignature() *utils.Signature {
	m, err := json.Marshal(t)
	if err != nil {
		log.Println("Error marshaling transaction:", err)
		return nil
	}
	h := sha256.Sum256(m)
	r, s, err := ecdsa.Sign(rand.Reader, t.senderPrivateKey, h[:])
	if err != nil {
		log.Println("Error generating signature:", err)
		return nil
	}
	log.Println("Generated Signature:", r, s)
	return &utils.Signature{R: r, S: s}
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_addrress"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}

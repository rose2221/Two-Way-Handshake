package eth

import (
	"context" 
	  "crypto/ecdsa"
"log"
	"math/big"
"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

// make signer that can send ethereum transactions
func NewTransactor(privHex string, chainID int64) *bind.TransactOpts {
	key, err := crypto.HexToECDSA(privHex[2:])
	if err != nil { log.Fatalf("bad pk: %v", err) }

	opts, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(chainID))
	if err != nil { log.Fatalf("tx opts: %v", err) }

	opts.GasTipCap = big.NewInt(1_500_000_000)  
	opts.Context   = context.TODO()              
	return opts
}

func PublicAddress(privHex string) string {
	key, _ := crypto.HexToECDSA(privHex[2:])
	pub := key.Public().(*ecdsa.PublicKey)
	return crypto.PubkeyToAddress(*pub).Hex()
}

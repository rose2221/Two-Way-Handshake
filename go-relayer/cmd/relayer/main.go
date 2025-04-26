package main

import (
	"context"
	"log"
"github.com/ethereum/go-ethereum/common"
	contracts "github.com/rose2221/Two-Way-Handshake/go-relayer/internal/contracts"
	"github.com/rose2221/Two-Way-Handshake/go-relayer/internal/config"
	"github.com/rose2221/Two-Way-Handshake/go-relayer/internal/eth"
"github.com/rose2221/Two-Way-Handshake/go-relayer/internal/relay"
	// "github.com/rose2221/Two-Way-Handshake/go-relayer/internal/persistence"
)

func main() {
	cfg := config.Load()
	ctx := context.Background()

	a_Addr := common.HexToAddress(cfg.Addr_a)
	b_Addr := common.HexToAddress(cfg.Addr_b)

cli_a := eth.Connect(ctx, cfg.SepoliaWs) 
	cli_b := eth.Connect(ctx, cfg.ArbitrumWs) 


	contracta, _ := contracts.NewContractA(a_Addr, cli_a)
	  contractb, _ := contracts.NewContractB(b_Addr, cli_b)

	authA := eth.NewTransactor(cfg.Privhex , cfg.Chain_a)
	authB := eth.NewTransactor(cfg.Privhex , cfg.Chain_b)

	loop := relay.NewLoop(ctx,  contracta, contractb, authA, authB)
	log.Printf("relayer has been initiated,addr=%s", eth.PublicAddress(cfg.Privhex ))
	loop.Start()
}

package eth

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

// Connect opens a websocket connection and auto-reconnects on drop.
func Connect(ctx context.Context, url string) *ethclient.Client {
	c, err := ethclient.DialContext(ctx, url)
	if err != nil { log.Fatalf("eth dial: %v", err) }
	return c
}

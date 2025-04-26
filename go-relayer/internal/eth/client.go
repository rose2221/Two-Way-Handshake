package eth

import (
	"context"
	"log"
"github.com/ethereum/go-ethereum/ethclient"
)

// opens a websocket connection & auto-reconnects on drop
func Connect(ctx context.Context, url string) *ethclient.Client {
	con, err := ethclient.DialContext(ctx, url)
	if err != nil { log.Fatalf("eth dial- %v", err) }
	return con
}

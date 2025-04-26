package config

import (
	"log"
	"os"
	"strconv"
"github.com/joho/godotenv"
)

type Config struct {
	SepoliaWs, ArbitrumWs   string
	Privhex           string
	Addr_a, Addr_b         string
	Chain_a, Chain_b   int64
}

func Load() Config {
	_ = godotenv.Load() 


get := func(k string) string {
		v := os.Getenv(k)
	if v == "" { log.Fatalf("missing env %s", k) }
		return v
	}

	toI := func(k string) int64 {
		i, err := strconv.ParseInt(get(k), 10, 64)
			if err != nil { log.Fatalf("bad int env %s", k) }
			return i
		}
return Config{
		SepoliaWs:  get("SEPOLIA_RPC"),
		ArbitrumWs:  get("RPC_B"),
		Privhex : get("RELAYER_PVT_KEY"),
		Addr_a:     get("CONTRACT_A_ADDR"),
		Addr_b:     get("CONTRACT_B_ADDR"),
		Chain_a:  toI("CHAIN_IDA"),
		Chain_b:  toI("CHAIN_IDB"),
	}
}

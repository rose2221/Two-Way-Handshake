package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	SepoliaWs, ArbitrumWs   string
	PrivKeyHex           string
	AddrA, AddrB         string
	ChainIdA, ChainIdB   int64
}

func Load() Config {
	_ = godotenv.Load() // ignores if no .env

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
		PrivKeyHex: get("RELAYER_PVT_KEY"),
		AddrA:     get("CONTRACT_A_ADDR"),
		AddrB:     get("CONTRACT_B_ADDR"),
		ChainIdA:  toI("CHAIN_IDA"),
		ChainIdB:  toI("CHAIN_IDB"),
	}
}

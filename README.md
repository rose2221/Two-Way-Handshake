
# Cross-chain Handshake 

This project demonstrates a two-way cross-chain handshake between two smart contracts deployed on different networks.  
A relayer service listens for events and sends transactions across chains to complete the handshake.

---

### Deployed contracts

- contract a (sepolia): [`0xddA72c449E9e8Cf5b440859c6AA78d8E1803FAc3`](https://sepolia.etherscan.io/address/0xddA72c449E9e8Cf5b440859c6AA78d8E1803FAc3)
- contract b (arbitrum sepolia): [`0x8c7A0E5A49f6f4ca7AA8EE10933Ab5E7bCE85FeF`](https://sepolia.arbiscan.io/address/0x8c7A0E5A49f6f4ca7AA8EE10933Ab5E7bCE85FeF)

---

### rpc urls used

- sepolia rpc: `https://eth-sepolia.g.alchemy.com/v2/cuJGeZ26qywahhkli6epc2XqIOoZJhFF`
- arbitrum sepolia rpc: `https://arb-sepolia.g.alchemy.com/v2/AJAOv3C3nutHF-2duNghabWWPgzD1hxw`

---

### Handshake initiation transaction

- initiated from contract a with session id `785`
- transaction hash: [`0x8c3025777753e38b5073e4e69097c32639588d4b91965e1c4a87856bba8d5cbd`](https://sepolia.etherscan.io/tx/0x8c3025777753e38b5073e4e69097c32639588d4b91965e1c4a87856bba8d5cbd)

---

### video demo

- [click here to watch the successful handshake demo](link-to-your-video)

---

### How to run the relayer

1. fill the `.env` file presnt in `./go-relayer` directory with the following variables:

```env
SEPOLIA_RPC=https://eth-sepolia.g.alchemy.com/v2/cuJGeZ26qywahhkli6epc2XqIOoZJhFF
RPC_B=https://arb-sepolia.g.alchemy.com/v2/AJAOv3C3nutHF-2duNghabWWPgzD1hxw
RELAYER_PVT_KEY=your_private_key
CONTRACT_A_ADDR=0xddA72c449E9e8Cf5b440859c6AA78d8E1803FAc3
CONTRACT_B_ADDR=0x8c7A0E5A49f6f4ca7AA8EE10933Ab5E7bCE85FeF
CHAIN_IDA=11155111    # sepolia chain id
CHAIN_IDB=421614      # arbitrum sepolia chain id
```

2. build and run the relayer:

```bash
go build ./cmd/relayer
./relayer
```



### how to initiate handshake manually

any user with a funded sepolia wallet can initiate a handshake by calling `initiateHandshake` on contract a.


```bash
export CONTRACT_A_ADDR=0xddA72c449E9e8Cf5b440859c6AA78d8E1803FAc3
export SEPOLIA_RPC=https://eth-sepolia.g.alchemy.com/v2/cuJGeZ26qywahhkli6epc2XqIOoZJhFF
export YOUR_PRIVATE_KEY=your_private_key_here
export SESSION=

cast send $CONTRACT_A_ADDR "initiateHandshake(uint256)" $SESSION \
    --rpc-url $SEPOLIA_RPC \
    --private-key $YOUR_PRIVATE_KEY
```

this will emit a `syn` event, and the relayer will automatically complete the handshake across the two networks.

---
Of course! Here's your content cleaned up, **humanized**, and **formatted nicely** for the README (without extending it, just polishing it):

---

## Future Enhancements

### 1. Persistence Layer (Replay Protection)
Currently, if the relayer restarts, it forgets which sessions it has already handled, which could cause duplicate transactions.

**Solution:**  
Integrate a simple LevelDB instance to store:
- Last processed block numbers
- Session IDs that have already been relayed

---

### 2. Dockerization & CI Pipeline
Right now, the relayer requires manual environment setup.

**Solution:**  
- Add a `Dockerfile` and `docker-compose.yaml` to simplify deployment.
- Add tests to validate the entire relayer flow during CI builds.

---

### 3. Transaction Retry with Exponential Backoff
If a transaction fails due to issues like RPC errors, nonce gaps, or mempool reorgs, the relayer currently just logs the error and moves on.

**Solution:**  
Use a retry mechanism such as `go-retry` or implement a simple exponential backoff system to retry failed transactions automatically.

---

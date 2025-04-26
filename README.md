```markdown
# cross-chain handshake demo

this project demonstrates a two-way cross-chain handshake between two smart contracts deployed on different networks.  
a relayer service listens for events and sends transactions across chains to complete the handshake.

---

### deployed contracts

- contract a (sepolia): [`0xddA72c449E9e8Cf5b440859c6AA78d8E1803FAc3`](https://sepolia.etherscan.io/address/0xddA72c449E9e8Cf5b440859c6AA78d8E1803FAc3)
- contract b (arbitrum sepolia): [`0x8c7A0E5A49f6f4ca7AA8EE10933Ab5E7bCE85FeF`](https://sepolia.arbiscan.io/address/0x8c7A0E5A49f6f4ca7AA8EE10933Ab5E7bCE85FeF)

---

### rpc urls used

- sepolia rpc: `https://eth-sepolia.g.alchemy.com/v2/cuJGeZ26qywahhkli6epc2XqIOoZJhFF`
- arbitrum sepolia rpc: `https://arb-sepolia.g.alchemy.com/v2/AJAOv3C3nutHF-2duNghabWWPgzD1hxw`

---

### handshake initiation transaction

- initiated from contract a with session id `785`
- transaction hash: [`0x8c3025777753e38b5073e4e69097c32639588d4b91965e1c4a87856bba8d5cbd`](https://sepolia.etherscan.io/tx/0x8c3025777753e38b5073e4e69097c32639588d4b91965e1c4a87856bba8d5cbd)

---

### video demo

- [click here to watch the successful handshake demo](link-to-your-video)

---

### how to run the relayer

1. fill the `.env` file with the following variables:

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

---

### how to initiate handshake manually

any user with a funded sepolia wallet can initiate a handshake by calling `initiateHandshake` on contract a.


```bash
export CONTRACT_A_ADDR=0xddA72c449E9e8Cf5b440859c6AA78d8E1803FAc3
export SEPOLIA_RPC=https://eth-sepolia.g.alchemy.com/v2/cuJGeZ26qywahhkli6epc2XqIOoZJhFF
export YOUR_PRIVATE_KEY=your_private_key_here
export SESSION=785

cast send $CONTRACT_A_ADDR "initiateHandshake(uint256)" $SESSION \
    --rpc-url $SEPOLIA_RPC \
    --private-key $YOUR_PRIVATE_KEY
```

this will emit a `syn` event, and the relayer will automatically complete the handshake across the two networks.

---
```




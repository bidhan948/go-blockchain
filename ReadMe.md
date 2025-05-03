# 🚀 Go-Blockchain: A Hands-On Blockchain in Go

![Go](https://img.shields.io/badge/language-Go-blue) ![License](https://img.shields.io/github/license/bidhan948/go-blockchain)

Welcome to **Go-Blockchain**, a comprehensive, end-to-end blockchain implementation written in Go and maintained in this repository: [https://github.com/bidhan948/go-blockchain](https://github.com/bidhan948/go-blockchain). This project is designed as a learning playground for developers who want to:

* Understand and build each piece of a blockchain from scratch
* Dive deep into Proof-of-Work mining and nonce discovery
* Construct and verify Merkle trees for transaction integrity
* Model UTXO-based transactions and manage unspent outputs
* Generate and manage digital wallets with ECDSA keys and Base58 addresses
* Broadcast blocks and transactions over a simple P2P network
* Interact with your own blockchain via a user-friendly CLI

---

## 🔎 Detailed Features

### 1. Core Blockchain Logic (`blockchain` package)

* **`Block` & `Blockchain`** structures: store data, link via hashes.
* **Proof-of-Work**: find a valid nonce to satisfy difficulty targets.
* **Merkle Trees**: generate and traverse trees for transaction sets.
* **UTXO Set**: track and handle unspent transaction outputs.
* **Transaction Builder**: create, sign, and verify transactions.

### 2. CLI Interface (`cli` package)

* **`createwallet`**: generate new wallet and addresses.
* **`listaddresses`** & **`getbalance`**: inspect your wallet.
* **`send`**: craft and broadcast transactions.
* **`printchain`**: view the full chain in human-readable form.
* **`startnode`**: launch a node that can mine and sync with peers.

### 3. Wallet Management (`wallet` package)

* **ECDSA Key Pairs**: secure private keys and derive public keys.
* **Base58Check Encoding**: produce Bitcoin-style addresses.
* **Address Validation**: check for tampering or invalid formats.

### 4. Networking (`network` package)

* **P2P Communication**: simple in-memory node discovery via `NODE_ID`.
* **Message Broadcasting**: relay new blocks and transactions to peers.
* **Sync Mechanism**: ensure each node stays in consensus.

---

## ⚙️ Getting Started

### Prerequisites

* Go 1.18 or later installed
* A configured `GOPATH`

### Quick Installation

```bash
# Clone the repo into your GOPATH
git clone https://github.com/bidhan948/go-blockchain $GOPATH/src/github.com/bidhan948/go-blockchain
cd $GOPATH/src/github.com/bidhan948/go-blockchain

# Install dependencies
go mod tidy
```

### Running the CLI

1. **Start a node & mine a genesis block**

   ```bash
   export NODE_ID=3000
   go run main.go startnode -miner YOUR_ADDRESS
   ```

2. **Print the blockchain**

   ```bash
   go run main.go printchain
   ```

3. **Create and list wallets**

   ```bash
   go run main.go createwallet
   go run main.go listaddresses
   ```

4. **Check balances**

   ```bash
   go run main.go getbalance -address YOUR_ADDRESS
   ```

5. **Send coins**

   ```bash
   go run main.go send -from SENDER_ADDRESS -to RECEIVER_ADDRESS -amount 10
   ```

---

## 📁 Repository Structure

```
go-blockchain/         # Root of the project
├── blockchain/        # Core block, pow, and merkle logic
├── cli/               # Command-line parsing and commands
├── wallet/            # Key generation and address utilities
├── network/           # P2P networking and message handling
├── main.go            # Entry point for CLI
└── go.mod             # Module definition
```

---

## 🛠️ Customization & Extensions

* **Consensus Algorithms**: swap Proof-of-Work for Proof-of-Stake or others
* **Persistent Storage**: integrate LevelDB, BoltDB, or Redis for block storage
* **Advanced Networking**: include NAT traversal, gossip protocols, and discovery
* **Smart Contracts**: experiment with a simple VM or scripting language

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Please open a pull request or issue on GitHub: [https://github.com/bidhan948/go-blockchain](https://github.com/bidhan948/go-blockchain)

---

## 📄 License

Distributed under the MIT License. See [LICENSE](LICENSE) for more information.

---

*Happy coding and may your blocks always validate!*

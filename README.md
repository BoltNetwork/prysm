<h1 align="center">⚡ Bolt Network: Beacon Chain Client</h1>

<p align="center">
  <strong>Fast Ethereum testnet with 5-second blocks — test faster, ship faster.</strong>
</p>

<div align="center">

[![Go Report Card](https://goreportcard.com/badge/github.com/BoltNetwork/prysm)](https://goreportcard.com/report/github.com/BoltNetwork/prysm)
[![Consensus_Spec_Version 1.4.0](https://img.shields.io/badge/Consensus%20Spec%20Version-v1.4.0-blue.svg)](https://github.com/ethereum/consensus-specs/tree/v1.4.0)
[![Block Time](https://img.shields.io/badge/Block%20Time-5%20seconds-brightgreen.svg)](https://boltnode.org)
[![Website](https://img.shields.io/badge/Website-boltnode.org-blue.svg)](https://boltnode.org)

</div>

---

## 📖 Overview

**Bolt Network** is a fast Ethereum testnet featuring **5-second block times** and Flashbots integration. Built for developers who need rapid iteration without waiting for mainnet's 12-second blocks.

This repository contains the Beacon Chain client for Bolt Network, forked from [Prysm](https://github.com/prysmaticlabs/prysm) and optimized for the Bolt testnet environment.

### 🎯 Use Cases

- **Smart Contract Testing** — Iterate 2.4x faster with 5-second blocks
- **MEV Research** — Experiment with Flashbots in a sandboxed environment
- **Validator Experimentation** — Test validator setups without mainnet risk
- **CI/CD Pipelines** — Faster test suites for blockchain applications
- **Load Testing dApps** — Stress test your applications at accelerated speeds

---

## 🌐 Network Information

| Property | Value |
|----------|-------|
| **Network Name** | Bolt Network Testnet |
| **Chain ID** | `1337` |
| **Block Time** | `5 seconds` |
| **RPC Endpoint** | `https://rpc.boltnode.org` |
| **Block Explorer** | [explorer.boltnode.org](https://explorer.boltnode.org) |
| **Validator Launchpad** | [launchpad.boltnode.org](https://launchpad.boltnode.org) |

---

## 🚀 Quick Start

### Prerequisites

- [RETH](https://github.com/paradigmxyz/reth) (Execution Client)
- Go 1.21+ (for building from source)

### Run Beacon Node

```bash
# Using the launcher script
./prysm.sh beacon-chain --bolt-testnet \
  --execution-endpoint=http://localhost:8551 \
  --jwt-secret=/path/to/jwt.hex

# Or build and run directly
go build -o beacon-chain ./cmd/beacon-chain
./beacon-chain --bolt-testnet \
  --execution-endpoint=http://localhost:8551 \
  --jwt-secret=/path/to/jwt.hex
```

### Run Validator

```bash
./prysm.sh validator --bolt-testnet \
  --wallet-dir=/path/to/wallet \
  --wallet-password-file=/path/to/password.txt
```

### Run with RETH (Execution Client)

```bash
# Start RETH first
reth node \
  --chain bolt \
  --http \
  --http.api eth,net,web3 \
  --authrpc.jwtsecret=/path/to/jwt.hex

# Then start the beacon chain
./prysm.sh beacon-chain --bolt-testnet \
  --execution-endpoint=http://localhost:8551 \
  --jwt-secret=/path/to/jwt.hex
```

---

## 🏆 Become a Validator

Ready to validate on Bolt Network? Visit the **[Bolt Validator Launchpad](https://launchpad.boltnode.org)** to get started.

🔍 Track your validator performance at **[explorer.boltnode.org](https://explorer.boltnode.org)**

---

## 🤝 Contributing

### 🔥 Branches

- **[`master`](https://github.com/BoltNetwork/prysm/tree/master)** — Latest stable release
- **[`develop`](https://github.com/BoltNetwork/prysm/tree/develop)** — Active development, base PRs here

### 🛠 How to Contribute

1. Fork this repository
2. Create a feature branch from `develop`
3. Submit a Pull Request

---

## 🔗 Resources

- 🌐 **Website:** [boltnode.org](https://boltnode.org)
- 🔎 **Explorer:** [explorer.boltnode.org](https://explorer.boltnode.org)
- 🚀 **Launchpad:** [launchpad.boltnode.org](https://launchpad.boltnode.org)
- 💻 **GitHub:** [github.com/BoltNetwork](https://github.com/BoltNetwork)

---

## 📜 License

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0.en.html)

This project is licensed under the **GNU General Public License v3.0**.

Based on [Prysm](https://github.com/prysmaticlabs/prysm) by Prysmatic Labs.

---

## ⚖️ Legal Disclaimer

📜 [Terms of Use](/TERMS_OF_SERVICE.md)

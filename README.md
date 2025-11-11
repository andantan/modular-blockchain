# Kangaroo (Migration)

This repository documents the migrated version of the original project now hosted at [https://github.com/andantan/kangaroo](https://github.com/andantan/kangaroo). It summarizes the content presented at **GopherCon Korea 2025** and provides essential commands for setup and execution.

---

## Overview

* `make run ARGS="-k c.hex"` : Runs a blockchain node. The `-k` flag specifies the path to the private key file.
* `make run-discovery` : Runs the **Peer DNS (Discovery)** server — **must be running** before any node.
* `make key_gen ARGS="-o test.hex"` : Generates a new private key.
* `make key_info ARGS="-k test.hex"` : Prints information about the given private key.
* `validators.json` : Located in the root directory, this file defines the registered validator addresses.
* Send random transactions via `GET http://127.0.0.1:51773/random` — returns a transaction hash in response.

---

## Quickstart

1. **Clone and install dependencies**

```bash
git clone https://github.com/andantan/modular-blockchain.git
cd modular-blockchain
# Build and install dependencies (e.g., go build, etc.)
```

2. **Start the Discovery Server (Required)**

```bash
make run-discovery
```

3. **Generate a Key**

```bash
make key_gen ARGS="-o test.hex"
# View key information
make key_info ARGS="-k test.hex"
```

4. **Run a Node**

```bash
make run ARGS="-k test.hex"
```

> The `-k` option specifies the path to the `.hex` private key file.

---

## Command Reference

### `make run`

* Runs a blockchain node.
* Example: `make run ARGS="-k path/to/privkey.hex"`

### `make run-discovery`

* Runs the peer discovery (DNS) server.
* Must be active to allow node connection and network discovery.

### `make key_gen`

* Generates a new private key in `.hex` format.
* Example: `make key_gen ARGS="-o mykey.hex"`

### `make key_info`

* Prints the key’s information (address, public key, etc.).
* Example: `make key_info ARGS="-k mykey.hex"`

---

## validators.json

Located in the project root, `validators.json` defines the validator set in the following format:

```json
[
  "0xd6ac9a9828d40df537a2e137401ec9e844146032",
  "0x8fbb14f9d1e5584093272ba76114fe8dc454667f",
  "0xb03fc15085e05046286f12b0d19bb63bab6da40c",
  "0x80a8b001136e08864f459e9218a2df845e187543"
]
```

* Used to register or initialize validator nodes at startup.
* In production, this file should be managed through an automated deployment process.

---

## Random Transaction API (for Testing)

You can generate a random transaction using:

```http
GET http://127.0.0.1:{api-server-port}/random
```

**Example Response:**

```json
{
  "transaction_hash": "0xf82118e6d3a2e9e2a7df8e7d634e1f026b18589f367c8dbd404c07020f66726a"
}
```

* This endpoint creates a random transaction inside the node and submits it to the mempool/network.
* Useful for load testing or automated QA.

---

## GopherCon Korea 2025



This project was presented at **GopherCon Korea 2025**, and the demo and materials are being integrated into the migrated repository:

👉 [https://github.com/andantan/kangaroo](https://github.com/andantan/kangaroo)

* **Speaker:** 전규빈(Qbean)
* **Session:** Go로 밑바닥부터 맨 땅에 헤딩하듯 만드는 P2P 블록체인 네트워크
* **Conference URL:** [https://gophercon.kr/2025](https://gophercon.kr/2025)
* **Presentation Slides:** [https://github.com/andantan/GopherCon-Korea-2025-Scripts](https://github.com/andantan/GopherCon-Korea-2025-Scripts)
* **Presentation materials:** See `docs/` or the presentation section in the migrated repo.

---

## Developer Notes & Operational Tips

* The Discovery server must be running for proper peer discovery.
* After modifying `validators.json`, you may need to restart the node depending on reload support.
* Keep private key files secure — **never commit them to version control**.
* The random transaction endpoint is for **testing only** and should be access-restricted in production.

---

## Example Execution Flow

1. `make run-discovery` — start the discovery server
2. `make key_gen ARGS="-o test.hex"` — generate a private key
3. `make key_info ARGS="-k test.hex"` — view key info
4. `make run ARGS="-k test.hex"` — run the node
5. `curl http://127.0.0.1:{api-server-port}/random` — generate a random transaction

---

## Migration & Contribution

* The code is being migrated to **andantan/kangaroo**. For PRs, issues, or demos, please refer to the migrated repository.
* This repository will eventually be fully integrated based on the migration plan.

---

## Contact

Maintainer / Speaker: kyubin2892@gmail.com

---

*If you’d like to add more details — such as screenshots, CI/CD setup, or environment variable descriptions — let me know, and I’ll include them here.*

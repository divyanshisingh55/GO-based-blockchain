
# Go-Based Blockchain

A lightweight **blockchain system built in Go**, designed for storing file metadata securely with SHA-256 hashing and AES-256 encryption. It uses the **Gin framework** for API handling and serves a simple frontend for interaction.

---

## Features

*  **Blockchain Structure** – Stores blocks with file data and metadata.
*  **SHA-256 Hashing** – Ensures file integrity and immutability.
*  **Genesis Block** – Automatically created as the first block.
*  **File Upload System** – Files are uploaded and stored locally with hashed references.
*  **Gin Web API** – Fast and lightweight API built with Go’s Gin framework.
*  **Custom Logger** – Logs all operations clearly in the terminal.
*  **AES-256 Encryption** – Blockchain data (blockchain.json) is securely encrypted.
*  **Frontend Pages** – Includes issuer, holder, and verifier HTML pages.

---

## Setup & Run

### 1. Clone the repo

```bash
git clone https://github.com/<your-username>/go-blockchain.git
cd go-blockchain
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Run the server

```bash
go run main.go
```

### 4. Access in browser

```
http://localhost:8080
```

---

##  API Endpoints

| Method | Endpoint              | Description            |
| ------ | --------------------- | ---------------------- |
| `GET`  | `/`                   | Load homepage          |
| `POST` | `/upload`             | Upload and hash a file |
| `GET`  | `/chain`              | View blockchain data   |
| `GET`  | `/download/:filename` | Download a stored file |

---

## Next Goals

* Integrate **Hyperledger Fabric** for permissioned blockchain.
* Add **smart contracts** for access control.
* Include **AI-based verification** for file authenticity.
* Host system with **multi-node simulation**.

---

## Tech Stack

* **Language:** Go (Golang)
* **Framework:** Gin
* **Encryption:** AES-256
* **Hashing:** SHA-256
* **Frontend:** HTML, CSS, JS

---


---

want me to make it formatted for **GitHub (markdown style with emojis and colors)** or keep it simple like docs version?

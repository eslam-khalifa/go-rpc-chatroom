# go-rpc-chatroom

A simple **half-duplex (polling) chatroom** implemented in **Go** using the standard `net/rpc` package.

---

## 🎥 Demo Video

Watch the project in action here:

👉 [**Demo Video**](https://drive.google.com/file/d/12fTzXtOM8FsNaM8HUVNMR-0N_VFDNmZO/view?usp=sharing)

---

## 🚀 Features

- Separate `server.go` and `client.go`.
- Shared `commons/shared.go` package for message types and server address.
- Clients provide **username** (e.g., `Eslam Khalifa`), and messages are stored as `Name: Message`.
- Server prints messages **immediately** on receipt.
- Clients can view the **entire chat history** after each message (polling mode).

---

## 🗂️ Project Structure

```txt

go-rpc-chatroom/
├── client.go
├── server.go
├── go.mod
├── .gitignore
├── test.sh
├── LICENSE
├── commons/
│   └── shared.go
└── README.md

````

---

## ⚙️ Requirements

- Go 1.20+ (or later)
- Working internet connection (for GitHub cloning, not required for running)
- Initialized Go module (see below)

---

## 🧱 Module Setup

If this is your first time running the project, initialize a Go module in your project root:

```bash
go mod init go-rpc-chatroom
go mod tidy
````

If you plan to publish it on GitHub, use your repo path instead:

```bash
go mod init github.com/<your-username>/go-rpc-chatroom
go mod tidy
```

---

## 🏃‍♂️ How to Run

### 1️⃣ Start the Server

```bash
go run server.go
```

Expected output:

```txt
Chat server running on 0.0.0.0:9999
```

---

### 2️⃣ Start the Client (in a new terminal)

```bash
go run client.go
```

Then follow the prompts:

```txt
Connected to chat server at 0.0.0.0:9999
Enter your name: Eslam Khalifa
Welcome, Eslam Khalifa! Type your messages below (type 'exit' to quit)

You: Hello everyone!

--- Chat History ---
Eslam Khalifa: Hello everyone!
--------------------
```

---

## 🧩 Example Interaction

**Client 1 (Eslam Khalifa):**

```txt
You: Hello everyone!
```

**Client 2 (Omar):**

```txt
You: Hey Eslam Khalifa!
```

**Server terminal output:**

```txt
💬 Eslam Khalifa: Hello everyone!
💬 Omar: Hey Eslam Khalifa!
```

---

## 📦 Files Explained

| File                | Description                                                           |
| ------------------- | --------------------------------------------------------------------- |
| `server.go`         | RPC server that coordinates messages and broadcasts chat history      |
| `client.go`         | CLI client that sends messages and fetches history                    |
| `commons/shared.go` | Shared structs (`Message`, `Reply`) and `GetServerAddress()` function |

---

## 🧾 License

MIT License — free to use and modify.
See the [LICENSE](LICENSE) file for details.

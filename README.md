# gshell

![License](https://img.shields.io/badge/license-MIT-blue.svg)

**gshell — Secure peer-to-peer terminal sharing.**

`gshell` lets you share a live terminal session with someone using a single command.
No servers, no accounts, no infrastructure.

It generates a **secure encrypted invite token** that hides the host IP and port. Anyone with the token (and password) can join the session.

---

# Features

* **Peer-to-peer terminal sharing**
* **Encrypted invite tokens**
* **Password-protected sessions**
* **Session expiry**
* **Secure handshake authentication**
* **Single static Go binary**
* **Works directly from the terminal**
* **No backend or central server required**

---

# How It Works

When a user starts sharing:

1. `gshell` creates a local terminal session (PTY).
2. A **session payload** is created containing:

   * Host IP
   * Port
   * Session name
   * Expiry time
3. The payload is **encrypted using a secret embedded in the binary**.
4. An invite token is generated.

Example flow:

Host:

```
gshell share --password dev123
```

Output:

```
Session: default-session
Expires in: 30 minutes

Join using:
gshell join eyJ3Y...
```

Client:

```
gshell join eyJ3Y... --password dev123
```

The client decrypts the token and connects automatically.

The client **never sees the IP address**.

---

# Installation

### Build locally

```bash
git clone https://github.com/pawannn/gshell
cd gshell
go build -o gshell
```

Run:

```bash
./gshell
```

---

# Usage

### Share a terminal session

```bash
gshell share --password dev123
```

Options:

```
--port       Port to listen on (default: 9000)
--password   Session password
```

Example output:

```
Session: default-session
Expires in: 30 minutes

gshell join eyJhbGc...
```

---

### Join a session

```bash
gshell join <token> --password dev123
```

Example:

```bash
gshell join eyJhbGc... --password dev123
```

Once connected, both users share the same terminal session.

---

# Security Model

`gshell` uses several layers of security:

### Encrypted session tokens

Session information is encrypted before being shared.

Payload example:

```json
{
  "ip": "192.168.1.10",
  "port": "9000",
  "session": "dev-shell",
  "exp": 1712000000,
  "pwd": "dev123"
}
```

The encrypted token is what gets shared.

---

### Build-time secret key

A random secret key is embedded in every build using Go linker flags.

Example:

```
go build -ldflags "-X github.com/pawannn/gshell/internal/security.SecretKey=RANDOM_SECRET"
```

In CI, the key is generated automatically.

---

### Password handshake

When a client connects:

```
client → send password
server → verify password
server → allow shell session
```

Unauthorized users cannot start the shell.

---

### Session expiry

Invite tokens automatically expire after the configured time.

If no one joins before expiry, the listener shuts down automatically.

---

# CI/CD

The project uses **GitHub Actions**.

On every push:

1. A **random secret key is generated**
2. The binary is built with that key embedded
3. The compiled binary is uploaded as an artifact

This ensures every build uses a **different encryption key**.

---

# Example Session

Host:

```
gshell share --password secret
```

Client:

```
gshell join eyJhbGciOi... --password secret
```

Both terminals now control the same shell.

---

# Limitations

* If two users type simultaneously, characters may overlap (same as shared tmux sessions).
* The secret key is embedded in the binary and can theoretically be extracted with advanced reverse engineering.
* Designed for **ephemeral P2P sessions**, not permanent infrastructure.

---

# Roadmap

Planned improvements:

* multi-client sessions
* shorter invite tokens
* session recording
* terminal resize sync
* cross-platform binary releases
* TLS support

---

# License

MIT License

---

# Contributing

Contributions, suggestions, and improvements are welcome.

If you find a bug or have an idea:

```
open an issue
submit a pull request
```

---

# Author

Created by **Pawan Kalyan**

GitHub:
https://github.com/pawannn

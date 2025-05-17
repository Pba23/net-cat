# 🌐 net-cat

> **Note:** This project is one of my first coding projects written in Go.  
> As such, some best practices and code conventions may not be fully respected yet.  
> This is a learning journey and I'm improving step by step!

<div align="center">
  
![TCP Chat](https://img.shields.io/badge/TCP-Chat-blue)
![Go](https://img.shields.io/badge/Go-1.16+-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green)

</div>

## 📋 Project Overview

**net-cat** is a reimplementation of the classic NetCat (`nc`) tool in Go, based on a Server-Client architecture. It runs either as a server listening on a specified port for incoming TCP connections, or as a client connecting to a server and exchanging messages.

NetCat is a command-line utility used to read and write data across network connections using TCP or UDP protocols. It's often used for network debugging, testing, or creating simple chat systems.

This project simulates a multi-client group chat with features inspired by the original NetCat tool.

## ✨ Features

- 🔄 TCP server that supports multiple client connections (1-to-many)
- 👤 Client name registration on connection
- 🔟 Maximum of 10 simultaneous connections
- 💬 Clients can send messages to the chat; empty messages are ignored
- 🕒 Each message is timestamped and identified with the sender's username, e.g.:  
  `[2020-01-20 15:48:41][client.name]:[client.message]`
- 📜 When a new client joins, all previous chat messages are sent to them
- 🔔 Other clients are notified when a client joins or leaves the chat
- 🛡️ The server handles client disconnections gracefully without interrupting other clients
- 🔢 Default port is 8989 if none is specified
- ❓ Usage message shown if incorrect arguments are provided

## 🚀 Usage

### Running the server

```bash
go run . [port]
```

- If no port is provided, it defaults to 8989
- Example: `go run . 2525`

### Connecting clients

```bash
nc <server_ip> <port>
```

Upon connection, clients see a welcome ASCII art and are prompted to enter their name.

### Example client session

```
Welcome to TCP-Chat!
         *nnnn*
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
*)      \.*__.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'

[ENTER YOUR NAME]:
```

## 🌟 Bonus Features (Optional)

- 🔄 Allow clients to change their usernames
- 📢 Notify all clients when a user changes their name
- 👥 Support multiple chat groups
- 🚩 Implement additional NetCat command flags
- 💾 Save all chat logs to a file

## 📚 What I Learned

- 🧩 How to manipulate Go structures
- 🌐 Working with TCP and UDP connections and sockets
- ⚡ Go concurrency primitives: goroutines and channels
- 🔒 Synchronization using mutexes
- 🔌 Managing IP addresses and ports in Go networking

## 📦 Allowed Packages

This project uses only the following standard Go packages:

- `io` - Basic input/output operations
- `log` - Logging events and errors
- `os` - OS-dependent functionality
- `fmt` - Formatted input/output
- `net` - Basic network interfaces
- `sync` - Synchronization primitives
- `time` - Measuring and displaying time
- `bufio` - Buffered I/O
- `errors` - Error handling
- `strings` - String manipulation
- `reflect` - Runtime type inspection

## 🔧 Installation

```bash
# Clone the repository
git clone https://github.com/your-username/net-cat.git
cd net-cat

# Build the project
go build

# Run the server
./net-cat [port]
```

## 🤝 Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## 📝 License

This project is licensed under the MIT License. See the LICENSE file for details.

package main

import (
	"fmt"
	"log"
	"net"
	"netcat/utils"
	"os"
	"sync"
)

type client struct {
	conn     net.Conn
	username string
}

var clients = make(map[net.Conn]client)

var mutex sync.Mutex

func main() {
	port := "8989"

	if len(os.Args) == 2 {
		port = os.Args[1]
	}

	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	fmt.Println("Serveur en écoute sur le port " + port + "...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		mutex.Lock()
		if len(clients) > 10 {
			conn.Close()
			mutex.Unlock()
			continue
		}
		mutex.Unlock()

		go utils.HandleClient(conn)
	}
}

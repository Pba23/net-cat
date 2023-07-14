package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type client struct {
	conn     net.Conn
	username string
}

var clients = make(map[net.Conn]client)
var tabchat []string
var tabUserName []string
var mutex sync.Mutex

func HandleClient(conn net.Conn) {
	defer conn.Close()

	data, _ := ioutil.ReadFile("linux.txt")
	linux := string(data)

	conn.Write([]byte("Welcome to TCP-Chat!\n"))
	conn.Write([]byte(linux))

restart:

	conn.Write([]byte("\n[ENTER YOUR NAME]: "))
	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	username := scanner.Text()

	if len(username) < 2 || len(username) > 12 {
		conn.Write([]byte("Username not allowed !!!, Retry...\n"))
		goto restart
	}

	if IsPresent(tabUserName, username) {
		conn.Write([]byte("This username already exists !!!, Retry...\n"))
		goto restart
	}
	tabUserName = append(tabUserName, username)
	c := client{
		conn:     conn,
		username: username,
	}

	fmt.Println(tabchat)
	for _, chat := range tabchat {
		conn.Write([]byte(chat + "\n"))
	}

	mutex.Lock()
	clients[conn] = c
	mutex.Unlock()

	broadcastMessage(fmt.Sprintf("\r\033[K%s has joined our chat...✔", username), username)
	conn.Write([]byte("[" + time.Now().Format("2006-01-02 15:04:05") + "][" + username + "]: "))

	for scanner.Scan() {
		msg := scanner.Text()

		if len(msg) == 0 {
			conn.Write([]byte("⛔ No empty Message \n"))
			conn.Write([]byte("[" + time.Now().Format("2006-01-02 15:04:05") + "][" + username + "]: "))
			continue
		}

		emojiMess := "\U0001F4AC"
		broadcastMessage(fmt.Sprintf("\r\033[K[%s][%s] %s : %s", time.Now().Format("2006-01-02 15:04:05"), username, emojiMess, msg), username)
		conn.Write([]byte("[" + time.Now().Format("2006-01-02 15:04:05") + "][" + username + "]: "))

		chat := fmt.Sprintf("\r\033[K[%s][%s] %s : %s", time.Now().Format("2006-01-02 15:04:05"), username, emojiMess, msg)
		tabchat = append(tabchat, chat)

		log.Println(chat)

		file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err)
		}
		log.SetOutput(file)
	}

	mutex.Lock()
	delete(clients, conn)
	mutex.Unlock()
	broadcastMessage(fmt.Sprintf("\r\033[K%s has left our chat...🏃", username), username)
	tabUserName = removeIndex(tabUserName, index(tabUserName, username))
	conn.Close()
	// mutex.Unlock()
	// return
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
func index(s []string, str string) int {
	for i, val := range s {
		if strings.ToLower(val) == strings.ToLower(str) {
			return i
		}
	}
	return -1
}

// func Red(message string) string {
// 	return fmt.Sprintf("\x1b[31m%s\x1b[0m", message)
// }

// // Green returns a green string
// func Green(message string) string {
// 	return fmt.Sprintf("\x1b[32m%s\x1b[0m", message)
// }

// func Orange(message string) string {
// 	return fmt.Sprintf("\x1b[91m%s\x1b[0m", message)
// }

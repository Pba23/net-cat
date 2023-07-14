package utils

import (
	"log"
	"time"
)

func broadcastMessage(msg string, client string) {
	mutex.Lock()
	defer mutex.Unlock()
	for _, c := range clients {
		if c.username != client {
			_, err := c.conn.Write([]byte("📧" + msg + "\n"))
			if err != nil {
				log.Println(err)
			}
			c.conn.Write([]byte("[" + time.Now().Format("2006-01-02 15:04:05") + "][" + c.username + "]: "))

		}
	}
}

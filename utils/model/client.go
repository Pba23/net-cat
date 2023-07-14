package model

import "net"

type Client struct {
	conn     net.Conn
	username string
}

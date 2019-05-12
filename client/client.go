package client

import (
	"net"
)

type Client struct {
	Conn net.Conn
	Username string
}

func (c *Client) SetUsername(username string){
	c.Username = username
}

func (c *Client) JoinChannel(){

}

func (c *Client) SendMessage(){
}


func (c *Client) username() string{
	return c.Username
}





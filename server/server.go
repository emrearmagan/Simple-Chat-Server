package server

import (
	"Go/ChatServer/client"
	"bufio"
	"fmt"
	"net"
)

type TCPSever struct {
	listener net.Listener
	clients  []*client.Client
}

func (s *TCPSever) serve(message string, client *client.Client) {
	for _, v := range s.clients {
		//Client-Side
		fmt.Fprintf(v.Conn, "<%s> %s\n", client.Username, message)
		//Server-Side
		fmt.Printf("<%s> %s\n", client.Username,message)
	}
}

func (s *TCPSever) handleConn(client *client.Client) {
	defer s.removeClient(client)

	scanner := bufio.NewScanner(client.Conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if len(ln) != 0 {
			s.serve(ln, client)
		}
	}
}

func (s *TCPSever) removeClient(client *client.Client) {

	for i, v := range s.clients {
		if v == client {
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
			break
		}
	}

	s.serve("left the channel", client)

	client.Conn.Close()
}

func (s *TCPSever) getClient(conn net.Conn) *client.Client {
	var input string

	//clear terminal
	fmt.Fprint(conn,"\033[2J")

	scanner := bufio.NewScanner(conn)
	fmt.Fprint(conn, "Enter name here: ")

	for scanner.Scan() {
		input = scanner.Text()
		if len(input) != 0 {
			break
		}
		fmt.Fprint(conn, "invalid name\n")
		fmt.Fprint(conn, "Enter name here: ")
	}

	client := &client.Client{
		Conn:     conn,
		Username: input,
	}

	s.clients = append(s.clients, client)
	s.serve("joined the channel", client)

	return client
}

func (s *TCPSever) Start(port string) {
	li, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	} else {
		s.listener = li
	}

	defer li.Close()
	fmt.Printf("Starting Chat-Server on Port%s...\n", port)

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println(err)
		} else {
			client := s.getClient(conn)
			go s.handleConn(client)
		}
	}
}

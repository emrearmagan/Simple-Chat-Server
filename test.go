/*
 start with go run main.go
 in second terminal "nc localhost 8080"
 */

package main

import "Go/ChatServer/server"

func main() {
	var s server.TCPSever
	s.Start(":8080")

}

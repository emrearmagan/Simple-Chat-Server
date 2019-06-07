/*
 start with go run main.go
 in second terminal "nc localhost 8080"
 */

package main

import "Chat-Server/server"

func main() {
	var s server.TCPSever
	s.Start(":8888")
}

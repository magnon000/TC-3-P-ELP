package main

import (
	"fmt"
	"net"
)

// save no file on server

const (
	ip_addr  = "127.0.0.1"
	port     = "50000"
	protocol = "tcp"
)

// handle request, type: net.Conn
func process(conn net.Conn) {
	// process here
	// close connection
	defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("Connection reading failed:\n	%v\n", err)
			break
		}
		fmt.Printf("Content from client:%v\n", string(buf[:n]))
	}

}

func main() {
	// 1. start listening
	listener, err := net.Listen(protocol, ip_addr+":"+port)
	if err != nil {
		fmt.Println("Listening error:\n 	", err) // do not just panic here
		return
	}

	fmt.Println("Start listening ...:")

	// 2. wait for connection
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Connection to client failed:\n	%v\n", err) // print the value in a default format
			continue
		}
		// 3. open connection, handle request
		fmt.Printf("Client connected:\n		%v\n", conn) // todo: distinguish clients
		go process(conn)
	}
	defer listener.Close() // maybe not needed
}

package main

import (
	"bufio"
	"fmt"
	. "matrix/file_io"
	"net"
	"os"
	"strings"
)

// reuse func in input.go, read local text files and send bytes to server
// then receive bytes from server, show them and store in txt file

const (
	ip_addr      = "127.0.0.1"
	port         = "50000"
	protocol     = "tcp"
	input_method = "file" //options: "file" "console"
	matriceA_raw = "a.txt"
	matriceB_raw = "b.txt"
	start_phrase = "\nsend_start\n"
	next_phrase  = "\nsend_next\n"
	end_phrase   = "\nsend_end\n"
)

func trans(connection net.Conn, data []byte) {
	_, err := connection.Write(data)
	if err != nil {
		fmt.Printf("Write failed:\n		%v\n", err)
		// break // todo: fix break
	}
}

func main() {
	//1. connect to server
	conn, err := net.Dial(protocol, ip_addr+":"+port)
	if err != nil {
		fmt.Printf("Connection to server failed:\n		%v\n", err)
		return
	}
	fmt.Println("Connection Established...:")

	//2. read file or console input
	if input_method == "console" {
		reader := bufio.NewReader(os.Stdin)
		for {
			data, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Read from console failed:\n		%v\n", err)
				break
			}

			data = strings.TrimSpace(data)
			// 3. transmit to server
			trans(conn, []byte(data))
		}
	} else if input_method == "file" {
		// var mA, mB []byte
		mA, mB := Input(matriceA_raw, matriceB_raw) //todo: a list for multiple matrix
		// var data_list [2]uint8{mA, mB}
		// for { todo: use for loop here
		trans(conn, []byte(start_phrase))
		trans(conn, mA)
		trans(conn, []byte(next_phrase))
		// trans(conn, []byte(start_phrase))
		trans(conn, mB)
		trans(conn, []byte(end_phrase))
		defer conn.Close()
	}
}

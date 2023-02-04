package main

import (
	"bytes"
	"fmt"
	. "matrix/file_io"
	. "matrix/matrix_ops"
	"net"
	"regexp"
	"strings"
	// "strconv"
)

// save no file on server

const (
	ip_addr      = "127.0.0.1"
	port         = "50000"
	protocol     = "tcp"
	buffer_size  = 10240 // max size for all matrix
	start_phrase = "\nsend_start\n"
	end_phrase   = "\nsend_end\n"
)

func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

func trans(connection net.Conn, data []byte) {
	_, err := connection.Write(data)
	if err != nil {
		fmt.Printf("Write failed:\n		%v\n", err)
		// break // todo: fix break
	}
}

// handle request, type: net.Conn
func process(conn net.Conn) {
	var full_buf []byte
	// defer conn.Close()
	for {
		var buf [buffer_size]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("Connection reading failed %v:\n	%v\n", conn, err)
			break
		}
		// fmt.Printf("Content from client %v:\n%v\n", conn, string(buf[:n]))
		full_buf = BytesCombine(full_buf, buf[:n])
	}
	fmt.Printf("Full content from %v in buffer:\n%v\n", conn, string(full_buf))

	regex := regexp.MustCompile(start_phrase + `(\s|.)+?` + end_phrase)
	matrix_list := regex.FindAllString(string(full_buf), -1)
	fmt.Print(matrix_list)
	matrix_list[0] = strings.Replace(matrix_list[0], "\nsend_start\n", "", -1)
	matrix_list[0] = strings.Replace(matrix_list[0], "\nsend_end\n", "", -1)

	matrix_list[1] = strings.Replace(matrix_list[1], "\nsend_start\n", "", -1)
	matrix_list[1] = strings.Replace(matrix_list[1], "\nsend_end\n", "", -1)
	mA, mB := Output(matrix_list[0], matrix_list[1])
	// fmt.Println(mA, mB)
	matC := MultiplicationMatricielle(mA, mB) // then to all matrix in list
	fmt.Println(matC)
	// var out_str string
	// float64 to string
	// for _, one_list := range matC {
	// 	for _, value := range one_list{
	// 		out_str = out_str + strconv.FormatFloat(value, ?, 64)
	// 	}
	// }
	// string to byte
	// trans(conn, matC)
	defer conn.Close()
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
			fmt.Printf("Connection to client failed %v:\n	%v\n", conn, err) // print the value in a default format
			continue
		}
		// 3. open connection, handle request
		fmt.Printf("Client connected:\n		%v\n", conn) // todo: distinguish clients
		go process(conn)
	}
	defer listener.Close() // maybe not needed
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	. "matrix/file_io"
	"net"
	"os"
	"strings"
)

const (
	ip_addr      = "127.0.0.1"
	port         = "50000"
	protocol     = "tcp"
	input_method = "file" //options: "file" "console"
	// input_method   = "console" //options: "file" "console"
	buffer_size    = 10240 // max size for one buffer
	start_phrase   = "\nsend_start\n"
	end_phrase     = "\nsend_end\n"
	matriceA_raw   = "matrix_input_txt/a.txt"
	matriceB_raw   = "matrix_input_txt/b.txt"
	matriceC_raw   = "matrix_input_txt/c.txt"
	matriceFin_raw = "matrix_input_txt/end.txt" // use this if matrix number is not pair
	resultFile     = "./res.txt"
)

// combine all bytes from n buffers
func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

// send
func trans(connection net.Conn, data []byte) {
	_, err := connection.Write(data)
	if err != nil {
		fmt.Printf("Write failed:\n		%v\n", err)
		panic(err)
	}
}

// shutdown TCP connection to indicate EOF to the other side
func shutdownWrite(conn net.Conn) {
	if v, ok := conn.(interface{ CloseWrite() error }); ok {
		v.CloseWrite()
	}
}

func main() {
	//1. Dial server
	conn, err := net.Dial(protocol, ip_addr+":"+port)
	if err != nil {
		fmt.Printf("Connection to server failed:\n		%v\n", err)
		return
	}
	fmt.Println("Connection Established...:\nSending:")

	//2. Read file or console input
	if input_method == "console" {
		index := 1
		for {
			fmt.Printf("#	N.%v matrix:\n", index)
			index++
			reader := bufio.NewReader(os.Stdin)
			data, err := reader.ReadString('/')
			if err != nil {
				fmt.Printf("Read from console failed:\n		%v\n", err)
				continue
			}
			data = strings.TrimSpace(data)
			data = strings.Replace(data, "/", "", -1)
			if strings.Replace(data, "\r\n", "", -1) == "end" { // when input == 'end/', one matrix input is done
				// trans(conn, []byte(data))
				shutdownWrite(conn)
				break
			} else if strings.Replace(data, "\n", "", -1) == "end" {
				// trans(conn, []byte(data))
				shutdownWrite(conn)
				break
			}
			// 3. Transmit to server
			trans(conn, []byte(start_phrase))
			trans(conn, []byte(data))
			trans(conn, []byte(end_phrase))
		}
	} else if input_method == "file" {
		// test
		// mA, mB := Input(matriceA_raw, matriceB_raw) //todo: a list for multiple matrix
		// mC, mD := Input(matriceC_raw, matriceFin_raw)
		// // var data_list [2]uint8{mA, mB}
		// // for { todo: use for loop here
		// // 3. Transmit to server
		// trans(conn, []byte(start_phrase))
		// trans(conn, mA)
		// trans(conn, []byte(end_phrase))
		// trans(conn, []byte(start_phrase))
		// trans(conn, mB)
		// trans(conn, []byte(end_phrase))
		// trans(conn, []byte(start_phrase))
		// trans(conn, mC)
		// trans(conn, []byte(end_phrase))
		// trans(conn, []byte(start_phrase))
		// trans(conn, mD)
		// trans(conn, []byte(end_phrase))
		// shutdownWrite(conn)

		// ! assert matrix_raw_list pair
		var matrix_raw_list = [...]string{matriceA_raw, matriceB_raw, matriceC_raw, matriceFin_raw}
		var matrix_trans_list [][]byte
		var trans_a, trans_b []byte
		var trans_temp string
		for num, matrix_trans := range matrix_raw_list {
			if num%2 != 0 {
				trans_a, trans_b = Input(trans_temp, matrix_trans)
				matrix_trans_list = append(matrix_trans_list, trans_a, trans_b)
			} else {
				trans_temp = matrix_trans
			}
		}
		for _, matrix_trans := range matrix_trans_list {
			trans(conn, []byte(start_phrase))
			trans(conn, matrix_trans)
			trans(conn, []byte(end_phrase))
		}
		shutdownWrite(conn)
	}

	// 4. Read from server
	var full_buf []byte
	for {
		var buf [buffer_size]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("Connection reading failed %v:\n		%v\n", conn, err)
			break
		}
		full_buf = BytesCombine(full_buf, buf[:n])
	}
	fmt.Printf("Full content from %v in buffer:\n%v\n", conn, string(full_buf))

	defer conn.Close()

	Save(resultFile, string(full_buf))
}

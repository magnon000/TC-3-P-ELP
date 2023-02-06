package main

import (
	"bytes"
	"fmt"
	. "matrix/file_io"
	. "matrix/matrix_ops"
	"net"
	"regexp"
	"strconv"
	"strings"
)

const (
	ip_addr      = "127.0.0.1"
	port         = "50000"
	protocol     = "tcp"
	buffer_size  = 10240 // max size for one buffer
	start_phrase = "\nsend_start\n"
	end_phrase   = "\nsend_end\n"
	num_split    = ","
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

// handle request, type: net.Conn
func process(conn net.Conn) {
	var full_buf []byte
	// a. Read from client
	for {
		var buf [buffer_size]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("Connection reading failed %v:\n		%v\n", conn, err)
			break
		}
		// fmt.Printf("Content from client %v:\n%v\n", conn, string(buf[:n]))
		full_buf = BytesCombine(full_buf, buf[:n])
		// if string(full_buf)[len(full_buf)-4:len(full_buf)-1] == "end" {
		// 	break
		// }
	}
	fmt.Printf("Full content from %v in buffer:\n%v\n", conn, string(full_buf))
	// regex each matrix
	regex := regexp.MustCompile(start_phrase + `(\s|.)+?` + end_phrase)
	matrix_raw_list := regex.FindAllString(string(full_buf), -1)
	// test
	// matrix_raw_list[0] = strings.Replace(matrix_raw_list[0], "\nsend_start\n", "", -1)
	// matrix_raw_list[0] = strings.Replace(matrix_raw_list[0], "\nsend_end\n", "", -1)

	// matrix_raw_list[1] = strings.Replace(matrix_raw_list[1], "\nsend_start\n", "", -1)
	// matrix_raw_list[1] = strings.Replace(matrix_raw_list[1], "\nsend_end\n", "", -1)
	// mA, mB := Output(matrix_raw_list[0], matrix_raw_list[1])
	// // fmt.Println("mA:", mA, "\nmB:", mB)
	// matC := MultiplicationMatricielle(mA, mB) // then to all matrix in list

	// matrix raw string list
	var matrix_output_list []string
	for _, matrix_raw := range matrix_raw_list {
		matrix_raw = strings.Replace(matrix_raw, "\nsend_start\n", "", -1)
		matrix_raw = strings.Replace(matrix_raw, "\nsend_end\n", "", -1)
		matrix_output_list = append(matrix_output_list, matrix_raw)
	}
	if len(matrix_output_list)%2 == 1 { // handle not pair matrix list
		matrix_output_list = append(matrix_output_list, "")
	}
	// matrix Output string list (list of [][]float64)
	var str_temp string
	var out_a, out_b [][]float64
	var matrix_mul_list [][][]float64
	for num, matrix_out := range matrix_output_list { // ! assert list has pair num of matrix
		if num%2 != 0 {
			out_a, out_b = Output(str_temp, matrix_out)
			if len(out_b) == 0 || out_b == nil {
				matrix_mul_list = append(matrix_mul_list, out_a)
			} else {
				matrix_mul_list = append(matrix_mul_list, out_a, out_b)
			}
			str_temp = ""
		} else if num%2 == 0 {
			str_temp = matrix_out
		}
	}
	// fmt.Println("mul list:", matrix_mul_list)
	// matrix ops
	var mRes [][]float64
	for num, matrix_mul := range matrix_mul_list {
		if matrix_mul == nil {
			continue
		}
		if num != 0 {
			mRes = MultiplicationMatricielle(mRes, matrix_mul)
		} else if num == 0 {
			mRes = matrix_mul
		}
	}

	fmt.Printf("### Result for %v: %v\n", conn, mRes)
	fmt.Println("##################################################")

	// b. Transmist to client
	// float64 to string
	var out_str string
	for _, one_list := range mRes {
		for _, value := range one_list {
			out_str = out_str + strconv.FormatFloat(value, 'f', 2, 64) + num_split
		}
		out_str = out_str[0:len(out_str)-1] + string("\n")
	}
	// string to byte
	trans(conn, []byte(out_str))
	defer conn.Close()
	fmt.Printf("!!! Client %v job done, disconnected.\n", conn)
}

func main() {
	// 1. start listening
	listener, err := net.Listen(protocol, ip_addr+":"+port)
	if err != nil {
		fmt.Println("Listening error:\n 	", err) // do not panic here
		return
	}

	fmt.Println("Start listening ...:")

	// 2. wait for connection
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Connection to client failed %v:\n 	%v\n", conn, err) // print the value in a default format
			continue
		}
		// 3. open connection, handle request
		fmt.Printf("Client connected:\n		%v\n", conn) // distinguish clients
		go process(conn)                              // todo: handle error
	}
	defer listener.Close() // maybe not needed
}

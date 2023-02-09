package main

import (
	"bufio"
	"fmt"
	. "matrix/file_io"
	"net"
	"os"
	"strings"
	"time"
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
	matriceA_raw   = "matrix_input_txt/8236_A.txt"
	matriceB_raw   = "matrix_input_txt/8236_B.txt"
	matriceC_raw   = "matrix_input_txt/c.txt"
	matriceD_raw   = "matrix_input_txt/d.txt"
	matriceE_raw   = "matrix_input_txt/e.txt"
	matriceF_raw   = "matrix_input_txt/f.txt"
	matriceG_raw   = "matrix_input_txt/g.txt"
	matriceFin_raw = "matrix_input_txt/end.txt" // use this if matrix number is not pair
	matriceNbr     = 2                          // ! assert matriceNbr pair
	resultFile     = "./res.txt"
)

// ! assert matrix_raw_list pair
// var matrix_raw_list = [...]string{matriceA_raw, matriceB_raw}
var matrix_raw_list = [...]string{ // TODO: read .txt names from a file?
	matriceA_raw,
	matriceB_raw,
	matriceC_raw,
	matriceD_raw,
	matriceE_raw,
	matriceF_raw,
	matriceG_raw,
	matriceFin_raw}

// shutdown TCP connection to indicate EOF to the other side
func shutdownWrite(conn net.Conn) {
	if v, ok := conn.(interface{ CloseWrite() error }); ok {
		v.CloseWrite()
	}
}

func main() {
	//0. start chrono
	startClientProgram := time.Now()
	
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
			Trans(conn, []byte(start_phrase))
			Trans(conn, []byte(data))
			Trans(conn, []byte(end_phrase))
		}
	} else if input_method == "file" {
		// test with 2 matrix
		// mA, mB := Input(matriceA_raw, matriceB_raw)
		// mC, mD := Input(matriceC_raw, matriceFin_raw)
		// // var data_list [2]uint8{mA, mB}
		// Trans(conn, []byte(start_phrase))
		// Trans(conn, mA)
		// Trans(conn, []byte(end_phrase))
		// Trans(conn, []byte(start_phrase))
		// Trans(conn, mB)
		// Trans(conn, []byte(end_phrase))
		// Trans(conn, []byte(start_phrase))
		// Trans(conn, mC)
		// Trans(conn, []byte(end_phrase))
		// Trans(conn, []byte(start_phrase))
		// Trans(conn, mD)
		// Trans(conn, []byte(end_phrase))
		// shutdownWrite(conn)

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
		for _, matrix_trans := range matrix_trans_list[:matriceNbr] {
			Trans(conn, []byte(start_phrase))
			Trans(conn, matrix_trans)
			Trans(conn, []byte(end_phrase))
		}
		shutdownWrite(conn)
	}

	// 4. Read from server
	startReception := time.Now()
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
	endReception := time.Now()
	receptionTime := endReception.Sub(startReception)
	fmt.Println("Durée de la réception :",receptionTime)

	defer conn.Close()

	Save(resultFile, string(full_buf))
	
	endClientProgram := time.Now()
	fmt.Println("Durée totale complète :",endClientProgram.Sub(startClientProgram))
}

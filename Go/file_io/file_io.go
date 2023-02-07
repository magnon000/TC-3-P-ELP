package file_io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// const (
// 	matriceA = "a.txt"
// 	matriceB = "b.txt"
// )

func Input(matrice_a string, matrice_b string) ([]byte, []byte) {
	contentA, err := os.ReadFile(matrice_a)
	if err != nil {
		panic(err)
	}
	fmt.Println("A =")
	fmt.Println(string(contentA))
	contentB, err := os.ReadFile(matrice_b)
	if err != nil {
		panic(err)
	}
	fmt.Println("B =")
	fmt.Println(string(contentB))
	return contentA, contentB
}

func Output(content_a string, content_b string) ([][]float64, [][]float64) {
	// mA := len(content_a) /* len(contentA[0])*/
	// mB := len(content_b) /* len(contentB[0])*/
	// fmt.Println(mA, mB)
	if content_b == "" {
		A_lines := strings.Split(content_a, "\n")
		A_line_len := len(A_lines)
		A_col_len := len(strings.Split(A_lines[0], ","))
		fmt.Println("\nmatrice A de taille: (", A_line_len, A_col_len, ")")
		var matriceA [][]float64
		for y := 0; y < A_line_len; y++ {
			temp_line := A_lines[y]
			temp_line = strings.Replace(temp_line, "\n", "", -1)
			temp_line = strings.Replace(temp_line, "\r", "", -1)
			a_line := strings.Split(temp_line, ",")
			fmt.Println("a_line:", a_line)
			var out_line []float64
			for index := 0; index < A_col_len; index++ {
				num, erreur := strconv.ParseFloat(a_line[index], 64)
				if erreur != nil {
					fmt.Println("")
					panic(erreur)
				}
				out_line = append(out_line, num)
			}
			matriceA = append(matriceA, out_line)
		}
		fmt.Println("\nmatrice A de taille: (", A_line_len, A_col_len, ")")
		return matriceA, [][]float64{}
	} else {
		// // fmt.Println(A_lines[0])
		// // fmt.Printf("type: %T\n", A_lines[0]) // one line in matrix list is of type string
		// A_lines := strings.SplitAfter(content_a, "\n") // do not use string.Split here, \n causes assignment error
		// B_lines := strings.SplitAfter(content_b, "\n") // list contains all lines of matrix
		// A_col_len := len(A_lines[0]) / 2 // \n counts in len, \r\n or \n cause error
		// B_col_len := len(B_lines[0]) / 2
		A_lines := strings.Split(content_a, "\n")
		B_lines := strings.Split(content_b, "\n") // list contains all lines of matrix
		A_line_len := len(A_lines)                // ligne de matrice A
		B_line_len := len(B_lines)                // ligne de matrice B
		A_col_len := len(strings.Split(A_lines[0], ","))
		B_col_len := len(strings.Split(B_lines[0], ","))
		fmt.Println("\nmatrice A de taille: (", A_line_len, A_col_len, ")")
		fmt.Println("matrice B de taille: (", B_line_len, B_col_len, ")")

		// todo: verify this in server script
		// if A_line_len != B_col_len {
		// 	panic("Pas possible de multiplier: Longueur de ligne de A != Longueur de colonne de B !")
		// }

		var matriceA [][]float64
		var matriceB [][]float64
		//matrice A
		for y := 0; y < A_line_len; y++ {
			temp_line := A_lines[y]
			temp_line = strings.Replace(temp_line, "\n", "", -1) // in linux a line ends with "\n"
			temp_line = strings.Replace(temp_line, "\r", "", -1) // in win10 a line ends with "\r\n"
			a_line := strings.Split(temp_line, ",")              // list of string
			var out_line []float64
			for index := 0; index < A_col_len; index++ {
				num, erreur := strconv.ParseFloat(a_line[index], 64)
				if erreur != nil {
					panic(erreur)
				}
				out_line = append(out_line, num)
			}
			matriceA = append(matriceA, out_line)
		}
		fmt.Println("matriceA =", matriceA)
		//matrice B
		for y := 0; y < B_line_len; y++ {
			temp_line := B_lines[y]
			temp_line = strings.Replace(temp_line, "\n", "", -1) // in linux a line ends with "\n"
			temp_line = strings.Replace(temp_line, "\r", "", -1) // in win10 a line ends with "\r\n"
			a_line := strings.Split(temp_line, ",")              // list of string
			var out_line []float64
			for index := 0; index < B_col_len; index++ {
				num, erreur := strconv.ParseFloat(a_line[index], 64)
				if erreur != nil {
					panic(erreur)
				}
				out_line = append(out_line, num)
			}
			matriceB = append(matriceB, out_line)
		}
		fmt.Println("matriceB =", matriceB)
		// A_line_len = len(matriceA)
		// A_col_len = len(matriceA[0])
		// fmt.Println("\nmatrice A de taille: (", A_line_len, A_col_len, ")")
		// B_line_len = len(matriceB)
		// B_col_len = len(matriceB[0])
		// fmt.Println("matrice B de taille: (", B_line_len, B_col_len, ")")
		return matriceA, matriceB
	}
}

func Save(fileName string, content string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Fail to open file: %v\n		%v", fileName, err)
		panic(err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(content)
	// Flush write buffer to real file
	write.Flush()
}

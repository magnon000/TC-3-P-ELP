package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	matriceA = "a.txt"
	matriceB = "b.txt"
)

func Input() {
	contentA, err := os.ReadFile(matriceA)
	if err != nil {
		panic(err)
	}
	fmt.Println("A =")
	fmt.Println(string(contentA))
	contentB, err := os.ReadFile(matriceB)
	if err != nil {
		panic(err)
	}
	fmt.Println("B =")
	fmt.Println(string(contentB), "\n")

	// mA := len(contentA) /* len(contentA[0])*/
	// mB := len(contentB)  /* len(contentB[0])*/
	// fmt.Println(mA, mB)

	A_lines := strings.SplitAfter(string(contentA), "\n") // do not use string.Split here, \n causes assignment error
	B_lines := strings.SplitAfter(string(contentB), "\n") // list contains all lines of matrix
	// fmt.Println(A_lines[0])
	// fmt.Printf("type: %T\n", A_lines[0]) // one line in matrix list is of type string
	A_line_len := len(A_lines)       // ligne de matrice A
	B_line_len := len(B_lines)       // ligne de matrice B
	A_col_len := len(A_lines[0]) / 2 // \n counts
	B_col_len := len(B_lines[0]) / 2
	// fmt.Println(A_lines)
	// fmt.Println(B_lines, "\n")
	fmt.Println("matric A de taille: (", A_line_len, A_col_len, ")")
	fmt.Println("matric B de taille: (", B_line_len, B_col_len, ")\n")

	if A_line_len != B_col_len {
		panic("Pas possible de multiplier: Longueur de ligne de A != Longueur de colonne de B !")
	}

	var matriceA [][]float64
	var matriceB [][]float64
	//matrice A
	for y := 0; y < A_line_len-1; y++ {
		temp_line := strings.Split(A_lines[y], ",") // list of string
		var out_line []float64
		for index := 0; index < A_col_len-1; index++ {
			num, erreur := strconv.ParseFloat(temp_line[index], 64)
			if erreur != nil {
				panic(erreur)
			}
			out_line = append(out_line, num)
		}
		matriceA = append(matriceA, out_line)
	}
	fmt.Println("matriceA =", matriceA)
	// //matrice B
	for y := 0; y < B_line_len-1; y++ {
		temp_line := strings.Split(B_lines[y], ",") // list of string
		var out_line []float64
		for index := 0; index < B_col_len-1; index++ {
			num, erreur := strconv.ParseFloat(temp_line[index], 64)
			if erreur != nil {
				panic(erreur)
			}
			out_line = append(out_line, num)
		}
		matriceB = append(matriceB, out_line)
	}
	fmt.Println("matriceB =", matriceB)

}

func main() {
	Input()
}

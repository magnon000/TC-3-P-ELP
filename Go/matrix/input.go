package main

import (
	"fmt"
	"os"
)

const (
	matriceA = "a.txt"
	matriceB = "b.txt"
)

func input() {
	contentA, err := os.ReadFile(matriceA)
	if err != nil {
		panic(err)
	}
	fmt.Println("A =\n", string(contentA))
	contentB, err := os.ReadFile(matriceB)
	if err != nil {
		panic(err)
	}
	fmt.Println("B =\n", string(contentB))




	//blablabla
	mA, nA := len(A), len(A[0])
	nB, pB := len(B), len(B[0])
	matriceA := make([][]float64, tailleAy)
	matriceB := make([][]float64, tailleBy)

	//matrice A
	for y := 0; y < tailleAy; y++ {
		ligne := make([]float64, tailleAx)
		for x := 0; x < tailleAx; x++ {
			ligne[x] = float64(coefficientnumérique)
		}
		matriceA[y] = ligne
	
	}

	//matrice B
	for y := 0; y < tailleAy; y++ {
		ligne := make([]float64, tailleBx)
		for x := 0; x < tailleBx; x++ {
			ligne[x] = float64(coefficientnumérique)
		}
		matriceB[y] = ligne
	}
}

package main

import (
	"fmt"
)

func multiplication(matriceA [][]float64, matriceB [][]float64) [][]float64{
	//récupérer tailles (vérifié dans input : tailleAx==tailleBy)
	tailleAy := len(matriceA)
	tailleBx := len(matriceB[0])
	tailleCommune := len(matriceB)

	//créer et remplir matrice résultat C
	matriceC := make([][]float64, tailleAy)
	for y:= 0; y<tailleAy;y++ {
		ligne := make([]float64, tailleBx)
		for x := 0; x<tailleBx; x++ {
			//ligne[x] = float64(3.14)
			for i := 0; i<tailleCommune; i++ {
				ligne[x] += matriceA[y][i]*matriceB[i][x]
			}
		}
		matriceC[y] = ligne
	}

	return matriceC

}

func main(){
	fmt.Println("Test")
	//input("test2")

	tailleAx:=3
	tailleAy:=3
	tailleBx:=3
	tailleBy:=3

	matriceA := make([][]float64, tailleAy)
	matriceB := make([][]float64, tailleBy)

	//matrice A
	for y:= 0; y<tailleAy;y++ {
		ligne := make([]float64, tailleAx)
		for x := 0; x<tailleAx; x++ {
			if x==y {
				ligne[x] = 1
			} else {
				ligne[x] = 0
			}
		}
		matriceA[y] = ligne
	}

	//matrice B
	for y:= 0; y<tailleAy;y++ {
		ligne := make([]float64, tailleBx)
		for x := 0; x<tailleBx; x++ {
			if x==y {
				ligne[x] = 1
			} else {
				ligne[x] = 0
			}
		}
		matriceB[y] = ligne
	}

	fmt.Println(multiplication(matriceA, matriceB))

}
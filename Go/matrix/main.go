package main

import (
	"fmt"
	"sync"
	"time"
)

type portionTodo struct {
	positionY int
	matriceB *[][]float64 //on a besoin de toutes les colonnes de B pour calculer 1 ligne de la mat produit
	ligneYmatriceA *[]float64 //sera parcourue à chaque colonne X de B
}

type portionFinished struct {
	positionY int
	ligneFinale *[]float64 //la ligne Y remplie de la matrice produit
}

func multiplicationTotale(matriceA [][]float64, matriceB [][]float64) [][]float64{
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

//calcule 1 ligne de la matrice produit
func workerMultiplicationPartielle(wg *sync.WaitGroup, jobs chan *portionTodo, result chan *portionFinished){
	//récupération des infos du channel
	job :=<- jobs
	tailleRecurrente := len(*job.ligneYmatriceA)
	nombreColonnes := len((*job.matriceB)[0])

	//partie calcul
	ligne := make([]float64, nombreColonnes)
	for x:=0; x<nombreColonnes; x++{
		ligne[x] = 0
		for i:=0; i<tailleRecurrente; i++ {
			ligne[i] += (*job.ligneYmatriceA)[i] * (*job.matriceB)[i][x]
		}
	}

	//rendu du travail effectué
	rendu := portionFinished{positionY: job.positionY}
	rendu.ligneFinale = &ligne
	result <- &rendu
	wg.Done()
}

func managerMultiplicationPartielle(matriceA [][]float64, matriceB [][]float64) [][]float64{
	nombreWorkers := len(matriceA)
	matriceC := make([][]float64, nombreWorkers)

	var wg sync.WaitGroup
	jobsChannel := make(chan *portionTodo, nombreWorkers)
	resultChannel := make(chan *portionFinished, nombreWorkers)

	for i:=0; i<nombreWorkers; i++ {
		wg.Add(1)
		go workerMultiplicationPartielle(&wg,jobsChannel,resultChannel)
	}

	go func(){
		for i:=0; i<nombreWorkers; i++ {
			var portion portionTodo
			portion = portionTodo{positionY: i, matriceB: &matriceB, ligneYmatriceA: &(matriceA[i])}
			jobsChannel <- &portion
		}
	}()

	wg.Wait()

	for i:=0; i<nombreWorkers; i++ {
		resultat :=<- resultChannel
		pos := resultat.positionY
		matriceC[pos] = *resultat.ligneFinale
	}

	return matriceC
}

func main(){
	fmt.Println("Test")
	//input("test2")

	tailleAx:=1000
	tailleAy:=1000
	tailleBx:=1000
	tailleBy:=1000

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

	//test 1 thread
	start1 := time.Now()
	fmt.Println(multiplicationTotale(matriceA, matriceB))
	end1 := time.Now()
	fmt.Println("Durée du calcul 1 thread :",end1.Sub(start1))

	//test parallélisme
	start2 := time.Now()
	fmt.Println(managerMultiplicationPartielle(matriceA, matriceB))
	end2 := time.Now()
	fmt.Println("Durée du calcul 1 thread :",end1.Sub(start1))
	fmt.Println("Durée du calcul avec goroutines :",end2.Sub(start2))


}
package matrix_ops

import (
	"fmt"
	"time"
)

var (
	NOMBRE_GOROUTINES = 16
)

type portionTodo struct {
	stop           bool
	positionY      int
	matriceB       *[][]float64 //on a besoin de toutes les colonnes de B pour calculer 1 ligne de la mat produit
	ligneYmatriceA *[]float64   //sera parcourue à chaque colonne X de B
}

type portionFinished struct {
	positionY   int
	ligneFinale *[]float64 //la ligne Y remplie de la matrice produit
}

// phase 0 (à titre de comparaison)
func multiplicationTotale(matriceA [][]float64, matriceB [][]float64) [][]float64 {
	//récupérer tailles (vérifié dans input : tailleAx==tailleBy)
	tailleAy := len(matriceA)
	tailleBx := len(matriceB[0])
	tailleCommune := len(matriceB)

	//créer et remplir matrice résultat C
	matriceC := make([][]float64, tailleAy)
	for y := 0; y < tailleAy; y++ {
		ligne := make([]float64, tailleBx)
		for x := 0; x < tailleBx; x++ {
			for i := 0; i < tailleCommune; i++ {
				ligne[x] += matriceA[y][i] * matriceB[i][x]
			}
		}
		matriceC[y] = ligne
	}

	return matriceC

}

// worker calcule 1 ligne de la matrice produit par tour
func workerMultiplicationPartielle(jobs chan *portionTodo, result chan *portionFinished) {
	travail_a_faire := true
	for travail_a_faire {
		job := <-jobs
		if job.stop {
			travail_a_faire = false
		} else {
			//récupération des infos du channel
			tailleRecurrente := len(*job.ligneYmatriceA)
			nombreColonnes := len((*job.matriceB)[0])

			//partie calcul
			ligne := make([]float64, nombreColonnes)
			for x := 0; x < nombreColonnes; x++ {
				ligne[x] = 0
				for i := 0; i < tailleRecurrente; i++ {
					ligne[i] += (*job.ligneYmatriceA)[i] * (*job.matriceB)[i][x]
				}
			}

			//rendu du travail effectué
			rendu := portionFinished{positionY: job.positionY}
			rendu.ligneFinale = &ligne
			result <- &rendu
		}
	}
}

// manager des workers
func MultiplicationMatricielle(matriceA [][]float64, matriceB [][]float64) [][]float64 {
	nombre_total_travaux := len(matriceA)
	matriceC := make([][]float64, nombre_total_travaux)

	jobsChannel := make(chan *portionTodo, nombre_total_travaux+NOMBRE_GOROUTINES)
	resultChannel := make(chan *portionFinished, nombre_total_travaux)

	//lancer les goroutines
	for i := 0; i < NOMBRE_GOROUTINES; i++ {
		go workerMultiplicationPartielle(jobsChannel, resultChannel)
	}

	go func() {
		//pousser les travaux
		for i := 0; i < nombre_total_travaux; i++ {
			var portion portionTodo
			portion = portionTodo{stop: false, positionY: i, matriceB: &matriceB, ligneYmatriceA: &(matriceA[i])}
			jobsChannel <- &portion
		}

		//pousser les STOP afin que les goroutines arrêtent d'écouter
		//besoin de remplir tous les champs de portionTodo cependant (donc baits)
		bait1 := make([][]float64, 1)
		bait2 := make([]float64, 1)
		for i := 0; i < NOMBRE_GOROUTINES; i++ {
			var portion portionTodo
			portion = portionTodo{stop: true, positionY: -1, matriceB: &bait1, ligneYmatriceA: &bait2}
			jobsChannel <- &portion
		}
	}()

	//récupérer les travaux et les mettre dans la matrice produit
	for i := 0; i < nombre_total_travaux; i++ {
		resultat := <-resultChannel
		pos := resultat.positionY
		matriceC[pos] = *resultat.ligneFinale
	}

	return matriceC
}

func creationMatricesTests() {
	tailleAx := 500
	tailleAy := 500
	tailleBx := 500
	tailleBy := 500

	matriceA := make([][]float64, tailleAy)
	matriceB := make([][]float64, tailleBy)

	//matrice A
	for y := 0; y < tailleAy; y++ {
		ligne := make([]float64, tailleAx)
		for x := 0; x < tailleAx; x++ {
			if x == y {
				ligne[x] = 1
			} else {
				ligne[x] = 0
			}
		}
		matriceA[y] = ligne
	}

	//matrice B
	for y := 0; y < tailleAy; y++ {
		ligne := make([]float64, tailleBx)
		for x := 0; x < tailleBx; x++ {
			if x == y {
				ligne[x] = 1
			} else {
				ligne[x] = 0
			}
		}
		matriceB[y] = ligne
	}
}

func timeTest(matriceA [][]float64, matriceB [][]float64) {
	fmt.Println("Test")

	//test 1 thread
	start1 := time.Now()
	fmt.Println(multiplicationTotale(matriceA, matriceB))
	end1 := time.Now()

	//test parallélisme
	start2 := time.Now()
	fmt.Println(MultiplicationMatricielle(matriceA, matriceB))
	end2 := time.Now()

	fmt.Println("Durée du calcul 1 thread :", end1.Sub(start1))
	fmt.Println("Durée du calcul avec goroutines :", end2.Sub(start2))

}

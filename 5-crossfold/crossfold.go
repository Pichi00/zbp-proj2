/*
	Przykład  ilustruje użycie biblioteki goLearn do przeprowadzenia klasyfikacji na zbiorze danych irysów (Iris) przy użyciu algorytmu lasu losowego
	(Random Forest). Celem tego przykładu jest zademonstrowanie wpływu liczby drzew w lesie losowym na dokładność klasyfikacji oraz porównanie wyników z
	oraz bez dyskretyzacji danych.

	Rezultatem jest wydrukowanie tabeli wyników, która pokazuje, jak zmienia się dokładność klasyfikacji (wraz z odchyleniem)
	w zależności od liczby drzew w lesie losowym.
	Oczekuje się, że wraz ze wzrostem liczby drzew, dokładność modelu będzie się poprawiać, a odchylenie standardowe będzie się zmniejszać,
	co oznacza większą stabilność modelu.
*/

package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/ensemble"
	"github.com/sjwhitworth/golearn/evaluation"
)

func main() {

	var tree base.Classifier

	// Załadowanie danych z datasetu
	iris, err := base.ParseCSVToInstances("datasets/iris_headers.csv", true)
	if err != nil {
		panic(err)
	}

	for i := 1; i < 60; i += 2 {
		rand.Seed(44111342)

		// Tworzenie nowego lasu z coraz większą ilością drzew
		tree = ensemble.NewRandomForest(i, 4)
		// Walidacja krzyżowa z pięcioma podziałami (5-fold cross-validation)
		cfs, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(iris, tree, 5)
		if err != nil {
			panic(err)
		}

		// Obliczenie średniej, wariancji i odchylenia standardowego dokładności klasyfikacji
		mean, variance := evaluation.GetCrossValidatedMetric(cfs, evaluation.GetAccuracy)
		stdev := math.Sqrt(variance)

		fmt.Printf("%d\t%.2f\t(+/- %.2f)\n", i, mean, stdev*2)
	}
}

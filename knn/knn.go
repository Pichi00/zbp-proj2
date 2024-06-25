package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func main() {
	// Ustawienie ziarna losowości
	rand.Seed(time.Now().UnixNano())

	// Załadowanie danych z pliku ARFF
	rawData, err := base.ParseCSVToInstances("../datasets/iris_headers.csv", true)
	if err != nil {
		panic(fmt.Sprintf("Nie udało się załadować danych: %v", err))
	}

	// Podzielenie danych na zestaw treningowy i testowy
	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.70)

	// Inicjalizacja klasyfikatora k-NN z k = 3
	cls := knn.NewKnnClassifier("euclidean", "linear", 3)

	// Trenowanie klasyfikatora
	cls.Fit(trainData)

	// Przewidywanie klas dla danych testowych
	predictions, err := cls.Predict(testData)
	if err != nil {
		panic(fmt.Sprintf("Nie udało się przeprowadzić predykcji: %v", err))
	}

	// Ocena dokładności klasyfikatora
	confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		panic(fmt.Sprintf("Nie udało się obliczyć macierzy pomyłek: %v", err))
	}

	accuracy := evaluation.GetAccuracy(confusionMat)
	fmt.Printf("Dokładność klasyfikatora k-NN: %.2f%%\n", accuracy*100)

}

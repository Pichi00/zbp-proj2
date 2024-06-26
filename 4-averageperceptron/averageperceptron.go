/*
Przykład trenowania i oceny klasyfikatora perceptronu średniego na zbiorze danych dotyczących głosów w Kongresie USA

Rezultatem jest wyświetlenie zestawu danych treningowych i testowych oraz podsumowanie wydajności modelu na danych testowych.
Podsumowanie zawiera wskaźniki takie jak precision, recall i f1-score, które pozwalają ocenić skuteczność klasyfikatora perceptronu średniego.
*/

package main

import (
	"fmt"
	"math/rand"

	base "github.com/sjwhitworth/golearn/base"
	evaluation "github.com/sjwhitworth/golearn/evaluation"
	perceptron "github.com/sjwhitworth/golearn/perceptron"
)

func main() {

	rand.Seed(4402201)

	// Wczytanie danych z datasetu
	rawData, err := base.ParseCSVToInstances("datasets/house-votes-84.csv", true)
	if err != nil {
		panic(err)
	}

	// Parametry klasyfikatora
	features := 10           // Liczba cech (atrybutów) w danych wejściowych
	learningRate := 1.2      // Współczynnik uczenia
	startingThreshold := 0.5 // Początkowy próg decyzyjny
	trainError := 0.3        // Błąd treningowy

	// Utworzenie klasyfikatora perceptronu średniego
	cls := perceptron.NewAveragePerceptron(features, learningRate, startingThreshold, trainError)

	// Podział na dane treningowe i testowe
	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.50)
	fmt.Println(trainData)
	fmt.Println(testData)

	// Trenowanie klasyfikatora
	cls.Fit(trainData)

	// Wykonanie predykcji
	predictions := cls.Predict(testData)

	// Ocena modelu
	confusionMat, _ := evaluation.GetConfusionMatrix(testData, predictions)
	fmt.Println(evaluation.GetSummary(confusionMat))
}

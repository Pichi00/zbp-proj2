/*
Celem przykładu jest zademonstrowanie, jak z pomocą biblioteki goLearn wykorzystać model Isolation Forest do wykrywania wartości odstających.

Rezultatem jest wyświetlenie wartości z pliku gaussian_outliers.csv wraz z przypisanymi wynikami anomalii.
Wyniki te pokazują, które wartości są uznawane za odstające przez model Isolation Forest.
Im wyższa wartość 'Anomaly score' tym większe jest prawdopodobieństwo odstawania wartości

*/

package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/trees"
)

func main() {
	// Wczytanie danych z pliku CSV
	csvData, err := base.ParseCSVToInstances("datasets/gaussian_outliers.csv", true)
	if err != nil {
		panic(err)
	}

	// Inicjalizacja modelu Isolation Forest
	forest := trees.NewIsolationForest(100, 100, 850)

	// Trenowanie modelu na danych
	forest.Fit(csvData)

	// Dokonywanie predykcji (wykrywanie anomalii)
	preds := forest.Predict(csvData)

	// Wyświetlanie wyników dla wartości odstających
	fmt.Println("Wartości odstające:")
	for i := 0; i < 5; i++ {
		fmt.Print("      ")
		fmt.Println("Wartości:", csvData.RowString(i), "Anomaly score:", preds[i])
	}

	// Wyświetlanie wyników dla poprawnych wartości
	fmt.Println("Poprawne wartości:")
	for i := 5; i < 15; i++ {
		fmt.Print("      ")
		fmt.Println("Wartości:", csvData.RowString(i), "Anomaly score:", preds[i])
	}
}

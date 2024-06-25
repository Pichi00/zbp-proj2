// Example of how to use Isolation Forest for outlier detection

package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/trees"
)

func main() {
	csvData, err := base.ParseCSVToInstances("datasets/gaussian_outliers.csv", true)
	if err != nil {
		panic(err)
	}

	forest := trees.NewIsolationForest(100, 100, 850)
	forest.Fit(csvData)
	preds := forest.Predict(csvData)

	fmt.Println("Wartości odstające:")
	for i := 0; i < 5; i++ {
		fmt.Print("      ")
		fmt.Println("Wartości:", csvData.RowString(i), "Anomaly score:", preds[i])
	}

	fmt.Println("Poprawne wartości:")
	for i := 5; i < 15; i++ {
		fmt.Print("      ")
		fmt.Println("Wartości:", csvData.RowString(i), "Anomaly score:", preds[i])
	}
}

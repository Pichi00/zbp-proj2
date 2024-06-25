package main

import (
	"encoding/json"
	"fmt"

	"github.com/sjwhitworth/golearn/base"
)

func main() {
	// Załadowanie danych z datasetu
	iris, err := base.ParseCSVToInstances("datasets/iris_headers.csv", true)
	if err != nil {
		panic(err)
	}

	// Serializacja atrybutów
	for _, a := range iris.AllAttributes() {
		s, err := json.Marshal(a)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(s))
	}
}

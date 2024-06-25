package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
)

func main() {
	// Definicja atrybutów: pierwszy to liczba zmiennoprzecinkowa (Ocena), drugi to atrybut kategoryczny (Przedmiot)
	attrs := make([]base.Attribute, 3)
	attrs[0] = base.NewFloatAttribute("X")
	attrs[1] = base.NewFloatAttribute("Y")
	attrs[2] = base.NewCategoricalAttribute()
	attrs[2].SetName("Klasa")

	// Utworzenie nowego zestawu danych ze zdefiniowanymi atrybutami
	newInst := base.NewDenseInstances()

	// Zdefiniowanie specyfikacji atrybutów dla zestawu danych
	newSpecs := make([]base.AttributeSpec, len(attrs))
	for i, a := range attrs {
		newSpecs[i] = newInst.AddAttribute(a)
	}

	// Rozszerzenie zestawu danych o dwa wiersze
	newInst.Extend(3)

	// Ustawienie wartości dla pierwszego wiersza
	newInst.Set(newSpecs[0], 0, newSpecs[0].GetAttribute().GetSysValFromString("0.5"))
	newInst.Set(newSpecs[1], 0, newSpecs[1].GetAttribute().GetSysValFromString("1.5"))
	newInst.Set(newSpecs[2], 0, newSpecs[2].GetAttribute().GetSysValFromString("A"))

	// Ustawienie wartości dla drugiego wiersza
	newInst.Set(newSpecs[0], 1, newSpecs[0].GetAttribute().GetSysValFromString("3.0"))
	newInst.Set(newSpecs[1], 1, newSpecs[1].GetAttribute().GetSysValFromString("2.0"))
	newInst.Set(newSpecs[2], 1, newSpecs[2].GetAttribute().GetSysValFromString("B"))

	// Wyświetlenie zawartości zestawu danych
	fmt.Println(newInst)
}

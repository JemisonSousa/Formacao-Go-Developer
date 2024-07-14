//Funçãop é também é chamada de procedimento ou sub-rotina

package main

import "fmt"

func media(lista []float64) float64 {
	total := 0.0
	for _, valor := range lista {
		total += valor
	}
	quatidade_itens := len(lista)
	media := total / float64(quatidade_itens)
	return media
}

func main() {
	lista := []float64{98, 93, 77, 82, 83}
	fmt.Println(media(lista))
}

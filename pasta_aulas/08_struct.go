package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	cidada    string
	idade     int
}

func main__() {

	fmt.Println(pessoa{"Ana", "Goular", "BSB", 25})
	fmt.Println(pessoa{"Julia", "Avart", "RIO", 23})

}

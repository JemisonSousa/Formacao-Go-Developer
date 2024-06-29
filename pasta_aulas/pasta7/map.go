package main

import (
	"fmt"
)

func main() {

	x := make(map[string]int)
	x["chave"] = 10

	fmt.Println("chave")

	elementos := make(map[string]string)
	elementos["H"] = "Hidrogênio"
	elementos["He"] = "Hélio"
	elementos["Li"] = "Lítio"

	fmt.Println(elementos)
	fmt.Println(elementos["H"])

}

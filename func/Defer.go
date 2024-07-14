//Denfer: escalona nossas funções

package main

import "fmt"

func dia1() {
	fmt.Println("Domingo")
}

func dia2() {
	fmt.Println("Segunda-feira")
}

func main() {
	//sistema de execução do Denfer é LIFO, último que entrou na lista de espera é o primeiro a sair.
	dia1()
	dia2()
}

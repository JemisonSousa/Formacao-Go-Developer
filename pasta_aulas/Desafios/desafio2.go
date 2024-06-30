package main

import "fmt"

func main() {

	/*Desafio:
	Você e seus colegas querem criar um código
	que exiba todos os números entre 1 e 100 que
	são divisíveis pode 3.*/
	fmt.Print("Lista dos números que são divis[iveis por 3: \n")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print(i, ", ")
		}

	}

	/*Você e seus colegas querem criar um formato de código de uma antiga
	brincadeira: ao se falar os números multiplos de 3, o participante
	deve falar 'PIN' e nos multiplos de 5 o participante deve falar 'PAN'.
	Então, seu programa deve imprimir números de 1 a 100 e nos citados, não deve
	aparecer os números, mas sim o que o que o participante deve falar*/
	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "PIN-PAN")
		}
		if i%3 == 0 {
			fmt.Println(i, "Pin")
		}
		if i%5 == 0 {
			fmt.Println(i, "Pan")
		}
	}

}

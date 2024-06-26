package main

import "fmt"

const constante_de_Kelvin float64 = 373.15

func main(){

	valor_em_Kelvin := 459.89
	resutado_em_C := valor_em_Kelvin - constante_de_Kelvin
	fmt.Printf("O valor de %.2fK, convertido em Graus é %.2f°", valor_em_Kelvin, resutado_em_C)

}
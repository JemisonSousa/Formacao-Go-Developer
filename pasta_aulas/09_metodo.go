package main

import "fmt"

type retangulo struct {
	comprimento int
	altura      int
}

// Este método 'area' posssui um tipo retangulo
func (r *retangulo) area() int {
	return r.comprimento * r.altura
}

// Métodos podem ser definidos por qualquer tipo de receptor
// ponteitos ou valor. Aqui temos um exemplo do receptor de valor
func (k *retangulo) perimetro() int {
	return 2*k.comprimento + 2*k.altura
}

func main() {
	r := retangulo{comprimento: 10, altura: 5}

	//Aqui chamamos os 2 métodos definidos para a nossa estrutura
	fmt.Println("Área:", r.area())
	fmt.Println("Área:", r.area())
}

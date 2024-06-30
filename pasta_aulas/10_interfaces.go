package main

import (
	"fmt"
	"math" //utiizar o PI
)

type geometria interface {
	area() float64
	perimetro() float64
}

// Para nossa aula vamos usar essa interface nos
// tipos QUADRADOS e CÍRCULOS.
// área de um qudrado: lado² ou lado*lado
// perímetro = soma dos lados/
type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

// A implementação dos `quadrado`
func (q *quadrado) area() float64 {
	return q.lado * q.lado
}
func (p *quadrado) perimetro() float64 {
	return 4 * p.lado
}

// A implementação dos `circulos`
func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}
func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

// Se a variável tem um tipo interface, então nós podemos chamar
// métodos que estão na interface nomeada. Aqui temos uma
// função genérica "medida" tomando vantegam desse
// trabalho em qualquer "geometria"
func medir(g geometria) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())
}

func main() {
	v_q := &quadrado{lado: 10}
	v_c := &circulo{raio: 5}

	// Os tipos de estruturas 'circulo' e 'quadrado' ambos
	// implementam a interface 'geometria' então nós podemos usar
	// instâncias dessas estruturas comom argumentos para 'medir'

	medir(v_q)
	medir(v_c)

}

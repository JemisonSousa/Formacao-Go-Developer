// Closure: função "chamar" uma variável que está em outra função

package main

import "fmt"

func main() {
	x := 0

	numero := func() int {
		x++ //incrementa somar 1
		return x
	}

	fmt.Println(numero())
	fmt.Println(numero())
	fmt.Println(numero())
}

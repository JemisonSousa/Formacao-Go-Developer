//Ponteiro (ptr) armazenar um valor na memória, mas não é o valor propriamente escrito

package main

import "fmt"

func inicial(xPtr *int) {
	*xPtr = 2
}

func main() {
	x := 5
	inicial(&x)
	fmt.Println(x)
}

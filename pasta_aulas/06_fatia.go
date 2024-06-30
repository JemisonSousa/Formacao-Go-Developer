package main

import "fmt"

func mudar____() {

	// x := make([]float64, 5)
	// fmt.Println(x)

	arr := [7]float64{1, 2, 3, 4, 5, 6, 7}
	fatia := arr[0:3]
	fatia2 := arr[0:6]
	fatia3 := arr[3:7]
	fmt.Println(fatia)
	fmt.Println(fatia2)
	fmt.Println(fatia3)

	fatia4 := append(fatia, 10, 69)
	fmt.Println(fatia4)

	fatia5 := []int{1, 2, 3}
	fatia6 := make([]int, 4)
	fmt.Println(fatia6)
	copy(fatia6, fatia5)
	fmt.Println(fatia5, fatia6)

}

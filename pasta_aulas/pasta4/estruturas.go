package main

import "fmt"

func main(){
	
	// instrução FOR   <- Considerado o WHILE, pois não há no GO
	i := 0
	 for i <= 10{
		fmt.Print(i, ", ")
		i = i + 1
	 }

	 // outra forma de FOR
	 for i:=0; i<=10; i++{
		fmt.Println(i)
	 }

	 // CONDIÇÕES (SE)
	 
	 for y:=0; y<=10; y++ {
		if y%2 == 0 {
			fmt.Println("Número par: ",y)
			} else {
			// fmt.Println("Número impar: ",y)
		}
	 }
	 
	 for y:=0; y<=5; y++ {
		switch y {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("Um")
		case 2:
			fmt.Println("Dois")
		case 3:
			fmt.Println("Três")
		case 4:	
			fmt.Println("Quatro")
		case 5:	
			fmt.Println("Cinco")
		}
	 }

	 a := 1
	 b := 20
	 for a <= b {
		 fmt.Println(a)
		 if a == 5{
			fmt.Println("Parei, achei o", a)
			break
		 }
		a++
	}
	
	//  c := 1
	//  d := 20
	//  for c <= d {
	// 	 fmt.Println(c)
	// 	 if c == 5{
	// 		fmt.Println("Achei o ", c, "e continuei sem imprimí-lo")
	// 		continue
	// 	 }
	// 	c++
	//  }
}
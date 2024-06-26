package main

import "fmt"

func main(){
	
	//instrução FOR
	i := 0
	 for i <= 10{
		fmt.Print(i, ", ")
		i = i + 1
	 }

	 //outra forma de FOR
	 for i:=0; i<=10; i++{
		fmt.Println(i)
	 }

	 //CONDIÇÕES (SE)
	 
	 for y:=0; y<=10; y++ {
		if y%2 == 0 {
			fmt.Println("Número par: ",y)
			} else {
			// fmt.Println("Número impar: ",y)
		}
	 }

}
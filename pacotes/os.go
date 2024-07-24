package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("/workspaces/Formacao-Go-Developer/Como_Executar_Go.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
}

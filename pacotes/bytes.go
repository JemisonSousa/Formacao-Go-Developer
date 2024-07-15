package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("%s", bytes.Title([]byte("Pacote bytes")))
}

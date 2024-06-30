package main

import "fmt"

// Declaração da variável CONST do ponto de ebulição da água em F.
const abulicaoF = 212.0
func mudar_____xxxxx() {
	fmt.Println()
	//var tempF = abulicaoF
	//var tempC = (tempF - 32) * 5 / 9
	tempF := abulicaoF
	tempC := (tempF - 32) * 5 / 9

	fmt.Println("A temperatura de ebulição da água em °F =", tempF)
	fmt.Println("A temperatura de ebulição da água em °C =", tempC)

	fmt.Printf("A temperatura de ebulição da água em °F = %.2f (%T). E a temperatura de ebulição da água em °C = %.2f (%T)", tempF, tempF, tempC, tempC)
	fmt.Println()
	fmt.Println()

	//fmt.Sprintln("Aqui é um teste %v", abulicaoF)
}

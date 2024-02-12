package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/calculadora/operadores"
	//"github.com/Luis/calculadora/operadores"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("La Calculadora")
	scanner.Scan()
	operacion := scanner.Text()
	// Borrarlo - fmt.Println(operacion)

	resultado := 0
	if strings.Contains(operacion, "+") {
		resultado = operadores.Sumar(operacion)
	} else if strings.Contains(operacion, "-") {
		resultado = operadores.Restar(operacion)
	} else if strings.Contains(operacion, "*") {
		resultado = operadores.Multiplicar(operacion)
	} else if strings.Contains(operacion, "/") {
		resultado = operadores.Division(operacion)
	} else {
		fmt.Println("Error: El operador esta mal ingresado")
		fmt.Println("o ¡¡Solo debes realizar con un operador!!")
		fmt.Println("o ¡¡Este operador no esta implementado!!")
	}

	//borrar estas tres lineas
	//valores := strings.Split(operacion, "+")
	//operador1, _ := strconv.Atoi(valores[0])
	//operador2, _ := strconv.Atoi(valores[1])

	//fmt.Println(operador1 + operador2)
	fmt.Println(resultado)
}

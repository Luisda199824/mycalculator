package mycalculator

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calculator struct {
	operador  string
	operador1 int
	operador2 int
}

func main() {
	operacion := readEntry()

	operador := detectOperador(operacion)
	valores := strings.Split(operacion, operador)

	// Cast valores from text to number
	operador1 := strToInt(valores[0])
	operador2 := strToInt(valores[1])

	calc := calculator{operador, operador1, operador2}

	fmt.Println("Resultado: ", calc.operar(
		calc.operador,
		calc.operador1,
		calc.operador2,
	))
}

func readEntry() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese la operacion (suma, de la forma numero + numero, ej: 2+2): ")
	scanner.Scan()

	return scanner.Text()
}

func strToInt(value string) int {
	intValue, errValue := strconv.Atoi(value)
	if errValue != nil {
		fmt.Println("Operador 2 no es válido")
		panic(errValue)
	}

	return intValue
}

func (calculator) operar(operador string, operador1 int, operador2 int) float64 {
	var result float64

	switch operador {
	case "+":
		result = float64(operador1 + operador2)
	case "-":
		result = float64(operador1 - operador2)
	case "*":
		result = float64(operador1 * operador2)
	case "/":
		if operador2 == 0 {
			result = 0.0
		} else {
			result = float64(operador1 / operador2)
		}
	default:
		fmt.Println("Operación no válida")
		result = 0
	}

	return result
}

func detectOperador(operation string) string {
	if strings.Contains(operation, "+") {
		return "+"
	} else if strings.Contains(operation, "-") {
		return "-"
	} else if strings.Contains(operation, "*") {
		return "*"
	} else if strings.Contains(operation, "/") {
		return "/"
	}

	panic(errors.New("No valid operation found"))
}

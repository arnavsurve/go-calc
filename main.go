package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	var number1, number2 float64
	var operator string

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&number1)
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&number2)
	fmt.Println("Enter the operator:")
	fmt.Println("+ for addition\n- for subtraction\n* for multiplication\n/ for division\n^ for power\n% for modulus")
	fmt.Scanln(&operator)

	var sum float64 = calculate(operator, number1, number2)
	// trim trailing zeros
	formattedSum := strconv.FormatFloat(sum, 'f', -1, 64)
	formattedNum1 := strconv.FormatFloat(number1, 'f', -1, 64)
	formattedNum2 := strconv.FormatFloat(number2, 'f', -1, 64)

	fmt.Printf("%s "+operator+" %s = %s\n", formattedNum1, formattedNum2, formattedSum)

}

func calculate(operator string, num1 float64, num2 float64) float64 {
	if operator == "+" {
		return num1 + num2
	} else if operator == "-" {
		return num1 - num2
	} else if operator == "*" {
		return num1 * num2
	} else if operator == "/" && num2 != 0 {
		return num1 / num2
	} else if operator == "/" && num2 == 0 {
		fmt.Println("division by 0 error")
		return 0
	} else if operator == "^" {
		return math.Pow(num1, num2)
	} else if operator == "%" {
		return math.Mod(num1, num2)
	} else {
		fmt.Println("invalid operator")
		return 0
	}
}

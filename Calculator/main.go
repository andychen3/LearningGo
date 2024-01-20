package main

import "fmt"

func main() {
	var operation string
	var number1, number2 int
	fmt.Println("Calculator")
	fmt.Println("==================")
	fmt.Println("Which operation do you want to perform? (add, subtract, divide, or multiply?)")
	fmt.Scanf("%s", &operation)
	fmt.Println("Enter the first number")
	fmt.Scanf("%d", &number1)
	fmt.Println("Enter the second number")
	fmt.Scanf("%d", &number2)
	switch operation {
	case "add":
		fmt.Println(number1 + number2)
	case "subtract":
		fmt.Println(number1 - number2)
	case "multiply":
		fmt.Println(number1 * number2)
	case "divide":
		fmt.Println(number1 / number2)
	}
}

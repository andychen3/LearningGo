package main

import "fmt"

func birthday(age *int) {
	*age++
	print(*age)
}

func main() {

	var message string = "hello"
	price := 34.4
	fmt.Println(message, price)
	age := 1
	birthday(&age)

}

package main

import (
	"fmt"
	"strconv"

	"computator/fibonacci"
)

func main() {
	var input string
	fmt.Println("This program calculates a number from Fibonacci sequence.")
	fmt.Println("Enter \"stop\" to interrupt the program")
	for {
		fmt.Print("Enter a number: ")
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			fmt.Printf("Failed to parse input: %v", err)
		} else if input == "stop" {
			fmt.Println("Exiting")
			return
		} else if n, err := strconv.ParseUint(input, 0, 64); err == nil {
			fmt.Println(fibonacci.Fibonacci(uint(n)))
		} else if n, err := strconv.ParseInt(input, 0, 64); err == nil {
			fmt.Println(fibonacci.Fibonacci(int(n)))
		} else if n, err := strconv.ParseFloat(input, 64); err == nil {
			fmt.Println(fibonacci.Fibonacci(float64(n)))
		} else {
			fmt.Println("Error: non-numeric value was entered!")
		}
	}
}

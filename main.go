package main

import (
	"errors"
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	fmt.Println("My Calculator App")
	fmt.Println("Add(2, 3) =", Add(2, 3))
	fmt.Println("Subtract(5, 2) =", Subtract(5, 2))
	fmt.Println("Multiply(3, 4) =", Multiply(3, 4))
	result, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Divide(10, 2) error:", err)
	} else {
		fmt.Println("Divide(10, 2) =", result)
	}
}

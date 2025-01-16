package main

import (
	"fmt"
	"./circle"
	"./factorial"
	"./arithmetic"
)

func main() {
	fmt.Println("Hello, World!")

	// Circle calculations
	radius := 5.0
	circleArea := circle.CalculateArea(radius)
	fmt.Printf("The area of a circle with radius %f is %.2f\n", radius, circleArea)

	// Factorial calculation
	num := 5
	factorialResult := factorial.Calculate(num)
	fmt.Printf("The factorial of %d is %d\n", num, factorialResult)

	// Arithmetic operation
	a := 10
	b := 3
	result := arithmetic.Add(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, result)
}

package main

import "fmt"

func main() {
    // Call functions from other files directly
    radius := 5.0
    area := CalculateCircleArea(radius)
    fmt.Printf("The area of a circle with radius %f is %.2f\n", radius, area)

    num := 5
    factorialResult := Factorial(num)
    fmt.Printf("The factorial of %d is %d\n", num, factorialResult)

    a := 10
    b := 3
    sum := Add(a, b)
    fmt.Printf("%d + %d = %d\n", a, b, sum)
}

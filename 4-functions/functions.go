package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func matematicalCalculation(a, b int) (int, int, int, int) {
	sum1 := a + b
	sub := a - b
	mul := a * b
	div := a / b
	return sum1, sub, mul, div
}

func main() {
	a := 5
	b := 10
	result := sum(a, b)
	fmt.Println("The sum is:", result)

	f := func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultfunction := f("text to function")
	fmt.Println(resultfunction)

	sum1, sub, mul, div := matematicalCalculation(a, b)
	fmt.Println("The sum is:", sum1)
	fmt.Println("The substraction is:", sub)
	fmt.Println("The multiplication is:", mul)
	fmt.Println("The division is:", div)

	sum2, sub1, _, _ := matematicalCalculation(20, 10)
	fmt.Println("The sum is:", sum2)
	fmt.Println("The substraction is:", sub1)
}

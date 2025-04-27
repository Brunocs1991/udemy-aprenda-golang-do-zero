package main

import "fmt"

func main() {
	slice := make([]float32, 10, 15)
	fmt.Println(slice)

	slice = append(slice, 1.0, 2.0, 3.0, 4.0, 5.0)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 6.0)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice2 := make([]float32, 5)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	slice2 = append(slice2, 1.0)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

}

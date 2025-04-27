package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Declare and initialize an array of integers
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)

	// Declare and initialize a slice of integers
	slice := []int{6, 7, 8, 9, 10}
	fmt.Println("Slice:", slice)

	// Append to the slice
	slice = append(slice, 11)
	fmt.Println("Slice after append:", slice)

	// Slice the array
	subSlice := arr[1:4] // This will create a slice from index 1 to 3 (4 is exclusive)
	fmt.Println("Sub-slice of array:", subSlice)

	fmt.Println(reflect.TypeOf(arr))
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(subSlice))
}
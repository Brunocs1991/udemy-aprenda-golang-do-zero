package main

import "fmt"

var n int

func init() {
	fmt.Println("Function init started")
	n = 10
}
func main(){
	fmt.Println("Function main started")
	fmt.Println("The value of n is: ", n)
}
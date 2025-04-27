package main

import "fmt"

func main(){
	txt := func(txt string)string {
		return fmt.Sprintf("Received %s", txt)
	}("Anonimous function") // Call the anonymous function immediat
	fmt.Println(txt)
}
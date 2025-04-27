package main
import  "fmt"
func main() {
	// if statement
	number := -5
	if number > 15 {
		fmt.Println("numer is greater than 15")
	} else{
		fmt.Println("numer is less than or equal to 15")
	}

	if otherNumber := number; otherNumber >0 {
		fmt.Println("otherNumber is greater than 0")
	} else if otherNumber < -10 {
		fmt.Println("otherNumber is less than -10")
	} else {
		fmt.Println("otherNumber is between -10 and 0")
	} 
}
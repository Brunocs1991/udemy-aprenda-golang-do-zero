package main

import "fmt"

func dayOfWeek(number int) string {
	var dayOfWeek string
	switch number {
	case 1:
		dayOfWeek = "Sunday"
		fallthrough
	case 2:
		dayOfWeek = "Monday"
	case 3:
		dayOfWeek = "Tuesday"
	case 4:
		dayOfWeek = "Wednesday"
	case 5:
		dayOfWeek = "Thursday"
	case 6:
		dayOfWeek = "Friday"
	case 7:
		dayOfWeek = "Saturday"
	default:
		dayOfWeek = "invalid day of week"
	}
	return dayOfWeek
}

func dayOfWeek2(number int) string {
	switch {
	case number == 1:
		return "Sunday"
	case number == 2:
		return "Monday"
	case number == 3:
		return "Tuesday"
	case number == 5:
		return "Thursday"
	case number == 6:
		return "Friday"
	case number == 7:
		return "Saturday"
	default:
		return "invalid day of week"
	}
}

func main() {
	fmt.Println("Switch Example")
	day := dayOfWeek(3)
	fmt.Println("Day of the week is:", day)
	day = dayOfWeek(5)
	fmt.Println("Day of the week is:", day)
	day = dayOfWeek(8)
	fmt.Println("Day of the week is:", day)
	day = dayOfWeek2(3)
	fmt.Println("Day of the week is:", day)
	day = dayOfWeek(1)
	fmt.Println("Day of the week is:", day)
}

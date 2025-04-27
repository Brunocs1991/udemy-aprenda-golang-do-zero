package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")
	// Creating a map
	user := map[string]string{
		"first_name": "John",
		"last_name":  "Doe",
		"age":       "30",
	}
	fmt.Println(user)
	fmt.Println(user["first_name"])
	fmt.Println(user["last_name"])
	fmt.Println(user["age"])

	user2 := map[string]map[string]string{
		"name": {
			"first_name": "John",
			"last_name":  "Doe",
		},
		"course": {
			"course_name": "Go",
			"duration":    "2 months",
		},
	}
	fmt.Println(user2)
	fmt.Println(user2["name"]["first_name"])
	fmt.Println(user2["name"]["last_name"])
	fmt.Println(user2["course"]["course_name"])
	fmt.Println(user2["course"]["duration"])

	delete(user2, "name")
	fmt.Println(user2)

	user2["name"] = map[string]string{
		"first_name": "Jane",
		"last_name":  "Smith",
	}
	fmt.Println(user2)

}
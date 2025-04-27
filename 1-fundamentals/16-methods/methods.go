package main

import "fmt"

type user struct {
	name string
	age  int
}

func (u user) salve() {
	fmt.Printf("Save the user %s in database\n", u.name)
}

func (u user) ofAge() bool {
	return u.age >= 18
}

func (u *user) birthday() {
	u.age++
	fmt.Printf("Happy birthday %s, you are now %d years old\n", u.name, u.age)
}

func main() {
	firstUser := user{name: "John", age: 17}
	fmt.Println(firstUser)
	firstUser.salve()
	fmt.Println("Is the user of age?", firstUser.ofAge())
	firstUser.birthday()
	fmt.Println(firstUser)
	fmt.Println("Is the user of age?", firstUser.ofAge())
	fmt.Println("\n======================================")
	secondUser := user{name: "Jane", age: 25}
	fmt.Println(secondUser)
	secondUser.salve()
	fmt.Println("Is the user of age?", secondUser.ofAge())
	secondUser.birthday()
	fmt.Println(secondUser)
	fmt.Println("Is the user of age?", secondUser.ofAge())
}

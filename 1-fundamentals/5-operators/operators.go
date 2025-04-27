package main

func main() {
	// Arithmetic Operators
	a := 10
	b := 20

	sum := a + b
	diff := a - b
	prod := a * b
	quot := a / b
	mod := a % b

	// Comparison Operators
	isEqual := a == b
	isNotEqual := a != b
	isGreater := a > b
	isLess := a < b
	isGreaterOrEqual := a >= b
	isLessOrEqual := a <= b

	// Logical Operators
	andResult := (a > 5) && (b < 30)
	orResult := (a < 5) || (b > 15)
	notResult := !(a == 10)

	// Bitwise Operators
	andBitwise := a & b
	orBitwise := a | b
	xorBitwise := a ^ b
	leftShiftBitwise := a << 1
	rightShiftBitwise := a >> 1

	// Assignment Operators
	c := 5 // Assigns value of 5 to c
	c += 2 // c = c + 2, now c is 7
	c -= 3 // c = c - 3, now c is 4

	// Ternary Operator (using if-else as Go doesn't have it)
	var max int
	if a > b {
		max = a
	} else {
		max = b
	}

	println("Arithmetic Results: ", sum, diff, prod, quot, mod)
	println("Comparison Results: ", isEqual, isNotEqual, isGreater, isLess, isGreaterOrEqual, isLessOrEqual)
	println("Logical Results: ", andResult, orResult, notResult)
	println("Bitwise Results: ", andBitwise, orBitwise, xorBitwise, leftShiftBitwise, rightShiftBitwise)
	println("Assignment Result: ", c)
	println("Max Value: ", max)
}
package main

import (
	"errors"
	"fmt"
)

func main() {
	// INIT INTEGERS NUMBERS
	// The int8 type is an 8-bit signed integer, which means it can hold values from -128 to 127.
	// If you try to assign a value outside this range, it will result in an overflow.
	var number int16 = 100
	fmt.Println(number)

	// The int type is platform-dependent, meaning it can be either 32-bit or 64-bit depending on the architecture of the machine.
	// On a 32-bit machine, int is 32 bits, and on a 64-bit machine, int is 64 bits.
	var number2 int = 1000000000
	fmt.Println(number2)

	// The uint type is an unsigned integer, which means it can only hold non-negative values.
	// The range of uint is from 0 to 2^32-1 on a 32-bit machine and from 0 to 2^64-1 on a 64-bit machine.
	var number3 uint = 1000000000
	fmt.Println(number3)

	// alias
	// The rune type is an alias for int32 and is used to represent a Unicode code point.
	// A rune can hold any character in the Unicode standard, which includes characters from various languages and symbols.
	var number4 rune = 123456
	fmt.Println(number4)

	// The byte type is an alias for uint8 and is used to represent a single byte of data.
	var number6 byte = 123
	fmt.Println(number6)

	// END iNTEGERS NUMBERS
	// INIT FLOATING NUMBERS

	// The float32 type is a 32-bit floating-point number, which means it can hold decimal values with a certain level of precision.
	var realNumber float32 = 123.456
	fmt.Println(realNumber)

	// The float64 type is a 64-bit floating-point number, which means it can hold decimal values with a higher level of precision than float32.
	var realNumber2 float64 = 123.4567890123456789
	fmt.Println(realNumber2)

	realNumber3 := 123.44
	fmt.Println(realNumber3)

	// END FLOATING NUMBERS
	// INIT STRINGS

	var str string = "Hello, World!"
	fmt.Println(str)

	str2 := "Hello, World!"
	fmt.Println(str2)

	// The string type in Go is a sequence of bytes, and it can hold any UTF-8 encoded text.
	character := '¨'
	fmt.Println(character)

	// END STRINGS

	// INIT DEFAULT VALUES
	// In Go, when you declare a variable without initializing it, it gets a default value based on its type.
	var text string
	fmt.Println(text)

	var int int16
	fmt.Println(int)

	var float float32
	fmt.Println(float)

	var boolean bool
	fmt.Println(boolean)

	var error1 error
	fmt.Println(error1)

	// END DEFAULT VALUES
	// INIT BOOLEAN VALUES
	var booleanValue bool = true
	fmt.Println(booleanValue)

	booleanValue2 := false
	fmt.Println(booleanValue2)

	// END BOOLEAN VALUES

	// INIT ERRORS
	var error2 error
	fmt.Println(error2)

	var error3 error = errors.New("this is an error")
	fmt.Println(error3)
}

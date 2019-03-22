package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	type any struct{}

	// Using var
	var name = "Bala"
	var age int32 = 20
	const isCool = true
	var size float32 = 2.3

	// Shorthand
	// name := "Brad"
	// email := "brad@gmail.com"

	username, email, password := "Brad", "brad@gmail.com", "secret"

	fmt.Println(name, age, isCool, email, username, password)
	fmt.Printf("%T\n", size)
}

package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int int8   int16  int32  int64
	// uint uint8  uint16  uint32  uint64  (unsigned int - only positive numbers or 0)
	// byte - alias for uint8
	// rune - alias for int32
	// float32  float64
	// complex64 complex128

	//  USING VAR:
	var name string = "Brad"
	var age int32 = 37
	const isCool = true

	//Shorthand
	name2 := "Stevia"
	// shorter way to sign multiple vars
	name3, email := "Mikica", "email@gmail.com"

	fmt.Println(name, age, isCool, name2)
	fmt.Printf("%T\n", age)
	fmt.Println("====>", name3, email)
}
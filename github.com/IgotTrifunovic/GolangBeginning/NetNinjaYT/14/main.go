package main

import "fmt"

func updateName(x *string) {
	*x = "wedge"
}

func main() {
	// group A types -> strings, ints, bools, floats, arrays, structs
	name := "tifa"

	// name = updateName(name)
	// fmt.Println(name)
	m := &name
	// fmt.Println("memory adress:", m)
	// fmt.Println("value at memmory address:", *m)

	fmt.Println(name)
	updateName(m)
	fmt.Println(name)
	fmt.Println(name)
}

//  maybe demaged code

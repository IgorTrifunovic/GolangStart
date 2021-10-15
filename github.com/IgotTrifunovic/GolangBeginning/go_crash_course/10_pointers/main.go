package main

import "fmt"

func main() {
	a := 5
	b := &a // &(ampersand)"a" is assigning "b" to be pointer of "a"

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// Use * to read val from address
	fmt.Println(*b) // *b  printuje vrednost od b, tj na sta pointer ukazuje
	//fmt.Println(*&a)
	// isti nacin da uzmemo vrednost pointera "a"

	// Change val with pointer
	*b = 10					// * (asteriSC. not asterix)
	fmt.Println(a)

}										// POINTERS are used to increase performance

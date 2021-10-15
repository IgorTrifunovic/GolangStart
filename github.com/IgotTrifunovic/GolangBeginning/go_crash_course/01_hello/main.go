package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// format specifiers from tuturial on YT
	age := 23
	name := "Zolika"
	score := 3.456133
	fmt.Println("=================================")
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T and name is of type %T \n", age, name)
	fmt.Printf("you scored %0.3f points!\n", score)
	fmt.Printf("you scored %0.5f points!", 235.34947658656383956555555)

}

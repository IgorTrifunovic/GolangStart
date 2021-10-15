package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if x > y {
		fmt.Printf("%d is less than %d\n", y, x)
	} else {
		fmt.Printf("%d is equal to %d\n", y, x)
	}

	// else if

	color := "asdasd"
	//colorType := reflect.TypeOf(color).Kind()   //ne radi (check if input is string)

	// &&  je and; ||   je or;
	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
		// } else if colorType != reflect.String() {    //ne radi (check if input is string)
		// 	fmt.Println("Please enter a string.")
	} else {
		fmt.Println("color is NOT blue or red")
	}
}

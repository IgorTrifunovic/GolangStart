package main

import "fmt"

var score = 99.5 // THIS VAR MUST BE GLOBAL (not in func main) SO THAT PACKAGE COULD USE IT

func main() {
	sayHello("mario")

	for _, v := range points {
		fmt.Println(v)
	}
	showScore()
}

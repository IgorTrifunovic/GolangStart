package main

import "fmt"

func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func main() {
	fmt.Println("Go Testing Tuturial")
	result := Calculate(2)
	fmt.Println(result)
}

// RUN THESE TESTS WITH "go test" or "go test -v"

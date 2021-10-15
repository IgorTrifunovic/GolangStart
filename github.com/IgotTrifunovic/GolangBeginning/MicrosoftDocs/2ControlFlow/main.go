package main

import (
	"fmt"
	"strconv"
)

func findprimes(number int) bool {
	for i := 2; i < number; i++ {
		if number % i == 0 {
			return false
		}
	}
    if number > 1 {
        return true
    } else {
        return false
	}
}

func fizzbuzz(num int) string {
	switch {
	case num%15 == 0:
		return "FizzBuzz"
	case num%3 == 0:
		return "Fizz"
	case num%5 == 0:
		return "Buzz"
	}
	return strconv.Itoa(num)
}

func main() {
	for num := 1; num <= 100; num++ {
		fmt.Println(fizzbuzz(num))
	}

    fmt.Println("Prime numbers less than 20:")

    for number := 1; number <= 20; number++ {
        if findprimes(number) {
            fmt.Printf("%v ", number)
        }
    }
}
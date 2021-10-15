package main

import (
	"bufio"
	"fmt"
	"os"
)

// func fibonacci(number int)  [100]int {
// 	i, k := 0, 1
// 	returnfunc () int {
// 		a, b = b, a+func Benchmark(b *testing.B) {
// 			for i := 0; i < b.N; i++ {

// 			}
// 		}

func fibonacci() func() int {
	i, k := 0, 1
	return func() int {
		i, k = k, i+k
		return i
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Enter febonacci number: ", reader) 
	switch opt{}
	
}

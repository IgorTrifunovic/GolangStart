package main

import "fmt"

func main() {
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30") // redom ce gledati if izjave dok ne nadje true, nakon cega nece gledati dalje
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is NOT less than 45")
	}
	fmt.Println("==================================================")

	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue // BREAK from current iterattion and continue with loop (from the top)
		}
		if index == 2 {
			fmt.Println("breaking at pos", index)
			break
		}
		fmt.Printf("the value at pos %v is %v \n", index, value)
	}

}

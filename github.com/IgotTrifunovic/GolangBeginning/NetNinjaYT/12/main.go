package main

import "fmt"

func main() {

	menu := map[string]float64{ // all keys must be of the same type as same as the values
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//looping maps
	for k, v := range menu {
		fmt.Println(k, " - ", v)
	}

	//ints as key type
	phonebook := map[int]string{
		3565454: "mario",
		6345821: "luigi",
		9236728: "peach",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[6345821])

	phonebook[9236728] = "bowser"
	fmt.Println(phonebook)

	phonebook[3565454] = "yoshi"
	fmt.Println(phonebook)
}

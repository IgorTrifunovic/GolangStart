package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	// loop trough ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// not using index
	for _, id := range ids { // ako neces koristiti var "i" stavi _
		fmt.Printf("ID: %d\n", id)
	}
	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	//range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com",
		"Moma": "moma@gmail.com", "Petar": "petar@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
	for k := range emails {
		fmt.Println("Name: " + k)
	}
	// vezba
	fmt.Println("=====================")
	for o, n := range emails {	//SA OBZIROM DA JE OVO LOOP TROUGH MAP BIRAMO EL. SA KEY-VALUE (o-n u ovom sl.)
		if o == "Moma" {
			fmt.Println(n)
		}
	}
}

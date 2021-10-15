package main

import (
	"fmt"
	"strconv"
)

// Define person struct  (nesto kao klasa)
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// VALUE RECIEVER je metoda koja ne menja struct, vec kopira struct-ove vrednosti (vise smeca losiji performance)
// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// POINTER RECIEVER je metoda koja direktno koristi/menja struct-ove vrednosti (manje smeca bolji performance)
// Has Birthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

// Init person using struct  (nesto kao objekat)
func main() {

	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 27}

	//Alternative
	person2 := Person{"Bob", "Smith", "Boston", "m", 26}

	person1.firstName = "Maria"

	fmt.Println(person1.firstName, "=============================", person2.firstName)

	person1.age = 29
	person1.firstName = "Samantha"
	fmt.Println(person1.greet())
	fmt.Println("++++++++++++++++++++++++++++++++++++")

	person1.hasBirthday()
	person2.hasBirthday()
	person1.hasBirthday()

	person1.getMarried("Williams")
	person2.getMarried("Tomphson")
	fmt.Println(person1.greet(), "\n", person2.greet())
}

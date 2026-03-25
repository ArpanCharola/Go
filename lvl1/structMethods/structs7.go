package main

import "fmt"

// Define a struct
type Person struct {
	Name  string
	Age   int
	Email string
}

// Method with value receiver
func (p Person) Greet() {
	fmt.Printf("Hi, I'm %s and I'm %d years old\n", p.Name, p.Age)
}

// Method with pointer receiver (can modify)
func (p *Person) HaveBirthday() {
	p.Age++
}

// Constructor-like function
func NewPerson(name string, age int) *Person {
	return &Person{
		Name:  name,
		Age:   age,
		Email: fmt.Sprintf("%s@example.com", name),
	}
}

func main() {
	// Creating structs
	person1 := Person{"Alice", 30, "alice@example.com"}
	person2 := Person{Name: "Bob", Age: 25}
	person3 := NewPerson("Charlie", 35)

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)

	// Accessing fields
	fmt.Println(person1.Name)

	// Using methods
	person1.Greet()
	person1.HaveBirthday()
	person1.Greet()
}

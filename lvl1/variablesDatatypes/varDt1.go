package main

import "fmt"

func main() {
	//variables declaration
	var name string = "Ajay"
	var age int = 26
	var height float64 = 5.8
	var isStudent bool = false

	// short declaration
	country := "India"
	weight := 67.9

	// multiple variables declaration
	var x, y int = 10, 23
	a, b := 5.5, "Learning Go"

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Is Student:", isStudent)
	fmt.Println("Country:", country)
	fmt.Println("Weight:", weight)
	fmt.Println("x:", x, "y:", y)
	fmt.Println("a:", a, "b:", b)
}

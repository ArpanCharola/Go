package main

import "fmt"

func main() {
	// Creating maps
	var scores map[string]int
	scores = make(map[string]int)

	// Adding values
	scores["Alice"] = 95
	scores["Bob"] = 87

	// Map literal
	grades := map[string]string{
		"Alice":   "A",
		"Bob":     "B",
		"Charlie": "C",
	}

	fmt.Println("Scores:", scores)
	fmt.Println("Grades:", grades)

	// Accessing values
	fmt.Println("Alice's grade:", grades["Alice"])

	// Check if key exists
	value, exists := grades["David"]
	if exists {
		fmt.Println("David's grade:", value)
	} else {
		fmt.Println("David not found")
	}

	// Delete from map
	delete(grades, "Charlie")
	fmt.Println("After deletion:", grades)

	// Iterate over map
	for name, grade := range grades {
		fmt.Printf("%s: %s\n", name, grade)
	}
}

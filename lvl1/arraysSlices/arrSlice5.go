package main

import "fmt"

func main() {
	// Arrays (fixed size)
	var arr1 [3]int = [3]int{1, 2, 3}
	arr2 := [4]string{"a", "b", "c", "d"}

	fmt.Println("Array1:", arr1)
	fmt.Println("Array2:", arr2)

	// Slices (dynamic size)
	var slice1 []int = []int{1, 2, 3}
	slice2 := make([]string, 3) // length 3, capacity 3
	slice2[0] = "hello"
	slice2 = append(slice2, "world") // grows dynamically

	fmt.Println("Slice1:", slice1)
	fmt.Println("Slice2:", slice2)

	// Slice operations
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Original:", numbers)
	fmt.Println("First three:", numbers[:3])
	fmt.Println("Last two:", numbers[3:])
	fmt.Println("Middle:", numbers[1:4])

	// Modifying slices
	numbers = append(numbers, 60, 70)
	fmt.Println("After append:", numbers)

	// Multi-dimensional slice
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Matrix:", matrix)
}

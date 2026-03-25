package main

import "fmt"

//basic func
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

//func with return value
func add(x int, y int) int {
	return x + y
}

//multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("number can't be divided by 0")
	}
	return a / b, nil
}

// named return values
func getCoordinates() (z, y int) {
	z = 10
	y = 20
	return // naked return, returns z and y
}

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	greet("Deepak")

	result := add(5, 7)
	fmt.Println("Sum:", result)

	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Quotient:", quotient)
	}

	z, y := getCoordinates()
	fmt.Printf("Coordinates: (%d, %d)\n", z, y)

	total := sum(1, 2, 3, 4, 5)
	fmt.Println("Total:", total)
}

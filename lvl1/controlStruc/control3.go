package main

import (
	"fmt"
	"time"
)

func main() {
	num := 5

	// basic if-else
	if num > 0 {
		fmt.Println("The number is positive.")
	} else if num < 0 {
		fmt.Println("The number is negative.")
	} else {
		fmt.Println("The number is zero.")
	}

	//if with initialization
	if score := 86; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}

	//for loop (go only has for loop, no while)
	for i := 0; i < 5; i++ {
		fmt.Printf("Count: %d\n", i)
	}
	fmt.Printf("\n")
	// for in while style
	count := 0
	for count < 3 {
		fmt.Println("Count:", count)
		count++
	}
	// Infinite loop with break
	x := 0
	for {
		if x >= 3 {
			break
		}
		fmt.Println("x:", x)
		x++
	}

	//range loop
	fruits := []string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Printf("%d, %s\n", index, fruit)
	}

	//switch case
	day := time.Now().Weekday()
	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend!")
	default:
		fmt.Println("Weekday.")
	}
}

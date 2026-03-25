package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name   string
	Grades []int
}

func (s Student) Average() float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	return float64(sum) / float64(len(s.Grades))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	students := make(map[string]Student)

	for {
		fmt.Println("\n--- Student Grade Manager ---")
		fmt.Println("1. Add student")
		fmt.Println("2. Add grade")
		fmt.Println("3. Show student average")
		fmt.Println("4. List all students")
		fmt.Println("5. Exit")
		fmt.Print("Choose option: ")

		optionStr, _ := reader.ReadString('\n')
		optionStr = strings.TrimSpace(optionStr)
		option, _ := strconv.Atoi(optionStr)

		switch option {
		case 1:
			fmt.Print("Enter student name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			students[name] = Student{Name: name, Grades: []int{}}
			fmt.Printf("Added student: %s\n", name)

		case 2:
			fmt.Print("Enter student name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			if student, exists := students[name]; exists {
				fmt.Print("Enter grade: ")
				gradeStr, _ := reader.ReadString('\n')
				gradeStr = strings.TrimSpace(gradeStr)
				grade, _ := strconv.Atoi(gradeStr)

				student.Grades = append(student.Grades, grade)
				students[name] = student
				fmt.Printf("Added grade %d to %s\n", grade, name)
			} else {
				fmt.Println("Student not found!")
			}

		case 3:
			fmt.Print("Enter student name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			if student, exists := students[name]; exists {
				avg := student.Average()
				fmt.Printf("%s's average: %.2f\n", name, avg)
			} else {
				fmt.Println("Student not found!")
			}

		case 4:
			fmt.Println("\n--- All Students ---")
			for name, student := range students {
				fmt.Printf("%s: %v (Avg: %.2f)\n", name, student.Grades, student.Average())
			}

		case 5:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option!")
		}
	}
}

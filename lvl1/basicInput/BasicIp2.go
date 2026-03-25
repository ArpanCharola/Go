package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BasicIp2() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter your age: ")
	ageStr, _ := reader.ReadString('\n')
	ageStr = strings.TrimSpace(ageStr)
	age, _ := strconv.Atoi(ageStr)

	fmt.Printf("Hello %s, your age is %d\n", name, age)

}
func main() {
	BasicIp2()
}

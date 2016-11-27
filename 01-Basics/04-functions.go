package main

// "Notice that the type comes after the variable name."

import "fmt"

func add(x int, y int) int {
	return y + x
}

func main() {
	fmt.Println(add(3, 2))
}

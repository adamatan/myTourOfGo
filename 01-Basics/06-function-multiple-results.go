package main

// "A function can return any number of results."

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("a", "b")
	fmt.Println(swap("a", "b"))

	s1, s2 := swap("World", "Hello")
	fmt.Println(s1,s2)
}

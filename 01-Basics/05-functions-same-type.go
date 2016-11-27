package main

// "When two or more consecutive named function parameters share a type, you can omit the type from all but the last."

import "fmt"

func mul(y, x int) int {
	return (x * y)
}

func main() {
	fmt.Println(mul(9, 9))
}

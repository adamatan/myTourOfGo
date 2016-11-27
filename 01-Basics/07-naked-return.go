package main

/*
"These names should be used to document the meaning of the return values.
...
A return statement without arguments returns the named return values. This is known as a "naked" return."
*/

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}

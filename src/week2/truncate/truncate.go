package main

import "fmt"

/*
 * Teaches how to read from standard input, how to do simple type conversions and use pointers.
 */
func main() {
	x := new(int64)
	fmt.Println("Type in float number, I print it as int: ")
	fmt.Scan(x)
	fmt.Println("Your float converted to int is: ", *x)
}

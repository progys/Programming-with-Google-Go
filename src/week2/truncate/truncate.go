package main

import "fmt"

/*
* Teaches how to read from standard input, how do simple type conversions and use pointers.
 */
func main() {
	x := new(float64)
	fmt.Println("Type in float number, I will convert it to int: ")
	fmt.Scan(x)
	fmt.Println("Your float converted to int is: ", int64(*x))
}

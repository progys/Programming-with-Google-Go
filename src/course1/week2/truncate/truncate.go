package main

import "fmt"

/*
 * Write a program which prompts the user to enter a floating point number and prints the integer
 * which is a truncated version of the floating point number that was entered.
 * Truncation is the process of removing the digits to the right of the decimal place.
 *
 * Teaches how to read from standard input and use pointers.
 */
func main() {
	x := new(int64)
	fmt.Println("Type in float number, I print it as int: ")
	fmt.Scan(x)
	fmt.Println("Your float converted to int is: ", *x)
}

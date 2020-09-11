package main

import "fmt"

func main() {
	x := new(float64)
	fmt.Println("Type in float number, I will convert it to int: ")
	fmt.Scan(x)
	fmt.Println("Your float converted to int is: ", int64(*x))
}

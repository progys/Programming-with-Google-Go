/*
* Let us assume the following formula for displacement s as a function of time t, acceleration a, initial velocity vo, and initial displacement so.

s =½ a t2 + vot + so

Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement. Then the program should prompt the user to enter a value for time and the program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so. GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity, and initial displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume the following values for acceleration, initial velocity, and initial displacement: a = 10, vo = 2, so = 1. I can use the following statement to call GenDisplaceFn() to generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print the displacement after 5 seconds.

fmt.Println(fn(5))
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var accel float64
	var velocity float64
	var displacement float64

	fmt.Printf("Please enter acceleration value: ")
	fmt.Scan(&accel)

	fmt.Printf("Please enter initial velocity: ")
	fmt.Scan(&velocity)

	fmt.Printf("Please enter initial displacement: ")
	fmt.Scan(&displacement)

	displAtGivenTime := GenDisplaceFn(accel, velocity, displacement)

	for {
		time := ReadTime()
		fmt.Printf("time=%v, displacement=%v\n", time, displAtGivenTime(time))
	}
}

//GenDisplaceFn returns a function to calculate displacement at a given time
func GenDisplaceFn(a, v, d float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v*t + d
	}
}

//ReadTime reads user input as a time
func ReadTime() float64 {
	var input string
	fmt.Print("Enter a time ('x' to exit): ")
	fmt.Scan(&input)
	if input == "X" || input == "x" {
		os.Exit(0)
	}
	time, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Printf("Cannot convert \"%v\" to float", input)
		os.Exit(1)
	}
	return time
}

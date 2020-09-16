package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
* Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
* The program should be written as a loop.
* Before entering the loop, the program should create an empty integer slice of size (length) 3.
* During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
* The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
* The slice must grow in size to accommodate any number of integers which the user decides to enter.
* The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
 */

func swap(i int, integers []int64) {
	integers[i], integers[i-1] = integers[i-1], integers[i]
}

func sortEnd(integers []int64) {
	for i := len(integers) - 1; i >= 1; i-- {
		if integers[i] < integers[i-1] {
			swap(i, integers)
		}
	}
}

func sortFromIndex(integers []int64, start int) {
	for i := start; i < len(integers)-1; i++ {
		if integers[i] > integers[i+1] {
			swap(i+1, integers)
		}
	}
}

func main() {
	integers := make([]int64, 3)

	var input string
	numbersEntered := 0
	for {
		fmt.Println("Please enter an integer: ")
		fmt.Scan(&input)

		if strings.ToLower(input) == "x" {
			fmt.Println("Got 'X' exiting the program.")
			return
		}

		number, error := strconv.ParseInt(input, 10, 64)
		if error != nil {
			fmt.Println("Wrong value, not an integer!")
			continue
		}

		if numbersEntered >= len(integers) {
			integers = append(integers, number)
		} else {
			for i := range integers {
				if integers[i] == 0 {
					integers[i] = number
					sortFromIndex(integers, i)
					break
				}
			}
		}

		numbersEntered++

		sortEnd(integers)

		fmt.Println("Sorted: ", integers)
	}
}

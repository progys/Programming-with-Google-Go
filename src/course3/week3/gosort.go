package main

/*
Write a program to sort an array of integers. The program should partition the array into 4 parts,
each of which is sorted by a different goroutine. Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers.
Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
When sorting is complete, the main goroutine should print the entire sorted list.
*/
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	const partitionSize = 4
	chl := make(chan []int, partitionSize)
	var results []int

	groups := readInputIntoGroups(partitionSize)

	for i := 0; i < len(groups); i++ {
		go sortSlice(groups[i], chl)
		sortedSlice := <-chl
		results = mergeResult(results, sortedSlice)
	}

	fmt.Println("Final sorted:", results)
}

func readInputIntoGroups(groups int) [][]int {
	ints := make([][]int, groups, groups)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a series of integers> ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	inputs := strings.Split(input, " ")

	counter := 0
	for _, value := range inputs {
		parsed, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Wrong input: ", value)
			os.Exit(2)
		}
		currentGroup := counter % groups
		ints[currentGroup] = append(ints[currentGroup], parsed)
		counter++
	}

	return ints
}

func sortSlice(ints []int, c chan []int) {
	fmt.Println("Sorting", ints)
	sort.Ints(ints)
	c <- ints
}

func mergeResult(results []int, sortedSlice []int) []int {
	var temporary []int
	indx1, indx2 := 0, 0
	for {
		if len(sortedSlice) == indx1 {
			temporary = copySlice(indx2, results, temporary)
			break
		}
		if len(results) == indx2 {
			temporary = copySlice(indx1, sortedSlice, temporary)
			break
		}
		if sortedSlice[indx1] < results[indx2] {
			temporary = append(temporary, sortedSlice[indx1])
			indx1++
		} else {
			temporary = append(temporary, results[indx2])
			indx2++
		}
	}
	return temporary
}

func copySlice(indx int, results []int, temporary []int) []int {
	for ; indx < len(results); indx++ {
		temporary = append(temporary, results[indx])
	}
	return temporary
}

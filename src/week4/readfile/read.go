/**
* Write a program which reads information from a file and represents it in a slice of structs.
* Assume that there is a text file which contains a series of names.
* Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.
*
* Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
* Each field will be a string of size 20 (characters).
*
* Your program should prompt the user for the name of the text file.
* Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
* Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file.
* After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
 */
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	type Name struct {
		fname [20]rune
		lname [20]rune
	}
	var filename string
	var names []Name

	fmt.Print("Please enter the filename: ")
	fmt.Scan(&filename)

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Problem opening the file", filename)
	}

	finished := false
	for {
		name := Name{}
		buffer := make([]byte, 1)

		var word string
		for {
			_, err := file.Read(buffer)

			if err != nil {
				if err == io.EOF {
					for i, char := range word {
						name.lname[i] = char
					}
					finished = true
					break
				}
				fmt.Println("Error reading the file", err)
				return
			}

			character := string(buffer)
			if character == " " {
				for i, char := range word {
					name.fname[i] = char
				}
				word = ""
				continue
			}
			if character == "\r" {
				continue
			}
			if character == "\n" {
				for i, char := range word {
					name.lname[i] = char
				}
				break
			}
			word = word + character
		}
		names = append(names, name)
		if finished {
			break
		}
	}

	for idx := range names {
		fmt.Println(string([]rune(names[idx].fname[:])), string([]rune(names[idx].lname[:])))
	}

	file.Close()

}

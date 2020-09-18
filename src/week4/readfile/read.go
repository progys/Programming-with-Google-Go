/**
* Write a program which reads information from a file and represents it in a slice of structs.
* Assume that there is a fixedLengthName file which contains a series of names.
* Each line of the fixedLengthName file has a first name and a last name, in that order, separated by a single space on the line.
*
* Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
* Each field will be a string of size 20 (characters).
*
* Your program should prompt the user for the name of the fixedLengthName file.
* Your program will successively read each line of the fixedLengthName file and create a struct which contains the first and last names found in the file.
* Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file.
* After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
 */
package main

import (
	"fmt"
	"io"
	"os"
)

type fixedLengthName [20]rune

type nameType struct {
	fname fixedLengthName
	lname fixedLengthName
}

func main() {
	var filename string
	var names []nameType

	fmt.Print("Please enter the filename: ")
	fmt.Scan(&filename)

	file := openFile(filename)

	for {
		name, reachedEndOfFile := readLine(file)
		names = append(names, name)

		if reachedEndOfFile {
			break
		}
	}

	for _, name := range names {
		fmt.Println(nameToString(name.fname), nameToString(name.lname))
	}

	file.Close()
}

func nameToString(runes fixedLengthName) string {
	return string([]rune(runes[:]))
}

func readWord(file *os.File) (string, bool, error) {
	var word string
	buffer := make([]byte, 1)
	for {
		_, err := file.Read(buffer)

		if err != nil {
			if err == io.EOF {
				return word, true, err
			}
			fmt.Println("Error reading the file: ", err)
			file.Close()
			os.Exit(3)
		}

		character := string(buffer)

		isSpace := character == " "
		isEndOfLine := character == "\n"
		isCarriegeReturn := character == "\r"

		if isSpace || isEndOfLine {
			return word, isEndOfLine, err
		}

		if !isCarriegeReturn {
			word = word + character
		}
	}
}

func readLine(file *os.File) (nameType, bool) {
	var name nameType
	for {
		word, endOfLine, err := readWord(file)

		switch {
		case err == io.EOF:
			name.lname = stringToName(word)
			return name, true

		case endOfLine:
			name.lname = stringToName(word)
			return name, false
		}

		name.fname = stringToName(word)
	}
}

func stringToName(word string) fixedLengthName {
	var result [20]rune
	for i, char := range word {
		result[i] = char
	}
	return result
}

func openFile(filename string) *os.File {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Can't open the file \"%v\": %v\n", filename, err)
		os.Exit(2)
	}
	return file
}

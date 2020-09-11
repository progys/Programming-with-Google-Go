package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
 * Write a program which prompts the user to enter a string.
 * The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
 * The program should print “Found!” if the entered string starts with the character ‘i’,
 * ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise.
 *
 * The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.
 */
func main() {
	consoleReader := bufio.NewReader(os.Stdin)
	const (
		a = "a"
		n = "n"
		i = "i"
	)

	fmt.Println("Enter some text, I will check if it starts with 'i', ends with 'n' and contains 'a': ")
	text, err := consoleReader.ReadString('\r')

	if err != nil {
		fmt.Printf("Some error happened '%v'. Exiting program.\n", err)
		return
	}

	lowerCaseInput := strings.ToLower(text[0 : len(text)-1])

	switch {
	case !strings.HasPrefix(lowerCaseInput, i):
		fmt.Println("Not Found! does not start with letter 'i'")
	case !strings.HasSuffix(lowerCaseInput, n):
		fmt.Println("Not Found! does not end with 'n'")
	case !strings.Contains(lowerCaseInput, a):
		fmt.Println("Not Found! does not contain letter 'a'")
	default:
		fmt.Println("Found!")
	}
}

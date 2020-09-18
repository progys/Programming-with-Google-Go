package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/**
 * Write a program which prompts the user to first enter a name, and then enter an address.
 * Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
 * Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
 */
func main() {
	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter your name: ")
	name, _ := consoleReader.ReadString('\r')

	fmt.Print("Please enter your address: ")
	address, _ := consoleReader.ReadString('\r')

	userMap := map[string]string{"name": strings.TrimSpace(name), "address": strings.TrimSpace(address)}
	json, err := json.Marshal(userMap)

	if err != nil {
		panic(err)
	}

	fmt.Println("JSON: ", string(json))
}

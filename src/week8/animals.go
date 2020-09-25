/*
* Write a program which allows the user to create a set of animals and to get information about those animals. Each animal has a name and can be either a cow, bird, or snake. With each command, the user can either create a new animal of one of the three types, or the user can request information about an animal that he/she has already created. Each animal has a unique name, defined by the user. Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake. The following table contains the three types of animals and their associated data.
*
* Animal	Food eaten	Locomtion method	Spoken sound
* cow		grass		walk					moo
* bird		worms		fly						peep
* snake		mice		slither					hsss
*
* Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
* Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line.
* Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command
* or a “query” command.
*
* Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”.
* The second string is an arbitrary string which will be the name of the new animal.
* The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
* Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.
*
* Each “query” command must be a single line containing 3 strings. The first string is “query”.
* The second string is the name of the animal.
* The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
* Your program should process each query command by printing out the requested data.
*
* Define an interface type called Animal which describes the methods of an animal.
* Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
* The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak()
* method should print the animal’s spoken sound. Define three types Cow, Bird, and Snake.
* For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake
* all satisfy the Animal interface. When the user creates an animal, create an object of the appropriate type.
* Your program should call the appropriate method when the user issues a query command.
 */
package main

import (
	"fmt"
	"strings"
)

//Animal describes an animal
type Animal interface {
	Speak()
	Eat()
	Move()
}

type animalData struct {
	food       string
	locomotion string
	noise      string
}

func main() {
	animals := map[string]Animal{}
	var command, name, typeOrQuery string

	for {
		fmt.Print("> ")
		fmt.Scanln(&command, &name, &typeOrQuery)

		command := strings.ToLower(command)

		switch command {
		case "newanimal":
			addAnimal(animals, name, typeOrQuery)
		case "query":
			query(animals, name, typeOrQuery)
		case "x":
			fmt.Println("Exiting.")
			return
		default:
			fmt.Println("Unrecognized command. Supported: newanimal or query. Type x to exit.")
		}
		command, name, typeOrQuery = "", "", ""
	}
}

func query(animals map[string]Animal, name string, command string) {
	animal := animals[name]
	if animal == nil {
		fmt.Printf("No such animal with name \"%s\". Create it with \"newanimal\" command first.\n", name)
		return
	}
	switch {
	case command == "eat":
		animal.Eat()
	case command == "move":
		animal.Move()
	case command == "speak":
		animal.Speak()
	default:
		fmt.Println("Unrecognized query for an animal. Supported - eat, move or speak.")
	}
}

func addAnimal(animals map[string]Animal, name string, animalType string) {
	var animal Animal
	switch strings.ToLower(animalType) {
	case "cow":
		animal = Cow{animalData{"grass", "walk", "moo"}}
	case "bird":
		animal = Bird{animalData{"worms", "fly", "peep"}}
	case "snake":
		animal = Snake{animalData{"mice", "slither", "hsss"}}
	default:
		fmt.Println("Unrecognized type of the animal. E.g.: newanimal <name> <cow|bird|snake>")
	}
	if animal != nil {
		animals[name] = animal
		fmt.Println("Created it!")
	}
}

//Cow represents cow
type Cow struct {
	data animalData
}

//Bird represents bird
type Bird struct {
	data animalData
}

//Snake represents snake
type Snake struct {
	data animalData
}

//Eat prints out what cow eats
func (animal Cow) Eat() {
	fmt.Println(animal.data.food)
}

//Move prints out how cow moves
func (animal Cow) Move() {
	fmt.Println(animal.data.locomotion)
}

//Speak prints out what sound cow makes
func (animal Cow) Speak() {
	fmt.Println(animal.data.noise)
}

//Eat prints out what bird eats
func (animal Bird) Eat() {
	fmt.Println(animal.data.food)
}

//Move prints out how bird moves
func (animal Bird) Move() {
	fmt.Println(animal.data.locomotion)
}

//Speak prints out what sound bird makes
func (animal Bird) Speak() {
	fmt.Println(animal.data.noise)
}

//Eat prints out what snake eats
func (animal Snake) Eat() {
	fmt.Println(animal.data.food)
}

//Move prints out how snake moves
func (animal Snake) Move() {
	fmt.Println(animal.data.locomotion)
}

//Speak prints out what sound snake makes
func (animal Snake) Speak() {
	fmt.Println(animal.data.noise)
}

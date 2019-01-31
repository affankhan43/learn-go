/*
Week3 Assignment:
Write a program which allows the user to get information about a predefined set of animals. Three animals are predefined,
 cow, bird, and snake. Each animal can eat, move, and speak. The user can issue a request to find out one of three
 things about an animal: 1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks.
 The following table contains the three animals and their associated data which should be hard-coded into your program.

Animal	Food eaten	Locomotion 	Spoken sound
cow		grass			walk		moo
bird	worms			fly			peep
snake	mice			slither		hsss
Your program should present the user with a prompt, â€œ>â€, to indicate that the user can type a request.
Your program accepts one request at a time from the user, prints out the answer to the request,
and prints out a new prompt. Your program should continue in this loop forever. Every request from the user
must be a single line containing 2 strings. The first string is the name of an animal, either â€œcowâ€, â€œbirdâ€, or â€œsnakeâ€.
The second string is the name of the information requested about the animal, either â€œeatâ€, â€œmoveâ€, or â€œspeakâ€.
Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal. Make a type called Animal which is a
struct containing three fields:food, locomotion, and noise, all of which are strings. Make three methods
called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type.
The Eat() method should print the animalâ€™s food, the Move() method should print the animalâ€™s locomotion,
and the Speak() method should print the animalâ€™s spoken sound. Your program should call the appropriate
method when the user makes a request.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal struct
type Animal struct {
	Food       string
	Locomotion string
	Sound      string
}

var (
	// list of animals
	animals = map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}
)

// Eat prints the Animal food
func (a Animal) Eat() {
	fmt.Println(a.Food)
}

// Move prints the Animal locomotion
func (a Animal) Move() {
	fmt.Println(a.Locomotion)
}

// Speak prints the Animal sound
func (a Animal) Speak() {
	fmt.Println(a.Sound)
}

// commands for animals
func (a Animal) Do(commands string) {
	switch commands {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Printf("%s:Action Not Exists. %s can do only: eat, move or speak\n", commands,a)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	for scanner.Scan() {

		in := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading input: %v", err)
			break
		}

		// Validate input commands
		commands := strings.Fields(strings.ToLower(in))
		if len(commands) == 0 || len(commands) > 2 || len(commands) == 1 {
			if (len(commands) == 1 && commands[0] == "quit"){
				fmt.Println("Exitting ...")
				break	
			}else{
				fmt.Println("incorrect input. Usage: <animal> <action> e.g.: cow speak. Or 'quit' to exit.")
				fmt.Print("> ")
				continue
			}
		}

		// Separate animal and action
		animal := commands[0]
		cmd := commands[1]

		// run action on know animal
		if a, exists := animals[animal]; exists {
			a.Do(cmd)
		} else {
			fmt.Printf("Unknown Animal %s \n", animal)
		}

		fmt.Print("> ")
	}

}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var characteristics []charact
var animal map[string]interface{}
var newanimalorquery, newanimalnameorname, typeoracitivity string

type Animal interface {
	Eat()
	Move()
	Sound()
}

type cow struct {
	Foodeaten string
	Locomtion string
	Spoken    string
}
type bird struct {
	Foodeaten string
	Locomtion string
	Spoken    string
}

type snake struct {
	Foodeaten string
	Locomtion string
	Spoken    string
}

type charact struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (c *cow) Eat() {
	fmt.Println(c.Foodeaten, "by the animal")
}

func (c *cow) Move() {
	fmt.Println(c.Locomtion, "by the animal")
}

func (c *cow) Speak() {
	fmt.Println(c.Spoken, "by the animal")
}
func (b *bird) Eat() {

	fmt.Println(b.Foodeaten, "by the animal")
}

func (b *bird) Move() {
	fmt.Println(b.Locomtion, "by the animal")
}

func (b *bird) Speak() {
	fmt.Println(b.Spoken, "by the animal")

}
func (s *snake) Eat() {
	fmt.Println(s.Foodeaten, "by the animal")
}

func (s *snake) Move() {
	fmt.Println(s.Locomtion, "by the animal")
}

func (s *snake) Speak() {
	fmt.Println(s.Spoken, "by the animal")
}

func newAnimal(typeAnimal string, Newanimlaname string) {
	for _, k := range characteristics {
		if k.name == strings.ToLower(typeAnimal) {
			temp := charact{}
			temp.name = Newanimlaname
			temp.noise = k.noise
			temp.locomotion = k.locomotion
			temp.food = k.food
			characteristics = append(characteristics, temp)
		}
	}
	fmt.Println(characteristics)
	fmt.Println("New Animal Added!")
}

func query(typeAnimal string, activity string) {
	for _, k := range characteristics {
		str := animal[typeAnimal]
		switch v := str.(type) {
		case cow:
			if k.name == typeAnimal {
				v.Foodeaten = k.food
				v.Locomtion = k.locomotion
				v.Spoken = k.noise
				if activity == "eat" {
					v.Eat()
				} else if activity == "move" {
					v.Move()
				} else if activity == "sound" {
					v.Speak()
				} else {
					fmt.Println("activity is not listed in the list")
				}
			}
		case snake:
			if k.name == typeAnimal {
				v.Foodeaten = k.food
				v.Locomtion = k.locomotion
				v.Spoken = k.noise
				if activity == "eat" {
					v.Eat()
				} else if activity == "move" {
					v.Move()
				} else if activity == "sound" {
					v.Speak()
				} else {
					fmt.Println("activity is not listed ")
				}
			}
		case bird:
			if k.name == typeAnimal {
				v.Foodeaten = k.food
				v.Locomtion = k.locomotion
				v.Spoken = k.noise
				if activity == "eat" {
					v.Eat()
				} else if activity == "move" {
					v.Move()
				} else if activity == "sound" {
					v.Speak()
				} else {
					fmt.Println("activity is not listed in the list")
				}
			}
		default:
			fmt.Println("no data", v)
		}
	}
}

func main() {
	characteristics = []charact{
		{"cow", "grass", "walk", "moo"},
		{"bird", "worms", "fly", "peep"},
		{"snake", "mice", "slither", "hsss"},
	}
	animal = make(map[string]interface{})
	cow := cow{}
	bird := bird{}
	snake := snake{}

	animal["cow"] = cow
	animal["bird"] = bird
	animal["snake"] = snake

	for {
		fmt.Println("Enter 3 String : QUERY or NEWANIMAL/AnimalName or NewAnimalName/Activity or Type:")
		fmt.Print("> ")
		var line string
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line = scanner.Text()
			if len(strings.Fields(line)) == 3 {

				newanimalorquery = strings.Fields(line)[0]
				newanimalnameorname = strings.Fields(line)[1]
				typeoracitivity = strings.Fields(line)[2]

				if strings.ToUpper(newanimalorquery) == "NEWANIMAL" {
					if typeoracitivity == "cow" {
						animal[newanimalnameorname] = cow
					} else if typeoracitivity == "snake" {
						animal[newanimalnameorname] = snake
					} else if typeoracitivity == "bird" {
						animal[newanimalnameorname] = bird
					} else {
						fmt.Println("Pass either cow/snake/bird")
					}

					newAnimal(typeoracitivity, newanimalnameorname)
					var answer string
					fmt.Println("Do you want to continue yes/no ")
					_, err := fmt.Scan(&answer)
					if err!=nil{
						os.Exit(0)
					} else{
						if answer == "yes" {
							break
						} else {
							os.Exit(0)
						}
					}
				} else if strings.ToUpper(newanimalorquery) == "QUERY" {

					query(newanimalnameorname, typeoracitivity)
					var answer string
					fmt.Println("Do you want to continue yes/no ")
					_, err := fmt.Scan(&answer)
					if err!=nil{
						os.Exit(0)
					} else{
						if answer == "yes" {
							break
						} else {
							os.Exit(0)
						}
					}
				} else {
					fmt.Println("Wrong value")
					os.Exit(0)
				}
			} else {
				fmt.Println("Enter 3 string in same line")
				fmt.Print("> ")
			}
		}
	}
}
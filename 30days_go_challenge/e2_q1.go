// MODIFY THIS AREA
// [filename]: e2_q1.go
// [question]: 1
// [instagram username]: affan.ak43

package main

import (
	"fmt"
)

type Cat struct{
	Color string
	Name  string
	Legs  int
}
type Owl struct{
	Color string
	Name  string
	Legs  int
}

func main() {
	Cats:= Cat{
		Color: "White",
		Name:  "Ollie",
		Legs:  4,
	}
	Owls:=Owl{
		Color: "Black",
		Name:  "Tawny",
		Legs:  2,
	}
	//Full Output of struts
	fmt.Println(Cats,Owls)
	fmt.Println("Cat Name is",Cats.GetCatName())
	fmt.Println(Cats.GetCatName(),"has ",Cats.GetCatLegs(),"legs.")
	Cats.SetCatName("Bowie")
	fmt.Println(Cats.GetCatName(),"is cat's new name.")
	fmt.Println("Owl Name is",Owls.GetOwlName())
	fmt.Println(Owls.GetOwlName(),"has ",Owls.GetOwlLegs(),"legs.")
	Owls.SetOwlName("Screech")
	fmt.Println(Owls.GetOwlName(),"is owl's new name.")

}

func(c Cat) GetCatName() string{
	return c.Name
}
func(o Owl) GetOwlName() string{
	return o.Name
}
func(c Cat) GetCatLegs() int{
	return c.Legs
}
func(o Owl) GetOwlLegs() int{
	return o.Legs
}
func(c *Cat) SetCatName(name string){
	c.Name = name
}
func(o *Owl) SetOwlName(name string){
	o.Name = name
}
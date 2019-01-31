// MODIFY THIS AREA
// [filename]: e1_q4.go
// [question]: 4
// [instagram username]: affan.ak43

package main

import (
	"fmt"
)
 
func main() {
	planets := []string{"Mars","Earth","Saturn","Jupiter","Venus"}
	for i, planet := range planets{
		fmt.Printf("Planet: %s (index: %v) \n",planet,i)
	}
	food:="Pommes Frites"
	for j,ch:= range food{
		fmt.Printf("Character: %s, Position:%v \n",string(ch),j)
	}
}

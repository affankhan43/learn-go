// MODIFY THIS AREA
// [filename]: e1_q3.go
// [question]: 3
// [instagram username]: affan.ak43

package main

import (
	"fmt"
)
 
func main() {
	fmt.Println("prints out integers from 0 to 10.")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("prints out even numbers counting down from 100 to 0.")
	for j := 100; j >= 0; j-=2 {
		fmt.Println(j)
	}
	fmt.Println("Nested Loops")
	for a := 0; a <= 5 ; a++ {
		for b := 0; b <= 10; b+=2 {
			fmt.Printf("Outer Loop: %v , ",a)
			fmt.Printf("Inner Loop: %v \n",b)
		}
		//fmt.Println()
	}
}

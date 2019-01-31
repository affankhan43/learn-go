// MODIFY THIS AREA
// [filename]: e1_q2.go
// [question]: 2
// [instagram username]: affan.ak43

package main

import (
	"fmt"
)
 
func main() {
	var slice1 []string
	slice1 = make([]string,2)
	slice1[0],slice1[1]="Affan Ahmed","Khan"
	fmt.Printf("%s %s \n ",slice1[0],slice1[1])
	var(
		b1 bool= true
		b2 bool= false
		f1 float32= 3.11
		f2 float32= 6.22
	)
	fmt.Println(b1,b2,f1,f2)
}

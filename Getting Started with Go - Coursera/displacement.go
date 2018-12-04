package main

import (
	"fmt"
	"os"
)

// acceleration a float64, initial velocity vo float64, and initial displacement so float64
func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	fn := func(t float64) float64 {
		s := so + vo*t + 0.5*a*t*t
		return s
	}
	return fn
}

func main() {

	var a, v0, s0, t float64
	fmt.Printf("Enter Acceleration (a) -| ")
	n, err := fmt.Scanf("%f\n", &a)
	if err != nil || n != 1 {
		fmt.Println("Invalid Input Error", n, err)
		os.Exit(1)
	}
	fmt.Printf("Enter Initial Velocity (v0) -| ")
	n, err = fmt.Scanf("%f\n", &v0)
	if err != nil || n != 1 {
		fmt.Println("Invalid Input Error", n, err)
		os.Exit(1)
	}
	fmt.Printf("Enter Initial Displacement (s0) -| ")
	n, err = fmt.Scanf("%f\n", &s0)
	if err != nil || n != 1 {
		fmt.Println("Invalid Input Error", n, err)
		os.Exit(1)
	}
	fmt.Printf("Enter Time(secs) to Compute Total Displacement (t) -| ")
	n, err = fmt.Scanf("%f\n", &t)
	if err != nil || n != 1 {
		fmt.Println("Invalid Input Error", n, err)
		os.Exit(1)
	}
	if t < 0 {
		fmt.Println("Invalid Time Input Error", t)
		os.Exit(1)
	}

	fmt.Println(a, v0, s0, t)
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println("Total Computed Displacement (s) -|", fn(t))
}
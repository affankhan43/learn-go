package main

import (
	"fmt"
)

func takeInputs(numberList []int) []int {
	var number int
	for i := 0; i < 10; i++ {
		fmt.Printf("Type a number (%d to go)\n", 10-i)
		fmt.Scan(&number)
		numberList = append(numberList, number)
	}
	return numberList
}

func swap(slice []int, index int) {
	valueAtIndex := slice[index]
	slice[index] = slice[index+1]
	slice[index+1] = valueAtIndex
}

func boubleSort(numberList []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(numberList)-1; i++ {
			if numberList[i] > numberList[i+1] {
				swap(numberList, i)
				swapped = true
			}
		}
	}
}

func main() {
	numberList := []int{}
	numberList = takeInputs(numberList)
	fmt.Printf("The Slice: %v\n", numberList)
	boubleSort(numberList)
	fmt.Printf("The Sorted Slice: %v\n", numberList)
}
// MODIFY THIS AREA
// [filename]: a2_q1.go
// [question]: 1
// [instagram username]: affan,ak43
package main

import (
	"fmt"
	"math/rand"
)


func main() {
	slice1 := []int{1, 4, 4, 7, 9, 12, 7, 61, 44, 2, 8, 6, 9, 11}
	slice2 := []int{2, 14, 6, 88, 12, 91, 57, 23, 33, 30, 14, 2, 30}

	exists:=func (sl []int, no int) bool{
		for _, x := range sl {
			if no == x {
				return true
			}
		}
		return false
	}
	find:=func (sl []int) []int{
		keys := make(map[int]bool)
		list := []int{}
		for _, entry := range sl {
			if _, value := keys[entry]; !value {
				keys[entry] = true
				list = append(list, entry)
			}
		}
		return list
	}
	merge:=func (sl1 []int, sl2 []int) []int{
		final:=[]int{}
		final=append(sl1,sl2...)
		final=find(final)
		return final
	}
	var sort func(a []int) []int
	sort=func (a []int) []int {
		if len(a) < 2 {
			return a
		}
		left, right := 0, len(a)-1
		// Pick a pivot
		pivotIndex := rand.Int() % len(a)
		// Move the pivot to the right
		a[pivotIndex], a[right] = a[right], a[pivotIndex]
		// Pile elements smaller than the pivot on the left
		for i := range a {
			if a[i] < a[right] {
				a[i], a[left] = a[left], a[i]
				left++
			}
		}
		// Place the pivot after the last smaller element
		a[left], a[right] = a[right], a[left]
		// Go down the rabbit hole
		sort(a[:left])
		sort(a[left+1:])
		return a
	}
	merged_slice:=merge(slice1,slice2)
	fmt.Println("Removed Duplicates From Slice1",find(slice1))
	fmt.Println("Removed Duplicates From Slice2",find(slice2))
	fmt.Println("Removed duplicates from merge of Slice 1 and 2:",merged_slice)
	fmt.Println("Sorted and Removed duplicates from merge of Slice 1 and 2:",sort(merged_slice))
	fmt.Println("Search \"1\" From Slice2: ",exists(slice2,1))
}

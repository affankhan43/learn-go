package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	n:=2
	arry := []int{2, 2, 2, 2, 2, 2}

	var wg sync.WaitGroup
	wg.Add(n)

	go func() {
		a := rand.Intn(6)
		arry[a]++
		wg.Done()
	}()

	go func() {
		b := rand.Intn(6)
		arry[b]++
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(arry)
}

// Race Condition:
// It happens when two goroutines access the same arry at the same time and they are accessing the same item of this arry while at least one of them is modifying that item
// Communication is the source of race condition
// Some Outputs:
//  E:\Coursera>go run gorutines2.go
//  [3 3 2 2 2 2]

//  E:\Coursera>go run gorutines2.go
//  [4 2 2 2 2 2]

//  E:\Coursera>go run gorutines2.go
//  [2 2 4 2 2 2]
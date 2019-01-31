package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ReadIntegers - reads a list of integers from user
func ReadIntegers() []int {
	fmt.Print("Enter integers separated by space: ")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	input = regexp.MustCompile(`[\s\p{Zs}]{2,}`).ReplaceAllString(input, " ")
	input = regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`).ReplaceAllString(input, "")
	inputs := strings.Split(input, " ")
	inputSize := len(inputs)

	list := make([]int, 0)

	for i := 0; i < inputSize; i++ {
		numb, err := strconv.Atoi(inputs[i])
		if err == nil {
			list = append(list, numb)
		}
	}

	return list
}

// QuarterList - divides an array of integers to roughly 4 equal arrays
func QuarterList(list []int) ([4][]int, error) {
	quarters := [4][]int{[]int{}, []int{}, []int{}, []int{}}
	if len(list) < 4 {
		return quarters, errors.New("List must have at least 4 items to quarter")
	}

	size := len(list)
	quarter := int(math.Floor(float64(size) / 4.0))

	for i := 0; i < 4; i++ {
		i0 := i * quarter
		i1 := (i + 1) * quarter
		if i == 3 {
			i1 = size
		}
		quarters[i] = list[i0:i1]
	}

	return quarters, nil
}

// Swap - modifieds as list by swapping the indexed slice item
// with the adjacent item (index + 1)
func Swap(list []int, index int) {
	temp := list[index]
	list[index] = list[index+1]
	list[index+1] = temp
}

// BubbleSort - modifies a list with bubble sort algorithm
func BubbleSort(list []int, channel chan int) {
	size := len(list)

	for i := 0; i < size-1; i++ {
		if list[i] > list[i+1] {
			Swap(list, i)
		}
	}

	if size > 2 {
		subchannel := make(chan int)
		go BubbleSort(list[0:size-1], subchannel)
		<-subchannel
	}

	channel <- 1
}

func main() {
	channel := make(chan int)
	numbers := ReadIntegers()
	quarters, _ := QuarterList(numbers)

	for i := 0; i < len(quarters); i++ {
		fmt.Println("Sorting thread ", quarters[i])
		go BubbleSort(quarters[i], channel)
	}

	<-channel
	<-channel
	<-channel
	<-channel

	fulllist := make([]int, 0)
	fulllist = append(fulllist, quarters[0]...)
	fulllist = append(fulllist, quarters[1]...)
	fulllist = append(fulllist, quarters[2]...)
	fulllist = append(fulllist, quarters[3]...)

	fmt.Println("Sorting thread ", fulllist)
	go BubbleSort(fulllist, channel)
	<-channel

	fmt.Println(fulllist)
}
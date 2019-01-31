package main

import (
  "bufio"
  "fmt"
  "os"
  "sort"
  "strconv"
  "strings"
  "sync"
)

func groups_sorting(x []string, wg *sync.WaitGroup) []int {
  y := make([]int, 0)
  for _, i := range x {
    new_i, _ := strconv.Atoi(i)
    y = append(y, new_i)
  }
  sort.Ints(y)
  fmt.Println(y)
  wg.Done()
  return y
}

func merged_sort(args ...[]int) []int {
  total := make([]int,0)
  for _, z := range args {
    total = append(total, z...)
  }
  sort.Ints(total)
  return total
}


func main() {
  scanner := bufio.NewReader(os.Stdin)
  fmt.Println("Please enter Integers seperated with spaces.")
  fmt.Printf(">")
  line, err := scanner.ReadString('\n')
  if err != nil {
    fmt.Println(err)
  }

  g_r1 := make([]int,0)
  g_r2 := make([]int,0)
  g_r3 := make([]int,0)
  g_r4 := make([]int,0)
  var wg sync.WaitGroup
  wg.Add(4)

  input_ints       := strings.Fields(line)
  inp_nums         := len(input_ints)
  divide           := inp_nums / 4
  group_one        := input_ints[0:divide]
  input_ints        = input_ints[divide:]
  group_two        := input_ints[0:divide]
  input_ints        = input_ints[divide:]
  group_three      := input_ints[0:divide]
  input_ints        = input_ints[divide:]
  group_four       := input_ints[:]

  go func() {
    fmt.Println("Gorutine#1 ")
    g_r1 = groups_sorting(group_one, &wg)
  }()
  go func() {
    fmt.Println("Gorutine#2 ")
    g_r2 = groups_sorting(group_two, &wg)
  }()
  go func() {
    fmt.Println("Gorutine#3 ")
    g_r3 = groups_sorting(group_three, &wg)
  }()
  go func() {
    fmt.Println("Gorutine#4 ")
    g_r4 =  groups_sorting(group_four, &wg)
  }()

  wg.Wait()
  sorted     := merged_sort(g_r1,g_r2,g_r3,g_r4)
  fmt.Println("Proccesses Completed...")
  fmt.Println(sorted)
}
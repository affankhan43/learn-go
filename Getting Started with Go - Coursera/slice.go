package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

func main() {
    enter := bufio.NewScanner(os.Stdin)
    ss := make([]int, 3)
    var count = 0

    fmt.Println("Press 'X' to exit the program or enter an integer:")
    for enter.Scan() {
        ans, err := strconv.Atoi(enter.Text())
        if err != nil {
            if strings.ToUpper(enter.Text()) == "X" {
                fmt.Println("Quit the program.")
                return
            }
            fmt.Println("Invalid Format! Try it again!")
            continue
        }
        fmt.Printf("You entered: %T %v\n", ans, ans)

        if count < 3 {
            // Check the index of slice elements is 0, then replace it
            ss[sort.SearchInts(ss, 0)] = ans
            count++
        } else {
            ss = append(ss, ans)
        }
        sort.Ints(ss)
        fmt.Println("Slice:", ss)
    }
}
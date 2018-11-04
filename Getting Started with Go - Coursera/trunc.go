package main

import (
    "fmt"
    "strconv"
)

func main() {
    var enter string
    fmt.Printf("Enter a Floating Point Number: ")
    _, err := fmt.Scan(&enter)

    if err == nil {
        if result, err := strconv.ParseFloat(enter, 64); err == nil {
            truncat := int64(result)
            fmt.Printf(strconv.FormatInt(truncat, 10))
            return
        }
    }
    fmt.Println("Please Enter Floating Number")
}
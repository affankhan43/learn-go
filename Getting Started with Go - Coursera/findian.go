package main

import (
    "fmt"
    "strings"
)

func main() {
    var enter string
    fmt.Printf("Enter String: ") 
    if _, err := fmt.Scan(&enter); err == nil {
        if strings.ContainsAny(enter, "i") && strings.ContainsAny(enter, "a") && strings.ContainsAny(enter, "n") {
            ln:=len(enter)
            if ln >= 3 && strings.Index(enter,"i") == 0 && strings.Index(enter,"n") == ln-1 {
                fmt.Println("FOUND!")
            } else {
                fmt.Println("NOT FOUND!")
            }
        } else {
            fmt.Println("NOT FOUND!")
        }
    }
}
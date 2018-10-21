package main

import (
"fmt"
// "bytes"
// "encoding/json"
// "math/big"
//"math"
)

func main() {

    /* Simple Arrays */
        var arr [5]string
        arr[0] = "a"
        arr[1] = "b"
        arr[2] = "c"
        arr[3] = "d"
        arr[4] = "e"
        fmt.Println(arr)

        arr1:=[]string{}
        row1:="abcd"
        row2:="abcde"
        arr1=append(arr1,row1)
        arr1=append(arr1,row2)
        fmt.Println(arr1)
    /* Simple Array END */

}
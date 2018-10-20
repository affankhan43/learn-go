package main

import (
"fmt"
"encoding/json"
)

func main() {
    /* Integer */
        z:=30
        fmt.Println(z)
        var n1 uint8 // Unsigned 8-bit integers (0 to 255)
        n1 = 200
        fmt.Println(n1)
         
        var n2 uint16 // Unsigned 16-bit integers (0 to 65535)
        n2 = 54200
        fmt.Println(n2)
         
        var n3 uint32 // Unsigned 32-bit integers (0 to 4294967295)
        n3 = 98765214
        fmt.Println(n3)
         
        var n4 uint64 // Unsigned 64-bit integers (0 to 18446744073709551615)
        n4 = 1844674073709551615
        fmt.Println(n4)
         
        var n5 int8 //Signed 8-bit integers (-128 to 127)
        n5 = -52
        fmt.Println(n5)
        fmt.Println(n5*-1)
         
        var n6 int16 // Signed 16-bit integers (-32768 to 32767)
        n6 = -32552
        fmt.Println(n6)
        fmt.Println(n6*-1)
         
        var n7 int32 // Signed 32-bit integers (-2147483648 to 2147483647)
        n7 = -98658754
        fmt.Println(n7)
        fmt.Println(n7*-1)
         
        var n8 int64 // Signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
        n8 = -92211111111111111
        fmt.Println(n8)
        fmt.Println(n8*-1)
    /* Integer End */

    /*Float*/
        var f1 float32 = 1677.7216 // IEEE-754 32-bit floating-point numbers
        fmt.Println(f1)
     
        var f2 float64 = 18787677.878716 // IEEE-754 64-bit floating-point numbers
        fmt.Println(f2)
         
        var f3 complex64 = 8789579877.721 // Complex numbers with float32 real and imaginary parts
        fmt.Println(f3)
         
        var f4 complex128 = 985214745.7216 // Complex numbers with float64 real and imaginary parts
        fmt.Println(f4)
    /*Float End*/

    /* String */
        var y string = "Hello World"
        fmt.Println(y)
    /* String End*/

    /* Boolean */    
        var a,b bool // Boolean type declaration
        a=true
        b=false
        fmt.Println(a,b)
    /* Boolean End */

    mapD := make(map[string]string)
    mapD["k"]="a"
    mapD["aa"]="asd"
    mapB,_ := json.Marshal(mapD)
    fmt.Println(string(mapB))
}
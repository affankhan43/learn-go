package main

import(
	"fmt"
	"encoding/json"
)

func main(){
	var name string
	fmt.Printf("Please Enter Name: ")
	fmt.Scan(&name)
	var address string
	fmt.Printf("Please Enter Address: ")
	fmt.Scan(&address)
	p1:=map[string]string{"name":name,"address":address}
	jsonString,err := json.Marshal(p1)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(jsonString))

}
// package main

// import (
// 	"bufio"
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	var jsonMap = map[string]string{}

// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Enter name: ")
// 	if str, err := reader.ReadString('\n'); err == nil {
// 		jsonMap["name"] = strings.TrimSpace(str)
// 	} else {
// 		panic("Input Error")
// 	}

// 	fmt.Print("Enter address: ")
// 	if str, err := reader.ReadString('\n'); err == nil {
// 		jsonMap["address"] = strings.TrimSpace(str)
// 	} else {
// 		panic("Input Error")
// 	}

// 	if jsonBytes, err := json.Marshal(jsonMap); err == nil {
// 		fmt.Println(string(jsonBytes))
// 	} else {
// 		panic("Marshalling error")
// 	}
// }
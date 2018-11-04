package main

import (
"fmt"
"encoding/json"
)

type withdraw_data struct{
	Address string `json:"address"`
	Amount int `json:"amount"`
}
type Data struct{
	Coin string `json:"coin"`
	Withdraw []withdraw_data `json:"withdraw_data"`
}


func main() {
	
	first:=map[int]string{}
	first[0]="hello"
	fmt.Println(first)

	//var post_data Data
	var w withdraw_data
	w.Address="asd"
	w.Amount=2
	wd:= []withdraw_data{}
	wd=append(wd,w)
	full := &Data{Coin: "KMD",Withdraw:wd}
	fmt.Println(full)
	jsonString,_:=json.Marshal(full)
	fmt.Println(string(jsonString))
}
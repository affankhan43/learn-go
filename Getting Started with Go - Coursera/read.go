package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main(){
	var name string
	lines:=[]Person{}
	fmt.Printf("Please File Name: ")
	fmt.Scan(&name)
	inFile, err := os.Open(name)
	if err != nil {
		fmt.Println(err.Error() + `: ` + name)
		return
	} else {
		defer inFile.Close()
	}
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)       
	for scanner.Scan() {
		str:=scanner.Text()
		words := strings.Fields(str)
		line:=Person{fname:words[0],lname:words[1]}
		lines=append(lines,line)
		fmt.Println(words, len(words))
	}
	fmt.Println(lines)
}
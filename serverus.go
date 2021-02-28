package main

import (
	"fmt"
	"strings"
)

func main(){
	path := "/home/abel/programming/serverusSnape//server-connector.go"
	_, err := getFile(path)
	if err != nil {
		fmt.Printf("Error: $s\n", err)
	}
	err = putFile("/home/abel/programming/serverusSnape/hello.txt", strings.NewReader("Hello from hello.txt\n"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File recieved")
}

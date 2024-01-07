package main

import (
	"fmt"
	"io"
	"os"
)



func main(){
	myString :=  ""
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Print("Please give me argument")
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, "This is Standart output\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")
}
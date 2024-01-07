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

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}
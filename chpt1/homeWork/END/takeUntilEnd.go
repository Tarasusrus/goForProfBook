/*
Напишите Go-программу, которая считывает целые числа до тех пор, пока не встретит во входных данных слово END.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	var f *os.File
	defer f.Close()
	f = os.Stdin
	
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if scanner.Text() == "END" {
			os.Exit(0)
		}
		fmt.Println(scanner.Text())
	}
}

package main

import (
	"bytes"
)


func main()  {
	var str string
	var run rune

	str = "New string"
	run = 'N'

	buf := []byte{1,2,3,4} // he can change
	buf[0] = 1
	// he also can change cap
	buf = append(buf, 100)

	//string can't change 
	//str[0] = "d"

}

func NewReader(b []byte) *bytes.Reader
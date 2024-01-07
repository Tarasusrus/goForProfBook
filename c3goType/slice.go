package main

import "fmt"

func main()  {
	aSlice := []int{1,2,3,4,5}
	fmt.Println(aSlice) // 1,2,3,4,5

	integers :=make([]int,2)
	fmt.Println(integers) // 0 0

	integers = nil
	fmt.Println(integers) // nil nil

	anArray := [5]int{-1, -2, -3, -4, -5}
	refAnArray := anArray[:]

	fmt.Println(anArray) //-1, -2, -3, -4, -5
	fmt.Println(refAnArray) // -1, -2, -3, -4, -5
	anArray[4] = -100
	fmt.Println(refAnArray) // -1, -2, -3, -4, -100

	s := make([]byte, 5)
	fmt.Println(s) // 00000

	twoD := make([][]int, 3)
	fmt.Println(twoD) // [[000][000]]
	fmt.Println()





}
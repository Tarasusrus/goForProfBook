/*
Создайте массив и слайс в Golang.
Попробуйте изменить их размеры и элементы. Объясните разницу в их поведении
*/

package main

import "fmt"

var arr = [5]int{1,2,3,4,5} 

var slice = make([]int,0,10)

func main(){
	fmt.Println("Array: ",arr)
	fmt.Println("Slice: ",slice)
	var sliceArr = arr[1:]
	var coppyArr = arr
	var copySlice = slice
	fmt.Printf("Adress of oldSlice: %p\n", &slice)
	
	slice = append(slice, arr[1])
	fmt.Printf("Adress of arr: %p\n", &arr)
	fmt.Printf("Adress of copy arr: %p\n", &coppyArr)
	fmt.Printf("Adress of slice: %p\n", &slice)
	fmt.Printf("Adress of copy slice: %p\n", &copySlice)
	fmt.Printf("Adress of arr: %p\n", &arr)
	fmt.Printf("Adress of sliceArr: %p\n", &sliceArr)

	newSlice := arr
	fmt.Println("newSlice = ", newSlice)
	newSlice[1] = 20
	fmt.Println(newSlice, arr )

	slice2 := newSlice[:2]
	fmt.Println(slice2)
	slice2[1] = 0
	fmt.Println(slice2,newSlice)
	slice2 = slice2[0:cap(slice2)]
	fmt.Println(slice2)

}


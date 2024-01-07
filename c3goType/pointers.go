package main

import "fmt"


func getPointer(n *int) { // указатель содержит адресс, ничего возвращать не нужно
	*n = *n * *n
}

func returnPointer(n int) *int {
	v := n * n
	return &v
}

func main()  {
	i := -10
	j := 25

	pI := &i // тут хранится адрес 
	pJ := &j //тут хранится адрес

	fmt.Println("pI memory:", pI)
    fmt.Println("pI memory:", pJ)
    fmt.Println("pI value:", *pI) // разыыменовываем и получаем значение по адресу 
    fmt.Println("pI value:", *pJ) //

	//Здесь показано, как можно изменить переменную i с помощью указателя pI, 
	//который ссылается на i, двумя разными способами: во-первых, путем непосредственного присвоения нового значения и, 
	//во-вторых, с помощью оператора --.
	*pI = 123456
	*pI--
	fmt.Println("i: ", i)
}
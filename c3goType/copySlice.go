package main

import (
    "fmt"
)

func main() {
    // Копирование элементов из a4 в a6.
    a6 := []int{-10, 1, 2, 3, 4, 5}
    a4 := []int{-1, -2, -3, -4}
    fmt.Println("a6:", a6)
    fmt.Println("a4:", a4)
    copy(a6, a4)  // Копирует все элементы a4 в a6.
	fmt.Println("--- Копирует все элементы a4 в a6---")
    fmt.Println("a6:", a6)  // a6 будет изменен.
    fmt.Println("a4:", a4)  // a4 останется без изменений.
    fmt.Println()

    // Копирование элементов из b6 в b4.
    b6 := []int{-10, 1, 2, 3, 4, 5}
    b4 := []int{-1, -2, -3, -4}
    fmt.Println("b6:", b6)
    fmt.Println("b4:", b4)
    copy(b4, b6)  // Копирует первые 4 элемента b6 в b4.
	fmt.Println("--- Копирует первые 4 элемента b6 в b4 ---")
    fmt.Println("b6:", b6)  // b6 останется без изменений.
    fmt.Println("b4:", b4)  // b4 будет изменен.
    fmt.Println()

    // Копирование элементов из массива array4 в срез s6.
    array4 := [4]int{4, -4, 4, -4}
    s6 := []int{1, 1, -1, -1, 5, -5}
    copy(s6, array4[0:])  // Копирует все элементы array4 в s6.
	fmt.Println("--- Копирует все элементы array4 в s6 ---")
    fmt.Println("array4:", array4[0:])
    fmt.Println("s6:", s6)  // s6 будет изменен.
    fmt.Println()

    // Копирование элементов из среза s7 в массив array5.
    array5 := [5]int{5, -5, 5, -5, 5}
    s7 := []int{7, 7, -7, -7, 7, -7, 7}
    copy(array5[0:], s7)  // Копирует первые 5 элементов s7 в array5.
	fmt.Println("--- Копирует первые 5 элементов s7 в array5 ---")
    fmt.Println("array5:", array5)  // array5 будет изменен.
    fmt.Println("s7:", s7)  // s7 останется без изменений.
}

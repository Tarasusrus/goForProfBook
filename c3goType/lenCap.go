package main

import "fmt"

// Функция printSlice принимает срез целых чисел и печатает их.
func printSlice(x []int) {
    for _, number := range x {
        fmt.Print(number, " ")  // Выводит каждый элемент среза с пробелом.
    }
    fmt.Println()  // Завершает вывод с новой строки.
}

func main() {
    aSlice := []int{-1, 0, 4}  // Инициализация среза с тремя элементами.
    fmt.Printf("aSlice: ")
    printSlice(aSlice)  // Вывод элементов среза.
    fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))  // Вывод емкости и длины среза.

    aSlice = append(aSlice, -100)  // Добавление одного элемента в срез.
    fmt.Printf("aSlice: ")
    printSlice(aSlice)  // Вывод элементов среза после добавления.
    fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))  // Вывод обновленной емкости и длины.

    // Добавление еще трех элементов в срез.
    aSlice = append(aSlice, -2)
    aSlice = append(aSlice, -3)
    aSlice = append(aSlice, -4)
    printSlice(aSlice)  // Вывод элементов среза после дополнительных добавлений.
    fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))  // Вывод финальной емкости и длины.
}
// aSlice: -1 0 4 
// Cap: 3, Length: 3
// aSlice: -1 0 4 -100 
// Cap: 6, Length: 4
// -1 0 4 -100 -2 -3 -4 
// Cap: 12, Length: 7
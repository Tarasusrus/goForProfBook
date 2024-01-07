package main

import (
    "fmt"
    "unsafe"  // Импорт пакета unsafe для работы с указателями и арифметикой указателей.
)

func main() {
    array := [...]int{0, 1, -2, 3, 4}  // Объявление массива из пяти целых чисел.

    // Получение указателя на первый элемент массива.
    pointer := &array[0]
    fmt.Print(*pointer, " ")  // Печать значения первого элемента массива.

    // Вычисление адреса второго элемента массива, используя адрес первого элемента.
    memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])

    for i := 0; i < len(array)-1; i++ {
        // Приведение вычисленного адреса обратно к указателю на int и разыменование его для получения значения.
        pointer = (*int)(unsafe.Pointer(memoryAddress))
        fmt.Print(*pointer, " ")  // Печать значения текущего элемента.

        // Вычисление адреса следующего элемента.
        memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
    }
    fmt.Println()

    // Попытка доступа к памяти за пределами массива (potentially unsafe!).
    pointer = (*int)(unsafe.Pointer(memoryAddress))
    fmt.Print("One more: ", *pointer, " ")
    memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
    fmt.Println()
}

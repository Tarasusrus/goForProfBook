package main

import (
    "fmt"
)

func main() {
    // Цикл for с использованием continue и break.
    for i := 0; i < 100; i++ {
        if i%20 == 0 {
            continue  // Пропускаем итерации, где i делится на 20 без остатка.
        }
        if i == 95 {
            break  // Прекращаем цикл, когда i достигает 95.
        }
        fmt.Print(i, " ")  // Печатаем значение i и пробел.
    }
    fmt.Println()

    // Обратный цикл от 10 до 0.
    i := 10
    for {
        if i < 0 {
            break  // Прекращаем цикл, когда i меньше 0.
        }
        fmt.Print(i, " ")  // Печатаем значение i и пробел.
        i--
    }
    fmt.Println()

    // Цикл с использованием внешнего условия.
    i = 0
    anExpression := true
    for ok := true; ok; ok = anExpression {
        if i > 10 {
            anExpression = false  // Меняем внешнее условие для прекращения цикла.
        }
        fmt.Print(i, " ")  // Печатаем значение i и пробел.
        i++
    }
    fmt.Println()

    // Цикл range для итерации по массиву.
    anArray := [5]int{0, 1, -1, 2, -2}
    for i, value := range anArray {
        fmt.Println("index:", i, "value: ", value)  // Печатаем индекс и значение каждого элемента массива.
    }
}

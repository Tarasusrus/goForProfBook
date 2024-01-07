package main

import (
    "fmt"
    "sort"
)

// aStructure определяет структуру с именем, ростом и весом человека.
type aStructure struct {
    person string
    height int
    weight int
}

func main() {
    // Создание среза структур aStructure.
    mySlice := make([]aStructure, 0)

    // Добавление нескольких структур в срез.
    mySlice = append(mySlice, aStructure{"Mihalis", 180, 90})
    mySlice = append(mySlice, aStructure{"Bill", 134, 45})
    mySlice = append(mySlice, aStructure{"Marietta", 155, 45})
    mySlice = append(mySlice, aStructure{"Epifanios", 144, 50})
    mySlice = append(mySlice, aStructure{"Athina", 134, 40})

    fmt.Println("0:", mySlice)  // Вывод исходного набора данных.

    // Сортировка среза по весу с использованием sort.Slice.
    sort.Slice(mySlice, func(i, j int) bool {
        return mySlice[i].height < mySlice[j].height
    })
    fmt.Println("<:", mySlice)  // Вывод среза, отсортированного по весу.

    // Сортировка среза по росту с использованием sort.Slice.
    sort.Slice(mySlice, func(i, j int) bool {
        return mySlice[i].height > mySlice[j].height
    })
    fmt.Println(">:", mySlice)  // Вывод среза, отсортированного по росту.
}

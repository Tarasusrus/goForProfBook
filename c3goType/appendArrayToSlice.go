package main

import "fmt"

func main() {
    s := []int{1, 2, 3}  // Инициализация среза s.
    a := [3]int{4, 5, 6}  // Инициализация массива a.

    // Получение среза из массива a.
    ref := a[:]  // ref теперь ссылается на все элементы массива a.
    fmt.Println("Existing array:\t", ref)  // Вывод среза, созданного из массива.

    // Создание нового среза путем добавления элементов ref к s.
    t := append(s, ref...)  // Использование оператора ..., чтобы добавить каждый элемент ref к s.
    fmt.Println("New slice:\t", t)  // Вывод нового среза t.

    // Добавление элементов из ref в существующий срез s.
    s = append(s, ref...)
    fmt.Println("Existing slice:\t", s)  // Вывод измененного среза s.

    // Добавление всех элементов среза s к самому себе.
    s = append(s, s...)
    fmt.Println("s+s:\t\t", s)  // Вывод среза s после добавления к нему самого себя.
}
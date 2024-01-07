package main

import (
    "fmt"
)

func main() {
    // Создание карты с использованием make и добавление элементов.
    iMap := make(map[string]int)
    iMap["k1"] = 12
    iMap["k2"] = 13
    fmt.Println("iMap:", iMap)  // Вывод содержимого iMap.

    // Создание карты с использованием литерала карты.
    anotherMap := map[string]int{
        "k1": 12,
        "k2": 13,
    }
    fmt.Println("anotherMap:", anotherMap)  // Вывод содержимого anotherMap.

    // Удаление элемента с ключом "k1" из anotherMap.
    delete(anotherMap, "k1")  // Удаляет элемент "k1" из карты.
    delete(anotherMap, "k1")  // Повторное удаление не имеет эффекта, если ключ не существует.
    delete(anotherMap, "k1")  // Это безопасная операция, даже если ключ не существует.
    fmt.Println("anotherMap:", anotherMap)  // Вывод содержимого anotherMap после удаления.

    // Проверка наличия ключа в карте.
    _, ok := iMap["doesItExist"]
    if ok {
        fmt.Println("Exists!")  // Выводится, если ключ существует.
    } else {
        fmt.Println("Does NOT exist")  // Выводится, если ключ не существует.
    }

    // Итерация по элементам карты.
    for key, value := range iMap {
        fmt.Println(key, value)  // Выводит ключ и значение каждого элемента карты.
    }
}

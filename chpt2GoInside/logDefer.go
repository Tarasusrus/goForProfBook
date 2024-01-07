package main

import (
    "fmt"
    "log"
    "os"
)

var LOGFILE = "mGo.log"  // Путь к файлу лога.

// Функция one записывает числа от 0 до 9 в лог.
func one(aLog *log.Logger) {
    aLog.Println("-- FUNCTION one ------")  // Запись в лог при входе в функцию.
    defer aLog.Println("-- FUNCTION one ------")  // Отложенная запись в лог при выходе из функции.
    for i := 0; i < 10; i++ {
        aLog.Println(i)  // Запись в лог текущего значения i.
    }
}

// Функция two записывает числа от 10 до 1 в лог.
func two(aLog *log.Logger) {
    aLog.Println("---- FUNCTION two")  // Запись в лог при входе в функцию.
    defer aLog.Println("FUNCTION two ------")  // Отложенная запись в лог при выходе из функции.
    for i := 10; i > 0; i-- {
        aLog.Println(i)  // Запись в лог текущего значения i.
    }
}

func main() {
    // Открытие файла лога для записи.
    f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()  // Отложенное закрытие файла лога при выходе из main().

    // Создание нового логгера.
    iLog := log.New(f, "logDefer ", log.LstdFlags)
    iLog.Println("Hello there!")  // Запись в лог.
    iLog.Println("Another log entry!")  // Еще одна запись в лог.

    one(iLog)  // Вызов функции one.
    two(iLog)  // Вызов функции two.
}

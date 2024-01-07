package main

import (
    "fmt"
)

// Функция a() использует defer для обработки паники.
func a() {
    fmt.Println("Inside a()")
    
    // Отложенная анонимная функция вызывает recover().
    defer func() {
        if c := recover(); c != nil {
            fmt.Println("Recover inside a()!")
        }
    }()
    
    fmt.Println("About to call b()")
    b()  // Вызов функции b, которая вызывает панику.
    fmt.Println("b() exited!")  // Эта строка не будет выполнена из-за паники в b().
    fmt.Println("Exiting a()")
}

// Функция b() вызывает панику.
func b() {
    fmt.Println("Inside b()")
    panic("Panic in b()!")  // Паника с сообщением.
    fmt.Println("Exiting b()")  // Эта строка не будет выполнена из-за паники.
}

func main() {
    a()  // Вызов функции a.
    fmt.Println("main() ended!")  // Эта строка будет выполнена, так как паника была обработана.
}

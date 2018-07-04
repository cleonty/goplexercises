package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    abort := make(chan struct{})
    go func() {
        os.Stdin.Read(make([]byte, 1))
        abort <- struct{}{}
    }()
    fmt.Println("Начинаю отсчет. Нажмите <enter> для отказа...")
    select {
    case <- time.After(10 * time.Second):
    case <- abort:
        fmt.Println("Запуск отменен")
        return
    }
    fmt.Println("Запуск!")

    ch := make(chan int, 1)
    for i := 0; i < 10; i++ {
        fmt.Print("i = ", i)
        select {
        case x := <-ch:
            fmt.Println(" Приняли из канала", x)
        case ch <- i:
            fmt.Println(" Отправили i в канал")
        }
    }
}
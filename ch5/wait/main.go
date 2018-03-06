package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// WaitForServer пытается соединиться с сервером заданного URL.
// Попытки предпринимаются в течение минуты с растущими интервалами.
// Сообщает об ошибке, если все попытки неудачны
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // Успешное соединение
		}
		log.Printf("Сервер не отвечает (%s); повтор...", err)
		time.Sleep(time.Second << uint(tries)) // Увеличение задержки
	}
	return fmt.Errorf("Сервер %s не отвечает; время %s ", url, timeout)
}

func main() {
	url := "http://x.e2.ru"
	log.SetPrefix("wait: ")
	log.SetFlags(0)
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Сервер не работает: %v\n", err)
		os.Exit(1)
	}
}

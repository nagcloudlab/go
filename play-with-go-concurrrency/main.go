package main

import (
	"fmt"
	"time"
)

func logger(logs <-chan string) {
	for {
		fmt.Println("Waiting for logs...")
		select {
		case msg := <-logs:
			fmt.Println("LOG:", msg)
		case <-time.After(3 * time.Second):
			fmt.Println("No logs received in 3 seconds. Waiting...")
		}
	}
}

func main() {

	logs := make(chan string)
	go logger(logs)

	go func() {
		time.Sleep(2 * time.Second)
		logs <- "User logged in"
		time.Sleep(4 * time.Second)
		logs <- "User performed action"
	}()

	time.Sleep(10 * time.Second)

}

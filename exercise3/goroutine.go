package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	ch := make(chan string)
	go sendToMisterA(ch)
	messageFromMisterB := <-ch
	messageFromMisterC := <-ch
	if messageFromMisterB == "กำลังส่งของให้นาย A" {
		fmt.Println("นาย A ได้รับของแล้ว")
		fmt.Println(messageFromMisterC)
		fmt.Println("นาย A นั่งรอมา ", time.Since(startTime), " วินาที")
	}
}

func sendToMisterA(message chan<- string) {
	time.Sleep(1 * time.Second)
	message <- "กำลังส่งของให้นาย A"
}

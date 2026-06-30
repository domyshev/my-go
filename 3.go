package main

import "fmt"
import "time"

func worker(inCh chan int) {
	fmt.Println("я воркер пришел к вам на работу")
	time.Sleep(1 * time.Second)
	fmt.Println("это текст из воркера")
	inCh <- 12
}

func main() {
	x := make(chan int)
	fmt.Println("привет мир")	
	go worker(x)
	fmt.Println("текст сразу после вызова воркера")
	fmt.Println(<-x)
}

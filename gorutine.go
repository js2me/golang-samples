package main

import (
	"fmt"
	"strconv"
	"time"
	"math/rand"
)

func pinger(c chan string) {
	for i := 0; ; {
		randNumber := rand.Intn(15000)
		fmt.Println("randNumber : ",randNumber)
		i+=randNumber
		c <- "ping " + strconv.Itoa(i)
	}
}
func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 159)
	}
}
func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
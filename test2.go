package main

import (
	"fmt"
	"time"
)

//import (
//	"fmt"
////	"math"
//)

func time_operations(){

	fmt.Println("...")

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Выходные ура!")
	case time.Monday:
		fmt.Println("Сегодня понедельник :ССС")
	case time.Thursday:
		fmt.Println("Сегодня четверг, скоро пятница :З")
	case time.Friday:
		fmt.Println("Емае пятница")
	default:
		fmt.Println("Сегодня будние дни:(")
	}

	fmt.Println("...")

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Сейчас до полудня")
	default:
		fmt.Println("Сейчас после полудня")
	}
}


func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func learn_channels(){
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}


func learn_go_and_select(){

	// В нашем примере будем делать выбор из двух каналов.
	c1 := make(chan string)
	c2 := make(chan string)

	// Каждый канал будет получать значение после некоторого
	// промежутка времени, например, для имитации
	// блокирования RPC операций в одновременно
	// выполняющихся горутинах.
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	// Будем использовать `select` для ожидания обоих
	// значений одновременно и печати каждого по мере
	// получения.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func main() {
	defer fmt.Println("Функция main завершена")

	time_operations()
	learn_channels()
	learn_go_and_select()


	var trainSets = [4][2]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	for x := 0; x < len(trainSets); x++ {
		fmt.Println(trainSets[x])
	}
}

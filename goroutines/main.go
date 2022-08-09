package main

import (
	"fmt"
)

func worker0(done chan<- bool) {
	fmt.Println("working...0")
	fmt.Println("done")
	done <- true
}

func worker1(done chan<- bool) {
	fmt.Println("working...1")
	fmt.Println("done")
	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker0(done)
	go worker1(done)

	<-done
	<-done

}

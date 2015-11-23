package main

import (
	"fmt"
	"time"
)

func pt(a int, ci chan int) {
	time.Sleep(time.Duration(a) * time.Second)
	fmt.Printf("I'm %d\n", a)
	ci <- a
	fmt.Printf("%d push into the chan\n", a)
}

func main() {
	ci := make(chan int, 2)
	go pt(1, ci)
	go pt(2, ci)
	fmt.Println("a")
	time.Sleep(time.Duration(5) * time.Second)
	fmt.Println("b")
	<-ci
	fmt.Println("c")
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("d")
	<-ci
	fmt.Println("e")
	time.Sleep(time.Duration(5) * time.Second)
	fmt.Println("f")
}

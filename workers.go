package main

import (
	"fmt"
	"time"
)

const count = 10

type Numbers struct {
	Out chan int
}

func (n Numbers) DoWork() []int {
	res := make([]int, 0, 0)
	for i := 0; i < count; i++ {
		res = append(res, i)
	}
	return res
}

func (n Numbers) Start() {
	for _, i := range n.DoWork() {
		n.Out <- i
	}
}

type Fibonacci struct {
	In    chan int
	Out   chan int
	ToLog chan string
}

func (f Fibonacci) DoWork(i int) int {
	if i == 0 || i == 1 {
		return i
	}
	return (i - 1) + (i - 2)
}

func (f Fibonacci) Start() {
	for nmbr := range f.In {
		fnmbr := f.DoWork(nmbr)
		f.Out <- fnmbr
		f.ToLog <- fmt.Sprintf("Fibonacci number: %d", fnmbr)
	}
}

type Factorial struct {
	In    chan int
	ToLog chan string
}

func (f Factorial) DoWork(i int) int {
	if i > 0 {
		return i * f.DoWork(i-1)
	}
	return 1
}

func (f Factorial) Start() {
	for fnmbr := range f.In {
		fctrl := f.DoWork(fnmbr)
		f.ToLog <- fmt.Sprintf("Factorial of %d is: %d", fnmbr, fctrl)
	}
}

type Logger struct {
	In  chan string
	Out chan interface{}
}

func (l Logger) DoWork(msg string) {
	fmt.Println(msg)
}

func (l Logger) Start() {
	timer := time.NewTimer(1 * time.Second)
	for {
		select {
		case msg := <-l.In:
			l.DoWork(msg)
			timer.Reset(1 * time.Second)
		case <-timer.C:
			l.Out <- ""
			return
		}

	}
}

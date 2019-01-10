package main

import (
	"fmt"
	"time"
)

const count = 10

// Numbers produces numbers to create fibonacci numbers from
type Numbers struct {
	Out chan int
}

// DoWork is standart function for do some work.
// In this case it produces slice of numbers
func (n Numbers) DoWork() []int {
	res := make([]int, 0, 0)
	for i := 0; i < count; i++ {
		res = append(res, i)
	}
	return res
}

// Start is a standart function for receiving and sending data.
// In this case it sends numbers to channel
func (n Numbers) Start() {
	for _, i := range n.DoWork() {
		n.Out <- i
	}
}

// Fibonacci produce fibonacci numbers
type Fibonacci struct {
	In    chan int
	Out   chan int
	ToLog chan string
}

// DoWork is standart function for do some work.
// In this case it calculates fibonacci numbers
func (f Fibonacci) DoWork(i int) int {
	if i == 0 || i == 1 {
		return i
	}
	return (i - 1) + (i - 2)
}

// Start is a standart function for receiving and sending data.
// In this case it gets number from in channel, calculates fibonacci number for it and sends it to out channel.
// Also it sends data to logger
func (f Fibonacci) Start() {
	for nmbr := range f.In {
		fnmbr := f.DoWork(nmbr)
		f.Out <- fnmbr
		f.ToLog <- fmt.Sprintf("Fibonacci number: %d", fnmbr)
	}
}

// Factorial calculates factorial
type Factorial struct {
	In    chan int
	ToLog chan string
}

// DoWork is standart function for do some work.
// In this case it calculates factorial of number
func (f Factorial) DoWork(i int) int {
	if i > 0 {
		return i * f.DoWork(i-1)
	}
	return 1
}

// Start is a standart function for receiving and sending data.
// In this case it gets number from in channel and calculates its factorial.
// Also it sends data to logger
func (f Factorial) Start() {
	for fnmbr := range f.In {
		fctrl := f.DoWork(fnmbr)
		f.ToLog <- fmt.Sprintf("Factorial of %d is: %d", fnmbr, fctrl)
	}
}

// Logger prints incoming data
type Logger struct {
	In  chan string
	Out chan interface{}
}

// DoWork is standart function for do some work.
// In this case it prints incoming string
func (l Logger) DoWork(msg string) {
	fmt.Println(msg)
}

// Start is a standart function for receiving and sending data.
// In this case it prints incoming.
// Also it determines the finishing of calculations by timer
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

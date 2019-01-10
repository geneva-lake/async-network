package main

func start(endChan chan interface{}) {
	fbncChn := make(chan int, count)
	fctrlChn := make(chan int, count)
	lgChn := make(chan string, 2*count)
	nmbrs := Numbers{fbncChn}
	fbncs := make([]*Fibonacci, 0, 0)
	for i := 0; i < count; i++ {
		fbnc := &Fibonacci{
			In:    fbncChn,
			Out:   fctrlChn,
			ToLog: lgChn,
		}
		fbncs = append(fbncs, fbnc)
	}
	fctrls := make([]*Factorial, 0, 0)
	for i := 0; i < count; i++ {
		fctrl := &Factorial{
			In:    fctrlChn,
			ToLog: lgChn,
		}
		fctrls = append(fctrls, fctrl)
	}

	lgr := Logger{
		In:  lgChn,
		Out: endChan,
	}
	go lgr.Start()
	for _, fctrl := range fctrls {
		go fctrl.Start()
	}
	//go fbnc.Start()
	for _, fbnc := range fbncs {
		go fbnc.Start()
	}
	go nmbrs.Start()
}

// func fibonacci(out chan interface{}) {
// 	for i := 0; i < 10; i++ {
// 		if i == 0 || i == 1 {
// 			out <- i
// 		} else {
// 			out <- (i - 1) + (i - 2)
// 		}
// 	}
// }

// func factorial(n int) int {
// 	if n > 0 {
// 		return n * factorial(n-1)
// 	}
// 	return 1
// }

// func logging(in chan interface{}) {
// 	for msg := range in {
// 		fmt.Println(msg.(*Dto).Name, msg.(*Dto).Data)
// 	}
// }

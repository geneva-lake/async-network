package main

func start(endChan chan interface{}) {
	// creation channels for communication between modules
	fbncChn := make(chan int)
	fctrlChn := make(chan int)
	lgChn := make(chan string, 2*count)

	// creation calculation modules
	nmbrs := Numbers{fbncChn}

	// creation slice of Fibonacci modules for asynchronous
	// calculation fibonacci numbers
	fbncs := make([]*Fibonacci, 0, 0)
	for i := 0; i < count; i++ {
		fbnc := &Fibonacci{
			In:    fbncChn,
			Out:   fctrlChn,
			ToLog: lgChn,
		}
		fbncs = append(fbncs, fbnc)
	}

	// creation slice of Factorial modules for asynchronous
	// calculation of factorials
	fctrls := make([]*Factorial, 0, 0)
	for i := 0; i < count; i++ {
		fctrl := &Factorial{
			In:    fctrlChn,
			ToLog: lgChn,
		}
		fctrls = append(fctrls, fctrl)
	}

	// creation logger
	lgr := Logger{
		In:  lgChn,
		Out: endChan,
	}

	// start goroutines containing functions for calculations
	go lgr.Start()
	for _, fctrl := range fctrls {
		go fctrl.Start()
	}
	for _, fbnc := range fbncs {
		go fbnc.Start()
	}
	go nmbrs.Start()
}

This is example of asynchronous calculation network (async network pattern) powered by golang goroutines.<br/>
We have initial data as ten integers from 0 to 9. We need calculate for each of them sum of two previous numbers (Fibonacci number) and then find factorial of resulting number. Also we want log results. Construct computational network. The first module simulate initial data and produce numbers from 0 to 9. Through the channel it sends numbers to calculation layer which finds Fibonacci number for received integer. It consists of ten modules. Second calculation layer finds factorials. Also there is logger module. It gets results from both layers and prints it. The scheme is depicted on figure.<br/>
<br/>
<br/>
<img style="width: 80%; margin-top: 50px;" src="https://github.com/geneva-lake/async-network/blob/master/network.png"/>
<br/>
<br/>
Let’s code. Initial element is realized as Numbers struct

```
type Numbers struct {
    Out chan int
}

func (n Numbers) Start() {
    for _, i := range n.DoWork() {
        n.Out <- i
    }
}
```

DoWork function produce slice of ten integers. 

Modules in first layer are realized as Fibonacci struct. It has incoming channel for receiving data, out coming channel for result and functions. Function DoWork find Fibonacci number and function Start gets data from channel, calls DoWork and send result to next computation layer and string to logger. 

```
type Fibonacci struct {
    In chan int
    Out chan int
    ToLog chan string
}

func (f Fibonacci) Start() {
    for nmbr := range f.In {
        fnmbr := f.DoWork(nmbr)
        f.Out <- fnmbr
        f.ToLog <- fmt.Sprintf("Fibonacci number: %d", fnmbr)
    }
}
```

Second calculation level organized similarly. 

```
type Factorial struct {
    In chan int
    ToLog chan string
}

func (f Factorial) Start() {
    for fnmbr := range f.In {
        fctrl := f.DoWork(fnmbr)
        f.ToLog <- fmt.Sprintf("Factorial of %d is: %d", fnmbr, fctrl)
    }
}
```


So we need make channels and create modules

```
    fbncChn := make(chan int)
    fctrlChn := make(chan int)
    lgChn := make(chan string, 20)
    nmbrs := Numbers{fbncChn}

    fbncs := make([]*Fibonacci, 0, 0)
    for i := 0; i < 10; i++ {
        fbnc := &Fibonacci{
            In: fbncChn,
            Out: fctrlChn,
            ToLog: lgChn,
        }
        fbncs = append(fbncs, fbnc)
    }

fctrls := make([]*Factorial, 0, 0)
    for i := 0; i < 10; i++ {
        fctrl := &Factorial{
            In: fctrlChn,
            ToLog: lgChn,
        }
        fctrls = append(fctrls, fctrl)
    }

    lgr := Logger{
        In: lgChn,
        Out: endChan,
    }
```


And run Start functions in goroutines from end to beginning

```
go lgr.Start()
    for _, fctrl := range fctrls {
        go fctrl.Start()
    }
    for _, fbnc := range fbncs {
        go fbnc.Start()
    }
    go nmbrs.Start()
```
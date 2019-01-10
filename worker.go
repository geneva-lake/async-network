package main

type Dto struct {
	Name string
	Data interface{}
}

// type Worker struct {
// 	Name     string
// 	In       chan interface{}
// 	Out      []chan interface{}
// 	Function func(interface{}) interface{}
// }

// func (w Worker) Start() {
// 	for msg := range w.In {
// 		ans := w.Function(msg)
// 		for _, c := range w.Out {
// 			c <- &Dto{w.Name, ans}
// 		}
// 	}
// }

type Worker interface {
	DoWork(interface{}) interface{}
	Start()
}

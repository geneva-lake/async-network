package main

func main() {
	endChan := make(chan interface{})
	start(endChan)
	<-endChan
}

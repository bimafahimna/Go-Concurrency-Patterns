package main

import "sync"

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	// go SingleGeneratorImplementation(&wg)
	// go MultipleGeneratorImplementation(&wg)
	go FanOutFanInImplementation(&wg)

	wg.Wait()
}

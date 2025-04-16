package main

import "sync"

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go singleGeneratorImplementation(&wg)
	go multipleGeneratorImplementation(&wg)

	wg.Wait()
}

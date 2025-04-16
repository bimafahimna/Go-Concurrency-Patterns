package main

import (
	"basic_patterns/generator"
	"fmt"
	"sync"
	"time"
)

func singleGeneratorImplementation(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Start of Single Generator Implementation")

	res := generator.Generator(6)
	for val := range res {
		fmt.Println("S Gen:", val)
	}

	fmt.Println("End of Single Generator Implementation")
}

func multipleGeneratorImplementation(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Start of Multiple Generator Implementation")

	res1 := generator.Generator(6)
	time.Sleep(1 * time.Second)
	res2 := generator.Generator(11)
	for {
		select {
		case val, ok := <-res1:
			if !ok {
				res1 = nil
			} else {
				fmt.Println("M Gen 1:", val)
			}
		case val, ok := <-res2:
			if !ok {
				res2 = nil
			} else {
				fmt.Println("M Gen 2:", val)
			}
		}

		if res1 == nil && res2 == nil {
			break
		}
	}

	fmt.Println("End of Multiple Generator Implementation")
}

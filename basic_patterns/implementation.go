package main

import (
	"basic_patterns/generator"
	"fmt"
	"time"
)

func singleGeneratorImplementation() {
	res := generator.Generator(6)
	for val := range res {
		fmt.Println(val)
	}
}

func multipleGeneratorImplementation() {
	res1 := generator.Generator(6)
	time.Sleep(1 * time.Second)
	res2 := generator.Generator(11)
	for {
		select {
		case val, ok := <-res1:
			if !ok {
				res1 = nil
			} else {
				fmt.Println(val)
			}
		case val, ok := <-res2:
			if !ok {
				res2 = nil
			} else {
				fmt.Println(val)
			}
		}

		if res1 == nil && res2 == nil {
			break
		}
	}
}

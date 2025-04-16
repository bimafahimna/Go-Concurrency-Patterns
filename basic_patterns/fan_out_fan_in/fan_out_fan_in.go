package fanoutfanin

import (
	"fmt"
	"sync"
)

// Fan Out Fan In function

// FanOut split received data to n workers and do concurrent process
func FanOut(in <-chan any, n int) []chan any {
	channels := []chan any{}

	for i := 0; i < n; i++ {
		channel := make(chan any)
		go func() {
			for val := range in {
				switch x := val.(type) {
				case int:
					channel <- fmt.Sprintf("wrap generator results: %d", x)
				case string:
					channel <- x[len(x)-1]
				}
			}
			close(channel)
		}()
		channels = append(channels, channel)
	}

	return channels
}

// FanIn function that only process string value from mutiple goroutines
func FanIn(channels ...chan any) <-chan string {
	out := make(chan string, len(channels))

	var wg sync.WaitGroup
	for _, channel := range channels {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for val := range channel {
				switch x := val.(type) {
				case string:
					out <- fmt.Sprintf("string from Fan Out Process: %s", x)
				default:
					out <- "error, not a string"
				}
			}
		}()
	}

	// close channel after all goroutines finished processing
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

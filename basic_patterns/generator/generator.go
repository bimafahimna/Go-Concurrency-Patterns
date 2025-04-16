package generator

import (
	"fmt"
	"math/rand"
	"time"
)

// Generator func which produces data which is computationally expensive.
func Generator(n int) <-chan any {
	out := make(chan any)

	go func() {
		// imitating expensive computation when do query from database
		for i := 0; i < n; i++ {
			time.Sleep(2 * time.Second)

			if i%2 == 0 {
				// generating integer from 10 to 100
				out <- rand.Intn(100-10) + 10
			} else {
				// generating string with odd number
				out <- fmt.Sprintf("Odd number: %d", i)
			}
		}

		// close open channel
		close(out)
	}()
	return out
}

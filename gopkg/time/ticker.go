package time

import (
	"fmt"
	"time"
)

func NewTicker() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	var counter = 0
	for {
		select {
		case <-ticker.C:
			counter += 1
		}

		fmt.Printf("counter: %d", counter)
	}
}

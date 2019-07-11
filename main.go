package main

import (
	"github.com/robfig/cron"
	"math"
	"time"
)

func sum(id int) {
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}

	println(id, x)
}

func main() {
	c := cron.New()

	time.Parse()

	time1.Add(1 * time.Month())

	time1.AddDate

}
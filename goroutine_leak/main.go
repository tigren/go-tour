package main

import (
	"runtime/pprof"
	"time"
)

func potentiallyTimeConsumingCall(result chan struct{}) {
	<-time.After(10 * time.Millisecond)
	result <- struct{}{}
}

func runALotOfTasks() int {
	r := make(chan struct{})
	for i := 0; i < 100; i++ {
		go potentiallyTimeConsumingCall(r)
		select {
		case <-r:
		case <-time.After(5 * time.Millisecond):
		}
	}

	return pprof.Lookup("goroutine").Count()
}

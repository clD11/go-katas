package scratchpad

import (
	"log"
	"sync/atomic"
)

func GoRoutineChannelStop() {
	q := make(chan bool)
	go func(msg string) {
		for {
			select {
			case <-q:
				return
			default:
				log.Print(msg)
			}
		}
	}("test")
}

func CompareAndSwapInt32(addr *int32, old, new int32) bool {
	return atomic.CompareAndSwapInt32(addr, old, new)
}

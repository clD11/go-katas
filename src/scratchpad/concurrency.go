package scratchpad

import (
	"log"
	"sync/atomic"
)

type Mutant struct {
	name    string
	version int
}

func (m Mutant) Version(i int) {
	m.version += 1
}

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

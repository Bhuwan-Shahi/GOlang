package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	count int
	lock  sync.Mutex
}

func cont(counter *Counter) {
	counter.lock.Lock()
	defer counter.lock.Unlock()
	counter.count++
	fmt.Println(counter.count)
}
func main() {
	counter := Counter{}
	for i := 0; i < 100; i++ {
		go cont(&counter)
	}
	time.Sleep(100 * time.Millisecond)
}

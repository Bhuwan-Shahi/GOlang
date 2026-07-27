package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	lock  sync.Mutex
}

func cont(counter *Counter, wg *sync.WaitGroup) {
	counter.lock.Lock()
	defer counter.lock.Unlock()
	counter.count++
	fmt.Println(counter.count)
	wg.Done()
}
func main() {
	counter := Counter{}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go cont(&counter, &wg)
	}
	wg.Wait()
}

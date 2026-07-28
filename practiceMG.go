package main

import (
	"fmt"
	"sync"
)

type counts struct {
	counter int
	lock    sync.Mutex
}

func Increment(c *counts, wg *sync.WaitGroup) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.counter++
	fmt.Println(c.counter)
	wg.Done()

}
func main() {
	c := counts{}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go Increment(&c, &wg)
	}
	wg.Wait()
}

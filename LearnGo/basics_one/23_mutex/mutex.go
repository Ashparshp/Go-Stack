package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Inc(wg *sync.WaitGroup) {
	
	defer func ()  {
		wg.Done()
		c.mu.Unlock()
	} ()


	c.mu.Lock()
	c.value++
	
}

func main() {

	var wg sync.WaitGroup
	count := Counter{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go count.Inc(&wg)
	}

	wg.Wait()
	fmt.Println(count.value)

}
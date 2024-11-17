package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}

	wg.Wait()
}

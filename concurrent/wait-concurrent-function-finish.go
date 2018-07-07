package concurrent

import (
	"fmt"
	"sync"
)

func init() {
	fmt.Println("===== Waiting for All Concurrent Functions to Finish =====")

	var wg sync.WaitGroup
	for i :=0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("Hello World")
			wg.Done()
		}()
	}

	wg.Wait()

	//time.Sleep(10 * time.Second)
}

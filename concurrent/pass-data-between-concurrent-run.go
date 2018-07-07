package concurrent

import (
	"fmt"
)

func init() {
	fmt.Println("===== Passing Data Between Concurrently Running Functions =====")
	nameChannel := make(chan string, 3)
	done := make(chan string)

	go func() {
		names := []string {"tarik", "john", "michael", "jessica"}
		for _, name := range names {
			// doing some operation
			fmt.Println("Processing the first stage of :" + name)
			nameChannel <- name
		}
		close(nameChannel)
	}()

	go func() {
		for name := range nameChannel {
			fmt.Println("Processing the second stage of :" + name)
		}
		done <- ""
	}()

	//time.Sleep(time.Second * 5)
	<- done
}

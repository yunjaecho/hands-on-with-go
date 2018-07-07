package concurrent

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("====== Running Multiple Functions Concurrently =====")

	nameChannel := make(chan string)
	go func() {
		names := []string{"tarik", "john", "michael", "jessica"}
		for _, name := range names {
			time.Sleep(time.Second * 1)
			//fmt.Println(name)
			nameChannel <- name
		}
		close(nameChannel)
	}()

	for data := range nameChannel {
		fmt.Println(data)
	}


	/*go func() {
		ages := []int{1,2, 3, 4}
		for _, age := range ages {
			time.Sleep(time.Second * 1)
			fmt.Println(age)
		}
	}()

	time.Sleep(time.Second * 5)*/


}

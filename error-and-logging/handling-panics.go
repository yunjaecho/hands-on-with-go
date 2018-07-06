package error_and_logging

import "fmt"

func writeSomething() {
	panic("Write operation error")
}

func init() {
	fmt.Println("===== Handling Panics =====")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	writeSomething()
}

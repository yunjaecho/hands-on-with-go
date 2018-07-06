package error_and_logging

import "fmt"

type MyError struct {
	ShortMessage string
	DetailedMessage string
}

func (e *MyError) Error() string {
	return e.ShortMessage + "\n" + e.DetailedMessage
}

func doSomething() error {
	return &MyError{ShortMessage:"Wohoo something happend!", DetailedMessage:"File cannot find"}
}

func init() {
	fmt.Println("===== Create custom error types =====")
	err := doSomething()
	fmt.Println(err)

}

package file_directory

import (
	"fmt"
	"io/ioutil"
)

func init() {
	fmt.Println("===== Write File =====")

	helloWorld := "Hello, World"
	err := ioutil.WriteFile("hello_world",[]byte(helloWorld), 0644 )
	if err != nil {
		fmt.Println(err)
	}
}

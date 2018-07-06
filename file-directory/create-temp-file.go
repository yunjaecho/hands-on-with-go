package file_directory

import (
	"fmt"
	"io/ioutil"
)

func init() {
	fmt.Println("===== Creating temp files =====")
	helloWorld := "Hello, World"
	file, err := ioutil.TempFile("", "hello_world_temp")
	if err != nil {
		panic(err)
	}

	//defer os.Remove(file.Name())

	if _, err := file.Write([]byte(helloWorld)); err != nil {
		panic(err)
	}

	fmt.Println(file.Name())
}
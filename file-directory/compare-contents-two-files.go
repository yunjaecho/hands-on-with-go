package file_directory

import (
	"fmt"
	"io/ioutil"
	"bytes"
)

func init() {
	fmt.Println("===== Comparing the Contents of Two Files =====")

	one, err1 := ioutil.ReadFile("one.txt")
	if err1 != nil {
		panic(err1)
	}

	two, err2 := ioutil.ReadFile("two.txt")
	if err2 != nil {
		panic(err2)
	}

	same := bytes.Equal(one, two)
	fmt.Println(same)
}

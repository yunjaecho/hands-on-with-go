package file_directory

import (
	"fmt"
	"io/ioutil"
)

func init() {
	fmt.Println("===== Reading the Entire of a Text File =====")
	contentBytes, err := ioutil.ReadFile("log_file")

	if err == nil {
		var contentStr string = string(contentBytes)
		fmt.Println(contentStr)
	}
}

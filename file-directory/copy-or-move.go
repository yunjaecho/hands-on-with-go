package file_directory

import (
	"fmt"
	"os"
	"io"
)

func init() {
	fmt.Println("===== Copying or Moving =====")

	original, err := os.Open("original.txt")
	defer original.Close()
	if err != nil {
		panic(err)
	}


	original_copy, err2 := os.Create("target/original.txt")
	defer original_copy.Close()
	if err2 != nil {
		panic(err2)
	}

	_, err3 := io.Copy(original_copy, original)
	if err3 != nil {
		panic(err3)
	}

}

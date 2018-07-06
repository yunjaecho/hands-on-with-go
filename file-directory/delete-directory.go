package file_directory

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("===== Deleting a Directory and Its Contents =====")

	//err := os.Remove("hello")
	err := os.RemoveAll("hello")
	if err != nil {
		fmt.Println(err)
	}

}


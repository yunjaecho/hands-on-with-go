package file_directory

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("===== Checking the Existence of a File =====")

	if _, err := os.Stat("log_file"); os.IsNotExist(err) {
		fmt.Println("log_file file dose not exists")
	} else {
		fmt.Println("log_file file exists")
	}
}
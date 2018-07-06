package file_directory

import (
	"fmt"
	"os"
	"bufio"
)

func init() {
	fmt.Println("===== Counting lines of a file =====")
	file, _ := os.Open("log_file")

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}


	fmt.Println(lineCount)
}

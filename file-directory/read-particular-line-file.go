package file_directory

import (
	"fmt"
	"os"
	"bufio"
)

func init() {
	fmt.Println("===== Reading a Praticular Line in a file  =====")
	fmt.Println(ReadLine(2))

}

func ReadLine(lineNumber int) string {
	file, _ := os.Open("log_file")

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	lineCount := 1
	for fileScanner.Scan() {
		if lineCount == lineNumber {
			return fileScanner.Text()
		}
		lineCount++
	}

	return ""
}

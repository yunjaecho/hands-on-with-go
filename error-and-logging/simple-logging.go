package error_and_logging

import (
	"fmt"
	"os"
	"log"
)

func init() {
	fmt.Println("===== Doing a Simple Logging in Your App =====")

	log_file, err := os.Create("log_file")

	if err != nil {
		fmt.Println("An error occured")
	}

	defer log_file.Close()

	log.SetOutput(log_file)

	log.Println("Doing some logging here....")
	log.Fatalln("Fatal Application crashed!")
}
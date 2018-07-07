package system

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("===== Processing Command-line Arguments =====")

	realArgs := os.Args[1:]
	fmt.Println(realArgs)


}

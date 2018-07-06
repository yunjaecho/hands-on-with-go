package file_directory

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("===== Renameing Files ======")
	os.Rename("old.txt", "new.txt")
}

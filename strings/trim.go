package strings

import (
	"fmt"
	"strings"
)

func init() {
	fmt.Println("====== String Trim =====")
	greetings := "\t Hello World "
	fmt.Printf("%s is length %d\n", greetings, len(greetings))
	fmt.Printf("%s is length %d\n", strings.TrimSpace(greetings), len(strings.TrimSpace(greetings)))
}
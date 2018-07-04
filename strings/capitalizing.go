package strings

import (
	"fmt"
	"strings"
)

func init() {
	fmt.Println("====== Capitalizing ======")
	helloWorld := "hello world, how are you today"
	helloWorldTitle := strings.Title(helloWorld)
	fmt.Println(helloWorldTitle)

	helloworldUpper := strings.ToUpper(helloWorld)
	fmt.Println(helloworldUpper)

}

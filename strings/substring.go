package strings

import (
	"fmt"
)

func init() {
	fmt.Println("====== String Substring =====")
	greetings := "Hello World and Mars"
	helloWorld := greetings[0: 12]
	fmt.Println(helloWorld)

}
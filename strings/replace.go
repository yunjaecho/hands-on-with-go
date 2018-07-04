package strings

import (
	"strings"
	"fmt"
)

func init() {
	fmt.Println("===== String Replace =====")
	helloWorld := "Hello, World. How are you World, I am good, thanks World"
	//helloMars := strings.Replace(helloWorld, "World", "Mars", 1)
	helloMars := strings.Replace(helloWorld, "World", "Mars", -1)
	fmt.Println(helloMars)


}
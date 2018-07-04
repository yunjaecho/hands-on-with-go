package strings

import "fmt"

func init() {
	fmt.Println("===== String Escaping =====")
	helloWorld := "Hello World, this is \"Tarik. \" \t\nHello again"
	fmt.Println(helloWorld)
}
package typeConversions

import "fmt"

func init() {
	fmt.Println("===== Type Convert Byte Array to String =====")
	helloWordByte := []byte {72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100}
	fmt.Println(string(helloWordByte))
	fmt.Println([]byte(string(helloWordByte)))
}
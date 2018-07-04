package typeConversions

import (
	"fmt"
	"strconv"
)

func init() {
	fmt.Println("===== Type Convert Boolean to String =====")
	isNew := true
	message := "Purchased item is " + strconv.FormatBool(isNew)
	fmt.Println(message)
}
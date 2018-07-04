package typeConversions

import (
	"fmt"
	"strconv"
)

func init() {
	fmt.Println("===== Type Convert Number to String =====")
	number := int64(1000)
	numberStr := strconv.FormatInt(number, 10)
	fmt.Println(numberStr)


	number2 := 1000
	numberStr2 := strconv.Itoa(number2)
	fmt.Println(numberStr2)

	numberFloat := 34234234.12312
	numberFloatStr := strconv.FormatFloat(numberFloat, 'f', -1, 64)
	fmt.Println(numberFloatStr)
}

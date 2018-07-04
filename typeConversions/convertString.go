package typeConversions

import (
	"fmt"
	"strconv"
)

func init() {
	fmt.Println("===== Type Convert String to Boolean =====")
	isNew := "true"
	isNewBoolean, err := strconv.ParseBool(isNew)
	if(err != nil) {
		fmt.Println("failed")
	} else {
		if (isNewBoolean) {
			fmt.Println("IsNew")
		} else {
			fmt.Println("Not New")
		}
	}

	fmt.Println("===== Type Convert String to Integer =====")
	number := "2"
	valueInt, err := strconv.ParseInt(number, 10, 64)

	if err != nil {
		fmt.Println("Error happended...")
	} else {
		if valueInt == 2 {
			fmt.Println("Success")
		}
	}


	fmt.Println("===== Type Convert String to Float =====")
	numberFloat := "2.2"
	valueFloat, errFloat := strconv.ParseFloat(numberFloat, 64)

	if errFloat != nil {
		fmt.Println("Error happended...")
	} else {
		if valueFloat == 2.2 {
			fmt.Println("Success")
		}
	}

}

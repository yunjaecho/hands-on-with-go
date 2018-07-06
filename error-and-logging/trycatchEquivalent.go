package error_and_logging

import (
	"fmt"
	"time"
	"errors"
)



func init() {
	fmt.Println("===== Try/Catch Equivalent in Go =====")
	parseDate, err := time.Parse("2006", "2008 abc")

	if (err != nil) {
		fmt.Println("An error occured", err.Error())
	} else {
		fmt.Println(parseDate)
	}

	_, err2 := doSometing2()
	if err2 != nil {
		fmt.Println(err2)
	}
}

func doSometing2() (string, error)  {
	return "", errors.New("Something happened")

}

package dateAndTime

import (
	"fmt"
	"time"
)

/**
	2006-01-02 15:04:06": Since MST is GMT-0700, the reference time can be thought of as
 */
func init() {
	fmt.Println("===== Time =====")
	current := time.Now()
	fmt.Println(current.String())
	fmt.Println(current.Format("2006"))
	fmt.Println(current.Format("01"))
	fmt.Println(current.Format("02"))

	fmt.Println("MM-DD-YYYY : " , current.Format("01-02-2006"))
	fmt.Println("YYYY-MM-DD hh:mm:ss  : ", current.Format("2006-01-02 15:04:06"))

	aprilDate := current.AddDate(1, 1, 1)
	fmt.Println(aprilDate.String())

	tenMoreMinutes := aprilDate.Add(10 * time.Hour)
	fmt.Println(tenMoreMinutes.String())

	//  Finding the Difference Between Two Dates
	first := time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC)
	second := time.Date(2011, 1, 1, 0, 0, 0, 0, time.UTC)

	difference := second.Sub(first)
	fmt.Printf("Difference %v\n", difference)
	fmt.Printf("Difference %v\n", difference.Hours() / 24)

	// Parsing Dates and Times from Strings
	str := "2018-03-12T11:45:26.371Z"
	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, str)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t.String())

}


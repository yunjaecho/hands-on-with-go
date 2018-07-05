package map_array

import (
	"fmt"
	"sort"
)

func init() {
	fmt.Println("===== Finding an Element from an Array =====")

	str :=[]string{"Sandy", "Provo", "St.", "George", "Salt Lak City", "Draper", "South Jordan", "Murray"}

	for i, v := range str {
		if v == "Sandy" {
			fmt.Println(i)
		}
	}

	sortedList := sort.StringSlice(str)
	sortedList.Sort()
	index := sortedList.Search("Provo")
	fmt.Println(sortedList)
	fmt.Println(index)
}

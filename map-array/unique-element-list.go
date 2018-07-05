package map_array

import "fmt"

func init() {
	fmt.Println("===== Unique Elements in a list =====")
	intSlice := []int{1,5,5,5,7,8,6,6,6}
	fmt.Println(intSlice)

	uniqueIntSlice := unique(intSlice)
	fmt.Println(uniqueIntSlice)
}

func unique(intSlice []int) interface{} {
	keys := make(map[int]bool)
	uniqueElements := []int{}

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			uniqueElements = append(uniqueElements, entry)
		}
	}
	return uniqueElements
}

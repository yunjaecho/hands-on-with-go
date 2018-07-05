package map_array

import "fmt"

func init() {
	fmt.Println("===== Merging Arrays =====")
	item1 := []int{3,4}
	item2 := []int{1,2}

	result := append(item1, item2...)
	fmt.Println(result)
}
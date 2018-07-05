package map_array

import "fmt"

type NameAge struct {
	Name string
	Age int
}

func init() {
	fmt.Println("===== Convert Map to keys value list =====")
	var nameAgeSlice []NameAge

	nameAges := map[string]int {
		"Michanel": 30,
		"John": 25,
		"Jessica": 36,
		"Ali": 28,
	}

	for key, value := range nameAges {
		nameAgeSlice = append(nameAgeSlice, NameAge{key, value})
	}

	fmt.Println(nameAgeSlice)

}

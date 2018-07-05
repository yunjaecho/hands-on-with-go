package map_array

import "fmt"

func init() {
	fmt.Println("===== Merging Maps =====")
	map1 := map[string]int {
		"Michanel": 30,
		"John": 25,
		"Jessica": 36,
		"Ali": 28,
	}

	map2 := map[string]int {
		"Lord": 11,
		"Of": 22,
		"The": 36,
		"Rings": 23,
	}

	for key, value := range map2 {
		map1[key] = value
	}

	fmt.Println(map1)
}

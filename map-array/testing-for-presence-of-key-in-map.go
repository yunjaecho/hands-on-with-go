package map_array

import "fmt"

func init() {
	fmt.Println("===== Testing for presence of key in a map =====")

	nameAges := map[string]int {
		"Michanel": 30,
		"John": 25,
		"Jessica": 36,
		"Ali": 28,
	}

	fmt.Println(nameAges["Michanel"])
	value, exists := nameAges["Tarik"]
	fmt.Println(value)
	fmt.Println(exists)

	if value , exists := nameAges["Jessica"]; exists {
		fmt.Println("Jessica value : ", value)
	} else {
		fmt.Println("Jessica cannot be found ")
	}
}

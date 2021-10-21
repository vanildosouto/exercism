package raindrops

import (
	"sort"
	"strconv"
)

// Convert return a raindrops strings
func Convert(number int) (raindrops string) {
	factors := map[int]string{3: "Pling", 5: "Plang", 7: "Plong"}
	keys := []int{}

	// Ordenando o map
	for k := range factors {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, index := range keys {
		if (number % index) == 0 {
			raindrops += factors[index]
		}
	}

	if raindrops == "" {
		raindrops = strconv.Itoa(number)
	}

	return
}

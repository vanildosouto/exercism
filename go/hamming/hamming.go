package hamming

import (
	"errors"
)

// Distance calcule the Hamming Distance
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("a and b are different")
	}

	distanceCount := 0

	for index := range a {
		if a[index] != b[index] {
			distanceCount++
		}
	}

	return distanceCount, nil
}

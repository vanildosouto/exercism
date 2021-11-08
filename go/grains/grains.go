package grains

import (
	"errors"
)

func Square(number int) (square uint64, err error) {
	if number > 64 || number < 1 {
		return square, errors.New("number out of scale")
	}

	square = 1
	for x := 1; x < number; x++ {
		square = square * 2
	}

	return
}

func Total() (total uint64) {
	for x := 1; x <= 64; x++ {
		square, _ := Square(x)
		total += square
	}

	return
}

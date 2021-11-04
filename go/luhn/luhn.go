package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	withoutSpace := strings.Replace(id, " ", "", -1)

	if len(withoutSpace) < 2 {
		return false
	}

	match, _ := regexp.Match(`\D`, []byte(withoutSpace))

	if match {
		return false
	}

	reversed := reverse(withoutSpace)
	sum := 0

	for idx, r := range reversed {
		number, _ := strconv.Atoi(string(r))
		if (idx % 2) != 0 {
			double := number * 2
			if double > 9 {
				double -= 9
			}
			sum += double
		} else {
			sum += number
		}
	}

	return (sum % 10) == 0
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}

	return
}

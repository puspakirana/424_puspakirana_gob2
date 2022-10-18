package helper

import "strconv"

func myFunction() string {

	result := ""

	for i := 0; i <= 10; i++ {
		if i%3 == 0 {
			result += strconv.Itoa(i) + " "
		}
	}

	return result
}

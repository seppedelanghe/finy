package utils

import (
	"strconv"
)


func StringDecimalInt(s string) int {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}

	return int(f * 100)
}

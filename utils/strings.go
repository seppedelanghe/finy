package utils

import (
	"strconv"
	"strings"
)


func StringDecimalInt(s string) int {
	v := strings.ReplaceAll(strings.ReplaceAll(s, ".", ""), ",", "")
	i, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return i
}

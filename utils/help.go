package utils

import "strconv"

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func IntToString(a int) string {
	return strconv.Itoa(a)
}

package utils

import "strconv"

// parse string to uint
func ParseUint(s string) uint {
	n, _ := strconv.ParseUint(s, 10, 64)
	return uint(n)
}

// parse string to int
func ParseInt(s string) int {
	n, _ := strconv.ParseInt(s, 10, 64)
	return int(n)
}

func FormatUint(i uint) (s string) {
	s = strconv.FormatUint(uint64(i), 10)
	return s
}

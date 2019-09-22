package utils

import "strconv"

// parse string to uint
func ParseUint(s string) (n uint64, err error) {
	n, err = strconv.ParseUint(s, 10, 64)
	return n, err
}

func FormatUint(i uint) (s string) {
	s = strconv.FormatUint(uint64(i), 10)
	return s
}

package main

import "strconv"

func stringToInt64(s string) (int64, error) {
	no, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, err
	}
	return no, err
}
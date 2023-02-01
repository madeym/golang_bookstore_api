package helper

import (
	"strconv"
)

func StringToInt(str string) (int, error) {
	s, err := strconv.Atoi(str)
	if err != nil {
		return -1, err
	}
	s = int(s)
	return s, nil
}

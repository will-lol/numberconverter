package util

import (
	"math"
)

type Integer interface {
	int8 | int16 | int32 | int64 | int
}

func Insert[T any](a []T, index int, value T) []T {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func GetDigitLength[T Integer](num int64) T {
	var l T = 1
	if num == 0 {
		return l
	}
	return l + T(math.Log10(float64(num)))
}


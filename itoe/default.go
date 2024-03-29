package itoe

import (
	"math"
	"strings"

	"github.com/will-lol/numberconverter/util"
)

func DigitArrToString(digits []int, negative bool) string {
	arrs := splitArr(digits)

	words := make([]string, 0, len(arrs))

	if negative {
		words = append(words, "negative")
	}

	for i, arr := range arrs {
		words = append(words, fragmentToStrings(arr)...)
		if len(arrs)-i-1 > 0 {
			words = append(words, itoeNumbers[len(arrs)-i-1])
		}
	}

	return strings.Join(words, " ")
}

func fragmentToStrings(arr []int) []string {
	out := make([]string, 0, 3)

	hundreds := 0
	tens := 0
	ones := 0
	if len(arr) == 3 {
		hundreds = arr[0]
		tens = arr[1]
		ones = arr[2]
	} else if len(arr) == 2 {
		tens = arr[0]
		ones = arr[1]
	} else if len(arr) == 1 {
		ones = arr[0]
	} else {
		return []string{"zero"}
	}

	// hundreds
	if hundreds != 0 {
		out = append(out, itoeUniques[hundreds]+" hundred")
	}

	// tens and ones
	tensAndOnes := make([]string, 0, 2)
	lessThanTwenty := 0
	if tens != 0 {
		if tens == 1 {
			lessThanTwenty = 10
		} else {
			tensAndOnes = append(tensAndOnes, itoeTens[tens])
		}
	}
	lessThanTwenty += ones
	if lessThanTwenty != 0 {
		tensAndOnes = append(tensAndOnes, itoeUniques[lessThanTwenty])
	}

	if len(tensAndOnes) > 0 {
		out = append(out, strings.Join(tensAndOnes, "-"))
	}
	return out
}

func splitArr(arr []int) [][]int {
	l := int(math.Ceil(float64(len(arr)) / 3.0))
	out := make([][]int, l, l)

	offset := len(arr) % 3
	if offset == 0 {
		offset = 3
	}

	for i := range out {
		low := i*3 - (3 - offset)
		high := low + 3
		out[i] = arr[max(low, 0):min(high, len(arr))]
	}

	return out
}

func ToDigitArr(num int64) []int {
	length := util.GetDigitLength[int](num)
	arr := make([]int, length, length)

	prev := math.Inf(1)
	for i := range arr {
		divisor := math.Pow10(length - 1 - i)
		arr[i] = int((num % int64(prev)) / int64(divisor))
		prev = divisor
	}

	return arr
}

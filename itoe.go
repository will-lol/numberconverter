package numberconverter

import (
	"math"
	"strings"

	"golang.org/x/exp/constraints"
)

func Itoe(num int64) string {
	// handle zero and negative cases
	negative := ""
	if num < 0 {
		negative = "negative "
		num = num * -1
	}
	if num == 0 {
		return "zero"
	}

	arrs := splitArr(toDigitArr(num))

	output := make([]string, len(arrs), len(arrs))

	for i, arr := range arrs {
		output[i] = fragmentToString(arr) + ItoeNumbers[len(arrs)-i-1]
	}

	return negative + strings.Join(output, " ")
}

func ItoeGeneric[T constraints.Signed](num T) string {
	return Itoe(int64(num))
}

func fragmentToString(arr []int) string {
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
		return "zero"
	}

	// hundreds
	if hundreds != 0 {
		out = append(out, ItoeUniques[hundreds]+" hundred")
	}

	// tens and ones
	tensAndOnes := make([]string, 0, 2)
	lessThanTwenty := 0
	if tens != 0 {
		if tens == 1 {
			lessThanTwenty = 10
		} else {
			tensAndOnes = append(tensAndOnes, ItoeTens[tens])
		}
	}
	lessThanTwenty += ones
	if lessThanTwenty != 0 {
		tensAndOnes = append(tensAndOnes, ItoeUniques[lessThanTwenty])
	}

	out = append(out, strings.Join(tensAndOnes, "-"))
	return strings.Join(out, " ")
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

func toDigitArr(num int64) []int {
	length := getDigitLength(num)
	arr := make([]int, length, length)

	prev := math.Inf(1)
	for i := range arr {
		divisor := math.Pow10(length - 1 - i)
		arr[i] = int((num % int64(prev)) / int64(divisor))
		prev = divisor
	}

	return arr
}

func getDigitLength(num int64) int {
	l := 1
	if num == 0 {
		return l
	}
	return l + int(math.Log10(float64(num)))
}

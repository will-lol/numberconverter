// Package numberconverter converts contemporary English language numbers ("negative one million three hundred thousand") into signed integers (-1_300_000), and back.
package numberconverter

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"strings"
	"unicode"
)

// Etoi (English to Integer) will convert an english language string into an int64. An input of "five" would return 5. This function may not error on some English syntax errors. It assumes correct English.
func Etoi(str string) (int64, error) {
	// handle empty string case
	if str == "" {
		return 0, errors.New("Received empty string")
	}

	words := strings.Fields(strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) {
			return 32
		}
		return r
	}, strings.ToLower(str)))

	var processed []string
	for _, word := range words {
		if word != "and" {
			if word == "a" {
				word = "one"
			}
			processed = append(processed, word)
		}
	}

	var negative int64 = 1
	if processed[0] == "negative" || processed[0] == "minus" {
		negative = -1
		processed = processed[1:]
	}

	nums, err := toNums(processed)
	if err != nil {
		return 0, err
	}

	out, err := recurse(imply(nums), false)
	if err != nil {
		return 0, err
	}

	return negative * out, nil
}

// EtoiGeneric also converts an english language string into a signed integer, but is generic. This function uses bare type coercion and may result in funky numbers being returned! If you want to guarantee a conversion, consider the non generic version.
func EtoiGeneric[T Integer](str string) (T, error) {
	i, err := Etoi(str)
	if err != nil {
		return 0, err
	}
	return T(i), nil
}

func recurse(arr []int64, isMultiplying bool) (int64, error) {
	if len(arr) == 0 {
		if isMultiplying {
			return 1, nil
		} else {
			return 0, nil
		}
	}
	if len(arr) == 1 {
		return arr[0], nil
	}
	i, err := findMaxIndex(arr)
	if err != nil {
		return 0, err
	}
	multiply := arr[0:i]
	add := arr[i+1:]

	multiplicand, err := recurse(multiply, true)
	if err != nil {
		return 0, err
	}
	addend, err := recurse(add, false)
	if err != nil {
		return 0, err
	}
	return arr[i]*multiplicand + addend, nil
}

func imply(arr []int64) []int64 {
	placeValues := make([]int, len(arr), len(arr))
	for i, val := range arr {
		placeValues[i] = getDigitLength[int](val)
	}

	m := slices.Max(placeValues)
	implyLocs := make([]int, 0, len(arr)/2)
	for i, val := range placeValues {
		if val == m {
			implyLocs = append(implyLocs, i)
		}
	}

	for i, val := range implyLocs {
		pow := int64(math.Pow10(m * (len(implyLocs) - 1 - i)))
		if pow != 1 {
			arr = insert(arr, val+i+1, pow)
		}
	}

	return arr
}

func findMaxIndex(arr []int64) (int, error) {
	maxIndex := 0
	flag := false
	for i, num := range arr {
		if num > arr[maxIndex] {
			maxIndex = i
			flag = false
		} else if num == arr[maxIndex] && i != maxIndex {
			flag = true
		}
	}
	if flag {
		return 0, errors.New("Nested duplicates unsupported i.e. 'one hundred two hundred thousand'")
	}
	return maxIndex, nil
}

func toNums(strs []string) ([]int64, error) {
	ints := make([]int64, len(strs), len(strs))

	for i, word := range strs {
		n, err := convertNum(word)
		if err != nil {
			return nil, err
		}
		ints[i] = n
	}

	return ints, nil
}

func convertNum(str string) (int64, error) {
	i, found := etoiTokens[str]
	if !found {
		return 0, errors.New(fmt.Sprintf("%s incorrect or unsupported", str))
	}
	return i, nil
}

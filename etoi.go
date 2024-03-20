// Package numberconverter converts contemporary English language numbers ("negative one million three hundred thousand") into signed integers (-1_300_000), and back.
package numberconverter

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var word = regexp.MustCompile(`\w+`)

// Etoi (English to Integer) will convert an english language string into an int64. An input of "five" would return 5. This function may not error on some English syntax errors. It assumes correct English.
func Etoi(str string) (int64, error) {
	// handle empty string case
	if str == "" {
		return 0, errors.New("Received empty string")
	}

	str = strings.ToLower(str)
	words := word.FindAllString(str, -1)

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

	return negative * recurse(nums, false), nil
}

// EtoiGeneric also converts an english language string into a signed integer, but is generic. This function uses bare type coercion and may result in funky numbers being returned! If you want to guarantee a conversion, consider the non generic version.
func EtoiGeneric[T Integer](str string) (T, error) {
	i, err := Etoi(str)
	if err != nil {
		return 0, err
	}
	return T(i), nil
}

func recurse(arr []int64, isMultiplying bool) int64 {
	if len(arr) == 0 {
		if isMultiplying {
			return 1
		} else {
			return 0
		}
	}
	if len(arr) == 1 {
		return arr[0]
	}
	i := findMaxIndex(arr)
	multiply := arr[0:i]
	add := arr[i+1:]
	return arr[i]*recurse(multiply, true) + recurse(add, false)
}

func imply(arr []int64) []int64 {
	return []int64{}
}

func findMaxIndex(arr []int64) int {
	maxIndex := 0
	for i, num := range arr {
		if num > arr[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
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

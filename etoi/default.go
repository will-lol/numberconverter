package etoi

import (
	"errors"
	"fmt"
	"math"
	"slices"

	"github.com/will-lol/numberconverter/util"
)

func TokensToInt(arr []int64) (int64, error) {
	if arr[0] == -1 {
		val, err := recurse(imply(arr[1:]), false)
		return -1*val, err
	}
	return recurse(imply(arr), false)
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
		placeValues[i] = util.GetDigitLength[int](val)
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
			arr = util.Insert(arr, val+i+1, pow)
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

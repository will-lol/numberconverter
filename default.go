// Package numberconverter converts contemporary English language numbers ("negative one million three hundred thousand") into signed integers (-1_300_000), and back.
package numberconverter

import (
	"github.com/will-lol/numberconverter/etoi"
	"github.com/will-lol/numberconverter/itoe"
	"github.com/will-lol/numberconverter/util"
)

// Etoi (English to Integer) will convert an english language string into an int64. An input of "five" would return 5. This function may not error on some English syntax errors. It assumes correct English.
func Etoi(str string) (int64, error) {
	tokens, err := etoi.Tokenize(str)	
	if err != nil {
		return 0, err
	}
	i, err := etoi.TokensToInt(tokens)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// EtoiGeneric also converts an english language string into a signed integer, but is generic. This function uses bare type coercion and may result in funky numbers being returned! If you want to guarantee a conversion, consider the non generic version.
func EtoiGeneric[T util.Integer](str string) (T, error) {
	i, err := Etoi(str)
	if err != nil {
		return 0, err
	}
	return T(i), nil
}

// Itoe (Integer to English) will convert an int64 into an English language string. For example, an input of 5 would produce "five". The style of the output is always consistent—lower case, no 'and', and hyphenation of numbers 21 to 99.
func Itoe(num int64) string {
	// handle zero and negative cases
	negative := false
	if num < 0 {
		negative = true
		num = num * -1
	}
	if num == 0 {
		return "zero"
	}

	return itoe.DigitArrToString(itoe.ToDigitArr(num), negative)
}

// Function ItoeGeneric performs the same function as Itoe but is generic.
func ItoeGeneric[T util.Integer](num T) string {
	return Itoe(int64(num))
}

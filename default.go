// Package numberconverter converts contemporary English language numbers ("negative one million three hundred thousand") into signed integers (-1_300_000), and back. numtowords and wordstonum.
package numberconverter

import (
	"bytes"
	"io"

	"github.com/will-lol/numberconverter/etoi"
	"github.com/will-lol/numberconverter/itoe"
	"github.com/will-lol/numberconverter/tokenizer"
	"github.com/will-lol/numberconverter/util"
)

// Etoi (English to Integer) will convert the first instance of an english language number string into an int64. An input of "five" would return 5. This function may not error on some English syntax errors. It assumes correct English.
func Etoi(r io.Reader) (int64, error) {
	tokenize := tokenizer.NewTokenizer(r)

	tokens := make([]int64, 0, 15)
	// read leading delims
	for {
		token, err := tokenize.Next()
		if err != nil {
			return 0, err
		}
		if token != tokenizer.DelimToken && token != tokenizer.AndToken {
			tokens = append(tokens, int64(token))
			break
		}
	}

	// read nums
	for {
		token, err := tokenize.Next()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return 0, err
		}
		if token != tokenizer.DelimToken {
			if token != tokenizer.AndToken {
				tokens = append(tokens, int64(token))
			}
		} else {
			break	
		}
	}

	i, err := etoi.TokensToInt(tokens)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// EtoiGeneric also converts an english language string into a signed integer, but is generic. This function uses bare type coercion and may result in funky numbers being returned! If you want to guarantee a conversion, consider the non generic version.
func EtoiGeneric[T util.Integer](r io.Reader) (T, error) {
	i, err := Etoi(r)
	if err != nil {
		return 0, err
	}
	return T(i), nil
}

func EtoiString(str string) (int64, error) {
	return Etoi(bytes.NewBufferString(str))
}

func EtoiStringGeneric[T util.Integer](str string) (T, error) {
	i, err := EtoiString(str)
	if err != nil {
		return 0, err
	}
	return T(i), nil
}

func EtoiBytes(buf []byte) (int64, error) {
	return Etoi(bytes.NewBuffer(buf))
}

func EtoiBytesGeneric[T util.Integer](buf []byte) (T, error) {
	i, err := EtoiBytes(buf)
	if err != nil {
		return 0, err
	}
	return T(i), nil
}

func EtoiReplace(str string) string {
	

}

func EtoiReplaceAll(str string) string {

}

func FindEnglishNumberIndex(str string) []int {
	tokenize := tokenizer.NewTokenizer(bytes.NewBufferString(str))

	tokens := make([]int64, 0, 15)
	// read leading delims
	for {
		token, err := tokenize.Next()
		if err != nil {
			return nil
		}
		if token != tokenizer.DelimToken && token != tokenizer.AndToken {
			tokens = append(tokens, int64(token))
			break
		}
	}

	// read nums
	for {
		token, err := tokenize.Next()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil
		}
		if token != tokenizer.DelimToken {
			if token != tokenizer.AndToken {
				tokens = append(tokens, int64(token))
			}
		} else {
			break	
		}
	}

}

func FindAllEnglishNumbers(str string) [][]int {

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

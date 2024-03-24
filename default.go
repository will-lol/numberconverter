// Package numberconverter converts contemporary English language numbers ("negative one million three hundred thousand") into signed integers (-1_300_000), and back. numtowords and wordstonum.
package numberconverter

import (
	"container/ring"
	"io"
	"strconv"

	"github.com/will-lol/numberconverter/etoi"
	"github.com/will-lol/numberconverter/itoe"
	"github.com/will-lol/numberconverter/tokenizer"
	"github.com/will-lol/numberconverter/util"
)

// Etoi (English to Integer) will convert the first instance of an english language number string into an int64. An input of "five" would return 5. This function may not error on some English syntax errors. It assumes correct English.
func Etoi(buf []rune) (int64, error) {
	tokenize := tokenizer.NewTokenizer(tokenizer.NewInputRunes(buf))
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
			if err == io.EOF {
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
func EtoiGeneric[T util.Integer](buf []rune) (T, error) {
	i, err := Etoi(buf)
	if err != nil {
		return 0, err
	}
	return T(i), nil
}

// EtoiString wraps Etoi, converting the given string to a rune array
func EtoiString(str string) (int64, error) {
	return Etoi([]rune(str))
}

// EtoiStringGeneric allows generic usage of EtoiString. Like EtoiGeneric, it also uses bare type coercion: use types with a lower maximum than int64 at your own risk!
func EtoiStringGeneric[T util.Integer](str string) (T, error) {
	i, err := EtoiString(str)
	if err != nil {
		return 0, err
	}
	return T(i), nil
}

// EtoiReplaceAll wraps EtoiReplaceAllFunc by always replacing matches.
func EtoiReplaceAll(str string) string {
	return EtoiReplaceAllFunc(str, func(_ int64) bool { return true })
}

// EtoiReplaceAllFunc allows a user replace all the instances of an English number in a given string with its numerical representation if the given function returns true. The given function is passed the integer representation of the matched English number.
func EtoiReplaceAllFunc(str string, f func(i int64) bool) string {
	tokenize := tokenizer.NewTokenizer(tokenizer.NewInputString(str))

	diff := 0
all:
	for {
		tokens := make([]int64, 0, 15)
		cur := []int{0, 0}
		// read leading delims
		for {
			token, err := tokenize.Next()
			if err != nil {
				if err == io.EOF {
					break all
				}
				return str
			}
			if token != tokenizer.DelimToken && token != tokenizer.AndToken {
				tokens = append(tokens, int64(token))
				cur[0] = tokenize.Position()[0]
				break
			}
		}

		// read nums
		andCheck := false
		for {
			cur[1] = tokenize.Position()[1]
			token, err := tokenize.Next()
			if err != nil {
				if err == io.EOF {
					break
				}
				return str
			}
			if token != tokenizer.DelimToken {
				if token != tokenizer.AndToken {
					if andCheck {
						prev := tokens[len(tokens)-1]
						if util.GetDigitLength[int64](prev) == util.GetDigitLength[int64](int64(token)) {
							tokenize.Set(cur[1] + 1) // rewind
							cur[1] = cur[1] - 4      // remove trailing ' and'
							break
						}
						andCheck = false
					}

					tokens = append(tokens, int64(token))
				} else {
					andCheck = true
				}
			} else {
				break
			}
		}

		i, err := etoi.TokensToInt(tokens)
		if err != nil {
			return str
		}

		if f(i) {
			strLen := len(str)
			str = str[:(cur[0]-diff)] + strconv.FormatInt(i, 10) + str[(cur[1]-diff):]
			diff += strLen - len(str)
		}
	}

	return str
}

// FindEnglishNumber wraps FindAllEnglishNumber with n=1
func FindEnglishNumber(str string) string {
	return FindAllEnglishNumber(str, 1)[0]
}

// FindAllEnglishNumber wraps FindAllEnglishNumberIndex but returns a list of strings rather than indexes.
func FindAllEnglishNumber(str string, n int) []string {
	nums := make([]string, 0)
	indices := FindAllEnglishNumberIndex(str, n)

	for _, val := range indices {
		nums = append(nums, str[val[0]:val[1]])
	}

	return nums
}

// FindEnglishNumberIndex wraps FindAllEnglishNumberIndex with n=1
func FindEnglishNumberIndex(str string) []int {
	return FindAllEnglishNumberIndex(str, 1)[0]
}

// FindAllEnglishNumberIndex searches a given string for English numbers, returning their indexes in an array. The parameter 'n' specifies the maximum number of numbers returned. If n=-1, all numbers are returned.
func FindAllEnglishNumberIndex(str string, n int) [][]int {
	tokenize := tokenizer.NewTokenizer(tokenizer.NewInputString(str))
	out := make([][]int, 0)

	if n < 0 {
		n = len(str) + 1
	}

all:
	for i := 0; i < n; i++ {
		cur := []int{0, 0}
		tokens := ring.New(2)

		// read leading delims
		for {
			token, err := tokenize.Next()
			if err != nil {
				if err == io.EOF {
					break all
				}
				return out
			}
			if token != tokenizer.DelimToken && token != tokenizer.AndToken {
				cur[0] = tokenize.Position()[0]
				tokens.Value = token
				break
			}
		}

		// read nums
		andCheck := false
		for {
			cur[1] = tokenize.Position()[1]
			token, err := tokenize.Next()
			if err != nil {
				if err == io.EOF {
					break
				}
				return out
			}
			if token != tokenizer.DelimToken {
				if token != tokenizer.AndToken {
					tokens = tokens.Next()

					if andCheck {
						prev := tokens.Prev().Value
						if util.GetDigitLength[int64](int64(prev.(tokenizer.TokenType))) == util.GetDigitLength[int64](int64(token)) {
							tokenize.Set(cur[1] + 1) // rewind
							cur[1] = cur[1] - 4      // remove trailing ' and'
							break
						}
						andCheck = false
					}

					tokens.Value = token
				} else {
					andCheck = true
				}
			} else {
				break
			}
		}

		out = append(out, cur)
	}

	return out
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

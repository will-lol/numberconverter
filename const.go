package numberconverter

var itoeNumbers = []string{"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion"}

var itoeUniques = map[int]string{
	0:  "zero",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}

var itoeTens = map[int]string{
	1: "ten",
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

var etoiTokens = map[string]int64{
	"oh":          0,
	"zero":        0,
	"one":         1,
	"two":         2,
	"three":       3,
	"four":        4,
	"five":        5,
	"six":         6,
	"seven":       7,
	"eight":       8,
	"nine":        9,
	"ten":         10,
	"eleven":      11,
	"twelve":      12,
	"thirteen":    13,
	"fourteen":    14,
	"fifteen":     15,
	"sixteen":     16,
	"seventeen":   17,
	"eighteen":    18,
	"nineteen":    19,
	"twenty":      20,
	"thirty":      30,
	"forty":       40,
	"fourty":      40,
	"fifty":       50,
	"sixty":       60,
	"seventy":     70,
	"eighty":      80,
	"ninety":      90,
	"hundred":     100,
	"thousand":    1_000,
	"million":     1_000_000,
	"billion":     1_000_000_000,
	"trillion":    1_000_000_000_000,
	"quadrillion": 1_000_000_000_000_000,
	"quintillion": 1_000_000_000_000_000_000,
}

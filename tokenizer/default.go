package tokenizer

import (
	"math"
	"unicode"
)

type TokenType int64

const (
	DelimToken  TokenType = math.MaxInt64
	AndToken    TokenType = math.MaxInt64 - 1
	Negative    TokenType = -1
	Zero        TokenType = 0
	One         TokenType = 1
	Two         TokenType = 2
	Three       TokenType = 3
	Four        TokenType = 4
	Five        TokenType = 5
	Six         TokenType = 6
	Seven       TokenType = 7
	Eight       TokenType = 8
	Nine        TokenType = 9
	Ten         TokenType = 10
	Eleven      TokenType = 11
	Twelve      TokenType = 12
	Thirteen    TokenType = 13
	Fourteen    TokenType = 14
	Fifteen     TokenType = 15
	Sixteen     TokenType = 16
	Seventeen   TokenType = 17
	Eighteen    TokenType = 18
	Nineteen    TokenType = 19
	Twenty      TokenType = 20
	Thirty      TokenType = 30
	Forty       TokenType = 40
	Fifty       TokenType = 50
	Sixty       TokenType = 60
	Seventy     TokenType = 70
	Eighty      TokenType = 80
	Ninety      TokenType = 90
	Hundred     TokenType = 100
	Thousand    TokenType = 1_000
	Million     TokenType = 1_000_000
	Billion     TokenType = 1_000_000_000
	Trillion    TokenType = 1_000_000_000_000
	Quadrillion TokenType = 1_000_000_000_000_000
	Quintillion TokenType = 1_000_000_000_000_000_000
)

const (
	AndHash         = 136963
	BillionHash     = 258930331441
	EightHash       = 194753797
	EighteenHash    = 9864864221557
	EightyHash      = 7205890610
	ElevenHash      = 7211430387
	FifteenHash     = 269182391489
	FiftyHash       = 196627038
	FiveHash        = 5314818
	FortyHash       = 196947384
	FourHash        = 5323008
	FourteenHash    = 9976180014152
	FourtyHash      = 7287202365
	HundredHash     = 275160172418
	MillionHash     = 287153321940
	MinusHash       = 209757148
	NegativeHash    = 10708975097087
	NineHash        = 5719746
	NineteenHash    = 10719730900970
	NinetyHash      = 7830336687
	OhHash          = 4211
	OneHash         = 156130
	QuadrillionHash = 558926061315938568
	QuintillionHash = 558955115601794084
	SevenHash       = 220809857
	SeventeenHash   = 413833228422841
	SeventyHash     = 302288698646
	SixHash         = 161440
	SixteenHash     = 302570569704
	SixtyHash       = 221015773
	TenHash         = 162651
	ThirteenHash    = 11286434081667
	ThirtyHash      = 8244290800
	ThousandHash    = 11286855712086
	ThreeHash       = 222830492
	TrillionHash    = 11312079701413
	TwelveHash      = 8272192443
	TwentyHash      = 8272195127
	TwoHash         = 163318
	ZeroHash        = 6322264
)

var Tokens = map[int]TokenType{
	AndHash:         AndToken,
	MinusHash:       Negative,
	NegativeHash:    Negative,
	ZeroHash:        Zero,
	OneHash:         One,
	TwoHash:         Two,
	ThreeHash:       Three,
	FourHash:        Four,
	FiveHash:        Five,
	SixHash:         Six,
	SevenHash:       Seven,
	EightHash:       Eight,
	NineHash:        Nine,
	TenHash:         Ten,
	ElevenHash:      Eleven,
	TwelveHash:      Twelve,
	ThirteenHash:    Thirteen,
	FourteenHash:    Fourteen,
	FifteenHash:     Fifteen,
	SixteenHash:     Sixteen,
	SeventeenHash:   Seventeen,
	EighteenHash:    Eighteen,
	NineteenHash:    Nineteen,
	TwentyHash:      Twenty,
	ThirtyHash:      Thirty,
	FourtyHash:      Forty,
	FortyHash:       Forty,
	FiftyHash:       Fifty,
	SixtyHash:       Sixty,
	SeventyHash:     Seventy,
	EightyHash:      Eighty,
	NinetyHash:      Ninety,
	OhHash:          Hundred,
	HundredHash:     Hundred,
	ThousandHash:    Thousand,
	MillionHash:     Million,
	BillionHash:     Billion,
	TrillionHash:    Trillion,
	QuadrillionHash: Quadrillion,
	QuintillionHash: Quintillion,
}

func hash(buf []rune) int {
	h := 0
	for _, byte := range buf {
		if byte != 0 {
			h = 37*h + int(unicode.ToLower(byte))
		}
	}
	return h
}

type Tokenizer interface {
	Next() (TokenType, error)
	Position() []int
	Set(pos int)
}

type tokenizer struct {
	Input Input
	Pos []int
}

func NewTokenizer(input Input) Tokenizer {
	return &tokenizer{
		Input: input,
		Pos: []int{0, 0},
	}
}

// Returns the position of the last token
func (t *tokenizer) Position() ([]int) {
	return t.Pos
}

var significant = []*unicode.RangeTable{unicode.Letter, unicode.Sentence_Terminal}

func (t *tokenizer) Set(pos int) {
	t.Input.Set(pos)
}

func (t *tokenizer) Next() (TokenType, error) {
	val := t.Input.Peek(0)
	if err := t.Input.Error(); err != nil {
		return 0, err
	}
	if !unicode.IsOneOf(significant, val) {
		t.consumeInsignificant()
	}

	val = t.Input.Peek(0)
	if unicode.Is(unicode.Sentence_Terminal, val) {
		t.Input.Advance(1)
		return DelimToken, nil
	}

	t.Pos[0] = t.Input.Position()
	word, err := t.consumeWord()
	t.Pos[1] = t.Input.Position()

	if err != nil {
		return 0, err
	}

	if word[len(word)-1] == 's' {
		word = word[:len(word)-1]
	}
	token, ok := Tokens[hash(word)]
	if !ok {
		return DelimToken, nil
	}
	return token, nil
}

func (t *tokenizer) consumeInsignificant() error {
	for {
		r := t.Input.Peek(0)
		if unicode.IsOneOf(significant, r) || t.Input.Error() != nil {
			break
		}
		t.Input.Advance(1)
	}
	return t.Input.Error()
}

func (t *tokenizer) consumeWord() ([]rune, error) {
	s := t.Input.ReadSliceFunc(func(r rune) bool {
		return unicode.IsLetter(r)
	})
	if err := t.Input.Error(); err != nil {
		return nil, err
	}
	for i := range s {
		if unicode.IsOneOf([]*unicode.RangeTable{unicode.Space, unicode.Dash}, s[i]) {
			s[i] = 0
		}
	}
	return s, nil
}

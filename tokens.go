package numberconverter

type Token interface {
	Order() int
	Place() int
	Multiplier() int
	SetMultiplier() int
}

type token struct {
}

package numberconverter_test

import (
	"testing"

	"github.com/will-lol/numberconverter"
)

func BenchmarkItoe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numberconverter.Itoe(21474836473)
	}
}

func TestItoeSimple(t *testing.T) {
	var in int64 = 21474836473
	var want string = "twenty-one billion four hundred seventy-four million eight hundred thirty-six thousand four hundred seventy-three"
	val := numberconverter.Itoe(in)
	if want != val {
		t.Fatalf("Expected \"%s\" but got \"%s\"", want, val)
	}
}

func TestItoeNum(t *testing.T) {
	var in int64 = 5
	var want string = "five"
	val := numberconverter.Itoe(in)
	if want != val {
		t.Fatalf("Expected \"%s\" but got \"%s\"", want, val)
	}
}

func FuzzItoe(f *testing.F) {
	f.Fuzz(func(t *testing.T, i int) {
		out := numberconverter.ItoeGeneric(i)
		res, err := numberconverter.EtoiGeneric[int](out)
		if err != nil {
			t.Error(err.Error(), "input", i, "got", out)
		}
		if res != i {
			t.Error("input", i, "went to", out, "expected", i, "but got", res)
		}
	})
}

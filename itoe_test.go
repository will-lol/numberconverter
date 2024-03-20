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

func TestItoe(t *testing.T) {
	cases := map[int64]string{
		21474836473: "twenty-one billion four hundred seventy-four million eight hundred thirty-six thousand four hundred seventy-three",
		5: "five",
	}

	for in, out := range cases {
		val := numberconverter.Itoe(in)
		if out != val {
			t.Fatalf("Expected %s but got %s", out, val)
		}
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

package numberconverter_test

import (
	"github.com/will-lol/numberconverter"
	"testing"
)

func BenchmarkEtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numberconverter.Etoi("one hundred and fourty five million two hundred thousand two hundred and fourty five")
	}
}

func FuzzEtoi(f *testing.F) {
	f.Add(int64(9_223_372_036_554_775_806))
	f.Add(int64(4))
	f.Add(int64(2312312))
	f.Fuzz(func(t *testing.T, i int64) {
		str := numberconverter.Itoe(i)
		val, err := numberconverter.Etoi(str)
		if err != nil {
			t.Error(err)
		}
		if val != i {
			t.Errorf("Expected %d from converted %s but got %d", i, str, val)
		}
	})
}

func TestEtoiGeneric(t *testing.T) {
	in := "fifty five"
	var want int8 = 55
	val, err := numberconverter.EtoiGeneric[int8](in)
	if err != nil {
		t.Fatal(err)
	}
	if want != val {
		t.Fatalf("Expected %d but got %d", want, val)
	}
}

func TestEtoiError(t *testing.T) {
	cases := []string{
		"one hundred two hundred thousand",
	}
	for _, val := range cases {
		out, err := numberconverter.Etoi(val)
		if err == nil {
			t.Fatalf("Expected err but got %d", out)
		}
	}
}

func TestEtoi(t *testing.T) {
	cases := map[string]int64{
		"zero":                                0,
		"one hundred and twenty three":        123,
		"Two-million, four hundred, and five": 2_000_405,
		"one hundred and fourty five million two hundred thousand two hundred and fourty five":          145_200_245,
		"negative one hundred and fourty five million two hundred thousand two hundred and fourty five": -145_200_245,
		"hundred million":             100_000_000,
		"three ten":                   30,
		"nineteen oh five":            1905,
		"one two five":                125,
		"six five four three two one": 654321,
		"twenty twenty four":          2024,
		"twenty twenty":               2020,
		"one hundred two hundred":     100200,
		"twenty twenty twenty":        202020,
		"one thousand million":        1_000_000_000,
		"nineteen thirty five":        1935,
		"negative twenty twenty":      -2020,
	}

	for in, out := range cases {
		val, err := numberconverter.Etoi(in)
		if err != nil {
			t.Fatal(err)
		}
		if out != val {
			t.Fatalf("Expected %d but got %d", out, val)
		}
	}
}

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

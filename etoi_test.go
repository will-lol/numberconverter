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

func TestEtoiInvalid(t *testing.T) {
	in := "five five"
	val, err := numberconverter.Etoi(in)
	if err == nil {
		t.Fatalf("Should have errored, instead got %d", val)
	}
}

func TestEtoi(t *testing.T) {
	cases := map[string]int64{
		"zero":                                0,
		"one hundred and twenty three":        123,
		"Two-million, four hundred, and five": 2_000_405,
		"one hundred and fourty five million two hundred thousand two hundred and fourty five":          145_200_245,
		"negative one hundred and fourty five million two hundred thousand two hundred and fourty five": -145_200_245,
		"hundred million":      100_000_000,
		"three ten": 30,
		"oh fourty five": 45,
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

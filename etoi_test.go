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
	f.Fuzz(func (t *testing.T, i int64) {
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

func TestEtoiZero(t *testing.T) {
	in := "zero"
	var want int64 = 0
	val, err := numberconverter.Etoi(in)
	if err != nil {
		t.Fatal(err)
	}
	if want != val {
		t.Fatalf("Expected %d but got %d", want, val)
	}
}

func TestEtoiSimple(t *testing.T) {
	in := "one hundred and twenty three"
	var want int64 = 123
	val, err := numberconverter.Etoi(in)
	if err != nil {
		t.Fatal(err)
	}
	if want != val {
		t.Fatalf("Expected %d but got %d", want, val)
	}
}

func TestEtoiGaps(t *testing.T) {
	in := "Two-million, four hundred, and five"
	var want int64 = 2_000_405
	val, err := numberconverter.Etoi(in)
	if err != nil {
		t.Fatal(err)
	}
	if want != val {
		t.Fatalf("Expected %d but got %d", want, val)
	}
}

func TestEtoiHarder(t *testing.T) {
	in := "one hundred and fourty five million two hundred thousand two hundred and fourty five"
	var want int64 = 145_200_245
	val, err := numberconverter.Etoi(in)
	if err != nil {
		t.Fatal(err)
	}
	if want != val {
		t.Fatalf("Expected %d but got %d", want, val)
	}
}

func TestEtoiNegative(t *testing.T) {
	in := "negative one hundred and fourty five million two hundred thousand two hundred and fourty five"
	var want int64 = -145_200_245
	val, err := numberconverter.Etoi(in)
	if err != nil {
		t.Fatal(err)
	}
	if want != val {
		t.Fatalf("Expected %d but got %d", want, val)
	}
}

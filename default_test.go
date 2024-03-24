package numberconverter_test

import (
	"slices"
	"testing"

	"github.com/will-lol/numberconverter"
)

func BenchmarkEtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numberconverter.EtoiString("one hundred and fourty five million two hundred thousand two hundred and fourty five")
	}
}

func TestFindEnglishNumber(t *testing.T) {
	cases := map[string]string {
		"I have three dogs and five cats": "three",
		"If we talk about dogs, I have three. Two of them live in a box": "three",
		"Between three and five dogs live in my house": "three",
	}
	for in, out := range cases {
		val := numberconverter.FindEnglishNumber(in)
		if val != out {
			t.Fatalf("Expected string %q but got %q from %q", out, val, in)
		}
	}
}

func TestFindAllEnglishNumber(t *testing.T) {
	cases := map[string][]string{
		"There are fifty five dogs": {"fifty five"},
		"five": {"five"},
		"Fifty five dogs. Three hundred and twenty three geese. Next level bears.": {"Fifty five", "Three hundred and twenty three"},
		"If we talk about dogs, I have three. Two of them": {"three", "Two"},
	}
	for in, out := range cases {
		val := numberconverter.FindAllEnglishNumber(in, -1)
		for _, str := range out {
			if !slices.Contains(val, str) {
				t.Fatalf("Expected slice %#v to contain %q from %q", val, str, in)
			}
			
		}
	}
}

func TestEtoiReplaceAll(t *testing.T) {
	cases := map[string]string{
		"If we talk about dogs, I have three. Two of them": "If we talk about dogs, I have 3. 2 of them",
		"Between three and five dogs live in my house": "Between 3 and 5 dogs live in my house",
		"I have fifty five dogs in my house":                                                "I have 55 dogs in my house",
		"I have number fifty five-for each number it is a dog.":                             "I have number 55-for each number it is a dog.",
		"There are Three hundred and seventy dogs in my house, and eight of them are dead!": "There are 370 dogs in my house, and 8 of them are dead!",
		"Three hundred and fifty three dogs are EATING my one pizza and I just want them to stop it! Two dogs ate me as well :(": "353 dogs are EATING my 1 pizza and I just want them to stop it! 2 dogs ate me as well :(",
		"five": "5",
	}

	for in, out := range cases {
		val := numberconverter.EtoiReplaceAll(in)
		if val != out {
			t.Fatalf("Expected %q from %q but got %q", out, in, val)
		}
	}
}

func FuzzEtoi(f *testing.F) {
	f.Add(int64(9_223_372_036_554_775_806))
	f.Add(int64(4))
	f.Add(int64(2312312))
	f.Fuzz(func(t *testing.T, i int64) {
		str := numberconverter.Itoe(i)
		val, err := numberconverter.EtoiString(str)
		if err != nil {
			t.Error(err)
		}
		if val != i {
			t.Errorf("Expected %d from converted %q but got %d", i, str, val)
		}
	})
}

func TestEtoiGeneric(t *testing.T) {
	in := "fifty five"
	var want int8 = 55
	val, err := numberconverter.EtoiStringGeneric[int8](in)
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
		out, err := numberconverter.EtoiString(val)
		if err == nil {
			t.Fatalf("Expected err but got %d", out)
		}
	}
}

func TestEtoi(t *testing.T) {
	cases := map[string]int64{
		"zero": 0,
		"my dog has one hundred and twenty three bones":                                                 123,
		"Two-million, four hundred, and five":                                                           2_000_405,
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
		val, err := numberconverter.EtoiString(in)
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
		5:           "five",
	}

	for in, out := range cases {
		val := numberconverter.Itoe(in)
		if out != val {
			t.Fatalf("Expected %q but got %q", out, val)
		}
	}
}

func FuzzItoe(f *testing.F) {
	f.Fuzz(func(t *testing.T, i int) {
		out := numberconverter.ItoeGeneric(i)
		res, err := numberconverter.EtoiStringGeneric[int](out)
		if err != nil {
			t.Error(err.Error(), "input", i, "got", out)
		}
		if res != i {
			t.Error("input", i, "went to", out, "expected", i, "but got", res)
		}
	})
}

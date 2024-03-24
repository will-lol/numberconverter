package tokenizer_test

import (
	"fmt"
	"testing"

	"github.com/will-lol/numberconverter/tokenizer"
)

func TestMain(t *testing.T) {
	tokenize := tokenizer.NewTokenizer(tokenizer.NewInputString("Between three and five dogs live in my house"))
	for {
		token, err := tokenize.Next()
		if err != nil {
			break
		}
		fmt.Println(token)
	}
}

package tokenizer

import (
	"io"
)

type Input interface {
	Peek(pos int) rune
	Advance(n int)
	Retreat(n int)
	Error() error
	Position() int
	Buffer() *[]rune
	ReadSliceFunc(f func(r rune) bool) []rune
	Set(pos int)
}

type input struct {
	Buf []rune
	Pos int	
	Err error
}

func NewInputString(s string) Input {
	return NewInputRunes([]rune(s))
}

func NewInputRunes(r []rune) Input {
	return &input{
		Buf: r,
		Pos: 0,
		Err: nil,
	}
}

func (in *input) Peek(pos int) rune {
	pos += in.Pos
	if pos >= len(in.Buf) - 1 || pos < 0 {
		in.Err = io.EOF
		return 0
	}
	return in.Buf[pos]
}

func (in *input) Advance(n int) {
	in.Pos += n
}

func (in *input) Retreat(n int) {
	in.Pos -= n
}

func (in *input) Position() int {
	return in.Pos
}

func (in *input) Buffer() *[]rune {
	return &in.Buf
}

func (in *input) ReadSliceFunc(f func(r rune) bool) []rune {
	s := 0
	if in.Pos > len(in.Buf) || in.Pos < 0 {
		in.Err = io.EOF
		return nil
	}
	for _, val := range in.Buf[in.Pos:] {
		if f(val) {
			s++
		} else {
			break
		}
	}
	res := in.Buf[in.Pos:in.Pos + s]
	in.Pos += s
	return res
}

func (in *input) Error() error {
	return in.Err
}

func (in *input) Set(pos int) {
	in.Pos = pos
}

package vq

import (
	"fmt"

	"github.com/leorolland/vq/parser"
)

func isCloseBracket(r rune) bool {
	return r == ']'
}

func isNotCloseBrackets(r rune) bool {
	return !isCloseBracket(r)
}

func Brackets() parser.Parser[string] {
	s := parser.StartSkipping(parser.Exactly("["))
	s1 := parser.AppendKeeping(s, parser.GetString(parser.ConsumeWhile(isNotCloseBrackets)))
	s2 := parser.AppendSkipping(s1, parser.Exactly("]"))
	return parser.Apply(s2, func(inside string) string {
		return fmt.Sprintf("[%s]", inside)
	})
}

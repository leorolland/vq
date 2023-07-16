package parsers

import (
	"fmt"
	"strings"

	"github.com/leorolland/vq/parser"
)

func isCloseBracket(r rune) bool {
	return r == ']'
}

func isNotCloseBrackets(r rune) bool {
	return !isCloseBracket(r)
}

func Brackets(recursion int) parser.Parser[string] {
	s := parser.StartSkipping(parser.Exactly("["))

	var s1 parser.Parser[parser.Seq[parser.Empty, string]]
	if recursion == 0 {
		s1 = parser.AppendKeeping(s, Text())
	} else {
		s1 = parser.AppendKeeping(s, parser.Map(
			Anythings(recursion-1),
			func(things []string) string {
				return strings.Join(things, ", ")
			},
		))
	}

	s2 := parser.AppendSkipping(s1, parser.Exactly("]"))

	return parser.Apply(s2, func(inside string) string {
		return fmt.Sprintf("Brackets(%s)", inside)
	})
}

package parsers

import (
	"github.com/leorolland/vq"
	"github.com/leorolland/vq/parser"
)

func isCloseBracket(r rune) bool {
	return r == ']'
}

func isNotCloseBrackets(r rune) bool {
	return !isCloseBracket(r)
}

func Brackets(recursion int) parser.Parser[vq.Node] {
	s := parser.StartSkipping(parser.Exactly("["))

	var s1 parser.Parser[parser.Seq[parser.Empty, vq.Node]]
	if recursion == 0 {
		s1 = parser.AppendKeeping(s, Text())
	} else {
		s1 = parser.AppendKeeping(s, parser.Map(
			Anythings(recursion-1),
			func(things []vq.Node) vq.Node {
				return vq.Node{
					Kind:     vq.Brackets,
					Children: things,
				}
			},
		))
	}

	s2 := parser.AppendSkipping(s1, parser.Exactly("]"))

	return parser.Apply(s2, func(inside vq.Node) vq.Node {
		return vq.Node{
			Kind:     vq.Brackets,
			Children: inside.Children,
		}
	})
}

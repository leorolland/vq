package parsers

import (
	"github.com/leorolland/vq"
	"github.com/leorolland/vq/parser"
)

func isNotDelimiter(r rune) bool {
	return r != '[' && r != ']'
}

func Text() parser.Parser[vq.Node] {
	return parser.Map(
		parser.GetString(parser.ConsumeSome(isNotDelimiter)),
		func(s string) vq.Node {
			return vq.Node{
				Kind:     vq.Text,
				Value:    s,
				Children: nil,
			}
		},
	)
}

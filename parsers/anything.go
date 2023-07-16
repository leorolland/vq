package parsers

import (
	"github.com/leorolland/vq"
	"github.com/leorolland/vq/parser"
)

func Anything(recursion int) parser.Parser[vq.Node] {
	return parser.OneOf(
		Brackets(recursion),
		Text(),
	)
}

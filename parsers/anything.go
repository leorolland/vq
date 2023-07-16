package parsers

import "github.com/leorolland/vq/parser"

func Anything(recursion int) parser.Parser[string] {
	return parser.OneOf(
		Brackets(recursion),
		Text(),
	)
}

package parsers

import "github.com/leorolland/vq/parser"

func Anything() parser.Parser[string] {
	return parser.OneOf(
		Brackets(),
		Text(),
	)
}

package parsers

import "github.com/leorolland/vq/parser"

func isNotDelimiter(r rune) bool {
	return r != '[' && r != ']'
}

func Text() parser.Parser[string] {
	return parser.GetString(parser.ConsumeSome(isNotDelimiter))
}

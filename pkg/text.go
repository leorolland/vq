package vq

import "github.com/leorolland/vq/parser"

func isAsciiLetter(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}

func Text() parser.Parser[string] {
	return parser.GetString(parser.ConsumeSome(isAsciiLetter))
}

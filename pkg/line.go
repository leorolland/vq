package vq

import "github.com/leorolland/vq/parser"

type Loglevel string

type LineParsers struct {
	// infoParser       parser.Parser[string]
	// warnParser       parser.Parser[string]
	// loglevelParser   parser.Parser[string]
	// whitespaceParser   parser.Parser[parser.Empty]
	word       parser.Parser[string]
	brackets   parser.Parser[string]
	LineParser parser.Parser[string]
}

// func isWhitespace(r rune) bool {
// 	return r == ' ' || r == '\n' || r == '\t'
// }]

func isAsciiLetter(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}

func isOpenBracket(r rune) bool {
	return r == '['
}

func isCloseBracket(r rune) bool {
	return r == ']'
}

func isNotCloseBrackets(r rune) bool {
	return r != ']'
}

func NewLineParser() LineParsers {
	var p LineParsers

	// p.infoParser = parser.Map(
	// 	parser.Exactly("INFO"),
	// 	func(parser.Empty) string {
	// 		return "info"
	// 	},
	// )
	// p.warnParser = parser.Map(
	// 	parser.Exactly("WARN"),
	// 	func(parser.Empty) string {
	// 		return "warn"
	// 	},
	// )

	// p.loglevelParser = parser.OneOf(
	// 	p.infoParser,
	// 	p.warnParser,
	// )

	// p.whitespaceParser = parser.ConsumeWhile(isWhitespace)

	{
		p.word = parser.GetString(parser.ConsumeWhile(isAsciiLetter))
	}

	{
		s := parser.StartSkipping(parser.Exactly("["))
		s1 := parser.AppendKeeping(s, parser.ConsumeWhile(isNotCloseBrackets))
		s2 := parser.AppendSkipping(s1, parser.Exactly("]"))
		p.brackets = parser.Apply(s2, func(x parser.Empty) string { return "foo" })
	}

	{
		p.LineParser = parser.OneOf(
			p.brackets,
			p.word,
		)
	}

	return p
}

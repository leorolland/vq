package parsers

import (
	"github.com/leorolland/vq/parser"
)

func Anythings() parser.Parser[[]string] {
	type anythingList struct {
		thing string
		next  *anythingList
	}
	return parser.Loop(nil,
		func(things *anythingList) parser.Parser[parser.Step[*anythingList, []string]] {
			extend := parser.Map(
				Anything(),
				func(thing string) parser.Step[*anythingList, []string] {
					return parser.Step[*anythingList, []string]{
						Accum: &anythingList{thing: thing, next: things},
						Done:  false,
					}
				},
			)

			var thingsSlice []string
			t := things
			for {
				if t == nil {
					break
				}
				thingsSlice = append(thingsSlice, t.thing)
				t = t.next
			}

			return parser.OneOf(
				extend,
				parser.Succeed(parser.Step[*anythingList, []string]{Value: thingsSlice, Done: true}),
			)
		},
	)
}

package parsers

import (
	"github.com/leorolland/vq"
	"github.com/leorolland/vq/parser"
)

func Anythings(recursion int) parser.Parser[[]vq.Node] {
	type anythingList struct {
		thing vq.Node
		next  *anythingList
	}
	return parser.Loop(nil,
		func(things *anythingList) parser.Parser[parser.Step[*anythingList, []vq.Node]] {
			extend := parser.Map(
				Anything(recursion),
				func(thing vq.Node) parser.Step[*anythingList, []vq.Node] {
					return parser.Step[*anythingList, []vq.Node]{
						Accum: &anythingList{thing: thing, next: things},
						Done:  false,
					}
				},
			)

			var thingsSlice []vq.Node
			t := things
			for {
				if t == nil {
					break
				}
				thingsSlice = append([]vq.Node{t.thing}, thingsSlice...)
				t = t.next
			}

			return parser.OneOf(
				extend,
				parser.Succeed(parser.Step[*anythingList, []vq.Node]{Value: thingsSlice, Done: true}),
			)
		},
	)
}

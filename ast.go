package vq

type Kind string

const (
	Text     Kind = "Text"
	Brackets      = "Brackets"
)

type Node struct {
	Kind     Kind
	Value    any
	Children []Node
}

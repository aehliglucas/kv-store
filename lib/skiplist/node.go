package skiplist

type SkiplistNode struct {
	links []*SkiplistNode
	levels int
	key string
	value string
}

func NewSkiplistNode(levels int, key string, value string) *SkiplistNode {
	return &SkiplistNode{
		levels: levels,
		key: key,
		value: value,
		links: make([]*SkiplistNode, levels),
	}
}
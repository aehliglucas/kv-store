package skiplist

import (
	"fmt"
)

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

func (node *SkiplistNode) Dump() {
	fmt.Printf("SkiplistNode[%d] [%s:%s] AT %p\n",
		node.levels, node.key, node.value, node)

	for i := node.levels - 1; i >= 0; i-- {
		fmt.Printf("Node %p pointing to %p on level %d\n",
			node, node.links[i], i)
	}
}
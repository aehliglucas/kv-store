package skiplist

import "errors"

type PredecessorResult struct {
	PredecessorNode *SkiplistNode
	Predecessors []*SkiplistNode
	Error error
}

var ErrKeyNotFound = errors.New("Key not found")

func (list *Skiplist) IdentifyPredecessorNodes(key string) PredecessorResult {
	predecessors := make([]*SkiplistNode, list.maxLevel)
	current := list.head
	
	for level := list.maxLevel - 1; level >= 0; level-- {
		// Move forward on current level as long as the next key is smaller than the provided one
		for current.links[level] != nil && current.links[level] != list.tail && current.links[level].key < key {
			current = current.links[level]
		}
		predecessors[level] = current
	}

	var predecessorNode *SkiplistNode
	if current != list.head {
		predecessorNode = current
	}

	next := current.links[0]
	err := ErrKeyNotFound
	if next != nil && next != list.tail && next.key == key {
		err = nil
	}

	return PredecessorResult{
		PredecessorNode: predecessorNode,
		Predecessors: predecessors,
		Error: err,
	}
}
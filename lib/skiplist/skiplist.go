package skiplist

import (
	"errors"
	"math/rand"
	"time"
)

type Skiplist struct {
	head *SkiplistNode
	tail *SkiplistNode
	maxLevel int
	currentMaxLevel int
	p float64
	rng *rand.Rand
}

type PredecessorResult struct {
	PredecessorNode *SkiplistNode
	Predecessors []*SkiplistNode
	Error error
}

var ErrKeyNotFound = errors.New("Key not found")

func NewSkipList(maxLevel int, p float64) *Skiplist {
	head := NewSkiplistNode(maxLevel, "HEAD_KEY", "")
	tail := NewSkiplistNode(maxLevel, "\xff", "")
	randSource := rand.NewSource(time.Now().UnixNano())
	 
	for i := range maxLevel {
		head.links[i] = tail
	}

	return &Skiplist{
		head: head,
		tail: tail,
		maxLevel: maxLevel,
		currentMaxLevel: 0,
		p: p,
		rng: rand.New(randSource),
	}
}

func (list *Skiplist) Search(key string) (string, error) {
	return "", nil
}

func (list *Skiplist) Insert(key string, value string) error {
	level := list.getRandomLevel()
	if level > list.currentMaxLevel {
		list.currentMaxLevel = level
	}
	return nil
}

func (list *Skiplist) Delete(key string) error {
	return nil
}

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

func (list *Skiplist) getRandomLevel() int {
	level := 1
	for list.rng.Float64() < list.p && level < list.maxLevel {
		level++
	}

	return level
}
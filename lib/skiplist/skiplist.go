package skiplist

import (
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

// TODO: Replace with "flush" later on
func (list *Skiplist) Delete(key string) error {
	return nil
}

func (list *Skiplist) getRandomLevel() int {
	level := 1
	for list.rng.Float64() < list.p && level < list.maxLevel {
		level++
	}

	return level
}
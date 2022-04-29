package kdb

import (
	"bytes"
	"github.com/dgraph-io/ristretto/z"
	"math"
	"sync/atomic"
)

const (
	defaultLevel   = 20
	heightIncrease = math.MaxUint32 / 3
)

type SimpleSkl struct {
	height int32
	// head is nil and its indexer
	head *SimpleSklNode
	storage
}

type Level struct {
	next *SimpleSklNode
}

func NewLevel(level int) []*Level {
	levels := make([]*Level, level)
	for i := 0; i < level; i++ {
		levels[i] = new(Level)
	}
	return levels
}

type SimpleSklNode struct {
	Key   []byte
	Value []byte
	Level []*Level
}

func NewSimpleSkl() *SimpleSkl {
	head := newNode(nil, nil, defaultLevel)
	return &SimpleSkl{
		height: 1,
		head:   head,
	}
}

func newNode(key, value []byte, level int) *SimpleSklNode {
	return &SimpleSklNode{
		Key:   key,
		Value: value,
		Level: NewLevel(level),
	}
}

func (s *SimpleSkl) randomHeight() int {
	h := 1
	for h < defaultLevel && z.FastRand() <= heightIncrease {
		h++
	}
	return h
}

// Get  从 skl.height 最高的地方开始 朝下查找
func (s *SimpleSkl) Get(key []byte) ([]byte, bool) {
	return nil, false
}

func (s *SimpleSkl) Put(key, value []byte) {
	currentHeight := s.getHeight()
	newNodeHeight := s.randomHeight()
	node := newNode(key, value, newNodeHeight)
	var prev [defaultLevel + 1]*SimpleSklNode
	var next [defaultLevel + 1]*SimpleSklNode
	// 新节点
	for int(currentHeight) < newNodeHeight {
		if atomic.CompareAndSwapInt32(&s.height, currentHeight, int32(newNodeHeight)) {
			break
		}
		currentHeight = s.getHeight()
	}

	// 寻找 新节点插入的位置
	// 1. 当  before < key < next
	// 2. 当  key > latest node key
	for i := 0; i < newNodeHeight; i++ {
		for {
			prev[i], next[i] = s.findSpliceForLevel(key, s.head, i)
			prev[i].Level[i].next = node
			if next[i] != nil {
				node.Level[i] = next[i].Level[i]
			}
			break
		}
	}
}

func (s *SimpleSkl) findSpliceForLevel(key []byte, beforeNode *SimpleSklNode, level int) (before, next *SimpleSklNode) {
	for {
		splice := beforeNode.Level[level]
		if splice == nil {
			return beforeNode, nil
		}

		if splice.next != nil {
			cmp := bytes.Compare(splice.next.Key, key)
			if cmp == 0 {
				return splice.next, splice.next
			}

			if cmp < 0 {
				return beforeNode, splice.next
			}
			beforeNode = splice.next
			continue
		}
		return beforeNode, nil
	}
}

func (s *SimpleSkl) findNear() {

}

func (s *SimpleSkl) getHeight() int32 {
	return atomic.LoadInt32(&s.height)
}

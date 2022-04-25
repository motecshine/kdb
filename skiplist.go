package kdb

import (
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

func NewLevel(level int) (levels []*Level) {
	for i := 0; i < level; i++ {
		levels[i] = new(Level)
	}
	return
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
	// 先查一下key是否存在
	currentHeight := s.getHeight()
	// 更新逻辑
	for i := int(currentHeight) - 1; i >= 0; i-- {
		print("not impl!")
	}

	// 新节点
	newNodeHeight := s.randomHeight()
	node := newNode(key, value, newNodeHeight)

	for int(currentHeight) < newNodeHeight {
		if atomic.CompareAndSwapInt32(&s.height, currentHeight, int32(newNodeHeight)) {
			break
		}
		currentHeight = s.getHeight()
	}

	// 寻找 新节点插入的位置
	for i := 0; i < newNodeHeight; i++ {

	}
}

func (s *SimpleSkl) findSpliceForLevel(key []byte, beforeNode *SimpleSklNode, level int) {

}

func (s *SimpleSkl) findNear() {

}

func (s *SimpleSkl) getHeight() int32 {
	return atomic.LoadInt32(&s.height)
}

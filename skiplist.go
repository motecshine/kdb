package kdb

import (
	"bytes"
	"github.com/dgraph-io/ristretto/z"
	"github.com/rs/zerolog/log"
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
	Key    []byte
	Value  []byte
	height int32
	Level  []*Level
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
		Key:    key,
		Value:  value,
		Level:  NewLevel(level),
		height: int32(level),
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
	x := s.head
	level := s.getHeight() - 1
	for {
		next := x.Level[level].next
		if next == nil {
			if level > 0 {
				log.Debug().Msgf("current level: %d next node is nil ", level)
				level--
				continue
			}
			return nil, false
		}

		cmp := s.CompareKeys(key, next.Key)
		log.Debug().Msgf("current level: %d, cmp: %d, key:%s next key: %s", level, cmp, string(key), string(next.Key))
		if cmp > 0 {
			x = next
			continue
		}
		if cmp == 0 {
			return x.Value, false
		}

		// cmp < 0. In other words, x.key < key < next.
		if level > 0 {
			level--
			continue
		}

		// Try to return x. Make sure it is not a head node.
		if x == s.head {
			return nil, false
		}
		return nil, false
	}
}

func (s *SimpleSkl) Put(key, value []byte) {
	currentMaxHeight := s.getHeight()
	var prev [defaultLevel + 1]*SimpleSklNode
	var next [defaultLevel + 1]*SimpleSklNode
	// 先看看key 存不存在
	prev[currentMaxHeight] = s.head
	for i := currentMaxHeight - 1; i >= 0; i-- {
		// 遍历同level 所有的节点
		var canUpdate bool
		prev[i], next[i], canUpdate = s.findSpliceNode(prev[i+1], key, i)
		if canUpdate && (prev[i] == next[i]) {
			log.Info().Msgf("find same key, value will update.")
			next[i].Value = value
			return
		}
	}
	// 如果没找到 创建新的node
	newNodeHeight := s.randomHeight()
	node := newNode(key, value, newNodeHeight)

	// 更新高度
	for newNodeHeight > int(currentMaxHeight) {
		if atomic.CompareAndSwapInt32(&s.height, currentMaxHeight, int32(newNodeHeight)) {
			break
		}
		currentMaxHeight = s.getHeight()
	}
	for i := 0; i < newNodeHeight; i++ {
		for {
			if prev[i] == nil {
				if i <= 1 {
					panic("can not append level0 ")
				}
				var canUpdate bool
				prev[i], next[i], canUpdate = s.findSpliceNode(s.head, key, int32(i))
				if canUpdate {
					panic("update node expect")
				}
			}
			node.Level[i].next = next[i]
			prev[i].Level[i].next = node
			break
		}
	}
}

func (s *SimpleSkl) findSpliceNode(beforeNode *SimpleSklNode, key []byte, level int32) (startNode, nextNode *SimpleSklNode, update bool) {
	for {
		if beforeNode == nil {
			panic("unexpected level")
		}
		nextNode := beforeNode.Level[level].next
		if nextNode == nil {
			return beforeNode, nextNode, false
		}
		cmp := s.CompareKeys(key, nextNode.Key)
		if cmp == 0 {
			return nextNode, nextNode, true
		}
		if cmp < 0 {
			return beforeNode, nextNode, false
		}
		if nextNode == nil {
			panic("unexpected level")
		}
		beforeNode = nextNode
	}
	return nil, nil, false
}

func (s *SimpleSkl) CompareKeys(prev, current []byte) int {
	return bytes.Compare(prev, current)
}

func (s *SimpleSkl) getHeight() int32 {
	return atomic.LoadInt32(&s.height)
}

func (s *SimpleSkl) Print() {
	currentMaxHeight := s.getHeight()
	for i := currentMaxHeight - 1; i >= 0; i-- {
		current := s.head.Level[i].next
		log.Info().Msgf("%d", i)
		for {
			if current == nil {
				break
			}
			log.Info().Msgf("height:%d, key:%s, value:%s", i, string(current.Key), string(current.Value))
			current = current.Level[i].next
		}
	}
}

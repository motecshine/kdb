package kdb

import (
	"github.com/rs/zerolog/log"
	"go.uber.org/atomic"
)

type color byte

const (
	red color = iota
	black
)

type rbNode struct {
	left   *rbNode
	right  *rbNode
	parent *rbNode
	key    []byte
	value  []byte
	color  color
	height int64
}

type Avl interface {
	RotateLeft(node *node)
	RotateRight(node *node)
	Insert(newNode *node)
}

// RB
// Degree: max(left.balance, right.balance)
type RB struct {
	iterator
	Avl
	storage
	root             *rbNode
	IteratorNextNode *rbNode
	size             atomic.Int64
}

// NewRB 红黑树
// 1. 节点是红色或黑色。
// 2. 根是黑色。
// 3. 所有叶子都是黑色（叶子是NIL节点）。
// 4. 每个红色节点必须有两个黑色的子节点。（从每个叶子到根的所有路径上不能有两个连续的红色节点。）
// 5. 从任一节点到其每个叶子的所有简单路径都包含相同数目的黑色节点。
func NewRB() *RB {
	log.Info().Msg("New RB.")
	return &RB{}
}

func (rb *RB) RotateLeft(node *node) {

}

func (rb *RB) RotateRight(node *node) {

}

func (rb *RB) Put(key []byte, value []byte) ([]byte, bool) {

	return nil, false
}

// Insert
// 1. 检测当前分支是否满足：
// 	  a. 所有叶子都是黑色（叶子是NIL节点）。
//    b. 每个红色节点必须有两个黑色的子节点。（从每个叶子到根的所有路径上不能有两个连续的红色节点。）
//    c. 从任一节点到其每个叶子的所有简单路径都包含相同数目的黑色节点。
// 2. 使当前树满足 Avl Tree 特征
func (rb *RB) Insert(newNode *node) {

	return
}

func (rb *RB) Get(key []byte) ([]byte, bool) {
	return nil, false
}

func (rb *RB) Next() *node {
	return nil
}

func (rb *RB) HasNext() bool {
	return rb.IteratorNextNode != nil
}

func (rb *RB) Rewind() {
	return
}

func (rb *RB) IntoIter() *RB {
	if rb.root == nil {
		return nil
	}
	rb.IteratorNextNode = rb.root
	return rb
}

func (rb *RB) Walk(fn func(node *node)) {

}

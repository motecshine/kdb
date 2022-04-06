package kdb

import (
	"bytes"
	"github.com/rs/zerolog/log"
	"go.uber.org/atomic"
)

type color byte

const (
	red color = iota
	black
)

type node struct {
	left   *node
	right  *node
	parent *node
	key    []byte
	value  []byte
	color  color
	height int64
}

func (n *node) SetValue(value []byte) {
	n.value = value
}

func (n *node) Value() []byte {
	return n.value
}

func (n *node) Key() []byte {
	return n.key
}

func (n *node) Height() int64 {
	return n.height
}

func (n *node) SetHeight(height int64) {
	n.height = height
}

func (n *node) Color() color {
	return n.color
}

func (n *node) SetColor(color color) {
	n.color = color
}

func (n *node) Right() *node {
	return n.right
}

func (n *node) SetRight(right *node) {
	n.right = right
}

func (n *node) Left() *node {
	return n.left
}

func (n *node) SetLeft(left *node) {
	n.left = left
}

func (n *node) Parent() *node {
	return n.parent
}

func (n *node) SetParent(parent *node) {
	n.parent = parent
}

type Avl interface {
	RotateLeft(node *node)
	RotateRight(node *node)
	Put(key []byte, value []byte) ([]byte, bool)
	Get(key []byte) ([]byte, bool)
	Insert(newNode *node)
	Update(parent *node, key []byte, value []byte)
}

// RB
// Degree: max(left.balance, right.balance)
type RB struct {
	Iterator
	Avl
	root             *node
	IteratorNextNode *node
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
	if rb.root == nil {
		rb.root = &node{
			left:   nil,
			right:  nil,
			parent: nil,
			key:    key,
			value:  value,
			color:  black,
		}
		return nil, false
	}

	var (
		cmp    int
		parent *node
	)
	rb.Walk(func(node *node) {
		cmp = bytes.Compare(key, node.Key())
		if cmp == 0 {
			log.Info().Msgf("update current: %s, old: %s", string(key), string(node.Key()))
			node.SetValue(value)
		}
	})
	// 更新
	{
		current := rb.root
		for current != nil {
			parent = current
			cmp = bytes.Compare(key, current.key)
			// 找到 old key
			if cmp == 0 {

			}
			// 如果当前key值小于0,那我们就遍历左枝,否则我们遍历右枝.
			if cmp < 0 {
				log.Info().Msgf("update current: %s, old: %s, 遍历左支，", string(key), string(current.key))
				current = current.left
			} else {
				log.Info().Msgf("update current: %s, old: %s, 遍历右支，", string(key), string(current.key))
				current = current.right
			}
		}
	}

	// 新增
	{
		newNode := &node{nil, nil, nil, key, value, red, 0}
		if cmp < 0 {
			parent.left = newNode
		} else {
			parent.right = newNode
		}
		newNode.parent = parent
		log.Info().Msgf("insert current: %+v.", newNode)
		rb.Insert(newNode)
		rb.size.Add(1)
	}
	return nil, false
}

func (rb *RB) Update() {

}

// Insert
// 1. 检测当前分支是否满足：
// 	  a. 所有叶子都是黑色（叶子是NIL节点）。
//    b. 每个红色节点必须有两个黑色的子节点。（从每个叶子到根的所有路径上不能有两个连续的红色节点。）
//    c. 从任一节点到其每个叶子的所有简单路径都包含相同数目的黑色节点。
// 2. 使当前树满足 Avl Tree 特征
func (rb *RB) Insert(newNode *node) {
	// 我们先旋转，然后再更新颜色.
	current := newNode
	// @todo: 提到上面逻辑
	// 在左枝
	if bytes.Compare(current.key, current.parent.left.key) == 0 {
		// 当前左枝高度大于parent右子树。
		if (current.height - current.Parent().Right().Height()) == 2 {

		}
	}
	// 在右枝
	if bytes.Compare(current.key, current.parent.right.key) == 0 {
		// 当前右枝高度大于parent左子树。
		if (current.height - current.Parent().Left().Height()) == 2 {

		}
	}

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

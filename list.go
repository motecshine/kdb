package kdb

import "bytes"

type list struct {
	storage
	iterator
	node   *listNode
	cursor *listNode
	head   *listNode
	tail   *listNode
}

type listNode struct {
	node
	next  *listNode
	key   []byte
	value []byte
}

func (ln *listNode) Key() ([]byte, bool) {
	if ln != nil {
		return ln.key, true
	}
	return nil, false
}

func (ln *listNode) Value() []byte {
	if ln != nil {
		return ln.value
	}
	return nil
}

func NewList() *list {
	return &list{}
}

func (l *list) Next() node {
	current := l.cursor.next
	l.cursor = current
	return current
}

func (l *list) HasNext() bool {
	if l.cursor == nil {
		return false
	}
	return l.cursor.next != nil
}

func (l *list) Visit(action func(key, value []byte) bool) bool {
	// rewind cursor
	l.Rewind()
	current := l.node
	for current != nil {
		if exists := action(current.key, current.value); exists {
			return true
		}
		current = l.Next().(*listNode)
	}
	return false
}

func (l *list) Rewind() {
	l.cursor = l.node
	return
}

func (l *list) Get(key []byte) ([]byte, bool) {
	exists := l.Visit(func(newKey, value []byte) bool {
		if bytes.Compare(newKey, key) == 0 {
			return true
		}
		return false
	})
	if l.cursor == nil {
		return nil, false
	}
	if exists {
		return l.cursor.Value(), true
	}
	return nil, false
}

func (l *list) Put(key []byte, value []byte) {
	newNode := &listNode{
		next:  nil,
		key:   key,
		value: value,
	}
	if l.head == nil {
		l.node = newNode
		l.head = newNode
		l.cursor = newNode
		l.tail = newNode
		return
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
	return
}

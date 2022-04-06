package kdb

type Iterator interface {
	Next() *node
	HasNext() bool
	Walk(action func(key, value []byte))
	Rewind()
}

package kdb

type iterator interface {
	Next() node
	HasNext() bool
	Visit(action func(key, value []byte) bool) bool
	Rewind()
}

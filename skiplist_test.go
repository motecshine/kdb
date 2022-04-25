package kdb

import "testing"

func TestSimpleSkl_Get(t *testing.T) {

}

func TestSimpleSkl_Put(t *testing.T) {
	skl := NewSimpleSkl()
	skl.Put([]byte("hello"), []byte("world"))
}

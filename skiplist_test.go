package kdb

import "testing"

func TestSimpleSkl_Get(t *testing.T) {

}

func TestSimpleSkl_Put(t *testing.T) {
	skl := NewSimpleSkl()
	skl.Put([]byte("hello"), []byte("world"))
	skl.Put([]byte("hello1"), []byte("world1"))
	skl.Put([]byte("hello2"), []byte("world2"))
	skl.Put([]byte("hello3"), []byte("world3"))
	skl.Put([]byte("hello4"), []byte("world4"))
	skl.Put([]byte("hello5"), []byte("world5"))
	skl.Put([]byte("hello6"), []byte("world6"))
	skl.Put([]byte("hello7"), []byte("world7"))
}

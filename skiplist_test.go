package kdb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleSkl(t *testing.T) {
	skl := NewSimpleSkl()
	skl.Put([]byte("hello"), []byte("world"))
	skl.Put([]byte("hello1"), []byte("world1"))
	skl.Put([]byte("hello2"), []byte("world2"))
	skl.Put([]byte("hello3"), []byte("world3"))
	skl.Put([]byte("hello4"), []byte("world4"))
	skl.Put([]byte("hello5"), []byte("world5"))
	skl.Put([]byte("hello6"), []byte("world6"))
	skl.Put([]byte("hello7"), []byte("world7"))

	// test find
	v, b := skl.Get([]byte("hello"))
	if !b {
		t.Failed()
	}
	assert.Equal(t, []byte("world"), v)

	v, b = skl.Get([]byte("hello1"))
	if !b {
		t.Failed()
	}
	assert.Equal(t, []byte("world1"), v)

	// update
	skl.Put([]byte("hello"), []byte("new world"))

	v, b = skl.Get([]byte("hello"))
	if !b {
		t.Failed()
	}
	assert.Equal(t, []byte("new world"), v)
	skl.Print()
}

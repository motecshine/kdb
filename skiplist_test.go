package kdb

import "testing"

func TestSimpleSkl_Get(t *testing.T) {
	skl := fakeData()
	value, _ := skl.Get([]byte("hello1"))
	t.Logf("%s", string(value))
	value2, _ := skl.Get([]byte("hello 2"))
	t.Logf("%s", string(value2))

	value3, _ := skl.Get([]byte("hello5"))
	t.Logf("%s", string(value3))
}

func TestSimpleSkl_Put(t *testing.T) {
	skl := fakeData()
	skl.Print()
}

func TestSimpleSkl_Put_2(t *testing.T) {
	skl := fakeData()
	skl.Print()

	skl.Put([]byte("hello7"), []byte("new hello7 value"))
	skl.Print()
}

func fakeData() *SimpleSkl {
	skl := NewSimpleSkl()
	skl.Put([]byte("hello"), []byte("world"))
	skl.Put([]byte("hello1"), []byte("world1"))
	skl.Put([]byte("hello2"), []byte("world2"))
	skl.Put([]byte("hello3"), []byte("world3"))
	skl.Put([]byte("hello4"), []byte("world4"))
	skl.Put([]byte("hello5"), []byte("world5"))
	skl.Put([]byte("hello6"), []byte("world6"))
	skl.Put([]byte("hello7"), []byte("world7"))
	return skl
}

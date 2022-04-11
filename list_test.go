package kdb

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestListPut(t *testing.T) {
	data := [][]string{
		{"key1", "value1"},
		{"key2", "value2"},
		{"key3", "value3"},
		{"key4", "value4"},
		{"key5", "value5"},
		{"key6", "value6"},
		{"key7", "value7"},
	}
	l := NewList()
	for _, d := range data {
		l.Put([]byte(d[0]), []byte(d[1]))
	}
	i := 1
	l.Visit(func(key, value []byte) bool {
		require.Equal(t, string(key), fmt.Sprintf("key%d", i))
		require.Equal(t, string(value), fmt.Sprintf("value%d", i))
		i++
		return false
	})
	i = 1
	// test rewind
	l.Visit(func(key, value []byte) bool {
		require.Equal(t, string(key), fmt.Sprintf("key%d", i))
		require.Equal(t, string(value), fmt.Sprintf("value%d", i))
		i++
		return false
	})
}

func TestListGet(t *testing.T) {
	data := [][]string{
		{"key1", "value1"},
		{"key2", "value2"},
		{"key3", "value3"},
		{"key4", "value4"},
		{"key5", "value5"},
		{"key6", "value6"},
		{"key7", "value7"},
	}
	l := NewList()
	for _, d := range data {
		l.Put([]byte(d[0]), []byte(d[1]))
	}

	for _, d := range data {
		v, exists := l.Get([]byte(d[0]))
		if !exists {
			t.Fail()
		}
		if bytes.Compare(v, []byte(d[1])) != 0 {
			t.Fail()
		}
	}
}

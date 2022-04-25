package kdb

import "github.com/dgraph-io/ristretto/z"

type wal struct {
	*z.MmapFile
}

type walEntry struct {
	key   []byte
	value []byte
}

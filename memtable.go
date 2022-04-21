package kdb

import "bytes"

type memoryTable struct {
	storage storage
	wal     file
	buf     *bytes.Buffer
}

func newMemoryTable(storage storage) *memoryTable {
	return &memoryTable{
		storage: storage,
	}
}

func (m *memoryTable) Put(key, value []byte) error {
	return nil
}

func (m *memoryTable) Get() {

}

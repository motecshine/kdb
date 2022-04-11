package kdb

type memoryTable struct {
	storage storage
}

func newMemoryTable(storage storage) *memoryTable {
	return &memoryTable{
		storage: storage,
	}
}

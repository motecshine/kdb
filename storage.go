package kdb

const (
	RB_DRIVER = iota
	SKL_DRIVER
)

type storage interface {
	Put(key []byte, value []byte)
	Get(key []byte) ([]byte, bool)
}

type node interface {
	Key() ([]byte, bool)
	Value() []byte
}

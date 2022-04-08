package kdb

const (
	RB_DRIVER = iota
	SKL_DRIVER
)

type Storage interface {
	Put(key []byte, value []byte) ([]byte, bool)
	Get(key []byte) ([]byte, bool)
}

# KDB



## SkipList Arena 内存对齐问题

### 场景

假设我们有一结构体长下面这样：

```golang 
const (
	maxHeight      = 20
	heightIncrease = math.MaxUint32 / 3
)

const (
	offsetSize = int(unsafe.Sizeof(uint32(0)))
	nodeAlign = int(unsafe.Sizeof(uint64(0))) - 1
	MaxNodeSize = int(unsafe.Sizeof(node{}))
)

type node struct {
	
	value uint64

	keyOffset uint32 // Immutable. No need to lock to access key.
	keySize   uint16 // Immutable. No need to lock to access key.

	height uint16

	tower [maxHeight]uint32
}

func putNode(height int) uint32 {
	// Compute the amount of the tower that will never be used, since the height
	// is less than maxHeight.
	// 假设新node 的高度10
	// unusedSize = (20 - 10) * 4 = 40
	// 假设新node 的高度3
	// unusedSize = (20 - 3) * 4 = 68
	// 76
	unusedSize := (maxHeight - height) * offsetSize

	// Pad the allocation with enough bytes to ensure pointer alignment.
	// 	假设新node 的高度10
	//   96 - 40 + 7  = 63
	// 	 假设新node 的高度3
	//   96 - 68 + 7  = 35
	// MaxNodeSize - unusedSize  是tower[height] 要使用的 size 。
	l := uint32(MaxNodeSize - unusedSize + nodeAlign)
	n := s.allocate(l)

	// Return the aligned offset.

	m := (n + uint32(nodeAlign)) & ^uint32(nodeAlign)
	return m
}

```





## Special thanks
- [Red-Black Tree](https://github.com/krasun/rbytree)
- [深入浅出分析LSM树] (https://zhuanlan.zhihu.com/p/415799237)
- [skiplist] (https://lengrongfu.github.io/2019-07-03-Badger%E6%BA%90%E7%A0%81%E5%AD%A6%E4%B9%A0%E4%B9%8B-SkipList%E5%9F%BA%E7%A1%80%E7%AF%87/)
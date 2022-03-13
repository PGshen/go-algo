package container

/**
 * 基于bitmap实现hashset
 */
type MyHashSet struct {
	bitset []uint64
}

func NewHashSet() MyHashSet {
	return MyHashSet{
		bitset: []uint64{},
	}
}

func (s *MyHashSet) Add(key int) {
	bit := key % 64
	length := key / 64
	for i := len(s.bitset); i < length; i++ { // 扩容
		s.bitset = append(s.bitset, 0)
	}
	s.bitset[length] = s.bitset[length] | (1 << bit) // 将对应位设置为1
}

func (s *MyHashSet) Remove(key int) {
	bit := key % 64
	length := key / 64
	if length >= len(s.bitset) {
		return
	}
	s.bitset[length] = s.bitset[length] & ^(1 << bit) // 将对应位设置为0
}

func (s *MyHashSet) Contains(key int) bool {
	bit := key % 64
	length := key / 64
	if length >= len(s.bitset) {
		return false
	}
	return s.bitset[length]&(1<<bit) != 0
}

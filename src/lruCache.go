package leetcode

import (
	"container/list"
)

type KV struct {
	key   string
	value string
}
type LRUCache struct {
	size  int
	cache map[string]*list.Element
	list  *list.List
}

// 初始化设置缓存容量
func Constructor(s int) LRUCache {
	return LRUCache{
		size:  s,
		list:  list.New(),
		cache: make(map[string]*list.Element),
	}
}
func (this *LRUCache) Get(key string) string {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		return elem.Value.(KV).value
	}
	return ""
}

func (this *LRUCache) Put(key string, value string) {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		elem.Value = KV{key, value}
	} else {
		if this.list.Len() >= this.size {
			delete(this.cache, this.list.Back().Value.(KV).key)
			this.list.Remove(this.list.Back())
		}
		this.list.PushFront(KV{key, value})
		this.cache[key] = this.list.Front()
	}
}

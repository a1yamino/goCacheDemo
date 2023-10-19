package LRU

import "container/list"

type Cache struct {
	maxLength int64
	length    int64
	Ll        *list.List
	Cache     map[string]*list.Element
	OnEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

func New(maxLen int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxLength: maxLen,
		Ll:        list.New(),
		Cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (Value, bool) {
	if ele, isOk := c.Cache[key]; isOk {
		c.Ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return nil, false
}

func (c *Cache) RemoveOldest() {
	ele := c.Ll.Back()
	if ele != nil {
		c.Ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.Cache, kv.key)
		c.length -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

func (c *Cache) Add(key string, value Value) {
	if ele, isOk := c.Cache[key]; isOk {
		c.Ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.length += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.Ll.PushFront(&entry{key, value})
		c.Cache[key] = ele
		c.length += int64(len(key)) + int64(value.Len())
	}
	for c.maxLength != 0 && c.maxLength < c.length {
		c.RemoveOldest()
	}

}

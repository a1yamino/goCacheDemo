package goCacheDemo

import (
	"fmt"
	"goCacheDemo/LRU"
)

func main() {
	cache := LRU.New(64, nil)

	cache.Add("a", String("aaa"))
	cache.Add("b", String("bbb"))
	cache.Add("c", String("ccc"))

	for e := cache.Ll.Front(); e != nil; e = e.Next() {
		fmt.Printf("%+v\n", e.Value)
	}
}

type String string

func (s String) Len() int {
	return len(s)
}

package util

import "strings"

type CachMap struct {
	Cache map[string]interface{}
}

//在这里接口怎么用啊
type Cacher interface {
	Contain(key string) bool
	Empty() bool
	GetVal(key string) interface{}
	PutVal(key string, val interface{}) bool
	Size() int
}

func GetIns() *CachMap {
	CachMap := new(CachMap)
	CachMap.Cache = make(map[string]interface{}, 1000)
	return CachMap
}
func (c *CachMap) Size() int {
	if c.Cache != nil {
		return len(c.Cache)
	} else {
		return -1
	}
}

func (c *CachMap) Contain(key string) bool {
	val, ok := c.Cache[key]
	if ok && val != nil {
		return true
	} else {
		return false
	}
}

func (c *CachMap) Remove(key string) bool {
	if c.Contain(key) {
		delete(c.Cache, key)
		return true
	}
	return false

}
func (c *CachMap) RemoveAll(key string) {
	c.Cache = make(map[string]interface{}, 1000)
}

func (c *CachMap) Empty() bool {
	if len(c.Cache) > 0 {
		return false
	} else {
		return true
	}
}

func (c *CachMap) Get(key string) interface{} {
	val, ok := c.Cache[key]
	if ok {
		return val
	} else {
		return ok
	}
}

func (c *CachMap) Put(key string, val interface{}) interface{} {
	if len(strings.TrimSpace(key)) == 0 {
		return false
	} else {
		c.Cache[key] = val
	}
	return true
}

package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	myMap map[string]interface{}
	rw    sync.RWMutex
}

// NewSafeMap 创建 SafeMap 实例
func NewSafeMap() *SafeMap {
	return &SafeMap{
		myMap: make(map[string]interface{}),
	}
}

// Set 设置键值对
func (m *SafeMap) Set(str string, val interface{}) {
	m.rw.Lock()
	defer m.rw.Unlock()
	m.myMap[str] = val
}

// Get 获取键对应的值和键是否存在
func (m *SafeMap) Get(key string) (val interface{}, b bool) {
	m.rw.RLock()
	defer m.rw.RUnlock()
	val, b = m.myMap[key]
	return val, b
}

// Del 删除指定的键
func (m *SafeMap) Del(key string) {
	m.rw.Lock()
	defer m.rw.Unlock()
	delete(m.myMap, key)
}

// Len 获取 Map 中键值对数量
func (m *SafeMap) Len() (l int) {
	m.rw.RLock()
	defer m.rw.RUnlock()
	l = len(m.myMap)
	return l
}

func main() {
	m := NewSafeMap()
	m.Set("name", "JustGopher")
	v, exists := m.Get("name")
	fmt.Println(v, exists)
	fmt.Println(m.Len())
	m.Del("name")
	fmt.Println(m.Len())
}

package main

import "fmt"

type SafeMap struct {
	myMap map[string]interface{}
	ops   chan func()
}

func NewSafeMap() *SafeMap {
	sm := &SafeMap{
		myMap: make(map[string]interface{}),
		ops:   make(chan func()),
	}
	go sm.run()
	return sm
}

func (sm *SafeMap) run() {
	for op := range sm.ops {
		op()
	}
}

func (sm *SafeMap) Set(key string, val interface{}) {
	sm.ops <- func() {
		sm.myMap[key] = val
	}
}

func (sm *SafeMap) Get(key string) (val interface{}, exists bool) {
	type res struct {
		val    interface{}
		exists bool
	}
	resChan := make(chan res)
	sm.ops <- func() {
		val, exists = sm.myMap[key]
		resChan <- res{val, exists}
	}
	r := <-resChan
	return r.val, r.exists
}

func (sm *SafeMap) Del(key string) {
	sm.ops <- func() {
		delete(sm.myMap, key)
	}
}

func (sm *SafeMap) Len() int {
	l := make(chan int)
	sm.ops <- func() {
		l <- len(sm.myMap)
	}
	return <-l
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

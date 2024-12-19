package byChan

// --- todo -- 不会写

type SafeMap struct {
	myMap map[string]interface{}
	c     chan struct{}
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		myMap: make(map[string]interface{}),
		c:     make(chan struct{}),
	}
}

//func (m *SafeMap) Get() {
//
//}
//
//func (m *SafeMap) Set() {
//
//}
//
//func (m *SafeMap) Del() {
//
//}
//
//func (m *SafeMap) Len() {
//
//}

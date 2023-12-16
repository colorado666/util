package amap

import (
	"container/list"
	"sync"
)

type ListMap struct {
	Mu       sync.RWMutex
	DataList *list.List
	DataMap  map[string]interface{}
}

func NewListMap() *ListMap {
	l := new(ListMap)
	l.DataMap = make(map[string]interface{})
	l.DataList = list.New()
	return l
}

func (l *ListMap) Lock() {
	l.Mu.Lock()
}

func (l *ListMap) Unlock() {
	l.Mu.Unlock()
}

func (l *ListMap) RLock() {
	l.Mu.RLock()
}

func (l *ListMap) RUnlock() {
	l.Mu.RUnlock()
}

func (l *ListMap) Len() int {
	return l.DataList.Len()
}

//是否已经存在
func (l *ListMap) IsExist(key string) bool {
	_, ok := l.DataMap[key]
	return ok
}

func (l *ListMap) PushFront(item interface{}, key string, value interface{}) {
	if !l.IsExist(key) {
		l.DataList.PushFront(item)
		l.DataMap[key] = value
	}
}

func (l *ListMap) PushBack(item interface{}, key string, value interface{}) {
	if !l.IsExist(key) {
		l.DataList.PushBack(item)
		l.DataMap[key] = value
	}
}

func (l *ListMap) Front() *list.Element {
	return l.DataList.Front()
}

func (l *ListMap) Back() *list.Element {
	return l.DataList.Back()
}

func (l *ListMap) Remove(e *list.Element, key string) {
	l.DataList.Remove(e)
	delete(l.DataMap, key)
}

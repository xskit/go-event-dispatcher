package event

import (
	"sort"
	"sync"
)

var (
	once sync.Once

	dispatcher *EventDispatcher
)

// EventDispatcher 事件管理器
type EventDispatcher struct {
	listeners map[string]*Listener
}

type listenerSlice []*Listener

func (l listenerSlice) Len() int {
	return len(l)
}
func (l listenerSlice) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
func (l listenerSlice) Less(i, j int) bool {
	return (*l[i]).Order() > (*l[j]).Order()
}

// GetEventDispatcher 事件管理器
func GetEventDispatcher() *EventDispatcher {
	once.Do(func() {
		dispatcher = &EventDispatcher{
			listeners: make(map[string]*Listener),
		}
	})
	return dispatcher
}

// AddListener 添加事件监听器
func (e *EventDispatcher) AddListener(listener *Listener) {
	for _, v := range (*listener).Listen() {
		e.listeners[v] = listener
	}
}

// Dispatch 触发一次事件
func (e EventDispatcher) Dispatch(event *Event) {
	name := (*event).Name()
	if name == "" {
		return
	}
	listeners := make([]*Listener, 0)
	for k, listener := range e.listeners {
		if name == k || name == "*" {
			listeners = append(listeners, listener)
		}
	}
	// 优化级
	sort.Sort(listenerSlice(listeners))
	// 开始调度
	for _, listener := range listeners {
		if false == (*listener).Process(event) {
			break
		}
	}

}

package event

import (
	"log"
	"testing"
)

type e1 struct {
}

func (e e1) Name() string {
	return "*"
}

func (e e1) Handler() {
	log.Println(e.Name(), ",Handler")
}

type e2 struct {
}

func (e e2) Name() string {
	return ""
}

type e3 struct{}

func (e e3) Name() string {
	return "TestEvent"
}

// 定义监听器一
type l1 struct{}

func (l l1) Order() int {
	return 0
}
func (l l1) Listen() []string {
	return []string{
		"TestEvent",
	}
}
func (l l1) Process(event *Event) bool {
	log.Println("listener 1")

	log.Println(
		"Dispatch:",
		(*event).Name(),
	)

	return true
}

// 监听二
type l2 struct{}

func (l l2) Order() int {
	return 1
}
func (l l2) Listen() []string {
	return []string{
		"TestEvent2",
	}
}
func (l l2) Process(event *Event) bool {
	log.Println("listener 2")
	log.Println(
		"Dispatch:",
		(*event).Name(),
	)

	if e,ok := (*event).(e1); ok {
		// 如果是 事件 e1
		e.Handler()
		// 阻止事件冒泡
		return false
	}

	return true
}



var ed *EventDispatcher

func TestMain(m *testing.M) {
	ed = GetEventDispatcher()
	m.Run()
}

func addListener(t *testing.T) {
	var lis1 ,lis2 Listener = &l1{},&l2{}
	ed.AddListener(&lis1)
	ed.AddListener(&lis2)
	t.Log("addListener done.")
}

func TestDispatcher(t *testing.T) {
	addListener(t)
	var event1 Event = &e1{}
	ed.Dispatch(&event1)
	var event2 Event = &e2{}
	ed.Dispatch(&event2)
	var event3 Event = &e3{}
	ed.Dispatch(&event3)

}

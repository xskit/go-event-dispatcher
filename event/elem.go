package event

type Listener interface {
	// Order 监听器优先级,越大越优化，默认0，按添加时的顺序
	Order() int
	// Listen 监听的事件集
	Listen() []string
	// Process 执行处理程序
	Process(event *Event) bool
}

type Event interface {
	// Name 事件名，*：会触发所有监听器, ""：空值不触发监听
	Name() string
}



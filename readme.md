## 事件机制
##### 概念
事件模式是一种经过了充分测试的可靠机制，是一种非常适用于解耦的机制，分别存在以下 3 种角色：
- 事件(Event) 是传递于应用代码与 监听器(Listener) 之间的通讯对象
- 监听器(Listener) 是用于监听 事件(Event) 的发生的监听对象
- 事件调度器(EventDispatcher) 是用于触发 事件(Event) 和管理 监听器(Listener) 与 事件(Event) 之间的关系的管理者对象

> 用通俗易懂的例子来说明就是，假设我们存在一个 user.Register() 方法用于注册一个账号，在账号注册成功后我们可以通过事件调度器触发 UserRegistered 事件，由监听器监听该事件的发生，在触发时进行某些操作，比如发送用户注册成功短信，在业务发展的同时我们可能会希望在用户注册成功之后做更多的事情，比如发送用户注册成功的邮件等待，此时我们就可以通过再增加一个监听器监听 UserRegistered 事件即可，无需在 user.register() 方法内部增加与之无关的代码。

#### 示例
```go
// 定义事件
// todo: Name() 返回事件名，如果事件名：“*”，该事件会触发所有的监听器，如果事件名：“” 空字符，不会触发监听器
type eventA struct {
}

func (e eventA) Name() string {
return "事件A"
}

// 定义监听器
// todo: 
type lentener1 struct{}

// 定义监听器相同事件下的执行顺序，越大越优先
func (l lentener1) Order() int {
    return 0
}
// 返回定义监听的事件名
func (l lentener1) Listen() []string {
    return []string{
        "事件A",
    }
}
func (l lentener1) Process(event *Event) bool {
    log.Println("listener 1")
    log.Println(
        "Dispatch:",
        (*event).Name(),
    )
    return true
}
// 获取事件调度器
dispatcher := event.GetEventDispatcher()

// 注册事件监听器
var l1 Listener = &lentener1{}
dispatcher.AddEventListener(&l1)

// 触发事件
var eA Event = &eventA{}
dispatcher.Dispatch(&eA)

```

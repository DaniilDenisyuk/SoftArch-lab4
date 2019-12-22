package engine

type Command interface {
	Execute(handler Handler)
}

// Handler allows to send commands to an event loop it's associated with.
type Handler interface {
	Post(cmd Command)
}

type EventLoop struct {
	messageQueue []Command
}

func (el *EventLoop) Start() {
  el.messageQueue := []Command
  while(!el.AwaitFinish())
}

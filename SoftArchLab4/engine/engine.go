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
	go func() {
		for {
			if len(el.messageQueue) > 0 {
				el.messageQueue[0].Execute(el)
				el.messageQueue = el.messageQueue[1:]
			}
		}
	}()
}

func (el *EventLoop) Post(cmd Command) {
	el.messageQueue = append(el.messageQueue, cmd)
}

func (el *EventLoop) AwaitFinish() {
	for len(el.messageQueue) > 0 {
	}
	el = nil
}

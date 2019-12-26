package engine
// Старався зробити максимально коротко та читабельно
import "time"

type Command interface {
	Execute(handler Handler)
	IsChild() bool
}

type Handler interface {
	Post(cmd Command)
}

type EventLoop struct {
	messageQueue chan Command
	isFinished   bool
	StopRcv      bool
}

func (el *EventLoop) Start() {
	el.messageQueue = make(chan Command, 5)
	el.isFinished = false
	go func() {
		for !el.isFinished {
			cmd := <- el.messageQueue
			cmd.Execute(el)
			}
	}()
}

func (el *EventLoop) Post(cmd Command) {
	if el.StopRcv {
		if cmd.IsChild() {
			el.messageQueue <- cmd
		}
		return
	}
	el.messageQueue <- cmd
}

func (el *EventLoop) AwaitFinish() {
	el.StopRcv = true
	for len(el.messageQueue) > 0 {
    time.Sleep(50 * time.Millisecond)
	}
	close(el.messageQueue)
	el.isFinished = true
}

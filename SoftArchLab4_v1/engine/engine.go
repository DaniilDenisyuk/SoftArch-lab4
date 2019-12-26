package engine

import "sync"

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(cmd Command)
}

type EventLoop struct {
	sync.WaitGroup
	mq   *messageQueue
	done bool
}

type messageQueue struct {
	sync.Mutex
	sync.WaitGroup
	data []Command
	wait bool
}

func (mq *messageQueue) push(cmd Command) {
	mq.Lock()
	defer mq.Unlock()
	mq.data = append(mq.data, cmd)
	if mq.wait {
		mq.wait = false
		mq.Done()
	}
}

func (mq *messageQueue) pull() Command {
	mq.Lock()
	defer mq.Unlock()
	if len(mq.data) == 0 {
		mq.wait = true
		mq.Add(1)
		mq.Unlock()
		mq.Wait()
		mq.Lock()
	}
	res := mq.data[0]
	mq.data = mq.data[1:]
	return res
}

func (el *EventLoop) Start() {
	el.mq = new(messageQueue)
	el.Add(1)
	go func() {
		defer el.Done()
		for (!el.done) || (len(el.mq.data) != 0) {
			cmd := el.mq.pull()
			cmd.Execute(el)
		}
	}()
}

type exitFunc func(loop Handler)

func (cmd exitFunc) Execute(loop Handler) {
	cmd(loop)
}

func (el *EventLoop) Post(cmd Command) {
	el.mq.push(cmd)
}

func (el *EventLoop) AwaitFinish() {
	el.Post(exitFunc(func(loop Handler) {
		loop.(*EventLoop).done = true
	}))
	el.Wait()
}

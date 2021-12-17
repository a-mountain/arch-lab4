package engine

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(command Command)
}

type EventLoop struct {
	stop       bool
	isFinished chan bool
	queue      SyncCommandQueue
	sleep      bool
}

func NewEventLoop() EventLoop {
	return EventLoop{
		stop:       false,
		isFinished: make(chan bool, 1),
		queue:      SyncCommandQueue{},
		sleep:      false,
	}
}

func (l *EventLoop) Post(command Command) {
	if l.sleep {
		l.Start()
	}
	l.queue.Push(command)
}

func (l *EventLoop) Start() {
	l.sleep = false
	go func() {
		for {
			stopEventLoop := l.queue.isEmpty() && l.stop
			if stopEventLoop {
				l.isFinished <- true
				return
			}
			executeNextCommand := !l.queue.isEmpty()
			if executeNextCommand {
				command := l.queue.Pull()
				command.Execute(l)
			} else {
				l.sleep = true
				return
			}
		}
	}()
}

func (l *EventLoop) AwaitFinish() {
	l.stop = true
	<-l.isFinished
}

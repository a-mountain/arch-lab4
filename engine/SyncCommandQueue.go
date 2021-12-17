package engine

import "sync"

type SyncCommandQueue struct {
	commands []Command
	lock     sync.Mutex
}

func (q *SyncCommandQueue) Push(command Command) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.commands = append(q.commands, command)
}

func (q *SyncCommandQueue) Pull() Command {
	q.lock.Lock()
	defer q.lock.Unlock()
	head := q.commands[0]
	q.commands[0] = nil
	q.commands = q.commands[1:]
	return head
}

func (q *SyncCommandQueue) isEmpty() bool {
	return len(q.commands) == 0
}

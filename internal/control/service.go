package control

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Service struct {
	queue *queue
}

type queue struct {
	counter  int64
	mu       sync.RWMutex
	programs []program
}

type program struct {
	commands []command
}

type command struct {
	command string
}

func NewService() *Service {
	s := Service{
		queue: newQueue(),
	}

	go s.serial()

	fmt.Println("Queue", s.queue)

	return &s
}

func newQueue() *queue {
	return &queue{
		programs: []program{},
	}
}

func (q *queue) enqueue(prog program) int64 {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.programs = append(q.programs, prog)
	q.counter++
	id := q.counter

	return id
}

func (q *queue) dequeue() (program, error) {
	q.mu.Lock()
	if len(q.programs) == 0 {
		q.mu.Unlock()
		return program{}, errors.New("queue is empty")
	}

	prog := q.programs[0]
	q.programs = q.programs[1:]

	q.mu.Unlock()

	return prog, nil
}

func (s *Service) serial() {
	for {
		program, err := s.queue.dequeue()
		if err != nil {
			time.Sleep(time.Millisecond * 8)
			continue
		}

		time.Sleep(time.Second * 10)

		_ = program // TODO: Do something with program
	}
}

func (s *Service) addCommand(cmd command) int64 {
	cmds := []command{}
	cmds = append(cmds, cmd)

	prog := program{
		commands: cmds,
	}

	return s.queue.enqueue(prog)
}

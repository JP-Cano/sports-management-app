package worker

import (
	"log"
	"sync"
)

type TaskFunc func(row []string) error

type Pool struct {
	numWorkers int
	taskFunc   TaskFunc
	wg         sync.WaitGroup
	errorChan  chan error
}

func New(numWorkers int, taskFunc TaskFunc, errorChan chan error) *Pool {
	return &Pool{
		numWorkers: numWorkers,
		taskFunc:   taskFunc,
		errorChan:  errorChan,
	}
}

func (p *Pool) Start(tasks <-chan []string) {
	log.Printf("Starting %d workers", p.numWorkers)
	for i := 0; i < p.numWorkers; i++ {
		p.wg.Add(1)
		go p.worker(tasks)
	}
}

func (p *Pool) Wait() {
	p.wg.Wait()
}

func (p *Pool) worker(tasks <-chan []string) {
	defer p.wg.Done()

	for task := range tasks {
		if err := p.taskFunc(task); err != nil {
			p.errorChan <- err
		}
	}
}

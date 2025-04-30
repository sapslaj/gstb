// Relatively simple worker pool implementation.
package workerpool

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrNoMessageHandler = errors.New("workerpool.WorkerPool.MessageHandler is not set")
)

type WorkerPool[V any] struct {
	WorkerCount    int
	MessageHandler func(V)
	MessageChannel chan V
	WaitGroup      sync.WaitGroup
	MesssageBuffer int
}

func (wp *WorkerPool[V]) Start() (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = fmt.Errorf("unknown panic: %v", r)
			}
		}
	}()
	if wp.MessageChannel == nil {
		wp.MessageChannel = make(chan V, wp.MesssageBuffer)
	}
	if wp.MessageHandler == nil {
		return ErrNoMessageHandler
	}
	for i := 0; i < wp.WorkerCount; i++ {
		wp.WaitGroup.Add(1)
		go func(i int) {
			defer wp.WaitGroup.Done()
			for msg := range wp.MessageChannel {
				wp.MessageHandler(msg)
			}
		}(i)
	}
	return
}

func (wp *WorkerPool[V]) Stop() (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = fmt.Errorf("unknown panic: %v", r)
			}
		}
	}()
	close(wp.MessageChannel)
	wp.WaitGroup.Wait()
	return
}

func (wp *WorkerPool[V]) Push(message V) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = fmt.Errorf("unknown panic: %v", r)
			}
		}
	}()
	wp.MessageChannel <- message
	return
}

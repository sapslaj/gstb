// Concurrency-safe simple ring buffer
package ring

import (
	"sync"
)

type Ring[T any] struct {
	sync.RWMutex
	Size       uint64
	Buffer     []T
	ReadIndex  uint64
	WriteIndex uint64
}

func (rb *Ring[T]) Push(v T) T {
	if rb.Buffer == nil {
		rb.Buffer = make([]T, rb.Size)
	}
	rb.Lock()
	defer rb.Unlock()
	rb.WriteIndex = (rb.WriteIndex + 1) % rb.Size
	rb.Buffer[rb.WriteIndex] = v
	return rb.Buffer[rb.WriteIndex]
}

func (rb *Ring[T]) Pop() T {
	if rb.Buffer == nil {
		rb.Buffer = make([]T, rb.Size)
	}
	rb.RLock()
	defer rb.RUnlock()
	rb.ReadIndex = (rb.ReadIndex + 1) % rb.Size
	return rb.Buffer[rb.ReadIndex]
}

func (rb *Ring[T]) HasNext() bool {
	return rb.ReadIndex != rb.WriteIndex
}

func (rb *Ring[T]) Next() (T, bool) {
	if !rb.HasNext() {
		var def T
		return def, false
	}
	return rb.Pop(), true
}

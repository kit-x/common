package syncx

import "sync"

// WaitGroupWrapper make WaitGroup easier
type WaitGroupWrapper struct {
	sync.WaitGroup
}

// Wrap callback and will done after handled (cb) func
func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}

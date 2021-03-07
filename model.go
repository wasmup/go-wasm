package main

import "sync/atomic"

// Model struct for a domain business logic.
// separate the presentation layer from the business logic.
type Model struct {
	counter int64
}

// Counter returns a Model counter value.
func (p *Model) Counter() int64 {
	return atomic.LoadInt64(&p.counter)
}

// SetCounter sets the Model counter value.
func (p *Model) SetCounter(counter int64) {
	atomic.StoreInt64(&p.counter, counter)
}

// Add to the Model counter value and return the new value.
func (p *Model) Add(counter int64) int64 {
	return atomic.AddInt64(&p.counter, counter)
}

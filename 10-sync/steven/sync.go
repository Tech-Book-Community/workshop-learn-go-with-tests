package main

import "sync"

type Counter struct {
	sync.RWMutex
	value int
}

func (c *Counter) Inc() int {
	c.Lock()
	defer c.Unlock()
	c.value++
	return c.value
}

func (c *Counter) Value() int {
	c.RLock()
	defer c.RUnlock()
	return c.value
}

// NewCounter returns a new Counter.
func NewCounter() *Counter {
	return &Counter{}
}

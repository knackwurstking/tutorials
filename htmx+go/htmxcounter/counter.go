package main

import "sync"

type Counter struct {
	value int
	mu    sync.Mutex
}

func NewCounter(initialValue int) *Counter {
	return &Counter{
		value: initialValue,
		mu:    sync.Mutex{},
	}
}

func (c *Counter) Decrease() {
	c.mu.Lock()
	c.value -= 1
	c.mu.Unlock()
}

func (c *Counter) Increase() {
	c.mu.Lock()
	c.value += 1
	c.mu.Unlock()
}

func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

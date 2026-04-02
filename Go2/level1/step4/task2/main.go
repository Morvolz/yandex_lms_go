package main

import "sync"

type Counter struct {
	value int
	mu    sync.RWMutex
}

type Сount interface {
	Increment()    // увеличение счётчика на единицу
	GetValue() int // получение текущего значения
}

func (c Counter) Increment() {
	c.mu.RLock()
	c.value++
	c.mu.RUnlock()
}

func (c Counter) GetValue() int {
	c.mu.RLock()
	val := c.value
	c.mu.RUnlock()
	return val
}

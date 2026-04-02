package main

import "sync"

type ConcurrentQueue struct {
	queue []interface{}
	mutex sync.Mutex
}

type Queue interface {
	Enqueue(element interface{}) // положить элемент в очередь
	Dequeue() interface{}        // забрать первый элемент из очереди
}

func (cq *ConcurrentQueue) Enqueue(element interface{}) {
	cq.mutex.Lock()
	cq.queue = append(cq.queue, element)
	cq.mutex.Unlock()
}

func (cq *ConcurrentQueue) Dequeue() interface{} {
	cq.mutex.Lock()
	data := cq.queue[0]
	cq.queue = cq.queue[1:]
	cq.mutex.Unlock()
	return data

}

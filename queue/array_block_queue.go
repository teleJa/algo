package queue

import "sync"

type ArrayBlockingQueue[T any] struct {
	data                       []T
	cap                        int // 容量
	notFull                    *sync.Cond
	notEmpty                   *sync.Cond
	putIndex, takeIndex, count int
}

func NewArrayBlockingQueue[T any](cap int) *ArrayBlockingQueue[T] {
	var mutex sync.Mutex
	return &ArrayBlockingQueue[T]{
		data:     make([]T, cap),
		cap:      cap,
		notFull:  sync.NewCond(&mutex),
		notEmpty: sync.NewCond(&mutex),
	}
}

func (q *ArrayBlockingQueue[T]) Offer(data T) bool {
	q.notFull.L.Lock()
	defer q.notFull.L.Unlock()
	if q.count == len(q.data) {
		return false
	}
	q.enqueue(data)
	return true
}

func (q *ArrayBlockingQueue[T]) Put(data T) {
	q.notFull.L.Lock()
	defer q.notFull.L.Unlock()
	for q.count == len(q.data) {
		q.notFull.Wait()
	}
	q.enqueue(data)
}

func (q *ArrayBlockingQueue[T]) enqueue(data T) {
	q.data[q.putIndex] = data
	q.putIndex++
	if q.putIndex == len(q.data) {
		q.putIndex = 0
	}
	q.count++
	q.notEmpty.Signal()
}

func (q *ArrayBlockingQueue[T]) dequeue() T {
	t := q.data[q.takeIndex]
	// TODO 移除元素
	q.takeIndex++
	if q.takeIndex == len(q.data) {
		q.takeIndex = 0
	}
	q.notFull.Signal()
	q.count--
	return t
}

func (q *ArrayBlockingQueue[T]) Take() T {
	q.notEmpty.L.Lock()
	defer q.notEmpty.L.Unlock()
	for q.count == 0 {
		q.notEmpty.Wait()
	}
	data := q.dequeue()
	return data
}

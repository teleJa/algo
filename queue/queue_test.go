package queue

import (
	"testing"
	"time"
)

func TestArrayBlockingQueue1(t *testing.T) {

	queue := NewArrayBlockingQueue[int](5)
	queue.Put(1)
	queue.Put(2)
	queue.Put(3)
	queue.Put(4)
	queue.Put(5)

	t.Log("el:", queue.Take())
	t.Log("el:", queue.Take())
	t.Log("el:", queue.Take())

}

func TestArrayBlockingQueue2(t *testing.T) {

	queue := NewArrayBlockingQueue[int](5)
	queue.Put(1)
	queue.Put(2)
	queue.Put(3)
	queue.Put(4)
	queue.Put(5)

	t.Log("offer:", queue.Offer(6))

}

func TestArrayBlockingQueue3(t *testing.T) {

	queue := NewArrayBlockingQueue[int](5)

	go func() {
		for {
			t.Log("take from queue:", queue.Take())
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			queue.Put(i)
		}
	}()

	select {}

}

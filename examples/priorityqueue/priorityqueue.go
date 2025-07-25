// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"cmp"

	pq "github.com/Arvin619/gods/queues/priorityqueue"
)

// Element is an entry in the priority queue
type Element struct {
	name     string
	priority int
}

// Compare method (sort by element's priority value in descending order)
func (e Element) Compare(another Element) int {
	return -cmp.Compare(e.priority, another.priority) // "-" descending order
}

// PriorityQueueExample to demonstrate basic usage of BinaryHeap
func main() {
	a := Element{name: "a", priority: 1}
	b := Element{name: "b", priority: 2}
	c := Element{name: "c", priority: 3}

	queue := pq.NewWith(Element.Compare) // empty
	queue.Enqueue(a)                     // {a 1}
	queue.Enqueue(c)                     // {c 3}, {a 1}
	queue.Enqueue(b)                     // {c 3}, {b 2}, {a 1}
	_ = queue.Values()                   // [{c 3} {b 2} {a 1}]
	_, _ = queue.Peek()                  // {c 3} true
	_, _ = queue.Dequeue()               // {c 3} true
	_, _ = queue.Dequeue()               // {b 2} true
	_, _ = queue.Dequeue()               // {a 1} true
	_, _ = queue.Dequeue()               // <nil> false (nothing to dequeue)
	queue.Clear()                        // empty
	_ = queue.Empty()                    // true
	_ = queue.Size()                     // 0
}

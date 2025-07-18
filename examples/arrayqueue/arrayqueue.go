// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import aq "github.com/Arvin619/gods/queues/arrayqueue"

// ArrayQueueExample to demonstrate basic usage of ArrayQueue
func main() {
	queue := aq.New[int]() // empty
	queue.Enqueue(1)       // 1
	queue.Enqueue(2)       // 1, 2
	_ = queue.Values()     // 1, 2 (FIFO order)
	_, _ = queue.Peek()    // 1,true
	_, _ = queue.Dequeue() // 1, true
	_, _ = queue.Dequeue() // 2, true
	_, _ = queue.Dequeue() // nil, false (nothing to dequeue)
	queue.Enqueue(1)       // 1
	queue.Clear()          // empty
	queue.Empty()          // true
	_ = queue.Size()       // 0
}

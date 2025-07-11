// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binaryheap

import (
	"github.com/Arvin619/gods/containers"
)

// Assert Serialization implementation
var (
	_ containers.JSONSerializer   = (*Heap[int])(nil)
	_ containers.JSONDeserializer = (*Heap[int])(nil)
)

// ToJSON outputs the JSON representation of the heap.
func (heap *Heap[T]) ToJSON() ([]byte, error) {
	return heap.list.ToJSON()
}

// FromJSON populates the heap from the input JSON representation.
func (heap *Heap[T]) FromJSON(data []byte) error {
	return heap.list.FromJSON(data)
}

// UnmarshalJSON @implements json.Unmarshaler
func (heap *Heap[T]) UnmarshalJSON(bytes []byte) error {
	return heap.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (heap *Heap[T]) MarshalJSON() ([]byte, error) {
	return heap.ToJSON()
}

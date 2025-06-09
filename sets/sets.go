// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sets provides an abstract Set interface.
//
// In computer science, a set is an abstract data type that can store certain values and no repeated values. It is a computer implementation of the mathematical concept of a finite set. Unlike most other collection types, rather than retrieving a specific element from a set, one typically tests a value for membership in a set.
//
// Reference: https://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package sets

import (
	"github.com/Arvin619/gods/containers"
)

// Set interface that all sets implement
type Set[T comparable] interface {
	Add(elements ...T)
	Remove(elements ...T)
	Contains(elements ...T) bool

	containers.Container[T]
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []T
	// String() string
}

// RichSet is a generic extension of Set with basic set operations.
// It assumes S is a pointer to a type implementing Set[T] and these operations.
type RichSet[T comparable, S any] interface {
	*S // pointer constraint

	Set[T]
	Intersection(another *S) *S
	Union(another *S) *S
	Difference(another *S) *S
}

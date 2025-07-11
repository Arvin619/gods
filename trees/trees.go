// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package trees provides an abstract Tree interface.
//
// In computer science, a tree is a widely used abstract data type (ADT) or data structure implementing this ADT that simulates a hierarchical tree structure, with a root value and subtrees of children with a parent node, represented as a set of linked nodes.
//
// Reference: https://en.wikipedia.org/wiki/Tree_%28data_structure%29
package trees

import "github.com/Arvin619/gods/containers"

// Tree interface that all trees implement
type Tree[V any] interface {
	containers.Container[V]
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []T
	// String() string
}

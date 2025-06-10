// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/Arvin619/gods/sets/treeset"
)

// User model (id and name)
type User struct {
	id   int
	name string
}

// compare method(sort by id)
func (u User) Compare(another User) int {
	switch {
	case u.id > another.id:
		return 1
	case u.id < another.id:
		return -1
	default:
		return 0
	}
}

// CustomComparatorExample to demonstrate basic usage of CustomComparator
func main() {
	set := treeset.NewWith(User.Compare)

	set.Add(User{2, "Second"})
	set.Add(User{3, "Third"})
	set.Add(User{1, "First"})
	set.Add(User{4, "Fourth"})

	fmt.Println(set) // {1 First}, {2 Second}, {3 Third}, {4 Fourth}
}

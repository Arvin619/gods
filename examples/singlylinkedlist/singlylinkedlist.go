// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"cmp"

	sll "github.com/Arvin619/gods/lists/singlylinkedlist"
)

// SinglyLinkedListExample to demonstrate basic usage of SinglyLinkedList
func main() {
	list := sll.New[string]()
	list.Add("a")                         // ["a"]
	list.Append("b")                      // ["a","b"] (same as Add())
	list.Prepend("c")                     // ["c","a","b"]
	list.Sort(cmp.Compare)                // ["a","b","c"]
	_, _ = list.Get(0)                    // "a",true
	_, _ = list.Get(100)                  // nil,false
	_ = list.Contains("a", "b", "c")      // true
	_ = list.Contains("a", "b", "c", "d") // false
	list.Remove(2)                        // ["a","b"]
	list.Remove(1)                        // ["a"]
	list.Remove(0)                        // []
	list.Remove(0)                        // [] (ignored)
	_ = list.Empty()                      // true
	_ = list.Size()                       // 0
	list.Add("a")                         // ["a"]
	list.Clear()                          // []
}

package main

import (
	"fmt"

	"github.com/Arvin619/gods/sets"
	"github.com/Arvin619/gods/sets/hashset"
	"github.com/Arvin619/gods/sets/treeset"
)

func combineSetsExcluding[T comparable, S any, RS sets.RichSet[T, S]](combine1, combine2, exclude RS) RS {
	rs := RS(combine1.Union(combine2))

	rs = rs.Difference(exclude)

	return rs
}

func main() {
	hs1 := hashset.New(1, 2, 3)
	hs2 := hashset.New(3, 4, 5)
	hs3 := hashset.New(2, 5, 6)

	resultHashSet := combineSetsExcluding(hs1, hs2, hs3)
	fmt.Println("combineSetsExcluding HashSet:", resultHashSet)
	// combineSetsExcluding HashSet: HashSet
	// 1, 3, 4 (unordered)

	ts1 := treeset.New("a", "b", "c")
	ts2 := treeset.New("c", "d", "e")
	ts3 := treeset.New("b", "e", "f")
	resultTreeSet := combineSetsExcluding(ts1, ts2, ts3)
	fmt.Println("combineSetsExcluding TreeSet:", resultTreeSet)
	// combineSetsExcluding TreeSet: TreeSet
	// a, c, d (ordered)
}

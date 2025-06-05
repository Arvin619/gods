package check

import "github.com/Arvin619/gods/sets"

// ImplementRichSet is used to statically assert that *S implements sets.RichSet[T, S].
func ImplementRichSet[T comparable, S any, RS sets.RichSet[T, S]](rs RS) RS { return rs }

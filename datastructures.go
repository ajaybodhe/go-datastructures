/*
Package datastructures exists solely to aid consumers of the go-datastructures
library when using dependency managers.  Depman, for instance, will work
correctly with any datastructure by simply importing this package instead of
each subpackage individually.

For more information about the datastructures package, see the README at

	http://github.com/ajaybodhe/go-datastructures

*/
package datastructures

import (
	_ "github.com/ajaybodhe/go-datastructures/augmentedtree"
	_ "github.com/ajaybodhe/go-datastructures/bitarray"
	_ "github.com/ajaybodhe/go-datastructures/btree/palm"
	_ "github.com/ajaybodhe/go-datastructures/btree/plus"
	_ "github.com/ajaybodhe/go-datastructures/fibheap"
	_ "github.com/ajaybodhe/go-datastructures/futures"
	_ "github.com/ajaybodhe/go-datastructures/hashmap/fastinteger"
	_ "github.com/ajaybodhe/go-datastructures/numerics/optimization"
	_ "github.com/ajaybodhe/go-datastructures/queue"
	_ "github.com/ajaybodhe/go-datastructures/rangetree"
	_ "github.com/ajaybodhe/go-datastructures/rangetree/skiplist"
	_ "github.com/ajaybodhe/go-datastructures/set"
	_ "github.com/ajaybodhe/go-datastructures/slice"
	_ "github.com/ajaybodhe/go-datastructures/slice/skip"
	_ "github.com/ajaybodhe/go-datastructures/sort"
	_ "github.com/ajaybodhe/go-datastructures/threadsafe/err"
	_ "github.com/ajaybodhe/go-datastructures/tree/avl"
	_ "github.com/ajaybodhe/go-datastructures/trie/xfast"
	_ "github.com/ajaybodhe/go-datastructures/trie/yfast"
)

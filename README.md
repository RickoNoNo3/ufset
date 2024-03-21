# Generic Disjoint Sets (Union-Find Sets)
[![Go](https://github.com/RickoNoNo3/ufset/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/RickoNoNo3/ufset/actions/workflows/go.yml)

Implementation of disjoint set in Go(>= 1.18 with generic features). For the algorithm, see https://en.wikipedia.org/wiki/Disjoint-set_data_structure


## Example

```go
package main

import (
	"fmt"
	"github.com/rickonono3/ufset"
)

func main() {
	sets := ufset.New[int]()
	sets.Union(0, 1)
	sets.Union(2, 1)
	sets.Union(0, 0)
	sets.Union(3, 4)
	sets.Union(5, 3)
	fmt.Println(sets.InSameSet(0, 2))
	fmt.Println(sets.InSameSet(0, 3))

	setsList := make([]int, 6)
	for i := 0; i < 6; i++ {
		setsList[i] = sets.Find(i)
	}
	fmt.Println(setsList)

	sets2 := ufset.New[string]()
	sets.Union("hello", "world")
	sets.Union("world", "peace")
	sets.Union("foo", "bar")
	fmt.Println(sets.Find("hello"), sets.Find("peace"), sets.Find("bar"))
}
```
`output:`
```text
true
false
[0, 0, 0, 3, 3, 3]
["hello", "hello", "foo"]
```

## Methods

### `ufset.New[K]() (*DisjointSets[K comparable])`

Instantiates a new sets struct.

### `(*DisjointSets[K]) Union(k1, k2 K)`

Unions two sets contain the provided keys, keys not exist will be created automatically as single-key sets.

### `(*DisjointSets[K]) Find(key K) K`

Returns the root key of a set contains the provided key.

### `(*DisjointSets[K]) InSameSet(k1, k2 K) bool`

Returns whether the set contains `k1` is exactly the set contains `k2`. Equals to `Find(k1) == Find(k2)`.

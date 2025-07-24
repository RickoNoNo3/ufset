package ufset

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func print(sets *DisjointSets[int]) {
	setsList := make([]int, 5)
	for i := 0; i < 5; i++ {
		setsList[i] = sets.Find(i)
	}
	fmt.Println(setsList)
}

func Test1(t *testing.T) {
	sets := New[int]()
	BasicTest(t, sets)
}

func Test2(t *testing.T) {
	sets := NewRigid[int]()
	BasicTest(t, sets)
	tree, err := json.Marshal(GetTree(sets))
	assert.NoError(t, err)
	fmt.Println(string(tree))
}

func BasicTest(t *testing.T, sets *DisjointSets[int]) {
	sets.add(0)
	sets.add(1)
	sets.add(2)
	sets.Union(0, 1)
	sets.Union(2, 1)
	sets.Union(0, 0)
	sets.Union(3, 4)
	sets.Union(5, 3)
	sets.Union(5, 5)
	print(sets)
	assert.Equal(t, sets.Find(0), sets.Find(1))
	assert.Equal(t, sets.Find(0), sets.Find(2))
	assert.Equal(t, sets.Find(3), sets.Find(4))
	assert.Equal(t, sets.Find(3), sets.Find(5))
	assert.True(t, sets.InSameSet(0, 1))
	assert.True(t, sets.InSameSet(3, 5))
	assert.NotEqual(t, sets.Find(0), sets.Find(3))
	assert.False(t, sets.InSameSet(0, 3))

	sets.Union(3, 0)
	print(sets)
	assert.Equal(t, sets.Find(0), sets.Find(3))
	assert.True(t, sets.InSameSet(0, 3))
}
